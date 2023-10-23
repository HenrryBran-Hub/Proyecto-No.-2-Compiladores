// Code generated from C3DGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parserc3d // C3DGrammar
import "github.com/antlr4-go/antlr/v4"

// C3DGrammarListener is a complete listener for a parse tree produced by C3DGrammarParser.
type C3DGrammarListener interface {
	antlr.ParseTreeListener

	// EnterZ is called when entering the z production.
	EnterZ(c *ZContext)

	// EnterEncabezadoa is called when entering the encabezadoa production.
	EnterEncabezadoa(c *EncabezadoaContext)

	// EnterTemporales is called when entering the temporales production.
	EnterTemporales(c *TemporalesContext)

	// EnterBlocktemporales is called when entering the blocktemporales production.
	EnterBlocktemporales(c *BlocktemporalesContext)

	// EnterBloquetemps is called when entering the bloquetemps production.
	EnterBloquetemps(c *BloquetempsContext)

	// EnterBlockfuncions is called when entering the blockfuncions production.
	EnterBlockfuncions(c *BlockfuncionsContext)

	// EnterBloquefunciones is called when entering the bloquefunciones production.
	EnterBloquefunciones(c *BloquefuncionesContext)

	// EnterFuncionmain is called when entering the funcionmain production.
	EnterFuncionmain(c *FuncionmainContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterHead_op is called when entering the head_op production.
	EnterHead_op(c *Head_opContext)

	// EnterStack_op is called when entering the stack_op production.
	EnterStack_op(c *Stack_opContext)

	// EnterPrintff is called when entering the printff production.
	EnterPrintff(c *PrintffContext)

	// EnterEmbebida is called when entering the embebida production.
	EnterEmbebida(c *EmbebidaContext)

	// EnterTipodata is called when entering the tipodata production.
	EnterTipodata(c *TipodataContext)

	// ExitZ is called when exiting the z production.
	ExitZ(c *ZContext)

	// ExitEncabezadoa is called when exiting the encabezadoa production.
	ExitEncabezadoa(c *EncabezadoaContext)

	// ExitTemporales is called when exiting the temporales production.
	ExitTemporales(c *TemporalesContext)

	// ExitBlocktemporales is called when exiting the blocktemporales production.
	ExitBlocktemporales(c *BlocktemporalesContext)

	// ExitBloquetemps is called when exiting the bloquetemps production.
	ExitBloquetemps(c *BloquetempsContext)

	// ExitBlockfuncions is called when exiting the blockfuncions production.
	ExitBlockfuncions(c *BlockfuncionsContext)

	// ExitBloquefunciones is called when exiting the bloquefunciones production.
	ExitBloquefunciones(c *BloquefuncionesContext)

	// ExitFuncionmain is called when exiting the funcionmain production.
	ExitFuncionmain(c *FuncionmainContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitHead_op is called when exiting the head_op production.
	ExitHead_op(c *Head_opContext)

	// ExitStack_op is called when exiting the stack_op production.
	ExitStack_op(c *Stack_opContext)

	// ExitPrintff is called when exiting the printff production.
	ExitPrintff(c *PrintffContext)

	// ExitEmbebida is called when exiting the embebida production.
	ExitEmbebida(c *EmbebidaContext)

	// ExitTipodata is called when exiting the tipodata production.
	ExitTipodata(c *TipodataContext)
}
