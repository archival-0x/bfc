// Code generated from Brainfuck.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Brainfuck

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBrainfuckListener is a complete listener for a parse tree produced by BrainfuckParser.
type BaseBrainfuckListener struct{}

var _ BrainfuckListener = &BaseBrainfuckListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBrainfuckListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBrainfuckListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBrainfuckListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBrainfuckListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BaseBrainfuckListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseBrainfuckListener) ExitFile(ctx *FileContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseBrainfuckListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseBrainfuckListener) ExitStatement(ctx *StatementContext) {}

// EnterOpcode is called when production opcode is entered.
func (s *BaseBrainfuckListener) EnterOpcode(ctx *OpcodeContext) {}

// ExitOpcode is called when production opcode is exited.
func (s *BaseBrainfuckListener) ExitOpcode(ctx *OpcodeContext) {}
