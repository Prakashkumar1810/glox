package util

import "fmt"

var HadError bool

func Error(line uint, message string) {
	report(line, "", message)
}

func report(line uint, where string, message string) {
	fmt.Println("[line " + fmt.Sprint(line) + "] Error" + where + ": " + message)
	HadError = true
}
