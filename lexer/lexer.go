package Lexer

import (
	"fmt"
	"log"
	"strconv"
	"unicode"
)

const (
	ADD = iota
	SUB
	MUL
	DIV
	LP
	RP
	INT_NUM
	INVALID
)

type Token struct {
	TokenType int
	NumValue  interface{}
}

func Lexer(input []rune) []Token {
	pos := 0
	tokens := make([]Token, 0)

	for pos < len(input) {
		SkipSpace(&pos, input)

		if isDigit(input[pos]) {
			stapos := pos
			var endpos int
			for {
				if !NextPos(&pos, input) {
					endpos = pos
					break
				}
				if !(isDigit(input[pos])) {
					endpos = pos
					break
				}
			}

			numstr := string(input[stapos:endpos])

			num, err := strconv.Atoi(numstr)

			if err != nil {
				log.Println("str = ", numstr)
				panic("stringからnumに変換できませんでした")
			}
			tokens = append(tokens, Token{TokenType: INT_NUM, NumValue: num})
			continue
		}

		switch input[pos] {
		case '+':
			tokens = append(tokens, Token{TokenType: ADD})
		case '-':
			tokens = append(tokens, Token{TokenType: SUB})
		case '*':
			tokens = append(tokens, Token{TokenType: MUL})
		case '/':
			tokens = append(tokens, Token{TokenType: DIV})
		case '(':
			tokens = append(tokens, Token{TokenType: LP})
		case ')':
			tokens = append(tokens, Token{TokenType: RP})
		default:
			tokens = append(tokens, Token{TokenType: INVALID})

		}

		NextPos(&pos, input)
	}

	return tokens
}

func SkipSpace(pos *int, input []rune) {
	for unicode.IsSpace([]rune(input)[*pos]) {
		NextPos(pos, input)
	}
}

func NextPos(pos *int, input []rune) bool {
	*pos += 1
	fmt.Println(*pos)

	if *pos < len(input) {
		return true
	} else {
		return false
	}
}

func isDigit(ch rune) bool {
	if '1' <= ch && ch <= '9' {
		return true
	} else {
		return false
	}
}
