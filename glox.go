package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var hadError bool = false

func main() {
	args := os.Args
	if len(args) > 2 {
		fmt.Println("Usage: glox [script]")
	} else if len(args) == 2 {
		runFile(args[1])
	} else {
		runPrompt()
	}
}

func runFile(fileName string) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file")
	}
	run(string(content))
    if hadError {
        os.Exit(65)
    }
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
        ok := scanner.Scan()
        if !ok {
            os.Exit(1)
        }
        line := scanner.Text()
        if line == "exit" {
            break
        }
        run(line)
        hadError = false
	}
}

func run(source string) {
    tokens := strings.Fields(source)
    for i := range tokens {
        fmt.Println(tokens[i])
    }
	fmt.Println(tokens)
}

func error(line uint, message string) {
    report(line, "", message)
}

func report(line uint, where string, message string) {
    fmt.Println("[line " + fmt.Sprint(line) + "] Error" + where + ": " + message)
    hadError = true
}
