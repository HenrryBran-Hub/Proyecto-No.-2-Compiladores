// Generated from c:/Users/Henrr/OneDrive/Escritorio/Compiladores/Proyecto No. 2/Proyecto No. 2 Compiladores/Backend/SwiftGrammar.g4 by ANTLR 4.13.1

    import "Backend/interfaces"
    import "Backend/environment"
    import "Backend/expressions"
    import "Backend/instructions/datoscompuestos"
    import "Backend/instructions/datosprimitivos"
    import "Backend/instructions/funciones"
    import "Backend/instructions/sentencias"
    import "strings"
   

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class SwiftGrammarParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		INT=1, FLOAT=2, STRING=3, BOOL=4, CHARACT=5, TRU=6, FAL=7, VAR=8, LET=9, 
		NULO=10, IF=11, ELSE=12, SWITCH=13, CASE=14, DEFAULT=15, BREAK=16, CONTINUE=17, 
		FOR=18, IN=19, RANGO=20, WHILE=21, GUARD=22, RETURN=23, FUNCION=24, PRINT=25, 
		INOUT=26, APPEND=27, REMOVE=28, REMOVELAST=29, COUNT=30, ISEMPTY=31, AT=32, 
		REPEATING=33, STRUCT=34, MUTATING=35, SELF=36, NUMBER=37, CADENA=38, ID_VALIDO=39, 
		CHARACTER=40, WS=41, IG=42, DOS_PUNTOS=43, PUNTOCOMA=44, CIERRE_INTE=45, 
		PARIZQ=46, PARDER=47, DIF=48, IG_IG=49, NOT=50, OR=51, AND=52, MAY_IG=53, 
		MEN_IG=54, MAYOR=55, MENOR=56, MODULO=57, MUL=58, DIV=59, ADD=60, SUB=61, 
		SUMA=62, RESTA=63, LLAVEIZQ=64, LLAVEDER=65, RETORNO=66, COMA=67, PUNTO=68, 
		GUIONBAJO=69, CORCHIZQ=70, CORCHDER=71, DIRME=72, WHITESPACE=73, COMMENT=74, 
		LINE_COMMENT=75;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_instruction = 2, RULE_blockinterno = 3, 
		RULE_instructionint = 4, RULE_declavarible = 5, RULE_declaconstante = 6, 
		RULE_asignacionvariable = 7, RULE_tipodato = 8, RULE_expr = 9, RULE_sentenciaifelse = 10, 
		RULE_switchcontrol = 11, RULE_blockcase = 12, RULE_bloquecase = 13, RULE_whilecontrol = 14, 
		RULE_forcontrol = 15, RULE_guardcontrol = 16, RULE_continuee = 17, RULE_breakk = 18, 
		RULE_retornos = 19, RULE_vectorcontrol = 20, RULE_blockparams = 21, RULE_bloqueparams = 22, 
		RULE_vectoragregar = 23, RULE_vectorremover = 24, RULE_vectorvacio = 25, 
		RULE_vectorcount = 26, RULE_vectoraccess = 27, RULE_matrizcontrol = 28, 
		RULE_tipomatriz = 29, RULE_defmatriz = 30, RULE_listavaloresmat = 31, 
		RULE_listavaloresmat2 = 32, RULE_listaexpresions = 33, RULE_listaexpresion = 34, 
		RULE_simplematriz = 35, RULE_listamatrizaddsubs = 36, RULE_listamatrizaddsub = 37, 
		RULE_funciondeclaracioncontrol = 38, RULE_listaparametro = 39, RULE_funcionllamadacontrol = 40, 
		RULE_funcionllamadacontrolConRetorno = 41, RULE_listaparametrosllamada = 42, 
		RULE_printstmt = 43, RULE_intembebida = 44, RULE_floatembebida = 45, RULE_stringembebida = 46;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "blockinterno", "instructionint", "declavarible", 
			"declaconstante", "asignacionvariable", "tipodato", "expr", "sentenciaifelse", 
			"switchcontrol", "blockcase", "bloquecase", "whilecontrol", "forcontrol", 
			"guardcontrol", "continuee", "breakk", "retornos", "vectorcontrol", "blockparams", 
			"bloqueparams", "vectoragregar", "vectorremover", "vectorvacio", "vectorcount", 
			"vectoraccess", "matrizcontrol", "tipomatriz", "defmatriz", "listavaloresmat", 
			"listavaloresmat2", "listaexpresions", "listaexpresion", "simplematriz", 
			"listamatrizaddsubs", "listamatrizaddsub", "funciondeclaracioncontrol", 
			"listaparametro", "funcionllamadacontrol", "funcionllamadacontrolConRetorno", 
			"listaparametrosllamada", "printstmt", "intembebida", "floatembebida", 
			"stringembebida"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'String'", "'Bool'", "'Character'", "'true'", 
			"'false'", "'var'", "'let'", "'nil'", "'if'", "'else'", "'switch'", "'case'", 
			"'default'", "'break'", "'continue'", "'for'", "'in'", "'...'", "'while'", 
			"'guard'", "'return'", "'func'", "'print'", "'inout'", "'append'", "'remove'", 
			"'removeLast'", "'count'", "'isEmpty'", "'at'", "'repeating'", "'struct'", 
			"'mutating'", "'self'", null, null, null, null, null, "'='", "':'", "';'", 
			"'?'", "'('", "')'", "'!='", "'=='", "'!'", "'||'", "'&&'", "'>='", "'<='", 
			"'>'", "'<'", "'%'", "'*'", "'/'", "'+'", "'-'", "'+='", "'-='", "'{'", 
			"'}'", "'->'", "','", "'.'", "'_'", "'['", "']'", "'&'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "STRING", "BOOL", "CHARACT", "TRU", "FAL", "VAR", 
			"LET", "NULO", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", "BREAK", "CONTINUE", 
			"FOR", "IN", "RANGO", "WHILE", "GUARD", "RETURN", "FUNCION", "PRINT", 
			"INOUT", "APPEND", "REMOVE", "REMOVELAST", "COUNT", "ISEMPTY", "AT", 
			"REPEATING", "STRUCT", "MUTATING", "SELF", "NUMBER", "CADENA", "ID_VALIDO", 
			"CHARACTER", "WS", "IG", "DOS_PUNTOS", "PUNTOCOMA", "CIERRE_INTE", "PARIZQ", 
			"PARDER", "DIF", "IG_IG", "NOT", "OR", "AND", "MAY_IG", "MEN_IG", "MAYOR", 
			"MENOR", "MODULO", "MUL", "DIV", "ADD", "SUB", "SUMA", "RESTA", "LLAVEIZQ", 
			"LLAVEDER", "RETORNO", "COMA", "PUNTO", "GUIONBAJO", "CORCHIZQ", "CORCHDER", 
			"DIRME", "WHITESPACE", "COMMENT", "LINE_COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "SwiftGrammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public SwiftGrammarParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SContext extends ParserRuleContext {
		public []interface{} code;
		public BlockContext block;
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode EOF() { return getToken(SwiftGrammarParser.EOF, 0); }
		public SContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_s; }
	}

	public final SContext s() throws RecognitionException {
		SContext _localctx = new SContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_s);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(94);
			((SContext)_localctx).block = block();
			setState(95);
			match(EOF);
			   
			        _localctx.code = ((SContext)_localctx).block.blk
			    
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockContext extends ParserRuleContext {
		public []interface{} blk;
		public InstructionContext instruction;
		public List<InstructionContext> ins = new ArrayList<InstructionContext>();
		public List<InstructionContext> instruction() {
			return getRuleContexts(InstructionContext.class);
		}
		public InstructionContext instruction(int i) {
			return getRuleContext(InstructionContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_block);

		    _localctx.blk = []interface{}{}
		    var listInt []IInstructionContext
		  
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(99); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(98);
				((BlockContext)_localctx).instruction = instruction();
				((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
				}
				}
				setState(101); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 549812710144L) != 0) );

			        listInt = localctx.(*BlockContext).GetIns()
			        for _, e := range listInt {
			            _localctx.blk = append(_localctx.blk, e.GetInst())
			        }
			    
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InstructionContext extends ParserRuleContext {
		public interfaces.Instruction inst;
		public DeclavaribleContext declavarible;
		public DeclaconstanteContext declaconstante;
		public AsignacionvariableContext asignacionvariable;
		public SentenciaifelseContext sentenciaifelse;
		public SwitchcontrolContext switchcontrol;
		public WhilecontrolContext whilecontrol;
		public ForcontrolContext forcontrol;
		public GuardcontrolContext guardcontrol;
		public VectorcontrolContext vectorcontrol;
		public VectoragregarContext vectoragregar;
		public VectorremoverContext vectorremover;
		public PrintstmtContext printstmt;
		public MatrizcontrolContext matrizcontrol;
		public FunciondeclaracioncontrolContext funciondeclaracioncontrol;
		public FuncionllamadacontrolContext funcionllamadacontrol;
		public DeclavaribleContext declavarible() {
			return getRuleContext(DeclavaribleContext.class,0);
		}
		public TerminalNode PUNTOCOMA() { return getToken(SwiftGrammarParser.PUNTOCOMA, 0); }
		public DeclaconstanteContext declaconstante() {
			return getRuleContext(DeclaconstanteContext.class,0);
		}
		public AsignacionvariableContext asignacionvariable() {
			return getRuleContext(AsignacionvariableContext.class,0);
		}
		public SentenciaifelseContext sentenciaifelse() {
			return getRuleContext(SentenciaifelseContext.class,0);
		}
		public SwitchcontrolContext switchcontrol() {
			return getRuleContext(SwitchcontrolContext.class,0);
		}
		public WhilecontrolContext whilecontrol() {
			return getRuleContext(WhilecontrolContext.class,0);
		}
		public ForcontrolContext forcontrol() {
			return getRuleContext(ForcontrolContext.class,0);
		}
		public GuardcontrolContext guardcontrol() {
			return getRuleContext(GuardcontrolContext.class,0);
		}
		public VectorcontrolContext vectorcontrol() {
			return getRuleContext(VectorcontrolContext.class,0);
		}
		public VectoragregarContext vectoragregar() {
			return getRuleContext(VectoragregarContext.class,0);
		}
		public VectorremoverContext vectorremover() {
			return getRuleContext(VectorremoverContext.class,0);
		}
		public PrintstmtContext printstmt() {
			return getRuleContext(PrintstmtContext.class,0);
		}
		public MatrizcontrolContext matrizcontrol() {
			return getRuleContext(MatrizcontrolContext.class,0);
		}
		public FunciondeclaracioncontrolContext funciondeclaracioncontrol() {
			return getRuleContext(FunciondeclaracioncontrolContext.class,0);
		}
		public FuncionllamadacontrolContext funcionllamadacontrol() {
			return getRuleContext(FuncionllamadacontrolContext.class,0);
		}
		public InstructionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruction; }
	}

	public final InstructionContext instruction() throws RecognitionException {
		InstructionContext _localctx = new InstructionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruction);
		int _la;
		try {
			setState(168);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(105);
				((InstructionContext)_localctx).declavarible = declavarible();
				setState(107);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(106);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).declavarible.decvbl
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(111);
				((InstructionContext)_localctx).declaconstante = declaconstante();
				setState(113);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(112);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).declaconstante.deccon
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(117);
				((InstructionContext)_localctx).asignacionvariable = asignacionvariable();
				setState(119);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(118);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).asignacionvariable.asgvbl
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(123);
				((InstructionContext)_localctx).sentenciaifelse = sentenciaifelse();
				 _localctx.inst = ((InstructionContext)_localctx).sentenciaifelse.myIfElse
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(126);
				((InstructionContext)_localctx).switchcontrol = switchcontrol();
				 _localctx.inst = ((InstructionContext)_localctx).switchcontrol.mySwitch
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(129);
				((InstructionContext)_localctx).whilecontrol = whilecontrol();
				 _localctx.inst = ((InstructionContext)_localctx).whilecontrol.whict
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(132);
				((InstructionContext)_localctx).forcontrol = forcontrol();
				 _localctx.inst = ((InstructionContext)_localctx).forcontrol.forct
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(135);
				((InstructionContext)_localctx).guardcontrol = guardcontrol();
				 _localctx.inst = ((InstructionContext)_localctx).guardcontrol.guct
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(138);
				((InstructionContext)_localctx).vectorcontrol = vectorcontrol();
				setState(140);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(139);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).vectorcontrol.vect 
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(144);
				((InstructionContext)_localctx).vectoragregar = vectoragregar();
				 _localctx.inst = ((InstructionContext)_localctx).vectoragregar.veadct 
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(147);
				((InstructionContext)_localctx).vectorremover = vectorremover();
				 _localctx.inst = ((InstructionContext)_localctx).vectorremover.vermct 
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(150);
				((InstructionContext)_localctx).printstmt = printstmt();
				setState(152);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(151);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(156);
				((InstructionContext)_localctx).matrizcontrol = matrizcontrol();
				setState(158);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(157);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).matrizcontrol.matct
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(162);
				((InstructionContext)_localctx).funciondeclaracioncontrol = funciondeclaracioncontrol();
				 _localctx.inst = ((InstructionContext)_localctx).funciondeclaracioncontrol.fdc
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(165);
				((InstructionContext)_localctx).funcionllamadacontrol = funcionllamadacontrol();
				 _localctx.inst = ((InstructionContext)_localctx).funcionllamadacontrol.flctl
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockinternoContext extends ParserRuleContext {
		public []interface{} blkint;
		public InstructionintContext instructionint;
		public List<InstructionintContext> insint = new ArrayList<InstructionintContext>();
		public List<InstructionintContext> instructionint() {
			return getRuleContexts(InstructionintContext.class);
		}
		public InstructionintContext instructionint(int i) {
			return getRuleContext(InstructionintContext.class,i);
		}
		public BlockinternoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_blockinterno; }
	}

	public final BlockinternoContext blockinterno() throws RecognitionException {
		BlockinternoContext _localctx = new BlockinternoContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_blockinterno);

		    _localctx.blkint = []interface{}{}
		    var listInt []IInstructionintContext
		  
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(171); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(170);
				((BlockinternoContext)_localctx).instructionint = instructionint();
				((BlockinternoContext)_localctx).insint.add(((BlockinternoContext)_localctx).instructionint);
				}
				}
				setState(173); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 549804518144L) != 0) );

			        listInt = localctx.(*BlockinternoContext).GetInsint()
			        for _, e := range listInt {
			            _localctx.blkint = append(_localctx.blkint, e.GetInsint())
			        }
			    
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InstructionintContext extends ParserRuleContext {
		public interfaces.Instruction insint;
		public DeclavaribleContext declavarible;
		public DeclaconstanteContext declaconstante;
		public AsignacionvariableContext asignacionvariable;
		public SentenciaifelseContext sentenciaifelse;
		public SwitchcontrolContext switchcontrol;
		public WhilecontrolContext whilecontrol;
		public ForcontrolContext forcontrol;
		public GuardcontrolContext guardcontrol;
		public ContinueeContext continuee;
		public BreakkContext breakk;
		public RetornosContext retornos;
		public VectorcontrolContext vectorcontrol;
		public VectoragregarContext vectoragregar;
		public VectorremoverContext vectorremover;
		public PrintstmtContext printstmt;
		public MatrizcontrolContext matrizcontrol;
		public FuncionllamadacontrolContext funcionllamadacontrol;
		public DeclavaribleContext declavarible() {
			return getRuleContext(DeclavaribleContext.class,0);
		}
		public TerminalNode PUNTOCOMA() { return getToken(SwiftGrammarParser.PUNTOCOMA, 0); }
		public DeclaconstanteContext declaconstante() {
			return getRuleContext(DeclaconstanteContext.class,0);
		}
		public AsignacionvariableContext asignacionvariable() {
			return getRuleContext(AsignacionvariableContext.class,0);
		}
		public SentenciaifelseContext sentenciaifelse() {
			return getRuleContext(SentenciaifelseContext.class,0);
		}
		public SwitchcontrolContext switchcontrol() {
			return getRuleContext(SwitchcontrolContext.class,0);
		}
		public WhilecontrolContext whilecontrol() {
			return getRuleContext(WhilecontrolContext.class,0);
		}
		public ForcontrolContext forcontrol() {
			return getRuleContext(ForcontrolContext.class,0);
		}
		public GuardcontrolContext guardcontrol() {
			return getRuleContext(GuardcontrolContext.class,0);
		}
		public ContinueeContext continuee() {
			return getRuleContext(ContinueeContext.class,0);
		}
		public BreakkContext breakk() {
			return getRuleContext(BreakkContext.class,0);
		}
		public RetornosContext retornos() {
			return getRuleContext(RetornosContext.class,0);
		}
		public VectorcontrolContext vectorcontrol() {
			return getRuleContext(VectorcontrolContext.class,0);
		}
		public VectoragregarContext vectoragregar() {
			return getRuleContext(VectoragregarContext.class,0);
		}
		public VectorremoverContext vectorremover() {
			return getRuleContext(VectorremoverContext.class,0);
		}
		public PrintstmtContext printstmt() {
			return getRuleContext(PrintstmtContext.class,0);
		}
		public MatrizcontrolContext matrizcontrol() {
			return getRuleContext(MatrizcontrolContext.class,0);
		}
		public FuncionllamadacontrolContext funcionllamadacontrol() {
			return getRuleContext(FuncionllamadacontrolContext.class,0);
		}
		public InstructionintContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instructionint; }
	}

	public final InstructionintContext instructionint() throws RecognitionException {
		InstructionintContext _localctx = new InstructionintContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_instructionint);
		int _la;
		try {
			setState(261);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(177);
				((InstructionintContext)_localctx).declavarible = declavarible();
				setState(179);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(178);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).declavarible.decvbl
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(183);
				((InstructionintContext)_localctx).declaconstante = declaconstante();
				setState(185);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(184);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).declaconstante.deccon
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(189);
				((InstructionintContext)_localctx).asignacionvariable = asignacionvariable();
				setState(191);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(190);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).asignacionvariable.asgvbl
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(195);
				((InstructionintContext)_localctx).sentenciaifelse = sentenciaifelse();
				 _localctx.insint = ((InstructionintContext)_localctx).sentenciaifelse.myIfElse
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(198);
				((InstructionintContext)_localctx).switchcontrol = switchcontrol();
				 _localctx.insint = ((InstructionintContext)_localctx).switchcontrol.mySwitch
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(201);
				((InstructionintContext)_localctx).whilecontrol = whilecontrol();
				 _localctx.insint = ((InstructionintContext)_localctx).whilecontrol.whict
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(204);
				((InstructionintContext)_localctx).forcontrol = forcontrol();
				 _localctx.insint = ((InstructionintContext)_localctx).forcontrol.forct
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(207);
				((InstructionintContext)_localctx).guardcontrol = guardcontrol();
				 _localctx.insint = ((InstructionintContext)_localctx).guardcontrol.guct
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(210);
				((InstructionintContext)_localctx).continuee = continuee();
				setState(212);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(211);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).continuee.coct
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(216);
				((InstructionintContext)_localctx).breakk = breakk();
				setState(218);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(217);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).breakk.brkct
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(222);
				((InstructionintContext)_localctx).retornos = retornos();
				setState(224);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(223);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).retornos.rect 
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(228);
				((InstructionintContext)_localctx).vectorcontrol = vectorcontrol();
				setState(230);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(229);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).vectorcontrol.vect 
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(234);
				((InstructionintContext)_localctx).vectoragregar = vectoragregar();
				setState(236);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(235);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).vectoragregar.veadct 
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(240);
				((InstructionintContext)_localctx).vectorremover = vectorremover();
				setState(242);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(241);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).vectorremover.vermct 
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(246);
				((InstructionintContext)_localctx).printstmt = printstmt();
				setState(248);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(247);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).printstmt.prnt
				}
				break;
			case 16:
				enterOuterAlt(_localctx, 16);
				{
				setState(252);
				((InstructionintContext)_localctx).matrizcontrol = matrizcontrol();
				setState(254);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(253);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).matrizcontrol.matct
				}
				break;
			case 17:
				enterOuterAlt(_localctx, 17);
				{
				setState(258);
				((InstructionintContext)_localctx).funcionllamadacontrol = funcionllamadacontrol();
				 _localctx.insint = ((InstructionintContext)_localctx).funcionllamadacontrol.flctl
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DeclavaribleContext extends ParserRuleContext {
		public interfaces.Instruction decvbl;
		public Token VAR;
		public Token ID_VALIDO;
		public TipodatoContext tipodato;
		public ExprContext expr;
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TerminalNode ID_VALIDO() { return getToken(SwiftGrammarParser.ID_VALIDO, 0); }
		public TerminalNode DOS_PUNTOS() { return getToken(SwiftGrammarParser.DOS_PUNTOS, 0); }
		public TipodatoContext tipodato() {
			return getRuleContext(TipodatoContext.class,0);
		}
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode CIERRE_INTE() { return getToken(SwiftGrammarParser.CIERRE_INTE, 0); }
		public DeclavaribleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declavarible; }
	}

	public final DeclavaribleContext declavarible() throws RecognitionException {
		DeclavaribleContext _localctx = new DeclavaribleContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_declavarible);
		try {
			setState(284);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(263);
				((DeclavaribleContext)_localctx).VAR = match(VAR);
				setState(264);
				((DeclavaribleContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(265);
				match(DOS_PUNTOS);
				setState(266);
				((DeclavaribleContext)_localctx).tipodato = tipodato();
				setState(267);
				match(IG);
				setState(268);
				((DeclavaribleContext)_localctx).expr = expr(0);
				 _localctx.decvbl = datosprimitivos.NewVariableDeclaration((((DeclavaribleContext)_localctx).VAR!=null?((DeclavaribleContext)_localctx).VAR.getLine():0), (((DeclavaribleContext)_localctx).VAR!=null?((DeclavaribleContext)_localctx).VAR.getCharPositionInLine():0), (((DeclavaribleContext)_localctx).ID_VALIDO!=null?((DeclavaribleContext)_localctx).ID_VALIDO.getText():null), ((DeclavaribleContext)_localctx).tipodato.tipo, ((DeclavaribleContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(271);
				((DeclavaribleContext)_localctx).VAR = match(VAR);
				setState(272);
				((DeclavaribleContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(273);
				match(IG);
				setState(274);
				((DeclavaribleContext)_localctx).expr = expr(0);
				_localctx.decvbl = datosprimitivos.NewVariableDeclaracionSinTipo((((DeclavaribleContext)_localctx).VAR!=null?((DeclavaribleContext)_localctx).VAR.getLine():0), (((DeclavaribleContext)_localctx).VAR!=null?((DeclavaribleContext)_localctx).VAR.getCharPositionInLine():0), (((DeclavaribleContext)_localctx).ID_VALIDO!=null?((DeclavaribleContext)_localctx).ID_VALIDO.getText():null), ((DeclavaribleContext)_localctx).expr.e)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(277);
				((DeclavaribleContext)_localctx).VAR = match(VAR);
				setState(278);
				((DeclavaribleContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(279);
				match(DOS_PUNTOS);
				setState(280);
				((DeclavaribleContext)_localctx).tipodato = tipodato();
				setState(281);
				match(CIERRE_INTE);
				_localctx.decvbl = datosprimitivos.NewVariableDeclaracionSinExp((((DeclavaribleContext)_localctx).VAR!=null?((DeclavaribleContext)_localctx).VAR.getLine():0), (((DeclavaribleContext)_localctx).VAR!=null?((DeclavaribleContext)_localctx).VAR.getCharPositionInLine():0), (((DeclavaribleContext)_localctx).ID_VALIDO!=null?((DeclavaribleContext)_localctx).ID_VALIDO.getText():null), ((DeclavaribleContext)_localctx).tipodato.tipo)
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DeclaconstanteContext extends ParserRuleContext {
		public interfaces.Instruction deccon;
		public Token LET;
		public Token ID_VALIDO;
		public TipodatoContext tipodato;
		public ExprContext expr;
		public TerminalNode LET() { return getToken(SwiftGrammarParser.LET, 0); }
		public TerminalNode ID_VALIDO() { return getToken(SwiftGrammarParser.ID_VALIDO, 0); }
		public TerminalNode DOS_PUNTOS() { return getToken(SwiftGrammarParser.DOS_PUNTOS, 0); }
		public TipodatoContext tipodato() {
			return getRuleContext(TipodatoContext.class,0);
		}
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public DeclaconstanteContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaconstante; }
	}

	public final DeclaconstanteContext declaconstante() throws RecognitionException {
		DeclaconstanteContext _localctx = new DeclaconstanteContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_declaconstante);
		try {
			setState(300);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(286);
				((DeclaconstanteContext)_localctx).LET = match(LET);
				setState(287);
				((DeclaconstanteContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(288);
				match(DOS_PUNTOS);
				setState(289);
				((DeclaconstanteContext)_localctx).tipodato = tipodato();
				setState(290);
				match(IG);
				setState(291);
				((DeclaconstanteContext)_localctx).expr = expr(0);
				_localctx.deccon = datosprimitivos.NewConstanteDeclaration((((DeclaconstanteContext)_localctx).LET!=null?((DeclaconstanteContext)_localctx).LET.getLine():0), (((DeclaconstanteContext)_localctx).LET!=null?((DeclaconstanteContext)_localctx).LET.getCharPositionInLine():0), (((DeclaconstanteContext)_localctx).ID_VALIDO!=null?((DeclaconstanteContext)_localctx).ID_VALIDO.getText():null), ((DeclaconstanteContext)_localctx).tipodato.tipo, ((DeclaconstanteContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(294);
				((DeclaconstanteContext)_localctx).LET = match(LET);
				setState(295);
				((DeclaconstanteContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(296);
				match(IG);
				setState(297);
				((DeclaconstanteContext)_localctx).expr = expr(0);
				_localctx.deccon = datosprimitivos.NewConstanteDeclaracionSinTipo((((DeclaconstanteContext)_localctx).LET!=null?((DeclaconstanteContext)_localctx).LET.getLine():0), (((DeclaconstanteContext)_localctx).LET!=null?((DeclaconstanteContext)_localctx).LET.getCharPositionInLine():0), (((DeclaconstanteContext)_localctx).ID_VALIDO!=null?((DeclaconstanteContext)_localctx).ID_VALIDO.getText():null), ((DeclaconstanteContext)_localctx).expr.e)
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AsignacionvariableContext extends ParserRuleContext {
		public interfaces.Instruction asgvbl;
		public Token ID_VALIDO;
		public ExprContext expr;
		public TerminalNode ID_VALIDO() { return getToken(SwiftGrammarParser.ID_VALIDO, 0); }
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode SUMA() { return getToken(SwiftGrammarParser.SUMA, 0); }
		public TerminalNode RESTA() { return getToken(SwiftGrammarParser.RESTA, 0); }
		public AsignacionvariableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacionvariable; }
	}

	public final AsignacionvariableContext asignacionvariable() throws RecognitionException {
		AsignacionvariableContext _localctx = new AsignacionvariableContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_asignacionvariable);
		try {
			setState(317);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(302);
				((AsignacionvariableContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(303);
				match(IG);
				setState(304);
				((AsignacionvariableContext)_localctx).expr = expr(0);
				 _localctx.asgvbl = sentencias.NewAsignacionVariable((((AsignacionvariableContext)_localctx).ID_VALIDO!=null?((AsignacionvariableContext)_localctx).ID_VALIDO.getLine():0), (((AsignacionvariableContext)_localctx).ID_VALIDO!=null?((AsignacionvariableContext)_localctx).ID_VALIDO.getCharPositionInLine():0), (((AsignacionvariableContext)_localctx).ID_VALIDO!=null?((AsignacionvariableContext)_localctx).ID_VALIDO.getText():null), ((AsignacionvariableContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(307);
				((AsignacionvariableContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(308);
				match(SUMA);
				setState(309);
				((AsignacionvariableContext)_localctx).expr = expr(0);
				 _localctx.asgvbl = sentencias.NewAsignacionSuma((((AsignacionvariableContext)_localctx).ID_VALIDO!=null?((AsignacionvariableContext)_localctx).ID_VALIDO.getLine():0), (((AsignacionvariableContext)_localctx).ID_VALIDO!=null?((AsignacionvariableContext)_localctx).ID_VALIDO.getCharPositionInLine():0), (((AsignacionvariableContext)_localctx).ID_VALIDO!=null?((AsignacionvariableContext)_localctx).ID_VALIDO.getText():null), ((AsignacionvariableContext)_localctx).expr.e)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(312);
				((AsignacionvariableContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(313);
				match(RESTA);
				setState(314);
				((AsignacionvariableContext)_localctx).expr = expr(0);
				 _localctx.asgvbl = sentencias.NewAsignacionResta((((AsignacionvariableContext)_localctx).ID_VALIDO!=null?((AsignacionvariableContext)_localctx).ID_VALIDO.getLine():0), (((AsignacionvariableContext)_localctx).ID_VALIDO!=null?((AsignacionvariableContext)_localctx).ID_VALIDO.getCharPositionInLine():0), (((AsignacionvariableContext)_localctx).ID_VALIDO!=null?((AsignacionvariableContext)_localctx).ID_VALIDO.getText():null), ((AsignacionvariableContext)_localctx).expr.e)
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TipodatoContext extends ParserRuleContext {
		public environment.TipoExpresion tipo;
		public TerminalNode INT() { return getToken(SwiftGrammarParser.INT, 0); }
		public TerminalNode FLOAT() { return getToken(SwiftGrammarParser.FLOAT, 0); }
		public TerminalNode STRING() { return getToken(SwiftGrammarParser.STRING, 0); }
		public TerminalNode BOOL() { return getToken(SwiftGrammarParser.BOOL, 0); }
		public TerminalNode CHARACT() { return getToken(SwiftGrammarParser.CHARACT, 0); }
		public TipodatoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipodato; }
	}

	public final TipodatoContext tipodato() throws RecognitionException {
		TipodatoContext _localctx = new TipodatoContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_tipodato);
		try {
			setState(329);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(319);
				match(INT);
				 _localctx.tipo = environment.INTEGER 
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(321);
				match(FLOAT);
				 _localctx.tipo = environment.FLOAT 
				}
				break;
			case STRING:
				enterOuterAlt(_localctx, 3);
				{
				setState(323);
				match(STRING);
				 _localctx.tipo = environment.STRING 
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(325);
				match(BOOL);
				 _localctx.tipo = environment.BOOLEAN 
				}
				break;
			case CHARACT:
				enterOuterAlt(_localctx, 5);
				{
				setState(327);
				match(CHARACT);
				 _localctx.tipo = environment.CHARACTER 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExprContext extends ParserRuleContext {
		public interfaces.Expression e;
		public ExprContext left;
		public Token op;
		public ExprContext right;
		public ExprContext expr;
		public Token NUMBER;
		public Token CADENA;
		public Token TRU;
		public Token FAL;
		public Token CHARACTER;
		public Token ID_VALIDO;
		public Token NULO;
		public VectorvacioContext vectorvacio;
		public VectorcountContext vectorcount;
		public VectoraccessContext vectoraccess;
		public IntembebidaContext intembebida;
		public FloatembebidaContext floatembebida;
		public StringembebidaContext stringembebida;
		public FuncionllamadacontrolConRetornoContext funcionllamadacontrolConRetorno;
		public TerminalNode NOT() { return getToken(SwiftGrammarParser.NOT, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public TerminalNode SUB() { return getToken(SwiftGrammarParser.SUB, 0); }
		public TerminalNode NUMBER() { return getToken(SwiftGrammarParser.NUMBER, 0); }
		public TerminalNode CADENA() { return getToken(SwiftGrammarParser.CADENA, 0); }
		public TerminalNode TRU() { return getToken(SwiftGrammarParser.TRU, 0); }
		public TerminalNode FAL() { return getToken(SwiftGrammarParser.FAL, 0); }
		public TerminalNode CHARACTER() { return getToken(SwiftGrammarParser.CHARACTER, 0); }
		public TerminalNode ID_VALIDO() { return getToken(SwiftGrammarParser.ID_VALIDO, 0); }
		public TerminalNode NULO() { return getToken(SwiftGrammarParser.NULO, 0); }
		public VectorvacioContext vectorvacio() {
			return getRuleContext(VectorvacioContext.class,0);
		}
		public VectorcountContext vectorcount() {
			return getRuleContext(VectorcountContext.class,0);
		}
		public VectoraccessContext vectoraccess() {
			return getRuleContext(VectoraccessContext.class,0);
		}
		public IntembebidaContext intembebida() {
			return getRuleContext(IntembebidaContext.class,0);
		}
		public FloatembebidaContext floatembebida() {
			return getRuleContext(FloatembebidaContext.class,0);
		}
		public StringembebidaContext stringembebida() {
			return getRuleContext(StringembebidaContext.class,0);
		}
		public FuncionllamadacontrolConRetornoContext funcionllamadacontrolConRetorno() {
			return getRuleContext(FuncionllamadacontrolConRetornoContext.class,0);
		}
		public TerminalNode MODULO() { return getToken(SwiftGrammarParser.MODULO, 0); }
		public TerminalNode MUL() { return getToken(SwiftGrammarParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(SwiftGrammarParser.DIV, 0); }
		public TerminalNode ADD() { return getToken(SwiftGrammarParser.ADD, 0); }
		public TerminalNode MAY_IG() { return getToken(SwiftGrammarParser.MAY_IG, 0); }
		public TerminalNode MAYOR() { return getToken(SwiftGrammarParser.MAYOR, 0); }
		public TerminalNode MEN_IG() { return getToken(SwiftGrammarParser.MEN_IG, 0); }
		public TerminalNode MENOR() { return getToken(SwiftGrammarParser.MENOR, 0); }
		public TerminalNode IG_IG() { return getToken(SwiftGrammarParser.IG_IG, 0); }
		public TerminalNode DIF() { return getToken(SwiftGrammarParser.DIF, 0); }
		public TerminalNode AND() { return getToken(SwiftGrammarParser.AND, 0); }
		public TerminalNode OR() { return getToken(SwiftGrammarParser.OR, 0); }
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 18;
		enterRecursionRule(_localctx, 18, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(379);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				{
				setState(332);
				((ExprContext)_localctx).op = match(NOT);
				setState(333);
				((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(25);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).right!=null?(((ExprContext)_localctx).right.start):null).GetLine(), (((ExprContext)_localctx).right!=null?(((ExprContext)_localctx).right.start):null).GetColumn(), ((ExprContext)_localctx).right.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
				}
				break;
			case 2:
				{
				setState(336);
				match(PARIZQ);
				setState(337);
				((ExprContext)_localctx).expr = expr(0);
				setState(338);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case 3:
				{
				setState(341);
				match(SUB);
				setState(342);
				((ExprContext)_localctx).NUMBER = match(NUMBER);

				        if (strings.Contains((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getText():null),".")){
				            num,err := strconv.ParseFloat((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getText():null), 64);
				            if err!= nil{
				                fmt.Println(err)
				            }
					        num2 := fmt.Sprintf("%.6f", num)
				            num3,err := strconv.ParseFloat(num2, 64);
				            if err!= nil{
				                fmt.Println(err)
				            }
				            _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getLine():0),(((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getCharPositionInLine():0),-num3,environment.FLOAT)
				        }else{
				            num,err := strconv.Atoi((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getText():null))
				            if err!= nil{
				                fmt.Println(err)
				            }
				            _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getLine():0),(((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getCharPositionInLine():0),-num,environment.INTEGER)
				        }
				    
				}
				break;
			case 4:
				{
				setState(344);
				((ExprContext)_localctx).NUMBER = match(NUMBER);

				        if (strings.Contains((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getText():null),".")){
				            num,err := strconv.ParseFloat((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getText():null), 64);
				            if err!= nil{
				                fmt.Println(err)
				            }
					        num2 := fmt.Sprintf("%.6f", num)
				            num3,err := strconv.ParseFloat(num2, 64);
				            if err!= nil{
				                fmt.Println(err)
				            }
					        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getLine():0),(((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getCharPositionInLine():0),num3,environment.FLOAT)
				        }else{
				            num,err := strconv.Atoi((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getText():null))
				            if err!= nil{
				                fmt.Println(err)
				            }            
				            _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getLine():0),(((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getCharPositionInLine():0),num,environment.INTEGER)
				        }
				    
				}
				break;
			case 5:
				{
				setState(346);
				((ExprContext)_localctx).CADENA = match(CADENA);

				        str := (((ExprContext)_localctx).CADENA!=null?((ExprContext)_localctx).CADENA.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).CADENA!=null?((ExprContext)_localctx).CADENA.getLine():0), (((ExprContext)_localctx).CADENA!=null?((ExprContext)_localctx).CADENA.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case 6:
				{
				setState(348);
				((ExprContext)_localctx).TRU = match(TRU);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getLine():0), (((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case 7:
				{
				setState(350);
				((ExprContext)_localctx).FAL = match(FAL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getLine():0), (((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			case 8:
				{
				setState(352);
				((ExprContext)_localctx).CHARACTER = match(CHARACTER);
				 
				        str := (((ExprContext)_localctx).CHARACTER!=null?((ExprContext)_localctx).CHARACTER.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).CHARACTER!=null?((ExprContext)_localctx).CHARACTER.getLine():0), (((ExprContext)_localctx).CHARACTER!=null?((ExprContext)_localctx).CHARACTER.getCharPositionInLine():0), str[1:len(str)-1], environment.CHARACTER) 
				    
				}
				break;
			case 9:
				{
				setState(354);
				((ExprContext)_localctx).ID_VALIDO = match(ID_VALIDO);

				        id := (((ExprContext)_localctx).ID_VALIDO!=null?((ExprContext)_localctx).ID_VALIDO.getText():null)
				        _localctx.e = sentencias.NewCallid((((ExprContext)_localctx).ID_VALIDO!=null?((ExprContext)_localctx).ID_VALIDO.getLine():0),(((ExprContext)_localctx).ID_VALIDO!=null?((ExprContext)_localctx).ID_VALIDO.getCharPositionInLine():0),id)
				    
				}
				break;
			case 10:
				{
				setState(356);
				((ExprContext)_localctx).NULO = match(NULO);
				_localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NULO!=null?((ExprContext)_localctx).NULO.getLine():0), (((ExprContext)_localctx).NULO!=null?((ExprContext)_localctx).NULO.getCharPositionInLine():0), (((ExprContext)_localctx).NULO!=null?((ExprContext)_localctx).NULO.getText():null),environment.NULL)
				}
				break;
			case 11:
				{
				setState(358);
				((ExprContext)_localctx).vectorvacio = vectorvacio();
				 _localctx.e = ((ExprContext)_localctx).vectorvacio.veemct
				}
				break;
			case 12:
				{
				setState(361);
				((ExprContext)_localctx).vectorcount = vectorcount();
				 _localctx.e = ((ExprContext)_localctx).vectorcount.vecnct
				}
				break;
			case 13:
				{
				setState(364);
				((ExprContext)_localctx).vectoraccess = vectoraccess();
				 _localctx.e = ((ExprContext)_localctx).vectoraccess.vepposct
				}
				break;
			case 14:
				{
				setState(367);
				((ExprContext)_localctx).intembebida = intembebida();
				 _localctx.e = ((ExprContext)_localctx).intembebida.intemb
				}
				break;
			case 15:
				{
				setState(370);
				((ExprContext)_localctx).floatembebida = floatembebida();
				 _localctx.e = ((ExprContext)_localctx).floatembebida.floemb
				}
				break;
			case 16:
				{
				setState(373);
				((ExprContext)_localctx).stringembebida = stringembebida();
				 _localctx.e = ((ExprContext)_localctx).stringembebida.stremb
				}
				break;
			case 17:
				{
				setState(376);
				((ExprContext)_localctx).funcionllamadacontrolConRetorno = funcionllamadacontrolConRetorno();
				 _localctx.e = ((ExprContext)_localctx).funcionllamadacontrolConRetorno.flctlret
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(423);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(421);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(381);
						if (!(precpred(_ctx, 24))) throw new FailedPredicateException(this, "precpred(_ctx, 24)");
						setState(382);
						((ExprContext)_localctx).op = match(MODULO);
						setState(383);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(25);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(386);
						if (!(precpred(_ctx, 23))) throw new FailedPredicateException(this, "precpred(_ctx, 23)");
						setState(387);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MUL || _la==DIV) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(388);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(24);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(391);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(392);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==ADD || _la==SUB) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(393);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(23);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(396);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(397);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MAY_IG || _la==MAYOR) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(398);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(22);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(401);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(402);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MEN_IG || _la==MENOR) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(403);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(21);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(406);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(407);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==DIF || _la==IG_IG) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(408);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(20);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(411);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(412);
						((ExprContext)_localctx).op = match(AND);
						setState(413);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(19);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 8:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(416);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(417);
						((ExprContext)_localctx).op = match(OR);
						setState(418);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(18);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(425);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SentenciaifelseContext extends ParserRuleContext {
		public interfaces.Instruction myIfElse;
		public Token IF;
		public ExprContext expr;
		public BlockinternoContext blockinterno;
		public BlockinternoContext ifop;
		public BlockinternoContext elseop;
		public SentenciaifelseContext sentenciaifelse;
		public TerminalNode IF() { return getToken(SwiftGrammarParser.IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public List<TerminalNode> LLAVEIZQ() { return getTokens(SwiftGrammarParser.LLAVEIZQ); }
		public TerminalNode LLAVEIZQ(int i) {
			return getToken(SwiftGrammarParser.LLAVEIZQ, i);
		}
		public List<BlockinternoContext> blockinterno() {
			return getRuleContexts(BlockinternoContext.class);
		}
		public BlockinternoContext blockinterno(int i) {
			return getRuleContext(BlockinternoContext.class,i);
		}
		public List<TerminalNode> LLAVEDER() { return getTokens(SwiftGrammarParser.LLAVEDER); }
		public TerminalNode LLAVEDER(int i) {
			return getToken(SwiftGrammarParser.LLAVEDER, i);
		}
		public TerminalNode ELSE() { return getToken(SwiftGrammarParser.ELSE, 0); }
		public SentenciaifelseContext sentenciaifelse() {
			return getRuleContext(SentenciaifelseContext.class,0);
		}
		public SentenciaifelseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sentenciaifelse; }
	}

	public final SentenciaifelseContext sentenciaifelse() throws RecognitionException {
		SentenciaifelseContext _localctx = new SentenciaifelseContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_sentenciaifelse);
		try {
			setState(453);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(426);
				((SentenciaifelseContext)_localctx).IF = match(IF);
				setState(427);
				((SentenciaifelseContext)_localctx).expr = expr(0);
				setState(428);
				match(LLAVEIZQ);
				setState(429);
				((SentenciaifelseContext)_localctx).blockinterno = blockinterno();
				setState(430);
				match(LLAVEDER);
				 _localctx.myIfElse = sentencias.NewSentenciaIf((((SentenciaifelseContext)_localctx).IF!=null?((SentenciaifelseContext)_localctx).IF.getLine():0), (((SentenciaifelseContext)_localctx).IF!=null?((SentenciaifelseContext)_localctx).IF.getCharPositionInLine():0), ((SentenciaifelseContext)_localctx).expr.e, ((SentenciaifelseContext)_localctx).blockinterno.blkint)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(433);
				((SentenciaifelseContext)_localctx).IF = match(IF);
				setState(434);
				((SentenciaifelseContext)_localctx).expr = expr(0);
				setState(435);
				match(LLAVEIZQ);
				setState(436);
				((SentenciaifelseContext)_localctx).ifop = blockinterno();
				setState(437);
				match(LLAVEDER);
				setState(438);
				match(ELSE);
				setState(439);
				match(LLAVEIZQ);
				setState(440);
				((SentenciaifelseContext)_localctx).elseop = blockinterno();
				setState(441);
				match(LLAVEDER);
				 _localctx.myIfElse = sentencias.NewSentenciaIfElse((((SentenciaifelseContext)_localctx).IF!=null?((SentenciaifelseContext)_localctx).IF.getLine():0), (((SentenciaifelseContext)_localctx).IF!=null?((SentenciaifelseContext)_localctx).IF.getCharPositionInLine():0), ((SentenciaifelseContext)_localctx).expr.e, ((SentenciaifelseContext)_localctx).ifop.blkint , ((SentenciaifelseContext)_localctx).elseop.blkint)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(444);
				((SentenciaifelseContext)_localctx).IF = match(IF);
				setState(445);
				((SentenciaifelseContext)_localctx).expr = expr(0);
				setState(446);
				match(LLAVEIZQ);
				setState(447);
				((SentenciaifelseContext)_localctx).blockinterno = blockinterno();
				setState(448);
				match(LLAVEDER);
				setState(449);
				match(ELSE);
				setState(450);
				((SentenciaifelseContext)_localctx).sentenciaifelse = sentenciaifelse();
				 _localctx.myIfElse = sentencias.NewSentenciaIfElseIf((((SentenciaifelseContext)_localctx).IF!=null?((SentenciaifelseContext)_localctx).IF.getLine():0), (((SentenciaifelseContext)_localctx).IF!=null?((SentenciaifelseContext)_localctx).IF.getCharPositionInLine():0), ((SentenciaifelseContext)_localctx).expr.e, ((SentenciaifelseContext)_localctx).blockinterno.blkint, ((SentenciaifelseContext)_localctx).sentenciaifelse.myIfElse)
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SwitchcontrolContext extends ParserRuleContext {
		public interfaces.Instruction mySwitch;
		public Token SWITCH;
		public ExprContext expr;
		public BlockcaseContext blockcase;
		public Token DEFAULT;
		public BlockinternoContext blockinterno;
		public TerminalNode SWITCH() { return getToken(SwiftGrammarParser.SWITCH, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockcaseContext blockcase() {
			return getRuleContext(BlockcaseContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public TerminalNode DEFAULT() { return getToken(SwiftGrammarParser.DEFAULT, 0); }
		public TerminalNode DOS_PUNTOS() { return getToken(SwiftGrammarParser.DOS_PUNTOS, 0); }
		public BlockinternoContext blockinterno() {
			return getRuleContext(BlockinternoContext.class,0);
		}
		public SwitchcontrolContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switchcontrol; }
	}

	public final SwitchcontrolContext switchcontrol() throws RecognitionException {
		SwitchcontrolContext _localctx = new SwitchcontrolContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_switchcontrol);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(455);
			((SwitchcontrolContext)_localctx).SWITCH = match(SWITCH);
			setState(456);
			((SwitchcontrolContext)_localctx).expr = expr(0);
			setState(457);
			match(LLAVEIZQ);
			setState(458);
			((SwitchcontrolContext)_localctx).blockcase = blockcase();
			setState(462);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(459);
				((SwitchcontrolContext)_localctx).DEFAULT = match(DEFAULT);
				setState(460);
				match(DOS_PUNTOS);
				setState(461);
				((SwitchcontrolContext)_localctx).blockinterno = blockinterno();
				}
			}

			setState(464);
			match(LLAVEDER);

			    if (((SwitchcontrolContext)_localctx).DEFAULT != nil) {
			        _localctx.mySwitch = sentencias.NewSentenciaSwitchDefault((((SwitchcontrolContext)_localctx).SWITCH!=null?((SwitchcontrolContext)_localctx).SWITCH.getLine():0), (((SwitchcontrolContext)_localctx).SWITCH!=null?((SwitchcontrolContext)_localctx).SWITCH.getCharPositionInLine():0), ((SwitchcontrolContext)_localctx).expr.e, ((SwitchcontrolContext)_localctx).blockcase.blkcase, ((SwitchcontrolContext)_localctx).blockinterno.blkint)
			    } else {
			        _localctx.mySwitch = sentencias.NewSentenciaSwitch((((SwitchcontrolContext)_localctx).SWITCH!=null?((SwitchcontrolContext)_localctx).SWITCH.getLine():0), (((SwitchcontrolContext)_localctx).SWITCH!=null?((SwitchcontrolContext)_localctx).SWITCH.getCharPositionInLine():0), ((SwitchcontrolContext)_localctx).expr.e, ((SwitchcontrolContext)_localctx).blockcase.blkcase)
			    }

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockcaseContext extends ParserRuleContext {
		public []interface{} blkcase;
		public BloquecaseContext bloquecase;
		public List<BloquecaseContext> blocas = new ArrayList<BloquecaseContext>();
		public List<BloquecaseContext> bloquecase() {
			return getRuleContexts(BloquecaseContext.class);
		}
		public BloquecaseContext bloquecase(int i) {
			return getRuleContext(BloquecaseContext.class,i);
		}
		public BlockcaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_blockcase; }
	}

	public final BlockcaseContext blockcase() throws RecognitionException {
		BlockcaseContext _localctx = new BlockcaseContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_blockcase);

		    _localctx.blkcase = []interface{}{}
		    var listInt []IBloquecaseContext

		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(468); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(467);
				((BlockcaseContext)_localctx).bloquecase = bloquecase();
				((BlockcaseContext)_localctx).blocas.add(((BlockcaseContext)_localctx).bloquecase);
				}
				}
				setState(470); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==CASE );

			    listInt = localctx.(*BlockcaseContext).GetBlocas()
			    for _, e := range listInt {
			        _localctx.blkcase = append(_localctx.blkcase, e.GetBlocas())
			    }

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BloquecaseContext extends ParserRuleContext {
		public interfaces.Instruction blocas;
		public Token CASE;
		public ExprContext expr;
		public BlockinternoContext blockinterno;
		public TerminalNode CASE() { return getToken(SwiftGrammarParser.CASE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode DOS_PUNTOS() { return getToken(SwiftGrammarParser.DOS_PUNTOS, 0); }
		public BlockinternoContext blockinterno() {
			return getRuleContext(BlockinternoContext.class,0);
		}
		public BloquecaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bloquecase; }
	}

	public final BloquecaseContext bloquecase() throws RecognitionException {
		BloquecaseContext _localctx = new BloquecaseContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_bloquecase);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(474);
			((BloquecaseContext)_localctx).CASE = match(CASE);
			setState(475);
			((BloquecaseContext)_localctx).expr = expr(0);
			setState(476);
			match(DOS_PUNTOS);
			setState(477);
			((BloquecaseContext)_localctx).blockinterno = blockinterno();

			    _localctx.blocas=sentencias.NewSentenciaSwitchCase((((BloquecaseContext)_localctx).CASE!=null?((BloquecaseContext)_localctx).CASE.getLine():0) ,(((BloquecaseContext)_localctx).CASE!=null?((BloquecaseContext)_localctx).CASE.getCharPositionInLine():0), ((BloquecaseContext)_localctx).expr.e, ((BloquecaseContext)_localctx).blockinterno.blkint)

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class WhilecontrolContext extends ParserRuleContext {
		public interfaces.Instruction whict;
		public Token WHILE;
		public ExprContext expr;
		public BlockinternoContext blockinterno;
		public TerminalNode WHILE() { return getToken(SwiftGrammarParser.WHILE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockinternoContext blockinterno() {
			return getRuleContext(BlockinternoContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public WhilecontrolContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whilecontrol; }
	}

	public final WhilecontrolContext whilecontrol() throws RecognitionException {
		WhilecontrolContext _localctx = new WhilecontrolContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_whilecontrol);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(480);
			((WhilecontrolContext)_localctx).WHILE = match(WHILE);
			setState(481);
			((WhilecontrolContext)_localctx).expr = expr(0);
			setState(482);
			match(LLAVEIZQ);
			setState(483);
			((WhilecontrolContext)_localctx).blockinterno = blockinterno();
			setState(484);
			match(LLAVEDER);
			 _localctx.whict = sentencias.NewSentenciaWhile((((WhilecontrolContext)_localctx).WHILE!=null?((WhilecontrolContext)_localctx).WHILE.getLine():0), (((WhilecontrolContext)_localctx).WHILE!=null?((WhilecontrolContext)_localctx).WHILE.getCharPositionInLine():0), ((WhilecontrolContext)_localctx).expr.e, ((WhilecontrolContext)_localctx).blockinterno.blkint)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ForcontrolContext extends ParserRuleContext {
		public interfaces.Instruction forct;
		public Token FOR;
		public Token ID_VALIDO;
		public ExprContext left;
		public ExprContext right;
		public BlockinternoContext blockinterno;
		public Token op1;
		public Token op2;
		public ExprContext expr;
		public TerminalNode FOR() { return getToken(SwiftGrammarParser.FOR, 0); }
		public List<TerminalNode> ID_VALIDO() { return getTokens(SwiftGrammarParser.ID_VALIDO); }
		public TerminalNode ID_VALIDO(int i) {
			return getToken(SwiftGrammarParser.ID_VALIDO, i);
		}
		public TerminalNode IN() { return getToken(SwiftGrammarParser.IN, 0); }
		public TerminalNode RANGO() { return getToken(SwiftGrammarParser.RANGO, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockinternoContext blockinterno() {
			return getRuleContext(BlockinternoContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public ForcontrolContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forcontrol; }
	}

	public final ForcontrolContext forcontrol() throws RecognitionException {
		ForcontrolContext _localctx = new ForcontrolContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_forcontrol);
		try {
			setState(516);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(487);
				((ForcontrolContext)_localctx).FOR = match(FOR);
				setState(488);
				((ForcontrolContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(489);
				match(IN);
				setState(490);
				((ForcontrolContext)_localctx).left = expr(0);
				setState(491);
				match(RANGO);
				setState(492);
				((ForcontrolContext)_localctx).right = expr(0);
				setState(493);
				match(LLAVEIZQ);
				setState(494);
				((ForcontrolContext)_localctx).blockinterno = blockinterno();
				setState(495);
				match(LLAVEDER);
				 _localctx.forct = sentencias.NewSentenciaForRango((((ForcontrolContext)_localctx).FOR!=null?((ForcontrolContext)_localctx).FOR.getLine():0), (((ForcontrolContext)_localctx).FOR!=null?((ForcontrolContext)_localctx).FOR.getCharPositionInLine():0), (((ForcontrolContext)_localctx).ID_VALIDO!=null?((ForcontrolContext)_localctx).ID_VALIDO.getText():null), ((ForcontrolContext)_localctx).left.e, ((ForcontrolContext)_localctx).right.e,((ForcontrolContext)_localctx).blockinterno.blkint)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(498);
				((ForcontrolContext)_localctx).FOR = match(FOR);
				setState(499);
				((ForcontrolContext)_localctx).op1 = match(ID_VALIDO);
				setState(500);
				match(IN);
				setState(501);
				((ForcontrolContext)_localctx).op2 = match(ID_VALIDO);
				setState(502);
				match(LLAVEIZQ);
				setState(503);
				((ForcontrolContext)_localctx).blockinterno = blockinterno();
				setState(504);
				match(LLAVEDER);
				 _localctx.forct = sentencias.NewSentenciaForId((((ForcontrolContext)_localctx).FOR!=null?((ForcontrolContext)_localctx).FOR.getLine():0), (((ForcontrolContext)_localctx).FOR!=null?((ForcontrolContext)_localctx).FOR.getCharPositionInLine():0), (((ForcontrolContext)_localctx).op1!=null?((ForcontrolContext)_localctx).op1.getText():null), (((ForcontrolContext)_localctx).op2!=null?((ForcontrolContext)_localctx).op2.getText():null), ((ForcontrolContext)_localctx).blockinterno.blkint)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(507);
				((ForcontrolContext)_localctx).FOR = match(FOR);
				setState(508);
				((ForcontrolContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(509);
				match(IN);
				setState(510);
				((ForcontrolContext)_localctx).expr = expr(0);
				setState(511);
				match(LLAVEIZQ);
				setState(512);
				((ForcontrolContext)_localctx).blockinterno = blockinterno();
				setState(513);
				match(LLAVEDER);
				 _localctx.forct = sentencias.NewSentenciaForCadena((((ForcontrolContext)_localctx).FOR!=null?((ForcontrolContext)_localctx).FOR.getLine():0), (((ForcontrolContext)_localctx).FOR!=null?((ForcontrolContext)_localctx).FOR.getCharPositionInLine():0), (((ForcontrolContext)_localctx).ID_VALIDO!=null?((ForcontrolContext)_localctx).ID_VALIDO.getText():null), ((ForcontrolContext)_localctx).expr.e, ((ForcontrolContext)_localctx).blockinterno.blkint)
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class GuardcontrolContext extends ParserRuleContext {
		public interfaces.Instruction guct;
		public Token GUARD;
		public ExprContext expr;
		public BlockinternoContext blockinterno;
		public TerminalNode GUARD() { return getToken(SwiftGrammarParser.GUARD, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(SwiftGrammarParser.ELSE, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockinternoContext blockinterno() {
			return getRuleContext(BlockinternoContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public GuardcontrolContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_guardcontrol; }
	}

	public final GuardcontrolContext guardcontrol() throws RecognitionException {
		GuardcontrolContext _localctx = new GuardcontrolContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_guardcontrol);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(518);
			((GuardcontrolContext)_localctx).GUARD = match(GUARD);
			setState(519);
			((GuardcontrolContext)_localctx).expr = expr(0);
			setState(520);
			match(ELSE);
			setState(521);
			match(LLAVEIZQ);
			setState(522);
			((GuardcontrolContext)_localctx).blockinterno = blockinterno();
			setState(523);
			match(LLAVEDER);
			 
			    _localctx.guct = sentencias.NewSentenciaGuard((((GuardcontrolContext)_localctx).GUARD!=null?((GuardcontrolContext)_localctx).GUARD.getLine():0), (((GuardcontrolContext)_localctx).GUARD!=null?((GuardcontrolContext)_localctx).GUARD.getCharPositionInLine():0), ((GuardcontrolContext)_localctx).expr.e, ((GuardcontrolContext)_localctx).blockinterno.blkint)

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ContinueeContext extends ParserRuleContext {
		public interfaces.Instruction coct;
		public Token CONTINUE;
		public TerminalNode CONTINUE() { return getToken(SwiftGrammarParser.CONTINUE, 0); }
		public ContinueeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_continuee; }
	}

	public final ContinueeContext continuee() throws RecognitionException {
		ContinueeContext _localctx = new ContinueeContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_continuee);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(526);
			((ContinueeContext)_localctx).CONTINUE = match(CONTINUE);
			_localctx.coct = sentencias.NewTransferenciaContinue((((ContinueeContext)_localctx).CONTINUE!=null?((ContinueeContext)_localctx).CONTINUE.getLine():0), (((ContinueeContext)_localctx).CONTINUE!=null?((ContinueeContext)_localctx).CONTINUE.getCharPositionInLine():0))
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BreakkContext extends ParserRuleContext {
		public interfaces.Instruction brkct;
		public Token BREAK;
		public TerminalNode BREAK() { return getToken(SwiftGrammarParser.BREAK, 0); }
		public BreakkContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_breakk; }
	}

	public final BreakkContext breakk() throws RecognitionException {
		BreakkContext _localctx = new BreakkContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_breakk);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(529);
			((BreakkContext)_localctx).BREAK = match(BREAK);
			 _localctx.brkct = sentencias.NewTransferenciaBreak((((BreakkContext)_localctx).BREAK!=null?((BreakkContext)_localctx).BREAK.getLine():0), (((BreakkContext)_localctx).BREAK!=null?((BreakkContext)_localctx).BREAK.getCharPositionInLine():0))
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RetornosContext extends ParserRuleContext {
		public interfaces.Instruction rect;
		public Token RETURN;
		public ExprContext op;
		public TerminalNode RETURN() { return getToken(SwiftGrammarParser.RETURN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public RetornosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_retornos; }
	}

	public final RetornosContext retornos() throws RecognitionException {
		RetornosContext _localctx = new RetornosContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_retornos);
		try {
			setState(538);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(532);
				((RetornosContext)_localctx).RETURN = match(RETURN);
				setState(533);
				((RetornosContext)_localctx).op = expr(0);

				    ((RetornosContext)_localctx).rect =  sentencias.NewTransferenciaReturnExp((((RetornosContext)_localctx).RETURN!=null?((RetornosContext)_localctx).RETURN.getLine():0), (((RetornosContext)_localctx).RETURN!=null?((RetornosContext)_localctx).RETURN.getCharPositionInLine():0), ((RetornosContext)_localctx).op.e);

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(536);
				((RetornosContext)_localctx).RETURN = match(RETURN);

				    ((RetornosContext)_localctx).rect =  sentencias.NewTransferenciaReturn((((RetornosContext)_localctx).RETURN!=null?((RetornosContext)_localctx).RETURN.getLine():0), (((RetornosContext)_localctx).RETURN!=null?((RetornosContext)_localctx).RETURN.getCharPositionInLine():0));

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VectorcontrolContext extends ParserRuleContext {
		public interfaces.Instruction vect;
		public Token VAR;
		public Token ID_VALIDO;
		public TipodatoContext tipodato;
		public BlockparamsContext blockparams;
		public Token prin;
		public Token secu;
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public List<TerminalNode> ID_VALIDO() { return getTokens(SwiftGrammarParser.ID_VALIDO); }
		public TerminalNode ID_VALIDO(int i) {
			return getToken(SwiftGrammarParser.ID_VALIDO, i);
		}
		public TerminalNode DOS_PUNTOS() { return getToken(SwiftGrammarParser.DOS_PUNTOS, 0); }
		public List<TerminalNode> CORCHIZQ() { return getTokens(SwiftGrammarParser.CORCHIZQ); }
		public TerminalNode CORCHIZQ(int i) {
			return getToken(SwiftGrammarParser.CORCHIZQ, i);
		}
		public TipodatoContext tipodato() {
			return getRuleContext(TipodatoContext.class,0);
		}
		public List<TerminalNode> CORCHDER() { return getTokens(SwiftGrammarParser.CORCHDER); }
		public TerminalNode CORCHDER(int i) {
			return getToken(SwiftGrammarParser.CORCHDER, i);
		}
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public BlockparamsContext blockparams() {
			return getRuleContext(BlockparamsContext.class,0);
		}
		public VectorcontrolContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vectorcontrol; }
	}

	public final VectorcontrolContext vectorcontrol() throws RecognitionException {
		VectorcontrolContext _localctx = new VectorcontrolContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_vectorcontrol);
		try {
			setState(573);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(540);
				((VectorcontrolContext)_localctx).VAR = match(VAR);
				setState(541);
				((VectorcontrolContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(542);
				match(DOS_PUNTOS);
				setState(543);
				match(CORCHIZQ);
				setState(544);
				((VectorcontrolContext)_localctx).tipodato = tipodato();
				setState(545);
				match(CORCHDER);
				setState(546);
				match(IG);
				setState(547);
				match(CORCHIZQ);
				setState(548);
				((VectorcontrolContext)_localctx).blockparams = blockparams();
				setState(549);
				match(CORCHDER);
				 _localctx.vect = datoscompuestos.NewArregloDeclaracionLista((((VectorcontrolContext)_localctx).VAR!=null?((VectorcontrolContext)_localctx).VAR.getLine():0) ,(((VectorcontrolContext)_localctx).VAR!=null?((VectorcontrolContext)_localctx).VAR.getCharPositionInLine():0), (((VectorcontrolContext)_localctx).ID_VALIDO!=null?((VectorcontrolContext)_localctx).ID_VALIDO.getText():null) , ((VectorcontrolContext)_localctx).tipodato.tipo, ((VectorcontrolContext)_localctx).blockparams.blkpar)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(552);
				((VectorcontrolContext)_localctx).VAR = match(VAR);
				setState(553);
				((VectorcontrolContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(554);
				match(DOS_PUNTOS);
				setState(555);
				match(CORCHIZQ);
				setState(556);
				((VectorcontrolContext)_localctx).tipodato = tipodato();
				setState(557);
				match(CORCHDER);
				setState(558);
				match(IG);
				setState(559);
				match(CORCHIZQ);
				setState(560);
				match(CORCHDER);
				 _localctx.vect = datoscompuestos.NewArregloDeclaracionSinLista((((VectorcontrolContext)_localctx).VAR!=null?((VectorcontrolContext)_localctx).VAR.getLine():0) ,(((VectorcontrolContext)_localctx).VAR!=null?((VectorcontrolContext)_localctx).VAR.getCharPositionInLine():0), (((VectorcontrolContext)_localctx).ID_VALIDO!=null?((VectorcontrolContext)_localctx).ID_VALIDO.getText():null) , ((VectorcontrolContext)_localctx).tipodato.tipo)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(563);
				((VectorcontrolContext)_localctx).VAR = match(VAR);
				setState(564);
				((VectorcontrolContext)_localctx).prin = match(ID_VALIDO);
				setState(565);
				match(DOS_PUNTOS);
				setState(566);
				match(CORCHIZQ);
				setState(567);
				((VectorcontrolContext)_localctx).tipodato = tipodato();
				setState(568);
				match(CORCHDER);
				setState(569);
				match(IG);
				setState(570);
				((VectorcontrolContext)_localctx).secu = match(ID_VALIDO);
				 _localctx.vect = datoscompuestos.NewArregloDeclaracionId((((VectorcontrolContext)_localctx).VAR!=null?((VectorcontrolContext)_localctx).VAR.getLine():0) ,(((VectorcontrolContext)_localctx).VAR!=null?((VectorcontrolContext)_localctx).VAR.getCharPositionInLine():0), (((VectorcontrolContext)_localctx).prin!=null?((VectorcontrolContext)_localctx).prin.getText():null) , ((VectorcontrolContext)_localctx).tipodato.tipo, (((VectorcontrolContext)_localctx).secu!=null?((VectorcontrolContext)_localctx).secu.getText():null))
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockparamsContext extends ParserRuleContext {
		public []interface{} blkpar;
		public BloqueparamsContext bloqueparams;
		public List<BloqueparamsContext> blopas = new ArrayList<BloqueparamsContext>();
		public List<BloqueparamsContext> bloqueparams() {
			return getRuleContexts(BloqueparamsContext.class);
		}
		public BloqueparamsContext bloqueparams(int i) {
			return getRuleContext(BloqueparamsContext.class,i);
		}
		public BlockparamsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_blockparams; }
	}

	public final BlockparamsContext blockparams() throws RecognitionException {
		BlockparamsContext _localctx = new BlockparamsContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_blockparams);

		    _localctx.blkpar = []interface{}{}
		    var listInt []IBloqueparamsContext

		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(576); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(575);
				((BlockparamsContext)_localctx).bloqueparams = bloqueparams();
				((BlockparamsContext)_localctx).blopas.add(((BlockparamsContext)_localctx).bloqueparams);
				}
				}
				setState(578); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 2307041339449017550L) != 0) || _la==COMA );

			    listInt = localctx.(*BlockparamsContext).GetBlopas()
			    for _, e := range listInt {
			        _localctx.blkpar = append(_localctx.blkpar, e.GetBlopas())
			    }

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BloqueparamsContext extends ParserRuleContext {
		public interfaces.Expression blopas;
		public Token COMA;
		public ExprContext expr;
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public BloqueparamsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bloqueparams; }
	}

	public final BloqueparamsContext bloqueparams() throws RecognitionException {
		BloqueparamsContext _localctx = new BloqueparamsContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_bloqueparams);
		try {
			setState(589);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case COMA:
				enterOuterAlt(_localctx, 1);
				{
				setState(582);
				((BloqueparamsContext)_localctx).COMA = match(COMA);
				setState(583);
				((BloqueparamsContext)_localctx).expr = expr(0);

				    _localctx.blopas = datoscompuestos.NewArregloParametros((((BloqueparamsContext)_localctx).COMA!=null?((BloqueparamsContext)_localctx).COMA.getLine():0) ,(((BloqueparamsContext)_localctx).COMA!=null?((BloqueparamsContext)_localctx).COMA.getCharPositionInLine():0), ((BloqueparamsContext)_localctx).expr.e)

				}
				break;
			case INT:
			case FLOAT:
			case STRING:
			case TRU:
			case FAL:
			case NULO:
			case NUMBER:
			case CADENA:
			case ID_VALIDO:
			case CHARACTER:
			case PARIZQ:
			case NOT:
			case SUB:
				enterOuterAlt(_localctx, 2);
				{
				setState(586);
				((BloqueparamsContext)_localctx).expr = expr(0);

				    _localctx.blopas = datoscompuestos.NewArregloParametro(((BloqueparamsContext)_localctx).expr.e)

				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VectoragregarContext extends ParserRuleContext {
		public interfaces.Instruction veadct;
		public Token ID_VALIDO;
		public ExprContext expr;
		public Token prin;
		public ExprContext pop;
		public Token secu;
		public ExprContext sop;
		public ExprContext op1;
		public ExprContext op2;
		public ListamatrizaddsubsContext listamatrizaddsubs;
		public ExprContext op3;
		public List<TerminalNode> ID_VALIDO() { return getTokens(SwiftGrammarParser.ID_VALIDO); }
		public TerminalNode ID_VALIDO(int i) {
			return getToken(SwiftGrammarParser.ID_VALIDO, i);
		}
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
		public TerminalNode APPEND() { return getToken(SwiftGrammarParser.APPEND, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public List<TerminalNode> CORCHIZQ() { return getTokens(SwiftGrammarParser.CORCHIZQ); }
		public TerminalNode CORCHIZQ(int i) {
			return getToken(SwiftGrammarParser.CORCHIZQ, i);
		}
		public List<TerminalNode> CORCHDER() { return getTokens(SwiftGrammarParser.CORCHDER); }
		public TerminalNode CORCHDER(int i) {
			return getToken(SwiftGrammarParser.CORCHDER, i);
		}
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ListamatrizaddsubsContext listamatrizaddsubs() {
			return getRuleContext(ListamatrizaddsubsContext.class,0);
		}
		public VectoragregarContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vectoragregar; }
	}

	public final VectoragregarContext vectoragregar() throws RecognitionException {
		VectoragregarContext _localctx = new VectoragregarContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_vectoragregar);
		try {
			setState(641);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,36,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(591);
				((VectoragregarContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(592);
				match(PUNTO);
				setState(593);
				match(APPEND);
				setState(594);
				match(PARIZQ);
				setState(595);
				((VectoragregarContext)_localctx).expr = expr(0);
				setState(596);
				match(PARDER);
				 _localctx.veadct = datoscompuestos.NewArregloAppend((((VectoragregarContext)_localctx).ID_VALIDO!=null?((VectoragregarContext)_localctx).ID_VALIDO.getText():null) , ((VectoragregarContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(599);
				((VectoragregarContext)_localctx).prin = match(ID_VALIDO);
				setState(600);
				match(CORCHIZQ);
				setState(601);
				((VectoragregarContext)_localctx).pop = expr(0);
				setState(602);
				match(CORCHDER);
				setState(603);
				match(IG);
				setState(604);
				((VectoragregarContext)_localctx).secu = match(ID_VALIDO);
				setState(605);
				match(CORCHIZQ);
				setState(606);
				((VectoragregarContext)_localctx).sop = expr(0);
				setState(607);
				match(CORCHDER);
				 _localctx.veadct = datoscompuestos.NewArregloAppendArreglo((((VectoragregarContext)_localctx).prin!=null?((VectoragregarContext)_localctx).prin.getText():null) , ((VectoragregarContext)_localctx).pop.e, (((VectoragregarContext)_localctx).secu!=null?((VectoragregarContext)_localctx).secu.getText():null), ((VectoragregarContext)_localctx).sop.e)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(610);
				((VectoragregarContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(611);
				match(CORCHIZQ);
				setState(612);
				((VectoragregarContext)_localctx).op1 = expr(0);
				setState(613);
				match(CORCHDER);
				setState(614);
				match(CORCHIZQ);
				setState(615);
				((VectoragregarContext)_localctx).op2 = expr(0);
				setState(616);
				match(CORCHDER);
				setState(617);
				((VectoragregarContext)_localctx).listamatrizaddsubs = listamatrizaddsubs();
				setState(618);
				match(IG);
				setState(619);
				((VectoragregarContext)_localctx).op3 = expr(0);
				 _localctx.veadct = datoscompuestos.NewMatrizAsignacionList((((VectoragregarContext)_localctx).ID_VALIDO!=null?((VectoragregarContext)_localctx).ID_VALIDO.getText():null), ((VectoragregarContext)_localctx).op1.e, ((VectoragregarContext)_localctx).op2.e, ((VectoragregarContext)_localctx).listamatrizaddsubs.blklimatas, ((VectoragregarContext)_localctx).op3.e) 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(622);
				((VectoragregarContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(623);
				match(CORCHIZQ);
				setState(624);
				((VectoragregarContext)_localctx).op1 = expr(0);
				setState(625);
				match(CORCHDER);
				setState(626);
				match(CORCHIZQ);
				setState(627);
				((VectoragregarContext)_localctx).op2 = expr(0);
				setState(628);
				match(CORCHDER);
				setState(629);
				match(IG);
				setState(630);
				((VectoragregarContext)_localctx).op3 = expr(0);
				 _localctx.veadct = datoscompuestos.NewMatrizAsignacion((((VectoragregarContext)_localctx).ID_VALIDO!=null?((VectoragregarContext)_localctx).ID_VALIDO.getText():null), ((VectoragregarContext)_localctx).op1.e, ((VectoragregarContext)_localctx).op2.e, ((VectoragregarContext)_localctx).op3.e) 
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(633);
				((VectoragregarContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(634);
				match(CORCHIZQ);
				setState(635);
				((VectoragregarContext)_localctx).pop = expr(0);
				setState(636);
				match(CORCHDER);
				setState(637);
				match(IG);
				setState(638);
				((VectoragregarContext)_localctx).sop = expr(0);
				 _localctx.veadct = datoscompuestos.NewArregloAppendExp((((VectoragregarContext)_localctx).ID_VALIDO!=null?((VectoragregarContext)_localctx).ID_VALIDO.getText():null) , ((VectoragregarContext)_localctx).pop.e, ((VectoragregarContext)_localctx).sop.e)
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VectorremoverContext extends ParserRuleContext {
		public interfaces.Instruction vermct;
		public Token ID_VALIDO;
		public Token PUNTO;
		public ExprContext expr;
		public TerminalNode ID_VALIDO() { return getToken(SwiftGrammarParser.ID_VALIDO, 0); }
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
		public TerminalNode REMOVELAST() { return getToken(SwiftGrammarParser.REMOVELAST, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public TerminalNode REMOVE() { return getToken(SwiftGrammarParser.REMOVE, 0); }
		public TerminalNode AT() { return getToken(SwiftGrammarParser.AT, 0); }
		public TerminalNode DOS_PUNTOS() { return getToken(SwiftGrammarParser.DOS_PUNTOS, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public VectorremoverContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vectorremover; }
	}

	public final VectorremoverContext vectorremover() throws RecognitionException {
		VectorremoverContext _localctx = new VectorremoverContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_vectorremover);
		try {
			setState(659);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,37,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(643);
				((VectorremoverContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(644);
				((VectorremoverContext)_localctx).PUNTO = match(PUNTO);
				setState(645);
				match(REMOVELAST);
				setState(646);
				match(PARIZQ);
				setState(647);
				match(PARDER);
				 _localctx.vermct = datoscompuestos.NewArregloRemoveLast((((VectorremoverContext)_localctx).PUNTO!=null?((VectorremoverContext)_localctx).PUNTO.getLine():0), (((VectorremoverContext)_localctx).PUNTO!=null?((VectorremoverContext)_localctx).PUNTO.getCharPositionInLine():0), (((VectorremoverContext)_localctx).ID_VALIDO!=null?((VectorremoverContext)_localctx).ID_VALIDO.getText():null))
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(649);
				((VectorremoverContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(650);
				((VectorremoverContext)_localctx).PUNTO = match(PUNTO);
				setState(651);
				match(REMOVE);
				setState(652);
				match(PARIZQ);
				setState(653);
				match(AT);
				setState(654);
				match(DOS_PUNTOS);
				setState(655);
				((VectorremoverContext)_localctx).expr = expr(0);
				setState(656);
				match(PARDER);
				 _localctx.vermct = datoscompuestos.NewArregloRemovePos((((VectorremoverContext)_localctx).PUNTO!=null?((VectorremoverContext)_localctx).PUNTO.getLine():0), (((VectorremoverContext)_localctx).PUNTO!=null?((VectorremoverContext)_localctx).PUNTO.getCharPositionInLine():0), (((VectorremoverContext)_localctx).ID_VALIDO!=null?((VectorremoverContext)_localctx).ID_VALIDO.getText():null), ((VectorremoverContext)_localctx).expr.e)
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VectorvacioContext extends ParserRuleContext {
		public interfaces.Expression veemct;
		public Token ID_VALIDO;
		public Token PUNTO;
		public TerminalNode ID_VALIDO() { return getToken(SwiftGrammarParser.ID_VALIDO, 0); }
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
		public TerminalNode ISEMPTY() { return getToken(SwiftGrammarParser.ISEMPTY, 0); }
		public VectorvacioContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vectorvacio; }
	}

	public final VectorvacioContext vectorvacio() throws RecognitionException {
		VectorvacioContext _localctx = new VectorvacioContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_vectorvacio);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(661);
			((VectorvacioContext)_localctx).ID_VALIDO = match(ID_VALIDO);
			setState(662);
			((VectorvacioContext)_localctx).PUNTO = match(PUNTO);
			setState(663);
			match(ISEMPTY);
			 _localctx.veemct = datoscompuestos.NewArregloIsEmpty((((VectorvacioContext)_localctx).PUNTO!=null?((VectorvacioContext)_localctx).PUNTO.getLine():0), (((VectorvacioContext)_localctx).PUNTO!=null?((VectorvacioContext)_localctx).PUNTO.getCharPositionInLine():0), (((VectorvacioContext)_localctx).ID_VALIDO!=null?((VectorvacioContext)_localctx).ID_VALIDO.getText():null))
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VectorcountContext extends ParserRuleContext {
		public interfaces.Expression vecnct;
		public Token ID_VALIDO;
		public Token PUNTO;
		public TerminalNode ID_VALIDO() { return getToken(SwiftGrammarParser.ID_VALIDO, 0); }
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
		public TerminalNode COUNT() { return getToken(SwiftGrammarParser.COUNT, 0); }
		public VectorcountContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vectorcount; }
	}

	public final VectorcountContext vectorcount() throws RecognitionException {
		VectorcountContext _localctx = new VectorcountContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_vectorcount);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(666);
			((VectorcountContext)_localctx).ID_VALIDO = match(ID_VALIDO);
			setState(667);
			((VectorcountContext)_localctx).PUNTO = match(PUNTO);
			setState(668);
			match(COUNT);
			 _localctx.vecnct = datoscompuestos.NewArregloCount((((VectorcountContext)_localctx).PUNTO!=null?((VectorcountContext)_localctx).PUNTO.getLine():0), (((VectorcountContext)_localctx).PUNTO!=null?((VectorcountContext)_localctx).PUNTO.getCharPositionInLine():0), (((VectorcountContext)_localctx).ID_VALIDO!=null?((VectorcountContext)_localctx).ID_VALIDO.getText():null))
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VectoraccessContext extends ParserRuleContext {
		public interfaces.Expression vepposct;
		public Token ID_VALIDO;
		public ExprContext op1;
		public ExprContext op2;
		public ListamatrizaddsubsContext listamatrizaddsubs;
		public ExprContext expr;
		public Token CORCHDER;
		public TerminalNode ID_VALIDO() { return getToken(SwiftGrammarParser.ID_VALIDO, 0); }
		public List<TerminalNode> CORCHIZQ() { return getTokens(SwiftGrammarParser.CORCHIZQ); }
		public TerminalNode CORCHIZQ(int i) {
			return getToken(SwiftGrammarParser.CORCHIZQ, i);
		}
		public List<TerminalNode> CORCHDER() { return getTokens(SwiftGrammarParser.CORCHDER); }
		public TerminalNode CORCHDER(int i) {
			return getToken(SwiftGrammarParser.CORCHDER, i);
		}
		public ListamatrizaddsubsContext listamatrizaddsubs() {
			return getRuleContext(ListamatrizaddsubsContext.class,0);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public VectoraccessContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vectoraccess; }
	}

	public final VectoraccessContext vectoraccess() throws RecognitionException {
		VectoraccessContext _localctx = new VectoraccessContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_vectoraccess);
		try {
			setState(696);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,38,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(671);
				((VectoraccessContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(672);
				match(CORCHIZQ);
				setState(673);
				((VectoraccessContext)_localctx).op1 = expr(0);
				setState(674);
				match(CORCHDER);
				setState(675);
				match(CORCHIZQ);
				setState(676);
				((VectoraccessContext)_localctx).op2 = expr(0);
				setState(677);
				match(CORCHDER);
				setState(678);
				((VectoraccessContext)_localctx).listamatrizaddsubs = listamatrizaddsubs();
				 _localctx.vepposct = datoscompuestos.NewMatrizObtencionList((((VectoraccessContext)_localctx).ID_VALIDO!=null?((VectoraccessContext)_localctx).ID_VALIDO.getText():null), ((VectoraccessContext)_localctx).op1.e, ((VectoraccessContext)_localctx).op2.e, ((VectoraccessContext)_localctx).listamatrizaddsubs.blklimatas) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(681);
				((VectoraccessContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(682);
				match(CORCHIZQ);
				setState(683);
				((VectoraccessContext)_localctx).op1 = expr(0);
				setState(684);
				match(CORCHDER);
				setState(685);
				match(CORCHIZQ);
				setState(686);
				((VectoraccessContext)_localctx).op2 = expr(0);
				setState(687);
				match(CORCHDER);
				 _localctx.vepposct = datoscompuestos.NewMatrizObtencion((((VectoraccessContext)_localctx).ID_VALIDO!=null?((VectoraccessContext)_localctx).ID_VALIDO.getText():null), ((VectoraccessContext)_localctx).op1.e, ((VectoraccessContext)_localctx).op2.e) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(690);
				((VectoraccessContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(691);
				match(CORCHIZQ);
				setState(692);
				((VectoraccessContext)_localctx).expr = expr(0);
				setState(693);
				((VectoraccessContext)_localctx).CORCHDER = match(CORCHDER);
				 _localctx.vepposct = datoscompuestos.NewArregloAccess((((VectoraccessContext)_localctx).CORCHDER!=null?((VectoraccessContext)_localctx).CORCHDER.getLine():0), (((VectoraccessContext)_localctx).CORCHDER!=null?((VectoraccessContext)_localctx).CORCHDER.getCharPositionInLine():0), (((VectoraccessContext)_localctx).ID_VALIDO!=null?((VectoraccessContext)_localctx).ID_VALIDO.getText():null), ((VectoraccessContext)_localctx).expr.e)
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MatrizcontrolContext extends ParserRuleContext {
		public interfaces.Instruction matct;
		public Token VAR;
		public Token ID_VALIDO;
		public Token DOS_PUNTOS;
		public TipomatrizContext tipomatriz;
		public DefmatrizContext defmatriz;
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TerminalNode ID_VALIDO() { return getToken(SwiftGrammarParser.ID_VALIDO, 0); }
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public DefmatrizContext defmatriz() {
			return getRuleContext(DefmatrizContext.class,0);
		}
		public TerminalNode DOS_PUNTOS() { return getToken(SwiftGrammarParser.DOS_PUNTOS, 0); }
		public TipomatrizContext tipomatriz() {
			return getRuleContext(TipomatrizContext.class,0);
		}
		public MatrizcontrolContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_matrizcontrol; }
	}

	public final MatrizcontrolContext matrizcontrol() throws RecognitionException {
		MatrizcontrolContext _localctx = new MatrizcontrolContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_matrizcontrol);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(698);
			((MatrizcontrolContext)_localctx).VAR = match(VAR);
			setState(699);
			((MatrizcontrolContext)_localctx).ID_VALIDO = match(ID_VALIDO);
			setState(702);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DOS_PUNTOS) {
				{
				setState(700);
				((MatrizcontrolContext)_localctx).DOS_PUNTOS = match(DOS_PUNTOS);
				setState(701);
				((MatrizcontrolContext)_localctx).tipomatriz = tipomatriz();
				}
			}

			setState(704);
			match(IG);
			setState(705);
			((MatrizcontrolContext)_localctx).defmatriz = defmatriz();

			    if (((MatrizcontrolContext)_localctx).DOS_PUNTOS != nil) {
			        _localctx.matct = datoscompuestos.NewMatrizDeclaracion((((MatrizcontrolContext)_localctx).VAR!=null?((MatrizcontrolContext)_localctx).VAR.getLine():0), (((MatrizcontrolContext)_localctx).VAR!=null?((MatrizcontrolContext)_localctx).VAR.getCharPositionInLine():0), (((MatrizcontrolContext)_localctx).ID_VALIDO!=null?((MatrizcontrolContext)_localctx).ID_VALIDO.getText():null) ,((MatrizcontrolContext)_localctx).tipomatriz.tipomat, ((MatrizcontrolContext)_localctx).defmatriz.defmat)
			    } else {
			        fmt.Println("Nada")
			        //_localctx.matct = datoscompuestos.NewMatrizDeclaracionSinTipo((((MatrizcontrolContext)_localctx).VAR!=null?((MatrizcontrolContext)_localctx).VAR.getLine():0), (((MatrizcontrolContext)_localctx).VAR!=null?((MatrizcontrolContext)_localctx).VAR.getCharPositionInLine():0), (((MatrizcontrolContext)_localctx).ID_VALIDO!=null?((MatrizcontrolContext)_localctx).ID_VALIDO.getText():null) , ((MatrizcontrolContext)_localctx).defmatriz.defmat)
			    }

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TipomatrizContext extends ParserRuleContext {
		public interfaces.Expression tipomat;
		public Token CORCHIZQ;
		public TipomatrizContext tipomatriz;
		public TipodatoContext tipodato;
		public TerminalNode CORCHIZQ() { return getToken(SwiftGrammarParser.CORCHIZQ, 0); }
		public TipomatrizContext tipomatriz() {
			return getRuleContext(TipomatrizContext.class,0);
		}
		public TerminalNode CORCHDER() { return getToken(SwiftGrammarParser.CORCHDER, 0); }
		public TipodatoContext tipodato() {
			return getRuleContext(TipodatoContext.class,0);
		}
		public TipomatrizContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipomatriz; }
	}

	public final TipomatrizContext tipomatriz() throws RecognitionException {
		TipomatrizContext _localctx = new TipomatrizContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_tipomatriz);
		try {
			setState(718);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,40,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(708);
				((TipomatrizContext)_localctx).CORCHIZQ = match(CORCHIZQ);
				setState(709);
				((TipomatrizContext)_localctx).tipomatriz = tipomatriz();
				setState(710);
				match(CORCHDER);
				 
				    _localctx.tipomat = datoscompuestos.NewMatrizDimension((((TipomatrizContext)_localctx).CORCHIZQ!=null?((TipomatrizContext)_localctx).CORCHIZQ.getLine():0), (((TipomatrizContext)_localctx).CORCHIZQ!=null?((TipomatrizContext)_localctx).CORCHIZQ.getCharPositionInLine():0), ((TipomatrizContext)_localctx).tipomatriz.tipomat)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(713);
				((TipomatrizContext)_localctx).CORCHIZQ = match(CORCHIZQ);
				setState(714);
				((TipomatrizContext)_localctx).tipodato = tipodato();
				setState(715);
				match(CORCHDER);
				 
				    _localctx.tipomat = datoscompuestos.NewMatrizTipo((((TipomatrizContext)_localctx).CORCHIZQ!=null?((TipomatrizContext)_localctx).CORCHIZQ.getLine():0), (((TipomatrizContext)_localctx).CORCHIZQ!=null?((TipomatrizContext)_localctx).CORCHIZQ.getCharPositionInLine():0), ((TipomatrizContext)_localctx).tipodato.tipo)

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DefmatrizContext extends ParserRuleContext {
		public interfaces.Instruction defmat;
		public ListavaloresmatContext listavaloresmat;
		public ListavaloresmatContext listavaloresmat() {
			return getRuleContext(ListavaloresmatContext.class,0);
		}
		public DefmatrizContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_defmatriz; }
	}

	public final DefmatrizContext defmatriz() throws RecognitionException {
		DefmatrizContext _localctx = new DefmatrizContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_defmatriz);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(720);
			((DefmatrizContext)_localctx).listavaloresmat = listavaloresmat();
			 _localctx.defmat = ((DefmatrizContext)_localctx).listavaloresmat.listvlamat
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ListavaloresmatContext extends ParserRuleContext {
		public interfaces.Instruction listvlamat;
		public Listavaloresmat2Context listavaloresmat2;
		public SimplematrizContext simplematriz;
		public TerminalNode CORCHIZQ() { return getToken(SwiftGrammarParser.CORCHIZQ, 0); }
		public Listavaloresmat2Context listavaloresmat2() {
			return getRuleContext(Listavaloresmat2Context.class,0);
		}
		public TerminalNode CORCHDER() { return getToken(SwiftGrammarParser.CORCHDER, 0); }
		public SimplematrizContext simplematriz() {
			return getRuleContext(SimplematrizContext.class,0);
		}
		public ListavaloresmatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listavaloresmat; }
	}

	public final ListavaloresmatContext listavaloresmat() throws RecognitionException {
		ListavaloresmatContext _localctx = new ListavaloresmatContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_listavaloresmat);
		try {
			setState(731);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,41,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(723);
				match(CORCHIZQ);
				setState(724);
				((ListavaloresmatContext)_localctx).listavaloresmat2 = listavaloresmat2(0);
				setState(725);
				match(CORCHDER);
				 _localctx.listvlamat = ((ListavaloresmatContext)_localctx).listavaloresmat2.mylisttmatt
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(728);
				((ListavaloresmatContext)_localctx).simplematriz = simplematriz();
				 _localctx.listvlamat = ((ListavaloresmatContext)_localctx).simplematriz.simmat
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Listavaloresmat2Context extends ParserRuleContext {
		public interfaces.Instruction mylisttmatt;
		public Listavaloresmat2Context op;
		public ListavaloresmatContext listavaloresmat;
		public ListaexpresionsContext listaexpresions;
		public ListavaloresmatContext listavaloresmat() {
			return getRuleContext(ListavaloresmatContext.class,0);
		}
		public ListaexpresionsContext listaexpresions() {
			return getRuleContext(ListaexpresionsContext.class,0);
		}
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public Listavaloresmat2Context listavaloresmat2() {
			return getRuleContext(Listavaloresmat2Context.class,0);
		}
		public Listavaloresmat2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listavaloresmat2; }
	}

	public final Listavaloresmat2Context listavaloresmat2() throws RecognitionException {
		return listavaloresmat2(0);
	}

	private Listavaloresmat2Context listavaloresmat2(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Listavaloresmat2Context _localctx = new Listavaloresmat2Context(_ctx, _parentState);
		Listavaloresmat2Context _prevctx = _localctx;
		int _startState = 64;
		enterRecursionRule(_localctx, 64, RULE_listavaloresmat2, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(740);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CORCHIZQ:
				{
				setState(734);
				((Listavaloresmat2Context)_localctx).listavaloresmat = listavaloresmat();
				 _localctx.mylisttmatt = datoscompuestos.NewMatrizListaNivel(((Listavaloresmat2Context)_localctx).listavaloresmat.listvlamat)
				}
				break;
			case INT:
			case FLOAT:
			case STRING:
			case TRU:
			case FAL:
			case NULO:
			case NUMBER:
			case CADENA:
			case ID_VALIDO:
			case CHARACTER:
			case PARIZQ:
			case NOT:
			case SUB:
			case COMA:
				{
				setState(737);
				((Listavaloresmat2Context)_localctx).listaexpresions = listaexpresions();
				 _localctx.mylisttmatt = datoscompuestos.NewMatrizListaExpresion(((Listavaloresmat2Context)_localctx).listaexpresions.blkparf)
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(749);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,43,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Listavaloresmat2Context(_parentctx, _parentState);
					_localctx.op = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listavaloresmat2);
					setState(742);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(743);
					match(COMA);
					setState(744);
					((Listavaloresmat2Context)_localctx).listavaloresmat = listavaloresmat();
					 _localctx.mylisttmatt = datoscompuestos.NewMatrizListaExpresionList(((Listavaloresmat2Context)_localctx).op.mylisttmatt, ((Listavaloresmat2Context)_localctx).listavaloresmat.listvlamat)
					}
					} 
				}
				setState(751);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,43,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ListaexpresionsContext extends ParserRuleContext {
		public []interface{} blkparf;
		public ListaexpresionContext listaexpresion;
		public List<ListaexpresionContext> funpar = new ArrayList<ListaexpresionContext>();
		public List<ListaexpresionContext> listaexpresion() {
			return getRuleContexts(ListaexpresionContext.class);
		}
		public ListaexpresionContext listaexpresion(int i) {
			return getRuleContext(ListaexpresionContext.class,i);
		}
		public ListaexpresionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listaexpresions; }
	}

	public final ListaexpresionsContext listaexpresions() throws RecognitionException {
		ListaexpresionsContext _localctx = new ListaexpresionsContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_listaexpresions);

		    _localctx.blkparf = []interface{}{}
		    var listInt []IListaexpresionContext

		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(753); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(752);
					((ListaexpresionsContext)_localctx).listaexpresion = listaexpresion();
					((ListaexpresionsContext)_localctx).funpar.add(((ListaexpresionsContext)_localctx).listaexpresion);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(755); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,44,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );

			    listInt = localctx.(*ListaexpresionsContext).GetFunpar()
			    for _, e := range listInt {
			        _localctx.blkparf = append(_localctx.blkparf, e.GetFunpar())
			    }

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ListaexpresionContext extends ParserRuleContext {
		public interfaces.Expression funpar;
		public Token COMA;
		public ExprContext expr;
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ListaexpresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listaexpresion; }
	}

	public final ListaexpresionContext listaexpresion() throws RecognitionException {
		ListaexpresionContext _localctx = new ListaexpresionContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_listaexpresion);
		try {
			setState(766);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case COMA:
				enterOuterAlt(_localctx, 1);
				{
				setState(759);
				((ListaexpresionContext)_localctx).COMA = match(COMA);
				setState(760);
				((ListaexpresionContext)_localctx).expr = expr(0);

				    _localctx.funpar = datoscompuestos.NewArregloParametros((((ListaexpresionContext)_localctx).COMA!=null?((ListaexpresionContext)_localctx).COMA.getLine():0) ,(((ListaexpresionContext)_localctx).COMA!=null?((ListaexpresionContext)_localctx).COMA.getCharPositionInLine():0), ((ListaexpresionContext)_localctx).expr.e)

				}
				break;
			case INT:
			case FLOAT:
			case STRING:
			case TRU:
			case FAL:
			case NULO:
			case NUMBER:
			case CADENA:
			case ID_VALIDO:
			case CHARACTER:
			case PARIZQ:
			case NOT:
			case SUB:
				enterOuterAlt(_localctx, 2);
				{
				setState(763);
				((ListaexpresionContext)_localctx).expr = expr(0);

				    _localctx.funpar = datoscompuestos.NewArregloParametro(((ListaexpresionContext)_localctx).expr.e)

				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SimplematrizContext extends ParserRuleContext {
		public interfaces.Instruction simmat;
		public TipomatrizContext tipomatriz;
		public SimplematrizContext op;
		public Token NUMBER;
		public ExprContext expr;
		public TipomatrizContext tipomatriz() {
			return getRuleContext(TipomatrizContext.class,0);
		}
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public TerminalNode REPEATING() { return getToken(SwiftGrammarParser.REPEATING, 0); }
		public List<TerminalNode> DOS_PUNTOS() { return getTokens(SwiftGrammarParser.DOS_PUNTOS); }
		public TerminalNode DOS_PUNTOS(int i) {
			return getToken(SwiftGrammarParser.DOS_PUNTOS, i);
		}
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public TerminalNode COUNT() { return getToken(SwiftGrammarParser.COUNT, 0); }
		public TerminalNode NUMBER() { return getToken(SwiftGrammarParser.NUMBER, 0); }
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public SimplematrizContext simplematriz() {
			return getRuleContext(SimplematrizContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public SimplematrizContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_simplematriz; }
	}

	public final SimplematrizContext simplematriz() throws RecognitionException {
		SimplematrizContext _localctx = new SimplematrizContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_simplematriz);
		try {
			setState(792);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,46,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(768);
				((SimplematrizContext)_localctx).tipomatriz = tipomatriz();
				setState(769);
				match(PARIZQ);
				setState(770);
				match(REPEATING);
				setState(771);
				match(DOS_PUNTOS);
				setState(772);
				((SimplematrizContext)_localctx).op = simplematriz();
				setState(773);
				match(COMA);
				setState(774);
				match(COUNT);
				setState(775);
				match(DOS_PUNTOS);
				setState(776);
				((SimplematrizContext)_localctx).NUMBER = match(NUMBER);
				setState(777);
				match(PARDER);
				 _localctx.simmat = datoscompuestos.NewMatrizSimpleUno(((SimplematrizContext)_localctx).tipomatriz.tipomat, ((SimplematrizContext)_localctx).op.simmat, (((SimplematrizContext)_localctx).NUMBER!=null?((SimplematrizContext)_localctx).NUMBER.getText():null), (((SimplematrizContext)_localctx).NUMBER!=null?((SimplematrizContext)_localctx).NUMBER.getLine():0),(((SimplematrizContext)_localctx).NUMBER!=null?((SimplematrizContext)_localctx).NUMBER.getCharPositionInLine():0))
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(780);
				((SimplematrizContext)_localctx).tipomatriz = tipomatriz();
				setState(781);
				match(PARIZQ);
				setState(782);
				match(REPEATING);
				setState(783);
				match(DOS_PUNTOS);
				setState(784);
				((SimplematrizContext)_localctx).expr = expr(0);
				setState(785);
				match(COMA);
				setState(786);
				match(COUNT);
				setState(787);
				match(DOS_PUNTOS);
				setState(788);
				((SimplematrizContext)_localctx).NUMBER = match(NUMBER);
				setState(789);
				match(PARDER);
				 _localctx.simmat = datoscompuestos.NewMatrizSimpleDos(((SimplematrizContext)_localctx).tipomatriz.tipomat, ((SimplematrizContext)_localctx).expr.e, (((SimplematrizContext)_localctx).NUMBER!=null?((SimplematrizContext)_localctx).NUMBER.getText():null), (((SimplematrizContext)_localctx).NUMBER!=null?((SimplematrizContext)_localctx).NUMBER.getLine():0),(((SimplematrizContext)_localctx).NUMBER!=null?((SimplematrizContext)_localctx).NUMBER.getCharPositionInLine():0))
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ListamatrizaddsubsContext extends ParserRuleContext {
		public []interface{} blklimatas;
		public ListamatrizaddsubContext listamatrizaddsub;
		public List<ListamatrizaddsubContext> lmas = new ArrayList<ListamatrizaddsubContext>();
		public List<ListamatrizaddsubContext> listamatrizaddsub() {
			return getRuleContexts(ListamatrizaddsubContext.class);
		}
		public ListamatrizaddsubContext listamatrizaddsub(int i) {
			return getRuleContext(ListamatrizaddsubContext.class,i);
		}
		public ListamatrizaddsubsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listamatrizaddsubs; }
	}

	public final ListamatrizaddsubsContext listamatrizaddsubs() throws RecognitionException {
		ListamatrizaddsubsContext _localctx = new ListamatrizaddsubsContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_listamatrizaddsubs);

		    _localctx.blklimatas = []interface{}{}
		    var listInt []IListamatrizaddsubContext

		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(795); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(794);
					((ListamatrizaddsubsContext)_localctx).listamatrizaddsub = listamatrizaddsub();
					((ListamatrizaddsubsContext)_localctx).lmas.add(((ListamatrizaddsubsContext)_localctx).listamatrizaddsub);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(797); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,47,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );

			    listInt = localctx.(*ListamatrizaddsubsContext).GetLmas()
			    for _, e := range listInt {
			        _localctx.blklimatas = append(_localctx.blklimatas, e.GetLmas())
			    }

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ListamatrizaddsubContext extends ParserRuleContext {
		public interfaces.Expression lmas;
		public Token CORCHIZQ;
		public ExprContext expr;
		public TerminalNode CORCHIZQ() { return getToken(SwiftGrammarParser.CORCHIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode CORCHDER() { return getToken(SwiftGrammarParser.CORCHDER, 0); }
		public ListamatrizaddsubContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listamatrizaddsub; }
	}

	public final ListamatrizaddsubContext listamatrizaddsub() throws RecognitionException {
		ListamatrizaddsubContext _localctx = new ListamatrizaddsubContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_listamatrizaddsub);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(801);
			((ListamatrizaddsubContext)_localctx).CORCHIZQ = match(CORCHIZQ);
			setState(802);
			((ListamatrizaddsubContext)_localctx).expr = expr(0);
			setState(803);
			match(CORCHDER);

			    _localctx.lmas = datoscompuestos.NewArregloParametros((((ListamatrizaddsubContext)_localctx).CORCHIZQ!=null?((ListamatrizaddsubContext)_localctx).CORCHIZQ.getLine():0) ,(((ListamatrizaddsubContext)_localctx).CORCHIZQ!=null?((ListamatrizaddsubContext)_localctx).CORCHIZQ.getCharPositionInLine():0), ((ListamatrizaddsubContext)_localctx).expr.e)

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FunciondeclaracioncontrolContext extends ParserRuleContext {
		public interfaces.Instruction fdc;
		public Token ID_VALIDO;
		public ListaparametroContext listaparametro;
		public TipodatoContext tipodato;
		public BlockinternoContext blockinterno;
		public TerminalNode FUNCION() { return getToken(SwiftGrammarParser.FUNCION, 0); }
		public TerminalNode ID_VALIDO() { return getToken(SwiftGrammarParser.ID_VALIDO, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ListaparametroContext listaparametro() {
			return getRuleContext(ListaparametroContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public TerminalNode RETORNO() { return getToken(SwiftGrammarParser.RETORNO, 0); }
		public TipodatoContext tipodato() {
			return getRuleContext(TipodatoContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockinternoContext blockinterno() {
			return getRuleContext(BlockinternoContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public FunciondeclaracioncontrolContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funciondeclaracioncontrol; }
	}

	public final FunciondeclaracioncontrolContext funciondeclaracioncontrol() throws RecognitionException {
		FunciondeclaracioncontrolContext _localctx = new FunciondeclaracioncontrolContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_funciondeclaracioncontrol);
		try {
			setState(848);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,48,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(806);
				match(FUNCION);
				setState(807);
				((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(808);
				match(PARIZQ);
				setState(809);
				((FunciondeclaracioncontrolContext)_localctx).listaparametro = listaparametro();
				setState(810);
				match(PARDER);
				setState(811);
				match(RETORNO);
				setState(812);
				((FunciondeclaracioncontrolContext)_localctx).tipodato = tipodato();
				setState(813);
				match(LLAVEIZQ);
				setState(814);
				((FunciondeclaracioncontrolContext)_localctx).blockinterno = blockinterno();
				setState(815);
				match(LLAVEDER);

				    _localctx.fdc = funciones.NewFuncionesDeclaracionRP((((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO!=null?((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO.getLine():0), (((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO!=null?((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO.getCharPositionInLine():0), (((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO!=null?((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO.getText():null), ((FunciondeclaracioncontrolContext)_localctx).listaparametro.listparfun, ((FunciondeclaracioncontrolContext)_localctx).tipodato.tipo, ((FunciondeclaracioncontrolContext)_localctx).blockinterno.blkint)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(818);
				match(FUNCION);
				setState(819);
				((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(820);
				match(PARIZQ);
				setState(821);
				match(PARDER);
				setState(822);
				match(RETORNO);
				setState(823);
				((FunciondeclaracioncontrolContext)_localctx).tipodato = tipodato();
				setState(824);
				match(LLAVEIZQ);
				setState(825);
				((FunciondeclaracioncontrolContext)_localctx).blockinterno = blockinterno();
				setState(826);
				match(LLAVEDER);

				    _localctx.fdc = funciones.NewFuncionesDeclaracionR((((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO!=null?((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO.getLine():0), (((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO!=null?((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO.getCharPositionInLine():0), (((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO!=null?((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO.getText():null), ((FunciondeclaracioncontrolContext)_localctx).tipodato.tipo, ((FunciondeclaracioncontrolContext)_localctx).blockinterno.blkint)

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(829);
				match(FUNCION);
				setState(830);
				((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(831);
				match(PARIZQ);
				setState(832);
				((FunciondeclaracioncontrolContext)_localctx).listaparametro = listaparametro();
				setState(833);
				match(PARDER);
				setState(834);
				match(LLAVEIZQ);
				setState(835);
				((FunciondeclaracioncontrolContext)_localctx).blockinterno = blockinterno();
				setState(836);
				match(LLAVEDER);

				   _localctx.fdc = funciones.NewFuncionesDeclaracionP((((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO!=null?((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO.getLine():0), (((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO!=null?((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO.getCharPositionInLine():0), (((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO!=null?((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO.getText():null), ((FunciondeclaracioncontrolContext)_localctx).listaparametro.listparfun, ((FunciondeclaracioncontrolContext)_localctx).blockinterno.blkint)

				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(839);
				match(FUNCION);
				setState(840);
				((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(841);
				match(PARIZQ);
				setState(842);
				match(PARDER);
				setState(843);
				match(LLAVEIZQ);
				setState(844);
				((FunciondeclaracioncontrolContext)_localctx).blockinterno = blockinterno();
				setState(845);
				match(LLAVEDER);

				    _localctx.fdc = funciones.NewFuncionesDeclaracion((((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO!=null?((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO.getLine():0), (((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO!=null?((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO.getCharPositionInLine():0), (((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO!=null?((FunciondeclaracioncontrolContext)_localctx).ID_VALIDO.getText():null), ((FunciondeclaracioncontrolContext)_localctx).blockinterno.blkint)

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ListaparametroContext extends ParserRuleContext {
		public interfaces.Instruction listparfun;
		public Token op;
		public Token op2;
		public Token INOUT;
		public TipodatoContext tipodato;
		public ListaparametroContext op3;
		public TerminalNode DOS_PUNTOS() { return getToken(SwiftGrammarParser.DOS_PUNTOS, 0); }
		public TipodatoContext tipodato() {
			return getRuleContext(TipodatoContext.class,0);
		}
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public List<TerminalNode> ID_VALIDO() { return getTokens(SwiftGrammarParser.ID_VALIDO); }
		public TerminalNode ID_VALIDO(int i) {
			return getToken(SwiftGrammarParser.ID_VALIDO, i);
		}
		public ListaparametroContext listaparametro() {
			return getRuleContext(ListaparametroContext.class,0);
		}
		public TerminalNode INOUT() { return getToken(SwiftGrammarParser.INOUT, 0); }
		public TerminalNode GUIONBAJO() { return getToken(SwiftGrammarParser.GUIONBAJO, 0); }
		public ListaparametroContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listaparametro; }
	}

	public final ListaparametroContext listaparametro() throws RecognitionException {
		ListaparametroContext _localctx = new ListaparametroContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_listaparametro);
		int _la;
		try {
			setState(874);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,53,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(851);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,49,_ctx) ) {
				case 1:
					{
					setState(850);
					((ListaparametroContext)_localctx).op = _input.LT(1);
					_la = _input.LA(1);
					if ( !(_la==ID_VALIDO || _la==GUIONBAJO) ) {
						((ListaparametroContext)_localctx).op = (Token)_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
					break;
				}
				setState(853);
				((ListaparametroContext)_localctx).op2 = match(ID_VALIDO);
				setState(854);
				match(DOS_PUNTOS);
				setState(856);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==INOUT) {
					{
					setState(855);
					((ListaparametroContext)_localctx).INOUT = match(INOUT);
					}
				}

				setState(858);
				((ListaparametroContext)_localctx).tipodato = tipodato();
				setState(859);
				match(COMA);
				setState(860);
				((ListaparametroContext)_localctx).op3 = listaparametro();

				    if ((ListaparametroContext)_localctx).op != nil{
				        if ((ListaparametroContext)_localctx).INOUT != nil{
				            _localctx.listparfun = funciones.NewFuncionesListaParametro((((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getLine():0), (((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getCharPositionInLine():0), (((ListaparametroContext)_localctx).op!=null?((ListaparametroContext)_localctx).op.getText():null), (((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getText():null), ((ListaparametroContext)_localctx).tipodato.tipo, true, true, ((ListaparametroContext)_localctx).op3.listparfun )
				        }else {
				            _localctx.listparfun = funciones.NewFuncionesListaParametro((((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getLine():0), (((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getCharPositionInLine():0), (((ListaparametroContext)_localctx).op!=null?((ListaparametroContext)_localctx).op.getText():null), (((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getText():null), ((ListaparametroContext)_localctx).tipodato.tipo, false, true, ((ListaparametroContext)_localctx).op3.listparfun )
				        } 
				    }else{
				        if ((ListaparametroContext)_localctx).INOUT != nil{
				            _localctx.listparfun = funciones.NewFuncionesListaParametro((((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getLine():0), (((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getCharPositionInLine():0), "", (((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getText():null), ((ListaparametroContext)_localctx).tipodato.tipo, true, false, ((ListaparametroContext)_localctx).op3.listparfun )
				        }else {
				            _localctx.listparfun = funciones.NewFuncionesListaParametro((((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getLine():0), (((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getCharPositionInLine():0), "", (((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getText():null), ((ListaparametroContext)_localctx).tipodato.tipo, false, false,((ListaparametroContext)_localctx).op3.listparfun )
				        } 
				    }      

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(864);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,51,_ctx) ) {
				case 1:
					{
					setState(863);
					((ListaparametroContext)_localctx).op = _input.LT(1);
					_la = _input.LA(1);
					if ( !(_la==ID_VALIDO || _la==GUIONBAJO) ) {
						((ListaparametroContext)_localctx).op = (Token)_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
					break;
				}
				setState(866);
				((ListaparametroContext)_localctx).op2 = match(ID_VALIDO);
				setState(867);
				match(DOS_PUNTOS);
				setState(869);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==INOUT) {
					{
					setState(868);
					((ListaparametroContext)_localctx).INOUT = match(INOUT);
					}
				}

				setState(871);
				((ListaparametroContext)_localctx).tipodato = tipodato();

				    if ((ListaparametroContext)_localctx).op != nil{
				        if ((ListaparametroContext)_localctx).INOUT != nil{
				            _localctx.listparfun = funciones.NewFuncionesParametro((((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getLine():0), (((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getCharPositionInLine():0), (((ListaparametroContext)_localctx).op!=null?((ListaparametroContext)_localctx).op.getText():null), (((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getText():null), ((ListaparametroContext)_localctx).tipodato.tipo, true , true)
				        }else {
				            _localctx.listparfun = funciones.NewFuncionesParametro((((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getLine():0), (((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getCharPositionInLine():0), (((ListaparametroContext)_localctx).op!=null?((ListaparametroContext)_localctx).op.getText():null), (((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getText():null), ((ListaparametroContext)_localctx).tipodato.tipo, false, true)
				        } 
				    }else{
				        if ((ListaparametroContext)_localctx).INOUT != nil{
				            _localctx.listparfun = funciones.NewFuncionesParametro((((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getLine():0), (((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getCharPositionInLine():0), "", (((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getText():null), ((ListaparametroContext)_localctx).tipodato.tipo, true, false)
				        }else {
				            _localctx.listparfun = funciones.NewFuncionesParametro((((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getLine():0), (((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getCharPositionInLine():0), "", (((ListaparametroContext)_localctx).op2!=null?((ListaparametroContext)_localctx).op2.getText():null), ((ListaparametroContext)_localctx).tipodato.tipo, false, false)
				    } 
				    }
				    

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FuncionllamadacontrolContext extends ParserRuleContext {
		public interfaces.Instruction flctl;
		public Token ID_VALIDO;
		public ListaparametrosllamadaContext listaparametrosllamada;
		public TerminalNode ID_VALIDO() { return getToken(SwiftGrammarParser.ID_VALIDO, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ListaparametrosllamadaContext listaparametrosllamada() {
			return getRuleContext(ListaparametrosllamadaContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public FuncionllamadacontrolContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcionllamadacontrol; }
	}

	public final FuncionllamadacontrolContext funcionllamadacontrol() throws RecognitionException {
		FuncionllamadacontrolContext _localctx = new FuncionllamadacontrolContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_funcionllamadacontrol);
		try {
			setState(886);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,54,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(876);
				((FuncionllamadacontrolContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(877);
				match(PARIZQ);
				setState(878);
				((FuncionllamadacontrolContext)_localctx).listaparametrosllamada = listaparametrosllamada();
				setState(879);
				match(PARDER);

				    _localctx.flctl = funciones.NewFuncionesControlP((((FuncionllamadacontrolContext)_localctx).ID_VALIDO!=null?((FuncionllamadacontrolContext)_localctx).ID_VALIDO.getLine():0), (((FuncionllamadacontrolContext)_localctx).ID_VALIDO!=null?((FuncionllamadacontrolContext)_localctx).ID_VALIDO.getCharPositionInLine():0), (((FuncionllamadacontrolContext)_localctx).ID_VALIDO!=null?((FuncionllamadacontrolContext)_localctx).ID_VALIDO.getText():null), ((FuncionllamadacontrolContext)_localctx).listaparametrosllamada.lpll)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(882);
				((FuncionllamadacontrolContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(883);
				match(PARIZQ);
				setState(884);
				match(PARDER);

				    _localctx.flctl = funciones.NewFuncionesControl((((FuncionllamadacontrolContext)_localctx).ID_VALIDO!=null?((FuncionllamadacontrolContext)_localctx).ID_VALIDO.getLine():0), (((FuncionllamadacontrolContext)_localctx).ID_VALIDO!=null?((FuncionllamadacontrolContext)_localctx).ID_VALIDO.getCharPositionInLine():0), (((FuncionllamadacontrolContext)_localctx).ID_VALIDO!=null?((FuncionllamadacontrolContext)_localctx).ID_VALIDO.getText():null) )

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FuncionllamadacontrolConRetornoContext extends ParserRuleContext {
		public interfaces.Expression flctlret;
		public Token ID_VALIDO;
		public ListaparametrosllamadaContext listaparametrosllamada;
		public TerminalNode ID_VALIDO() { return getToken(SwiftGrammarParser.ID_VALIDO, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ListaparametrosllamadaContext listaparametrosllamada() {
			return getRuleContext(ListaparametrosllamadaContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public FuncionllamadacontrolConRetornoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcionllamadacontrolConRetorno; }
	}

	public final FuncionllamadacontrolConRetornoContext funcionllamadacontrolConRetorno() throws RecognitionException {
		FuncionllamadacontrolConRetornoContext _localctx = new FuncionllamadacontrolConRetornoContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_funcionllamadacontrolConRetorno);
		try {
			setState(898);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,55,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(888);
				((FuncionllamadacontrolConRetornoContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(889);
				match(PARIZQ);
				setState(890);
				((FuncionllamadacontrolConRetornoContext)_localctx).listaparametrosllamada = listaparametrosllamada();
				setState(891);
				match(PARDER);

				    _localctx.flctlret = funciones.NewFuncionesControlPR((((FuncionllamadacontrolConRetornoContext)_localctx).ID_VALIDO!=null?((FuncionllamadacontrolConRetornoContext)_localctx).ID_VALIDO.getLine():0), (((FuncionllamadacontrolConRetornoContext)_localctx).ID_VALIDO!=null?((FuncionllamadacontrolConRetornoContext)_localctx).ID_VALIDO.getCharPositionInLine():0), (((FuncionllamadacontrolConRetornoContext)_localctx).ID_VALIDO!=null?((FuncionllamadacontrolConRetornoContext)_localctx).ID_VALIDO.getText():null), ((FuncionllamadacontrolConRetornoContext)_localctx).listaparametrosllamada.lpll)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(894);
				((FuncionllamadacontrolConRetornoContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(895);
				match(PARIZQ);
				setState(896);
				match(PARDER);

				    _localctx.flctlret = funciones.NewFuncionesControlR((((FuncionllamadacontrolConRetornoContext)_localctx).ID_VALIDO!=null?((FuncionllamadacontrolConRetornoContext)_localctx).ID_VALIDO.getLine():0), (((FuncionllamadacontrolConRetornoContext)_localctx).ID_VALIDO!=null?((FuncionllamadacontrolConRetornoContext)_localctx).ID_VALIDO.getCharPositionInLine():0), (((FuncionllamadacontrolConRetornoContext)_localctx).ID_VALIDO!=null?((FuncionllamadacontrolConRetornoContext)_localctx).ID_VALIDO.getText():null) )

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ListaparametrosllamadaContext extends ParserRuleContext {
		public interfaces.Instruction lpll;
		public Token DIRME;
		public Token ID_VALIDO;
		public ListaparametrosllamadaContext op2;
		public Token op;
		public ExprContext expr;
		public Token COMA;
		public TerminalNode DIRME() { return getToken(SwiftGrammarParser.DIRME, 0); }
		public TerminalNode ID_VALIDO() { return getToken(SwiftGrammarParser.ID_VALIDO, 0); }
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public ListaparametrosllamadaContext listaparametrosllamada() {
			return getRuleContext(ListaparametrosllamadaContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode DOS_PUNTOS() { return getToken(SwiftGrammarParser.DOS_PUNTOS, 0); }
		public ListaparametrosllamadaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listaparametrosllamada; }
	}

	public final ListaparametrosllamadaContext listaparametrosllamada() throws RecognitionException {
		ListaparametrosllamadaContext _localctx = new ListaparametrosllamadaContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_listaparametrosllamada);
		try {
			setState(925);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,58,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(900);
				((ListaparametrosllamadaContext)_localctx).DIRME = match(DIRME);
				setState(901);
				((ListaparametrosllamadaContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(902);
				match(COMA);
				setState(903);
				((ListaparametrosllamadaContext)_localctx).op2 = listaparametrosllamada();

				    _localctx.lpll = funciones.NewFuncionesLlamadaList1((((ListaparametrosllamadaContext)_localctx).DIRME!=null?((ListaparametrosllamadaContext)_localctx).DIRME.getLine():0), (((ListaparametrosllamadaContext)_localctx).DIRME!=null?((ListaparametrosllamadaContext)_localctx).DIRME.getCharPositionInLine():0), (((ListaparametrosllamadaContext)_localctx).ID_VALIDO!=null?((ListaparametrosllamadaContext)_localctx).ID_VALIDO.getText():null), ((ListaparametrosllamadaContext)_localctx).op2.lpll)    

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(906);
				((ListaparametrosllamadaContext)_localctx).DIRME = match(DIRME);
				setState(907);
				((ListaparametrosllamadaContext)_localctx).ID_VALIDO = match(ID_VALIDO);

				    _localctx.lpll = funciones.NewFuncionesLlamadaList2((((ListaparametrosllamadaContext)_localctx).DIRME!=null?((ListaparametrosllamadaContext)_localctx).DIRME.getLine():0), (((ListaparametrosllamadaContext)_localctx).DIRME!=null?((ListaparametrosllamadaContext)_localctx).DIRME.getCharPositionInLine():0), (((ListaparametrosllamadaContext)_localctx).ID_VALIDO!=null?((ListaparametrosllamadaContext)_localctx).ID_VALIDO.getText():null))    

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(911);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,56,_ctx) ) {
				case 1:
					{
					setState(909);
					((ListaparametrosllamadaContext)_localctx).ID_VALIDO = match(ID_VALIDO);
					setState(910);
					((ListaparametrosllamadaContext)_localctx).op = match(DOS_PUNTOS);
					}
					break;
				}
				setState(913);
				((ListaparametrosllamadaContext)_localctx).expr = expr(0);
				setState(914);
				((ListaparametrosllamadaContext)_localctx).COMA = match(COMA);
				setState(915);
				((ListaparametrosllamadaContext)_localctx).op2 = listaparametrosllamada();

				    if ((ListaparametrosllamadaContext)_localctx).op != nil{
				        _localctx.lpll = funciones.NewFuncionesLlamadaList3((((ListaparametrosllamadaContext)_localctx).ID_VALIDO!=null?((ListaparametrosllamadaContext)_localctx).ID_VALIDO.getLine():0), (((ListaparametrosllamadaContext)_localctx).ID_VALIDO!=null?((ListaparametrosllamadaContext)_localctx).ID_VALIDO.getCharPositionInLine():0), (((ListaparametrosllamadaContext)_localctx).ID_VALIDO!=null?((ListaparametrosllamadaContext)_localctx).ID_VALIDO.getText():null), ((ListaparametrosllamadaContext)_localctx).expr.e, ((ListaparametrosllamadaContext)_localctx).op2.lpll)
				    }else{
				        _localctx.lpll = funciones.NewFuncionesLlamadaList4((((ListaparametrosllamadaContext)_localctx).COMA!=null?((ListaparametrosllamadaContext)_localctx).COMA.getLine():0), (((ListaparametrosllamadaContext)_localctx).COMA!=null?((ListaparametrosllamadaContext)_localctx).COMA.getCharPositionInLine():0), ((ListaparametrosllamadaContext)_localctx).expr.e, ((ListaparametrosllamadaContext)_localctx).op2.lpll)
				    }

				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(920);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,57,_ctx) ) {
				case 1:
					{
					setState(918);
					((ListaparametrosllamadaContext)_localctx).ID_VALIDO = match(ID_VALIDO);
					setState(919);
					((ListaparametrosllamadaContext)_localctx).op = match(DOS_PUNTOS);
					}
					break;
				}
				setState(922);
				((ListaparametrosllamadaContext)_localctx).expr = expr(0);

				    if ((ListaparametrosllamadaContext)_localctx).op != nil{
				        _localctx.lpll = funciones.NewFuncionesLlamadaList5((((ListaparametrosllamadaContext)_localctx).ID_VALIDO!=null?((ListaparametrosllamadaContext)_localctx).ID_VALIDO.getLine():0), (((ListaparametrosllamadaContext)_localctx).ID_VALIDO!=null?((ListaparametrosllamadaContext)_localctx).ID_VALIDO.getCharPositionInLine():0), (((ListaparametrosllamadaContext)_localctx).ID_VALIDO!=null?((ListaparametrosllamadaContext)_localctx).ID_VALIDO.getText():null), ((ListaparametrosllamadaContext)_localctx).expr.e)
				    }else{
				        _localctx.lpll = funciones.NewFuncionesLlamadaList6(((ListaparametrosllamadaContext)_localctx).expr.e)
				    }     

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PrintstmtContext extends ParserRuleContext {
		public interfaces.Instruction prnt;
		public Token PRINT;
		public ListaexpresionsContext listaexpresions;
		public TerminalNode PRINT() { return getToken(SwiftGrammarParser.PRINT, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ListaexpresionsContext listaexpresions() {
			return getRuleContext(ListaexpresionsContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public PrintstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printstmt; }
	}

	public final PrintstmtContext printstmt() throws RecognitionException {
		PrintstmtContext _localctx = new PrintstmtContext(_ctx, getState());
		enterRule(_localctx, 86, RULE_printstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(927);
			((PrintstmtContext)_localctx).PRINT = match(PRINT);
			setState(928);
			match(PARIZQ);
			setState(929);
			((PrintstmtContext)_localctx).listaexpresions = listaexpresions();
			setState(930);
			match(PARDER);
			 _localctx.prnt = funciones.NewPrint((((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getLine():0),(((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getCharPositionInLine():0),((PrintstmtContext)_localctx).listaexpresions.blkparf)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IntembebidaContext extends ParserRuleContext {
		public interfaces.Expression intemb;
		public ExprContext expr;
		public TerminalNode INT() { return getToken(SwiftGrammarParser.INT, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public IntembebidaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_intembebida; }
	}

	public final IntembebidaContext intembebida() throws RecognitionException {
		IntembebidaContext _localctx = new IntembebidaContext(_ctx, getState());
		enterRule(_localctx, 88, RULE_intembebida);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(933);
			match(INT);
			setState(934);
			match(PARIZQ);
			setState(935);
			((IntembebidaContext)_localctx).expr = expr(0);
			setState(936);
			match(PARDER);
			 _localctx.intemb = funciones.NewFuncionIntEmbebida(((IntembebidaContext)_localctx).expr.e)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FloatembebidaContext extends ParserRuleContext {
		public interfaces.Expression floemb;
		public ExprContext expr;
		public TerminalNode FLOAT() { return getToken(SwiftGrammarParser.FLOAT, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public FloatembebidaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_floatembebida; }
	}

	public final FloatembebidaContext floatembebida() throws RecognitionException {
		FloatembebidaContext _localctx = new FloatembebidaContext(_ctx, getState());
		enterRule(_localctx, 90, RULE_floatembebida);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(939);
			match(FLOAT);
			setState(940);
			match(PARIZQ);
			setState(941);
			((FloatembebidaContext)_localctx).expr = expr(0);
			setState(942);
			match(PARDER);
			 _localctx.floemb = funciones.NewFuncionFloatEmbebida(((FloatembebidaContext)_localctx).expr.e)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StringembebidaContext extends ParserRuleContext {
		public interfaces.Expression stremb;
		public ExprContext expr;
		public TerminalNode STRING() { return getToken(SwiftGrammarParser.STRING, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public StringembebidaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stringembebida; }
	}

	public final StringembebidaContext stringembebida() throws RecognitionException {
		StringembebidaContext _localctx = new StringembebidaContext(_ctx, getState());
		enterRule(_localctx, 92, RULE_stringembebida);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(945);
			match(STRING);
			setState(946);
			match(PARIZQ);
			setState(947);
			((StringembebidaContext)_localctx).expr = expr(0);
			setState(948);
			match(PARDER);
			 _localctx.stremb = funciones.NewFuncionStringEmbebida(((StringembebidaContext)_localctx).expr.e)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 9:
			return expr_sempred((ExprContext)_localctx, predIndex);
		case 32:
			return listavaloresmat2_sempred((Listavaloresmat2Context)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 24);
		case 1:
			return precpred(_ctx, 23);
		case 2:
			return precpred(_ctx, 22);
		case 3:
			return precpred(_ctx, 21);
		case 4:
			return precpred(_ctx, 20);
		case 5:
			return precpred(_ctx, 19);
		case 6:
			return precpred(_ctx, 18);
		case 7:
			return precpred(_ctx, 17);
		}
		return true;
	}
	private boolean listavaloresmat2_sempred(Listavaloresmat2Context _localctx, int predIndex) {
		switch (predIndex) {
		case 8:
			return precpred(_ctx, 3);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001K\u03b8\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0002"+
		"#\u0007#\u0002$\u0007$\u0002%\u0007%\u0002&\u0007&\u0002\'\u0007\'\u0002"+
		"(\u0007(\u0002)\u0007)\u0002*\u0007*\u0002+\u0007+\u0002,\u0007,\u0002"+
		"-\u0007-\u0002.\u0007.\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0001\u0004\u0001d\b\u0001\u000b\u0001\f\u0001e\u0001\u0001\u0001"+
		"\u0001\u0001\u0002\u0001\u0002\u0003\u0002l\b\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0003\u0002r\b\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0003\u0002x\b\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003"+
		"\u0002\u008d\b\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003"+
		"\u0002\u0099\b\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003"+
		"\u0002\u009f\b\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002\u00a9\b\u0002\u0001"+
		"\u0003\u0004\u0003\u00ac\b\u0003\u000b\u0003\f\u0003\u00ad\u0001\u0003"+
		"\u0001\u0003\u0001\u0004\u0001\u0004\u0003\u0004\u00b4\b\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004\u00ba\b\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004\u00c0\b\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0003\u0004\u00d5\b\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0003\u0004\u00db\b\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0003\u0004\u00e1\b\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0003\u0004\u00e7\b\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0003\u0004\u00ed\b\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0003\u0004\u00f3\b\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0003\u0004\u00f9\b\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0003\u0004\u00ff\b\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0003\u0004\u0106\b\u0004\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0003\u0005\u011d\b\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006\u012d\b\u0006"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007\u013e\b\u0007\u0001\b"+
		"\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0003\b\u014a\b\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0003\t\u017c\b\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0005\t\u01a6\b\t\n\t\f\t\u01a9\t\t\u0001\n\u0001\n"+
		"\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0003\n\u01c6"+
		"\b\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0003\u000b\u01cf\b\u000b\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\f\u0004\f\u01d5\b\f\u000b\f\f\f\u01d6\u0001\f\u0001\f\u0001"+
		"\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u0205\b\u000f\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0003\u0013\u021b\b\u0013\u0001\u0014\u0001\u0014\u0001"+
		"\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001"+
		"\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001"+
		"\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001"+
		"\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001"+
		"\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001"+
		"\u0014\u0003\u0014\u023e\b\u0014\u0001\u0015\u0004\u0015\u0241\b\u0015"+
		"\u000b\u0015\f\u0015\u0242\u0001\u0015\u0001\u0015\u0001\u0016\u0001\u0016"+
		"\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0003\u0016"+
		"\u024e\b\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0003\u0017\u0282\b\u0017\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0003\u0018\u0294\b\u0018\u0001\u0019"+
		"\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u001a\u0001\u001a"+
		"\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001b\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0003\u001b\u02b9\b\u001b"+
		"\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0003\u001c\u02bf\b\u001c"+
		"\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001d\u0001\u001d"+
		"\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d"+
		"\u0001\u001d\u0001\u001d\u0003\u001d\u02cf\b\u001d\u0001\u001e\u0001\u001e"+
		"\u0001\u001e\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f"+
		"\u0001\u001f\u0001\u001f\u0001\u001f\u0003\u001f\u02dc\b\u001f\u0001 "+
		"\u0001 \u0001 \u0001 \u0001 \u0001 \u0001 \u0003 \u02e5\b \u0001 \u0001"+
		" \u0001 \u0001 \u0001 \u0005 \u02ec\b \n \f \u02ef\t \u0001!\u0004!\u02f2"+
		"\b!\u000b!\f!\u02f3\u0001!\u0001!\u0001\"\u0001\"\u0001\"\u0001\"\u0001"+
		"\"\u0001\"\u0001\"\u0003\"\u02ff\b\"\u0001#\u0001#\u0001#\u0001#\u0001"+
		"#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001"+
		"#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0003"+
		"#\u0319\b#\u0001$\u0004$\u031c\b$\u000b$\f$\u031d\u0001$\u0001$\u0001"+
		"%\u0001%\u0001%\u0001%\u0001%\u0001&\u0001&\u0001&\u0001&\u0001&\u0001"+
		"&\u0001&\u0001&\u0001&\u0001&\u0001&\u0001&\u0001&\u0001&\u0001&\u0001"+
		"&\u0001&\u0001&\u0001&\u0001&\u0001&\u0001&\u0001&\u0001&\u0001&\u0001"+
		"&\u0001&\u0001&\u0001&\u0001&\u0001&\u0001&\u0001&\u0001&\u0001&\u0001"+
		"&\u0001&\u0001&\u0001&\u0001&\u0001&\u0001&\u0003&\u0351\b&\u0001\'\u0003"+
		"\'\u0354\b\'\u0001\'\u0001\'\u0001\'\u0003\'\u0359\b\'\u0001\'\u0001\'"+
		"\u0001\'\u0001\'\u0001\'\u0001\'\u0003\'\u0361\b\'\u0001\'\u0001\'\u0001"+
		"\'\u0003\'\u0366\b\'\u0001\'\u0001\'\u0001\'\u0003\'\u036b\b\'\u0001("+
		"\u0001(\u0001(\u0001(\u0001(\u0001(\u0001(\u0001(\u0001(\u0001(\u0003"+
		"(\u0377\b(\u0001)\u0001)\u0001)\u0001)\u0001)\u0001)\u0001)\u0001)\u0001"+
		")\u0001)\u0003)\u0383\b)\u0001*\u0001*\u0001*\u0001*\u0001*\u0001*\u0001"+
		"*\u0001*\u0001*\u0001*\u0001*\u0003*\u0390\b*\u0001*\u0001*\u0001*\u0001"+
		"*\u0001*\u0001*\u0001*\u0003*\u0399\b*\u0001*\u0001*\u0001*\u0003*\u039e"+
		"\b*\u0001+\u0001+\u0001+\u0001+\u0001+\u0001+\u0001,\u0001,\u0001,\u0001"+
		",\u0001,\u0001,\u0001-\u0001-\u0001-\u0001-\u0001-\u0001-\u0001.\u0001"+
		".\u0001.\u0001.\u0001.\u0001.\u0001.\u0000\u0002\u0012@/\u0000\u0002\u0004"+
		"\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \""+
		"$&(*,.02468:<>@BDFHJLNPRTVXZ\\\u0000\u0006\u0001\u0000:;\u0001\u0000<"+
		"=\u0002\u00005577\u0002\u00006688\u0001\u000001\u0002\u0000\'\'EE\u0404"+
		"\u0000^\u0001\u0000\u0000\u0000\u0002c\u0001\u0000\u0000\u0000\u0004\u00a8"+
		"\u0001\u0000\u0000\u0000\u0006\u00ab\u0001\u0000\u0000\u0000\b\u0105\u0001"+
		"\u0000\u0000\u0000\n\u011c\u0001\u0000\u0000\u0000\f\u012c\u0001\u0000"+
		"\u0000\u0000\u000e\u013d\u0001\u0000\u0000\u0000\u0010\u0149\u0001\u0000"+
		"\u0000\u0000\u0012\u017b\u0001\u0000\u0000\u0000\u0014\u01c5\u0001\u0000"+
		"\u0000\u0000\u0016\u01c7\u0001\u0000\u0000\u0000\u0018\u01d4\u0001\u0000"+
		"\u0000\u0000\u001a\u01da\u0001\u0000\u0000\u0000\u001c\u01e0\u0001\u0000"+
		"\u0000\u0000\u001e\u0204\u0001\u0000\u0000\u0000 \u0206\u0001\u0000\u0000"+
		"\u0000\"\u020e\u0001\u0000\u0000\u0000$\u0211\u0001\u0000\u0000\u0000"+
		"&\u021a\u0001\u0000\u0000\u0000(\u023d\u0001\u0000\u0000\u0000*\u0240"+
		"\u0001\u0000\u0000\u0000,\u024d\u0001\u0000\u0000\u0000.\u0281\u0001\u0000"+
		"\u0000\u00000\u0293\u0001\u0000\u0000\u00002\u0295\u0001\u0000\u0000\u0000"+
		"4\u029a\u0001\u0000\u0000\u00006\u02b8\u0001\u0000\u0000\u00008\u02ba"+
		"\u0001\u0000\u0000\u0000:\u02ce\u0001\u0000\u0000\u0000<\u02d0\u0001\u0000"+
		"\u0000\u0000>\u02db\u0001\u0000\u0000\u0000@\u02e4\u0001\u0000\u0000\u0000"+
		"B\u02f1\u0001\u0000\u0000\u0000D\u02fe\u0001\u0000\u0000\u0000F\u0318"+
		"\u0001\u0000\u0000\u0000H\u031b\u0001\u0000\u0000\u0000J\u0321\u0001\u0000"+
		"\u0000\u0000L\u0350\u0001\u0000\u0000\u0000N\u036a\u0001\u0000\u0000\u0000"+
		"P\u0376\u0001\u0000\u0000\u0000R\u0382\u0001\u0000\u0000\u0000T\u039d"+
		"\u0001\u0000\u0000\u0000V\u039f\u0001\u0000\u0000\u0000X\u03a5\u0001\u0000"+
		"\u0000\u0000Z\u03ab\u0001\u0000\u0000\u0000\\\u03b1\u0001\u0000\u0000"+
		"\u0000^_\u0003\u0002\u0001\u0000_`\u0005\u0000\u0000\u0001`a\u0006\u0000"+
		"\uffff\uffff\u0000a\u0001\u0001\u0000\u0000\u0000bd\u0003\u0004\u0002"+
		"\u0000cb\u0001\u0000\u0000\u0000de\u0001\u0000\u0000\u0000ec\u0001\u0000"+
		"\u0000\u0000ef\u0001\u0000\u0000\u0000fg\u0001\u0000\u0000\u0000gh\u0006"+
		"\u0001\uffff\uffff\u0000h\u0003\u0001\u0000\u0000\u0000ik\u0003\n\u0005"+
		"\u0000jl\u0005,\u0000\u0000kj\u0001\u0000\u0000\u0000kl\u0001\u0000\u0000"+
		"\u0000lm\u0001\u0000\u0000\u0000mn\u0006\u0002\uffff\uffff\u0000n\u00a9"+
		"\u0001\u0000\u0000\u0000oq\u0003\f\u0006\u0000pr\u0005,\u0000\u0000qp"+
		"\u0001\u0000\u0000\u0000qr\u0001\u0000\u0000\u0000rs\u0001\u0000\u0000"+
		"\u0000st\u0006\u0002\uffff\uffff\u0000t\u00a9\u0001\u0000\u0000\u0000"+
		"uw\u0003\u000e\u0007\u0000vx\u0005,\u0000\u0000wv\u0001\u0000\u0000\u0000"+
		"wx\u0001\u0000\u0000\u0000xy\u0001\u0000\u0000\u0000yz\u0006\u0002\uffff"+
		"\uffff\u0000z\u00a9\u0001\u0000\u0000\u0000{|\u0003\u0014\n\u0000|}\u0006"+
		"\u0002\uffff\uffff\u0000}\u00a9\u0001\u0000\u0000\u0000~\u007f\u0003\u0016"+
		"\u000b\u0000\u007f\u0080\u0006\u0002\uffff\uffff\u0000\u0080\u00a9\u0001"+
		"\u0000\u0000\u0000\u0081\u0082\u0003\u001c\u000e\u0000\u0082\u0083\u0006"+
		"\u0002\uffff\uffff\u0000\u0083\u00a9\u0001\u0000\u0000\u0000\u0084\u0085"+
		"\u0003\u001e\u000f\u0000\u0085\u0086\u0006\u0002\uffff\uffff\u0000\u0086"+
		"\u00a9\u0001\u0000\u0000\u0000\u0087\u0088\u0003 \u0010\u0000\u0088\u0089"+
		"\u0006\u0002\uffff\uffff\u0000\u0089\u00a9\u0001\u0000\u0000\u0000\u008a"+
		"\u008c\u0003(\u0014\u0000\u008b\u008d\u0005,\u0000\u0000\u008c\u008b\u0001"+
		"\u0000\u0000\u0000\u008c\u008d\u0001\u0000\u0000\u0000\u008d\u008e\u0001"+
		"\u0000\u0000\u0000\u008e\u008f\u0006\u0002\uffff\uffff\u0000\u008f\u00a9"+
		"\u0001\u0000\u0000\u0000\u0090\u0091\u0003.\u0017\u0000\u0091\u0092\u0006"+
		"\u0002\uffff\uffff\u0000\u0092\u00a9\u0001\u0000\u0000\u0000\u0093\u0094"+
		"\u00030\u0018\u0000\u0094\u0095\u0006\u0002\uffff\uffff\u0000\u0095\u00a9"+
		"\u0001\u0000\u0000\u0000\u0096\u0098\u0003V+\u0000\u0097\u0099\u0005,"+
		"\u0000\u0000\u0098\u0097\u0001\u0000\u0000\u0000\u0098\u0099\u0001\u0000"+
		"\u0000\u0000\u0099\u009a\u0001\u0000\u0000\u0000\u009a\u009b\u0006\u0002"+
		"\uffff\uffff\u0000\u009b\u00a9\u0001\u0000\u0000\u0000\u009c\u009e\u0003"+
		"8\u001c\u0000\u009d\u009f\u0005,\u0000\u0000\u009e\u009d\u0001\u0000\u0000"+
		"\u0000\u009e\u009f\u0001\u0000\u0000\u0000\u009f\u00a0\u0001\u0000\u0000"+
		"\u0000\u00a0\u00a1\u0006\u0002\uffff\uffff\u0000\u00a1\u00a9\u0001\u0000"+
		"\u0000\u0000\u00a2\u00a3\u0003L&\u0000\u00a3\u00a4\u0006\u0002\uffff\uffff"+
		"\u0000\u00a4\u00a9\u0001\u0000\u0000\u0000\u00a5\u00a6\u0003P(\u0000\u00a6"+
		"\u00a7\u0006\u0002\uffff\uffff\u0000\u00a7\u00a9\u0001\u0000\u0000\u0000"+
		"\u00a8i\u0001\u0000\u0000\u0000\u00a8o\u0001\u0000\u0000\u0000\u00a8u"+
		"\u0001\u0000\u0000\u0000\u00a8{\u0001\u0000\u0000\u0000\u00a8~\u0001\u0000"+
		"\u0000\u0000\u00a8\u0081\u0001\u0000\u0000\u0000\u00a8\u0084\u0001\u0000"+
		"\u0000\u0000\u00a8\u0087\u0001\u0000\u0000\u0000\u00a8\u008a\u0001\u0000"+
		"\u0000\u0000\u00a8\u0090\u0001\u0000\u0000\u0000\u00a8\u0093\u0001\u0000"+
		"\u0000\u0000\u00a8\u0096\u0001\u0000\u0000\u0000\u00a8\u009c\u0001\u0000"+
		"\u0000\u0000\u00a8\u00a2\u0001\u0000\u0000\u0000\u00a8\u00a5\u0001\u0000"+
		"\u0000\u0000\u00a9\u0005\u0001\u0000\u0000\u0000\u00aa\u00ac\u0003\b\u0004"+
		"\u0000\u00ab\u00aa\u0001\u0000\u0000\u0000\u00ac\u00ad\u0001\u0000\u0000"+
		"\u0000\u00ad\u00ab\u0001\u0000\u0000\u0000\u00ad\u00ae\u0001\u0000\u0000"+
		"\u0000\u00ae\u00af\u0001\u0000\u0000\u0000\u00af\u00b0\u0006\u0003\uffff"+
		"\uffff\u0000\u00b0\u0007\u0001\u0000\u0000\u0000\u00b1\u00b3\u0003\n\u0005"+
		"\u0000\u00b2\u00b4\u0005,\u0000\u0000\u00b3\u00b2\u0001\u0000\u0000\u0000"+
		"\u00b3\u00b4\u0001\u0000\u0000\u0000\u00b4\u00b5\u0001\u0000\u0000\u0000"+
		"\u00b5\u00b6\u0006\u0004\uffff\uffff\u0000\u00b6\u0106\u0001\u0000\u0000"+
		"\u0000\u00b7\u00b9\u0003\f\u0006\u0000\u00b8\u00ba\u0005,\u0000\u0000"+
		"\u00b9\u00b8\u0001\u0000\u0000\u0000\u00b9\u00ba\u0001\u0000\u0000\u0000"+
		"\u00ba\u00bb\u0001\u0000\u0000\u0000\u00bb\u00bc\u0006\u0004\uffff\uffff"+
		"\u0000\u00bc\u0106\u0001\u0000\u0000\u0000\u00bd\u00bf\u0003\u000e\u0007"+
		"\u0000\u00be\u00c0\u0005,\u0000\u0000\u00bf\u00be\u0001\u0000\u0000\u0000"+
		"\u00bf\u00c0\u0001\u0000\u0000\u0000\u00c0\u00c1\u0001\u0000\u0000\u0000"+
		"\u00c1\u00c2\u0006\u0004\uffff\uffff\u0000\u00c2\u0106\u0001\u0000\u0000"+
		"\u0000\u00c3\u00c4\u0003\u0014\n\u0000\u00c4\u00c5\u0006\u0004\uffff\uffff"+
		"\u0000\u00c5\u0106\u0001\u0000\u0000\u0000\u00c6\u00c7\u0003\u0016\u000b"+
		"\u0000\u00c7\u00c8\u0006\u0004\uffff\uffff\u0000\u00c8\u0106\u0001\u0000"+
		"\u0000\u0000\u00c9\u00ca\u0003\u001c\u000e\u0000\u00ca\u00cb\u0006\u0004"+
		"\uffff\uffff\u0000\u00cb\u0106\u0001\u0000\u0000\u0000\u00cc\u00cd\u0003"+
		"\u001e\u000f\u0000\u00cd\u00ce\u0006\u0004\uffff\uffff\u0000\u00ce\u0106"+
		"\u0001\u0000\u0000\u0000\u00cf\u00d0\u0003 \u0010\u0000\u00d0\u00d1\u0006"+
		"\u0004\uffff\uffff\u0000\u00d1\u0106\u0001\u0000\u0000\u0000\u00d2\u00d4"+
		"\u0003\"\u0011\u0000\u00d3\u00d5\u0005,\u0000\u0000\u00d4\u00d3\u0001"+
		"\u0000\u0000\u0000\u00d4\u00d5\u0001\u0000\u0000\u0000\u00d5\u00d6\u0001"+
		"\u0000\u0000\u0000\u00d6\u00d7\u0006\u0004\uffff\uffff\u0000\u00d7\u0106"+
		"\u0001\u0000\u0000\u0000\u00d8\u00da\u0003$\u0012\u0000\u00d9\u00db\u0005"+
		",\u0000\u0000\u00da\u00d9\u0001\u0000\u0000\u0000\u00da\u00db\u0001\u0000"+
		"\u0000\u0000\u00db\u00dc\u0001\u0000\u0000\u0000\u00dc\u00dd\u0006\u0004"+
		"\uffff\uffff\u0000\u00dd\u0106\u0001\u0000\u0000\u0000\u00de\u00e0\u0003"+
		"&\u0013\u0000\u00df\u00e1\u0005,\u0000\u0000\u00e0\u00df\u0001\u0000\u0000"+
		"\u0000\u00e0\u00e1\u0001\u0000\u0000\u0000\u00e1\u00e2\u0001\u0000\u0000"+
		"\u0000\u00e2\u00e3\u0006\u0004\uffff\uffff\u0000\u00e3\u0106\u0001\u0000"+
		"\u0000\u0000\u00e4\u00e6\u0003(\u0014\u0000\u00e5\u00e7\u0005,\u0000\u0000"+
		"\u00e6\u00e5\u0001\u0000\u0000\u0000\u00e6\u00e7\u0001\u0000\u0000\u0000"+
		"\u00e7\u00e8\u0001\u0000\u0000\u0000\u00e8\u00e9\u0006\u0004\uffff\uffff"+
		"\u0000\u00e9\u0106\u0001\u0000\u0000\u0000\u00ea\u00ec\u0003.\u0017\u0000"+
		"\u00eb\u00ed\u0005,\u0000\u0000\u00ec\u00eb\u0001\u0000\u0000\u0000\u00ec"+
		"\u00ed\u0001\u0000\u0000\u0000\u00ed\u00ee\u0001\u0000\u0000\u0000\u00ee"+
		"\u00ef\u0006\u0004\uffff\uffff\u0000\u00ef\u0106\u0001\u0000\u0000\u0000"+
		"\u00f0\u00f2\u00030\u0018\u0000\u00f1\u00f3\u0005,\u0000\u0000\u00f2\u00f1"+
		"\u0001\u0000\u0000\u0000\u00f2\u00f3\u0001\u0000\u0000\u0000\u00f3\u00f4"+
		"\u0001\u0000\u0000\u0000\u00f4\u00f5\u0006\u0004\uffff\uffff\u0000\u00f5"+
		"\u0106\u0001\u0000\u0000\u0000\u00f6\u00f8\u0003V+\u0000\u00f7\u00f9\u0005"+
		",\u0000\u0000\u00f8\u00f7\u0001\u0000\u0000\u0000\u00f8\u00f9\u0001\u0000"+
		"\u0000\u0000\u00f9\u00fa\u0001\u0000\u0000\u0000\u00fa\u00fb\u0006\u0004"+
		"\uffff\uffff\u0000\u00fb\u0106\u0001\u0000\u0000\u0000\u00fc\u00fe\u0003"+
		"8\u001c\u0000\u00fd\u00ff\u0005,\u0000\u0000\u00fe\u00fd\u0001\u0000\u0000"+
		"\u0000\u00fe\u00ff\u0001\u0000\u0000\u0000\u00ff\u0100\u0001\u0000\u0000"+
		"\u0000\u0100\u0101\u0006\u0004\uffff\uffff\u0000\u0101\u0106\u0001\u0000"+
		"\u0000\u0000\u0102\u0103\u0003P(\u0000\u0103\u0104\u0006\u0004\uffff\uffff"+
		"\u0000\u0104\u0106\u0001\u0000\u0000\u0000\u0105\u00b1\u0001\u0000\u0000"+
		"\u0000\u0105\u00b7\u0001\u0000\u0000\u0000\u0105\u00bd\u0001\u0000\u0000"+
		"\u0000\u0105\u00c3\u0001\u0000\u0000\u0000\u0105\u00c6\u0001\u0000\u0000"+
		"\u0000\u0105\u00c9\u0001\u0000\u0000\u0000\u0105\u00cc\u0001\u0000\u0000"+
		"\u0000\u0105\u00cf\u0001\u0000\u0000\u0000\u0105\u00d2\u0001\u0000\u0000"+
		"\u0000\u0105\u00d8\u0001\u0000\u0000\u0000\u0105\u00de\u0001\u0000\u0000"+
		"\u0000\u0105\u00e4\u0001\u0000\u0000\u0000\u0105\u00ea\u0001\u0000\u0000"+
		"\u0000\u0105\u00f0\u0001\u0000\u0000\u0000\u0105\u00f6\u0001\u0000\u0000"+
		"\u0000\u0105\u00fc\u0001\u0000\u0000\u0000\u0105\u0102\u0001\u0000\u0000"+
		"\u0000\u0106\t\u0001\u0000\u0000\u0000\u0107\u0108\u0005\b\u0000\u0000"+
		"\u0108\u0109\u0005\'\u0000\u0000\u0109\u010a\u0005+\u0000\u0000\u010a"+
		"\u010b\u0003\u0010\b\u0000\u010b\u010c\u0005*\u0000\u0000\u010c\u010d"+
		"\u0003\u0012\t\u0000\u010d\u010e\u0006\u0005\uffff\uffff\u0000\u010e\u011d"+
		"\u0001\u0000\u0000\u0000\u010f\u0110\u0005\b\u0000\u0000\u0110\u0111\u0005"+
		"\'\u0000\u0000\u0111\u0112\u0005*\u0000\u0000\u0112\u0113\u0003\u0012"+
		"\t\u0000\u0113\u0114\u0006\u0005\uffff\uffff\u0000\u0114\u011d\u0001\u0000"+
		"\u0000\u0000\u0115\u0116\u0005\b\u0000\u0000\u0116\u0117\u0005\'\u0000"+
		"\u0000\u0117\u0118\u0005+\u0000\u0000\u0118\u0119\u0003\u0010\b\u0000"+
		"\u0119\u011a\u0005-\u0000\u0000\u011a\u011b\u0006\u0005\uffff\uffff\u0000"+
		"\u011b\u011d\u0001\u0000\u0000\u0000\u011c\u0107\u0001\u0000\u0000\u0000"+
		"\u011c\u010f\u0001\u0000\u0000\u0000\u011c\u0115\u0001\u0000\u0000\u0000"+
		"\u011d\u000b\u0001\u0000\u0000\u0000\u011e\u011f\u0005\t\u0000\u0000\u011f"+
		"\u0120\u0005\'\u0000\u0000\u0120\u0121\u0005+\u0000\u0000\u0121\u0122"+
		"\u0003\u0010\b\u0000\u0122\u0123\u0005*\u0000\u0000\u0123\u0124\u0003"+
		"\u0012\t\u0000\u0124\u0125\u0006\u0006\uffff\uffff\u0000\u0125\u012d\u0001"+
		"\u0000\u0000\u0000\u0126\u0127\u0005\t\u0000\u0000\u0127\u0128\u0005\'"+
		"\u0000\u0000\u0128\u0129\u0005*\u0000\u0000\u0129\u012a\u0003\u0012\t"+
		"\u0000\u012a\u012b\u0006\u0006\uffff\uffff\u0000\u012b\u012d\u0001\u0000"+
		"\u0000\u0000\u012c\u011e\u0001\u0000\u0000\u0000\u012c\u0126\u0001\u0000"+
		"\u0000\u0000\u012d\r\u0001\u0000\u0000\u0000\u012e\u012f\u0005\'\u0000"+
		"\u0000\u012f\u0130\u0005*\u0000\u0000\u0130\u0131\u0003\u0012\t\u0000"+
		"\u0131\u0132\u0006\u0007\uffff\uffff\u0000\u0132\u013e\u0001\u0000\u0000"+
		"\u0000\u0133\u0134\u0005\'\u0000\u0000\u0134\u0135\u0005>\u0000\u0000"+
		"\u0135\u0136\u0003\u0012\t\u0000\u0136\u0137\u0006\u0007\uffff\uffff\u0000"+
		"\u0137\u013e\u0001\u0000\u0000\u0000\u0138\u0139\u0005\'\u0000\u0000\u0139"+
		"\u013a\u0005?\u0000\u0000\u013a\u013b\u0003\u0012\t\u0000\u013b\u013c"+
		"\u0006\u0007\uffff\uffff\u0000\u013c\u013e\u0001\u0000\u0000\u0000\u013d"+
		"\u012e\u0001\u0000\u0000\u0000\u013d\u0133\u0001\u0000\u0000\u0000\u013d"+
		"\u0138\u0001\u0000\u0000\u0000\u013e\u000f\u0001\u0000\u0000\u0000\u013f"+
		"\u0140\u0005\u0001\u0000\u0000\u0140\u014a\u0006\b\uffff\uffff\u0000\u0141"+
		"\u0142\u0005\u0002\u0000\u0000\u0142\u014a\u0006\b\uffff\uffff\u0000\u0143"+
		"\u0144\u0005\u0003\u0000\u0000\u0144\u014a\u0006\b\uffff\uffff\u0000\u0145"+
		"\u0146\u0005\u0004\u0000\u0000\u0146\u014a\u0006\b\uffff\uffff\u0000\u0147"+
		"\u0148\u0005\u0005\u0000\u0000\u0148\u014a\u0006\b\uffff\uffff\u0000\u0149"+
		"\u013f\u0001\u0000\u0000\u0000\u0149\u0141\u0001\u0000\u0000\u0000\u0149"+
		"\u0143\u0001\u0000\u0000\u0000\u0149\u0145\u0001\u0000\u0000\u0000\u0149"+
		"\u0147\u0001\u0000\u0000\u0000\u014a\u0011\u0001\u0000\u0000\u0000\u014b"+
		"\u014c\u0006\t\uffff\uffff\u0000\u014c\u014d\u00052\u0000\u0000\u014d"+
		"\u014e\u0003\u0012\t\u0019\u014e\u014f\u0006\t\uffff\uffff\u0000\u014f"+
		"\u017c\u0001\u0000\u0000\u0000\u0150\u0151\u0005.\u0000\u0000\u0151\u0152"+
		"\u0003\u0012\t\u0000\u0152\u0153\u0005/\u0000\u0000\u0153\u0154\u0006"+
		"\t\uffff\uffff\u0000\u0154\u017c\u0001\u0000\u0000\u0000\u0155\u0156\u0005"+
		"=\u0000\u0000\u0156\u0157\u0005%\u0000\u0000\u0157\u017c\u0006\t\uffff"+
		"\uffff\u0000\u0158\u0159\u0005%\u0000\u0000\u0159\u017c\u0006\t\uffff"+
		"\uffff\u0000\u015a\u015b\u0005&\u0000\u0000\u015b\u017c\u0006\t\uffff"+
		"\uffff\u0000\u015c\u015d\u0005\u0006\u0000\u0000\u015d\u017c\u0006\t\uffff"+
		"\uffff\u0000\u015e\u015f\u0005\u0007\u0000\u0000\u015f\u017c\u0006\t\uffff"+
		"\uffff\u0000\u0160\u0161\u0005(\u0000\u0000\u0161\u017c\u0006\t\uffff"+
		"\uffff\u0000\u0162\u0163\u0005\'\u0000\u0000\u0163\u017c\u0006\t\uffff"+
		"\uffff\u0000\u0164\u0165\u0005\n\u0000\u0000\u0165\u017c\u0006\t\uffff"+
		"\uffff\u0000\u0166\u0167\u00032\u0019\u0000\u0167\u0168\u0006\t\uffff"+
		"\uffff\u0000\u0168\u017c\u0001\u0000\u0000\u0000\u0169\u016a\u00034\u001a"+
		"\u0000\u016a\u016b\u0006\t\uffff\uffff\u0000\u016b\u017c\u0001\u0000\u0000"+
		"\u0000\u016c\u016d\u00036\u001b\u0000\u016d\u016e\u0006\t\uffff\uffff"+
		"\u0000\u016e\u017c\u0001\u0000\u0000\u0000\u016f\u0170\u0003X,\u0000\u0170"+
		"\u0171\u0006\t\uffff\uffff\u0000\u0171\u017c\u0001\u0000\u0000\u0000\u0172"+
		"\u0173\u0003Z-\u0000\u0173\u0174\u0006\t\uffff\uffff\u0000\u0174\u017c"+
		"\u0001\u0000\u0000\u0000\u0175\u0176\u0003\\.\u0000\u0176\u0177\u0006"+
		"\t\uffff\uffff\u0000\u0177\u017c\u0001\u0000\u0000\u0000\u0178\u0179\u0003"+
		"R)\u0000\u0179\u017a\u0006\t\uffff\uffff\u0000\u017a\u017c\u0001\u0000"+
		"\u0000\u0000\u017b\u014b\u0001\u0000\u0000\u0000\u017b\u0150\u0001\u0000"+
		"\u0000\u0000\u017b\u0155\u0001\u0000\u0000\u0000\u017b\u0158\u0001\u0000"+
		"\u0000\u0000\u017b\u015a\u0001\u0000\u0000\u0000\u017b\u015c\u0001\u0000"+
		"\u0000\u0000\u017b\u015e\u0001\u0000\u0000\u0000\u017b\u0160\u0001\u0000"+
		"\u0000\u0000\u017b\u0162\u0001\u0000\u0000\u0000\u017b\u0164\u0001\u0000"+
		"\u0000\u0000\u017b\u0166\u0001\u0000\u0000\u0000\u017b\u0169\u0001\u0000"+
		"\u0000\u0000\u017b\u016c\u0001\u0000\u0000\u0000\u017b\u016f\u0001\u0000"+
		"\u0000\u0000\u017b\u0172\u0001\u0000\u0000\u0000\u017b\u0175\u0001\u0000"+
		"\u0000\u0000\u017b\u0178\u0001\u0000\u0000\u0000\u017c\u01a7\u0001\u0000"+
		"\u0000\u0000\u017d\u017e\n\u0018\u0000\u0000\u017e\u017f\u00059\u0000"+
		"\u0000\u017f\u0180\u0003\u0012\t\u0019\u0180\u0181\u0006\t\uffff\uffff"+
		"\u0000\u0181\u01a6\u0001\u0000\u0000\u0000\u0182\u0183\n\u0017\u0000\u0000"+
		"\u0183\u0184\u0007\u0000\u0000\u0000\u0184\u0185\u0003\u0012\t\u0018\u0185"+
		"\u0186\u0006\t\uffff\uffff\u0000\u0186\u01a6\u0001\u0000\u0000\u0000\u0187"+
		"\u0188\n\u0016\u0000\u0000\u0188\u0189\u0007\u0001\u0000\u0000\u0189\u018a"+
		"\u0003\u0012\t\u0017\u018a\u018b\u0006\t\uffff\uffff\u0000\u018b\u01a6"+
		"\u0001\u0000\u0000\u0000\u018c\u018d\n\u0015\u0000\u0000\u018d\u018e\u0007"+
		"\u0002\u0000\u0000\u018e\u018f\u0003\u0012\t\u0016\u018f\u0190\u0006\t"+
		"\uffff\uffff\u0000\u0190\u01a6\u0001\u0000\u0000\u0000\u0191\u0192\n\u0014"+
		"\u0000\u0000\u0192\u0193\u0007\u0003\u0000\u0000\u0193\u0194\u0003\u0012"+
		"\t\u0015\u0194\u0195\u0006\t\uffff\uffff\u0000\u0195\u01a6\u0001\u0000"+
		"\u0000\u0000\u0196\u0197\n\u0013\u0000\u0000\u0197\u0198\u0007\u0004\u0000"+
		"\u0000\u0198\u0199\u0003\u0012\t\u0014\u0199\u019a\u0006\t\uffff\uffff"+
		"\u0000\u019a\u01a6\u0001\u0000\u0000\u0000\u019b\u019c\n\u0012\u0000\u0000"+
		"\u019c\u019d\u00054\u0000\u0000\u019d\u019e\u0003\u0012\t\u0013\u019e"+
		"\u019f\u0006\t\uffff\uffff\u0000\u019f\u01a6\u0001\u0000\u0000\u0000\u01a0"+
		"\u01a1\n\u0011\u0000\u0000\u01a1\u01a2\u00053\u0000\u0000\u01a2\u01a3"+
		"\u0003\u0012\t\u0012\u01a3\u01a4\u0006\t\uffff\uffff\u0000\u01a4\u01a6"+
		"\u0001\u0000\u0000\u0000\u01a5\u017d\u0001\u0000\u0000\u0000\u01a5\u0182"+
		"\u0001\u0000\u0000\u0000\u01a5\u0187\u0001\u0000\u0000\u0000\u01a5\u018c"+
		"\u0001\u0000\u0000\u0000\u01a5\u0191\u0001\u0000\u0000\u0000\u01a5\u0196"+
		"\u0001\u0000\u0000\u0000\u01a5\u019b\u0001\u0000\u0000\u0000\u01a5\u01a0"+
		"\u0001\u0000\u0000\u0000\u01a6\u01a9\u0001\u0000\u0000\u0000\u01a7\u01a5"+
		"\u0001\u0000\u0000\u0000\u01a7\u01a8\u0001\u0000\u0000\u0000\u01a8\u0013"+
		"\u0001\u0000\u0000\u0000\u01a9\u01a7\u0001\u0000\u0000\u0000\u01aa\u01ab"+
		"\u0005\u000b\u0000\u0000\u01ab\u01ac\u0003\u0012\t\u0000\u01ac\u01ad\u0005"+
		"@\u0000\u0000\u01ad\u01ae\u0003\u0006\u0003\u0000\u01ae\u01af\u0005A\u0000"+
		"\u0000\u01af\u01b0\u0006\n\uffff\uffff\u0000\u01b0\u01c6\u0001\u0000\u0000"+
		"\u0000\u01b1\u01b2\u0005\u000b\u0000\u0000\u01b2\u01b3\u0003\u0012\t\u0000"+
		"\u01b3\u01b4\u0005@\u0000\u0000\u01b4\u01b5\u0003\u0006\u0003\u0000\u01b5"+
		"\u01b6\u0005A\u0000\u0000\u01b6\u01b7\u0005\f\u0000\u0000\u01b7\u01b8"+
		"\u0005@\u0000\u0000\u01b8\u01b9\u0003\u0006\u0003\u0000\u01b9\u01ba\u0005"+
		"A\u0000\u0000\u01ba\u01bb\u0006\n\uffff\uffff\u0000\u01bb\u01c6\u0001"+
		"\u0000\u0000\u0000\u01bc\u01bd\u0005\u000b\u0000\u0000\u01bd\u01be\u0003"+
		"\u0012\t\u0000\u01be\u01bf\u0005@\u0000\u0000\u01bf\u01c0\u0003\u0006"+
		"\u0003\u0000\u01c0\u01c1\u0005A\u0000\u0000\u01c1\u01c2\u0005\f\u0000"+
		"\u0000\u01c2\u01c3\u0003\u0014\n\u0000\u01c3\u01c4\u0006\n\uffff\uffff"+
		"\u0000\u01c4\u01c6\u0001\u0000\u0000\u0000\u01c5\u01aa\u0001\u0000\u0000"+
		"\u0000\u01c5\u01b1\u0001\u0000\u0000\u0000\u01c5\u01bc\u0001\u0000\u0000"+
		"\u0000\u01c6\u0015\u0001\u0000\u0000\u0000\u01c7\u01c8\u0005\r\u0000\u0000"+
		"\u01c8\u01c9\u0003\u0012\t\u0000\u01c9\u01ca\u0005@\u0000\u0000\u01ca"+
		"\u01ce\u0003\u0018\f\u0000\u01cb\u01cc\u0005\u000f\u0000\u0000\u01cc\u01cd"+
		"\u0005+\u0000\u0000\u01cd\u01cf\u0003\u0006\u0003\u0000\u01ce\u01cb\u0001"+
		"\u0000\u0000\u0000\u01ce\u01cf\u0001\u0000\u0000\u0000\u01cf\u01d0\u0001"+
		"\u0000\u0000\u0000\u01d0\u01d1\u0005A\u0000\u0000\u01d1\u01d2\u0006\u000b"+
		"\uffff\uffff\u0000\u01d2\u0017\u0001\u0000\u0000\u0000\u01d3\u01d5\u0003"+
		"\u001a\r\u0000\u01d4\u01d3\u0001\u0000\u0000\u0000\u01d5\u01d6\u0001\u0000"+
		"\u0000\u0000\u01d6\u01d4\u0001\u0000\u0000\u0000\u01d6\u01d7\u0001\u0000"+
		"\u0000\u0000\u01d7\u01d8\u0001\u0000\u0000\u0000\u01d8\u01d9\u0006\f\uffff"+
		"\uffff\u0000\u01d9\u0019\u0001\u0000\u0000\u0000\u01da\u01db\u0005\u000e"+
		"\u0000\u0000\u01db\u01dc\u0003\u0012\t\u0000\u01dc\u01dd\u0005+\u0000"+
		"\u0000\u01dd\u01de\u0003\u0006\u0003\u0000\u01de\u01df\u0006\r\uffff\uffff"+
		"\u0000\u01df\u001b\u0001\u0000\u0000\u0000\u01e0\u01e1\u0005\u0015\u0000"+
		"\u0000\u01e1\u01e2\u0003\u0012\t\u0000\u01e2\u01e3\u0005@\u0000\u0000"+
		"\u01e3\u01e4\u0003\u0006\u0003\u0000\u01e4\u01e5\u0005A\u0000\u0000\u01e5"+
		"\u01e6\u0006\u000e\uffff\uffff\u0000\u01e6\u001d\u0001\u0000\u0000\u0000"+
		"\u01e7\u01e8\u0005\u0012\u0000\u0000\u01e8\u01e9\u0005\'\u0000\u0000\u01e9"+
		"\u01ea\u0005\u0013\u0000\u0000\u01ea\u01eb\u0003\u0012\t\u0000\u01eb\u01ec"+
		"\u0005\u0014\u0000\u0000\u01ec\u01ed\u0003\u0012\t\u0000\u01ed\u01ee\u0005"+
		"@\u0000\u0000\u01ee\u01ef\u0003\u0006\u0003\u0000\u01ef\u01f0\u0005A\u0000"+
		"\u0000\u01f0\u01f1\u0006\u000f\uffff\uffff\u0000\u01f1\u0205\u0001\u0000"+
		"\u0000\u0000\u01f2\u01f3\u0005\u0012\u0000\u0000\u01f3\u01f4\u0005\'\u0000"+
		"\u0000\u01f4\u01f5\u0005\u0013\u0000\u0000\u01f5\u01f6\u0005\'\u0000\u0000"+
		"\u01f6\u01f7\u0005@\u0000\u0000\u01f7\u01f8\u0003\u0006\u0003\u0000\u01f8"+
		"\u01f9\u0005A\u0000\u0000\u01f9\u01fa\u0006\u000f\uffff\uffff\u0000\u01fa"+
		"\u0205\u0001\u0000\u0000\u0000\u01fb\u01fc\u0005\u0012\u0000\u0000\u01fc"+
		"\u01fd\u0005\'\u0000\u0000\u01fd\u01fe\u0005\u0013\u0000\u0000\u01fe\u01ff"+
		"\u0003\u0012\t\u0000\u01ff\u0200\u0005@\u0000\u0000\u0200\u0201\u0003"+
		"\u0006\u0003\u0000\u0201\u0202\u0005A\u0000\u0000\u0202\u0203\u0006\u000f"+
		"\uffff\uffff\u0000\u0203\u0205\u0001\u0000\u0000\u0000\u0204\u01e7\u0001"+
		"\u0000\u0000\u0000\u0204\u01f2\u0001\u0000\u0000\u0000\u0204\u01fb\u0001"+
		"\u0000\u0000\u0000\u0205\u001f\u0001\u0000\u0000\u0000\u0206\u0207\u0005"+
		"\u0016\u0000\u0000\u0207\u0208\u0003\u0012\t\u0000\u0208\u0209\u0005\f"+
		"\u0000\u0000\u0209\u020a\u0005@\u0000\u0000\u020a\u020b\u0003\u0006\u0003"+
		"\u0000\u020b\u020c\u0005A\u0000\u0000\u020c\u020d\u0006\u0010\uffff\uffff"+
		"\u0000\u020d!\u0001\u0000\u0000\u0000\u020e\u020f\u0005\u0011\u0000\u0000"+
		"\u020f\u0210\u0006\u0011\uffff\uffff\u0000\u0210#\u0001\u0000\u0000\u0000"+
		"\u0211\u0212\u0005\u0010\u0000\u0000\u0212\u0213\u0006\u0012\uffff\uffff"+
		"\u0000\u0213%\u0001\u0000\u0000\u0000\u0214\u0215\u0005\u0017\u0000\u0000"+
		"\u0215\u0216\u0003\u0012\t\u0000\u0216\u0217\u0006\u0013\uffff\uffff\u0000"+
		"\u0217\u021b\u0001\u0000\u0000\u0000\u0218\u0219\u0005\u0017\u0000\u0000"+
		"\u0219\u021b\u0006\u0013\uffff\uffff\u0000\u021a\u0214\u0001\u0000\u0000"+
		"\u0000\u021a\u0218\u0001\u0000\u0000\u0000\u021b\'\u0001\u0000\u0000\u0000"+
		"\u021c\u021d\u0005\b\u0000\u0000\u021d\u021e\u0005\'\u0000\u0000\u021e"+
		"\u021f\u0005+\u0000\u0000\u021f\u0220\u0005F\u0000\u0000\u0220\u0221\u0003"+
		"\u0010\b\u0000\u0221\u0222\u0005G\u0000\u0000\u0222\u0223\u0005*\u0000"+
		"\u0000\u0223\u0224\u0005F\u0000\u0000\u0224\u0225\u0003*\u0015\u0000\u0225"+
		"\u0226\u0005G\u0000\u0000\u0226\u0227\u0006\u0014\uffff\uffff\u0000\u0227"+
		"\u023e\u0001\u0000\u0000\u0000\u0228\u0229\u0005\b\u0000\u0000\u0229\u022a"+
		"\u0005\'\u0000\u0000\u022a\u022b\u0005+\u0000\u0000\u022b\u022c\u0005"+
		"F\u0000\u0000\u022c\u022d\u0003\u0010\b\u0000\u022d\u022e\u0005G\u0000"+
		"\u0000\u022e\u022f\u0005*\u0000\u0000\u022f\u0230\u0005F\u0000\u0000\u0230"+
		"\u0231\u0005G\u0000\u0000\u0231\u0232\u0006\u0014\uffff\uffff\u0000\u0232"+
		"\u023e\u0001\u0000\u0000\u0000\u0233\u0234\u0005\b\u0000\u0000\u0234\u0235"+
		"\u0005\'\u0000\u0000\u0235\u0236\u0005+\u0000\u0000\u0236\u0237\u0005"+
		"F\u0000\u0000\u0237\u0238\u0003\u0010\b\u0000\u0238\u0239\u0005G\u0000"+
		"\u0000\u0239\u023a\u0005*\u0000\u0000\u023a\u023b\u0005\'\u0000\u0000"+
		"\u023b\u023c\u0006\u0014\uffff\uffff\u0000\u023c\u023e\u0001\u0000\u0000"+
		"\u0000\u023d\u021c\u0001\u0000\u0000\u0000\u023d\u0228\u0001\u0000\u0000"+
		"\u0000\u023d\u0233\u0001\u0000\u0000\u0000\u023e)\u0001\u0000\u0000\u0000"+
		"\u023f\u0241\u0003,\u0016\u0000\u0240\u023f\u0001\u0000\u0000\u0000\u0241"+
		"\u0242\u0001\u0000\u0000\u0000\u0242\u0240\u0001\u0000\u0000\u0000\u0242"+
		"\u0243\u0001\u0000\u0000\u0000\u0243\u0244\u0001\u0000\u0000\u0000\u0244"+
		"\u0245\u0006\u0015\uffff\uffff\u0000\u0245+\u0001\u0000\u0000\u0000\u0246"+
		"\u0247\u0005C\u0000\u0000\u0247\u0248\u0003\u0012\t\u0000\u0248\u0249"+
		"\u0006\u0016\uffff\uffff\u0000\u0249\u024e\u0001\u0000\u0000\u0000\u024a"+
		"\u024b\u0003\u0012\t\u0000\u024b\u024c\u0006\u0016\uffff\uffff\u0000\u024c"+
		"\u024e\u0001\u0000\u0000\u0000\u024d\u0246\u0001\u0000\u0000\u0000\u024d"+
		"\u024a\u0001\u0000\u0000\u0000\u024e-\u0001\u0000\u0000\u0000\u024f\u0250"+
		"\u0005\'\u0000\u0000\u0250\u0251\u0005D\u0000\u0000\u0251\u0252\u0005"+
		"\u001b\u0000\u0000\u0252\u0253\u0005.\u0000\u0000\u0253\u0254\u0003\u0012"+
		"\t\u0000\u0254\u0255\u0005/\u0000\u0000\u0255\u0256\u0006\u0017\uffff"+
		"\uffff\u0000\u0256\u0282\u0001\u0000\u0000\u0000\u0257\u0258\u0005\'\u0000"+
		"\u0000\u0258\u0259\u0005F\u0000\u0000\u0259\u025a\u0003\u0012\t\u0000"+
		"\u025a\u025b\u0005G\u0000\u0000\u025b\u025c\u0005*\u0000\u0000\u025c\u025d"+
		"\u0005\'\u0000\u0000\u025d\u025e\u0005F\u0000\u0000\u025e\u025f\u0003"+
		"\u0012\t\u0000\u025f\u0260\u0005G\u0000\u0000\u0260\u0261\u0006\u0017"+
		"\uffff\uffff\u0000\u0261\u0282\u0001\u0000\u0000\u0000\u0262\u0263\u0005"+
		"\'\u0000\u0000\u0263\u0264\u0005F\u0000\u0000\u0264\u0265\u0003\u0012"+
		"\t\u0000\u0265\u0266\u0005G\u0000\u0000\u0266\u0267\u0005F\u0000\u0000"+
		"\u0267\u0268\u0003\u0012\t\u0000\u0268\u0269\u0005G\u0000\u0000\u0269"+
		"\u026a\u0003H$\u0000\u026a\u026b\u0005*\u0000\u0000\u026b\u026c\u0003"+
		"\u0012\t\u0000\u026c\u026d\u0006\u0017\uffff\uffff\u0000\u026d\u0282\u0001"+
		"\u0000\u0000\u0000\u026e\u026f\u0005\'\u0000\u0000\u026f\u0270\u0005F"+
		"\u0000\u0000\u0270\u0271\u0003\u0012\t\u0000\u0271\u0272\u0005G\u0000"+
		"\u0000\u0272\u0273\u0005F\u0000\u0000\u0273\u0274\u0003\u0012\t\u0000"+
		"\u0274\u0275\u0005G\u0000\u0000\u0275\u0276\u0005*\u0000\u0000\u0276\u0277"+
		"\u0003\u0012\t\u0000\u0277\u0278\u0006\u0017\uffff\uffff\u0000\u0278\u0282"+
		"\u0001\u0000\u0000\u0000\u0279\u027a\u0005\'\u0000\u0000\u027a\u027b\u0005"+
		"F\u0000\u0000\u027b\u027c\u0003\u0012\t\u0000\u027c\u027d\u0005G\u0000"+
		"\u0000\u027d\u027e\u0005*\u0000\u0000\u027e\u027f\u0003\u0012\t\u0000"+
		"\u027f\u0280\u0006\u0017\uffff\uffff\u0000\u0280\u0282\u0001\u0000\u0000"+
		"\u0000\u0281\u024f\u0001\u0000\u0000\u0000\u0281\u0257\u0001\u0000\u0000"+
		"\u0000\u0281\u0262\u0001\u0000\u0000\u0000\u0281\u026e\u0001\u0000\u0000"+
		"\u0000\u0281\u0279\u0001\u0000\u0000\u0000\u0282/\u0001\u0000\u0000\u0000"+
		"\u0283\u0284\u0005\'\u0000\u0000\u0284\u0285\u0005D\u0000\u0000\u0285"+
		"\u0286\u0005\u001d\u0000\u0000\u0286\u0287\u0005.\u0000\u0000\u0287\u0288"+
		"\u0005/\u0000\u0000\u0288\u0294\u0006\u0018\uffff\uffff\u0000\u0289\u028a"+
		"\u0005\'\u0000\u0000\u028a\u028b\u0005D\u0000\u0000\u028b\u028c\u0005"+
		"\u001c\u0000\u0000\u028c\u028d\u0005.\u0000\u0000\u028d\u028e\u0005 \u0000"+
		"\u0000\u028e\u028f\u0005+\u0000\u0000\u028f\u0290\u0003\u0012\t\u0000"+
		"\u0290\u0291\u0005/\u0000\u0000\u0291\u0292\u0006\u0018\uffff\uffff\u0000"+
		"\u0292\u0294\u0001\u0000\u0000\u0000\u0293\u0283\u0001\u0000\u0000\u0000"+
		"\u0293\u0289\u0001\u0000\u0000\u0000\u02941\u0001\u0000\u0000\u0000\u0295"+
		"\u0296\u0005\'\u0000\u0000\u0296\u0297\u0005D\u0000\u0000\u0297\u0298"+
		"\u0005\u001f\u0000\u0000\u0298\u0299\u0006\u0019\uffff\uffff\u0000\u0299"+
		"3\u0001\u0000\u0000\u0000\u029a\u029b\u0005\'\u0000\u0000\u029b\u029c"+
		"\u0005D\u0000\u0000\u029c\u029d\u0005\u001e\u0000\u0000\u029d\u029e\u0006"+
		"\u001a\uffff\uffff\u0000\u029e5\u0001\u0000\u0000\u0000\u029f\u02a0\u0005"+
		"\'\u0000\u0000\u02a0\u02a1\u0005F\u0000\u0000\u02a1\u02a2\u0003\u0012"+
		"\t\u0000\u02a2\u02a3\u0005G\u0000\u0000\u02a3\u02a4\u0005F\u0000\u0000"+
		"\u02a4\u02a5\u0003\u0012\t\u0000\u02a5\u02a6\u0005G\u0000\u0000\u02a6"+
		"\u02a7\u0003H$\u0000\u02a7\u02a8\u0006\u001b\uffff\uffff\u0000\u02a8\u02b9"+
		"\u0001\u0000\u0000\u0000\u02a9\u02aa\u0005\'\u0000\u0000\u02aa\u02ab\u0005"+
		"F\u0000\u0000\u02ab\u02ac\u0003\u0012\t\u0000\u02ac\u02ad\u0005G\u0000"+
		"\u0000\u02ad\u02ae\u0005F\u0000\u0000\u02ae\u02af\u0003\u0012\t\u0000"+
		"\u02af\u02b0\u0005G\u0000\u0000\u02b0\u02b1\u0006\u001b\uffff\uffff\u0000"+
		"\u02b1\u02b9\u0001\u0000\u0000\u0000\u02b2\u02b3\u0005\'\u0000\u0000\u02b3"+
		"\u02b4\u0005F\u0000\u0000\u02b4\u02b5\u0003\u0012\t\u0000\u02b5\u02b6"+
		"\u0005G\u0000\u0000\u02b6\u02b7\u0006\u001b\uffff\uffff\u0000\u02b7\u02b9"+
		"\u0001\u0000\u0000\u0000\u02b8\u029f\u0001\u0000\u0000\u0000\u02b8\u02a9"+
		"\u0001\u0000\u0000\u0000\u02b8\u02b2\u0001\u0000\u0000\u0000\u02b97\u0001"+
		"\u0000\u0000\u0000\u02ba\u02bb\u0005\b\u0000\u0000\u02bb\u02be\u0005\'"+
		"\u0000\u0000\u02bc\u02bd\u0005+\u0000\u0000\u02bd\u02bf\u0003:\u001d\u0000"+
		"\u02be\u02bc\u0001\u0000\u0000\u0000\u02be\u02bf\u0001\u0000\u0000\u0000"+
		"\u02bf\u02c0\u0001\u0000\u0000\u0000\u02c0\u02c1\u0005*\u0000\u0000\u02c1"+
		"\u02c2\u0003<\u001e\u0000\u02c2\u02c3\u0006\u001c\uffff\uffff\u0000\u02c3"+
		"9\u0001\u0000\u0000\u0000\u02c4\u02c5\u0005F\u0000\u0000\u02c5\u02c6\u0003"+
		":\u001d\u0000\u02c6\u02c7\u0005G\u0000\u0000\u02c7\u02c8\u0006\u001d\uffff"+
		"\uffff\u0000\u02c8\u02cf\u0001\u0000\u0000\u0000\u02c9\u02ca\u0005F\u0000"+
		"\u0000\u02ca\u02cb\u0003\u0010\b\u0000\u02cb\u02cc\u0005G\u0000\u0000"+
		"\u02cc\u02cd\u0006\u001d\uffff\uffff\u0000\u02cd\u02cf\u0001\u0000\u0000"+
		"\u0000\u02ce\u02c4\u0001\u0000\u0000\u0000\u02ce\u02c9\u0001\u0000\u0000"+
		"\u0000\u02cf;\u0001\u0000\u0000\u0000\u02d0\u02d1\u0003>\u001f\u0000\u02d1"+
		"\u02d2\u0006\u001e\uffff\uffff\u0000\u02d2=\u0001\u0000\u0000\u0000\u02d3"+
		"\u02d4\u0005F\u0000\u0000\u02d4\u02d5\u0003@ \u0000\u02d5\u02d6\u0005"+
		"G\u0000\u0000\u02d6\u02d7\u0006\u001f\uffff\uffff\u0000\u02d7\u02dc\u0001"+
		"\u0000\u0000\u0000\u02d8\u02d9\u0003F#\u0000\u02d9\u02da\u0006\u001f\uffff"+
		"\uffff\u0000\u02da\u02dc\u0001\u0000\u0000\u0000\u02db\u02d3\u0001\u0000"+
		"\u0000\u0000\u02db\u02d8\u0001\u0000\u0000\u0000\u02dc?\u0001\u0000\u0000"+
		"\u0000\u02dd\u02de\u0006 \uffff\uffff\u0000\u02de\u02df\u0003>\u001f\u0000"+
		"\u02df\u02e0\u0006 \uffff\uffff\u0000\u02e0\u02e5\u0001\u0000\u0000\u0000"+
		"\u02e1\u02e2\u0003B!\u0000\u02e2\u02e3\u0006 \uffff\uffff\u0000\u02e3"+
		"\u02e5\u0001\u0000\u0000\u0000\u02e4\u02dd\u0001\u0000\u0000\u0000\u02e4"+
		"\u02e1\u0001\u0000\u0000\u0000\u02e5\u02ed\u0001\u0000\u0000\u0000\u02e6"+
		"\u02e7\n\u0003\u0000\u0000\u02e7\u02e8\u0005C\u0000\u0000\u02e8\u02e9"+
		"\u0003>\u001f\u0000\u02e9\u02ea\u0006 \uffff\uffff\u0000\u02ea\u02ec\u0001"+
		"\u0000\u0000\u0000\u02eb\u02e6\u0001\u0000\u0000\u0000\u02ec\u02ef\u0001"+
		"\u0000\u0000\u0000\u02ed\u02eb\u0001\u0000\u0000\u0000\u02ed\u02ee\u0001"+
		"\u0000\u0000\u0000\u02eeA\u0001\u0000\u0000\u0000\u02ef\u02ed\u0001\u0000"+
		"\u0000\u0000\u02f0\u02f2\u0003D\"\u0000\u02f1\u02f0\u0001\u0000\u0000"+
		"\u0000\u02f2\u02f3\u0001\u0000\u0000\u0000\u02f3\u02f1\u0001\u0000\u0000"+
		"\u0000\u02f3\u02f4\u0001\u0000\u0000\u0000\u02f4\u02f5\u0001\u0000\u0000"+
		"\u0000\u02f5\u02f6\u0006!\uffff\uffff\u0000\u02f6C\u0001\u0000\u0000\u0000"+
		"\u02f7\u02f8\u0005C\u0000\u0000\u02f8\u02f9\u0003\u0012\t\u0000\u02f9"+
		"\u02fa\u0006\"\uffff\uffff\u0000\u02fa\u02ff\u0001\u0000\u0000\u0000\u02fb"+
		"\u02fc\u0003\u0012\t\u0000\u02fc\u02fd\u0006\"\uffff\uffff\u0000\u02fd"+
		"\u02ff\u0001\u0000\u0000\u0000\u02fe\u02f7\u0001\u0000\u0000\u0000\u02fe"+
		"\u02fb\u0001\u0000\u0000\u0000\u02ffE\u0001\u0000\u0000\u0000\u0300\u0301"+
		"\u0003:\u001d\u0000\u0301\u0302\u0005.\u0000\u0000\u0302\u0303\u0005!"+
		"\u0000\u0000\u0303\u0304\u0005+\u0000\u0000\u0304\u0305\u0003F#\u0000"+
		"\u0305\u0306\u0005C\u0000\u0000\u0306\u0307\u0005\u001e\u0000\u0000\u0307"+
		"\u0308\u0005+\u0000\u0000\u0308\u0309\u0005%\u0000\u0000\u0309\u030a\u0005"+
		"/\u0000\u0000\u030a\u030b\u0006#\uffff\uffff\u0000\u030b\u0319\u0001\u0000"+
		"\u0000\u0000\u030c\u030d\u0003:\u001d\u0000\u030d\u030e\u0005.\u0000\u0000"+
		"\u030e\u030f\u0005!\u0000\u0000\u030f\u0310\u0005+\u0000\u0000\u0310\u0311"+
		"\u0003\u0012\t\u0000\u0311\u0312\u0005C\u0000\u0000\u0312\u0313\u0005"+
		"\u001e\u0000\u0000\u0313\u0314\u0005+\u0000\u0000\u0314\u0315\u0005%\u0000"+
		"\u0000\u0315\u0316\u0005/\u0000\u0000\u0316\u0317\u0006#\uffff\uffff\u0000"+
		"\u0317\u0319\u0001\u0000\u0000\u0000\u0318\u0300\u0001\u0000\u0000\u0000"+
		"\u0318\u030c\u0001\u0000\u0000\u0000\u0319G\u0001\u0000\u0000\u0000\u031a"+
		"\u031c\u0003J%\u0000\u031b\u031a\u0001\u0000\u0000\u0000\u031c\u031d\u0001"+
		"\u0000\u0000\u0000\u031d\u031b\u0001\u0000\u0000\u0000\u031d\u031e\u0001"+
		"\u0000\u0000\u0000\u031e\u031f\u0001\u0000\u0000\u0000\u031f\u0320\u0006"+
		"$\uffff\uffff\u0000\u0320I\u0001\u0000\u0000\u0000\u0321\u0322\u0005F"+
		"\u0000\u0000\u0322\u0323\u0003\u0012\t\u0000\u0323\u0324\u0005G\u0000"+
		"\u0000\u0324\u0325\u0006%\uffff\uffff\u0000\u0325K\u0001\u0000\u0000\u0000"+
		"\u0326\u0327\u0005\u0018\u0000\u0000\u0327\u0328\u0005\'\u0000\u0000\u0328"+
		"\u0329\u0005.\u0000\u0000\u0329\u032a\u0003N\'\u0000\u032a\u032b\u0005"+
		"/\u0000\u0000\u032b\u032c\u0005B\u0000\u0000\u032c\u032d\u0003\u0010\b"+
		"\u0000\u032d\u032e\u0005@\u0000\u0000\u032e\u032f\u0003\u0006\u0003\u0000"+
		"\u032f\u0330\u0005A\u0000\u0000\u0330\u0331\u0006&\uffff\uffff\u0000\u0331"+
		"\u0351\u0001\u0000\u0000\u0000\u0332\u0333\u0005\u0018\u0000\u0000\u0333"+
		"\u0334\u0005\'\u0000\u0000\u0334\u0335\u0005.\u0000\u0000\u0335\u0336"+
		"\u0005/\u0000\u0000\u0336\u0337\u0005B\u0000\u0000\u0337\u0338\u0003\u0010"+
		"\b\u0000\u0338\u0339\u0005@\u0000\u0000\u0339\u033a\u0003\u0006\u0003"+
		"\u0000\u033a\u033b\u0005A\u0000\u0000\u033b\u033c\u0006&\uffff\uffff\u0000"+
		"\u033c\u0351\u0001\u0000\u0000\u0000\u033d\u033e\u0005\u0018\u0000\u0000"+
		"\u033e\u033f\u0005\'\u0000\u0000\u033f\u0340\u0005.\u0000\u0000\u0340"+
		"\u0341\u0003N\'\u0000\u0341\u0342\u0005/\u0000\u0000\u0342\u0343\u0005"+
		"@\u0000\u0000\u0343\u0344\u0003\u0006\u0003\u0000\u0344\u0345\u0005A\u0000"+
		"\u0000\u0345\u0346\u0006&\uffff\uffff\u0000\u0346\u0351\u0001\u0000\u0000"+
		"\u0000\u0347\u0348\u0005\u0018\u0000\u0000\u0348\u0349\u0005\'\u0000\u0000"+
		"\u0349\u034a\u0005.\u0000\u0000\u034a\u034b\u0005/\u0000\u0000\u034b\u034c"+
		"\u0005@\u0000\u0000\u034c\u034d\u0003\u0006\u0003\u0000\u034d\u034e\u0005"+
		"A\u0000\u0000\u034e\u034f\u0006&\uffff\uffff\u0000\u034f\u0351\u0001\u0000"+
		"\u0000\u0000\u0350\u0326\u0001\u0000\u0000\u0000\u0350\u0332\u0001\u0000"+
		"\u0000\u0000\u0350\u033d\u0001\u0000\u0000\u0000\u0350\u0347\u0001\u0000"+
		"\u0000\u0000\u0351M\u0001\u0000\u0000\u0000\u0352\u0354\u0007\u0005\u0000"+
		"\u0000\u0353\u0352\u0001\u0000\u0000\u0000\u0353\u0354\u0001\u0000\u0000"+
		"\u0000\u0354\u0355\u0001\u0000\u0000\u0000\u0355\u0356\u0005\'\u0000\u0000"+
		"\u0356\u0358\u0005+\u0000\u0000\u0357\u0359\u0005\u001a\u0000\u0000\u0358"+
		"\u0357\u0001\u0000\u0000\u0000\u0358\u0359\u0001\u0000\u0000\u0000\u0359"+
		"\u035a\u0001\u0000\u0000\u0000\u035a\u035b\u0003\u0010\b\u0000\u035b\u035c"+
		"\u0005C\u0000\u0000\u035c\u035d\u0003N\'\u0000\u035d\u035e\u0006\'\uffff"+
		"\uffff\u0000\u035e\u036b\u0001\u0000\u0000\u0000\u035f\u0361\u0007\u0005"+
		"\u0000\u0000\u0360\u035f\u0001\u0000\u0000\u0000\u0360\u0361\u0001\u0000"+
		"\u0000\u0000\u0361\u0362\u0001\u0000\u0000\u0000\u0362\u0363\u0005\'\u0000"+
		"\u0000\u0363\u0365\u0005+\u0000\u0000\u0364\u0366\u0005\u001a\u0000\u0000"+
		"\u0365\u0364\u0001\u0000\u0000\u0000\u0365\u0366\u0001\u0000\u0000\u0000"+
		"\u0366\u0367\u0001\u0000\u0000\u0000\u0367\u0368\u0003\u0010\b\u0000\u0368"+
		"\u0369\u0006\'\uffff\uffff\u0000\u0369\u036b\u0001\u0000\u0000\u0000\u036a"+
		"\u0353\u0001\u0000\u0000\u0000\u036a\u0360\u0001\u0000\u0000\u0000\u036b"+
		"O\u0001\u0000\u0000\u0000\u036c\u036d\u0005\'\u0000\u0000\u036d\u036e"+
		"\u0005.\u0000\u0000\u036e\u036f\u0003T*\u0000\u036f\u0370\u0005/\u0000"+
		"\u0000\u0370\u0371\u0006(\uffff\uffff\u0000\u0371\u0377\u0001\u0000\u0000"+
		"\u0000\u0372\u0373\u0005\'\u0000\u0000\u0373\u0374\u0005.\u0000\u0000"+
		"\u0374\u0375\u0005/\u0000\u0000\u0375\u0377\u0006(\uffff\uffff\u0000\u0376"+
		"\u036c\u0001\u0000\u0000\u0000\u0376\u0372\u0001\u0000\u0000\u0000\u0377"+
		"Q\u0001\u0000\u0000\u0000\u0378\u0379\u0005\'\u0000\u0000\u0379\u037a"+
		"\u0005.\u0000\u0000\u037a\u037b\u0003T*\u0000\u037b\u037c\u0005/\u0000"+
		"\u0000\u037c\u037d\u0006)\uffff\uffff\u0000\u037d\u0383\u0001\u0000\u0000"+
		"\u0000\u037e\u037f\u0005\'\u0000\u0000\u037f\u0380\u0005.\u0000\u0000"+
		"\u0380\u0381\u0005/\u0000\u0000\u0381\u0383\u0006)\uffff\uffff\u0000\u0382"+
		"\u0378\u0001\u0000\u0000\u0000\u0382\u037e\u0001\u0000\u0000\u0000\u0383"+
		"S\u0001\u0000\u0000\u0000\u0384\u0385\u0005H\u0000\u0000\u0385\u0386\u0005"+
		"\'\u0000\u0000\u0386\u0387\u0005C\u0000\u0000\u0387\u0388\u0003T*\u0000"+
		"\u0388\u0389\u0006*\uffff\uffff\u0000\u0389\u039e\u0001\u0000\u0000\u0000"+
		"\u038a\u038b\u0005H\u0000\u0000\u038b\u038c\u0005\'\u0000\u0000\u038c"+
		"\u039e\u0006*\uffff\uffff\u0000\u038d\u038e\u0005\'\u0000\u0000\u038e"+
		"\u0390\u0005+\u0000\u0000\u038f\u038d\u0001\u0000\u0000\u0000\u038f\u0390"+
		"\u0001\u0000\u0000\u0000\u0390\u0391\u0001\u0000\u0000\u0000\u0391\u0392"+
		"\u0003\u0012\t\u0000\u0392\u0393\u0005C\u0000\u0000\u0393\u0394\u0003"+
		"T*\u0000\u0394\u0395\u0006*\uffff\uffff\u0000\u0395\u039e\u0001\u0000"+
		"\u0000\u0000\u0396\u0397\u0005\'\u0000\u0000\u0397\u0399\u0005+\u0000"+
		"\u0000\u0398\u0396\u0001\u0000\u0000\u0000\u0398\u0399\u0001\u0000\u0000"+
		"\u0000\u0399\u039a\u0001\u0000\u0000\u0000\u039a\u039b\u0003\u0012\t\u0000"+
		"\u039b\u039c\u0006*\uffff\uffff\u0000\u039c\u039e\u0001\u0000\u0000\u0000"+
		"\u039d\u0384\u0001\u0000\u0000\u0000\u039d\u038a\u0001\u0000\u0000\u0000"+
		"\u039d\u038f\u0001\u0000\u0000\u0000\u039d\u0398\u0001\u0000\u0000\u0000"+
		"\u039eU\u0001\u0000\u0000\u0000\u039f\u03a0\u0005\u0019\u0000\u0000\u03a0"+
		"\u03a1\u0005.\u0000\u0000\u03a1\u03a2\u0003B!\u0000\u03a2\u03a3\u0005"+
		"/\u0000\u0000\u03a3\u03a4\u0006+\uffff\uffff\u0000\u03a4W\u0001\u0000"+
		"\u0000\u0000\u03a5\u03a6\u0005\u0001\u0000\u0000\u03a6\u03a7\u0005.\u0000"+
		"\u0000\u03a7\u03a8\u0003\u0012\t\u0000\u03a8\u03a9\u0005/\u0000\u0000"+
		"\u03a9\u03aa\u0006,\uffff\uffff\u0000\u03aaY\u0001\u0000\u0000\u0000\u03ab"+
		"\u03ac\u0005\u0002\u0000\u0000\u03ac\u03ad\u0005.\u0000\u0000\u03ad\u03ae"+
		"\u0003\u0012\t\u0000\u03ae\u03af\u0005/\u0000\u0000\u03af\u03b0\u0006"+
		"-\uffff\uffff\u0000\u03b0[\u0001\u0000\u0000\u0000\u03b1\u03b2\u0005\u0003"+
		"\u0000\u0000\u03b2\u03b3\u0005.\u0000\u0000\u03b3\u03b4\u0003\u0012\t"+
		"\u0000\u03b4\u03b5\u0005/\u0000\u0000\u03b5\u03b6\u0006.\uffff\uffff\u0000"+
		"\u03b6]\u0001\u0000\u0000\u0000;ekqw\u008c\u0098\u009e\u00a8\u00ad\u00b3"+
		"\u00b9\u00bf\u00d4\u00da\u00e0\u00e6\u00ec\u00f2\u00f8\u00fe\u0105\u011c"+
		"\u012c\u013d\u0149\u017b\u01a5\u01a7\u01c5\u01ce\u01d6\u0204\u021a\u023d"+
		"\u0242\u024d\u0281\u0293\u02b8\u02be\u02ce\u02db\u02e4\u02ed\u02f3\u02fe"+
		"\u0318\u031d\u0350\u0353\u0358\u0360\u0365\u036a\u0376\u0382\u038f\u0398"+
		"\u039d";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}