package ics

import (
	"bytes"
	"io"
)

type unfolder struct {
	r      io.Reader
	buf    [2]byte
	bufLen int
}

func (u *unfolder) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	var n int
	if u.bufLen > 0 {
		n = copy(p, u.buf[:u.bufLen])
		if n == 1 && u.bufLen == 2 {
			u.buf[0] = u.buf[1]
			u.bufLen = 1
			return 1, nil
		}
		p = p[n:]
		u.bufLen = 0
	}
	m, err := u.r.Read(p)
	n += m
	p = p[:m]
	for {
		pos := bytes.IndexByte(p, '\r')
		if pos < 0 {
			return n, err
		}
		if len(p) > pos+1 {
			if p[pos+1] == '\n' {
				if len(p) > pos+2 {
					if p[pos+2] == ' ' {
						copy(p[pos:], p[pos+3:])
						p = p[pos:]
						m, _ := io.ReadFull(u.r, p[len(p)-3:])
						p = p[:len(p)-3+m]
						n = n - 3 + m
					} else {
						p = p[pos+2:]
					}
				} else if err != nil {
					return n, err
				} else {
					io.ReadFull(u.r, u.buf[:1])
					if u.buf[0] == ' ' {
						p = p[pos:]
						m, e := io.ReadFull(u.r, p[len(p)-2:])
						n = n - 2 + m
						if e != nil {
							return n, err
						}
					} else {
						u.bufLen = 1
						return n, err
					}
				}
			} else {
				p = p[pos+1:]
			}
		} else {
			_, err := io.ReadFull(u.r, u.buf[:2])
			if err != nil {
				return n, err
			}
			if u.buf[0] == '\n' && u.buf[1] == ' ' {
				_, e := io.ReadFull(u.r, u.buf[:1])
				if e != nil {
					return n, err
				}
				p[pos] = u.buf[0]
			} else {
				u.bufLen = 2
			}
			return n, err
		}
	}
}
