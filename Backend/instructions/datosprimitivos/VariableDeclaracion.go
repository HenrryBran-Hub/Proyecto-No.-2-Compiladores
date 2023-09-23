package datosprimitivos

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
)

type VariableDeclaracion struct {
	Lin   int
	Col   int
	Name  string
	Type  environment.TipoExpresion
	Value interfaces.Expression
}

func NewVariableDeclaration(lin int, col int, name string, tipo environment.TipoExpresion, value interfaces.Expression) VariableDeclaracion {
	return VariableDeclaracion{lin, col, name, tipo, value}
}

func (v VariableDeclaracion) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
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
		Mutable:     true,
		TipoSimbolo: "Variable",
	}

	var tipoexp int = -1
	var tipoexpstr string
	switch value.Type {
	case 0:
		tipoexp = 0
		tipoexpstr = "Int"
	case 1:
		tipoexp = 1
		tipoexpstr = "Float"
	case 2:
		tipoexp = 2
		tipoexpstr = "String"
	case 3:
		tipoexp = 3
		tipoexpstr = "Boolean"
	case 4:
		tipoexp = 4
		tipoexpstr = "Character"
	default:
		tipoexp = 5
		tipoexpstr = "nil"
	}

	var tipoexpstr2 string
	switch v.Type {
	case 0:
		tipoexpstr2 = "Int"
	case 1:
		tipoexpstr2 = "Float"
	case 2:
		tipoexpstr2 = "String"
	case 3:
		tipoexpstr2 = "Boolean"
	case 4:
		tipoexpstr2 = "Character"
	default:
		tipoexpstr2 = "nil"
	}

	if v.Type == 0 && value.Type == 0 && v.Type != environment.BOOLEAN {
		Variable.Symbols.Tipo = environment.INTEGER
		tipoexp = 0
	}

	if v.Type == 1 && tipoexp == 0 {
		Variable.Symbols.Tipo = environment.FLOAT
		tipoexp = 1
	}

	if v.Type == 4 && value.Type == 4 {
		Variable.Symbols.Tipo = environment.CHARACTER
		tipoexp = 4
	}

	if tipoexp != int(Variable.Symbols.Tipo) {
		Errores := environment.Errores{
			Descripcion: "Se ha querido asignar un valor no correspondiente a el tipo de dato: \nTipo de dato:" + tipoexpstr2 + "\nTipo de Valor:" + tipoexpstr + ".",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		Variable.Symbols.Valor = nil
		Variable.Symbols.Tipo = environment.INTEGER
	}

	gen.AddComment("Declaracion de Variable")

	if value.Type == environment.BOOLEAN {
		//si no es temp (boolean)
		newLabel := gen.NewLabel()
		Variable.TEti = newLabel
		newLabel2 := gen.NewLabel()
		Variable.FEti = newLabel2
		if value.Value == "1" {
			gen.AddGoto(value.Val.TEti)
		} else {
			gen.AddGoto(value.Val.FEti)
		}

		for e := value.TrueLabel.Front(); e != nil; e = e.Next() {
			gen.AddLabel(e.Value.(string))
		}
		gen.AddSetStack(strconv.Itoa(ast.Lista_Variables.Len()), "1")
		gen.AddGoto(newLabel)

		for e := value.FalseLabel.Front(); e != nil; e = e.Next() {
			gen.AddLabel(e.Value.(string))
		}
		gen.AddSetStack(strconv.Itoa(ast.Lista_Variables.Len()), "0")
		gen.AddGoto(newLabel2)
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