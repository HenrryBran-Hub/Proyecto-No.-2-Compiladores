package interfacesc3d

import "Optimizar/environmentc3d"

type Instruction interface {
	Ejecutar(ast *environmentc3d.AST) interface{}
}
