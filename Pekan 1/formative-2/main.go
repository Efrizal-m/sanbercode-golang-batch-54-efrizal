package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// soal 1
	var word1 string = "Bootcamp"
	var word2 string = "Digital"
	var word3 string = "Skill"
	var word4 string = "Sanbercode"
	var word5 string = "Golang"
	fmt.Println(word1 + " " + word2 + " " + word3 + " " + word4 + " " + word5)

	// soal 2
	halo := "Halo Dunia"
	var newHaloText = strings.Replace(halo, "Dunia", "Golang", 1)
	fmt.Println(newHaloText)

	// soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"
	fmt.Println(kataPertama + " " + strings.Replace(kataKedua, string(kataKedua[0]), strings.ToUpper(string(kataKedua[0])), 1) + " " + strings.Replace(kataKetiga, string(kataKetiga[len(kataKetiga)-1]), strings.ToUpper(string(kataKetiga[len(kataKetiga)-1])), 1) + " " + strings.ToUpper(kataKeempat))

	// soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	var angkaPertamaConverted, _ = strconv.Atoi(angkaPertama)
	var angkaKeduaConverted, _ = strconv.Atoi(angkaKedua)
	var angkaKetigaConverted, _ = strconv.Atoi(angkaKetiga)
	var angkaKeempatConverted, _ = strconv.Atoi(angkaKeempat)
	var jumlahTotal = angkaPertamaConverted + angkaKeduaConverted + angkaKetigaConverted + angkaKeempatConverted
	fmt.Println(jumlahTotal)

	// soal 5
	kalimat := "halo halo bandung"
	angka := 2021
	fmt.Printf("\"%s\" - %d \n", strings.ReplaceAll(kalimat, "halo", "HI"), angka)
}
