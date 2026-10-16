// ===============================================================
// Goroutine
// ---------------------------------------------------------------
// - goroutine = ฟังก์ชันที่รันคู่ขนาน (lightweight thread ของ Go)
// - ใส่ go หน้าการเรียกฟังก์ชัน: go f()
// - main จบ = โปรแกรมจบ แม้ goroutine อื่นยังทำไม่เสร็จ เลยต้องรอ (WaitGroup / channel)
// รันทั้งโฟลเดอร์: go run ./channel
// ===============================================================
package main

import (
	"fmt"
	"sync"
	"time"
)

func say(msg string) {
	fmt.Println(msg)
}

func demoGoroutine() {
	// 1) เรียกธรรมดา: ทำเสร็จก่อน แล้วค่อยไปบรรทัดถัดไป
	say("sync: hello")

	// 2) go: สั่งแล้วไม่รอ ไปทำบรรทัดถัดไปเลย
	go say("async: จาก goroutine")

	// รอสั้นๆ กัน main จบก่อน goroutine พิมพ์ (เดโม่เท่านั้น)
	time.Sleep(50 * time.Millisecond)

	// 3) WaitGroup: รอหลาย goroutine ให้ครบก่อนเดินต่อ (ใช้จริงกว่า Sleep)
	var wg sync.WaitGroup
	names := []string{"A", "B", "C"}
	for _, name := range names {
		wg.Add(1) // บอกว่าจะรออีก 1 งาน
		go func(n string) {
			defer wg.Done() // งานเสร็จ ลดตัวนับ 1
			fmt.Println("งานของ", n)
		}(name) // ส่ง name เข้าไป กัน loop ใช้ค่าตัวสุดท้าย
	}
	wg.Wait() // บล็อกจนกว่า Done ครบตาม Add
	fmt.Println("goroutine ครบแล้ว")
}
