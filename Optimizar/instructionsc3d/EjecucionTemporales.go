package instructionsc3d

import (
	"Optimizar/environmentc3d"
	"Optimizar/interfacesc3d"
)

type EjecucionTemporales struct {
	Bloque []interface{}
}

func NewEjecucionTemporales(bloque []interface{}) EjecucionTemporales {
	return EjecucionTemporales{bloque}
}

func (v EjecucionTemporales) Ejecutar(ast *environmentc3d.AST) interface{} {
	ast.FinalCode.PushBack("double ")
	for _, inst := range v.Bloque {
		if inst == nil {
			continue
		}
		instruction, ok := inst.(interfacesc3d.Instruction)
		if !ok {
			continue
		}
		instruction.Ejecutar(ast)
	}
	for e := ast.Lista_Temporales.Front(); e != nil; e = e.Next() {
		valor := e.Value.(environmentc3d.Temporal)
		if e.Next() != nil {
			ast.FinalCode.PushBack(valor.Id + ",")
		} else {
			ast.FinalCode.PushBack(valor.Id)
		}
	}
	ast.FinalCode.PushBack(";\n")
	return nil
}
