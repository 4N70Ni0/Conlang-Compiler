package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("-- Conlang Compiler --")

	if len(os.Args) < 2 {
		log.Fatal("ERROR: ConCom needs source file as argument.")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	buf, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	source := string(buf)

	// source := "(    C1 ) (G) V(C2) (C3) "
	lex := Lexer{}
	lex.Init(source)

	par := Parser{}
	par.Init(lex)

	par.Program()
	fmt.Println("Parsing completed.")

	// token := lex.GetToken()
	// for ; token.Kind != EOF; token = lex.GetToken() {
	// 	fmt.Println("LEX: ", token.Text, token.Kind)
	// }
}
