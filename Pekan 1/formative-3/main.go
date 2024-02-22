package main

import (
	"fmt"
	"strconv"
)

func main() {
	// soal 1

	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var panjangPersegiPanjangInteger, _ = strconv.Atoi(panjangPersegiPanjang)
	var lebarPersegiPanjangInteger, _ = strconv.Atoi(lebarPersegiPanjang)
	var alasSegitigaInteger, _ = strconv.Atoi(alasSegitiga)
	var tinggiSegitigaInteger, _ = strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int = panjangPersegiPanjangInteger * lebarPersegiPanjangInteger
	var kelilingPersegiPanjang int = 2 * (panjangPersegiPanjangInteger + lebarPersegiPanjangInteger)
	var luasSegitiga int = (alasSegitigaInteger * tinggiSegitigaInteger) / 2

	fmt.Printf("luas persegi panjang: %d\n", luasPersegiPanjang)
	fmt.Printf("keliling persegi panjang: %d\n", kelilingPersegiPanjang)
	fmt.Printf("luas segitiga: %d\n", luasSegitiga)

	// soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	if nilaiJohn >= 80 {
		fmt.Println("indeks nilai John: A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("indeks  nilai John: B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("indeks  nilai John: C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("indeks  nilai John: D")
	} else if nilaiJohn < 50 {
		fmt.Println("indeks  nilai John: E")
	}

	if nilaiDoe >= 80 {
		fmt.Println("indeks nilai Doe: A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("indeks  nilai Doe: B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("indeks  nilai Doe: C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("indeks  nilai Doe: D")
	} else if nilaiDoe < 50 {
		fmt.Println("indeks  nilai Doe: E")
	}

	// soal 3
	var tanggal = 24
	var bulan = 1
	var tahun = 1995
	var bulanString string

	switch bulan {
	case 1:
		bulanString = "Januari"
	case 2:
		bulanString = "Februari"
	case 3:
		bulanString = "Maret"
	case 4:
		bulanString = "April"
	case 5:
		bulanString = "Mei"
	case 6:
		bulanString = "Juni"
	case 7:
		bulanString = "Juli"
	case 8:
		bulanString = "Agustus"
	case 9:
		bulanString = "September"
	case 10:
		bulanString = "Oktober"
	case 11:
		bulanString = "November"
	case 12:
		bulanString = "Desember"
	}
	tanggalLengkap := fmt.Sprintf(`%d %s %d`, tanggal, bulanString, tahun)
	fmt.Println(tanggalLengkap)

	// soal 4
	var generasi string
	if tahun >= 1944 && tahun < 1965 {
		generasi = "Baby boomer"
	} else if tahun >= 1965 && tahun < 1980 {
		generasi = "X"
	} else if tahun >= 1980 && tahun < 1995 {
		generasi = "Y (Millenials)"
	} else if tahun >= 1995 && tahun < 2015 {
		generasi = "Z"
	}
	fmt.Printf("Anda kelahiran tahun %d. sehingga anda adalah generasi %s\n", tahun, generasi)
}
