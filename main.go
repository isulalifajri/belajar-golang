package main

import "fmt"

// constant
const AppName string = "Belajar Golang Dasar"

// type declaration (membuat alias tipe data baru)
type Age int
type Status bool

// === Struct ===
type Mahasiswa struct {
	Nama   string
	Umur   int
	Prodi  string
	IPK    float64
	Active bool
}

// Method untuk struct Mahasiswa
func (m Mahasiswa) SapaMahasiswa() {
	fmt.Printf("Halo, saya %s dari prodi %s dengan IPK %.2f!\n", m.Nama, m.Prodi, m.IPK)
}

// === Function tanpa parameter ===
func sapa() {
	fmt.Println("Halo dari function sapa()!")
}

// === Function dengan parameter ===
func hitungLuasPersegiPanjang(panjang int, lebar int) {
	luas := panjang * lebar
	fmt.Printf("Luas persegi panjang (%d x %d) = %d\n", panjang, lebar, luas)
}

// === Function dengan return value ===
func tambah(a int, b int) int {
	return a + b
}

// === Function dengan defer, panic, dan recover ===
func prosesData() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi error:", r)
			fmt.Println("Program tetap lanjut berkat recover()")
		}
		fmt.Println("Selesai menjalankan defer di prosesData()")
	}()

	fmt.Println("Mulai proses data...")

	var pembagi int = 0
	if pembagi == 0 {
		panic("Tidak bisa membagi dengan nol!")
	}

	fmt.Println("Hasil:", 10/pembagi)
}

func main() {
	// variable string
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
	var y float64 = float64(x)
	var z int = int(y)
	var char byte = 'A'
	var str string = string(char)

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
	fmt.Printf("%d %% %d = %d\n", a, b, modulo)

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
			continue
		}
		if i == 8 {
			fmt.Println("Berhenti di angka", i, "(break)")
			break
		}
		fmt.Println("Angka:", i)
	}

	// === Function ===
	fmt.Println("\n=== Function ===")
	sapa()
	hitungLuasPersegiPanjang(5, 10)
	hitungLuasPersegiPanjang(7, 3)
	hasil := tambah(12, 8)
	fmt.Println("Hasil penjumlahan (12 + 8) =", hasil)

	// === Struct ===
	fmt.Println("\n=== Struct ===")
	m1 := Mahasiswa{
		Nama:   "Bunga Melati",
		Umur:   20,
		Prodi:  "Teknik Informatika",
		IPK:    3.85,
		Active: true,
	}
	fmt.Printf("Data Mahasiswa: %+v\n", m1)
	m1.SapaMahasiswa()

	// === Defer, Panic, Recover ===
	fmt.Println("\n=== Defer, Panic, dan Recover ===")
	prosesData()

	fmt.Println("Program tetap berjalan setelah recover()")
}
