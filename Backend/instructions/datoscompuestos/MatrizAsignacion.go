package datoscompuestos

import (
	"Backend/interfaces"
)

type MatrizAsignacion struct {
	Name  string
	Expr1 interfaces.Expression
	Expr2 interfaces.Expression
	Expr3 interfaces.Expression
}

func NewMatrizAsignacion(name string, exp1, exp2, exp3 interfaces.Expression) MatrizAsignacion {
	return MatrizAsignacion{name, exp1, exp2, exp3}
}

/*
func (v MatrizAsignacion) Ejecutar(ast *environment.AST) interface{} {

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
		return nil
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
		return nil
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
		return nil
	}

	entrada := v.Expr3.Ejecutar(ast)

	if entrada.Tipo != matriz.Symbols.Tipo {
		Errores := environment.Errores{
			Descripcion: "El valor que intenta ingresar no es del mismo tipo que el de la matriz",
			Fila:        strconv.Itoa(primerval2.Lin),
			Columna:     strconv.Itoa(primerval2.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	ast.IngresarValor(matriz, valoresSlice, entrada.Valor)
	return nil

}
*/
