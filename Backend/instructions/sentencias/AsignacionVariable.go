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
		Variable.Symbols.Valor = value.Value
		Variable.Symbols.Scope = ast.ObtenerAmbito()
		ast.ActualizarVariable(Variable)

		gen.AddComment("Asignacion de Variable")

		if value.Type == environment.BOOLEAN {
			//si no es temp (boolean)
			newLabel := gen.NewLabel()
			//add labels
			for e := value.TrueLabel.Front(); e != nil; e = e.Next() {
				gen.AddLabel(e.Value.(string))
			}
			gen.AddSetStack(strconv.Itoa(Variable.Symbols.Posicion), "1")
			gen.AddGoto(newLabel)
			//add labels
			for e := value.FalseLabel.Front(); e != nil; e = e.Next() {
				gen.AddLabel(e.Value.(string))
			}
			gen.AddSetStack(strconv.Itoa(Variable.Symbols.Posicion), "0")
			gen.AddGoto(newLabel)
			gen.AddLabel(newLabel)
			gen.AddBr()
		} else {
			//si es temp (num,string,etc)
			gen.AddSetStack(strconv.Itoa(Variable.Symbols.Posicion), value.Value)
			gen.AddBr()
		}

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
