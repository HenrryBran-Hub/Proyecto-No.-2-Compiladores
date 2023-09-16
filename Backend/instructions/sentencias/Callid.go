package sentencias

type Callid struct {
	Lin  int
	Col  int
	Name string
}

func NewCallid(lin int, col int, name string) Callid {
	exp := Callid{Lin: lin, Col: col, Name: name}
	return exp
}

/*
func (o Callid) Ejecutar(ast *environment.AST) environment.Symbol {
	variable := ast.GetVariable(o.Name)
	if variable != nil {
		return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: variable.Symbols.Tipo, Valor: variable.Symbols.Valor}
	} else {
		Errores := environment.Errores{
			Descripcion: "La variale que esta intentando llamar no existe: \n Variable: " + o.Name,
			Fila:        strconv.Itoa(o.Lin),
			Columna:     strconv.Itoa(o.Col),
			Tipo:        "Error Semantico",
		}
		ast.ErroresHTML(Errores)
		var result interface{}
		return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.NULL, Valor: result}
	}
}
*/
