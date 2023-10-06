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
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	symbol := environment.Symbol{
		Lin:      v.Lin,
		Col:      v.Col,
		Tipo:     v.Type,
		Scope:    ast.ObtenerAmbito(),
		TipoDato: environment.VARIABLE,
		Posicion: ast.PosicionStack,
		Valor:    201314439,
		ValorInt: 0,
	}
	Variable := environment.Variable{
		Name:        v.Name,
		Symbols:     symbol,
		Mutable:     true,
		TipoSimbolo: "Variable",
	}

	var result environment.Value

	if v.Type == environment.INTEGER {
		gen.AddComment("Primitivo Int")
		result = environment.NewValue(fmt.Sprintf("%v", symbol.Valor), false, v.Type, false, false, false, Variable)
	} else if v.Type == environment.FLOAT {
		gen.AddComment("Primitivo Float")
		result = environment.NewValue(fmt.Sprintf("%v", symbol.Valor), false, v.Type, false, false, false, Variable)
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
		result = environment.NewValue(newTemp, true, v.Type, false, false, false, Variable)
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
		result = environment.NewValue(newTemp, true, v.Type, false, false, false, Variable)
	} else if v.Type == environment.BOOLEAN {
		gen.AddComment("Primitivo bool")
		aux := "0"
		result = environment.NewValue(aux, false, environment.BOOLEAN, false, false, false, environment.Variable{})
	} else {
		gen.AddComment("Primitivo nil")
		result = environment.NewValue(fmt.Sprintf("%v", symbol.Valor), false, v.Type, false, false, false, Variable)
	}

	gen.AddComment("Declaracion de Variable")

	if result.Type == environment.BOOLEAN {
		gen.AddSetStack(strconv.Itoa(symbol.Posicion), result.Value)
		gen.AddBr()
	} else {
		//si es temp (num,string,etc)
		gen.AddSetStack(strconv.Itoa(symbol.Posicion), result.Value)
		gen.AddBr()
	}

	ast.GuardarVariable(Variable)
	gen.MainCodeF()
	return result
}
