package repeated

type Inner struct {
	Text string
}

type Outer struct {
	ID     int
	Inners []Inner
}
