package main

import (
	"fmt"
	"os"
)

func main() {
	if err := loadTask(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else if err := Execute(os.Args); err != nil {
		fmt.Println("error:", err)
		ShowHelp()
	} else if err := saveTask(); err != nil {
		fmt.Fprintln(os.Stderr, saveTask())
	}
}
