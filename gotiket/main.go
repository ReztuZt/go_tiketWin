package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Menetapkan handler untuk route "/" untuk menampilkan index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Menetapkan handler untuk route "/static/" untuk menyajikan file statis (CSS dan JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Menjalankan server pada port 8080
	port := 8080
	fmt.Printf("Server running on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
