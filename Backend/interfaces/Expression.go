package interfaces

import (
	"Backend/environment"
	"Backend/generator"
)

type Expression interface {
	Ejecutar(ast *environment.AST, gen *generator.Generator) environment.Value
}
