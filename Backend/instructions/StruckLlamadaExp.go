package instructions

import (
	"Backend/environment"
	"strconv"
)

type StruckLlamadaExp struct {
	Lin      int
	Col      int
	NameStru string
	NameVar  string
}

func NewStruckLlamadaExp(lin, col int, stru, name string) StruckLlamadaExp {
	return StruckLlamadaExp{lin, col, stru, name}
}

func (v StruckLlamadaExp) Ejecutar(ast *environment.AST) environment.Symbol {
	existvarstruc := ast.GetVariableStruc(v.NameStru)
	if existvarstruc == nil {
		Errores := environment.Errores{
			Descripcion: "No existe el struc que esta intentando ingresar " + v.NameStru,
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return environment.Symbol{Lin: v.Lin, Col: v.Col, Tipo: environment.NULL, Valor: nil}
	}

	symbol := environment.Symbol{
		Lin:   v.Lin,
		Col:   v.Col,
		Tipo:  environment.NULL,
		Valor: nil,
		Scope: ast.ObtenerAmbito(),
	}

	for e := existvarstruc.Strukt.Variables.Front(); e != nil; e = e.Next() {
		valor := e.Value.(environment.Variable)
		if valor.Name == v.NameVar {
			symbol.Tipo = valor.Symbols.Tipo
			symbol.Valor = valor.Symbols.Valor
		}
	}

	return symbol
}
