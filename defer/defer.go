// ===============================================================
// defer
// ---------------------------------------------------------------
// - defer เลื่อนการเรียกฟังก์ชันไปจนฟังก์ชันที่ครอบอยู่ return (ปกติหรือ panic)
// - ใช้ปิดของ / ปลดล็อก / wg.Done() ให้ชัวร์ แม้ฟังก์ชันออกกลางคัน
// - หลาย defer ในฟังก์ชันเดียวกัน ทำงานแบบกองซ้อน (LIFO): ตัวหลังถูกเรียกก่อน
// รัน: go run ./defer
// ===============================================================
package main

import "fmt"

func hello() {
	fmt.Println("hello")
}

func main() {
	// 1) เรียกทันที vs defer: พิมพ์ world ก่อน แล้วค่อย hello เมื่อ main ใกล้จบ
	defer hello()
	fmt.Println("world")

	demoOrder()
	demoArgs()
	fmt.Println("ผลคูณ:", multiply(3, 4))
}

func demoOrder() {
	fmt.Println("--- ลำดับ defer (LIFO) ---")
	defer fmt.Println("หนึ่ง") // ถูกเรียกทีหลังสุดในฟังก์ชันนี้
	defer fmt.Println("สอง")
	defer fmt.Println("สาม") // ถูกเรียกก่อน
	fmt.Println("ตัวฟังก์ชันยังทำอยู่")
}

func demoArgs() {
	fmt.Println("--- อาร์กิวเมนต์ถูกประเมินทันที ---")
	n := 1
	defer fmt.Println("defer ได้ n =", n) // ส่ง 1 เข้าไปแล้ว แม้เดี๋ยว n จะเป็น 2
	n = 2
	fmt.Println("ตอนนี้ n =", n)
}

// 2) defer กับฟังก์ชันที่มี return: ทำงานหลังคำนวณค่าคืน แต่ก่อนผู้เรียกได้รับค่า
func multiply(a, b int) int {
	defer fmt.Println("คูณเสร็จแล้ว")
	return a * b
}
