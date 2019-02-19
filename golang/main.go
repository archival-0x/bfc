package main

import (
    "fmt"
    "io/ioutil"
    "os"

    "github.com/antlr/antlr4/runtime/Go/antlr"
    "./parser"
)

type BFListener struct {
    *parser.BasebrainfuckListener
    tape []int
}


type BFListener interface {
    antlr ParseTreeListener

    // TODO interface logic
}


func main() {

    // read file from arg
    fileName := os.Args[1]
    code, err := ioutil.ReadFile(fileName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %s\n", err)
        os.Exit(-1)
    }

    // initialize input with ANTLR
    in_stream = antlr.NewInputStream(code)

    // create lexer
    lexer := parser.NewbrainfuckLexer(in_stream)
    stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

    // create parser
    parser := parser.NewbrainfuckParser(stream)
}
