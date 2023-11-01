// Generated from c:/Users/Henrr/OneDrive/Escritorio/Compiladores/Proyecto No. 2/Proyecto No. 2 Compiladores/Backend/SwiftGrammar.g4 by ANTLR 4.13.1

    import "Backend/interfaces"
    import "Backend/environment"
    import "Backend/expressions"
    import "Backend/instructions/datoscompuestos"
    import "Backend/instructions/datosprimitivos"
    import "Backend/instructions/funciones"
    import "Backend/instructions/sentencias"
    import "strings"
   

import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link SwiftGrammarParser}.
 */
public interface SwiftGrammarListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#s}.
	 * @param ctx the parse tree
	 */
	void enterS(SwiftGrammarParser.SContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#s}.
	 * @param ctx the parse tree
	 */
	void exitS(SwiftGrammarParser.SContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(SwiftGrammarParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(SwiftGrammarParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#instruction}.
	 * @param ctx the parse tree
	 */
	void enterInstruction(SwiftGrammarParser.InstructionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#instruction}.
	 * @param ctx the parse tree
	 */
	void exitInstruction(SwiftGrammarParser.InstructionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#blockinterno}.
	 * @param ctx the parse tree
	 */
	void enterBlockinterno(SwiftGrammarParser.BlockinternoContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#blockinterno}.
	 * @param ctx the parse tree
	 */
	void exitBlockinterno(SwiftGrammarParser.BlockinternoContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#instructionint}.
	 * @param ctx the parse tree
	 */
	void enterInstructionint(SwiftGrammarParser.InstructionintContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#instructionint}.
	 * @param ctx the parse tree
	 */
	void exitInstructionint(SwiftGrammarParser.InstructionintContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#declavarible}.
	 * @param ctx the parse tree
	 */
	void enterDeclavarible(SwiftGrammarParser.DeclavaribleContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#declavarible}.
	 * @param ctx the parse tree
	 */
	void exitDeclavarible(SwiftGrammarParser.DeclavaribleContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#declaconstante}.
	 * @param ctx the parse tree
	 */
	void enterDeclaconstante(SwiftGrammarParser.DeclaconstanteContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#declaconstante}.
	 * @param ctx the parse tree
	 */
	void exitDeclaconstante(SwiftGrammarParser.DeclaconstanteContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#asignacionvariable}.
	 * @param ctx the parse tree
	 */
	void enterAsignacionvariable(SwiftGrammarParser.AsignacionvariableContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#asignacionvariable}.
	 * @param ctx the parse tree
	 */
	void exitAsignacionvariable(SwiftGrammarParser.AsignacionvariableContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#tipodato}.
	 * @param ctx the parse tree
	 */
	void enterTipodato(SwiftGrammarParser.TipodatoContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#tipodato}.
	 * @param ctx the parse tree
	 */
	void exitTipodato(SwiftGrammarParser.TipodatoContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr(SwiftGrammarParser.ExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr(SwiftGrammarParser.ExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#sentenciaifelse}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaifelse(SwiftGrammarParser.SentenciaifelseContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#sentenciaifelse}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaifelse(SwiftGrammarParser.SentenciaifelseContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#switchcontrol}.
	 * @param ctx the parse tree
	 */
	void enterSwitchcontrol(SwiftGrammarParser.SwitchcontrolContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#switchcontrol}.
	 * @param ctx the parse tree
	 */
	void exitSwitchcontrol(SwiftGrammarParser.SwitchcontrolContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#blockcase}.
	 * @param ctx the parse tree
	 */
	void enterBlockcase(SwiftGrammarParser.BlockcaseContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#blockcase}.
	 * @param ctx the parse tree
	 */
	void exitBlockcase(SwiftGrammarParser.BlockcaseContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#bloquecase}.
	 * @param ctx the parse tree
	 */
	void enterBloquecase(SwiftGrammarParser.BloquecaseContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#bloquecase}.
	 * @param ctx the parse tree
	 */
	void exitBloquecase(SwiftGrammarParser.BloquecaseContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#whilecontrol}.
	 * @param ctx the parse tree
	 */
	void enterWhilecontrol(SwiftGrammarParser.WhilecontrolContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#whilecontrol}.
	 * @param ctx the parse tree
	 */
	void exitWhilecontrol(SwiftGrammarParser.WhilecontrolContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#forcontrol}.
	 * @param ctx the parse tree
	 */
	void enterForcontrol(SwiftGrammarParser.ForcontrolContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#forcontrol}.
	 * @param ctx the parse tree
	 */
	void exitForcontrol(SwiftGrammarParser.ForcontrolContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#guardcontrol}.
	 * @param ctx the parse tree
	 */
	void enterGuardcontrol(SwiftGrammarParser.GuardcontrolContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#guardcontrol}.
	 * @param ctx the parse tree
	 */
	void exitGuardcontrol(SwiftGrammarParser.GuardcontrolContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#continuee}.
	 * @param ctx the parse tree
	 */
	void enterContinuee(SwiftGrammarParser.ContinueeContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#continuee}.
	 * @param ctx the parse tree
	 */
	void exitContinuee(SwiftGrammarParser.ContinueeContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#breakk}.
	 * @param ctx the parse tree
	 */
	void enterBreakk(SwiftGrammarParser.BreakkContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#breakk}.
	 * @param ctx the parse tree
	 */
	void exitBreakk(SwiftGrammarParser.BreakkContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#retornos}.
	 * @param ctx the parse tree
	 */
	void enterRetornos(SwiftGrammarParser.RetornosContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#retornos}.
	 * @param ctx the parse tree
	 */
	void exitRetornos(SwiftGrammarParser.RetornosContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#vectorcontrol}.
	 * @param ctx the parse tree
	 */
	void enterVectorcontrol(SwiftGrammarParser.VectorcontrolContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#vectorcontrol}.
	 * @param ctx the parse tree
	 */
	void exitVectorcontrol(SwiftGrammarParser.VectorcontrolContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#blockparams}.
	 * @param ctx the parse tree
	 */
	void enterBlockparams(SwiftGrammarParser.BlockparamsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#blockparams}.
	 * @param ctx the parse tree
	 */
	void exitBlockparams(SwiftGrammarParser.BlockparamsContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#bloqueparams}.
	 * @param ctx the parse tree
	 */
	void enterBloqueparams(SwiftGrammarParser.BloqueparamsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#bloqueparams}.
	 * @param ctx the parse tree
	 */
	void exitBloqueparams(SwiftGrammarParser.BloqueparamsContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#vectoragregar}.
	 * @param ctx the parse tree
	 */
	void enterVectoragregar(SwiftGrammarParser.VectoragregarContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#vectoragregar}.
	 * @param ctx the parse tree
	 */
	void exitVectoragregar(SwiftGrammarParser.VectoragregarContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#vectorremover}.
	 * @param ctx the parse tree
	 */
	void enterVectorremover(SwiftGrammarParser.VectorremoverContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#vectorremover}.
	 * @param ctx the parse tree
	 */
	void exitVectorremover(SwiftGrammarParser.VectorremoverContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#vectorvacio}.
	 * @param ctx the parse tree
	 */
	void enterVectorvacio(SwiftGrammarParser.VectorvacioContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#vectorvacio}.
	 * @param ctx the parse tree
	 */
	void exitVectorvacio(SwiftGrammarParser.VectorvacioContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#vectorcount}.
	 * @param ctx the parse tree
	 */
	void enterVectorcount(SwiftGrammarParser.VectorcountContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#vectorcount}.
	 * @param ctx the parse tree
	 */
	void exitVectorcount(SwiftGrammarParser.VectorcountContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#vectoraccess}.
	 * @param ctx the parse tree
	 */
	void enterVectoraccess(SwiftGrammarParser.VectoraccessContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#vectoraccess}.
	 * @param ctx the parse tree
	 */
	void exitVectoraccess(SwiftGrammarParser.VectoraccessContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#matrizcontrol}.
	 * @param ctx the parse tree
	 */
	void enterMatrizcontrol(SwiftGrammarParser.MatrizcontrolContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#matrizcontrol}.
	 * @param ctx the parse tree
	 */
	void exitMatrizcontrol(SwiftGrammarParser.MatrizcontrolContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#tipomatriz}.
	 * @param ctx the parse tree
	 */
	void enterTipomatriz(SwiftGrammarParser.TipomatrizContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#tipomatriz}.
	 * @param ctx the parse tree
	 */
	void exitTipomatriz(SwiftGrammarParser.TipomatrizContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#defmatriz}.
	 * @param ctx the parse tree
	 */
	void enterDefmatriz(SwiftGrammarParser.DefmatrizContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#defmatriz}.
	 * @param ctx the parse tree
	 */
	void exitDefmatriz(SwiftGrammarParser.DefmatrizContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listavaloresmat}.
	 * @param ctx the parse tree
	 */
	void enterListavaloresmat(SwiftGrammarParser.ListavaloresmatContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listavaloresmat}.
	 * @param ctx the parse tree
	 */
	void exitListavaloresmat(SwiftGrammarParser.ListavaloresmatContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listavaloresmat2}.
	 * @param ctx the parse tree
	 */
	void enterListavaloresmat2(SwiftGrammarParser.Listavaloresmat2Context ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listavaloresmat2}.
	 * @param ctx the parse tree
	 */
	void exitListavaloresmat2(SwiftGrammarParser.Listavaloresmat2Context ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listaexpresions}.
	 * @param ctx the parse tree
	 */
	void enterListaexpresions(SwiftGrammarParser.ListaexpresionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listaexpresions}.
	 * @param ctx the parse tree
	 */
	void exitListaexpresions(SwiftGrammarParser.ListaexpresionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listaexpresion}.
	 * @param ctx the parse tree
	 */
	void enterListaexpresion(SwiftGrammarParser.ListaexpresionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listaexpresion}.
	 * @param ctx the parse tree
	 */
	void exitListaexpresion(SwiftGrammarParser.ListaexpresionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#simplematriz}.
	 * @param ctx the parse tree
	 */
	void enterSimplematriz(SwiftGrammarParser.SimplematrizContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#simplematriz}.
	 * @param ctx the parse tree
	 */
	void exitSimplematriz(SwiftGrammarParser.SimplematrizContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listamatrizaddsubs}.
	 * @param ctx the parse tree
	 */
	void enterListamatrizaddsubs(SwiftGrammarParser.ListamatrizaddsubsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listamatrizaddsubs}.
	 * @param ctx the parse tree
	 */
	void exitListamatrizaddsubs(SwiftGrammarParser.ListamatrizaddsubsContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listamatrizaddsub}.
	 * @param ctx the parse tree
	 */
	void enterListamatrizaddsub(SwiftGrammarParser.ListamatrizaddsubContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listamatrizaddsub}.
	 * @param ctx the parse tree
	 */
	void exitListamatrizaddsub(SwiftGrammarParser.ListamatrizaddsubContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#structcontrol}.
	 * @param ctx the parse tree
	 */
	void enterStructcontrol(SwiftGrammarParser.StructcontrolContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#structcontrol}.
	 * @param ctx the parse tree
	 */
	void exitStructcontrol(SwiftGrammarParser.StructcontrolContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listaatributos}.
	 * @param ctx the parse tree
	 */
	void enterListaatributos(SwiftGrammarParser.ListaatributosContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listaatributos}.
	 * @param ctx the parse tree
	 */
	void exitListaatributos(SwiftGrammarParser.ListaatributosContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listaatributo}.
	 * @param ctx the parse tree
	 */
	void enterListaatributo(SwiftGrammarParser.ListaatributoContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listaatributo}.
	 * @param ctx the parse tree
	 */
	void exitListaatributo(SwiftGrammarParser.ListaatributoContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#funciondeclaracioncontrol}.
	 * @param ctx the parse tree
	 */
	void enterFunciondeclaracioncontrol(SwiftGrammarParser.FunciondeclaracioncontrolContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#funciondeclaracioncontrol}.
	 * @param ctx the parse tree
	 */
	void exitFunciondeclaracioncontrol(SwiftGrammarParser.FunciondeclaracioncontrolContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listaparametro}.
	 * @param ctx the parse tree
	 */
	void enterListaparametro(SwiftGrammarParser.ListaparametroContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listaparametro}.
	 * @param ctx the parse tree
	 */
	void exitListaparametro(SwiftGrammarParser.ListaparametroContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#funcionllamadacontrol}.
	 * @param ctx the parse tree
	 */
	void enterFuncionllamadacontrol(SwiftGrammarParser.FuncionllamadacontrolContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#funcionllamadacontrol}.
	 * @param ctx the parse tree
	 */
	void exitFuncionllamadacontrol(SwiftGrammarParser.FuncionllamadacontrolContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#funcionllamadacontrolConRetorno}.
	 * @param ctx the parse tree
	 */
	void enterFuncionllamadacontrolConRetorno(SwiftGrammarParser.FuncionllamadacontrolConRetornoContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#funcionllamadacontrolConRetorno}.
	 * @param ctx the parse tree
	 */
	void exitFuncionllamadacontrolConRetorno(SwiftGrammarParser.FuncionllamadacontrolConRetornoContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listaparametrosllamada}.
	 * @param ctx the parse tree
	 */
	void enterListaparametrosllamada(SwiftGrammarParser.ListaparametrosllamadaContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listaparametrosllamada}.
	 * @param ctx the parse tree
	 */
	void exitListaparametrosllamada(SwiftGrammarParser.ListaparametrosllamadaContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#printstmt}.
	 * @param ctx the parse tree
	 */
	void enterPrintstmt(SwiftGrammarParser.PrintstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#printstmt}.
	 * @param ctx the parse tree
	 */
	void exitPrintstmt(SwiftGrammarParser.PrintstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#intembebida}.
	 * @param ctx the parse tree
	 */
	void enterIntembebida(SwiftGrammarParser.IntembebidaContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#intembebida}.
	 * @param ctx the parse tree
	 */
	void exitIntembebida(SwiftGrammarParser.IntembebidaContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#floatembebida}.
	 * @param ctx the parse tree
	 */
	void enterFloatembebida(SwiftGrammarParser.FloatembebidaContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#floatembebida}.
	 * @param ctx the parse tree
	 */
	void exitFloatembebida(SwiftGrammarParser.FloatembebidaContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#stringembebida}.
	 * @param ctx the parse tree
	 */
	void enterStringembebida(SwiftGrammarParser.StringembebidaContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#stringembebida}.
	 * @param ctx the parse tree
	 */
	void exitStringembebida(SwiftGrammarParser.StringembebidaContext ctx);
}