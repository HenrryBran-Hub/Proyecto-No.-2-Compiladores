// Generated from c:/Users/Henrr/OneDrive/Escritorio/Compiladores/Proyecto No. 2/Proyecto No. 2 Compiladores/Optimizar/C3DGrammar.g4 by ANTLR 4.13.1

    import "Optimizar/interfacesc3d"    
    import "Optimizar/instructionsc3d"
    // import "Optimizar/environmentc3d"
    // import "Optimizar/expressionsc3d"
    // import "strings"
   

import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link C3DGrammarParser}.
 */
public interface C3DGrammarListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#z}.
	 * @param ctx the parse tree
	 */
	void enterZ(C3DGrammarParser.ZContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#z}.
	 * @param ctx the parse tree
	 */
	void exitZ(C3DGrammarParser.ZContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#encabezadoa}.
	 * @param ctx the parse tree
	 */
	void enterEncabezadoa(C3DGrammarParser.EncabezadoaContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#encabezadoa}.
	 * @param ctx the parse tree
	 */
	void exitEncabezadoa(C3DGrammarParser.EncabezadoaContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#temporales}.
	 * @param ctx the parse tree
	 */
	void enterTemporales(C3DGrammarParser.TemporalesContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#temporales}.
	 * @param ctx the parse tree
	 */
	void exitTemporales(C3DGrammarParser.TemporalesContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#blocktemporales}.
	 * @param ctx the parse tree
	 */
	void enterBlocktemporales(C3DGrammarParser.BlocktemporalesContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#blocktemporales}.
	 * @param ctx the parse tree
	 */
	void exitBlocktemporales(C3DGrammarParser.BlocktemporalesContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#bloquetemps}.
	 * @param ctx the parse tree
	 */
	void enterBloquetemps(C3DGrammarParser.BloquetempsContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#bloquetemps}.
	 * @param ctx the parse tree
	 */
	void exitBloquetemps(C3DGrammarParser.BloquetempsContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#blockfuncions}.
	 * @param ctx the parse tree
	 */
	void enterBlockfuncions(C3DGrammarParser.BlockfuncionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#blockfuncions}.
	 * @param ctx the parse tree
	 */
	void exitBlockfuncions(C3DGrammarParser.BlockfuncionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#bloquefunciones}.
	 * @param ctx the parse tree
	 */
	void enterBloquefunciones(C3DGrammarParser.BloquefuncionesContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#bloquefunciones}.
	 * @param ctx the parse tree
	 */
	void exitBloquefunciones(C3DGrammarParser.BloquefuncionesContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#funcionmain}.
	 * @param ctx the parse tree
	 */
	void enterFuncionmain(C3DGrammarParser.FuncionmainContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#funcionmain}.
	 * @param ctx the parse tree
	 */
	void exitFuncionmain(C3DGrammarParser.FuncionmainContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(C3DGrammarParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(C3DGrammarParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#instruction}.
	 * @param ctx the parse tree
	 */
	void enterInstruction(C3DGrammarParser.InstructionContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#instruction}.
	 * @param ctx the parse tree
	 */
	void exitInstruction(C3DGrammarParser.InstructionContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#head_op}.
	 * @param ctx the parse tree
	 */
	void enterHead_op(C3DGrammarParser.Head_opContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#head_op}.
	 * @param ctx the parse tree
	 */
	void exitHead_op(C3DGrammarParser.Head_opContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#stack_op}.
	 * @param ctx the parse tree
	 */
	void enterStack_op(C3DGrammarParser.Stack_opContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#stack_op}.
	 * @param ctx the parse tree
	 */
	void exitStack_op(C3DGrammarParser.Stack_opContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#printff}.
	 * @param ctx the parse tree
	 */
	void enterPrintff(C3DGrammarParser.PrintffContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#printff}.
	 * @param ctx the parse tree
	 */
	void exitPrintff(C3DGrammarParser.PrintffContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#embebida}.
	 * @param ctx the parse tree
	 */
	void enterEmbebida(C3DGrammarParser.EmbebidaContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#embebida}.
	 * @param ctx the parse tree
	 */
	void exitEmbebida(C3DGrammarParser.EmbebidaContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#tipodata}.
	 * @param ctx the parse tree
	 */
	void enterTipodata(C3DGrammarParser.TipodataContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#tipodata}.
	 * @param ctx the parse tree
	 */
	void exitTipodata(C3DGrammarParser.TipodataContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#operaritme}.
	 * @param ctx the parse tree
	 */
	void enterOperaritme(C3DGrammarParser.OperaritmeContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#operaritme}.
	 * @param ctx the parse tree
	 */
	void exitOperaritme(C3DGrammarParser.OperaritmeContext ctx);
}