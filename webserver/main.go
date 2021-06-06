package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		names := r.URL.Query()["name"]
		var name string
		if len(names) == 1 {
			name = names[0]
		}

		//w.Write([]byte("Web Server in GO - Your name is " + name))

		m := map[string]string{"name": name}
		enc := json.NewEncoder(w)
		enc.Encode(m)
	})

	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
