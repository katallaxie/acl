package runtime

import (
	"fmt"

	parser "github.com/katallaxie/acl/parser"

	"github.com/antlr4-go/antlr/v4"
)

type interp struct {
	parser *parser.AclRulesParser
	parser.BaseAclRulesListener
}

// FromString ...
func FromString(s string) (*interp, error) {
	input := antlr.NewInputStream(s)
	lexer := parser.NewAclRulesLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewAclRulesParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	in := new(interp)
	in.parser = p

	service := p.Service()
	antlr.ParseTreeWalkerDefault.Walk(in, service)

	return in, nil
}

// NewInterp ...
func NewInterp() *interp {
	return &interp{}
}

// EnterService ...
func (i *interp) EnterService(ctx *parser.ServiceContext) {
	fmt.Println("EnterService")
}

// HasAccess ...
func (i *interp) HasAccess(ctx Context, path string) bool {
	return true
}
