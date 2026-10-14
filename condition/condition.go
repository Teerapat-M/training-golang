package main

import "fmt"

var score int

func main() {

	fmt.Println("Enter your score:")
	fmt.Scanf("%d", &score)
	fmt.Println("Your score is: ", score)
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else {
		fmt.Println("F")
	}
}
