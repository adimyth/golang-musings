package basics

const pi = 3.14

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * pi * c.Radius
}

type Rectangle struct {
	Length  float32
	Breadth float32
}

func (r Rectangle) Area() float32 {
	return r.Length * r.Breadth
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Length + r.Breadth)
}
