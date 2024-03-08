package Testing

import (
	"fmt"
)

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func LearningMethods() {
	r := rect{width: 10, height: 20}
	fmt.Println("area", r.area())
	fmt.Println(r.perim())
	rp := r
	fmt.Println("area", rp.area())
	fmt.Println(rp.perim())
}
