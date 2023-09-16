package sentencias

type SentenciaForId struct {
	Lin   int
	Col   int
	Id    string
	Id2   string
	slice []interface{}
}

func NewSentenciaForId(lin int, col int, id string, id2 string, bloque []interface{}) SentenciaForId {
	return SentenciaForId{lin, col, id, id2, bloque}
}

/*
func (v SentenciaForId) Ejecutar(ast *environment.AST) interface{} {
	variables := ast.GetVariable(v.Id2)
	arreglos := ast.GetArreglo(v.Id2)
	if variables != nil {
		if variables.Symbols.Tipo == environment.STRING {
			ast.AumentarAmbito("For Cadena")
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
			var retornable int = 0
			var reexp environment.Symbol
			ast.GuardarVariable(Variable)

			for i := 0; i < len(Variable.Symbols.Valor.(string)); i++ {
				vari := ast.GetVariable(v.Id)
				cade := Variable.Symbols.Valor.(string)[i]
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
				Descripcion: "El id ingresado no a una variable tipo string",
				Fila:        strconv.Itoa(v.Lin),
				Columna:     strconv.Itoa(v.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
			return nil
		}
	}

	if arreglos != nil {

		ast.AumentarAmbito("For Cadena")
		symbol := environment.Symbol{
			Lin:   v.Lin,
			Col:   v.Col,
			Tipo:  arreglos.Symbols.Tipo,
			Valor: nil,
			Scope: ast.ObtenerAmbito(),
		}
		Variable := environment.Variable{
			Name:        v.Id,
			Symbols:     symbol,
			Mutable:     true,
			TipoSimbolo: "Variable",
		}
		var retornable int = 0
		var reexp environment.Symbol
		ast.GuardarVariable(Variable)

		for i := arreglos.Elements.Front(); i != nil; i = i.Next() {
			vari := ast.GetVariable(v.Id)
			cade := i.Value.(environment.Symbol).Valor
			symbol := environment.Symbol{
				Lin:   vari.Symbols.Lin,
				Col:   vari.Symbols.Col,
				Tipo:  arreglos.Symbols.Tipo,
				Valor: cade,
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
			Descripcion: "El id ingresado no pertenece ni a una variable ni aun vector",
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
