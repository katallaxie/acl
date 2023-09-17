// Code generated from AclRules.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // AclRules
import "github.com/antlr4-go/antlr/v4"

// BaseAclRulesListener is a complete listener for a parse tree produced by AclRulesParser.
type BaseAclRulesListener struct{}

var _ AclRulesListener = &BaseAclRulesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAclRulesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAclRulesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAclRulesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAclRulesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterService is called when production service is entered.
func (s *BaseAclRulesListener) EnterService(ctx *ServiceContext) {}

// ExitService is called when production service is exited.
func (s *BaseAclRulesListener) ExitService(ctx *ServiceContext) {}

// EnterNamespaceIdentifier is called when production namespaceIdentifier is entered.
func (s *BaseAclRulesListener) EnterNamespaceIdentifier(ctx *NamespaceIdentifierContext) {}

// ExitNamespaceIdentifier is called when production namespaceIdentifier is exited.
func (s *BaseAclRulesListener) ExitNamespaceIdentifier(ctx *NamespaceIdentifierContext) {}

// EnterNamespaceBlock is called when production namespaceBlock is entered.
func (s *BaseAclRulesListener) EnterNamespaceBlock(ctx *NamespaceBlockContext) {}

// ExitNamespaceBlock is called when production namespaceBlock is exited.
func (s *BaseAclRulesListener) ExitNamespaceBlock(ctx *NamespaceBlockContext) {}

// EnterMatchBlock is called when production matchBlock is entered.
func (s *BaseAclRulesListener) EnterMatchBlock(ctx *MatchBlockContext) {}

// ExitMatchBlock is called when production matchBlock is exited.
func (s *BaseAclRulesListener) ExitMatchBlock(ctx *MatchBlockContext) {}

// EnterAllowKey is called when production allowKey is entered.
func (s *BaseAclRulesListener) EnterAllowKey(ctx *AllowKeyContext) {}

// ExitAllowKey is called when production allowKey is exited.
func (s *BaseAclRulesListener) ExitAllowKey(ctx *AllowKeyContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseAclRulesListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseAclRulesListener) ExitComment(ctx *CommentContext) {}

// EnterMatcher is called when production matcher is entered.
func (s *BaseAclRulesListener) EnterMatcher(ctx *MatcherContext) {}

// ExitMatcher is called when production matcher is exited.
func (s *BaseAclRulesListener) ExitMatcher(ctx *MatcherContext) {}

// EnterAllow is called when production allow is entered.
func (s *BaseAclRulesListener) EnterAllow(ctx *AllowContext) {}

// ExitAllow is called when production allow is exited.
func (s *BaseAclRulesListener) ExitAllow(ctx *AllowContext) {}

// EnterGetPathVariable is called when production getPathVariable is entered.
func (s *BaseAclRulesListener) EnterGetPathVariable(ctx *GetPathVariableContext) {}

// ExitGetPathVariable is called when production getPathVariable is exited.
func (s *BaseAclRulesListener) ExitGetPathVariable(ctx *GetPathVariableContext) {}

// EnterPathVariable is called when production pathVariable is entered.
func (s *BaseAclRulesListener) EnterPathVariable(ctx *PathVariableContext) {}

// ExitPathVariable is called when production pathVariable is exited.
func (s *BaseAclRulesListener) ExitPathVariable(ctx *PathVariableContext) {}

// EnterArg is called when production arg is entered.
func (s *BaseAclRulesListener) EnterArg(ctx *ArgContext) {}

// ExitArg is called when production arg is exited.
func (s *BaseAclRulesListener) ExitArg(ctx *ArgContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseAclRulesListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseAclRulesListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterMemberArg is called when production memberArg is entered.
func (s *BaseAclRulesListener) EnterMemberArg(ctx *MemberArgContext) {}

// ExitMemberArg is called when production memberArg is exited.
func (s *BaseAclRulesListener) ExitMemberArg(ctx *MemberArgContext) {}

// EnterMemberArguments is called when production memberArguments is entered.
func (s *BaseAclRulesListener) EnterMemberArguments(ctx *MemberArgumentsContext) {}

// ExitMemberArguments is called when production memberArguments is exited.
func (s *BaseAclRulesListener) ExitMemberArguments(ctx *MemberArgumentsContext) {}

// EnterArgDeclaration is called when production argDeclaration is entered.
func (s *BaseAclRulesListener) EnterArgDeclaration(ctx *ArgDeclarationContext) {}

// ExitArgDeclaration is called when production argDeclaration is exited.
func (s *BaseAclRulesListener) ExitArgDeclaration(ctx *ArgDeclarationContext) {}

// EnterArgDeclarations is called when production argDeclarations is entered.
func (s *BaseAclRulesListener) EnterArgDeclarations(ctx *ArgDeclarationsContext) {}

// ExitArgDeclarations is called when production argDeclarations is exited.
func (s *BaseAclRulesListener) ExitArgDeclarations(ctx *ArgDeclarationsContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseAclRulesListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseAclRulesListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFieldReferenceWithIdentifier is called when production fieldReferenceWithIdentifier is entered.
func (s *BaseAclRulesListener) EnterFieldReferenceWithIdentifier(ctx *FieldReferenceWithIdentifierContext) {
}

// ExitFieldReferenceWithIdentifier is called when production fieldReferenceWithIdentifier is exited.
func (s *BaseAclRulesListener) ExitFieldReferenceWithIdentifier(ctx *FieldReferenceWithIdentifierContext) {
}

// EnterFieldReferenceWithMemberRef is called when production fieldReferenceWithMemberRef is entered.
func (s *BaseAclRulesListener) EnterFieldReferenceWithMemberRef(ctx *FieldReferenceWithMemberRefContext) {
}

// ExitFieldReferenceWithMemberRef is called when production fieldReferenceWithMemberRef is exited.
func (s *BaseAclRulesListener) ExitFieldReferenceWithMemberRef(ctx *FieldReferenceWithMemberRefContext) {
}

// EnterId is called when production id is entered.
func (s *BaseAclRulesListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseAclRulesListener) ExitId(ctx *IdContext) {}

// EnterArrayExpression is called when production arrayExpression is entered.
func (s *BaseAclRulesListener) EnterArrayExpression(ctx *ArrayExpressionContext) {}

// ExitArrayExpression is called when production arrayExpression is exited.
func (s *BaseAclRulesListener) ExitArrayExpression(ctx *ArrayExpressionContext) {}

// EnterNumberExpression is called when production numberExpression is entered.
func (s *BaseAclRulesListener) EnterNumberExpression(ctx *NumberExpressionContext) {}

// ExitNumberExpression is called when production numberExpression is exited.
func (s *BaseAclRulesListener) ExitNumberExpression(ctx *NumberExpressionContext) {}

// EnterObjectReferenceExpression is called when production objectReferenceExpression is entered.
func (s *BaseAclRulesListener) EnterObjectReferenceExpression(ctx *ObjectReferenceExpressionContext) {
}

// ExitObjectReferenceExpression is called when production objectReferenceExpression is exited.
func (s *BaseAclRulesListener) ExitObjectReferenceExpression(ctx *ObjectReferenceExpressionContext) {}

// EnterParenthesisExpression is called when production parenthesisExpression is entered.
func (s *BaseAclRulesListener) EnterParenthesisExpression(ctx *ParenthesisExpressionContext) {}

// ExitParenthesisExpression is called when production parenthesisExpression is exited.
func (s *BaseAclRulesListener) ExitParenthesisExpression(ctx *ParenthesisExpressionContext) {}

// EnterArithmeticExpression is called when production arithmeticExpression is entered.
func (s *BaseAclRulesListener) EnterArithmeticExpression(ctx *ArithmeticExpressionContext) {}

// ExitArithmeticExpression is called when production arithmeticExpression is exited.
func (s *BaseAclRulesListener) ExitArithmeticExpression(ctx *ArithmeticExpressionContext) {}

// EnterMemberReferenceExpression is called when production memberReferenceExpression is entered.
func (s *BaseAclRulesListener) EnterMemberReferenceExpression(ctx *MemberReferenceExpressionContext) {
}

// ExitMemberReferenceExpression is called when production memberReferenceExpression is exited.
func (s *BaseAclRulesListener) ExitMemberReferenceExpression(ctx *MemberReferenceExpressionContext) {}

// EnterBooleanExpression is called when production booleanExpression is entered.
func (s *BaseAclRulesListener) EnterBooleanExpression(ctx *BooleanExpressionContext) {}

// ExitBooleanExpression is called when production booleanExpression is exited.
func (s *BaseAclRulesListener) ExitBooleanExpression(ctx *BooleanExpressionContext) {}

// EnterFunctionExpression is called when production functionExpression is entered.
func (s *BaseAclRulesListener) EnterFunctionExpression(ctx *FunctionExpressionContext) {}

// ExitFunctionExpression is called when production functionExpression is exited.
func (s *BaseAclRulesListener) ExitFunctionExpression(ctx *FunctionExpressionContext) {}

// EnterCompareExpression is called when production compareExpression is entered.
func (s *BaseAclRulesListener) EnterCompareExpression(ctx *CompareExpressionContext) {}

// ExitCompareExpression is called when production compareExpression is exited.
func (s *BaseAclRulesListener) ExitCompareExpression(ctx *CompareExpressionContext) {}

// EnterBinaryExpression is called when production binaryExpression is entered.
func (s *BaseAclRulesListener) EnterBinaryExpression(ctx *BinaryExpressionContext) {}

// ExitBinaryExpression is called when production binaryExpression is exited.
func (s *BaseAclRulesListener) ExitBinaryExpression(ctx *BinaryExpressionContext) {}

// EnterLogicalExpression is called when production LogicalExpression is entered.
func (s *BaseAclRulesListener) EnterLogicalExpression(ctx *LogicalExpressionContext) {}

// ExitLogicalExpression is called when production LogicalExpression is exited.
func (s *BaseAclRulesListener) ExitLogicalExpression(ctx *LogicalExpressionContext) {}

// EnterGetExpression is called when production getExpression is entered.
func (s *BaseAclRulesListener) EnterGetExpression(ctx *GetExpressionContext) {}

// ExitGetExpression is called when production getExpression is exited.
func (s *BaseAclRulesListener) ExitGetExpression(ctx *GetExpressionContext) {}

// EnterInExpression is called when production inExpression is entered.
func (s *BaseAclRulesListener) EnterInExpression(ctx *InExpressionContext) {}

// ExitInExpression is called when production inExpression is exited.
func (s *BaseAclRulesListener) ExitInExpression(ctx *InExpressionContext) {}

// EnterStringExpression is called when production stringExpression is entered.
func (s *BaseAclRulesListener) EnterStringExpression(ctx *StringExpressionContext) {}

// ExitStringExpression is called when production stringExpression is exited.
func (s *BaseAclRulesListener) ExitStringExpression(ctx *StringExpressionContext) {}

// EnterNullExpression is called when production nullExpression is entered.
func (s *BaseAclRulesListener) EnterNullExpression(ctx *NullExpressionContext) {}

// ExitNullExpression is called when production nullExpression is exited.
func (s *BaseAclRulesListener) ExitNullExpression(ctx *NullExpressionContext) {}

// EnterRangeExpression is called when production rangeExpression is entered.
func (s *BaseAclRulesListener) EnterRangeExpression(ctx *RangeExpressionContext) {}

// ExitRangeExpression is called when production rangeExpression is exited.
func (s *BaseAclRulesListener) ExitRangeExpression(ctx *RangeExpressionContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseAclRulesListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseAclRulesListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterMemberFunctionExpression is called when production memberFunctionExpression is entered.
func (s *BaseAclRulesListener) EnterMemberFunctionExpression(ctx *MemberFunctionExpressionContext) {}

// ExitMemberFunctionExpression is called when production memberFunctionExpression is exited.
func (s *BaseAclRulesListener) ExitMemberFunctionExpression(ctx *MemberFunctionExpressionContext) {}

// EnterObjectReference is called when production objectReference is entered.
func (s *BaseAclRulesListener) EnterObjectReference(ctx *ObjectReferenceContext) {}

// ExitObjectReference is called when production objectReference is exited.
func (s *BaseAclRulesListener) ExitObjectReference(ctx *ObjectReferenceContext) {}

// EnterGetPathExpressionVariable is called when production getPathExpressionVariable is entered.
func (s *BaseAclRulesListener) EnterGetPathExpressionVariable(ctx *GetPathExpressionVariableContext) {
}

// ExitGetPathExpressionVariable is called when production getPathExpressionVariable is exited.
func (s *BaseAclRulesListener) ExitGetPathExpressionVariable(ctx *GetPathExpressionVariableContext) {}

// EnterGetPath is called when production getPath is entered.
func (s *BaseAclRulesListener) EnterGetPath(ctx *GetPathContext) {}

// ExitGetPath is called when production getPath is exited.
func (s *BaseAclRulesListener) ExitGetPath(ctx *GetPathContext) {}

// EnterRuleFunctionCall is called when production ruleFunctionCall is entered.
func (s *BaseAclRulesListener) EnterRuleFunctionCall(ctx *RuleFunctionCallContext) {}

// ExitRuleFunctionCall is called when production ruleFunctionCall is exited.
func (s *BaseAclRulesListener) ExitRuleFunctionCall(ctx *RuleFunctionCallContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseAclRulesListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseAclRulesListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterMatchPath is called when production matchPath is entered.
func (s *BaseAclRulesListener) EnterMatchPath(ctx *MatchPathContext) {}

// ExitMatchPath is called when production matchPath is exited.
func (s *BaseAclRulesListener) ExitMatchPath(ctx *MatchPathContext) {}
