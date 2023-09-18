package datosprimitivos

import (
	"Backend/environment"
	"Backend/generator"
	"fmt"
	"strconv"
)

type VariableDeclaracionSinExp struct {
	Lin  int
	Col  int
	Name string
	Type environment.TipoExpresion
}

func NewVariableDeclaracionSinExp(lin int, col int, name string, tipo environment.TipoExpresion) VariableDeclaracionSinExp {
	return VariableDeclaracionSinExp{Lin: lin, Col: col, Name: name, Type: tipo}
}

func (v VariableDeclaracionSinExp) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	if ast.ObtenerAmbito() == "Global" {
		gen.MainCodeT()
	}
	gen.MainCodeT()
	symbol := environment.Symbol{
		Lin:      v.Lin,
		Col:      v.Col,
		Tipo:     v.Type,
		Scope:    ast.ObtenerAmbito(),
		TipoDato: environment.VARIABLE,
		Posicion: ast.Lista_Variables.Len(),
	}
	Variable := environment.Variable{
		Name:        v.Name,
		Symbols:     symbol,
		Mutable:     true,
		TipoSimbolo: "Variable",
	}
	safe := ast.GuardarVariable(Variable)
	var result environment.Value
	if safe {
		if v.Type == environment.INTEGER {
			gen.AddComment("Primitivo Int")
			result = environment.NewValue(fmt.Sprintf("%v", symbol.Valor), false, v.Type, false, false, false)
		} else if v.Type == environment.FLOAT {
			gen.AddComment("Primitivo Float")
			result = environment.NewValue(fmt.Sprintf("%v", symbol.Valor), false, v.Type, false, false, false)
		} else if v.Type == environment.STRING {
			gen.AddComment("Primitivo String")
			newTemp := gen.NewTemp()    //Cremos un nuevo temporal en esta va a ir el Head osea la asignacion de temp = head
			gen.AddAssign(newTemp, "H") //Creamos un nuevo puntero Head para que en este vayan los valores
			symbol.Valor = ""
			myString := symbol.Valor.(string) //Pasamos el valor a un string
			byteArray := []byte(myString)     //hacemos un vector del string anterior
			for _, asc := range byteArray {
				gen.AddSetHeap("(int)H", strconv.Itoa(int(asc))) //agregamos el ascii en el head Ej. heap[(int)H] = 111;
				//suma heap pointer
				gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion en el head Ej. H = H + 1; este es para poner un nuevo espacio
			}
			gen.AddSetHeap("(int)H", "-1")        // Agregamos el caracter de escape
			gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion igual que arriba de addexpression
			gen.AddBr()
			result = environment.NewValue(newTemp, true, v.Type, false, false, false)
		} else if v.Type == environment.CHARACTER {
			gen.AddComment("Primitivo Char")
			newTemp := gen.NewTemp()    //Cremos un nuevo temporal en esta va a ir el Head osea la asignacion de temp = head
			gen.AddAssign(newTemp, "H") //Creamos un nuevo puntero Head para que en este vayan los valores
			symbol.Valor = ""
			myString := symbol.Valor.(string) //Pasamos el valor a un string
			byteArray := []byte(myString)     //hacemos un vector del string anterior
			for _, asc := range byteArray {
				gen.AddSetHeap("(int)H", strconv.Itoa(int(asc))) //agregamos el ascii en el head Ej. heap[(int)H] = 111;
				//suma heap pointer
				gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion en el head Ej. H = H + 1; este es para poner un nuevo espacio
			}
			gen.AddSetHeap("(int)H", "-1")        // Agregamos el caracter de escape
			gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion igual que arriba de addexpression
			gen.AddBr()
			result = environment.NewValue(newTemp, true, v.Type, false, false, false)
		} else if v.Type == environment.BOOLEAN {
			gen.AddComment("Primitivo bool")
			trueLabel := gen.NewLabel()
			falseLabel := gen.NewLabel()
			symbol.Valor = false
			if symbol.Valor.(bool) {
				gen.AddGoto(trueLabel)
			} else {
				gen.AddGoto(falseLabel)
			}
			result = environment.NewValue("", false, environment.BOOLEAN, false, false, false)
			result.TrueLabel.PushBack(trueLabel)
			result.FalseLabel.PushBack(falseLabel)
		} else {
			gen.AddComment("Primitivo nil")
			result = environment.NewValue(fmt.Sprintf("%v", symbol.Valor), false, v.Type, false, false, false)
		}

		gen.AddComment("Declaracion de Variable")

		if v.Type == environment.BOOLEAN {
			//si no es temp (boolean)
			newLabel := gen.NewLabel()
			//add labels
			for e := result.TrueLabel.Front(); e != nil; e = e.Next() {
				gen.AddLabel(e.Value.(string))
			}
			gen.AddSetStack(strconv.Itoa(symbol.Posicion), "1")
			gen.AddGoto(newLabel)
			//add labels
			for e := result.FalseLabel.Front(); e != nil; e = e.Next() {
				gen.AddLabel(e.Value.(string))
			}
			gen.AddSetStack(strconv.Itoa(symbol.Posicion), "0")
			gen.AddGoto(newLabel)
			gen.AddLabel(newLabel)
			gen.AddBr()
		} else {
			//si es temp (num,string,etc)
			gen.AddSetStack(strconv.Itoa(symbol.Posicion), result.Value)
			gen.AddBr()
		}
	}

	gen.MainCodeT()
	return result
}
