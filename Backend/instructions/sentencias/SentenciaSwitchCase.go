package sentencias

import (
	"Backend/interfaces"
)

type SentenciaSwitchCase struct {
	Lin  int
	Col  int
	Exp1 interfaces.Expression
	Case []interface{}
}

func NewSentenciaSwitchCase(lin int, col int, expr1 interfaces.Expression, operations []interface{}) SentenciaSwitchCase {
	return SentenciaSwitchCase{lin, col, expr1, operations}
}

/*
func (v SentenciaSwitchCase) Ejecutar(ast *environment.AST) interface{} {
	return nil
}
*/
