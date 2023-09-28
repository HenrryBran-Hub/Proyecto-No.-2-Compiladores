package sentencias

import (
	"Backend/environment"
	"Backend/generator"
	"strconv"
)

type TransferenciaContinue struct {
	Lin int
	Col int
}

func NewTransferenciaContinue(lin int, col int) TransferenciaContinue {
	return TransferenciaContinue{Lin: lin, Col: col}
}

func (v TransferenciaContinue) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
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
		Name:        "Continue",
		Symbols:     symbol,
		Mutable:     false,
		TipoSimbolo: "Sentencia de Transferencia",
	}

	if ast.Lista_For_Rango.Len() != 0 {
		variableaux := ast.Lista_For_Rango.Back().Value.(environment.Variable)
		newTemp1 := gen.NewTemp()
		gen.AddGetStack(newTemp1, strconv.Itoa(variableaux.Symbols.Posicion))
		newTemp2 := gen.NewTemp()
		gen.AddExpression(newTemp2, newTemp1, "1", "+")
		gen.AddSetStack(strconv.Itoa(variableaux.Symbols.Posicion), newTemp2)
	}

	etiquetas := ast.Lista_Tranferencias.Back().Value.(environment.SentenciasdeTransferencia)
	gen.AddGoto(etiquetas.ETrue)
	ast.GuardarVariable(Variable)
	gen.MainCodeF()
	return nil
}
