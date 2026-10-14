// ===============================================================
// บทที่ 2: Slice (อาร์เรย์ที่ยืดหดได้)
// ---------------------------------------------------------------
// - ใช้บ่อยที่สุดในงานจริง คล้าย array ใน JavaScript
// - เขียนว่า []string (ไม่ใส่ตัวเลขในวงเล็บ)
// - ข้างในคือ "หน้าต่าง" ที่มองไปยัง array ด้านหลัง
//   มี len (จำนวนที่ใช้อยู่) และ cap (ความจุที่จองไว้)
// รัน: go run ./02-slice
// ===============================================================
package main

import (
	"fmt"
	"slices" // แพ็กเกจมาตรฐาน (Go 1.21+) มีฟังก์ชันช่วยจัดการ slice
)

func main() {
	// ---------- สร้าง ----------
	products := []string{"Macbook", "iPhone"} // slice literal
	// อีกแบบ: make(ชนิด, len, cap) จองพื้นที่ล่วงหน้า
	buffer := make([]int, 0, 10)
	fmt.Println("products:", products, "| buffer len/cap:", len(buffer), cap(buffer))

	// ---------- ADD (เพิ่ม) ----------
	// append คืน slice ใหม่ ต้องรับค่ากลับใส่ตัวแปรเสมอ!
	// (เหมือน push ใน JS แต่ต้องเขียน products = ...)
	products = append(products, "iPad")
	products = append(products, "AirPods", "Watch") // เพิ่มหลายตัวพร้อมกัน
	fmt.Println("หลัง append:", products)

	// แทรกตรงกลาง: slices.Insert(slice, index, ค่า...)
	products = slices.Insert(products, 1, "iMac")
	fmt.Println("หลังแทรก index 1:", products)

	// ---------- ACCESS (เข้าถึง) ----------
	fmt.Println("ตัวแรก:", products[0])
	fmt.Println("ตัวสุดท้าย:", products[len(products)-1]) // Go ไม่มี index ติดลบแบบ [-1]
	fmt.Println("ช่วง [1:3]:", products[1:3])             // ตั้งแต่ index 1 ถึงก่อน 3 (เหมือน slice() ใน JS)

	// ค้นหาตำแหน่ง: คืน -1 ถ้าไม่เจอ (เหมือน indexOf)
	fmt.Println("iPad อยู่ index:", slices.Index(products, "iPad"))
	fmt.Println("มี Watch ไหม:", slices.Contains(products, "Watch"))

	for i, p := range products {
		fmt.Printf("  %d: %s\n", i, p)
	}

	// ---------- UPDATE (แก้ไข) ----------
	products[0] = "Macbook Pro"
	// แก้ทุกตัวในลูป: ต้องใช้ index เพราะ p ใน range เป็น "สำเนา"
	for i := range products {
		products[i] = "[Apple] " + products[i]
	}
	fmt.Println("หลังแก้ไข:", products)

	// ---------- DELETE (ลบ) ----------
	// วิธีที่ 1 (แนะนำ, Go 1.21+): slices.Delete(slice, ตั้งแต่, ถึงก่อน)
	products = slices.Delete(products, 1, 2) // ลบ index 1
	fmt.Println("หลังลบ index 1:", products)

	// วิธีที่ 2 (แบบดั้งเดิม ควรรู้ไว้เพราะเจอในโค้ดเก่า):
	// เอาส่วนหน้า + ส่วนหลัง มาต่อกัน โดยข้ามตัวที่จะลบ
	i := 0
	products = append(products[:i], products[i+1:]...) // ... คือกระจายสมาชิก (เหมือน spread ใน JS)
	fmt.Println("หลังลบ index 0:", products)

	// ลบทั้งหมด (ล้าง slice)
	products = products[:0]
	fmt.Println("หลังล้าง:", products, "len =", len(products))

	// ---------- ข้อควรระวัง ----------
	// slice ที่ตัดมาจากกัน "แชร์ข้อมูลเดียวกัน" แก้ตัวหนึ่ง อีกตัวเปลี่ยนด้วย
	a := []int{1, 2, 3}
	b := a[:2]
	b[0] = 99
	fmt.Println("a:", a, "b:", b) // a ก็เป็น 99 ด้วย
	// ถ้าอยากได้สำเนาแยกจริงๆ ใช้ slices.Clone(a)
}
