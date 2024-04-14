package main

import (
	"AvitoTestTask/cmd/app"
	"log"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}

}
