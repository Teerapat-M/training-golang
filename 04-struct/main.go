// ===============================================================
// บทที่ 4: Struct + Slice ของ Struct (ใกล้เคียงงานจริงที่สุด)
// ---------------------------------------------------------------
// - struct คือการรวมข้อมูลหลายชนิดเป็นก้อนเดียว คล้าย object / interface ใน TS
// - ตัวอย่างนี้ทำ CRUD รายการสินค้า แบบที่ backend ทำกับข้อมูลจริง
// - ชื่อขึ้นต้นตัวใหญ่ (Name) = public ใช้ข้ามแพ็กเกจได้, ตัวเล็ก = private
// รัน: go run ./04-struct
// ===============================================================
package main

import "fmt"

// Product คือ "แม่แบบ" ของสินค้าหนึ่งชิ้น
type Product struct {
	ID    int
	Name  string
	Price float64
}

// ---------- ADD ----------
// รับ slice เข้ามา แล้วคืน slice ใหม่กลับไป (เพราะ append อาจได้ slice ใหม่)
func addProduct(list []Product, p Product) []Product {
	return append(list, p)
}

// ---------- ACCESS ----------
// คืนค่า 2 ตัว: (สินค้า, เจอไหม) เป็นสไตล์ที่ Go ใช้บ่อยมาก
func findByID(list []Product, id int) (Product, bool) {
	for _, p := range list { // _ คือไม่สนใจ index
		if p.ID == id {
			return p, true
		}
	}
	return Product{}, false // ไม่เจอ คืน struct ว่าง + false
}

// ---------- UPDATE ----------
// ต้องแก้ผ่าน list[i] เพราะ p ใน range เป็นแค่ "สำเนา"
func updatePrice(list []Product, id int, newPrice float64) bool {
	for i := range list {
		if list[i].ID == id {
			list[i].Price = newPrice
			return true
		}
	}
	return false
}

// ---------- DELETE ----------
func deleteByID(list []Product, id int) []Product {
	for i, p := range list {
		if p.ID == id {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list // ไม่เจอ คืนของเดิม
}

func main() {
	var products []Product // slice ว่าง (nil) แต่ append ได้เลย

	// ADD: สร้าง struct โดยระบุชื่อ field (อ่านง่าย แนะนำ)
	products = addProduct(products, Product{ID: 1, Name: "Macbook", Price: 45900})
	products = addProduct(products, Product{ID: 2, Name: "iPhone", Price: 32900})
	products = addProduct(products, Product{ID: 3, Name: "iPad", Price: 19900})
	fmt.Printf("ทั้งหมด: %+v\n", products) // %+v พิมพ์ชื่อ field ด้วย

	// ACCESS: เข้าถึง field ด้วยจุด .
	fmt.Println("ชื่อชิ้นแรก:", products[0].Name)
	if p, ok := findByID(products, 2); ok {
		fmt.Printf("เจอ ID 2: %s ราคา %.2f บาท\n", p.Name, p.Price)
	}
	if _, ok := findByID(products, 99); !ok {
		fmt.Println("ไม่เจอ ID 99")
	}

	// UPDATE
	updatePrice(products, 1, 42900)
	products[2].Name = "iPad Air" // แก้ตรงๆ ผ่าน index ก็ได้
	fmt.Printf("หลังแก้ไข: %+v\n", products)

	// DELETE
	products = deleteByID(products, 2)
	fmt.Printf("หลังลบ ID 2: %+v\n", products)

	// ---------- แถม: pointer ----------
	// ถ้าอยากให้ฟังก์ชันแก้ struct ตัวจริง ให้ส่ง pointer (*Product)
	item := Product{ID: 9, Name: "AirPods", Price: 8990}
	applyDiscount(&item, 10) // & = ส่งที่อยู่ของตัวแปร
	fmt.Printf("หลังลด 10%%: %+v\n", item)
}

// *Product = รับ pointer เลยแก้ค่าตัวจริงได้ (ถ้ารับ Product เฉยๆ จะแก้แค่สำเนา)
func applyDiscount(p *Product, percent float64) {
	p.Price = p.Price * (100 - percent) / 100
}
