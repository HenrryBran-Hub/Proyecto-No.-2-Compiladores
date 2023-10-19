package environment

import (
	list "container/list"
)

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
	FloatValue   float64
	StringValue  string
	Val          Variable
}

func NewValue(Val string, tmp bool, typ TipoExpresion, ret, br, cont bool, val Variable) Value {
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
		FloatValue:   0.0,
		StringValue:  "",
		Val:          val,
	}
	return result
}
