package instructionsc3d

import (
	"Optimizar/environmentc3d"
	"Optimizar/interfacesc3d"
)

type FuncionVoid struct {
	Id     string
	Bloque []interface{}
}

func NewFuncionVoid(id string, bloque []interface{}) FuncionVoid {
	return FuncionVoid{id, bloque}
}

func (v FuncionVoid) Ejecutar(ast *environmentc3d.AST) interface{} {
	ast.FinalCode.PushBack("void " + v.Id + "(){\n")

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
	ast.FinalCode.PushBack("}\n")
	return nil
}
