package main

import (
	"fmt"
	"math"
)

func main() {
	//soal 1
	var luasLigkaran float64
	var kelilingLingkaran float64
	var radius float64 = 0.32

	fmt.Printf("before | luas : %f | keliling: : %f \n", luasLigkaran, kelilingLingkaran)
	updateLingkaran(&luasLigkaran, &kelilingLingkaran, &radius)
	fmt.Printf("after | luas : %f | keliling: : %f \n", luasLigkaran, kelilingLingkaran)

	//soal 2
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)

	//soal 3
	var buah = []string{}
	addFruits(&buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")
	viewFruits(buah)

	//soal 4
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)
	viewFilm(dataFilm)
}

func updateLingkaran(luasLigkaran *float64, kelilingLingkaran *float64, radius *float64) {
	*luasLigkaran = (3.14) * math.Pow(*radius, 2)
	*kelilingLingkaran = 2 * (3.14) * (*radius)
}

func introduce(sentence *string, name, gender, role, age string) {
	if gender == "laki-laki" {
		*sentence = fmt.Sprintf(`Pak %s adalah seorang %s yang berusia %s tahun`, name, role, age)
	} else if gender == "perempuan" {
		*sentence = fmt.Sprintf(`Bu %s adalah seorang %s yang berusia %s tahun`, name, role, age)
	}
}

func addFruits(fruits *[]string, fruit ...string) {
	*fruits = append(*fruits, fruit...)
}

func viewFruits(fruits []string) {
	for i, b := range fruits {
		fmt.Printf("%d. %s \n", i+1, b)
	}
}

func tambahDataFilm(title, duration, genre, tahun string, dataFilm *[]map[string]string) {
	*dataFilm = append(*dataFilm, map[string]string{"genre": genre, "jam": duration, "tahun": tahun, "title": title})
}

func viewFilm(dataFilm []map[string]string) {
	for i, b := range dataFilm {
		fmt.Printf("%d.\ttittle : %s \n\tduration : %s \n\tgenre : %s \n\tyear : %s \n", i+1, b["title"], b["jam"], b["genre"], b["tahun"])
	}
}
