package sentencias

import (
	"Backend/interfaces"
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

/*
func (v SentenciaForRango) Ejecutar(ast *environment.AST) interface{} {
	var left, right environment.Symbol
	left = v.Left.Ejecutar(ast)
	right = v.Right.Ejecutar(ast)

	if left.Tipo != environment.INTEGER && right.Tipo != environment.INTEGER {
		Errores := environment.Errores{
			Descripcion: "El rango de for solo admite valores de tipo Int",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	if left.Valor.(int) > right.Valor.(int) {
		Errores := environment.Errores{
			Descripcion: "El rango de for tiene que ser el lado izquierdo (<=) derecho no mayor",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	symbol := environment.Symbol{
		Lin:   v.Lin,
		Col:   v.Col,
		Tipo:  environment.INTEGER,
		Valor: left.Valor,
		Scope: ast.ObtenerAmbito(),
	}
	Variable := environment.Variable{
		Name:        v.Id,
		Symbols:     symbol,
		Mutable:     true,
		TipoSimbolo: "Variable",
	}
	ast.AumentarAmbito("For Rango")
	var retornable int = 0
	var reexp environment.Symbol
	ast.GuardarVariable(Variable)
	for i := left.Valor.(int); i <= right.Valor.(int); i++ {
		for _, inst := range v.slice {
			if inst == nil {
				continue
			}
			instruction, ok := inst.(interfaces.Instruction)
			if !ok {
				continue
			}
			instruction.Ejecutar(ast)
			bvari := ast.GetVariable("Break")
			if bvari != nil {
				retornable = 1
				break
			}
			rvari := ast.GetVariable("Return")
			if rvari != nil {
				retornable = 2
				break
			}
			revari := ast.GetVariable("ReturnExp")
			if revari != nil {
				retornable = 3
				reexp = revari.Symbols
				break
			}
			cvari := ast.GetVariable("Continue")
			if cvari != nil {
				continue
			}
		}
		if retornable == 1 || retornable == 2 || retornable == 3 {
			break
		}
		vari := ast.GetVariable(v.Id)
		symbol := environment.Symbol{
			Lin:   vari.Symbols.Lin,
			Col:   vari.Symbols.Col,
			Tipo:  environment.INTEGER,
			Valor: vari.Symbols.Valor.(int) + 1,
		}
		Variable := environment.Variable{
			Name:        vari.Name,
			Symbols:     symbol,
			Mutable:     true,
			TipoSimbolo: "Variable",
		}
		ast.ActualizarVariable(&Variable)
	}
	ast.DisminuirAmbito()
	tamanio := ast.Pila_Variables.Len()
	if tamanio > 1 {
		if retornable == 2 {
			symbol := environment.Symbol{
				Lin:   v.Lin,
				Col:   v.Col,
				Tipo:  environment.BOOLEAN,
				Valor: true,
			}
			Variable := environment.Variable{
				Name:        "Return",
				Symbols:     symbol,
				Mutable:     false,
				TipoSimbolo: "Sentencia de Transferencia",
			}
			ast.GuardarVariable(Variable)
		}
		if retornable == 3 {
			symbol := environment.Symbol{
				Lin:   v.Lin,
				Col:   v.Col,
				Tipo:  reexp.Tipo,
				Valor: reexp.Valor,
				Scope: ast.ObtenerAmbito(),
			}
			Variable := environment.Variable{
				Name:        "ReturnExp",
				Symbols:     symbol,
				Mutable:     false,
				TipoSimbolo: "Sentencia de Transferencia",
			}
			ast.GuardarVariable(Variable)
		}
	}
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
	return nil
}
*/
