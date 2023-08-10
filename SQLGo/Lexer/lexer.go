package lexer

import (
	"regexp"
	"strings"
)

type Token struct {
	Type  string
	Value string
}

func Lex(statement string) []Token {
	// Tokenize the query
	tokens := []Token{}
	content := strings.Fields(statement)

	// Regular expressions for basic tokens
	keywordPattern := regexp.MustCompile(`(?i)(SELECT|FROM|WHERE|AND|OR)`)
	identifierPattern := regexp.MustCompile(`[a-zA-Z_]\w*`)
	operatorPattern := regexp.MustCompile(`[=><]`)
	stringLiteralPattern := regexp.MustCompile(`'[^']*'`)

	for _, word := range content {
		if keywordPattern.MatchString(word) {
			tokens = append(tokens, Token{Type: "KEYWORD", Value: word})
		} else if identifierPattern.MatchString(word) {
			tokens = append(tokens, Token{Type: "IDENTITY", Value: word})
		} else if operatorPattern.MatchString(word) {
			tokens = append(tokens, Token{Type: "OPERATOR", Value: word})
		} else if stringLiteralPattern.MatchString(word) {
			tokens = append(tokens, Token{Type: "LITERAL", Value: word})
		} else {
			tokens = append(tokens, Token{Type: "UNKNOWN", Value: word})
		}
	}

	return tokens
}
