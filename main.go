package main

import (
	"os"
)

func main(){
	args := os.Args[1:]

	if len(args) != 2{
		println("Use: go run main.go sample.txt result.txt")
		return
	}

	sampleText := args[0]
	resultText := args[1]

	readSample, err := os.ReadFile(sampleText)
	if err != nil{
		println("Error reading from File: ",err)
		return
	}
}