package main

import (
	"fmt"
	"lab3/lexer"
	"os"
)

func main(){

	input, _ := os.ReadFile("input/example.txt")
	l := lexer.New(string(input))

	for i := 0; i < 33; i++ {
		tok:= l.NextToken()
		fmt.Printf("%d, %s - %s - %d\n", i+1, tok.Type, tok.Value, tok.Line)
	}

}
