package ics

import (
	"bytes"
	"io"
)

type unfolder struct {
	r      io.Reader
	bufLen int
	buf    [2]byte
}

func (u *unfolder) Read(p []byte) (int, error) {
	var (
		m, n int
		err  error
	)
	q := p
	if u.bufLen > 0 {
		n = copy(p, u.buf[:u.bufLen])
		if n > 0 {
			u.buf[0] = u.buf[1]
			u.bufLen -= n
			p = p[n:]
			m += n
		}
	}
	for len(p) > 0 && err == nil {
		n, err = u.r.Read(p)
		m += n
		p = p[:n]
		var toRead int
		for {
			pos := bytes.Index(p, eol[:])
			if pos == -1 {
				break
			}
			copy(p[pos:], p[pos+3:])
			p = p[pos:]
			toRead += 3
		}
		p = p[len(p)-toRead:]
		m -= toRead
	}
	q = q[:m]
	for err == nil {
		if lq := len(q); lq > 0 && q[lq-1] == '\r' {
			u.bufLen, err = u.r.Read(u.buf[:])
			if u.bufLen < 2 {
				break
			}
			if u.buf[0] == '\n' && u.buf[1] == ' ' {
				m -= 1
				if err == nil {
					n, err = u.r.Read(q[lq-1:])
					m += n
					u.bufLen = 0
				}
			} else {
				break
			}
		} else if lq > 1 && q[lq-2] == '\r' && q[lq-1] == '\n' {
			u.bufLen, err = u.r.Read(u.buf[:1])
			if u.bufLen < 1 {
				break
			}
			if u.buf[0] == ' ' {
				m -= 2
				if err == nil {
					n, err = u.Read(q[lq-2:])
					m += n
					u.bufLen = 0
				}
			} else {
				break
			}
		} else {
			break
		}
	}
	return m, err
}
