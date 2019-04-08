// Code generated from Brainfuck.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Brainfuck

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 11, 28, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 7, 2, 10, 10, 2, 12, 2, 14, 2, 13,
	11, 2, 3, 3, 3, 3, 3, 3, 7, 3, 18, 10, 3, 12, 3, 14, 3, 21, 11, 3, 3, 3,
	5, 3, 24, 10, 3, 3, 4, 3, 4, 3, 4, 2, 2, 5, 2, 4, 6, 2, 3, 3, 2, 3, 8,
	2, 27, 2, 11, 3, 2, 2, 2, 4, 23, 3, 2, 2, 2, 6, 25, 3, 2, 2, 2, 8, 10,
	5, 4, 3, 2, 9, 8, 3, 2, 2, 2, 10, 13, 3, 2, 2, 2, 11, 9, 3, 2, 2, 2, 11,
	12, 3, 2, 2, 2, 12, 3, 3, 2, 2, 2, 13, 11, 3, 2, 2, 2, 14, 24, 5, 6, 4,
	2, 15, 19, 7, 9, 2, 2, 16, 18, 5, 4, 3, 2, 17, 16, 3, 2, 2, 2, 18, 21,
	3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 22, 3, 2, 2, 2,
	21, 19, 3, 2, 2, 2, 22, 24, 7, 10, 2, 2, 23, 14, 3, 2, 2, 2, 23, 15, 3,
	2, 2, 2, 24, 5, 3, 2, 2, 2, 25, 26, 9, 2, 2, 2, 26, 7, 3, 2, 2, 2, 5, 11,
	19, 23,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'>'", "'<'", "'+'", "'-'", "'.'", "','", "'['", "']'",
}
var symbolicNames = []string{
	"", "GT", "LT", "PLUS", "MINUS", "DOT", "COMMA", "LPAREN", "RPAREN", "WS",
}

var ruleNames = []string{
	"file", "statement", "opcode",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type BrainfuckParser struct {
	*antlr.BaseParser
}

func NewBrainfuckParser(input antlr.TokenStream) *BrainfuckParser {
	this := new(BrainfuckParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Brainfuck.g4"

	return this
}

// BrainfuckParser tokens.
const (
	BrainfuckParserEOF    = antlr.TokenEOF
	BrainfuckParserGT     = 1
	BrainfuckParserLT     = 2
	BrainfuckParserPLUS   = 3
	BrainfuckParserMINUS  = 4
	BrainfuckParserDOT    = 5
	BrainfuckParserCOMMA  = 6
	BrainfuckParserLPAREN = 7
	BrainfuckParserRPAREN = 8
	BrainfuckParserWS     = 9
)

// BrainfuckParser rules.
const (
	BrainfuckParserRULE_file      = 0
	BrainfuckParserRULE_statement = 1
	BrainfuckParserRULE_opcode    = 2
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BrainfuckParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BrainfuckParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *FileContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrainfuckListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrainfuckListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *BrainfuckParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BrainfuckParserRULE_file)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(9)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BrainfuckParserGT)|(1<<BrainfuckParserLT)|(1<<BrainfuckParserPLUS)|(1<<BrainfuckParserMINUS)|(1<<BrainfuckParserDOT)|(1<<BrainfuckParserCOMMA)|(1<<BrainfuckParserLPAREN))) != 0 {
		{
			p.SetState(6)
			p.Statement()
		}

		p.SetState(11)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BrainfuckParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BrainfuckParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Opcode() IOpcodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpcodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpcodeContext)
}

func (s *StatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BrainfuckParserLPAREN, 0)
}

func (s *StatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BrainfuckParserRPAREN, 0)
}

func (s *StatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrainfuckListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrainfuckListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *BrainfuckParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BrainfuckParserRULE_statement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(21)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BrainfuckParserGT, BrainfuckParserLT, BrainfuckParserPLUS, BrainfuckParserMINUS, BrainfuckParserDOT, BrainfuckParserCOMMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(12)
			p.Opcode()
		}

	case BrainfuckParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(13)
			p.Match(BrainfuckParserLPAREN)
		}
		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BrainfuckParserGT)|(1<<BrainfuckParserLT)|(1<<BrainfuckParserPLUS)|(1<<BrainfuckParserMINUS)|(1<<BrainfuckParserDOT)|(1<<BrainfuckParserCOMMA)|(1<<BrainfuckParserLPAREN))) != 0 {
			{
				p.SetState(14)
				p.Statement()
			}

			p.SetState(19)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(20)
			p.Match(BrainfuckParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOpcodeContext is an interface to support dynamic dispatch.
type IOpcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpcodeContext differentiates from other interfaces.
	IsOpcodeContext()
}

type OpcodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpcodeContext() *OpcodeContext {
	var p = new(OpcodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BrainfuckParserRULE_opcode
	return p
}

func (*OpcodeContext) IsOpcodeContext() {}

func NewOpcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpcodeContext {
	var p = new(OpcodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BrainfuckParserRULE_opcode

	return p
}

func (s *OpcodeContext) GetParser() antlr.Parser { return s.parser }

func (s *OpcodeContext) GT() antlr.TerminalNode {
	return s.GetToken(BrainfuckParserGT, 0)
}

func (s *OpcodeContext) LT() antlr.TerminalNode {
	return s.GetToken(BrainfuckParserLT, 0)
}

func (s *OpcodeContext) PLUS() antlr.TerminalNode {
	return s.GetToken(BrainfuckParserPLUS, 0)
}

func (s *OpcodeContext) MINUS() antlr.TerminalNode {
	return s.GetToken(BrainfuckParserMINUS, 0)
}

func (s *OpcodeContext) DOT() antlr.TerminalNode {
	return s.GetToken(BrainfuckParserDOT, 0)
}

func (s *OpcodeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(BrainfuckParserCOMMA, 0)
}

func (s *OpcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrainfuckListener); ok {
		listenerT.EnterOpcode(s)
	}
}

func (s *OpcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrainfuckListener); ok {
		listenerT.ExitOpcode(s)
	}
}

func (p *BrainfuckParser) Opcode() (localctx IOpcodeContext) {
	localctx = NewOpcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BrainfuckParserRULE_opcode)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BrainfuckParserGT)|(1<<BrainfuckParserLT)|(1<<BrainfuckParserPLUS)|(1<<BrainfuckParserMINUS)|(1<<BrainfuckParserDOT)|(1<<BrainfuckParserCOMMA))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
