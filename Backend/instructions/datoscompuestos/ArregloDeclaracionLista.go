package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
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

func (v ArregloDeclaracionLista) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	listavalores := list.New()
	for _, inst := range v.Lista {
		if inst == nil {
			continue
		}
		instruction, ok := inst.(interfaces.Expression)
		if !ok {
			continue
		}
		values := instruction.Ejecutar(ast, gen)
		if !ast.IsMain(ast.ObtenerAmbito()) {
			gen.MainCodeT()
		}
		listavalores.PushBack(values)
	}

	for e := listavalores.Front(); e != nil; e = e.Next() {
		symbol := e.Value.(environment.Value)
		if symbol.Type != v.Type {
			Errores := environment.Errores{
				Descripcion: "Existen parametros que no son compatibles con el tipo de arreglo definido",
				Fila:        strconv.Itoa(symbol.Val.Symbols.Lin),
				Columna:     strconv.Itoa(symbol.Val.Symbols.Lin),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
			return nil
		}
	}

	symbol := environment.Symbol{
		Lin:      v.Lin,
		Col:      v.Col,
		Tipo:     v.Type,
		Valor:    nil,
		Scope:    ast.ObtenerAmbito(),
		Posicion: ast.PosicionStack,
	}

	gen.AddComment("Declaracion de Vector")

	listavalorespunteros := list.New()
	for e := listavalores.Front(); e != nil; e = e.Next() {
		value := e.Value.(environment.Value)
		newTemp := gen.NewTemp()
		gen.AddAssign(newTemp, "H")           //Creamos un nuevo puntero Head para que en este vayan los valores
		gen.AddSetHeap("(int)H", value.Value) // Agregamos el caracter de escape
		gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion igual que arriba de addexpression
		gen.AddBr()
		listavalorespunteros.PushBack(newTemp)
	}

	vector := environment.Vector{
		Name:        v.Name,
		Symbols:     symbol,
		Mutable:     true,
		Elements:    listavalores,
		ElementsPt:  listavalorespunteros,
		TipoSimbolo: "Vector",
	}

	ast.GuardarArreglo(vector)
	gen.MainCodeF()
	return nil
}
