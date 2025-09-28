package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Println("Hello, application startup.")
		time.Sleep(1 * time.Second)
	}
}
