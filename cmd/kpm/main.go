package main

import (
	"fmt"
	"os"

	"github.com/edsonmichaque/kabula/internal/cmd/kpm"
)

func main() {
	if err := kpm.New("kpm"); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
