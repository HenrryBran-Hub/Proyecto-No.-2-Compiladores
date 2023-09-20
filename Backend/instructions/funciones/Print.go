package funciones

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"container/list"
	"fmt"
	"strconv"
)

type Print struct {
	Lin          int
	Col          int
	ListaValores []interface{}
}

func NewPrint(lin int, col int, val []interface{}) Print {
	return Print{lin, col, val}
}

func (p Print) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {

	lista := list.New()
	for _, inst := range p.ListaValores {
		if inst == nil {
			continue
		}
		instruction, ok := inst.(interfaces.Expression)
		if !ok {
			continue
		}
		if ast.ObtenerAmbito() == "Global" {
			gen.MainCode = true
		} else {
			gen.MainCode = false
		}
		valor := instruction.Ejecutar(ast, gen)
		lista.PushBack(valor)
	}
	gen.AddComment("----------Imprimimos----------")
	var result environment.Value
	for e := lista.Front(); e != nil; e = e.Next() {
		result = e.Value.(environment.Value)
		if result.Type == environment.INTEGER {
			newTemp := gen.NewTemp()
			gen.AddGetStack(newTemp, strconv.Itoa(result.Val.Symbols.Posicion))
			gen.AddPrintf("d", "(int)"+fmt.Sprintf("%v", newTemp))
			gen.AddPrintf("c", "10")
			gen.AddBr()
		} else if result.Type == environment.FLOAT {
			newTemp := gen.NewTemp()
			gen.AddGetStack(newTemp, strconv.Itoa(result.Val.Symbols.Posicion))
			gen.AddPrintf("d", "(int)"+fmt.Sprintf("%v", newTemp))
			gen.AddPrintf("c", "10")
			gen.AddBr()
		} else if result.Type == environment.BOOLEAN {
			newLabel := gen.NewLabel()
			for i := result.TrueLabel.Front(); i != nil; i = i.Next() {
				gen.AddLabel(i.Value.(string))
			}
			if result.Val.TEti != "" {
				gen.AddLabel(result.Val.TEti)
			}
			gen.AddPrintf("c", "(char)116")
			gen.AddPrintf("c", "(char)114")
			gen.AddPrintf("c", "(char)117")
			gen.AddPrintf("c", "(char)101")
			gen.AddPrintf("c", "10")
			gen.AddGoto(newLabel)
			for i := result.FalseLabel.Front(); i != nil; i = i.Next() {
				gen.AddLabel(i.Value.(string))
			}
			if result.Val.TEti != "" {
				gen.AddLabel(result.Val.FEti)
			}
			gen.AddPrintf("c", "(char)102")
			gen.AddPrintf("c", "(char)97")
			gen.AddPrintf("c", "(char)108")
			gen.AddPrintf("c", "(char)115")
			gen.AddPrintf("c", "(char)101")
			gen.AddPrintf("c", "10")
			gen.AddLabel(newLabel)
			gen.AddBr()
		} else if result.Type == environment.STRING || result.Type == environment.CHARACTER {
			//llamar a generar printstring
			gen.GeneratePrintString()
			//agregar codigo en el main
			newTemp1 := gen.NewTemp()
			newTemp2 := gen.NewTemp()
			size := strconv.Itoa(ast.Lista_Variables.Len())
			gen.AddExpression(newTemp1, "P", size, "+")     //nuevo temporal en pos vacia
			gen.AddExpression(newTemp1, newTemp1, "1", "+") //se deja espacio de retorno
			gen.AddSetStack("(int)"+newTemp1, result.Value) //se coloca string en parametro que se manda
			gen.AddExpression("P", "P", size, "+")          // cambio de entorno
			gen.AddCall("printString")                      //Llamada
			gen.AddGetStack(newTemp2, "(int)P")             //obtencion retorno
			gen.AddExpression("P", "P", size, "-")          //regreso del entorno
			gen.AddPrintf("c", "10")                        //salto de linea
			gen.AddBr()
		}
	}

	return result
}
