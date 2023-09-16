package funciones

type FuncionesControlR struct {
	Lin  int
	Col  int
	Name string
}

func NewFuncionesControlR(lin int, col int, name string) FuncionesControlR {
	return FuncionesControlR{lin, col, name}
}

/*
func (v FuncionesControlR) Ejecutar(ast *environment.AST) environment.Symbol {
	existfun := ast.GetFuncion(v.Name)
	if existfun == nil {
		Errores := environment.Errores{
			Descripcion: "No existe la funcion",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		ast.Lista_Funciones_Par.Init()
		return environment.Symbol{Lin: v.Lin, Col: v.Col, Tipo: existfun.Tipo, Valor: nil}
	}

	if existfun.IsParame == true && ast.Lista_Funciones_Par.Len() > 0 {
		Errores := environment.Errores{
			Descripcion: "La funcion no debe de tener parametros",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		ast.Lista_Funciones_Par.Init()
		return environment.Symbol{Lin: v.Lin, Col: v.Col, Tipo: existfun.Tipo, Valor: nil}
	}

	ast.AumentarAmbito("Funcion-" + v.Name)
	var retorno environment.Symbol
	for _, inst := range existfun.CodigoFuncion {
		if inst == nil {
			continue
		}
		instruction, ok := inst.(interfaces.Instruction)
		if !ok {
			continue
		}
		instruction.Ejecutar(ast)
		rvari := ast.GetVariable("Return")
		if rvari != nil {
			retorno = rvari.Symbols
			break
		}
		revari := ast.GetVariable("ReturnExp")
		if revari != nil {
			retorno = revari.Symbols
			break
		}

	}

	listavariablesinterna2 := list.New()
	for e := ast.Lista_Variables.Front(); e != nil; e = e.Next() {
		listavariablesinterna2.PushBack(e.Value)
	}
	ast.DisminuirAmbito()

	if existfun.IsReturn == true {
		if retorno.Tipo == existfun.Tipo {
			ast.Lista_Funciones_Par.Init()
			return environment.Symbol{Lin: retorno.Lin, Col: retorno.Col, Tipo: retorno.Tipo, Valor: retorno.Valor}
		} else {
			Errores := environment.Errores{
				Descripcion: "El tipo de dato devuelto no es del tipo de la funcion",
				Fila:        strconv.Itoa(v.Lin),
				Columna:     strconv.Itoa(v.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
			ast.Lista_Funciones_Par.Init()
			return environment.Symbol{Lin: retorno.Lin, Col: retorno.Col, Tipo: retorno.Tipo, Valor: nil}
		}
	} else {
		Errores := environment.Errores{
			Descripcion: "Esta llamando una funcion sin retorno a la operacion",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		ast.Lista_Funciones_Par.Init()
		return environment.Symbol{Lin: retorno.Lin, Col: retorno.Col, Tipo: retorno.Tipo, Valor: nil}
	}
}
*/
