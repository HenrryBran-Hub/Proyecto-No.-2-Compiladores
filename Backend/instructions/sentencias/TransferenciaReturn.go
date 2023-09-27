package sentencias

import (
	"Backend/environment"
	"Backend/generator"
)

type TransferenciaReturn struct {
	Lin int
	Col int
}

func NewTransferenciaReturn(lin int, col int) TransferenciaReturn {
	return TransferenciaReturn{Lin: lin, Col: col}
}

func (v TransferenciaReturn) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	symbol := environment.Symbol{
		Lin:   v.Lin,
		Col:   v.Col,
		Tipo:  environment.BOOLEAN,
		Valor: true,
		Scope: ast.ObtenerAmbito(),
	}
	Variable := environment.Variable{
		Name:        "Return",
		Symbols:     symbol,
		Mutable:     false,
		TipoSimbolo: "Sentencia de Transferencia",
	}

	etiquetas := ast.Lista_Tranferencias.Back().Value.(environment.SentenciasdeTransferencia)
	gen.AddGoto(etiquetas.EFalse)
	ast.GuardarVariable(Variable)
	gen.MainCodeF()
	return nil
}
