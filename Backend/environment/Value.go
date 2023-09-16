package environment

import list "container/list"

type Value struct {
	Value        string
	IsTemp       bool
	Type         TipoExpresion
	ArrType      TipoExpresion
	TrueLabel    *list.List
	FalseLabel   *list.List
	OutLabel     *list.List
	ReturnFlag   bool
	BreakFlag    bool
	ContinueFlag bool
	IntValue     int
}

func NewValue(Val string, tmp bool, typ TipoExpresion, ret, br, cont bool) Value {
	result := Value{
		Value:        Val,
		IsTemp:       tmp,
		Type:         typ,
		ArrType:      NULL,
		TrueLabel:    list.New(),
		FalseLabel:   list.New(),
		OutLabel:     list.New(),
		ReturnFlag:   ret,
		BreakFlag:    br,
		ContinueFlag: cont,
		IntValue:     0,
	}
	return result
}
