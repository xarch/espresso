// Command espresso parses Espresso as defined by [1].
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alkchr/espresso/parser"
)

func main() {
	in := os.Stdin
	nm := "stdin"
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		in = f
		nm = os.Args[1]
	}

	got, err := parser.ParseReader(nm, in)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(got)
}
