package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"container/list"
	"strconv"
	"strings"
)

type MatrizSimpleDos struct {
	Type   interfaces.Expression
	Expre  interfaces.Expression
	Numero string
	Lin    int
	Col    int
}

func NewMatrizSimpleDos(tipo interfaces.Expression, expre interfaces.Expression, numero string, line int, col int) MatrizSimpleDos {
	exp := MatrizSimpleDos{tipo, expre, numero, line, col}
	return exp
}

func (o MatrizSimpleDos) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	tipo := o.Type.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	expre := o.Expre.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}

	if tipo.Type != expre.Type {
		Errores := environment.Errores{
			Descripcion: "No es posible crear la matriz, esta colocando diferentes tipos de datos",
			Fila:        strconv.Itoa(o.Lin),
			Columna:     strconv.Itoa(o.Col),
			Tipo:        "Error Semantico",
			Ambito:      expre.Val.Symbols.Scope,
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	if strings.Contains(o.Numero, ".") {
		Errores := environment.Errores{
			Descripcion: "No es posible crear la matriz, esta colocando numeros no enteros",
			Fila:        strconv.Itoa(o.Lin),
			Columna:     strconv.Itoa(o.Col),
			Tipo:        "Error Semantico",
			Ambito:      expre.Val.Symbols.Scope,
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	num, err := strconv.Atoi(o.Numero)
	if err != nil {
		Errores := environment.Errores{
			Descripcion: "No es posible crear la matriz, Error en la conversion",
			Fila:        strconv.Itoa(o.Lin),
			Columna:     strconv.Itoa(o.Col),
			Tipo:        "Error Semantico",
			Ambito:      expre.Val.Symbols.Scope,
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	if num <= 0 {
		Errores := environment.Errores{
			Descripcion: "No es posible crear la matriz, el numero es menor o igual a 0",
			Fila:        strconv.Itoa(o.Lin),
			Columna:     strconv.Itoa(o.Col),
			Tipo:        "Error Semantico",
			Ambito:      expre.Val.Symbols.Scope,
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	lista := list.New()
	lista.PushBack(num)
	expre.Val.Symbols.Valor = expre.Value
	expre.Val.Symbols.Tipo = expre.Type
	arreglo := environment.Valores_Matriz{
		Tipo:      tipo.Type,
		Valor:     expre.Val.Symbols,
		Matriztam: lista,
	}
	ast.Lista_Matriz_Val.PushFront(arreglo)
	gen.MainCodeF()
	return nil
}
