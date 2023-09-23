package generator

import (
	"container/list"
	"fmt"
)

type Generator struct {
	Temporal         int
	Label            int
	Code             *list.List
	FinalCode        *list.List
	Natives          *list.List
	FuncCode         *list.List
	TempList         *list.List
	PrintStringFlag  bool
	ConcatStringFlag bool
	BreakLabel       string
	ContinueLabel    string
	MainCode         bool
	MenorFlag        bool
	MayorFlag        bool
	MenorIgFlag      bool
	MayorIgFlag      bool
	DifFlag          bool
	IgFlag           bool
	AndFlag          bool
	OrFlag           bool
}

func NewGenerator() Generator {
	generator := Generator{
		Temporal:         0,
		Label:            0,
		Code:             list.New(),
		FinalCode:        list.New(),
		Natives:          list.New(),
		FuncCode:         list.New(),
		TempList:         list.New(),
		PrintStringFlag:  true,
		ConcatStringFlag: true,
		BreakLabel:       "",
		ContinueLabel:    "",
		MainCode:         false,
		MenorFlag:        true,
		MayorFlag:        true,
		MenorIgFlag:      true,
		MayorIgFlag:      true,
		DifFlag:          true,
		IgFlag:           true,
		AndFlag:          true,
		OrFlag:           true,
	}
	return generator
}

func (g Generator) GetCode() *list.List {
	return g.Code
}

func (g Generator) GetFinalCode() *list.List {
	return g.FinalCode
}

func (g Generator) GetTemps() *list.List {
	return g.TempList
}

// PushBack break lvl
func (g *Generator) AddBreak(lvl string) {
	g.BreakLabel = lvl
}

// PushBack continue lvl
func (g *Generator) AddContinue(lvl string) {
	g.ContinueLabel = lvl
}

// Generar un nuevo temporal
func (g *Generator) NewTemp() string {
	temp := "t" + fmt.Sprintf("%v", g.Temporal)
	g.Temporal = g.Temporal + 1
	//Lo guardamos para declararlo
	g.TempList.PushBack(temp)
	return temp
}

// Generador de Label
func (g *Generator) NewLabel() string {
	temp := g.Label
	g.Label = g.Label + 1
	return "L" + fmt.Sprintf("%v", temp)
}

// Obtener maincode true
func (g *Generator) MainCodeT() {
	g.MainCode = true
}

// Cambiar maincode false
func (g *Generator) MainCodeF() {
	g.MainCode = false
}

// Añade Label al codigo
func (g *Generator) AddLabel(Label string) {
	if g.MainCode {
		g.Code.PushBack(Label + ":\n")
	} else {
		g.FuncCode.PushBack(Label + ":\n")
	}
}

func (g *Generator) AddIf(left string, right string, operator string, Label string) {
	if g.MainCode {
		g.Code.PushBack("if(" + left + " " + operator + " " + right + ") goto " + Label + ";\n")
	} else {
		g.FuncCode.PushBack("if(" + left + " " + operator + " " + right + ") goto " + Label + ";\n")
	}
}

func (g *Generator) AddGoto(Label string) {
	if g.MainCode {
		g.Code.PushBack("goto " + Label + ";\n")
	} else {
		g.FuncCode.PushBack("goto " + Label + ";\n")
	}
}

func (g *Generator) AddExpression(target string, left string, right string, operator string) {
	if g.MainCode {
		g.Code.PushBack(target + " = " + left + " " + operator + " " + right + ";\n")
	} else {
		g.FuncCode.PushBack(target + " = " + left + " " + operator + " " + right + ";\n")
	}
}

func (g *Generator) AddAssign(target, val string) {
	if g.MainCode {
		g.Code.PushBack(target + " = " + val + ";\n")
	} else {
		g.FuncCode.PushBack(target + " = " + val + ";\n")
	}
}

func (g *Generator) AddPrintf(typePrint string, value string) {
	if g.MainCode {
		g.Code.PushBack("printf(\"%" + typePrint + "\", " + value + ");\n")
	} else {
		g.FuncCode.PushBack("printf(\"%" + typePrint + "\", " + value + ");\n")
	}
}

func (g *Generator) AddSetHeap(index string, value string) {
	if g.MainCode {
		g.Code.PushBack("heap[" + index + "] = " + value + ";\n")
	} else {
		g.FuncCode.PushBack("heap[" + index + "] = " + value + ";\n")
	}
}

func (g *Generator) AddGetHeap(target string, index string) {
	if g.MainCode {
		g.Code.PushBack(target + " = heap[" + index + "];\n")
	} else {
		g.FuncCode.PushBack(target + " = heap[" + index + "];\n")
	}
}

func (g *Generator) AddSetStack(index string, value string) {
	if g.MainCode {
		g.Code.PushBack("stack[" + index + "] = " + value + ";\n")
	} else {
		g.FuncCode.PushBack("stack[" + index + "] = " + value + ";\n")
	}
}

func (g *Generator) AddGetStack(target string, index string) {
	if g.MainCode {
		g.Code.PushBack(target + " = stack[" + index + "];\n")
	} else {
		g.FuncCode.PushBack(target + " = stack[" + index + "];\n")
	}
}

func (g *Generator) AddCall(target string) {
	if g.MainCode {
		g.Code.PushBack(target + "();\n")
	} else {
		g.FuncCode.PushBack(target + "();\n")
	}
}

func (g *Generator) AddBr() {
	if g.MainCode {
		g.Code.PushBack("\n")
	} else {
		g.FuncCode.PushBack("\n")
	}

}

func (g *Generator) AddComment(target string) {
	if g.MainCode {
		g.Code.PushBack("//" + target + "\n")
	} else {
		g.FuncCode.PushBack("//" + target + "\n")
	}
}

func (g *Generator) AddTittle(target string) {
	if g.MainCode {
		g.Code.PushBack("void " + target + "() {\n")
	} else {
		g.FuncCode.PushBack("void " + target + "() {\n")
	}
}

func (g *Generator) AddEnd() {
	if g.MainCode {
		g.Code.PushBack("\treturn;\n")
		g.Code.PushBack("}\n\n")
	} else {
		g.FuncCode.PushBack("\treturn;\n")
		g.FuncCode.PushBack("}\n\n")
	}
}

// agregar headers
func (g *Generator) GenerateFinalCode() {
	//****************** PushBack head
	g.FinalCode.PushBack("/*------HEADER------*/\n")
	g.FinalCode.PushBack("#include <stdio.h>\n")
	g.FinalCode.PushBack("double heap[26121992];\n")
	g.FinalCode.PushBack("double stack[26121992];\n")
	g.FinalCode.PushBack("double P;\n")
	g.FinalCode.PushBack("double H;\n")
	//****************** PushBack temporal declaration
	if g.GetTemps().Len() > 0 {
		g.FinalCode.PushBack("double ")
		tmpDec := fmt.Sprintf("%v", g.GetTemps().Front().Value)
		g.GetTemps().Remove(g.GetTemps().Front())
		for e := g.GetTemps().Front(); e != nil; e = e.Next() {
			valor := e.Value
			tmpDec += ", "
			tmpDec += fmt.Sprintf("%v", valor)
		}
		tmpDec += ";\n\n"
		g.FinalCode.PushBack(tmpDec)
	}
	//****************** PushBack natives functions
	if g.Natives.Len() > 0 {
		g.FinalCode.PushBack("/*------NATIVES------*/\n")
		for e := g.Natives.Front(); e != nil; e = e.Next() {
			g.FinalCode.PushBack(e.Value.(string))
		}
	}
	//****************** PushBack functions
	if g.FuncCode.Len() > 0 {
		g.FinalCode.PushBack("/*------FUNCTIONS------*/\n")
		for e := g.FuncCode.Front(); e != nil; e = e.Next() {
			g.FinalCode.PushBack(e.Value.(string))
		}
	}
	//****************** PushBack main
	g.FinalCode.PushBack("/*------MAIN------*/\n")
	g.FinalCode.PushBack("void main() {\n")
	g.FinalCode.PushBack("\tP = 0; H = 0;\n\n")
	for e := g.Code.Front(); e != nil; e = e.Next() {
		s := e.Value
		g.FinalCode.PushBack("\t" + s.(string))
	}
	g.FinalCode.PushBack("\n\treturn;\n}\n")
}

func (g *Generator) GeneratePrintString() {
	if g.PrintStringFlag {
		//generando temporales y etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newTemp3 := g.NewTemp()
		newLvl1 := g.NewLabel()
		newLvl2 := g.NewLabel()
		//se genera la funcion printstring
		g.Natives.PushBack("void printString() {\n")
		g.Natives.PushBack("\t" + newTemp1 + " = P + 1;\n")
		g.Natives.PushBack("\t" + newTemp2 + " = stack[(int)" + newTemp1 + "];\n")
		g.Natives.PushBack("\t" + newLvl2 + ":\n")
		g.Natives.PushBack("\t" + newTemp3 + " = heap[(int)" + newTemp2 + "];\n")
		g.Natives.PushBack("\tif(" + newTemp3 + " == -1) goto " + newLvl1 + ";\n")
		g.Natives.PushBack("\tprintf(\"%c\", (char)" + newTemp3 + ");\n")
		g.Natives.PushBack("\t" + newTemp2 + " = " + newTemp2 + " + 1;\n")
		g.Natives.PushBack("\tgoto " + newLvl2 + ";\n")
		g.Natives.PushBack("\t" + newLvl1 + ":\n")
		g.Natives.PushBack("\treturn;\n")
		g.Natives.PushBack("}\n\n")
		g.PrintStringFlag = false
	}
}

func (g *Generator) GenerateConcatString() {
	if g.ConcatStringFlag {
		//generando temporales y etiquetas
		tmp1 := g.NewTemp()
		tmp2 := g.NewTemp()
		tmp3 := g.NewTemp()
		tmp4 := g.NewTemp()
		tmp5 := g.NewTemp()
		lvl1 := g.NewLabel()
		lvl2 := g.NewLabel()
		lvl3 := g.NewLabel()
		lvl4 := g.NewLabel()
		//se genera la funcion printstring
		g.Natives.PushBack("void concatString() {\n")
		g.Natives.PushBack("\t" + tmp1 + " = H;" + "\n")
		g.Natives.PushBack("\t" + tmp2 + " = P + 1;" + "\n")
		g.Natives.PushBack("\t" + tmp4 + " = stack[(int)" + tmp2 + "];" + "\n")
		g.Natives.PushBack("\t" + tmp3 + " = P + 2;" + "\n")
		g.Natives.PushBack("\t" + lvl2 + ":" + "\n")
		g.Natives.PushBack("\t" + tmp5 + " = heap[(int)" + tmp4 + "];" + "\n")
		g.Natives.PushBack("\t" + "if(" + tmp5 + " == -1) goto " + lvl3 + ";" + "\n")
		g.Natives.PushBack("\t" + "heap[(int)H] = " + tmp5 + ";" + "\n")
		g.Natives.PushBack("\t" + "H = H + 1;" + "\n")
		g.Natives.PushBack("\t" + tmp4 + " = " + tmp4 + " + 1;" + "\n")
		g.Natives.PushBack("\t" + "goto " + lvl2 + ";" + "\n")
		g.Natives.PushBack("\t" + lvl3 + ":" + "\n")
		g.Natives.PushBack("\t" + tmp4 + " = stack[(int)" + tmp3 + "];" + "\n")
		g.Natives.PushBack("\t" + lvl4 + ":" + "\n")
		g.Natives.PushBack("\t" + tmp5 + " = heap[(int)" + tmp4 + "];" + "\n")
		g.Natives.PushBack("\t" + "if(" + tmp5 + " == -1) goto " + lvl1 + ";" + "\n")
		g.Natives.PushBack("\t" + "heap[(int)H] = " + tmp5 + ";" + "\n")
		g.Natives.PushBack("\t" + "H = H + 1;" + "\n")
		g.Natives.PushBack("\t" + tmp4 + " = " + tmp4 + " + 1;" + "\n")
		g.Natives.PushBack("\t" + "goto " + lvl4 + ";" + "\n")
		g.Natives.PushBack("\t" + lvl1 + ":" + "\n")
		g.Natives.PushBack("\t" + "heap[(int)H] = -1;" + "\n")
		g.Natives.PushBack("\t" + "H = H + 1;" + "\n")
		g.Natives.PushBack("\t" + "stack[(int)P] = " + tmp1 + ";" + "\n")
		g.Natives.PushBack("\treturn;\n")
		g.Natives.PushBack("}\n\n")
		g.ConcatStringFlag = false
	}
}

func (g *Generator) GenerateCompareMaryorString() {
	if g.MayorFlag {
		//generando temporales y etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newTemp3 := g.NewTemp()
		newTemp4 := g.NewTemp()
		newTemp5 := g.NewTemp()
		newTemp6 := g.NewTemp()
		newLvl1 := g.NewLabel()
		newLvl2 := g.NewLabel()
		newLvl3 := g.NewLabel()
		newLvl4 := g.NewLabel()
		newLvl5 := g.NewLabel()
		newLvl6 := g.NewLabel()
		newLvl7 := g.NewLabel()
		newLvl8 := g.NewLabel()
		newLvl9 := g.NewLabel()
		newLvl10 := g.NewLabel()
		newLvl11 := g.NewLabel()
		newLvl12 := g.NewLabel()
		newLvl13 := g.NewLabel()
		newLvl14 := g.NewLabel()
		newLvl15 := g.NewLabel()
		newLvl16 := g.NewLabel()
		newLvl17 := g.NewLabel()
		newLvlend := g.NewLabel()
		//se genera la funcion printstring
		g.Natives.PushBack("void comparemayorString() {\n")
		g.Natives.PushBack("\t" + newTemp1 + " = P + 1;\n")
		g.Natives.PushBack("\t" + newTemp2 + " = stack[(int)" + newTemp1 + "];\n")
		g.Natives.PushBack("\t" + newTemp3 + " = P + 2;\n")
		g.Natives.PushBack("\t" + newTemp4 + " = stack[(int)" + newTemp3 + "];\n")
		g.Natives.PushBack("\t" + newTemp5 + " = heap[(int)" + newTemp2 + "];\n")
		g.Natives.PushBack("\t" + newTemp6 + " = heap[(int)" + newTemp4 + "];\n")
		g.Natives.PushBack("\tif(" + newTemp5 + " == -1) goto " + newLvl2 + ";\n")
		g.Natives.PushBack("\tgoto " + newLvl3 + ";\n")
		g.Natives.PushBack("\t" + newLvl2 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp6 + " == -1) goto " + newLvl4 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl5 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl4 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t" + newLvl5 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t" + newLvl3 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp6 + " == -1) goto " + newLvl6 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl7 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl6 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t" + newLvl7 + ":\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl1 + ";\n")
		g.Natives.PushBack("\t" + newLvl1 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp5 + " >= " + newTemp6 + ") goto " + newLvl8 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl9 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl8 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\t" + newTemp2 + " = " + newTemp2 + " + 1;\n")
		g.Natives.PushBack("\t\t\t" + newTemp4 + " = " + newTemp4 + " + 1;\n")
		g.Natives.PushBack("\t\t\t" + newTemp5 + " = heap[(int)" + newTemp2 + "];\n")
		g.Natives.PushBack("\t\t\t" + newTemp6 + " = heap[(int)" + newTemp4 + "];\n")
		g.Natives.PushBack("\t\t\tif(" + newTemp5 + " == -1) goto " + newLvl10 + ";\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl11 + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl10 + ":\n")
		g.Natives.PushBack("\t\t\t\tif(" + newTemp6 + " == -1) goto " + newLvl12 + ";\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvl13 + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl12 + ":\n")
		g.Natives.PushBack("\t\t\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl13 + ":\n")
		g.Natives.PushBack("\t\t\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl11 + ":\n")
		g.Natives.PushBack("\t\t\t\tif(" + newTemp6 + " == -1) goto " + newLvl14 + ";\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvl15 + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl14 + ":\n")
		g.Natives.PushBack("\t\t\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl15 + ":\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvl1 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl9 + ":\n")
		g.Natives.PushBack("\t\t\t" + newTemp2 + " = " + newTemp2 + " - 1;\n")
		g.Natives.PushBack("\t\t\t" + newTemp4 + " = " + newTemp4 + " - 1;\n")
		g.Natives.PushBack("\t\t\t" + newTemp5 + " = heap[(int)" + newTemp2 + "];\n")
		g.Natives.PushBack("\t\t\t" + newTemp6 + " = heap[(int)" + newTemp4 + "];\n")
		g.Natives.PushBack("\t\t\tif(" + newTemp5 + " >= " + newTemp6 + ") goto " + newLvl16 + ";\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl17 + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl16 + ":\n")
		g.Natives.PushBack("\t\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl17 + ":\n")
		g.Natives.PushBack("\t\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t" + newLvlend + ":\n")
		g.Natives.PushBack("\t\treturn;\n")
		g.Natives.PushBack("}\n\n")
		g.MayorFlag = false
	}
}

func (g *Generator) GenerateCompareMenorString() {
	if g.MenorFlag {
		//generando temporales y etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newTemp3 := g.NewTemp()
		newTemp4 := g.NewTemp()
		newTemp5 := g.NewTemp()
		newTemp6 := g.NewTemp()
		newLvl1 := g.NewLabel()
		newLvl2 := g.NewLabel()
		newLvl3 := g.NewLabel()
		newLvl4 := g.NewLabel()
		newLvl5 := g.NewLabel()
		newLvl6 := g.NewLabel()
		newLvl7 := g.NewLabel()
		newLvl8 := g.NewLabel()
		newLvl9 := g.NewLabel()
		newLvl10 := g.NewLabel()
		newLvl11 := g.NewLabel()
		newLvl12 := g.NewLabel()
		newLvl13 := g.NewLabel()
		newLvl14 := g.NewLabel()
		newLvl15 := g.NewLabel()
		newLvl16 := g.NewLabel()
		newLvl17 := g.NewLabel()
		newLvlend := g.NewLabel()
		//se genera la funcion printstring
		g.Natives.PushBack("void comparemenorString() {\n")
		g.Natives.PushBack("\t" + newTemp1 + " = P + 1;\n")
		g.Natives.PushBack("\t" + newTemp2 + " = stack[(int)" + newTemp1 + "];\n")
		g.Natives.PushBack("\t" + newTemp3 + " = P + 2;\n")
		g.Natives.PushBack("\t" + newTemp4 + " = stack[(int)" + newTemp3 + "];\n")
		g.Natives.PushBack("\t" + newTemp5 + " = heap[(int)" + newTemp2 + "];\n")
		g.Natives.PushBack("\t" + newTemp6 + " = heap[(int)" + newTemp4 + "];\n")
		g.Natives.PushBack("\tif(" + newTemp5 + " == -1) goto " + newLvl2 + ";\n")
		g.Natives.PushBack("\tgoto " + newLvl3 + ";\n")
		g.Natives.PushBack("\t" + newLvl2 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp6 + " == -1) goto " + newLvl4 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl5 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl4 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t" + newLvl5 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t" + newLvl3 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp6 + " == -1) goto " + newLvl6 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl7 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl6 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t" + newLvl7 + ":\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl1 + ";\n")
		g.Natives.PushBack("\t" + newLvl1 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp5 + " <= " + newTemp6 + ") goto " + newLvl8 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl9 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl8 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\t" + newTemp2 + " = " + newTemp2 + " + 1;\n")
		g.Natives.PushBack("\t\t\t" + newTemp4 + " = " + newTemp4 + " + 1;\n")
		g.Natives.PushBack("\t\t\t" + newTemp5 + " = heap[(int)" + newTemp2 + "];\n")
		g.Natives.PushBack("\t\t\t" + newTemp6 + " = heap[(int)" + newTemp4 + "];\n")
		g.Natives.PushBack("\t\t\tif(" + newTemp5 + " == -1) goto " + newLvl10 + ";\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl11 + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl10 + ":\n")
		g.Natives.PushBack("\t\t\t\tif(" + newTemp6 + " == -1) goto " + newLvl12 + ";\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvl13 + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl12 + ":\n")
		g.Natives.PushBack("\t\t\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl13 + ":\n")
		g.Natives.PushBack("\t\t\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl11 + ":\n")
		g.Natives.PushBack("\t\t\t\tif(" + newTemp6 + " == -1) goto " + newLvl14 + ";\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvl15 + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl14 + ":\n")
		g.Natives.PushBack("\t\t\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl15 + ":\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvl1 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl9 + ":\n")
		g.Natives.PushBack("\t\t\t" + newTemp2 + " = " + newTemp2 + " - 1;\n")
		g.Natives.PushBack("\t\t\t" + newTemp4 + " = " + newTemp4 + " - 1;\n")
		g.Natives.PushBack("\t\t\t" + newTemp5 + " = heap[(int)" + newTemp2 + "];\n")
		g.Natives.PushBack("\t\t\t" + newTemp6 + " = heap[(int)" + newTemp4 + "];\n")
		g.Natives.PushBack("\t\t\tif(" + newTemp5 + " <= " + newTemp6 + ") goto " + newLvl16 + ";\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl17 + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl16 + ":\n")
		g.Natives.PushBack("\t\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl17 + ":\n")
		g.Natives.PushBack("\t\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t" + newLvlend + ":\n")
		g.Natives.PushBack("\t\treturn;\n")
		g.Natives.PushBack("}\n\n")
		g.MenorFlag = false
	}
}

func (g *Generator) GenerateCompareMenorigString() {
	if g.MenorIgFlag {
		//generando temporales y etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newTemp3 := g.NewTemp()
		newTemp4 := g.NewTemp()
		newTemp5 := g.NewTemp()
		newTemp6 := g.NewTemp()
		newLvl1 := g.NewLabel()
		newLvl2 := g.NewLabel()
		newLvl3 := g.NewLabel()
		newLvl4 := g.NewLabel()
		newLvl5 := g.NewLabel()
		newLvl6 := g.NewLabel()
		newLvl7 := g.NewLabel()
		newLvl8 := g.NewLabel()
		newLvl9 := g.NewLabel()
		newLvl10 := g.NewLabel()
		newLvl11 := g.NewLabel()
		newLvl12 := g.NewLabel()
		newLvl13 := g.NewLabel()
		newLvl14 := g.NewLabel()
		newLvl15 := g.NewLabel()
		newLvl16 := g.NewLabel()
		newLvl17 := g.NewLabel()
		newLvlend := g.NewLabel()
		//se genera la funcion printstring
		g.Natives.PushBack("void comparemenorigString() {\n")
		g.Natives.PushBack("\t" + newTemp1 + " = P + 1;\n")
		g.Natives.PushBack("\t" + newTemp2 + " = stack[(int)" + newTemp1 + "];\n")
		g.Natives.PushBack("\t" + newTemp3 + " = P + 2;\n")
		g.Natives.PushBack("\t" + newTemp4 + " = stack[(int)" + newTemp3 + "];\n")
		g.Natives.PushBack("\t" + newTemp5 + " = heap[(int)" + newTemp2 + "];\n")
		g.Natives.PushBack("\t" + newTemp6 + " = heap[(int)" + newTemp4 + "];\n")
		g.Natives.PushBack("\tif(" + newTemp5 + " == -1) goto " + newLvl2 + ";\n")
		g.Natives.PushBack("\tgoto " + newLvl3 + ";\n")
		g.Natives.PushBack("\t" + newLvl2 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp6 + " == -1) goto " + newLvl4 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl5 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl4 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t" + newLvl5 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t" + newLvl3 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp6 + " == -1) goto " + newLvl6 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl7 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl6 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t" + newLvl7 + ":\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl1 + ";\n")
		g.Natives.PushBack("\t" + newLvl1 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp5 + " <= " + newTemp6 + ") goto " + newLvl8 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl9 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl8 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\t" + newTemp2 + " = " + newTemp2 + " + 1;\n")
		g.Natives.PushBack("\t\t\t" + newTemp4 + " = " + newTemp4 + " + 1;\n")
		g.Natives.PushBack("\t\t\t" + newTemp5 + " = heap[(int)" + newTemp2 + "];\n")
		g.Natives.PushBack("\t\t\t" + newTemp6 + " = heap[(int)" + newTemp4 + "];\n")
		g.Natives.PushBack("\t\t\tif(" + newTemp5 + " == -1) goto " + newLvl10 + ";\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl11 + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl10 + ":\n")
		g.Natives.PushBack("\t\t\t\tif(" + newTemp6 + " == -1) goto " + newLvl12 + ";\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvl13 + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl12 + ":\n")
		g.Natives.PushBack("\t\t\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl13 + ":\n")
		g.Natives.PushBack("\t\t\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl11 + ":\n")
		g.Natives.PushBack("\t\t\t\tif(" + newTemp6 + " == -1) goto " + newLvl14 + ";\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvl15 + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl14 + ":\n")
		g.Natives.PushBack("\t\t\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl15 + ":\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvl1 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl9 + ":\n")
		g.Natives.PushBack("\t\t\t" + newTemp2 + " = " + newTemp2 + " - 1;\n")
		g.Natives.PushBack("\t\t\t" + newTemp4 + " = " + newTemp4 + " - 1;\n")
		g.Natives.PushBack("\t\t\t" + newTemp5 + " = heap[(int)" + newTemp2 + "];\n")
		g.Natives.PushBack("\t\t\t" + newTemp6 + " = heap[(int)" + newTemp4 + "];\n")
		g.Natives.PushBack("\t\t\tif(" + newTemp5 + " <= " + newTemp6 + ") goto " + newLvl16 + ";\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl17 + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl16 + ":\n")
		g.Natives.PushBack("\t\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl17 + ":\n")
		g.Natives.PushBack("\t\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t" + newLvlend + ":\n")
		g.Natives.PushBack("\t\treturn;\n")
		g.Natives.PushBack("}\n\n")
		g.MenorIgFlag = false
	}
}

func (g *Generator) GenerateCompareMayorigString() {
	if g.MayorIgFlag {
		//generando temporales y etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newTemp3 := g.NewTemp()
		newTemp4 := g.NewTemp()
		newTemp5 := g.NewTemp()
		newTemp6 := g.NewTemp()
		newLvl1 := g.NewLabel()
		newLvl2 := g.NewLabel()
		newLvl3 := g.NewLabel()
		newLvl4 := g.NewLabel()
		newLvl5 := g.NewLabel()
		newLvl6 := g.NewLabel()
		newLvl7 := g.NewLabel()
		newLvl8 := g.NewLabel()
		newLvl9 := g.NewLabel()
		newLvlend := g.NewLabel()
		//se genera la funcion printstring
		g.Natives.PushBack("void comparemayorigString() {\n")
		g.Natives.PushBack("\t" + newTemp1 + " = P + 1;\n")
		g.Natives.PushBack("\t" + newTemp2 + " = stack[(int)" + newTemp1 + "];\n")
		g.Natives.PushBack("\t" + newTemp3 + " = P + 2;\n")
		g.Natives.PushBack("\t" + newTemp4 + " = stack[(int)" + newTemp3 + "];\n")
		g.Natives.PushBack("\t" + newTemp5 + " = heap[(int)" + newTemp2 + "];\n")
		g.Natives.PushBack("\t" + newTemp6 + " = heap[(int)" + newTemp4 + "];\n")
		g.Natives.PushBack("\tif(" + newTemp5 + " == -1) goto " + newLvl2 + ";\n")
		g.Natives.PushBack("\tgoto " + newLvl3 + ";\n")
		g.Natives.PushBack("\t" + newLvl2 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp6 + " == -1) goto " + newLvl4 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl5 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl4 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t" + newLvl5 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t" + newLvl3 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp6 + " == -1) goto " + newLvl6 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl7 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl6 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t" + newLvl7 + ":\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl1 + ";\n")
		g.Natives.PushBack("\t" + newLvl1 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp5 + " >= " + newTemp6 + ") goto " + newLvl8 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl9 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl8 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t" + newLvl9 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t" + newLvlend + ":\n")
		g.Natives.PushBack("\t\treturn;\n")
		g.Natives.PushBack("}\n\n")
		g.MayorIgFlag = false
	}
}

func (g *Generator) GenerateCompareIgualString() {
	if g.IgFlag {
		//generando temporales y etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newTemp3 := g.NewTemp()
		newTemp4 := g.NewTemp()
		newTemp5 := g.NewTemp()
		newTemp6 := g.NewTemp()
		newLvl1 := g.NewLabel()
		newLvl2 := g.NewLabel()
		newLvl3 := g.NewLabel()
		newLvl4 := g.NewLabel()
		newLvl5 := g.NewLabel()
		newLvl6 := g.NewLabel()
		newLvl7 := g.NewLabel()
		newLvl8 := g.NewLabel()
		newLvl9 := g.NewLabel()
		newLvl10 := g.NewLabel()
		newLvl11 := g.NewLabel()
		newLvl12 := g.NewLabel()
		newLvl13 := g.NewLabel()
		newLvl14 := g.NewLabel()
		newLvl15 := g.NewLabel()
		newLvlend := g.NewLabel()
		//se genera la funcion printstring
		g.Natives.PushBack("void compareigualString() {\n")
		g.Natives.PushBack("\t" + newTemp1 + " = P + 1;\n")
		g.Natives.PushBack("\t" + newTemp2 + " = stack[(int)" + newTemp1 + "];\n")
		g.Natives.PushBack("\t" + newTemp3 + " = P + 2;\n")
		g.Natives.PushBack("\t" + newTemp4 + " = stack[(int)" + newTemp3 + "];\n")
		g.Natives.PushBack("\t" + newTemp5 + " = heap[(int)" + newTemp2 + "];\n")
		g.Natives.PushBack("\t" + newTemp6 + " = heap[(int)" + newTemp4 + "];\n")
		g.Natives.PushBack("\tif(" + newTemp5 + " == -1) goto " + newLvl2 + ";\n")
		g.Natives.PushBack("\tgoto " + newLvl3 + ";\n")
		g.Natives.PushBack("\t" + newLvl2 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp6 + " == -1) goto " + newLvl4 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl5 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl4 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t" + newLvl5 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t" + newLvl3 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp6 + " == -1) goto " + newLvl6 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl7 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl6 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t" + newLvl7 + ":\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl1 + ";\n")
		g.Natives.PushBack("\t" + newLvl1 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp5 + " == " + newTemp6 + ") goto " + newLvl8 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl9 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl8 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\t" + newTemp2 + " = " + newTemp2 + " + 1;\n")
		g.Natives.PushBack("\t\t\t" + newTemp4 + " = " + newTemp4 + " + 1;\n")
		g.Natives.PushBack("\t\t\t" + newTemp5 + " = heap[(int)" + newTemp2 + "];\n")
		g.Natives.PushBack("\t\t\t" + newTemp6 + " = heap[(int)" + newTemp4 + "];\n")
		g.Natives.PushBack("\t\t\tif(" + newTemp5 + " == -1) goto " + newLvl10 + ";\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl11 + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl10 + ":\n")
		g.Natives.PushBack("\t\t\t\tif(" + newTemp6 + " == -1) goto " + newLvl12 + ";\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvl13 + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl12 + ":\n")
		g.Natives.PushBack("\t\t\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl13 + ":\n")
		g.Natives.PushBack("\t\t\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl11 + ":\n")
		g.Natives.PushBack("\t\t\t\tif(" + newTemp6 + " == -1) goto " + newLvl14 + ";\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvl15 + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl14 + ":\n")
		g.Natives.PushBack("\t\t\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl15 + ":\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvl1 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl9 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t" + newLvlend + ":\n")
		g.Natives.PushBack("\t\treturn;\n")
		g.Natives.PushBack("}\n\n")
		g.IgFlag = false
	}
}

func (g *Generator) GenerateCompareDifeString() {
	if g.DifFlag {
		//generando temporales y etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newTemp3 := g.NewTemp()
		newTemp4 := g.NewTemp()
		newTemp5 := g.NewTemp()
		newTemp6 := g.NewTemp()
		newLvl1 := g.NewLabel()
		newLvl2 := g.NewLabel()
		newLvl3 := g.NewLabel()
		newLvl4 := g.NewLabel()
		newLvl5 := g.NewLabel()
		newLvl6 := g.NewLabel()
		newLvl7 := g.NewLabel()
		newLvl8 := g.NewLabel()
		newLvl9 := g.NewLabel()
		newLvl10 := g.NewLabel()
		newLvl11 := g.NewLabel()
		newLvl12 := g.NewLabel()
		newLvl13 := g.NewLabel()
		newLvl14 := g.NewLabel()
		newLvl15 := g.NewLabel()
		newLvlend := g.NewLabel()
		//se genera la funcion printstring
		g.Natives.PushBack("void comparedifeString() {\n")
		g.Natives.PushBack("\t" + newTemp1 + " = P + 1;\n")
		g.Natives.PushBack("\t" + newTemp2 + " = stack[(int)" + newTemp1 + "];\n")
		g.Natives.PushBack("\t" + newTemp3 + " = P + 2;\n")
		g.Natives.PushBack("\t" + newTemp4 + " = stack[(int)" + newTemp3 + "];\n")
		g.Natives.PushBack("\t" + newTemp5 + " = heap[(int)" + newTemp2 + "];\n")
		g.Natives.PushBack("\t" + newTemp6 + " = heap[(int)" + newTemp4 + "];\n")
		g.Natives.PushBack("\tif(" + newTemp5 + " == -1) goto " + newLvl2 + ";\n")
		g.Natives.PushBack("\tgoto " + newLvl3 + ";\n")
		g.Natives.PushBack("\t" + newLvl2 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp6 + " == -1) goto " + newLvl4 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl5 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl4 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t" + newLvl5 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t" + newLvl3 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp6 + " == -1) goto " + newLvl6 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl7 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl6 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t" + newLvl7 + ":\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl1 + ";\n")
		g.Natives.PushBack("\t" + newLvl1 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp5 + " != " + newTemp6 + ") goto " + newLvl8 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl9 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl8 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t" + newLvl9 + ":\n")
		g.Natives.PushBack("\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\t" + newTemp2 + " = " + newTemp2 + " + 1;\n")
		g.Natives.PushBack("\t\t\t" + newTemp4 + " = " + newTemp4 + " + 1;\n")
		g.Natives.PushBack("\t\t\t" + newTemp5 + " = heap[(int)" + newTemp2 + "];\n")
		g.Natives.PushBack("\t\t\t" + newTemp6 + " = heap[(int)" + newTemp4 + "];\n")
		g.Natives.PushBack("\t\t\tif(" + newTemp5 + " == -1) goto " + newLvl10 + ";\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl11 + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl10 + ":\n")
		g.Natives.PushBack("\t\t\t\tif(" + newTemp6 + " == -1) goto " + newLvl12 + ";\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvl13 + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl12 + ":\n")
		g.Natives.PushBack("\t\t\t\t\t" + "stack[(int)P] = 0;" + "\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl13 + ":\n")
		g.Natives.PushBack("\t\t\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl11 + ":\n")
		g.Natives.PushBack("\t\t\t\tif(" + newTemp6 + " == -1) goto " + newLvl14 + ";\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvl15 + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl14 + ":\n")
		g.Natives.PushBack("\t\t\t\t\t" + "stack[(int)P] = 1;" + "\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvlend + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl15 + ":\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvl1 + ";\n")
		g.Natives.PushBack("\t" + newLvlend + ":\n")
		g.Natives.PushBack("\t\treturn;\n")
		g.Natives.PushBack("}\n\n")
		g.DifFlag = false
	}
}