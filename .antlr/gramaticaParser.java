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
		public DeclaracionContext declaracion;
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
			setState(36);
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
				((InstruccionesContext)_localctx).declaracion = declaracion();
				fmt.Println("mensaje en declaracion: ",((InstruccionesContext)_localctx).declaracion.decla)
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(35);
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
		public interface{} decla;
		public Token idd;
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
		public TerminalNode TK_igual() { return getToken(gramaticaParser.TK_igual, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public DeclaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion; }
	}

	public final DeclaracionContext declaracion() throws RecognitionException {
		DeclaracionContext _localctx = new DeclaracionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_declaracion);
		try {
			setState(66);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(38);
				match(TKR_let);
				setState(39);
				match(TKR_mut);
				setState(40);
				match(TK_id);
				setState(41);
				match(TK_dosPuntos);
				setState(42);
				tipovariable();
				setState(43);
				igualacion();
				setState(44);
				match(TK_pcoma);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(46);
				match(TKR_let);
				setState(47);
				match(TKR_mut);
				setState(48);
				match(TK_id);
				setState(49);
				igualacion();
				setState(50);
				match(TK_pcoma);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(52);
				match(TKR_let);
				setState(53);
				match(TK_id);
				setState(54);
				match(TK_dosPuntos);
				setState(55);
				tipovariable();
				setState(56);
				igualacion();
				setState(57);
				match(TK_pcoma);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(59);
				match(TKR_let);
				setState(60);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(61);
				match(TK_igual);
				setState(62);
				expresion(0);
				setState(63);
				match(TK_pcoma);
				_localctx.decla = Interfaces.ConstructorSimbolo((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),50,false,2)
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
			setState(68);
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
			setState(70);
			match(TK_igual);
			setState(71);
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
			setState(73);
			match(TK_id);
			setState(74);
			igualacion();
			setState(75);
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
			setState(77);
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
			setState(109);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				{
				setState(80);
				match(TK_menos);
				setState(81);
				expresion(21);
				}
				break;
			case 2:
				{
				setState(82);
				match(TKR_pow);
				setState(83);
				match(TK_par_apertura);
				setState(84);
				expresion(0);
				setState(85);
				match(TK_coma);
				setState(86);
				expresion(0);
				setState(87);
				match(TK_par_cierre);
				}
				break;
			case 3:
				{
				setState(89);
				match(TKR_powf);
				setState(90);
				match(TK_par_apertura);
				setState(91);
				expresion(0);
				setState(92);
				match(TK_coma);
				setState(93);
				expresion(0);
				setState(94);
				match(TK_par_cierre);
				}
				break;
			case 4:
				{
				setState(96);
				match(TK_par_apertura);
				setState(97);
				expresion(0);
				setState(98);
				match(TK_par_cierre);
				}
				break;
			case 5:
				{
				setState(100);
				valores();
				setState(101);
				match(TKR_as);
				setState(102);
				match(TKR_numericos_enteros);
				}
				break;
			case 6:
				{
				setState(104);
				valores();
				setState(105);
				match(TKR_as);
				setState(106);
				match(TKR_numericos_flotantes);
				}
				break;
			case 7:
				{
				setState(108);
				valores();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(163);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(161);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
					case 1:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(111);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(112);
						((ExpresionContext)_localctx).op = match(TK_suma);
						setState(113);
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
						setState(116);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(117);
						((ExpresionContext)_localctx).op = match(TK_menos);
						setState(118);
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
						setState(121);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(122);
						((ExpresionContext)_localctx).op = match(TK_por);
						setState(123);
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
						setState(126);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(127);
						((ExpresionContext)_localctx).op = match(TK_diagonal);
						setState(128);
						((ExpresionContext)_localctx).e2 = expresion(18);
						_localctx.valorexpresion=Interfaces.OperacionAritmetica((((ExpresionContext)_localctx).e1!=null?_input.getText(((ExpresionContext)_localctx).e1.start,((ExpresionContext)_localctx).e1.stop):null),(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),(((ExpresionContext)_localctx).e2!=null?_input.getText(((ExpresionContext)_localctx).e2.start,((ExpresionContext)_localctx).e2.stop):null))
						}
						break;
					case 5:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(131);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(132);
						match(TK_porcentaje);
						setState(133);
						expresion(15);
						}
						break;
					case 6:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(134);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(135);
						match(TK_menor);
						setState(136);
						expresion(14);
						}
						break;
					case 7:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(137);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(138);
						match(TK_mayor);
						setState(139);
						expresion(13);
						}
						break;
					case 8:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(140);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(141);
						match(TK_mayor_igual);
						setState(142);
						expresion(12);
						}
						break;
					case 9:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(143);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(144);
						match(TK_menor_igual);
						setState(145);
						expresion(11);
						}
						break;
					case 10:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(146);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(147);
						match(TK_igualacion);
						setState(148);
						expresion(10);
						}
						break;
					case 11:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(149);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(150);
						match(TK_diferente);
						setState(151);
						expresion(9);
						}
						break;
					case 12:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(152);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(153);
						match(TK_or);
						setState(154);
						expresion(8);
						}
						break;
					case 13:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(155);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(156);
						match(TK_and);
						setState(157);
						expresion(7);
						}
						break;
					case 14:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(158);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(159);
						match(TK_sig_admiracion);
						setState(160);
						expresion(6);
						}
						break;
					}
					} 
				}
				setState(165);
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
			setState(179);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(166);
				match(TKR_println);
				setState(167);
				match(TK_par_apertura);
				setState(168);
				expresion(0);
				setState(169);
				match(TK_par_cierre);
				setState(170);
				match(TK_pcoma);
				fmt.Println("Impresion")
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(173);
				match(TKR_println);
				setState(174);
				match(TK_par_apertura);
				setState(175);
				expresion(0);
				setState(176);
				impresioncomas(0);
				setState(177);
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
			setState(182);
			match(TK_coma);
			setState(183);
			expresion(0);
			setState(184);
			match(TK_par_cierre);
			setState(185);
			match(TK_pcoma);
			fmt.Println("Impresion especial")
			}
			_ctx.stop = _input.LT(-1);
			setState(197);
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
					setState(188);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(189);
					match(TK_coma);
					setState(190);
					expresion(0);
					setState(191);
					match(TK_par_cierre);
					setState(192);
					match(TK_pcoma);
					fmt.Println("Impresion especial")
					}
					} 
				}
				setState(199);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3F\u00cb\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\3\2\7\2\30\n\2\f\2\16\2\33\13\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\5\3\'\n\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4E\n\4\3"+
		"\5\3\5\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\5\tp\n\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\7\t\u00a4\n\t\f\t\16\t\u00a7\13\t\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n\u00b6\n\n\3\13\3\13\3\13\3"+
		"\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\7\13\u00c6\n\13"+
		"\f\13\16\13\u00c9\13\13\3\13\2\4\20\24\f\2\4\6\b\n\f\16\20\22\24\2\4\4"+
		"\2#$(+\5\2\61\62ACEF\2\u00dd\2\31\3\2\2\2\4&\3\2\2\2\6D\3\2\2\2\bF\3\2"+
		"\2\2\nH\3\2\2\2\fK\3\2\2\2\16O\3\2\2\2\20o\3\2\2\2\22\u00b5\3\2\2\2\24"+
		"\u00b7\3\2\2\2\26\30\5\4\3\2\27\26\3\2\2\2\30\33\3\2\2\2\31\27\3\2\2\2"+
		"\31\32\3\2\2\2\32\34\3\2\2\2\33\31\3\2\2\2\34\35\7\2\2\3\35\3\3\2\2\2"+
		"\36\37\5\20\t\2\37 \b\3\1\2 \'\3\2\2\2!\'\5\22\n\2\"#\5\6\4\2#$\b\3\1"+
		"\2$\'\3\2\2\2%\'\5\f\7\2&\36\3\2\2\2&!\3\2\2\2&\"\3\2\2\2&%\3\2\2\2\'"+
		"\5\3\2\2\2()\7,\2\2)*\7-\2\2*+\7C\2\2+,\7\16\2\2,-\5\b\5\2-.\5\n\6\2."+
		"/\7\20\2\2/E\3\2\2\2\60\61\7,\2\2\61\62\7-\2\2\62\63\7C\2\2\63\64\5\n"+
		"\6\2\64\65\7\20\2\2\65E\3\2\2\2\66\67\7,\2\2\678\7C\2\289\7\16\2\29:\5"+
		"\b\5\2:;\5\n\6\2;<\7\20\2\2<E\3\2\2\2=>\7,\2\2>?\7C\2\2?@\7\26\2\2@A\5"+
		"\20\t\2AB\7\20\2\2BC\b\4\1\2CE\3\2\2\2D(\3\2\2\2D\60\3\2\2\2D\66\3\2\2"+
		"\2D=\3\2\2\2E\7\3\2\2\2FG\t\2\2\2G\t\3\2\2\2HI\7\26\2\2IJ\5\20\t\2J\13"+
		"\3\2\2\2KL\7C\2\2LM\5\n\6\2MN\7\20\2\2N\r\3\2\2\2OP\t\3\2\2P\17\3\2\2"+
		"\2QR\b\t\1\2RS\7\30\2\2Sp\5\20\t\27TU\7%\2\2UV\7\24\2\2VW\5\20\t\2WX\7"+
		"\17\2\2XY\5\20\t\2YZ\7\25\2\2Zp\3\2\2\2[\\\7\'\2\2\\]\7\24\2\2]^\5\20"+
		"\t\2^_\7\17\2\2_`\5\20\t\2`a\7\25\2\2ap\3\2\2\2bc\7\24\2\2cd\5\20\t\2"+
		"de\7\25\2\2ep\3\2\2\2fg\5\16\b\2gh\7/\2\2hi\7#\2\2ip\3\2\2\2jk\5\16\b"+
		"\2kl\7/\2\2lm\7$\2\2mp\3\2\2\2np\5\16\b\2oQ\3\2\2\2oT\3\2\2\2o[\3\2\2"+
		"\2ob\3\2\2\2of\3\2\2\2oj\3\2\2\2on\3\2\2\2p\u00a5\3\2\2\2qr\f\26\2\2r"+
		"s\7\27\2\2st\5\20\t\27tu\b\t\1\2u\u00a4\3\2\2\2vw\f\25\2\2wx\7\30\2\2"+
		"xy\5\20\t\26yz\b\t\1\2z\u00a4\3\2\2\2{|\f\24\2\2|}\7\31\2\2}~\5\20\t\25"+
		"~\177\b\t\1\2\177\u00a4\3\2\2\2\u0080\u0081\f\23\2\2\u0081\u0082\7\32"+
		"\2\2\u0082\u0083\5\20\t\24\u0083\u0084\b\t\1\2\u0084\u00a4\3\2\2\2\u0085"+
		"\u0086\f\20\2\2\u0086\u0087\7\33\2\2\u0087\u00a4\5\20\t\21\u0088\u0089"+
		"\f\17\2\2\u0089\u008a\7\21\2\2\u008a\u00a4\5\20\t\20\u008b\u008c\f\16"+
		"\2\2\u008c\u008d\7\22\2\2\u008d\u00a4\5\20\t\17\u008e\u008f\f\r\2\2\u008f"+
		"\u0090\7\b\2\2\u0090\u00a4\5\20\t\16\u0091\u0092\f\f\2\2\u0092\u0093\7"+
		"\t\2\2\u0093\u00a4\5\20\t\r\u0094\u0095\f\13\2\2\u0095\u0096\7\6\2\2\u0096"+
		"\u00a4\5\20\t\f\u0097\u0098\f\n\2\2\u0098\u0099\7\7\2\2\u0099\u00a4\5"+
		"\20\t\13\u009a\u009b\f\t\2\2\u009b\u009c\7\4\2\2\u009c\u00a4\5\20\t\n"+
		"\u009d\u009e\f\b\2\2\u009e\u009f\7\5\2\2\u009f\u00a4\5\20\t\t\u00a0\u00a1"+
		"\f\7\2\2\u00a1\u00a2\7\36\2\2\u00a2\u00a4\5\20\t\b\u00a3q\3\2\2\2\u00a3"+
		"v\3\2\2\2\u00a3{\3\2\2\2\u00a3\u0080\3\2\2\2\u00a3\u0085\3\2\2\2\u00a3"+
		"\u0088\3\2\2\2\u00a3\u008b\3\2\2\2\u00a3\u008e\3\2\2\2\u00a3\u0091\3\2"+
		"\2\2\u00a3\u0094\3\2\2\2\u00a3\u0097\3\2\2\2\u00a3\u009a\3\2\2\2\u00a3"+
		"\u009d\3\2\2\2\u00a3\u00a0\3\2\2\2\u00a4\u00a7\3\2\2\2\u00a5\u00a3\3\2"+
		"\2\2\u00a5\u00a6\3\2\2\2\u00a6\21\3\2\2\2\u00a7\u00a5\3\2\2\2\u00a8\u00a9"+
		"\7\60\2\2\u00a9\u00aa\7\24\2\2\u00aa\u00ab\5\20\t\2\u00ab\u00ac\7\25\2"+
		"\2\u00ac\u00ad\7\20\2\2\u00ad\u00ae\b\n\1\2\u00ae\u00b6\3\2\2\2\u00af"+
		"\u00b0\7\60\2\2\u00b0\u00b1\7\24\2\2\u00b1\u00b2\5\20\t\2\u00b2\u00b3"+
		"\5\24\13\2\u00b3\u00b4\7\20\2\2\u00b4\u00b6\3\2\2\2\u00b5\u00a8\3\2\2"+
		"\2\u00b5\u00af\3\2\2\2\u00b6\23\3\2\2\2\u00b7\u00b8\b\13\1\2\u00b8\u00b9"+
		"\7\17\2\2\u00b9\u00ba\5\20\t\2\u00ba\u00bb\7\25\2\2\u00bb\u00bc\7\20\2"+
		"\2\u00bc\u00bd\b\13\1\2\u00bd\u00c7\3\2\2\2\u00be\u00bf\f\4\2\2\u00bf"+
		"\u00c0\7\17\2\2\u00c0\u00c1\5\20\t\2\u00c1\u00c2\7\25\2\2\u00c2\u00c3"+
		"\7\20\2\2\u00c3\u00c4\b\13\1\2\u00c4\u00c6\3\2\2\2\u00c5\u00be\3\2\2\2"+
		"\u00c6\u00c9\3\2\2\2\u00c7\u00c5\3\2\2\2\u00c7\u00c8\3\2\2\2\u00c8\25"+
		"\3\2\2\2\u00c9\u00c7\3\2\2\2\n\31&Do\u00a3\u00a5\u00b5\u00c7";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}