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

	// === Array ===
	var angka [3]int = [3]int{1, 2, 3}
	fmt.Println("\n=== Array ===")
	fmt.Println("Isi array:", angka)
	fmt.Println("Panjang array:", len(angka))

	// === Slice ===
	slice := []string{"apel", "jeruk", "mangga"}
	fmt.Println("\n=== Slice ===")
	fmt.Println("Isi slice:", slice)
	slice = append(slice, "pisang")
	fmt.Println("Setelah ditambah:", slice)

	// === Map ===
	mahasiswa := map[string]int{
		"Melati": 20,
		"Budi":   22,
	}
	fmt.Println("\n=== Map ===")
	fmt.Println("Isi map:", mahasiswa)
	fmt.Println("Umur Melati:", mahasiswa["Melati"])

	// === If Expression ===
	fmt.Println("\n=== If Expression ===")
	if umur > 18 {
		fmt.Println("Status: Dewasa")
	} else {
		fmt.Println("Status: Anak-anak")
	}

	// === Switch Expression ===
	fmt.Println("\n=== Switch Expression ===")
	hari := "Senin"
	switch hari {
	case "Senin":
		fmt.Println("Hari ini Senin, semangat kerja!")
	case "Sabtu", "Minggu":
		fmt.Println("Weekend, saatnya libur!")
	default:
		fmt.Println("Hari biasa.")
	}

	// === Break & Continue ===
	fmt.Println("\n=== Break & Continue ===")
	for i := 1; i <= 10; i++ {
		if i == 3 {
			fmt.Println("Lewati angka", i, "(continue)")
			continue // loncat ke iterasi berikutnya
		}
		if i == 8 {
			fmt.Println("Berhenti di angka", i, "(break)")
			break // hentikan loop sepenuhnya
		}
		fmt.Println("Angka:", i)
	}
}
