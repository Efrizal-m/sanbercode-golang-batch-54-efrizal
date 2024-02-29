package main

import (
	"fmt"
	"math"
)

func main() {
	//soal 1
	s := segitigaSamaSisi{3, 4}
	p := persegiPanjang{4, 6}
	hitungDatar(s)
	hitungDatar(p)

	t := tabung{10, 8}
	b := balok{6, 7, 8}
	hitungRuang(t)
	hitungRuang(b)

	//soal 2
	mobile := phone{name: "Samsung Galaxy Note 20", brand: "Samsung Galaxy Note 20", year: 2020, colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"}}
	showPhoneInfo(mobile)

	//soal 3
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	//soal 4
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	prefixString := prefix.(string)
	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)

	total := sum(angkaPertama) + sum(angkaKedua)

	output := fmt.Sprintf("%s%s = %d", prefixString, appendValues(angkaPertama, angkaKedua...), total)

	fmt.Println(output)
}

// ---------------------nomor 1------------------------//
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return int((s.alas * s.tinggi) / 2)
}

func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (t tabung) volume() float64 {
	return float64(3.14 * math.Pow(t.jariJari, 2) * t.tinggi)
}

func (t tabung) luasPermukaan() float64 {
	return float64(2 * 3.14 * t.jariJari * t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return float64(2 * (b.panjang*b.lebar + b.panjang*b.tinggi + b.lebar*b.tinggi))
}

func hitungDatar(d hitungBangunDatar) {
	fmt.Println(d)
	fmt.Println(d.luas())
	fmt.Println(d.keliling())
}

func hitungRuang(r hitungBangunRuang) {
	fmt.Println(r)
	fmt.Println(r.volume())
	fmt.Println(r.luasPermukaan())
}

//----------------------------------------------------//

// ---------------------nomor 2------------------------//
type phone struct {
	name, brand string
	year        int
	colors      []string
}

type phoneInfo interface {
	show() string
}

func (p phone) show() string {
	result := fmt.Sprintf("name: %s \nbrand: %s \nyear: %d \ncolors: ", p.name, p.brand, p.year)
	for i, b := range p.colors {
		if i == len(p.colors)-1 {
			result += fmt.Sprintf(` %s`, b)
		} else {
			result += fmt.Sprintf(` %s,`, b)
		}

	}
	return result
}

func showPhoneInfo(p phoneInfo) {
	fmt.Println(p.show())
}

// ----------------------------nomor 3----------------------//
func luasPersegi(rusuk int, status bool) (result interface{}) {
	if rusuk > 0 {
		if status {
			result = fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d", rusuk, int(math.Pow(float64(rusuk), 2)))
		} else {
			result = math.Pow(float64(rusuk), 2)
		}
	} else if rusuk == 0 {
		if status {
			result = "Maaf anda belum menginput sisi dari persegi"
		} else {
			result = nil
		}
	} else {
		result = "Maaf, input angka tidak valid"
	}
	return
}

// ----------------------------nomor 4----------------------//
func sum(numbers []int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}
	return result
}

func appendValues(numbers []int, values ...int) string {
	result := ""
	for i, num := range numbers {
		result += fmt.Sprint(num)
		if i < len(numbers)-1 || len(values) > 0 {
			result += " + "
		}
	}
	for i, value := range values {
		result += fmt.Sprint(value)
		if i < len(values)-1 {
			result += " + "
		}
	}
	return result
}
