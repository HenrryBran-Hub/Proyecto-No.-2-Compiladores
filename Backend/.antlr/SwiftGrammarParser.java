// Generated from c:\Users\Henrr\OneDrive\Escritorio\Compiladores\Proyecto No. 2\Proyecto No. 2 Compiladores\Backend\SwiftGrammar.g4 by ANTLR 4.9.2

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

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SwiftGrammarParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

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
		RULE_retornos = 19, RULE_listaexpresions = 20, RULE_listaexpresion = 21, 
		RULE_printstmt = 22;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "blockinterno", "instructionint", "declavarible", 
			"declaconstante", "asignacionvariable", "tipodato", "expr", "sentenciaifelse", 
			"switchcontrol", "blockcase", "bloquecase", "whilecontrol", "forcontrol", 
			"guardcontrol", "continuee", "breakk", "retornos", "listaexpresions", 
			"listaexpresion", "printstmt"
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
			setState(46);
			((SContext)_localctx).block = block();
			setState(47);
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
			setState(51); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(50);
				((BlockContext)_localctx).instruction = instruction();
				((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
				}
				}
				setState(53); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << VAR) | (1L << LET) | (1L << IF) | (1L << SWITCH) | (1L << FOR) | (1L << WHILE) | (1L << GUARD) | (1L << PRINT) | (1L << ID_VALIDO))) != 0) );

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
		public PrintstmtContext printstmt;
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
		public PrintstmtContext printstmt() {
			return getRuleContext(PrintstmtContext.class,0);
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
			setState(96);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VAR:
				enterOuterAlt(_localctx, 1);
				{
				setState(57);
				((InstructionContext)_localctx).declavarible = declavarible();
				setState(59);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(58);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).declavarible.decvbl
				}
				break;
			case LET:
				enterOuterAlt(_localctx, 2);
				{
				setState(63);
				((InstructionContext)_localctx).declaconstante = declaconstante();
				setState(65);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(64);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).declaconstante.deccon
				}
				break;
			case ID_VALIDO:
				enterOuterAlt(_localctx, 3);
				{
				setState(69);
				((InstructionContext)_localctx).asignacionvariable = asignacionvariable();
				setState(71);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(70);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).asignacionvariable.asgvbl
				}
				break;
			case IF:
				enterOuterAlt(_localctx, 4);
				{
				setState(75);
				((InstructionContext)_localctx).sentenciaifelse = sentenciaifelse();
				 _localctx.inst = ((InstructionContext)_localctx).sentenciaifelse.myIfElse
				}
				break;
			case SWITCH:
				enterOuterAlt(_localctx, 5);
				{
				setState(78);
				((InstructionContext)_localctx).switchcontrol = switchcontrol();
				 _localctx.inst = ((InstructionContext)_localctx).switchcontrol.mySwitch
				}
				break;
			case WHILE:
				enterOuterAlt(_localctx, 6);
				{
				setState(81);
				((InstructionContext)_localctx).whilecontrol = whilecontrol();
				 _localctx.inst = ((InstructionContext)_localctx).whilecontrol.whict
				}
				break;
			case FOR:
				enterOuterAlt(_localctx, 7);
				{
				setState(84);
				((InstructionContext)_localctx).forcontrol = forcontrol();
				 _localctx.inst = ((InstructionContext)_localctx).forcontrol.forct
				}
				break;
			case GUARD:
				enterOuterAlt(_localctx, 8);
				{
				setState(87);
				((InstructionContext)_localctx).guardcontrol = guardcontrol();
				 _localctx.inst = ((InstructionContext)_localctx).guardcontrol.guct
				}
				break;
			case PRINT:
				enterOuterAlt(_localctx, 9);
				{
				setState(90);
				((InstructionContext)_localctx).printstmt = printstmt();
				setState(92);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(91);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
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
			setState(99); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(98);
				((BlockinternoContext)_localctx).instructionint = instructionint();
				((BlockinternoContext)_localctx).insint.add(((BlockinternoContext)_localctx).instructionint);
				}
				}
				setState(101); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << VAR) | (1L << LET) | (1L << IF) | (1L << SWITCH) | (1L << BREAK) | (1L << CONTINUE) | (1L << FOR) | (1L << WHILE) | (1L << GUARD) | (1L << RETURN) | (1L << PRINT) | (1L << ID_VALIDO))) != 0) );

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
		public PrintstmtContext printstmt;
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
		public PrintstmtContext printstmt() {
			return getRuleContext(PrintstmtContext.class,0);
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
			setState(162);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VAR:
				enterOuterAlt(_localctx, 1);
				{
				setState(105);
				((InstructionintContext)_localctx).declavarible = declavarible();
				setState(107);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(106);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).declavarible.decvbl
				}
				break;
			case LET:
				enterOuterAlt(_localctx, 2);
				{
				setState(111);
				((InstructionintContext)_localctx).declaconstante = declaconstante();
				setState(113);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(112);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).declaconstante.deccon
				}
				break;
			case ID_VALIDO:
				enterOuterAlt(_localctx, 3);
				{
				setState(117);
				((InstructionintContext)_localctx).asignacionvariable = asignacionvariable();
				setState(119);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(118);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).asignacionvariable.asgvbl
				}
				break;
			case IF:
				enterOuterAlt(_localctx, 4);
				{
				setState(123);
				((InstructionintContext)_localctx).sentenciaifelse = sentenciaifelse();
				 _localctx.insint = ((InstructionintContext)_localctx).sentenciaifelse.myIfElse
				}
				break;
			case SWITCH:
				enterOuterAlt(_localctx, 5);
				{
				setState(126);
				((InstructionintContext)_localctx).switchcontrol = switchcontrol();
				 _localctx.insint = ((InstructionintContext)_localctx).switchcontrol.mySwitch
				}
				break;
			case WHILE:
				enterOuterAlt(_localctx, 6);
				{
				setState(129);
				((InstructionintContext)_localctx).whilecontrol = whilecontrol();
				 _localctx.insint = ((InstructionintContext)_localctx).whilecontrol.whict
				}
				break;
			case FOR:
				enterOuterAlt(_localctx, 7);
				{
				setState(132);
				((InstructionintContext)_localctx).forcontrol = forcontrol();
				 _localctx.insint = ((InstructionintContext)_localctx).forcontrol.forct
				}
				break;
			case GUARD:
				enterOuterAlt(_localctx, 8);
				{
				setState(135);
				((InstructionintContext)_localctx).guardcontrol = guardcontrol();
				 _localctx.insint = ((InstructionintContext)_localctx).guardcontrol.guct
				}
				break;
			case CONTINUE:
				enterOuterAlt(_localctx, 9);
				{
				setState(138);
				((InstructionintContext)_localctx).continuee = continuee();
				setState(140);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(139);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).continuee.coct
				}
				break;
			case BREAK:
				enterOuterAlt(_localctx, 10);
				{
				setState(144);
				((InstructionintContext)_localctx).breakk = breakk();
				setState(146);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(145);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).breakk.brkct
				}
				break;
			case RETURN:
				enterOuterAlt(_localctx, 11);
				{
				setState(150);
				((InstructionintContext)_localctx).retornos = retornos();
				setState(152);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(151);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).retornos.rect 
				}
				break;
			case PRINT:
				enterOuterAlt(_localctx, 12);
				{
				setState(156);
				((InstructionintContext)_localctx).printstmt = printstmt();
				setState(158);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(157);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).printstmt.prnt
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
			setState(185);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(164);
				((DeclavaribleContext)_localctx).VAR = match(VAR);
				setState(165);
				((DeclavaribleContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(166);
				match(DOS_PUNTOS);
				setState(167);
				((DeclavaribleContext)_localctx).tipodato = tipodato();
				setState(168);
				match(IG);
				setState(169);
				((DeclavaribleContext)_localctx).expr = expr(0);
				 _localctx.decvbl = datosprimitivos.NewVariableDeclaration((((DeclavaribleContext)_localctx).VAR!=null?((DeclavaribleContext)_localctx).VAR.getLine():0), (((DeclavaribleContext)_localctx).VAR!=null?((DeclavaribleContext)_localctx).VAR.getCharPositionInLine():0), (((DeclavaribleContext)_localctx).ID_VALIDO!=null?((DeclavaribleContext)_localctx).ID_VALIDO.getText():null), ((DeclavaribleContext)_localctx).tipodato.tipo, ((DeclavaribleContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(172);
				((DeclavaribleContext)_localctx).VAR = match(VAR);
				setState(173);
				((DeclavaribleContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(174);
				match(IG);
				setState(175);
				((DeclavaribleContext)_localctx).expr = expr(0);
				_localctx.decvbl = datosprimitivos.NewVariableDeclaracionSinTipo((((DeclavaribleContext)_localctx).VAR!=null?((DeclavaribleContext)_localctx).VAR.getLine():0), (((DeclavaribleContext)_localctx).VAR!=null?((DeclavaribleContext)_localctx).VAR.getCharPositionInLine():0), (((DeclavaribleContext)_localctx).ID_VALIDO!=null?((DeclavaribleContext)_localctx).ID_VALIDO.getText():null), ((DeclavaribleContext)_localctx).expr.e)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(178);
				((DeclavaribleContext)_localctx).VAR = match(VAR);
				setState(179);
				((DeclavaribleContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(180);
				match(DOS_PUNTOS);
				setState(181);
				((DeclavaribleContext)_localctx).tipodato = tipodato();
				setState(182);
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
			setState(201);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(187);
				((DeclaconstanteContext)_localctx).LET = match(LET);
				setState(188);
				((DeclaconstanteContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(189);
				match(DOS_PUNTOS);
				setState(190);
				((DeclaconstanteContext)_localctx).tipodato = tipodato();
				setState(191);
				match(IG);
				setState(192);
				((DeclaconstanteContext)_localctx).expr = expr(0);
				_localctx.deccon = datosprimitivos.NewConstanteDeclaration((((DeclaconstanteContext)_localctx).LET!=null?((DeclaconstanteContext)_localctx).LET.getLine():0), (((DeclaconstanteContext)_localctx).LET!=null?((DeclaconstanteContext)_localctx).LET.getCharPositionInLine():0), (((DeclaconstanteContext)_localctx).ID_VALIDO!=null?((DeclaconstanteContext)_localctx).ID_VALIDO.getText():null), ((DeclaconstanteContext)_localctx).tipodato.tipo, ((DeclaconstanteContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(195);
				((DeclaconstanteContext)_localctx).LET = match(LET);
				setState(196);
				((DeclaconstanteContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(197);
				match(IG);
				setState(198);
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
			setState(218);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(203);
				((AsignacionvariableContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(204);
				match(IG);
				setState(205);
				((AsignacionvariableContext)_localctx).expr = expr(0);
				 _localctx.asgvbl = sentencias.NewAsignacionVariable((((AsignacionvariableContext)_localctx).ID_VALIDO!=null?((AsignacionvariableContext)_localctx).ID_VALIDO.getLine():0), (((AsignacionvariableContext)_localctx).ID_VALIDO!=null?((AsignacionvariableContext)_localctx).ID_VALIDO.getCharPositionInLine():0), (((AsignacionvariableContext)_localctx).ID_VALIDO!=null?((AsignacionvariableContext)_localctx).ID_VALIDO.getText():null), ((AsignacionvariableContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(208);
				((AsignacionvariableContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(209);
				match(SUMA);
				setState(210);
				((AsignacionvariableContext)_localctx).expr = expr(0);
				 _localctx.asgvbl = sentencias.NewAsignacionSuma((((AsignacionvariableContext)_localctx).ID_VALIDO!=null?((AsignacionvariableContext)_localctx).ID_VALIDO.getLine():0), (((AsignacionvariableContext)_localctx).ID_VALIDO!=null?((AsignacionvariableContext)_localctx).ID_VALIDO.getCharPositionInLine():0), (((AsignacionvariableContext)_localctx).ID_VALIDO!=null?((AsignacionvariableContext)_localctx).ID_VALIDO.getText():null), ((AsignacionvariableContext)_localctx).expr.e)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(213);
				((AsignacionvariableContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(214);
				match(RESTA);
				setState(215);
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
			setState(230);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(220);
				match(INT);
				 _localctx.tipo = environment.INTEGER 
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(222);
				match(FLOAT);
				 _localctx.tipo = environment.FLOAT 
				}
				break;
			case STRING:
				enterOuterAlt(_localctx, 3);
				{
				setState(224);
				match(STRING);
				 _localctx.tipo = environment.STRING 
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(226);
				match(BOOL);
				 _localctx.tipo = environment.BOOLEAN 
				}
				break;
			case CHARACT:
				enterOuterAlt(_localctx, 5);
				{
				setState(228);
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
			setState(259);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NOT:
				{
				setState(233);
				((ExprContext)_localctx).op = match(NOT);
				setState(234);
				((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(18);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).right!=null?(((ExprContext)_localctx).right.start):null).GetLine(), (((ExprContext)_localctx).right!=null?(((ExprContext)_localctx).right.start):null).GetColumn(), ((ExprContext)_localctx).right.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
				}
				break;
			case PARIZQ:
				{
				setState(237);
				match(PARIZQ);
				setState(238);
				((ExprContext)_localctx).expr = expr(0);
				setState(239);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case SUB:
				{
				setState(242);
				match(SUB);
				setState(243);
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
			case NUMBER:
				{
				setState(245);
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
			case CADENA:
				{
				setState(247);
				((ExprContext)_localctx).CADENA = match(CADENA);

				        str := (((ExprContext)_localctx).CADENA!=null?((ExprContext)_localctx).CADENA.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).CADENA!=null?((ExprContext)_localctx).CADENA.getLine():0), (((ExprContext)_localctx).CADENA!=null?((ExprContext)_localctx).CADENA.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case TRU:
				{
				setState(249);
				((ExprContext)_localctx).TRU = match(TRU);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getLine():0), (((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case FAL:
				{
				setState(251);
				((ExprContext)_localctx).FAL = match(FAL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getLine():0), (((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			case CHARACTER:
				{
				setState(253);
				((ExprContext)_localctx).CHARACTER = match(CHARACTER);
				 
				        str := (((ExprContext)_localctx).CHARACTER!=null?((ExprContext)_localctx).CHARACTER.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).CHARACTER!=null?((ExprContext)_localctx).CHARACTER.getLine():0), (((ExprContext)_localctx).CHARACTER!=null?((ExprContext)_localctx).CHARACTER.getCharPositionInLine():0), str[1:len(str)-1], environment.CHARACTER) 
				    
				}
				break;
			case ID_VALIDO:
				{
				setState(255);
				((ExprContext)_localctx).ID_VALIDO = match(ID_VALIDO);

				        id := (((ExprContext)_localctx).ID_VALIDO!=null?((ExprContext)_localctx).ID_VALIDO.getText():null)
				        _localctx.e = sentencias.NewCallid((((ExprContext)_localctx).ID_VALIDO!=null?((ExprContext)_localctx).ID_VALIDO.getLine():0),(((ExprContext)_localctx).ID_VALIDO!=null?((ExprContext)_localctx).ID_VALIDO.getCharPositionInLine():0),id)
				    
				}
				break;
			case NULO:
				{
				setState(257);
				((ExprContext)_localctx).NULO = match(NULO);
				_localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NULO!=null?((ExprContext)_localctx).NULO.getLine():0), (((ExprContext)_localctx).NULO!=null?((ExprContext)_localctx).NULO.getCharPositionInLine():0), (((ExprContext)_localctx).NULO!=null?((ExprContext)_localctx).NULO.getText():null),environment.NULL)
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(303);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,21,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(301);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(261);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(262);
						((ExprContext)_localctx).op = match(MODULO);
						setState(263);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(18);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(266);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(267);
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
						setState(268);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(17);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(271);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(272);
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
						setState(273);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(16);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(276);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(277);
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
						setState(278);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(15);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(281);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(282);
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
						setState(283);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(14);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(286);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(287);
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
						setState(288);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(13);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(291);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(292);
						((ExprContext)_localctx).op = match(AND);
						setState(293);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(12);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 8:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(296);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(297);
						((ExprContext)_localctx).op = match(OR);
						setState(298);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(11);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(305);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,21,_ctx);
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
			setState(333);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(306);
				((SentenciaifelseContext)_localctx).IF = match(IF);
				setState(307);
				((SentenciaifelseContext)_localctx).expr = expr(0);
				setState(308);
				match(LLAVEIZQ);
				setState(309);
				((SentenciaifelseContext)_localctx).blockinterno = blockinterno();
				setState(310);
				match(LLAVEDER);
				 _localctx.myIfElse = sentencias.NewSentenciaIf((((SentenciaifelseContext)_localctx).IF!=null?((SentenciaifelseContext)_localctx).IF.getLine():0), (((SentenciaifelseContext)_localctx).IF!=null?((SentenciaifelseContext)_localctx).IF.getCharPositionInLine():0), ((SentenciaifelseContext)_localctx).expr.e, ((SentenciaifelseContext)_localctx).blockinterno.blkint)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(313);
				((SentenciaifelseContext)_localctx).IF = match(IF);
				setState(314);
				((SentenciaifelseContext)_localctx).expr = expr(0);
				setState(315);
				match(LLAVEIZQ);
				setState(316);
				((SentenciaifelseContext)_localctx).ifop = blockinterno();
				setState(317);
				match(LLAVEDER);
				setState(318);
				match(ELSE);
				setState(319);
				match(LLAVEIZQ);
				setState(320);
				((SentenciaifelseContext)_localctx).elseop = blockinterno();
				setState(321);
				match(LLAVEDER);
				 _localctx.myIfElse = sentencias.NewSentenciaIfElse((((SentenciaifelseContext)_localctx).IF!=null?((SentenciaifelseContext)_localctx).IF.getLine():0), (((SentenciaifelseContext)_localctx).IF!=null?((SentenciaifelseContext)_localctx).IF.getCharPositionInLine():0), ((SentenciaifelseContext)_localctx).expr.e, ((SentenciaifelseContext)_localctx).ifop.blkint , ((SentenciaifelseContext)_localctx).elseop.blkint)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(324);
				((SentenciaifelseContext)_localctx).IF = match(IF);
				setState(325);
				((SentenciaifelseContext)_localctx).expr = expr(0);
				setState(326);
				match(LLAVEIZQ);
				setState(327);
				((SentenciaifelseContext)_localctx).blockinterno = blockinterno();
				setState(328);
				match(LLAVEDER);
				setState(329);
				match(ELSE);
				setState(330);
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
			setState(335);
			((SwitchcontrolContext)_localctx).SWITCH = match(SWITCH);
			setState(336);
			((SwitchcontrolContext)_localctx).expr = expr(0);
			setState(337);
			match(LLAVEIZQ);
			setState(338);
			((SwitchcontrolContext)_localctx).blockcase = blockcase();
			setState(342);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(339);
				((SwitchcontrolContext)_localctx).DEFAULT = match(DEFAULT);
				setState(340);
				match(DOS_PUNTOS);
				setState(341);
				((SwitchcontrolContext)_localctx).blockinterno = blockinterno();
				}
			}

			setState(344);
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
			setState(348); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(347);
				((BlockcaseContext)_localctx).bloquecase = bloquecase();
				((BlockcaseContext)_localctx).blocas.add(((BlockcaseContext)_localctx).bloquecase);
				}
				}
				setState(350); 
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
			setState(354);
			((BloquecaseContext)_localctx).CASE = match(CASE);
			setState(355);
			((BloquecaseContext)_localctx).expr = expr(0);
			setState(356);
			match(DOS_PUNTOS);
			setState(357);
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
			setState(360);
			((WhilecontrolContext)_localctx).WHILE = match(WHILE);
			setState(361);
			((WhilecontrolContext)_localctx).expr = expr(0);
			setState(362);
			match(LLAVEIZQ);
			setState(363);
			((WhilecontrolContext)_localctx).blockinterno = blockinterno();
			setState(364);
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
			setState(396);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(367);
				((ForcontrolContext)_localctx).FOR = match(FOR);
				setState(368);
				((ForcontrolContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(369);
				match(IN);
				setState(370);
				((ForcontrolContext)_localctx).left = expr(0);
				setState(371);
				match(RANGO);
				setState(372);
				((ForcontrolContext)_localctx).right = expr(0);
				setState(373);
				match(LLAVEIZQ);
				setState(374);
				((ForcontrolContext)_localctx).blockinterno = blockinterno();
				setState(375);
				match(LLAVEDER);
				 _localctx.forct = sentencias.NewSentenciaForRango((((ForcontrolContext)_localctx).FOR!=null?((ForcontrolContext)_localctx).FOR.getLine():0), (((ForcontrolContext)_localctx).FOR!=null?((ForcontrolContext)_localctx).FOR.getCharPositionInLine():0), (((ForcontrolContext)_localctx).ID_VALIDO!=null?((ForcontrolContext)_localctx).ID_VALIDO.getText():null), ((ForcontrolContext)_localctx).left.e, ((ForcontrolContext)_localctx).right.e,((ForcontrolContext)_localctx).blockinterno.blkint)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(378);
				((ForcontrolContext)_localctx).FOR = match(FOR);
				setState(379);
				((ForcontrolContext)_localctx).op1 = match(ID_VALIDO);
				setState(380);
				match(IN);
				setState(381);
				((ForcontrolContext)_localctx).op2 = match(ID_VALIDO);
				setState(382);
				match(LLAVEIZQ);
				setState(383);
				((ForcontrolContext)_localctx).blockinterno = blockinterno();
				setState(384);
				match(LLAVEDER);
				 _localctx.forct = sentencias.NewSentenciaForId((((ForcontrolContext)_localctx).FOR!=null?((ForcontrolContext)_localctx).FOR.getLine():0), (((ForcontrolContext)_localctx).FOR!=null?((ForcontrolContext)_localctx).FOR.getCharPositionInLine():0), (((ForcontrolContext)_localctx).op1!=null?((ForcontrolContext)_localctx).op1.getText():null), (((ForcontrolContext)_localctx).op2!=null?((ForcontrolContext)_localctx).op2.getText():null), ((ForcontrolContext)_localctx).blockinterno.blkint)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(387);
				((ForcontrolContext)_localctx).FOR = match(FOR);
				setState(388);
				((ForcontrolContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(389);
				match(IN);
				setState(390);
				((ForcontrolContext)_localctx).expr = expr(0);
				setState(391);
				match(LLAVEIZQ);
				setState(392);
				((ForcontrolContext)_localctx).blockinterno = blockinterno();
				setState(393);
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
			setState(398);
			((GuardcontrolContext)_localctx).GUARD = match(GUARD);
			setState(399);
			((GuardcontrolContext)_localctx).expr = expr(0);
			setState(400);
			match(ELSE);
			setState(401);
			match(LLAVEIZQ);
			setState(402);
			((GuardcontrolContext)_localctx).blockinterno = blockinterno();
			setState(403);
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
			setState(406);
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
			setState(409);
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
			setState(418);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(412);
				((RetornosContext)_localctx).RETURN = match(RETURN);
				setState(413);
				((RetornosContext)_localctx).op = expr(0);

				    ((RetornosContext)_localctx).rect =  sentencias.NewTransferenciaReturnExp((((RetornosContext)_localctx).RETURN!=null?((RetornosContext)_localctx).RETURN.getLine():0), (((RetornosContext)_localctx).RETURN!=null?((RetornosContext)_localctx).RETURN.getCharPositionInLine():0), ((RetornosContext)_localctx).op.e);

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(416);
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
		enterRule(_localctx, 40, RULE_listaexpresions);

		    _localctx.blkparf = []interface{}{}
		    var listInt []IListaexpresionContext

		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(421); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(420);
				((ListaexpresionsContext)_localctx).listaexpresion = listaexpresion();
				((ListaexpresionsContext)_localctx).funpar.add(((ListaexpresionsContext)_localctx).listaexpresion);
				}
				}
				setState(423); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( ((((_la - 6)) & ~0x3f) == 0 && ((1L << (_la - 6)) & ((1L << (TRU - 6)) | (1L << (FAL - 6)) | (1L << (NULO - 6)) | (1L << (NUMBER - 6)) | (1L << (CADENA - 6)) | (1L << (ID_VALIDO - 6)) | (1L << (CHARACTER - 6)) | (1L << (PARIZQ - 6)) | (1L << (NOT - 6)) | (1L << (SUB - 6)) | (1L << (COMA - 6)))) != 0) );

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
		enterRule(_localctx, 42, RULE_listaexpresion);
		try {
			setState(434);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case COMA:
				enterOuterAlt(_localctx, 1);
				{
				setState(427);
				((ListaexpresionContext)_localctx).COMA = match(COMA);
				setState(428);
				((ListaexpresionContext)_localctx).expr = expr(0);

				    _localctx.funpar = datoscompuestos.NewArregloParametros((((ListaexpresionContext)_localctx).COMA!=null?((ListaexpresionContext)_localctx).COMA.getLine():0) ,(((ListaexpresionContext)_localctx).COMA!=null?((ListaexpresionContext)_localctx).COMA.getCharPositionInLine():0), ((ListaexpresionContext)_localctx).expr.e)

				}
				break;
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
				setState(431);
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
		enterRule(_localctx, 44, RULE_printstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(436);
			((PrintstmtContext)_localctx).PRINT = match(PRINT);
			setState(437);
			match(PARIZQ);
			setState(438);
			((PrintstmtContext)_localctx).listaexpresions = listaexpresions();
			setState(439);
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 9:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 17);
		case 1:
			return precpred(_ctx, 16);
		case 2:
			return precpred(_ctx, 15);
		case 3:
			return precpred(_ctx, 14);
		case 4:
			return precpred(_ctx, 13);
		case 5:
			return precpred(_ctx, 12);
		case 6:
			return precpred(_ctx, 11);
		case 7:
			return precpred(_ctx, 10);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3M\u01bd\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\3\2\3\2\3"+
		"\2\3\2\3\3\6\3\66\n\3\r\3\16\3\67\3\3\3\3\3\4\3\4\5\4>\n\4\3\4\3\4\3\4"+
		"\3\4\5\4D\n\4\3\4\3\4\3\4\3\4\5\4J\n\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4_\n\4\3\4\3\4\5\4c\n\4"+
		"\3\5\6\5f\n\5\r\5\16\5g\3\5\3\5\3\6\3\6\5\6n\n\6\3\6\3\6\3\6\3\6\5\6t"+
		"\n\6\3\6\3\6\3\6\3\6\5\6z\n\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6\u008f\n\6\3\6\3\6\3\6\3\6\5\6"+
		"\u0095\n\6\3\6\3\6\3\6\3\6\5\6\u009b\n\6\3\6\3\6\3\6\3\6\5\6\u00a1\n\6"+
		"\3\6\3\6\5\6\u00a5\n\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7"+
		"\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7\u00bc\n\7\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\5\b\u00cc\n\b\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t\u00dd\n\t\3\n\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\5\n\u00e9\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3"+
		"\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3"+
		"\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13\u0106\n\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\7\13\u0130\n\13\f\13\16\13\u0133"+
		"\13\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\5\f\u0150\n\f\3\r\3\r\3\r"+
		"\3\r\3\r\3\r\3\r\5\r\u0159\n\r\3\r\3\r\3\r\3\16\6\16\u015f\n\16\r\16\16"+
		"\16\u0160\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\5\21\u018f\n\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\23\3\23\3\23\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\5\25"+
		"\u01a5\n\25\3\26\6\26\u01a8\n\26\r\26\16\26\u01a9\3\26\3\26\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\5\27\u01b5\n\27\3\30\3\30\3\30\3\30\3\30\3\30"+
		"\3\30\2\3\24\31\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\2\7\3"+
		"\2<=\3\2>?\4\2\67\6799\4\288::\3\2\62\63\2\u01e8\2\60\3\2\2\2\4\65\3\2"+
		"\2\2\6b\3\2\2\2\be\3\2\2\2\n\u00a4\3\2\2\2\f\u00bb\3\2\2\2\16\u00cb\3"+
		"\2\2\2\20\u00dc\3\2\2\2\22\u00e8\3\2\2\2\24\u0105\3\2\2\2\26\u014f\3\2"+
		"\2\2\30\u0151\3\2\2\2\32\u015e\3\2\2\2\34\u0164\3\2\2\2\36\u016a\3\2\2"+
		"\2 \u018e\3\2\2\2\"\u0190\3\2\2\2$\u0198\3\2\2\2&\u019b\3\2\2\2(\u01a4"+
		"\3\2\2\2*\u01a7\3\2\2\2,\u01b4\3\2\2\2.\u01b6\3\2\2\2\60\61\5\4\3\2\61"+
		"\62\7\2\2\3\62\63\b\2\1\2\63\3\3\2\2\2\64\66\5\6\4\2\65\64\3\2\2\2\66"+
		"\67\3\2\2\2\67\65\3\2\2\2\678\3\2\2\289\3\2\2\29:\b\3\1\2:\5\3\2\2\2;"+
		"=\5\f\7\2<>\7.\2\2=<\3\2\2\2=>\3\2\2\2>?\3\2\2\2?@\b\4\1\2@c\3\2\2\2A"+
		"C\5\16\b\2BD\7.\2\2CB\3\2\2\2CD\3\2\2\2DE\3\2\2\2EF\b\4\1\2Fc\3\2\2\2"+
		"GI\5\20\t\2HJ\7.\2\2IH\3\2\2\2IJ\3\2\2\2JK\3\2\2\2KL\b\4\1\2Lc\3\2\2\2"+
		"MN\5\26\f\2NO\b\4\1\2Oc\3\2\2\2PQ\5\30\r\2QR\b\4\1\2Rc\3\2\2\2ST\5\36"+
		"\20\2TU\b\4\1\2Uc\3\2\2\2VW\5 \21\2WX\b\4\1\2Xc\3\2\2\2YZ\5\"\22\2Z[\b"+
		"\4\1\2[c\3\2\2\2\\^\5.\30\2]_\7.\2\2^]\3\2\2\2^_\3\2\2\2_`\3\2\2\2`a\b"+
		"\4\1\2ac\3\2\2\2b;\3\2\2\2bA\3\2\2\2bG\3\2\2\2bM\3\2\2\2bP\3\2\2\2bS\3"+
		"\2\2\2bV\3\2\2\2bY\3\2\2\2b\\\3\2\2\2c\7\3\2\2\2df\5\n\6\2ed\3\2\2\2f"+
		"g\3\2\2\2ge\3\2\2\2gh\3\2\2\2hi\3\2\2\2ij\b\5\1\2j\t\3\2\2\2km\5\f\7\2"+
		"ln\7.\2\2ml\3\2\2\2mn\3\2\2\2no\3\2\2\2op\b\6\1\2p\u00a5\3\2\2\2qs\5\16"+
		"\b\2rt\7.\2\2sr\3\2\2\2st\3\2\2\2tu\3\2\2\2uv\b\6\1\2v\u00a5\3\2\2\2w"+
		"y\5\20\t\2xz\7.\2\2yx\3\2\2\2yz\3\2\2\2z{\3\2\2\2{|\b\6\1\2|\u00a5\3\2"+
		"\2\2}~\5\26\f\2~\177\b\6\1\2\177\u00a5\3\2\2\2\u0080\u0081\5\30\r\2\u0081"+
		"\u0082\b\6\1\2\u0082\u00a5\3\2\2\2\u0083\u0084\5\36\20\2\u0084\u0085\b"+
		"\6\1\2\u0085\u00a5\3\2\2\2\u0086\u0087\5 \21\2\u0087\u0088\b\6\1\2\u0088"+
		"\u00a5\3\2\2\2\u0089\u008a\5\"\22\2\u008a\u008b\b\6\1\2\u008b\u00a5\3"+
		"\2\2\2\u008c\u008e\5$\23\2\u008d\u008f\7.\2\2\u008e\u008d\3\2\2\2\u008e"+
		"\u008f\3\2\2\2\u008f\u0090\3\2\2\2\u0090\u0091\b\6\1\2\u0091\u00a5\3\2"+
		"\2\2\u0092\u0094\5&\24\2\u0093\u0095\7.\2\2\u0094\u0093\3\2\2\2\u0094"+
		"\u0095\3\2\2\2\u0095\u0096\3\2\2\2\u0096\u0097\b\6\1\2\u0097\u00a5\3\2"+
		"\2\2\u0098\u009a\5(\25\2\u0099\u009b\7.\2\2\u009a\u0099\3\2\2\2\u009a"+
		"\u009b\3\2\2\2\u009b\u009c\3\2\2\2\u009c\u009d\b\6\1\2\u009d\u00a5\3\2"+
		"\2\2\u009e\u00a0\5.\30\2\u009f\u00a1\7.\2\2\u00a0\u009f\3\2\2\2\u00a0"+
		"\u00a1\3\2\2\2\u00a1\u00a2\3\2\2\2\u00a2\u00a3\b\6\1\2\u00a3\u00a5\3\2"+
		"\2\2\u00a4k\3\2\2\2\u00a4q\3\2\2\2\u00a4w\3\2\2\2\u00a4}\3\2\2\2\u00a4"+
		"\u0080\3\2\2\2\u00a4\u0083\3\2\2\2\u00a4\u0086\3\2\2\2\u00a4\u0089\3\2"+
		"\2\2\u00a4\u008c\3\2\2\2\u00a4\u0092\3\2\2\2\u00a4\u0098\3\2\2\2\u00a4"+
		"\u009e\3\2\2\2\u00a5\13\3\2\2\2\u00a6\u00a7\7\n\2\2\u00a7\u00a8\7)\2\2"+
		"\u00a8\u00a9\7-\2\2\u00a9\u00aa\5\22\n\2\u00aa\u00ab\7,\2\2\u00ab\u00ac"+
		"\5\24\13\2\u00ac\u00ad\b\7\1\2\u00ad\u00bc\3\2\2\2\u00ae\u00af\7\n\2\2"+
		"\u00af\u00b0\7)\2\2\u00b0\u00b1\7,\2\2\u00b1\u00b2\5\24\13\2\u00b2\u00b3"+
		"\b\7\1\2\u00b3\u00bc\3\2\2\2\u00b4\u00b5\7\n\2\2\u00b5\u00b6\7)\2\2\u00b6"+
		"\u00b7\7-\2\2\u00b7\u00b8\5\22\n\2\u00b8\u00b9\7/\2\2\u00b9\u00ba\b\7"+
		"\1\2\u00ba\u00bc\3\2\2\2\u00bb\u00a6\3\2\2\2\u00bb\u00ae\3\2\2\2\u00bb"+
		"\u00b4\3\2\2\2\u00bc\r\3\2\2\2\u00bd\u00be\7\13\2\2\u00be\u00bf\7)\2\2"+
		"\u00bf\u00c0\7-\2\2\u00c0\u00c1\5\22\n\2\u00c1\u00c2\7,\2\2\u00c2\u00c3"+
		"\5\24\13\2\u00c3\u00c4\b\b\1\2\u00c4\u00cc\3\2\2\2\u00c5\u00c6\7\13\2"+
		"\2\u00c6\u00c7\7)\2\2\u00c7\u00c8\7,\2\2\u00c8\u00c9\5\24\13\2\u00c9\u00ca"+
		"\b\b\1\2\u00ca\u00cc\3\2\2\2\u00cb\u00bd\3\2\2\2\u00cb\u00c5\3\2\2\2\u00cc"+
		"\17\3\2\2\2\u00cd\u00ce\7)\2\2\u00ce\u00cf\7,\2\2\u00cf\u00d0\5\24\13"+
		"\2\u00d0\u00d1\b\t\1\2\u00d1\u00dd\3\2\2\2\u00d2\u00d3\7)\2\2\u00d3\u00d4"+
		"\7@\2\2\u00d4\u00d5\5\24\13\2\u00d5\u00d6\b\t\1\2\u00d6\u00dd\3\2\2\2"+
		"\u00d7\u00d8\7)\2\2\u00d8\u00d9\7A\2\2\u00d9\u00da\5\24\13\2\u00da\u00db"+
		"\b\t\1\2\u00db\u00dd\3\2\2\2\u00dc\u00cd\3\2\2\2\u00dc\u00d2\3\2\2\2\u00dc"+
		"\u00d7\3\2\2\2\u00dd\21\3\2\2\2\u00de\u00df\7\3\2\2\u00df\u00e9\b\n\1"+
		"\2\u00e0\u00e1\7\4\2\2\u00e1\u00e9\b\n\1\2\u00e2\u00e3\7\5\2\2\u00e3\u00e9"+
		"\b\n\1\2\u00e4\u00e5\7\6\2\2\u00e5\u00e9\b\n\1\2\u00e6\u00e7\7\7\2\2\u00e7"+
		"\u00e9\b\n\1\2\u00e8\u00de\3\2\2\2\u00e8\u00e0\3\2\2\2\u00e8\u00e2\3\2"+
		"\2\2\u00e8\u00e4\3\2\2\2\u00e8\u00e6\3\2\2\2\u00e9\23\3\2\2\2\u00ea\u00eb"+
		"\b\13\1\2\u00eb\u00ec\7\64\2\2\u00ec\u00ed\5\24\13\24\u00ed\u00ee\b\13"+
		"\1\2\u00ee\u0106\3\2\2\2\u00ef\u00f0\7\60\2\2\u00f0\u00f1\5\24\13\2\u00f1"+
		"\u00f2\7\61\2\2\u00f2\u00f3\b\13\1\2\u00f3\u0106\3\2\2\2\u00f4\u00f5\7"+
		"?\2\2\u00f5\u00f6\7\'\2\2\u00f6\u0106\b\13\1\2\u00f7\u00f8\7\'\2\2\u00f8"+
		"\u0106\b\13\1\2\u00f9\u00fa\7(\2\2\u00fa\u0106\b\13\1\2\u00fb\u00fc\7"+
		"\b\2\2\u00fc\u0106\b\13\1\2\u00fd\u00fe\7\t\2\2\u00fe\u0106\b\13\1\2\u00ff"+
		"\u0100\7*\2\2\u0100\u0106\b\13\1\2\u0101\u0102\7)\2\2\u0102\u0106\b\13"+
		"\1\2\u0103\u0104\7\f\2\2\u0104\u0106\b\13\1\2\u0105\u00ea\3\2\2\2\u0105"+
		"\u00ef\3\2\2\2\u0105\u00f4\3\2\2\2\u0105\u00f7\3\2\2\2\u0105\u00f9\3\2"+
		"\2\2\u0105\u00fb\3\2\2\2\u0105\u00fd\3\2\2\2\u0105\u00ff\3\2\2\2\u0105"+
		"\u0101\3\2\2\2\u0105\u0103\3\2\2\2\u0106\u0131\3\2\2\2\u0107\u0108\f\23"+
		"\2\2\u0108\u0109\7;\2\2\u0109\u010a\5\24\13\24\u010a\u010b\b\13\1\2\u010b"+
		"\u0130\3\2\2\2\u010c\u010d\f\22\2\2\u010d\u010e\t\2\2\2\u010e\u010f\5"+
		"\24\13\23\u010f\u0110\b\13\1\2\u0110\u0130\3\2\2\2\u0111\u0112\f\21\2"+
		"\2\u0112\u0113\t\3\2\2\u0113\u0114\5\24\13\22\u0114\u0115\b\13\1\2\u0115"+
		"\u0130\3\2\2\2\u0116\u0117\f\20\2\2\u0117\u0118\t\4\2\2\u0118\u0119\5"+
		"\24\13\21\u0119\u011a\b\13\1\2\u011a\u0130\3\2\2\2\u011b\u011c\f\17\2"+
		"\2\u011c\u011d\t\5\2\2\u011d\u011e\5\24\13\20\u011e\u011f\b\13\1\2\u011f"+
		"\u0130\3\2\2\2\u0120\u0121\f\16\2\2\u0121\u0122\t\6\2\2\u0122\u0123\5"+
		"\24\13\17\u0123\u0124\b\13\1\2\u0124\u0130\3\2\2\2\u0125\u0126\f\r\2\2"+
		"\u0126\u0127\7\66\2\2\u0127\u0128\5\24\13\16\u0128\u0129\b\13\1\2\u0129"+
		"\u0130\3\2\2\2\u012a\u012b\f\f\2\2\u012b\u012c\7\65\2\2\u012c\u012d\5"+
		"\24\13\r\u012d\u012e\b\13\1\2\u012e\u0130\3\2\2\2\u012f\u0107\3\2\2\2"+
		"\u012f\u010c\3\2\2\2\u012f\u0111\3\2\2\2\u012f\u0116\3\2\2\2\u012f\u011b"+
		"\3\2\2\2\u012f\u0120\3\2\2\2\u012f\u0125\3\2\2\2\u012f\u012a\3\2\2\2\u0130"+
		"\u0133\3\2\2\2\u0131\u012f\3\2\2\2\u0131\u0132\3\2\2\2\u0132\25\3\2\2"+
		"\2\u0133\u0131\3\2\2\2\u0134\u0135\7\r\2\2\u0135\u0136\5\24\13\2\u0136"+
		"\u0137\7B\2\2\u0137\u0138\5\b\5\2\u0138\u0139\7C\2\2\u0139\u013a\b\f\1"+
		"\2\u013a\u0150\3\2\2\2\u013b\u013c\7\r\2\2\u013c\u013d\5\24\13\2\u013d"+
		"\u013e\7B\2\2\u013e\u013f\5\b\5\2\u013f\u0140\7C\2\2\u0140\u0141\7\16"+
		"\2\2\u0141\u0142\7B\2\2\u0142\u0143\5\b\5\2\u0143\u0144\7C\2\2\u0144\u0145"+
		"\b\f\1\2\u0145\u0150\3\2\2\2\u0146\u0147\7\r\2\2\u0147\u0148\5\24\13\2"+
		"\u0148\u0149\7B\2\2\u0149\u014a\5\b\5\2\u014a\u014b\7C\2\2\u014b\u014c"+
		"\7\16\2\2\u014c\u014d\5\26\f\2\u014d\u014e\b\f\1\2\u014e\u0150\3\2\2\2"+
		"\u014f\u0134\3\2\2\2\u014f\u013b\3\2\2\2\u014f\u0146\3\2\2\2\u0150\27"+
		"\3\2\2\2\u0151\u0152\7\17\2\2\u0152\u0153\5\24\13\2\u0153\u0154\7B\2\2"+
		"\u0154\u0158\5\32\16\2\u0155\u0156\7\21\2\2\u0156\u0157\7-\2\2\u0157\u0159"+
		"\5\b\5\2\u0158\u0155\3\2\2\2\u0158\u0159\3\2\2\2\u0159\u015a\3\2\2\2\u015a"+
		"\u015b\7C\2\2\u015b\u015c\b\r\1\2\u015c\31\3\2\2\2\u015d\u015f\5\34\17"+
		"\2\u015e\u015d\3\2\2\2\u015f\u0160\3\2\2\2\u0160\u015e\3\2\2\2\u0160\u0161"+
		"\3\2\2\2\u0161\u0162\3\2\2\2\u0162\u0163\b\16\1\2\u0163\33\3\2\2\2\u0164"+
		"\u0165\7\20\2\2\u0165\u0166\5\24\13\2\u0166\u0167\7-\2\2\u0167\u0168\5"+
		"\b\5\2\u0168\u0169\b\17\1\2\u0169\35\3\2\2\2\u016a\u016b\7\27\2\2\u016b"+
		"\u016c\5\24\13\2\u016c\u016d\7B\2\2\u016d\u016e\5\b\5\2\u016e\u016f\7"+
		"C\2\2\u016f\u0170\b\20\1\2\u0170\37\3\2\2\2\u0171\u0172\7\24\2\2\u0172"+
		"\u0173\7)\2\2\u0173\u0174\7\25\2\2\u0174\u0175\5\24\13\2\u0175\u0176\7"+
		"\26\2\2\u0176\u0177\5\24\13\2\u0177\u0178\7B\2\2\u0178\u0179\5\b\5\2\u0179"+
		"\u017a\7C\2\2\u017a\u017b\b\21\1\2\u017b\u018f\3\2\2\2\u017c\u017d\7\24"+
		"\2\2\u017d\u017e\7)\2\2\u017e\u017f\7\25\2\2\u017f\u0180\7)\2\2\u0180"+
		"\u0181\7B\2\2\u0181\u0182\5\b\5\2\u0182\u0183\7C\2\2\u0183\u0184\b\21"+
		"\1\2\u0184\u018f\3\2\2\2\u0185\u0186\7\24\2\2\u0186\u0187\7)\2\2\u0187"+
		"\u0188\7\25\2\2\u0188\u0189\5\24\13\2\u0189\u018a\7B\2\2\u018a\u018b\5"+
		"\b\5\2\u018b\u018c\7C\2\2\u018c\u018d\b\21\1\2\u018d\u018f\3\2\2\2\u018e"+
		"\u0171\3\2\2\2\u018e\u017c\3\2\2\2\u018e\u0185\3\2\2\2\u018f!\3\2\2\2"+
		"\u0190\u0191\7\30\2\2\u0191\u0192\5\24\13\2\u0192\u0193\7\16\2\2\u0193"+
		"\u0194\7B\2\2\u0194\u0195\5\b\5\2\u0195\u0196\7C\2\2\u0196\u0197\b\22"+
		"\1\2\u0197#\3\2\2\2\u0198\u0199\7\23\2\2\u0199\u019a\b\23\1\2\u019a%\3"+
		"\2\2\2\u019b\u019c\7\22\2\2\u019c\u019d\b\24\1\2\u019d\'\3\2\2\2\u019e"+
		"\u019f\7\31\2\2\u019f\u01a0\5\24\13\2\u01a0\u01a1\b\25\1\2\u01a1\u01a5"+
		"\3\2\2\2\u01a2\u01a3\7\31\2\2\u01a3\u01a5\b\25\1\2\u01a4\u019e\3\2\2\2"+
		"\u01a4\u01a2\3\2\2\2\u01a5)\3\2\2\2\u01a6\u01a8\5,\27\2\u01a7\u01a6\3"+
		"\2\2\2\u01a8\u01a9\3\2\2\2\u01a9\u01a7\3\2\2\2\u01a9\u01aa\3\2\2\2\u01aa"+
		"\u01ab\3\2\2\2\u01ab\u01ac\b\26\1\2\u01ac+\3\2\2\2\u01ad\u01ae\7E\2\2"+
		"\u01ae\u01af\5\24\13\2\u01af\u01b0\b\27\1\2\u01b0\u01b5\3\2\2\2\u01b1"+
		"\u01b2\5\24\13\2\u01b2\u01b3\b\27\1\2\u01b3\u01b5\3\2\2\2\u01b4\u01ad"+
		"\3\2\2\2\u01b4\u01b1\3\2\2\2\u01b5-\3\2\2\2\u01b6\u01b7\7\33\2\2\u01b7"+
		"\u01b8\7\60\2\2\u01b8\u01b9\5*\26\2\u01b9\u01ba\7\61\2\2\u01ba\u01bb\b"+
		"\30\1\2\u01bb/\3\2\2\2\37\67=CI^bgmsy\u008e\u0094\u009a\u00a0\u00a4\u00bb"+
		"\u00cb\u00dc\u00e8\u0105\u012f\u0131\u014f\u0158\u0160\u018e\u01a4\u01a9"+
		"\u01b4";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}