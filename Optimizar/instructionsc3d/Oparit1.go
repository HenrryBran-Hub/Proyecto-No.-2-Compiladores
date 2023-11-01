package instructionsc3d

import (
	"Optimizar/environmentc3d"
	"strconv"
)

type Oparit1 struct {
	Id    string
	Valor string
}

func NewOparit1(id, valor string) Oparit1 {
	return Oparit1{id, valor}
}

func (v Oparit1) Ejecutar(ast *environmentc3d.AST) interface{} {
	valor := ast.ObtenerTemporal(v.Id)
	if valor != nil {
		valor.Valor = v.Valor
		valor.Tipo = ast.DeterminarTipo(v.Valor)
		if valor.Tipo == 0 || valor.Tipo == 1 {
			if valor.Tipo == 0 {
				intValue, _ := strconv.Atoi(v.Valor)
				if intValue < 0 {
					valor.IsOp = false
					valor.IsNeg = true
				}
			}
			if valor.Tipo == 1 {
				floatValue, _ := strconv.ParseFloat(v.Valor, 64)
				if floatValue < 0 {
					valor.IsOp = false
					valor.IsNeg = true
				}
			}
		}
		if valor.Tipo == 2 {
			valor.IsOp = true
		}
		ast.FinalCode.PushBack("\t" + valor.Id + " =  " + valor.Valor.(string) + ";\n")
		ast.ActualizarTemporal(valor)
	}

	return nil
}
