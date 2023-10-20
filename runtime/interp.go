package runtime

import (
	"log"

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
	l := antlr.NewDiagnosticErrorListener(true)
	p.AddErrorListener(l)

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
	log.Print("EnterService")
}

// EnterMatcher ...
func (i *interp) EnterMatcher(ctx *parser.MatcherContext) {
	log.Print("EnterMatcher")
}

// ExitMatcher ...
func (i *interp) ExitMatcher(ctx *parser.MatcherContext) {
	log.Print("ExitMatcher")
}

// EnterAllow ...
func (i *interp) EnterAllow(ctx *parser.AllowContext) {
	log.Print("EnterAllow")
}

// ExitAllow ...
func (i *interp) ExitAllow(ctx *parser.AllowContext) {
	log.Print("ExitAllow")
}

// EnterAllowKey ...
func (i *interp) EnterAllowKey(ctx *parser.AllowKeyContext) {
	log.Print("EnterAllowKey")
}

// ExitAllowKey ...
func (i *interp) ExitAllowKey(ctx *parser.AllowKeyContext) {
	log.Print("ExitAllowKey")
}

// HasAccess ...
func (i *interp) HasAccess(ctx Context, path string) bool {
	return true
}
