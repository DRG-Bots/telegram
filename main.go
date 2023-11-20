package main

import (
	"os"
	"log"

	"github.com/mymmrac/telego"
)

func main() {
	token := ""

	_, err := telego.NewBot(token, telego.WithDefaultDebugLogger())

	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}


