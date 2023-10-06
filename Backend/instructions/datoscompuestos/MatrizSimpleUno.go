package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"container/list"
	"strconv"
	"strings"
)

type MatrizSimpleUno struct {
	Type   interfaces.Expression
	Op     interfaces.Instruction
	Numero string
	Lin    int
	Col    int
}

func NewMatrizSimpleUno(tipo interfaces.Expression, op interfaces.Instruction, numero string, line int, col int) MatrizSimpleUno {
	exp := MatrizSimpleUno{tipo, op, numero, line, col}
	return exp
}

func (o MatrizSimpleUno) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	o.Op.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	tipo := o.Type.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	if strings.Contains(o.Numero, ".") {
		Errores := environment.Errores{
			Descripcion: "No es posible crear la matriz, esta colocando numeros no enteros",
			Fila:        strconv.Itoa(o.Lin),
			Columna:     strconv.Itoa(o.Col),
			Tipo:        "Error Semantico",
			Ambito:      tipo.Val.Symbols.Scope,
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
			Ambito:      tipo.Val.Symbols.Scope,
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
			Ambito:      tipo.Val.Symbols.Scope,
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	lista := list.New()
	lista.PushBack(num)
	arreglo := environment.Valores_Matriz{
		Matriztam: lista,
		Tipo:      tipo.Type,
		Valor:     tipo.Val.Symbols,
	}
	ast.Lista_Matriz_Val.PushFront(arreglo)
	gen.MainCodeF()
	return nil
}
