package environmentc3d

import (
	list "container/list"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/skratchdot/open-golang/open"
)

type AST struct {
	Instructions       []interface{}
	Lista_Temporales   *list.List
	FinalCode          *list.List
	Lista_Optimizacion *list.List
}

type Temporal struct {
	Id    string
	Valor interface{}
	Tipo  TipoExpresion
	IsNeg bool
	IsOp  bool
}

type Optimizacion struct {
	Operacion     string
	OP_Optimizada string
	Regla         TipoRegla
	Fila          int
}

func NewAST(inst []interface{}) AST {
	ast := AST{
		Instructions: inst}
	return ast
}

func (a *AST) Iniciar() {
	a.Lista_Temporales = list.New()
	a.FinalCode = list.New()
}

func (a *AST) ActualizarTemporal(temp *Temporal) {
	for elem := a.Lista_Temporales.Front(); elem != nil; elem = elem.Next() {
		t := elem.Value.(Temporal)
		if t.Id == temp.Id {
			t.Valor = temp.Valor
			t.Tipo = temp.Tipo
			elem.Value = t
			return
		}
	}
}

func (a *AST) ObtenerTemporal(id string) *Temporal {
	for elem := a.Lista_Temporales.Front(); elem != nil; elem = elem.Next() {
		temp := elem.Value.(Temporal)
		if temp.Id == id {
			return &temp
		}
	}
	return nil
}

func (a *AST) DeterminarTipo(valor interface{}) TipoExpresion {
	cadena := valor.(string)
	if strings.Contains(cadena, ".") {
		return 1
	}

	if strings.Contains(cadena, "t") {
		return 2
	}

	return 0

}

func (a *AST) TablaOptimizacion() {
	// Create a new HTML file
	file, err := os.Create("TablaOptimizacion.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write the HTML header and table structure
	fmt.Fprintln(file, `<!DOCTYPE html>
	   <html>
	   <head>
	   	<title>Tabla de Optimizacion</title>
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
	   	<h1>Tabla de Optimizaciones</h1>
	   	<table>
	   		<tr>
	   			<th>No</th>
				<th>Operacion</th>
	   			<th>Operacion Optimizada</th>
	   			<th>Regla aplicada</th>
				<th>Descripcion</th>				
	   			<th>Fila</th>
	   		</tr>`)

	rowNumber := 1
	for e := a.Lista_Optimizacion.Front(); e != nil; e = e.Next() {
		variable := e.Value.(Optimizacion)
		fmt.Fprintf(file, `
   					<tr>
   						<td>%d</td>
						<td>%s</td>
   						<td>%s</td>
						<td>%d</td>
   						<td>%v</td>
   						<td>%d</td>
   					</tr>`,
			rowNumber,
			variable.Operacion,
			variable.OP_Optimizada,
			variable.Regla,
			TipoReglaDescripciones[variable.Regla],
			variable.Fila,
		)
		rowNumber++
	}

	// Write the HTML footer
	fmt.Fprintln(file, `
           </table>
       </body>
       </html>`)

	// Open the HTML file in the default web browser
	open.Start("TablaOptimizacion.html")
}
