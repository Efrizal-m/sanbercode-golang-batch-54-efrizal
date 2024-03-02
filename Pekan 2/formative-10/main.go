package main

import (
	"errors"
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	//soal 1
	kalimat := "Golang Backend Development"
	tahun := 2021
	defer twoParamsFunc(kalimat, tahun)

	//soal 2
	fmt.Println(LuasPersegi(4, true))
	fmt.Println(LuasPersegi(8, false))
	fmt.Println(LuasPersegi(0, true))
	fmt.Println(LuasPersegi(0, false))

	//soal 3
	angka := 1
	defer cetakAngka(&angka)
	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	//soal 4
	var phones = []string{}
	tambahMerek(&phones, "Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo")
	showText(&phones)

	//soal 5
	var wg sync.WaitGroup

	wg.Add(1)
	showTextwithWG(&phones, &wg)
	wg.Wait()
}

func twoParamsFunc(kalimat string, tahun int) {
	fmt.Println(kalimat)
	fmt.Println(tahun)
}

func LuasPersegiCalc(Rusuk int, Status bool) (interface{}, error) {
	if Rusuk > 0 {
		if Status {
			return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d", Rusuk, int(math.Pow(float64(Rusuk), 2))), nil
		} else {
			return math.Pow(float64(Rusuk), 2), nil
		}
	} else if Rusuk == 0 {
		if Status {
			return nil, errors.New("maaf anda belum menginput sisi dari persegi")
		} else {
			setPanic(true)
			return nil, errors.New("maaf anda belum menginput sisi dari persegi")
		}
	} else {
		return nil, errors.New("maaf, input angka tidak valid")
	}
}

func LuasPersegi(number int, status bool) (result interface{}) {
	res, err := LuasPersegiCalc(number, status)
	if err != nil {
		result = err
	} else {
		result = res
	}
	return
}

func setPanic(err bool) {
	defer endApp()
	if err {
		panic("Panic Error")
	}

}

func endApp() {
	message := recover()
	fmt.Println("Recover from:", message)
}

func cetakAngka(angka *int) {
	fmt.Println(int(*angka))
}

func tambahAngka(newNumber int, angka *int) {
	*angka += newNumber
}

func tambahMerek(phones *[]string, phone ...string) {
	*phones = append(*phones, phone...)
}

func showText(phones *[]string) {
	for i, p := range *phones {
		time.Sleep(1000 * time.Millisecond)
		go fmt.Printf("%d. %s\n", i+1, p)
	}
}

func showTextwithWG(phones *[]string, wg *sync.WaitGroup) {
	for i, p := range *phones {
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("%d. %s\n", i+1, p)
	}
	wg.Done()
}
