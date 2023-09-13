package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
	"container/list"
	"strconv"
)

type ArregloDeclaracionLista struct {
	Lin   int
	Col   int
	Name  string
	Type  environment.TipoExpresion
	Lista []interface{}
}

func NewArregloDeclaracionLista(lin int, col int, name string, tipo environment.TipoExpresion, values []interface{}) ArregloDeclaracionLista {
	return ArregloDeclaracionLista{lin, col, name, tipo, values}
}

func (v ArregloDeclaracionLista) Ejecutar(ast *environment.AST) interface{} {
	listavalores := list.New()
	for _, inst := range v.Lista {
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
		if symbol.Tipo != v.Type {
			Errores := environment.Errores{
				Descripcion: "Existen parametros que no son compatibles con el tipo de arreglo definido",
				Fila:        strconv.Itoa(symbol.Lin),
				Columna:     strconv.Itoa(symbol.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
			return nil
		}
	}

	symbol := environment.Symbol{
		Lin:   v.Lin,
		Col:   v.Col,
		Tipo:  v.Type,
		Valor: nil,
		Scope: ast.ObtenerAmbito(),
	}
	vector := environment.Vector{
		Name:        v.Name,
		Symbols:     symbol,
		Mutable:     true,
		Elements:    listavalores,
		TipoSimbolo: "Vector",
	}

	ast.GuardarArreglo(vector)
	return nil
}
