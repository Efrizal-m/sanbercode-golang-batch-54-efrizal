package main

import (
	h "f9/helpers"
	"fmt"
)

func main() {
	//soal 1
	s := h.SegitigaSamaSisi{Alas: 3, Tinggi: 4}
	p := h.PersegiPanjang{Panjang: 4, Lebar: 6}
	h.HitungDatar(s)
	h.HitungDatar(p)

	t := h.Tabung{JariJari: 10, Tinggi: 8}
	b := h.Balok{Panjang: 6, Lebar: 7, Tinggi: 8}
	h.HitungRuang(t)
	h.HitungRuang(b)

	//soal 2
	mobile := h.Phone{Name: "Samsung Galaxy Note 20", Brand: "Samsung Galaxy Note 20", Year: 2020, Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"}}
	h.ShowPhoneInfo(mobile)

	// //soal 3
	fmt.Println(h.LuasPersegi(4, true))
	fmt.Println(h.LuasPersegi(8, false))
	fmt.Println(h.LuasPersegi(0, true))
	fmt.Println(h.LuasPersegi(0, false))

	// //soal 4
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	prefixString := prefix.(string)
	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)

	total := h.Sum(angkaPertama) + h.Sum(angkaKedua)

	output := fmt.Sprintf("%s%s = %d", prefixString, h.AppendValues(angkaPertama, angkaKedua...), total)

	fmt.Println(output)
}
