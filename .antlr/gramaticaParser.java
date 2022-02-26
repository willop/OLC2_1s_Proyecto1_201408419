// Generated from e:\WillOP\u005Cu\Compi2\Laboratorio\Proyecto1_201408419\gramatica.g4 by ANTLR 4.8

    import "proyecto1/Interfaces"

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class gramaticaParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		TK_flecha=1, TK_or=2, TK_and=3, TK_igualacion=4, TK_diferente=5, TK_mayor_igual=6, 
		TK_menor_igual=7, TK_corchete_apertura=8, TK_corchete_cierre=9, TK_llave_apertura=10, 
		TK_llave_cierre=11, TK_dosPuntos=12, TK_coma=13, TK_pcoma=14, TK_menor=15, 
		TK_mayor=16, TK_punto=17, TK_par_apertura=18, TK_par_cierre=19, TK_igual=20, 
		TK_suma=21, TK_menos=22, TK_por=23, TK_diagonal=24, TK_porcentaje=25, 
		TK_linea=26, TK_amp=27, TK_sig_admiracion=28, TK_sig_interrogacion=29, 
		WHITESPACE=30, TK_comentario_multi=31, TK_comentario_lineal=32, TKR_numericos_enteros=33, 
		TKR_numericos_flotantes=34, TKR_pow=35, TKR_vec=36, TKR_powf=37, TKR_bool=38, 
		TKR_char=39, TKR_String=40, TKR_usize=41, TKR_let=42, TKR_mut=43, TKR_struct=44, 
		TKR_as=45, TKR_println=46, TKR_true=47, TKR_false=48, TKR_fn=49, TKR_return=50, 
		TKR_abs=51, TKR_sqrt=52, TKR_to_string=53, TKR_clone=54, TKR_new=55, TKR_len=56, 
		TKR_push=57, TKR_remove=58, TKR_contains=59, TKR_insert=60, TKR_capacity=61, 
		TKR_with_capacity=62, TK_entero=63, TK_decimal=64, TK_id=65, Letra=66, 
		TK_cadena=67, TK_caracter=68;
	public static final int
		RULE_start = 0, RULE_instrucciones = 1, RULE_declaracion = 2, RULE_tipovariable = 3, 
		RULE_igualacion = 4, RULE_identificadores = 5, RULE_valores = 6, RULE_expresion = 7, 
		RULE_impresion = 8, RULE_impresioncomas = 9;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "instrucciones", "declaracion", "tipovariable", "igualacion", 
			"identificadores", "valores", "expresion", "impresion", "impresioncomas"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'->'", "'||'", "'&&'", "'=='", "'!='", "'>='", "'<='", "'{'", 
			"'}'", "'['", "']'", "':'", "','", "';'", "'<'", "'>'", "'.'", "'('", 
			"')'", "'='", "'+'", "'-'", "'*'", "'/'", "'%'", "'|'", "'&'", "'!'", 
			"'?'", null, null, null, "'i64'", "'f64'", "'pow'", "'vec'", "'powf'", 
			"'bool'", "'char'", null, "'usize'", "'let'", "'mut'", "'struct'", "'as'", 
			"'println!'", "'true'", "'false'", "'fn'", "'return'", "'abs'", "'sqrt'", 
			"'to_string()'", "'clone()'", "'new'", "'len'", "'push'", "'remove'", 
			"'contains'", "'insert'", "'capacity'", "'witch_capacity'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "TK_flecha", "TK_or", "TK_and", "TK_igualacion", "TK_diferente", 
			"TK_mayor_igual", "TK_menor_igual", "TK_corchete_apertura", "TK_corchete_cierre", 
			"TK_llave_apertura", "TK_llave_cierre", "TK_dosPuntos", "TK_coma", "TK_pcoma", 
			"TK_menor", "TK_mayor", "TK_punto", "TK_par_apertura", "TK_par_cierre", 
			"TK_igual", "TK_suma", "TK_menos", "TK_por", "TK_diagonal", "TK_porcentaje", 
			"TK_linea", "TK_amp", "TK_sig_admiracion", "TK_sig_interrogacion", "WHITESPACE", 
			"TK_comentario_multi", "TK_comentario_lineal", "TKR_numericos_enteros", 
			"TKR_numericos_flotantes", "TKR_pow", "TKR_vec", "TKR_powf", "TKR_bool", 
			"TKR_char", "TKR_String", "TKR_usize", "TKR_let", "TKR_mut", "TKR_struct", 
			"TKR_as", "TKR_println", "TKR_true", "TKR_false", "TKR_fn", "TKR_return", 
			"TKR_abs", "TKR_sqrt", "TKR_to_string", "TKR_clone", "TKR_new", "TKR_len", 
			"TKR_push", "TKR_remove", "TKR_contains", "TKR_insert", "TKR_capacity", 
			"TKR_with_capacity", "TK_entero", "TK_decimal", "TK_id", "Letra", "TK_cadena", 
			"TK_caracter"
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
	public String getGrammarFileName() { return "gramatica.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }



	public gramaticaParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(gramaticaParser.EOF, 0); }
		public List<InstruccionesContext> instrucciones() {
			return getRuleContexts(InstruccionesContext.class);
		}
		public InstruccionesContext instrucciones(int i) {
			return getRuleContext(InstruccionesContext.class,i);
		}
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(23);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 18)) & ~0x3f) == 0 && ((1L << (_la - 18)) & ((1L << (TK_par_apertura - 18)) | (1L << (TK_menos - 18)) | (1L << (TKR_pow - 18)) | (1L << (TKR_powf - 18)) | (1L << (TKR_let - 18)) | (1L << (TKR_println - 18)) | (1L << (TKR_true - 18)) | (1L << (TKR_false - 18)) | (1L << (TK_entero - 18)) | (1L << (TK_decimal - 18)) | (1L << (TK_id - 18)) | (1L << (TK_cadena - 18)) | (1L << (TK_caracter - 18)))) != 0)) {
				{
				{
				setState(20);
				instrucciones();
				}
				}
				setState(25);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(26);
			match(EOF);
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

	public static class InstruccionesContext extends ParserRuleContext {
		public ExpresionContext expresion;
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public ImpresionContext impresion() {
			return getRuleContext(ImpresionContext.class,0);
		}
		public DeclaracionContext declaracion() {
			return getRuleContext(DeclaracionContext.class,0);
		}
		public IdentificadoresContext identificadores() {
			return getRuleContext(IdentificadoresContext.class,0);
		}
		public InstruccionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instrucciones; }
	}

	public final InstruccionesContext instrucciones() throws RecognitionException {
		InstruccionesContext _localctx = new InstruccionesContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_instrucciones);
		try {
			setState(34);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(28);
				((InstruccionesContext)_localctx).expresion = expresion(0);
				fmt.Println("mensaje en instrucciones: ",((InstruccionesContext)_localctx).expresion.valorexpresion)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(31);
				impresion();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(32);
				declaracion();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(33);
				identificadores();
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

	public static class DeclaracionContext extends ParserRuleContext {
		public TerminalNode TKR_let() { return getToken(gramaticaParser.TKR_let, 0); }
		public TerminalNode TKR_mut() { return getToken(gramaticaParser.TKR_mut, 0); }
		public TerminalNode TK_id() { return getToken(gramaticaParser.TK_id, 0); }
		public TerminalNode TK_dosPuntos() { return getToken(gramaticaParser.TK_dosPuntos, 0); }
		public TipovariableContext tipovariable() {
			return getRuleContext(TipovariableContext.class,0);
		}
		public IgualacionContext igualacion() {
			return getRuleContext(IgualacionContext.class,0);
		}
		public TerminalNode TK_pcoma() { return getToken(gramaticaParser.TK_pcoma, 0); }
		public DeclaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion; }
	}

	public final DeclaracionContext declaracion() throws RecognitionException {
		DeclaracionContext _localctx = new DeclaracionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_declaracion);
		try {
			setState(62);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(36);
				match(TKR_let);
				setState(37);
				match(TKR_mut);
				setState(38);
				match(TK_id);
				setState(39);
				match(TK_dosPuntos);
				setState(40);
				tipovariable();
				setState(41);
				igualacion();
				setState(42);
				match(TK_pcoma);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(44);
				match(TKR_let);
				setState(45);
				match(TKR_mut);
				setState(46);
				match(TK_id);
				setState(47);
				igualacion();
				setState(48);
				match(TK_pcoma);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(50);
				match(TKR_let);
				setState(51);
				match(TK_id);
				setState(52);
				match(TK_dosPuntos);
				setState(53);
				tipovariable();
				setState(54);
				igualacion();
				setState(55);
				match(TK_pcoma);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(57);
				match(TKR_let);
				setState(58);
				match(TK_id);
				setState(59);
				igualacion();
				setState(60);
				match(TK_pcoma);
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

	public static class TipovariableContext extends ParserRuleContext {
		public TerminalNode TKR_numericos_enteros() { return getToken(gramaticaParser.TKR_numericos_enteros, 0); }
		public TerminalNode TKR_numericos_flotantes() { return getToken(gramaticaParser.TKR_numericos_flotantes, 0); }
		public TerminalNode TKR_String() { return getToken(gramaticaParser.TKR_String, 0); }
		public TerminalNode TKR_bool() { return getToken(gramaticaParser.TKR_bool, 0); }
		public TerminalNode TKR_char() { return getToken(gramaticaParser.TKR_char, 0); }
		public TerminalNode TKR_usize() { return getToken(gramaticaParser.TKR_usize, 0); }
		public TipovariableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipovariable; }
	}

	public final TipovariableContext tipovariable() throws RecognitionException {
		TipovariableContext _localctx = new TipovariableContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_tipovariable);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(64);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << TKR_numericos_enteros) | (1L << TKR_numericos_flotantes) | (1L << TKR_bool) | (1L << TKR_char) | (1L << TKR_String) | (1L << TKR_usize))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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

	public static class IgualacionContext extends ParserRuleContext {
		public TerminalNode TK_igual() { return getToken(gramaticaParser.TK_igual, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public IgualacionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_igualacion; }
	}

	public final IgualacionContext igualacion() throws RecognitionException {
		IgualacionContext _localctx = new IgualacionContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_igualacion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(66);
			match(TK_igual);
			setState(67);
			expresion(0);
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

	public static class IdentificadoresContext extends ParserRuleContext {
		public TerminalNode TK_id() { return getToken(gramaticaParser.TK_id, 0); }
		public IgualacionContext igualacion() {
			return getRuleContext(IgualacionContext.class,0);
		}
		public TerminalNode TK_pcoma() { return getToken(gramaticaParser.TK_pcoma, 0); }
		public IdentificadoresContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identificadores; }
	}

	public final IdentificadoresContext identificadores() throws RecognitionException {
		IdentificadoresContext _localctx = new IdentificadoresContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_identificadores);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(69);
			match(TK_id);
			setState(70);
			igualacion();
			setState(71);
			match(TK_pcoma);
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

	public static class ValoresContext extends ParserRuleContext {
		public TerminalNode TK_entero() { return getToken(gramaticaParser.TK_entero, 0); }
		public TerminalNode TK_decimal() { return getToken(gramaticaParser.TK_decimal, 0); }
		public TerminalNode TK_cadena() { return getToken(gramaticaParser.TK_cadena, 0); }
		public TerminalNode TK_caracter() { return getToken(gramaticaParser.TK_caracter, 0); }
		public TerminalNode TKR_true() { return getToken(gramaticaParser.TKR_true, 0); }
		public TerminalNode TKR_false() { return getToken(gramaticaParser.TKR_false, 0); }
		public TerminalNode TK_id() { return getToken(gramaticaParser.TK_id, 0); }
		public ValoresContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_valores; }
	}

	public final ValoresContext valores() throws RecognitionException {
		ValoresContext _localctx = new ValoresContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_valores);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(73);
			_la = _input.LA(1);
			if ( !(((((_la - 47)) & ~0x3f) == 0 && ((1L << (_la - 47)) & ((1L << (TKR_true - 47)) | (1L << (TKR_false - 47)) | (1L << (TK_entero - 47)) | (1L << (TK_decimal - 47)) | (1L << (TK_id - 47)) | (1L << (TK_cadena - 47)) | (1L << (TK_caracter - 47)))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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

	public static class ExpresionContext extends ParserRuleContext {
		public string valorexpresion;
		public ExpresionContext e1;
		public Token op;
		public ExpresionContext e2;
		public TerminalNode TK_menos() { return getToken(gramaticaParser.TK_menos, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode TKR_pow() { return getToken(gramaticaParser.TKR_pow, 0); }
		public TerminalNode TK_par_apertura() { return getToken(gramaticaParser.TK_par_apertura, 0); }
		public TerminalNode TK_coma() { return getToken(gramaticaParser.TK_coma, 0); }
		public TerminalNode TK_par_cierre() { return getToken(gramaticaParser.TK_par_cierre, 0); }
		public TerminalNode TKR_powf() { return getToken(gramaticaParser.TKR_powf, 0); }
		public ValoresContext valores() {
			return getRuleContext(ValoresContext.class,0);
		}
		public TerminalNode TKR_as() { return getToken(gramaticaParser.TKR_as, 0); }
		public TerminalNode TKR_numericos_enteros() { return getToken(gramaticaParser.TKR_numericos_enteros, 0); }
		public TerminalNode TKR_numericos_flotantes() { return getToken(gramaticaParser.TKR_numericos_flotantes, 0); }
		public TerminalNode TK_suma() { return getToken(gramaticaParser.TK_suma, 0); }
		public TerminalNode TK_por() { return getToken(gramaticaParser.TK_por, 0); }
		public TerminalNode TK_diagonal() { return getToken(gramaticaParser.TK_diagonal, 0); }
		public TerminalNode TK_porcentaje() { return getToken(gramaticaParser.TK_porcentaje, 0); }
		public TerminalNode TK_menor() { return getToken(gramaticaParser.TK_menor, 0); }
		public TerminalNode TK_mayor() { return getToken(gramaticaParser.TK_mayor, 0); }
		public TerminalNode TK_mayor_igual() { return getToken(gramaticaParser.TK_mayor_igual, 0); }
		public TerminalNode TK_menor_igual() { return getToken(gramaticaParser.TK_menor_igual, 0); }
		public TerminalNode TK_igualacion() { return getToken(gramaticaParser.TK_igualacion, 0); }
		public TerminalNode TK_diferente() { return getToken(gramaticaParser.TK_diferente, 0); }
		public TerminalNode TK_or() { return getToken(gramaticaParser.TK_or, 0); }
		public TerminalNode TK_and() { return getToken(gramaticaParser.TK_and, 0); }
		public TerminalNode TK_sig_admiracion() { return getToken(gramaticaParser.TK_sig_admiracion, 0); }
		public ExpresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expresion; }
	}

	public final ExpresionContext expresion() throws RecognitionException {
		return expresion(0);
	}

	private ExpresionContext expresion(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpresionContext _localctx = new ExpresionContext(_ctx, _parentState);
		ExpresionContext _prevctx = _localctx;
		int _startState = 14;
		enterRecursionRule(_localctx, 14, RULE_expresion, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(105);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				{
				setState(76);
				match(TK_menos);
				setState(77);
				expresion(21);
				}
				break;
			case 2:
				{
				setState(78);
				match(TKR_pow);
				setState(79);
				match(TK_par_apertura);
				setState(80);
				expresion(0);
				setState(81);
				match(TK_coma);
				setState(82);
				expresion(0);
				setState(83);
				match(TK_par_cierre);
				}
				break;
			case 3:
				{
				setState(85);
				match(TKR_powf);
				setState(86);
				match(TK_par_apertura);
				setState(87);
				expresion(0);
				setState(88);
				match(TK_coma);
				setState(89);
				expresion(0);
				setState(90);
				match(TK_par_cierre);
				}
				break;
			case 4:
				{
				setState(92);
				match(TK_par_apertura);
				setState(93);
				expresion(0);
				setState(94);
				match(TK_par_cierre);
				}
				break;
			case 5:
				{
				setState(96);
				valores();
				setState(97);
				match(TKR_as);
				setState(98);
				match(TKR_numericos_enteros);
				}
				break;
			case 6:
				{
				setState(100);
				valores();
				setState(101);
				match(TKR_as);
				setState(102);
				match(TKR_numericos_flotantes);
				}
				break;
			case 7:
				{
				setState(104);
				valores();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(159);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(157);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
					case 1:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(107);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(108);
						((ExpresionContext)_localctx).op = match(TK_suma);
						setState(109);
						((ExpresionContext)_localctx).e2 = expresion(21);
						_localctx.valorexpresion=Interfaces.OperacionAritmetica((((ExpresionContext)_localctx).e1!=null?_input.getText(((ExpresionContext)_localctx).e1.start,((ExpresionContext)_localctx).e1.stop):null),(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),(((ExpresionContext)_localctx).e2!=null?_input.getText(((ExpresionContext)_localctx).e2.start,((ExpresionContext)_localctx).e2.stop):null))
						}
						break;
					case 2:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(112);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(113);
						((ExpresionContext)_localctx).op = match(TK_menos);
						setState(114);
						((ExpresionContext)_localctx).e2 = expresion(20);
						_localctx.valorexpresion=Interfaces.OperacionAritmetica((((ExpresionContext)_localctx).e1!=null?_input.getText(((ExpresionContext)_localctx).e1.start,((ExpresionContext)_localctx).e1.stop):null),(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),(((ExpresionContext)_localctx).e2!=null?_input.getText(((ExpresionContext)_localctx).e2.start,((ExpresionContext)_localctx).e2.stop):null))
						}
						break;
					case 3:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(117);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(118);
						((ExpresionContext)_localctx).op = match(TK_por);
						setState(119);
						((ExpresionContext)_localctx).e2 = expresion(19);
						_localctx.valorexpresion=Interfaces.OperacionAritmetica((((ExpresionContext)_localctx).e1!=null?_input.getText(((ExpresionContext)_localctx).e1.start,((ExpresionContext)_localctx).e1.stop):null),(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),(((ExpresionContext)_localctx).e2!=null?_input.getText(((ExpresionContext)_localctx).e2.start,((ExpresionContext)_localctx).e2.stop):null))
						}
						break;
					case 4:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(122);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(123);
						((ExpresionContext)_localctx).op = match(TK_diagonal);
						setState(124);
						((ExpresionContext)_localctx).e2 = expresion(18);
						_localctx.valorexpresion=Interfaces.OperacionAritmetica((((ExpresionContext)_localctx).e1!=null?_input.getText(((ExpresionContext)_localctx).e1.start,((ExpresionContext)_localctx).e1.stop):null),(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),(((ExpresionContext)_localctx).e2!=null?_input.getText(((ExpresionContext)_localctx).e2.start,((ExpresionContext)_localctx).e2.stop):null))
						}
						break;
					case 5:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(127);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(128);
						match(TK_porcentaje);
						setState(129);
						expresion(15);
						}
						break;
					case 6:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(130);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(131);
						match(TK_menor);
						setState(132);
						expresion(14);
						}
						break;
					case 7:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(133);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(134);
						match(TK_mayor);
						setState(135);
						expresion(13);
						}
						break;
					case 8:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(136);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(137);
						match(TK_mayor_igual);
						setState(138);
						expresion(12);
						}
						break;
					case 9:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(139);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(140);
						match(TK_menor_igual);
						setState(141);
						expresion(11);
						}
						break;
					case 10:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(142);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(143);
						match(TK_igualacion);
						setState(144);
						expresion(10);
						}
						break;
					case 11:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(145);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(146);
						match(TK_diferente);
						setState(147);
						expresion(9);
						}
						break;
					case 12:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(148);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(149);
						match(TK_or);
						setState(150);
						expresion(8);
						}
						break;
					case 13:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(151);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(152);
						match(TK_and);
						setState(153);
						expresion(7);
						}
						break;
					case 14:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(154);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(155);
						match(TK_sig_admiracion);
						setState(156);
						expresion(6);
						}
						break;
					}
					} 
				}
				setState(161);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
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

	public static class ImpresionContext extends ParserRuleContext {
		public TerminalNode TKR_println() { return getToken(gramaticaParser.TKR_println, 0); }
		public TerminalNode TK_par_apertura() { return getToken(gramaticaParser.TK_par_apertura, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode TK_par_cierre() { return getToken(gramaticaParser.TK_par_cierre, 0); }
		public TerminalNode TK_pcoma() { return getToken(gramaticaParser.TK_pcoma, 0); }
		public ImpresioncomasContext impresioncomas() {
			return getRuleContext(ImpresioncomasContext.class,0);
		}
		public ImpresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_impresion; }
	}

	public final ImpresionContext impresion() throws RecognitionException {
		ImpresionContext _localctx = new ImpresionContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_impresion);
		try {
			setState(175);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(162);
				match(TKR_println);
				setState(163);
				match(TK_par_apertura);
				setState(164);
				expresion(0);
				setState(165);
				match(TK_par_cierre);
				setState(166);
				match(TK_pcoma);
				fmt.Println("Impresion")
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(169);
				match(TKR_println);
				setState(170);
				match(TK_par_apertura);
				setState(171);
				expresion(0);
				setState(172);
				impresioncomas(0);
				setState(173);
				match(TK_pcoma);
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

	public static class ImpresioncomasContext extends ParserRuleContext {
		public TerminalNode TK_coma() { return getToken(gramaticaParser.TK_coma, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode TK_par_cierre() { return getToken(gramaticaParser.TK_par_cierre, 0); }
		public TerminalNode TK_pcoma() { return getToken(gramaticaParser.TK_pcoma, 0); }
		public ImpresioncomasContext impresioncomas() {
			return getRuleContext(ImpresioncomasContext.class,0);
		}
		public ImpresioncomasContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_impresioncomas; }
	}

	public final ImpresioncomasContext impresioncomas() throws RecognitionException {
		return impresioncomas(0);
	}

	private ImpresioncomasContext impresioncomas(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ImpresioncomasContext _localctx = new ImpresioncomasContext(_ctx, _parentState);
		ImpresioncomasContext _prevctx = _localctx;
		int _startState = 18;
		enterRecursionRule(_localctx, 18, RULE_impresioncomas, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(178);
			match(TK_coma);
			setState(179);
			expresion(0);
			setState(180);
			match(TK_par_cierre);
			setState(181);
			match(TK_pcoma);
			fmt.Println("Impresion especial")
			}
			_ctx.stop = _input.LT(-1);
			setState(193);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ImpresioncomasContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_impresioncomas);
					setState(184);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(185);
					match(TK_coma);
					setState(186);
					expresion(0);
					setState(187);
					match(TK_par_cierre);
					setState(188);
					match(TK_pcoma);
					fmt.Println("Impresion especial")
					}
					} 
				}
				setState(195);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 7:
			return expresion_sempred((ExpresionContext)_localctx, predIndex);
		case 9:
			return impresioncomas_sempred((ImpresioncomasContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expresion_sempred(ExpresionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 20);
		case 1:
			return precpred(_ctx, 19);
		case 2:
			return precpred(_ctx, 18);
		case 3:
			return precpred(_ctx, 17);
		case 4:
			return precpred(_ctx, 14);
		case 5:
			return precpred(_ctx, 13);
		case 6:
			return precpred(_ctx, 12);
		case 7:
			return precpred(_ctx, 11);
		case 8:
			return precpred(_ctx, 10);
		case 9:
			return precpred(_ctx, 9);
		case 10:
			return precpred(_ctx, 8);
		case 11:
			return precpred(_ctx, 7);
		case 12:
			return precpred(_ctx, 6);
		case 13:
			return precpred(_ctx, 5);
		}
		return true;
	}
	private boolean impresioncomas_sempred(ImpresioncomasContext _localctx, int predIndex) {
		switch (predIndex) {
		case 14:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3F\u00c7\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\3\2\7\2\30\n\2\f\2\16\2\33\13\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\5"+
		"\3%\n\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4A\n\4\3\5\3\5\3\6\3\6\3"+
		"\6\3\7\3\7\3\7\3\7\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\5\tl\n\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\7\t\u00a0\n\t\f\t\16\t\u00a3\13\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\5\n\u00b2\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\7\13\u00c2\n\13\f\13\16\13\u00c5\13"+
		"\13\3\13\2\4\20\24\f\2\4\6\b\n\f\16\20\22\24\2\4\4\2#$(+\5\2\61\62ACE"+
		"F\2\u00d9\2\31\3\2\2\2\4$\3\2\2\2\6@\3\2\2\2\bB\3\2\2\2\nD\3\2\2\2\fG"+
		"\3\2\2\2\16K\3\2\2\2\20k\3\2\2\2\22\u00b1\3\2\2\2\24\u00b3\3\2\2\2\26"+
		"\30\5\4\3\2\27\26\3\2\2\2\30\33\3\2\2\2\31\27\3\2\2\2\31\32\3\2\2\2\32"+
		"\34\3\2\2\2\33\31\3\2\2\2\34\35\7\2\2\3\35\3\3\2\2\2\36\37\5\20\t\2\37"+
		" \b\3\1\2 %\3\2\2\2!%\5\22\n\2\"%\5\6\4\2#%\5\f\7\2$\36\3\2\2\2$!\3\2"+
		"\2\2$\"\3\2\2\2$#\3\2\2\2%\5\3\2\2\2&\'\7,\2\2\'(\7-\2\2()\7C\2\2)*\7"+
		"\16\2\2*+\5\b\5\2+,\5\n\6\2,-\7\20\2\2-A\3\2\2\2./\7,\2\2/\60\7-\2\2\60"+
		"\61\7C\2\2\61\62\5\n\6\2\62\63\7\20\2\2\63A\3\2\2\2\64\65\7,\2\2\65\66"+
		"\7C\2\2\66\67\7\16\2\2\678\5\b\5\289\5\n\6\29:\7\20\2\2:A\3\2\2\2;<\7"+
		",\2\2<=\7C\2\2=>\5\n\6\2>?\7\20\2\2?A\3\2\2\2@&\3\2\2\2@.\3\2\2\2@\64"+
		"\3\2\2\2@;\3\2\2\2A\7\3\2\2\2BC\t\2\2\2C\t\3\2\2\2DE\7\26\2\2EF\5\20\t"+
		"\2F\13\3\2\2\2GH\7C\2\2HI\5\n\6\2IJ\7\20\2\2J\r\3\2\2\2KL\t\3\2\2L\17"+
		"\3\2\2\2MN\b\t\1\2NO\7\30\2\2Ol\5\20\t\27PQ\7%\2\2QR\7\24\2\2RS\5\20\t"+
		"\2ST\7\17\2\2TU\5\20\t\2UV\7\25\2\2Vl\3\2\2\2WX\7\'\2\2XY\7\24\2\2YZ\5"+
		"\20\t\2Z[\7\17\2\2[\\\5\20\t\2\\]\7\25\2\2]l\3\2\2\2^_\7\24\2\2_`\5\20"+
		"\t\2`a\7\25\2\2al\3\2\2\2bc\5\16\b\2cd\7/\2\2de\7#\2\2el\3\2\2\2fg\5\16"+
		"\b\2gh\7/\2\2hi\7$\2\2il\3\2\2\2jl\5\16\b\2kM\3\2\2\2kP\3\2\2\2kW\3\2"+
		"\2\2k^\3\2\2\2kb\3\2\2\2kf\3\2\2\2kj\3\2\2\2l\u00a1\3\2\2\2mn\f\26\2\2"+
		"no\7\27\2\2op\5\20\t\27pq\b\t\1\2q\u00a0\3\2\2\2rs\f\25\2\2st\7\30\2\2"+
		"tu\5\20\t\26uv\b\t\1\2v\u00a0\3\2\2\2wx\f\24\2\2xy\7\31\2\2yz\5\20\t\25"+
		"z{\b\t\1\2{\u00a0\3\2\2\2|}\f\23\2\2}~\7\32\2\2~\177\5\20\t\24\177\u0080"+
		"\b\t\1\2\u0080\u00a0\3\2\2\2\u0081\u0082\f\20\2\2\u0082\u0083\7\33\2\2"+
		"\u0083\u00a0\5\20\t\21\u0084\u0085\f\17\2\2\u0085\u0086\7\21\2\2\u0086"+
		"\u00a0\5\20\t\20\u0087\u0088\f\16\2\2\u0088\u0089\7\22\2\2\u0089\u00a0"+
		"\5\20\t\17\u008a\u008b\f\r\2\2\u008b\u008c\7\b\2\2\u008c\u00a0\5\20\t"+
		"\16\u008d\u008e\f\f\2\2\u008e\u008f\7\t\2\2\u008f\u00a0\5\20\t\r\u0090"+
		"\u0091\f\13\2\2\u0091\u0092\7\6\2\2\u0092\u00a0\5\20\t\f\u0093\u0094\f"+
		"\n\2\2\u0094\u0095\7\7\2\2\u0095\u00a0\5\20\t\13\u0096\u0097\f\t\2\2\u0097"+
		"\u0098\7\4\2\2\u0098\u00a0\5\20\t\n\u0099\u009a\f\b\2\2\u009a\u009b\7"+
		"\5\2\2\u009b\u00a0\5\20\t\t\u009c\u009d\f\7\2\2\u009d\u009e\7\36\2\2\u009e"+
		"\u00a0\5\20\t\b\u009fm\3\2\2\2\u009fr\3\2\2\2\u009fw\3\2\2\2\u009f|\3"+
		"\2\2\2\u009f\u0081\3\2\2\2\u009f\u0084\3\2\2\2\u009f\u0087\3\2\2\2\u009f"+
		"\u008a\3\2\2\2\u009f\u008d\3\2\2\2\u009f\u0090\3\2\2\2\u009f\u0093\3\2"+
		"\2\2\u009f\u0096\3\2\2\2\u009f\u0099\3\2\2\2\u009f\u009c\3\2\2\2\u00a0"+
		"\u00a3\3\2\2\2\u00a1\u009f\3\2\2\2\u00a1\u00a2\3\2\2\2\u00a2\21\3\2\2"+
		"\2\u00a3\u00a1\3\2\2\2\u00a4\u00a5\7\60\2\2\u00a5\u00a6\7\24\2\2\u00a6"+
		"\u00a7\5\20\t\2\u00a7\u00a8\7\25\2\2\u00a8\u00a9\7\20\2\2\u00a9\u00aa"+
		"\b\n\1\2\u00aa\u00b2\3\2\2\2\u00ab\u00ac\7\60\2\2\u00ac\u00ad\7\24\2\2"+
		"\u00ad\u00ae\5\20\t\2\u00ae\u00af\5\24\13\2\u00af\u00b0\7\20\2\2\u00b0"+
		"\u00b2\3\2\2\2\u00b1\u00a4\3\2\2\2\u00b1\u00ab\3\2\2\2\u00b2\23\3\2\2"+
		"\2\u00b3\u00b4\b\13\1\2\u00b4\u00b5\7\17\2\2\u00b5\u00b6\5\20\t\2\u00b6"+
		"\u00b7\7\25\2\2\u00b7\u00b8\7\20\2\2\u00b8\u00b9\b\13\1\2\u00b9\u00c3"+
		"\3\2\2\2\u00ba\u00bb\f\4\2\2\u00bb\u00bc\7\17\2\2\u00bc\u00bd\5\20\t\2"+
		"\u00bd\u00be\7\25\2\2\u00be\u00bf\7\20\2\2\u00bf\u00c0\b\13\1\2\u00c0"+
		"\u00c2\3\2\2\2\u00c1\u00ba\3\2\2\2\u00c2\u00c5\3\2\2\2\u00c3\u00c1\3\2"+
		"\2\2\u00c3\u00c4\3\2\2\2\u00c4\25\3\2\2\2\u00c5\u00c3\3\2\2\2\n\31$@k"+
		"\u009f\u00a1\u00b1\u00c3";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}