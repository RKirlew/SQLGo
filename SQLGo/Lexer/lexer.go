package lexer

import (
	"strings"
	"text/scanner"
)

type Token struct {
	Type  string
	Value string
}

func Lex(statement string) []Token {
	// Tokenize the query
	var tokens []Token
	var s scanner.Scanner
	s.Init(strings.NewReader(statement))
	s.Mode = scanner.ScanStrings | scanner.ScanIdents | scanner.ScanFloats | scanner.ScanInts

	for {
		tok := s.Scan()
		if tok == scanner.EOF {
			break
		}

		tokValue := s.TokenText()
		switch s.TokenText() {
		case "SELECT", "FROM", "WHERE", "AND", "OR":
			tokens = append(tokens, Token{Type: "KEYWORD", Value: tokValue})
		case "=", ">", "<":
			tokens = append(tokens, Token{Type: "OPERATOR", Value: tokValue})
		case "'":
			// Handle string literals
			literal := tokValue
			for {
				tok = s.Scan()
				tokValue = s.TokenText()
				if tok == scanner.EOF || tokValue == "'" {
					break
				}
				literal += tokValue
			}
			literal = literal[1:len(literal)]
			tokens = append(tokens, Token{Type: "LITERAL", Value: literal})
		default:
			// Handle identifiers, numbers, and other cases
			tokens = append(tokens, Token{Type: "IDENTIFIER", Value: tokValue})
		}
	}

	return tokens
}
