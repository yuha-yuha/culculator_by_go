package Evaluator

import (
	"calculator/Lexer"
	"calculator/Parser"
)

func Evaluate(ast Parser.Node) int {
	switch ast.Token.TokenType {
	case Lexer.ADD:
		return Evaluate(*ast.LeftNode) + Evaluate(*ast.RightNode)
	case Lexer.SUB:
		return Evaluate(*ast.LeftNode) - Evaluate(*ast.RightNode)
	case Lexer.MUL:
		return Evaluate(*ast.LeftNode) * Evaluate(*ast.RightNode)
	case Lexer.DIV:
		return Evaluate(*ast.LeftNode) / Evaluate(*ast.RightNode)
	case Lexer.INT_NUM:
		return ast.Token.NumValue
	default:
		panic("evaluator: invalid syntax")
	}
}
