package interfaces

import (
	"Backend/environment"
	"Backend/generator"
)

type Instruction interface {
	Ejecutar(ast *environment.AST, gen *generator.Generator) interface{}
}
