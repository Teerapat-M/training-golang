// ===============================================================
// เขียนไฟล์ข้อความ
// ---------------------------------------------------------------
// - os.WriteFile เขียนทั้งก้อน (ทับของเดิม) สั้นสุด
// - os.Create / OpenFile ใช้เมื่ออยากเขียนทีละส่วน หรือต่อท้าย (append)
// - เปิดไฟล์แล้ว defer Close() เสมอ
// รันในโฟลเดอร์นี้: go run .
// ===============================================================
package main

import (
	"fmt"
	"os"
)

const dataPath = "data.txt"

func writeAll() {
	text := "บรรทัดที่ 1\nบรรทัดที่ 2\n"
	// 0644 = เจ้าของอ่าน/เขียน คนอื่นอ่านได้
	err := os.WriteFile(dataPath, []byte(text), 0644)
	if err != nil {
		fmt.Println("เขียนไม่สำเร็จ:", err)
		return
	}
	fmt.Println("WriteFile แล้ว:", dataPath)
}

func writeAppend() {
	// O_APPEND = ต่อท้ายของเดิม, O_CREATE = ไม่มีไฟล์ให้สร้าง
	f, err := os.OpenFile(dataPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("เปิดไม่สำเร็จ:", err)
		return
	}
	defer f.Close() // ปิดไฟล์เมื่อฟังก์ชันจบ

	_, err = f.WriteString("บรรทัดต่อท้าย\n")
	if err != nil {
		fmt.Println("ต่อท้ายไม่สำเร็จ:", err)
		return
	}
	fmt.Println("ต่อท้ายแล้ว")
}
