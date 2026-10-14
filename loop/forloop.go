package main

import "fmt"

func main() {
	// Go มีแค่ for (ไม่มี while / do-while)
	// รูปแบบเต็ม: for init; condition; post { ... }

	// 1) for แบบ C: นับจาก 0 ถึง 4
	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}

	// 2) เหมือน while: ใส่แค่เงื่อนไข
	n := 3
	for n > 0 {
		fmt.Println("n =", n)
		n--
	}

	// 3) infinite loop: for { } ออกด้วย break
	count := 0
	for {
		count++
		if count == 3 {
			break
		}
	}
	fmt.Println("count =", count)

	// 4) continue: ข้ามรอบนี้ ไปรอบถัดไป
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue // ข้ามเลขคู่
		}
		fmt.Println("odd =", i)
	}

	// 5) range กับ slice: index + value
	names := []string{"A", "B", "C"}
	for i, name := range names {
		fmt.Println(i, name)
	}

	// ใช้ _ ถ้าไม่ต้องการ index
	for _, name := range names {
		fmt.Println(name)
	}

	// 6) range กับ map: key + value (ลำดับไม่รับประกัน)
	ages := map[string]int{"Ann": 20, "Bob": 25}
	for name, age := range ages {
		fmt.Println(name, age)
	}
}
