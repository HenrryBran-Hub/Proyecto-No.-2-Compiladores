package datoscompuestos

import (
	"Backend/interfaces"
)

type MatrizSimpleDos struct {
	Type   interfaces.Expression
	Expre  interfaces.Expression
	Numero string
	Lin    int
	Col    int
}

func NewMatrizSimpleDos(tipo interfaces.Expression, expre interfaces.Expression, numero string, line int, col int) MatrizSimpleDos {
	exp := MatrizSimpleDos{tipo, expre, numero, line, col}
	return exp
}

/*
func (o MatrizSimpleDos) Ejecutar(ast *environment.AST) interface{} {
	tipo := o.Type.Ejecutar(ast)
	expre := o.Expre.Ejecutar(ast)

	if tipo.Tipo != expre.Tipo {
		Errores := environment.Errores{
			Descripcion: "No es posible crear la matriz, esta colocando diferentes tipos de datos",
			Fila:        strconv.Itoa(o.Lin),
			Columna:     strconv.Itoa(o.Col),
			Tipo:        "Error Semantico",
			Ambito:      expre.Scope,
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	if strings.Contains(o.Numero, ".") {
		Errores := environment.Errores{
			Descripcion: "No es posible crear la matriz, esta colocando numeros no enteros",
			Fila:        strconv.Itoa(o.Lin),
			Columna:     strconv.Itoa(o.Col),
			Tipo:        "Error Semantico",
			Ambito:      expre.Scope,
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	num, err := strconv.Atoi(o.Numero)
	if err != nil {
		Errores := environment.Errores{
			Descripcion: "No es posible crear la matriz, Error en la conversion",
			Fila:        strconv.Itoa(o.Lin),
			Columna:     strconv.Itoa(o.Col),
			Tipo:        "Error Semantico",
			Ambito:      expre.Scope,
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	if num <= 0 {
		Errores := environment.Errores{
			Descripcion: "No es posible crear la matriz, el numero es menor o igual a 0",
			Fila:        strconv.Itoa(o.Lin),
			Columna:     strconv.Itoa(o.Col),
			Tipo:        "Error Semantico",
			Ambito:      expre.Scope,
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	lista := list.New()
	lista.PushBack(num)
	arreglo := environment.Valores_Matriz{
		Tipo:      tipo.Tipo,
		Valor:     expre,
		Matriztam: lista,
	}
	ast.Lista_Matriz_Val.PushFront(arreglo)
	return nil
}
*/
