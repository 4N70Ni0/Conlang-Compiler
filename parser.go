package main

import (
	"fmt"
	"log"
	"slices"
)

type Symbol struct {
	Name       string
	IsOptional bool
}

type Parser struct {
	lex       Lexer
	CurToken  Token
	PeekToken Token
	Symbols   []Symbol // Variables declared.
	ErrLine   int
}

func (par *Parser) Init(lexer Lexer) {
	par.lex = lexer
	par.CurToken = Token{}
	par.PeekToken = Token{}
	par.Symbols = []Symbol{}
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

func (par *Parser) CheckDeclaredIdent() {
	tk := par.CurToken
	par.Match(IDENT)
	for _, sym := range par.Symbols {
		if tk.Text == sym.Name {
			return
		}
	}
	par.Abort("Undeclared identifier '" + tk.Text + "'")
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

		// OPIDENT used by STATEMENT-DECLARATION.
	} else if par.CheckToken(IDENT) || par.CheckToken(OPIDENT) {
		// (OPIDENT | IDENT)+ VALUES
		if par.CheckPeek(COMMA) || par.CheckPeek(VALUES) {
			fmt.Println("STATEMENT-DEFINE")
			par.CheckDeclaredIdent()

			// Loop through the idents until the values are reached.
			for !par.CheckToken(VALUES) {
				par.Match(COMMA)
				par.CheckDeclaredIdent()
			}
			par.Match(VALUES)

			// (IDENT | OPIDENT) "IF" (IDENT | OPIDENT) "IS" ["NOT"] VALUES
		} else if par.CheckPeek(IF) {
			fmt.Println("STATEMENT-IF")
			par.CheckDeclaredIdent()
			par.Match(IF)

			par.CheckDeclaredIdent()
			par.Match(IS)
			if par.CheckToken(NOT) {
				par.NextToken()
			}
			par.Match(VALUES)

		} else if (par.CheckPeek(IDENT) || par.CheckPeek(OPIDENT)) || par.CheckPeek(NEWLINE) {
			fmt.Println("STATEMENT-DECLARATION")

			// par.Match(IDENT) // or just NextToken?

			for par.CheckToken(IDENT) || par.CheckToken(OPIDENT) {
				symbol := Symbol{
					Name:       par.CurToken.Text,
					IsOptional: par.CheckToken(OPIDENT),
				}
				if !slices.Contains(par.Symbols, symbol) {
					par.Symbols = append(par.Symbols, symbol)
				} else {
					par.Abort("Variable '" + symbol.Name + "' cannot be re-declared.")
				}
				par.NextToken()
			}
			// fmt.Println(par.Symbols)
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
