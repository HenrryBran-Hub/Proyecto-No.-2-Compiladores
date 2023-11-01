// Generated from c:/Users/Henrr/OneDrive/Escritorio/Compiladores/Proyecto No. 2/Proyecto No. 2 Compiladores/Optimizar/C3DGrammar.g4 by ANTLR 4.13.1

    import "Optimizar/interfacesc3d"    
    import "Optimizar/instructionsc3d"
    // import "Optimizar/environmentc3d"
    // import "Optimizar/expressionsc3d"
    // import "strings"
   

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class C3DGrammarParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		INT=1, FLOAT=2, DOUBLE=3, CHAR=4, VOID=5, INCLUDE=6, STDIO=7, HEAP=8, 
		STACK=9, IF=10, GOTO=11, RETURN=12, PRINTF=13, PHEAD=14, PSTACK=15, NUMERO=16, 
		CADENA=17, ID_VALIDO=18, CHARACTER=19, WS=20, COMMA=21, DOS_PUNTOS=22, 
		PUNTOCOMA=23, IG=24, DIF=25, IG_IG=26, NOT=27, MAY_IG=28, MEN_IG=29, MAYOR=30, 
		MENOR=31, MODULO=32, MUL=33, DIV=34, ADD=35, SUB=36, PARIZQ=37, PARDER=38, 
		LLAVEIZQ=39, LLAVEDER=40, CORCHIZQ=41, CORCHDER=42, WHITESPACE=43, COMMENT=44, 
		LINE_COMMENT=45;
	public static final int
		RULE_z = 0, RULE_encabezadoa = 1, RULE_temporales = 2, RULE_blocktemporales = 3, 
		RULE_bloquetemps = 4, RULE_blockfuncions = 5, RULE_bloquefunciones = 6, 
		RULE_funcionmain = 7, RULE_block = 8, RULE_instruction = 9, RULE_head_op = 10, 
		RULE_stack_op = 11, RULE_printff = 12, RULE_embebida = 13, RULE_tipodata = 14, 
		RULE_operaritme = 15;
	private static String[] makeRuleNames() {
		return new String[] {
			"z", "encabezadoa", "temporales", "blocktemporales", "bloquetemps", "blockfuncions", 
			"bloquefunciones", "funcionmain", "block", "instruction", "head_op", 
			"stack_op", "printff", "embebida", "tipodata", "operaritme"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'int'", "'float'", "'double'", "'char'", "'void'", "'#include'", 
			"'<stdio.h>'", "'heap'", "'stack'", "'if'", "'goto'", "'return'", "'printf'", 
			"'H'", "'P'", null, null, null, null, null, "','", "':'", "';'", "'='", 
			"'!='", "'=='", "'!'", "'\\u0009\\u0009>='", "'<='", "'>'", "'<'", "'%'", 
			"'*'", "'/'", "'+'", "'-'", "'('", "')'", "'{'", "'}'", "'['", "']'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "DOUBLE", "CHAR", "VOID", "INCLUDE", "STDIO", "HEAP", 
			"STACK", "IF", "GOTO", "RETURN", "PRINTF", "PHEAD", "PSTACK", "NUMERO", 
			"CADENA", "ID_VALIDO", "CHARACTER", "WS", "COMMA", "DOS_PUNTOS", "PUNTOCOMA", 
			"IG", "DIF", "IG_IG", "NOT", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MODULO", 
			"MUL", "DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", 
			"CORCHIZQ", "CORCHDER", "WHITESPACE", "COMMENT", "LINE_COMMENT"
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
	public String getGrammarFileName() { return "C3DGrammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public C3DGrammarParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ZContext extends ParserRuleContext {
		public []interface{} code;
		public EncabezadoaContext encabezadoa;
		public TemporalesContext temporales;
		public Token PUNTOCOMA;
		public BlockfuncionsContext blockfuncions;
		public FuncionmainContext funcionmain;
		public EncabezadoaContext encabezadoa() {
			return getRuleContext(EncabezadoaContext.class,0);
		}
		public FuncionmainContext funcionmain() {
			return getRuleContext(FuncionmainContext.class,0);
		}
		public TerminalNode EOF() { return getToken(C3DGrammarParser.EOF, 0); }
		public TemporalesContext temporales() {
			return getRuleContext(TemporalesContext.class,0);
		}
		public TerminalNode PUNTOCOMA() { return getToken(C3DGrammarParser.PUNTOCOMA, 0); }
		public BlockfuncionsContext blockfuncions() {
			return getRuleContext(BlockfuncionsContext.class,0);
		}
		public ZContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_z; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).enterZ(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).exitZ(this);
		}
	}

	public final ZContext z() throws RecognitionException {
		ZContext _localctx = new ZContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_z);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			    
			        var mySlice []interface{}
			        mySlice = make([]interface{}, 0) // Inicializa el slice vacío
			    
			setState(33);
			((ZContext)_localctx).encabezadoa = encabezadoa();
			setState(37);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DOUBLE) {
				{
				setState(34);
				((ZContext)_localctx).temporales = temporales();
				setState(35);
				((ZContext)_localctx).PUNTOCOMA = match(PUNTOCOMA);
				}
			}

			setState(40);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==VOID) {
				{
				setState(39);
				((ZContext)_localctx).blockfuncions = blockfuncions();
				}
			}

			setState(42);
			((ZContext)_localctx).funcionmain = funcionmain();
			setState(43);
			match(EOF);

			        mySlice = append(mySlice, ((ZContext)_localctx).encabezadoa.encaa)
			        if ((ZContext)_localctx).PUNTOCOMA != nil {
			            mySlice = append(mySlice, ((ZContext)_localctx).temporales.tinst)
			        }

			        for _, item := range ((ZContext)_localctx).blockfuncions.blkfunc {
			            mySlice = append(mySlice, item)
			        }

			        mySlice = append(mySlice, ((ZContext)_localctx).funcionmain.funmain)

			        _localctx.code = mySlice
			    
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
	public static class EncabezadoaContext extends ParserRuleContext {
		public interfacesc3d.Instruction encaa;
		public Token h;
		public Token s;
		public TerminalNode INCLUDE() { return getToken(C3DGrammarParser.INCLUDE, 0); }
		public TerminalNode STDIO() { return getToken(C3DGrammarParser.STDIO, 0); }
		public List<TerminalNode> DOUBLE() { return getTokens(C3DGrammarParser.DOUBLE); }
		public TerminalNode DOUBLE(int i) {
			return getToken(C3DGrammarParser.DOUBLE, i);
		}
		public TerminalNode HEAP() { return getToken(C3DGrammarParser.HEAP, 0); }
		public List<TerminalNode> CORCHIZQ() { return getTokens(C3DGrammarParser.CORCHIZQ); }
		public TerminalNode CORCHIZQ(int i) {
			return getToken(C3DGrammarParser.CORCHIZQ, i);
		}
		public List<TerminalNode> CORCHDER() { return getTokens(C3DGrammarParser.CORCHDER); }
		public TerminalNode CORCHDER(int i) {
			return getToken(C3DGrammarParser.CORCHDER, i);
		}
		public List<TerminalNode> PUNTOCOMA() { return getTokens(C3DGrammarParser.PUNTOCOMA); }
		public TerminalNode PUNTOCOMA(int i) {
			return getToken(C3DGrammarParser.PUNTOCOMA, i);
		}
		public TerminalNode STACK() { return getToken(C3DGrammarParser.STACK, 0); }
		public TerminalNode PSTACK() { return getToken(C3DGrammarParser.PSTACK, 0); }
		public TerminalNode PHEAD() { return getToken(C3DGrammarParser.PHEAD, 0); }
		public List<TerminalNode> NUMERO() { return getTokens(C3DGrammarParser.NUMERO); }
		public TerminalNode NUMERO(int i) {
			return getToken(C3DGrammarParser.NUMERO, i);
		}
		public EncabezadoaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_encabezadoa; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).enterEncabezadoa(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).exitEncabezadoa(this);
		}
	}

	public final EncabezadoaContext encabezadoa() throws RecognitionException {
		EncabezadoaContext _localctx = new EncabezadoaContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_encabezadoa);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(46);
			match(INCLUDE);
			setState(47);
			match(STDIO);
			setState(48);
			match(DOUBLE);
			setState(49);
			match(HEAP);
			setState(50);
			match(CORCHIZQ);
			setState(51);
			((EncabezadoaContext)_localctx).h = match(NUMERO);
			setState(52);
			match(CORCHDER);
			setState(53);
			match(PUNTOCOMA);
			setState(54);
			match(DOUBLE);
			setState(55);
			match(STACK);
			setState(56);
			match(CORCHIZQ);
			setState(57);
			((EncabezadoaContext)_localctx).s = match(NUMERO);
			setState(58);
			match(CORCHDER);
			setState(59);
			match(PUNTOCOMA);
			setState(60);
			match(DOUBLE);
			setState(61);
			match(PSTACK);
			setState(62);
			match(PUNTOCOMA);
			setState(63);
			match(DOUBLE);
			setState(64);
			match(PHEAD);
			setState(65);
			match(PUNTOCOMA);

			        _localctx.encaa = instructionsc3d.NewAcumuladorEncabezado((((EncabezadoaContext)_localctx).h!=null?((EncabezadoaContext)_localctx).h.getText():null),(((EncabezadoaContext)_localctx).s!=null?((EncabezadoaContext)_localctx).s.getText():null))
			    
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
	public static class TemporalesContext extends ParserRuleContext {
		public interfacesc3d.Instruction tinst;
		public BlocktemporalesContext blocktemporales;
		public TerminalNode DOUBLE() { return getToken(C3DGrammarParser.DOUBLE, 0); }
		public BlocktemporalesContext blocktemporales() {
			return getRuleContext(BlocktemporalesContext.class,0);
		}
		public TemporalesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_temporales; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).enterTemporales(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).exitTemporales(this);
		}
	}

	public final TemporalesContext temporales() throws RecognitionException {
		TemporalesContext _localctx = new TemporalesContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_temporales);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(68);
			match(DOUBLE);
			setState(69);
			((TemporalesContext)_localctx).blocktemporales = blocktemporales();

			        _localctx.tinst = instructionsc3d.NewEjecucionTemporales(((TemporalesContext)_localctx).blocktemporales.blktmps)
			    
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
	public static class BlocktemporalesContext extends ParserRuleContext {
		public []interface{} blktmps;
		public BloquetempsContext bloquetemps;
		public List<BloquetempsContext> temps = new ArrayList<BloquetempsContext>();
		public List<BloquetempsContext> bloquetemps() {
			return getRuleContexts(BloquetempsContext.class);
		}
		public BloquetempsContext bloquetemps(int i) {
			return getRuleContext(BloquetempsContext.class,i);
		}
		public BlocktemporalesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_blocktemporales; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).enterBlocktemporales(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).exitBlocktemporales(this);
		}
	}

	public final BlocktemporalesContext blocktemporales() throws RecognitionException {
		BlocktemporalesContext _localctx = new BlocktemporalesContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_blocktemporales);

		    _localctx.blktmps = []interface{}{}
		    var listTemp []IBloquetempsContext

		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(73); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(72);
				((BlocktemporalesContext)_localctx).bloquetemps = bloquetemps();
				((BlocktemporalesContext)_localctx).temps.add(((BlocktemporalesContext)_localctx).bloquetemps);
				}
				}
				setState(75); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID_VALIDO || _la==COMMA );

			    listTemp = localctx.(*BlocktemporalesContext).GetTemps()
			    for _, e := range listTemp {
			        _localctx.blktmps = append(_localctx.blktmps, e.GetTemps())
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
	public static class BloquetempsContext extends ParserRuleContext {
		public interfacesc3d.Instruction temps;
		public Token ID_VALIDO;
		public TerminalNode COMMA() { return getToken(C3DGrammarParser.COMMA, 0); }
		public TerminalNode ID_VALIDO() { return getToken(C3DGrammarParser.ID_VALIDO, 0); }
		public BloquetempsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bloquetemps; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).enterBloquetemps(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).exitBloquetemps(this);
		}
	}

	public final BloquetempsContext bloquetemps() throws RecognitionException {
		BloquetempsContext _localctx = new BloquetempsContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_bloquetemps);
		try {
			setState(84);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case COMMA:
				enterOuterAlt(_localctx, 1);
				{
				setState(79);
				match(COMMA);
				setState(80);
				((BloquetempsContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				 _localctx.temps = instructionsc3d.NewArregloTemporales((((BloquetempsContext)_localctx).ID_VALIDO!=null?((BloquetempsContext)_localctx).ID_VALIDO.getText():null))
				}
				break;
			case ID_VALIDO:
				enterOuterAlt(_localctx, 2);
				{
				setState(82);
				((BloquetempsContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				 _localctx.temps = instructionsc3d.NewArregloTemporales((((BloquetempsContext)_localctx).ID_VALIDO!=null?((BloquetempsContext)_localctx).ID_VALIDO.getText():null))
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
	public static class BlockfuncionsContext extends ParserRuleContext {
		public []interface{} blkfunc;
		public BloquefuncionesContext bloquefunciones;
		public List<BloquefuncionesContext> func = new ArrayList<BloquefuncionesContext>();
		public List<BloquefuncionesContext> bloquefunciones() {
			return getRuleContexts(BloquefuncionesContext.class);
		}
		public BloquefuncionesContext bloquefunciones(int i) {
			return getRuleContext(BloquefuncionesContext.class,i);
		}
		public BlockfuncionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_blockfuncions; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).enterBlockfuncions(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).exitBlockfuncions(this);
		}
	}

	public final BlockfuncionsContext blockfuncions() throws RecognitionException {
		BlockfuncionsContext _localctx = new BlockfuncionsContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_blockfuncions);

		    _localctx.blkfunc = []interface{}{}
		    var listFunc []IBloquefuncionesContext

		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(87); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(86);
				((BlockfuncionsContext)_localctx).bloquefunciones = bloquefunciones();
				((BlockfuncionsContext)_localctx).func.add(((BlockfuncionsContext)_localctx).bloquefunciones);
				}
				}
				setState(89); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==VOID );

			    listFunc = localctx.(*BlockfuncionsContext).GetFunc_()
			    for _, e := range listFunc {
			        _localctx.blkfunc = append(_localctx.blkfunc, e.GetFunc_())
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
	public static class BloquefuncionesContext extends ParserRuleContext {
		public interfacesc3d.Instruction func;
		public Token ID_VALIDO;
		public BlockContext block;
		public TerminalNode VOID() { return getToken(C3DGrammarParser.VOID, 0); }
		public TerminalNode ID_VALIDO() { return getToken(C3DGrammarParser.ID_VALIDO, 0); }
		public TerminalNode PARIZQ() { return getToken(C3DGrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(C3DGrammarParser.PARDER, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(C3DGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(C3DGrammarParser.LLAVEDER, 0); }
		public BloquefuncionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bloquefunciones; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).enterBloquefunciones(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).exitBloquefunciones(this);
		}
	}

	public final BloquefuncionesContext bloquefunciones() throws RecognitionException {
		BloquefuncionesContext _localctx = new BloquefuncionesContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_bloquefunciones);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(93);
			match(VOID);
			setState(94);
			((BloquefuncionesContext)_localctx).ID_VALIDO = match(ID_VALIDO);
			setState(95);
			match(PARIZQ);
			setState(96);
			match(PARDER);
			setState(97);
			match(LLAVEIZQ);
			setState(98);
			((BloquefuncionesContext)_localctx).block = block();
			setState(99);
			match(LLAVEDER);

			    _localctx.func = instructionsc3d.NewFuncionVoid((((BloquefuncionesContext)_localctx).ID_VALIDO!=null?((BloquefuncionesContext)_localctx).ID_VALIDO.getText():null),((BloquefuncionesContext)_localctx).block.blk)

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
	public static class FuncionmainContext extends ParserRuleContext {
		public interfacesc3d.Instruction funmain;
		public BlockContext block;
		public TerminalNode INT() { return getToken(C3DGrammarParser.INT, 0); }
		public TerminalNode ID_VALIDO() { return getToken(C3DGrammarParser.ID_VALIDO, 0); }
		public TerminalNode PARIZQ() { return getToken(C3DGrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(C3DGrammarParser.PARDER, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(C3DGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode RETURN() { return getToken(C3DGrammarParser.RETURN, 0); }
		public TerminalNode NUMERO() { return getToken(C3DGrammarParser.NUMERO, 0); }
		public TerminalNode PUNTOCOMA() { return getToken(C3DGrammarParser.PUNTOCOMA, 0); }
		public TerminalNode LLAVEDER() { return getToken(C3DGrammarParser.LLAVEDER, 0); }
		public FuncionmainContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcionmain; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).enterFuncionmain(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).exitFuncionmain(this);
		}
	}

	public final FuncionmainContext funcionmain() throws RecognitionException {
		FuncionmainContext _localctx = new FuncionmainContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_funcionmain);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(102);
			match(INT);
			setState(103);
			match(ID_VALIDO);
			setState(104);
			match(PARIZQ);
			setState(105);
			match(PARDER);
			setState(106);
			match(LLAVEIZQ);
			setState(107);
			((FuncionmainContext)_localctx).block = block();
			setState(108);
			match(RETURN);
			setState(109);
			match(NUMERO);
			setState(110);
			match(PUNTOCOMA);
			setState(111);
			match(LLAVEDER);

			    _localctx.funmain = instructionsc3d.NewFuncionMain(((FuncionmainContext)_localctx).block.blk)

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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).enterBlock(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).exitBlock(this);
		}
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_block);

		    _localctx.blk = []interface{}{}
		    var listBlo []IInstructionContext
		  
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(115); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(114);
				((BlockContext)_localctx).instruction = instruction();
				((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
				}
				}
				setState(117); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 320256L) != 0) );

			        listBlo = localctx.(*BlockContext).GetIns()
			        for _, e := range listBlo {
			            _localctx.blk = append(_localctx.blk, e.GetInstr())
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
		public interfacesc3d.Instruction instr;
		public Head_opContext head_op;
		public Stack_opContext stack_op;
		public PrintffContext printff;
		public OperaritmeContext operaritme;
		public Head_opContext head_op() {
			return getRuleContext(Head_opContext.class,0);
		}
		public Stack_opContext stack_op() {
			return getRuleContext(Stack_opContext.class,0);
		}
		public PrintffContext printff() {
			return getRuleContext(PrintffContext.class,0);
		}
		public OperaritmeContext operaritme() {
			return getRuleContext(OperaritmeContext.class,0);
		}
		public InstructionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruction; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).enterInstruction(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).exitInstruction(this);
		}
	}

	public final InstructionContext instruction() throws RecognitionException {
		InstructionContext _localctx = new InstructionContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_instruction);
		try {
			setState(133);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(121);
				((InstructionContext)_localctx).head_op = head_op();
				 _localctx.instr = ((InstructionContext)_localctx).head_op.heapop
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(124);
				((InstructionContext)_localctx).stack_op = stack_op();
				 _localctx.instr = ((InstructionContext)_localctx).stack_op.stackop
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(127);
				((InstructionContext)_localctx).printff = printff();
				 _localctx.instr = ((InstructionContext)_localctx).printff.prtff 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(130);
				((InstructionContext)_localctx).operaritme = operaritme();
				 _localctx.instr = ((InstructionContext)_localctx).operaritme.oparit 
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
	public static class Head_opContext extends ParserRuleContext {
		public interfacesc3d.Instruction heapop;
		public Token NUMERO;
		public Token ID_VALIDO;
		public EmbebidaContext embebida;
		public Token op;
		public List<TerminalNode> PHEAD() { return getTokens(C3DGrammarParser.PHEAD); }
		public TerminalNode PHEAD(int i) {
			return getToken(C3DGrammarParser.PHEAD, i);
		}
		public TerminalNode IG() { return getToken(C3DGrammarParser.IG, 0); }
		public TerminalNode NUMERO() { return getToken(C3DGrammarParser.NUMERO, 0); }
		public TerminalNode PUNTOCOMA() { return getToken(C3DGrammarParser.PUNTOCOMA, 0); }
		public TerminalNode ID_VALIDO() { return getToken(C3DGrammarParser.ID_VALIDO, 0); }
		public TerminalNode HEAP() { return getToken(C3DGrammarParser.HEAP, 0); }
		public TerminalNode CORCHIZQ() { return getToken(C3DGrammarParser.CORCHIZQ, 0); }
		public EmbebidaContext embebida() {
			return getRuleContext(EmbebidaContext.class,0);
		}
		public TerminalNode CORCHDER() { return getToken(C3DGrammarParser.CORCHDER, 0); }
		public TerminalNode ADD() { return getToken(C3DGrammarParser.ADD, 0); }
		public Head_opContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_head_op; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).enterHead_op(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).exitHead_op(this);
		}
	}

	public final Head_opContext head_op() throws RecognitionException {
		Head_opContext _localctx = new Head_opContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_head_op);
		int _la;
		try {
			setState(162);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(135);
				match(PHEAD);
				setState(136);
				match(IG);
				setState(137);
				((Head_opContext)_localctx).NUMERO = match(NUMERO);
				setState(138);
				match(PUNTOCOMA);

				    _localctx.heapop = instructionsc3d.NewHeapop1((((Head_opContext)_localctx).NUMERO!=null?((Head_opContext)_localctx).NUMERO.getText():null))

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(140);
				((Head_opContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(141);
				match(IG);
				setState(142);
				match(PHEAD);
				setState(143);
				match(PUNTOCOMA);

				    _localctx.heapop = instructionsc3d.NewHeapop2((((Head_opContext)_localctx).ID_VALIDO!=null?((Head_opContext)_localctx).ID_VALIDO.getText():null))

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(145);
				match(HEAP);
				setState(146);
				match(CORCHIZQ);
				setState(147);
				((Head_opContext)_localctx).embebida = embebida();
				setState(148);
				match(PHEAD);
				setState(149);
				match(CORCHDER);
				setState(150);
				match(IG);
				setState(151);
				((Head_opContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==NUMERO || _la==ID_VALIDO) ) {
					((Head_opContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(152);
				match(PUNTOCOMA);

				    _localctx.heapop = instructionsc3d.NewHeapop3(((Head_opContext)_localctx).embebida.embe,(((Head_opContext)_localctx).op!=null?((Head_opContext)_localctx).op.getText():null))

				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(155);
				match(PHEAD);
				setState(156);
				match(IG);
				setState(157);
				match(PHEAD);
				setState(158);
				match(ADD);
				setState(159);
				((Head_opContext)_localctx).NUMERO = match(NUMERO);
				setState(160);
				match(PUNTOCOMA);

				    _localctx.heapop = instructionsc3d.NewHeapop4((((Head_opContext)_localctx).NUMERO!=null?((Head_opContext)_localctx).NUMERO.getText():null))

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
	public static class Stack_opContext extends ParserRuleContext {
		public interfacesc3d.Instruction stackop;
		public Token NUMERO;
		public Token ID_VALIDO;
		public Token op1;
		public EmbebidaContext tip1;
		public Token op2;
		public Token PARIZQ;
		public TipodataContext tipodata;
		public Token op3;
		public Token op;
		public List<TerminalNode> PSTACK() { return getTokens(C3DGrammarParser.PSTACK); }
		public TerminalNode PSTACK(int i) {
			return getToken(C3DGrammarParser.PSTACK, i);
		}
		public TerminalNode IG() { return getToken(C3DGrammarParser.IG, 0); }
		public List<TerminalNode> NUMERO() { return getTokens(C3DGrammarParser.NUMERO); }
		public TerminalNode NUMERO(int i) {
			return getToken(C3DGrammarParser.NUMERO, i);
		}
		public TerminalNode PUNTOCOMA() { return getToken(C3DGrammarParser.PUNTOCOMA, 0); }
		public List<TerminalNode> ID_VALIDO() { return getTokens(C3DGrammarParser.ID_VALIDO); }
		public TerminalNode ID_VALIDO(int i) {
			return getToken(C3DGrammarParser.ID_VALIDO, i);
		}
		public TerminalNode STACK() { return getToken(C3DGrammarParser.STACK, 0); }
		public TerminalNode CORCHIZQ() { return getToken(C3DGrammarParser.CORCHIZQ, 0); }
		public TerminalNode CORCHDER() { return getToken(C3DGrammarParser.CORCHDER, 0); }
		public TerminalNode PARIZQ() { return getToken(C3DGrammarParser.PARIZQ, 0); }
		public TipodataContext tipodata() {
			return getRuleContext(TipodataContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(C3DGrammarParser.PARDER, 0); }
		public EmbebidaContext embebida() {
			return getRuleContext(EmbebidaContext.class,0);
		}
		public TerminalNode ADD() { return getToken(C3DGrammarParser.ADD, 0); }
		public TerminalNode SUB() { return getToken(C3DGrammarParser.SUB, 0); }
		public Stack_opContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stack_op; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).enterStack_op(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).exitStack_op(this);
		}
	}

	public final Stack_opContext stack_op() throws RecognitionException {
		Stack_opContext _localctx = new Stack_opContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_stack_op);
		int _la;
		try {
			setState(200);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(164);
				match(PSTACK);
				setState(165);
				match(IG);
				setState(166);
				((Stack_opContext)_localctx).NUMERO = match(NUMERO);
				setState(167);
				match(PUNTOCOMA);

				    _localctx.stackop = instructionsc3d.NewStackop1((((Stack_opContext)_localctx).NUMERO!=null?((Stack_opContext)_localctx).NUMERO.getText():null))

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(169);
				((Stack_opContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(170);
				match(IG);
				setState(171);
				match(PSTACK);
				setState(172);
				match(PUNTOCOMA);

				    _localctx.stackop = instructionsc3d.NewStackop2((((Stack_opContext)_localctx).ID_VALIDO!=null?((Stack_opContext)_localctx).ID_VALIDO.getText():null))

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(174);
				match(STACK);
				setState(175);
				match(CORCHIZQ);
				setState(180);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case NUMERO:
					{
					setState(176);
					((Stack_opContext)_localctx).op1 = match(NUMERO);
					}
					break;
				case PARIZQ:
					{
					{
					setState(177);
					((Stack_opContext)_localctx).tip1 = embebida();
					setState(178);
					((Stack_opContext)_localctx).op2 = _input.LT(1);
					_la = _input.LA(1);
					if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 360448L) != 0)) ) {
						((Stack_opContext)_localctx).op2 = (Token)_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(182);
				match(CORCHDER);
				setState(183);
				match(IG);
				setState(188);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PARIZQ) {
					{
					setState(184);
					((Stack_opContext)_localctx).PARIZQ = match(PARIZQ);
					setState(185);
					((Stack_opContext)_localctx).tipodata = tipodata();
					setState(186);
					match(PARDER);
					}
				}

				setState(190);
				((Stack_opContext)_localctx).op3 = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==NUMERO || _la==ID_VALIDO) ) {
					((Stack_opContext)_localctx).op3 = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(191);
				match(PUNTOCOMA);

				    if ((Stack_opContext)_localctx).op1 != nil {
				        if ((Stack_opContext)_localctx).PARIZQ != nil {
				             _localctx.stackop = instructionsc3d.NewStackop31((((Stack_opContext)_localctx).op1!=null?((Stack_opContext)_localctx).op1.getText():null), ((Stack_opContext)_localctx).tipodata.tipdat, (((Stack_opContext)_localctx).op3!=null?((Stack_opContext)_localctx).op3.getText():null))
				        }else{
				            _localctx.stackop = instructionsc3d.NewStackop32((((Stack_opContext)_localctx).op1!=null?((Stack_opContext)_localctx).op1.getText():null), (((Stack_opContext)_localctx).op3!=null?((Stack_opContext)_localctx).op3.getText():null))
				        }       
				    }else{
				         if ((Stack_opContext)_localctx).PARIZQ != nil {
				             _localctx.stackop = instructionsc3d.NewStackop33(((Stack_opContext)_localctx).tip1.embe, (((Stack_opContext)_localctx).op2!=null?((Stack_opContext)_localctx).op2.getText():null), ((Stack_opContext)_localctx).tipodata.tipdat, (((Stack_opContext)_localctx).op3!=null?((Stack_opContext)_localctx).op3.getText():null))
				        }else{
				            _localctx.stackop = instructionsc3d.NewStackop34(((Stack_opContext)_localctx).tip1.embe, (((Stack_opContext)_localctx).op2!=null?((Stack_opContext)_localctx).op2.getText():null), (((Stack_opContext)_localctx).op3!=null?((Stack_opContext)_localctx).op3.getText():null))
				        } 
				    }    

				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(193);
				match(PSTACK);
				setState(194);
				match(IG);
				setState(195);
				match(PSTACK);
				setState(196);
				((Stack_opContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==ADD || _la==SUB) ) {
					((Stack_opContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(197);
				((Stack_opContext)_localctx).NUMERO = match(NUMERO);
				setState(198);
				match(PUNTOCOMA);

				    _localctx.stackop = instructionsc3d.NewStackop4((((Stack_opContext)_localctx).op!=null?((Stack_opContext)_localctx).op.getText():null),(((Stack_opContext)_localctx).NUMERO!=null?((Stack_opContext)_localctx).NUMERO.getText():null))

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
	public static class PrintffContext extends ParserRuleContext {
		public interfacesc3d.Instruction prtff;
		public Token CADENA;
		public Token NUMERO;
		public EmbebidaContext embebida;
		public Token op;
		public TerminalNode PRINTF() { return getToken(C3DGrammarParser.PRINTF, 0); }
		public TerminalNode PARIZQ() { return getToken(C3DGrammarParser.PARIZQ, 0); }
		public TerminalNode CADENA() { return getToken(C3DGrammarParser.CADENA, 0); }
		public TerminalNode COMMA() { return getToken(C3DGrammarParser.COMMA, 0); }
		public TerminalNode NUMERO() { return getToken(C3DGrammarParser.NUMERO, 0); }
		public TerminalNode PARDER() { return getToken(C3DGrammarParser.PARDER, 0); }
		public TerminalNode PUNTOCOMA() { return getToken(C3DGrammarParser.PUNTOCOMA, 0); }
		public EmbebidaContext embebida() {
			return getRuleContext(EmbebidaContext.class,0);
		}
		public TerminalNode ID_VALIDO() { return getToken(C3DGrammarParser.ID_VALIDO, 0); }
		public PrintffContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printff; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).enterPrintff(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).exitPrintff(this);
		}
	}

	public final PrintffContext printff() throws RecognitionException {
		PrintffContext _localctx = new PrintffContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_printff);
		int _la;
		try {
			setState(220);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(202);
				match(PRINTF);
				setState(203);
				match(PARIZQ);
				setState(204);
				((PrintffContext)_localctx).CADENA = match(CADENA);
				setState(205);
				match(COMMA);
				setState(206);
				((PrintffContext)_localctx).NUMERO = match(NUMERO);
				setState(207);
				match(PARDER);
				setState(208);
				match(PUNTOCOMA);

				    _localctx.prtff = instructionsc3d.NewPrint1((((PrintffContext)_localctx).CADENA!=null?((PrintffContext)_localctx).CADENA.getText():null), (((PrintffContext)_localctx).NUMERO!=null?((PrintffContext)_localctx).NUMERO.getText():null))

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(210);
				match(PRINTF);
				setState(211);
				match(PARIZQ);
				setState(212);
				((PrintffContext)_localctx).CADENA = match(CADENA);
				setState(213);
				match(COMMA);
				setState(214);
				((PrintffContext)_localctx).embebida = embebida();
				setState(215);
				((PrintffContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==NUMERO || _la==ID_VALIDO) ) {
					((PrintffContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(216);
				match(PARDER);
				setState(217);
				match(PUNTOCOMA);

				    _localctx.prtff = instructionsc3d.NewPrint2((((PrintffContext)_localctx).CADENA!=null?((PrintffContext)_localctx).CADENA.getText():null), ((PrintffContext)_localctx).embebida.embe, (((PrintffContext)_localctx).op!=null?((PrintffContext)_localctx).op.getText():null))

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
	public static class EmbebidaContext extends ParserRuleContext {
		public string embe;
		public TerminalNode PARIZQ() { return getToken(C3DGrammarParser.PARIZQ, 0); }
		public TerminalNode INT() { return getToken(C3DGrammarParser.INT, 0); }
		public TerminalNode PARDER() { return getToken(C3DGrammarParser.PARDER, 0); }
		public TerminalNode FLOAT() { return getToken(C3DGrammarParser.FLOAT, 0); }
		public TerminalNode CHAR() { return getToken(C3DGrammarParser.CHAR, 0); }
		public EmbebidaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_embebida; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).enterEmbebida(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).exitEmbebida(this);
		}
	}

	public final EmbebidaContext embebida() throws RecognitionException {
		EmbebidaContext _localctx = new EmbebidaContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_embebida);
		try {
			setState(234);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(222);
				match(PARIZQ);
				setState(223);
				match(INT);
				setState(224);
				match(PARDER);

				    str := "(int)"
				    _localctx.embe = str

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(226);
				match(PARIZQ);
				setState(227);
				match(FLOAT);
				setState(228);
				match(PARDER);

				    str := "(float)"
				    _localctx.embe = str

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(230);
				match(PARIZQ);
				setState(231);
				match(CHAR);
				setState(232);
				match(PARDER);

				    str := "(char)"
				    _localctx.embe = str

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
	public static class TipodataContext extends ParserRuleContext {
		public string tipdat;
		public TerminalNode INT() { return getToken(C3DGrammarParser.INT, 0); }
		public TerminalNode FLOAT() { return getToken(C3DGrammarParser.FLOAT, 0); }
		public TerminalNode CHAR() { return getToken(C3DGrammarParser.CHAR, 0); }
		public TipodataContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipodata; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).enterTipodata(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).exitTipodata(this);
		}
	}

	public final TipodataContext tipodata() throws RecognitionException {
		TipodataContext _localctx = new TipodataContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_tipodata);
		try {
			setState(242);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(236);
				match(INT);

				    str := "int"
				    _localctx.tipdat = str

				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(238);
				match(FLOAT);

				    str := "float"
				    _localctx.tipdat = str

				}
				break;
			case CHAR:
				enterOuterAlt(_localctx, 3);
				{
				setState(240);
				match(CHAR);

				    str := "char"
				    _localctx.tipdat = str

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
	public static class OperaritmeContext extends ParserRuleContext {
		public interfacesc3d.Instruction oparit;
		public Token ID_VALIDO;
		public Token op;
		public Token opp;
		public Token tip1;
		public TipodataContext tip11;
		public Token val1;
		public Token tip2;
		public TipodataContext tip22;
		public Token val2;
		public List<TerminalNode> ID_VALIDO() { return getTokens(C3DGrammarParser.ID_VALIDO); }
		public TerminalNode ID_VALIDO(int i) {
			return getToken(C3DGrammarParser.ID_VALIDO, i);
		}
		public TerminalNode IG() { return getToken(C3DGrammarParser.IG, 0); }
		public TerminalNode PUNTOCOMA() { return getToken(C3DGrammarParser.PUNTOCOMA, 0); }
		public List<TerminalNode> NUMERO() { return getTokens(C3DGrammarParser.NUMERO); }
		public TerminalNode NUMERO(int i) {
			return getToken(C3DGrammarParser.NUMERO, i);
		}
		public TerminalNode ADD() { return getToken(C3DGrammarParser.ADD, 0); }
		public TerminalNode SUB() { return getToken(C3DGrammarParser.SUB, 0); }
		public TerminalNode MUL() { return getToken(C3DGrammarParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(C3DGrammarParser.DIV, 0); }
		public TerminalNode MODULO() { return getToken(C3DGrammarParser.MODULO, 0); }
		public List<TerminalNode> PARIZQ() { return getTokens(C3DGrammarParser.PARIZQ); }
		public TerminalNode PARIZQ(int i) {
			return getToken(C3DGrammarParser.PARIZQ, i);
		}
		public List<TipodataContext> tipodata() {
			return getRuleContexts(TipodataContext.class);
		}
		public TipodataContext tipodata(int i) {
			return getRuleContext(TipodataContext.class,i);
		}
		public OperaritmeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operaritme; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).enterOperaritme(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof C3DGrammarListener ) ((C3DGrammarListener)listener).exitOperaritme(this);
		}
	}

	public final OperaritmeContext operaritme() throws RecognitionException {
		OperaritmeContext _localctx = new OperaritmeContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_operaritme);
		int _la;
		try {
			setState(268);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(244);
				((OperaritmeContext)_localctx).ID_VALIDO = match(ID_VALIDO);
				setState(245);
				match(IG);
				setState(246);
				((OperaritmeContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==NUMERO || _la==ID_VALIDO) ) {
					((OperaritmeContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(247);
				match(PUNTOCOMA);

				    _localctx.oparit = instructionsc3d.NewOparit1((((OperaritmeContext)_localctx).ID_VALIDO!=null?((OperaritmeContext)_localctx).ID_VALIDO.getText():null), (((OperaritmeContext)_localctx).op!=null?((OperaritmeContext)_localctx).op.getText():null))

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(249);
				((OperaritmeContext)_localctx).opp = match(ID_VALIDO);
				setState(250);
				match(IG);
				setState(255);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PARIZQ) {
					{
					setState(251);
					((OperaritmeContext)_localctx).tip1 = match(PARIZQ);
					setState(252);
					((OperaritmeContext)_localctx).tip11 = tipodata();
					setState(253);
					match(PARIZQ);
					}
				}

				setState(257);
				((OperaritmeContext)_localctx).val1 = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==NUMERO || _la==ID_VALIDO) ) {
					((OperaritmeContext)_localctx).val1 = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(258);
				((OperaritmeContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 133143986176L) != 0)) ) {
					((OperaritmeContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(263);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PARIZQ) {
					{
					setState(259);
					((OperaritmeContext)_localctx).tip2 = match(PARIZQ);
					setState(260);
					((OperaritmeContext)_localctx).tip22 = tipodata();
					setState(261);
					match(PARIZQ);
					}
				}

				setState(265);
				((OperaritmeContext)_localctx).val2 = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==NUMERO || _la==ID_VALIDO) ) {
					((OperaritmeContext)_localctx).val2 = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(266);
				match(PUNTOCOMA);

				    if ((OperaritmeContext)_localctx).tip1 != nil && ((OperaritmeContext)_localctx).tip2 != nil{
				        _localctx.oparit = instructionsc3d.NewOparit21((((OperaritmeContext)_localctx).opp!=null?((OperaritmeContext)_localctx).opp.getText():null), ((OperaritmeContext)_localctx).tip11.tipdat, (((OperaritmeContext)_localctx).val1!=null?((OperaritmeContext)_localctx).val1.getText():null), (((OperaritmeContext)_localctx).op!=null?((OperaritmeContext)_localctx).op.getText():null), ((OperaritmeContext)_localctx).tip22.tipdat, (((OperaritmeContext)_localctx).val2!=null?((OperaritmeContext)_localctx).val2.getText():null))
				    }else if ((OperaritmeContext)_localctx).tip1 == nil && ((OperaritmeContext)_localctx).tip2 == nil{
				        _localctx.oparit = instructionsc3d.NewOparit22((((OperaritmeContext)_localctx).opp!=null?((OperaritmeContext)_localctx).opp.getText():null), (((OperaritmeContext)_localctx).val1!=null?((OperaritmeContext)_localctx).val1.getText():null), (((OperaritmeContext)_localctx).op!=null?((OperaritmeContext)_localctx).op.getText():null), (((OperaritmeContext)_localctx).val2!=null?((OperaritmeContext)_localctx).val2.getText():null))
				    }else if ((OperaritmeContext)_localctx).tip1 != nil && ((OperaritmeContext)_localctx).tip2 == nil{
				        _localctx.oparit = instructionsc3d.NewOparit24((((OperaritmeContext)_localctx).opp!=null?((OperaritmeContext)_localctx).opp.getText():null), ((OperaritmeContext)_localctx).tip11.tipdat, (((OperaritmeContext)_localctx).val1!=null?((OperaritmeContext)_localctx).val1.getText():null), (((OperaritmeContext)_localctx).op!=null?((OperaritmeContext)_localctx).op.getText():null), (((OperaritmeContext)_localctx).val2!=null?((OperaritmeContext)_localctx).val2.getText():null))
				    }else if ((OperaritmeContext)_localctx).tip1 == nil && ((OperaritmeContext)_localctx).tip2 != nil{
				        _localctx.oparit = instructionsc3d.NewOparit23((((OperaritmeContext)_localctx).opp!=null?((OperaritmeContext)_localctx).opp.getText():null), (((OperaritmeContext)_localctx).val1!=null?((OperaritmeContext)_localctx).val1.getText():null), (((OperaritmeContext)_localctx).op!=null?((OperaritmeContext)_localctx).op.getText():null), ((OperaritmeContext)_localctx).tip22.tipdat, (((OperaritmeContext)_localctx).val2!=null?((OperaritmeContext)_localctx).val2.getText():null))                
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

	public static final String _serializedATN =
		"\u0004\u0001-\u010f\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0003\u0000"+
		"&\b\u0000\u0001\u0000\u0003\u0000)\b\u0000\u0001\u0000\u0001\u0000\u0001"+
		"\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0004\u0003J\b"+
		"\u0003\u000b\u0003\f\u0003K\u0001\u0003\u0001\u0003\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004U\b\u0004\u0001"+
		"\u0005\u0004\u0005X\b\u0005\u000b\u0005\f\u0005Y\u0001\u0005\u0001\u0005"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0004\bt\b\b\u000b\b\f\b"+
		"u\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0003\t\u0086\b\t\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0003\n\u00a3"+
		"\b\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0003\u000b\u00b5"+
		"\b\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0003\u000b\u00bd\b\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0003\u000b\u00c9\b\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0001\f\u0003\f\u00dd\b\f\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0003\r\u00eb\b\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0003\u000e\u00f3\b\u000e\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u0100\b\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u0108"+
		"\b\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u010d\b\u000f"+
		"\u0001\u000f\u0000\u0000\u0010\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010"+
		"\u0012\u0014\u0016\u0018\u001a\u001c\u001e\u0000\u0004\u0002\u0000\u0010"+
		"\u0010\u0012\u0012\u0002\u0000\u000f\u0010\u0012\u0012\u0001\u0000#$\u0001"+
		"\u0000 $\u0117\u0000 \u0001\u0000\u0000\u0000\u0002.\u0001\u0000\u0000"+
		"\u0000\u0004D\u0001\u0000\u0000\u0000\u0006I\u0001\u0000\u0000\u0000\b"+
		"T\u0001\u0000\u0000\u0000\nW\u0001\u0000\u0000\u0000\f]\u0001\u0000\u0000"+
		"\u0000\u000ef\u0001\u0000\u0000\u0000\u0010s\u0001\u0000\u0000\u0000\u0012"+
		"\u0085\u0001\u0000\u0000\u0000\u0014\u00a2\u0001\u0000\u0000\u0000\u0016"+
		"\u00c8\u0001\u0000\u0000\u0000\u0018\u00dc\u0001\u0000\u0000\u0000\u001a"+
		"\u00ea\u0001\u0000\u0000\u0000\u001c\u00f2\u0001\u0000\u0000\u0000\u001e"+
		"\u010c\u0001\u0000\u0000\u0000 !\u0006\u0000\uffff\uffff\u0000!%\u0003"+
		"\u0002\u0001\u0000\"#\u0003\u0004\u0002\u0000#$\u0005\u0017\u0000\u0000"+
		"$&\u0001\u0000\u0000\u0000%\"\u0001\u0000\u0000\u0000%&\u0001\u0000\u0000"+
		"\u0000&(\u0001\u0000\u0000\u0000\')\u0003\n\u0005\u0000(\'\u0001\u0000"+
		"\u0000\u0000()\u0001\u0000\u0000\u0000)*\u0001\u0000\u0000\u0000*+\u0003"+
		"\u000e\u0007\u0000+,\u0005\u0000\u0000\u0001,-\u0006\u0000\uffff\uffff"+
		"\u0000-\u0001\u0001\u0000\u0000\u0000./\u0005\u0006\u0000\u0000/0\u0005"+
		"\u0007\u0000\u000001\u0005\u0003\u0000\u000012\u0005\b\u0000\u000023\u0005"+
		")\u0000\u000034\u0005\u0010\u0000\u000045\u0005*\u0000\u000056\u0005\u0017"+
		"\u0000\u000067\u0005\u0003\u0000\u000078\u0005\t\u0000\u000089\u0005)"+
		"\u0000\u00009:\u0005\u0010\u0000\u0000:;\u0005*\u0000\u0000;<\u0005\u0017"+
		"\u0000\u0000<=\u0005\u0003\u0000\u0000=>\u0005\u000f\u0000\u0000>?\u0005"+
		"\u0017\u0000\u0000?@\u0005\u0003\u0000\u0000@A\u0005\u000e\u0000\u0000"+
		"AB\u0005\u0017\u0000\u0000BC\u0006\u0001\uffff\uffff\u0000C\u0003\u0001"+
		"\u0000\u0000\u0000DE\u0005\u0003\u0000\u0000EF\u0003\u0006\u0003\u0000"+
		"FG\u0006\u0002\uffff\uffff\u0000G\u0005\u0001\u0000\u0000\u0000HJ\u0003"+
		"\b\u0004\u0000IH\u0001\u0000\u0000\u0000JK\u0001\u0000\u0000\u0000KI\u0001"+
		"\u0000\u0000\u0000KL\u0001\u0000\u0000\u0000LM\u0001\u0000\u0000\u0000"+
		"MN\u0006\u0003\uffff\uffff\u0000N\u0007\u0001\u0000\u0000\u0000OP\u0005"+
		"\u0015\u0000\u0000PQ\u0005\u0012\u0000\u0000QU\u0006\u0004\uffff\uffff"+
		"\u0000RS\u0005\u0012\u0000\u0000SU\u0006\u0004\uffff\uffff\u0000TO\u0001"+
		"\u0000\u0000\u0000TR\u0001\u0000\u0000\u0000U\t\u0001\u0000\u0000\u0000"+
		"VX\u0003\f\u0006\u0000WV\u0001\u0000\u0000\u0000XY\u0001\u0000\u0000\u0000"+
		"YW\u0001\u0000\u0000\u0000YZ\u0001\u0000\u0000\u0000Z[\u0001\u0000\u0000"+
		"\u0000[\\\u0006\u0005\uffff\uffff\u0000\\\u000b\u0001\u0000\u0000\u0000"+
		"]^\u0005\u0005\u0000\u0000^_\u0005\u0012\u0000\u0000_`\u0005%\u0000\u0000"+
		"`a\u0005&\u0000\u0000ab\u0005\'\u0000\u0000bc\u0003\u0010\b\u0000cd\u0005"+
		"(\u0000\u0000de\u0006\u0006\uffff\uffff\u0000e\r\u0001\u0000\u0000\u0000"+
		"fg\u0005\u0001\u0000\u0000gh\u0005\u0012\u0000\u0000hi\u0005%\u0000\u0000"+
		"ij\u0005&\u0000\u0000jk\u0005\'\u0000\u0000kl\u0003\u0010\b\u0000lm\u0005"+
		"\f\u0000\u0000mn\u0005\u0010\u0000\u0000no\u0005\u0017\u0000\u0000op\u0005"+
		"(\u0000\u0000pq\u0006\u0007\uffff\uffff\u0000q\u000f\u0001\u0000\u0000"+
		"\u0000rt\u0003\u0012\t\u0000sr\u0001\u0000\u0000\u0000tu\u0001\u0000\u0000"+
		"\u0000us\u0001\u0000\u0000\u0000uv\u0001\u0000\u0000\u0000vw\u0001\u0000"+
		"\u0000\u0000wx\u0006\b\uffff\uffff\u0000x\u0011\u0001\u0000\u0000\u0000"+
		"yz\u0003\u0014\n\u0000z{\u0006\t\uffff\uffff\u0000{\u0086\u0001\u0000"+
		"\u0000\u0000|}\u0003\u0016\u000b\u0000}~\u0006\t\uffff\uffff\u0000~\u0086"+
		"\u0001\u0000\u0000\u0000\u007f\u0080\u0003\u0018\f\u0000\u0080\u0081\u0006"+
		"\t\uffff\uffff\u0000\u0081\u0086\u0001\u0000\u0000\u0000\u0082\u0083\u0003"+
		"\u001e\u000f\u0000\u0083\u0084\u0006\t\uffff\uffff\u0000\u0084\u0086\u0001"+
		"\u0000\u0000\u0000\u0085y\u0001\u0000\u0000\u0000\u0085|\u0001\u0000\u0000"+
		"\u0000\u0085\u007f\u0001\u0000\u0000\u0000\u0085\u0082\u0001\u0000\u0000"+
		"\u0000\u0086\u0013\u0001\u0000\u0000\u0000\u0087\u0088\u0005\u000e\u0000"+
		"\u0000\u0088\u0089\u0005\u0018\u0000\u0000\u0089\u008a\u0005\u0010\u0000"+
		"\u0000\u008a\u008b\u0005\u0017\u0000\u0000\u008b\u00a3\u0006\n\uffff\uffff"+
		"\u0000\u008c\u008d\u0005\u0012\u0000\u0000\u008d\u008e\u0005\u0018\u0000"+
		"\u0000\u008e\u008f\u0005\u000e\u0000\u0000\u008f\u0090\u0005\u0017\u0000"+
		"\u0000\u0090\u00a3\u0006\n\uffff\uffff\u0000\u0091\u0092\u0005\b\u0000"+
		"\u0000\u0092\u0093\u0005)\u0000\u0000\u0093\u0094\u0003\u001a\r\u0000"+
		"\u0094\u0095\u0005\u000e\u0000\u0000\u0095\u0096\u0005*\u0000\u0000\u0096"+
		"\u0097\u0005\u0018\u0000\u0000\u0097\u0098\u0007\u0000\u0000\u0000\u0098"+
		"\u0099\u0005\u0017\u0000\u0000\u0099\u009a\u0006\n\uffff\uffff\u0000\u009a"+
		"\u00a3\u0001\u0000\u0000\u0000\u009b\u009c\u0005\u000e\u0000\u0000\u009c"+
		"\u009d\u0005\u0018\u0000\u0000\u009d\u009e\u0005\u000e\u0000\u0000\u009e"+
		"\u009f\u0005#\u0000\u0000\u009f\u00a0\u0005\u0010\u0000\u0000\u00a0\u00a1"+
		"\u0005\u0017\u0000\u0000\u00a1\u00a3\u0006\n\uffff\uffff\u0000\u00a2\u0087"+
		"\u0001\u0000\u0000\u0000\u00a2\u008c\u0001\u0000\u0000\u0000\u00a2\u0091"+
		"\u0001\u0000\u0000\u0000\u00a2\u009b\u0001\u0000\u0000\u0000\u00a3\u0015"+
		"\u0001\u0000\u0000\u0000\u00a4\u00a5\u0005\u000f\u0000\u0000\u00a5\u00a6"+
		"\u0005\u0018\u0000\u0000\u00a6\u00a7\u0005\u0010\u0000\u0000\u00a7\u00a8"+
		"\u0005\u0017\u0000\u0000\u00a8\u00c9\u0006\u000b\uffff\uffff\u0000\u00a9"+
		"\u00aa\u0005\u0012\u0000\u0000\u00aa\u00ab\u0005\u0018\u0000\u0000\u00ab"+
		"\u00ac\u0005\u000f\u0000\u0000\u00ac\u00ad\u0005\u0017\u0000\u0000\u00ad"+
		"\u00c9\u0006\u000b\uffff\uffff\u0000\u00ae\u00af\u0005\t\u0000\u0000\u00af"+
		"\u00b4\u0005)\u0000\u0000\u00b0\u00b5\u0005\u0010\u0000\u0000\u00b1\u00b2"+
		"\u0003\u001a\r\u0000\u00b2\u00b3\u0007\u0001\u0000\u0000\u00b3\u00b5\u0001"+
		"\u0000\u0000\u0000\u00b4\u00b0\u0001\u0000\u0000\u0000\u00b4\u00b1\u0001"+
		"\u0000\u0000\u0000\u00b5\u00b6\u0001\u0000\u0000\u0000\u00b6\u00b7\u0005"+
		"*\u0000\u0000\u00b7\u00bc\u0005\u0018\u0000\u0000\u00b8\u00b9\u0005%\u0000"+
		"\u0000\u00b9\u00ba\u0003\u001c\u000e\u0000\u00ba\u00bb\u0005&\u0000\u0000"+
		"\u00bb\u00bd\u0001\u0000\u0000\u0000\u00bc\u00b8\u0001\u0000\u0000\u0000"+
		"\u00bc\u00bd\u0001\u0000\u0000\u0000\u00bd\u00be\u0001\u0000\u0000\u0000"+
		"\u00be\u00bf\u0007\u0000\u0000\u0000\u00bf\u00c0\u0005\u0017\u0000\u0000"+
		"\u00c0\u00c9\u0006\u000b\uffff\uffff\u0000\u00c1\u00c2\u0005\u000f\u0000"+
		"\u0000\u00c2\u00c3\u0005\u0018\u0000\u0000\u00c3\u00c4\u0005\u000f\u0000"+
		"\u0000\u00c4\u00c5\u0007\u0002\u0000\u0000\u00c5\u00c6\u0005\u0010\u0000"+
		"\u0000\u00c6\u00c7\u0005\u0017\u0000\u0000\u00c7\u00c9\u0006\u000b\uffff"+
		"\uffff\u0000\u00c8\u00a4\u0001\u0000\u0000\u0000\u00c8\u00a9\u0001\u0000"+
		"\u0000\u0000\u00c8\u00ae\u0001\u0000\u0000\u0000\u00c8\u00c1\u0001\u0000"+
		"\u0000\u0000\u00c9\u0017\u0001\u0000\u0000\u0000\u00ca\u00cb\u0005\r\u0000"+
		"\u0000\u00cb\u00cc\u0005%\u0000\u0000\u00cc\u00cd\u0005\u0011\u0000\u0000"+
		"\u00cd\u00ce\u0005\u0015\u0000\u0000\u00ce\u00cf\u0005\u0010\u0000\u0000"+
		"\u00cf\u00d0\u0005&\u0000\u0000\u00d0\u00d1\u0005\u0017\u0000\u0000\u00d1"+
		"\u00dd\u0006\f\uffff\uffff\u0000\u00d2\u00d3\u0005\r\u0000\u0000\u00d3"+
		"\u00d4\u0005%\u0000\u0000\u00d4\u00d5\u0005\u0011\u0000\u0000\u00d5\u00d6"+
		"\u0005\u0015\u0000\u0000\u00d6\u00d7\u0003\u001a\r\u0000\u00d7\u00d8\u0007"+
		"\u0000\u0000\u0000\u00d8\u00d9\u0005&\u0000\u0000\u00d9\u00da\u0005\u0017"+
		"\u0000\u0000\u00da\u00db\u0006\f\uffff\uffff\u0000\u00db\u00dd\u0001\u0000"+
		"\u0000\u0000\u00dc\u00ca\u0001\u0000\u0000\u0000\u00dc\u00d2\u0001\u0000"+
		"\u0000\u0000\u00dd\u0019\u0001\u0000\u0000\u0000\u00de\u00df\u0005%\u0000"+
		"\u0000\u00df\u00e0\u0005\u0001\u0000\u0000\u00e0\u00e1\u0005&\u0000\u0000"+
		"\u00e1\u00eb\u0006\r\uffff\uffff\u0000\u00e2\u00e3\u0005%\u0000\u0000"+
		"\u00e3\u00e4\u0005\u0002\u0000\u0000\u00e4\u00e5\u0005&\u0000\u0000\u00e5"+
		"\u00eb\u0006\r\uffff\uffff\u0000\u00e6\u00e7\u0005%\u0000\u0000\u00e7"+
		"\u00e8\u0005\u0004\u0000\u0000\u00e8\u00e9\u0005&\u0000\u0000\u00e9\u00eb"+
		"\u0006\r\uffff\uffff\u0000\u00ea\u00de\u0001\u0000\u0000\u0000\u00ea\u00e2"+
		"\u0001\u0000\u0000\u0000\u00ea\u00e6\u0001\u0000\u0000\u0000\u00eb\u001b"+
		"\u0001\u0000\u0000\u0000\u00ec\u00ed\u0005\u0001\u0000\u0000\u00ed\u00f3"+
		"\u0006\u000e\uffff\uffff\u0000\u00ee\u00ef\u0005\u0002\u0000\u0000\u00ef"+
		"\u00f3\u0006\u000e\uffff\uffff\u0000\u00f0\u00f1\u0005\u0004\u0000\u0000"+
		"\u00f1\u00f3\u0006\u000e\uffff\uffff\u0000\u00f2\u00ec\u0001\u0000\u0000"+
		"\u0000\u00f2\u00ee\u0001\u0000\u0000\u0000\u00f2\u00f0\u0001\u0000\u0000"+
		"\u0000\u00f3\u001d\u0001\u0000\u0000\u0000\u00f4\u00f5\u0005\u0012\u0000"+
		"\u0000\u00f5\u00f6\u0005\u0018\u0000\u0000\u00f6\u00f7\u0007\u0000\u0000"+
		"\u0000\u00f7\u00f8\u0005\u0017\u0000\u0000\u00f8\u010d\u0006\u000f\uffff"+
		"\uffff\u0000\u00f9\u00fa\u0005\u0012\u0000\u0000\u00fa\u00ff\u0005\u0018"+
		"\u0000\u0000\u00fb\u00fc\u0005%\u0000\u0000\u00fc\u00fd\u0003\u001c\u000e"+
		"\u0000\u00fd\u00fe\u0005%\u0000\u0000\u00fe\u0100\u0001\u0000\u0000\u0000"+
		"\u00ff\u00fb\u0001\u0000\u0000\u0000\u00ff\u0100\u0001\u0000\u0000\u0000"+
		"\u0100\u0101\u0001\u0000\u0000\u0000\u0101\u0102\u0007\u0000\u0000\u0000"+
		"\u0102\u0107\u0007\u0003\u0000\u0000\u0103\u0104\u0005%\u0000\u0000\u0104"+
		"\u0105\u0003\u001c\u000e\u0000\u0105\u0106\u0005%\u0000\u0000\u0106\u0108"+
		"\u0001\u0000\u0000\u0000\u0107\u0103\u0001\u0000\u0000\u0000\u0107\u0108"+
		"\u0001\u0000\u0000\u0000\u0108\u0109\u0001\u0000\u0000\u0000\u0109\u010a"+
		"\u0007\u0000\u0000\u0000\u010a\u010b\u0005\u0017\u0000\u0000\u010b\u010d"+
		"\u0006\u000f\uffff\uffff\u0000\u010c\u00f4\u0001\u0000\u0000\u0000\u010c"+
		"\u00f9\u0001\u0000\u0000\u0000\u010d\u001f\u0001\u0000\u0000\u0000\u0011"+
		"%(KTYu\u0085\u00a2\u00b4\u00bc\u00c8\u00dc\u00ea\u00f2\u00ff\u0107\u010c";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}