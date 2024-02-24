package main

import "fmt"

func main() {
	// soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	// soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john)

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)

	// soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("John", buah...)
	fmt.Println(buahFavoritJohn)

	// soal 4
	var dataFilm = []map[string]string{}
	var tambahDataFilm = func(title, duration, genre, tahun string) (result []map[string]string) {
		result = append(result, map[string]string{"genre": genre, "jam": duration, "tahun": tahun, "title": title})
		return result
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}

// function soal 1
// ----------------------------------------------------------//
func luasPersegiPanjang(panjang, lebar int) int {
	return panjang * lebar
}
func kelilingPersegiPanjang(panjang, lebar int) int {
	return 2 * (panjang + lebar)
}
func volumeBalok(panjang, lebar, tinggi int) int {
	return panjang * lebar * tinggi
}

//----------------------------------------------------------//

// function soal 2
// ----------------------------------------------------------//
func introduce(name, gender, role, age string) (result string) {
	if gender == "laki-laki" {
		result = fmt.Sprintf(`Pak %s adalah seorang %s yang berusia %s tahun`, name, role, age)
	} else if gender == "perempuan" {
		result = fmt.Sprintf(`Bu %s adalah seorang %s yang berusia %s tahun`, name, role, age)
	}
	return
}

//----------------------------------------------------------//

// function soal 3
// ----------------------------------------------------------//
func buahFavorit(name string, buah ...string) (result string) {
	result = fmt.Sprintf(`halo nama saya %s dan buah favorit saya adalah `, name)
	for i, b := range buah {
		if i == len(buah)-1 {
			result += fmt.Sprintf(`"%s"`, b)
		} else {
			result += fmt.Sprintf(`"%s", `, b)
		}
	}
	return result
}

//----------------------------------------------------------//
