package lexer

const (
	ADD = iota
	SUB
	MUL
	DIV
	LP
	RP
	INT_NUM
)

type Token struct {
	TokenType int
	NumValue  interface{}
}

func lexer(input string) []Token {
	pos := 0
	tokens := make([]Token, len(input))

	for pos < len(input) {
		switch input[pos] {

		}
	}

	return
}
