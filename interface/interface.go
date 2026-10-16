// ===============================================================
// Interface ใน Go
// ---------------------------------------------------------------
// - interface คือ "สัญญา" ว่าต้องมี method อะไรบ้าง ไม่ระบุว่าทำยังไง
// - ไม่ต้องเขียน implements: มี method ครบตามสัญญา = ใช้เป็น type นั้นได้เลย
// - ใช้ตอนอยากเรียกพฤติกรรมเดียวกันกับหลายชนิด (เช่น Dog / Cat ทั้งคู่ Speak ได้)
// รัน: go run ./interface
// ===============================================================
package main

import "fmt"

// Speaker คือสัญญา: ใครมี Speak() string ก็เป็น Speaker ได้
type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says woof"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return c.Name + " says meow"
}

// รับเป็น Speaker ไม่สนว่า Dog หรือ Cat — สนแค่ว่า Speak ได้
func announce(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	dog := Dog{Name: "Lucky"}
	cat := Cat{Name: "Mimi"}

	// 1) ส่ง struct ที่ implement interface เข้าฟังก์ชัน
	announce(dog)
	announce(cat)

	// 2) เก็บหลายชนิดใน slice ของ interface เดียวกัน
	pets := []Speaker{dog, cat}
	for _, p := range pets {
		fmt.Println(p.Speak())
	}

	// 3) interface ว่าง (any) รับค่าทุกชนิด
	printAnything("hello")
	printAnything(42)

	// 4) type assertion: หยิบค่าจริงออกจาก interface
	var s Speaker = dog
	d, ok := s.(Dog) // ok = true ถ้า s เป็น Dog จริง
	if ok {
		fmt.Println("เป็นหมาชื่อ", d.Name)
	}

	// 5) type switch: แยกชนิดใน interface
	describe(dog)
	describe(cat)
	describe(10)
}

func printAnything(v any) { // any = interface{}
	fmt.Println("ค่า:", v)
}

func describe(v any) {
	switch x := v.(type) {
	case Dog:
		fmt.Println("Dog:", x.Name)
	case Cat:
		fmt.Println("Cat:", x.Name)
	default:
		fmt.Println("ชนิดอื่น:", x)
	}
}

