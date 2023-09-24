package funciones

type FuncionesControl struct {
	Lin  int
	Col  int
	Name string
}

func NewFuncionesControl(lin int, col int, name string) FuncionesControl {
	return FuncionesControl{lin, col, name}
}

/*
func (v FuncionesControl) Ejecutar(ast *environment.AST) interface{} {
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
		return nil
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
		return nil
	}

	ast.AumentarAmbito("Funcion-" + v.Name)
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
			break
		}
		revari := ast.GetVariable("ReturnExp")
		if revari != nil {
			Errores := environment.Errores{
				Descripcion: "Se esta devolviendo un valor en una funcion que no debe de tener retorno",
				Fila:        strconv.Itoa(v.Lin),
				Columna:     strconv.Itoa(v.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
			ast.Lista_Funciones_Par.Init()
			ast.DisminuirAmbito()
			return nil
		}
	}
	ast.DisminuirAmbito()

	if existfun.IsReturn == true {
		Errores := environment.Errores{
			Descripcion: "Se esta declarando una funcion con retorno fuera de un ambiente que acepte esto",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		ast.Lista_Funciones_Par.Init()
		return nil
	} else {
		ast.Lista_Funciones_Par.Init()
		return nil
	}
}
*/
