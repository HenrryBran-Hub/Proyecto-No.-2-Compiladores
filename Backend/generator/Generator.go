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
	IsString         bool
	CadtoNumber      bool
	IsAppend         bool
	IsCount          bool
	IsEmpty          bool
	IsRemovePos      bool
	IsRemoveLast     bool
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
		IsString:         true,
		CadtoNumber:      true,
		IsAppend:         true,
		IsCount:          true,
		IsEmpty:          true,
		IsRemovePos:      true,
		IsRemoveLast:     true,
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
	} else {
		g.FinalCode.PushBack("\n\n")
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
	g.FinalCode.PushBack("int main() {\n")
	g.FinalCode.PushBack("\tP = 0; H = 0;\n\n")
	for e := g.Code.Front(); e != nil; e = e.Next() {
		s := e.Value
		g.FinalCode.PushBack("\t" + s.(string))
	}
	g.FinalCode.PushBack("\n\treturn 0;\n}\n")
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

func (g *Generator) IsNumber() {
	if g.IsString {
		//generando temporales y etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newTemp3 := g.NewTemp()
		newTemp5 := g.NewTemp()
		newLvl1 := g.NewLabel()
		newLvl2 := g.NewLabel()
		newLvl3 := g.NewLabel()
		newLvl4 := g.NewLabel()
		newLvl5 := g.NewLabel()
		newLvl6 := g.NewLabel()
		newLvl7 := g.NewLabel()
		newLvl8 := g.NewLabel()
		//se genera la funcion printstring
		g.Natives.PushBack("void IsNumber() {\n")
		g.Natives.PushBack("\t" + newTemp1 + " = P + 1;\n")
		g.Natives.PushBack("\t" + newTemp2 + " = stack[(int)" + newTemp1 + "];\n")
		g.Natives.PushBack("\t" + newTemp3 + " = 1;\n")
		g.Natives.PushBack("\t" + newLvl1 + ":\n")
		g.Natives.PushBack("\t" + newTemp5 + " = heap[(int)" + newTemp2 + "];\n")
		g.Natives.PushBack("\tif(" + newTemp5 + " == -1) goto " + newLvl2 + ";\n")
		g.Natives.PushBack("\tgoto " + newLvl3 + ";\n")
		g.Natives.PushBack("\t" + newLvl2 + ":\n")
		g.Natives.PushBack("\t\tstack[(int)P] = " + newTemp3 + ";\n")
		g.Natives.PushBack("\t\treturn;\n")
		g.Natives.PushBack("\t" + newLvl3 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp5 + " < 48 ) goto " + newLvl4 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl5 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl4 + ":\n")
		g.Natives.PushBack("\t\t\t" + newTemp3 + " = 0;\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl8 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl5 + ":\n")
		g.Natives.PushBack("\t\t\tif(" + newTemp5 + " > 57 ) goto " + newLvl6 + ";\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl7 + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl6 + ":\n")
		g.Natives.PushBack("\t\t\t\t" + newTemp3 + " = 0;\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvl8 + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl7 + ":\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvl8 + ";\n")
		g.Natives.PushBack("\t" + newLvl8 + ":\n")
		g.Natives.PushBack("\t\t" + newTemp2 + " = " + newTemp2 + " + 1 ;\n")
		g.Natives.PushBack("\t\tgoto " + newLvl1 + ";\n")
		g.Natives.PushBack("}\n\n")
		g.IsString = false
	}
}

func (g *Generator) ConvertirCadenaANumero() {
	if g.CadtoNumber {
		//generando temporales y etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newTemp3 := g.NewTemp()
		newTemp4 := g.NewTemp()
		newTemp5 := g.NewTemp()
		newTemp6 := g.NewTemp()
		newTemp7 := g.NewTemp()
		newTemp8 := g.NewTemp()
		newTemp9 := g.NewTemp()
		newTemp10 := g.NewTemp()

		newTemp20 := g.NewTemp()
		newTemp21 := g.NewTemp()
		newTemp22 := g.NewTemp()
		newTemp23 := g.NewTemp()
		newTemp24 := g.NewTemp()

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

		newLvl20 := g.NewLabel()
		//newLvl21 := g.NewLabel()
		newLvl22 := g.NewLabel()
		newLvl23 := g.NewLabel()
		newLvl24 := g.NewLabel()
		newLvl25 := g.NewLabel()
		newLvl26 := g.NewLabel()
		newLvl27 := g.NewLabel()
		newLvl28 := g.NewLabel()
		newLvl29 := g.NewLabel()
		newLvl30 := g.NewLabel()
		newLvl31 := g.NewLabel()

		newLvlfin := g.NewLabel()
		newLvlexit := g.NewLabel()
		//se genera la funcion printstring
		g.Natives.PushBack("void ConvertCadtoNumber() {\n")
		g.Natives.PushBack("\t" + newTemp1 + " = P + 1;\n")
		g.Natives.PushBack("\t" + newTemp2 + " = stack[(int)" + newTemp1 + "];\n")
		g.Natives.PushBack("\t" + newTemp20 + " = stack[(int)" + newTemp1 + "];\n")
		g.Natives.PushBack("\t" + newTemp3 + " = 0.0;\n")
		g.Natives.PushBack("\t" + newTemp24 + " = 0.0;\n")
		g.Natives.PushBack("\t" + newTemp4 + " = 1.0;\n")
		g.Natives.PushBack("\t" + newTemp5 + " = 0;\n")
		g.Natives.PushBack("\t" + newTemp10 + " = 0;\n")
		g.Natives.PushBack("\t" + newTemp7 + " = 0;\n")
		g.Natives.PushBack("\t" + newTemp8 + " = 0.1;\n")
		g.Natives.PushBack("\t" + newTemp22 + " = 0;\n")
		g.Natives.PushBack("\t" + newTemp23 + " = -1;\n")
		g.Natives.PushBack("\t" + newLvl20 + ":\n")
		g.Natives.PushBack("\t" + newTemp21 + " = heap[(int)" + newTemp20 + "];\n")
		g.Natives.PushBack("\tif(" + newTemp21 + " != -1) goto " + newLvl22 + ";\n")
		g.Natives.PushBack("\tgoto " + newLvl23 + ";\n")
		g.Natives.PushBack("\t" + newLvl22 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp21 + " == 45) goto " + newLvl24 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl25 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl24 + ":\n")
		g.Natives.PushBack("\t\t\t" + newTemp20 + " = " + newTemp20 + " + 1;\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl20 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl25 + ":\n")
		g.Natives.PushBack("\t\t\tif(" + newTemp21 + " == 46) goto " + newLvl26 + ";\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl27 + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl26 + ":\n")
		g.Natives.PushBack("\t\t\t\t" + newTemp22 + " = 1;\n")
		g.Natives.PushBack("\t\t\t\t" + newTemp20 + " = " + newTemp20 + " + 1;\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvl20 + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl27 + ":\n")

		g.Natives.PushBack("\t\t\t\tif(" + newTemp22 + " == 1) goto " + newLvl28 + ";\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvl29 + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl28 + ":\n")
		g.Natives.PushBack("\t\t\t\t\t" + newTemp8 + " = " + newTemp8 + " * 10;\n")
		g.Natives.PushBack("\t\t\t\t\t" + newTemp20 + " = " + newTemp20 + " + 1;\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvl20 + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl29 + ":\n")
		g.Natives.PushBack("\t\t\t\t\t" + newTemp4 + " = " + newTemp4 + " * 10;\n")
		g.Natives.PushBack("\t\t\t\t\t" + newTemp20 + " = " + newTemp20 + " + 1;\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvl20 + ";\n")

		g.Natives.PushBack("\t" + newLvl23 + ":\n")
		g.Natives.PushBack("\t" + newTemp23 + " = " + newTemp8 + ";\n")
		g.Natives.PushBack("\t" + newTemp8 + " = " + newTemp8 + " / 10;\n")
		g.Natives.PushBack("\t" + newTemp4 + " = " + newTemp4 + " / 10;\n")

		g.Natives.PushBack("\t" + newLvl1 + ":\n")
		g.Natives.PushBack("\t" + newTemp9 + " = heap[(int)" + newTemp2 + "];\n")
		g.Natives.PushBack("\tif(" + newTemp9 + " == -1) goto " + newLvl2 + ";\n")
		g.Natives.PushBack("\tgoto " + newLvl3 + ";\n")
		g.Natives.PushBack("\t" + newLvl2 + ":\n")

		g.Natives.PushBack("\t\tif(" + newTemp23 + " != -1) goto " + newLvl30 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl31 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl30 + ":\n")
		g.Natives.PushBack("\t\t\t" + newTemp24 + " = " + newTemp24 + " / " + newTemp23 + ";\n")
		g.Natives.PushBack("\t\t\t" + newTemp3 + " = " + newTemp3 + " + " + newTemp24 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl31 + ":\n")

		g.Natives.PushBack("\t\tif(" + newTemp5 + " == 1) goto " + newLvl4 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl5 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl4 + ":\n")
		g.Natives.PushBack("\t\t\t" + newTemp3 + " = -" + newTemp3 + ";\n")
		g.Natives.PushBack("\t\t\tstack[(int)P] = " + newTemp3 + ";\n")
		g.Natives.PushBack("\t\t\treturn;\n")
		g.Natives.PushBack("\t\t" + newLvl5 + ":\n")
		g.Natives.PushBack("\t\t\tstack[(int)P] = " + newTemp3 + ";\n")
		g.Natives.PushBack("\t\t\treturn;\n")
		g.Natives.PushBack("\t" + newLvl3 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp9 + " == 45 ) goto " + newLvl6 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl7 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl6 + ":\n")
		g.Natives.PushBack("\t\t\t" + newTemp5 + " = 1;\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlfin + ";\n")
		g.Natives.PushBack("\t\t" + newLvl7 + ":\n")

		g.Natives.PushBack("\t\tif(" + newTemp9 + " >= 48 ) goto " + newLvl8 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl9 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl8 + ":\n")
		g.Natives.PushBack("\t\t\tif(" + newTemp9 + " <= 57 ) goto " + newLvl14 + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl14 + ":\n")
		g.Natives.PushBack("\t\t\t\t" + newTemp6 + " = " + newTemp9 + " - 48;\n")
		g.Natives.PushBack("\t\t\t\tif(" + newTemp7 + " == 1 ) goto " + newLvl10 + ";\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvl11 + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl10 + ":\n")
		g.Natives.PushBack("\t\t\t\t\t" + newTemp6 + " = " + newTemp6 + " * " + newTemp8 + ";\n")
		g.Natives.PushBack("\t\t\t\t\t" + newTemp24 + " = " + newTemp24 + " + " + newTemp6 + ";\n")
		g.Natives.PushBack("\t\t\t\t\t" + newTemp8 + " = " + newTemp8 + " / 10;\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvlfin + ";\n")
		g.Natives.PushBack("\t\t\t\t" + newLvl11 + ":\n")
		g.Natives.PushBack("\t\t\t\t\t" + newTemp6 + " = " + newTemp6 + " * " + newTemp4 + ";\n")
		g.Natives.PushBack("\t\t\t\t\t" + newTemp3 + " = " + newTemp3 + " + " + newTemp6 + ";\n")
		g.Natives.PushBack("\t\t\t\t\t" + newTemp4 + " = " + newTemp4 + " / 10;\n")
		g.Natives.PushBack("\t\t\t\t\tgoto " + newLvlfin + ";\n")
		g.Natives.PushBack("\t\t" + newLvl9 + ":\n")

		g.Natives.PushBack("\t\tif(" + newTemp9 + " == 46 ) goto " + newLvl15 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl16 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl15 + ":\n")
		g.Natives.PushBack("\t\t\tif(" + newTemp10 + " == 0 ) goto " + newLvl12 + ";\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl13 + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl12 + ":\n")
		g.Natives.PushBack("\t\t\t\t" + newTemp10 + " = 1;\n")
		g.Natives.PushBack("\t\t\t\t" + newTemp7 + " = 1;\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvlfin + ";\n")
		g.Natives.PushBack("\t\t\t" + newLvl13 + ":\n")
		g.Natives.PushBack("\t\t\t\tgoto " + newLvlexit + ";\n")
		g.Natives.PushBack("\t\t" + newLvl16 + ":\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvlexit + ";\n")
		g.Natives.PushBack("\t" + newLvlfin + ":\n")
		g.Natives.PushBack("\t\t" + newTemp2 + " = " + newTemp2 + " + 1 ;\n")
		g.Natives.PushBack("\t\tgoto " + newLvl1 + ";\n")
		g.Natives.PushBack("\t" + newLvlexit + ":\n")
		g.Natives.PushBack("\t\tstack[(int)P] = 201314439;\n")
		g.Natives.PushBack("\t\treturn;\n")
		g.Natives.PushBack("}\n\n")
		g.CadtoNumber = false
	}
}

func (g *Generator) AppendVector() {
	if g.IsAppend {
		//generando temporales y etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newTemp3 := g.NewTemp()
		//se genera la funcion printstring
		g.Natives.PushBack("void AppendVector() {\n")
		g.Natives.PushBack("\t" + newTemp1 + " = P + 0;\n")
		g.Natives.PushBack("\t" + newTemp2 + " = stack[(int)" + newTemp1 + "];\n")
		g.Natives.PushBack("\t" + newTemp3 + "= H;\n")
		g.Natives.PushBack("\theap[(int)H] = " + newTemp2 + ";\n")
		g.Natives.PushBack("\tH = H + 1;\n")
		g.Natives.PushBack("\tstack[(int)P] = " + newTemp3 + ";\n")
		g.Natives.PushBack("\treturn;\n")
		g.Natives.PushBack("}\n\n")
		g.IsAppend = false
	}
}

func (g *Generator) CountVector() {
	if g.IsCount {
		//generando temporales y etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newTemp3 := g.NewTemp()
		newLvl1 := g.NewLabel()
		newLvl2 := g.NewLabel()
		newLvl3 := g.NewLabel()
		newLvl4 := g.NewLabel()
		//se genera la funcion printstring
		g.Natives.PushBack("void CountVector() {\n")
		g.Natives.PushBack("\t" + newTemp1 + " = P + 1;\n")
		g.Natives.PushBack("\t" + newTemp2 + " = stack[(int)" + newTemp1 + "];\n")
		g.Natives.PushBack("\t" + newTemp3 + "= 0;\n")
		g.Natives.PushBack("\t" + newLvl1 + ":\n")
		g.Natives.PushBack("\t\tif(" + newTemp3 + " < " + newTemp2 + ") goto " + newLvl2 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl3 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl2 + ":\n")
		g.Natives.PushBack("\t\t\t" + newTemp3 + " = " + newTemp3 + " + 1;\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl1 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl3 + ":\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl4 + ";\n")
		g.Natives.PushBack("\t" + newLvl4 + ":\n")
		g.Natives.PushBack("\t\tstack[(int)P] = " + newTemp3 + ";\n")
		g.Natives.PushBack("\t\treturn;\n")
		g.Natives.PushBack("}\n\n")
		g.IsCount = false
	}
}

func (g *Generator) IsEmptyVector() {
	if g.IsEmpty {
		//generando temporales y etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newLvl2 := g.NewLabel()
		newLvl3 := g.NewLabel()
		newLvl4 := g.NewLabel()
		//se genera la funcion printstring
		g.Natives.PushBack("void IsEmptyVector() {\n")
		g.Natives.PushBack("\t" + newTemp1 + " = P + 1;\n")
		g.Natives.PushBack("\t" + newTemp2 + " = stack[(int)" + newTemp1 + "];\n")
		g.Natives.PushBack("\t\tif(" + newTemp2 + " == 0) goto " + newLvl2 + ";\n")
		g.Natives.PushBack("\t\tgoto " + newLvl3 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl2 + ":\n")
		g.Natives.PushBack("\t\t\tstack[(int)P] = 1;\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl4 + ";\n")
		g.Natives.PushBack("\t\t" + newLvl3 + ":\n")
		g.Natives.PushBack("\t\t\tstack[(int)P] = 0;\n")
		g.Natives.PushBack("\t\t\tgoto " + newLvl4 + ";\n")
		g.Natives.PushBack("\t" + newLvl4 + ":\n")
		g.Natives.PushBack("\t\treturn;\n")
		g.Natives.PushBack("}\n\n")
		g.IsEmpty = false
	}
}

func (g *Generator) RemovePosVector() {
	if g.IsRemovePos {
		//generando temporales y etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		//se genera la funcion printstring
		g.Natives.PushBack("void RemovePosVector() {\n")
		g.Natives.PushBack("\t" + newTemp1 + " = P + 1;\n")
		g.Natives.PushBack("\t" + newTemp2 + " = stack[(int)" + newTemp1 + "];\n")
		g.Natives.PushBack("\tprintf(\"%c\", 10);\n")                    // Agrega un espacio
		g.Natives.PushBack("\tprintf(\"%c\", (char)69);\n")              // E
		g.Natives.PushBack("\tprintf(\"%c\", (char)108);\n")             // l
		g.Natives.PushBack("\tprintf(\"%c\", (char)105);\n")             // i
		g.Natives.PushBack("\tprintf(\"%c\", (char)109);\n")             // m
		g.Natives.PushBack("\tprintf(\"%c\", (char)105);\n")             // i
		g.Natives.PushBack("\tprintf(\"%c\", (char)110);\n")             // n
		g.Natives.PushBack("\tprintf(\"%c\", (char)97);\n")              // a
		g.Natives.PushBack("\tprintf(\"%c\", (char)115);\n")             // s
		g.Natives.PushBack("\tprintf(\"%c\", (char)116);\n")             // t
		g.Natives.PushBack("\tprintf(\"%c\", (char)101);\n")             // e
		g.Natives.PushBack("\tprintf(\"%c\", (char)32);\n")              // Agrega un espacio
		g.Natives.PushBack("\tprintf(\"%c\", (char)117);\n")             // u
		g.Natives.PushBack("\tprintf(\"%c\", (char)110);\n")             // n
		g.Natives.PushBack("\tprintf(\"%c\", (char)32);\n")              // Agrega un espacio
		g.Natives.PushBack("\tprintf(\"%c\", (char)118);\n")             // v
		g.Natives.PushBack("\tprintf(\"%c\", (char)97);\n")              // a
		g.Natives.PushBack("\tprintf(\"%c\", (char)108);\n")             // l
		g.Natives.PushBack("\tprintf(\"%c\", (char)111);\n")             // o
		g.Natives.PushBack("\tprintf(\"%c\", (char)114);\n")             // r
		g.Natives.PushBack("\tprintf(\"%c\", (char)32);\n")              // Agrega un espacio
		g.Natives.PushBack("\tprintf(\"%c\", (char)100);\n")             // d
		g.Natives.PushBack("\tprintf(\"%c\", (char)101);\n")             // e
		g.Natives.PushBack("\tprintf(\"%c\", (char)32);\n")              // Agrega un espacio
		g.Natives.PushBack("\tprintf(\"%c\", (char)86);\n")              // V
		g.Natives.PushBack("\tprintf(\"%c\", (char)101);\n")             // e
		g.Natives.PushBack("\tprintf(\"%c\", (char)99);\n")              // c
		g.Natives.PushBack("\tprintf(\"%c\", (char)116);\n")             // t
		g.Natives.PushBack("\tprintf(\"%c\", (char)111);\n")             // o
		g.Natives.PushBack("\tprintf(\"%c\", (char)114);\n")             // r
		g.Natives.PushBack("\tprintf(\"%c\", (char)32);\n")              // Agrega un espacio
		g.Natives.PushBack("\tprintf(\"%c\", (char)101);\n")             // e
		g.Natives.PushBack("\tprintf(\"%c\", (char)110);\n")             // n
		g.Natives.PushBack("\tprintf(\"%c\", (char)32);\n")              // Agrega un espacio
		g.Natives.PushBack("\tprintf(\"%d\", (int)" + newTemp2 + ");\n") // X
		g.Natives.PushBack("\tprintf(\"%c\", (char)32);\n")              // Agrega un espacio
		g.Natives.PushBack("\tprintf(\"%c\", 10);\n")                    // Agrega un espacio
		g.Natives.PushBack("\tprintf(\"%c\", 10);\n")                    // Agrega un espacio
		g.Natives.PushBack("\treturn;\n")
		g.Natives.PushBack("}\n\n")
		g.IsRemovePos = false
	}
}

func (g *Generator) RemoveLastVector() {
	if g.IsRemoveLast {
		g.Natives.PushBack("void RemoveLastVector() {\n")
		g.Natives.PushBack("\tprintf(\"%c\", 10);\n")        // Agrega un espacio
		g.Natives.PushBack("\tprintf(\"%c\", (char)69);\n")  // E
		g.Natives.PushBack("\tprintf(\"%c\", (char)108);\n") // l
		g.Natives.PushBack("\tprintf(\"%c\", (char)105);\n") // i
		g.Natives.PushBack("\tprintf(\"%c\", (char)109);\n") // m
		g.Natives.PushBack("\tprintf(\"%c\", (char)105);\n") // i
		g.Natives.PushBack("\tprintf(\"%c\", (char)110);\n") // n
		g.Natives.PushBack("\tprintf(\"%c\", (char)97);\n")  // a
		g.Natives.PushBack("\tprintf(\"%c\", (char)115);\n") // s
		g.Natives.PushBack("\tprintf(\"%c\", (char)116);\n") // t
		g.Natives.PushBack("\tprintf(\"%c\", (char)101);\n") // e
		g.Natives.PushBack("\tprintf(\"%c\", (char)32);\n")  // Agrega un espacio
		g.Natives.PushBack("\tprintf(\"%c\", (char)117);\n") // u
		g.Natives.PushBack("\tprintf(\"%c\", (char)110);\n") // n
		g.Natives.PushBack("\tprintf(\"%c\", (char)32);\n")  // Agrega un espacio
		g.Natives.PushBack("\tprintf(\"%c\", (char)118);\n") // v
		g.Natives.PushBack("\tprintf(\"%c\", (char)97);\n")  // a
		g.Natives.PushBack("\tprintf(\"%c\", (char)108);\n") // l
		g.Natives.PushBack("\tprintf(\"%c\", (char)111);\n") // o
		g.Natives.PushBack("\tprintf(\"%c\", (char)114);\n") // r
		g.Natives.PushBack("\tprintf(\"%c\", (char)32);\n")  // Agrega un espacio
		g.Natives.PushBack("\tprintf(\"%c\", (char)250);\n") // ú
		g.Natives.PushBack("\tprintf(\"%c\", (char)108);\n") // l
		g.Natives.PushBack("\tprintf(\"%c\", (char)116);\n") // t
		g.Natives.PushBack("\tprintf(\"%c\", (char)105);\n") // i
		g.Natives.PushBack("\tprintf(\"%c\", (char)109);\n") // m
		g.Natives.PushBack("\tprintf(\"%c\", (char)111);\n") // o
		g.Natives.PushBack("\tprintf(\"%c\", (char)32);\n")  // Agrega un espacio
		g.Natives.PushBack("\tprintf(\"%c\", (char)10);\n")  // Agrega un salto de línea
		g.Natives.PushBack("\treturn;\n")
		g.Natives.PushBack("}\n\n")
		g.IsRemoveLast = false
	}
}
