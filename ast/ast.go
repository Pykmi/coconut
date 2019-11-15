package ast

// Identifier
func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// LetStatement
func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// ReturnStatement
func (ls *ReturnStatement) statementNode() {}

func (ls *ReturnStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Program
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}
