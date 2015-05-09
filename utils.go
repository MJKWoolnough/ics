package ics

func unescape(p []byte) []byte {
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
