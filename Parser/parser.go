package Parser

import (
	"calculator/Lexer"
	"encoding/json"
	"fmt"
)

type Node struct {
	Token     Lexer.Token
	LeftNode  *Node
	RightNode *Node
}

func NewNode(token Lexer.Token, leftNode *Node, rightNode *Node) Node {
	node := Node{}
	node.Token = token
	node.LeftNode = leftNode
	node.RightNode = rightNode
	return node
}

func NodeMap(n Node) {

	j, _ := json.Marshal(n)

	fmt.Println(string(j))

}

type Parser struct {
	Pos    int
	tokens []Lexer.Token
}

func NewParser(tokens []Lexer.Token) Parser {
	p := Parser{}

	p.Pos = -1
	p.tokens = tokens

	return p
}

func (p *Parser) Parse() Node {
	// あとでtokenのlenがゼロの時の処理かく
	node := p.expr()
	return node
}

func (p *Parser) expr() Node {
	node := p.factor()

	for {

		if !p.IsExistNextPos() {
			return node
		}

		if (p.tokens[p.Pos+1].TokenType == Lexer.ADD) || (p.tokens[p.Pos+1].TokenType == Lexer.SUB) {
			LeftNode := node
			p.NextPos()
			node.Token = p.tokens[p.Pos]
			node.LeftNode = &LeftNode
			RightNode := p.factor()
			node.RightNode = &RightNode
		} else {
			return node
		}

		fmt.Println("expr:", p.tokens[p.Pos].Ch)
	}
}

func (p *Parser) factor() Node {
	node := p.primary()

	for {
		if !p.IsExistNextPos() {
			return node
		}

		if (p.tokens[p.Pos+1].TokenType == Lexer.MUL) || (p.tokens[p.Pos+1].TokenType == Lexer.DIV) {
			LeftNode := node
			p.NextPos()
			node.Token = p.tokens[p.Pos]
			RightNode := p.primary()
			node.LeftNode = &LeftNode
			node.RightNode = &RightNode
		} else {
			return node
		}
		fmt.Println("factor:", p.tokens[p.Pos].Ch)
	}

}

func (p *Parser) primary() Node {
	node := Node{}
	fmt.Println("primary:", p.tokens[p.Pos+1].Ch)
	if p.tokens[p.Pos+1].TokenType == Lexer.INT_NUM {
		p.NextPos()
		node.Token = p.tokens[p.Pos]
	} else if p.tokens[p.Pos+1].TokenType == Lexer.LP {
		p.NextPos()
		node = p.expr()
		if !p.IsExistNextPos() || p.tokens[p.Pos+1].TokenType != Lexer.RP {
			panic("syntax error: expected token ')'")
		}
		p.NextPos()

	} else {
		panic("syntax error")
	}

	return node
}

func (p *Parser) NextPos() bool {
	if p.Pos+1 < len(p.tokens) {
		p.Pos += 1
		return true
	} else {
		return false
	}
}
func (p *Parser) IsExistNextPos() bool {
	if p.Pos+1 < len(p.tokens) {
		return true
	} else {
		return false
	}
}
