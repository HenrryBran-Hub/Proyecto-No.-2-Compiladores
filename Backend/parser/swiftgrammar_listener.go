// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// SwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type SwiftGrammarListener interface {
	antlr.ParseTreeListener

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterBlockinterno is called when entering the blockinterno production.
	EnterBlockinterno(c *BlockinternoContext)

	// EnterInstructionint is called when entering the instructionint production.
	EnterInstructionint(c *InstructionintContext)

	// EnterDeclavarible is called when entering the declavarible production.
	EnterDeclavarible(c *DeclavaribleContext)

	// EnterDeclaconstante is called when entering the declaconstante production.
	EnterDeclaconstante(c *DeclaconstanteContext)

	// EnterAsignacionvariable is called when entering the asignacionvariable production.
	EnterAsignacionvariable(c *AsignacionvariableContext)

	// EnterTipodato is called when entering the tipodato production.
	EnterTipodato(c *TipodatoContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterSentenciaifelse is called when entering the sentenciaifelse production.
	EnterSentenciaifelse(c *SentenciaifelseContext)

	// EnterSwitchcontrol is called when entering the switchcontrol production.
	EnterSwitchcontrol(c *SwitchcontrolContext)

	// EnterBlockcase is called when entering the blockcase production.
	EnterBlockcase(c *BlockcaseContext)

	// EnterBloquecase is called when entering the bloquecase production.
	EnterBloquecase(c *BloquecaseContext)

	// EnterWhilecontrol is called when entering the whilecontrol production.
	EnterWhilecontrol(c *WhilecontrolContext)

	// EnterForcontrol is called when entering the forcontrol production.
	EnterForcontrol(c *ForcontrolContext)

	// EnterGuardcontrol is called when entering the guardcontrol production.
	EnterGuardcontrol(c *GuardcontrolContext)

	// EnterContinuee is called when entering the continuee production.
	EnterContinuee(c *ContinueeContext)

	// EnterBreakk is called when entering the breakk production.
	EnterBreakk(c *BreakkContext)

	// EnterRetornos is called when entering the retornos production.
	EnterRetornos(c *RetornosContext)

	// EnterVectorcontrol is called when entering the vectorcontrol production.
	EnterVectorcontrol(c *VectorcontrolContext)

	// EnterBlockparams is called when entering the blockparams production.
	EnterBlockparams(c *BlockparamsContext)

	// EnterBloqueparams is called when entering the bloqueparams production.
	EnterBloqueparams(c *BloqueparamsContext)

	// EnterVectoragregar is called when entering the vectoragregar production.
	EnterVectoragregar(c *VectoragregarContext)

	// EnterVectorremover is called when entering the vectorremover production.
	EnterVectorremover(c *VectorremoverContext)

	// EnterVectorvacio is called when entering the vectorvacio production.
	EnterVectorvacio(c *VectorvacioContext)

	// EnterVectorcount is called when entering the vectorcount production.
	EnterVectorcount(c *VectorcountContext)

	// EnterVectoraccess is called when entering the vectoraccess production.
	EnterVectoraccess(c *VectoraccessContext)

	// EnterMatrizcontrol is called when entering the matrizcontrol production.
	EnterMatrizcontrol(c *MatrizcontrolContext)

	// EnterTipomatriz is called when entering the tipomatriz production.
	EnterTipomatriz(c *TipomatrizContext)

	// EnterDefmatriz is called when entering the defmatriz production.
	EnterDefmatriz(c *DefmatrizContext)

	// EnterListavaloresmat is called when entering the listavaloresmat production.
	EnterListavaloresmat(c *ListavaloresmatContext)

	// EnterListavaloresmat2 is called when entering the listavaloresmat2 production.
	EnterListavaloresmat2(c *Listavaloresmat2Context)

	// EnterListaexpresions is called when entering the listaexpresions production.
	EnterListaexpresions(c *ListaexpresionsContext)

	// EnterListaexpresion is called when entering the listaexpresion production.
	EnterListaexpresion(c *ListaexpresionContext)

	// EnterSimplematriz is called when entering the simplematriz production.
	EnterSimplematriz(c *SimplematrizContext)

	// EnterListamatrizaddsubs is called when entering the listamatrizaddsubs production.
	EnterListamatrizaddsubs(c *ListamatrizaddsubsContext)

	// EnterListamatrizaddsub is called when entering the listamatrizaddsub production.
	EnterListamatrizaddsub(c *ListamatrizaddsubContext)

	// EnterStructcontrol is called when entering the structcontrol production.
	EnterStructcontrol(c *StructcontrolContext)

	// EnterListaatributos is called when entering the listaatributos production.
	EnterListaatributos(c *ListaatributosContext)

	// EnterListaatributo is called when entering the listaatributo production.
	EnterListaatributo(c *ListaatributoContext)

	// EnterStructexpr is called when entering the structexpr production.
	EnterStructexpr(c *StructexprContext)

	// EnterLdupla is called when entering the ldupla production.
	EnterLdupla(c *LduplaContext)

	// EnterLlamadastruct is called when entering the llamadastruct production.
	EnterLlamadastruct(c *LlamadastructContext)

	// EnterAsignacionparametrostruct is called when entering the asignacionparametrostruct production.
	EnterAsignacionparametrostruct(c *AsignacionparametrostructContext)

	// EnterLlamadafuncionstructcontrol is called when entering the llamadafuncionstructcontrol production.
	EnterLlamadafuncionstructcontrol(c *LlamadafuncionstructcontrolContext)

	// EnterLlamadafuncionstructcontrolret is called when entering the llamadafuncionstructcontrolret production.
	EnterLlamadafuncionstructcontrolret(c *LlamadafuncionstructcontrolretContext)

	// EnterFunciondeclaracioncontrol is called when entering the funciondeclaracioncontrol production.
	EnterFunciondeclaracioncontrol(c *FunciondeclaracioncontrolContext)

	// EnterListaparametro is called when entering the listaparametro production.
	EnterListaparametro(c *ListaparametroContext)

	// EnterFuncionllamadacontrol is called when entering the funcionllamadacontrol production.
	EnterFuncionllamadacontrol(c *FuncionllamadacontrolContext)

	// EnterFuncionllamadacontrolConRetorno is called when entering the funcionllamadacontrolConRetorno production.
	EnterFuncionllamadacontrolConRetorno(c *FuncionllamadacontrolConRetornoContext)

	// EnterListaparametrosllamada is called when entering the listaparametrosllamada production.
	EnterListaparametrosllamada(c *ListaparametrosllamadaContext)

	// EnterPrintstmt is called when entering the printstmt production.
	EnterPrintstmt(c *PrintstmtContext)

	// EnterIntembebida is called when entering the intembebida production.
	EnterIntembebida(c *IntembebidaContext)

	// EnterFloatembebida is called when entering the floatembebida production.
	EnterFloatembebida(c *FloatembebidaContext)

	// EnterStringembebida is called when entering the stringembebida production.
	EnterStringembebida(c *StringembebidaContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitBlockinterno is called when exiting the blockinterno production.
	ExitBlockinterno(c *BlockinternoContext)

	// ExitInstructionint is called when exiting the instructionint production.
	ExitInstructionint(c *InstructionintContext)

	// ExitDeclavarible is called when exiting the declavarible production.
	ExitDeclavarible(c *DeclavaribleContext)

	// ExitDeclaconstante is called when exiting the declaconstante production.
	ExitDeclaconstante(c *DeclaconstanteContext)

	// ExitAsignacionvariable is called when exiting the asignacionvariable production.
	ExitAsignacionvariable(c *AsignacionvariableContext)

	// ExitTipodato is called when exiting the tipodato production.
	ExitTipodato(c *TipodatoContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitSentenciaifelse is called when exiting the sentenciaifelse production.
	ExitSentenciaifelse(c *SentenciaifelseContext)

	// ExitSwitchcontrol is called when exiting the switchcontrol production.
	ExitSwitchcontrol(c *SwitchcontrolContext)

	// ExitBlockcase is called when exiting the blockcase production.
	ExitBlockcase(c *BlockcaseContext)

	// ExitBloquecase is called when exiting the bloquecase production.
	ExitBloquecase(c *BloquecaseContext)

	// ExitWhilecontrol is called when exiting the whilecontrol production.
	ExitWhilecontrol(c *WhilecontrolContext)

	// ExitForcontrol is called when exiting the forcontrol production.
	ExitForcontrol(c *ForcontrolContext)

	// ExitGuardcontrol is called when exiting the guardcontrol production.
	ExitGuardcontrol(c *GuardcontrolContext)

	// ExitContinuee is called when exiting the continuee production.
	ExitContinuee(c *ContinueeContext)

	// ExitBreakk is called when exiting the breakk production.
	ExitBreakk(c *BreakkContext)

	// ExitRetornos is called when exiting the retornos production.
	ExitRetornos(c *RetornosContext)

	// ExitVectorcontrol is called when exiting the vectorcontrol production.
	ExitVectorcontrol(c *VectorcontrolContext)

	// ExitBlockparams is called when exiting the blockparams production.
	ExitBlockparams(c *BlockparamsContext)

	// ExitBloqueparams is called when exiting the bloqueparams production.
	ExitBloqueparams(c *BloqueparamsContext)

	// ExitVectoragregar is called when exiting the vectoragregar production.
	ExitVectoragregar(c *VectoragregarContext)

	// ExitVectorremover is called when exiting the vectorremover production.
	ExitVectorremover(c *VectorremoverContext)

	// ExitVectorvacio is called when exiting the vectorvacio production.
	ExitVectorvacio(c *VectorvacioContext)

	// ExitVectorcount is called when exiting the vectorcount production.
	ExitVectorcount(c *VectorcountContext)

	// ExitVectoraccess is called when exiting the vectoraccess production.
	ExitVectoraccess(c *VectoraccessContext)

	// ExitMatrizcontrol is called when exiting the matrizcontrol production.
	ExitMatrizcontrol(c *MatrizcontrolContext)

	// ExitTipomatriz is called when exiting the tipomatriz production.
	ExitTipomatriz(c *TipomatrizContext)

	// ExitDefmatriz is called when exiting the defmatriz production.
	ExitDefmatriz(c *DefmatrizContext)

	// ExitListavaloresmat is called when exiting the listavaloresmat production.
	ExitListavaloresmat(c *ListavaloresmatContext)

	// ExitListavaloresmat2 is called when exiting the listavaloresmat2 production.
	ExitListavaloresmat2(c *Listavaloresmat2Context)

	// ExitListaexpresions is called when exiting the listaexpresions production.
	ExitListaexpresions(c *ListaexpresionsContext)

	// ExitListaexpresion is called when exiting the listaexpresion production.
	ExitListaexpresion(c *ListaexpresionContext)

	// ExitSimplematriz is called when exiting the simplematriz production.
	ExitSimplematriz(c *SimplematrizContext)

	// ExitListamatrizaddsubs is called when exiting the listamatrizaddsubs production.
	ExitListamatrizaddsubs(c *ListamatrizaddsubsContext)

	// ExitListamatrizaddsub is called when exiting the listamatrizaddsub production.
	ExitListamatrizaddsub(c *ListamatrizaddsubContext)

	// ExitStructcontrol is called when exiting the structcontrol production.
	ExitStructcontrol(c *StructcontrolContext)

	// ExitListaatributos is called when exiting the listaatributos production.
	ExitListaatributos(c *ListaatributosContext)

	// ExitListaatributo is called when exiting the listaatributo production.
	ExitListaatributo(c *ListaatributoContext)

	// ExitStructexpr is called when exiting the structexpr production.
	ExitStructexpr(c *StructexprContext)

	// ExitLdupla is called when exiting the ldupla production.
	ExitLdupla(c *LduplaContext)

	// ExitLlamadastruct is called when exiting the llamadastruct production.
	ExitLlamadastruct(c *LlamadastructContext)

	// ExitAsignacionparametrostruct is called when exiting the asignacionparametrostruct production.
	ExitAsignacionparametrostruct(c *AsignacionparametrostructContext)

	// ExitLlamadafuncionstructcontrol is called when exiting the llamadafuncionstructcontrol production.
	ExitLlamadafuncionstructcontrol(c *LlamadafuncionstructcontrolContext)

	// ExitLlamadafuncionstructcontrolret is called when exiting the llamadafuncionstructcontrolret production.
	ExitLlamadafuncionstructcontrolret(c *LlamadafuncionstructcontrolretContext)

	// ExitFunciondeclaracioncontrol is called when exiting the funciondeclaracioncontrol production.
	ExitFunciondeclaracioncontrol(c *FunciondeclaracioncontrolContext)

	// ExitListaparametro is called when exiting the listaparametro production.
	ExitListaparametro(c *ListaparametroContext)

	// ExitFuncionllamadacontrol is called when exiting the funcionllamadacontrol production.
	ExitFuncionllamadacontrol(c *FuncionllamadacontrolContext)

	// ExitFuncionllamadacontrolConRetorno is called when exiting the funcionllamadacontrolConRetorno production.
	ExitFuncionllamadacontrolConRetorno(c *FuncionllamadacontrolConRetornoContext)

	// ExitListaparametrosllamada is called when exiting the listaparametrosllamada production.
	ExitListaparametrosllamada(c *ListaparametrosllamadaContext)

	// ExitPrintstmt is called when exiting the printstmt production.
	ExitPrintstmt(c *PrintstmtContext)

	// ExitIntembebida is called when exiting the intembebida production.
	ExitIntembebida(c *IntembebidaContext)

	// ExitFloatembebida is called when exiting the floatembebida production.
	ExitFloatembebida(c *FloatembebidaContext)

	// ExitStringembebida is called when exiting the stringembebida production.
	ExitStringembebida(c *StringembebidaContext)
}