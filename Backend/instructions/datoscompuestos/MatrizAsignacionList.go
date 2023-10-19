package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
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

func (v MatrizAsignacionList) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {

	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	primerval := v.Expr1.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}

	if primerval.Type != environment.INTEGER {
		Errores := environment.Errores{
			Descripcion: "Las posiciones ingresadas deben de ser Enteros o el resultado de una operacion que de entero",
			Fila:        strconv.Itoa(primerval.Val.Symbols.Lin),
			Columna:     strconv.Itoa(primerval.Val.Symbols.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	primerval2 := v.Expr2.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}

	if primerval2.Type != environment.INTEGER {
		Errores := environment.Errores{
			Descripcion: "Las posiciones ingresadas deben de ser Enteros o el resultado de una operacion que de entero",
			Fila:        strconv.Itoa(primerval2.Val.Symbols.Lin),
			Columna:     strconv.Itoa(primerval2.Val.Symbols.Col),
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
		if !ast.IsMain(ast.ObtenerAmbito()) {
			gen.MainCodeT()
		}
		values := instruction.Ejecutar(ast, gen)
		if !ast.IsMain(ast.ObtenerAmbito()) {
			gen.MainCodeT()
		}
		listavalores.PushBack(values)
	}

	for e := listavalores.Front(); e != nil; e = e.Next() {
		symbol := e.Value.(environment.Value)
		if symbol.Type != environment.INTEGER {
			Errores := environment.Errores{
				Descripcion: "Las posiciones ingresadas deben de ser Enteros o el resultado de una operacion que de entero",
				Fila:        strconv.Itoa(symbol.Val.Symbols.Lin),
				Columna:     strconv.Itoa(symbol.Val.Symbols.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
			return nil
		}
	}

	var valoresSlice []int
	valoresSlice = append(valoresSlice, primerval.IntValue)
	valoresSlice = append(valoresSlice, primerval2.IntValue)
	for e := listavalores.Front(); e != nil; e = e.Next() {
		symbol := e.Value.(environment.Value)
		valoresSlice = append(valoresSlice, symbol.IntValue)
	}

	matriz := ast.GetMatriz(v.Name)
	if matriz == nil {
		Errores := environment.Errores{
			Descripcion: "Las posiciones ingresadas deben de ser Enteros o el resultado de una operacion que de entero",
			Fila:        strconv.Itoa(primerval.Val.Symbols.Lin),
			Columna:     strconv.Itoa(primerval.Val.Symbols.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	entrada := v.Expr3.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}

	if entrada.Type != matriz.Symbols.Tipo {
		Errores := environment.Errores{
			Descripcion: "El valor que intenta ingresar no es del mismo tipo que el de la matriz",
			Fila:        strconv.Itoa(primerval2.Val.Symbols.Lin),
			Columna:     strconv.Itoa(primerval2.Val.Symbols.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	gen.AddComment("Datos Compuestos Matriz-Asignacion-List")
	if matriz.Symbols.Tipo == environment.INTEGER || matriz.Symbols.Tipo == environment.FLOAT {
		var ingresado interface{}
		ingresado = entrada.Value
		gen.AddBr()
		newTemp := gen.NewTemp()
		gen.AddAssign(newTemp, "H")
		gen.AddSetHeap("(int)H", entrada.Value)
		gen.AddExpression("H", "H", "1", "+")
		gen.AddBr()
		ingresado = newTemp
		ast.IngresarValor(matriz, valoresSlice, ingresado)
	} else {
		var ingresado interface{}
		ingresado = entrada.Value
		gen.AddBr()
		newTemp := gen.NewTemp()
		gen.AddAssign(newTemp, entrada.Value)
		gen.AddBr()
		ingresado = newTemp
		ast.IngresarValor(matriz, valoresSlice, ingresado)
	}

	gen.AddBr()
	gen.MainCodeF()
	return nil
}
