package main

import (
	"fmt"
	"net/http"
)

func main() {
	// konfigurasi server
	server := &http.Server{
		Addr: ":8080",
	}

	// routing
	http.Handle("/", Auth(http.HandlerFunc(GetNilai)))
	http.Handle("/add", Auth(http.HandlerFunc(PostNilai)))

	// jalankan server
	fmt.Println("server running at http://localhost:8080")
	server.ListenAndServe()
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "admin" && pwd == "123" {
			next.ServeHTTP(w, r)
			return
		}
		w.Write([]byte("Username atau Password tidak sesuai"))
	})
}
