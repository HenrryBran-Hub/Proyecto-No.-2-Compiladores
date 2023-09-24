package sentencias

type TransferenciaContinue struct {
	Lin int
	Col int
}

func NewTransferenciaContinue(lin int, col int) TransferenciaContinue {
	return TransferenciaContinue{Lin: lin, Col: col}
}

/*
func (v TransferenciaContinue) Ejecutar(ast *environment.AST) interface{} {
	symbol := environment.Symbol{
		Lin:   v.Lin,
		Col:   v.Col,
		Tipo:  environment.BOOLEAN,
		Valor: true,
		Scope: ast.ObtenerAmbito(),
	}
	Variable := environment.Variable{
		Name:        "Continue",
		Symbols:     symbol,
		Mutable:     false,
		TipoSimbolo: "Sentencia de Transferencia",
	}

	ast.GuardarVariable(Variable)
	return nil
}

*/
