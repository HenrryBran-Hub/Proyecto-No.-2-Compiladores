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

	// EnterListaexpresions is called when entering the listaexpresions production.
	EnterListaexpresions(c *ListaexpresionsContext)

	// EnterListaexpresion is called when entering the listaexpresion production.
	EnterListaexpresion(c *ListaexpresionContext)

	// EnterPrintstmt is called when entering the printstmt production.
	EnterPrintstmt(c *PrintstmtContext)

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

	// ExitListaexpresions is called when exiting the listaexpresions production.
	ExitListaexpresions(c *ListaexpresionsContext)

	// ExitListaexpresion is called when exiting the listaexpresion production.
	ExitListaexpresion(c *ListaexpresionContext)

	// ExitPrintstmt is called when exiting the printstmt production.
	ExitPrintstmt(c *PrintstmtContext)
}
