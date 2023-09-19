package runtime

import (
	parser "github.com/katallaxie/acl/parser"

	"github.com/antlr4-go/antlr/v4"
)

type interp struct {
	parser *parser.AclRulesParser
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

	// service := p.Service()
	// antlr.ParseTreeWalkerDefault.Walk(p, service)

	return in, nil
}

// NewInterp ...
func NewInterp() *interp {
	return &interp{}
}

// HasAccess ...
func (i *interp) HasAccess(ctx Context, path string) bool {
	return true
}
