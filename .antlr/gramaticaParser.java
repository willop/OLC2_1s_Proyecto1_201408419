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
		RULE_tipovariable = 4, RULE_valores = 5, RULE_expresion = 6, RULE_impresion = 7, 
		RULE_impresioncomas = 8;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "instrucciones", "instruccion", "declaracion", "tipovariable", 
			"valores", "expresion", "impresion", "impresioncomas"
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
			setState(18);
			instrucciones();
			setState(19);
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
			setState(24);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 18)) & ~0x3f) == 0 && ((1L << (_la - 18)) & ((1L << (TK_par_apertura - 18)) | (1L << (TK_menos - 18)) | (1L << (TKR_pow - 18)) | (1L << (TKR_powf - 18)) | (1L << (TKR_let - 18)) | (1L << (TKR_println - 18)) | (1L << (TKR_true - 18)) | (1L << (TKR_false - 18)) | (1L << (TK_entero - 18)) | (1L << (TK_decimal - 18)) | (1L << (TK_id - 18)) | (1L << (TK_cadena - 18)) | (1L << (TK_caracter - 18)))) != 0)) {
				{
				{
				setState(21);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).e.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(26);
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
		public Interfaces.Instruccion inst;
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
		public InstruccionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion; }
	}

	public final InstruccionContext instruccion() throws RecognitionException {
		InstruccionContext _localctx = new InstruccionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruccion);
		try {
			setState(34);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_par_apertura:
			case TK_menos:
			case TKR_pow:
			case TKR_powf:
			case TKR_true:
			case TKR_false:
			case TK_entero:
			case TK_decimal:
			case TK_id:
			case TK_cadena:
			case TK_caracter:
				enterOuterAlt(_localctx, 1);
				{
				setState(27);
				((InstruccionContext)_localctx).expresion = expresion(0);
				fmt.Println("mensaje en instrucciones: ",((InstruccionContext)_localctx).expresion.valorexpresion)
				}
				break;
			case TKR_println:
				enterOuterAlt(_localctx, 2);
				{
				setState(30);
				impresion();
				}
				break;
			case TKR_let:
				enterOuterAlt(_localctx, 3);
				{
				setState(31);
				((InstruccionContext)_localctx).declaracion = declaracion();
				fmt.Println("mensaje en declaracion: ",((InstruccionContext)_localctx).declaracion.decla)
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
			setState(70);
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
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(39);
				match(TK_dosPuntos);
				setState(40);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(41);
				match(TK_igual);
				setState(42);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(43);
				match(TK_pcoma);
				_localctx.decla = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar,((DeclaracionContext)_localctx).expresion.valorexpresion,true,false,false) 
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
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(49);
				match(TK_igual);
				setState(50);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(51);
				match(TK_pcoma);
				_localctx.decla = ((DeclaracionContext)_localctx).expresion.valorexpresion
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(54);
				match(TKR_let);
				setState(55);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(56);
				match(TK_dosPuntos);
				setState(57);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(58);
				match(TK_igual);
				setState(59);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(60);
				match(TK_pcoma);
				_localctx.decla = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar,((DeclaracionContext)_localctx).expresion.valorexpresion,false,false,false) 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(63);
				match(TKR_let);
				setState(64);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(65);
				match(TK_igual);
				setState(66);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(67);
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
			setState(84);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TKR_numericos_enteros:
				enterOuterAlt(_localctx, 1);
				{
				setState(72);
				match(TKR_numericos_enteros);
				_localctx.tipovar = Interfaces.INTEGER
				}
				break;
			case TKR_numericos_flotantes:
				enterOuterAlt(_localctx, 2);
				{
				setState(74);
				match(TKR_numericos_flotantes);
				_localctx.tipovar = Interfaces.FLOAT
				}
				break;
			case TKR_String:
				enterOuterAlt(_localctx, 3);
				{
				setState(76);
				match(TKR_String);
				_localctx.tipovar = Interfaces.STRING
				}
				break;
			case TKR_bool:
				enterOuterAlt(_localctx, 4);
				{
				setState(78);
				match(TKR_bool);
				_localctx.tipovar = Interfaces.BOOLEAN
				}
				break;
			case TKR_char:
				enterOuterAlt(_localctx, 5);
				{
				setState(80);
				match(TKR_char);
				_localctx.tipovar = Interfaces.CHAR
				}
				break;
			case TKR_usize:
				enterOuterAlt(_localctx, 6);
				{
				setState(82);
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
		enterRule(_localctx, 10, RULE_valores);
		try {
			setState(99);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_entero:
				enterOuterAlt(_localctx, 1);
				{
				setState(86);
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
				setState(88);
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
				setState(90);
				((ValoresContext)_localctx).TK_cadena = match(TK_cadena);

				                str:= (((ValoresContext)_localctx).TK_cadena!=null?((ValoresContext)_localctx).TK_cadena.getText():null)[1:len((((ValoresContext)_localctx).TK_cadena!=null?((ValoresContext)_localctx).TK_cadena.getText():null))-1]
				                _localctx.lit = Expresion.NuevoPrimitivo(str,Interfaces.STRING)
				}
				break;
			case TK_caracter:
				enterOuterAlt(_localctx, 4);
				{
				setState(92);
				((ValoresContext)_localctx).TK_caracter = match(TK_caracter);
				str:= (((ValoresContext)_localctx).TK_caracter!=null?((ValoresContext)_localctx).TK_caracter.getText():null)[1:len((((ValoresContext)_localctx).TK_caracter!=null?((ValoresContext)_localctx).TK_caracter.getText():null))-1]
				                _localctx.lit = Expresion.NuevoPrimitivo(str,Interfaces.CHAR)
				}
				break;
			case TKR_true:
				enterOuterAlt(_localctx, 5);
				{
				setState(94);
				match(TKR_true);
				_localctx.lit = Expresion.NuevoPrimitivo(true, Interfaces.BOOLEAN)
				}
				break;
			case TKR_false:
				enterOuterAlt(_localctx, 6);
				{
				setState(96);
				match(TKR_false);
				_localctx.lit = Expresion.NuevoPrimitivo(false,Interfaces.BOOLEAN)
				}
				break;
			case TK_id:
				enterOuterAlt(_localctx, 7);
				{
				setState(98);
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
		int _startState = 12;
		enterRecursionRule(_localctx, 12, RULE_expresion, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(133);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				{
				setState(102);
				match(TK_menos);
				setState(103);
				expresion(21);
				}
				break;
			case 2:
				{
				setState(104);
				match(TKR_pow);
				setState(105);
				match(TK_par_apertura);
				setState(106);
				expresion(0);
				setState(107);
				match(TK_coma);
				setState(108);
				expresion(0);
				setState(109);
				match(TK_par_cierre);
				}
				break;
			case 3:
				{
				setState(111);
				match(TKR_powf);
				setState(112);
				match(TK_par_apertura);
				setState(113);
				expresion(0);
				setState(114);
				match(TK_coma);
				setState(115);
				expresion(0);
				setState(116);
				match(TK_par_cierre);
				}
				break;
			case 4:
				{
				setState(118);
				match(TK_par_apertura);
				setState(119);
				expresion(0);
				setState(120);
				match(TK_par_cierre);
				}
				break;
			case 5:
				{
				setState(122);
				valores();
				setState(123);
				match(TKR_as);
				setState(124);
				match(TKR_numericos_enteros);
				}
				break;
			case 6:
				{
				setState(126);
				valores();
				setState(127);
				match(TKR_as);
				setState(128);
				match(TKR_numericos_flotantes);
				}
				break;
			case 7:
				{
				setState(130);
				((ExpresionContext)_localctx).vall = valores();
				_localctx.valorexpresion = ((ExpresionContext)_localctx).vall.lit
				                                                                                        fmt.Println(_localctx.valorexpresion)
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(179);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(177);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
					case 1:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(135);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(136);
						((ExpresionContext)_localctx).op = match(TK_por);
						setState(137);
						((ExpresionContext)_localctx).e2 = expresion(21);
						}
						break;
					case 2:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(138);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(139);
						((ExpresionContext)_localctx).op = match(TK_menos);
						setState(140);
						((ExpresionContext)_localctx).e2 = expresion(20);
						}
						break;
					case 3:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(141);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(142);
						((ExpresionContext)_localctx).op = match(TK_diagonal);
						setState(143);
						((ExpresionContext)_localctx).e2 = expresion(19);
						}
						break;
					case 4:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(144);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(145);
						((ExpresionContext)_localctx).op = match(TK_suma);
						setState(146);
						((ExpresionContext)_localctx).e2 = expresion(18);
						}
						break;
					case 5:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(147);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(148);
						match(TK_porcentaje);
						setState(149);
						expresion(15);
						}
						break;
					case 6:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(150);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(151);
						match(TK_menor);
						setState(152);
						expresion(14);
						}
						break;
					case 7:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(153);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(154);
						match(TK_mayor);
						setState(155);
						expresion(13);
						}
						break;
					case 8:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(156);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(157);
						match(TK_mayor_igual);
						setState(158);
						expresion(12);
						}
						break;
					case 9:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(159);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(160);
						match(TK_menor_igual);
						setState(161);
						expresion(11);
						}
						break;
					case 10:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(162);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(163);
						match(TK_igualacion);
						setState(164);
						expresion(10);
						}
						break;
					case 11:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(165);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(166);
						match(TK_diferente);
						setState(167);
						expresion(9);
						}
						break;
					case 12:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(168);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(169);
						match(TK_or);
						setState(170);
						expresion(8);
						}
						break;
					case 13:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(171);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(172);
						match(TK_and);
						setState(173);
						expresion(7);
						}
						break;
					case 14:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(174);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(175);
						match(TK_sig_admiracion);
						setState(176);
						expresion(6);
						}
						break;
					}
					} 
				}
				setState(181);
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
		public Interfaces.Instruccion impr;
		public ExpresionContext expresion;
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
		enterRule(_localctx, 14, RULE_impresion);
		try {
			setState(195);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(182);
				match(TKR_println);
				setState(183);
				match(TK_par_apertura);
				setState(184);
				((ImpresionContext)_localctx).expresion = expresion(0);
				setState(185);
				match(TK_par_cierre);
				setState(186);
				match(TK_pcoma);
				fmt.Println("Impresion")
				                                                                                                        _localctx.impr = Instruccion.NuevoPrint(((ImpresionContext)_localctx).expresion.valorexpresion)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(189);
				match(TKR_println);
				setState(190);
				match(TK_par_apertura);
				setState(191);
				expresion(0);
				setState(192);
				impresioncomas(0);
				setState(193);
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
		int _startState = 16;
		enterRecursionRule(_localctx, 16, RULE_impresioncomas, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(198);
			match(TK_coma);
			setState(199);
			expresion(0);
			setState(200);
			match(TK_par_cierre);
			setState(201);
			match(TK_pcoma);
			fmt.Println("Impresion especial")
			}
			_ctx.stop = _input.LT(-1);
			setState(213);
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
					setState(204);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(205);
					match(TK_coma);
					setState(206);
					expresion(0);
					setState(207);
					match(TK_par_cierre);
					setState(208);
					match(TK_pcoma);
					fmt.Println("Impresion especial")
					}
					} 
				}
				setState(215);
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
		case 6:
			return expresion_sempred((ExpresionContext)_localctx, predIndex);
		case 8:
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3F\u00db\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\3\2\3\2"+
		"\3\2\3\3\7\3\31\n\3\f\3\16\3\34\13\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4%"+
		"\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5"+
		"\5\5I\n\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6W\n\6\3\7"+
		"\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7f\n\7\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\5\b\u0088\n\b\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\7\b\u00b4\n\b\f\b\16\b\u00b7\13\b\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t\u00c6\n\t\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\7\n\u00d6\n\n\f\n\16\n\u00d9\13\n\3"+
		"\n\2\4\16\22\13\2\4\6\b\n\f\16\20\22\2\2\2\u00f8\2\24\3\2\2\2\4\32\3\2"+
		"\2\2\6$\3\2\2\2\bH\3\2\2\2\nV\3\2\2\2\fe\3\2\2\2\16\u0087\3\2\2\2\20\u00c5"+
		"\3\2\2\2\22\u00c7\3\2\2\2\24\25\5\4\3\2\25\26\7\2\2\3\26\3\3\2\2\2\27"+
		"\31\5\6\4\2\30\27\3\2\2\2\31\34\3\2\2\2\32\30\3\2\2\2\32\33\3\2\2\2\33"+
		"\5\3\2\2\2\34\32\3\2\2\2\35\36\5\16\b\2\36\37\b\4\1\2\37%\3\2\2\2 %\5"+
		"\20\t\2!\"\5\b\5\2\"#\b\4\1\2#%\3\2\2\2$\35\3\2\2\2$ \3\2\2\2$!\3\2\2"+
		"\2%\7\3\2\2\2&\'\7,\2\2\'(\7-\2\2()\7C\2\2)*\7\16\2\2*+\5\n\6\2+,\7\26"+
		"\2\2,-\5\16\b\2-.\7\20\2\2./\b\5\1\2/I\3\2\2\2\60\61\7,\2\2\61\62\7-\2"+
		"\2\62\63\7C\2\2\63\64\7\26\2\2\64\65\5\16\b\2\65\66\7\20\2\2\66\67\b\5"+
		"\1\2\67I\3\2\2\289\7,\2\29:\7C\2\2:;\7\16\2\2;<\5\n\6\2<=\7\26\2\2=>\5"+
		"\16\b\2>?\7\20\2\2?@\b\5\1\2@I\3\2\2\2AB\7,\2\2BC\7C\2\2CD\7\26\2\2DE"+
		"\5\16\b\2EF\7\20\2\2FG\b\5\1\2GI\3\2\2\2H&\3\2\2\2H\60\3\2\2\2H8\3\2\2"+
		"\2HA\3\2\2\2I\t\3\2\2\2JK\7#\2\2KW\b\6\1\2LM\7$\2\2MW\b\6\1\2NO\7*\2\2"+
		"OW\b\6\1\2PQ\7(\2\2QW\b\6\1\2RS\7)\2\2SW\b\6\1\2TU\7+\2\2UW\b\6\1\2VJ"+
		"\3\2\2\2VL\3\2\2\2VN\3\2\2\2VP\3\2\2\2VR\3\2\2\2VT\3\2\2\2W\13\3\2\2\2"+
		"XY\7A\2\2Yf\b\7\1\2Z[\7B\2\2[f\b\7\1\2\\]\7E\2\2]f\b\7\1\2^_\7F\2\2_f"+
		"\b\7\1\2`a\7\61\2\2af\b\7\1\2bc\7\62\2\2cf\b\7\1\2df\7C\2\2eX\3\2\2\2"+
		"eZ\3\2\2\2e\\\3\2\2\2e^\3\2\2\2e`\3\2\2\2eb\3\2\2\2ed\3\2\2\2f\r\3\2\2"+
		"\2gh\b\b\1\2hi\7\30\2\2i\u0088\5\16\b\27jk\7%\2\2kl\7\24\2\2lm\5\16\b"+
		"\2mn\7\17\2\2no\5\16\b\2op\7\25\2\2p\u0088\3\2\2\2qr\7\'\2\2rs\7\24\2"+
		"\2st\5\16\b\2tu\7\17\2\2uv\5\16\b\2vw\7\25\2\2w\u0088\3\2\2\2xy\7\24\2"+
		"\2yz\5\16\b\2z{\7\25\2\2{\u0088\3\2\2\2|}\5\f\7\2}~\7/\2\2~\177\7#\2\2"+
		"\177\u0088\3\2\2\2\u0080\u0081\5\f\7\2\u0081\u0082\7/\2\2\u0082\u0083"+
		"\7$\2\2\u0083\u0088\3\2\2\2\u0084\u0085\5\f\7\2\u0085\u0086\b\b\1\2\u0086"+
		"\u0088\3\2\2\2\u0087g\3\2\2\2\u0087j\3\2\2\2\u0087q\3\2\2\2\u0087x\3\2"+
		"\2\2\u0087|\3\2\2\2\u0087\u0080\3\2\2\2\u0087\u0084\3\2\2\2\u0088\u00b5"+
		"\3\2\2\2\u0089\u008a\f\26\2\2\u008a\u008b\7\31\2\2\u008b\u00b4\5\16\b"+
		"\27\u008c\u008d\f\25\2\2\u008d\u008e\7\30\2\2\u008e\u00b4\5\16\b\26\u008f"+
		"\u0090\f\24\2\2\u0090\u0091\7\32\2\2\u0091\u00b4\5\16\b\25\u0092\u0093"+
		"\f\23\2\2\u0093\u0094\7\27\2\2\u0094\u00b4\5\16\b\24\u0095\u0096\f\20"+
		"\2\2\u0096\u0097\7\33\2\2\u0097\u00b4\5\16\b\21\u0098\u0099\f\17\2\2\u0099"+
		"\u009a\7\21\2\2\u009a\u00b4\5\16\b\20\u009b\u009c\f\16\2\2\u009c\u009d"+
		"\7\22\2\2\u009d\u00b4\5\16\b\17\u009e\u009f\f\r\2\2\u009f\u00a0\7\b\2"+
		"\2\u00a0\u00b4\5\16\b\16\u00a1\u00a2\f\f\2\2\u00a2\u00a3\7\t\2\2\u00a3"+
		"\u00b4\5\16\b\r\u00a4\u00a5\f\13\2\2\u00a5\u00a6\7\6\2\2\u00a6\u00b4\5"+
		"\16\b\f\u00a7\u00a8\f\n\2\2\u00a8\u00a9\7\7\2\2\u00a9\u00b4\5\16\b\13"+
		"\u00aa\u00ab\f\t\2\2\u00ab\u00ac\7\4\2\2\u00ac\u00b4\5\16\b\n\u00ad\u00ae"+
		"\f\b\2\2\u00ae\u00af\7\5\2\2\u00af\u00b4\5\16\b\t\u00b0\u00b1\f\7\2\2"+
		"\u00b1\u00b2\7\36\2\2\u00b2\u00b4\5\16\b\b\u00b3\u0089\3\2\2\2\u00b3\u008c"+
		"\3\2\2\2\u00b3\u008f\3\2\2\2\u00b3\u0092\3\2\2\2\u00b3\u0095\3\2\2\2\u00b3"+
		"\u0098\3\2\2\2\u00b3\u009b\3\2\2\2\u00b3\u009e\3\2\2\2\u00b3\u00a1\3\2"+
		"\2\2\u00b3\u00a4\3\2\2\2\u00b3\u00a7\3\2\2\2\u00b3\u00aa\3\2\2\2\u00b3"+
		"\u00ad\3\2\2\2\u00b3\u00b0\3\2\2\2\u00b4\u00b7\3\2\2\2\u00b5\u00b3\3\2"+
		"\2\2\u00b5\u00b6\3\2\2\2\u00b6\17\3\2\2\2\u00b7\u00b5\3\2\2\2\u00b8\u00b9"+
		"\7\60\2\2\u00b9\u00ba\7\24\2\2\u00ba\u00bb\5\16\b\2\u00bb\u00bc\7\25\2"+
		"\2\u00bc\u00bd\7\20\2\2\u00bd\u00be\b\t\1\2\u00be\u00c6\3\2\2\2\u00bf"+
		"\u00c0\7\60\2\2\u00c0\u00c1\7\24\2\2\u00c1\u00c2\5\16\b\2\u00c2\u00c3"+
		"\5\22\n\2\u00c3\u00c4\7\20\2\2\u00c4\u00c6\3\2\2\2\u00c5\u00b8\3\2\2\2"+
		"\u00c5\u00bf\3\2\2\2\u00c6\21\3\2\2\2\u00c7\u00c8\b\n\1\2\u00c8\u00c9"+
		"\7\17\2\2\u00c9\u00ca\5\16\b\2\u00ca\u00cb\7\25\2\2\u00cb\u00cc\7\20\2"+
		"\2\u00cc\u00cd\b\n\1\2\u00cd\u00d7\3\2\2\2\u00ce\u00cf\f\4\2\2\u00cf\u00d0"+
		"\7\17\2\2\u00d0\u00d1\5\16\b\2\u00d1\u00d2\7\25\2\2\u00d2\u00d3\7\20\2"+
		"\2\u00d3\u00d4\b\n\1\2\u00d4\u00d6\3\2\2\2\u00d5\u00ce\3\2\2\2\u00d6\u00d9"+
		"\3\2\2\2\u00d7\u00d5\3\2\2\2\u00d7\u00d8\3\2\2\2\u00d8\23\3\2\2\2\u00d9"+
		"\u00d7\3\2\2\2\f\32$HVe\u0087\u00b3\u00b5\u00c5\u00d7";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}