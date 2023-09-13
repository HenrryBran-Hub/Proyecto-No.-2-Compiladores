package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
	"strconv"
)

type MatrizObtencion struct {
	Name  string
	Expr1 interfaces.Expression
	Expr2 interfaces.Expression
}

func NewMatrizObtencion(name string, exp1, exp2 interfaces.Expression) MatrizObtencion {
	return MatrizObtencion{name, exp1, exp2}
}

func (v MatrizObtencion) Ejecutar(ast *environment.AST) environment.Symbol {

	primerval := v.Expr1.Ejecutar(ast)

	if primerval.Tipo != environment.INTEGER {
		Errores := environment.Errores{
			Descripcion: "Las posiciones ingresadas deben de ser Enteros o el resultado de una operacion que de entero",
			Fila:        strconv.Itoa(primerval.Lin),
			Columna:     strconv.Itoa(primerval.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return environment.Symbol{Lin: primerval.Lin, Col: primerval.Col, Tipo: environment.INTEGER, Valor: nil}
	}

	primerval2 := v.Expr2.Ejecutar(ast)

	if primerval2.Tipo != environment.INTEGER {
		Errores := environment.Errores{
			Descripcion: "Las posiciones ingresadas deben de ser Enteros o el resultado de una operacion que de entero",
			Fila:        strconv.Itoa(primerval2.Lin),
			Columna:     strconv.Itoa(primerval2.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return environment.Symbol{Lin: primerval2.Lin, Col: primerval2.Col, Tipo: environment.INTEGER, Valor: nil}
	}

	var valoresSlice []int
	valoresSlice = append(valoresSlice, primerval.Valor.(int))
	valoresSlice = append(valoresSlice, primerval2.Valor.(int))

	matriz := ast.GetMatriz(v.Name)
	if matriz == nil {
		Errores := environment.Errores{
			Descripcion: "Las posiciones ingresadas deben de ser Enteros o el resultado de una operacion que de entero",
			Fila:        strconv.Itoa(primerval.Lin),
			Columna:     strconv.Itoa(primerval.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return environment.Symbol{Lin: primerval.Lin, Col: primerval.Col, Tipo: environment.INTEGER, Valor: nil}
	}

	valor := ast.ObtenerValor(*matriz, valoresSlice)
	if valor != nil {
		return environment.Symbol{Lin: primerval.Lin, Col: primerval.Col, Tipo: matriz.Symbols.Tipo, Valor: valor}
	} else {
		return environment.Symbol{Lin: primerval.Lin, Col: primerval.Col, Tipo: matriz.Symbols.Tipo, Valor: nil}
	}
}
