package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
	"container/list"
	"strconv"
)

type StruckFuncionesControlR struct {
	Lin       int
	Col       int
	IdStruk   string
	IdFuncion string
}

func NewStruckFuncionesControlR(lin int, col int, str, fun string) StruckFuncionesControlR {
	return StruckFuncionesControlR{lin, col, str, fun}
}

func (v StruckFuncionesControlR) Ejecutar(ast *environment.AST) environment.Symbol {
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
		return environment.Symbol{Lin: v.Lin, Col: v.Col, Tipo: environment.NULL, Valor: nil}
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
		return environment.Symbol{Lin: v.Lin, Col: v.Col, Tipo: funcion.Tipo, Valor: nil}
	}

	ast.AumentarAmbito("Funcion-" + v.IdStruk + "-" + v.IdFuncion)
	var retorno environment.Symbol
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

	if funcion.IsReturn == true {
		if retorno.Tipo == funcion.Tipo {
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
