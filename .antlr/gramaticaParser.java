// Generated from e:\WillOP\u005Cu\Compi2\Laboratorio\Proyecto1_201408419\gramatica.g4 by ANTLR 4.8

    import "proyecto1/Interfaces"
    import "proyecto1/Expresion"
    import "proyecto1/Instruccion"
    import "proyecto1/Expresiones"
    import "proyecto1/Enum"
    import "proyecto1/Funcion"
    //import "proyecto1/Simbolo"
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
		RULE_start = 0, RULE_funciones = 1, RULE_declararParametros = 2, RULE_main = 3, 
		RULE_instrucciones = 4, RULE_instruccion = 5, RULE_asignacion = 6, RULE_control = 7, 
		RULE_declaracion = 8, RULE_tipovariable = 9, RULE_valores = 10, RULE_arreglo = 11, 
		RULE_expresion = 12, RULE_impresion = 13, RULE_impresionexpresion = 14, 
		RULE_expcoma = 15, RULE_condicionales = 16, RULE_funcionif = 17, RULE_funcionelseif = 18, 
		RULE_listaelseif = 19, RULE_bloque = 20, RULE_bucles = 21, RULE_fwhile = 22, 
		RULE_floop = 23;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "funciones", "declararParametros", "main", "instrucciones", 
			"instruccion", "asignacion", "control", "declaracion", "tipovariable", 
			"valores", "arreglo", "expresion", "impresion", "impresionexpresion", 
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
			"'?'", null, null, null, "'i64'", "'f64'", "'pow'", "'Vec'", "'powf'", 
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
		public FuncionesContext funciones;
		public List<FuncionesContext> e = new ArrayList<FuncionesContext>();
		public MainContext listtt;
		public List<FuncionesContext> r = new ArrayList<FuncionesContext>();
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
			setState(51);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(48);
					((StartContext)_localctx).funciones = funciones();
					((StartContext)_localctx).e.add(((StartContext)_localctx).funciones);
					}
					} 
				}
				setState(53);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			}
			setState(54);
			((StartContext)_localctx).listtt = main();
			setState(58);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==TKR_fn) {
				{
				{
				setState(55);
				((StartContext)_localctx).funciones = funciones();
				((StartContext)_localctx).r.add(((StartContext)_localctx).funciones);
				}
				}
				setState(60);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			_localctx.listainstrucciones = arrayList.New()
			                                                                       listaInst := localctx.(*StartContext).GetE() 
			                                                                                for _, e:= range listaInst {
			                                                                                        _localctx.listainstrucciones.Add(e.GetInst())
			                                                                                }
			                                                                                //fmt.Println("La lista es: --------",_localctx.listainstrucciones)
			                                                                        listaInst2 := localctx.(*StartContext).GetR() 
			                                                                                for _, r:= range listaInst2 {
			                                                                                        _localctx.listainstrucciones.Add(r.GetInst())
			                                                                                }
			                                                                                //fmt.Println("La lista es: --------",_localctx.listainstrucciones)
			                                                                        
			                                                                        //fmt.Println("\n\n\nLa lista es: ",((StartContext)_localctx).listtt.listainstrucciones)
			                                                                        listatemp :=((StartContext)_localctx).listtt.listainstrucciones
			                                                                                for _, s:= range listatemp.ToArray(){
			                                                                                        _localctx.listainstrucciones.Add(s)
			                                                                                }
			                                                                                fmt.Println("Fin del for")                                                                                
			                                                                                fmt.Println("Lista final: ", _localctx.listainstrucciones)
			                                                                        
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
		public DeclararParametrosContext declararParametros;
		public TerminalNode TKR_fn() { return getToken(gramaticaParser.TKR_fn, 0); }
		public TerminalNode TK_par_apertura() { return getToken(gramaticaParser.TK_par_apertura, 0); }
		public TerminalNode TK_par_cierre() { return getToken(gramaticaParser.TK_par_cierre, 0); }
		public TerminalNode TK_corchete_apertura() { return getToken(gramaticaParser.TK_corchete_apertura, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode TK_corchete_cierre() { return getToken(gramaticaParser.TK_corchete_cierre, 0); }
		public TerminalNode TK_id() { return getToken(gramaticaParser.TK_id, 0); }
		public DeclararParametrosContext declararParametros() {
			return getRuleContext(DeclararParametrosContext.class,0);
		}
		public FuncionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funciones; }
	}

	public final FuncionesContext funciones() throws RecognitionException {
		FuncionesContext _localctx = new FuncionesContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_funciones);
		try {
			setState(82);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(63);
				match(TKR_fn);
				setState(64);
				((FuncionesContext)_localctx).idd = match(TK_id);
				setState(65);
				match(TK_par_apertura);
				setState(66);
				match(TK_par_cierre);
				setState(67);
				match(TK_corchete_apertura);
				setState(68);
				((FuncionesContext)_localctx).instrucciones = instrucciones();
				setState(69);
				match(TK_corchete_cierre);
				_localctx.inst = Funcion.NewFuncion((((FuncionesContext)_localctx).idd!=null?((FuncionesContext)_localctx).idd.getText():null),Enum.STRING,nil,((FuncionesContext)_localctx).instrucciones.lista)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(72);
				match(TKR_fn);
				setState(73);
				((FuncionesContext)_localctx).idd = match(TK_id);
				setState(74);
				match(TK_par_apertura);
				setState(75);
				((FuncionesContext)_localctx).declararParametros = declararParametros(0);
				setState(76);
				match(TK_par_cierre);
				setState(77);
				match(TK_corchete_apertura);
				setState(78);
				((FuncionesContext)_localctx).instrucciones = instrucciones();
				setState(79);
				match(TK_corchete_cierre);
				_localctx.inst = Funcion.NewFuncion((((FuncionesContext)_localctx).idd!=null?((FuncionesContext)_localctx).idd.getText():null),Enum.STRING,((FuncionesContext)_localctx).declararParametros.lista,((FuncionesContext)_localctx).instrucciones.lista)
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

	public static class DeclararParametrosContext extends ParserRuleContext {
		public *arrayList.List lista;
		public DeclararParametrosContext list;
		public Token e1;
		public TipovariableContext e2;
		public TerminalNode TK_dosPuntos() { return getToken(gramaticaParser.TK_dosPuntos, 0); }
		public TerminalNode TK_id() { return getToken(gramaticaParser.TK_id, 0); }
		public TipovariableContext tipovariable() {
			return getRuleContext(TipovariableContext.class,0);
		}
		public TerminalNode TK_coma() { return getToken(gramaticaParser.TK_coma, 0); }
		public DeclararParametrosContext declararParametros() {
			return getRuleContext(DeclararParametrosContext.class,0);
		}
		public DeclararParametrosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declararParametros; }
	}

	public final DeclararParametrosContext declararParametros() throws RecognitionException {
		return declararParametros(0);
	}

	private DeclararParametrosContext declararParametros(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		DeclararParametrosContext _localctx = new DeclararParametrosContext(_ctx, _parentState);
		DeclararParametrosContext _prevctx = _localctx;
		int _startState = 4;
		enterRecursionRule(_localctx, 4, RULE_declararParametros, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(85);
			((DeclararParametrosContext)_localctx).e1 = match(TK_id);
			setState(86);
			match(TK_dosPuntos);
			setState(87);
			((DeclararParametrosContext)_localctx).e2 = tipovariable();
			_localctx.lista = arrayList.New()
			                                                                                                        _localctx.lista.Add(Funcion.NewParam((((DeclararParametrosContext)_localctx).e1!=null?((DeclararParametrosContext)_localctx).e1.getText():null),((DeclararParametrosContext)_localctx).e2.tipovar))
			}
			_ctx.stop = _input.LT(-1);
			setState(99);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new DeclararParametrosContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_declararParametros);
					setState(90);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(91);
					match(TK_coma);
					setState(92);
					((DeclararParametrosContext)_localctx).e1 = match(TK_id);
					setState(93);
					match(TK_dosPuntos);
					setState(94);
					((DeclararParametrosContext)_localctx).e2 = tipovariable();
					 ((DeclararParametrosContext)_localctx).list.lista.Add(Funcion.NewParam((((DeclararParametrosContext)_localctx).e1!=null?((DeclararParametrosContext)_localctx).e1.getText():null),((DeclararParametrosContext)_localctx).e2.tipovar))
					                                                                                                                  _localctx.lista = ((DeclararParametrosContext)_localctx).list.lista
					}
					} 
				}
				setState(101);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
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
		enterRule(_localctx, 6, RULE_main);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(102);
			match(TKR_fn);
			setState(103);
			match(TKR_main);
			setState(104);
			match(TK_par_apertura);
			setState(105);
			match(TK_par_cierre);
			setState(106);
			match(TK_corchete_apertura);
			setState(107);
			((MainContext)_localctx).instrucciones = instrucciones();
			setState(108);
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
		enterRule(_localctx, 8, RULE_instrucciones);

		        _localctx.lista = arrayList.New()

		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(114);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 43)) & ~0x3f) == 0 && ((1L << (_la - 43)) & ((1L << (TKR_let - 43)) | (1L << (TKR_println - 43)) | (1L << (TKR_if - 43)) | (1L << (TKR_while - 43)) | (1L << (TKR_break - 43)) | (1L << (TKR_loop - 43)) | (1L << (TK_id - 43)))) != 0)) {
				{
				{
				setState(111);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).e.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(116);
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
		public ImpresionContext impresion;
		public AsignacionContext asignacion;
		public DeclaracionContext declaracion;
		public CondicionalesContext condicionales;
		public BuclesContext bucles;
		public ControlContext control;
		public ImpresionContext impresion() {
			return getRuleContext(ImpresionContext.class,0);
		}
		public AsignacionContext asignacion() {
			return getRuleContext(AsignacionContext.class,0);
		}
		public DeclaracionContext declaracion() {
			return getRuleContext(DeclaracionContext.class,0);
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
		enterRule(_localctx, 10, RULE_instruccion);
		try {
			setState(137);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TKR_println:
				enterOuterAlt(_localctx, 1);
				{
				setState(119);
				((InstruccionContext)_localctx).impresion = impresion();
				_localctx.inst = ((InstruccionContext)_localctx).impresion.inst
				}
				break;
			case TK_id:
				enterOuterAlt(_localctx, 2);
				{
				setState(122);
				((InstruccionContext)_localctx).asignacion = asignacion();
				_localctx.inst = ((InstruccionContext)_localctx).asignacion.inst
				}
				break;
			case TKR_let:
				enterOuterAlt(_localctx, 3);
				{
				setState(125);
				((InstruccionContext)_localctx).declaracion = declaracion();
				_localctx.inst = ((InstruccionContext)_localctx).declaracion.inst
				}
				break;
			case TKR_if:
				enterOuterAlt(_localctx, 4);
				{
				setState(128);
				((InstruccionContext)_localctx).condicionales = condicionales();
				_localctx.inst = ((InstruccionContext)_localctx).condicionales.inst
				}
				break;
			case TKR_while:
			case TKR_loop:
				enterOuterAlt(_localctx, 5);
				{
				setState(131);
				((InstruccionContext)_localctx).bucles = bucles();
				_localctx.inst = ((InstruccionContext)_localctx).bucles.inst
				}
				break;
			case TKR_break:
				enterOuterAlt(_localctx, 6);
				{
				setState(134);
				((InstruccionContext)_localctx).control = control();
				_localctx.inst = ((InstruccionContext)_localctx).control.inst
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

	public static class AsignacionContext extends ParserRuleContext {
		public Interfaces.Instruccion inst;
		public Token TK_id;
		public Token e1;
		public ExpresionContext e2;
		public ArregloContext arreglo;
		public ExpresionContext ee;
		public TerminalNode TK_id() { return getToken(gramaticaParser.TK_id, 0); }
		public TerminalNode TK_par_apertura() { return getToken(gramaticaParser.TK_par_apertura, 0); }
		public TerminalNode TK_par_cierre() { return getToken(gramaticaParser.TK_par_cierre, 0); }
		public TerminalNode TK_pcoma() { return getToken(gramaticaParser.TK_pcoma, 0); }
		public TerminalNode TK_igual() { return getToken(gramaticaParser.TK_igual, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public ArregloContext arreglo() {
			return getRuleContext(ArregloContext.class,0);
		}
		public AsignacionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacion; }
	}

	public final AsignacionContext asignacion() throws RecognitionException {
		AsignacionContext _localctx = new AsignacionContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_asignacion);
		try {
			setState(156);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(139);
				((AsignacionContext)_localctx).TK_id = match(TK_id);
				setState(140);
				match(TK_par_apertura);
				setState(141);
				match(TK_par_cierre);
				setState(142);
				match(TK_pcoma);
				_localctx.inst = Funcion.NewLlamadafuncion((((AsignacionContext)_localctx).TK_id!=null?((AsignacionContext)_localctx).TK_id.getText():null),true,nil)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(144);
				((AsignacionContext)_localctx).e1 = match(TK_id);
				setState(145);
				match(TK_igual);
				setState(146);
				((AsignacionContext)_localctx).e2 = expresion(0);
				setState(147);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NewAsignacion((((AsignacionContext)_localctx).e1!=null?((AsignacionContext)_localctx).e1.getText():null),((AsignacionContext)_localctx).e2.exp)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(150);
				((AsignacionContext)_localctx).arreglo = arreglo(0);
				setState(151);
				match(TK_igual);
				setState(152);
				((AsignacionContext)_localctx).ee = expresion(0);
				setState(153);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NewAsignacionArray(((AsignacionContext)_localctx).arreglo.exp,((AsignacionContext)_localctx).ee.exp)
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
		enterRule(_localctx, 14, RULE_control);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(158);
			match(TKR_break);
			setState(159);
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
		public TerminalNode TKR_vec() { return getToken(gramaticaParser.TKR_vec, 0); }
		public TerminalNode TK_menor() { return getToken(gramaticaParser.TK_menor, 0); }
		public TerminalNode TK_mayor() { return getToken(gramaticaParser.TK_mayor, 0); }
		public DeclaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion; }
	}

	public final DeclaracionContext declaracion() throws RecognitionException {
		DeclaracionContext _localctx = new DeclaracionContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_declaracion);
		try {
			setState(248);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(162);
				match(TKR_let);
				setState(163);
				match(TKR_mut);
				setState(164);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(165);
				match(TK_dosPuntos);
				setState(166);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(167);
				match(TK_igual);
				setState(168);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(169);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar,((DeclaracionContext)_localctx).expresion.exp,false,false,true) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(172);
				match(TKR_let);
				setState(173);
				match(TKR_mut);
				setState(174);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(175);
				match(TK_igual);
				setState(176);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(177);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),Enum.SINTIPO,((DeclaracionContext)_localctx).expresion.exp,false,false,true) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(180);
				match(TKR_let);
				setState(181);
				match(TKR_mut);
				setState(182);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(183);
				match(TK_dosPuntos);
				setState(184);
				match(TK_llave_apertura);
				setState(185);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(186);
				match(TK_pcoma);
				setState(187);
				((DeclaracionContext)_localctx).cant = valores();
				setState(188);
				match(TK_llave_cierre);
				setState(189);
				match(TK_igual);
				setState(190);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(191);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracionArray((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar,((DeclaracionContext)_localctx).cant.exp,((DeclaracionContext)_localctx).expresion.exp,true)
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(194);
				match(TKR_let);
				setState(195);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(196);
				match(TK_dosPuntos);
				setState(197);
				match(TK_llave_apertura);
				setState(198);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(199);
				match(TK_pcoma);
				setState(200);
				((DeclaracionContext)_localctx).cant = valores();
				setState(201);
				match(TK_llave_cierre);
				setState(202);
				match(TK_igual);
				setState(203);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(204);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracionArray((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar,((DeclaracionContext)_localctx).cant.exp,((DeclaracionContext)_localctx).expresion.exp,false)
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(207);
				match(TKR_let);
				setState(208);
				match(TKR_mut);
				setState(209);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(210);
				match(TK_dosPuntos);
				setState(211);
				match(TKR_vec);
				setState(212);
				match(TK_menor);
				setState(213);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(214);
				match(TK_mayor);
				setState(215);
				match(TK_igual);
				setState(216);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(217);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracionVector((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar, ((DeclaracionContext)_localctx).expresion.exp,true) 
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(220);
				match(TKR_let);
				setState(221);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(222);
				match(TK_dosPuntos);
				setState(223);
				match(TKR_vec);
				setState(224);
				match(TK_menor);
				setState(225);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(226);
				match(TK_mayor);
				setState(227);
				match(TK_igual);
				setState(228);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(229);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracionVector((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar, ((DeclaracionContext)_localctx).expresion.exp,false) 
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(232);
				match(TKR_let);
				setState(233);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(234);
				match(TK_dosPuntos);
				setState(235);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(236);
				match(TK_igual);
				setState(237);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(238);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar,((DeclaracionContext)_localctx).expresion.exp,false,false,false) 
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(241);
				match(TKR_let);
				setState(242);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(243);
				match(TK_igual);
				setState(244);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(245);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),Enum.SINTIPO,((DeclaracionContext)_localctx).expresion.exp,false,false,false) 
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
		public Enum.Tipoexpresion tipovar;
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
		enterRule(_localctx, 18, RULE_tipovariable);
		try {
			setState(262);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TKR_numericos_enteros:
				enterOuterAlt(_localctx, 1);
				{
				setState(250);
				match(TKR_numericos_enteros);
				_localctx.tipovar = Enum.INTEGER
				}
				break;
			case TKR_numericos_flotantes:
				enterOuterAlt(_localctx, 2);
				{
				setState(252);
				match(TKR_numericos_flotantes);
				_localctx.tipovar = Enum.FLOAT
				}
				break;
			case TKR_String:
				enterOuterAlt(_localctx, 3);
				{
				setState(254);
				match(TKR_String);
				_localctx.tipovar = Enum.STRING
				}
				break;
			case TKR_bool:
				enterOuterAlt(_localctx, 4);
				{
				setState(256);
				match(TKR_bool);
				_localctx.tipovar = Enum.BOOLEAN
				}
				break;
			case TKR_char:
				enterOuterAlt(_localctx, 5);
				{
				setState(258);
				match(TKR_char);
				_localctx.tipovar = Enum.CHAR
				}
				break;
			case TKR_usize:
				enterOuterAlt(_localctx, 6);
				{
				setState(260);
				match(TKR_usize);
				_localctx.tipovar = Enum.USIZE
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
		public ArregloContext ar;
		public TerminalNode TK_entero() { return getToken(gramaticaParser.TK_entero, 0); }
		public TerminalNode TK_decimal() { return getToken(gramaticaParser.TK_decimal, 0); }
		public TerminalNode TK_cadena() { return getToken(gramaticaParser.TK_cadena, 0); }
		public TerminalNode TK_caracter() { return getToken(gramaticaParser.TK_caracter, 0); }
		public TerminalNode TK_amp() { return getToken(gramaticaParser.TK_amp, 0); }
		public TerminalNode TKR_Str() { return getToken(gramaticaParser.TKR_Str, 0); }
		public TerminalNode TKR_true() { return getToken(gramaticaParser.TKR_true, 0); }
		public TerminalNode TKR_false() { return getToken(gramaticaParser.TKR_false, 0); }
		public ArregloContext arreglo() {
			return getRuleContext(ArregloContext.class,0);
		}
		public ValoresContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_valores; }
	}

	public final ValoresContext valores() throws RecognitionException {
		ValoresContext _localctx = new ValoresContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_valores);
		try {
			setState(282);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_entero:
				enterOuterAlt(_localctx, 1);
				{
				setState(264);
				((ValoresContext)_localctx).TK_entero = match(TK_entero);

				                numero, err := strconv.Atoi((((ValoresContext)_localctx).TK_entero!=null?((ValoresContext)_localctx).TK_entero.getText():null))
				                if err != nil {
				                        fmt.Println(err)
				                }
				                _localctx.exp = Expresion.NuevoPrimitivo(numero, Enum.INTEGER)
				}
				break;
			case TK_decimal:
				enterOuterAlt(_localctx, 2);
				{
				setState(266);
				((ValoresContext)_localctx).TK_decimal = match(TK_decimal);
				decimal, err := strconv.ParseFloat((((ValoresContext)_localctx).TK_decimal!=null?((ValoresContext)_localctx).TK_decimal.getText():null),8)
				                if err != nil {
				                        fmt.Println(err)
				                }
				                _localctx.exp = Expresion.NuevoPrimitivo(decimal, Enum.FLOAT)
				}
				break;
			case TK_cadena:
				enterOuterAlt(_localctx, 3);
				{
				setState(268);
				((ValoresContext)_localctx).TK_cadena = match(TK_cadena);

				                str:= (((ValoresContext)_localctx).TK_cadena!=null?((ValoresContext)_localctx).TK_cadena.getText():null)[1:len((((ValoresContext)_localctx).TK_cadena!=null?((ValoresContext)_localctx).TK_cadena.getText():null))-1]
				                _localctx.exp = Expresion.NuevoPrimitivo(str,Enum.STRING)
				}
				break;
			case TK_caracter:
				enterOuterAlt(_localctx, 4);
				{
				setState(270);
				((ValoresContext)_localctx).TK_caracter = match(TK_caracter);
				str:= (((ValoresContext)_localctx).TK_caracter!=null?((ValoresContext)_localctx).TK_caracter.getText():null)[1:len((((ValoresContext)_localctx).TK_caracter!=null?((ValoresContext)_localctx).TK_caracter.getText():null))-1]
				                _localctx.exp = Expresion.NuevoPrimitivo(str,Enum.CHAR)
				}
				break;
			case TK_amp:
				enterOuterAlt(_localctx, 5);
				{
				setState(272);
				match(TK_amp);
				setState(273);
				((ValoresContext)_localctx).TKR_Str = match(TKR_Str);
				str:= (((ValoresContext)_localctx).TKR_Str!=null?((ValoresContext)_localctx).TKR_Str.getText():null)[1:len((((ValoresContext)_localctx).TKR_Str!=null?((ValoresContext)_localctx).TKR_Str.getText():null))-1]
				                _localctx.exp = Expresion.NuevoPrimitivo(str,Enum.STR)
				}
				break;
			case TKR_true:
				enterOuterAlt(_localctx, 6);
				{
				setState(275);
				match(TKR_true);
				_localctx.exp = Expresion.NuevoPrimitivo(true, Enum.BOOLEAN)
				}
				break;
			case TKR_false:
				enterOuterAlt(_localctx, 7);
				{
				setState(277);
				match(TKR_false);
				_localctx.exp = Expresion.NuevoPrimitivo(false,Enum.BOOLEAN)
				}
				break;
			case TK_id:
				enterOuterAlt(_localctx, 8);
				{
				setState(279);
				((ValoresContext)_localctx).ar = arreglo(0);
				_localctx.exp = ((ValoresContext)_localctx).ar.exp
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

	public static class ArregloContext extends ParserRuleContext {
		public Interfaces.Expresion exp;
		public ArregloContext ar;
		public Token TK_id;
		public ExpresionContext expresion;
		public TerminalNode TK_id() { return getToken(gramaticaParser.TK_id, 0); }
		public TerminalNode TK_llave_apertura() { return getToken(gramaticaParser.TK_llave_apertura, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode TK_llave_cierre() { return getToken(gramaticaParser.TK_llave_cierre, 0); }
		public ArregloContext arreglo() {
			return getRuleContext(ArregloContext.class,0);
		}
		public ArregloContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arreglo; }
	}

	public final ArregloContext arreglo() throws RecognitionException {
		return arreglo(0);
	}

	private ArregloContext arreglo(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ArregloContext _localctx = new ArregloContext(_ctx, _parentState);
		ArregloContext _prevctx = _localctx;
		int _startState = 22;
		enterRecursionRule(_localctx, 22, RULE_arreglo, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(285);
			((ArregloContext)_localctx).TK_id = match(TK_id);
			_localctx.exp = Expresiones.NewLlamarvariable((((ArregloContext)_localctx).TK_id!=null?((ArregloContext)_localctx).TK_id.getText():null))
			}
			_ctx.stop = _input.LT(-1);
			setState(296);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ArregloContext(_parentctx, _parentState);
					_localctx.ar = _prevctx;
					_localctx.ar = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_arreglo);
					setState(288);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(289);
					match(TK_llave_apertura);
					setState(290);
					((ArregloContext)_localctx).expresion = expresion(0);
					setState(291);
					match(TK_llave_cierre);
					_localctx.exp = Expresiones.NewAccesoArray(((ArregloContext)_localctx).ar.exp,((ArregloContext)_localctx).expresion.exp)
					}
					} 
				}
				setState(298);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
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
		public TerminalNode TKR_vec() { return getToken(gramaticaParser.TKR_vec, 0); }
		public List<TerminalNode> TK_dosPuntos() { return getTokens(gramaticaParser.TK_dosPuntos); }
		public TerminalNode TK_dosPuntos(int i) {
			return getToken(gramaticaParser.TK_dosPuntos, i);
		}
		public TerminalNode TKR_new() { return getToken(gramaticaParser.TKR_new, 0); }
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
		int _startState = 24;
		enterRecursionRule(_localctx, 24, RULE_expresion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(379);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				{
				setState(300);
				match(TK_menos);
				setState(301);
				((ExpresionContext)_localctx).e1 = expresion(27);
				numero := -1
				                                                                                                        numm := Expresion.NuevoPrimitivo(numero, Enum.INTEGER)
				                                                                                                                _localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp, numm ,Enum.MULTIPLICACION)
				                                                                                                
				}
				break;
			case 2:
				{
				setState(304);
				match(TKR_pow);
				setState(305);
				match(TK_par_apertura);
				setState(306);
				((ExpresionContext)_localctx).e1 = expresion(0);
				setState(307);
				match(TK_coma);
				setState(308);
				((ExpresionContext)_localctx).e2 = expresion(0);
				setState(309);
				match(TK_par_cierre);
				_localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.POW)
				}
				break;
			case 3:
				{
				setState(312);
				match(TKR_powf);
				setState(313);
				match(TK_par_apertura);
				setState(314);
				((ExpresionContext)_localctx).e1 = expresion(0);
				setState(315);
				match(TK_coma);
				setState(316);
				((ExpresionContext)_localctx).e2 = expresion(0);
				setState(317);
				match(TK_par_cierre);
				_localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.POWF)
				}
				break;
			case 4:
				{
				setState(320);
				match(TK_sig_admiracion);
				setState(321);
				((ExpresionContext)_localctx).e1 = expresion(14);
				_localctx.exp = Expresiones.NuevaLogica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e1.exp,Enum.NOT)
				}
				break;
			case 5:
				{
				setState(324);
				match(TK_par_apertura);
				setState(325);
				((ExpresionContext)_localctx).va = expresion(0);
				setState(326);
				match(TK_par_cierre);
				_localctx.exp = ((ExpresionContext)_localctx).va.exp
				}
				break;
			case 6:
				{
				setState(329);
				((ExpresionContext)_localctx).val = valores();
				setState(330);
				match(TKR_as);
				setState(331);
				match(TKR_numericos_enteros);
				_localctx.exp = Expresiones.NewAsi64(((ExpresionContext)_localctx).val.exp)
				}
				break;
			case 7:
				{
				setState(334);
				((ExpresionContext)_localctx).val = valores();
				setState(335);
				match(TKR_as);
				setState(336);
				match(TKR_numericos_flotantes);
				_localctx.exp = Expresiones.NewAsf64(((ExpresionContext)_localctx).val.exp)
				}
				break;
			case 8:
				{
				setState(339);
				match(TK_llave_apertura);
				setState(340);
				((ExpresionContext)_localctx).e1 = expresion(0);
				setState(341);
				match(TK_pcoma);
				setState(342);
				((ExpresionContext)_localctx).e2 = expresion(0);
				setState(343);
				match(TK_llave_cierre);
				_localctx.exp = Expresion.NewArray(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,nil,Enum.MULTIPLE)
				}
				break;
			case 9:
				{
				setState(346);
				match(TK_llave_apertura);
				setState(347);
				((ExpresionContext)_localctx).e1 = expresion(0);
				setState(348);
				((ExpresionContext)_localctx).l1 = impresionexpresion();
				setState(349);
				match(TK_llave_cierre);
				_localctx.exp = Expresion.NewArray(((ExpresionContext)_localctx).e1.exp,nil,((ExpresionContext)_localctx).l1.lista,Enum.NORMAL)
				}
				break;
			case 10:
				{
				setState(352);
				match(TKR_vec);
				setState(353);
				match(TK_sig_admiracion);
				setState(354);
				match(TK_llave_apertura);
				setState(355);
				((ExpresionContext)_localctx).e1 = expresion(0);
				setState(356);
				match(TK_pcoma);
				setState(357);
				((ExpresionContext)_localctx).e2 = expresion(0);
				setState(358);
				match(TK_llave_cierre);
				_localctx.exp = Expresion.NewVector(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,nil,Enum.MULTIPLE)
				}
				break;
			case 11:
				{
				setState(361);
				match(TKR_vec);
				setState(362);
				match(TK_sig_admiracion);
				setState(363);
				match(TK_llave_apertura);
				setState(364);
				((ExpresionContext)_localctx).e1 = expresion(0);
				setState(365);
				((ExpresionContext)_localctx).l1 = impresionexpresion();
				setState(366);
				match(TK_llave_cierre);
				_localctx.exp = Expresion.NewVector(((ExpresionContext)_localctx).e1.exp,nil,((ExpresionContext)_localctx).l1.lista,Enum.NORMAL)
				}
				break;
			case 12:
				{
				setState(369);
				match(TKR_vec);
				setState(370);
				match(TK_dosPuntos);
				setState(371);
				match(TK_dosPuntos);
				setState(372);
				match(TKR_new);
				setState(373);
				match(TK_par_apertura);
				setState(374);
				match(TK_par_cierre);
				_localctx.exp = Expresion.NewVector(((ExpresionContext)_localctx).e1.exp,nil,((ExpresionContext)_localctx).l1.lista,Enum.NORMAL)
				}
				break;
			case 13:
				{
				setState(376);
				((ExpresionContext)_localctx).vall = valores();
				_localctx.exp = ((ExpresionContext)_localctx).vall.exp
				                                                                                        fmt.Println(_localctx.exp)
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(456);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(454);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
					case 1:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(381);
						if (!(precpred(_ctx, 26))) throw new FailedPredicateException(this, "precpred(_ctx, 26)");
						setState(382);
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
						setState(383);
						((ExpresionContext)_localctx).e2 = expresion(27);

						                                                                                                                  if (((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null) == "*" {
						                                                                                                                          _localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.MULTIPLICACION)
						                                                                                                                  }else if (((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null) == "%" {
						                                                                                                                          _localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.MODULO)
						                                                                                                                  }else{
						                                                                                                                          _localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.DIVISION)
						                                                                                                                  }
						                                                                                                          
						}
						break;
					case 2:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(386);
						if (!(precpred(_ctx, 25))) throw new FailedPredicateException(this, "precpred(_ctx, 25)");
						setState(387);
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
						setState(388);
						((ExpresionContext)_localctx).e2 = expresion(26);

						                                                                                                                  if (((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null) == "+" {
						                                                                                                                          _localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.SUMA)
						                                                                                                                  }else{
						                                                                                                                          _localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.RESTA)
						                                                                                                                  }
						                                                                                                          
						}
						break;
					case 3:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(391);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(392);
						match(TK_menor);
						setState(393);
						((ExpresionContext)_localctx).e2 = expresion(23);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.MENOR_QUE)
						}
						break;
					case 4:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(396);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(397);
						match(TK_mayor);
						setState(398);
						((ExpresionContext)_localctx).e2 = expresion(22);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.MAYOR_QUE)
						}
						break;
					case 5:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(401);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(402);
						match(TK_mayor_igual);
						setState(403);
						((ExpresionContext)_localctx).e2 = expresion(21);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.MAYOR_IGUAL)
						}
						break;
					case 6:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(406);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(407);
						match(TK_menor_igual);
						setState(408);
						((ExpresionContext)_localctx).e2 = expresion(20);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.MENOR_IGUAL)
						}
						break;
					case 7:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(411);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(412);
						match(TK_igualacion);
						setState(413);
						((ExpresionContext)_localctx).e2 = expresion(19);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.IGUALDAD)
						}
						break;
					case 8:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(416);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(417);
						match(TK_diferente);
						setState(418);
						((ExpresionContext)_localctx).e2 = expresion(18);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.DESIGUAL)
						}
						break;
					case 9:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(421);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(422);
						match(TK_or);
						setState(423);
						((ExpresionContext)_localctx).e2 = expresion(17);
						_localctx.exp = Expresiones.NuevaLogica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.OR)
						}
						break;
					case 10:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(426);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(427);
						match(TK_and);
						setState(428);
						((ExpresionContext)_localctx).e2 = expresion(16);
						_localctx.exp = Expresiones.NuevaLogica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.AND)
						}
						break;
					case 11:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.vll = _prevctx;
						_localctx.vll = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(431);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(432);
						match(TK_punto);
						setState(433);
						match(TKR_abs);
						setState(434);
						match(TK_par_apertura);
						setState(435);
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
						setState(437);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(438);
						match(TK_punto);
						setState(439);
						match(TKR_sqrt);
						setState(440);
						match(TK_par_apertura);
						setState(441);
						match(TK_par_cierre);
						_localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e1.exp,Enum.MULTIPLICACION)
						}
						break;
					case 13:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.vll = _prevctx;
						_localctx.vll = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(443);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(444);
						match(TK_punto);
						setState(445);
						match(TKR_to_string);
						setState(446);
						match(TK_par_apertura);
						setState(447);
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
						setState(449);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(450);
						match(TK_punto);
						setState(451);
						match(TKR_clone);
						setState(452);
						match(TK_par_apertura);
						setState(453);
						match(TK_par_cierre);
						}
						break;
					}
					} 
				}
				setState(458);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
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
		enterRule(_localctx, 26, RULE_impresion);
		try {
			setState(474);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(459);
				match(TKR_println);
				setState(460);
				match(TK_par_apertura);
				setState(461);
				((ImpresionContext)_localctx).e1 = expresion(0);
				setState(462);
				match(TK_par_cierre);
				setState(463);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevoPrint(((ImpresionContext)_localctx).e1.exp,nil)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(466);
				match(TKR_println);
				setState(467);
				match(TK_par_apertura);
				setState(468);
				((ImpresionContext)_localctx).e2 = expresion(0);
				setState(469);
				((ImpresionContext)_localctx).li = impresionexpresion();
				setState(470);
				match(TK_par_cierre);
				setState(471);
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
		enterRule(_localctx, 28, RULE_impresionexpresion);
		_localctx.lista = arrayList.New()
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(479);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==TK_coma) {
				{
				{
				setState(476);
				((ImpresionexpresionContext)_localctx).expcoma = expcoma();
				((ImpresionexpresionContext)_localctx).list.add(((ImpresionexpresionContext)_localctx).expcoma);
				}
				}
				setState(481);
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
		enterRule(_localctx, 30, RULE_expcoma);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(484);
			match(TK_coma);
			setState(485);
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
		enterRule(_localctx, 32, RULE_condicionales);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(488);
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
		enterRule(_localctx, 34, RULE_funcionif);
		try {
			setState(517);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(491);
				match(TKR_if);
				setState(492);
				((FuncionifContext)_localctx).e1 = expresion(0);
				setState(493);
				((FuncionifContext)_localctx).ee = bloque();
				_localctx.inst = Instruccion.NewIf(((FuncionifContext)_localctx).e1.exp,((FuncionifContext)_localctx).ee.lista,nil,nil )
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(496);
				match(TKR_if);
				setState(497);
				((FuncionifContext)_localctx).e1 = expresion(0);
				setState(498);
				((FuncionifContext)_localctx).e5 = bloque();
				setState(499);
				match(TKR_else);
				setState(500);
				((FuncionifContext)_localctx).b2 = bloque();
				_localctx.inst = Instruccion.NewIf(((FuncionifContext)_localctx).e1.exp,((FuncionifContext)_localctx).e5.lista,nil,((FuncionifContext)_localctx).b2.lista)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(503);
				match(TKR_if);
				setState(504);
				((FuncionifContext)_localctx).e1 = expresion(0);
				setState(505);
				((FuncionifContext)_localctx).b1 = bloque();
				setState(506);
				((FuncionifContext)_localctx).listaelseif = listaelseif();
				setState(507);
				match(TKR_else);
				setState(508);
				((FuncionifContext)_localctx).b2 = bloque();
				_localctx.inst = Instruccion.NewIf(((FuncionifContext)_localctx).e1.exp,((FuncionifContext)_localctx).b1.lista,((FuncionifContext)_localctx).listaelseif.lista,((FuncionifContext)_localctx).b2.lista)
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(511);
				match(TKR_if);
				setState(512);
				((FuncionifContext)_localctx).e1 = expresion(0);
				setState(513);
				((FuncionifContext)_localctx).b1 = bloque();
				setState(514);
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
		enterRule(_localctx, 36, RULE_funcionelseif);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(519);
			match(TKR_elseif);
			setState(520);
			((FuncionelseifContext)_localctx).e1 = expresion(0);
			setState(521);
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
		enterRule(_localctx, 38, RULE_listaelseif);
		 _localctx.lista = arrayList.New()
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(527);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==TKR_elseif) {
				{
				{
				setState(524);
				((ListaelseifContext)_localctx).funcionelseif = funcionelseif();
				((ListaelseifContext)_localctx).list.add(((ListaelseifContext)_localctx).funcionelseif);
				}
				}
				setState(529);
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
		enterRule(_localctx, 40, RULE_bloque);
		try {
			setState(540);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(532);
				match(TK_corchete_apertura);
				setState(533);
				((BloqueContext)_localctx).instrucciones = instrucciones();
				setState(534);
				match(TK_corchete_cierre);
				_localctx.lista = ((BloqueContext)_localctx).instrucciones.lista
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(537);
				match(TK_corchete_apertura);
				setState(538);
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
		enterRule(_localctx, 42, RULE_bucles);
		try {
			setState(548);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TKR_while:
				enterOuterAlt(_localctx, 1);
				{
				setState(542);
				((BuclesContext)_localctx).fwhile = fwhile();
				_localctx.inst = ((BuclesContext)_localctx).fwhile.inst
				}
				break;
			case TKR_loop:
				enterOuterAlt(_localctx, 2);
				{
				setState(545);
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
		enterRule(_localctx, 44, RULE_fwhile);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(550);
			match(TKR_while);
			setState(551);
			((FwhileContext)_localctx).e1 = expresion(0);
			setState(552);
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
		enterRule(_localctx, 46, RULE_floop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(555);
			match(TKR_loop);
			setState(556);
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
		case 2:
			return declararParametros_sempred((DeclararParametrosContext)_localctx, predIndex);
		case 11:
			return arreglo_sempred((ArregloContext)_localctx, predIndex);
		case 12:
			return expresion_sempred((ExpresionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean declararParametros_sempred(DeclararParametrosContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean arreglo_sempred(ArregloContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expresion_sempred(ExpresionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 26);
		case 3:
			return precpred(_ctx, 25);
		case 4:
			return precpred(_ctx, 22);
		case 5:
			return precpred(_ctx, 21);
		case 6:
			return precpred(_ctx, 20);
		case 7:
			return precpred(_ctx, 19);
		case 8:
			return precpred(_ctx, 18);
		case 9:
			return precpred(_ctx, 17);
		case 10:
			return precpred(_ctx, 16);
		case 11:
			return precpred(_ctx, 15);
		case 12:
			return precpred(_ctx, 10);
		case 13:
			return precpred(_ctx, 9);
		case 14:
			return precpred(_ctx, 8);
		case 15:
			return precpred(_ctx, 7);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3N\u0232\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\3\2\7\2\64\n\2\f\2\16\2\67\13\2\3\2\3\2\7\2;\n\2\f\2\16\2>\13\2\3\2\3"+
		"\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\3\5\3U\n\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\7\4d\n\4\f\4\16\4g\13\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6\7\6s"+
		"\n\6\f\6\16\6v\13\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3"+
		"\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7\u008c\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\5\b\u009f\n\b\3\t\3\t\3\t\3"+
		"\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n"+
		"\u00fb\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\5\13\u0109\n\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\f\5\f\u011d\n\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3"+
		"\r\7\r\u0129\n\r\f\r\16\r\u012c\13\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\5\16\u017e\n\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\7\16\u01c9\n\16\f\16\16"+
		"\16\u01cc\13\16\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\5\17\u01dd\n\17\3\20\7\20\u01e0\n\20\f\20\16\20\u01e3"+
		"\13\20\3\20\3\20\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\23\3\23\3\23\3\23"+
		"\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23"+
		"\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\5\23\u0208\n\23\3\24\3\24\3\24"+
		"\3\24\3\24\3\25\7\25\u0210\n\25\f\25\16\25\u0213\13\25\3\25\3\25\3\26"+
		"\3\26\3\26\3\26\3\26\3\26\3\26\3\26\5\26\u021f\n\26\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\5\27\u0227\n\27\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31"+
		"\3\31\2\5\6\30\32\32\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60"+
		"\2\4\3\2\31\33\3\2\27\30\2\u025b\2\65\3\2\2\2\4T\3\2\2\2\6V\3\2\2\2\b"+
		"h\3\2\2\2\nt\3\2\2\2\f\u008b\3\2\2\2\16\u009e\3\2\2\2\20\u00a0\3\2\2\2"+
		"\22\u00fa\3\2\2\2\24\u0108\3\2\2\2\26\u011c\3\2\2\2\30\u011e\3\2\2\2\32"+
		"\u017d\3\2\2\2\34\u01dc\3\2\2\2\36\u01e1\3\2\2\2 \u01e6\3\2\2\2\"\u01ea"+
		"\3\2\2\2$\u0207\3\2\2\2&\u0209\3\2\2\2(\u0211\3\2\2\2*\u021e\3\2\2\2,"+
		"\u0226\3\2\2\2.\u0228\3\2\2\2\60\u022d\3\2\2\2\62\64\5\4\3\2\63\62\3\2"+
		"\2\2\64\67\3\2\2\2\65\63\3\2\2\2\65\66\3\2\2\2\668\3\2\2\2\67\65\3\2\2"+
		"\28<\5\b\5\29;\5\4\3\2:9\3\2\2\2;>\3\2\2\2<:\3\2\2\2<=\3\2\2\2=?\3\2\2"+
		"\2><\3\2\2\2?@\b\2\1\2@\3\3\2\2\2AB\7\64\2\2BC\7K\2\2CD\7\24\2\2DE\7\25"+
		"\2\2EF\7\n\2\2FG\5\n\6\2GH\7\13\2\2HI\b\3\1\2IU\3\2\2\2JK\7\64\2\2KL\7"+
		"K\2\2LM\7\24\2\2MN\5\6\4\2NO\7\25\2\2OP\7\n\2\2PQ\5\n\6\2QR\7\13\2\2R"+
		"S\b\3\1\2SU\3\2\2\2TA\3\2\2\2TJ\3\2\2\2U\5\3\2\2\2VW\b\4\1\2WX\7K\2\2"+
		"XY\7\16\2\2YZ\5\24\13\2Z[\b\4\1\2[e\3\2\2\2\\]\f\4\2\2]^\7\17\2\2^_\7"+
		"K\2\2_`\7\16\2\2`a\5\24\13\2ab\b\4\1\2bd\3\2\2\2c\\\3\2\2\2dg\3\2\2\2"+
		"ec\3\2\2\2ef\3\2\2\2f\7\3\2\2\2ge\3\2\2\2hi\7\64\2\2ij\7B\2\2jk\7\24\2"+
		"\2kl\7\25\2\2lm\7\n\2\2mn\5\n\6\2no\7\13\2\2op\b\5\1\2p\t\3\2\2\2qs\5"+
		"\f\7\2rq\3\2\2\2sv\3\2\2\2tr\3\2\2\2tu\3\2\2\2uw\3\2\2\2vt\3\2\2\2wx\b"+
		"\6\1\2x\13\3\2\2\2yz\5\34\17\2z{\b\7\1\2{\u008c\3\2\2\2|}\5\16\b\2}~\b"+
		"\7\1\2~\u008c\3\2\2\2\177\u0080\5\22\n\2\u0080\u0081\b\7\1\2\u0081\u008c"+
		"\3\2\2\2\u0082\u0083\5\"\22\2\u0083\u0084\b\7\1\2\u0084\u008c\3\2\2\2"+
		"\u0085\u0086\5,\27\2\u0086\u0087\b\7\1\2\u0087\u008c\3\2\2\2\u0088\u0089"+
		"\5\20\t\2\u0089\u008a\b\7\1\2\u008a\u008c\3\2\2\2\u008by\3\2\2\2\u008b"+
		"|\3\2\2\2\u008b\177\3\2\2\2\u008b\u0082\3\2\2\2\u008b\u0085\3\2\2\2\u008b"+
		"\u0088\3\2\2\2\u008c\r\3\2\2\2\u008d\u008e\7K\2\2\u008e\u008f\7\24\2\2"+
		"\u008f\u0090\7\25\2\2\u0090\u0091\7\20\2\2\u0091\u009f\b\b\1\2\u0092\u0093"+
		"\7K\2\2\u0093\u0094\7\26\2\2\u0094\u0095\5\32\16\2\u0095\u0096\7\20\2"+
		"\2\u0096\u0097\b\b\1\2\u0097\u009f\3\2\2\2\u0098\u0099\5\30\r\2\u0099"+
		"\u009a\7\26\2\2\u009a\u009b\5\32\16\2\u009b\u009c\7\20\2\2\u009c\u009d"+
		"\b\b\1\2\u009d\u009f\3\2\2\2\u009e\u008d\3\2\2\2\u009e\u0092\3\2\2\2\u009e"+
		"\u0098\3\2\2\2\u009f\17\3\2\2\2\u00a0\u00a1\7G\2\2\u00a1\u00a2\7\20\2"+
		"\2\u00a2\u00a3\b\t\1\2\u00a3\21\3\2\2\2\u00a4\u00a5\7-\2\2\u00a5\u00a6"+
		"\7.\2\2\u00a6\u00a7\7K\2\2\u00a7\u00a8\7\16\2\2\u00a8\u00a9\5\24\13\2"+
		"\u00a9\u00aa\7\26\2\2\u00aa\u00ab\5\32\16\2\u00ab\u00ac\7\20\2\2\u00ac"+
		"\u00ad\b\n\1\2\u00ad\u00fb\3\2\2\2\u00ae\u00af\7-\2\2\u00af\u00b0\7.\2"+
		"\2\u00b0\u00b1\7K\2\2\u00b1\u00b2\7\26\2\2\u00b2\u00b3\5\32\16\2\u00b3"+
		"\u00b4\7\20\2\2\u00b4\u00b5\b\n\1\2\u00b5\u00fb\3\2\2\2\u00b6\u00b7\7"+
		"-\2\2\u00b7\u00b8\7.\2\2\u00b8\u00b9\7K\2\2\u00b9\u00ba\7\16\2\2\u00ba"+
		"\u00bb\7\f\2\2\u00bb\u00bc\5\24\13\2\u00bc\u00bd\7\20\2\2\u00bd\u00be"+
		"\5\26\f\2\u00be\u00bf\7\r\2\2\u00bf\u00c0\7\26\2\2\u00c0\u00c1\5\32\16"+
		"\2\u00c1\u00c2\7\20\2\2\u00c2\u00c3\b\n\1\2\u00c3\u00fb\3\2\2\2\u00c4"+
		"\u00c5\7-\2\2\u00c5\u00c6\7K\2\2\u00c6\u00c7\7\16\2\2\u00c7\u00c8\7\f"+
		"\2\2\u00c8\u00c9\5\24\13\2\u00c9\u00ca\7\20\2\2\u00ca\u00cb\5\26\f\2\u00cb"+
		"\u00cc\7\r\2\2\u00cc\u00cd\7\26\2\2\u00cd\u00ce\5\32\16\2\u00ce\u00cf"+
		"\7\20\2\2\u00cf\u00d0\b\n\1\2\u00d0\u00fb\3\2\2\2\u00d1\u00d2\7-\2\2\u00d2"+
		"\u00d3\7.\2\2\u00d3\u00d4\7K\2\2\u00d4\u00d5\7\16\2\2\u00d5\u00d6\7&\2"+
		"\2\u00d6\u00d7\7\21\2\2\u00d7\u00d8\5\24\13\2\u00d8\u00d9\7\22\2\2\u00d9"+
		"\u00da\7\26\2\2\u00da\u00db\5\32\16\2\u00db\u00dc\7\20\2\2\u00dc\u00dd"+
		"\b\n\1\2\u00dd\u00fb\3\2\2\2\u00de\u00df\7-\2\2\u00df\u00e0\7K\2\2\u00e0"+
		"\u00e1\7\16\2\2\u00e1\u00e2\7&\2\2\u00e2\u00e3\7\21\2\2\u00e3\u00e4\5"+
		"\24\13\2\u00e4\u00e5\7\22\2\2\u00e5\u00e6\7\26\2\2\u00e6\u00e7\5\32\16"+
		"\2\u00e7\u00e8\7\20\2\2\u00e8\u00e9\b\n\1\2\u00e9\u00fb\3\2\2\2\u00ea"+
		"\u00eb\7-\2\2\u00eb\u00ec\7K\2\2\u00ec\u00ed\7\16\2\2\u00ed\u00ee\5\24"+
		"\13\2\u00ee\u00ef\7\26\2\2\u00ef\u00f0\5\32\16\2\u00f0\u00f1\7\20\2\2"+
		"\u00f1\u00f2\b\n\1\2\u00f2\u00fb\3\2\2\2\u00f3\u00f4\7-\2\2\u00f4\u00f5"+
		"\7K\2\2\u00f5\u00f6\7\26\2\2\u00f6\u00f7\5\32\16\2\u00f7\u00f8\7\20\2"+
		"\2\u00f8\u00f9\b\n\1\2\u00f9\u00fb\3\2\2\2\u00fa\u00a4\3\2\2\2\u00fa\u00ae"+
		"\3\2\2\2\u00fa\u00b6\3\2\2\2\u00fa\u00c4\3\2\2\2\u00fa\u00d1\3\2\2\2\u00fa"+
		"\u00de\3\2\2\2\u00fa\u00ea\3\2\2\2\u00fa\u00f3\3\2\2\2\u00fb\23\3\2\2"+
		"\2\u00fc\u00fd\7#\2\2\u00fd\u0109\b\13\1\2\u00fe\u00ff\7$\2\2\u00ff\u0109"+
		"\b\13\1\2\u0100\u0101\7+\2\2\u0101\u0109\b\13\1\2\u0102\u0103\7(\2\2\u0103"+
		"\u0109\b\13\1\2\u0104\u0105\7)\2\2\u0105\u0109\b\13\1\2\u0106\u0107\7"+
		",\2\2\u0107\u0109\b\13\1\2\u0108\u00fc\3\2\2\2\u0108\u00fe\3\2\2\2\u0108"+
		"\u0100\3\2\2\2\u0108\u0102\3\2\2\2\u0108\u0104\3\2\2\2\u0108\u0106\3\2"+
		"\2\2\u0109\25\3\2\2\2\u010a\u010b\7I\2\2\u010b\u011d\b\f\1\2\u010c\u010d"+
		"\7J\2\2\u010d\u011d\b\f\1\2\u010e\u010f\7M\2\2\u010f\u011d\b\f\1\2\u0110"+
		"\u0111\7N\2\2\u0111\u011d\b\f\1\2\u0112\u0113\7\35\2\2\u0113\u0114\7*"+
		"\2\2\u0114\u011d\b\f\1\2\u0115\u0116\7\62\2\2\u0116\u011d\b\f\1\2\u0117"+
		"\u0118\7\63\2\2\u0118\u011d\b\f\1\2\u0119\u011a\5\30\r\2\u011a\u011b\b"+
		"\f\1\2\u011b\u011d\3\2\2\2\u011c\u010a\3\2\2\2\u011c\u010c\3\2\2\2\u011c"+
		"\u010e\3\2\2\2\u011c\u0110\3\2\2\2\u011c\u0112\3\2\2\2\u011c\u0115\3\2"+
		"\2\2\u011c\u0117\3\2\2\2\u011c\u0119\3\2\2\2\u011d\27\3\2\2\2\u011e\u011f"+
		"\b\r\1\2\u011f\u0120\7K\2\2\u0120\u0121\b\r\1\2\u0121\u012a\3\2\2\2\u0122"+
		"\u0123\f\4\2\2\u0123\u0124\7\f\2\2\u0124\u0125\5\32\16\2\u0125\u0126\7"+
		"\r\2\2\u0126\u0127\b\r\1\2\u0127\u0129\3\2\2\2\u0128\u0122\3\2\2\2\u0129"+
		"\u012c\3\2\2\2\u012a\u0128\3\2\2\2\u012a\u012b\3\2\2\2\u012b\31\3\2\2"+
		"\2\u012c\u012a\3\2\2\2\u012d\u012e\b\16\1\2\u012e\u012f\7\30\2\2\u012f"+
		"\u0130\5\32\16\35\u0130\u0131\b\16\1\2\u0131\u017e\3\2\2\2\u0132\u0133"+
		"\7%\2\2\u0133\u0134\7\24\2\2\u0134\u0135\5\32\16\2\u0135\u0136\7\17\2"+
		"\2\u0136\u0137\5\32\16\2\u0137\u0138\7\25\2\2\u0138\u0139\b\16\1\2\u0139"+
		"\u017e\3\2\2\2\u013a\u013b\7\'\2\2\u013b\u013c\7\24\2\2\u013c\u013d\5"+
		"\32\16\2\u013d\u013e\7\17\2\2\u013e\u013f\5\32\16\2\u013f\u0140\7\25\2"+
		"\2\u0140\u0141\b\16\1\2\u0141\u017e\3\2\2\2\u0142\u0143\7\36\2\2\u0143"+
		"\u0144\5\32\16\20\u0144\u0145\b\16\1\2\u0145\u017e\3\2\2\2\u0146\u0147"+
		"\7\24\2\2\u0147\u0148\5\32\16\2\u0148\u0149\7\25\2\2\u0149\u014a\b\16"+
		"\1\2\u014a\u017e\3\2\2\2\u014b\u014c\5\26\f\2\u014c\u014d\7\60\2\2\u014d"+
		"\u014e\7#\2\2\u014e\u014f\b\16\1\2\u014f\u017e\3\2\2\2\u0150\u0151\5\26"+
		"\f\2\u0151\u0152\7\60\2\2\u0152\u0153\7$\2\2\u0153\u0154\b\16\1\2\u0154"+
		"\u017e\3\2\2\2\u0155\u0156\7\f\2\2\u0156\u0157\5\32\16\2\u0157\u0158\7"+
		"\20\2\2\u0158\u0159\5\32\16\2\u0159\u015a\7\r\2\2\u015a\u015b\b\16\1\2"+
		"\u015b\u017e\3\2\2\2\u015c\u015d\7\f\2\2\u015d\u015e\5\32\16\2\u015e\u015f"+
		"\5\36\20\2\u015f\u0160\7\r\2\2\u0160\u0161\b\16\1\2\u0161\u017e\3\2\2"+
		"\2\u0162\u0163\7&\2\2\u0163\u0164\7\36\2\2\u0164\u0165\7\f\2\2\u0165\u0166"+
		"\5\32\16\2\u0166\u0167\7\20\2\2\u0167\u0168\5\32\16\2\u0168\u0169\7\r"+
		"\2\2\u0169\u016a\b\16\1\2\u016a\u017e\3\2\2\2\u016b\u016c\7&\2\2\u016c"+
		"\u016d\7\36\2\2\u016d\u016e\7\f\2\2\u016e\u016f\5\32\16\2\u016f\u0170"+
		"\5\36\20\2\u0170\u0171\7\r\2\2\u0171\u0172\b\16\1\2\u0172\u017e\3\2\2"+
		"\2\u0173\u0174\7&\2\2\u0174\u0175\7\16\2\2\u0175\u0176\7\16\2\2\u0176"+
		"\u0177\7:\2\2\u0177\u0178\7\24\2\2\u0178\u0179\7\25\2\2\u0179\u017e\b"+
		"\16\1\2\u017a\u017b\5\26\f\2\u017b\u017c\b\16\1\2\u017c\u017e\3\2\2\2"+
		"\u017d\u012d\3\2\2\2\u017d\u0132\3\2\2\2\u017d\u013a\3\2\2\2\u017d\u0142"+
		"\3\2\2\2\u017d\u0146\3\2\2\2\u017d\u014b\3\2\2\2\u017d\u0150\3\2\2\2\u017d"+
		"\u0155\3\2\2\2\u017d\u015c\3\2\2\2\u017d\u0162\3\2\2\2\u017d\u016b\3\2"+
		"\2\2\u017d\u0173\3\2\2\2\u017d\u017a\3\2\2\2\u017e\u01ca\3\2\2\2\u017f"+
		"\u0180\f\34\2\2\u0180\u0181\t\2\2\2\u0181\u0182\5\32\16\35\u0182\u0183"+
		"\b\16\1\2\u0183\u01c9\3\2\2\2\u0184\u0185\f\33\2\2\u0185\u0186\t\3\2\2"+
		"\u0186\u0187\5\32\16\34\u0187\u0188\b\16\1\2\u0188\u01c9\3\2\2\2\u0189"+
		"\u018a\f\30\2\2\u018a\u018b\7\21\2\2\u018b\u018c\5\32\16\31\u018c\u018d"+
		"\b\16\1\2\u018d\u01c9\3\2\2\2\u018e\u018f\f\27\2\2\u018f\u0190\7\22\2"+
		"\2\u0190\u0191\5\32\16\30\u0191\u0192\b\16\1\2\u0192\u01c9\3\2\2\2\u0193"+
		"\u0194\f\26\2\2\u0194\u0195\7\b\2\2\u0195\u0196\5\32\16\27\u0196\u0197"+
		"\b\16\1\2\u0197\u01c9\3\2\2\2\u0198\u0199\f\25\2\2\u0199\u019a\7\t\2\2"+
		"\u019a\u019b\5\32\16\26\u019b\u019c\b\16\1\2\u019c\u01c9\3\2\2\2\u019d"+
		"\u019e\f\24\2\2\u019e\u019f\7\6\2\2\u019f\u01a0\5\32\16\25\u01a0\u01a1"+
		"\b\16\1\2\u01a1\u01c9\3\2\2\2\u01a2\u01a3\f\23\2\2\u01a3\u01a4\7\7\2\2"+
		"\u01a4\u01a5\5\32\16\24\u01a5\u01a6\b\16\1\2\u01a6\u01c9\3\2\2\2\u01a7"+
		"\u01a8\f\22\2\2\u01a8\u01a9\7\4\2\2\u01a9\u01aa\5\32\16\23\u01aa\u01ab"+
		"\b\16\1\2\u01ab\u01c9\3\2\2\2\u01ac\u01ad\f\21\2\2\u01ad\u01ae\7\5\2\2"+
		"\u01ae\u01af\5\32\16\22\u01af\u01b0\b\16\1\2\u01b0\u01c9\3\2\2\2\u01b1"+
		"\u01b2\f\f\2\2\u01b2\u01b3\7\23\2\2\u01b3\u01b4\7\66\2\2\u01b4\u01b5\7"+
		"\24\2\2\u01b5\u01b6\7\25\2\2\u01b6\u01c9\b\16\1\2\u01b7\u01b8\f\13\2\2"+
		"\u01b8\u01b9\7\23\2\2\u01b9\u01ba\7\67\2\2\u01ba\u01bb\7\24\2\2\u01bb"+
		"\u01bc\7\25\2\2\u01bc\u01c9\b\16\1\2\u01bd\u01be\f\n\2\2\u01be\u01bf\7"+
		"\23\2\2\u01bf\u01c0\78\2\2\u01c0\u01c1\7\24\2\2\u01c1\u01c2\7\25\2\2\u01c2"+
		"\u01c9\b\16\1\2\u01c3\u01c4\f\t\2\2\u01c4\u01c5\7\23\2\2\u01c5\u01c6\7"+
		"9\2\2\u01c6\u01c7\7\24\2\2\u01c7\u01c9\7\25\2\2\u01c8\u017f\3\2\2\2\u01c8"+
		"\u0184\3\2\2\2\u01c8\u0189\3\2\2\2\u01c8\u018e\3\2\2\2\u01c8\u0193\3\2"+
		"\2\2\u01c8\u0198\3\2\2\2\u01c8\u019d\3\2\2\2\u01c8\u01a2\3\2\2\2\u01c8"+
		"\u01a7\3\2\2\2\u01c8\u01ac\3\2\2\2\u01c8\u01b1\3\2\2\2\u01c8\u01b7\3\2"+
		"\2\2\u01c8\u01bd\3\2\2\2\u01c8\u01c3\3\2\2\2\u01c9\u01cc\3\2\2\2\u01ca"+
		"\u01c8\3\2\2\2\u01ca\u01cb\3\2\2\2\u01cb\33\3\2\2\2\u01cc\u01ca\3\2\2"+
		"\2\u01cd\u01ce\7\61\2\2\u01ce\u01cf\7\24\2\2\u01cf\u01d0\5\32\16\2\u01d0"+
		"\u01d1\7\25\2\2\u01d1\u01d2\7\20\2\2\u01d2\u01d3\b\17\1\2\u01d3\u01dd"+
		"\3\2\2\2\u01d4\u01d5\7\61\2\2\u01d5\u01d6\7\24\2\2\u01d6\u01d7\5\32\16"+
		"\2\u01d7\u01d8\5\36\20\2\u01d8\u01d9\7\25\2\2\u01d9\u01da\7\20\2\2\u01da"+
		"\u01db\b\17\1\2\u01db\u01dd\3\2\2\2\u01dc\u01cd\3\2\2\2\u01dc\u01d4\3"+
		"\2\2\2\u01dd\35\3\2\2\2\u01de\u01e0\5 \21\2\u01df\u01de\3\2\2\2\u01e0"+
		"\u01e3\3\2\2\2\u01e1\u01df\3\2\2\2\u01e1\u01e2\3\2\2\2\u01e2\u01e4\3\2"+
		"\2\2\u01e3\u01e1\3\2\2\2\u01e4\u01e5\b\20\1\2\u01e5\37\3\2\2\2\u01e6\u01e7"+
		"\7\17\2\2\u01e7\u01e8\5\32\16\2\u01e8\u01e9\b\21\1\2\u01e9!\3\2\2\2\u01ea"+
		"\u01eb\5$\23\2\u01eb\u01ec\b\22\1\2\u01ec#\3\2\2\2\u01ed\u01ee\7C\2\2"+
		"\u01ee\u01ef\5\32\16\2\u01ef\u01f0\5*\26\2\u01f0\u01f1\b\23\1\2\u01f1"+
		"\u0208\3\2\2\2\u01f2\u01f3\7C\2\2\u01f3\u01f4\5\32\16\2\u01f4\u01f5\5"+
		"*\26\2\u01f5\u01f6\7E\2\2\u01f6\u01f7\5*\26\2\u01f7\u01f8\b\23\1\2\u01f8"+
		"\u0208\3\2\2\2\u01f9\u01fa\7C\2\2\u01fa\u01fb\5\32\16\2\u01fb\u01fc\5"+
		"*\26\2\u01fc\u01fd\5(\25\2\u01fd\u01fe\7E\2\2\u01fe\u01ff\5*\26\2\u01ff"+
		"\u0200\b\23\1\2\u0200\u0208\3\2\2\2\u0201\u0202\7C\2\2\u0202\u0203\5\32"+
		"\16\2\u0203\u0204\5*\26\2\u0204\u0205\5(\25\2\u0205\u0206\b\23\1\2\u0206"+
		"\u0208\3\2\2\2\u0207\u01ed\3\2\2\2\u0207\u01f2\3\2\2\2\u0207\u01f9\3\2"+
		"\2\2\u0207\u0201\3\2\2\2\u0208%\3\2\2\2\u0209\u020a\7D\2\2\u020a\u020b"+
		"\5\32\16\2\u020b\u020c\5*\26\2\u020c\u020d\b\24\1\2\u020d\'\3\2\2\2\u020e"+
		"\u0210\5&\24\2\u020f\u020e\3\2\2\2\u0210\u0213\3\2\2\2\u0211\u020f\3\2"+
		"\2\2\u0211\u0212\3\2\2\2\u0212\u0214\3\2\2\2\u0213\u0211\3\2\2\2\u0214"+
		"\u0215\b\25\1\2\u0215)\3\2\2\2\u0216\u0217\7\n\2\2\u0217\u0218\5\n\6\2"+
		"\u0218\u0219\7\13\2\2\u0219\u021a\b\26\1\2\u021a\u021f\3\2\2\2\u021b\u021c"+
		"\7\n\2\2\u021c\u021d\7\13\2\2\u021d\u021f\b\26\1\2\u021e\u0216\3\2\2\2"+
		"\u021e\u021b\3\2\2\2\u021f+\3\2\2\2\u0220\u0221\5.\30\2\u0221\u0222\b"+
		"\27\1\2\u0222\u0227\3\2\2\2\u0223\u0224\5\60\31\2\u0224\u0225\b\27\1\2"+
		"\u0225\u0227\3\2\2\2\u0226\u0220\3\2\2\2\u0226\u0223\3\2\2\2\u0227-\3"+
		"\2\2\2\u0228\u0229\7F\2\2\u0229\u022a\5\32\16\2\u022a\u022b\5*\26\2\u022b"+
		"\u022c\b\30\1\2\u022c/\3\2\2\2\u022d\u022e\7H\2\2\u022e\u022f\5*\26\2"+
		"\u022f\u0230\b\31\1\2\u0230\61\3\2\2\2\26\65<Tet\u008b\u009e\u00fa\u0108"+
		"\u011c\u012a\u017d\u01c8\u01ca\u01dc\u01e1\u0207\u0211\u021e\u0226";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}