package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
	"strconv"
)

type StruckFuncionesControl struct {
	Lin       int
	Col       int
	IdStruk   string
	IdFuncion string
}

func NewStruckFuncionesControl(lin int, col int, str, fun string) StruckFuncionesControl {
	return StruckFuncionesControl{lin, col, str, fun}
}

func (v StruckFuncionesControl) Ejecutar(ast *environment.AST) interface{} {
	existfun := ast.GetVariableStruc(v.IdStruk)
	if existfun == nil {
		Errores := environment.Errores{
			Descripcion: "No existe el struck con ese nombre",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		ast.Lista_Funciones_Par.Init()
		return nil
	}

	funcion := environment.Funcion{}
	for e := existfun.Strukt.Funciones.Front(); e != nil; e = e.Next() {
		val := e.Value.(environment.Funcion)
		if val.Nombre == v.IdFuncion {
			funcion = val
		}
	}

	if funcion.IsParame == true && ast.Lista_Funciones_Par.Len() > 0 {
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

	ast.AumentarAmbito("Funcion-" + v.IdStruk + "-" + v.IdFuncion)
	for _, inst := range funcion.CodigoFuncion {
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

	if funcion.IsReturn == true {
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
