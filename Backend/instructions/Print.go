package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
	"container/list"
	"fmt"
)

type Print struct {
	Lin          int
	Col          int
	ListaValores []interface{}
}

func NewPrint(lin int, col int, val []interface{}) Print {
	return Print{lin, col, val}
}

func (p Print) Ejecutar(ast *environment.AST) interface{} {
	lista := list.New()
	for _, inst := range p.ListaValores {
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

	var consoleOut string = ""
	for e := lista.Front(); e != nil; e = e.Next() {
		symbol := e.Value.(environment.Symbol)
		consoleOut += fmt.Sprintf(" %v ", symbol.Valor)
	}
	ast.SetPrint(consoleOut + "\n")
	return nil
}
