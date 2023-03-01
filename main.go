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
		message, err := os.ReadFile("/cache/fichier.doc")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(err)
		fmt.Println(message)

		response := struct {
			Message string
		}{
			Message: string(message),
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
