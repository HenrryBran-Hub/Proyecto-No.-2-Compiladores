package sentencias

import (
	"Backend/environment"
	"Backend/generator"
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

func (v AsignacionVariable) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	if ast.ObtenerAmbito() == "Global" {
		gen.MainCodeT()
	}
	value := v.Value.Ejecutar(ast, gen)

	Variable := ast.GetVariable(v.Name)
	if Variable != nil && Variable.Mutable && Variable.Symbols.Tipo == value.Type {
		Variable.Symbols.Lin = v.Lin
		Variable.Symbols.Col = v.Col

		Variable.Symbols.Scope = ast.ObtenerAmbito()

		gen.AddComment("Asignacion de Variable")

		if value.Type == environment.BOOLEAN {
			newlabel := gen.NewLabel()
			gen.AddLabel(value.Val.TEti)
			gen.AddSetStack(strconv.Itoa(Variable.Symbols.Posicion), "1")
			Variable.Symbols.Valor = "1"
			gen.AddGoto(newlabel)
			gen.AddLabel(value.Val.FEti)
			gen.AddSetStack(strconv.Itoa(Variable.Symbols.Posicion), "0")
			Variable.Symbols.Valor = "0"
			gen.AddGoto(newlabel)
			gen.AddLabel(newlabel)
			gen.AddBr()
			Variable.FEti = value.Val.FEti
			Variable.TEti = value.Val.TEti
		} else {
			//si es temp (num,string,etc)
			gen.AddSetStack(strconv.Itoa(Variable.Symbols.Posicion), value.Value)
			gen.AddBr()
			Variable.Symbols.Valor = value.Value
		}
		ast.ActualizarVariable(Variable)
	}

	if !Variable.Mutable {
		Errores := environment.Errores{
			Descripcion: "Se ha querido asignar un valor a una constante.",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		gen.MainCodeF()
		return nil
	}

	if Variable.Symbols.Tipo != value.Type {
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
		switch value.Type {
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
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		Variable.Symbols.Valor = nil
		gen.MainCodeF()
		return nil
	}

	gen.MainCodeF()
	return nil
}
