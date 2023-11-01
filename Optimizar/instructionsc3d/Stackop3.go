package instructionsc3d

import (
	"Optimizar/environmentc3d"
)

type Stackop3 struct {
	Op1  string
	Op2  string
	Op3  string
	Tip1 string
	Tip2 string
	Fun1 bool
	Fun2 bool
	Fun3 bool
	Fun4 bool
}

func NewStackop31(op1, tip2, op3 string) Stackop3 {
	result := Stackop3{
		Op1:  op1,
		Op2:  "",
		Op3:  op3,
		Tip1: "",
		Tip2: tip2,
		Fun1: true,
		Fun2: false,
		Fun3: false,
		Fun4: false,
	}
	return result
}

func NewStackop32(op1, op3 string) Stackop3 {
	result := Stackop3{
		Op1:  op1,
		Op2:  "",
		Op3:  op3,
		Tip1: "",
		Tip2: "",
		Fun1: false,
		Fun2: true,
		Fun3: false,
		Fun4: false,
	}
	return result
}

func NewStackop33(tip1, op2, tip2, op3 string) Stackop3 {
	result := Stackop3{
		Op1:  "",
		Op2:  op2,
		Op3:  op3,
		Tip1: tip1,
		Tip2: tip2,
		Fun1: false,
		Fun2: false,
		Fun3: true,
		Fun4: false,
	}
	return result
}

func NewStackop34(tip1, op2, op3 string) Stackop3 {
	result := Stackop3{
		Op1:  "",
		Op2:  op2,
		Op3:  op3,
		Tip1: tip1,
		Tip2: "",
		Fun1: false,
		Fun2: false,
		Fun3: false,
		Fun4: true,
	}
	return result
}

func (v Stackop3) Ejecutar(ast *environmentc3d.AST) interface{} {
	if v.Fun1 {
		ast.FinalCode.PushBack("\tstack[" + v.Op1 + "] = (" + v.Tip2 + ")" + v.Op3 + ";\n")
	} else if v.Fun2 {
		ast.FinalCode.PushBack("\tstack[" + v.Op1 + "] = " + v.Op3 + ";\n")
	} else if v.Fun3 {
		ast.FinalCode.PushBack("\tstack[" + v.Tip1 + v.Op2 + "] = (" + v.Tip2 + ")" + v.Op3 + ";\n")
	} else {
		ast.FinalCode.PushBack("\tstack[" + v.Tip1 + v.Op2 + "] = " + v.Op3 + ";\n")
	}
	return nil
}
