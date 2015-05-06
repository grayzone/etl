package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type CityLevel struct {
	City  string
	Total int
}

type ProvinceLevel struct {
	Province string
	Total    int
}

type CountryLevel struct {
	Country string
	Total   int
}

func renderHandler(w http.ResponseWriter, r *http.Request, templatepath string) {
	t, err := template.ParseFiles(templatepath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	t.Execute(w, nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	renderHandler(w, r, "template/index.html")
}

func viewWorldHandler(w http.ResponseWriter, r *http.Request) {


	fmt.Fprint(w, "Hello")
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Println(r.URL.Path[1:])

	http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/template/", staticHandler)
	http.HandleFunc("/", viewHandler)

	http.HandleFunc("/world", viewWorldHandler)
























	http.ListenAndServe(":137", nil)

}
