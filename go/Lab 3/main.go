package main

import (
	"fmt"
	"lab3/lexer"
	"os"
)

func main(){

	// code, _:= os.ReadFile("output/example.txt")


	// l:= lexer.New(input)
	// token:= l.NextToken()

	// for (string(token.Type) != "EOF") {
	// 	token = l.NextToken()
	// 	fmt.Println(token.Type, token.Literal)
	// }

	input, _ := os.ReadFile("output/example.txt")
	l := lexer.New(string(input))

	for i := 0; i < 31; i++ {
		tok:= l.NextToken()
		fmt.Printf("%s - %s - %d\n", tok.Type, tok.Value, tok.Line)
	}

}
