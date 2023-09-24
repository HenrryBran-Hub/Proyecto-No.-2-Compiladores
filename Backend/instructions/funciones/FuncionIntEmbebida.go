package funciones

import (
	"Backend/interfaces"
)

type FuncionIntEmbebida struct {
	Op interfaces.Expression
}

func NewFuncionIntEmbebida(op interfaces.Expression) FuncionIntEmbebida {
	exp := FuncionIntEmbebida{Op: op}
	return exp
}

/*
func (o FuncionIntEmbebida) Ejecutar(ast *environment.AST) environment.Symbol {
	var op environment.Symbol
	op = o.Op.Ejecutar(ast)

	if op.Tipo == environment.INTEGER {
		return environment.Symbol{Lin: op.Lin, Col: op.Col, Tipo: environment.INTEGER, Valor: op.Valor}
	} else if op.Tipo == environment.FLOAT {
		stringValue, _ := strconv.ParseFloat(fmt.Sprintf("%v", op.Valor), 64)
		intValue := int(stringValue)
		return environment.Symbol{Lin: op.Lin, Col: op.Col, Tipo: environment.INTEGER, Valor: intValue}
	} else if op.Tipo == environment.STRING {
		stringValue, ok := op.Valor.(string)
		if !ok {
			Errores := environment.Errores{
				Descripcion: "No es posible convertir el dato a tipo Int, el dato es tipo String pero no ha de tener solo valores numericos, retornamos nil",
				Fila:        strconv.Itoa(op.Lin),
				Columna:     strconv.Itoa(op.Col),
				Tipo:        "Error Semantico",
				Ambito:      op.Scope,
			}
			ast.ErroresHTML(Errores)
			return environment.Symbol{Lin: op.Lin, Col: op.Col, Tipo: environment.INTEGER, Valor: nil}
		}
		floatValue, err := strconv.ParseFloat(stringValue, 64)
		if err != nil {
			Errores := environment.Errores{
				Descripcion: "No es posible convertir el dato a tipo Int, el dato es tipo String pero no ha de tener solo valores numericos, retornamos nil",
				Fila:        strconv.Itoa(op.Lin),
				Columna:     strconv.Itoa(op.Col),
				Tipo:        "Error Semantico",
				Ambito:      op.Scope,
			}
			ast.ErroresHTML(Errores)
			return environment.Symbol{Lin: op.Lin, Col: op.Col, Tipo: environment.INTEGER, Valor: nil}
		}
		intValue := int(floatValue)
		return environment.Symbol{Lin: op.Lin, Col: op.Col, Tipo: environment.INTEGER, Valor: intValue}
	} else {
		Errores := environment.Errores{
			Descripcion: "No es posible convertir el dato a tipo Int, no es compatible el tipo de dato ingresado.(Character,Bool)",
			Fila:        strconv.Itoa(op.Lin),
			Columna:     strconv.Itoa(op.Col),
			Tipo:        "Error Semantico",
			Ambito:      op.Scope,
		}
		ast.ErroresHTML(Errores)
		return environment.Symbol{Lin: op.Lin, Col: op.Col, Tipo: environment.INTEGER, Valor: nil}
	}
}
*/
