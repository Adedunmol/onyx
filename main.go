package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	repl := Repl{}
	repl.Run()
}

type Repl struct{}

func (r *Repl) Print() {
	fmt.Print("User> ")
}

func (r *Repl) Run() {
	for {
		r.Print()
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			err := scanner.Err()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Scanner error: %v\n", err)
			} else {
				fmt.Print("\nExiting.")
			}
			break
		}

		line := strings.TrimSpace(scanner.Text())

		fmt.Println("Echo:", line)
	}
}
