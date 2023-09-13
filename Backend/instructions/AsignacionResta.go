package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
	"fmt"
	"strconv"
)

type AsignacionResta struct {
	Lin   int
	Col   int
	Name  string
	Type  environment.TipoExpresion
	Value interfaces.Expression
}

func NewAsignacionResta(lin int, col int, name string, value interfaces.Expression) AsignacionResta {
	return AsignacionResta{Lin: lin, Col: col, Name: name, Value: value}
}

func (v AsignacionResta) Ejecutar(ast *environment.AST) interface{} {
	value := v.Value.Ejecutar(ast)
	Variable := ast.GetVariable(v.Name)
	if Variable != nil && Variable.Mutable && Variable.Symbols.Tipo == value.Tipo {
		//valida el tipo
		if Variable.Symbols.Tipo == environment.INTEGER {
			val1, _ := Variable.Symbols.Valor.(int)
			val2, _ := value.Valor.(int)
			num2 := val1 - val2
			Variable.Symbols.Lin = v.Lin
			Variable.Symbols.Col = v.Col
			Variable.Symbols.Valor = num2
			Variable.Symbols.Scope = ast.ObtenerAmbito()
			ast.ActualizarVariable(Variable)
			return nil
		} else if Variable.Symbols.Tipo == environment.FLOAT {
			val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", Variable.Symbols.Valor), 64)
			val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", value.Valor), 64)
			num2 := fmt.Sprintf("%.6f", val1-val2)
			num3, err := strconv.ParseFloat(num2, 64)
			if err != nil {
				fmt.Println(err)
			}
			Variable.Symbols.Lin = v.Lin
			Variable.Symbols.Col = v.Col
			Variable.Symbols.Valor = num3
			Variable.Symbols.Scope = ast.ObtenerAmbito()
			ast.ActualizarVariable(Variable)
			return nil
		} else {
			Errores := environment.Errores{
				Descripcion: "Esta operacion -= solo se puede hacer con datos de tipo Int o Float y la variable que estas intentando hacerlo es de otro tipo.",
				Fila:        strconv.Itoa(value.Lin),
				Columna:     strconv.Itoa(value.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
			Variable.Symbols.Valor = nil
			return nil
		}
	}

	if Variable != nil && Variable.Mutable && Variable.Symbols.Tipo != value.Tipo {
		//valida el tipo
		if Variable.Symbols.Tipo == environment.FLOAT && value.Tipo == environment.INTEGER {
			val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", Variable.Symbols.Valor), 64)
			val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", value.Valor), 64)
			num2 := fmt.Sprintf("%.6f", val1-val2)
			num3, err := strconv.ParseFloat(num2, 64)
			if err != nil {
				fmt.Println(err)
			}
			Variable.Symbols.Lin = v.Lin
			Variable.Symbols.Col = v.Col
			Variable.Symbols.Valor = num3
			Variable.Symbols.Scope = ast.ObtenerAmbito()
			ast.ActualizarVariable(Variable)
			return nil
		}
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

	if Variable.Symbols.Tipo != value.Tipo {
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
