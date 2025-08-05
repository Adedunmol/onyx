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
	scanner := bufio.NewScanner(os.Stdin)

	for {
		r.Print()

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println(err)
				//fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}
			fmt.Println("EOF")
			break
		}

		line := strings.TrimSpace(scanner.Text())

		fmt.Println("echo:", line)
	}
}
