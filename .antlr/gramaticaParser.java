// Generated from e:\WillOP\u005Cu\Compi2\Laboratorio\Proyecto1_201408419\gramatica.g4 by ANTLR 4.8

    import "proyecto1/Interfaces"
    import "proyecto1/Expresion"
    import "proyecto1/Instruccion"
    //import arrayList "github.com/colegno/arraylist"
    //import "proyecto1/Expresion"

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
		RULE_start = 0, RULE_instrucciones = 1, RULE_instruccion = 2, RULE_declaracion = 3, 
		RULE_tipovariable = 4, RULE_identificadores = 5, RULE_valores = 6, RULE_expresion = 7, 
		RULE_impresion = 8, RULE_impresioncomas = 9;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "instrucciones", "instruccion", "declaracion", "tipovariable", 
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


	        var temporal = Interfaces.SINTIPO

	public gramaticaParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode EOF() { return getToken(gramaticaParser.EOF, 0); }
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(20);
			instrucciones();
			setState(21);
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
		public InstruccionContext instruccion;
		public List<InstruccionContext> e = new ArrayList<InstruccionContext>();
		public List<InstruccionContext> instruccion() {
			return getRuleContexts(InstruccionContext.class);
		}
		public InstruccionContext instruccion(int i) {
			return getRuleContext(InstruccionContext.class,i);
		}
		public InstruccionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instrucciones; }
	}

	public final InstruccionesContext instrucciones() throws RecognitionException {
		InstruccionesContext _localctx = new InstruccionesContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_instrucciones);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(26);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 18)) & ~0x3f) == 0 && ((1L << (_la - 18)) & ((1L << (TK_par_apertura - 18)) | (1L << (TK_menos - 18)) | (1L << (TKR_pow - 18)) | (1L << (TKR_powf - 18)) | (1L << (TKR_let - 18)) | (1L << (TKR_println - 18)) | (1L << (TKR_true - 18)) | (1L << (TKR_false - 18)) | (1L << (TK_entero - 18)) | (1L << (TK_decimal - 18)) | (1L << (TK_id - 18)) | (1L << (TK_cadena - 18)) | (1L << (TK_caracter - 18)))) != 0)) {
				{
				{
				setState(23);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).e.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(28);
				_errHandler.sync(this);
				_la = _input.LA(1);
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

	public static class InstruccionContext extends ParserRuleContext {
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
		public InstruccionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion; }
	}

	public final InstruccionContext instruccion() throws RecognitionException {
		InstruccionContext _localctx = new InstruccionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruccion);
		try {
			setState(37);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(29);
				((InstruccionContext)_localctx).expresion = expresion(0);
				fmt.Println("mensaje en instrucciones: ",((InstruccionContext)_localctx).expresion.valorexpresion)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(32);
				impresion();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(33);
				((InstruccionContext)_localctx).declaracion = declaracion();
				fmt.Println("mensaje en declaracion: ",((InstruccionContext)_localctx).declaracion.decla)
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(36);
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
		public TipovariableContext tipovariable;
		public ExpresionContext expresion;
		public TerminalNode TKR_let() { return getToken(gramaticaParser.TKR_let, 0); }
		public TerminalNode TKR_mut() { return getToken(gramaticaParser.TKR_mut, 0); }
		public TerminalNode TK_dosPuntos() { return getToken(gramaticaParser.TK_dosPuntos, 0); }
		public TipovariableContext tipovariable() {
			return getRuleContext(TipovariableContext.class,0);
		}
		public TerminalNode TK_igual() { return getToken(gramaticaParser.TK_igual, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode TK_pcoma() { return getToken(gramaticaParser.TK_pcoma, 0); }
		public TerminalNode TK_id() { return getToken(gramaticaParser.TK_id, 0); }
		public DeclaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion; }
	}

	public final DeclaracionContext declaracion() throws RecognitionException {
		DeclaracionContext _localctx = new DeclaracionContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_declaracion);
		try {
			setState(73);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(39);
				match(TKR_let);
				setState(40);
				match(TKR_mut);
				setState(41);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(42);
				match(TK_dosPuntos);
				setState(43);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(44);
				match(TK_igual);
				setState(45);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(46);
				match(TK_pcoma);
				_localctx.decla = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar,((DeclaracionContext)_localctx).expresion.valorexpresion,true,false,false) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(49);
				match(TKR_let);
				setState(50);
				match(TKR_mut);
				setState(51);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(52);
				match(TK_igual);
				setState(53);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(54);
				match(TK_pcoma);
				_localctx.decla = ((DeclaracionContext)_localctx).expresion.valorexpresion
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(57);
				match(TKR_let);
				setState(58);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(59);
				match(TK_dosPuntos);
				setState(60);
				tipovariable();
				setState(61);
				match(TK_igual);
				setState(62);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(63);
				match(TK_pcoma);
				_localctx.decla = ((DeclaracionContext)_localctx).expresion.valorexpresion
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(66);
				match(TKR_let);
				setState(67);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(68);
				match(TK_igual);
				setState(69);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(70);
				match(TK_pcoma);
				_localctx.decla = ((DeclaracionContext)_localctx).expresion.valorexpresion
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
		public Interfaces.Tipoexpresion tipovar;
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
		enterRule(_localctx, 8, RULE_tipovariable);
		try {
			setState(87);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TKR_numericos_enteros:
				enterOuterAlt(_localctx, 1);
				{
				setState(75);
				match(TKR_numericos_enteros);
				_localctx.tipovar = Interfaces.INTEGER
				}
				break;
			case TKR_numericos_flotantes:
				enterOuterAlt(_localctx, 2);
				{
				setState(77);
				match(TKR_numericos_flotantes);
				_localctx.tipovar = Interfaces.FLOAT
				}
				break;
			case TKR_String:
				enterOuterAlt(_localctx, 3);
				{
				setState(79);
				match(TKR_String);
				_localctx.tipovar = Interfaces.STRING
				}
				break;
			case TKR_bool:
				enterOuterAlt(_localctx, 4);
				{
				setState(81);
				match(TKR_bool);
				_localctx.tipovar = Interfaces.BOOLEAN
				}
				break;
			case TKR_char:
				enterOuterAlt(_localctx, 5);
				{
				setState(83);
				match(TKR_char);
				_localctx.tipovar = Interfaces.CHAR
				}
				break;
			case TKR_usize:
				enterOuterAlt(_localctx, 6);
				{
				setState(85);
				match(TKR_usize);
				_localctx.tipovar = Interfaces.USIZE
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

	public static class IdentificadoresContext extends ParserRuleContext {
		public TerminalNode TK_id() { return getToken(gramaticaParser.TK_id, 0); }
		public TerminalNode TK_igual() { return getToken(gramaticaParser.TK_igual, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
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
			setState(89);
			match(TK_id);
			setState(90);
			match(TK_igual);
			setState(91);
			expresion(0);
			setState(92);
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
		public Interfaces.Expresion lit;
		public Token TK_entero;
		public Token TK_decimal;
		public Token TK_cadena;
		public Token TK_caracter;
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
		try {
			setState(107);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_entero:
				enterOuterAlt(_localctx, 1);
				{
				setState(94);
				((ValoresContext)_localctx).TK_entero = match(TK_entero);

				                numero, err := strconv.Atoi((((ValoresContext)_localctx).TK_entero!=null?((ValoresContext)_localctx).TK_entero.getText():null))
				                if err != nil {
				                        fmt.Println(err)
				                }
				                _localctx.lit = Expresion.NuevoPrimitivo(numero, Interfaces.INTEGER)
				}
				break;
			case TK_decimal:
				enterOuterAlt(_localctx, 2);
				{
				setState(96);
				((ValoresContext)_localctx).TK_decimal = match(TK_decimal);
				decimal, err := strconv.ParseFloat((((ValoresContext)_localctx).TK_decimal!=null?((ValoresContext)_localctx).TK_decimal.getText():null),8)
				                if err != nil {
				                        fmt.Println(err)
				                }
				                _localctx.lit = Expresion.NuevoPrimitivo(decimal, Interfaces.FLOAT)
				}
				break;
			case TK_cadena:
				enterOuterAlt(_localctx, 3);
				{
				setState(98);
				((ValoresContext)_localctx).TK_cadena = match(TK_cadena);

				                str:= (((ValoresContext)_localctx).TK_cadena!=null?((ValoresContext)_localctx).TK_cadena.getText():null)[1:len((((ValoresContext)_localctx).TK_cadena!=null?((ValoresContext)_localctx).TK_cadena.getText():null))-1]
				                _localctx.lit = Expresion.NuevoPrimitivo(str,Interfaces.STRING)
				}
				break;
			case TK_caracter:
				enterOuterAlt(_localctx, 4);
				{
				setState(100);
				((ValoresContext)_localctx).TK_caracter = match(TK_caracter);
				str:= (((ValoresContext)_localctx).TK_caracter!=null?((ValoresContext)_localctx).TK_caracter.getText():null)[1:len((((ValoresContext)_localctx).TK_caracter!=null?((ValoresContext)_localctx).TK_caracter.getText():null))-1]
				                _localctx.lit = Expresion.NuevoPrimitivo(str,Interfaces.CHAR)
				}
				break;
			case TKR_true:
				enterOuterAlt(_localctx, 5);
				{
				setState(102);
				match(TKR_true);
				_localctx.lit = Expresion.NuevoPrimitivo(true, Interfaces.BOOLEAN)
				}
				break;
			case TKR_false:
				enterOuterAlt(_localctx, 6);
				{
				setState(104);
				match(TKR_false);
				_localctx.lit = Expresion.NuevoPrimitivo(false,Interfaces.BOOLEAN)
				}
				break;
			case TK_id:
				enterOuterAlt(_localctx, 7);
				{
				setState(106);
				match(TK_id);
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

	public static class ExpresionContext extends ParserRuleContext {
		public Interfaces.Expresion valorexpresion;
		public ExpresionContext e1;
		public ValoresContext vall;
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
		public TerminalNode TK_por() { return getToken(gramaticaParser.TK_por, 0); }
		public TerminalNode TK_diagonal() { return getToken(gramaticaParser.TK_diagonal, 0); }
		public TerminalNode TK_suma() { return getToken(gramaticaParser.TK_suma, 0); }
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
			setState(141);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				{
				setState(110);
				match(TK_menos);
				setState(111);
				expresion(21);
				}
				break;
			case 2:
				{
				setState(112);
				match(TKR_pow);
				setState(113);
				match(TK_par_apertura);
				setState(114);
				expresion(0);
				setState(115);
				match(TK_coma);
				setState(116);
				expresion(0);
				setState(117);
				match(TK_par_cierre);
				}
				break;
			case 3:
				{
				setState(119);
				match(TKR_powf);
				setState(120);
				match(TK_par_apertura);
				setState(121);
				expresion(0);
				setState(122);
				match(TK_coma);
				setState(123);
				expresion(0);
				setState(124);
				match(TK_par_cierre);
				}
				break;
			case 4:
				{
				setState(126);
				match(TK_par_apertura);
				setState(127);
				expresion(0);
				setState(128);
				match(TK_par_cierre);
				}
				break;
			case 5:
				{
				setState(130);
				valores();
				setState(131);
				match(TKR_as);
				setState(132);
				match(TKR_numericos_enteros);
				}
				break;
			case 6:
				{
				setState(134);
				valores();
				setState(135);
				match(TKR_as);
				setState(136);
				match(TKR_numericos_flotantes);
				}
				break;
			case 7:
				{
				setState(138);
				((ExpresionContext)_localctx).vall = valores();
				_localctx.valorexpresion = ((ExpresionContext)_localctx).vall.lit
				                                                                                        fmt.Println(_localctx.valorexpresion)
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(187);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(185);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
					case 1:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(143);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(144);
						((ExpresionContext)_localctx).op = match(TK_por);
						setState(145);
						((ExpresionContext)_localctx).e2 = expresion(21);
						}
						break;
					case 2:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(146);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(147);
						((ExpresionContext)_localctx).op = match(TK_menos);
						setState(148);
						((ExpresionContext)_localctx).e2 = expresion(20);
						}
						break;
					case 3:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(149);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(150);
						((ExpresionContext)_localctx).op = match(TK_diagonal);
						setState(151);
						((ExpresionContext)_localctx).e2 = expresion(19);
						}
						break;
					case 4:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(152);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(153);
						((ExpresionContext)_localctx).op = match(TK_suma);
						setState(154);
						((ExpresionContext)_localctx).e2 = expresion(18);
						}
						break;
					case 5:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(155);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(156);
						match(TK_porcentaje);
						setState(157);
						expresion(15);
						}
						break;
					case 6:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(158);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(159);
						match(TK_menor);
						setState(160);
						expresion(14);
						}
						break;
					case 7:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(161);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(162);
						match(TK_mayor);
						setState(163);
						expresion(13);
						}
						break;
					case 8:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(164);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(165);
						match(TK_mayor_igual);
						setState(166);
						expresion(12);
						}
						break;
					case 9:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(167);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(168);
						match(TK_menor_igual);
						setState(169);
						expresion(11);
						}
						break;
					case 10:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(170);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(171);
						match(TK_igualacion);
						setState(172);
						expresion(10);
						}
						break;
					case 11:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(173);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(174);
						match(TK_diferente);
						setState(175);
						expresion(9);
						}
						break;
					case 12:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(176);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(177);
						match(TK_or);
						setState(178);
						expresion(8);
						}
						break;
					case 13:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(179);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(180);
						match(TK_and);
						setState(181);
						expresion(7);
						}
						break;
					case 14:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(182);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(183);
						match(TK_sig_admiracion);
						setState(184);
						expresion(6);
						}
						break;
					}
					} 
				}
				setState(189);
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
			setState(203);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(190);
				match(TKR_println);
				setState(191);
				match(TK_par_apertura);
				setState(192);
				expresion(0);
				setState(193);
				match(TK_par_cierre);
				setState(194);
				match(TK_pcoma);
				fmt.Println("Impresion")
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(197);
				match(TKR_println);
				setState(198);
				match(TK_par_apertura);
				setState(199);
				expresion(0);
				setState(200);
				impresioncomas(0);
				setState(201);
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
			setState(206);
			match(TK_coma);
			setState(207);
			expresion(0);
			setState(208);
			match(TK_par_cierre);
			setState(209);
			match(TK_pcoma);
			fmt.Println("Impresion especial")
			}
			_ctx.stop = _input.LT(-1);
			setState(221);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ImpresioncomasContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_impresioncomas);
					setState(212);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(213);
					match(TK_coma);
					setState(214);
					expresion(0);
					setState(215);
					match(TK_par_cierre);
					setState(216);
					match(TK_pcoma);
					fmt.Println("Impresion especial")
					}
					} 
				}
				setState(223);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3F\u00e3\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\3\2\3\2\3\2\3\3\7\3\33\n\3\f\3\16\3\36\13\3\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\5\4(\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5"+
		"\3\5\3\5\3\5\3\5\5\5L\n\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\5\6Z\n\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\5\bn\n\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\5\t\u0090\n\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\7\t\u00bc\n\t"+
		"\f\t\16\t\u00bf\13\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\5\n\u00ce\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\7\13\u00de\n\13\f\13\16\13\u00e1\13\13\3\13\2\4\20\24"+
		"\f\2\4\6\b\n\f\16\20\22\24\2\2\2\u0100\2\26\3\2\2\2\4\34\3\2\2\2\6\'\3"+
		"\2\2\2\bK\3\2\2\2\nY\3\2\2\2\f[\3\2\2\2\16m\3\2\2\2\20\u008f\3\2\2\2\22"+
		"\u00cd\3\2\2\2\24\u00cf\3\2\2\2\26\27\5\4\3\2\27\30\7\2\2\3\30\3\3\2\2"+
		"\2\31\33\5\6\4\2\32\31\3\2\2\2\33\36\3\2\2\2\34\32\3\2\2\2\34\35\3\2\2"+
		"\2\35\5\3\2\2\2\36\34\3\2\2\2\37 \5\20\t\2 !\b\4\1\2!(\3\2\2\2\"(\5\22"+
		"\n\2#$\5\b\5\2$%\b\4\1\2%(\3\2\2\2&(\5\f\7\2\'\37\3\2\2\2\'\"\3\2\2\2"+
		"\'#\3\2\2\2\'&\3\2\2\2(\7\3\2\2\2)*\7,\2\2*+\7-\2\2+,\7C\2\2,-\7\16\2"+
		"\2-.\5\n\6\2./\7\26\2\2/\60\5\20\t\2\60\61\7\20\2\2\61\62\b\5\1\2\62L"+
		"\3\2\2\2\63\64\7,\2\2\64\65\7-\2\2\65\66\7C\2\2\66\67\7\26\2\2\678\5\20"+
		"\t\289\7\20\2\29:\b\5\1\2:L\3\2\2\2;<\7,\2\2<=\7C\2\2=>\7\16\2\2>?\5\n"+
		"\6\2?@\7\26\2\2@A\5\20\t\2AB\7\20\2\2BC\b\5\1\2CL\3\2\2\2DE\7,\2\2EF\7"+
		"C\2\2FG\7\26\2\2GH\5\20\t\2HI\7\20\2\2IJ\b\5\1\2JL\3\2\2\2K)\3\2\2\2K"+
		"\63\3\2\2\2K;\3\2\2\2KD\3\2\2\2L\t\3\2\2\2MN\7#\2\2NZ\b\6\1\2OP\7$\2\2"+
		"PZ\b\6\1\2QR\7*\2\2RZ\b\6\1\2ST\7(\2\2TZ\b\6\1\2UV\7)\2\2VZ\b\6\1\2WX"+
		"\7+\2\2XZ\b\6\1\2YM\3\2\2\2YO\3\2\2\2YQ\3\2\2\2YS\3\2\2\2YU\3\2\2\2YW"+
		"\3\2\2\2Z\13\3\2\2\2[\\\7C\2\2\\]\7\26\2\2]^\5\20\t\2^_\7\20\2\2_\r\3"+
		"\2\2\2`a\7A\2\2an\b\b\1\2bc\7B\2\2cn\b\b\1\2de\7E\2\2en\b\b\1\2fg\7F\2"+
		"\2gn\b\b\1\2hi\7\61\2\2in\b\b\1\2jk\7\62\2\2kn\b\b\1\2ln\7C\2\2m`\3\2"+
		"\2\2mb\3\2\2\2md\3\2\2\2mf\3\2\2\2mh\3\2\2\2mj\3\2\2\2ml\3\2\2\2n\17\3"+
		"\2\2\2op\b\t\1\2pq\7\30\2\2q\u0090\5\20\t\27rs\7%\2\2st\7\24\2\2tu\5\20"+
		"\t\2uv\7\17\2\2vw\5\20\t\2wx\7\25\2\2x\u0090\3\2\2\2yz\7\'\2\2z{\7\24"+
		"\2\2{|\5\20\t\2|}\7\17\2\2}~\5\20\t\2~\177\7\25\2\2\177\u0090\3\2\2\2"+
		"\u0080\u0081\7\24\2\2\u0081\u0082\5\20\t\2\u0082\u0083\7\25\2\2\u0083"+
		"\u0090\3\2\2\2\u0084\u0085\5\16\b\2\u0085\u0086\7/\2\2\u0086\u0087\7#"+
		"\2\2\u0087\u0090\3\2\2\2\u0088\u0089\5\16\b\2\u0089\u008a\7/\2\2\u008a"+
		"\u008b\7$\2\2\u008b\u0090\3\2\2\2\u008c\u008d\5\16\b\2\u008d\u008e\b\t"+
		"\1\2\u008e\u0090\3\2\2\2\u008fo\3\2\2\2\u008fr\3\2\2\2\u008fy\3\2\2\2"+
		"\u008f\u0080\3\2\2\2\u008f\u0084\3\2\2\2\u008f\u0088\3\2\2\2\u008f\u008c"+
		"\3\2\2\2\u0090\u00bd\3\2\2\2\u0091\u0092\f\26\2\2\u0092\u0093\7\31\2\2"+
		"\u0093\u00bc\5\20\t\27\u0094\u0095\f\25\2\2\u0095\u0096\7\30\2\2\u0096"+
		"\u00bc\5\20\t\26\u0097\u0098\f\24\2\2\u0098\u0099\7\32\2\2\u0099\u00bc"+
		"\5\20\t\25\u009a\u009b\f\23\2\2\u009b\u009c\7\27\2\2\u009c\u00bc\5\20"+
		"\t\24\u009d\u009e\f\20\2\2\u009e\u009f\7\33\2\2\u009f\u00bc\5\20\t\21"+
		"\u00a0\u00a1\f\17\2\2\u00a1\u00a2\7\21\2\2\u00a2\u00bc\5\20\t\20\u00a3"+
		"\u00a4\f\16\2\2\u00a4\u00a5\7\22\2\2\u00a5\u00bc\5\20\t\17\u00a6\u00a7"+
		"\f\r\2\2\u00a7\u00a8\7\b\2\2\u00a8\u00bc\5\20\t\16\u00a9\u00aa\f\f\2\2"+
		"\u00aa\u00ab\7\t\2\2\u00ab\u00bc\5\20\t\r\u00ac\u00ad\f\13\2\2\u00ad\u00ae"+
		"\7\6\2\2\u00ae\u00bc\5\20\t\f\u00af\u00b0\f\n\2\2\u00b0\u00b1\7\7\2\2"+
		"\u00b1\u00bc\5\20\t\13\u00b2\u00b3\f\t\2\2\u00b3\u00b4\7\4\2\2\u00b4\u00bc"+
		"\5\20\t\n\u00b5\u00b6\f\b\2\2\u00b6\u00b7\7\5\2\2\u00b7\u00bc\5\20\t\t"+
		"\u00b8\u00b9\f\7\2\2\u00b9\u00ba\7\36\2\2\u00ba\u00bc\5\20\t\b\u00bb\u0091"+
		"\3\2\2\2\u00bb\u0094\3\2\2\2\u00bb\u0097\3\2\2\2\u00bb\u009a\3\2\2\2\u00bb"+
		"\u009d\3\2\2\2\u00bb\u00a0\3\2\2\2\u00bb\u00a3\3\2\2\2\u00bb\u00a6\3\2"+
		"\2\2\u00bb\u00a9\3\2\2\2\u00bb\u00ac\3\2\2\2\u00bb\u00af\3\2\2\2\u00bb"+
		"\u00b2\3\2\2\2\u00bb\u00b5\3\2\2\2\u00bb\u00b8\3\2\2\2\u00bc\u00bf\3\2"+
		"\2\2\u00bd\u00bb\3\2\2\2\u00bd\u00be\3\2\2\2\u00be\21\3\2\2\2\u00bf\u00bd"+
		"\3\2\2\2\u00c0\u00c1\7\60\2\2\u00c1\u00c2\7\24\2\2\u00c2\u00c3\5\20\t"+
		"\2\u00c3\u00c4\7\25\2\2\u00c4\u00c5\7\20\2\2\u00c5\u00c6\b\n\1\2\u00c6"+
		"\u00ce\3\2\2\2\u00c7\u00c8\7\60\2\2\u00c8\u00c9\7\24\2\2\u00c9\u00ca\5"+
		"\20\t\2\u00ca\u00cb\5\24\13\2\u00cb\u00cc\7\20\2\2\u00cc\u00ce\3\2\2\2"+
		"\u00cd\u00c0\3\2\2\2\u00cd\u00c7\3\2\2\2\u00ce\23\3\2\2\2\u00cf\u00d0"+
		"\b\13\1\2\u00d0\u00d1\7\17\2\2\u00d1\u00d2\5\20\t\2\u00d2\u00d3\7\25\2"+
		"\2\u00d3\u00d4\7\20\2\2\u00d4\u00d5\b\13\1\2\u00d5\u00df\3\2\2\2\u00d6"+
		"\u00d7\f\4\2\2\u00d7\u00d8\7\17\2\2\u00d8\u00d9\5\20\t\2\u00d9\u00da\7"+
		"\25\2\2\u00da\u00db\7\20\2\2\u00db\u00dc\b\13\1\2\u00dc\u00de\3\2\2\2"+
		"\u00dd\u00d6\3\2\2\2\u00de\u00e1\3\2\2\2\u00df\u00dd\3\2\2\2\u00df\u00e0"+
		"\3\2\2\2\u00e0\25\3\2\2\2\u00e1\u00df\3\2\2\2\f\34\'KYm\u008f\u00bb\u00bd"+
		"\u00cd\u00df";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}