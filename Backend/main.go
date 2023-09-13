package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"Backend/environment"
	"Backend/interfaces"
	"Backend/parser"

	"github.com/antlr4-go/antlr/v4"
)

type TreeShapeListener struct {
	*parser.BaseSwiftGrammarListener
	Code []interface{}
}

type MyErrorListener struct {
	*antlr.DefaultErrorListener
	errors []string
}

func NewMyErrorListener() *MyErrorListener {
	return new(MyErrorListener)
}

func (l *MyErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errors = append(l.errors, fmt.Sprintf("%d,%d,%s", line, column, msg))
}

func (l *MyErrorListener) ReportError(recognizer antlr.Lexer, e antlr.RecognitionException) {
	pos := recognizer.GetCharPositionInLine()
	line := recognizer.GetLine()
	l.errors = append(l.errors, fmt.Sprintf("%d,%d,%s", line, pos, e.GetMessage()))
}

func main() {
	http.Handle("/simbolos", corsMiddleware(http.HandlerFunc(handleSimbolos)))
	http.Handle("/ejecutar", corsMiddleware(http.HandlerFunc(handleEjecutar)))
	http.Handle("/errores", corsMiddleware(http.HandlerFunc(handleErrores)))
	http.Handle("/arbol", corsMiddleware(http.HandlerFunc(handleCst)))

	fmt.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleEjecutar(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la petición", http.StatusInternalServerError)
		return
	}

	var data map[string]string
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Error al procesar el cuerpo de la petición", http.StatusBadRequest)
		return
	}

	code := data["code"]
	result := ejecutar(code)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": result})
}

func ejecutar(code string) string {
	input := antlr.NewInputStream(code)
	lexer := parser.NewSwiftLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create a new instance of MyErrorListener
	errorListener := NewMyErrorListener()

	// Remove default error listeners and add our custom error listener
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	p := parser.NewSwiftGrammarParser(tokens)
	p.BuildParseTrees = true

	// Remove default error listeners and add our custom error listener
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	tree := p.S()
	var listener *TreeShapeListener = NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	Code := listener.Code
	var Ast environment.AST
	Ast.IniciarAmbito()
	for _, inst := range Code {
		if inst == nil {
			fmt.Println("Error: inst is nil")
			continue
		}
		instruction, ok := inst.(interfaces.Instruction)
		if !ok {
			fmt.Printf("Error: inst is not of type interfaces.Instruction (actual type: %T, value: %#v)\n", inst, inst)
			continue
		}
		instruction.Ejecutar(&Ast)
	}
	return Ast.GetPrint()
}

func handleSimbolos(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la petición", http.StatusInternalServerError)
		return
	}

	var data map[string]string
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Error al procesar el cuerpo de la petición", http.StatusBadRequest)
		return
	}

	code := data["code"]
	result := simbolos(code)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": result})
}

func simbolos(code string) string {
	input := antlr.NewInputStream(code)
	lexer := parser.NewSwiftLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create a new instance of MyErrorListener
	errorListener := NewMyErrorListener()

	// Remove default error listeners and add our custom error listener
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	p := parser.NewSwiftGrammarParser(tokens)
	p.BuildParseTrees = true

	// Remove default error listeners and add our custom error listener
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	tree := p.S()
	var listener *TreeShapeListener = NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	Code := listener.Code
	var Ast environment.AST
	Ast.IniciarAmbito()
	for _, inst := range Code {
		if inst == nil {
			fmt.Println("Error: inst is nil")
			continue
		}
		instruction, ok := inst.(interfaces.Instruction)
		if !ok {
			fmt.Printf("Error: inst is not of type interfaces.Instruction (actual type: %T, value: %#v)\n", inst, inst)
			continue
		}
		instruction.Ejecutar(&Ast)
	}

	Ast.TablaVariablesHTML()
	return Ast.GetPrint()
}

func handleErrores(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la petición", http.StatusInternalServerError)
		return
	}

	var data map[string]string
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Error al procesar el cuerpo de la petición", http.StatusBadRequest)
		return
	}

	code := data["code"]
	result := errores(code)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": result})
}

func errores(code string) string {
	input := antlr.NewInputStream(code)
	lexer := parser.NewSwiftLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create a new instance of MyErrorListener
	errorListener := NewMyErrorListener()

	// Remove default error listeners and add our custom error listener
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	p := parser.NewSwiftGrammarParser(tokens)
	p.BuildParseTrees = true

	// Remove default error listeners and add our custom error listener
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	tree := p.S()
	var listener *TreeShapeListener = NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	Code := listener.Code
	var Ast environment.AST
	Ast.IniciarAmbito()
	for _, inst := range Code {
		if inst == nil {
			fmt.Println("Error: inst is nil")
			continue
		}
		instruction, ok := inst.(interfaces.Instruction)
		if !ok {
			fmt.Printf("Error: inst is not of type interfaces.Instruction (actual type: %T, value: %#v)\n", inst, inst)
			continue
		}
		instruction.Ejecutar(&Ast)
	}

	if len(errorListener.errors) > 0 {
		for _, err := range errorListener.errors {
			subcadena := strings.Split(err, ",")
			Errores := environment.Errores{
				Descripcion: subcadena[2],
				Fila:        subcadena[0],
				Columna:     subcadena[1],
				Tipo:        "Error Sintactico",
				Ambito:      "",
			}
			Ast.ErroresHTML(Errores)
		}
	}
	Ast.TablaErroresHTML()
	return Ast.GetPrint()
}

func handleCst(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la petición", http.StatusInternalServerError)
		return
	}

	var data map[string]string
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Error al procesar el cuerpo de la petición", http.StatusBadRequest)
		return
	}

	code := data["code"]
	result := Cst(code)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": result})
}

func Cst(code string) string {
	input := antlr.NewInputStream(code)
	lexer := parser.NewSwiftLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create a new instance of MyErrorListener
	errorListener := NewMyErrorListener()

	// Remove default error listeners and add our custom error listener
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	p := parser.NewSwiftGrammarParser(tokens)
	p.BuildParseTrees = true

	// Remove default error listeners and add our custom error listener
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	tree := p.S()
	var listener *TreeShapeListener = NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	Code := listener.Code
	var Ast environment.AST
	Ast.IniciarAmbito()
	var contador int = 0
	for _, inst := range Code {
		if inst == nil {
			fmt.Println("Error: inst is nil")
			continue
		}
		instruction, ok := inst.(interfaces.Instruction)
		if !ok {
			fmt.Printf("Error: inst is not of type interfaces.Instruction (actual type: %T, value: %#v)\n", inst, inst)
			continue
		}
		instruction.Ejecutar(&Ast)
		contador++
	}

	var operands []*environment.Node
	for i := 0; i < contador; i++ {
		operands = append(operands, Ast.Pop())
	}
	Ast.Id += 1
	Ast.Push(&environment.Node{
		Id:       Ast.Id,
		Label:    "Instrucciones",
		Children: operands,
	})

	var operands2 []*environment.Node
	for i := 0; i < 1; i++ {
		operands2 = append(operands2, Ast.Pop())
	}
	Ast.Id += 1
	Ast.Push(&environment.Node{
		Id:       Ast.Id,
		Label:    "Root",
		Children: operands2,
	})

	Ast.Postorder(Ast.Pop())
	Ast.GenerateCST()
	return Ast.GetPrint()
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitS(ctx *parser.SContext) {
	this.Code = ctx.GetCode()
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
