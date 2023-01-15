package main

import (
	"bufio"
	"fmt"
	"glox/scanner"
	"glox/util"
	"io/ioutil"
	"os"
)

type hello string

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
	if util.HadError {
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
		util.HadError = false
	}
}

func run(source string) {
    scnr := scanner.NewScanner(source)
	tokens := scnr.ScanTokens()
	for i := range tokens {
		fmt.Println(tokens[i].Lexeme)
	}
}
