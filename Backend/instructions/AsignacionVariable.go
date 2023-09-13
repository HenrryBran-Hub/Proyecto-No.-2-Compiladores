package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
	"strconv"
)

type AsignacionVariable struct {
	Lin   int
	Col   int
	Name  string
	Type  environment.TipoExpresion
	Value interfaces.Expression
}

func NewAsignacionVariable(lin int, col int, name string, value interfaces.Expression) AsignacionVariable {
	return AsignacionVariable{Lin: lin, Col: col, Name: name, Value: value}
}

func (v AsignacionVariable) Ejecutar(ast *environment.AST) interface{} {
	value := v.Value.Ejecutar(ast)
	Variable := ast.GetVariable(v.Name)
	if Variable != nil && Variable.Mutable && Variable.Symbols.Tipo == value.Tipo {
		Variable.Symbols.Lin = v.Lin
		Variable.Symbols.Col = v.Col
		Variable.Symbols.Valor = value.Valor
		Variable.Symbols.Scope = ast.ObtenerAmbito()
		ast.ActualizarVariable(Variable)
	}

	if Variable.Mutable == false {
		Errores := environment.Errores{
			Descripcion: "Se ha querido asignar un valor a una constante.",
			Fila:        strconv.Itoa(value.Lin),
			Columna:     strconv.Itoa(value.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	if Variable.Symbols.Tipo != value.Tipo {
		var tipoexpstr string
		switch Variable.Symbols.Tipo {
		case 0:
			tipoexpstr = "Int"
		case 1:
			tipoexpstr = "Float"
		case 2:
			tipoexpstr = "String"
		case 3:
			tipoexpstr = "Boolean"
		case 4:
			tipoexpstr = "Character"
		default:
			tipoexpstr = "nil"
		}

		var tipoexpstr2 string
		switch value.Tipo {
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

		Errores := environment.Errores{
			Descripcion: "Se ha querido asignar un valor no correspondiente a el tipo de dato: \nTipo de dato:" + tipoexpstr2 + "\nTipo de Valor:" + tipoexpstr + ".",
			Fila:        strconv.Itoa(value.Lin),
			Columna:     strconv.Itoa(value.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		Variable.Symbols.Valor = nil
		return nil
	}

	return nil
}
