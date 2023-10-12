package funciones

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
)

type FuncionesDeclaracionR struct {
	Lin    int
	Col    int
	Name   string
	Tipo   environment.TipoExpresion
	Bloque []interface{}
}

func NewFuncionesDeclaracionR(lin int, col int, name string, tipo environment.TipoExpresion, bloque []interface{}) FuncionesDeclaracionR {
	return FuncionesDeclaracionR{Lin: lin, Col: col, Name: name, Tipo: tipo, Bloque: bloque}
}

func (v FuncionesDeclaracionR) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	funcion := environment.Funcion{
		Lin:           v.Lin,
		Col:           v.Col,
		Nombre:        v.Name,
		IsReturn:      true,
		IsParame:      false,
		CodigoFuncion: v.Bloque,
		Tipo:          v.Tipo,
	}

	ast.GuardarFuncion(funcion)
	ambito := ast.ObtenerAmbito()
	ambitonuevo := "funcion" + "-" + ambito
	ast.AumentarAmbito(ambitonuevo)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}

	newlabelr := gen.NewLabel()
	exitla := gen.NewLabel()
	var errorgeneral int = 0

	transferencia := environment.SentenciasdeTransferencia{
		Scope:  ambitonuevo,
		ETrue:  newlabelr,
		EFalse: exitla,
	}
	ast.Lista_Tranferencias.PushBack(transferencia)

	gen.AddTittle(v.Name)
	for _, inst := range v.Bloque {
		if inst == nil {
			continue
		}
		instruction, ok := inst.(interfaces.Instruction)
		if !ok {
			continue
		}
		if !ast.IsMain(ambitonuevo) {
			gen.MainCodeT()
		}
		instruction.Ejecutar(ast, gen)
		if !ast.IsMain(ambitonuevo) {
			gen.MainCodeT()
		}

		rvari := ast.GetVariable("Return")
		if rvari != nil {
			if ast.Lista_Tranferencias.Len() == 0 {
				errorgeneral = 1
			}
		}
		revari := ast.GetVariable("ReturnExp")
		if revari != nil {
			if ast.Lista_Tranferencias.Len() == 0 {
				errorgeneral = 1
			}
			if revari.Symbols.Tipo != funcion.Tipo {
				Errores := environment.Errores{
					Descripcion: "El tipo que esta retornando no es del mismo tipo de la funcion",
					Fila:        strconv.Itoa(v.Lin),
					Columna:     strconv.Itoa(v.Col),
					Tipo:        "Error Semantico",
					Ambito:      ast.ObtenerAmbito(),
				}
				ast.ErroresHTML(Errores)
			}
			gen.AddSetStack("(int)P", revari.Symbols.Valor.(string))
		}
	}

	gen.AddLabel(exitla)
	gen.AddEnd()
	ast.DisminuirAmbito()
	gen.MainCodeF()
	ast.Lista_Funciones_Var.Init()
	ast.Lista_Tranferencias.Remove(ast.Lista_Tranferencias.Back())

	if errorgeneral == 1 {
		Errores := environment.Errores{
			Descripcion: "Se han colocado sentencias de transferencia fuera de ciclos",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
	}

	return nil
}
