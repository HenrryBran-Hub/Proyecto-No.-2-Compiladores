package interfacesc3d

import (
	"Optimizar/environmentc3d"
)

type Expression interface {
	Ejecutar(ast *environmentc3d.AST) environmentc3d.Symbol
}
