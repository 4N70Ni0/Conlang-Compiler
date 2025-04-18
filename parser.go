package main

import (
	"fmt"
	"log"
)

type Parser struct {
	lex       Lexer
	CurToken  Token
	PeekToken Token
}

func (par *Parser) Init(lexer Lexer) {
	par.lex = lexer
	par.CurToken = Token{}
	par.PeekToken = Token{}
	par.NextToken()
	par.NextToken()
}

// Return true if the current token matches.
func (par *Parser) CheckToken(kind TokenType) bool {
	return kind == par.CurToken.Kind
}

// Return true if the next token matches.
func (par *Parser) CheckPeek(kind TokenType) bool {
	return kind == par.PeekToken.Kind
}

// Try to match current token. If not, error. Advances the current token.
func (par *Parser) Match(kind TokenType) {
	if !par.CheckToken(kind) {
		log.Fatal("Expected token ", kind, ", got token ", par.CurToken.Kind)
	}
	par.NextToken()
}

// Advances the current token.
func (par *Parser) NextToken() {
	par.CurToken = par.PeekToken
	par.PeekToken = par.lex.GetToken()
	// Lexer handles the EOF
}

func (par *Parser) Abort(message string) {
	log.Fatal("Error. " + message)
}

// Parsing

func (par *Parser) Nl() {
	fmt.Println("NEWLINE")

	if !par.CheckToken(EOF) {
		// Require at least one new line otherwise, gives error.
		par.Match(NEWLINE)
		// Allow extra newlines.
		for par.CheckToken(NEWLINE) {
			par.NextToken()
		}
	}
}

func (par *Parser) Statement() {
	// Check the first token to see what kind of statement this is.

	// "PRINT"
	if par.CheckToken(PRINT) { // "PRINT" expression ["WORDS"] ["WITH"] expression ["SYLLABLES"] NL
		fmt.Println("STATEMENT-PRINT")
		par.NextToken()

		par.Match(NUMBER)
		// Optional.
		if par.CheckToken(WORDS) {
			par.NextToken()
		}
		// Optional.
		if par.CheckToken(WITH) {
			par.NextToken()
		}
		par.Match(RANGE)
		// Optional.
		if par.CheckToken(SYLLABLES) {
			par.NextToken()
		}

		// (IDENT | OPIDENT) "IF" (IDENT | OPIDENT) "IS" ["NOT"] VALUES
	} else if (par.CheckToken(IDENT) || par.CheckToken(OPIDENT)) && par.CheckPeek(IF) {
		fmt.Println("STATEMENT-IF")
		par.NextToken()
		par.NextToken()

		if par.CheckToken(IDENT) || par.CheckToken(OPIDENT) {
			par.NextToken()
		}
		par.Match(IS)
		if par.CheckToken(NOT) {
			par.NextToken()
		}
		par.Match(VALUES)
	}

	par.Nl()
}

// program ::= {statement}
func (par *Parser) Program() {
	fmt.Println("PROGRAM")

	for !par.CheckToken(EOF) {
		par.Statement()
	}
}
