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
		RULE_start = 0, RULE_funciones = 1, RULE_main = 2, RULE_instrucciones = 3, 
		RULE_instruccion = 4, RULE_control = 5, RULE_declaracion = 6, RULE_tipovariable = 7, 
		RULE_identificadores = 8, RULE_valores = 9, RULE_expresion = 10, RULE_impresion = 11, 
		RULE_impresionexpresion = 12, RULE_expcoma = 13, RULE_condicionales = 14, 
		RULE_funcionif = 15, RULE_funcionelseif = 16, RULE_listaelseif = 17, RULE_bloque = 18, 
		RULE_bucles = 19, RULE_fwhile = 20, RULE_floop = 21;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "funciones", "main", "instrucciones", "instruccion", "control", 
			"declaracion", "tipovariable", "identificadores", "valores", "expresion", 
			"impresion", "impresionexpresion", "expcoma", "condicionales", "funcionif", 
			"funcionelseif", "listaelseif", "bloque", "bucles", "fwhile", "floop"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'->'", "'||'", "'&&'", "'=='", "'!='", "'>='", "'<='", "'{'", 
			"'}'", "'['", "']'", "':'", "','", "';'", "'<'", "'>'", "'.'", "'('", 
			"')'", "'='", "'+'", "'-'", "'*'", "'/'", "'%'", "'|'", "'&'", "'!'", 
			"'?'", null, null, null, "'i64'", "'f64'", "'pow'", "'vec'", "'powf'", 
			"'bool'", "'char'", "'str'", "'String'", "'usize'", "'let'", "'mut'", 
			"'struct'", "'as'", "'println!'", "'true'", "'false'", "'fn'", "'return'", 
			"'abs'", "'sqrt'", "'to_string'", "'clone'", "'new'", "'len'", "'push'", 
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


	        //Temporalgramatica := "Esta en temporal";


	public gramaticaParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public *arrayList.List listainstrucciones;
		public MainContext main;
		public MainContext main() {
			return getRuleContext(MainContext.class,0);
		}
		public List<FuncionesContext> funciones() {
			return getRuleContexts(FuncionesContext.class);
		}
		public FuncionesContext funciones(int i) {
			return getRuleContext(FuncionesContext.class,i);
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
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(47);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(44);
					funciones();
					}
					} 
				}
				setState(49);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			}
			setState(50);
			((StartContext)_localctx).main = main();
			setState(54);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==TKR_fn) {
				{
				{
				setState(51);
				funciones();
				}
				}
				setState(56);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			_localctx.listainstrucciones = ((StartContext)_localctx).main.listainstrucciones
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

	public static class FuncionesContext extends ParserRuleContext {
		public Interfaces.Instruccion inst;
		public Token idd;
		public InstruccionesContext instrucciones;
		public TerminalNode TKR_fn() { return getToken(gramaticaParser.TKR_fn, 0); }
		public TerminalNode TK_par_apertura() { return getToken(gramaticaParser.TK_par_apertura, 0); }
		public TerminalNode TK_par_cierre() { return getToken(gramaticaParser.TK_par_cierre, 0); }
		public TerminalNode TK_corchete_apertura() { return getToken(gramaticaParser.TK_corchete_apertura, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode TK_corchete_cierre() { return getToken(gramaticaParser.TK_corchete_cierre, 0); }
		public TerminalNode TK_id() { return getToken(gramaticaParser.TK_id, 0); }
		public FuncionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funciones; }
	}

	public final FuncionesContext funciones() throws RecognitionException {
		FuncionesContext _localctx = new FuncionesContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_funciones);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(59);
			match(TKR_fn);
			setState(60);
			((FuncionesContext)_localctx).idd = match(TK_id);
			setState(61);
			match(TK_par_apertura);
			setState(62);
			match(TK_par_cierre);
			setState(63);
			match(TK_corchete_apertura);
			setState(64);
			((FuncionesContext)_localctx).instrucciones = instrucciones();
			setState(65);
			match(TK_corchete_cierre);
			Interfaces.NewFuncion("",Interfaces.STRING,((FuncionesContext)_localctx).instrucciones.lista,nil)
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

	public static class MainContext extends ParserRuleContext {
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
		public MainContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_main; }
	}

	public final MainContext main() throws RecognitionException {
		MainContext _localctx = new MainContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_main);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(68);
			match(TKR_fn);
			setState(69);
			match(TKR_main);
			setState(70);
			match(TK_par_apertura);
			setState(71);
			match(TK_par_cierre);
			setState(72);
			match(TK_corchete_apertura);
			setState(73);
			((MainContext)_localctx).instrucciones = instrucciones();
			setState(74);
			match(TK_corchete_cierre);
			_localctx.listainstrucciones = ((MainContext)_localctx).instrucciones.lista
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
		enterRule(_localctx, 6, RULE_instrucciones);

		        _localctx.lista = arrayList.New()

		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(80);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << TK_llave_apertura) | (1L << TK_par_apertura) | (1L << TK_menos) | (1L << TK_amp) | (1L << TK_sig_admiracion) | (1L << TKR_pow) | (1L << TKR_powf) | (1L << TKR_let) | (1L << TKR_println) | (1L << TKR_true) | (1L << TKR_false))) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & ((1L << (TKR_if - 65)) | (1L << (TKR_while - 65)) | (1L << (TKR_break - 65)) | (1L << (TKR_loop - 65)) | (1L << (TK_entero - 65)) | (1L << (TK_decimal - 65)) | (1L << (TK_id - 65)) | (1L << (TK_cadena - 65)) | (1L << (TK_caracter - 65)))) != 0)) {
				{
				{
				setState(77);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).e.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(82);
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
		enterRule(_localctx, 8, RULE_instruccion);
		try {
			setState(106);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(85);
				((InstruccionContext)_localctx).expresion = expresion(0);
				fmt.Println("mensaje en instrucciones: ",((InstruccionContext)_localctx).expresion.exp)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(88);
				((InstruccionContext)_localctx).impresion = impresion();
				_localctx.inst = ((InstruccionContext)_localctx).impresion.inst
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(91);
				((InstruccionContext)_localctx).declaracion = declaracion();
				_localctx.inst = ((InstruccionContext)_localctx).declaracion.inst
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(94);
				((InstruccionContext)_localctx).identificadores = identificadores();
				_localctx.inst = ((InstruccionContext)_localctx).identificadores.inst
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(97);
				((InstruccionContext)_localctx).condicionales = condicionales();
				_localctx.inst = ((InstruccionContext)_localctx).condicionales.inst
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(100);
				((InstruccionContext)_localctx).bucles = bucles();
				_localctx.inst = ((InstruccionContext)_localctx).bucles.inst
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(103);
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
		enterRule(_localctx, 10, RULE_control);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(108);
			match(TKR_break);
			setState(109);
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
		public ValoresContext cant;
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
		public List<TerminalNode> TK_pcoma() { return getTokens(gramaticaParser.TK_pcoma); }
		public TerminalNode TK_pcoma(int i) {
			return getToken(gramaticaParser.TK_pcoma, i);
		}
		public TerminalNode TK_id() { return getToken(gramaticaParser.TK_id, 0); }
		public TerminalNode TK_llave_apertura() { return getToken(gramaticaParser.TK_llave_apertura, 0); }
		public TerminalNode TK_llave_cierre() { return getToken(gramaticaParser.TK_llave_cierre, 0); }
		public ValoresContext valores() {
			return getRuleContext(ValoresContext.class,0);
		}
		public DeclaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion; }
	}

	public final DeclaracionContext declaracion() throws RecognitionException {
		DeclaracionContext _localctx = new DeclaracionContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_declaracion);
		try {
			setState(173);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(112);
				match(TKR_let);
				setState(113);
				match(TKR_mut);
				setState(114);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(115);
				match(TK_dosPuntos);
				setState(116);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(117);
				match(TK_igual);
				setState(118);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(119);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar,((DeclaracionContext)_localctx).expresion.exp,false,false,true) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(122);
				match(TKR_let);
				setState(123);
				match(TKR_mut);
				setState(124);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(125);
				match(TK_igual);
				setState(126);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(127);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),Interfaces.SINTIPO,((DeclaracionContext)_localctx).expresion.exp,false,false,true) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(130);
				match(TKR_let);
				setState(131);
				match(TKR_mut);
				setState(132);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(133);
				match(TK_dosPuntos);
				setState(134);
				match(TK_llave_apertura);
				setState(135);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(136);
				match(TK_pcoma);
				setState(137);
				((DeclaracionContext)_localctx).cant = valores();
				setState(138);
				match(TK_llave_cierre);
				setState(139);
				match(TK_igual);
				setState(140);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(141);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracionArray((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar,((DeclaracionContext)_localctx).cant.exp,((DeclaracionContext)_localctx).expresion.exp,true)
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(144);
				match(TKR_let);
				setState(145);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(146);
				match(TK_dosPuntos);
				setState(147);
				match(TK_llave_apertura);
				setState(148);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(149);
				match(TK_pcoma);
				setState(150);
				((DeclaracionContext)_localctx).cant = valores();
				setState(151);
				match(TK_llave_cierre);
				setState(152);
				match(TK_igual);
				setState(153);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(154);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracionArray((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar,((DeclaracionContext)_localctx).cant.exp,((DeclaracionContext)_localctx).expresion.exp,false)
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(157);
				match(TKR_let);
				setState(158);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(159);
				match(TK_dosPuntos);
				setState(160);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(161);
				match(TK_igual);
				setState(162);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(163);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar,((DeclaracionContext)_localctx).expresion.exp,false,false,false) 
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(166);
				match(TKR_let);
				setState(167);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(168);
				match(TK_igual);
				setState(169);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(170);
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
		enterRule(_localctx, 14, RULE_tipovariable);
		try {
			setState(187);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TKR_numericos_enteros:
				enterOuterAlt(_localctx, 1);
				{
				setState(175);
				match(TKR_numericos_enteros);
				_localctx.tipovar = Interfaces.INTEGER
				}
				break;
			case TKR_numericos_flotantes:
				enterOuterAlt(_localctx, 2);
				{
				setState(177);
				match(TKR_numericos_flotantes);
				_localctx.tipovar = Interfaces.FLOAT
				}
				break;
			case TKR_String:
				enterOuterAlt(_localctx, 3);
				{
				setState(179);
				match(TKR_String);
				_localctx.tipovar = Interfaces.STRING
				}
				break;
			case TKR_bool:
				enterOuterAlt(_localctx, 4);
				{
				setState(181);
				match(TKR_bool);
				_localctx.tipovar = Interfaces.BOOLEAN
				}
				break;
			case TKR_char:
				enterOuterAlt(_localctx, 5);
				{
				setState(183);
				match(TKR_char);
				_localctx.tipovar = Interfaces.CHAR
				}
				break;
			case TKR_usize:
				enterOuterAlt(_localctx, 6);
				{
				setState(185);
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
		enterRule(_localctx, 16, RULE_identificadores);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(189);
			((IdentificadoresContext)_localctx).e1 = match(TK_id);
			setState(190);
			match(TK_igual);
			setState(191);
			((IdentificadoresContext)_localctx).e2 = expresion(0);
			setState(192);
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
		public TerminalNode TK_amp() { return getToken(gramaticaParser.TK_amp, 0); }
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
		enterRule(_localctx, 18, RULE_valores);
		try {
			setState(212);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_entero:
				enterOuterAlt(_localctx, 1);
				{
				setState(195);
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
				setState(197);
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
				setState(199);
				((ValoresContext)_localctx).TK_cadena = match(TK_cadena);

				                str:= (((ValoresContext)_localctx).TK_cadena!=null?((ValoresContext)_localctx).TK_cadena.getText():null)[1:len((((ValoresContext)_localctx).TK_cadena!=null?((ValoresContext)_localctx).TK_cadena.getText():null))-1]
				                _localctx.exp = Expresion.NuevoPrimitivo(str,Interfaces.STRING)
				}
				break;
			case TK_caracter:
				enterOuterAlt(_localctx, 4);
				{
				setState(201);
				((ValoresContext)_localctx).TK_caracter = match(TK_caracter);
				str:= (((ValoresContext)_localctx).TK_caracter!=null?((ValoresContext)_localctx).TK_caracter.getText():null)[1:len((((ValoresContext)_localctx).TK_caracter!=null?((ValoresContext)_localctx).TK_caracter.getText():null))-1]
				                _localctx.exp = Expresion.NuevoPrimitivo(str,Interfaces.CHAR)
				}
				break;
			case TK_amp:
				enterOuterAlt(_localctx, 5);
				{
				setState(203);
				match(TK_amp);
				setState(204);
				((ValoresContext)_localctx).TKR_Str = match(TKR_Str);
				str:= (((ValoresContext)_localctx).TKR_Str!=null?((ValoresContext)_localctx).TKR_Str.getText():null)[1:len((((ValoresContext)_localctx).TKR_Str!=null?((ValoresContext)_localctx).TKR_Str.getText():null))-1]
				                _localctx.exp = Expresion.NuevoPrimitivo(str,Interfaces.STR)
				}
				break;
			case TKR_true:
				enterOuterAlt(_localctx, 6);
				{
				setState(206);
				match(TKR_true);
				_localctx.exp = Expresion.NuevoPrimitivo(true, Interfaces.BOOLEAN)
				}
				break;
			case TKR_false:
				enterOuterAlt(_localctx, 7);
				{
				setState(208);
				match(TKR_false);
				_localctx.exp = Expresion.NuevoPrimitivo(false,Interfaces.BOOLEAN)
				}
				break;
			case TK_id:
				enterOuterAlt(_localctx, 8);
				{
				setState(210);
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
		public ImpresionexpresionContext l1;
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
		public TerminalNode TK_llave_apertura() { return getToken(gramaticaParser.TK_llave_apertura, 0); }
		public TerminalNode TK_pcoma() { return getToken(gramaticaParser.TK_pcoma, 0); }
		public TerminalNode TK_llave_cierre() { return getToken(gramaticaParser.TK_llave_cierre, 0); }
		public ImpresionexpresionContext impresionexpresion() {
			return getRuleContext(ImpresionexpresionContext.class,0);
		}
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
		int _startState = 20;
		enterRecursionRule(_localctx, 20, RULE_expresion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(270);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				{
				setState(215);
				match(TK_menos);
				setState(216);
				((ExpresionContext)_localctx).e1 = expresion(24);
				numero := -1
				                                                                                                        numm := Expresion.NuevoPrimitivo(numero, Interfaces.INTEGER)
				                                                                                                                _localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp, numm ,Interfaces.MULTIPLICACION)
				                                                                                                
				}
				break;
			case 2:
				{
				setState(219);
				match(TKR_pow);
				setState(220);
				match(TK_par_apertura);
				setState(221);
				((ExpresionContext)_localctx).e1 = expresion(0);
				setState(222);
				match(TK_coma);
				setState(223);
				((ExpresionContext)_localctx).e2 = expresion(0);
				setState(224);
				match(TK_par_cierre);
				_localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.POW)
				}
				break;
			case 3:
				{
				setState(227);
				match(TKR_powf);
				setState(228);
				match(TK_par_apertura);
				setState(229);
				((ExpresionContext)_localctx).e1 = expresion(0);
				setState(230);
				match(TK_coma);
				setState(231);
				((ExpresionContext)_localctx).e2 = expresion(0);
				setState(232);
				match(TK_par_cierre);
				_localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.POWF)
				}
				break;
			case 4:
				{
				setState(235);
				match(TK_sig_admiracion);
				setState(236);
				((ExpresionContext)_localctx).e1 = expresion(11);
				_localctx.exp = Expresiones.NuevaLogica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e1.exp,Interfaces.NOT)
				}
				break;
			case 5:
				{
				setState(239);
				match(TK_par_apertura);
				setState(240);
				((ExpresionContext)_localctx).va = expresion(0);
				setState(241);
				match(TK_par_cierre);
				_localctx.exp = ((ExpresionContext)_localctx).va.exp
				}
				break;
			case 6:
				{
				setState(244);
				((ExpresionContext)_localctx).val = valores();
				setState(245);
				match(TKR_as);
				setState(246);
				match(TKR_numericos_enteros);
				_localctx.exp = Expresiones.NewAsi64(((ExpresionContext)_localctx).val.exp)
				}
				break;
			case 7:
				{
				setState(249);
				((ExpresionContext)_localctx).val = valores();
				setState(250);
				match(TKR_as);
				setState(251);
				match(TKR_numericos_flotantes);
				_localctx.exp = Expresiones.NewAsf64(((ExpresionContext)_localctx).val.exp)
				}
				break;
			case 8:
				{
				setState(254);
				match(TK_llave_apertura);
				setState(255);
				((ExpresionContext)_localctx).e1 = expresion(0);
				setState(256);
				match(TK_pcoma);
				setState(257);
				((ExpresionContext)_localctx).e2 = expresion(0);
				setState(258);
				match(TK_llave_cierre);
				_localctx.exp = Expresion.NewArray(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,nil,Interfaces.MULTIPLE)
				}
				break;
			case 9:
				{
				setState(261);
				match(TK_llave_apertura);
				setState(262);
				((ExpresionContext)_localctx).e1 = expresion(0);
				setState(263);
				((ExpresionContext)_localctx).l1 = impresionexpresion();
				setState(264);
				match(TK_llave_cierre);
				_localctx.exp = Expresion.NewArray(((ExpresionContext)_localctx).e1.exp,nil,((ExpresionContext)_localctx).l1.lista,Interfaces.NORMAL)
				}
				break;
			case 10:
				{
				setState(267);
				((ExpresionContext)_localctx).vall = valores();
				_localctx.exp = ((ExpresionContext)_localctx).vall.exp
				                                                                                        fmt.Println(_localctx.exp)
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(347);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(345);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
					case 1:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(272);
						if (!(precpred(_ctx, 23))) throw new FailedPredicateException(this, "precpred(_ctx, 23)");
						setState(273);
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
						setState(274);
						((ExpresionContext)_localctx).e2 = expresion(24);

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
						setState(277);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(278);
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
						setState(279);
						((ExpresionContext)_localctx).e2 = expresion(23);

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
						setState(282);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(283);
						match(TK_menor);
						setState(284);
						((ExpresionContext)_localctx).e2 = expresion(20);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.MENOR_QUE)
						}
						break;
					case 4:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(287);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(288);
						match(TK_mayor);
						setState(289);
						((ExpresionContext)_localctx).e2 = expresion(19);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.MAYOR_QUE)
						}
						break;
					case 5:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(292);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(293);
						match(TK_mayor_igual);
						setState(294);
						((ExpresionContext)_localctx).e2 = expresion(18);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.MAYOR_IGUAL)
						}
						break;
					case 6:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(297);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(298);
						match(TK_menor_igual);
						setState(299);
						((ExpresionContext)_localctx).e2 = expresion(17);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.MENOR_IGUAL)
						}
						break;
					case 7:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(302);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(303);
						match(TK_igualacion);
						setState(304);
						((ExpresionContext)_localctx).e2 = expresion(16);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.IGUALDAD)
						}
						break;
					case 8:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(307);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(308);
						match(TK_diferente);
						setState(309);
						((ExpresionContext)_localctx).e2 = expresion(15);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.DESIGUAL)
						}
						break;
					case 9:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(312);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(313);
						match(TK_or);
						setState(314);
						((ExpresionContext)_localctx).e2 = expresion(14);
						_localctx.exp = Expresiones.NuevaLogica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.OR)
						}
						break;
					case 10:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(317);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(318);
						match(TK_and);
						setState(319);
						((ExpresionContext)_localctx).e2 = expresion(13);
						_localctx.exp = Expresiones.NuevaLogica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.AND)
						}
						break;
					case 11:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.vll = _prevctx;
						_localctx.vll = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(322);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(323);
						match(TK_punto);
						setState(324);
						match(TKR_abs);
						setState(325);
						match(TK_par_apertura);
						setState(326);
						match(TK_par_cierre);
						_localctx.exp = Expresiones.Newabs(((ExpresionContext)_localctx).vll.exp)
						}
						break;
					case 12:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(328);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(329);
						match(TK_punto);
						setState(330);
						match(TKR_sqrt);
						setState(331);
						match(TK_par_apertura);
						setState(332);
						match(TK_par_cierre);
						_localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e1.exp,Interfaces.MULTIPLICACION)
						}
						break;
					case 13:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.vll = _prevctx;
						_localctx.vll = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(334);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(335);
						match(TK_punto);
						setState(336);
						match(TKR_to_string);
						setState(337);
						match(TK_par_apertura);
						setState(338);
						match(TK_par_cierre);
						_localctx.exp = Expresiones.NewFto_string(((ExpresionContext)_localctx).vll.exp)
						}
						break;
					case 14:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.vll = _prevctx;
						_localctx.vll = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(340);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(341);
						match(TK_punto);
						setState(342);
						match(TKR_clone);
						setState(343);
						match(TK_par_apertura);
						setState(344);
						match(TK_par_cierre);
						}
						break;
					}
					} 
				}
				setState(349);
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
		enterRule(_localctx, 22, RULE_impresion);
		try {
			setState(365);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(350);
				match(TKR_println);
				setState(351);
				match(TK_par_apertura);
				setState(352);
				((ImpresionContext)_localctx).e1 = expresion(0);
				setState(353);
				match(TK_par_cierre);
				setState(354);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevoPrint(((ImpresionContext)_localctx).e1.exp,nil)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(357);
				match(TKR_println);
				setState(358);
				match(TK_par_apertura);
				setState(359);
				((ImpresionContext)_localctx).e2 = expresion(0);
				setState(360);
				((ImpresionContext)_localctx).li = impresionexpresion();
				setState(361);
				match(TK_par_cierre);
				setState(362);
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
		enterRule(_localctx, 24, RULE_impresionexpresion);
		_localctx.lista = arrayList.New()
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(370);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==TK_coma) {
				{
				{
				setState(367);
				((ImpresionexpresionContext)_localctx).expcoma = expcoma();
				((ImpresionexpresionContext)_localctx).list.add(((ImpresionexpresionContext)_localctx).expcoma);
				}
				}
				setState(372);
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
		enterRule(_localctx, 26, RULE_expcoma);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(375);
			match(TK_coma);
			setState(376);
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
		enterRule(_localctx, 28, RULE_condicionales);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(379);
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
		enterRule(_localctx, 30, RULE_funcionif);
		try {
			setState(408);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(382);
				match(TKR_if);
				setState(383);
				((FuncionifContext)_localctx).e1 = expresion(0);
				setState(384);
				((FuncionifContext)_localctx).ee = bloque();
				_localctx.inst = Instruccion.NewIf(((FuncionifContext)_localctx).e1.exp,((FuncionifContext)_localctx).ee.lista,nil,nil )
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(387);
				match(TKR_if);
				setState(388);
				((FuncionifContext)_localctx).e1 = expresion(0);
				setState(389);
				((FuncionifContext)_localctx).e5 = bloque();
				setState(390);
				match(TKR_else);
				setState(391);
				((FuncionifContext)_localctx).b2 = bloque();
				_localctx.inst = Instruccion.NewIf(((FuncionifContext)_localctx).e1.exp,((FuncionifContext)_localctx).e5.lista,nil,((FuncionifContext)_localctx).b2.lista)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(394);
				match(TKR_if);
				setState(395);
				((FuncionifContext)_localctx).e1 = expresion(0);
				setState(396);
				((FuncionifContext)_localctx).b1 = bloque();
				setState(397);
				((FuncionifContext)_localctx).listaelseif = listaelseif();
				setState(398);
				match(TKR_else);
				setState(399);
				((FuncionifContext)_localctx).b2 = bloque();
				_localctx.inst = Instruccion.NewIf(((FuncionifContext)_localctx).e1.exp,((FuncionifContext)_localctx).b1.lista,((FuncionifContext)_localctx).listaelseif.lista,((FuncionifContext)_localctx).b2.lista)
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(402);
				match(TKR_if);
				setState(403);
				((FuncionifContext)_localctx).e1 = expresion(0);
				setState(404);
				((FuncionifContext)_localctx).b1 = bloque();
				setState(405);
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
		enterRule(_localctx, 32, RULE_funcionelseif);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(410);
			match(TKR_elseif);
			setState(411);
			((FuncionelseifContext)_localctx).e1 = expresion(0);
			setState(412);
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
		enterRule(_localctx, 34, RULE_listaelseif);
		 _localctx.lista = arrayList.New()
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(418);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==TKR_elseif) {
				{
				{
				setState(415);
				((ListaelseifContext)_localctx).funcionelseif = funcionelseif();
				((ListaelseifContext)_localctx).list.add(((ListaelseifContext)_localctx).funcionelseif);
				}
				}
				setState(420);
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
		enterRule(_localctx, 36, RULE_bloque);
		try {
			setState(431);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(423);
				match(TK_corchete_apertura);
				setState(424);
				((BloqueContext)_localctx).instrucciones = instrucciones();
				setState(425);
				match(TK_corchete_cierre);
				_localctx.lista = ((BloqueContext)_localctx).instrucciones.lista
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(428);
				match(TK_corchete_apertura);
				setState(429);
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
		enterRule(_localctx, 38, RULE_bucles);
		try {
			setState(439);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TKR_while:
				enterOuterAlt(_localctx, 1);
				{
				setState(433);
				((BuclesContext)_localctx).fwhile = fwhile();
				_localctx.inst = ((BuclesContext)_localctx).fwhile.inst
				}
				break;
			case TKR_loop:
				enterOuterAlt(_localctx, 2);
				{
				setState(436);
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
		enterRule(_localctx, 40, RULE_fwhile);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(441);
			match(TKR_while);
			setState(442);
			((FwhileContext)_localctx).e1 = expresion(0);
			setState(443);
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
		enterRule(_localctx, 42, RULE_floop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(446);
			match(TKR_loop);
			setState(447);
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
		case 10:
			return expresion_sempred((ExpresionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expresion_sempred(ExpresionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 23);
		case 1:
			return precpred(_ctx, 22);
		case 2:
			return precpred(_ctx, 19);
		case 3:
			return precpred(_ctx, 18);
		case 4:
			return precpred(_ctx, 17);
		case 5:
			return precpred(_ctx, 16);
		case 6:
			return precpred(_ctx, 15);
		case 7:
			return precpred(_ctx, 14);
		case 8:
			return precpred(_ctx, 13);
		case 9:
			return precpred(_ctx, 12);
		case 10:
			return precpred(_ctx, 7);
		case 11:
			return precpred(_ctx, 6);
		case 12:
			return precpred(_ctx, 5);
		case 13:
			return precpred(_ctx, 4);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3N\u01c5\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\3\2\7\2\60\n\2\f\2"+
		"\16\2\63\13\2\3\2\3\2\7\2\67\n\2\f\2\16\2:\13\2\3\2\3\2\3\3\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5\7\5Q\n"+
		"\5\f\5\16\5T\13\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6m\n\6\3\7\3\7\3\7\3\7\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\5\b\u00b0\n\b\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\5\t\u00be\n\t\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\5\13\u00d7\n\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\5\f\u0111\n\f\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\7\f\u015c"+
		"\n\f\f\f\16\f\u015f\13\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3"+
		"\r\3\r\3\r\3\r\5\r\u0170\n\r\3\16\7\16\u0173\n\16\f\16\16\16\u0176\13"+
		"\16\3\16\3\16\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3"+
		"\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3"+
		"\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\5\21\u019b\n\21\3\22\3\22\3\22"+
		"\3\22\3\22\3\23\7\23\u01a3\n\23\f\23\16\23\u01a6\13\23\3\23\3\23\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\5\24\u01b2\n\24\3\25\3\25\3\25\3\25"+
		"\3\25\3\25\5\25\u01ba\n\25\3\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27"+
		"\3\27\2\3\26\30\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,\2\4\3\2"+
		"\31\33\3\2\27\30\2\u01e7\2\61\3\2\2\2\4=\3\2\2\2\6F\3\2\2\2\bR\3\2\2\2"+
		"\nl\3\2\2\2\fn\3\2\2\2\16\u00af\3\2\2\2\20\u00bd\3\2\2\2\22\u00bf\3\2"+
		"\2\2\24\u00d6\3\2\2\2\26\u0110\3\2\2\2\30\u016f\3\2\2\2\32\u0174\3\2\2"+
		"\2\34\u0179\3\2\2\2\36\u017d\3\2\2\2 \u019a\3\2\2\2\"\u019c\3\2\2\2$\u01a4"+
		"\3\2\2\2&\u01b1\3\2\2\2(\u01b9\3\2\2\2*\u01bb\3\2\2\2,\u01c0\3\2\2\2."+
		"\60\5\4\3\2/.\3\2\2\2\60\63\3\2\2\2\61/\3\2\2\2\61\62\3\2\2\2\62\64\3"+
		"\2\2\2\63\61\3\2\2\2\648\5\6\4\2\65\67\5\4\3\2\66\65\3\2\2\2\67:\3\2\2"+
		"\28\66\3\2\2\289\3\2\2\29;\3\2\2\2:8\3\2\2\2;<\b\2\1\2<\3\3\2\2\2=>\7"+
		"\64\2\2>?\7K\2\2?@\7\24\2\2@A\7\25\2\2AB\7\n\2\2BC\5\b\5\2CD\7\13\2\2"+
		"DE\b\3\1\2E\5\3\2\2\2FG\7\64\2\2GH\7B\2\2HI\7\24\2\2IJ\7\25\2\2JK\7\n"+
		"\2\2KL\5\b\5\2LM\7\13\2\2MN\b\4\1\2N\7\3\2\2\2OQ\5\n\6\2PO\3\2\2\2QT\3"+
		"\2\2\2RP\3\2\2\2RS\3\2\2\2SU\3\2\2\2TR\3\2\2\2UV\b\5\1\2V\t\3\2\2\2WX"+
		"\5\26\f\2XY\b\6\1\2Ym\3\2\2\2Z[\5\30\r\2[\\\b\6\1\2\\m\3\2\2\2]^\5\16"+
		"\b\2^_\b\6\1\2_m\3\2\2\2`a\5\22\n\2ab\b\6\1\2bm\3\2\2\2cd\5\36\20\2de"+
		"\b\6\1\2em\3\2\2\2fg\5(\25\2gh\b\6\1\2hm\3\2\2\2ij\5\f\7\2jk\b\6\1\2k"+
		"m\3\2\2\2lW\3\2\2\2lZ\3\2\2\2l]\3\2\2\2l`\3\2\2\2lc\3\2\2\2lf\3\2\2\2"+
		"li\3\2\2\2m\13\3\2\2\2no\7G\2\2op\7\20\2\2pq\b\7\1\2q\r\3\2\2\2rs\7-\2"+
		"\2st\7.\2\2tu\7K\2\2uv\7\16\2\2vw\5\20\t\2wx\7\26\2\2xy\5\26\f\2yz\7\20"+
		"\2\2z{\b\b\1\2{\u00b0\3\2\2\2|}\7-\2\2}~\7.\2\2~\177\7K\2\2\177\u0080"+
		"\7\26\2\2\u0080\u0081\5\26\f\2\u0081\u0082\7\20\2\2\u0082\u0083\b\b\1"+
		"\2\u0083\u00b0\3\2\2\2\u0084\u0085\7-\2\2\u0085\u0086\7.\2\2\u0086\u0087"+
		"\7K\2\2\u0087\u0088\7\16\2\2\u0088\u0089\7\f\2\2\u0089\u008a\5\20\t\2"+
		"\u008a\u008b\7\20\2\2\u008b\u008c\5\24\13\2\u008c\u008d\7\r\2\2\u008d"+
		"\u008e\7\26\2\2\u008e\u008f\5\26\f\2\u008f\u0090\7\20\2\2\u0090\u0091"+
		"\b\b\1\2\u0091\u00b0\3\2\2\2\u0092\u0093\7-\2\2\u0093\u0094\7K\2\2\u0094"+
		"\u0095\7\16\2\2\u0095\u0096\7\f\2\2\u0096\u0097\5\20\t\2\u0097\u0098\7"+
		"\20\2\2\u0098\u0099\5\24\13\2\u0099\u009a\7\r\2\2\u009a\u009b\7\26\2\2"+
		"\u009b\u009c\5\26\f\2\u009c\u009d\7\20\2\2\u009d\u009e\b\b\1\2\u009e\u00b0"+
		"\3\2\2\2\u009f\u00a0\7-\2\2\u00a0\u00a1\7K\2\2\u00a1\u00a2\7\16\2\2\u00a2"+
		"\u00a3\5\20\t\2\u00a3\u00a4\7\26\2\2\u00a4\u00a5\5\26\f\2\u00a5\u00a6"+
		"\7\20\2\2\u00a6\u00a7\b\b\1\2\u00a7\u00b0\3\2\2\2\u00a8\u00a9\7-\2\2\u00a9"+
		"\u00aa\7K\2\2\u00aa\u00ab\7\26\2\2\u00ab\u00ac\5\26\f\2\u00ac\u00ad\7"+
		"\20\2\2\u00ad\u00ae\b\b\1\2\u00ae\u00b0\3\2\2\2\u00afr\3\2\2\2\u00af|"+
		"\3\2\2\2\u00af\u0084\3\2\2\2\u00af\u0092\3\2\2\2\u00af\u009f\3\2\2\2\u00af"+
		"\u00a8\3\2\2\2\u00b0\17\3\2\2\2\u00b1\u00b2\7#\2\2\u00b2\u00be\b\t\1\2"+
		"\u00b3\u00b4\7$\2\2\u00b4\u00be\b\t\1\2\u00b5\u00b6\7+\2\2\u00b6\u00be"+
		"\b\t\1\2\u00b7\u00b8\7(\2\2\u00b8\u00be\b\t\1\2\u00b9\u00ba\7)\2\2\u00ba"+
		"\u00be\b\t\1\2\u00bb\u00bc\7,\2\2\u00bc\u00be\b\t\1\2\u00bd\u00b1\3\2"+
		"\2\2\u00bd\u00b3\3\2\2\2\u00bd\u00b5\3\2\2\2\u00bd\u00b7\3\2\2\2\u00bd"+
		"\u00b9\3\2\2\2\u00bd\u00bb\3\2\2\2\u00be\21\3\2\2\2\u00bf\u00c0\7K\2\2"+
		"\u00c0\u00c1\7\26\2\2\u00c1\u00c2\5\26\f\2\u00c2\u00c3\7\20\2\2\u00c3"+
		"\u00c4\b\n\1\2\u00c4\23\3\2\2\2\u00c5\u00c6\7I\2\2\u00c6\u00d7\b\13\1"+
		"\2\u00c7\u00c8\7J\2\2\u00c8\u00d7\b\13\1\2\u00c9\u00ca\7M\2\2\u00ca\u00d7"+
		"\b\13\1\2\u00cb\u00cc\7N\2\2\u00cc\u00d7\b\13\1\2\u00cd\u00ce\7\35\2\2"+
		"\u00ce\u00cf\7*\2\2\u00cf\u00d7\b\13\1\2\u00d0\u00d1\7\62\2\2\u00d1\u00d7"+
		"\b\13\1\2\u00d2\u00d3\7\63\2\2\u00d3\u00d7\b\13\1\2\u00d4\u00d5\7K\2\2"+
		"\u00d5\u00d7\b\13\1\2\u00d6\u00c5\3\2\2\2\u00d6\u00c7\3\2\2\2\u00d6\u00c9"+
		"\3\2\2\2\u00d6\u00cb\3\2\2\2\u00d6\u00cd\3\2\2\2\u00d6\u00d0\3\2\2\2\u00d6"+
		"\u00d2\3\2\2\2\u00d6\u00d4\3\2\2\2\u00d7\25\3\2\2\2\u00d8\u00d9\b\f\1"+
		"\2\u00d9\u00da\7\30\2\2\u00da\u00db\5\26\f\32\u00db\u00dc\b\f\1\2\u00dc"+
		"\u0111\3\2\2\2\u00dd\u00de\7%\2\2\u00de\u00df\7\24\2\2\u00df\u00e0\5\26"+
		"\f\2\u00e0\u00e1\7\17\2\2\u00e1\u00e2\5\26\f\2\u00e2\u00e3\7\25\2\2\u00e3"+
		"\u00e4\b\f\1\2\u00e4\u0111\3\2\2\2\u00e5\u00e6\7\'\2\2\u00e6\u00e7\7\24"+
		"\2\2\u00e7\u00e8\5\26\f\2\u00e8\u00e9\7\17\2\2\u00e9\u00ea\5\26\f\2\u00ea"+
		"\u00eb\7\25\2\2\u00eb\u00ec\b\f\1\2\u00ec\u0111\3\2\2\2\u00ed\u00ee\7"+
		"\36\2\2\u00ee\u00ef\5\26\f\r\u00ef\u00f0\b\f\1\2\u00f0\u0111\3\2\2\2\u00f1"+
		"\u00f2\7\24\2\2\u00f2\u00f3\5\26\f\2\u00f3\u00f4\7\25\2\2\u00f4\u00f5"+
		"\b\f\1\2\u00f5\u0111\3\2\2\2\u00f6\u00f7\5\24\13\2\u00f7\u00f8\7\60\2"+
		"\2\u00f8\u00f9\7#\2\2\u00f9\u00fa\b\f\1\2\u00fa\u0111\3\2\2\2\u00fb\u00fc"+
		"\5\24\13\2\u00fc\u00fd\7\60\2\2\u00fd\u00fe\7$\2\2\u00fe\u00ff\b\f\1\2"+
		"\u00ff\u0111\3\2\2\2\u0100\u0101\7\f\2\2\u0101\u0102\5\26\f\2\u0102\u0103"+
		"\7\20\2\2\u0103\u0104\5\26\f\2\u0104\u0105\7\r\2\2\u0105\u0106\b\f\1\2"+
		"\u0106\u0111\3\2\2\2\u0107\u0108\7\f\2\2\u0108\u0109\5\26\f\2\u0109\u010a"+
		"\5\32\16\2\u010a\u010b\7\r\2\2\u010b\u010c\b\f\1\2\u010c\u0111\3\2\2\2"+
		"\u010d\u010e\5\24\13\2\u010e\u010f\b\f\1\2\u010f\u0111\3\2\2\2\u0110\u00d8"+
		"\3\2\2\2\u0110\u00dd\3\2\2\2\u0110\u00e5\3\2\2\2\u0110\u00ed\3\2\2\2\u0110"+
		"\u00f1\3\2\2\2\u0110\u00f6\3\2\2\2\u0110\u00fb\3\2\2\2\u0110\u0100\3\2"+
		"\2\2\u0110\u0107\3\2\2\2\u0110\u010d\3\2\2\2\u0111\u015d\3\2\2\2\u0112"+
		"\u0113\f\31\2\2\u0113\u0114\t\2\2\2\u0114\u0115\5\26\f\32\u0115\u0116"+
		"\b\f\1\2\u0116\u015c\3\2\2\2\u0117\u0118\f\30\2\2\u0118\u0119\t\3\2\2"+
		"\u0119\u011a\5\26\f\31\u011a\u011b\b\f\1\2\u011b\u015c\3\2\2\2\u011c\u011d"+
		"\f\25\2\2\u011d\u011e\7\21\2\2\u011e\u011f\5\26\f\26\u011f\u0120\b\f\1"+
		"\2\u0120\u015c\3\2\2\2\u0121\u0122\f\24\2\2\u0122\u0123\7\22\2\2\u0123"+
		"\u0124\5\26\f\25\u0124\u0125\b\f\1\2\u0125\u015c\3\2\2\2\u0126\u0127\f"+
		"\23\2\2\u0127\u0128\7\b\2\2\u0128\u0129\5\26\f\24\u0129\u012a\b\f\1\2"+
		"\u012a\u015c\3\2\2\2\u012b\u012c\f\22\2\2\u012c\u012d\7\t\2\2\u012d\u012e"+
		"\5\26\f\23\u012e\u012f\b\f\1\2\u012f\u015c\3\2\2\2\u0130\u0131\f\21\2"+
		"\2\u0131\u0132\7\6\2\2\u0132\u0133\5\26\f\22\u0133\u0134\b\f\1\2\u0134"+
		"\u015c\3\2\2\2\u0135\u0136\f\20\2\2\u0136\u0137\7\7\2\2\u0137\u0138\5"+
		"\26\f\21\u0138\u0139\b\f\1\2\u0139\u015c\3\2\2\2\u013a\u013b\f\17\2\2"+
		"\u013b\u013c\7\4\2\2\u013c\u013d\5\26\f\20\u013d\u013e\b\f\1\2\u013e\u015c"+
		"\3\2\2\2\u013f\u0140\f\16\2\2\u0140\u0141\7\5\2\2\u0141\u0142\5\26\f\17"+
		"\u0142\u0143\b\f\1\2\u0143\u015c\3\2\2\2\u0144\u0145\f\t\2\2\u0145\u0146"+
		"\7\23\2\2\u0146\u0147\7\66\2\2\u0147\u0148\7\24\2\2\u0148\u0149\7\25\2"+
		"\2\u0149\u015c\b\f\1\2\u014a\u014b\f\b\2\2\u014b\u014c\7\23\2\2\u014c"+
		"\u014d\7\67\2\2\u014d\u014e\7\24\2\2\u014e\u014f\7\25\2\2\u014f\u015c"+
		"\b\f\1\2\u0150\u0151\f\7\2\2\u0151\u0152\7\23\2\2\u0152\u0153\78\2\2\u0153"+
		"\u0154\7\24\2\2\u0154\u0155\7\25\2\2\u0155\u015c\b\f\1\2\u0156\u0157\f"+
		"\6\2\2\u0157\u0158\7\23\2\2\u0158\u0159\79\2\2\u0159\u015a\7\24\2\2\u015a"+
		"\u015c\7\25\2\2\u015b\u0112\3\2\2\2\u015b\u0117\3\2\2\2\u015b\u011c\3"+
		"\2\2\2\u015b\u0121\3\2\2\2\u015b\u0126\3\2\2\2\u015b\u012b\3\2\2\2\u015b"+
		"\u0130\3\2\2\2\u015b\u0135\3\2\2\2\u015b\u013a\3\2\2\2\u015b\u013f\3\2"+
		"\2\2\u015b\u0144\3\2\2\2\u015b\u014a\3\2\2\2\u015b\u0150\3\2\2\2\u015b"+
		"\u0156\3\2\2\2\u015c\u015f\3\2\2\2\u015d\u015b\3\2\2\2\u015d\u015e\3\2"+
		"\2\2\u015e\27\3\2\2\2\u015f\u015d\3\2\2\2\u0160\u0161\7\61\2\2\u0161\u0162"+
		"\7\24\2\2\u0162\u0163\5\26\f\2\u0163\u0164\7\25\2\2\u0164\u0165\7\20\2"+
		"\2\u0165\u0166\b\r\1\2\u0166\u0170\3\2\2\2\u0167\u0168\7\61\2\2\u0168"+
		"\u0169\7\24\2\2\u0169\u016a\5\26\f\2\u016a\u016b\5\32\16\2\u016b\u016c"+
		"\7\25\2\2\u016c\u016d\7\20\2\2\u016d\u016e\b\r\1\2\u016e\u0170\3\2\2\2"+
		"\u016f\u0160\3\2\2\2\u016f\u0167\3\2\2\2\u0170\31\3\2\2\2\u0171\u0173"+
		"\5\34\17\2\u0172\u0171\3\2\2\2\u0173\u0176\3\2\2\2\u0174\u0172\3\2\2\2"+
		"\u0174\u0175\3\2\2\2\u0175\u0177\3\2\2\2\u0176\u0174\3\2\2\2\u0177\u0178"+
		"\b\16\1\2\u0178\33\3\2\2\2\u0179\u017a\7\17\2\2\u017a\u017b\5\26\f\2\u017b"+
		"\u017c\b\17\1\2\u017c\35\3\2\2\2\u017d\u017e\5 \21\2\u017e\u017f\b\20"+
		"\1\2\u017f\37\3\2\2\2\u0180\u0181\7C\2\2\u0181\u0182\5\26\f\2\u0182\u0183"+
		"\5&\24\2\u0183\u0184\b\21\1\2\u0184\u019b\3\2\2\2\u0185\u0186\7C\2\2\u0186"+
		"\u0187\5\26\f\2\u0187\u0188\5&\24\2\u0188\u0189\7E\2\2\u0189\u018a\5&"+
		"\24\2\u018a\u018b\b\21\1\2\u018b\u019b\3\2\2\2\u018c\u018d\7C\2\2\u018d"+
		"\u018e\5\26\f\2\u018e\u018f\5&\24\2\u018f\u0190\5$\23\2\u0190\u0191\7"+
		"E\2\2\u0191\u0192\5&\24\2\u0192\u0193\b\21\1\2\u0193\u019b\3\2\2\2\u0194"+
		"\u0195\7C\2\2\u0195\u0196\5\26\f\2\u0196\u0197\5&\24\2\u0197\u0198\5$"+
		"\23\2\u0198\u0199\b\21\1\2\u0199\u019b\3\2\2\2\u019a\u0180\3\2\2\2\u019a"+
		"\u0185\3\2\2\2\u019a\u018c\3\2\2\2\u019a\u0194\3\2\2\2\u019b!\3\2\2\2"+
		"\u019c\u019d\7D\2\2\u019d\u019e\5\26\f\2\u019e\u019f\5&\24\2\u019f\u01a0"+
		"\b\22\1\2\u01a0#\3\2\2\2\u01a1\u01a3\5\"\22\2\u01a2\u01a1\3\2\2\2\u01a3"+
		"\u01a6\3\2\2\2\u01a4\u01a2\3\2\2\2\u01a4\u01a5\3\2\2\2\u01a5\u01a7\3\2"+
		"\2\2\u01a6\u01a4\3\2\2\2\u01a7\u01a8\b\23\1\2\u01a8%\3\2\2\2\u01a9\u01aa"+
		"\7\n\2\2\u01aa\u01ab\5\b\5\2\u01ab\u01ac\7\13\2\2\u01ac\u01ad\b\24\1\2"+
		"\u01ad\u01b2\3\2\2\2\u01ae\u01af\7\n\2\2\u01af\u01b0\7\13\2\2\u01b0\u01b2"+
		"\b\24\1\2\u01b1\u01a9\3\2\2\2\u01b1\u01ae\3\2\2\2\u01b2\'\3\2\2\2\u01b3"+
		"\u01b4\5*\26\2\u01b4\u01b5\b\25\1\2\u01b5\u01ba\3\2\2\2\u01b6\u01b7\5"+
		",\27\2\u01b7\u01b8\b\25\1\2\u01b8\u01ba\3\2\2\2\u01b9\u01b3\3\2\2\2\u01b9"+
		"\u01b6\3\2\2\2\u01ba)\3\2\2\2\u01bb\u01bc\7F\2\2\u01bc\u01bd\5\26\f\2"+
		"\u01bd\u01be\5&\24\2\u01be\u01bf\b\26\1\2\u01bf+\3\2\2\2\u01c0\u01c1\7"+
		"H\2\2\u01c1\u01c2\5&\24\2\u01c2\u01c3\b\27\1\2\u01c3-\3\2\2\2\22\618R"+
		"l\u00af\u00bd\u00d6\u0110\u015b\u015d\u016f\u0174\u019a\u01a4\u01b1\u01b9";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}