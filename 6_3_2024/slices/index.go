package main

import (
	"fmt"
	"slices"
)

func main() {
	mySlices := []int{}
	fmt.Println(len(mySlices))
	fmt.Println("saw kyaw myint", mySlices)

	mySlices2 := []string{"hla", "mya", "daw"}
	fmt.Println(len(mySlices2))
	mySlices2[2] = "kyaw"
	fmt.Println(mySlices2)
	fmt.Println(cap(mySlices2))

	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	testingSlices := arr1[2:4]
	fmt.Printf("myslice = %v\n", testingSlices)
	fmt.Printf("length = %d\n", len(testingSlices))
	fmt.Printf("capacity = %d\n", cap(testingSlices))

	s := make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s = append(s, "d")
	s = append(s, "e")

	fmt.Println("make slice", s)
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy value", c)

	cut := s[2:5]
	fmt.Println("slice array s", cut)
	fmt.Println("original  value", s)

	cut2 := s[:3]
	fmt.Println(cut2)
	cut3 := s[2:]
	fmt.Println(cut3)

	sOne := []string{"one", "two", "three"}
	sTwo := []string{"one", "two", "three"}
	if slices.Equal(sOne, sTwo) {
		fmt.Println("equal slices")
	} else {
		fmt.Println("not equal slices")
	}
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
