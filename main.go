package main

import "fmt"

// constant
const AppName string = "Belajar Golang Dasar"

// type declaration (membuat alias tipe data baru)
type Age int
type Status bool

func main() {
	// variable string
	// var nama string = "Hening"
	nama := "Hening"

	// variable boolean pakai type declaration
	var aktif Status = true

	// variable integer pakai type declaration
	var umur Age = 25

	// operasi matematika dasar
	a := 10
	b := 3
	jumlah := a + b
	kurang := a - b
	kali := a * b
	bagi := a / b
	modulo := a % b

	// === Konversi tipe data ===
	var x int = 10
	var y float64 = float64(x)   // int ke float64
	var z int = int(y)           // float64 ke int

	var char byte = 'A'
	var str string = string(char) // byte ke string

	// cetak semua
	fmt.Println("=== Cetak Variable & Constant ===")
	fmt.Println("Constant :", AppName)
	fmt.Println("Nama     :", nama)
	fmt.Println("Umur     :", umur)
	fmt.Println("Aktif?   :", aktif)

	fmt.Println("\n=== Operasi Matematika ===")
	fmt.Printf("%d + %d = %d\n", a, b, jumlah)
	fmt.Printf("%d - %d = %d\n", a, b, kurang)
	fmt.Printf("%d * %d = %d\n", a, b, kali)
	fmt.Printf("%d / %d = %d\n", a, b, bagi)
	fmt.Printf("%d %% %d = %d\n", a, b, modulo) // %% untuk cetak %

	fmt.Println("\n=== Konversi Tipe Data ===")
	fmt.Println("int ke float64 :", y)
	fmt.Println("float64 ke int :", z)
	fmt.Println("byte ke string :", str)
}