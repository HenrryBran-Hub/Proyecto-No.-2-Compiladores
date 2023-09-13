package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
	"container/list"
	"strconv"
)

type StructVariable struct {
	Lin        int
	Col        int
	IdVarLet   string
	IdVariable string
	IdStruct   string
	Op         interfaces.Instruction
	Estado     bool
}

func NewStruckVariable(lin, col int, varlet, variable, strucname string, op interfaces.Instruction, estado bool) StructVariable {
	return StructVariable{lin, col, varlet, variable, strucname, op, estado}
}

func NewStruckVariable2(lin, col int, varlet, variable, strucname string, estado bool) StructVariable {
	return StructVariable{Lin: lin, Col: col, IdVarLet: varlet, IdVariable: variable, IdStruct: strucname, Estado: estado}
}

func (v StructVariable) Ejecutar(ast *environment.AST) interface{} {
	if v.Estado == true {
		v.Op.Ejecutar(ast)
	}

	existvarstruc := ast.GetVariableStruc(v.IdVariable)
	if existvarstruc != nil {
		Errores := environment.Errores{
			Descripcion: "Se ya existe el struc con este nombre. " + v.IdVariable,
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		ast.ListaParametrosStruct.Init()
		return nil
	}

	existstruck := ast.GetStruc(v.IdStruct)
	if existstruck == nil {
		Errores := environment.Errores{
			Descripcion: "No existe el struct que estas intentando definir. " + v.IdStruct,
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		ast.ListaParametrosStruct.Init()
		return nil
	}

	var variables = list.New()
	var funciones = list.New()

	e1 := existstruck.Variables.Front()
	e2 := ast.ListaParametrosStruct.Front()

	for e1 != nil && e2 != nil {
		valor1 := e1.Value.(environment.Variable)
		valor2 := e2.Value.(environment.Parametrostruct)

		symbol := environment.Symbol{
			Lin:   valor1.Symbols.Lin,
			Col:   valor1.Symbols.Col,
			Tipo:  valor1.Symbols.Tipo,
			Valor: valor2.Symbolo.Valor,
			Scope: "Struct-" + v.IdStruct + "-" + v.IdVariable,
		}

		if valor1.Symbols.Tipo != valor2.Symbolo.Tipo {
			Errores := environment.Errores{
				Descripcion: "El tipo ingresado a la funcion no es el mismo que esta definido en la funcion",
				Fila:        strconv.Itoa(v.Lin),
				Columna:     strconv.Itoa(v.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
			ast.ListaParametrosStruct.Init()
			return nil
		}

		if valor1.Mutable == false {
			Errores := environment.Errores{
				Descripcion: "El valor de la variable que quieres cambiar o ingresar no es mutable",
				Fila:        strconv.Itoa(v.Lin),
				Columna:     strconv.Itoa(v.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
			ast.ListaParametrosStruct.Init()
			symbol.Valor = valor1.Symbols.Valor
		}

		Variable := environment.Variable{
			Name:        valor1.Name,
			Symbols:     symbol,
			Mutable:     valor1.Mutable,
			TipoSimbolo: "Variable",
		}

		variables.PushBack(Variable)

		e1 = e1.Next()
		e2 = e2.Next()
	}

	for e := existstruck.Funciones.Front(); e != nil; e = e.Next() {
		val := e.Value.(environment.Funcion)
		funcion := environment.Funcion{
			Lin:           val.Lin,
			Col:           val.Col,
			Nombre:        val.Nombre,
			IsReturn:      val.IsReturn,
			IsParame:      val.IsParame,
			Tipo:          val.Tipo,
			Retorno:       val.Retorno,
			Parametros:    val.Parametros,
			CodigoFuncion: val.CodigoFuncion,
			Mutating:      val.Mutating,
		}
		funciones.PushBack(funcion)
	}

	copiastruck := environment.Struc{
		Lin:       v.Lin,
		Col:       v.Col,
		Nombre:    v.IdStruct,
		Variables: variables,
		Funciones: funciones,
	}

	copiasvarstruc := environment.VariableStruct{
		Lin:      v.Lin,
		Col:      v.Col,
		Nombre:   v.IdVariable,
		Strukt:   copiastruck,
		Mutating: true,
		Scope:    ast.ObtenerAmbito(),
	}

	ast.GuardarVariableStruc(copiasvarstruc)
	ast.ListaParametrosStruct.Init()
	return nil
}
