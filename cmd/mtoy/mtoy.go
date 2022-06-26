package main

import (
	"fmt"
	"os"

	"github.com/walterbio/mtoy"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage of mtoy:\n	mtoy \"question ?\"\n\n")
		os.Exit(1)
	}
	fmt.Println(mtoy.New().RevealAnswer())
}
