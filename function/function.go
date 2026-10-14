package main
import "fmt"


func hello() {
	fmt.Println("Hello, function!")
}

func add(a int, b int) int {
	 
	result := a + b
	fmt.Println("result:", result)
	return result
}

func add2(a int, b int)  {
	 
	result := a + b
	fmt.Println("result:", result)
}
func add3(a int, b int) int {
	 
	return  a + b
}


func main() {
	hello()
	add(4,5)
	add2(6,7)
	result := add3(8,9)
	fmt.Println("result:", result)
}


