package instructions

import (
	"Backend/environment"
	"container/list"
)

type ArregloDeclaracionSinLista struct {
	Lin  int
	Col  int
	Name string
	Type environment.TipoExpresion
}

func NewArregloDeclaracionSinLista(lin int, col int, name string, tipo environment.TipoExpresion) ArregloDeclaracionSinLista {
	return ArregloDeclaracionSinLista{lin, col, name, tipo}
}

func (v ArregloDeclaracionSinLista) Ejecutar(ast *environment.AST) interface{} {
	listavalores := list.New()
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
