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

func Parse(fetchedTokens []lexer.Token) (Node, error) {
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

				selectNode := &SelectNode{}
				_ = selectNode
				newIndex, err := parseFromClause(fetchedTokens, i+1, selectNode)
				if err != nil {
					return nil, err
				}
				i = newIndex

			case "WHERE":
				selectNode := &SelectNode{}
				whereNode, newIndex, err := parseWhereClause(fetchedTokens, i+1)
				if err != nil {
					return nil, err
				}
				i = newIndex - 1 // Decrement by 1 because the loop will increment it

				// Set the WHERE clause node in the SelectNode
				selectNode.Where = whereNode

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

	return rootNode, nil
}

/*
The parseColumns function is responsible for parsing and extracting the column names following the "SELECT" keyword in an SQL query.

	It iterates through the list of tokens, identifies column names (identifiers), and adds them to the Columns field of the SelectNode as ColumnNode instances.
*/
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
	// Ensure there are tokens left to parse
	if startIndex >= len(tokens) {
		return startIndex, fmt.Errorf("Expected table name, got end of input")
	}

	// Get the token representing the table name
	tableToken := tokens[startIndex]

	if tableToken.Type != "IDENTIFIER" {
		return startIndex, fmt.Errorf("Expected table name, got %s", tableToken.Value)
	}

	// Create a TableNode for the table reference
	tableNode := &TableNode{Name: tableToken.Value}

	// Set the TableNode as the From field of the SelectNode
	selectNode.From = tableNode

	// Return the index after parsing the FROM clause
	return startIndex + 1, nil
}

func parseWhereClause(tokens []lexer.Token, startIndex int) (Node, int, error) {
	// Initialize variables to keep track of parsing progress
	var currentNode Node
	i := startIndex

	// Loop through tokens starting from startIndex
	for i < len(tokens) {
		token := tokens[i]

		switch token.Type {
		case "IDENTIFIER":
			// Handle identifier (left operand of the operator)
			leftOperand := &ColumnNode{Name: token.Value}

			// Move to the next token
			i++

			// Expect the next token to be an operator
			if i >= len(tokens) {
				return nil, i, fmt.Errorf("Expected operator, got end of input")
			}
			operatorToken := tokens[i]

			if operatorToken.Type != "OPERATOR" {
				return nil, i, fmt.Errorf("Expected operator, got %s", operatorToken.Value)
			}

			// Handle operator
			operator := &OperatorNode{Operator: operatorToken.Value, Left: leftOperand}

			// Move to the next token
			i++

			// Expect the next token to be a literal
			if i >= len(tokens) {
				return nil, i, fmt.Errorf("Expected literal, got end of input")
			}
			literalToken := tokens[i]

			if literalToken.Type != "LITERAL" {
				return nil, i, fmt.Errorf("Expected literal, got %s", literalToken.Value)
			}

			// Handle literal (right operand of the operator)
			rightOperand := &LiteralNode{Value: literalToken.Value}

			// Set the right operand for the operator node
			operator.Right = rightOperand

			// Set the operator node as the current node
			currentNode = operator

			// Move to the next token
			i++

			// If the next token is AND or OR, continue parsing the next expression
			if i < len(tokens) && (tokens[i].Value == "AND" || tokens[i].Value == "OR") {
				// TODO: Handle AND / OR logic
				// For now, simply move to the next token
				i++
			}

		default:
			// Handle syntax errors or other cases
			return nil, i, fmt.Errorf("Expected identifier, got %s", token.Value)
		}
	}

	// Return the parsed WHERE clause node and the new index
	return currentNode, i, nil
}
