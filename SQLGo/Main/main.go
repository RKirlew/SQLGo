package main

import (
	lexer "SQLGo/Lexer"
	"fmt"
)

func main() {

	sqlQuery := "SELECT column1, column2 FROM table WHERE column3 = 'value';"
	tokens := lexer.Lex(sqlQuery)

	// Print the tokens
	for _, token := range tokens {
		fmt.Printf("%s: %s\n", token.Type, token.Value)
	}
}
