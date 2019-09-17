package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func getData() map[string]interface{} {
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var tmpl, err = template.ParseFiles("./template.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, getData())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
