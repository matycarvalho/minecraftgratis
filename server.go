package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. Expose the "static" folder to the web under the "/static/" URL path
	fileServer := http.FileServer(http.Dir("./minecraft_files"))
	http.Handle("/minecraft_files/", http.StripPrefix("/minecraft_files/", fileServer))

	// 2. Route to serve your HTML page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	fmt.Println("Server starting at http://localhost:80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}

}

