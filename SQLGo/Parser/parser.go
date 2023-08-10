package parser

import (
	lexer "SQLGo/Lexer"
	"fmt"
)

type Node interface {
}

type SelectNode struct {
	Columns []Node
	From    Node
	Where   Node
}

type ColumnNode struct {
	Name string
}

type TableNode struct {
	Name string
}

type OperatorNode struct {
	Operator string // Operator symbol
	Left     Node   // Left operand
	Right    Node   // Right operand
}

// LiteralNode represents a literal value in the AST
type LiteralNode struct {
	Value string // The literal value
}

func Parse(fetchedTokens []lexer.Token) Node {
	var rootNode Node

	for i := 0; i < len(fetchedTokens); i++ {
		token := fetchedTokens[i]

		switch token.Type {
		case "KEYWORD":
			// Handle different SQL keywords (SELECT, FROM, WHERE, etc.)
			// Create appropriate nodes and establish relationships

			switch token.Value {
			case "SELECT":
				// Create a SelectNode and set it as the root of your AST
				selectNode := &SelectNode{}
				rootNode = selectNode

				i, err := parseColumns(fetchedTokens, i+1, selectNode)
				_ = i
				if err != nil {
					// Handle parsing error
				}

				// Parsing logic for the rest of the SELECT statement
				// You need to consider columns, FROM clause, WHERE clause, etc.
				// You might want to define separate functions to handle different parts of the SELECT statement.

			case "FROM":
				// Handle the FROM clause
				// Create a TableNode or appropriate node type, set relationships, etc.

			case "WHERE":
				// Handle the WHERE clause
				// Create an appropriate node type, set relationships, etc.

			// ... Handle other SQL keywords here ...

			default:
				// Handle unknown keywords or syntax errors
			}

		case "IDENTIFIER":
			// Handle identifiers (columns, tables, etc.)
			// Create appropriate nodes and establish relationships
		case "OPERATOR":
			// Handle operators (=, >, <, etc.)
			// Create appropriate nodes and establish relationships
		case "LITERAL":
			// Handle literals ('value', 123, etc.)
			// Create appropriate nodes and establish relationships
		default:
			// Handle unknown tokens or syntax errors
		}
	}

	return rootNode
}
func parseColumns(tokens []lexer.Token, startIndex int, selectNode *SelectNode) (int, error) {
	i := startIndex

	// Loop through tokens and parse columns
	for i < len(tokens) {
		token := tokens[i]

		switch token.Type {
		case "IDENTIFIER":
			// Create a ColumnNode for the identifier
			columnNode := &ColumnNode{Name: token.Value}

			// Add the ColumnNode to the Columns field of the SelectNode
			selectNode.Columns = append(selectNode.Columns, columnNode)

			// Move to the next token
			i++

			// If the next token is a comma, continue parsing the next column
			if i < len(tokens) && tokens[i].Value == "," {
				i++
				continue
			}

		default:
			// Handle syntax errors or other cases
			return i, fmt.Errorf("Expected column identifier, got %s", token.Value)
		}
	}

	// Return the index after parsing columns
	return i, nil
}

func parseFromClause(tokens []lexer.Token, startIndex int, selectNode *SelectNode) (int, error) {
	// Parse the table reference and create a TableNode
	tableToken := tokens[startIndex]
	tableNode := &TableNode{Name: tableToken.Value}

	// Set the TableNode as the From field of the SelectNode
	selectNode.From = tableNode

	// Return the index after parsing the FROM clause
	return startIndex + 1, nil
}
