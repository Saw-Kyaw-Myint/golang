package Testing

import "fmt"

func parent() (int, int) {
	return 3, 7
}
func multiple() (int, int) {
	return parent()
}
func LearningMultiple() {
	a, b := multiple()
	fmt.Println("multiple return", a, b)
}
