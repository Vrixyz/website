package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"math/rand"
	"time"
	"log"
	"os"
	"bytes"
)

var templates = template.Must(template.ParseFiles("templates/index.html"))

func handler_root(w http.ResponseWriter, r *http.Request, buffer *bytes.Buffer) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//var files, _ = filepath.Glob("static\\/cv\\/.*\\.html")
var files, errfiles = filepath.Glob("static/cv/*\\.html")
var numberOfFiles = len(files)

var challengeFiles, err = filepath.Glob(".well-known/*")

var challenge = challengeFiles[0][1:len(challengeFiles[0])]

func handler_cv(w http.ResponseWriter, r *http.Request, buffer *bytes.Buffer) {
	randomIndex := rand.Int() % numberOfFiles
	buffer.WriteString(" | serving: /")
	buffer.WriteString(files[randomIndex])
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	w.Header().Set("Pragma", "no-cache") // HTTP 1.0.
	w.Header().Set("Expires", "0") // Proxies.
	http.ServeFile(w, r, files[randomIndex])
}

func ProfiledHandle(handler func(http.ResponseWriter, *http.Request, *bytes.Buffer)) func (w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var buffer bytes.Buffer
		handler(w, r, &buffer)

		// TODO: defer all this to manage possible crash in function handler
		end := time.Now()
		responseTime := end.Sub(start)
		log.Println(buffer.String(), "| request:", r, "| response time:", responseTime.Nanoseconds())
	}
}

func handler_tls(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path,";","."+challenge)
        if r.URL.Path[1:len(r.URL.Path)] == challenge {
          http.ServeFile(w, r, "." + challenge)
        } 
}

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
    // Redirect the incoming HTTP request.
    http.Redirect(w, r, "https://thierryberger.com:8081"+r.RequestURI, http.StatusMovedPermanently)
}

func main() {

	logFile, err := os.Create("/tmp/website-logs.txt")
	if (err != nil) {
		panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	http.HandleFunc("/well-known/", handler_tls)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/cv/", 	ProfiledHandle(handler_cv))
	http.HandleFunc("/", 			ProfiledHandle(handler_root))

	log.Println("server starts.");
	jsonfiles, _ := filepath.Glob("static/cv/*\\.json")
	files = append(files, jsonfiles...)
	numberOfFiles = len(files)
	log.Println("errfiles: ", errfiles, " ; files: ", files)

	go http.ListenAndServeTLS(":8083", "cert.pem", "key.pem", nil)
	http.ListenAndServe(":8082", nil)
	//http.ListenAndServeTLS(":8080", http.HandlerFunc(redirectToHttps))
}
