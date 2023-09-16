package sentencias

import (
	"Backend/interfaces"
)

type SentenciaForCadena struct {
	Lin   int
	Col   int
	Id    string
	Expre interfaces.Expression
	slice []interface{}
}

func NewSentenciaForCadena(lin int, col int, id string, expre interfaces.Expression, bloque []interface{}) SentenciaForCadena {
	return SentenciaForCadena{lin, col, id, expre, bloque}
}

/*
func (v SentenciaForCadena) Ejecutar(ast *environment.AST) interface{} {
	var cadena environment.Symbol
	cadena = v.Expre.Ejecutar(ast)

	if cadena.Tipo == environment.VECTOR || cadena.Tipo == environment.STRING {
		symbol := environment.Symbol{
			Lin:   v.Lin,
			Col:   v.Col,
			Tipo:  environment.CHARACTER,
			Valor: nil,
			Scope: ast.ObtenerAmbito(),
		}
		Variable := environment.Variable{
			Name:        v.Id,
			Symbols:     symbol,
			Mutable:     true,
			TipoSimbolo: "Variable",
		}
		ast.AumentarAmbito("For Cadena")
		var retornable int = 0
		var reexp environment.Symbol
		ast.GuardarVariable(Variable)

		if cadena.Tipo == environment.STRING {
			for i := 0; i < len(cadena.Valor.(string)); i++ {
				vari := ast.GetVariable(v.Id)
				cade := cadena.Valor.(string)[i]
				character := fmt.Sprintf("%c", cade)
				symbol := environment.Symbol{
					Lin:   vari.Symbols.Lin,
					Col:   vari.Symbols.Col,
					Tipo:  environment.CHARACTER,
					Valor: character,
					Scope: ast.ObtenerAmbito(),
				}
				Variable := environment.Variable{
					Name:        vari.Name,
					Symbols:     symbol,
					Mutable:     true,
					TipoSimbolo: "Variable",
				}
				ast.ActualizarVariable(&Variable)
				Variable.Mutable = false
				ast.ActualizarVariable(&Variable)
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
			}
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
					Scope: ast.ObtenerAmbito(),
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
				Ambito:      "For Cadena",
			}
			ast.ErroresHTML(Errores)
		}
	} else {
		Errores := environment.Errores{
			Descripcion: "El for que se esta ejecutando solo funciona con arreglos unidimensionales",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	return nil
}
*/
