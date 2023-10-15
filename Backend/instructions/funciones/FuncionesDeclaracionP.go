package funciones

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"container/list"
	"strconv"
)

type FuncionesDeclaracionP struct {
	Lin        int
	Col        int
	Name       string
	Parametros interfaces.Instruction
	Bloque     []interface{}
}

func NewFuncionesDeclaracionP(lin int, col int, name string, parametros interfaces.Instruction, bloque []interface{}) FuncionesDeclaracionP {
	return FuncionesDeclaracionP{lin, col, name, parametros, bloque}
}

func (v FuncionesDeclaracionP) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	listavalores := list.New()
	v.Parametros.Ejecutar(ast, gen)
	for e := ast.Lista_Funciones_Var.Front(); e != nil; e = e.Next() {
		valor := e.Value.(environment.VariableFuncion)
		valor.Symbols.Scope = "Funcion-" + v.Name
		listavalores.PushBack(valor)
	}

	for e := ast.Lista_Funciones_Var.Front(); e != nil; e = e.Next() {
		valore := e.Value.(environment.VariableFuncion)

		// Compara con los elementos anteriores en la lista
		for i := ast.Lista_Funciones_Var.Front(); i != e; i = i.Next() {
			valori := i.Value.(environment.VariableFuncion)
			if valore.EI && valore.ExternoInterno != "_" {
				if valore.ExternoInterno == valori.ExternoInterno {
					Errores := environment.Errores{
						Descripcion: "Esta ingresando variables externas/internas iguales debe de cambiarles los nombres",
						Fila:        strconv.Itoa(v.Lin),
						Columna:     strconv.Itoa(v.Col),
						Tipo:        "Error Semántico",
						Ambito:      "Funcion-" + v.Name,
					}
					ast.ErroresHTML(Errores)
					ast.Lista_Funciones_Var.Init()
					return nil
				}
			}
		}

		// Compara con los elementos siguientes en la lista
		for i := e.Next(); i != nil; i = i.Next() {
			valori := i.Value.(environment.VariableFuncion)
			if valore.EI && valore.ExternoInterno != "_" {
				if valore.ExternoInterno == valori.ExternoInterno {
					Errores := environment.Errores{
						Descripcion: "Esta ingresando variables externas/internas iguales debe de cambiarles los nombres",
						Fila:        strconv.Itoa(v.Lin),
						Columna:     strconv.Itoa(v.Col),
						Tipo:        "Error Semántico",
						Ambito:      "Funcion-" + v.Name,
					}
					ast.ErroresHTML(Errores)
					ast.Lista_Funciones_Var.Init()
					return nil
				}
			}
		}
	}

	for e := ast.Lista_Funciones_Var.Front(); e != nil; e = e.Next() {
		valore := e.Value.(environment.VariableFuncion)

		// Compara con los elementos anteriores en la lista
		for i := ast.Lista_Funciones_Var.Front(); i != e; i = i.Next() {
			valori := i.Value.(environment.VariableFuncion)
			if valore.Name == valori.Name {
				Errores := environment.Errores{
					Descripcion: "Estas identificando los parametros con los mismos nombres tienes que cambiarlos",
					Fila:        strconv.Itoa(v.Lin),
					Columna:     strconv.Itoa(v.Col),
					Tipo:        "Error Semántico",
					Ambito:      "Funcion-" + v.Name,
				}
				ast.ErroresHTML(Errores)
				ast.Lista_Funciones_Var.Init()
				return nil
			}
		}

		// Compara con los elementos siguientes en la lista
		for i := e.Next(); i != nil; i = i.Next() {
			valori := i.Value.(environment.VariableFuncion)
			if valore.Name == valori.Name {
				Errores := environment.Errores{
					Descripcion: "Estas identificando los parametros con los mismos nombres tienes que cambiarlos",
					Fila:        strconv.Itoa(v.Lin),
					Columna:     strconv.Itoa(v.Col),
					Tipo:        "Error Semántico",
					Ambito:      "Funcion-" + v.Name,
				}
				ast.ErroresHTML(Errores)
				ast.Lista_Funciones_Var.Init()
				return nil
			}
		}
	}

	funcion := environment.Funcion{
		Lin:           v.Lin,
		Col:           v.Col,
		Nombre:        v.Name,
		IsReturn:      false,
		IsParame:      true,
		Parametros:    listavalores,
		CodigoFuncion: v.Bloque,
		Inicio:        ast.PosicionStack,
	}

	ast.GuardarFuncion(funcion)
	ast.Lista_Funciones_Var.Init()

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

	e := listavalores.Front()
	for i := 0; e != nil; i++ {
		valor := e.Value.(environment.VariableFuncion)
		symbol := environment.Symbol{
			Lin:         valor.Symbols.Lin,
			Col:         valor.Symbols.Col,
			Tipo:        valor.Symbols.Tipo,
			Scope:       ast.ObtenerAmbito(),
			TipoDato:    environment.VARIABLE,
			Posicion:    ast.PosicionStack,
			ValorInt:    valor.Symbols.ValorInt,
			ValorFloat:  valor.Symbols.ValorFloat,
			ValorString: valor.Symbols.ValorString,
		}
		newtem1 := gen.NewTemp()
		newtem2 := gen.NewTemp()
		gen.AddExpression(newtem1, "P", strconv.Itoa(i+1), "+")
		gen.AddGetStack(newtem2, "(int)"+newtem1)
		symbol.Valor = newtem2

		gen.AddSetStack(strconv.Itoa(symbol.Posicion), newtem2)
		Variable := environment.Variable{
			Name:        valor.Name,
			Symbols:     symbol,
			Mutable:     true,
			TipoSimbolo: "Variable",
		}

		ast.GuardarVariable(Variable)
		e = e.Next()
	}

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
