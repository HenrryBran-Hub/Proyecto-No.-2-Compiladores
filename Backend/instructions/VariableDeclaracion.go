package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
	"strconv"
)

type VariableDeclaracion struct {
	Lin   int
	Col   int
	Name  string
	Type  environment.TipoExpresion
	Value interfaces.Expression
}

func NewVariableDeclaration(lin int, col int, name string, tipo environment.TipoExpresion, value interfaces.Expression) VariableDeclaracion {
	return VariableDeclaracion{lin, col, name, tipo, value}
}

func (v VariableDeclaracion) Ejecutar(ast *environment.AST) interface{} {
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
		Mutable:     true,
		TipoSimbolo: "Variable",
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

	if v.Type == 0 && value.Valor == 0 {
		Variable.Symbols.Tipo = environment.INTEGER
		tipoexp = 0
	}

	if v.Type == 1 && tipoexp == 0 {
		Variable.Symbols.Tipo = environment.FLOAT
		tipoexp = 1
	}

	if v.Type == 4 && value.Tipo == 4 {
		Variable.Symbols.Tipo = environment.CHARACTER
		tipoexp = 4
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
		Variable.Symbols.Valor = nil
		Variable.Symbols.Tipo = environment.INTEGER
	}
	v.arbol(ast, tipoexpstr)
	ast.GuardarVariable(Variable)
	return nil
}

func (v VariableDeclaracion) arbol(ast *environment.AST, tipoexpstr string) {
	var operands2 []*environment.Node
	for i := 0; i < 1; i++ {
		operands2 = append(operands2, ast.Pop())
	}

	ast.Id += 1
	ast.Push(&environment.Node{
		Id:       ast.Id,
		Label:    "Expresion",
		Children: operands2,
	})

	ast.Id += 1
	ast.Push(&environment.Node{Id: ast.Id, Label: "IG"})

	ast.Id += 1
	ast.Push(&environment.Node{Id: ast.Id, Label: tipoexpstr})

	var operands3 []*environment.Node
	for i := 0; i < 1; i++ {
		operands3 = append(operands3, ast.Pop())
	}
	ast.Id += 1
	ast.Push(&environment.Node{
		Id:       ast.Id,
		Label:    "TipoDato",
		Children: operands3,
	})

	ast.Id += 1
	ast.Push(&environment.Node{Id: ast.Id, Label: "DOS_PUNTOS(:)"})

	ast.Id += 1
	ast.Push(&environment.Node{Id: ast.Id, Label: v.Name})

	var operands4 []*environment.Node
	for i := 0; i < 1; i++ {
		operands4 = append(operands4, ast.Pop())
	}
	ast.Id += 1
	ast.Push(&environment.Node{
		Id:       ast.Id,
		Label:    "ID_VALIDO",
		Children: operands4,
	})

	ast.Id += 1
	ast.Push(&environment.Node{Id: ast.Id, Label: "Var"})

	var operands []*environment.Node
	for i := 0; i < 6; i++ {
		operands = append(operands, ast.Pop())
	}
	ast.Id += 1
	ast.Push(&environment.Node{
		Id:       ast.Id,
		Label:    "VariableDeclaracion",
		Children: operands,
	})
}
