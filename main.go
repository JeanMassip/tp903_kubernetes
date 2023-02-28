package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := struct {
			Message string
		}{
			Message: os.Getenv("MESSAGE"),
		}

		res, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(res)

	})

	port := os.Getenv("APPPORT")
	fmt.Printf("running on http://127.0.0.1:%v\n", port)
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), nil); err != nil {
		log.Fatal(err)
	}
}
