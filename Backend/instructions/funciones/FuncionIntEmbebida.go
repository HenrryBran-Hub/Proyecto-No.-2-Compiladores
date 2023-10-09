package funciones

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"fmt"
	"strconv"
)

type FuncionIntEmbebida struct {
	Op interfaces.Expression
}

func NewFuncionIntEmbebida(op interfaces.Expression) FuncionIntEmbebida {
	exp := FuncionIntEmbebida{Op: op}
	return exp
}

func (o FuncionIntEmbebida) Ejecutar(ast *environment.AST, gen *generator.Generator) environment.Value {
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	op := o.Op.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	if op.Type == environment.INTEGER {
		newtmp := gen.NewTemp()
		gen.AddGetStack(newtmp, strconv.Itoa(op.Val.Symbols.Posicion))
		gen.AddSetStack(strconv.Itoa(ast.PosicionStack), "(int)"+newtmp)
		op.Val.Symbols.Posicion = ast.PosicionStack
		result := environment.NewValue(newtmp, true, op.Type, false, false, false, op.Val)
		result.IntValue = op.IntValue
		gen.MainCodeF()
		return result
	} else if op.Type == environment.FLOAT {
		stringValue, _ := strconv.ParseFloat(fmt.Sprintf("%v", op.IntValue), 64)
		intValue := int(stringValue)
		op.Val.Symbols.Tipo = environment.INTEGER
		newtmp := gen.NewTemp()
		gen.AddGetStack(newtmp, strconv.Itoa(op.Val.Symbols.Posicion))
		gen.AddSetStack(strconv.Itoa(ast.PosicionStack), "(int)"+newtmp)
		op.Val.Symbols.Posicion = ast.PosicionStack
		result := environment.NewValue(newtmp, true, environment.INTEGER, false, false, false, op.Val)
		result.IntValue = intValue
		gen.MainCodeF()
		return result
	} else if op.Type == environment.STRING {
		gen.ConvertirCadenaANumero()
		//concat
		gen.AddComment("volviendo strings a int")
		envSize := strconv.Itoa(ast.PosicionStack)
		tmp3 := gen.NewTemp()
		tmp4 := gen.NewTemp()
		gen.AddExpression(tmp3, "P", envSize, "+")
		gen.AddExpression(tmp3, tmp3, "1", "+")
		gen.AddSetStack("(int)"+tmp3, op.Value)
		gen.AddExpression("P", "P", envSize, "+")
		gen.AddCall("ConvertCadtoNumber")
		gen.AddGetStack(tmp4, "(int)P")
		gen.AddExpression("P", "P", envSize, "-")
		gen.AddBr()
		//result = environment.NewValue(tmp2, true, environment.BOOLEAN, false, false, false, environment.Variable{})

		tlabel := gen.NewLabel()
		flabel := gen.NewLabel()
		exitlable := gen.NewLabel()

		gen.AddIf(tmp4, "201314439", "!=", tlabel)
		gen.AddGoto(flabel)
		gen.AddLabel(tlabel)

		tmp5 := gen.NewTemp()
		gen.AddAssign(tmp5, tmp4)
		gen.AddSetStack(strconv.Itoa(ast.PosicionStack), "(int)"+tmp5)

		gen.AddGoto(exitlable)
		gen.AddLabel(flabel)
		gen.AddAssign(tmp5, "0")
		gen.AddSetStack(strconv.Itoa(ast.PosicionStack), "(int)"+tmp5)

		gen.AddPrintf("c", "(char)69")  // E
		gen.AddPrintf("c", "(char)114") // r
		gen.AddPrintf("c", "(char)114") // r
		gen.AddPrintf("c", "(char)111") // o
		gen.AddPrintf("c", "(char)114") // r
		gen.AddPrintf("c", "(char)32")  // Agrega un espacio
		gen.AddPrintf("c", "(char)78")  // N
		gen.AddPrintf("c", "(char)111") // o
		gen.AddPrintf("c", "(char)32")  // Agrega un espacio
		gen.AddPrintf("c", "(char)69")  // E
		gen.AddPrintf("c", "(char)115") // s
		gen.AddPrintf("c", "(char)32")  // Agrega un espacio
		gen.AddPrintf("c", "(char)78")  // N
		gen.AddPrintf("c", "(char)117") // u
		gen.AddPrintf("c", "(char)109") // m
		gen.AddPrintf("c", "(char)101") // e
		gen.AddPrintf("c", "(char)114") // r
		gen.AddPrintf("c", "(char)111") // o
		gen.AddPrintf("c", "10")        // Agrega un espacio
		gen.AddGoto(exitlable)
		gen.AddLabel(exitlable)

		op.Val.Symbols.Posicion = ast.PosicionStack
		op.Val.Symbols.Tipo = environment.INTEGER
		result := environment.NewValue(tmp5, true, environment.INTEGER, false, false, false, op.Val)
		gen.MainCodeF()
		return result
	} else {
		Errores := environment.Errores{
			Descripcion: "No es posible convertir el dato a tipo Int, no es compatible el tipo de dato ingresado.(Character,Bool)",
			Fila:        strconv.Itoa(op.Val.Symbols.Lin),
			Columna:     strconv.Itoa(op.Val.Symbols.Col),
			Tipo:        "Error Semantico",
			Ambito:      op.Val.Symbols.Scope,
		}
		ast.ErroresHTML(Errores)
		result := environment.Value{}
		return result
	}
}
