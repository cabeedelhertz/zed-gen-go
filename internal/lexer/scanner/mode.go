package scanner

type Mode uint

const (
	ScanIdent Mode = 1 << iota
	ScanStrLit
	ScanKeyword
	ScanComment
	ScanLit = ScanStrLit
)
