// Generated from e:\WillOP\u005Cu\Compi2\Laboratorio\Proyecto1_201408419\gramatica.g4 by ANTLR 4.8

    import "proyecto1/Interfaces"
    import "proyecto1/Expresion"
    import "proyecto1/Instruccion"
    import "proyecto1/Expresiones"
    import arrayList "github.com/colegno/arraylist"
    //import "proyecto1/Operaciones"
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
		TKR_capacity=62, TKR_with_capacity=63, TKR_main=64, TKR_if=65, TKR_elseif=66, 
		TKR_else=67, TKR_while=68, TKR_break=69, TKR_loop=70, TK_entero=71, TK_decimal=72, 
		TK_id=73, Letra=74, TK_cadena=75, TK_caracter=76;
	public static final int
		RULE_start = 0, RULE_instrucciones = 1, RULE_instruccion = 2, RULE_control = 3, 
		RULE_declaracion = 4, RULE_tipovariable = 5, RULE_identificadores = 6, 
		RULE_valores = 7, RULE_expresion = 8, RULE_impresion = 9, RULE_impresionexpresion = 10, 
		RULE_expcoma = 11, RULE_condicionales = 12, RULE_funcionif = 13, RULE_funcionelseif = 14, 
		RULE_listaelseif = 15, RULE_bloque = 16, RULE_bucles = 17, RULE_fwhile = 18, 
		RULE_floop = 19;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "instrucciones", "instruccion", "control", "declaracion", "tipovariable", 
			"identificadores", "valores", "expresion", "impresion", "impresionexpresion", 
			"expcoma", "condicionales", "funcionif", "funcionelseif", "listaelseif", 
			"bloque", "bucles", "fwhile", "floop"
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
			"'main'", "'if'", "'else if'", "'else'", "'while'", "'break'", "'loop'"
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
			"TKR_with_capacity", "TKR_main", "TKR_if", "TKR_elseif", "TKR_else", 
			"TKR_while", "TKR_break", "TKR_loop", "TK_entero", "TK_decimal", "TK_id", 
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
			setState(40);
			match(TKR_fn);
			setState(41);
			match(TKR_main);
			setState(42);
			match(TK_par_apertura);
			setState(43);
			match(TK_par_cierre);
			setState(44);
			match(TK_corchete_apertura);
			setState(45);
			((StartContext)_localctx).instrucciones = instrucciones();
			setState(46);
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
			setState(52);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 18)) & ~0x3f) == 0 && ((1L << (_la - 18)) & ((1L << (TK_par_apertura - 18)) | (1L << (TK_menos - 18)) | (1L << (TK_sig_admiracion - 18)) | (1L << (TKR_pow - 18)) | (1L << (TKR_powf - 18)) | (1L << (TKR_Str - 18)) | (1L << (TKR_let - 18)) | (1L << (TKR_println - 18)) | (1L << (TKR_true - 18)) | (1L << (TKR_false - 18)) | (1L << (TKR_if - 18)) | (1L << (TKR_while - 18)) | (1L << (TKR_break - 18)) | (1L << (TKR_loop - 18)) | (1L << (TK_entero - 18)) | (1L << (TK_decimal - 18)) | (1L << (TK_id - 18)) | (1L << (TK_cadena - 18)) | (1L << (TK_caracter - 18)))) != 0)) {
				{
				{
				setState(49);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).e.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(54);
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
		public IdentificadoresContext identificadores;
		public CondicionalesContext condicionales;
		public BuclesContext bucles;
		public ControlContext control;
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
		public CondicionalesContext condicionales() {
			return getRuleContext(CondicionalesContext.class,0);
		}
		public BuclesContext bucles() {
			return getRuleContext(BuclesContext.class,0);
		}
		public ControlContext control() {
			return getRuleContext(ControlContext.class,0);
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
			setState(78);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(57);
				((InstruccionContext)_localctx).expresion = expresion(0);
				fmt.Println("mensaje en instrucciones: ",((InstruccionContext)_localctx).expresion.exp)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(60);
				((InstruccionContext)_localctx).impresion = impresion();
				_localctx.inst = ((InstruccionContext)_localctx).impresion.inst
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(63);
				((InstruccionContext)_localctx).declaracion = declaracion();
				_localctx.inst = ((InstruccionContext)_localctx).declaracion.inst
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(66);
				((InstruccionContext)_localctx).identificadores = identificadores();
				_localctx.inst = ((InstruccionContext)_localctx).identificadores.inst
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(69);
				((InstruccionContext)_localctx).condicionales = condicionales();
				_localctx.inst = ((InstruccionContext)_localctx).condicionales.inst
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(72);
				((InstruccionContext)_localctx).bucles = bucles();
				_localctx.inst = ((InstruccionContext)_localctx).bucles.inst
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(75);
				((InstruccionContext)_localctx).control = control();
				_localctx.inst = ((InstruccionContext)_localctx).control.inst
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

	public static class ControlContext extends ParserRuleContext {
		public Interfaces.Instruccion inst;
		public TerminalNode TKR_break() { return getToken(gramaticaParser.TKR_break, 0); }
		public TerminalNode TK_pcoma() { return getToken(gramaticaParser.TK_pcoma, 0); }
		public ControlContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_control; }
	}

	public final ControlContext control() throws RecognitionException {
		ControlContext _localctx = new ControlContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_control);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(80);
			match(TKR_break);
			setState(81);
			match(TK_pcoma);
			_localctx.inst = Instruccion.NewBreak()
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
		enterRule(_localctx, 8, RULE_declaracion);
		try {
			setState(118);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(84);
				match(TKR_let);
				setState(85);
				match(TKR_mut);
				setState(86);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(87);
				match(TK_dosPuntos);
				setState(88);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(89);
				match(TK_igual);
				setState(90);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(91);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar,((DeclaracionContext)_localctx).expresion.exp,false,false,true) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(94);
				match(TKR_let);
				setState(95);
				match(TKR_mut);
				setState(96);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(97);
				match(TK_igual);
				setState(98);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(99);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),Interfaces.SINTIPO,((DeclaracionContext)_localctx).expresion.exp,false,false,true) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(102);
				match(TKR_let);
				setState(103);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(104);
				match(TK_dosPuntos);
				setState(105);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(106);
				match(TK_igual);
				setState(107);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(108);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar,((DeclaracionContext)_localctx).expresion.exp,false,false,false) 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(111);
				match(TKR_let);
				setState(112);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(113);
				match(TK_igual);
				setState(114);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(115);
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
		enterRule(_localctx, 10, RULE_tipovariable);
		try {
			setState(132);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TKR_numericos_enteros:
				enterOuterAlt(_localctx, 1);
				{
				setState(120);
				match(TKR_numericos_enteros);
				_localctx.tipovar = Interfaces.INTEGER
				}
				break;
			case TKR_numericos_flotantes:
				enterOuterAlt(_localctx, 2);
				{
				setState(122);
				match(TKR_numericos_flotantes);
				_localctx.tipovar = Interfaces.FLOAT
				}
				break;
			case TKR_String:
				enterOuterAlt(_localctx, 3);
				{
				setState(124);
				match(TKR_String);
				_localctx.tipovar = Interfaces.STRING
				}
				break;
			case TKR_bool:
				enterOuterAlt(_localctx, 4);
				{
				setState(126);
				match(TKR_bool);
				_localctx.tipovar = Interfaces.BOOLEAN
				}
				break;
			case TKR_char:
				enterOuterAlt(_localctx, 5);
				{
				setState(128);
				match(TKR_char);
				_localctx.tipovar = Interfaces.CHAR
				}
				break;
			case TKR_usize:
				enterOuterAlt(_localctx, 6);
				{
				setState(130);
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
		public Interfaces.Instruccion inst;
		public Token e1;
		public ExpresionContext e2;
		public TerminalNode TK_igual() { return getToken(gramaticaParser.TK_igual, 0); }
		public TerminalNode TK_pcoma() { return getToken(gramaticaParser.TK_pcoma, 0); }
		public TerminalNode TK_id() { return getToken(gramaticaParser.TK_id, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public IdentificadoresContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identificadores; }
	}

	public final IdentificadoresContext identificadores() throws RecognitionException {
		IdentificadoresContext _localctx = new IdentificadoresContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_identificadores);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(134);
			((IdentificadoresContext)_localctx).e1 = match(TK_id);
			setState(135);
			match(TK_igual);
			setState(136);
			((IdentificadoresContext)_localctx).e2 = expresion(0);
			setState(137);
			match(TK_pcoma);
			_localctx.inst = Instruccion.NewAsignacion((((IdentificadoresContext)_localctx).e1!=null?((IdentificadoresContext)_localctx).e1.getText():null),((IdentificadoresContext)_localctx).e2.exp)
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
		enterRule(_localctx, 14, RULE_valores);
		try {
			setState(156);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_entero:
				enterOuterAlt(_localctx, 1);
				{
				setState(140);
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
				setState(142);
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
				setState(144);
				((ValoresContext)_localctx).TK_cadena = match(TK_cadena);

				                str:= (((ValoresContext)_localctx).TK_cadena!=null?((ValoresContext)_localctx).TK_cadena.getText():null)[1:len((((ValoresContext)_localctx).TK_cadena!=null?((ValoresContext)_localctx).TK_cadena.getText():null))-1]
				                _localctx.exp = Expresion.NuevoPrimitivo(str,Interfaces.STRING)
				}
				break;
			case TK_caracter:
				enterOuterAlt(_localctx, 4);
				{
				setState(146);
				((ValoresContext)_localctx).TK_caracter = match(TK_caracter);
				str:= (((ValoresContext)_localctx).TK_caracter!=null?((ValoresContext)_localctx).TK_caracter.getText():null)[1:len((((ValoresContext)_localctx).TK_caracter!=null?((ValoresContext)_localctx).TK_caracter.getText():null))-1]
				                _localctx.exp = Expresion.NuevoPrimitivo(str,Interfaces.CHAR)
				}
				break;
			case TKR_Str:
				enterOuterAlt(_localctx, 5);
				{
				setState(148);
				((ValoresContext)_localctx).TKR_Str = match(TKR_Str);
				str:= (((ValoresContext)_localctx).TKR_Str!=null?((ValoresContext)_localctx).TKR_Str.getText():null)[1:len((((ValoresContext)_localctx).TKR_Str!=null?((ValoresContext)_localctx).TKR_Str.getText():null))-1]
				                _localctx.exp = Expresion.NuevoPrimitivo(str,Interfaces.STR)
				}
				break;
			case TKR_true:
				enterOuterAlt(_localctx, 6);
				{
				setState(150);
				match(TKR_true);
				_localctx.exp = Expresion.NuevoPrimitivo(true, Interfaces.BOOLEAN)
				}
				break;
			case TKR_false:
				enterOuterAlt(_localctx, 7);
				{
				setState(152);
				match(TKR_false);
				_localctx.exp = Expresion.NuevoPrimitivo(false,Interfaces.BOOLEAN)
				}
				break;
			case TK_id:
				enterOuterAlt(_localctx, 8);
				{
				setState(154);
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
		public ExpresionContext vll;
		public ExpresionContext e2;
		public ExpresionContext va;
		public ValoresContext val;
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
		public TerminalNode TKR_as() { return getToken(gramaticaParser.TKR_as, 0); }
		public TerminalNode TKR_numericos_enteros() { return getToken(gramaticaParser.TKR_numericos_enteros, 0); }
		public ValoresContext valores() {
			return getRuleContext(ValoresContext.class,0);
		}
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
		public TerminalNode TK_punto() { return getToken(gramaticaParser.TK_punto, 0); }
		public TerminalNode TKR_abs() { return getToken(gramaticaParser.TKR_abs, 0); }
		public TerminalNode TKR_sqrt() { return getToken(gramaticaParser.TKR_sqrt, 0); }
		public TerminalNode TKR_to_string() { return getToken(gramaticaParser.TKR_to_string, 0); }
		public TerminalNode TKR_clone() { return getToken(gramaticaParser.TKR_clone, 0); }
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
		int _startState = 16;
		enterRecursionRule(_localctx, 16, RULE_expresion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(201);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				{
				setState(159);
				match(TK_menos);
				setState(160);
				((ExpresionContext)_localctx).e1 = expresion(22);
				numero := -1
				                                                                                                        numm := Expresion.NuevoPrimitivo(numero, Interfaces.INTEGER)
				                                                                                                                _localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp, numm ,Interfaces.MULTIPLICACION)
				                                                                                                
				}
				break;
			case 2:
				{
				setState(163);
				match(TKR_pow);
				setState(164);
				match(TK_par_apertura);
				setState(165);
				((ExpresionContext)_localctx).e1 = expresion(0);
				setState(166);
				match(TK_coma);
				setState(167);
				((ExpresionContext)_localctx).e2 = expresion(0);
				setState(168);
				match(TK_par_cierre);
				_localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.POW)
				}
				break;
			case 3:
				{
				setState(171);
				match(TKR_powf);
				setState(172);
				match(TK_par_apertura);
				setState(173);
				((ExpresionContext)_localctx).e1 = expresion(0);
				setState(174);
				match(TK_coma);
				setState(175);
				((ExpresionContext)_localctx).e2 = expresion(0);
				setState(176);
				match(TK_par_cierre);
				_localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.POWF)
				}
				break;
			case 4:
				{
				setState(179);
				match(TK_sig_admiracion);
				setState(180);
				((ExpresionContext)_localctx).e1 = expresion(9);
				_localctx.exp = Expresiones.NuevaLogica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e1.exp,Interfaces.NOT)
				}
				break;
			case 5:
				{
				setState(183);
				match(TK_par_apertura);
				setState(184);
				((ExpresionContext)_localctx).va = expresion(0);
				setState(185);
				match(TK_par_cierre);
				_localctx.exp = ((ExpresionContext)_localctx).va.exp
				}
				break;
			case 6:
				{
				setState(188);
				((ExpresionContext)_localctx).val = valores();
				setState(189);
				match(TKR_as);
				setState(190);
				match(TKR_numericos_enteros);
				_localctx.exp = Expresiones.NewAsi64(((ExpresionContext)_localctx).val.exp)
				}
				break;
			case 7:
				{
				setState(193);
				((ExpresionContext)_localctx).val = valores();
				setState(194);
				match(TKR_as);
				setState(195);
				match(TKR_numericos_flotantes);
				_localctx.exp = Expresiones.NewAsf64(((ExpresionContext)_localctx).val.exp)
				}
				break;
			case 8:
				{
				setState(198);
				((ExpresionContext)_localctx).vall = valores();
				_localctx.exp = ((ExpresionContext)_localctx).vall.exp
				                                                                                        fmt.Println(_localctx.exp)
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(270);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(268);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
					case 1:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(203);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(204);
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
						setState(205);
						((ExpresionContext)_localctx).e2 = expresion(22);

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
						setState(208);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(209);
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
						setState(210);
						((ExpresionContext)_localctx).e2 = expresion(21);

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
						setState(213);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(214);
						match(TK_menor);
						setState(215);
						((ExpresionContext)_localctx).e2 = expresion(18);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.MENOR_QUE)
						}
						break;
					case 4:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(218);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(219);
						match(TK_mayor);
						setState(220);
						((ExpresionContext)_localctx).e2 = expresion(17);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.MAYOR_QUE)
						}
						break;
					case 5:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(223);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(224);
						match(TK_mayor_igual);
						setState(225);
						((ExpresionContext)_localctx).e2 = expresion(16);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.MAYOR_IGUAL)
						}
						break;
					case 6:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(228);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(229);
						match(TK_menor_igual);
						setState(230);
						((ExpresionContext)_localctx).e2 = expresion(15);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.MENOR_IGUAL)
						}
						break;
					case 7:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(233);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(234);
						match(TK_igualacion);
						setState(235);
						((ExpresionContext)_localctx).e2 = expresion(14);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.IGUALDAD)
						}
						break;
					case 8:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(238);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(239);
						match(TK_diferente);
						setState(240);
						((ExpresionContext)_localctx).e2 = expresion(13);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.DESIGUAL)
						}
						break;
					case 9:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(243);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(244);
						match(TK_or);
						setState(245);
						((ExpresionContext)_localctx).e2 = expresion(12);
						_localctx.exp = Expresiones.NuevaLogica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.OR)
						}
						break;
					case 10:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(248);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(249);
						match(TK_and);
						setState(250);
						((ExpresionContext)_localctx).e2 = expresion(11);
						_localctx.exp = Expresiones.NuevaLogica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.AND)
						}
						break;
					case 11:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.vll = _prevctx;
						_localctx.vll = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(253);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(254);
						match(TK_punto);
						setState(255);
						match(TKR_abs);
						setState(256);
						match(TK_par_apertura);
						setState(257);
						match(TK_par_cierre);
						_localctx.exp = Expresiones.Newabs(((ExpresionContext)_localctx).vll.exp)
						}
						break;
					case 12:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.vll = _prevctx;
						_localctx.vll = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(259);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(260);
						match(TK_punto);
						setState(261);
						match(TKR_sqrt);
						}
						break;
					case 13:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.vll = _prevctx;
						_localctx.vll = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(262);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(263);
						match(TK_punto);
						setState(264);
						match(TKR_to_string);
						}
						break;
					case 14:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.vll = _prevctx;
						_localctx.vll = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(265);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(266);
						match(TK_punto);
						setState(267);
						match(TKR_clone);
						}
						break;
					}
					} 
				}
				setState(272);
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
		public ExpresionContext e1;
		public ExpresionContext e2;
		public ImpresionexpresionContext li;
		public TerminalNode TKR_println() { return getToken(gramaticaParser.TKR_println, 0); }
		public TerminalNode TK_par_apertura() { return getToken(gramaticaParser.TK_par_apertura, 0); }
		public TerminalNode TK_par_cierre() { return getToken(gramaticaParser.TK_par_cierre, 0); }
		public TerminalNode TK_pcoma() { return getToken(gramaticaParser.TK_pcoma, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public ImpresionexpresionContext impresionexpresion() {
			return getRuleContext(ImpresionexpresionContext.class,0);
		}
		public ImpresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_impresion; }
	}

	public final ImpresionContext impresion() throws RecognitionException {
		ImpresionContext _localctx = new ImpresionContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_impresion);
		try {
			setState(287);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(273);
				match(TKR_println);
				setState(274);
				match(TK_par_apertura);
				setState(275);
				((ImpresionContext)_localctx).e1 = expresion(0);
				setState(276);
				match(TK_par_cierre);
				setState(277);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevoPrint(((ImpresionContext)_localctx).e1.exp,nil)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(280);
				match(TKR_println);
				setState(281);
				match(TK_par_apertura);
				setState(282);
				((ImpresionContext)_localctx).e2 = expresion(0);
				setState(283);
				((ImpresionContext)_localctx).li = impresionexpresion();
				setState(284);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevoPrint(((ImpresionContext)_localctx).e2.exp,((ImpresionContext)_localctx).li.lista)
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

	public static class ImpresionexpresionContext extends ParserRuleContext {
		public *arrayList.List lista;
		public ExpcomaContext expcoma;
		public List<ExpcomaContext> list = new ArrayList<ExpcomaContext>();
		public List<ExpcomaContext> expcoma() {
			return getRuleContexts(ExpcomaContext.class);
		}
		public ExpcomaContext expcoma(int i) {
			return getRuleContext(ExpcomaContext.class,i);
		}
		public ImpresionexpresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_impresionexpresion; }
	}

	public final ImpresionexpresionContext impresionexpresion() throws RecognitionException {
		ImpresionexpresionContext _localctx = new ImpresionexpresionContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_impresionexpresion);
		_localctx.lista = arrayList.New()
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(292);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==TK_coma) {
				{
				{
				setState(289);
				((ImpresionexpresionContext)_localctx).expcoma = expcoma();
				((ImpresionexpresionContext)_localctx).list.add(((ImpresionexpresionContext)_localctx).expcoma);
				}
				}
				setState(294);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}

			                                                                                        listInt:= localctx.(*ImpresionexpresionContext).GetList()
			                                                                                        for _,e := range listInt{
			                                                                                                _localctx.lista.Add(e.GetExp())
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

	public static class ExpcomaContext extends ParserRuleContext {
		public Interfaces.Expresion exp;
		public ExpresionContext e;
		public TerminalNode TK_coma() { return getToken(gramaticaParser.TK_coma, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public ExpcomaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expcoma; }
	}

	public final ExpcomaContext expcoma() throws RecognitionException {
		ExpcomaContext _localctx = new ExpcomaContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_expcoma);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(297);
			match(TK_coma);
			setState(298);
			((ExpcomaContext)_localctx).e = expresion(0);
			_localctx.exp = ((ExpcomaContext)_localctx).e.exp
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

	public static class CondicionalesContext extends ParserRuleContext {
		public Interfaces.Instruccion inst;
		public FuncionifContext funcionif;
		public FuncionifContext funcionif() {
			return getRuleContext(FuncionifContext.class,0);
		}
		public CondicionalesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_condicionales; }
	}

	public final CondicionalesContext condicionales() throws RecognitionException {
		CondicionalesContext _localctx = new CondicionalesContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_condicionales);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(301);
			((CondicionalesContext)_localctx).funcionif = funcionif();
			_localctx.inst = ((CondicionalesContext)_localctx).funcionif.inst
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

	public static class FuncionifContext extends ParserRuleContext {
		public Interfaces.Instruccion inst;
		public ExpresionContext e1;
		public BloqueContext ee;
		public BloqueContext e5;
		public BloqueContext b2;
		public BloqueContext b1;
		public ListaelseifContext listaelseif;
		public TerminalNode TKR_if() { return getToken(gramaticaParser.TKR_if, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public List<BloqueContext> bloque() {
			return getRuleContexts(BloqueContext.class);
		}
		public BloqueContext bloque(int i) {
			return getRuleContext(BloqueContext.class,i);
		}
		public TerminalNode TKR_else() { return getToken(gramaticaParser.TKR_else, 0); }
		public ListaelseifContext listaelseif() {
			return getRuleContext(ListaelseifContext.class,0);
		}
		public FuncionifContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcionif; }
	}

	public final FuncionifContext funcionif() throws RecognitionException {
		FuncionifContext _localctx = new FuncionifContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_funcionif);
		try {
			setState(330);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(304);
				match(TKR_if);
				setState(305);
				((FuncionifContext)_localctx).e1 = expresion(0);
				setState(306);
				((FuncionifContext)_localctx).ee = bloque();
				_localctx.inst = Instruccion.NewIf(((FuncionifContext)_localctx).e1.exp,((FuncionifContext)_localctx).ee.lista,nil,nil )
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(309);
				match(TKR_if);
				setState(310);
				((FuncionifContext)_localctx).e1 = expresion(0);
				setState(311);
				((FuncionifContext)_localctx).e5 = bloque();
				setState(312);
				match(TKR_else);
				setState(313);
				((FuncionifContext)_localctx).b2 = bloque();
				_localctx.inst = Instruccion.NewIf(((FuncionifContext)_localctx).e1.exp,((FuncionifContext)_localctx).e5.lista,nil,((FuncionifContext)_localctx).b2.lista)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(316);
				match(TKR_if);
				setState(317);
				((FuncionifContext)_localctx).e1 = expresion(0);
				setState(318);
				((FuncionifContext)_localctx).b1 = bloque();
				setState(319);
				((FuncionifContext)_localctx).listaelseif = listaelseif();
				setState(320);
				match(TKR_else);
				setState(321);
				((FuncionifContext)_localctx).b2 = bloque();
				_localctx.inst = Instruccion.NewIf(((FuncionifContext)_localctx).e1.exp,((FuncionifContext)_localctx).b1.lista,((FuncionifContext)_localctx).listaelseif.lista,((FuncionifContext)_localctx).b2.lista)
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(324);
				match(TKR_if);
				setState(325);
				((FuncionifContext)_localctx).e1 = expresion(0);
				setState(326);
				((FuncionifContext)_localctx).b1 = bloque();
				setState(327);
				((FuncionifContext)_localctx).listaelseif = listaelseif();
				_localctx.inst = Instruccion.NewIf(((FuncionifContext)_localctx).e1.exp,((FuncionifContext)_localctx).b1.lista,((FuncionifContext)_localctx).listaelseif.lista,nil)
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

	public static class FuncionelseifContext extends ParserRuleContext {
		public Interfaces.Instruccion inst;
		public ExpresionContext e1;
		public BloqueContext bloque;
		public TerminalNode TKR_elseif() { return getToken(gramaticaParser.TKR_elseif, 0); }
		public BloqueContext bloque() {
			return getRuleContext(BloqueContext.class,0);
		}
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public FuncionelseifContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcionelseif; }
	}

	public final FuncionelseifContext funcionelseif() throws RecognitionException {
		FuncionelseifContext _localctx = new FuncionelseifContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_funcionelseif);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(332);
			match(TKR_elseif);
			setState(333);
			((FuncionelseifContext)_localctx).e1 = expresion(0);
			setState(334);
			((FuncionelseifContext)_localctx).bloque = bloque();
			_localctx.inst = Instruccion.NewIf(((FuncionelseifContext)_localctx).e1.exp,((FuncionelseifContext)_localctx).bloque.lista,nil,nil)
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

	public static class ListaelseifContext extends ParserRuleContext {
		public *arrayList.List lista;
		public FuncionelseifContext funcionelseif;
		public List<FuncionelseifContext> list = new ArrayList<FuncionelseifContext>();
		public List<FuncionelseifContext> funcionelseif() {
			return getRuleContexts(FuncionelseifContext.class);
		}
		public FuncionelseifContext funcionelseif(int i) {
			return getRuleContext(FuncionelseifContext.class,i);
		}
		public ListaelseifContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listaelseif; }
	}

	public final ListaelseifContext listaelseif() throws RecognitionException {
		ListaelseifContext _localctx = new ListaelseifContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_listaelseif);
		 _localctx.lista = arrayList.New()
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(340);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==TKR_elseif) {
				{
				{
				setState(337);
				((ListaelseifContext)_localctx).funcionelseif = funcionelseif();
				((ListaelseifContext)_localctx).list.add(((ListaelseifContext)_localctx).funcionelseif);
				}
				}
				setState(342);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}

			                                                        listInt := localctx.(*ListaelseifContext).GetList()
			                                                        for _,e := range listInt{
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

	public static class BloqueContext extends ParserRuleContext {
		public *arrayList.List lista;
		public InstruccionesContext instrucciones;
		public TerminalNode TK_corchete_apertura() { return getToken(gramaticaParser.TK_corchete_apertura, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode TK_corchete_cierre() { return getToken(gramaticaParser.TK_corchete_cierre, 0); }
		public BloqueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bloque; }
	}

	public final BloqueContext bloque() throws RecognitionException {
		BloqueContext _localctx = new BloqueContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_bloque);
		try {
			setState(353);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(345);
				match(TK_corchete_apertura);
				setState(346);
				((BloqueContext)_localctx).instrucciones = instrucciones();
				setState(347);
				match(TK_corchete_cierre);
				_localctx.lista = ((BloqueContext)_localctx).instrucciones.lista
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(350);
				match(TK_corchete_apertura);
				setState(351);
				match(TK_corchete_cierre);
				_localctx.lista = arrayList.New()
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

	public static class BuclesContext extends ParserRuleContext {
		public Interfaces.Instruccion inst;
		public FwhileContext fwhile;
		public FloopContext floop;
		public FwhileContext fwhile() {
			return getRuleContext(FwhileContext.class,0);
		}
		public FloopContext floop() {
			return getRuleContext(FloopContext.class,0);
		}
		public BuclesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bucles; }
	}

	public final BuclesContext bucles() throws RecognitionException {
		BuclesContext _localctx = new BuclesContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_bucles);
		try {
			setState(361);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TKR_while:
				enterOuterAlt(_localctx, 1);
				{
				setState(355);
				((BuclesContext)_localctx).fwhile = fwhile();
				_localctx.inst = ((BuclesContext)_localctx).fwhile.inst
				}
				break;
			case TKR_loop:
				enterOuterAlt(_localctx, 2);
				{
				setState(358);
				((BuclesContext)_localctx).floop = floop();
				_localctx.inst = ((BuclesContext)_localctx).floop.inst
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

	public static class FwhileContext extends ParserRuleContext {
		public Interfaces.Instruccion inst;
		public ExpresionContext e1;
		public BloqueContext bl;
		public TerminalNode TKR_while() { return getToken(gramaticaParser.TKR_while, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public BloqueContext bloque() {
			return getRuleContext(BloqueContext.class,0);
		}
		public FwhileContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fwhile; }
	}

	public final FwhileContext fwhile() throws RecognitionException {
		FwhileContext _localctx = new FwhileContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_fwhile);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(363);
			match(TKR_while);
			setState(364);
			((FwhileContext)_localctx).e1 = expresion(0);
			setState(365);
			((FwhileContext)_localctx).bl = bloque();
			_localctx.inst = Instruccion.NewWhile(((FwhileContext)_localctx).e1.exp, ((FwhileContext)_localctx).bl.lista)
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

	public static class FloopContext extends ParserRuleContext {
		public Interfaces.Instruccion inst;
		public BloqueContext bb;
		public TerminalNode TKR_loop() { return getToken(gramaticaParser.TKR_loop, 0); }
		public BloqueContext bloque() {
			return getRuleContext(BloqueContext.class,0);
		}
		public FloopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_floop; }
	}

	public final FloopContext floop() throws RecognitionException {
		FloopContext _localctx = new FloopContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_floop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(368);
			match(TKR_loop);
			setState(369);
			((FloopContext)_localctx).bb = bloque();
			_localctx.inst = Instruccion.NewLoop(((FloopContext)_localctx).bb.lista)
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
		case 8:
			return expresion_sempred((ExpresionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expresion_sempred(ExpresionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 21);
		case 1:
			return precpred(_ctx, 20);
		case 2:
			return precpred(_ctx, 17);
		case 3:
			return precpred(_ctx, 16);
		case 4:
			return precpred(_ctx, 15);
		case 5:
			return precpred(_ctx, 14);
		case 6:
			return precpred(_ctx, 13);
		case 7:
			return precpred(_ctx, 12);
		case 8:
			return precpred(_ctx, 11);
		case 9:
			return precpred(_ctx, 10);
		case 10:
			return precpred(_ctx, 5);
		case 11:
			return precpred(_ctx, 4);
		case 12:
			return precpred(_ctx, 3);
		case 13:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3N\u0177\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\3"+
		"\7\3\65\n\3\f\3\16\38\13\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4Q\n\4\3\5\3\5\3\5"+
		"\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\5\6y\n\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7\u0087\n"+
		"\7\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\5\t\u009f\n\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n"+
		"\u00cc\n\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\7\n\u010f"+
		"\n\n\f\n\16\n\u0112\13\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\5\13\u0122\n\13\3\f\7\f\u0125\n\f\f\f\16\f\u0128"+
		"\13\f\3\f\3\f\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\5\17\u014d\n\17\3\20\3\20\3\20\3\20"+
		"\3\20\3\21\7\21\u0155\n\21\f\21\16\21\u0158\13\21\3\21\3\21\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u0164\n\22\3\23\3\23\3\23\3\23\3\23"+
		"\3\23\5\23\u016c\n\23\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25"+
		"\2\3\22\26\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(\2\4\3\2\31\33"+
		"\3\2\27\30\2\u0195\2*\3\2\2\2\4\66\3\2\2\2\6P\3\2\2\2\bR\3\2\2\2\nx\3"+
		"\2\2\2\f\u0086\3\2\2\2\16\u0088\3\2\2\2\20\u009e\3\2\2\2\22\u00cb\3\2"+
		"\2\2\24\u0121\3\2\2\2\26\u0126\3\2\2\2\30\u012b\3\2\2\2\32\u012f\3\2\2"+
		"\2\34\u014c\3\2\2\2\36\u014e\3\2\2\2 \u0156\3\2\2\2\"\u0163\3\2\2\2$\u016b"+
		"\3\2\2\2&\u016d\3\2\2\2(\u0172\3\2\2\2*+\7\64\2\2+,\7B\2\2,-\7\24\2\2"+
		"-.\7\25\2\2./\7\n\2\2/\60\5\4\3\2\60\61\7\13\2\2\61\62\b\2\1\2\62\3\3"+
		"\2\2\2\63\65\5\6\4\2\64\63\3\2\2\2\658\3\2\2\2\66\64\3\2\2\2\66\67\3\2"+
		"\2\2\679\3\2\2\28\66\3\2\2\29:\b\3\1\2:\5\3\2\2\2;<\5\22\n\2<=\b\4\1\2"+
		"=Q\3\2\2\2>?\5\24\13\2?@\b\4\1\2@Q\3\2\2\2AB\5\n\6\2BC\b\4\1\2CQ\3\2\2"+
		"\2DE\5\16\b\2EF\b\4\1\2FQ\3\2\2\2GH\5\32\16\2HI\b\4\1\2IQ\3\2\2\2JK\5"+
		"$\23\2KL\b\4\1\2LQ\3\2\2\2MN\5\b\5\2NO\b\4\1\2OQ\3\2\2\2P;\3\2\2\2P>\3"+
		"\2\2\2PA\3\2\2\2PD\3\2\2\2PG\3\2\2\2PJ\3\2\2\2PM\3\2\2\2Q\7\3\2\2\2RS"+
		"\7G\2\2ST\7\20\2\2TU\b\5\1\2U\t\3\2\2\2VW\7-\2\2WX\7.\2\2XY\7K\2\2YZ\7"+
		"\16\2\2Z[\5\f\7\2[\\\7\26\2\2\\]\5\22\n\2]^\7\20\2\2^_\b\6\1\2_y\3\2\2"+
		"\2`a\7-\2\2ab\7.\2\2bc\7K\2\2cd\7\26\2\2de\5\22\n\2ef\7\20\2\2fg\b\6\1"+
		"\2gy\3\2\2\2hi\7-\2\2ij\7K\2\2jk\7\16\2\2kl\5\f\7\2lm\7\26\2\2mn\5\22"+
		"\n\2no\7\20\2\2op\b\6\1\2py\3\2\2\2qr\7-\2\2rs\7K\2\2st\7\26\2\2tu\5\22"+
		"\n\2uv\7\20\2\2vw\b\6\1\2wy\3\2\2\2xV\3\2\2\2x`\3\2\2\2xh\3\2\2\2xq\3"+
		"\2\2\2y\13\3\2\2\2z{\7#\2\2{\u0087\b\7\1\2|}\7$\2\2}\u0087\b\7\1\2~\177"+
		"\7+\2\2\177\u0087\b\7\1\2\u0080\u0081\7(\2\2\u0081\u0087\b\7\1\2\u0082"+
		"\u0083\7)\2\2\u0083\u0087\b\7\1\2\u0084\u0085\7,\2\2\u0085\u0087\b\7\1"+
		"\2\u0086z\3\2\2\2\u0086|\3\2\2\2\u0086~\3\2\2\2\u0086\u0080\3\2\2\2\u0086"+
		"\u0082\3\2\2\2\u0086\u0084\3\2\2\2\u0087\r\3\2\2\2\u0088\u0089\7K\2\2"+
		"\u0089\u008a\7\26\2\2\u008a\u008b\5\22\n\2\u008b\u008c\7\20\2\2\u008c"+
		"\u008d\b\b\1\2\u008d\17\3\2\2\2\u008e\u008f\7I\2\2\u008f\u009f\b\t\1\2"+
		"\u0090\u0091\7J\2\2\u0091\u009f\b\t\1\2\u0092\u0093\7M\2\2\u0093\u009f"+
		"\b\t\1\2\u0094\u0095\7N\2\2\u0095\u009f\b\t\1\2\u0096\u0097\7*\2\2\u0097"+
		"\u009f\b\t\1\2\u0098\u0099\7\62\2\2\u0099\u009f\b\t\1\2\u009a\u009b\7"+
		"\63\2\2\u009b\u009f\b\t\1\2\u009c\u009d\7K\2\2\u009d\u009f\b\t\1\2\u009e"+
		"\u008e\3\2\2\2\u009e\u0090\3\2\2\2\u009e\u0092\3\2\2\2\u009e\u0094\3\2"+
		"\2\2\u009e\u0096\3\2\2\2\u009e\u0098\3\2\2\2\u009e\u009a\3\2\2\2\u009e"+
		"\u009c\3\2\2\2\u009f\21\3\2\2\2\u00a0\u00a1\b\n\1\2\u00a1\u00a2\7\30\2"+
		"\2\u00a2\u00a3\5\22\n\30\u00a3\u00a4\b\n\1\2\u00a4\u00cc\3\2\2\2\u00a5"+
		"\u00a6\7%\2\2\u00a6\u00a7\7\24\2\2\u00a7\u00a8\5\22\n\2\u00a8\u00a9\7"+
		"\17\2\2\u00a9\u00aa\5\22\n\2\u00aa\u00ab\7\25\2\2\u00ab\u00ac\b\n\1\2"+
		"\u00ac\u00cc\3\2\2\2\u00ad\u00ae\7\'\2\2\u00ae\u00af\7\24\2\2\u00af\u00b0"+
		"\5\22\n\2\u00b0\u00b1\7\17\2\2\u00b1\u00b2\5\22\n\2\u00b2\u00b3\7\25\2"+
		"\2\u00b3\u00b4\b\n\1\2\u00b4\u00cc\3\2\2\2\u00b5\u00b6\7\36\2\2\u00b6"+
		"\u00b7\5\22\n\13\u00b7\u00b8\b\n\1\2\u00b8\u00cc\3\2\2\2\u00b9\u00ba\7"+
		"\24\2\2\u00ba\u00bb\5\22\n\2\u00bb\u00bc\7\25\2\2\u00bc\u00bd\b\n\1\2"+
		"\u00bd\u00cc\3\2\2\2\u00be\u00bf\5\20\t\2\u00bf\u00c0\7\60\2\2\u00c0\u00c1"+
		"\7#\2\2\u00c1\u00c2\b\n\1\2\u00c2\u00cc\3\2\2\2\u00c3\u00c4\5\20\t\2\u00c4"+
		"\u00c5\7\60\2\2\u00c5\u00c6\7$\2\2\u00c6\u00c7\b\n\1\2\u00c7\u00cc\3\2"+
		"\2\2\u00c8\u00c9\5\20\t\2\u00c9\u00ca\b\n\1\2\u00ca\u00cc\3\2\2\2\u00cb"+
		"\u00a0\3\2\2\2\u00cb\u00a5\3\2\2\2\u00cb\u00ad\3\2\2\2\u00cb\u00b5\3\2"+
		"\2\2\u00cb\u00b9\3\2\2\2\u00cb\u00be\3\2\2\2\u00cb\u00c3\3\2\2\2\u00cb"+
		"\u00c8\3\2\2\2\u00cc\u0110\3\2\2\2\u00cd\u00ce\f\27\2\2\u00ce\u00cf\t"+
		"\2\2\2\u00cf\u00d0\5\22\n\30\u00d0\u00d1\b\n\1\2\u00d1\u010f\3\2\2\2\u00d2"+
		"\u00d3\f\26\2\2\u00d3\u00d4\t\3\2\2\u00d4\u00d5\5\22\n\27\u00d5\u00d6"+
		"\b\n\1\2\u00d6\u010f\3\2\2\2\u00d7\u00d8\f\23\2\2\u00d8\u00d9\7\21\2\2"+
		"\u00d9\u00da\5\22\n\24\u00da\u00db\b\n\1\2\u00db\u010f\3\2\2\2\u00dc\u00dd"+
		"\f\22\2\2\u00dd\u00de\7\22\2\2\u00de\u00df\5\22\n\23\u00df\u00e0\b\n\1"+
		"\2\u00e0\u010f\3\2\2\2\u00e1\u00e2\f\21\2\2\u00e2\u00e3\7\b\2\2\u00e3"+
		"\u00e4\5\22\n\22\u00e4\u00e5\b\n\1\2\u00e5\u010f\3\2\2\2\u00e6\u00e7\f"+
		"\20\2\2\u00e7\u00e8\7\t\2\2\u00e8\u00e9\5\22\n\21\u00e9\u00ea\b\n\1\2"+
		"\u00ea\u010f\3\2\2\2\u00eb\u00ec\f\17\2\2\u00ec\u00ed\7\6\2\2\u00ed\u00ee"+
		"\5\22\n\20\u00ee\u00ef\b\n\1\2\u00ef\u010f\3\2\2\2\u00f0\u00f1\f\16\2"+
		"\2\u00f1\u00f2\7\7\2\2\u00f2\u00f3\5\22\n\17\u00f3\u00f4\b\n\1\2\u00f4"+
		"\u010f\3\2\2\2\u00f5\u00f6\f\r\2\2\u00f6\u00f7\7\4\2\2\u00f7\u00f8\5\22"+
		"\n\16\u00f8\u00f9\b\n\1\2\u00f9\u010f\3\2\2\2\u00fa\u00fb\f\f\2\2\u00fb"+
		"\u00fc\7\5\2\2\u00fc\u00fd\5\22\n\r\u00fd\u00fe\b\n\1\2\u00fe\u010f\3"+
		"\2\2\2\u00ff\u0100\f\7\2\2\u0100\u0101\7\23\2\2\u0101\u0102\7\66\2\2\u0102"+
		"\u0103\7\24\2\2\u0103\u0104\7\25\2\2\u0104\u010f\b\n\1\2\u0105\u0106\f"+
		"\6\2\2\u0106\u0107\7\23\2\2\u0107\u010f\7\67\2\2\u0108\u0109\f\5\2\2\u0109"+
		"\u010a\7\23\2\2\u010a\u010f\78\2\2\u010b\u010c\f\4\2\2\u010c\u010d\7\23"+
		"\2\2\u010d\u010f\79\2\2\u010e\u00cd\3\2\2\2\u010e\u00d2\3\2\2\2\u010e"+
		"\u00d7\3\2\2\2\u010e\u00dc\3\2\2\2\u010e\u00e1\3\2\2\2\u010e\u00e6\3\2"+
		"\2\2\u010e\u00eb\3\2\2\2\u010e\u00f0\3\2\2\2\u010e\u00f5\3\2\2\2\u010e"+
		"\u00fa\3\2\2\2\u010e\u00ff\3\2\2\2\u010e\u0105\3\2\2\2\u010e\u0108\3\2"+
		"\2\2\u010e\u010b\3\2\2\2\u010f\u0112\3\2\2\2\u0110\u010e\3\2\2\2\u0110"+
		"\u0111\3\2\2\2\u0111\23\3\2\2\2\u0112\u0110\3\2\2\2\u0113\u0114\7\61\2"+
		"\2\u0114\u0115\7\24\2\2\u0115\u0116\5\22\n\2\u0116\u0117\7\25\2\2\u0117"+
		"\u0118\7\20\2\2\u0118\u0119\b\13\1\2\u0119\u0122\3\2\2\2\u011a\u011b\7"+
		"\61\2\2\u011b\u011c\7\24\2\2\u011c\u011d\5\22\n\2\u011d\u011e\5\26\f\2"+
		"\u011e\u011f\7\20\2\2\u011f\u0120\b\13\1\2\u0120\u0122\3\2\2\2\u0121\u0113"+
		"\3\2\2\2\u0121\u011a\3\2\2\2\u0122\25\3\2\2\2\u0123\u0125\5\30\r\2\u0124"+
		"\u0123\3\2\2\2\u0125\u0128\3\2\2\2\u0126\u0124\3\2\2\2\u0126\u0127\3\2"+
		"\2\2\u0127\u0129\3\2\2\2\u0128\u0126\3\2\2\2\u0129\u012a\b\f\1\2\u012a"+
		"\27\3\2\2\2\u012b\u012c\7\17\2\2\u012c\u012d\5\22\n\2\u012d\u012e\b\r"+
		"\1\2\u012e\31\3\2\2\2\u012f\u0130\5\34\17\2\u0130\u0131\b\16\1\2\u0131"+
		"\33\3\2\2\2\u0132\u0133\7C\2\2\u0133\u0134\5\22\n\2\u0134\u0135\5\"\22"+
		"\2\u0135\u0136\b\17\1\2\u0136\u014d\3\2\2\2\u0137\u0138\7C\2\2\u0138\u0139"+
		"\5\22\n\2\u0139\u013a\5\"\22\2\u013a\u013b\7E\2\2\u013b\u013c\5\"\22\2"+
		"\u013c\u013d\b\17\1\2\u013d\u014d\3\2\2\2\u013e\u013f\7C\2\2\u013f\u0140"+
		"\5\22\n\2\u0140\u0141\5\"\22\2\u0141\u0142\5 \21\2\u0142\u0143\7E\2\2"+
		"\u0143\u0144\5\"\22\2\u0144\u0145\b\17\1\2\u0145\u014d\3\2\2\2\u0146\u0147"+
		"\7C\2\2\u0147\u0148\5\22\n\2\u0148\u0149\5\"\22\2\u0149\u014a\5 \21\2"+
		"\u014a\u014b\b\17\1\2\u014b\u014d\3\2\2\2\u014c\u0132\3\2\2\2\u014c\u0137"+
		"\3\2\2\2\u014c\u013e\3\2\2\2\u014c\u0146\3\2\2\2\u014d\35\3\2\2\2\u014e"+
		"\u014f\7D\2\2\u014f\u0150\5\22\n\2\u0150\u0151\5\"\22\2\u0151\u0152\b"+
		"\20\1\2\u0152\37\3\2\2\2\u0153\u0155\5\36\20\2\u0154\u0153\3\2\2\2\u0155"+
		"\u0158\3\2\2\2\u0156\u0154\3\2\2\2\u0156\u0157\3\2\2\2\u0157\u0159\3\2"+
		"\2\2\u0158\u0156\3\2\2\2\u0159\u015a\b\21\1\2\u015a!\3\2\2\2\u015b\u015c"+
		"\7\n\2\2\u015c\u015d\5\4\3\2\u015d\u015e\7\13\2\2\u015e\u015f\b\22\1\2"+
		"\u015f\u0164\3\2\2\2\u0160\u0161\7\n\2\2\u0161\u0162\7\13\2\2\u0162\u0164"+
		"\b\22\1\2\u0163\u015b\3\2\2\2\u0163\u0160\3\2\2\2\u0164#\3\2\2\2\u0165"+
		"\u0166\5&\24\2\u0166\u0167\b\23\1\2\u0167\u016c\3\2\2\2\u0168\u0169\5"+
		"(\25\2\u0169\u016a\b\23\1\2\u016a\u016c\3\2\2\2\u016b\u0165\3\2\2\2\u016b"+
		"\u0168\3\2\2\2\u016c%\3\2\2\2\u016d\u016e\7F\2\2\u016e\u016f\5\22\n\2"+
		"\u016f\u0170\5\"\22\2\u0170\u0171\b\24\1\2\u0171\'\3\2\2\2\u0172\u0173"+
		"\7H\2\2\u0173\u0174\5\"\22\2\u0174\u0175\b\25\1\2\u0175)\3\2\2\2\20\66"+
		"Px\u0086\u009e\u00cb\u010e\u0110\u0121\u0126\u014c\u0156\u0163\u016b";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}