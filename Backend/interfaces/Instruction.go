package interfaces

import "Backend/environment"

type Instruction interface {
	Ejecutar(ast *environment.AST) interface{}
}
