package sentencias

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
)

type AsignacionSuma struct {
	Lin   int
	Col   int
	Name  string
	Type  environment.TipoExpresion
	Value interfaces.Expression
}

func NewAsignacionSuma(lin int, col int, name string, value interfaces.Expression) AsignacionSuma {
	return AsignacionSuma{Lin: lin, Col: col, Name: name, Value: value}
}

func (v AsignacionSuma) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	if ast.ObtenerAmbito() == "Global" {
		gen.MainCodeT()
	}
	value := v.Value.Ejecutar(ast, gen)
	Variable := ast.GetVariable(v.Name)
	if Variable != nil && Variable.Mutable && Variable.Symbols.Tipo == value.Type {
		//valida el tipo
		if Variable.Symbols.Tipo == environment.INTEGER {
			gen.AddComment("Asignacion de Int")
			newTemp := gen.NewTemp()
			gen.AddGetStack(newTemp, strconv.Itoa(Variable.Symbols.Posicion))
			newTemp2 := gen.NewTemp()
			gen.AddExpression(newTemp2, newTemp, value.Value, "+")
			gen.AddSetStack(strconv.Itoa(Variable.Symbols.Posicion), newTemp2)
			Variable.Symbols.Lin = v.Lin
			Variable.Symbols.Col = v.Col
			Variable.Symbols.Valor = newTemp2
			Variable.Symbols.Scope = ast.ObtenerAmbito()
			ast.ActualizarVariable(Variable)
			return nil
		} else if Variable.Symbols.Tipo == environment.FLOAT {
			gen.AddComment("Asignacion de Float")
			newTemp := gen.NewTemp()
			gen.AddGetStack(newTemp, strconv.Itoa(Variable.Symbols.Posicion))
			newTemp2 := gen.NewTemp()
			gen.AddExpression(newTemp2, newTemp, value.Value, "+")
			gen.AddSetStack(strconv.Itoa(Variable.Symbols.Posicion), newTemp2)
			Variable.Symbols.Lin = v.Lin
			Variable.Symbols.Col = v.Col
			Variable.Symbols.Valor = newTemp2
			Variable.Symbols.Scope = ast.ObtenerAmbito()
			ast.ActualizarVariable(Variable)
		} else if Variable.Symbols.Tipo == environment.STRING {
			//llamar a generar concatstring
			gen.GenerateConcatString()
			//concat
			newTemp := gen.NewTemp()
			gen.AddComment("Asignacion de String")
			envSize := strconv.Itoa(ast.Lista_Variables.Len() + 1)
			tmp1 := gen.NewTemp()
			tmp2 := gen.NewTemp()
			gen.AddExpression(tmp1, "P", envSize, "+")
			gen.AddExpression(tmp1, tmp1, "1", "+")
			gen.AddGetStack(newTemp, strconv.Itoa(Variable.Symbols.Posicion))
			gen.AddSetStack("(int)"+tmp1, newTemp)
			gen.AddExpression(tmp1, tmp1, "1", "+")
			gen.AddSetStack("(int)"+tmp1, value.Value)
			gen.AddExpression("P", "P", envSize, "+")
			gen.AddCall("concatString")
			gen.AddGetStack(tmp2, "(int)P")
			gen.AddExpression("P", "P", envSize, "-")
			gen.AddSetStack(strconv.Itoa(Variable.Symbols.Posicion), tmp2)
			gen.AddBr()

			Variable.Symbols.Lin = v.Lin
			Variable.Symbols.Col = v.Col
			Variable.Symbols.Valor = tmp2
			Variable.Symbols.Scope = ast.ObtenerAmbito()
			ast.ActualizarVariable(Variable)
			return nil
		} else {
			Errores := environment.Errores{
				Descripcion: "Esta operacion += solo se puede hacer con datos de tipo Int o Float y la variable que estas intentando hacerlo es de otro tipo.",
				Fila:        strconv.Itoa(v.Lin),
				Columna:     strconv.Itoa(v.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
			Variable.Symbols.Valor = nil
			return nil
		}
	}

	if Variable != nil && Variable.Mutable && Variable.Symbols.Tipo != value.Type {
		//valida el tipo
		if Variable.Symbols.Tipo == environment.FLOAT && value.Type == environment.INTEGER {
			gen.AddComment("Asignacion de Int-Float")
			newTemp := gen.NewTemp()
			gen.AddGetStack(newTemp, strconv.Itoa(Variable.Symbols.Posicion))
			newTemp2 := gen.NewTemp()
			gen.AddExpression(newTemp2, newTemp, value.Value, "+")
			gen.AddSetStack(strconv.Itoa(Variable.Symbols.Posicion), newTemp2)
			Variable.Symbols.Lin = v.Lin
			Variable.Symbols.Col = v.Col
			Variable.Symbols.Valor = newTemp2
			Variable.Symbols.Scope = ast.ObtenerAmbito()
			ast.ActualizarVariable(Variable)
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
		return nil
	}

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

	if Variable.Symbols.Tipo != value.Type {
		Errores := environment.Errores{
			Descripcion: "Se ha querido asignar un valor no correspondiente a el tipo de dato: \nTipo de dato:" + tipoexpstr2 + "\nTipo de Valor:" + tipoexpstr + ".",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		Variable.Symbols.Valor = nil
		return nil
	}

	return nil
}
