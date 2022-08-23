package main

import (
	"fmt"
)

func isThisTheAgeOfMyMother(age int) string {
	if age > 57 {
		return "no, this is not the age of my mother"
	} else if age < 57 {
		return fmt.Sprintf("no, but it's true that she doesn't look a day over %d", age)
	}
	return "yes, this is the age of my mother"
}

func main() {
	fmt.Println("hello goci")
}
