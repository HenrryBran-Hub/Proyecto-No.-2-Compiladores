package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"container/list"
)

type MatrizListaExpresion struct {
	ListaValores []interface{}
}

func NewMatrizListaExpresion(lista []interface{}) MatrizListaExpresion {
	exp := MatrizListaExpresion{lista}
	return exp
}

func (o MatrizListaExpresion) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	lista := list.New()
	for _, inst := range o.ListaValores {
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
		valor := instruction.Ejecutar(ast, gen)
		if !ast.IsMain(ast.ObtenerAmbito()) {
			gen.MainCodeT()
		}
		lista.PushBack(valor)
	}
	arreglo := environment.Valores_Matriz{
		Tipo:     environment.NULL,
		Elements: lista,
	}
	ast.Lista_Matriz_Val.PushBack(arreglo)
	gen.MainCodeF()
	return nil
}
