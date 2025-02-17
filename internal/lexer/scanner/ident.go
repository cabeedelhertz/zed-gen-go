package scanner

func (s *Scanner) scanIdent() string {
	ident := string(s.read())

	for {
		next := s.peek()
		if isLetter(next) || isDecimalDigit(next) || next == '_' {
			ident += string(s.read())
		} else {
			return ident
		}
	}
}
