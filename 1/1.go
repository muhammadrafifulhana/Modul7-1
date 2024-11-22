package main

import (
	"fmt"
	"math"
)

// Struktur data untuk titik
type Titik struct {
	x, y int
}

// Struktur data untuk lingkaran
type Lingkaran struct {
	titikPusat Titik
	radius     int
}

// Fungsi untuk menghitung jarak antara dua titik
func jarak(p, q Titik) float64 {
	return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2))
}

// Fungsi untuk mengecek apakah sebuah titik berada di dalam lingkaran
func diDalamLingkaran(lingkaran Lingkaran, titik Titik) bool {
	return jarak(lingkaran.titikPusat, titik) <= float64(lingkaran.radius)
}

func main() {
	var lingkaran1, lingkaran2 Lingkaran
	var titik Titik

	// Input untuk lingkaran pertama
	fmt.Print("Masukkan titik pusat dan radius lingkaran 1 (cx cy r): ")
	fmt.Scan(&lingkaran1.titikPusat.x, &lingkaran1.titikPusat.y, &lingkaran1.radius)

	// Input untuk lingkaran kedua
	fmt.Print("Masukkan titik pusat dan radius lingkaran 2 (cx cy r): ")
	fmt.Scan(&lingkaran2.titikPusat.x, &lingkaran2.titikPusat.y, &lingkaran2.radius)

	// Input untuk titik sembarang
	fmt.Print("Masukkan koordinat titik (x y): ")
	fmt.Scan(&titik.x, &titik.y)

	// Mengecek posisi titik terhadap lingkaran 1 dan lingkaran 2
	dalamLingkaran1 := diDalamLingkaran(lingkaran1, titik)
	dalamLingkaran2 := diDalamLingkaran(lingkaran2, titik)

	// Menentukan output berdasarkan posisi titik
	if dalamLingkaran1 && dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
