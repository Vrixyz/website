package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"fmt"
	"math/rand"
	"time"
)

var templates = template.Must(template.ParseFiles("templates/index.html"))

func handler_root(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//var files, _ = filepath.Glob("static\\/cv\\/.*\\.html")
var files, errfiles = filepath.Glob("static/cv/*\\.html")
var numberOfFiles = len(files)

func handler_cv(w http.ResponseWriter, r *http.Request) {
	randomIndex := rand.Int() % numberOfFiles
	fmt.Println("randomIndex: ", randomIndex)
	fmt.Println("serving: ", files[randomIndex])
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	w.Header().Set("Pragma", "no-cache") // HTTP 1.0.
	w.Header().Set("Expires", "0") // Proxies.
	http.ServeFile(w, r, files[randomIndex])
}

func ProfiledHandle(handler func(http.ResponseWriter, *http.Request)) func (w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler(w, r) // TODO: pass something to the function so we input information, then output all data from request in one line (easy aggregation, avoid id..)

		// TODO: defer all this to manage possible crash in function handler
		end := time.Now()
		responseTime := end.Sub(start)
		fmt.Println("request: ", r, " | response time: ", responseTime)
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
  	http.Handle("/static/", http.StripPrefix("/static/", fs))	
	
	http.HandleFunc("/cv/", 	ProfiledHandle(handler_cv))
	http.HandleFunc("/", 			ProfiledHandle(handler_root))
	
	fmt.Println("server starts.");
	jsonfiles, _ := filepath.Glob("static/cv/*\\.json")	
	files = append(files, jsonfiles...)
	numberOfFiles = len(files)
	fmt.Println("errfiles: ", errfiles, " ; files: ", files)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println("error: ", err)
}
