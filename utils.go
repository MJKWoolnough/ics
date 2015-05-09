package ics

func escape(s string) []byte {
	p := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '\\':
			p = append(p, '\\', '\\')
		case ';':
			p = append(p, '\\', ';')
		case ',':
			p = append(p, '\\', ',')
		case '\n':
			p = append(p, '\\', 'n')
		default:
			p = append(p, s[i])
		}
	}
	return p
}

func unescape(p string) []byte {
	u := p[:0]
	for i := 0; i < len(p); i++ {
		if p[i] == '\\' && i+1 < len(p) {
			i++
			switch p[i] {
			case '\\':
				u = append(u, '\\')
			case ';':
				u = append(u, ';')
			case ',':
				u = append(u, ',')
			case 'N', 'n':
				u = append(u, '\n')
			default:
				u = append(u, p[i])
			}
		} else {
			u = append(u, p[i])
		}
	}
	return toRet
}

func escape6868(s string) []byte {
	p := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '\n':
			p = append(p, '^', 'n')
		case '^':
			p = append(p, '^', '^')
		case '"':
			p = append(p, '^', '\'')
		default:
			p = append(p, s[i])
		}
	}
	return p
}

func unescape6868(p []byte) []byte {
	u := p[:0]
	for i := 0; i < len(p); i++ {
		if p[i] == '^' && i+1 < len(p) {
			i++
			switch p[i] {
			case 'n':
				u = append(u, '\n') //crlf on windows?
			case '^':
				u = append(u, '^')
			case '\'':
				u = append(u, '"')
			default:
				u = append(u, '^', p[i])
			}
		} else {
			u = append(u, p[i])
		}
	}
	return u
}

func textSplit(s string, delim byte) []string {
	toRet := make([]string, 0, 1)
	lastPos := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '\\':
			i++
		case delim:
			toRet = append(toRet, unescape(s[lastPos:i]))
			lastPos = i + 2
		}
	}
	if lastPos <= len(s) {
		toRet = append(toRet, unescape(s[lastPos:len(s)]))
	}
	return toRet
}
