package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Go", "Python", "Java", "C++", "JavaScript"}
	fmt.Println(courseName)
	courseName = append(courseName, "Ruby", "PHP", "C#")
	fmt.Println(courseName)

	courseWeb := courseName[0:2]
	fmt.Println(courseWeb)

	courseWeb = courseName[:4]
	fmt.Println(courseWeb)

	courseWeb = courseName[4:4]
	fmt.Println(courseWeb)
}
