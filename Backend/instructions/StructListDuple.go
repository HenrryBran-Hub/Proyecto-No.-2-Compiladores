package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
)

type StructListDuple struct {
	Name   string
	Expr   interfaces.Expression
	Op     interfaces.Instruction
	Estado bool
}

func NewStructListDuple(name string, expr interfaces.Expression, op interfaces.Instruction, estado bool) StructListDuple {
	return StructListDuple{name, expr, op, estado}
}

func NewStructDuple(name string, expr interfaces.Expression, estado bool) StructListDuple {
	return StructListDuple{Name: name, Expr: expr, Estado: estado}
}

func (v StructListDuple) Ejecutar(ast *environment.AST) interface{} {
	valor := v.Expr.Ejecutar(ast)
	symbol := environment.Symbol{
		Lin:   valor.Lin,
		Col:   valor.Col,
		Tipo:  valor.Tipo,
		Valor: valor.Valor,
		Scope: ast.ObtenerAmbito(),
	}
	parametros := environment.Parametrostruct{
		Name:    v.Name,
		Symbolo: symbol,
	}
	ast.ListaParametrosStruct.PushBack(parametros)
	if v.Estado == true {
		v.Op.Ejecutar(ast)
	}
	return nil
}
