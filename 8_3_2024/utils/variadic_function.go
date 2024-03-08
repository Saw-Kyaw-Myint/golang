package Testing

import "fmt"

func sum(nums ...int) {
	fmt.Println("nums", "")
    total :=0
	for _,num :=range nums{
		total +=num
	}

	fmt.Println(total)
}

func LearningVariadicFunction(){
	sum(1,3);
	nums := []int{1,3,3,4};
	sum(nums...);
}