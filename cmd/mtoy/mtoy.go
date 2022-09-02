package main

import (
	"fmt"
	"io"
	"os"

	"github.com/walterbio/mtoy"
)

func revealAnswer(w io.Writer) {
	fmt.Fprintln(w, mtoy.RevealAnswer())
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage of mtoy:\n	mtoy \"question ?\"\n\n")
		os.Exit(123)
	}

	revealAnswer(os.Stdout)
}
