package environment

type TipoExpresion int

const (
	INTEGER   TipoExpresion = iota //0
	FLOAT                          //1
	STRING                         //2
	BOOLEAN                        //3
	CHARACTER                      //4
	NULL                           //5
	VECTOR                         //6
	MATRIZ                         //7
	STRUCT                         //8
	VARIABLE                       //9
)
