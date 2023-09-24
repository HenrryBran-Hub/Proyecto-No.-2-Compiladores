package datoscompuestos

import (
	"Backend/interfaces"
)

type ArregloAppend struct {
	Name  string
	Value interfaces.Expression
}

func NewArregloAppend(name string, value interfaces.Expression) ArregloAppend {
	return ArregloAppend{name, value}
}

/*
func (v ArregloAppend) Ejecutar(ast *environment.AST) interface{} {
	arreglo := ast.GetArreglo(v.Name)
	expre := v.Value.Ejecutar(ast)
	if arreglo == nil {
		Errores := environment.Errores{
			Descripcion: "El arreglo que esta intentando agregar no existe: \n Arreglo: " + v.Name,
			Fila:        strconv.Itoa(expre.Lin),
			Columna:     strconv.Itoa(expre.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return nil
	} else {
		if arreglo.Symbols.Tipo != expre.Tipo {
			Errores := environment.Errores{
				Descripcion: "El valor que esta intentando asignar es de diferente tipo que del arreglo",
				Fila:        strconv.Itoa(expre.Lin),
				Columna:     strconv.Itoa(expre.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
			return nil
		}
	}

	symbol := environment.Symbol{
		Lin:   expre.Lin,
		Col:   expre.Col,
		Tipo:  expre.Tipo,
		Valor: expre.Valor,
		Scope: arreglo.Symbols.Scope,
	}

	arreglo.Elements.PushBack(symbol)

	ast.ActualizarArreglo(v.Name, arreglo)
	return nil
}
*/
