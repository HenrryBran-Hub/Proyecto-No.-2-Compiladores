package environmentc3d

type TipoRegla int

const (
	REGLA1 TipoRegla = iota // 0
	REGLA2                  // 1
	REGLA3                  // 2
	REGLA4                  // 3
	REGLA5                  // 4
	REGLA6                  // 5
	REGLA7                  // 6
	REGLA8                  // 7
	REGLA9                  // 8
)

// Mapping de descripciones para cada valor de TipoRegla
var TipoReglaDescripciones = map[TipoRegla]string{
	REGLA1: "Si existe una asignación de valor de la forma a = b y posteriormente existe una asignación de forma b = a, se eliminará la segunda asignación siempre que a no haya cambiado su valor. Se deberá tener la seguridad de que no exista el cambio de valor y no existan etiquetas entre las 2 asignaciones.",
	REGLA2: "Esta regla sostiene que si existe un salto condicional Lx y una etiqueta Lx; todo el código que se encuentre entre ellos podrá ser eliminado siempre y cuando no exista una etiqueta en dicho código.",
	REGLA3: "Si se utilizan valores constantes dentro de las condiciones y el resultado de la condición es una constante verdadera, se podrá transformar en un salto incondicional y eliminarse el salto hacia la etiqueta falsa Lf.",
	REGLA4: "Si se utilizan valores constantes dentro de las condiciones y el resultado de la condición es una constante falsa, se podrá transformar en un salto incondicional y eliminarse el salto hacia la etiqueta verdadera Lv.",
	REGLA5: "Si tenemos un salto incondicional hacia una etiqueta Lx, seguidamente existen instrucciones y si inmediatamente después se tiene una etiqueta Lx en donde existe un salto incondicional hacia una etiqueta Ly, entonces el primer salto incondicional se debe cambiar a la etiqueta Ly.",
	REGLA6: "Si existe un salto incondicional hacia una etiqueta Lx de la forma if < cond > {goto Lx},	seguidamente existen instrucciones y si inmediatamente después se tiene una etiqueta Lx en donde existe un salto incondicional hacia una etiqueta Ly, entonces el primer salto incondicional se debe cambiar a la etiqueta Ly.",
	REGLA7: "Se elimina la instrucción si una variable se asigna una expresión con operaciones de Suma/resta con 0 o multiplicaciones/divisiones con 1 a la misma variable.",
	REGLA8: "Es aplicable la eliminación de instrucciones con operaciones de variable distinta a la	variable de asignación y una constante, si las operaciones son sumas/restas con 0 y multiplicaciones/divisiones con 1, la instrucción se transforma en una asignación.",
	REGLA9: "Se deberá realizar la eliminación de reducción por fuerza para sustituir por operaciones de	alto costo por expresiones equivalentes de menor costo.",
}

type TipoExpresion int

const (
	NUMERO  TipoExpresion = iota // 0
	DECIMAL                      // 1
	ID                           // 2
	OP                           // 3
)
