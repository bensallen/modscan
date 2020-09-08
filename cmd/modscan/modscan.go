package main

import (
	"fmt"
	"os"

	"github.com/bensallen/modscan/internal/cli/root"
)

func main() {
	if err := root.Run(os.Args[1:]); err != nil {
		fmt.Printf("Error: %v\n\n", err)
		os.Exit(1)
	}
}
