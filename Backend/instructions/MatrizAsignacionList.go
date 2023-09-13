package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
	"container/list"
	"strconv"
)

type MatrizAsignacionList struct {
	Name  string
	Expr1 interfaces.Expression
	Expr2 interfaces.Expression
	Op    []interface{}
	Expr3 interfaces.Expression
}

func NewMatrizAsignacionList(name string, exp1, exp2 interfaces.Expression, values []interface{}, exp3 interfaces.Expression) MatrizAsignacionList {
	return MatrizAsignacionList{name, exp1, exp2, values, exp3}
}

func (v MatrizAsignacionList) Ejecutar(ast *environment.AST) interface{} {

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

	listavalores := list.New()
	for _, inst := range v.Op {
		if inst == nil {
			continue
		}
		instruction, ok := inst.(interfaces.Expression)
		if !ok {
			continue
		}
		values := instruction.Ejecutar(ast)
		listavalores.PushBack(values)
	}

	for e := listavalores.Front(); e != nil; e = e.Next() {
		symbol := e.Value.(environment.Symbol)
		if symbol.Tipo != environment.INTEGER {
			Errores := environment.Errores{
				Descripcion: "Las posiciones ingresadas deben de ser Enteros o el resultado de una operacion que de entero",
				Fila:        strconv.Itoa(symbol.Lin),
				Columna:     strconv.Itoa(symbol.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
			return nil
		}
	}

	var valoresSlice []int
	valoresSlice = append(valoresSlice, primerval.Valor.(int))
	valoresSlice = append(valoresSlice, primerval2.Valor.(int))
	for e := listavalores.Front(); e != nil; e = e.Next() {
		symbol := e.Value.(environment.Symbol)
		valor, ok := symbol.Valor.(int)
		if ok {
			valoresSlice = append(valoresSlice, valor)
		}
	}

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
