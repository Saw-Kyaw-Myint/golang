package Testing

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	widht, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.widht
}

func (r rect) perim() float64 {
	return 2*r.height + 2*r.widht
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func LearnInterface() {
	r := rect{height: 4, widht: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
