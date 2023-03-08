package main

type Shape interface {
	Area() float64
	Volume() float64
}

func getArea(shape Shape) {
	shape.Area()
}

type Rectangle2D struct {
	x, y float64
}

func (s Rectangle2D) Area() float64 {
	return s.x * s.y
}

func main() {
	getArea(Rectangle2D{})
}

type Cube struct {
	x, y, z float64
}

func (s Cube) Area() float64 {
	return s.x * s.y
}

func (s Cube) Volume() float64 {
	return s.x * s.y * s.z
}

type HasArea interface {
	Area() float64
}

type HasVolume interface {
	Volume() float64
}
