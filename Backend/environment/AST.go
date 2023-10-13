package environment

import (
	list "container/list"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/skratchdot/open-golang/open"
)

type AST struct {
	Instructions           []interface{}
	Print                  string
	Errors                 string
	Pila_Variables         *list.List
	Lista_Variables        *list.List
	Lista_VariablesHTML    *list.List
	Lista_Errores          *list.List
	Lista_Ambitos_Var      *list.List
	Lista_Arreglos         *list.List
	Pila_Arreglos          *list.List
	Lista_VectorHTML       *list.List
	Variables              Variable
	Stack                  []*Node
	Id                     int
	Dot                    string
	Root                   *Node
	ContadorDimen          int
	Lista_Matriz_Val       *list.List
	Pila_Matriz_Val        *list.List
	Lista_Matriz           *list.List
	Pila_Matriz            *list.List
	Lista_MatrizHTML       *list.List
	AtributosStruct        *list.List
	FuncionesStruct        *list.List
	Lista_Struct           *list.List
	Lista_Funciones        *list.List
	Lista_Funciones_Var    *list.List
	Lista_Funciones_Par    *list.List
	Lista_Variables_Struct *list.List
	Pila_Variables_Struct  *list.List
	ListaParametrosStruct  *list.List
	Lista_Struct_HTML      *list.List
	PosicionStack          int
	Lista_Tranferencias    *list.List
	Lista_Switch_Case      *list.List
	Lista_Switch_Case_Eti  *list.List
	Lista_For_Rango        *list.List
}

type Variable struct {
	Name        string
	Symbols     Symbol
	Mutable     bool
	TipoSimbolo string
	FEti        string
	TEti        string
}

type Vector struct {
	Name        string
	Symbols     Symbol
	Mutable     bool
	TipoSimbolo string
	Elements    *list.List
	ElementsPt  *list.List
}

type Matriz struct {
	Name        string
	Symbols     Symbol
	Mutable     bool
	TipoSimbolo string
	Elements    interface{}
}

type Errores struct {
	Descripcion string
	Fila        string
	Columna     string
	Tipo        string
	Ambito      string
}

type TablaSymbolosHTML struct {
	Id          string
	TipoSymbolo string
	TipoDato    string
	Lin         string
	Col         string
	Ambito      string
}

type Valores_Matriz struct {
	Tipo      TipoExpresion
	Valor     Symbol
	Matriztam *list.List
	Elements  *list.List
}

type Funcion struct {
	Lin           int
	Col           int
	Nombre        string
	IsReturn      bool
	IsParame      bool
	Tipo          TipoExpresion
	Retorno       interface{}
	Parametros    *list.List
	CodigoFuncion []interface{}
	Mutating      bool
}

type VariableFuncion struct {
	Name           string
	Symbols        Symbol
	Mutable        bool
	TipoSimbolo    string
	Inout          bool
	EI             bool
	ExternoInterno string
}

type Struc struct {
	Lin       int
	Col       int
	Nombre    string
	Variables *list.List
	Funciones *list.List
}

type VariableStruct struct {
	Lin      int
	Col      int
	Nombre   string
	Strukt   Struc
	Mutating bool
	Scope    string
}

type Parametrostruct struct {
	Name    string
	Symbolo Symbol
}

type SentenciasdeTransferencia struct {
	Scope  string
	ETrue  string
	EFalse string
	Tipo   TipoExpresion
}

func NewAST(inst []interface{}, print string, err string) AST {
	ast := AST{Instructions: inst, Print: print, Errors: err, Id: 0, Dot: "graph{ \n graph [bgcolor=lightgray, fontname=\"Arial\", fontsize=12]; \n"}
	return ast
}

func (a *AST) GetPrint() string {
	return a.Print
}

func (a *AST) SetPrint(ToPrint string) {
	a.Print = a.Print + ToPrint
}

func (a *AST) GetErrors() string {
	return a.Errors
}

func (a *AST) SetError(ToErr string) {
	a.Errors = a.Errors + ToErr
}

func (a *AST) IniciarAmbito() {
	a.Pila_Variables = list.New()
	a.Lista_Variables = list.New()
	a.Lista_VariablesHTML = list.New()
	a.Pila_Variables.PushBack(a.Lista_Variables)

	a.Lista_Ambitos_Var = list.New()
	a.Pila_Arreglos = list.New()
	a.Lista_Arreglos = list.New()
	a.Lista_VectorHTML = list.New()
	a.Pila_Arreglos.PushBack(a.Lista_Arreglos)

	a.Lista_Matriz_Val = list.New()
	a.Pila_Matriz_Val = list.New()
	a.Pila_Matriz_Val.PushBack(a.Lista_Matriz_Val)
	a.Lista_Matriz = list.New()
	a.Pila_Matriz = list.New()
	a.Lista_MatrizHTML = list.New()
	a.Pila_Matriz.PushBack(a.Lista_Matriz)

	a.Lista_Errores = list.New()
	a.Lista_Ambitos_Var.PushFront("Global")

	a.AtributosStruct = list.New()
	a.FuncionesStruct = list.New()
	a.Lista_Struct = list.New()

	a.Lista_Funciones = list.New()
	a.Lista_Funciones_Var = list.New()
	a.Lista_Funciones_Par = list.New()

	a.Lista_Variables_Struct = list.New()
	a.Pila_Variables_Struct = list.New()
	a.Pila_Variables_Struct.PushBack(a.Lista_Variables_Struct)
	a.ListaParametrosStruct = list.New()
	a.Lista_Struct_HTML = list.New()

	a.PosicionStack = 0
	a.Lista_Tranferencias = list.New()

	a.Lista_Switch_Case = list.New()
	a.Lista_Switch_Case_Eti = list.New()

	a.Lista_For_Rango = list.New()

}

func (a *AST) AumentarAmbito(ambito string) {
	nuevaLista := list.New()
	a.Pila_Variables.PushFront(nuevaLista)
	a.Lista_Variables = nuevaLista

	nuevaArreglos_Val := list.New()
	a.Pila_Arreglos.PushFront(nuevaArreglos_Val)
	a.Lista_Arreglos = nuevaArreglos_Val

	nuevaMatriz_Val := list.New()
	a.Pila_Matriz.PushFront(nuevaMatriz_Val)
	a.Lista_Matriz = nuevaMatriz_Val

	nuevaStruct_Val := list.New()
	a.Pila_Variables_Struct.PushFront(nuevaStruct_Val)
	a.Lista_Variables_Struct = nuevaStruct_Val

	a.Lista_Ambitos_Var.PushBack(ambito)
}

func (a *AST) AumentarNivel() {
	nuevalista := list.New()
	a.Pila_Matriz_Val.PushFront(nuevalista)
	a.Lista_Matriz_Val = nuevalista
}

func (a *AST) QuitarNiveles() {
	a.Pila_Matriz_Val.Init()
	a.Lista_Matriz_Val.Init()
	nuevaArreglos_Val := list.New()
	a.Pila_Arreglos.PushFront(nuevaArreglos_Val)
	a.Lista_Arreglos = nuevaArreglos_Val
}

func (a *AST) DisminuirAmbito() {
	a.Pila_Variables.Remove(a.Pila_Variables.Front())
	a.Lista_Variables = a.Pila_Variables.Front().Value.(*list.List)

	a.Pila_Arreglos.Remove(a.Pila_Arreglos.Front())
	a.Lista_Arreglos = a.Pila_Arreglos.Front().Value.(*list.List)

	a.Pila_Matriz.Remove(a.Pila_Matriz.Front())
	a.Lista_Matriz = a.Pila_Matriz.Front().Value.(*list.List)

	if a.Lista_Ambitos_Var.Len() > 0 {
		a.Lista_Ambitos_Var.Remove(a.Lista_Ambitos_Var.Back())
	}

	a.Pila_Variables_Struct.Remove(a.Pila_Variables_Struct.Front())
	a.Lista_Variables_Struct = a.Pila_Variables_Struct.Front().Value.(*list.List)
}

func (a *AST) GuardarVariable(variable Variable) {
	for e := a.Lista_Variables.Front(); e != nil; e = e.Next() {
		if e.Value.(Variable).Name == variable.Name {
			Errores := Errores{
				Descripcion: "La variale que esta intentando guardar ya existe en este ambito: \n Variable: " + variable.Name,
				Fila:        strconv.Itoa(e.Value.(Variable).Symbols.Lin),
				Columna:     strconv.Itoa(e.Value.(Variable).Symbols.Col),
				Tipo:        "Error Semantico",
				Ambito:      variable.Symbols.Scope,
			}
			a.ErroresHTML(Errores)
			return
		}
	}
	a.Lista_Variables.PushBack(variable)
	a.PosicionStack = a.PosicionStack + 1
	if variable.Name == "Break" || variable.Name == "Continue" || variable.Name == "Return" || variable.Name == "ReturnExp" {
		return
	}
	a.Lista_VariablesHTML.PushBack(variable)
}

func (a *AST) ObtenerAmbito() string {
	var lastScope string
	if a.Lista_Ambitos_Var.Len() > 0 {
		lastScope = a.Lista_Ambitos_Var.Back().Value.(string)
	}
	return lastScope
}

func (a *AST) GuardarArreglo(vector Vector) {
	for e := a.Lista_Arreglos.Front(); e != nil; e = e.Next() {
		if e.Value.(Vector).Name == vector.Name {
			Errores := Errores{
				Descripcion: "El arreglo que esta intentando guardar ya existe en este ambito: \n Arreglo: " + vector.Name,
				Fila:        strconv.Itoa(e.Value.(Variable).Symbols.Lin),
				Columna:     strconv.Itoa(e.Value.(Variable).Symbols.Col),
				Tipo:        "Error Semantico",
				Ambito:      vector.Symbols.Scope,
			}
			a.ErroresHTML(Errores)
			return
		}
	}
	a.Lista_Arreglos.PushBack(vector)
	a.Lista_VectorHTML.PushBack(vector)
	a.PosicionStack = a.PosicionStack + 1
}

func (a *AST) GuardarMatriz(matriz Matriz) {
	for e := a.Lista_Arreglos.Front(); e != nil; e = e.Next() {
		if e.Value.(Vector).Name == matriz.Name {
			Errores := Errores{
				Descripcion: "La Matriz que esta intentando guardar ya existe en este ambito: \n Arreglo: " + matriz.Name,
				Fila:        strconv.Itoa(e.Value.(Variable).Symbols.Lin),
				Columna:     strconv.Itoa(e.Value.(Variable).Symbols.Col),
				Tipo:        "Error Semantico",
				Ambito:      matriz.Symbols.Scope,
			}
			a.ErroresHTML(Errores)
			return
		}
	}
	for e := a.Lista_Matriz.Front(); e != nil; e = e.Next() {
		if e.Value.(Matriz).Name == matriz.Name {
			Errores := Errores{
				Descripcion: "La Matriz que esta intentando guardar ya existe en este ambito: \n Arreglo: " + matriz.Name,
				Fila:        strconv.Itoa(e.Value.(Variable).Symbols.Lin),
				Columna:     strconv.Itoa(e.Value.(Variable).Symbols.Col),
				Tipo:        "Error Semantico",
				Ambito:      matriz.Symbols.Scope,
			}
			a.ErroresHTML(Errores)
			return
		}
	}

	a.Lista_Matriz.PushBack(matriz)
	a.Lista_MatrizHTML.PushBack(matriz)
	a.PosicionStack = a.PosicionStack + 1
}

func (a *AST) ActualizarVariable(mariable *Variable) {
	for e := a.Pila_Variables.Front(); e != nil; e = e.Next() {
		lista := e.Value.(*list.List)
		for v := lista.Front(); v != nil; v = v.Next() {
			if v.Value.(Variable).Name == mariable.Name && mariable.Mutable {
				variable := v.Value.(Variable)
				variable.Symbols.Col = mariable.Symbols.Col
				variable.Symbols.Lin = mariable.Symbols.Lin
				variable.Symbols.Scope = mariable.Symbols.Scope
				variable.Symbols.Tipo = mariable.Symbols.Tipo
				variable.Symbols.Valor = mariable.Symbols.Valor
				variable.Symbols.ValorInt = mariable.Symbols.ValorInt
				v.Value = variable // Actualizar la variable en la lista
				for j := a.Lista_VariablesHTML.Front(); j != nil; j = j.Next() {
					if j.Value.(Variable).Name == mariable.Name && mariable.Mutable && j.Value.(Variable).Symbols.Scope == mariable.Symbols.Scope {
						variablej := j.Value.(Variable)
						variablej.Symbols.Col = mariable.Symbols.Col
						variablej.Symbols.Lin = mariable.Symbols.Lin
						variablej.Symbols.Scope = mariable.Symbols.Scope
						variablej.Symbols.Tipo = mariable.Symbols.Tipo
						variablej.Symbols.Valor = mariable.Symbols.Valor
						variablej.Symbols.ValorInt = mariable.Symbols.ValorInt
						j.Value = variable // Actualizar la variable en la lista
					}
				}
				return
			}
		}
	}
	Errores := Errores{
		Descripcion: "La variale que esta intentando modificar no existe: \n Variable: " + mariable.Name,
		Fila:        strconv.Itoa(mariable.Symbols.Lin),
		Columna:     strconv.Itoa(mariable.Symbols.Col),
		Tipo:        "Error Semantico",
		Ambito:      mariable.Symbols.Scope,
	}
	a.ErroresHTML(Errores)
}

func (a *AST) GetVariable(nombre string) *Variable {
	for e := a.Pila_Variables.Front(); e != nil; e = e.Next() {
		lista := e.Value.(*list.List)
		for v := lista.Front(); v != nil; v = v.Next() {
			variable := v.Value.(Variable)
			if variable.Name == nombre {
				return &variable
			}
		}
	}
	return nil
}

func (a *AST) ActualizarArreglo(nombre string, nuevoValor *Vector) {
	for e := a.Pila_Arreglos.Front(); e != nil; e = e.Next() {
		lista := e.Value.(*list.List)
		for v := lista.Front(); v != nil; v = v.Next() {
			if v.Value.(Vector).Name == nombre && v.Value.(Vector).Mutable {
				vector := v.Value.(Vector)
				vector.Elements = nuevoValor.Elements
				for i := a.Lista_VectorHTML.Front(); i != nil; i = i.Next() {
					if i.Value.(Vector).Name == nombre && i.Value.(Vector).Mutable && i.Value.(Vector).Symbols.Scope == nuevoValor.Symbols.Scope {
						vectorj := i.Value.(Vector)
						vectorj.Elements = nuevoValor.Elements
						vectorj.ElementsPt = nuevoValor.ElementsPt
					}
				}
				return
			}
		}
	}
	Errores := Errores{
		Descripcion: "El arreglo que está intentando modificar no existe: \n Arreglo: " + nombre,
		Fila:        strconv.Itoa(nuevoValor.Symbols.Lin),
		Columna:     strconv.Itoa(nuevoValor.Symbols.Col),
		Tipo:        "Error Semantico",
		Ambito:      nuevoValor.Symbols.Scope,
	}
	a.ErroresHTML(Errores)
}

func (a *AST) GetArreglo(nombre string) *Vector {
	for e := a.Pila_Arreglos.Front(); e != nil; e = e.Next() {
		lista := e.Value.(*list.List)
		for v := lista.Front(); v != nil; v = v.Next() {
			vector := v.Value.(Vector)
			if vector.Name == nombre {
				return &vector
			}
		}
	}
	return nil
}

func (a *AST) GetMatriz(nombre string) *Matriz {
	for e := a.Pila_Matriz.Front(); e != nil; e = e.Next() {
		lista := e.Value.(*list.List)
		for v := lista.Front(); v != nil; v = v.Next() {
			matriz := v.Value.(Matriz)
			if matriz.Name == nombre {
				return &matriz
			}
		}
	}
	return nil
}

func (a *AST) ErroresHTML(errores Errores) {
	a.Lista_Errores.PushBack(errores)
}

func (a *AST) TablaVariablesHTML() {
	// Create a new HTML file
	file, err := os.Create("TablaDeSimbolos.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write the HTML header and table structure
	fmt.Fprintln(file, `<!DOCTYPE html>
	   <html>
	   <head>
	   	<title>Tabla de Simbolos</title>
	   	<style>
	   		body {
	   			background-color: #333;
	   			color: white;
	   			font-family: sans-serif;
	   		}
	   		h1 {
	   			text-align: center;
	   		}
	   		table {
	   			border-collapse: collapse;
	   			width: 100%;
	   			border: 2px double black;
	   		}
	   		th, td {
	   			text-align: center;
	   			padding: 8px;
	   			border-bottom: 1px solid #ddd;
	   		}
	   		th {
	   			background-color: goldenrod;
	   		}
	   		tr:hover {background-color: #555;}
	   	</style>
	   </head>
	   <body>
	   	<h1>Tabla de Simbolos</h1>
	   	<table>
	   		<tr>
	   			<th>No</th>
				<th>Tipo Simbolo</th>
	   			<th>Nombre</th>
	   			<th>Mutable</th>
	   			<th>Fila</th>
	   			<th>Columna</th>
	   			<th>Valor</th>
	   			<th>Tipo</th>
	   			<th>Ambito</th>
	   		</tr>`)

	rowNumber := 1
	for e := a.Lista_VariablesHTML.Front(); e != nil; e = e.Next() {
		variable := e.Value.(Variable)
		var tipoexpstr string
		switch variable.Symbols.Tipo {
		case 0:
			tipoexpstr = "Int"
		case 1:
			tipoexpstr = "Float"
		case 2:
			tipoexpstr = "String"
		case 3:
			tipoexpstr = "Boolean"
		case 4:
			tipoexpstr = "Character"
		default:
			tipoexpstr = "nil"
		}
		fmt.Fprintf(file, `
   					<tr>
   						<td>%d</td>
						<td>%s</td>
   						<td>%s</td>
   						<td>%t</td>
   						<td>%d</td>
   						<td>%d</td>
   						<td>%v</td>
   						<td>%s</td>
   						<td>%s</td>
   					</tr>`,
			rowNumber,
			variable.TipoSimbolo,
			variable.Name,
			variable.Mutable,
			variable.Symbols.Lin,
			variable.Symbols.Col,
			variable.Symbols.Valor,
			tipoexpstr,
			variable.Symbols.Scope,
		)
		rowNumber++
	}

	for e := a.Lista_VectorHTML.Front(); e != nil; e = e.Next() {
		vector := e.Value.(Vector)
		var tipoexpstr string
		switch vector.Symbols.Tipo {
		case 0:
			tipoexpstr = "Int"
		case 1:
			tipoexpstr = "Float"
		case 2:
			tipoexpstr = "String"
		case 3:
			tipoexpstr = "Boolean"
		case 4:
			tipoexpstr = "Character"
		default:
			tipoexpstr = "nil"
		}

		var acumulado string
		for element := vector.Elements.Front(); element != nil; element = element.Next() {
			symbolo := element.Value.(Value)
			stringValue := fmt.Sprintf("%v", symbolo.Value)
			acumulado += stringValue
			if element.Next() != nil {
				acumulado += ","
			}
		}

		acumulado += " Tamanio Vector: " + strconv.Itoa(vector.Elements.Len())

		fmt.Fprintf(file, `
   					<tr>
   						<td>%d</td>
						<td>%s</td>
   						<td>%s</td>
   						<td>%t</td>
   						<td>%d</td>
   						<td>%d</td>
   						<td>%v</td>
   						<td>%s</td>
   						<td>%s</td>
   					</tr>`,
			rowNumber,
			vector.TipoSimbolo,
			vector.Name,
			vector.Mutable,
			vector.Symbols.Lin,
			vector.Symbols.Col,
			acumulado,
			tipoexpstr,
			vector.Symbols.Scope,
		)
		rowNumber++
	}

	for e := a.Lista_MatrizHTML.Front(); e != nil; e = e.Next() {
		matriz := e.Value.(Matriz)
		var tipoexpstr string
		switch matriz.Symbols.Tipo {
		case 0:
			tipoexpstr = "Int"
		case 1:
			tipoexpstr = "Float"
		case 2:
			tipoexpstr = "String"
		case 3:
			tipoexpstr = "Boolean"
		case 4:
			tipoexpstr = "Character"
		default:
			tipoexpstr = "nil"
		}

		matrizStr := fmt.Sprintf("%v", matriz.Elements)

		fmt.Fprintf(file, `
   					<tr>
   						<td>%d</td>
						<td>%s</td>
   						<td>%s</td>
   						<td>%t</td>
   						<td>%d</td>
   						<td>%d</td>
   						<td>%v</td>
   						<td>%s</td>
   						<td>%s</td>
   					</tr>`,
			rowNumber,
			matriz.TipoSimbolo,
			matriz.Name,
			matriz.Mutable,
			matriz.Symbols.Lin,
			matriz.Symbols.Col,
			matrizStr,
			tipoexpstr,
			matriz.Symbols.Scope,
		)
		rowNumber++
	}

	for e := a.Lista_Struct_HTML.Front(); e != nil; e = e.Next() {
		stru := e.Value.(VariableStruct)
		//variables
		var tipoexpstr string = ""
		for e := stru.Strukt.Variables.Front(); e != nil; e = e.Next() {
			tipoexpstr += " " + e.Value.(Variable).Name + " - " + fmt.Sprintf("%v", e.Value.(Variable).Symbols.Valor)
		}

		//funciones
		var struStr string = ""
		for e := stru.Strukt.Funciones.Front(); e != nil; e = e.Next() {
			struStr += " " + e.Value.(Funcion).Nombre
		}

		fmt.Fprintf(file, `
   					<tr>
   						<td>%d</td>
						<td>%s</td>
   						<td>%s</td>
   						<td>%t</td>
   						<td>%d</td>
   						<td>%d</td>
   						<td>%v</td>
   						<td>%s</td>
   						<td>%s</td>
   					</tr>`,
			rowNumber,
			stru.Strukt.Nombre,
			stru.Nombre,
			stru.Mutating,
			stru.Lin,
			stru.Col,
			"Funciones: "+struStr,
			"Variables: "+tipoexpstr,
			stru.Scope,
		)
		rowNumber++
	}

	// Write the HTML footer
	fmt.Fprintln(file, `
           </table>
       </body>
       </html>`)

	// Open the HTML file in the default web browser
	open.Start("TablaDeSimbolos.html")
}

func (a *AST) TablaErroresHTML() {
	// Create a new HTML file
	file, err := os.Create("TablaErrores.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write the HTML header and table structure
	fmt.Fprintln(file, `<!DOCTYPE html>
	   <html>
	   <head>
	   	<title>Reporte de Errores</title>
	   	<style>
	   		body {
	   			background-color: #333;
	   			color: white;
	   			font-family: sans-serif;
	   		}
	   		h1 {
	   			text-align: center;
	   		}
	   		table {
	   			border-collapse: collapse;
	   			width: 100%;
	   			border: 2px double black;
	   		}
	   		th, td {
	   			text-align: center;
	   			padding: 8px;
	   			border-bottom: 1px solid #ddd;
	   		}
	   		th {
	   			background-color: firebrick;
	   		}
	   		tr:hover {background-color: #555;}
	   	</style>
	   </head>
	   <body>
	   	<h1>Tabla de Errores</h1>
	   	<table>
	   		<tr>
	   			<th>No</th>
				<th>Descripcion</th>
	   			<th>Ambito</th>
	   			<th>Fila</th>
	   			<th>Columna</th>
				<th>Tipo de Error</th>
	   		</tr>`)

	rowNumber := 1
	for e := a.Lista_Errores.Front(); e != nil; e = e.Next() {
		errorItem := e.Value.(Errores)
		fmt.Fprintf(file, `
   					<tr>
   						<td>%d</td>
   						<td>%s</td>
						<td>%s</td>
						<td>%s</td>
						<td>%s</td>
						<td>%s</td>
   					</tr>`,
			rowNumber,
			errorItem.Descripcion,
			errorItem.Ambito,
			errorItem.Fila,
			errorItem.Columna,
			errorItem.Tipo,
		)
		rowNumber++
	}
	// Write the HTML footer
	fmt.Fprintln(file, `
           </table>
       </body>
       </html>`)

	// Open the HTML file in the default web browser
	open.Start("TablaErrores.html")
}

func (a *AST) removeFromListFromBack(valor int) {
	for e := 0; e < valor; e++ {
		lastElement := a.Lista_VariablesHTML.Back()
		if lastElement != nil {
			a.Lista_VariablesHTML.Remove(lastElement)
		}
	}
}

func (a *AST) LimpiarLista() {
	numElementsToRemove := a.Lista_Variables.Len()
	for i := 0; i < numElementsToRemove; i++ {
		lastElement := a.Lista_Variables.Back()
		if lastElement != nil {
			a.Lista_Variables.Remove(lastElement)
		}
	}
	a.removeFromListFromBack(numElementsToRemove)
}

type Node struct {
	Id       int
	Label    string
	Children []*Node
}

func (t *AST) Postorder(tmp *Node) {
	if tmp != nil {
		for _, child := range tmp.Children {
			t.Postorder(child)
		}
		t.Dot += strconv.Itoa(tmp.Id) + "[label=\"" + tmp.Label + "\" , shape=box, style=\"filled,rounded\", fillcolor=\"#F4A460\"];\n"
		for _, child := range tmp.Children {
			t.Dot += strconv.Itoa(tmp.Id) + "->" + strconv.Itoa(child.Id) + ";\n"
		}
	}
}

func (l *AST) Push(i *Node) {
	l.Stack = append(l.Stack, i)
}

func (l *AST) Pop() *Node {
	if len(l.Stack) < 1 {
		panic("empty stack")
	}
	result := l.Stack[len(l.Stack)-1]
	l.Stack = l.Stack[:len(l.Stack)-1]
	return result
}

func (c *AST) GenerateCST() {
	cstFile, err := os.Create("cst.dot")
	if err != nil {
		fmt.Println("Error creating CST DOT file:", err)
		return
	}
	defer cstFile.Close()

	_, err = cstFile.WriteString("digraph CST {\n")
	if err != nil {
		fmt.Println("Error writing to CST DOT file:", err)
		return
	}
	_, err = cstFile.WriteString(c.Dot)
	if err != nil {
		fmt.Println("Error writing to CST DOT file:", err)
		return
	}
	_, err = cstFile.WriteString("}\n")
	if err != nil {
		fmt.Println("Error writing to CST DOT file:", err)
		return
	}

	// Ejecutar Graphviz para generar la imagen
	cmd := exec.Command("dot", "-Tpng", "cst.dot", "-o", "cst.png")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error generating image:", err)
		return
	}
	c.MostrarImagenCST()

	fmt.Println("CST DOT file generated successfully.")
}

func (c *AST) MostrarImagenCST() {
	// Create a new HTML file
	file, err := os.Create("ImagenCST.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write the HTML header and embed the image
	fmt.Fprintln(file, `<!DOCTYPE html>
	<html>
	<head>
		<title>Reporte de Errores</title>
		<style>
			body {
				background-color: #333;
				color: white;
				font-family: sans-serif;
			}
			h1 {
				text-align: center;
			}
		</style>
	</head>
	<body>
		<h1>Árbol de Sintaxis Concreta</h1>
           <img src="cst.png" alt="Árbol de Sintaxis Concreta" style="max-width: 100%;">
       </body>
       </html>`)

	// Open the HTML file in the default web browser
	open.Start("ImagenCST.html")
}

func (c *AST) NuevaMatriz(nombre string, mutable bool, simbolo Symbol, d ...int) Matriz {
	var matriz Matriz
	matriz.Elements = c.crearMatriz(d)
	matriz.Name = nombre
	matriz.Mutable = mutable
	matriz.Symbols = simbolo
	matriz.TipoSimbolo = "Matriz"
	return matriz
}

// Función para crear una matriz anidada con dimensiones d
func (c *AST) crearMatriz(d []int) interface{} {
	if len(d) == 1 {
		return make([]interface{}, d[0])
	}

	result := make([]interface{}, d[0])
	for i := range result {
		result[i] = c.crearMatriz(d[1:])
	}
	return result
}

// Función para ingresar un valor en la matriz en la posición pos
func (c *AST) IngresarValor(mat *Matriz, pos []int, valor interface{}) {
	c.ingresarValorEnPosicion(mat.Elements, pos, valor)
}

// Función auxiliar para ingresar un valor en la posición pos de la matriz
func (c *AST) ingresarValorEnPosicion(data interface{}, pos []int, valor interface{}) {
	if len(pos) == 1 {
		data.([]interface{})[pos[0]] = valor
		return
	}

	c.ingresarValorEnPosicion(data.([]interface{})[pos[0]], pos[1:], valor)
}

// Función para obtener un valor de la matriz en la posición pos
func (c *AST) ObtenerValor(mat Matriz, pos []int) interface{} {
	return c.obtenerValorEnPosicion(mat.Elements, pos)
}

// Función auxiliar para obtener un valor en la posición pos de la matriz
func (c *AST) obtenerValorEnPosicion(data interface{}, pos []int) interface{} {
	if len(pos) == 1 {
		return data.([]interface{})[pos[0]]
	}

	return c.obtenerValorEnPosicion(data.([]interface{})[pos[0]], pos[1:])
}

// Función para llenar toda la matriz con un valor
func (c *AST) LlenarMatriz(mat *Matriz, valor interface{}) {
	c.llenarValores(mat.Elements, []int{}, valor)
}

// Función auxiliar para llenar los valores de la matriz con un valor
func (c *AST) llenarValores(data interface{}, pos []int, valor interface{}) {
	switch t := data.(type) {
	case []interface{}:
		for i := range t {
			c.llenarValores(t[i], append(pos, i), valor)
		}
	default:
		c.ingresarValorEnPosicion(data, pos, valor)
	}
}

// Función para imprimir todos los valores de la matriz
func (c *AST) ImprimirMatriz(mat Matriz) {
	c.imprimirValores(mat.Elements, []int{})
}

// Función para sustituir todos los valores de una matriz 2x2
func (c *AST) SustituirValores2(mat *Matriz, valor interface{}, dim1 int, dim2 int) {
	for i := 0; i < dim1; i++ {
		for j := 0; j < dim2; j++ {
			c.IngresarValor(mat, []int{i, j}, valor)
		}
	}
}

// Función para sustituir todos los valores de una matriz 3x3x3
func (c *AST) SustituirValores3(mat *Matriz, valor interface{}, dim1 int, dim2 int, dim3 int) {
	for i := 0; i < dim1; i++ {
		for j := 0; j < dim2; j++ {
			for k := 0; k < dim3; k++ {
				c.IngresarValor(mat, []int{i, j, k}, valor)
			}
		}
	}
}

// Función para sustituir todos los valores de una matriz 4x4x4x4
func (c *AST) SustituirValores4(mat *Matriz, valor interface{}, dim1 int, dim2 int, dim3 int, dim4 int) {
	for i := 0; i < dim1; i++ {
		for j := 0; j < dim2; j++ {
			for k := 0; k < dim3; k++ {
				for l := 0; l < dim4; l++ {
					c.IngresarValor(mat, []int{i, j, k, l}, valor)
				}
			}
		}
	}
}

// Función para sustituir todos los valores de una matriz 5x5x5x5x5
func (c *AST) SustituirValores5(mat *Matriz, valor interface{}, dim1 int, dim2 int, dim3 int, dim4 int, dim5 int) {
	for i := 0; i < dim1; i++ {
		for j := 0; j < dim2; j++ {
			for k := 0; k < dim3; k++ {
				for l := 0; l < dim4; l++ {
					for m := 0; m < dim5; m++ {
						c.IngresarValor(mat, []int{i, j, k, l, m}, valor)
					}
				}
			}
		}
	}
}

// Función auxiliar para imprimir los valores de la matriz
func (c *AST) imprimirValores(data interface{}, pos []int) {
	switch t := data.(type) {
	case []interface{}:
		for i, v := range t {
			c.imprimirValores(v, append(pos, i))
		}
	default:
		fmt.Printf("Valor en la posición %v: %v\n", pos, data)
	}
}

func (c *AST) ImprimirArreglovalores() {
	contadorpila := 0
	for nivel := c.Pila_Matriz_Val.Front(); nivel != nil; nivel = nivel.Next() {
		lista := nivel.Value.(*list.List)
		fmt.Println("Nivel de pila:", contadorpila)
		fmt.Println("----------------------")
		for elem := lista.Front(); elem != nil; elem = elem.Next() {
			valores := elem.Value.(Valores_Matriz)
			fmt.Printf("Tipo: %d\n", valores.Tipo)
			fmt.Printf("tamaño lista: %d\n", lista.Len())
			if valores.Elements != nil {
				fmt.Println("Valores Elements")
				fmt.Printf("Tamaño: %d\n", valores.Elements.Len())
				fmt.Println("Valores en la lista:")
				contador := 0
				for e := valores.Elements.Front(); e != nil; e = e.Next() {
					fmt.Printf("Elemento: %+v\n", e.Value)
					fmt.Printf("contadro: %+v\n", contador)
					contador++
				}
			}
			if valores.Matriztam != nil {
				fmt.Printf("Valor : %+v\n", valores.Valor)
				fmt.Println("Valores Matriztam")
				fmt.Printf("Tamaño: %d\n", valores.Matriztam.Len())
				fmt.Println("Valores en la lista:")
				for e := valores.Matriztam.Front(); e != nil; e = e.Next() {
					fmt.Printf("Elemento: %+v\n", e.Value)

				}
			}
		}
		contadorpila++
	}
}

func (a *AST) GuardarFuncion(funcion Funcion) {
	for e := a.Lista_Funciones.Front(); e != nil; e = e.Next() {
		if e.Value.(Funcion).Nombre == funcion.Nombre {
			Errores := Errores{
				Descripcion: "La Funcion que esta intentando guardar ya existe\n Variable: " + funcion.Nombre,
				Fila:        strconv.Itoa(e.Value.(Funcion).Lin),
				Columna:     strconv.Itoa(e.Value.(Funcion).Col),
				Tipo:        "Error Semantico",
				Ambito:      funcion.Nombre,
			}
			a.ErroresHTML(Errores)
			return
		}
	}
	a.Lista_Funciones.PushBack(funcion)
}

func (a *AST) GetFuncion(nombre string) *Funcion {
	for v := a.Lista_Funciones.Front(); v != nil; v = v.Next() {
		funcion := v.Value.(Funcion)
		if funcion.Nombre == nombre {
			return &funcion
		}
	}
	return nil
}

func (a *AST) GuardarStruct(stru Struc) {
	for e := a.Lista_Struct.Front(); e != nil; e = e.Next() {
		if e.Value.(Struc).Nombre == stru.Nombre {
			Errores := Errores{
				Descripcion: "El Strut que esta intentando guardar ya existe\n Variable: " + stru.Nombre,
				Fila:        strconv.Itoa(e.Value.(Struc).Lin),
				Columna:     strconv.Itoa(e.Value.(Struc).Col),
				Tipo:        "Error Semantico",
				Ambito:      stru.Nombre,
			}
			a.ErroresHTML(Errores)
			return
		}
	}
	a.Lista_Struct.PushBack(stru)
}

func (a *AST) GetStruc(nombre string) *Struc {
	for v := a.Lista_Struct.Front(); v != nil; v = v.Next() {
		variablestruc := v.Value.(Struc)
		if variablestruc.Nombre == nombre {
			return &variablestruc
		}
	}
	return nil
}

func (a *AST) GuardarVariableStruc(stru VariableStruct) {
	for e := a.Lista_Variables_Struct.Front(); e != nil; e = e.Next() {
		if e.Value.(VariableStruct).Nombre == stru.Nombre {
			Errores := Errores{
				Descripcion: "La variale Struct que esta intentando guardar ya existe en este ambito: \n Variable: " + stru.Nombre,
				Fila:        strconv.Itoa(e.Value.(VariableStruct).Lin),
				Columna:     strconv.Itoa(e.Value.(VariableStruct).Col),
				Tipo:        "Error Semantico",
				Ambito:      stru.Scope,
			}
			a.ErroresHTML(Errores)
			return
		}
	}
	a.Lista_Variables_Struct.PushBack(stru)
	a.Lista_Struct_HTML.PushBack(stru)
}

func (a *AST) GetVariableStruc(nombre string) *VariableStruct {
	for v := a.Lista_Variables_Struct.Front(); v != nil; v = v.Next() {
		variablestruc := v.Value.(VariableStruct)
		if variablestruc.Nombre == nombre {
			return &variablestruc
		}
	}
	return nil
}
func (a *AST) ActualizarVariableStruc(mariable *VariableStruct) {
	for e := a.Pila_Variables_Struct.Front(); e != nil; e = e.Next() {
		lista := e.Value.(*list.List)
		for v := lista.Front(); v != nil; v = v.Next() {
			valor := v.Value.(VariableStruct)
			if valor.Nombre == mariable.Nombre {
				for e1 := valor.Strukt.Variables.Front(); e1 != nil; e1 = e1.Next() {
					v1 := e1.Value.(Variable)
					for e2 := mariable.Strukt.Variables.Front(); e2 != nil; e2 = e2.Next() {
						v2 := e2.Value.(Variable)
						if v1.Name == v2.Name {
							v1.Symbols.Valor = v2.Symbols.Valor
							e1.Value = v1
							for e3 := a.Lista_Struct_HTML.Front(); e3 != nil; e3 = e3.Next() {
								v3 := e3.Value.(VariableStruct)
								v3.Strukt = mariable.Strukt
								e3.Value = v3
							}
							return
						}
					}
				}
			}
		}
	}
	Errores := Errores{
		Descripcion: "La variable que está intentando modificar no existe: \n Variable: " + mariable.Nombre,
		Fila:        strconv.Itoa(mariable.Lin),
		Columna:     strconv.Itoa(mariable.Col),
		Tipo:        "Error Semántico",
		Ambito:      mariable.Scope,
	}
	a.ErroresHTML(Errores)
}

func (a *AST) IsMain(cadena string) bool {
	cadena = strings.ToLower(cadena)
	return strings.Contains(cadena, "funcion")
}

func (a *AST) IsTempT(cadena string) bool {
	cadena = strings.ToLower(cadena)
	return strings.Contains(cadena, "t")
}
