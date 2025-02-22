package main

import (
	"calculator/Lexer"
	"log"
)

func main() {
	tokens := Lexer.Lexer([]rune("(1 / 2) * 5"))
	log.Println(tokens)
}
