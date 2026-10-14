package main

import "fmt"

func zerovalue(ivalue int){
	ivalue = 0
}

func zeropointer(ippointer *int){
	*ippointer = 0
}
func main() {
	i := 1
	fmt.Println("i value:", i)

	zerovalue(i)
	fmt.Println("i from zerovalue:", i)

	zeropointer(&i)
	fmt.Println("i from zeropointer:", i)
	fmt.Println("i address from zeropointer:", &i)

}