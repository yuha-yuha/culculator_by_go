package Lexer

import (
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
	Ch        string
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
			tokens = append(tokens, Token{TokenType: INT_NUM, NumValue: num, Ch: numstr})
			continue
		}

		switch input[pos] {
		case '+':
			tokens = append(tokens, Token{TokenType: ADD, Ch: "+"})
		case '-':
			tokens = append(tokens, Token{TokenType: SUB, Ch: "-"})
		case '*':
			tokens = append(tokens, Token{TokenType: MUL, Ch: "*"})
		case '/':
			tokens = append(tokens, Token{TokenType: DIV, Ch: "/"})
		case '(':
			tokens = append(tokens, Token{TokenType: LP, Ch: "("})
		case ')':
			tokens = append(tokens, Token{TokenType: RP, Ch: ")"})
		default:
			tokens = append(tokens, Token{TokenType: INVALID, Ch: string(input[pos])})

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

	if *pos < len(input) {
		return true
	} else {
		return false
	}
}

func isDigit(ch rune) bool {
	if '0' <= ch && ch <= '9' {
		return true
	} else {
		return false
	}
}
