package meta

type Meta struct {
	Pos     Position
	LastPos Position
}

type Position struct {
	Filename string
	Offset   int
	Line     int
	Column   int
}
