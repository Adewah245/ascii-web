package main

import (
	"html/template"
	"net/http"
)

type Data struct {
	Art   string
	Error string
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	if err := LoadAllBanners(); err != nil {
		panic("Cannot load banners: " + err.Error())
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, Data{})
	})

	http.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Bad request", 400)
			return
		}

		text := r.FormValue("text")
		banner := r.FormValue("banner")

		if text == "" {
			tmpl.Execute(w, Data{Error: "Please type something!"})
			return
		}
		if banner == "" {
			tmpl.Execute(w, Data{Error: "Choose a banner"})
			return
		}

		b, ok := banners[banner]
		if !ok {
			tmpl.Execute(w, Data{Error: "Wrong banner"})
			return
		}

		lines := inputFile(text)
		art := Generate(lines, b)

		tmpl.Execute(w, Data{Art: art})
	})

	println("Server started → http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
