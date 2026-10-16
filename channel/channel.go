// ===============================================================
// Channel
// ---------------------------------------------------------------
// - channel = ท่อส่งค่าข้าม goroutine (type-safe)
// - ch <- v ส่งค่า, v := <-ch รับค่า — ส่ง/รับจะบล็อกจนกว่าอีกฝั่งพร้อม (ถ้าไม่ใส่ buffer)
// - ใช้แทนการแชร์ตัวแปรกันตรงๆ: "อย่าสื่อสารด้วยการแชร์หน่วยความจำ แต่แชร์หน่วยความจำด้วยการสื่อสาร"
// รัน: go run ./channel
// ===============================================================
package main

import "fmt"

func demoChannel() {
	// 1) unbuffered: ส่งได้เมื่อมีคนรอรับ (ต้องมี goroutine อีกตัว)
	ch := make(chan string)
	go func() {
		ch <- "ping" // ส่ง — บล็อกจนกว่า main จะรับ
	}()
	msg := <-ch // รับ
	fmt.Println("รับจาก channel:", msg)

	// 2) buffered: ส่งได้โดยยังไม่มีคนรับ จนกว่าช่องเต็ม
	buf := make(chan int, 2)
	buf <- 1
	buf <- 2
	fmt.Println(<-buf, <-buf)

	// 3) ส่งผลลัพธ์งานกลับมาที่ main โดยไม่ใช้ WaitGroup
	result := make(chan int)
	go func() {
		sum := 0
		for i := 1; i <= 5; i++ {
			sum += i
		}
		result <- sum
	}()
	fmt.Println("ผลรวม:", <-result)

	// 4) close + range: ผู้ส่งปิดท่อ เมื่อไม่มีค่าใหม่แล้ว
	nums := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			nums <- i
		}
		close(nums) // ลืม close = range จะวนค้าง
	}()
	for n := range nums { // รับจน channel ถูก close
		fmt.Println("range:", n)
	}

	// 5) select: รอหลาย channel พร้อมกัน ทำตัวที่พร้อมก่อน
	a := make(chan string)
	b := make(chan string)
	go func() { a <- "จาก A" }()
	go func() { b <- "จาก B" }()
	for i := 0; i < 2; i++ {
		select {
		case s := <-a:
			fmt.Println(s)
		case s := <-b:
			fmt.Println(s)
		}
	}
}

func main() {
	fmt.Println("--- goroutine ---")
	demoGoroutine()
	fmt.Println("--- channel ---")
	demoChannel()
}
