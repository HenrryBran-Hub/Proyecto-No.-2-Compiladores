package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
	"container/list"
	"strconv"
)

type Struck struct {
	Line int
	Col  int
	Name string
	Op   []interface{}
}

func NewStruck(line, col int, name string, op []interface{}) Struck {
	return Struck{line, col, name, op}
}

func (v Struck) Ejecutar(ast *environment.AST) interface{} {

	for _, inst := range v.Op {
		if inst == nil {
			continue
		}
		instruction, ok := inst.(interfaces.Instruction)
		if !ok {
			continue
		}
		instruction.Ejecutar(ast)
	}

	for e := ast.AtributosStruct.Front(); e != nil; e = e.Next() {
		valore := e.Value.(environment.Variable)

		// Compara con los elementos anteriores en la lista
		for i := ast.AtributosStruct.Front(); i != e; i = i.Next() {
			valori := i.Value.(environment.Variable)
			if valore.Name == valori.Name {
				Errores := environment.Errores{
					Descripcion: "Estas identificando los atributos con los mismos nombres tienes que cambiarlos",
					Fila:        strconv.Itoa(valore.Symbols.Lin),
					Columna:     strconv.Itoa(valore.Symbols.Col),
					Tipo:        "Error Semántico",
					Ambito:      "Struc-" + v.Name,
				}
				ast.ErroresHTML(Errores)
				ast.AtributosStruct.Init()
				return nil
			}
		}

		// Compara con los elementos siguientes en la lista
		for i := e.Next(); i != nil; i = i.Next() {
			valori := i.Value.(environment.Variable)
			if valore.Name == valori.Name {
				Errores := environment.Errores{
					Descripcion: "Estas identificando los atributos con los mismos nombres tienes que cambiarlos",
					Fila:        strconv.Itoa(valore.Symbols.Lin),
					Columna:     strconv.Itoa(valore.Symbols.Col),
					Tipo:        "Error Semántico",
					Ambito:      "Struc-" + v.Name,
				}
				ast.ErroresHTML(Errores)
				ast.AtributosStruct.Init()
				return nil
			}
		}
	}

	listaatributos := list.New()
	for e := ast.AtributosStruct.Front(); e != nil; e = e.Next() {
		listaatributos.PushBack(e.Value)
	}
	ast.AtributosStruct.Init()

	for e := ast.FuncionesStruct.Front(); e != nil; e = e.Next() {
		valore := e.Value.(environment.Funcion)

		// Compara con los elementos anteriores en la lista
		for i := ast.FuncionesStruct.Front(); i != e; i = i.Next() {
			valori := i.Value.(environment.Funcion)
			if valore.Nombre == valori.Nombre {
				Errores := environment.Errores{
					Descripcion: "Estas identificando las Funciones con los mismos nombres tienes que cambiarlos",
					Fila:        strconv.Itoa(valore.Lin),
					Columna:     strconv.Itoa(valore.Col),
					Tipo:        "Error Semántico",
					Ambito:      "Struc-Funcion-" + v.Name,
				}
				ast.ErroresHTML(Errores)
				ast.FuncionesStruct.Init()
				return nil
			}
		}

		// Compara con los elementos siguientes en la lista
		for i := e.Next(); i != nil; i = i.Next() {
			valori := i.Value.(environment.Funcion)
			if valore.Nombre == valori.Nombre {
				Errores := environment.Errores{
					Descripcion: "Estas identificando las Funciones con los mismos nombres tienes que cambiarlos",
					Fila:        strconv.Itoa(valore.Lin),
					Columna:     strconv.Itoa(valore.Col),
					Tipo:        "Error Semántico",
					Ambito:      "Struc-Funcion-" + v.Name,
				}
				ast.ErroresHTML(Errores)
				ast.FuncionesStruct.Init()
				return nil
			}
		}
	}

	listafunciones := list.New()
	for e := ast.FuncionesStruct.Front(); e != nil; e = e.Next() {
		listafunciones.PushBack(e.Value)
	}
	ast.FuncionesStruct.Init()

	struc := environment.Struc{
		Lin:       v.Line,
		Col:       v.Col,
		Nombre:    v.Name,
		Variables: listaatributos,
		Funciones: listafunciones,
	}

	ast.GuardarStruct(struc)
	return nil
}
