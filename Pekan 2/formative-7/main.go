package main

import (
	"fmt"
)

// untuk soal 2
type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (s segitiga) hitungLuasSegitiga() {
	fmt.Printf("Luas segitiga: %d\n", (s.alas*s.tinggi)/2)
}

func (p persegi) hitungLuasPersegi() {
	fmt.Printf("Luas persegi: %d\n", p.sisi*p.sisi)
}

func (pp persegiPanjang) hitungLuasPersegiPanjang() {
	fmt.Printf("Luas persegi panjang: %d\n", pp.panjang*pp.lebar)
}

//--------------------------------------------------

// untuk soal 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p *phone) addColor(color string) (result string) {
	p.colors = append(p.colors, color)
	for i, b := range p.colors {
		if i == len(p.colors)-1 {
			result += fmt.Sprintf(`"%s"`, b)
		} else {
			result += fmt.Sprintf(`"%s", `, b)
		}
	}
	return result
}

func main() {
	//soal 1
	type buah struct {
		nama, warna string
		adaBijinya  bool
		harga       int
	}
	var buah1 = buah{nama: "nanas", warna: "kuning", adaBijinya: false, harga: 9000}
	var buah2 = buah{nama: "jeruk", warna: "orange", adaBijinya: true, harga: 8000}
	var buah3 = buah{nama: "semangka", warna: "hijau & merah", adaBijinya: true, harga: 10000}
	var buah4 = buah{nama: "pisang", warna: "kuning", adaBijinya: false, harga: 5000}
	fmt.Println(buah1)
	fmt.Println(buah2)
	fmt.Println(buah3)
	fmt.Println(buah4)

	//soal 2
	var segitiga = segitiga{3, 7}
	segitiga.hitungLuasSegitiga()

	var persegi = persegi{7}
	persegi.hitungLuasPersegi()

	var persegiPanjang = persegiPanjang{2, 5}
	persegiPanjang.hitungLuasPersegiPanjang()

	//soal 3

	var phone = phone{name: "M23 5G", brand: "Samsung", year: 2022, colors: []string{"blue"}}
	resultColor := phone.addColor("black")
	fmt.Println(resultColor)

	//soal 4
	// var dataFilm = []movie{}
	// tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	// tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	// tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	// tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)
}

// func tambahDataFilm(judul string, duration int, )  {

// }
