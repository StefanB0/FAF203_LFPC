package main

import (
	"fmt"
	"lab3/lexer"
)

func main() {
	input:= `function appendlazy(x){
		return x = "lazy" + x;
	}

	Program{

	let filename = "file.txt";

	rename(filename, appendlazy(filename));
	}
`

	l:= lexer.New(input)
	token:= l.NextToken()

	for (string(token.Type) != "EOF") {
		fmt.Println(token.Type, token.Literal)
		token = l.NextToken()
	}

}
