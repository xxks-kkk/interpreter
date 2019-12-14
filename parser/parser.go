package parser

import "github.com/xxks-kkk/interpreter/token"
import "github.com/xxks-kkk/interpreter/lexer"
import "github.com/xxks-kkk/interpreter/ast"

type Parser struct {
     l *lexer.Lexer
     curToken token.Token
     peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
     p := &Parser{l: l}

     // Read two tokens, so curToken and peekToken are both set
     p.nextToken()
     p.nextToken()

     return p
}

func (p *Parser) nextToken() {
     p.curToken = p.peekToken
     p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
     return nil
}

