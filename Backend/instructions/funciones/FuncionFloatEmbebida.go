package funciones

import (
	"Backend/interfaces"
)

type FuncionFloatEmbebida struct {
	Op interfaces.Expression
}

func NewFuncionFloatEmbebida(op interfaces.Expression) FuncionFloatEmbebida {
	exp := FuncionFloatEmbebida{Op: op}
	return exp
}

/*
func (o FuncionFloatEmbebida) Ejecutar(ast *environment.AST) environment.Symbol {
	var op environment.Symbol
	op = o.Op.Ejecutar(ast)

	if op.Tipo == environment.INTEGER {
		stringValue, _ := strconv.ParseFloat(fmt.Sprintf("%v", op.Valor), 64)
		num2 := fmt.Sprintf("%.6f", stringValue)
		num3, err := strconv.ParseFloat(num2, 64)
		if err != nil {
			fmt.Println(err)
		}
		return environment.Symbol{Lin: op.Lin, Col: op.Col, Tipo: environment.FLOAT, Valor: num3}
	} else if op.Tipo == environment.FLOAT {
		return environment.Symbol{Lin: op.Lin, Col: op.Col, Tipo: environment.FLOAT, Valor: op.Valor}
	} else if op.Tipo == environment.STRING {
		stringValue, ok := op.Valor.(string)
		if !ok {
			Errores := environment.Errores{
				Descripcion: "No es posible convertir el dato a tipo Float, el dato es tipo String pero no ha de tener solo valores numericos, retornamos nil",
				Fila:        strconv.Itoa(op.Lin),
				Columna:     strconv.Itoa(op.Col),
				Tipo:        "Error Semantico",
				Ambito:      op.Scope,
			}
			ast.ErroresHTML(Errores)
			return environment.Symbol{Lin: op.Lin, Col: op.Col, Tipo: environment.FLOAT, Valor: nil}
		}
		floatValue, err := strconv.ParseFloat(stringValue, 64)
		if err != nil {
			Errores := environment.Errores{
				Descripcion: "No es posible convertir el dato a tipo Float, el dato es tipo String pero no ha de tener solo valores numericos, retornamos nil",
				Fila:        strconv.Itoa(op.Lin),
				Columna:     strconv.Itoa(op.Col),
				Tipo:        "Error Semantico",
				Ambito:      op.Scope,
			}
			ast.ErroresHTML(Errores)
			return environment.Symbol{Lin: op.Lin, Col: op.Col, Tipo: environment.FLOAT, Valor: nil}
		}
		num2 := fmt.Sprintf("%.6f", floatValue)
		num3, err := strconv.ParseFloat(num2, 64)
		if err != nil {
			fmt.Println(err)
		}
		return environment.Symbol{Lin: op.Lin, Col: op.Col, Tipo: environment.FLOAT, Valor: num3}
	} else {
		Errores := environment.Errores{
			Descripcion: "No es posible convertir el dato a tipo Float, no es compatible el tipo de dato ingresado.(Character,Bool)",
			Fila:        strconv.Itoa(op.Lin),
			Columna:     strconv.Itoa(op.Col),
			Tipo:        "Error Semantico",
			Ambito:      op.Scope,
		}
		ast.ErroresHTML(Errores)
		return environment.Symbol{Lin: op.Lin, Col: op.Col, Tipo: environment.FLOAT, Valor: nil}
	}
}

*/
