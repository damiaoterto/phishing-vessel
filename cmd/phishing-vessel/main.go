package main

import (
	"log"

	"github.com/damiaoterto/phishing-vessel/internal/cmd"
)

func main() {
	// application entrypoint
	err := cmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
