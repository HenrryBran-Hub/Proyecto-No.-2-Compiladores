package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
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

func (v MatrizObtencion) Ejecutar(ast *environment.AST, gen *generator.Generator) environment.Value {

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
		result := environment.Value{}
		return result
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
		result := environment.Value{}
		return result
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
		result := environment.Value{}
		return result
	}

	valor := ast.ObtenerValor(*matriz, valoresSlice)
	if valor != nil {
		symbol := environment.Symbol{
			Lin:      primerval.Val.Symbols.Lin,
			Col:      primerval.Val.Symbols.Col,
			Tipo:     matriz.Symbols.Tipo,
			Scope:    matriz.Symbols.Scope,
			TipoDato: environment.MATRIZ,
			Posicion: matriz.Symbols.Posicion,
		}
		Variable := environment.Variable{
			Name:        matriz.Name,
			Symbols:     symbol,
			Mutable:     false,
			TipoSimbolo: "Variable",
		}

		if matriz.Symbols.Tipo == environment.INTEGER || matriz.Symbols.Tipo == environment.FLOAT {
			gen.AddBr()
			newTemp := gen.NewTemp()
			gen.AddGetHeap(newTemp, "(int)"+valor.(string))
			gen.AddSetStack(strconv.Itoa(matriz.Symbols.Posicion), newTemp)
			gen.AddBr()
			gen.MainCodeF()
			return environment.NewValue(newTemp, false, matriz.Symbols.Tipo, false, false, false, Variable)
		} else {
			gen.AddBr()
			newTemp := gen.NewTemp()
			gen.AddAssign(newTemp, valor.(string))
			gen.AddSetStack(strconv.Itoa(matriz.Symbols.Posicion), newTemp)
			gen.AddBr()
			gen.MainCodeF()
			return environment.NewValue(newTemp, false, matriz.Symbols.Tipo, false, false, false, Variable)
		}
	} else {
		symbol := environment.Symbol{
			Lin:      primerval.Val.Symbols.Lin,
			Col:      primerval.Val.Symbols.Col,
			Tipo:     matriz.Symbols.Tipo,
			Scope:    matriz.Symbols.Scope,
			TipoDato: environment.MATRIZ,
			Posicion: matriz.Symbols.Posicion,
		}
		Variable := environment.Variable{
			Name:        matriz.Name,
			Symbols:     symbol,
			Mutable:     false,
			TipoSimbolo: "Variable",
		}
		gen.MainCodeF()
		return environment.NewValue("201314439", false, environment.NULL, false, false, false, Variable)
	}
}
