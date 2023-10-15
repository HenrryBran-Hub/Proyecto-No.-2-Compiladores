package sentencias

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
)

type TransferenciaReturnExp struct {
	Lin   int
	Col   int
	Value interfaces.Expression
}

func NewTransferenciaReturnExp(lin int, col int, value interfaces.Expression) TransferenciaReturnExp {
	return TransferenciaReturnExp{Lin: lin, Col: col, Value: value}
}

func (v TransferenciaReturnExp) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	value := v.Value.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	symbol := environment.Symbol{
		Lin:         v.Lin,
		Col:         v.Col,
		Tipo:        value.Type,
		Valor:       value.Value,
		ValorInt:    value.Val.Symbols.ValorInt,
		ValorFloat:  value.Val.Symbols.ValorFloat,
		ValorString: value.Val.Symbols.ValorString,
		Scope:       ast.ObtenerAmbito(),
		Posicion:    ast.PosicionStack,
	}
	Variable := environment.Variable{
		Name:        "ReturnExp",
		Symbols:     symbol,
		Mutable:     false,
		TipoSimbolo: "Sentencia de Transferencia",
	}

	ast.GuardarVariable(Variable)
	gen.AddComment("Retorno de variable")

	if value.Type == environment.BOOLEAN {
		gen.AddSetStack(strconv.Itoa(symbol.Posicion), value.Value)
		gen.AddBr()
	} else {
		//si es temp (num,string,etc)
		gen.AddSetStack(strconv.Itoa(symbol.Posicion), value.Value)
		gen.AddBr()
	}
	etiquetas := ast.Lista_Tranferencias.Back().Value.(environment.SentenciasdeTransferencia)
	newtemp := gen.NewTemp()
	gen.AddGetStack(newtemp, strconv.Itoa(symbol.Posicion))
	gen.AddSetStack("(int)P", newtemp)
	gen.AddGoto(etiquetas.EFalse)

	etiqueta := ast.Lista_Tranferencias.Front().Value.(environment.SentenciasdeTransferencia)
	if etiqueta.Tipo != Variable.Symbols.Tipo {
		Errores := environment.Errores{
			Descripcion: "El tipo que esta retornando no es del mismo tipo de la funcion",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
	}

	gen.MainCodeF()
	return nil
}
