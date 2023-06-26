package main

import (
	"fmt"
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
}
