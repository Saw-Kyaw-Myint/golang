package  Testing

import "fmt"

func LearningRange() {
	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("here range", i, num)
		}
	}

	mkv := map[string]string{"name": "hla hla", "age": "12", "gender": "female"}
	for k, v := range mkv {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "a" {
		fmt.Println(i, c)
	}
}
