package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for name, values := range r.Header {
			// Loop over all values for the name.
			for _, value := range values {
				fmt.Println(name, value)
				e := json.NewEncoder(w)
				e.Encode(name + ": " + value)
			}
		}
	})

	log.Println("Server listening on port 8080.")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
