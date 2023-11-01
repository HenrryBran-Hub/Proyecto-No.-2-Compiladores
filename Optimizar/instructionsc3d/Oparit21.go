package instructionsc3d

import (
	"Optimizar/environmentc3d"
)

type Oparit2 struct {
	Opp  string
	Tip1 string
	Val1 string
	Op   string
	Tip2 string
	Val2 string
	Fun1 bool
	Fun2 bool
	Fun3 bool
	Fun4 bool
}

func NewOparit21(opp, tip1, val1, op, tip2, val2 string) Oparit2 {
	result := Oparit2{
		Opp:  opp,
		Tip1: tip1,
		Val1: val1,
		Op:   op,
		Tip2: tip2,
		Val2: val2,
		Fun1: true,
		Fun2: false,
		Fun3: false,
		Fun4: false,
	}
	return result
}

func NewOparit22(opp, val1, op, val2 string) Oparit2 {
	result := Oparit2{
		Opp:  opp,
		Tip1: "",
		Val1: val1,
		Op:   op,
		Tip2: "",
		Val2: val2,
		Fun1: false,
		Fun2: true,
		Fun3: false,
		Fun4: false,
	}
	return result
}

func NewOparit23(opp, tip1, val1, op, val2 string) Oparit2 {
	result := Oparit2{
		Opp:  opp,
		Tip1: tip1,
		Val1: val1,
		Op:   op,
		Tip2: "",
		Val2: val2,
		Fun1: false,
		Fun2: false,
		Fun3: true,
		Fun4: false,
	}
	return result
}

func NewOparit24(opp, val1, op, tip2, val2 string) Oparit2 {
	result := Oparit2{
		Opp:  opp,
		Tip1: "",
		Val1: val1,
		Op:   op,
		Tip2: tip2,
		Val2: val2,
		Fun1: false,
		Fun2: false,
		Fun3: false,
		Fun4: true,
	}
	return result
}

func (v Oparit2) Ejecutar(ast *environmentc3d.AST) interface{} {
	valorprin := ast.ObtenerTemporal(v.Opp)
	if v.Fun1 {
		comval1 := ast.DeterminarTipo(v.Val1)
		comval2 := ast.DeterminarTipo(v.Val2)
		if comval1 == environmentc3d.ID && comval2 == environmentc3d.ID {
			valor1 := ast.ObtenerTemporal(v.Val1)
			valor2 := ast.ObtenerTemporal(v.Val2)
			if valorprin.Id == valor1.Id && valorprin.Id == valor2.Id {
				ast.FinalCode.PushBack("\t" + valorprin.Id + " =  (" + v.Tip1 + ")" + v.Val1 + " " + v.Op + " (" + v.Tip2 + ")" + v.Val2 + ";\n")
				valorprin.IsOp = true
				valorprin.Tipo = environmentc3d.OP
			} else if valorprin.Id == valor1.Id && valorprin.Id != valor2.Id {
				if valor2.Tipo == environmentc3d.NUMERO || valor2.Tipo == environmentc3d.DECIMAL {
					floatValue, _ := strconv.ParseFloat(v.Valor, 64)
				if floatValue < 0 {
					valor.IsOp = false
					valor.IsNeg = true
				}
					if 

					ast.FinalCode.PushBack("\t" + valorprin.Id + " =  (" + v.Tip1 + ")" + v.Val1 + " " + v.Op + " (" + v.Tip2 + ")" + valor2.Valor.(string) + ";\n")
					valorprin.IsOp = true
					valorprin.Tipo = environmentc3d.OP
				} else {
					ast.FinalCode.PushBack("\t" + valorprin.Id + " =  (" + v.Tip1 + ")" + v.Val1 + " " + v.Op + " (" + v.Tip2 + ")" + v.Val2 + ";\n")
					valorprin.IsOp = true
					valorprin.Tipo = environmentc3d.OP
				}
			} else if valorprin.Id != valor1.Id && valorprin.Id == valor2.Id {
				if valor1.Tipo == environmentc3d.NUMERO || valor1.Tipo == environmentc3d.DECIMAL {
					ast.FinalCode.PushBack("\t" + valorprin.Id + " =  (" + v.Tip1 + ")" + valor1.Valor.(string) + " " + v.Op + " (" + v.Tip2 + ")" + v.Val2 + ";\n")
					valorprin.IsOp = true
					valorprin.Tipo = environmentc3d.OP
				} else {
					ast.FinalCode.PushBack("\t" + valorprin.Id + " =  (" + v.Tip1 + ")" + v.Val1 + " " + v.Op + " (" + v.Tip2 + ")" + v.Val2 + ";\n")
					valorprin.IsOp = true
					valorprin.Tipo = environmentc3d.OP
				}
			} else {
				if (valor1.Tipo == environmentc3d.NUMERO || valor1.Tipo == environmentc3d.DECIMAL) && (valor2.Tipo == environmentc3d.NUMERO || valor2.Tipo == environmentc3d.DECIMAL) {
					ast.FinalCode.PushBack("\t" + valorprin.Id + " =  (" + v.Tip1 + ")" + valor1.Valor.(string) + " " + v.Op + " (" + v.Tip2 + ")" + valor2.Valor.(string) + ";\n")
					valorprin.IsOp = true
					valorprin.Tipo = environmentc3d.OP
				} else if (valor1.Tipo == environmentc3d.ID) && (valor2.Tipo == environmentc3d.NUMERO || valor2.Tipo == environmentc3d.DECIMAL) {
					ast.FinalCode.PushBack("\t" + valorprin.Id + " =  (" + v.Tip1 + ")" + v.Val1 + " " + v.Op + " (" + v.Tip2 + ")" + valor2.Valor.(string) + ";\n")
					valorprin.IsOp = true
					valorprin.Tipo = environmentc3d.OP
				} else if (valor1.Tipo == environmentc3d.NUMERO || valor1.Tipo == environmentc3d.DECIMAL) && (valor2.Tipo == environmentc3d.ID) {
					ast.FinalCode.PushBack("\t" + valorprin.Id + " =  (" + v.Tip1 + ")" + valor1.Valor.(string) + " " + v.Op + " (" + v.Tip2 + ")" + v.Val2 + ";\n")
					valorprin.IsOp = true
					valorprin.Tipo = environmentc3d.OP
				} else {
					ast.FinalCode.PushBack("\t" + valorprin.Id + " =  (" + v.Tip1 + ")" + v.Val1 + " " + v.Op + " (" + v.Tip2 + ")" + v.Val2 + ";\n")
					valorprin.IsOp = true
					valorprin.Tipo = environmentc3d.OP
				}
			}

		} else if comval1 != environmentc3d.ID && comval2 == environmentc3d.ID {

		} else if comval1 == environmentc3d.ID && comval2 != environmentc3d.ID {

		} else {
			
		}

	} else if v.Fun2 {

	} else if v.Fun3 {

	} else {

	}

	ast.ActualizarTemporal(valorprin)
	return nil
}

/*
if valor != nil {
			valor.Valor = v.Valor
			valor.Tipo = ast.DeterminarTipo(v.Valor)
			if valor.Tipo == 0 || valor.Tipo == 1 {
				if valor.Tipo == 0 {
					intValue, _ := strconv.Atoi(v.Valor)
					if intValue < 0 {
						valor.IsNeg = true
					}
				}
				if valor.Tipo == 1 {
					floatValue, _ := strconv.ParseFloat(v.Valor, 64)
					if floatValue < 0 {
						valor.IsNeg = true
					}
				}
			}
			ast.FinalCode.PushBack("\t" + valor.Id + " =  " + valor.Valor.(string) + ";\n")
			ast.ActualizarTemporal(valor)
		}
*/
