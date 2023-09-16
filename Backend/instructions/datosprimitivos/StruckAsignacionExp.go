package datosprimitivos

import (
	"Backend/interfaces"
)

type StruckAsignacionExp struct {
	Lin        int
	Col        int
	IdStruct   string
	IdVariable string
	Op         interfaces.Expression
}

func NewStruckAsignacionExpr(lin int, col int, varstr string, variable string, op interfaces.Expression) StruckAsignacionExp {
	return StruckAsignacionExp{lin, col, varstr, variable, op}
}

/*
func (v StruckAsignacionExp) Ejecutar(ast *environment.AST) interface{} {
	estructura := ast.GetVariableStruc(v.IdStruct)
	if estructura == nil {
		Errores := environment.Errores{
			Descripcion: "El struc que estas intentando ingresar no existe",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return nil
	} else {

		expresion := v.Op.Ejecutar(ast)
		for e := estructura.Strukt.Variables.Front(); e != nil; e = e.Next() {
			valor := e.Value.(environment.Variable)
			if valor.Name == v.IdVariable {
				if valor.Name == v.IdVariable && valor.Mutable && valor.Symbols.Tipo == expresion.Tipo {
					valor.Symbols.Valor = expresion.Valor
					e.Value = valor
					ast.ActualizarVariableStruc(estructura)
				} else {
					Errores := environment.Errores{
						Descripcion: "La variable del struc que estas intentando cambiar posiblemente no sea del mismo tipo o no sea mutable",
						Fila:        strconv.Itoa(v.Lin),
						Columna:     strconv.Itoa(v.Col),
						Tipo:        "Error Semantico",
						Ambito:      ast.ObtenerAmbito(),
					}
					ast.ErroresHTML(Errores)
					return nil
				}
			}
		}
		return nil
	}
}
*/
