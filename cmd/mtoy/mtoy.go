package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/walterbio/mtoy"
)

func revealAnswer(w io.Writer) {
	seed := time.Now().Unix()
	fmt.Fprintln(w, mtoy.New(seed).RevealAnswer())
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage of mtoy:\n	mtoy \"question ?\"\n\n")
		os.Exit(1)
	}

	revealAnswer(os.Stdout)
}
