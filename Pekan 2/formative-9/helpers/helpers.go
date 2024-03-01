package helpers

import (
	"f9/interfaces"
	"fmt"
	"math"
)

// =========================soal 1================================//
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

func (s SegitigaSamaSisi) Luas() int {
	return int((s.Alas * s.Tinggi) / 2)
}

func (s SegitigaSamaSisi) Keliling() int {
	return 3 * s.Alas
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

func (t Tabung) Volume() float64 {
	return float64(3.14 * math.Pow(t.JariJari, 2) * t.Tinggi)
}

func (t Tabung) LuasPermukaan() float64 {
	return float64(2 * 3.14 * t.JariJari * t.Tinggi)
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return float64(2 * (b.Panjang*b.Lebar + b.Panjang*b.Tinggi + b.Lebar*b.Tinggi))
}

func HitungDatar(d interfaces.HitungBangunDatar) {
	fmt.Println(d)
	fmt.Println(d.Luas())
	fmt.Println(d.Keliling())
}

func HitungRuang(r interfaces.HitungBangunRuang) {
	fmt.Println(r)
	fmt.Println(r.Volume())
	fmt.Println(r.LuasPermukaan())
}

//==============================================================//

// =========================soal 2================================//
type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

func (p Phone) Show() string {
	result := fmt.Sprintf("name: %s \nbrand: %s \nyear: %d \ncolors: ", p.Name, p.Brand, p.Year)
	for i, b := range p.Colors {
		if i == len(p.Colors)-1 {
			result += fmt.Sprintf(` %s`, b)
		} else {
			result += fmt.Sprintf(` %s,`, b)
		}

	}
	return result
}

func ShowPhoneInfo(p interfaces.PhoneInfo) {
	fmt.Println(p.Show())
}

//==============================================================//

// =========================soal 3================================//
func LuasPersegi(Rusuk int, Status bool) (result interface{}) {
	if Rusuk > 0 {
		if Status {
			result = fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d", Rusuk, int(math.Pow(float64(Rusuk), 2)))
		} else {
			result = math.Pow(float64(Rusuk), 2)
		}
	} else if Rusuk == 0 {
		if Status {
			result = "Maaf anda belum menginput sisi dari persegi"
		} else {
			result = nil
		}
	} else {
		result = "Maaf, input angka tidak valid"
	}
	return
}

//==============================================================//

// =========================soal 4================================//
func Sum(numbers []int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}
	return result
}

func AppendValues(numbers []int, values ...int) string {
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

//==============================================================//
