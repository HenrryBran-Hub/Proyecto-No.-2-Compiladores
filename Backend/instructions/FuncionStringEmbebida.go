package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
	"fmt"
	"strconv"
)

type FuncionStringEmbebida struct {
	Op interfaces.Expression
}

func NewFuncionStringEmbebida(op interfaces.Expression) FuncionStringEmbebida {
	exp := FuncionStringEmbebida{Op: op}
	return exp
}

func (o FuncionStringEmbebida) Ejecutar(ast *environment.AST) environment.Symbol {
	var op environment.Symbol
	op = o.Op.Ejecutar(ast)

	if op.Tipo == environment.INTEGER {
		r1 := fmt.Sprintf("%v", op.Valor)
		return environment.Symbol{Lin: op.Lin, Col: op.Col, Tipo: environment.STRING, Valor: r1}
	} else if op.Tipo == environment.FLOAT {
		num2 := fmt.Sprintf("%.6f", op.Valor)
		num3, err := strconv.ParseFloat(num2, 64)
		if err != nil {
			fmt.Println(err)
		}
		r1 := fmt.Sprintf("%v", num3)
		return environment.Symbol{Lin: op.Lin, Col: op.Col, Tipo: environment.STRING, Valor: r1}
	} else if op.Tipo == environment.STRING {
		return environment.Symbol{Lin: op.Lin, Col: op.Col, Tipo: environment.STRING, Valor: op.Valor}
	} else if op.Tipo == environment.BOOLEAN {
		r1 := fmt.Sprintf("%v", op.Valor)
		return environment.Symbol{Lin: op.Lin, Col: op.Col, Tipo: environment.STRING, Valor: r1}
	} else if op.Tipo == environment.CHARACTER {
		r1 := fmt.Sprintf("%v", op.Valor)
		return environment.Symbol{Lin: op.Lin, Col: op.Col, Tipo: environment.STRING, Valor: r1}
	} else {
		return environment.Symbol{Lin: op.Lin, Col: op.Col, Tipo: environment.STRING, Valor: nil}
	}
}
