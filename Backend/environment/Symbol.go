package environment

type Symbol struct {
	Lin         int
	Col         int
	Tipo        TipoExpresion
	Valor       interface{}
	Scope       string
	Struct      string
	TipoDato    TipoExpresion
	Posicion    int
	ValorInt    int
	ValorFloat  float64
	ValorString string
	Valorpos    string
}
