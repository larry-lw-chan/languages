package shapes

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.141592653589793 * c.Radius * c.Radius
}
