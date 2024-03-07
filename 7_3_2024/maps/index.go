package Map

import "fmt"

func LearningMap() {
	x := map[string]string{"name": "saw kyaw myint", "age": "age"}
	val, key := x["name"]
	_, checkKey := x["age"]
	fmt.Println("chaeck key", checkKey)
	fmt.Println(val, key)

	xx := make(map[int]int)
	xx[0] = 1
	xx[1] = 2
	xx[2] = 3
	fmt.Println("xx", xx)
	delete(xx, 2)
	fmt.Println("xx", xx)
	clear(xx)
	_, prs := xx[1]
	fmt.Println("prs:", prs)
	fmt.Println("xx", xx)
}
