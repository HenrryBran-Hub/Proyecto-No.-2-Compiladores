package instructionsc3d

import (
	"Optimizar/environmentc3d"
	"Optimizar/interfacesc3d"
)

type FuncionMain struct {
	Bloque []interface{}
}

func NewFuncionMain(bloque []interface{}) FuncionMain {
	return FuncionMain{bloque}
}

func (v FuncionMain) Ejecutar(ast *environmentc3d.AST) interface{} {
	ast.FinalCode.PushBack("int main(){\n")
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
	ast.FinalCode.PushBack("\treturn 0;\n")
	ast.FinalCode.PushBack("}\n")
	return nil
}
