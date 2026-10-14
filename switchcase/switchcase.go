package main

import "fmt"

func main() {
	// switch ใน Go ไม่ต้องเขียน break ทุก case
	// เจอ case ที่ตรงแล้วจบทันที (ไม่ fall through เหมือน C)

	// 1) switch แบบพื้นฐาน: เทียบค่าตัวแปร
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("other day")
	}

	// 2) หลายค่าใน case เดียว คั่นด้วย comma
	letter := "a"
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("vowel")
	default:
		fmt.Println("consonant")
	}

	// 3) switch ไม่ใส่ expression = เหมือน if / else if
	score := 85
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	default:
		fmt.Println("F")
	}

	// 4) fallthrough: บังคับให้ทำ case ถัดไปด้วย (ใช้น้อย)
	n := 1
	switch n {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two") // รันต่อจาก case 1 เพราะ fallthrough
	}

	// 5) short statement: ประกาศตัวแปรใน switch ได้
	switch x := 10; x {
	case 10:
		fmt.Println("x is 10")
	}

	// 6) type switch: เช็คชนิดของ interface
	var v any = "hello"
	switch t := v.(type) {
	case string:
		fmt.Println("string:", t)
	case int:
		fmt.Println("int:", t)
	default:
		fmt.Println("unknown type")
	}
}
