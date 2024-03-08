package Testing

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func ClosuresFunction() {
	nextSeq := intSeq()
	fmt.Println(nextSeq())
	fmt.Println(nextSeq())
	fmt.Println(nextSeq())
	nextSeqs := intSeq()
	fmt.Println(nextSeqs())
}
