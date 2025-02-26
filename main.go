package main

import (
	"calculator/Evaluator"
	"calculator/Lexer"
	"calculator/Parser"
	"fmt"
)

func main() {

	tokens := Lexer.Lexer([]rune("(12 + 38) / (10 + 23)"))
	fmt.Println(tokens)
	p := Parser.NewParser(tokens)
	ast := p.Parse()
	Parser.NodeMap(ast)
	res := Evaluator.Evaluate(ast)
	fmt.Println("result:", res)

}
