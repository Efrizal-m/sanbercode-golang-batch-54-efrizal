package main

import (
	"fmt"
)

func main() {
	// soal 1
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%2 == 1 {
			fmt.Printf("%d - I Love Coding\n", i)
		} else if i%2 == 0 {
			fmt.Printf("%d - Berkualitas\n", i)
		} else if i%2 == 1 {
			fmt.Printf("%d - Santai\n", i)
		}
	}

	// soal 2
	for i := 0; i < 7; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("#")
		}
		fmt.Printf("\n")
	}

	// soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	var outputKata = make([]string, 5)
	copy(outputKata, kalimat[2:])
	fmt.Println(outputKata)

	// soal 4
	var sayuran = []string{}
	sayuran = append(sayuran,
		"Bayam",
		"Buncis",
		"Kangkung",
		"Kubis",
		"Seledri",
		"Tauge",
		"Timun",
	)
	for i := 0; i < len(sayuran); i++ {
		fmt.Printf("%d.%s\n", i+1, sayuran[i])
	}

	//soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	var volumeBalok int = 1

	for k, s := range satuan {
		volumeBalok = volumeBalok * s
		fmt.Printf("%s = %d\n", k, s)
	}
	fmt.Printf("Volume Balok = %d\n", volumeBalok)
}
