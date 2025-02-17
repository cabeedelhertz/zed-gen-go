package scanner

import "errors"

func (s *Scanner) scanComment() (string, error) {
	lit := string(s.read())

	ch := s.read()
	switch ch {
	case '/':
		for ch != '\n' {
			lit += string(ch)
			if s.isEOF() {
				return lit, nil
			}
			ch = s.read()
		}
	case '*':
		for {
			if s.isEOF() {
				return lit, errors.New("unexpected EOF in multi-line comment")
			}
			lit += string(ch)
			ch = s.read()
			next := s.peek()
			if ch == '*' && next == '/' {
				lit += string(ch)
				lit += string(s.read()) // next
				break
			}
		}
	default:
		return "", errors.New("unexpected character in comment")
	}
	return lit, nil
}

func (s *Scanner) TryInlineComment() (string, error) {
	ch := s.read()
	for ch != '\n' {
		if s.isEOF() {
			return "", errors.New("unexpected EOF in inline comment")
		}
		if ch == '/' {
			s.unread(ch)
			return s.scanComment()
		}
		ch = s.read()
	}
	return "", nil
}
