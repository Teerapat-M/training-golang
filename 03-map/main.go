// ===============================================================
// บทที่ 3: Map (เก็บข้อมูลแบบ key -> value)
// ---------------------------------------------------------------
// - คล้าย object / Map ใน JavaScript
// - ชนิดเขียนว่า map[ชนิดของkey]ชนิดของvalue
// - ลำดับตอนวนลูป "ไม่แน่นอน" (สุ่มทุกครั้ง) อย่าพึ่งลำดับ
// รัน: go run ./03-map
// ===============================================================
package main

import "fmt"

func main() {
	// ---------- สร้าง ----------
	// แบบ literal
	stock := map[string]int{
		"Macbook": 5,
		"iPhone":  10, // ตัวสุดท้ายต้องมี , ด้วย (Go บังคับ)
	}
	// แบบ make (map ว่าง)
	prices := make(map[string]float64)

	// ระวัง: var m map[string]int แบบนี้ได้ map เป็น nil
	// อ่านได้ แต่ถ้าเขียน m["a"] = 1 จะ panic -> ต้อง make ก่อน

	// ---------- ADD (เพิ่ม) ----------
	stock["iPad"] = 3
	prices["Macbook"] = 45900
	prices["iPhone"] = 32900
	fmt.Println("stock:", stock)
	fmt.Println("prices:", prices)

	// ---------- ACCESS (เข้าถึง) ----------
	fmt.Println("Macbook เหลือ:", stock["Macbook"])

	// key ที่ไม่มี จะได้ zero value (0) ไม่ error -> แยกไม่ออกว่า "มีค่า 0" หรือ "ไม่มี"
	fmt.Println("AirPods:", stock["AirPods"]) // 0

	// วิธีที่ถูก: "comma ok" idiom รับค่าที่สองเป็น bool บอกว่ามี key ไหม
	if qty, ok := stock["AirPods"]; ok {
		fmt.Println("มี AirPods:", qty)
	} else {
		fmt.Println("ไม่มี AirPods ในสต็อก")
	}

	fmt.Println("จำนวน key:", len(stock))
	for name, qty := range stock {
		fmt.Printf("  %s = %d\n", name, qty)
	}

	// ---------- UPDATE (แก้ไข) ----------
	// ใช้ syntax เดียวกับ add: ถ้ามี key อยู่แล้วจะเขียนทับ
	stock["Macbook"] = 4
	stock["iPhone"] += 5 // บวกเพิ่มจากค่าเดิม
	stock["iPad"]--      // ลดลง 1
	fmt.Println("หลังแก้ไข:", stock)

	// ---------- DELETE (ลบ) ----------
	// delete(map, key) ถ้าไม่มี key ก็ไม่ error
	delete(stock, "iPad")
	delete(stock, "ไม่มีจริง")
	fmt.Println("หลังลบ iPad:", stock)

	// ล้างทั้งหมด (Go 1.21+)
	clear(prices)
	fmt.Println("หลัง clear prices:", prices, "len =", len(prices))
}
