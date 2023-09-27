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
		RULE_whilecontrol = 11, RULE_continuee = 12, RULE_breakk = 13, RULE_retornos = 14, 
		RULE_listaexpresions = 15, RULE_listaexpresion = 16, RULE_printstmt = 17;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "blockinterno", "instructionint", "declavarible", 
			"declaconstante", "asignacionvariable", "tipodato", "expr", "sentenciaifelse", 
			"whilecontrol", "continuee", "breakk", "retornos", "listaexpresions", 
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
			setState(36);
			((SContext)_localctx).block = block();
			setState(37);
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
			setState(41); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(40);
				((BlockContext)_localctx).instruction = instruction();
				((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
				}
				}
				setState(43); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << VAR) | (1L << LET) | (1L << IF) | (1L << WHILE) | (1L << PRINT) | (1L << ID_VALIDO))) != 0) );

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
		public WhilecontrolContext whilecontrol;
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
		public WhilecontrolContext whilecontrol() {
			return getRuleContext(WhilecontrolContext.class,0);
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
			setState(77);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VAR:
				enterOuterAlt(_localctx, 1);
				{
				setState(47);
				((InstructionContext)_localctx).declavarible = declavarible();
				setState(49);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(48);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).declavarible.decvbl
				}
				break;
			case LET:
				enterOuterAlt(_localctx, 2);
				{
				setState(53);
				((InstructionContext)_localctx).declaconstante = declaconstante();
				setState(55);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(54);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).declaconstante.deccon
				}
				break;
			case ID_VALIDO:
				enterOuterAlt(_localctx, 3);
				{
				setState(59);
				((InstructionContext)_localctx).asignacionvariable = asignacionvariable();
				setState(61);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(60);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).asignacionvariable.asgvbl
				}
				break;
			case IF:
				enterOuterAlt(_localctx, 4);
				{
				setState(65);
				((InstructionContext)_localctx).sentenciaifelse = sentenciaifelse();
				 _localctx.inst = ((InstructionContext)_localctx).sentenciaifelse.myIfElse
				}
				break;
			case WHILE:
				enterOuterAlt(_localctx, 5);
				{
				setState(68);
				((InstructionContext)_localctx).whilecontrol = whilecontrol();
				 _localctx.inst = ((InstructionContext)_localctx).whilecontrol.whict
				}
				break;
			case PRINT:
				enterOuterAlt(_localctx, 6);
				{
				setState(71);
				((InstructionContext)_localctx).printstmt = printstmt();
				setState(73);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(72);
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
			setState(80); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(79);
				((BlockinternoContext)_localctx).instructionint = instructionint();
				((BlockinternoContext)_localctx).insint.add(((BlockinternoContext)_localctx).instructionint);
				}
				}
				setState(82); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << VAR) | (1L << LET) | (1L << IF) | (1L << BREAK) | (1L << CONTINUE) | (1L << WHILE) | (1L << RETURN) | (1L << PRINT) | (1L << ID_VALIDO))) != 0) );

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
		public WhilecontrolContext whilecontrol;
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
		public WhilecontrolContext whilecontrol() {
			return getRuleContext(WhilecontrolContext.class,0);
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
			setState(134);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VAR:
				enterOuterAlt(_localctx, 1);
				{
				setState(86);
				((InstructionintContext)_localctx).declavarible = declavarible();
				setState(88);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(87);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).declavarible.decvbl
				}
				break;
			case LET:
				enterOuterAlt(_localctx, 2);
				{
				setState(92);
				((InstructionintContext)_localctx).declaconstante = declaconstante();
				setState(94);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(93);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).declaconstante.deccon
				}
				break;
			case ID_VALIDO:
				enterOuterAlt(_localctx, 3);
				{
				setState(98);
				((InstructionintContext)_localctx).asignacionvariable = asignacionvariable();
				setState(100);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(99);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).asignacionvariable.asgvbl
				}
				break;
			case IF:
				enterOuterAlt(_localctx, 4);
				{
				setState(104);
				((InstructionintContext)_localctx).sentenciaifelse = sentenciaifelse();
				 _localctx.insint = ((InstructionintContext)_localctx).sentenciaifelse.myIfElse
				}
				break;
			case WHILE:
				enterOuterAlt(_localctx, 5);
				{
				setState(107);
				((InstructionintContext)_localctx).whilecontrol = whilecontrol();
				 _localctx.insint = ((InstructionintContext)_localctx).whilecontrol.whict
				}
				break;
			case CONTINUE:
				enterOuterAlt(_localctx, 6);
				{
				setState(110);
				((InstructionintContext)_localctx).continuee = continuee();
				setState(112);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(111);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).continuee.coct
				}
				break;
			case BREAK:
				enterOuterAlt(_localctx, 7);
				{
				setState(116);
				((InstructionintContext)_localctx).breakk = breakk();
				setState(118);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(117);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).breakk.brkct
				}
				break;
			case RETURN:
				enterOuterAlt(_localctx, 8);
				{
				setState(122);
				((InstructionintContext)_localctx).retornos = retornos();
				setState(124);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(123);
					match(PUNTOCOMA);
					}
				}

				 _localctx.insint = ((InstructionintContext)_localctx).retornos.rect 
				}
				break;
			case PRINT:
				enterOuterAlt(_localctx, 9);
				{
				setState(128);
				((InstructionintContext)_localctx).printstmt = printstmt();
				setState(130);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(129);
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
			setState(157);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(136);
				((DeclavaribleContext)_localctx).VAR = match(VAR);
				setState(137);
				((DeclavaribleContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(138);
				match(DOS_PUNTOS);
				setState(139);
				((DeclavaribleContext)_localctx).tipodato = tipodato();
				setState(140);
				match(IG);
				setState(141);
				((DeclavaribleContext)_localctx).expr = expr(0);
				 _localctx.decvbl = datosprimitivos.NewVariableDeclaration((((DeclavaribleContext)_localctx).VAR!=null?((DeclavaribleContext)_localctx).VAR.getLine():0), (((DeclavaribleContext)_localctx).VAR!=null?((DeclavaribleContext)_localctx).VAR.getCharPositionInLine():0), (((DeclavaribleContext)_localctx).ID_VALIDO!=null?((DeclavaribleContext)_localctx).ID_VALIDO.getText():null), ((DeclavaribleContext)_localctx).tipodato.tipo, ((DeclavaribleContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(144);
				((DeclavaribleContext)_localctx).VAR = match(VAR);
				setState(145);
				((DeclavaribleContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(146);
				match(IG);
				setState(147);
				((DeclavaribleContext)_localctx).expr = expr(0);
				_localctx.decvbl = datosprimitivos.NewVariableDeclaracionSinTipo((((DeclavaribleContext)_localctx).VAR!=null?((DeclavaribleContext)_localctx).VAR.getLine():0), (((DeclavaribleContext)_localctx).VAR!=null?((DeclavaribleContext)_localctx).VAR.getCharPositionInLine():0), (((DeclavaribleContext)_localctx).ID_VALIDO!=null?((DeclavaribleContext)_localctx).ID_VALIDO.getText():null), ((DeclavaribleContext)_localctx).expr.e)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(150);
				((DeclavaribleContext)_localctx).VAR = match(VAR);
				setState(151);
				((DeclavaribleContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(152);
				match(DOS_PUNTOS);
				setState(153);
				((DeclavaribleContext)_localctx).tipodato = tipodato();
				setState(154);
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
			setState(173);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(159);
				((DeclaconstanteContext)_localctx).LET = match(LET);
				setState(160);
				((DeclaconstanteContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(161);
				match(DOS_PUNTOS);
				setState(162);
				((DeclaconstanteContext)_localctx).tipodato = tipodato();
				setState(163);
				match(IG);
				setState(164);
				((DeclaconstanteContext)_localctx).expr = expr(0);
				_localctx.deccon = datosprimitivos.NewConstanteDeclaration((((DeclaconstanteContext)_localctx).LET!=null?((DeclaconstanteContext)_localctx).LET.getLine():0), (((DeclaconstanteContext)_localctx).LET!=null?((DeclaconstanteContext)_localctx).LET.getCharPositionInLine():0), (((DeclaconstanteContext)_localctx).ID_VALIDO!=null?((DeclaconstanteContext)_localctx).ID_VALIDO.getText():null), ((DeclaconstanteContext)_localctx).tipodato.tipo, ((DeclaconstanteContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(167);
				((DeclaconstanteContext)_localctx).LET = match(LET);
				setState(168);
				((DeclaconstanteContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(169);
				match(IG);
				setState(170);
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
			setState(190);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(175);
				((AsignacionvariableContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(176);
				match(IG);
				setState(177);
				((AsignacionvariableContext)_localctx).expr = expr(0);
				 _localctx.asgvbl = sentencias.NewAsignacionVariable((((AsignacionvariableContext)_localctx).ID_VALIDO!=null?((AsignacionvariableContext)_localctx).ID_VALIDO.getLine():0), (((AsignacionvariableContext)_localctx).ID_VALIDO!=null?((AsignacionvariableContext)_localctx).ID_VALIDO.getCharPositionInLine():0), (((AsignacionvariableContext)_localctx).ID_VALIDO!=null?((AsignacionvariableContext)_localctx).ID_VALIDO.getText():null), ((AsignacionvariableContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(180);
				((AsignacionvariableContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(181);
				match(SUMA);
				setState(182);
				((AsignacionvariableContext)_localctx).expr = expr(0);
				 _localctx.asgvbl = sentencias.NewAsignacionSuma((((AsignacionvariableContext)_localctx).ID_VALIDO!=null?((AsignacionvariableContext)_localctx).ID_VALIDO.getLine():0), (((AsignacionvariableContext)_localctx).ID_VALIDO!=null?((AsignacionvariableContext)_localctx).ID_VALIDO.getCharPositionInLine():0), (((AsignacionvariableContext)_localctx).ID_VALIDO!=null?((AsignacionvariableContext)_localctx).ID_VALIDO.getText():null), ((AsignacionvariableContext)_localctx).expr.e)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(185);
				((AsignacionvariableContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(186);
				match(RESTA);
				setState(187);
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
			setState(202);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(192);
				match(INT);
				 _localctx.tipo = environment.INTEGER 
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(194);
				match(FLOAT);
				 _localctx.tipo = environment.FLOAT 
				}
				break;
			case STRING:
				enterOuterAlt(_localctx, 3);
				{
				setState(196);
				match(STRING);
				 _localctx.tipo = environment.STRING 
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(198);
				match(BOOL);
				 _localctx.tipo = environment.BOOLEAN 
				}
				break;
			case CHARACT:
				enterOuterAlt(_localctx, 5);
				{
				setState(200);
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
			setState(231);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NOT:
				{
				setState(205);
				((ExprContext)_localctx).op = match(NOT);
				setState(206);
				((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(18);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).right!=null?(((ExprContext)_localctx).right.start):null).GetLine(), (((ExprContext)_localctx).right!=null?(((ExprContext)_localctx).right.start):null).GetColumn(), ((ExprContext)_localctx).right.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
				}
				break;
			case PARIZQ:
				{
				setState(209);
				match(PARIZQ);
				setState(210);
				((ExprContext)_localctx).expr = expr(0);
				setState(211);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case SUB:
				{
				setState(214);
				match(SUB);
				setState(215);
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
				setState(217);
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
				setState(219);
				((ExprContext)_localctx).CADENA = match(CADENA);

				        str := (((ExprContext)_localctx).CADENA!=null?((ExprContext)_localctx).CADENA.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).CADENA!=null?((ExprContext)_localctx).CADENA.getLine():0), (((ExprContext)_localctx).CADENA!=null?((ExprContext)_localctx).CADENA.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case TRU:
				{
				setState(221);
				((ExprContext)_localctx).TRU = match(TRU);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getLine():0), (((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case FAL:
				{
				setState(223);
				((ExprContext)_localctx).FAL = match(FAL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getLine():0), (((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			case CHARACTER:
				{
				setState(225);
				((ExprContext)_localctx).CHARACTER = match(CHARACTER);
				 
				        str := (((ExprContext)_localctx).CHARACTER!=null?((ExprContext)_localctx).CHARACTER.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).CHARACTER!=null?((ExprContext)_localctx).CHARACTER.getLine():0), (((ExprContext)_localctx).CHARACTER!=null?((ExprContext)_localctx).CHARACTER.getCharPositionInLine():0), str[1:len(str)-1], environment.CHARACTER) 
				    
				}
				break;
			case ID_VALIDO:
				{
				setState(227);
				((ExprContext)_localctx).ID_VALIDO = match(ID_VALIDO);

				        id := (((ExprContext)_localctx).ID_VALIDO!=null?((ExprContext)_localctx).ID_VALIDO.getText():null)
				        _localctx.e = sentencias.NewCallid((((ExprContext)_localctx).ID_VALIDO!=null?((ExprContext)_localctx).ID_VALIDO.getLine():0),(((ExprContext)_localctx).ID_VALIDO!=null?((ExprContext)_localctx).ID_VALIDO.getCharPositionInLine():0),id)
				    
				}
				break;
			case NULO:
				{
				setState(229);
				((ExprContext)_localctx).NULO = match(NULO);
				_localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NULO!=null?((ExprContext)_localctx).NULO.getLine():0), (((ExprContext)_localctx).NULO!=null?((ExprContext)_localctx).NULO.getCharPositionInLine():0), (((ExprContext)_localctx).NULO!=null?((ExprContext)_localctx).NULO.getText():null),environment.NULL)
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(275);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,21,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(273);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(233);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(234);
						((ExprContext)_localctx).op = match(MODULO);
						setState(235);
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
						setState(238);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(239);
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
						setState(240);
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
						setState(243);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(244);
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
						setState(245);
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
						setState(248);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(249);
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
						setState(250);
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
						setState(253);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(254);
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
						setState(255);
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
						setState(258);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(259);
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
						setState(260);
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
						setState(263);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(264);
						((ExprContext)_localctx).op = match(AND);
						setState(265);
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
						setState(268);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(269);
						((ExprContext)_localctx).op = match(OR);
						setState(270);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(11);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(277);
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
			setState(305);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(278);
				((SentenciaifelseContext)_localctx).IF = match(IF);
				setState(279);
				((SentenciaifelseContext)_localctx).expr = expr(0);
				setState(280);
				match(LLAVEIZQ);
				setState(281);
				((SentenciaifelseContext)_localctx).blockinterno = blockinterno();
				setState(282);
				match(LLAVEDER);
				 _localctx.myIfElse = sentencias.NewSentenciaIf((((SentenciaifelseContext)_localctx).IF!=null?((SentenciaifelseContext)_localctx).IF.getLine():0), (((SentenciaifelseContext)_localctx).IF!=null?((SentenciaifelseContext)_localctx).IF.getCharPositionInLine():0), ((SentenciaifelseContext)_localctx).expr.e, ((SentenciaifelseContext)_localctx).blockinterno.blkint)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(285);
				((SentenciaifelseContext)_localctx).IF = match(IF);
				setState(286);
				((SentenciaifelseContext)_localctx).expr = expr(0);
				setState(287);
				match(LLAVEIZQ);
				setState(288);
				((SentenciaifelseContext)_localctx).ifop = blockinterno();
				setState(289);
				match(LLAVEDER);
				setState(290);
				match(ELSE);
				setState(291);
				match(LLAVEIZQ);
				setState(292);
				((SentenciaifelseContext)_localctx).elseop = blockinterno();
				setState(293);
				match(LLAVEDER);
				 _localctx.myIfElse = sentencias.NewSentenciaIfElse((((SentenciaifelseContext)_localctx).IF!=null?((SentenciaifelseContext)_localctx).IF.getLine():0), (((SentenciaifelseContext)_localctx).IF!=null?((SentenciaifelseContext)_localctx).IF.getCharPositionInLine():0), ((SentenciaifelseContext)_localctx).expr.e, ((SentenciaifelseContext)_localctx).ifop.blkint , ((SentenciaifelseContext)_localctx).elseop.blkint)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(296);
				((SentenciaifelseContext)_localctx).IF = match(IF);
				setState(297);
				((SentenciaifelseContext)_localctx).expr = expr(0);
				setState(298);
				match(LLAVEIZQ);
				setState(299);
				((SentenciaifelseContext)_localctx).blockinterno = blockinterno();
				setState(300);
				match(LLAVEDER);
				setState(301);
				match(ELSE);
				setState(302);
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
		enterRule(_localctx, 22, RULE_whilecontrol);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(307);
			((WhilecontrolContext)_localctx).WHILE = match(WHILE);
			setState(308);
			((WhilecontrolContext)_localctx).expr = expr(0);
			setState(309);
			match(LLAVEIZQ);
			setState(310);
			((WhilecontrolContext)_localctx).blockinterno = blockinterno();
			setState(311);
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
		enterRule(_localctx, 24, RULE_continuee);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(314);
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
		enterRule(_localctx, 26, RULE_breakk);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(317);
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
		enterRule(_localctx, 28, RULE_retornos);
		try {
			setState(326);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(320);
				((RetornosContext)_localctx).RETURN = match(RETURN);
				setState(321);
				((RetornosContext)_localctx).op = expr(0);

				    ((RetornosContext)_localctx).rect =  sentencias.NewTransferenciaReturnExp((((RetornosContext)_localctx).RETURN!=null?((RetornosContext)_localctx).RETURN.getLine():0), (((RetornosContext)_localctx).RETURN!=null?((RetornosContext)_localctx).RETURN.getCharPositionInLine():0), ((RetornosContext)_localctx).op.e);

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(324);
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
		enterRule(_localctx, 30, RULE_listaexpresions);

		    _localctx.blkparf = []interface{}{}
		    var listInt []IListaexpresionContext

		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(329); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(328);
				((ListaexpresionsContext)_localctx).listaexpresion = listaexpresion();
				((ListaexpresionsContext)_localctx).funpar.add(((ListaexpresionsContext)_localctx).listaexpresion);
				}
				}
				setState(331); 
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
		enterRule(_localctx, 32, RULE_listaexpresion);
		try {
			setState(342);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case COMA:
				enterOuterAlt(_localctx, 1);
				{
				setState(335);
				((ListaexpresionContext)_localctx).COMA = match(COMA);
				setState(336);
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
				setState(339);
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
		enterRule(_localctx, 34, RULE_printstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(344);
			((PrintstmtContext)_localctx).PRINT = match(PRINT);
			setState(345);
			match(PARIZQ);
			setState(346);
			((PrintstmtContext)_localctx).listaexpresions = listaexpresions();
			setState(347);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3M\u0161\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\3\2\3\2\3\2\3\2\3\3\6\3,\n\3\r\3\16\3-\3\3\3\3\3\4\3\4\5\4"+
		"\64\n\4\3\4\3\4\3\4\3\4\5\4:\n\4\3\4\3\4\3\4\3\4\5\4@\n\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4L\n\4\3\4\3\4\5\4P\n\4\3\5\6\5S\n\5\r"+
		"\5\16\5T\3\5\3\5\3\6\3\6\5\6[\n\6\3\6\3\6\3\6\3\6\5\6a\n\6\3\6\3\6\3\6"+
		"\3\6\5\6g\n\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6s\n\6\3\6\3\6"+
		"\3\6\3\6\5\6y\n\6\3\6\3\6\3\6\3\6\5\6\177\n\6\3\6\3\6\3\6\3\6\5\6\u0085"+
		"\n\6\3\6\3\6\5\6\u0089\n\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7"+
		"\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7\u00a0\n\7\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\5\b\u00b0\n\b\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t\u00c1\n\t\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n\u00cd\n\n\3\13\3\13\3\13\3\13\3\13\3"+
		"\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3"+
		"\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13\u00ea\n\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\7\13\u0114\n\13\f\13\16"+
		"\13\u0117\13\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\5\f\u0134\n\f\3"+
		"\r\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\17\3\17\3\17\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\5\20\u0149\n\20\3\21\6\21\u014c\n\21\r\21\16\21\u014d"+
		"\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u0159\n\22\3\23\3\23"+
		"\3\23\3\23\3\23\3\23\3\23\2\3\24\24\2\4\6\b\n\f\16\20\22\24\26\30\32\34"+
		"\36 \"$\2\7\3\2<=\3\2>?\4\2\67\6799\4\288::\3\2\62\63\2\u0187\2&\3\2\2"+
		"\2\4+\3\2\2\2\6O\3\2\2\2\bR\3\2\2\2\n\u0088\3\2\2\2\f\u009f\3\2\2\2\16"+
		"\u00af\3\2\2\2\20\u00c0\3\2\2\2\22\u00cc\3\2\2\2\24\u00e9\3\2\2\2\26\u0133"+
		"\3\2\2\2\30\u0135\3\2\2\2\32\u013c\3\2\2\2\34\u013f\3\2\2\2\36\u0148\3"+
		"\2\2\2 \u014b\3\2\2\2\"\u0158\3\2\2\2$\u015a\3\2\2\2&\'\5\4\3\2\'(\7\2"+
		"\2\3()\b\2\1\2)\3\3\2\2\2*,\5\6\4\2+*\3\2\2\2,-\3\2\2\2-+\3\2\2\2-.\3"+
		"\2\2\2./\3\2\2\2/\60\b\3\1\2\60\5\3\2\2\2\61\63\5\f\7\2\62\64\7.\2\2\63"+
		"\62\3\2\2\2\63\64\3\2\2\2\64\65\3\2\2\2\65\66\b\4\1\2\66P\3\2\2\2\679"+
		"\5\16\b\28:\7.\2\298\3\2\2\29:\3\2\2\2:;\3\2\2\2;<\b\4\1\2<P\3\2\2\2="+
		"?\5\20\t\2>@\7.\2\2?>\3\2\2\2?@\3\2\2\2@A\3\2\2\2AB\b\4\1\2BP\3\2\2\2"+
		"CD\5\26\f\2DE\b\4\1\2EP\3\2\2\2FG\5\30\r\2GH\b\4\1\2HP\3\2\2\2IK\5$\23"+
		"\2JL\7.\2\2KJ\3\2\2\2KL\3\2\2\2LM\3\2\2\2MN\b\4\1\2NP\3\2\2\2O\61\3\2"+
		"\2\2O\67\3\2\2\2O=\3\2\2\2OC\3\2\2\2OF\3\2\2\2OI\3\2\2\2P\7\3\2\2\2QS"+
		"\5\n\6\2RQ\3\2\2\2ST\3\2\2\2TR\3\2\2\2TU\3\2\2\2UV\3\2\2\2VW\b\5\1\2W"+
		"\t\3\2\2\2XZ\5\f\7\2Y[\7.\2\2ZY\3\2\2\2Z[\3\2\2\2[\\\3\2\2\2\\]\b\6\1"+
		"\2]\u0089\3\2\2\2^`\5\16\b\2_a\7.\2\2`_\3\2\2\2`a\3\2\2\2ab\3\2\2\2bc"+
		"\b\6\1\2c\u0089\3\2\2\2df\5\20\t\2eg\7.\2\2fe\3\2\2\2fg\3\2\2\2gh\3\2"+
		"\2\2hi\b\6\1\2i\u0089\3\2\2\2jk\5\26\f\2kl\b\6\1\2l\u0089\3\2\2\2mn\5"+
		"\30\r\2no\b\6\1\2o\u0089\3\2\2\2pr\5\32\16\2qs\7.\2\2rq\3\2\2\2rs\3\2"+
		"\2\2st\3\2\2\2tu\b\6\1\2u\u0089\3\2\2\2vx\5\34\17\2wy\7.\2\2xw\3\2\2\2"+
		"xy\3\2\2\2yz\3\2\2\2z{\b\6\1\2{\u0089\3\2\2\2|~\5\36\20\2}\177\7.\2\2"+
		"~}\3\2\2\2~\177\3\2\2\2\177\u0080\3\2\2\2\u0080\u0081\b\6\1\2\u0081\u0089"+
		"\3\2\2\2\u0082\u0084\5$\23\2\u0083\u0085\7.\2\2\u0084\u0083\3\2\2\2\u0084"+
		"\u0085\3\2\2\2\u0085\u0086\3\2\2\2\u0086\u0087\b\6\1\2\u0087\u0089\3\2"+
		"\2\2\u0088X\3\2\2\2\u0088^\3\2\2\2\u0088d\3\2\2\2\u0088j\3\2\2\2\u0088"+
		"m\3\2\2\2\u0088p\3\2\2\2\u0088v\3\2\2\2\u0088|\3\2\2\2\u0088\u0082\3\2"+
		"\2\2\u0089\13\3\2\2\2\u008a\u008b\7\n\2\2\u008b\u008c\7)\2\2\u008c\u008d"+
		"\7-\2\2\u008d\u008e\5\22\n\2\u008e\u008f\7,\2\2\u008f\u0090\5\24\13\2"+
		"\u0090\u0091\b\7\1\2\u0091\u00a0\3\2\2\2\u0092\u0093\7\n\2\2\u0093\u0094"+
		"\7)\2\2\u0094\u0095\7,\2\2\u0095\u0096\5\24\13\2\u0096\u0097\b\7\1\2\u0097"+
		"\u00a0\3\2\2\2\u0098\u0099\7\n\2\2\u0099\u009a\7)\2\2\u009a\u009b\7-\2"+
		"\2\u009b\u009c\5\22\n\2\u009c\u009d\7/\2\2\u009d\u009e\b\7\1\2\u009e\u00a0"+
		"\3\2\2\2\u009f\u008a\3\2\2\2\u009f\u0092\3\2\2\2\u009f\u0098\3\2\2\2\u00a0"+
		"\r\3\2\2\2\u00a1\u00a2\7\13\2\2\u00a2\u00a3\7)\2\2\u00a3\u00a4\7-\2\2"+
		"\u00a4\u00a5\5\22\n\2\u00a5\u00a6\7,\2\2\u00a6\u00a7\5\24\13\2\u00a7\u00a8"+
		"\b\b\1\2\u00a8\u00b0\3\2\2\2\u00a9\u00aa\7\13\2\2\u00aa\u00ab\7)\2\2\u00ab"+
		"\u00ac\7,\2\2\u00ac\u00ad\5\24\13\2\u00ad\u00ae\b\b\1\2\u00ae\u00b0\3"+
		"\2\2\2\u00af\u00a1\3\2\2\2\u00af\u00a9\3\2\2\2\u00b0\17\3\2\2\2\u00b1"+
		"\u00b2\7)\2\2\u00b2\u00b3\7,\2\2\u00b3\u00b4\5\24\13\2\u00b4\u00b5\b\t"+
		"\1\2\u00b5\u00c1\3\2\2\2\u00b6\u00b7\7)\2\2\u00b7\u00b8\7@\2\2\u00b8\u00b9"+
		"\5\24\13\2\u00b9\u00ba\b\t\1\2\u00ba\u00c1\3\2\2\2\u00bb\u00bc\7)\2\2"+
		"\u00bc\u00bd\7A\2\2\u00bd\u00be\5\24\13\2\u00be\u00bf\b\t\1\2\u00bf\u00c1"+
		"\3\2\2\2\u00c0\u00b1\3\2\2\2\u00c0\u00b6\3\2\2\2\u00c0\u00bb\3\2\2\2\u00c1"+
		"\21\3\2\2\2\u00c2\u00c3\7\3\2\2\u00c3\u00cd\b\n\1\2\u00c4\u00c5\7\4\2"+
		"\2\u00c5\u00cd\b\n\1\2\u00c6\u00c7\7\5\2\2\u00c7\u00cd\b\n\1\2\u00c8\u00c9"+
		"\7\6\2\2\u00c9\u00cd\b\n\1\2\u00ca\u00cb\7\7\2\2\u00cb\u00cd\b\n\1\2\u00cc"+
		"\u00c2\3\2\2\2\u00cc\u00c4\3\2\2\2\u00cc\u00c6\3\2\2\2\u00cc\u00c8\3\2"+
		"\2\2\u00cc\u00ca\3\2\2\2\u00cd\23\3\2\2\2\u00ce\u00cf\b\13\1\2\u00cf\u00d0"+
		"\7\64\2\2\u00d0\u00d1\5\24\13\24\u00d1\u00d2\b\13\1\2\u00d2\u00ea\3\2"+
		"\2\2\u00d3\u00d4\7\60\2\2\u00d4\u00d5\5\24\13\2\u00d5\u00d6\7\61\2\2\u00d6"+
		"\u00d7\b\13\1\2\u00d7\u00ea\3\2\2\2\u00d8\u00d9\7?\2\2\u00d9\u00da\7\'"+
		"\2\2\u00da\u00ea\b\13\1\2\u00db\u00dc\7\'\2\2\u00dc\u00ea\b\13\1\2\u00dd"+
		"\u00de\7(\2\2\u00de\u00ea\b\13\1\2\u00df\u00e0\7\b\2\2\u00e0\u00ea\b\13"+
		"\1\2\u00e1\u00e2\7\t\2\2\u00e2\u00ea\b\13\1\2\u00e3\u00e4\7*\2\2\u00e4"+
		"\u00ea\b\13\1\2\u00e5\u00e6\7)\2\2\u00e6\u00ea\b\13\1\2\u00e7\u00e8\7"+
		"\f\2\2\u00e8\u00ea\b\13\1\2\u00e9\u00ce\3\2\2\2\u00e9\u00d3\3\2\2\2\u00e9"+
		"\u00d8\3\2\2\2\u00e9\u00db\3\2\2\2\u00e9\u00dd\3\2\2\2\u00e9\u00df\3\2"+
		"\2\2\u00e9\u00e1\3\2\2\2\u00e9\u00e3\3\2\2\2\u00e9\u00e5\3\2\2\2\u00e9"+
		"\u00e7\3\2\2\2\u00ea\u0115\3\2\2\2\u00eb\u00ec\f\23\2\2\u00ec\u00ed\7"+
		";\2\2\u00ed\u00ee\5\24\13\24\u00ee\u00ef\b\13\1\2\u00ef\u0114\3\2\2\2"+
		"\u00f0\u00f1\f\22\2\2\u00f1\u00f2\t\2\2\2\u00f2\u00f3\5\24\13\23\u00f3"+
		"\u00f4\b\13\1\2\u00f4\u0114\3\2\2\2\u00f5\u00f6\f\21\2\2\u00f6\u00f7\t"+
		"\3\2\2\u00f7\u00f8\5\24\13\22\u00f8\u00f9\b\13\1\2\u00f9\u0114\3\2\2\2"+
		"\u00fa\u00fb\f\20\2\2\u00fb\u00fc\t\4\2\2\u00fc\u00fd\5\24\13\21\u00fd"+
		"\u00fe\b\13\1\2\u00fe\u0114\3\2\2\2\u00ff\u0100\f\17\2\2\u0100\u0101\t"+
		"\5\2\2\u0101\u0102\5\24\13\20\u0102\u0103\b\13\1\2\u0103\u0114\3\2\2\2"+
		"\u0104\u0105\f\16\2\2\u0105\u0106\t\6\2\2\u0106\u0107\5\24\13\17\u0107"+
		"\u0108\b\13\1\2\u0108\u0114\3\2\2\2\u0109\u010a\f\r\2\2\u010a\u010b\7"+
		"\66\2\2\u010b\u010c\5\24\13\16\u010c\u010d\b\13\1\2\u010d\u0114\3\2\2"+
		"\2\u010e\u010f\f\f\2\2\u010f\u0110\7\65\2\2\u0110\u0111\5\24\13\r\u0111"+
		"\u0112\b\13\1\2\u0112\u0114\3\2\2\2\u0113\u00eb\3\2\2\2\u0113\u00f0\3"+
		"\2\2\2\u0113\u00f5\3\2\2\2\u0113\u00fa\3\2\2\2\u0113\u00ff\3\2\2\2\u0113"+
		"\u0104\3\2\2\2\u0113\u0109\3\2\2\2\u0113\u010e\3\2\2\2\u0114\u0117\3\2"+
		"\2\2\u0115\u0113\3\2\2\2\u0115\u0116\3\2\2\2\u0116\25\3\2\2\2\u0117\u0115"+
		"\3\2\2\2\u0118\u0119\7\r\2\2\u0119\u011a\5\24\13\2\u011a\u011b\7B\2\2"+
		"\u011b\u011c\5\b\5\2\u011c\u011d\7C\2\2\u011d\u011e\b\f\1\2\u011e\u0134"+
		"\3\2\2\2\u011f\u0120\7\r\2\2\u0120\u0121\5\24\13\2\u0121\u0122\7B\2\2"+
		"\u0122\u0123\5\b\5\2\u0123\u0124\7C\2\2\u0124\u0125\7\16\2\2\u0125\u0126"+
		"\7B\2\2\u0126\u0127\5\b\5\2\u0127\u0128\7C\2\2\u0128\u0129\b\f\1\2\u0129"+
		"\u0134\3\2\2\2\u012a\u012b\7\r\2\2\u012b\u012c\5\24\13\2\u012c\u012d\7"+
		"B\2\2\u012d\u012e\5\b\5\2\u012e\u012f\7C\2\2\u012f\u0130\7\16\2\2\u0130"+
		"\u0131\5\26\f\2\u0131\u0132\b\f\1\2\u0132\u0134\3\2\2\2\u0133\u0118\3"+
		"\2\2\2\u0133\u011f\3\2\2\2\u0133\u012a\3\2\2\2\u0134\27\3\2\2\2\u0135"+
		"\u0136\7\27\2\2\u0136\u0137\5\24\13\2\u0137\u0138\7B\2\2\u0138\u0139\5"+
		"\b\5\2\u0139\u013a\7C\2\2\u013a\u013b\b\r\1\2\u013b\31\3\2\2\2\u013c\u013d"+
		"\7\23\2\2\u013d\u013e\b\16\1\2\u013e\33\3\2\2\2\u013f\u0140\7\22\2\2\u0140"+
		"\u0141\b\17\1\2\u0141\35\3\2\2\2\u0142\u0143\7\31\2\2\u0143\u0144\5\24"+
		"\13\2\u0144\u0145\b\20\1\2\u0145\u0149\3\2\2\2\u0146\u0147\7\31\2\2\u0147"+
		"\u0149\b\20\1\2\u0148\u0142\3\2\2\2\u0148\u0146\3\2\2\2\u0149\37\3\2\2"+
		"\2\u014a\u014c\5\"\22\2\u014b\u014a\3\2\2\2\u014c\u014d\3\2\2\2\u014d"+
		"\u014b\3\2\2\2\u014d\u014e\3\2\2\2\u014e\u014f\3\2\2\2\u014f\u0150\b\21"+
		"\1\2\u0150!\3\2\2\2\u0151\u0152\7E\2\2\u0152\u0153\5\24\13\2\u0153\u0154"+
		"\b\22\1\2\u0154\u0159\3\2\2\2\u0155\u0156\5\24\13\2\u0156\u0157\b\22\1"+
		"\2\u0157\u0159\3\2\2\2\u0158\u0151\3\2\2\2\u0158\u0155\3\2\2\2\u0159#"+
		"\3\2\2\2\u015a\u015b\7\33\2\2\u015b\u015c\7\60\2\2\u015c\u015d\5 \21\2"+
		"\u015d\u015e\7\61\2\2\u015e\u015f\b\23\1\2\u015f%\3\2\2\2\34-\639?KOT"+
		"Z`frx~\u0084\u0088\u009f\u00af\u00c0\u00cc\u00e9\u0113\u0115\u0133\u0148"+
		"\u014d\u0158";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}