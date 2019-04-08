// Code generated from Brainfuck.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Brainfuck

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BrainfuckListener is a complete listener for a parse tree produced by BrainfuckParser.
type BrainfuckListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterOpcode is called when entering the opcode production.
	EnterOpcode(c *OpcodeContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitOpcode is called when exiting the opcode production.
	ExitOpcode(c *OpcodeContext)
}
