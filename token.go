package main

import "strings"

type TokenType int

const (
	EOF     TokenType = -1
	NEWLINE           = 0
	NUMBER            = 1
	IDENT             = 2
	OPIDENT           = 3
	VALUES            = 4
	RANGE             = 5
	// Keywords
	IF        = 101
	IS        = 102
	NOT       = 103
	PRINT     = 104
	WORDS     = 105
	WITH      = 106
	SYLLABLES = 107
	ANY       = 108
	SKIP      = 109
	// Operator
	OPPAR = 201 // (
	CLPAR = 202 // )
	COLON = 203
	COMMA = 204
	DASH  = 205
)

func (tt TokenType) String() string {
	var name string

	switch tt {
	case EOF:
		name = "EOF"
	case NEWLINE:
		name = "NEWLINE"
	case NUMBER:
		name = "NUMBER"
	case IDENT:
		name = "IDENT"
	case OPIDENT:
		name = "OPIDENT"
	case VALUES:
		name = "VALUES"
	case RANGE:
		name = "RANGE"
	case IF:
		name = "IF"
	case IS:
		name = "IS"
	case NOT:
		name = "NOT"
	case PRINT:
		name = "PRINT"
	case WORDS:
		name = "WORDS"
	case WITH:
		name = "WITH"
	case SYLLABLES:
		name = "SYLLABLES"
	case ANY:
		name = "ANY"
	case OPPAR:
		name = "OPPAR"
	case CLPAR:
		name = "CLPAR"
	case COLON:
		name = "COLON"
	case COMMA:
		name = "COMMA"
	case DASH:
		name = "DASH"
	default:
		name = "???"
	}

	return name
}

type Token struct {
	Text string
	Kind TokenType
}

var keywords map[string]TokenType = map[string]TokenType{
	"if":        IF,
	"is":        IS,
	"not":       NOT,
	"print":     PRINT,
	"words":     WORDS,
	"with":      WITH,
	"syllables": SYLLABLES,
	"any":       ANY,
}

func IsKeyword(keyword string) bool {
	_, isOk := keywords[strings.ToLower(keyword)]
	return isOk
	// keywords := []string{"if", "is", "not", "print", "words", "with", "syllables", "any"}
	// return slices.Contains(keywords, strings.ToLower(keyword))
}

func GetKeywordKind(keyword string) TokenType {
	return keywords[strings.ToLower(keyword)]
}
