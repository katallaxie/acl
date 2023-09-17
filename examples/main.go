package main

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/katallaxie/acl/parser"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewAclRulesLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewAclRulesParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	service := p.Service()
	fmt.Println(service.GetText())
}
