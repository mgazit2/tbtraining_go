package main

import(
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func main() {
    println("Hello, world!")
	var count int = 0
	var port string = "3001"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		if err != nil {
			return 
		}
	})
	http.HandleFunc("/counter", func(w http.ResponseWriter, r *http.Request) {
		count++
		_, err := fmt.Fprintf(w, "You have been here %d times", count)
		if err != nil {
			return 
		}
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "About page")
		if err != nil {
			return 
		}
	})
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "The time is %v", time.Now())
		if err != nil {
			return 
		}
	})
	log.Printf("Listening on localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
