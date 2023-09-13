package interfaces

import "Backend/environment"

type Expression interface {
	Ejecutar(ast *environment.AST) environment.Symbol
}
