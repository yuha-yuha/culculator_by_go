package main

import (
	"calculator/Lexer"
	"calculator/Parser"
	"fmt"
)

func main() {

	tokens := Lexer.Lexer([]rune("(5 + 4) * 2"))
	fmt.Println(tokens)
	p := Parser.NewParser(tokens)
	ast := p.Parse()
	Parser.NodeMap(ast)

}
