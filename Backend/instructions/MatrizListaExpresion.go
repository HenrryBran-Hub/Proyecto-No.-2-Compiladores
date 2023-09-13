package instructions

import (
	"Backend/environment"
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

func (o MatrizListaExpresion) Ejecutar(ast *environment.AST) interface{} {
	lista := list.New()
	for _, inst := range o.ListaValores {
		if inst == nil {
			continue
		}
		instruction, ok := inst.(interfaces.Expression)
		if !ok {
			continue
		}
		valor := instruction.Ejecutar(ast)
		lista.PushBack(valor)
	}
	arreglo := environment.Valores_Matriz{
		Tipo:     environment.NULL,
		Elements: lista,
	}
	ast.Lista_Matriz_Val.PushBack(arreglo)
	return nil
}
