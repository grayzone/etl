package main

import (
	"fmt"
	"github.com/grayzone/etl/util"
	"html/template"
	"log"
	"net/http"
	"encoding/json"
)

func renderHandler(w http.ResponseWriter, r *http.Request, templatepath string) {
	t, err := template.ParseFiles(templatepath)
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	renderHandler(w, r, "template/index.html")
}

func viewWorldHandler(w http.ResponseWriter, r *http.Request) {

//	result := getDeviceInCountries()
//	fmt.Fprint(w, result)

result := ""

fmt.Fprint(w, "ok")

}

func staticHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path[1:])

	http.ServeFile(w, r, r.URL.Path[1:])
}

func getDeviceInCountries() string{

	var db util.DBOps
	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.GetDeviceInCountry()
	if err != nil{
		log.Fatal(err)
	}
	log.Println(result)
	log.Println(len(result))
	
	b, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}

func main() {
	http.HandleFunc("/template/", staticHandler)
	http.HandleFunc("/", viewHandler)

	http.HandleFunc("/world", viewWorldHandler)

	http.ListenAndServe(":137", nil)

}
