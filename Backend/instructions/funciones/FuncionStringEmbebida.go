package funciones

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
)

type FuncionStringEmbebida struct {
	Op interfaces.Expression
}

func NewFuncionStringEmbebida(op interfaces.Expression) FuncionStringEmbebida {
	exp := FuncionStringEmbebida{Op: op}
	return exp
}

func (o FuncionStringEmbebida) Ejecutar(ast *environment.AST, gen *generator.Generator) environment.Value {
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	op := o.Op.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}

	if op.Type == environment.INTEGER {

		gen.AddComment("Int to String ")
		newTemp := gen.NewTemp()            //Cremos un nuevo temporal en esta va a ir el Head osea la asignacion de temp = head
		gen.AddAssign(newTemp, "H")         //Creamos un nuevo puntero Head para que en este vayan los valores
		byteArray := []byte(op.StringValue) //hacemos un vector del string anterior
		for _, asc := range byteArray {
			gen.AddSetHeap("(int)H", strconv.Itoa(int(asc))) //agregamos el ascii en el head Ej. heap[(int)H] = 111;
			//suma heap pointer
			gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion en el head Ej. H = H + 1; este es para poner un nuevo espacio
		}
		gen.AddSetHeap("(int)H", "-1")        // Agregamos el caracter de escape
		gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion igual que arriba de addexpression
		gen.AddBr()

		result := environment.NewValue(newTemp, true, environment.STRING, false, false, false, op.Val)
		result.IntValue = op.IntValue
		result.FloatValue = op.FloatValue
		result.StringValue = op.StringValue
		return result

	} else if op.Type == environment.FLOAT {
		gen.AddComment("Float to String ")
		newTemp := gen.NewTemp()            //Cremos un nuevo temporal en esta va a ir el Head osea la asignacion de temp = head
		gen.AddAssign(newTemp, "H")         //Creamos un nuevo puntero Head para que en este vayan los valores
		byteArray := []byte(op.StringValue) //hacemos un vector del string anterior
		for _, asc := range byteArray {
			gen.AddSetHeap("(int)H", strconv.Itoa(int(asc))) //agregamos el ascii en el head Ej. heap[(int)H] = 111;
			//suma heap pointer
			gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion en el head Ej. H = H + 1; este es para poner un nuevo espacio
		}
		gen.AddSetHeap("(int)H", "-1")        // Agregamos el caracter de escape
		gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion igual que arriba de addexpression
		gen.AddBr()

		result := environment.NewValue(newTemp, true, environment.STRING, false, false, false, op.Val)
		result.IntValue = op.IntValue
		result.FloatValue = op.FloatValue
		result.StringValue = op.StringValue
		return result
	} else if op.Type == environment.STRING {
		result := environment.NewValue(op.Value, true, environment.STRING, false, false, false, op.Val)
		result.IntValue = op.IntValue
		result.FloatValue = op.FloatValue
		result.StringValue = op.StringValue
		return result
	} else if op.Type == environment.BOOLEAN {

		etv := gen.NewLabel()
		etf := gen.NewLabel()
		newLabel := gen.NewLabel()
		newtmp := gen.NewTemp()
		gen.AddSetStack(strconv.Itoa(op.Val.Symbols.Posicion), op.Value)
		aux := ""
		for i := op.TrueLabel.Front(); i != nil; i = i.Next() {
			gen.AddLabel(i.Value.(string))
			aux = i.Value.(string)
		}
		if op.Val.FEti == "" && aux == "" {
			temp := gen.NewTemp()
			gen.AddGetStack(temp, strconv.Itoa(op.Val.Symbols.Posicion))
			gen.AddIf(temp, "1", "==", etv)
			gen.AddGoto(etf)
		}

		if op.Val.TEti != aux {
			gen.AddLabel(op.Val.TEti)
		}

		if op.Val.TEti == "" && aux == "" {
			gen.AddLabel(etv)
		}
		gen.AddAssign(newtmp, "H")
		gen.AddSetHeap("(int)H", "116")
		gen.AddExpression("H", "H", "1", "+")
		gen.AddSetHeap("(int)H", "114")
		gen.AddExpression("H", "H", "1", "+")
		gen.AddSetHeap("(int)H", "117")
		gen.AddExpression("H", "H", "1", "+")
		gen.AddSetHeap("(int)H", "101")
		gen.AddExpression("H", "H", "1", "+")
		gen.AddSetHeap("(int)H", "-1")
		gen.AddExpression("H", "H", "1", "+")
		gen.AddBr()

		gen.AddGoto(newLabel)
		aux2 := ""
		for i := op.FalseLabel.Front(); i != nil; i = i.Next() {
			gen.AddLabel(i.Value.(string))
			aux2 = i.Value.(string)
		}
		if op.Val.FEti != aux2 {
			gen.AddLabel(op.Val.FEti)
		}
		if op.Val.FEti == "" && aux2 == "" {
			gen.AddLabel(etf)
		}
		gen.AddAssign(newtmp, "H")
		gen.AddSetHeap("(int)H", "102")
		gen.AddExpression("H", "H", "1", "+")
		gen.AddSetHeap("(int)H", "97")
		gen.AddExpression("H", "H", "1", "+")
		gen.AddSetHeap("(int)H", "108")
		gen.AddExpression("H", "H", "1", "+")
		gen.AddSetHeap("(int)H", "115")
		gen.AddExpression("H", "H", "1", "+")
		gen.AddSetHeap("(int)H", "101")
		gen.AddExpression("H", "H", "1", "+")
		gen.AddSetHeap("(int)H", "-1")
		gen.AddExpression("H", "H", "1", "+")
		gen.AddLabel(newLabel)
		gen.AddBr()

		result := environment.NewValue(newtmp, true, environment.STRING, false, false, false, op.Val)
		return result
	} else if op.Type == environment.CHARACTER {
		result := environment.NewValue(op.Value, true, environment.STRING, false, false, false, op.Val)
		result.IntValue = op.IntValue
		result.FloatValue = op.FloatValue
		result.StringValue = op.StringValue
		return result
	} else {
		result := environment.NewValue("201314439", true, environment.NULL, false, false, false, environment.Variable{})
		return result
	}
}
