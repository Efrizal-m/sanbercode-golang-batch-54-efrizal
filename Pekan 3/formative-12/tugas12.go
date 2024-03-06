package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     uint
}

var nilaiMahasiswa = []NilaiMahasiswa{}

func GetNilai(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataNilai, err := json.Marshal(nilaiMahasiswa)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataNilai)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

func PostNilai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var NilaiMhs NilaiMahasiswa
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&NilaiMhs); err != nil {
				log.Fatal(err)
			}
		} else {
			// parse dari form
			// getID := r.PostFormValue("id")
			getNilai := r.PostFormValue("nilai")
			// id, _ := strconv.Atoi(getID)
			nilai, _ := strconv.Atoi(getNilai)

			nama := r.PostFormValue("nama")
			mataKuliah := r.PostFormValue("mataKuliah")

			NilaiMhs = NilaiMahasiswa{
				Nama:        nama,
				MataKuliah:  mataKuliah,
				IndeksNilai: "",
				Nilai:       uint(nilai),
				ID:          0,
			}

		}

		uniqueNumber := time.Now().UnixNano() / (1 << 22)
		NilaiMhs.ID = uint(uniqueNumber)

		if NilaiMhs.Nilai > 100 {
			http.Error(w, "Bad request: Nilai tidak boleh > 100", http.StatusBadRequest)
		}
		NilaiMhs.IndeksNilai = getIndeksNilai(int(NilaiMhs.Nilai))

		nilaiMahasiswa = append(nilaiMahasiswa, NilaiMhs)

		dataNilai, _ := json.Marshal(NilaiMhs) // to byte
		w.Write(dataNilai)                     // cetak di browser
		return
	}

	http.Error(w, "NOT FOUND", http.StatusNotFound)
}

func getIndeksNilai(nilai int) string {
	if nilai >= 80 {
		return "A"
	} else if nilai >= 70 && nilai < 80 {
		return "B"
	} else if nilai >= 60 && nilai < 70 {
		return "C"
	} else if nilai >= 50 && nilai < 60 {
		return "D"
	} else {
		return "E"
	}
}
