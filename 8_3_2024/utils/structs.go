package Testing

import "fmt"

type Person struct {
	name string
	age  int
}

func LearningStruct(){
	var person1 Person
	var person2 Person
    
    person1.name="kyaw kyaw"
	person1.age=13

	person2.name="maw maw"
	person2.age=13

	fmt.Println("struct",person1.name)
	fmt.Println("struct",person1.age)
	fmt.Println("struct",person2.name)
	fmt.Println("struct",person2.age)
}