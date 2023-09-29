package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
	"container/list"
	"strconv"
)

type ArregloDeclaracionId struct {
	Lin  int
	Col  int
	Prin string
	Type environment.TipoExpresion
	Secu string
}

func NewArregloDeclaracionId(lin int, col int, prin string, tipo environment.TipoExpresion, secu string) ArregloDeclaracionId {
	return ArregloDeclaracionId{lin, col, prin, tipo, secu}
}

func (v ArregloDeclaracionId) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	listavalores := list.New()
	secund := ast.GetArreglo(v.Secu)
	listavalorespt := list.New()
	if secund == nil {
		Errores := environment.Errores{
			Descripcion: "El arreglo que esta intentando asignar no existe: \n Arreglo: " + v.Secu,
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
	} else {
		if secund.Symbols.Tipo != v.Type {
			Errores := environment.Errores{
				Descripcion: "El arreglo que esta intentando asignar es de diferente tipo",
				Fila:        strconv.Itoa(v.Lin),
				Columna:     strconv.Itoa(v.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
			return nil
		}
		for e := secund.Elements.Front(); e != nil; e = e.Next() {
			listavalores.PushBack(e.Value)
		}

		valores := secund.Elements.Front()
		for valores != nil {
			newTemp2 := gen.NewTemp()
			gen.AddAssign(newTemp2, valores.Value.(environment.Value).Value)
			newTemp := gen.NewTemp()
			gen.AddAssign(newTemp, "H")           //Creamos un nuevo puntero Head para que en este vayan los valores
			gen.AddSetHeap("(int)H", newTemp2)    // Agregamos el caracter de escape
			gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion igual que arriba de addexpression
			gen.AddBr()
			listavalorespt.PushBack(newTemp)
			valores = valores.Next()
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
		Name:        v.Prin,
		Symbols:     symbol,
		Mutable:     true,
		Elements:    listavalores,
		ElementsPt:  listavalorespt,
		TipoSimbolo: "Vector",
	}

	ast.GuardarArreglo(vector)
	gen.MainCodeF()
	return nil
}
