package datosprimitivos

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
)

type ConstanteDeclaracionSinTipo struct {
	Lin   int
	Col   int
	Name  string
	Type  environment.TipoExpresion
	Value interfaces.Expression
}

func NewConstanteDeclaracionSinTipo(lin int, col int, name string, value interfaces.Expression) ConstanteDeclaracionSinTipo {
	return ConstanteDeclaracionSinTipo{Lin: lin, Col: col, Name: name, Value: value}
}

func (v ConstanteDeclaracionSinTipo) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	if ast.ObtenerAmbito() == "Global" {
		gen.MainCodeT()
	}
	value := v.Value.Ejecutar(ast, gen)
	symbol := environment.Symbol{
		Lin:      v.Lin,
		Col:      v.Col,
		Tipo:     v.Type,
		Scope:    ast.ObtenerAmbito(),
		TipoDato: environment.VARIABLE,
		Posicion: ast.Lista_Variables.Len(),
		Valor:    value.Value,
	}
	Variable := environment.Variable{
		Name:        v.Name,
		Symbols:     symbol,
		Mutable:     false,
		TipoSimbolo: "Constante",
	}

	gen.AddComment("Declaracion de Constante")

	if value.Type == environment.BOOLEAN {
		//si no es temp (boolean)
		//add labels
		for e := value.TrueLabel.Front(); e != nil; e = e.Next() {
			gen.AddLabel(e.Value.(string))
		}
		gen.AddSetStack(strconv.Itoa(symbol.Posicion), "1")
		newLabel := gen.NewLabel()
		gen.AddGoto(newLabel)
		Variable.TEti = newLabel
		//add labels
		for e := value.FalseLabel.Front(); e != nil; e = e.Next() {
			gen.AddLabel(e.Value.(string))
		}
		gen.AddSetStack(strconv.Itoa(symbol.Posicion), "0")
		newLabel2 := gen.NewLabel()
		gen.AddGoto(newLabel2)
		Variable.FEti = newLabel2
		gen.AddBr()
	} else {
		//si es temp (num,string,etc)
		gen.AddSetStack(strconv.Itoa(symbol.Posicion), value.Value)
		gen.AddBr()
	}

	ast.GuardarVariable(Variable)
	gen.MainCodeF()
	return value
}
