package main

import (
	"fmt"
	"os"
	"simple-text-editing-tool/binhexa"
	"simple-text-editing-tool/capitalize"
	"simple-text-editing-tool/lower"
	"simple-text-editing-tool/upper"
	"simple-text-editing-tool/vowel"
	"strings"
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

	// COnvert into slice of string
	toString := strings.Fields(string(readSample))
	
	// Manipulate the data with required functions
	toString= capitalize.Cap(toString)
	toString = lower.Low(toString)
	toString = binhexa.Bin(toString)
	toString = binhexa.Hex(toString)
	toString = upper.Up(toString)
	toString = vowel.Vow(toString)
	// toString = punctuation.Punc(toString)
	// toString = apostro.Apos(toString)

	// Join the slice of strings
	toJoin := strings.Join(toString, " ")
	// COnvert to Byte for writing
	toByte := []byte(toJoin)

	err = os.WriteFile(resultText, toByte, 0o644)
	if err != nil{
		println("Error writing to file :", err)
		return
	}

	fmt.Printf("Successfully Updated, %v \n", resultText)
}