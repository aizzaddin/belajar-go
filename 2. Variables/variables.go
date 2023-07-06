package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// STRING
	var nama string
	nama = "wildan"
	fmt.Println("variable nama : " + nama)

	namaToUpper := strings.ToUpper(nama) // WILDAN
	fmt.Println("variable nama setelah di uppercase : " + namaToUpper)

	nama1 := nama[0:2]
	fmt.Println("slice karakter index ke-0 -> ke-2 pada variable nama : " + nama1) // wi

	namaBelakang := "aizzaddin"
	fmt.Println(strings.ToUpper(nama) + " " + strings.ToUpper(namaBelakang)) // WILDAN AIZZADDIN

	// ARRAY
	var namaArray = [5]string{"Ahmad", "Dila", "Wildan", "Caca", "Doni"}
	fmt.Println(namaArray[1]) // Dila

	namaArray[1] = "Indah"
	fmt.Println(namaArray[1]) // Indah

	// SLICES
	var namaSlices = []string{"Ahmad", "Dila", "Wildan", "Caca", "Doni"}
	fmt.Println(namaSlices) // [Ahmad Dila Wildan Caca Doni]

	namaSlicesNew := append(namaSlices, "Galang", "Dani")
	fmt.Println(namaSlicesNew) // [Ahmad Dila Wildan Caca Doni Galang Dani]

	namaSlices1 := namaSlicesNew[:1]
	fmt.Println(namaSlices1) // [Ahmad]

	namaSlices1 = namaSlicesNew[2:3]
	fmt.Println(namaSlices1) // [Wildan]

	// MAPS
	profil := make(map[string]string)
	profil["nama"] = "Ahmad"
	profil["umur"] = "20"

	fmt.Println(profil["nama"]) // Ahmad
	fmt.Println(profil["umur"]) // 20

	delete(profil, "umur")
	fmt.Println(profil["umur"]) // " "

	// For Loops
	for i := 0; i <= 10; i++ {
		fmt.Println("Index ke-" + strconv.Itoa(i))
	}
}
