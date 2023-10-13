package sentencias

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
)

type SentenciaForRango struct {
	Lin   int
	Col   int
	Id    string
	Left  interfaces.Expression
	Right interfaces.Expression
	slice []interface{}
}

func NewSentenciaForRango(lin int, col int, id string, left interfaces.Expression, right interfaces.Expression, bloque []interface{}) SentenciaForRango {
	return SentenciaForRango{lin, col, id, left, right, bloque}
}

func (v SentenciaForRango) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	ambito := ast.ObtenerAmbito()
	ambitonuevo := "for" + "-" + ambito
	ast.AumentarAmbito(ambitonuevo)
	if !ast.IsMain(ambitonuevo) {
		gen.MainCodeT()
	}
	var left, right environment.Value
	left = v.Left.Ejecutar(ast, gen)
	if !ast.IsMain(ambitonuevo) {
		gen.MainCodeT()
	}
	right = v.Right.Ejecutar(ast, gen)
	if !ast.IsMain(ambitonuevo) {
		gen.MainCodeT()
	}

	if left.Type != environment.INTEGER || right.Type != environment.INTEGER {
		Errores := environment.Errores{
			Descripcion: "El rango de for solo admite valores de tipo Int",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		gen.MainCodeF()
		return nil
	}

	symbol := environment.Symbol{
		Lin:      v.Lin,
		Col:      v.Col,
		Tipo:     environment.INTEGER,
		Valor:    0,
		Scope:    ast.ObtenerAmbito(),
		Posicion: ast.PosicionStack,
		TipoDato: environment.VARIABLE,
		ValorInt: 0,
	}

	Variable := environment.Variable{
		Name:        v.Id,
		Symbols:     symbol,
		Mutable:     true,
		TipoSimbolo: "Variable",
	}
	ast.GuardarVariable(Variable)
	gen.AddSetStack(strconv.Itoa(symbol.Posicion), left.Value)
	gen.AddBr()

	var retornable int = 0
	var errorgeneral int = 0
	looptl := gen.NewLabel()
	exitla := gen.NewLabel()
	transferencia := environment.SentenciasdeTransferencia{
		Scope:  ambitonuevo,
		ETrue:  looptl,
		EFalse: exitla,
	}
	ast.Lista_Tranferencias.PushBack(transferencia)
	ast.Lista_For_Rango.PushBack(Variable)

	op1 := left
	op2 := right
	tlabel := gen.NewLabel()
	flabel := gen.NewLabel()
	newTemp1 := gen.NewTemp()
	newTemp2 := gen.NewTemp()

	if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
		gen.AddAssign(newTemp1, op1.Value)
	} else if !op1.IsTemp && op1.Value != "" {
		gen.AddAssign(newTemp1, op1.Value)
	} else {
		gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
	}

	if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
		gen.AddAssign(newTemp2, op2.Value)
	} else if !op2.IsTemp && op2.Value != "" {
		gen.AddAssign(newTemp2, op2.Value)
	} else {
		gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
	}

	gen.AddIf(newTemp1, newTemp2, "<=", tlabel)
	gen.AddGoto(flabel)
	gen.AddLabel(tlabel)
	gen.AddLabel(looptl)
	for i := left.IntValue; i <= right.IntValue; i++ {
		newTemp3 := gen.NewTemp()
		tlabel2 := gen.NewLabel()
		flabel2 := gen.NewLabel()
		gen.AddGetStack(newTemp3, strconv.Itoa(Variable.Symbols.Posicion))
		gen.AddIf(newTemp3, strconv.Itoa(i), "==", tlabel2)
		gen.AddGoto(flabel2)
		gen.AddLabel(tlabel2)
		Variable.Symbols.Valor = i
		Variable.Symbols.ValorInt = i
		ast.ActualizarVariable(&Variable)
		ambito = ast.ObtenerAmbito()
		ambitonuevo = "for" + "-" + ambito
		ast.AumentarAmbito(ambitonuevo)
		for _, inst := range v.slice {

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
			bvari := ast.GetVariable("Break")
			if bvari != nil {
				retornable = 1
				if ast.Lista_Tranferencias.Len() == 0 {
					errorgeneral = 1
				}
			}
			rvari := ast.GetVariable("Return")
			if rvari != nil {
				retornable = 2
				if ast.Lista_Tranferencias.Len() == 0 {
					errorgeneral = 1
				}
			}
			revari := ast.GetVariable("ReturnExp")
			if revari != nil {
				retornable = 3
				if ast.Lista_Tranferencias.Len() == 0 {
					errorgeneral = 1
				}
			}
			cvari := ast.GetVariable("Continue")
			if cvari != nil {
				if ast.Lista_Tranferencias.Len() == 0 {
					errorgeneral = 1
				}
			}
		}
		ast.DisminuirAmbito()
		newTemp4 := gen.NewTemp()
		gen.AddGetStack(newTemp4, strconv.Itoa(Variable.Symbols.Posicion))
		newTemp5 := gen.NewTemp()
		gen.AddExpression(newTemp5, newTemp4, "1", "+")
		gen.AddSetStack(strconv.Itoa(Variable.Symbols.Posicion), newTemp5)
		gen.AddGoto(looptl)
		gen.AddLabel(flabel2)
	}

	gen.AddGoto(exitla)
	gen.AddLabel(flabel)
	gen.AddPrintf("c", "(char)32")  // Agrega un espacio
	gen.AddPrintf("c", "(char)69")  // E
	gen.AddPrintf("c", "(char)114") // r
	gen.AddPrintf("c", "(char)114") // r
	gen.AddPrintf("c", "(char)111") // o
	gen.AddPrintf("c", "(char)114") // r
	gen.AddPrintf("c", "(char)32")  // Agrega un espacio
	gen.AddPrintf("c", "(char)82")  // R
	gen.AddPrintf("c", "(char)97")  // a
	gen.AddPrintf("c", "(char)110") // n
	gen.AddPrintf("c", "(char)103") // g
	gen.AddPrintf("c", "(char)111") // o
	gen.AddPrintf("c", "(char)32")  // Agrega un espacio
	gen.AddPrintf("c", "(char)70")  // F
	gen.AddPrintf("c", "(char)111") // o
	gen.AddPrintf("c", "(char)114") // r
	gen.AddGoto(exitla)
	gen.AddLabel(exitla)

	ast.DisminuirAmbito()
	tamanio := ast.Pila_Variables.Len()
	if tamanio == 1 && retornable == 3 {
		Errores := environment.Errores{
			Descripcion: "Estas retornando un valor fuera de una funcion",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
	}

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

	ast.Lista_Tranferencias.Remove(ast.Lista_Tranferencias.Back())
	ast.Lista_For_Rango.Remove(ast.Lista_For_Rango.Back())
	gen.MainCodeF()
	return nil
}
