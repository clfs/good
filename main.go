package main

import (
	"log"
	"os"

	"github.com/clfs/good/uci"
)

func main() {
	// for i, v := range chess.TableBlackPawnPushes {
	// 	if i%4 == 0 {
	// 		fmt.Printf("\n")
	// 	}
	// 	fmt.Printf("%#016X, ", v)
	// }
	client := uci.New(os.Stdin, os.Stdout)
	if err := client.Run(); err != nil {
		log.Fatal(err)
	}
}
