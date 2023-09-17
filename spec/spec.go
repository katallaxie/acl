package spec

import (
	parsing "github.com/katallaxie/acl/parser"

	"github.com/antlr4-go/antlr/v4"
)

// New ...
func New(input antlr.TokenStream) *parsing.AclRulesParser {
	p := parsing.NewAclRulesParser(input)

	return p
}
