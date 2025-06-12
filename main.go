package main

import (
	"chain/tui"
	"fmt"
	"os"
)

func main() {
	p := tui.GetProgram()
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
