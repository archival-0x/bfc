package main

import (
    "fmt"
    "os"
    "io/ioutil"

    "./parser"
    "github.com/antlr/antlr4/runtime/Go/antlr"
)

type BFListener struct {
    *parser.BaseBrainfuckListener
    tape []int
}

/*
type BFListener interface {
    antlr.ParseTreeListener
}
*/

func main() {

    // read file from arg
    fileName := os.Args[1]
    code, err := ioutil.ReadFile(fileName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %s\n", err)
        os.Exit(-1)
    }

    // initialize input with ANTLR
    in_stream := antlr.NewInputStream(string(code))

    // create lexer
    lexer := parser.NewBrainfuckLexer(in_stream)
    stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

    // create parser
    parser := parser.NewBrainfuckParser(stream)

    // parse input expression
    antlr.ParseTreeWalkerDefault.Walk(&BFListener{}, parser.Start())
}
