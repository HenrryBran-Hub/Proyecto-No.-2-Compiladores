package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"Optimizar/environmentc3d"
	"Optimizar/interfacesc3d"
	parser "Optimizar/parserc3d"

	"github.com/antlr4-go/antlr/v4"
)

type TreeShapeListener struct {
	*parser.BaseC3DGrammarListener
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
	http.Handle("/optimizar", corsMiddleware(http.HandlerFunc(handleOptimizar)))

	fmt.Println("Servidor escuchando en http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}

func handleOptimizar(w http.ResponseWriter, r *http.Request) {
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
	result := Optimizar(code)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": result})
}

func Optimizar(code string) string {
	input := antlr.NewInputStream(code)
	lexer := parser.NewC3DLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create a new instance of MyErrorListener
	errorListener := NewMyErrorListener()

	// Remove default error listeners and add our custom error listener
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	p := parser.NewC3DGrammarParser(tokens)
	p.BuildParseTrees = true

	// Remove default error listeners and add our custom error listener
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	tree := p.Z()
	var listener *TreeShapeListener = NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	Code := listener.Code
	var Ast environmentc3d.AST
	Ast.Iniciar()
	for _, inst := range Code {
		if inst == nil {
			fmt.Println("Error: inst is nil")
			continue
		}
		instruction, ok := inst.(interfacesc3d.Instruction)
		if !ok {
			fmt.Printf("Error: inst is not of type interfaces.Instruction (actual type: %T, value: %#v)\n", inst, inst)
			continue
		}
		instruction.Ejecutar(&Ast)
	}

	ConsoleOut := ""
	for e := Ast.FinalCode.Front(); e != nil; e = e.Next() {
		valor := e.Value
		ConsoleOut += valor.(string)
	}
	//fmt.Println(ConsoleOut)
	return ConsoleOut
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitZ(ctx *parser.ZContext) {
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
