package main

import(
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func main() {
	// wait for user input of integer
	var input int
	fmt.Print("Enter an integer: ")
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		return
	}

	// switch on input
	switch input {
		case 1:
			fmt.Println("Assignment One")
			intSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			for _, value := range intSlice {
				if value % 2 == 0 {
					fmt.Println(value, "is even")
				} else {
					fmt.Println(value, "is odd")
				}
			}
			return
		case 2:
			fmt.Println("Assignment Two")
		case 3:
			fmt.Println("Assignment Three")
		default:
			println("LET'S SPIN UP A SERVER")
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
}
