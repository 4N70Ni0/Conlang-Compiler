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
	// Operator
	OPPAR = 201 // (
	CLPAR = 202 // )
	COLON = 203
	COMMA = 204
	DASH  = 205
)

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
