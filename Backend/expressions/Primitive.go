package expressions

import (
	"Backend/environment"
	"Backend/generator"
	"fmt"
	"strconv"
)

type Primitive struct {
	Lin   int
	Col   int
	Valor interface{}
	Tipo  environment.TipoExpresion
}

func NewPrimitive(lin int, col int, valor interface{}, tipo environment.TipoExpresion) Primitive {
	exp := Primitive{lin, col, valor, tipo}
	return exp
}

func (p Primitive) Ejecutar(ast *environment.AST, gen *generator.Generator) environment.Value {
	var result environment.Value
	if p.Tipo == environment.INTEGER {
		gen.AddComment("Primitivo Int")
		result = environment.NewValue(fmt.Sprintf("%v", p.Valor), false, p.Tipo, false, false, false, environment.Variable{})
		prueba := fmt.Sprintf("%v", p.Valor)
		intValue, err := strconv.Atoi(prueba)
		if err != nil {
			result.IntValue = 0
		} else {
			result.IntValue = intValue
		}
		floatValue, err := strconv.ParseFloat(prueba, 64)
		if err != nil {
			result.FloatValue = 0
		} else {
			result.FloatValue = floatValue
		}
		result.StringValue = prueba
	} else if p.Tipo == environment.FLOAT {
		gen.AddComment("Primitivo Float")
		result = environment.NewValue(fmt.Sprintf("%v", p.Valor), false, p.Tipo, false, false, false, environment.Variable{})
		prueba := fmt.Sprintf("%v", p.Valor)
		intValue, err := strconv.Atoi(prueba)
		if err != nil {
			result.IntValue = 0
		} else {
			result.IntValue = intValue
		}
		floatValue, err := strconv.ParseFloat(prueba, 64)
		if err != nil {
			result.FloatValue = 0
		} else {
			result.FloatValue = floatValue
		}
		result.StringValue = prueba

	} else if p.Tipo == environment.STRING {
		gen.AddComment("Primitivo String")
		newTemp := gen.NewTemp()      //Cremos un nuevo temporal en esta va a ir el Head osea la asignacion de temp = head
		gen.AddAssign(newTemp, "H")   //Creamos un nuevo puntero Head para que en este vayan los valores
		myString := p.Valor.(string)  //Pasamos el valor a un string
		byteArray := []byte(myString) //hacemos un vector del string anterior
		for _, asc := range byteArray {
			gen.AddSetHeap("(int)H", strconv.Itoa(int(asc))) //agregamos el ascii en el head Ej. heap[(int)H] = 111;
			//suma heap pointer
			gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion en el head Ej. H = H + 1; este es para poner un nuevo espacio
		}
		gen.AddSetHeap("(int)H", "-1")        // Agregamos el caracter de escape
		gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion igual que arriba de addexpression
		gen.AddBr()
		result = environment.NewValue(newTemp, true, p.Tipo, false, false, false, environment.Variable{})
		prueba := fmt.Sprintf("%v", p.Valor)
		intValue, err := strconv.Atoi(prueba)
		if err != nil {
			result.IntValue = 0
		} else {
			result.IntValue = intValue
		}
		floatValue, err := strconv.ParseFloat(prueba, 64)
		if err != nil {
			result.FloatValue = 0
		} else {
			result.FloatValue = floatValue
		}
		result.StringValue = prueba
	} else if p.Tipo == environment.CHARACTER {
		gen.AddComment("Primitivo Char")
		newTemp := gen.NewTemp()      //Cremos un nuevo temporal en esta va a ir el Head osea la asignacion de temp = head
		gen.AddAssign(newTemp, "H")   //Creamos un nuevo puntero Head para que en este vayan los valores
		myString := p.Valor.(string)  //Pasamos el valor a un string
		byteArray := []byte(myString) //hacemos un vector del string anterior
		for _, asc := range byteArray {
			gen.AddSetHeap("(int)H", strconv.Itoa(int(asc))) //agregamos el ascii en el head Ej. heap[(int)H] = 111;
			//suma heap pointer
			gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion en el head Ej. H = H + 1; este es para poner un nuevo espacio
		}
		gen.AddSetHeap("(int)H", "-1")        // Agregamos el caracter de escape
		gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion igual que arriba de addexpression
		gen.AddBr()
		result = environment.NewValue(newTemp, true, p.Tipo, false, false, false, environment.Variable{})
		prueba := fmt.Sprintf("%v", p.Valor)
		intValue, err := strconv.Atoi(prueba)
		if err != nil {
			result.IntValue = 0
		} else {
			result.IntValue = intValue
		}
		floatValue, err := strconv.ParseFloat(prueba, 64)
		if err != nil {
			result.FloatValue = 0
		} else {
			result.FloatValue = floatValue
		}
		result.StringValue = prueba
	} else if p.Tipo == environment.BOOLEAN {
		gen.AddComment("Primitivo bool")
		aux := ""
		if p.Valor.(bool) {
			aux = "1"
			result.IntValue = 1
			result.StringValue = "true"
		} else {
			aux = "0"
			result.IntValue = 0
			result.StringValue = "false"
		}
		result = environment.NewValue(aux, false, environment.BOOLEAN, false, false, false, environment.Variable{})
	} else {
		gen.AddComment("Primitivo nil")
		result = environment.NewValue("201314439", false, environment.NULL, false, false, false, environment.Variable{})
		result.IntValue = 0
		result.FloatValue = 0.0
		result.StringValue = "201314439"
	}
	return result
}
