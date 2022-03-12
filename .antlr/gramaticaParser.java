// Generated from e:\WillOP\u005Cu\Compi2\Laboratorio\Proyecto1_201408419\gramatica.g4 by ANTLR 4.8

    import "proyecto1/Interfaces"
    import "proyecto1/Expresion"
    import "proyecto1/Instruccion"
    import "proyecto1/Expresiones"
    import arrayList "github.com/colegno/arraylist"
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
		TKR_char=39, TKR_Str=40, TKR_String=41, TKR_usize=42, TKR_let=43, TKR_mut=44, 
		TKR_struct=45, TKR_as=46, TKR_println=47, TKR_true=48, TKR_false=49, TKR_fn=50, 
		TKR_return=51, TKR_abs=52, TKR_sqrt=53, TKR_to_string=54, TKR_clone=55, 
		TKR_new=56, TKR_len=57, TKR_push=58, TKR_remove=59, TKR_contains=60, TKR_insert=61, 
		TKR_capacity=62, TKR_with_capacity=63, TKR_main=64, TK_entero=65, TK_decimal=66, 
		TK_id=67, Letra=68, TK_cadena=69, TK_caracter=70;
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
			"'bool'", "'char'", "'&str'", "'String'", "'usize'", "'let'", "'mut'", 
			"'struct'", "'as'", "'println!'", "'true'", "'false'", "'fn'", "'return'", 
			"'abs'", "'sqrt'", "'to_string()'", "'clone()'", "'new'", "'len'", "'push'", 
			"'remove'", "'contains'", "'insert'", "'capacity'", "'witch_capacity'", 
			"'main'"
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
			"TKR_char", "TKR_Str", "TKR_String", "TKR_usize", "TKR_let", "TKR_mut", 
			"TKR_struct", "TKR_as", "TKR_println", "TKR_true", "TKR_false", "TKR_fn", 
			"TKR_return", "TKR_abs", "TKR_sqrt", "TKR_to_string", "TKR_clone", "TKR_new", 
			"TKR_len", "TKR_push", "TKR_remove", "TKR_contains", "TKR_insert", "TKR_capacity", 
			"TKR_with_capacity", "TKR_main", "TK_entero", "TK_decimal", "TK_id", 
			"Letra", "TK_cadena", "TK_caracter"
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
		public *arrayList.List listainstrucciones;
		public InstruccionesContext instrucciones;
		public TerminalNode TKR_fn() { return getToken(gramaticaParser.TKR_fn, 0); }
		public TerminalNode TKR_main() { return getToken(gramaticaParser.TKR_main, 0); }
		public TerminalNode TK_par_apertura() { return getToken(gramaticaParser.TK_par_apertura, 0); }
		public TerminalNode TK_par_cierre() { return getToken(gramaticaParser.TK_par_cierre, 0); }
		public TerminalNode TK_corchete_apertura() { return getToken(gramaticaParser.TK_corchete_apertura, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode TK_corchete_cierre() { return getToken(gramaticaParser.TK_corchete_cierre, 0); }
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
			match(TKR_fn);
			setState(19);
			match(TKR_main);
			setState(20);
			match(TK_par_apertura);
			setState(21);
			match(TK_par_cierre);
			setState(22);
			match(TK_corchete_apertura);
			setState(23);
			((StartContext)_localctx).instrucciones = instrucciones();
			setState(24);
			match(TK_corchete_cierre);
			_localctx.listainstrucciones = ((StartContext)_localctx).instrucciones.lista
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
		public *arrayList.List lista;
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

		        _localctx.lista = arrayList.New()

		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(30);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 18)) & ~0x3f) == 0 && ((1L << (_la - 18)) & ((1L << (TK_par_apertura - 18)) | (1L << (TK_menos - 18)) | (1L << (TK_sig_admiracion - 18)) | (1L << (TKR_pow - 18)) | (1L << (TKR_powf - 18)) | (1L << (TKR_Str - 18)) | (1L << (TKR_let - 18)) | (1L << (TKR_println - 18)) | (1L << (TKR_true - 18)) | (1L << (TKR_false - 18)) | (1L << (TK_entero - 18)) | (1L << (TK_decimal - 18)) | (1L << (TK_id - 18)) | (1L << (TK_cadena - 18)) | (1L << (TK_caracter - 18)))) != 0)) {
				{
				{
				setState(27);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).e.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(32);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			listaInst := localctx.(*InstruccionesContext).GetE() 
			                                    for _, e:= range listaInst {
			                                        _localctx.lista.Add(e.GetInst())
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
		public ImpresionContext impresion;
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
			setState(44);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_par_apertura:
			case TK_menos:
			case TK_sig_admiracion:
			case TKR_pow:
			case TKR_powf:
			case TKR_Str:
			case TKR_true:
			case TKR_false:
			case TK_entero:
			case TK_decimal:
			case TK_id:
			case TK_cadena:
			case TK_caracter:
				enterOuterAlt(_localctx, 1);
				{
				setState(35);
				((InstruccionContext)_localctx).expresion = expresion(0);
				fmt.Println("mensaje en instrucciones: ",((InstruccionContext)_localctx).expresion.exp)
				}
				break;
			case TKR_println:
				enterOuterAlt(_localctx, 2);
				{
				setState(38);
				((InstruccionContext)_localctx).impresion = impresion();
				_localctx.inst = ((InstruccionContext)_localctx).impresion.inst
				}
				break;
			case TKR_let:
				enterOuterAlt(_localctx, 3);
				{
				setState(41);
				((InstruccionContext)_localctx).declaracion = declaracion();
				_localctx.inst = ((InstruccionContext)_localctx).declaracion.inst
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
		public Interfaces.Instruccion inst;
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
			setState(80);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(46);
				match(TKR_let);
				setState(47);
				match(TKR_mut);
				setState(48);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(49);
				match(TK_dosPuntos);
				setState(50);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(51);
				match(TK_igual);
				setState(52);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(53);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar,((DeclaracionContext)_localctx).expresion.exp,false,false,true) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(56);
				match(TKR_let);
				setState(57);
				match(TKR_mut);
				setState(58);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(59);
				match(TK_igual);
				setState(60);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(61);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),Interfaces.SINTIPO,((DeclaracionContext)_localctx).expresion.exp,false,false,true) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(64);
				match(TKR_let);
				setState(65);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(66);
				match(TK_dosPuntos);
				setState(67);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(68);
				match(TK_igual);
				setState(69);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(70);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar,((DeclaracionContext)_localctx).expresion.exp,false,false,false) 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(73);
				match(TKR_let);
				setState(74);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(75);
				match(TK_igual);
				setState(76);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(77);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),Interfaces.SINTIPO,((DeclaracionContext)_localctx).expresion.exp,false,false,false) 
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
			setState(94);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TKR_numericos_enteros:
				enterOuterAlt(_localctx, 1);
				{
				setState(82);
				match(TKR_numericos_enteros);
				_localctx.tipovar = Interfaces.INTEGER
				}
				break;
			case TKR_numericos_flotantes:
				enterOuterAlt(_localctx, 2);
				{
				setState(84);
				match(TKR_numericos_flotantes);
				_localctx.tipovar = Interfaces.FLOAT
				}
				break;
			case TKR_String:
				enterOuterAlt(_localctx, 3);
				{
				setState(86);
				match(TKR_String);
				_localctx.tipovar = Interfaces.STRING
				}
				break;
			case TKR_bool:
				enterOuterAlt(_localctx, 4);
				{
				setState(88);
				match(TKR_bool);
				_localctx.tipovar = Interfaces.BOOLEAN
				}
				break;
			case TKR_char:
				enterOuterAlt(_localctx, 5);
				{
				setState(90);
				match(TKR_char);
				_localctx.tipovar = Interfaces.CHAR
				}
				break;
			case TKR_usize:
				enterOuterAlt(_localctx, 6);
				{
				setState(92);
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
		public Interfaces.Expresion exp;
		public Token TK_entero;
		public Token TK_decimal;
		public Token TK_cadena;
		public Token TK_caracter;
		public Token TKR_Str;
		public Token TK_id;
		public TerminalNode TK_entero() { return getToken(gramaticaParser.TK_entero, 0); }
		public TerminalNode TK_decimal() { return getToken(gramaticaParser.TK_decimal, 0); }
		public TerminalNode TK_cadena() { return getToken(gramaticaParser.TK_cadena, 0); }
		public TerminalNode TK_caracter() { return getToken(gramaticaParser.TK_caracter, 0); }
		public TerminalNode TKR_Str() { return getToken(gramaticaParser.TKR_Str, 0); }
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
			setState(112);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_entero:
				enterOuterAlt(_localctx, 1);
				{
				setState(96);
				((ValoresContext)_localctx).TK_entero = match(TK_entero);

				                numero, err := strconv.Atoi((((ValoresContext)_localctx).TK_entero!=null?((ValoresContext)_localctx).TK_entero.getText():null))
				                if err != nil {
				                        fmt.Println(err)
				                }
				                _localctx.exp = Expresion.NuevoPrimitivo(numero, Interfaces.INTEGER)
				}
				break;
			case TK_decimal:
				enterOuterAlt(_localctx, 2);
				{
				setState(98);
				((ValoresContext)_localctx).TK_decimal = match(TK_decimal);
				decimal, err := strconv.ParseFloat((((ValoresContext)_localctx).TK_decimal!=null?((ValoresContext)_localctx).TK_decimal.getText():null),8)
				                if err != nil {
				                        fmt.Println(err)
				                }
				                _localctx.exp = Expresion.NuevoPrimitivo(decimal, Interfaces.FLOAT)
				}
				break;
			case TK_cadena:
				enterOuterAlt(_localctx, 3);
				{
				setState(100);
				((ValoresContext)_localctx).TK_cadena = match(TK_cadena);

				                str:= (((ValoresContext)_localctx).TK_cadena!=null?((ValoresContext)_localctx).TK_cadena.getText():null)[1:len((((ValoresContext)_localctx).TK_cadena!=null?((ValoresContext)_localctx).TK_cadena.getText():null))-1]
				                _localctx.exp = Expresion.NuevoPrimitivo(str,Interfaces.STRING)
				}
				break;
			case TK_caracter:
				enterOuterAlt(_localctx, 4);
				{
				setState(102);
				((ValoresContext)_localctx).TK_caracter = match(TK_caracter);
				str:= (((ValoresContext)_localctx).TK_caracter!=null?((ValoresContext)_localctx).TK_caracter.getText():null)[1:len((((ValoresContext)_localctx).TK_caracter!=null?((ValoresContext)_localctx).TK_caracter.getText():null))-1]
				                _localctx.exp = Expresion.NuevoPrimitivo(str,Interfaces.CHAR)
				}
				break;
			case TKR_Str:
				enterOuterAlt(_localctx, 5);
				{
				setState(104);
				((ValoresContext)_localctx).TKR_Str = match(TKR_Str);
				str:= (((ValoresContext)_localctx).TKR_Str!=null?((ValoresContext)_localctx).TKR_Str.getText():null)[1:len((((ValoresContext)_localctx).TKR_Str!=null?((ValoresContext)_localctx).TKR_Str.getText():null))-1]
				                _localctx.exp = Expresion.NuevoPrimitivo(str,Interfaces.STR)
				}
				break;
			case TKR_true:
				enterOuterAlt(_localctx, 6);
				{
				setState(106);
				match(TKR_true);
				_localctx.exp = Expresion.NuevoPrimitivo(true, Interfaces.BOOLEAN)
				}
				break;
			case TKR_false:
				enterOuterAlt(_localctx, 7);
				{
				setState(108);
				match(TKR_false);
				_localctx.exp = Expresion.NuevoPrimitivo(false,Interfaces.BOOLEAN)
				}
				break;
			case TK_id:
				enterOuterAlt(_localctx, 8);
				{
				setState(110);
				((ValoresContext)_localctx).TK_id = match(TK_id);
				_localctx.exp = Expresiones.NewLlamarvariable((((ValoresContext)_localctx).TK_id!=null?((ValoresContext)_localctx).TK_id.getText():null))
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
		public Interfaces.Expresion exp;
		public ExpresionContext e1;
		public ExpresionContext e2;
		public ValoresContext vall;
		public Token op;
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
		public TerminalNode TK_sig_admiracion() { return getToken(gramaticaParser.TK_sig_admiracion, 0); }
		public ValoresContext valores() {
			return getRuleContext(ValoresContext.class,0);
		}
		public TerminalNode TKR_as() { return getToken(gramaticaParser.TKR_as, 0); }
		public TerminalNode TKR_numericos_enteros() { return getToken(gramaticaParser.TKR_numericos_enteros, 0); }
		public TerminalNode TKR_numericos_flotantes() { return getToken(gramaticaParser.TKR_numericos_flotantes, 0); }
		public TerminalNode TK_por() { return getToken(gramaticaParser.TK_por, 0); }
		public TerminalNode TK_diagonal() { return getToken(gramaticaParser.TK_diagonal, 0); }
		public TerminalNode TK_porcentaje() { return getToken(gramaticaParser.TK_porcentaje, 0); }
		public TerminalNode TK_suma() { return getToken(gramaticaParser.TK_suma, 0); }
		public TerminalNode TK_menor() { return getToken(gramaticaParser.TK_menor, 0); }
		public TerminalNode TK_mayor() { return getToken(gramaticaParser.TK_mayor, 0); }
		public TerminalNode TK_mayor_igual() { return getToken(gramaticaParser.TK_mayor_igual, 0); }
		public TerminalNode TK_menor_igual() { return getToken(gramaticaParser.TK_menor_igual, 0); }
		public TerminalNode TK_igualacion() { return getToken(gramaticaParser.TK_igualacion, 0); }
		public TerminalNode TK_diferente() { return getToken(gramaticaParser.TK_diferente, 0); }
		public TerminalNode TK_or() { return getToken(gramaticaParser.TK_or, 0); }
		public TerminalNode TK_and() { return getToken(gramaticaParser.TK_and, 0); }
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
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(154);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				{
				setState(115);
				match(TK_menos);
				setState(116);
				((ExpresionContext)_localctx).e1 = expresion(18);
				numero := -1
				                                                                                                        numm := Expresion.NuevoPrimitivo(numero, Interfaces.INTEGER)
				                                                                                                                _localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp, numm ,Interfaces.MULTIPLICACION)
				                                                                                                
				}
				break;
			case 2:
				{
				setState(119);
				match(TKR_pow);
				setState(120);
				match(TK_par_apertura);
				setState(121);
				((ExpresionContext)_localctx).e1 = expresion(0);
				setState(122);
				match(TK_coma);
				setState(123);
				((ExpresionContext)_localctx).e2 = expresion(0);
				setState(124);
				match(TK_par_cierre);
				_localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.POW)
				}
				break;
			case 3:
				{
				setState(127);
				match(TKR_powf);
				setState(128);
				match(TK_par_apertura);
				setState(129);
				((ExpresionContext)_localctx).e1 = expresion(0);
				setState(130);
				match(TK_coma);
				setState(131);
				((ExpresionContext)_localctx).e2 = expresion(0);
				setState(132);
				match(TK_par_cierre);
				_localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.POWF)
				}
				break;
			case 4:
				{
				setState(135);
				match(TK_sig_admiracion);
				setState(136);
				((ExpresionContext)_localctx).e1 = expresion(5);
				_localctx.exp = Expresiones.NuevaLogica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e1.exp,Interfaces.NOT)
				}
				break;
			case 5:
				{
				setState(139);
				match(TK_par_apertura);
				setState(140);
				expresion(0);
				setState(141);
				match(TK_par_cierre);
				}
				break;
			case 6:
				{
				setState(143);
				valores();
				setState(144);
				match(TKR_as);
				setState(145);
				match(TKR_numericos_enteros);
				}
				break;
			case 7:
				{
				setState(147);
				valores();
				setState(148);
				match(TKR_as);
				setState(149);
				match(TKR_numericos_flotantes);
				}
				break;
			case 8:
				{
				setState(151);
				((ExpresionContext)_localctx).vall = valores();
				_localctx.exp = ((ExpresionContext)_localctx).vall.exp
				                                                                                        fmt.Println(_localctx.exp)
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(208);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(206);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
					case 1:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(156);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(157);
						((ExpresionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << TK_por) | (1L << TK_diagonal) | (1L << TK_porcentaje))) != 0)) ) {
							((ExpresionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(158);
						((ExpresionContext)_localctx).e2 = expresion(18);

						                                                                                                                  if (((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null) == "*" {
						                                                                                                                          _localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.MULTIPLICACION)
						                                                                                                                  }else if (((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null) == "%" {
						                                                                                                                          _localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.MODULO)
						                                                                                                                  }else{
						                                                                                                                          _localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.DIVISION)
						                                                                                                                  }
						                                                                                                          
						}
						break;
					case 2:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(161);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(162);
						((ExpresionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==TK_suma || _la==TK_menos) ) {
							((ExpresionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(163);
						((ExpresionContext)_localctx).e2 = expresion(17);

						                                                                                                                  if (((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null) == "+" {
						                                                                                                                          _localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.SUMA)
						                                                                                                                  }else{
						                                                                                                                          _localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.RESTA)
						                                                                                                                  }
						                                                                                                          
						}
						break;
					case 3:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(166);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(167);
						match(TK_menor);
						setState(168);
						((ExpresionContext)_localctx).e2 = expresion(14);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.MENOR_QUE)
						}
						break;
					case 4:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(171);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(172);
						match(TK_mayor);
						setState(173);
						((ExpresionContext)_localctx).e2 = expresion(13);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.MAYOR_QUE)
						}
						break;
					case 5:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(176);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(177);
						match(TK_mayor_igual);
						setState(178);
						((ExpresionContext)_localctx).e2 = expresion(12);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.MAYOR_IGUAL)
						}
						break;
					case 6:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(181);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(182);
						match(TK_menor_igual);
						setState(183);
						((ExpresionContext)_localctx).e2 = expresion(11);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.MENOR_IGUAL)
						}
						break;
					case 7:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(186);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(187);
						match(TK_igualacion);
						setState(188);
						((ExpresionContext)_localctx).e2 = expresion(10);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.IGUALDAD)
						}
						break;
					case 8:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(191);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(192);
						match(TK_diferente);
						setState(193);
						((ExpresionContext)_localctx).e2 = expresion(9);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.DESIGUAL)
						}
						break;
					case 9:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(196);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(197);
						match(TK_or);
						setState(198);
						((ExpresionContext)_localctx).e2 = expresion(8);
						_localctx.exp = Expresiones.NuevaLogica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.OR)
						}
						break;
					case 10:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(201);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(202);
						match(TK_and);
						setState(203);
						((ExpresionContext)_localctx).e2 = expresion(7);
						_localctx.exp = Expresiones.NuevaLogica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.AND)
						}
						break;
					}
					} 
				}
				setState(210);
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
		public Interfaces.Instruccion inst;
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
			setState(224);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(211);
				match(TKR_println);
				setState(212);
				match(TK_par_apertura);
				setState(213);
				((ImpresionContext)_localctx).expresion = expresion(0);
				setState(214);
				match(TK_par_cierre);
				setState(215);
				match(TK_pcoma);
				fmt.Println("Impresion")
				                                                                                                        _localctx.inst = Instruccion.NuevoPrint(((ImpresionContext)_localctx).expresion.exp)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(218);
				match(TKR_println);
				setState(219);
				match(TK_par_apertura);
				setState(220);
				expresion(0);
				setState(221);
				impresioncomas(0);
				setState(222);
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
			setState(227);
			match(TK_coma);
			setState(228);
			expresion(0);
			setState(229);
			match(TK_par_cierre);
			setState(230);
			match(TK_pcoma);
			fmt.Println("Impresion especial")
			}
			_ctx.stop = _input.LT(-1);
			setState(242);
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
					setState(233);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(234);
					match(TK_coma);
					setState(235);
					expresion(0);
					setState(236);
					match(TK_par_cierre);
					setState(237);
					match(TK_pcoma);
					fmt.Println("Impresion especial")
					}
					} 
				}
				setState(244);
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
			return precpred(_ctx, 17);
		case 1:
			return precpred(_ctx, 16);
		case 2:
			return precpred(_ctx, 13);
		case 3:
			return precpred(_ctx, 12);
		case 4:
			return precpred(_ctx, 11);
		case 5:
			return precpred(_ctx, 10);
		case 6:
			return precpred(_ctx, 9);
		case 7:
			return precpred(_ctx, 8);
		case 8:
			return precpred(_ctx, 7);
		case 9:
			return precpred(_ctx, 6);
		}
		return true;
	}
	private boolean impresioncomas_sempred(ImpresioncomasContext _localctx, int predIndex) {
		switch (predIndex) {
		case 10:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3H\u00f8\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\3\2\3\2"+
		"\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\3\7\3\37\n\3\f\3\16\3\"\13\3\3\3\3\3\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4/\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5"+
		"\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\5\5S\n\5\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6a\n\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7"+
		"\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7s\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\5\b\u009d"+
		"\n\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\7\b\u00d1"+
		"\n\b\f\b\16\b\u00d4\13\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\5\t\u00e3\n\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\7\n\u00f3\n\n\f\n\16\n\u00f6\13\n\3\n\2\4\16\22\13\2\4\6\b\n\f"+
		"\16\20\22\2\4\3\2\31\33\3\2\27\30\2\u0113\2\24\3\2\2\2\4 \3\2\2\2\6.\3"+
		"\2\2\2\bR\3\2\2\2\n`\3\2\2\2\fr\3\2\2\2\16\u009c\3\2\2\2\20\u00e2\3\2"+
		"\2\2\22\u00e4\3\2\2\2\24\25\7\64\2\2\25\26\7B\2\2\26\27\7\24\2\2\27\30"+
		"\7\25\2\2\30\31\7\n\2\2\31\32\5\4\3\2\32\33\7\13\2\2\33\34\b\2\1\2\34"+
		"\3\3\2\2\2\35\37\5\6\4\2\36\35\3\2\2\2\37\"\3\2\2\2 \36\3\2\2\2 !\3\2"+
		"\2\2!#\3\2\2\2\" \3\2\2\2#$\b\3\1\2$\5\3\2\2\2%&\5\16\b\2&\'\b\4\1\2\'"+
		"/\3\2\2\2()\5\20\t\2)*\b\4\1\2*/\3\2\2\2+,\5\b\5\2,-\b\4\1\2-/\3\2\2\2"+
		".%\3\2\2\2.(\3\2\2\2.+\3\2\2\2/\7\3\2\2\2\60\61\7-\2\2\61\62\7.\2\2\62"+
		"\63\7E\2\2\63\64\7\16\2\2\64\65\5\n\6\2\65\66\7\26\2\2\66\67\5\16\b\2"+
		"\678\7\20\2\289\b\5\1\29S\3\2\2\2:;\7-\2\2;<\7.\2\2<=\7E\2\2=>\7\26\2"+
		"\2>?\5\16\b\2?@\7\20\2\2@A\b\5\1\2AS\3\2\2\2BC\7-\2\2CD\7E\2\2DE\7\16"+
		"\2\2EF\5\n\6\2FG\7\26\2\2GH\5\16\b\2HI\7\20\2\2IJ\b\5\1\2JS\3\2\2\2KL"+
		"\7-\2\2LM\7E\2\2MN\7\26\2\2NO\5\16\b\2OP\7\20\2\2PQ\b\5\1\2QS\3\2\2\2"+
		"R\60\3\2\2\2R:\3\2\2\2RB\3\2\2\2RK\3\2\2\2S\t\3\2\2\2TU\7#\2\2Ua\b\6\1"+
		"\2VW\7$\2\2Wa\b\6\1\2XY\7+\2\2Ya\b\6\1\2Z[\7(\2\2[a\b\6\1\2\\]\7)\2\2"+
		"]a\b\6\1\2^_\7,\2\2_a\b\6\1\2`T\3\2\2\2`V\3\2\2\2`X\3\2\2\2`Z\3\2\2\2"+
		"`\\\3\2\2\2`^\3\2\2\2a\13\3\2\2\2bc\7C\2\2cs\b\7\1\2de\7D\2\2es\b\7\1"+
		"\2fg\7G\2\2gs\b\7\1\2hi\7H\2\2is\b\7\1\2jk\7*\2\2ks\b\7\1\2lm\7\62\2\2"+
		"ms\b\7\1\2no\7\63\2\2os\b\7\1\2pq\7E\2\2qs\b\7\1\2rb\3\2\2\2rd\3\2\2\2"+
		"rf\3\2\2\2rh\3\2\2\2rj\3\2\2\2rl\3\2\2\2rn\3\2\2\2rp\3\2\2\2s\r\3\2\2"+
		"\2tu\b\b\1\2uv\7\30\2\2vw\5\16\b\24wx\b\b\1\2x\u009d\3\2\2\2yz\7%\2\2"+
		"z{\7\24\2\2{|\5\16\b\2|}\7\17\2\2}~\5\16\b\2~\177\7\25\2\2\177\u0080\b"+
		"\b\1\2\u0080\u009d\3\2\2\2\u0081\u0082\7\'\2\2\u0082\u0083\7\24\2\2\u0083"+
		"\u0084\5\16\b\2\u0084\u0085\7\17\2\2\u0085\u0086\5\16\b\2\u0086\u0087"+
		"\7\25\2\2\u0087\u0088\b\b\1\2\u0088\u009d\3\2\2\2\u0089\u008a\7\36\2\2"+
		"\u008a\u008b\5\16\b\7\u008b\u008c\b\b\1\2\u008c\u009d\3\2\2\2\u008d\u008e"+
		"\7\24\2\2\u008e\u008f\5\16\b\2\u008f\u0090\7\25\2\2\u0090\u009d\3\2\2"+
		"\2\u0091\u0092\5\f\7\2\u0092\u0093\7\60\2\2\u0093\u0094\7#\2\2\u0094\u009d"+
		"\3\2\2\2\u0095\u0096\5\f\7\2\u0096\u0097\7\60\2\2\u0097\u0098\7$\2\2\u0098"+
		"\u009d\3\2\2\2\u0099\u009a\5\f\7\2\u009a\u009b\b\b\1\2\u009b\u009d\3\2"+
		"\2\2\u009ct\3\2\2\2\u009cy\3\2\2\2\u009c\u0081\3\2\2\2\u009c\u0089\3\2"+
		"\2\2\u009c\u008d\3\2\2\2\u009c\u0091\3\2\2\2\u009c\u0095\3\2\2\2\u009c"+
		"\u0099\3\2\2\2\u009d\u00d2\3\2\2\2\u009e\u009f\f\23\2\2\u009f\u00a0\t"+
		"\2\2\2\u00a0\u00a1\5\16\b\24\u00a1\u00a2\b\b\1\2\u00a2\u00d1\3\2\2\2\u00a3"+
		"\u00a4\f\22\2\2\u00a4\u00a5\t\3\2\2\u00a5\u00a6\5\16\b\23\u00a6\u00a7"+
		"\b\b\1\2\u00a7\u00d1\3\2\2\2\u00a8\u00a9\f\17\2\2\u00a9\u00aa\7\21\2\2"+
		"\u00aa\u00ab\5\16\b\20\u00ab\u00ac\b\b\1\2\u00ac\u00d1\3\2\2\2\u00ad\u00ae"+
		"\f\16\2\2\u00ae\u00af\7\22\2\2\u00af\u00b0\5\16\b\17\u00b0\u00b1\b\b\1"+
		"\2\u00b1\u00d1\3\2\2\2\u00b2\u00b3\f\r\2\2\u00b3\u00b4\7\b\2\2\u00b4\u00b5"+
		"\5\16\b\16\u00b5\u00b6\b\b\1\2\u00b6\u00d1\3\2\2\2\u00b7\u00b8\f\f\2\2"+
		"\u00b8\u00b9\7\t\2\2\u00b9\u00ba\5\16\b\r\u00ba\u00bb\b\b\1\2\u00bb\u00d1"+
		"\3\2\2\2\u00bc\u00bd\f\13\2\2\u00bd\u00be\7\6\2\2\u00be\u00bf\5\16\b\f"+
		"\u00bf\u00c0\b\b\1\2\u00c0\u00d1\3\2\2\2\u00c1\u00c2\f\n\2\2\u00c2\u00c3"+
		"\7\7\2\2\u00c3\u00c4\5\16\b\13\u00c4\u00c5\b\b\1\2\u00c5\u00d1\3\2\2\2"+
		"\u00c6\u00c7\f\t\2\2\u00c7\u00c8\7\4\2\2\u00c8\u00c9\5\16\b\n\u00c9\u00ca"+
		"\b\b\1\2\u00ca\u00d1\3\2\2\2\u00cb\u00cc\f\b\2\2\u00cc\u00cd\7\5\2\2\u00cd"+
		"\u00ce\5\16\b\t\u00ce\u00cf\b\b\1\2\u00cf\u00d1\3\2\2\2\u00d0\u009e\3"+
		"\2\2\2\u00d0\u00a3\3\2\2\2\u00d0\u00a8\3\2\2\2\u00d0\u00ad\3\2\2\2\u00d0"+
		"\u00b2\3\2\2\2\u00d0\u00b7\3\2\2\2\u00d0\u00bc\3\2\2\2\u00d0\u00c1\3\2"+
		"\2\2\u00d0\u00c6\3\2\2\2\u00d0\u00cb\3\2\2\2\u00d1\u00d4\3\2\2\2\u00d2"+
		"\u00d0\3\2\2\2\u00d2\u00d3\3\2\2\2\u00d3\17\3\2\2\2\u00d4\u00d2\3\2\2"+
		"\2\u00d5\u00d6\7\61\2\2\u00d6\u00d7\7\24\2\2\u00d7\u00d8\5\16\b\2\u00d8"+
		"\u00d9\7\25\2\2\u00d9\u00da\7\20\2\2\u00da\u00db\b\t\1\2\u00db\u00e3\3"+
		"\2\2\2\u00dc\u00dd\7\61\2\2\u00dd\u00de\7\24\2\2\u00de\u00df\5\16\b\2"+
		"\u00df\u00e0\5\22\n\2\u00e0\u00e1\7\20\2\2\u00e1\u00e3\3\2\2\2\u00e2\u00d5"+
		"\3\2\2\2\u00e2\u00dc\3\2\2\2\u00e3\21\3\2\2\2\u00e4\u00e5\b\n\1\2\u00e5"+
		"\u00e6\7\17\2\2\u00e6\u00e7\5\16\b\2\u00e7\u00e8\7\25\2\2\u00e8\u00e9"+
		"\7\20\2\2\u00e9\u00ea\b\n\1\2\u00ea\u00f4\3\2\2\2\u00eb\u00ec\f\4\2\2"+
		"\u00ec\u00ed\7\17\2\2\u00ed\u00ee\5\16\b\2\u00ee\u00ef\7\25\2\2\u00ef"+
		"\u00f0\7\20\2\2\u00f0\u00f1\b\n\1\2\u00f1\u00f3\3\2\2\2\u00f2\u00eb\3"+
		"\2\2\2\u00f3\u00f6\3\2\2\2\u00f4\u00f2\3\2\2\2\u00f4\u00f5\3\2\2\2\u00f5"+
		"\23\3\2\2\2\u00f6\u00f4\3\2\2\2\f .R`r\u009c\u00d0\u00d2\u00e2\u00f4";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}