package main

import (
	"log"
	"os"

	"github.com/clfs/good/uci"
)

func main() {
	client := uci.New(os.Stdin, os.Stdout)
	if err := client.Run(); err != nil {
		log.Fatal(err)
	}
}
