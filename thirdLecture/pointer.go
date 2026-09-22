package main

import "fmt"

func main() {
	age := 27 //regular variable
	var userAge *int
	userAge = &age
	//getting the address
	fmt.Println("Age Address: ", userAge)
	//getting the value of an address
	fmt.Println("Age Value: ", *userAge)

	rewriteAgeToAdultAge(userAge)
	fmt.Println("Adult age: ", age)
}

// direct mutation
func rewriteAgeToAdultAge(age *int) {
	//over wrote the value of the address
	*age = *age - 18

	//could be using return with the return type
	//return *age - 19
}
