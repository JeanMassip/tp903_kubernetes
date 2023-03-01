package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("STARTING")
	file, err := os.OpenFile("/cache/fichier.doc", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write([]byte(os.Getenv("MESSAGE")))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ENDED")
}
