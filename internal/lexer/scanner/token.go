package scanner

type Token int

const (
	TILLEGAL Token = iota
	TEOF

	TIDENT

	TSTRLIT

	TCOMMENT

	TSEMICOLON   // ;
	TCOLON       // :
	TRIGHTARROW  // ->
	THASH        // #
	TPIPE        // |
	TPLUS        // +
	TEQUALS      // =
	TQUOTE       // " or '
	TLEFTPAREN   // (
	TRIGHTPAREN  // )
	TLEFTCURLY   // {
	TRIGHTCURLY  // }
	TLEFTSQUARE  // [
	TRIGHTSQUARE // ]
	TLESS        // <
	TGREATER     // >
	TCOMMA       // ,
	TDOT         // .
	TMINUS       // -
	TBOM         // Byte Order Mark

	TWILDCARD // *

	// Keywords
	TDEFINITION
	TRELATION
	TPERMISSION
)

func asMiscToken(ch rune) Token {
	m := map[rune]Token{
		'|':      TPIPE,
		';':      TSEMICOLON,
		':':      TCOLON,
		'=':      TEQUALS,
		'"':      TQUOTE,
		'\'':     TQUOTE,
		'(':      TLEFTPAREN,
		')':      TRIGHTPAREN,
		'{':      TLEFTCURLY,
		'}':      TRIGHTCURLY,
		'[':      TLEFTSQUARE,
		']':      TRIGHTSQUARE,
		'<':      TLESS,
		'>':      TGREATER,
		',':      TCOMMA,
		'.':      TDOT,
		'-':      TMINUS,
		'+':      TPLUS,
		'#':      THASH,
		'*':      TWILDCARD,
		'\uFEFF': TBOM,
	}
	if t, ok := m[ch]; ok {
		return t
	}
	return TILLEGAL
}

func asKeywordToken(st string) Token {
	m := map[string]Token{
		"definition": TDEFINITION,
		"relation":   TRELATION,
		"permission": TPERMISSION,
	}
	if t, ok := m[st]; ok {
		return t
	}
	return TILLEGAL
}
