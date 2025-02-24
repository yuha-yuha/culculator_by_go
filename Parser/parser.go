package Parser

import (
	"calculator/Lexer"
)

type Node struct {
	LeftNode  *Node
	RightNode *Node
}

func Parser(tokens []Lexer.Token) Node {
	node := Node{}
	expr(tokens)
	return
}

func expr(tokens []Lexer.Token) {

}
