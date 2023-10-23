// Generated from c:/Users/Henrr/OneDrive/Escritorio/Compiladores/Proyecto No. 2/Proyecto No. 2 Compiladores/Optimizar/C3DGrammar.g4 by ANTLR 4.13.1

    import "Optimizar/interfaces"
    import "Optimizar/environment"
    import "Optimizar/expressions"
    import "Optimizar/instructions/datoscompuestos"
    import "Optimizar/instructions/datosprimitivos"
    import "Optimizar/instructions/funciones"
    import "Optimizar/instructions/sentencias"
    import "strings"
   

import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link C3DGrammarParser}.
 */
public interface C3DGrammarListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#start}.
	 * @param ctx the parse tree
	 */
	void enterStart(C3DGrammarParser.StartContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#start}.
	 * @param ctx the parse tree
	 */
	void exitStart(C3DGrammarParser.StartContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#encabezado}.
	 * @param ctx the parse tree
	 */
	void enterEncabezado(C3DGrammarParser.EncabezadoContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#encabezado}.
	 * @param ctx the parse tree
	 */
	void exitEncabezado(C3DGrammarParser.EncabezadoContext ctx);
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
	 * Enter a parse tree produced by {@link C3DGrammarParser#assign_stack}.
	 * @param ctx the parse tree
	 */
	void enterAssign_stack(C3DGrammarParser.Assign_stackContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#assign_stack}.
	 * @param ctx the parse tree
	 */
	void exitAssign_stack(C3DGrammarParser.Assign_stackContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#assign_heap}.
	 * @param ctx the parse tree
	 */
	void enterAssign_heap(C3DGrammarParser.Assign_heapContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#assign_heap}.
	 * @param ctx the parse tree
	 */
	void exitAssign_heap(C3DGrammarParser.Assign_heapContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#assign}.
	 * @param ctx the parse tree
	 */
	void enterAssign(C3DGrammarParser.AssignContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#assign}.
	 * @param ctx the parse tree
	 */
	void exitAssign(C3DGrammarParser.AssignContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#if_instr}.
	 * @param ctx the parse tree
	 */
	void enterIf_instr(C3DGrammarParser.If_instrContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#if_instr}.
	 * @param ctx the parse tree
	 */
	void exitIf_instr(C3DGrammarParser.If_instrContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#goto_instr}.
	 * @param ctx the parse tree
	 */
	void enterGoto_instr(C3DGrammarParser.Goto_instrContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#goto_instr}.
	 * @param ctx the parse tree
	 */
	void exitGoto_instr(C3DGrammarParser.Goto_instrContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#label_instr}.
	 * @param ctx the parse tree
	 */
	void enterLabel_instr(C3DGrammarParser.Label_instrContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#label_instr}.
	 * @param ctx the parse tree
	 */
	void exitLabel_instr(C3DGrammarParser.Label_instrContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#printf_instr}.
	 * @param ctx the parse tree
	 */
	void enterPrintf_instr(C3DGrammarParser.Printf_instrContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#printf_instr}.
	 * @param ctx the parse tree
	 */
	void exitPrintf_instr(C3DGrammarParser.Printf_instrContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#return_instr}.
	 * @param ctx the parse tree
	 */
	void enterReturn_instr(C3DGrammarParser.Return_instrContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#return_instr}.
	 * @param ctx the parse tree
	 */
	void exitReturn_instr(C3DGrammarParser.Return_instrContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#getHeap_instr}.
	 * @param ctx the parse tree
	 */
	void enterGetHeap_instr(C3DGrammarParser.GetHeap_instrContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#getHeap_instr}.
	 * @param ctx the parse tree
	 */
	void exitGetHeap_instr(C3DGrammarParser.GetHeap_instrContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#getStack_instr}.
	 * @param ctx the parse tree
	 */
	void enterGetStack_instr(C3DGrammarParser.GetStack_instrContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#getStack_instr}.
	 * @param ctx the parse tree
	 */
	void exitGetStack_instr(C3DGrammarParser.GetStack_instrContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#call_instr}.
	 * @param ctx the parse tree
	 */
	void enterCall_instr(C3DGrammarParser.Call_instrContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#call_instr}.
	 * @param ctx the parse tree
	 */
	void exitCall_instr(C3DGrammarParser.Call_instrContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#expr_print}.
	 * @param ctx the parse tree
	 */
	void enterExpr_print(C3DGrammarParser.Expr_printContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#expr_print}.
	 * @param ctx the parse tree
	 */
	void exitExpr_print(C3DGrammarParser.Expr_printContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#convert}.
	 * @param ctx the parse tree
	 */
	void enterConvert(C3DGrammarParser.ConvertContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#convert}.
	 * @param ctx the parse tree
	 */
	void exitConvert(C3DGrammarParser.ConvertContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(C3DGrammarParser.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(C3DGrammarParser.ExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#expr_rel}.
	 * @param ctx the parse tree
	 */
	void enterExpr_rel(C3DGrammarParser.Expr_relContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#expr_rel}.
	 * @param ctx the parse tree
	 */
	void exitExpr_rel(C3DGrammarParser.Expr_relContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#expr_arit}.
	 * @param ctx the parse tree
	 */
	void enterExpr_arit(C3DGrammarParser.Expr_aritContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#expr_arit}.
	 * @param ctx the parse tree
	 */
	void exitExpr_arit(C3DGrammarParser.Expr_aritContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#data_type}.
	 * @param ctx the parse tree
	 */
	void enterData_type(C3DGrammarParser.Data_typeContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#data_type}.
	 * @param ctx the parse tree
	 */
	void exitData_type(C3DGrammarParser.Data_typeContext ctx);
	/**
	 * Enter a parse tree produced by {@link C3DGrammarParser#expr_valor}.
	 * @param ctx the parse tree
	 */
	void enterExpr_valor(C3DGrammarParser.Expr_valorContext ctx);
	/**
	 * Exit a parse tree produced by {@link C3DGrammarParser#expr_valor}.
	 * @param ctx the parse tree
	 */
	void exitExpr_valor(C3DGrammarParser.Expr_valorContext ctx);
}