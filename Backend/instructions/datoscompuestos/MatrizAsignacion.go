package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
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

func (v MatrizAsignacion) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {

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

	var valoresSlice []int
	valoresSlice = append(valoresSlice, primerval.IntValue)
	valoresSlice = append(valoresSlice, primerval2.IntValue)

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

	gen.AddComment("Datos Compuestos Matriz-Asignacion")

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
