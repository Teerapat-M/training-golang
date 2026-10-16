// ===============================================================
// อ่านไฟล์ข้อความ
// ---------------------------------------------------------------
// - os.ReadFile อ่านทั้งก้อนเป็น []byte แล้วแปลงเป็น string
// - bufio.Scanner อ่านทีละบรรทัด ไฟล์ใหญ่ใช้แบบนี้ดีกว่า
// รัน: go run ./file
// ===============================================================
package main

import (
	"bufio"
	"fmt"
	"os"
)

func readAll() {
	b, err := os.ReadFile(dataPath)
	if err != nil {
		fmt.Println("อ่านไม่สำเร็จ:", err)
		return
	}
	fmt.Println("--- ทั้งไฟล์ ---")
	fmt.Print(string(b))
}

func readLines() {
	f, err := os.Open(dataPath)
	if err != nil {
		fmt.Println("เปิดไม่สำเร็จ:", err)
		return
	}
	defer f.Close()

	fmt.Println("--- ทีละบรรทัด ---")
	s := bufio.NewScanner(f)
	n := 0
	for s.Scan() { // อ่านจนหมดไฟล์
		n++
		fmt.Println(n, s.Text())
	}
	if err := s.Err(); err != nil {
		fmt.Println("สแกนไม่สำเร็จ:", err)
	}
}

func main() {
	writeAll()
	writeAppend()
	readAll()
	readLines()
}
