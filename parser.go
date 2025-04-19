package main

import (
	"fmt"
	"log"
)

type Parser struct {
	lex       Lexer
	CurToken  Token
	PeekToken Token
	ErrLine   int
}

func (par *Parser) Init(lexer Lexer) {
	par.lex = lexer
	par.CurToken = Token{}
	par.PeekToken = Token{}
	par.ErrLine = 1
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
		log.Fatal("Expected token '", kind, "' in line ", par.ErrLine, ", got token '", par.CurToken.Kind, "'")
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
		par.ErrLine++
		// Allow extra newlines.
		for par.CheckToken(NEWLINE) {
			par.NextToken()
			par.ErrLine++
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

	} else if par.CheckToken(IDENT) || par.CheckToken(OPIDENT) {
		// (OPIDENT | IDENT)+ VALUES
		if par.CheckPeek(COMMA) || par.CheckPeek(VALUES) {
			fmt.Println("STATEMENT-DEFINE")
			par.NextToken()

			// Loop through the idents until the values are reached.
			for !par.CheckToken(VALUES) {
				par.Match(COMMA)
				if par.CheckToken(OPIDENT) || par.CheckToken(IDENT) {
					par.NextToken()
				}
			}
			par.Match(VALUES)

			// (IDENT | OPIDENT) "IF" (IDENT | OPIDENT) "IS" ["NOT"] VALUES
		} else if par.CheckPeek(IF) {
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

		} else if par.CheckPeek(IDENT) || par.CheckPeek(NEWLINE) {
			fmt.Println("STATEMENT-DECLARATION")
			par.Match(IDENT) // or just NextToken?

			for !par.CheckPeek(NEWLINE) {
				par.Match(IDENT)
			}
			par.Match(IDENT)
		}
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
