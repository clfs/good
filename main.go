package main

import (
	"fmt"

	"github.com/clfs/good/chess"
)

func main() {
	for _, v := range chess.TableQueenAttacks {
		fmt.Printf("%#016X\n", v)
	}

	// client := uci.New(os.Stdin, os.Stdout)
	// if err := client.Run(); err != nil {
	// 	log.Fatal(err)
	// }
}
