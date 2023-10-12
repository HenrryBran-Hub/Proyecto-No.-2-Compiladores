// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// BaseSwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type BaseSwiftGrammarListener struct{}

var _ SwiftGrammarListener = &BaseSwiftGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSwiftGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSwiftGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSwiftGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSwiftGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterS is called when production s is entered.
func (s *BaseSwiftGrammarListener) EnterS(ctx *SContext) {}

// ExitS is called when production s is exited.
func (s *BaseSwiftGrammarListener) ExitS(ctx *SContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseSwiftGrammarListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseSwiftGrammarListener) ExitBlock(ctx *BlockContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseSwiftGrammarListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseSwiftGrammarListener) ExitInstruction(ctx *InstructionContext) {}

// EnterBlockinterno is called when production blockinterno is entered.
func (s *BaseSwiftGrammarListener) EnterBlockinterno(ctx *BlockinternoContext) {}

// ExitBlockinterno is called when production blockinterno is exited.
func (s *BaseSwiftGrammarListener) ExitBlockinterno(ctx *BlockinternoContext) {}

// EnterInstructionint is called when production instructionint is entered.
func (s *BaseSwiftGrammarListener) EnterInstructionint(ctx *InstructionintContext) {}

// ExitInstructionint is called when production instructionint is exited.
func (s *BaseSwiftGrammarListener) ExitInstructionint(ctx *InstructionintContext) {}

// EnterDeclavarible is called when production declavarible is entered.
func (s *BaseSwiftGrammarListener) EnterDeclavarible(ctx *DeclavaribleContext) {}

// ExitDeclavarible is called when production declavarible is exited.
func (s *BaseSwiftGrammarListener) ExitDeclavarible(ctx *DeclavaribleContext) {}

// EnterDeclaconstante is called when production declaconstante is entered.
func (s *BaseSwiftGrammarListener) EnterDeclaconstante(ctx *DeclaconstanteContext) {}

// ExitDeclaconstante is called when production declaconstante is exited.
func (s *BaseSwiftGrammarListener) ExitDeclaconstante(ctx *DeclaconstanteContext) {}

// EnterAsignacionvariable is called when production asignacionvariable is entered.
func (s *BaseSwiftGrammarListener) EnterAsignacionvariable(ctx *AsignacionvariableContext) {}

// ExitAsignacionvariable is called when production asignacionvariable is exited.
func (s *BaseSwiftGrammarListener) ExitAsignacionvariable(ctx *AsignacionvariableContext) {}

// EnterTipodato is called when production tipodato is entered.
func (s *BaseSwiftGrammarListener) EnterTipodato(ctx *TipodatoContext) {}

// ExitTipodato is called when production tipodato is exited.
func (s *BaseSwiftGrammarListener) ExitTipodato(ctx *TipodatoContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSwiftGrammarListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSwiftGrammarListener) ExitExpr(ctx *ExprContext) {}

// EnterSentenciaifelse is called when production sentenciaifelse is entered.
func (s *BaseSwiftGrammarListener) EnterSentenciaifelse(ctx *SentenciaifelseContext) {}

// ExitSentenciaifelse is called when production sentenciaifelse is exited.
func (s *BaseSwiftGrammarListener) ExitSentenciaifelse(ctx *SentenciaifelseContext) {}

// EnterSwitchcontrol is called when production switchcontrol is entered.
func (s *BaseSwiftGrammarListener) EnterSwitchcontrol(ctx *SwitchcontrolContext) {}

// ExitSwitchcontrol is called when production switchcontrol is exited.
func (s *BaseSwiftGrammarListener) ExitSwitchcontrol(ctx *SwitchcontrolContext) {}

// EnterBlockcase is called when production blockcase is entered.
func (s *BaseSwiftGrammarListener) EnterBlockcase(ctx *BlockcaseContext) {}

// ExitBlockcase is called when production blockcase is exited.
func (s *BaseSwiftGrammarListener) ExitBlockcase(ctx *BlockcaseContext) {}

// EnterBloquecase is called when production bloquecase is entered.
func (s *BaseSwiftGrammarListener) EnterBloquecase(ctx *BloquecaseContext) {}

// ExitBloquecase is called when production bloquecase is exited.
func (s *BaseSwiftGrammarListener) ExitBloquecase(ctx *BloquecaseContext) {}

// EnterWhilecontrol is called when production whilecontrol is entered.
func (s *BaseSwiftGrammarListener) EnterWhilecontrol(ctx *WhilecontrolContext) {}

// ExitWhilecontrol is called when production whilecontrol is exited.
func (s *BaseSwiftGrammarListener) ExitWhilecontrol(ctx *WhilecontrolContext) {}

// EnterForcontrol is called when production forcontrol is entered.
func (s *BaseSwiftGrammarListener) EnterForcontrol(ctx *ForcontrolContext) {}

// ExitForcontrol is called when production forcontrol is exited.
func (s *BaseSwiftGrammarListener) ExitForcontrol(ctx *ForcontrolContext) {}

// EnterGuardcontrol is called when production guardcontrol is entered.
func (s *BaseSwiftGrammarListener) EnterGuardcontrol(ctx *GuardcontrolContext) {}

// ExitGuardcontrol is called when production guardcontrol is exited.
func (s *BaseSwiftGrammarListener) ExitGuardcontrol(ctx *GuardcontrolContext) {}

// EnterContinuee is called when production continuee is entered.
func (s *BaseSwiftGrammarListener) EnterContinuee(ctx *ContinueeContext) {}

// ExitContinuee is called when production continuee is exited.
func (s *BaseSwiftGrammarListener) ExitContinuee(ctx *ContinueeContext) {}

// EnterBreakk is called when production breakk is entered.
func (s *BaseSwiftGrammarListener) EnterBreakk(ctx *BreakkContext) {}

// ExitBreakk is called when production breakk is exited.
func (s *BaseSwiftGrammarListener) ExitBreakk(ctx *BreakkContext) {}

// EnterRetornos is called when production retornos is entered.
func (s *BaseSwiftGrammarListener) EnterRetornos(ctx *RetornosContext) {}

// ExitRetornos is called when production retornos is exited.
func (s *BaseSwiftGrammarListener) ExitRetornos(ctx *RetornosContext) {}

// EnterVectorcontrol is called when production vectorcontrol is entered.
func (s *BaseSwiftGrammarListener) EnterVectorcontrol(ctx *VectorcontrolContext) {}

// ExitVectorcontrol is called when production vectorcontrol is exited.
func (s *BaseSwiftGrammarListener) ExitVectorcontrol(ctx *VectorcontrolContext) {}

// EnterBlockparams is called when production blockparams is entered.
func (s *BaseSwiftGrammarListener) EnterBlockparams(ctx *BlockparamsContext) {}

// ExitBlockparams is called when production blockparams is exited.
func (s *BaseSwiftGrammarListener) ExitBlockparams(ctx *BlockparamsContext) {}

// EnterBloqueparams is called when production bloqueparams is entered.
func (s *BaseSwiftGrammarListener) EnterBloqueparams(ctx *BloqueparamsContext) {}

// ExitBloqueparams is called when production bloqueparams is exited.
func (s *BaseSwiftGrammarListener) ExitBloqueparams(ctx *BloqueparamsContext) {}

// EnterVectoragregar is called when production vectoragregar is entered.
func (s *BaseSwiftGrammarListener) EnterVectoragregar(ctx *VectoragregarContext) {}

// ExitVectoragregar is called when production vectoragregar is exited.
func (s *BaseSwiftGrammarListener) ExitVectoragregar(ctx *VectoragregarContext) {}

// EnterVectorremover is called when production vectorremover is entered.
func (s *BaseSwiftGrammarListener) EnterVectorremover(ctx *VectorremoverContext) {}

// ExitVectorremover is called when production vectorremover is exited.
func (s *BaseSwiftGrammarListener) ExitVectorremover(ctx *VectorremoverContext) {}

// EnterVectorvacio is called when production vectorvacio is entered.
func (s *BaseSwiftGrammarListener) EnterVectorvacio(ctx *VectorvacioContext) {}

// ExitVectorvacio is called when production vectorvacio is exited.
func (s *BaseSwiftGrammarListener) ExitVectorvacio(ctx *VectorvacioContext) {}

// EnterVectorcount is called when production vectorcount is entered.
func (s *BaseSwiftGrammarListener) EnterVectorcount(ctx *VectorcountContext) {}

// ExitVectorcount is called when production vectorcount is exited.
func (s *BaseSwiftGrammarListener) ExitVectorcount(ctx *VectorcountContext) {}

// EnterVectoraccess is called when production vectoraccess is entered.
func (s *BaseSwiftGrammarListener) EnterVectoraccess(ctx *VectoraccessContext) {}

// ExitVectoraccess is called when production vectoraccess is exited.
func (s *BaseSwiftGrammarListener) ExitVectoraccess(ctx *VectoraccessContext) {}

// EnterMatrizcontrol is called when production matrizcontrol is entered.
func (s *BaseSwiftGrammarListener) EnterMatrizcontrol(ctx *MatrizcontrolContext) {}

// ExitMatrizcontrol is called when production matrizcontrol is exited.
func (s *BaseSwiftGrammarListener) ExitMatrizcontrol(ctx *MatrizcontrolContext) {}

// EnterTipomatriz is called when production tipomatriz is entered.
func (s *BaseSwiftGrammarListener) EnterTipomatriz(ctx *TipomatrizContext) {}

// ExitTipomatriz is called when production tipomatriz is exited.
func (s *BaseSwiftGrammarListener) ExitTipomatriz(ctx *TipomatrizContext) {}

// EnterDefmatriz is called when production defmatriz is entered.
func (s *BaseSwiftGrammarListener) EnterDefmatriz(ctx *DefmatrizContext) {}

// ExitDefmatriz is called when production defmatriz is exited.
func (s *BaseSwiftGrammarListener) ExitDefmatriz(ctx *DefmatrizContext) {}

// EnterListavaloresmat is called when production listavaloresmat is entered.
func (s *BaseSwiftGrammarListener) EnterListavaloresmat(ctx *ListavaloresmatContext) {}

// ExitListavaloresmat is called when production listavaloresmat is exited.
func (s *BaseSwiftGrammarListener) ExitListavaloresmat(ctx *ListavaloresmatContext) {}

// EnterListavaloresmat2 is called when production listavaloresmat2 is entered.
func (s *BaseSwiftGrammarListener) EnterListavaloresmat2(ctx *Listavaloresmat2Context) {}

// ExitListavaloresmat2 is called when production listavaloresmat2 is exited.
func (s *BaseSwiftGrammarListener) ExitListavaloresmat2(ctx *Listavaloresmat2Context) {}

// EnterListaexpresions is called when production listaexpresions is entered.
func (s *BaseSwiftGrammarListener) EnterListaexpresions(ctx *ListaexpresionsContext) {}

// ExitListaexpresions is called when production listaexpresions is exited.
func (s *BaseSwiftGrammarListener) ExitListaexpresions(ctx *ListaexpresionsContext) {}

// EnterListaexpresion is called when production listaexpresion is entered.
func (s *BaseSwiftGrammarListener) EnterListaexpresion(ctx *ListaexpresionContext) {}

// ExitListaexpresion is called when production listaexpresion is exited.
func (s *BaseSwiftGrammarListener) ExitListaexpresion(ctx *ListaexpresionContext) {}

// EnterSimplematriz is called when production simplematriz is entered.
func (s *BaseSwiftGrammarListener) EnterSimplematriz(ctx *SimplematrizContext) {}

// ExitSimplematriz is called when production simplematriz is exited.
func (s *BaseSwiftGrammarListener) ExitSimplematriz(ctx *SimplematrizContext) {}

// EnterListamatrizaddsubs is called when production listamatrizaddsubs is entered.
func (s *BaseSwiftGrammarListener) EnterListamatrizaddsubs(ctx *ListamatrizaddsubsContext) {}

// ExitListamatrizaddsubs is called when production listamatrizaddsubs is exited.
func (s *BaseSwiftGrammarListener) ExitListamatrizaddsubs(ctx *ListamatrizaddsubsContext) {}

// EnterListamatrizaddsub is called when production listamatrizaddsub is entered.
func (s *BaseSwiftGrammarListener) EnterListamatrizaddsub(ctx *ListamatrizaddsubContext) {}

// ExitListamatrizaddsub is called when production listamatrizaddsub is exited.
func (s *BaseSwiftGrammarListener) ExitListamatrizaddsub(ctx *ListamatrizaddsubContext) {}

// EnterFunciondeclaracioncontrol is called when production funciondeclaracioncontrol is entered.
func (s *BaseSwiftGrammarListener) EnterFunciondeclaracioncontrol(ctx *FunciondeclaracioncontrolContext) {
}

// ExitFunciondeclaracioncontrol is called when production funciondeclaracioncontrol is exited.
func (s *BaseSwiftGrammarListener) ExitFunciondeclaracioncontrol(ctx *FunciondeclaracioncontrolContext) {
}

// EnterFuncionllamadacontrol is called when production funcionllamadacontrol is entered.
func (s *BaseSwiftGrammarListener) EnterFuncionllamadacontrol(ctx *FuncionllamadacontrolContext) {}

// ExitFuncionllamadacontrol is called when production funcionllamadacontrol is exited.
func (s *BaseSwiftGrammarListener) ExitFuncionllamadacontrol(ctx *FuncionllamadacontrolContext) {}

// EnterFuncionllamadacontrolConRetorno is called when production funcionllamadacontrolConRetorno is entered.
func (s *BaseSwiftGrammarListener) EnterFuncionllamadacontrolConRetorno(ctx *FuncionllamadacontrolConRetornoContext) {
}

// ExitFuncionllamadacontrolConRetorno is called when production funcionllamadacontrolConRetorno is exited.
func (s *BaseSwiftGrammarListener) ExitFuncionllamadacontrolConRetorno(ctx *FuncionllamadacontrolConRetornoContext) {
}

// EnterPrintstmt is called when production printstmt is entered.
func (s *BaseSwiftGrammarListener) EnterPrintstmt(ctx *PrintstmtContext) {}

// ExitPrintstmt is called when production printstmt is exited.
func (s *BaseSwiftGrammarListener) ExitPrintstmt(ctx *PrintstmtContext) {}

// EnterIntembebida is called when production intembebida is entered.
func (s *BaseSwiftGrammarListener) EnterIntembebida(ctx *IntembebidaContext) {}

// ExitIntembebida is called when production intembebida is exited.
func (s *BaseSwiftGrammarListener) ExitIntembebida(ctx *IntembebidaContext) {}

// EnterFloatembebida is called when production floatembebida is entered.
func (s *BaseSwiftGrammarListener) EnterFloatembebida(ctx *FloatembebidaContext) {}

// ExitFloatembebida is called when production floatembebida is exited.
func (s *BaseSwiftGrammarListener) ExitFloatembebida(ctx *FloatembebidaContext) {}

// EnterStringembebida is called when production stringembebida is entered.
func (s *BaseSwiftGrammarListener) EnterStringembebida(ctx *StringembebidaContext) {}

// ExitStringembebida is called when production stringembebida is exited.
func (s *BaseSwiftGrammarListener) ExitStringembebida(ctx *StringembebidaContext) {}
