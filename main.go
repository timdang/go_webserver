package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Page of the web site
type Page struct {
	Name string
}

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := Page{Name: "Gopher"}

		// Assigns name of the p object from the url
		if name := r.FormValue("name"); name != "" {
			p.Name = name
		}
		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "Hello from func main")
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
