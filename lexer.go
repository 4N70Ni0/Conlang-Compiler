package main

import (
	"fmt"
	"os"
)

func isAlphaNumeric(char string) bool {
	// Check if the byte value falls within the range of alphanumeric characters
	c := []byte(char)[0]
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func isDigit(char string) bool {
	// Check if the byte value falls within the range of alphanumeric characters
	c := []byte(char)[0]
	return (c >= '0' && c <= '9')
}

type Lexer struct {
	Source  string
	CurChar string
	CurPos  int
}

func (lex *Lexer) Init(source string) {
	lex.Source = source + "\n"
	lex.CurChar = ""
	lex.CurPos = -1
	lex.NextChar()
}

// Process the next character.
func (lex *Lexer) NextChar() {
	lex.CurPos++
	if lex.CurPos >= len(lex.Source) {
		lex.CurChar = "\000" // EOF
	} else {
		lex.CurChar = string([]rune(lex.Source)[lex.CurPos])
	}
}

// Return the lookahead character.
func (lex *Lexer) Peek() string {
	if lex.CurPos+1 >= len(lex.Source) {
		return "\000"
	}
	return string([]rune(lex.Source)[lex.CurPos+1])
}

// Invalid token found, print error message and exit.
func (lex *Lexer) Abort(message string) {
	fmt.Println("Lexing error. " + message)
	os.Exit(1)
}

// Skip whitespace except newlines, which will be used to indicate the end of a statement.
func (lex *Lexer) SkipWhitespace() {
	for lex.CurChar == " " || lex.CurChar == "\t" || lex.CurChar == "\r" {
		lex.NextChar()
	}
}

func (lex *Lexer) SkipComment() {
	if lex.CurChar == "#" {
		for lex.CurChar != "\n" {
			lex.NextChar()
		}
	}
}

func (lex *Lexer) GetToken() Token {
	lex.SkipWhitespace()
	lex.SkipComment()

	var token Token

	// Check the first character.
	switch lex.CurChar {
	// Newline token.
	case "\n":
		token = Token{lex.CurChar, NEWLINE}
	// EOF token.
	case "\000":
		token = Token{lex.CurChar, EOF}
	case ",":
		token = Token{lex.CurChar, COMMA}
	// Define values for identifier.
	case ":":
		token = Token{lex.CurChar, COLON}
		lex.NextChar()
		lex.SkipWhitespace()

		if lex.CurChar == "\n" {
			lex.Abort("No values defined for IDENT")
		}

		startPos := lex.CurPos
		for lex.CurChar != "\n" {
			lex.NextChar()
		}
		tokText := lex.Source[startPos : lex.CurPos-1]
		token = Token{tokText, VALUES}

	// Open parenthesis token.
	case "(":
		lex.NextChar()
		lex.SkipWhitespace()

		if !isAlphaNumeric(lex.CurChar) {
			lex.Abort("Invalid character '" + lex.CurChar + "' in an optional ident.")
		}
		// Get the identifier.
		startPos := lex.CurPos
		for isAlphaNumeric(lex.Peek()) {
			lex.NextChar()
		}
		tokText := lex.Source[startPos : lex.CurPos+1]
		lex.NextChar()

		lex.SkipWhitespace()
		if lex.CurChar != ")" {
			lex.Abort("Expecting ')', character '" + lex.CurChar + "' was found.")
		}

		if IsKeyword(tokText) {
			lex.Abort("Reserved word '" + tokText + "' cannot be inside a optional ident declaration.")
		}
		token = Token{tokText, OPIDENT}

	// Uknown token.
	default:
		// Check if it is a number.
		if isDigit(lex.CurChar) {
			// The leading character is a digit.
			startPos := lex.CurPos
			for isDigit(lex.Peek()) {
				lex.NextChar()
			}

			if lex.Peek() == "-" { // Range
				lex.NextChar()

				if !isDigit(lex.Peek()) {
					lex.Abort("Illegal character in range.")
				}
				for isDigit(lex.Peek()) {
					lex.NextChar()
				}

				tokText := lex.Source[startPos : lex.CurPos+1]
				token = Token{tokText, RANGE}
			} else {
				tokText := lex.Source[startPos : lex.CurPos+1]
				token = Token{tokText, NUMBER}
			}

		} else if isAlphaNumeric(lex.CurChar) {
			// Check if it is an IDENT of a KEYWORD.
			startPos := lex.CurPos
			for isAlphaNumeric(lex.Peek()) {
				lex.NextChar()
			}
			tokText := lex.Source[startPos : lex.CurPos+1]

			if lex.Peek() == " " {
				lex.SkipWhitespace()
			} else {
				lex.NextChar()
			}

			if lex.CurChar == "," { // RANGE
				for lex.Peek() != "\n" {
					lex.NextChar()

				}
				tokText := lex.Source[startPos:lex.CurPos]
				// fmt.Println(tokText, "VALUES")
				token = Token{tokText, VALUES}
			} else if IsKeyword(tokText) { // KEYWORD
				token = Token{tokText, GetKeywordKind(tokText)}
			} else { // IDENT
				token = Token{tokText, IDENT}
			}

		} else {
			lex.Abort("Uknown token: " + lex.CurChar)
		}
	}

	lex.NextChar()
	return token
}
