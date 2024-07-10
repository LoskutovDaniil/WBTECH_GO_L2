package pkg

type Visitor interface {
	VisitForSquare(*Square)
	VisitForCircle(*Circle)
	VisitForRectangle(*Rectangle)
}

// Определение структуры и метода для Circle
type Circle struct {
	Radius int
}

func (c *Circle) Accept(v Visitor) {
	v.VisitForCircle(c)
}

func (c *Circle) GetType() string {
	return "Circle"
}

// Определение структуры и метода для Square
type Square struct {
	Side int
}

func (s *Square) Accept(v Visitor) {
	v.VisitForSquare(s)
}

func (s *Square) GetType() string {
	return "Square"
}

// Определение структуры и метода для Rectangle
type Rectangle struct {
	Length int
	Width  int
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitForRectangle(r)
}

func (r *Rectangle) GetType()
