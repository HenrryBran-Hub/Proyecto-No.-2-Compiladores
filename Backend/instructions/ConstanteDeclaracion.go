package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
	"strconv"
)

type ConstanteDeclaracion struct {
	Lin   int
	Col   int
	Name  string
	Type  environment.TipoExpresion
	Value interfaces.Expression
}

func NewConstanteDeclaration(lin int, col int, name string, tipo environment.TipoExpresion, value interfaces.Expression) ConstanteDeclaracion {
	return ConstanteDeclaracion{lin, col, name, tipo, value}
}

func (v ConstanteDeclaracion) Ejecutar(ast *environment.AST) interface{} {
	value := v.Value.Ejecutar(ast)
	symbol := environment.Symbol{
		Lin:   v.Lin,
		Col:   v.Col,
		Tipo:  v.Type,
		Valor: value.Valor,
		Scope: ast.ObtenerAmbito(),
	}
	Variable := environment.Variable{
		Name:        v.Name,
		Symbols:     symbol,
		Mutable:     false,
		TipoSimbolo: "Constante",
	}

	var tipoexp int = -1
	var tipoexpstr string
	switch symbol.Valor.(type) {
	case int:
		tipoexp = 0
		tipoexpstr = "Int"
	case float64:
		tipoexp = 1
		tipoexpstr = "Float"
	case string:
		tipoexp = 2
		tipoexpstr = "String"
	case bool:
		tipoexp = 3
		tipoexpstr = "Boolean"
	case rune:
		tipoexp = 4
		tipoexpstr = "Character"
	default:
		tipoexp = 5
		tipoexpstr = "nil"
	}

	var tipoexpstr2 string
	switch v.Type {
	case 0:
		tipoexpstr2 = "Int"
	case 1:
		tipoexpstr2 = "Float"
	case 2:
		tipoexpstr2 = "String"
	case 3:
		tipoexpstr2 = "Boolean"
	case 4:
		tipoexpstr2 = "Character"
	default:
		tipoexpstr2 = "nil"
	}

	if tipoexp != int(Variable.Symbols.Tipo) {
		Errores := environment.Errores{
			Descripcion: "Se ha querido asignar un valor no correspondiente a el tipo de dato: \nTipo de dato:" + tipoexpstr2 + "\nTipo de Valor:" + tipoexpstr + ".",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	ast.GuardarVariable(Variable)
	return nil
}
