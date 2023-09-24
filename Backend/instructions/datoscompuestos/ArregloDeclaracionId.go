package datoscompuestos

import (
	"Backend/environment"
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

/*
func (v ArregloDeclaracionId) Ejecutar(ast *environment.AST) interface{} {
	listavalores := list.New()
	secund := ast.GetArreglo(v.Secu)
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
		TipoSimbolo: "Vector",
	}

	ast.GuardarArreglo(vector)
	return nil
}
*/
