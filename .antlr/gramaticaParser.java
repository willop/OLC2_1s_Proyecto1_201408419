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
		RULE_declaracion = 8, RULE_tipovariable = 9, RULE_identificadores = 10, 
		RULE_valores = 11, RULE_arreglo = 12, RULE_expresion = 13, RULE_impresion = 14, 
		RULE_impresionexpresion = 15, RULE_expcoma = 16, RULE_condicionales = 17, 
		RULE_funcionif = 18, RULE_funcionelseif = 19, RULE_listaelseif = 20, RULE_bloque = 21, 
		RULE_bucles = 22, RULE_fwhile = 23, RULE_floop = 24;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "funciones", "declararParametros", "main", "instrucciones", 
			"instruccion", "asignacion", "control", "declaracion", "tipovariable", 
			"identificadores", "valores", "arreglo", "expresion", "impresion", "impresionexpresion", 
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
			setState(53);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(50);
					((StartContext)_localctx).funciones = funciones();
					((StartContext)_localctx).e.add(((StartContext)_localctx).funciones);
					}
					} 
				}
				setState(55);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			}
			setState(56);
			((StartContext)_localctx).listtt = main();
			setState(60);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==TKR_fn) {
				{
				{
				setState(57);
				((StartContext)_localctx).funciones = funciones();
				((StartContext)_localctx).r.add(((StartContext)_localctx).funciones);
				}
				}
				setState(62);
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
			setState(84);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(65);
				match(TKR_fn);
				setState(66);
				((FuncionesContext)_localctx).idd = match(TK_id);
				setState(67);
				match(TK_par_apertura);
				setState(68);
				match(TK_par_cierre);
				setState(69);
				match(TK_corchete_apertura);
				setState(70);
				((FuncionesContext)_localctx).instrucciones = instrucciones();
				setState(71);
				match(TK_corchete_cierre);
				_localctx.inst = Funcion.NewFuncion((((FuncionesContext)_localctx).idd!=null?((FuncionesContext)_localctx).idd.getText():null),Enum.STRING,nil,((FuncionesContext)_localctx).instrucciones.lista)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(74);
				match(TKR_fn);
				setState(75);
				((FuncionesContext)_localctx).idd = match(TK_id);
				setState(76);
				match(TK_par_apertura);
				setState(77);
				((FuncionesContext)_localctx).declararParametros = declararParametros(0);
				setState(78);
				match(TK_par_cierre);
				setState(79);
				match(TK_corchete_apertura);
				setState(80);
				((FuncionesContext)_localctx).instrucciones = instrucciones();
				setState(81);
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
			setState(87);
			((DeclararParametrosContext)_localctx).e1 = match(TK_id);
			setState(88);
			match(TK_dosPuntos);
			setState(89);
			((DeclararParametrosContext)_localctx).e2 = tipovariable();
			_localctx.lista = arrayList.New()
			                                                                                                        _localctx.lista.Add(Funcion.NewParam((((DeclararParametrosContext)_localctx).e1!=null?((DeclararParametrosContext)_localctx).e1.getText():null),((DeclararParametrosContext)_localctx).e2.tipovar))
			}
			_ctx.stop = _input.LT(-1);
			setState(101);
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
					setState(92);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(93);
					match(TK_coma);
					setState(94);
					((DeclararParametrosContext)_localctx).e1 = match(TK_id);
					setState(95);
					match(TK_dosPuntos);
					setState(96);
					((DeclararParametrosContext)_localctx).e2 = tipovariable();
					 ((DeclararParametrosContext)_localctx).list.lista.Add(Funcion.NewParam((((DeclararParametrosContext)_localctx).e1!=null?((DeclararParametrosContext)_localctx).e1.getText():null),((DeclararParametrosContext)_localctx).e2.tipovar))
					                                                                                                                  _localctx.lista = ((DeclararParametrosContext)_localctx).list.lista
					}
					} 
				}
				setState(103);
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
			setState(104);
			match(TKR_fn);
			setState(105);
			match(TKR_main);
			setState(106);
			match(TK_par_apertura);
			setState(107);
			match(TK_par_cierre);
			setState(108);
			match(TK_corchete_apertura);
			setState(109);
			((MainContext)_localctx).instrucciones = instrucciones();
			setState(110);
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
			setState(116);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << TK_llave_apertura) | (1L << TK_par_apertura) | (1L << TK_menos) | (1L << TK_amp) | (1L << TK_sig_admiracion) | (1L << TKR_pow) | (1L << TKR_powf) | (1L << TKR_let) | (1L << TKR_println) | (1L << TKR_true) | (1L << TKR_false))) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & ((1L << (TKR_if - 65)) | (1L << (TKR_while - 65)) | (1L << (TKR_break - 65)) | (1L << (TKR_loop - 65)) | (1L << (TK_entero - 65)) | (1L << (TK_decimal - 65)) | (1L << (TK_id - 65)) | (1L << (TK_cadena - 65)) | (1L << (TK_caracter - 65)))) != 0)) {
				{
				{
				setState(113);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).e.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(118);
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
		public AsignacionContext asignacion;
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
		public AsignacionContext asignacion() {
			return getRuleContext(AsignacionContext.class,0);
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
			setState(145);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(121);
				((InstruccionContext)_localctx).expresion = expresion(0);
				fmt.Println("mensaje en instrucciones: ",((InstruccionContext)_localctx).expresion.exp)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(124);
				((InstruccionContext)_localctx).impresion = impresion();
				_localctx.inst = ((InstruccionContext)_localctx).impresion.inst
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(127);
				((InstruccionContext)_localctx).declaracion = declaracion();
				_localctx.inst = ((InstruccionContext)_localctx).declaracion.inst
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(130);
				((InstruccionContext)_localctx).identificadores = identificadores();
				_localctx.inst = ((InstruccionContext)_localctx).identificadores.inst
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(133);
				((InstruccionContext)_localctx).condicionales = condicionales();
				_localctx.inst = ((InstruccionContext)_localctx).condicionales.inst
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(136);
				((InstruccionContext)_localctx).bucles = bucles();
				_localctx.inst = ((InstruccionContext)_localctx).bucles.inst
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(139);
				((InstruccionContext)_localctx).control = control();
				_localctx.inst = ((InstruccionContext)_localctx).control.inst
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(142);
				((InstruccionContext)_localctx).asignacion = asignacion();
				_localctx.inst = ((InstruccionContext)_localctx).asignacion.inst
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

	public static class AsignacionContext extends ParserRuleContext {
		public Interfaces.Instruccion inst;
		public Token TK_id;
		public TerminalNode TK_id() { return getToken(gramaticaParser.TK_id, 0); }
		public TerminalNode TK_par_apertura() { return getToken(gramaticaParser.TK_par_apertura, 0); }
		public TerminalNode TK_par_cierre() { return getToken(gramaticaParser.TK_par_cierre, 0); }
		public TerminalNode TK_pcoma() { return getToken(gramaticaParser.TK_pcoma, 0); }
		public AsignacionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacion; }
	}

	public final AsignacionContext asignacion() throws RecognitionException {
		AsignacionContext _localctx = new AsignacionContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_asignacion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(147);
			((AsignacionContext)_localctx).TK_id = match(TK_id);
			setState(148);
			match(TK_par_apertura);
			setState(149);
			match(TK_par_cierre);
			setState(150);
			match(TK_pcoma);
			_localctx.inst = Funcion.NewLlamadafuncion((((AsignacionContext)_localctx).TK_id!=null?((AsignacionContext)_localctx).TK_id.getText():null),true,nil)
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
			setState(153);
			match(TKR_break);
			setState(154);
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
		enterRule(_localctx, 16, RULE_declaracion);
		try {
			setState(218);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(157);
				match(TKR_let);
				setState(158);
				match(TKR_mut);
				setState(159);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(160);
				match(TK_dosPuntos);
				setState(161);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(162);
				match(TK_igual);
				setState(163);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(164);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar,((DeclaracionContext)_localctx).expresion.exp,false,false,true) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(167);
				match(TKR_let);
				setState(168);
				match(TKR_mut);
				setState(169);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(170);
				match(TK_igual);
				setState(171);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(172);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),Enum.SINTIPO,((DeclaracionContext)_localctx).expresion.exp,false,false,true) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(175);
				match(TKR_let);
				setState(176);
				match(TKR_mut);
				setState(177);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(178);
				match(TK_dosPuntos);
				setState(179);
				match(TK_llave_apertura);
				setState(180);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(181);
				match(TK_pcoma);
				setState(182);
				((DeclaracionContext)_localctx).cant = valores();
				setState(183);
				match(TK_llave_cierre);
				setState(184);
				match(TK_igual);
				setState(185);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(186);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracionArray((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar,((DeclaracionContext)_localctx).cant.exp,((DeclaracionContext)_localctx).expresion.exp,true)
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(189);
				match(TKR_let);
				setState(190);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(191);
				match(TK_dosPuntos);
				setState(192);
				match(TK_llave_apertura);
				setState(193);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(194);
				match(TK_pcoma);
				setState(195);
				((DeclaracionContext)_localctx).cant = valores();
				setState(196);
				match(TK_llave_cierre);
				setState(197);
				match(TK_igual);
				setState(198);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(199);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracionArray((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar,((DeclaracionContext)_localctx).cant.exp,((DeclaracionContext)_localctx).expresion.exp,false)
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(202);
				match(TKR_let);
				setState(203);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(204);
				match(TK_dosPuntos);
				setState(205);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(206);
				match(TK_igual);
				setState(207);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(208);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar,((DeclaracionContext)_localctx).expresion.exp,false,false,false) 
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(211);
				match(TKR_let);
				setState(212);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(213);
				match(TK_igual);
				setState(214);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(215);
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
			setState(232);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TKR_numericos_enteros:
				enterOuterAlt(_localctx, 1);
				{
				setState(220);
				match(TKR_numericos_enteros);
				_localctx.tipovar = Enum.INTEGER
				}
				break;
			case TKR_numericos_flotantes:
				enterOuterAlt(_localctx, 2);
				{
				setState(222);
				match(TKR_numericos_flotantes);
				_localctx.tipovar = Enum.FLOAT
				}
				break;
			case TKR_String:
				enterOuterAlt(_localctx, 3);
				{
				setState(224);
				match(TKR_String);
				_localctx.tipovar = Enum.STRING
				}
				break;
			case TKR_bool:
				enterOuterAlt(_localctx, 4);
				{
				setState(226);
				match(TKR_bool);
				_localctx.tipovar = Enum.BOOLEAN
				}
				break;
			case TKR_char:
				enterOuterAlt(_localctx, 5);
				{
				setState(228);
				match(TKR_char);
				_localctx.tipovar = Enum.CHAR
				}
				break;
			case TKR_usize:
				enterOuterAlt(_localctx, 6);
				{
				setState(230);
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
		enterRule(_localctx, 20, RULE_identificadores);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(234);
			((IdentificadoresContext)_localctx).e1 = match(TK_id);
			setState(235);
			match(TK_igual);
			setState(236);
			((IdentificadoresContext)_localctx).e2 = expresion(0);
			setState(237);
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
		enterRule(_localctx, 22, RULE_valores);
		try {
			setState(258);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_entero:
				enterOuterAlt(_localctx, 1);
				{
				setState(240);
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
				setState(242);
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
				setState(244);
				((ValoresContext)_localctx).TK_cadena = match(TK_cadena);

				                str:= (((ValoresContext)_localctx).TK_cadena!=null?((ValoresContext)_localctx).TK_cadena.getText():null)[1:len((((ValoresContext)_localctx).TK_cadena!=null?((ValoresContext)_localctx).TK_cadena.getText():null))-1]
				                _localctx.exp = Expresion.NuevoPrimitivo(str,Enum.STRING)
				}
				break;
			case TK_caracter:
				enterOuterAlt(_localctx, 4);
				{
				setState(246);
				((ValoresContext)_localctx).TK_caracter = match(TK_caracter);
				str:= (((ValoresContext)_localctx).TK_caracter!=null?((ValoresContext)_localctx).TK_caracter.getText():null)[1:len((((ValoresContext)_localctx).TK_caracter!=null?((ValoresContext)_localctx).TK_caracter.getText():null))-1]
				                _localctx.exp = Expresion.NuevoPrimitivo(str,Enum.CHAR)
				}
				break;
			case TK_amp:
				enterOuterAlt(_localctx, 5);
				{
				setState(248);
				match(TK_amp);
				setState(249);
				((ValoresContext)_localctx).TKR_Str = match(TKR_Str);
				str:= (((ValoresContext)_localctx).TKR_Str!=null?((ValoresContext)_localctx).TKR_Str.getText():null)[1:len((((ValoresContext)_localctx).TKR_Str!=null?((ValoresContext)_localctx).TKR_Str.getText():null))-1]
				                _localctx.exp = Expresion.NuevoPrimitivo(str,Enum.STR)
				}
				break;
			case TKR_true:
				enterOuterAlt(_localctx, 6);
				{
				setState(251);
				match(TKR_true);
				_localctx.exp = Expresion.NuevoPrimitivo(true, Enum.BOOLEAN)
				}
				break;
			case TKR_false:
				enterOuterAlt(_localctx, 7);
				{
				setState(253);
				match(TKR_false);
				_localctx.exp = Expresion.NuevoPrimitivo(false,Enum.BOOLEAN)
				}
				break;
			case TK_id:
				enterOuterAlt(_localctx, 8);
				{
				setState(255);
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
		int _startState = 24;
		enterRecursionRule(_localctx, 24, RULE_arreglo, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(261);
			((ArregloContext)_localctx).TK_id = match(TK_id);
			_localctx.exp = Expresiones.NewLlamarvariable((((ArregloContext)_localctx).TK_id!=null?((ArregloContext)_localctx).TK_id.getText():null))
			}
			_ctx.stop = _input.LT(-1);
			setState(272);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
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
					setState(264);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(265);
					match(TK_llave_apertura);
					setState(266);
					((ArregloContext)_localctx).expresion = expresion(0);
					setState(267);
					match(TK_llave_cierre);
					_localctx.exp = Expresiones.NewAccesoArray(((ArregloContext)_localctx).ar.exp,((ArregloContext)_localctx).expresion.exp)
					}
					} 
				}
				setState(274);
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
		int _startState = 26;
		enterRecursionRule(_localctx, 26, RULE_expresion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(331);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				{
				setState(276);
				match(TK_menos);
				setState(277);
				((ExpresionContext)_localctx).e1 = expresion(24);
				numero := -1
				                                                                                                        numm := Expresion.NuevoPrimitivo(numero, Enum.INTEGER)
				                                                                                                                _localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp, numm ,Enum.MULTIPLICACION)
				                                                                                                
				}
				break;
			case 2:
				{
				setState(280);
				match(TKR_pow);
				setState(281);
				match(TK_par_apertura);
				setState(282);
				((ExpresionContext)_localctx).e1 = expresion(0);
				setState(283);
				match(TK_coma);
				setState(284);
				((ExpresionContext)_localctx).e2 = expresion(0);
				setState(285);
				match(TK_par_cierre);
				_localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.POW)
				}
				break;
			case 3:
				{
				setState(288);
				match(TKR_powf);
				setState(289);
				match(TK_par_apertura);
				setState(290);
				((ExpresionContext)_localctx).e1 = expresion(0);
				setState(291);
				match(TK_coma);
				setState(292);
				((ExpresionContext)_localctx).e2 = expresion(0);
				setState(293);
				match(TK_par_cierre);
				_localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.POWF)
				}
				break;
			case 4:
				{
				setState(296);
				match(TK_sig_admiracion);
				setState(297);
				((ExpresionContext)_localctx).e1 = expresion(11);
				_localctx.exp = Expresiones.NuevaLogica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e1.exp,Enum.NOT)
				}
				break;
			case 5:
				{
				setState(300);
				match(TK_par_apertura);
				setState(301);
				((ExpresionContext)_localctx).va = expresion(0);
				setState(302);
				match(TK_par_cierre);
				_localctx.exp = ((ExpresionContext)_localctx).va.exp
				}
				break;
			case 6:
				{
				setState(305);
				((ExpresionContext)_localctx).val = valores();
				setState(306);
				match(TKR_as);
				setState(307);
				match(TKR_numericos_enteros);
				_localctx.exp = Expresiones.NewAsi64(((ExpresionContext)_localctx).val.exp)
				}
				break;
			case 7:
				{
				setState(310);
				((ExpresionContext)_localctx).val = valores();
				setState(311);
				match(TKR_as);
				setState(312);
				match(TKR_numericos_flotantes);
				_localctx.exp = Expresiones.NewAsf64(((ExpresionContext)_localctx).val.exp)
				}
				break;
			case 8:
				{
				setState(315);
				match(TK_llave_apertura);
				setState(316);
				((ExpresionContext)_localctx).e1 = expresion(0);
				setState(317);
				match(TK_pcoma);
				setState(318);
				((ExpresionContext)_localctx).e2 = expresion(0);
				setState(319);
				match(TK_llave_cierre);
				_localctx.exp = Expresion.NewArray(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,nil,Enum.MULTIPLE)
				}
				break;
			case 9:
				{
				setState(322);
				match(TK_llave_apertura);
				setState(323);
				((ExpresionContext)_localctx).e1 = expresion(0);
				setState(324);
				((ExpresionContext)_localctx).l1 = impresionexpresion();
				setState(325);
				match(TK_llave_cierre);
				_localctx.exp = Expresion.NewArray(((ExpresionContext)_localctx).e1.exp,nil,((ExpresionContext)_localctx).l1.lista,Enum.NORMAL)
				}
				break;
			case 10:
				{
				setState(328);
				((ExpresionContext)_localctx).vall = valores();
				_localctx.exp = ((ExpresionContext)_localctx).vall.exp
				                                                                                        fmt.Println(_localctx.exp)
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(408);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(406);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
					case 1:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(333);
						if (!(precpred(_ctx, 23))) throw new FailedPredicateException(this, "precpred(_ctx, 23)");
						setState(334);
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
						setState(335);
						((ExpresionContext)_localctx).e2 = expresion(24);

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
						setState(338);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(339);
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
						setState(340);
						((ExpresionContext)_localctx).e2 = expresion(23);

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
						setState(343);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(344);
						match(TK_menor);
						setState(345);
						((ExpresionContext)_localctx).e2 = expresion(20);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.MENOR_QUE)
						}
						break;
					case 4:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(348);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(349);
						match(TK_mayor);
						setState(350);
						((ExpresionContext)_localctx).e2 = expresion(19);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.MAYOR_QUE)
						}
						break;
					case 5:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(353);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(354);
						match(TK_mayor_igual);
						setState(355);
						((ExpresionContext)_localctx).e2 = expresion(18);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.MAYOR_IGUAL)
						}
						break;
					case 6:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(358);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(359);
						match(TK_menor_igual);
						setState(360);
						((ExpresionContext)_localctx).e2 = expresion(17);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.MENOR_IGUAL)
						}
						break;
					case 7:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(363);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(364);
						match(TK_igualacion);
						setState(365);
						((ExpresionContext)_localctx).e2 = expresion(16);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.IGUALDAD)
						}
						break;
					case 8:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(368);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(369);
						match(TK_diferente);
						setState(370);
						((ExpresionContext)_localctx).e2 = expresion(15);
						_localctx.exp = Expresiones.NuevaRelacional(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.DESIGUAL)
						}
						break;
					case 9:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(373);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(374);
						match(TK_or);
						setState(375);
						((ExpresionContext)_localctx).e2 = expresion(14);
						_localctx.exp = Expresiones.NuevaLogica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.OR)
						}
						break;
					case 10:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(378);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(379);
						match(TK_and);
						setState(380);
						((ExpresionContext)_localctx).e2 = expresion(13);
						_localctx.exp = Expresiones.NuevaLogica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Enum.AND)
						}
						break;
					case 11:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.vll = _prevctx;
						_localctx.vll = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(383);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(384);
						match(TK_punto);
						setState(385);
						match(TKR_abs);
						setState(386);
						match(TK_par_apertura);
						setState(387);
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
						setState(389);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(390);
						match(TK_punto);
						setState(391);
						match(TKR_sqrt);
						setState(392);
						match(TK_par_apertura);
						setState(393);
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
						setState(395);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(396);
						match(TK_punto);
						setState(397);
						match(TKR_to_string);
						setState(398);
						match(TK_par_apertura);
						setState(399);
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
						setState(401);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(402);
						match(TK_punto);
						setState(403);
						match(TKR_clone);
						setState(404);
						match(TK_par_apertura);
						setState(405);
						match(TK_par_cierre);
						}
						break;
					}
					} 
				}
				setState(410);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
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
		enterRule(_localctx, 28, RULE_impresion);
		try {
			setState(426);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(411);
				match(TKR_println);
				setState(412);
				match(TK_par_apertura);
				setState(413);
				((ImpresionContext)_localctx).e1 = expresion(0);
				setState(414);
				match(TK_par_cierre);
				setState(415);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevoPrint(((ImpresionContext)_localctx).e1.exp,nil)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(418);
				match(TKR_println);
				setState(419);
				match(TK_par_apertura);
				setState(420);
				((ImpresionContext)_localctx).e2 = expresion(0);
				setState(421);
				((ImpresionContext)_localctx).li = impresionexpresion();
				setState(422);
				match(TK_par_cierre);
				setState(423);
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
		enterRule(_localctx, 30, RULE_impresionexpresion);
		_localctx.lista = arrayList.New()
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(431);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==TK_coma) {
				{
				{
				setState(428);
				((ImpresionexpresionContext)_localctx).expcoma = expcoma();
				((ImpresionexpresionContext)_localctx).list.add(((ImpresionexpresionContext)_localctx).expcoma);
				}
				}
				setState(433);
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
		enterRule(_localctx, 32, RULE_expcoma);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(436);
			match(TK_coma);
			setState(437);
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
		enterRule(_localctx, 34, RULE_condicionales);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(440);
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
		enterRule(_localctx, 36, RULE_funcionif);
		try {
			setState(469);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(443);
				match(TKR_if);
				setState(444);
				((FuncionifContext)_localctx).e1 = expresion(0);
				setState(445);
				((FuncionifContext)_localctx).ee = bloque();
				_localctx.inst = Instruccion.NewIf(((FuncionifContext)_localctx).e1.exp,((FuncionifContext)_localctx).ee.lista,nil,nil )
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(448);
				match(TKR_if);
				setState(449);
				((FuncionifContext)_localctx).e1 = expresion(0);
				setState(450);
				((FuncionifContext)_localctx).e5 = bloque();
				setState(451);
				match(TKR_else);
				setState(452);
				((FuncionifContext)_localctx).b2 = bloque();
				_localctx.inst = Instruccion.NewIf(((FuncionifContext)_localctx).e1.exp,((FuncionifContext)_localctx).e5.lista,nil,((FuncionifContext)_localctx).b2.lista)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(455);
				match(TKR_if);
				setState(456);
				((FuncionifContext)_localctx).e1 = expresion(0);
				setState(457);
				((FuncionifContext)_localctx).b1 = bloque();
				setState(458);
				((FuncionifContext)_localctx).listaelseif = listaelseif();
				setState(459);
				match(TKR_else);
				setState(460);
				((FuncionifContext)_localctx).b2 = bloque();
				_localctx.inst = Instruccion.NewIf(((FuncionifContext)_localctx).e1.exp,((FuncionifContext)_localctx).b1.lista,((FuncionifContext)_localctx).listaelseif.lista,((FuncionifContext)_localctx).b2.lista)
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(463);
				match(TKR_if);
				setState(464);
				((FuncionifContext)_localctx).e1 = expresion(0);
				setState(465);
				((FuncionifContext)_localctx).b1 = bloque();
				setState(466);
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
		enterRule(_localctx, 38, RULE_funcionelseif);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(471);
			match(TKR_elseif);
			setState(472);
			((FuncionelseifContext)_localctx).e1 = expresion(0);
			setState(473);
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
		enterRule(_localctx, 40, RULE_listaelseif);
		 _localctx.lista = arrayList.New()
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(479);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==TKR_elseif) {
				{
				{
				setState(476);
				((ListaelseifContext)_localctx).funcionelseif = funcionelseif();
				((ListaelseifContext)_localctx).list.add(((ListaelseifContext)_localctx).funcionelseif);
				}
				}
				setState(481);
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
		enterRule(_localctx, 42, RULE_bloque);
		try {
			setState(492);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(484);
				match(TK_corchete_apertura);
				setState(485);
				((BloqueContext)_localctx).instrucciones = instrucciones();
				setState(486);
				match(TK_corchete_cierre);
				_localctx.lista = ((BloqueContext)_localctx).instrucciones.lista
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(489);
				match(TK_corchete_apertura);
				setState(490);
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
		enterRule(_localctx, 44, RULE_bucles);
		try {
			setState(500);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TKR_while:
				enterOuterAlt(_localctx, 1);
				{
				setState(494);
				((BuclesContext)_localctx).fwhile = fwhile();
				_localctx.inst = ((BuclesContext)_localctx).fwhile.inst
				}
				break;
			case TKR_loop:
				enterOuterAlt(_localctx, 2);
				{
				setState(497);
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
		enterRule(_localctx, 46, RULE_fwhile);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(502);
			match(TKR_while);
			setState(503);
			((FwhileContext)_localctx).e1 = expresion(0);
			setState(504);
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
		enterRule(_localctx, 48, RULE_floop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(507);
			match(TKR_loop);
			setState(508);
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
		case 12:
			return arreglo_sempred((ArregloContext)_localctx, predIndex);
		case 13:
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
			return precpred(_ctx, 23);
		case 3:
			return precpred(_ctx, 22);
		case 4:
			return precpred(_ctx, 19);
		case 5:
			return precpred(_ctx, 18);
		case 6:
			return precpred(_ctx, 17);
		case 7:
			return precpred(_ctx, 16);
		case 8:
			return precpred(_ctx, 15);
		case 9:
			return precpred(_ctx, 14);
		case 10:
			return precpred(_ctx, 13);
		case 11:
			return precpred(_ctx, 12);
		case 12:
			return precpred(_ctx, 7);
		case 13:
			return precpred(_ctx, 6);
		case 14:
			return precpred(_ctx, 5);
		case 15:
			return precpred(_ctx, 4);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3N\u0202\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\3\2\7\2\66\n\2\f\2\16\29\13\2\3\2\3\2\7\2=\n\2\f\2\16\2@\13"+
		"\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\5\3W\n\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\7\4f\n\4\f\4\16\4i\13\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3"+
		"\6\7\6u\n\6\f\6\16\6x\13\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7"+
		"\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7\u0094"+
		"\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\3\n\5\n\u00dd\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3"+
		"\13\3\13\3\13\5\13\u00eb\n\13\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r"+
		"\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\5\r\u0105\n\r"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\7\16\u0111\n\16\f\16"+
		"\16\16\u0114\13\16\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3"+
		"\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3"+
		"\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3"+
		"\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3"+
		"\17\3\17\3\17\3\17\5\17\u014e\n\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\7\17\u0199\n\17\f\17"+
		"\16\17\u019c\13\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3"+
		"\20\3\20\3\20\3\20\3\20\5\20\u01ad\n\20\3\21\7\21\u01b0\n\21\f\21\16\21"+
		"\u01b3\13\21\3\21\3\21\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\24\3\24\3"+
		"\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3"+
		"\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\5\24\u01d8\n\24\3\25"+
		"\3\25\3\25\3\25\3\25\3\26\7\26\u01e0\n\26\f\26\16\26\u01e3\13\26\3\26"+
		"\3\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\5\27\u01ef\n\27\3\30\3\30"+
		"\3\30\3\30\3\30\3\30\5\30\u01f7\n\30\3\31\3\31\3\31\3\31\3\31\3\32\3\32"+
		"\3\32\3\32\3\32\2\5\6\32\34\33\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36"+
		" \"$&(*,.\60\62\2\4\3\2\31\33\3\2\27\30\2\u0225\2\67\3\2\2\2\4V\3\2\2"+
		"\2\6X\3\2\2\2\bj\3\2\2\2\nv\3\2\2\2\f\u0093\3\2\2\2\16\u0095\3\2\2\2\20"+
		"\u009b\3\2\2\2\22\u00dc\3\2\2\2\24\u00ea\3\2\2\2\26\u00ec\3\2\2\2\30\u0104"+
		"\3\2\2\2\32\u0106\3\2\2\2\34\u014d\3\2\2\2\36\u01ac\3\2\2\2 \u01b1\3\2"+
		"\2\2\"\u01b6\3\2\2\2$\u01ba\3\2\2\2&\u01d7\3\2\2\2(\u01d9\3\2\2\2*\u01e1"+
		"\3\2\2\2,\u01ee\3\2\2\2.\u01f6\3\2\2\2\60\u01f8\3\2\2\2\62\u01fd\3\2\2"+
		"\2\64\66\5\4\3\2\65\64\3\2\2\2\669\3\2\2\2\67\65\3\2\2\2\678\3\2\2\28"+
		":\3\2\2\29\67\3\2\2\2:>\5\b\5\2;=\5\4\3\2<;\3\2\2\2=@\3\2\2\2><\3\2\2"+
		"\2>?\3\2\2\2?A\3\2\2\2@>\3\2\2\2AB\b\2\1\2B\3\3\2\2\2CD\7\64\2\2DE\7K"+
		"\2\2EF\7\24\2\2FG\7\25\2\2GH\7\n\2\2HI\5\n\6\2IJ\7\13\2\2JK\b\3\1\2KW"+
		"\3\2\2\2LM\7\64\2\2MN\7K\2\2NO\7\24\2\2OP\5\6\4\2PQ\7\25\2\2QR\7\n\2\2"+
		"RS\5\n\6\2ST\7\13\2\2TU\b\3\1\2UW\3\2\2\2VC\3\2\2\2VL\3\2\2\2W\5\3\2\2"+
		"\2XY\b\4\1\2YZ\7K\2\2Z[\7\16\2\2[\\\5\24\13\2\\]\b\4\1\2]g\3\2\2\2^_\f"+
		"\4\2\2_`\7\17\2\2`a\7K\2\2ab\7\16\2\2bc\5\24\13\2cd\b\4\1\2df\3\2\2\2"+
		"e^\3\2\2\2fi\3\2\2\2ge\3\2\2\2gh\3\2\2\2h\7\3\2\2\2ig\3\2\2\2jk\7\64\2"+
		"\2kl\7B\2\2lm\7\24\2\2mn\7\25\2\2no\7\n\2\2op\5\n\6\2pq\7\13\2\2qr\b\5"+
		"\1\2r\t\3\2\2\2su\5\f\7\2ts\3\2\2\2ux\3\2\2\2vt\3\2\2\2vw\3\2\2\2wy\3"+
		"\2\2\2xv\3\2\2\2yz\b\6\1\2z\13\3\2\2\2{|\5\34\17\2|}\b\7\1\2}\u0094\3"+
		"\2\2\2~\177\5\36\20\2\177\u0080\b\7\1\2\u0080\u0094\3\2\2\2\u0081\u0082"+
		"\5\22\n\2\u0082\u0083\b\7\1\2\u0083\u0094\3\2\2\2\u0084\u0085\5\26\f\2"+
		"\u0085\u0086\b\7\1\2\u0086\u0094\3\2\2\2\u0087\u0088\5$\23\2\u0088\u0089"+
		"\b\7\1\2\u0089\u0094\3\2\2\2\u008a\u008b\5.\30\2\u008b\u008c\b\7\1\2\u008c"+
		"\u0094\3\2\2\2\u008d\u008e\5\20\t\2\u008e\u008f\b\7\1\2\u008f\u0094\3"+
		"\2\2\2\u0090\u0091\5\16\b\2\u0091\u0092\b\7\1\2\u0092\u0094\3\2\2\2\u0093"+
		"{\3\2\2\2\u0093~\3\2\2\2\u0093\u0081\3\2\2\2\u0093\u0084\3\2\2\2\u0093"+
		"\u0087\3\2\2\2\u0093\u008a\3\2\2\2\u0093\u008d\3\2\2\2\u0093\u0090\3\2"+
		"\2\2\u0094\r\3\2\2\2\u0095\u0096\7K\2\2\u0096\u0097\7\24\2\2\u0097\u0098"+
		"\7\25\2\2\u0098\u0099\7\20\2\2\u0099\u009a\b\b\1\2\u009a\17\3\2\2\2\u009b"+
		"\u009c\7G\2\2\u009c\u009d\7\20\2\2\u009d\u009e\b\t\1\2\u009e\21\3\2\2"+
		"\2\u009f\u00a0\7-\2\2\u00a0\u00a1\7.\2\2\u00a1\u00a2\7K\2\2\u00a2\u00a3"+
		"\7\16\2\2\u00a3\u00a4\5\24\13\2\u00a4\u00a5\7\26\2\2\u00a5\u00a6\5\34"+
		"\17\2\u00a6\u00a7\7\20\2\2\u00a7\u00a8\b\n\1\2\u00a8\u00dd\3\2\2\2\u00a9"+
		"\u00aa\7-\2\2\u00aa\u00ab\7.\2\2\u00ab\u00ac\7K\2\2\u00ac\u00ad\7\26\2"+
		"\2\u00ad\u00ae\5\34\17\2\u00ae\u00af\7\20\2\2\u00af\u00b0\b\n\1\2\u00b0"+
		"\u00dd\3\2\2\2\u00b1\u00b2\7-\2\2\u00b2\u00b3\7.\2\2\u00b3\u00b4\7K\2"+
		"\2\u00b4\u00b5\7\16\2\2\u00b5\u00b6\7\f\2\2\u00b6\u00b7\5\24\13\2\u00b7"+
		"\u00b8\7\20\2\2\u00b8\u00b9\5\30\r\2\u00b9\u00ba\7\r\2\2\u00ba\u00bb\7"+
		"\26\2\2\u00bb\u00bc\5\34\17\2\u00bc\u00bd\7\20\2\2\u00bd\u00be\b\n\1\2"+
		"\u00be\u00dd\3\2\2\2\u00bf\u00c0\7-\2\2\u00c0\u00c1\7K\2\2\u00c1\u00c2"+
		"\7\16\2\2\u00c2\u00c3\7\f\2\2\u00c3\u00c4\5\24\13\2\u00c4\u00c5\7\20\2"+
		"\2\u00c5\u00c6\5\30\r\2\u00c6\u00c7\7\r\2\2\u00c7\u00c8\7\26\2\2\u00c8"+
		"\u00c9\5\34\17\2\u00c9\u00ca\7\20\2\2\u00ca\u00cb\b\n\1\2\u00cb\u00dd"+
		"\3\2\2\2\u00cc\u00cd\7-\2\2\u00cd\u00ce\7K\2\2\u00ce\u00cf\7\16\2\2\u00cf"+
		"\u00d0\5\24\13\2\u00d0\u00d1\7\26\2\2\u00d1\u00d2\5\34\17\2\u00d2\u00d3"+
		"\7\20\2\2\u00d3\u00d4\b\n\1\2\u00d4\u00dd\3\2\2\2\u00d5\u00d6\7-\2\2\u00d6"+
		"\u00d7\7K\2\2\u00d7\u00d8\7\26\2\2\u00d8\u00d9\5\34\17\2\u00d9\u00da\7"+
		"\20\2\2\u00da\u00db\b\n\1\2\u00db\u00dd\3\2\2\2\u00dc\u009f\3\2\2\2\u00dc"+
		"\u00a9\3\2\2\2\u00dc\u00b1\3\2\2\2\u00dc\u00bf\3\2\2\2\u00dc\u00cc\3\2"+
		"\2\2\u00dc\u00d5\3\2\2\2\u00dd\23\3\2\2\2\u00de\u00df\7#\2\2\u00df\u00eb"+
		"\b\13\1\2\u00e0\u00e1\7$\2\2\u00e1\u00eb\b\13\1\2\u00e2\u00e3\7+\2\2\u00e3"+
		"\u00eb\b\13\1\2\u00e4\u00e5\7(\2\2\u00e5\u00eb\b\13\1\2\u00e6\u00e7\7"+
		")\2\2\u00e7\u00eb\b\13\1\2\u00e8\u00e9\7,\2\2\u00e9\u00eb\b\13\1\2\u00ea"+
		"\u00de\3\2\2\2\u00ea\u00e0\3\2\2\2\u00ea\u00e2\3\2\2\2\u00ea\u00e4\3\2"+
		"\2\2\u00ea\u00e6\3\2\2\2\u00ea\u00e8\3\2\2\2\u00eb\25\3\2\2\2\u00ec\u00ed"+
		"\7K\2\2\u00ed\u00ee\7\26\2\2\u00ee\u00ef\5\34\17\2\u00ef\u00f0\7\20\2"+
		"\2\u00f0\u00f1\b\f\1\2\u00f1\27\3\2\2\2\u00f2\u00f3\7I\2\2\u00f3\u0105"+
		"\b\r\1\2\u00f4\u00f5\7J\2\2\u00f5\u0105\b\r\1\2\u00f6\u00f7\7M\2\2\u00f7"+
		"\u0105\b\r\1\2\u00f8\u00f9\7N\2\2\u00f9\u0105\b\r\1\2\u00fa\u00fb\7\35"+
		"\2\2\u00fb\u00fc\7*\2\2\u00fc\u0105\b\r\1\2\u00fd\u00fe\7\62\2\2\u00fe"+
		"\u0105\b\r\1\2\u00ff\u0100\7\63\2\2\u0100\u0105\b\r\1\2\u0101\u0102\5"+
		"\32\16\2\u0102\u0103\b\r\1\2\u0103\u0105\3\2\2\2\u0104\u00f2\3\2\2\2\u0104"+
		"\u00f4\3\2\2\2\u0104\u00f6\3\2\2\2\u0104\u00f8\3\2\2\2\u0104\u00fa\3\2"+
		"\2\2\u0104\u00fd\3\2\2\2\u0104\u00ff\3\2\2\2\u0104\u0101\3\2\2\2\u0105"+
		"\31\3\2\2\2\u0106\u0107\b\16\1\2\u0107\u0108\7K\2\2\u0108\u0109\b\16\1"+
		"\2\u0109\u0112\3\2\2\2\u010a\u010b\f\4\2\2\u010b\u010c\7\f\2\2\u010c\u010d"+
		"\5\34\17\2\u010d\u010e\7\r\2\2\u010e\u010f\b\16\1\2\u010f\u0111\3\2\2"+
		"\2\u0110\u010a\3\2\2\2\u0111\u0114\3\2\2\2\u0112\u0110\3\2\2\2\u0112\u0113"+
		"\3\2\2\2\u0113\33\3\2\2\2\u0114\u0112\3\2\2\2\u0115\u0116\b\17\1\2\u0116"+
		"\u0117\7\30\2\2\u0117\u0118\5\34\17\32\u0118\u0119\b\17\1\2\u0119\u014e"+
		"\3\2\2\2\u011a\u011b\7%\2\2\u011b\u011c\7\24\2\2\u011c\u011d\5\34\17\2"+
		"\u011d\u011e\7\17\2\2\u011e\u011f\5\34\17\2\u011f\u0120\7\25\2\2\u0120"+
		"\u0121\b\17\1\2\u0121\u014e\3\2\2\2\u0122\u0123\7\'\2\2\u0123\u0124\7"+
		"\24\2\2\u0124\u0125\5\34\17\2\u0125\u0126\7\17\2\2\u0126\u0127\5\34\17"+
		"\2\u0127\u0128\7\25\2\2\u0128\u0129\b\17\1\2\u0129\u014e\3\2\2\2\u012a"+
		"\u012b\7\36\2\2\u012b\u012c\5\34\17\r\u012c\u012d\b\17\1\2\u012d\u014e"+
		"\3\2\2\2\u012e\u012f\7\24\2\2\u012f\u0130\5\34\17\2\u0130\u0131\7\25\2"+
		"\2\u0131\u0132\b\17\1\2\u0132\u014e\3\2\2\2\u0133\u0134\5\30\r\2\u0134"+
		"\u0135\7\60\2\2\u0135\u0136\7#\2\2\u0136\u0137\b\17\1\2\u0137\u014e\3"+
		"\2\2\2\u0138\u0139\5\30\r\2\u0139\u013a\7\60\2\2\u013a\u013b\7$\2\2\u013b"+
		"\u013c\b\17\1\2\u013c\u014e\3\2\2\2\u013d\u013e\7\f\2\2\u013e\u013f\5"+
		"\34\17\2\u013f\u0140\7\20\2\2\u0140\u0141\5\34\17\2\u0141\u0142\7\r\2"+
		"\2\u0142\u0143\b\17\1\2\u0143\u014e\3\2\2\2\u0144\u0145\7\f\2\2\u0145"+
		"\u0146\5\34\17\2\u0146\u0147\5 \21\2\u0147\u0148\7\r\2\2\u0148\u0149\b"+
		"\17\1\2\u0149\u014e\3\2\2\2\u014a\u014b\5\30\r\2\u014b\u014c\b\17\1\2"+
		"\u014c\u014e\3\2\2\2\u014d\u0115\3\2\2\2\u014d\u011a\3\2\2\2\u014d\u0122"+
		"\3\2\2\2\u014d\u012a\3\2\2\2\u014d\u012e\3\2\2\2\u014d\u0133\3\2\2\2\u014d"+
		"\u0138\3\2\2\2\u014d\u013d\3\2\2\2\u014d\u0144\3\2\2\2\u014d\u014a\3\2"+
		"\2\2\u014e\u019a\3\2\2\2\u014f\u0150\f\31\2\2\u0150\u0151\t\2\2\2\u0151"+
		"\u0152\5\34\17\32\u0152\u0153\b\17\1\2\u0153\u0199\3\2\2\2\u0154\u0155"+
		"\f\30\2\2\u0155\u0156\t\3\2\2\u0156\u0157\5\34\17\31\u0157\u0158\b\17"+
		"\1\2\u0158\u0199\3\2\2\2\u0159\u015a\f\25\2\2\u015a\u015b\7\21\2\2\u015b"+
		"\u015c\5\34\17\26\u015c\u015d\b\17\1\2\u015d\u0199\3\2\2\2\u015e\u015f"+
		"\f\24\2\2\u015f\u0160\7\22\2\2\u0160\u0161\5\34\17\25\u0161\u0162\b\17"+
		"\1\2\u0162\u0199\3\2\2\2\u0163\u0164\f\23\2\2\u0164\u0165\7\b\2\2\u0165"+
		"\u0166\5\34\17\24\u0166\u0167\b\17\1\2\u0167\u0199\3\2\2\2\u0168\u0169"+
		"\f\22\2\2\u0169\u016a\7\t\2\2\u016a\u016b\5\34\17\23\u016b\u016c\b\17"+
		"\1\2\u016c\u0199\3\2\2\2\u016d\u016e\f\21\2\2\u016e\u016f\7\6\2\2\u016f"+
		"\u0170\5\34\17\22\u0170\u0171\b\17\1\2\u0171\u0199\3\2\2\2\u0172\u0173"+
		"\f\20\2\2\u0173\u0174\7\7\2\2\u0174\u0175\5\34\17\21\u0175\u0176\b\17"+
		"\1\2\u0176\u0199\3\2\2\2\u0177\u0178\f\17\2\2\u0178\u0179\7\4\2\2\u0179"+
		"\u017a\5\34\17\20\u017a\u017b\b\17\1\2\u017b\u0199\3\2\2\2\u017c\u017d"+
		"\f\16\2\2\u017d\u017e\7\5\2\2\u017e\u017f\5\34\17\17\u017f\u0180\b\17"+
		"\1\2\u0180\u0199\3\2\2\2\u0181\u0182\f\t\2\2\u0182\u0183\7\23\2\2\u0183"+
		"\u0184\7\66\2\2\u0184\u0185\7\24\2\2\u0185\u0186\7\25\2\2\u0186\u0199"+
		"\b\17\1\2\u0187\u0188\f\b\2\2\u0188\u0189\7\23\2\2\u0189\u018a\7\67\2"+
		"\2\u018a\u018b\7\24\2\2\u018b\u018c\7\25\2\2\u018c\u0199\b\17\1\2\u018d"+
		"\u018e\f\7\2\2\u018e\u018f\7\23\2\2\u018f\u0190\78\2\2\u0190\u0191\7\24"+
		"\2\2\u0191\u0192\7\25\2\2\u0192\u0199\b\17\1\2\u0193\u0194\f\6\2\2\u0194"+
		"\u0195\7\23\2\2\u0195\u0196\79\2\2\u0196\u0197\7\24\2\2\u0197\u0199\7"+
		"\25\2\2\u0198\u014f\3\2\2\2\u0198\u0154\3\2\2\2\u0198\u0159\3\2\2\2\u0198"+
		"\u015e\3\2\2\2\u0198\u0163\3\2\2\2\u0198\u0168\3\2\2\2\u0198\u016d\3\2"+
		"\2\2\u0198\u0172\3\2\2\2\u0198\u0177\3\2\2\2\u0198\u017c\3\2\2\2\u0198"+
		"\u0181\3\2\2\2\u0198\u0187\3\2\2\2\u0198\u018d\3\2\2\2\u0198\u0193\3\2"+
		"\2\2\u0199\u019c\3\2\2\2\u019a\u0198\3\2\2\2\u019a\u019b\3\2\2\2\u019b"+
		"\35\3\2\2\2\u019c\u019a\3\2\2\2\u019d\u019e\7\61\2\2\u019e\u019f\7\24"+
		"\2\2\u019f\u01a0\5\34\17\2\u01a0\u01a1\7\25\2\2\u01a1\u01a2\7\20\2\2\u01a2"+
		"\u01a3\b\20\1\2\u01a3\u01ad\3\2\2\2\u01a4\u01a5\7\61\2\2\u01a5\u01a6\7"+
		"\24\2\2\u01a6\u01a7\5\34\17\2\u01a7\u01a8\5 \21\2\u01a8\u01a9\7\25\2\2"+
		"\u01a9\u01aa\7\20\2\2\u01aa\u01ab\b\20\1\2\u01ab\u01ad\3\2\2\2\u01ac\u019d"+
		"\3\2\2\2\u01ac\u01a4\3\2\2\2\u01ad\37\3\2\2\2\u01ae\u01b0\5\"\22\2\u01af"+
		"\u01ae\3\2\2\2\u01b0\u01b3\3\2\2\2\u01b1\u01af\3\2\2\2\u01b1\u01b2\3\2"+
		"\2\2\u01b2\u01b4\3\2\2\2\u01b3\u01b1\3\2\2\2\u01b4\u01b5\b\21\1\2\u01b5"+
		"!\3\2\2\2\u01b6\u01b7\7\17\2\2\u01b7\u01b8\5\34\17\2\u01b8\u01b9\b\22"+
		"\1\2\u01b9#\3\2\2\2\u01ba\u01bb\5&\24\2\u01bb\u01bc\b\23\1\2\u01bc%\3"+
		"\2\2\2\u01bd\u01be\7C\2\2\u01be\u01bf\5\34\17\2\u01bf\u01c0\5,\27\2\u01c0"+
		"\u01c1\b\24\1\2\u01c1\u01d8\3\2\2\2\u01c2\u01c3\7C\2\2\u01c3\u01c4\5\34"+
		"\17\2\u01c4\u01c5\5,\27\2\u01c5\u01c6\7E\2\2\u01c6\u01c7\5,\27\2\u01c7"+
		"\u01c8\b\24\1\2\u01c8\u01d8\3\2\2\2\u01c9\u01ca\7C\2\2\u01ca\u01cb\5\34"+
		"\17\2\u01cb\u01cc\5,\27\2\u01cc\u01cd\5*\26\2\u01cd\u01ce\7E\2\2\u01ce"+
		"\u01cf\5,\27\2\u01cf\u01d0\b\24\1\2\u01d0\u01d8\3\2\2\2\u01d1\u01d2\7"+
		"C\2\2\u01d2\u01d3\5\34\17\2\u01d3\u01d4\5,\27\2\u01d4\u01d5\5*\26\2\u01d5"+
		"\u01d6\b\24\1\2\u01d6\u01d8\3\2\2\2\u01d7\u01bd\3\2\2\2\u01d7\u01c2\3"+
		"\2\2\2\u01d7\u01c9\3\2\2\2\u01d7\u01d1\3\2\2\2\u01d8\'\3\2\2\2\u01d9\u01da"+
		"\7D\2\2\u01da\u01db\5\34\17\2\u01db\u01dc\5,\27\2\u01dc\u01dd\b\25\1\2"+
		"\u01dd)\3\2\2\2\u01de\u01e0\5(\25\2\u01df\u01de\3\2\2\2\u01e0\u01e3\3"+
		"\2\2\2\u01e1\u01df\3\2\2\2\u01e1\u01e2\3\2\2\2\u01e2\u01e4\3\2\2\2\u01e3"+
		"\u01e1\3\2\2\2\u01e4\u01e5\b\26\1\2\u01e5+\3\2\2\2\u01e6\u01e7\7\n\2\2"+
		"\u01e7\u01e8\5\n\6\2\u01e8\u01e9\7\13\2\2\u01e9\u01ea\b\27\1\2\u01ea\u01ef"+
		"\3\2\2\2\u01eb\u01ec\7\n\2\2\u01ec\u01ed\7\13\2\2\u01ed\u01ef\b\27\1\2"+
		"\u01ee\u01e6\3\2\2\2\u01ee\u01eb\3\2\2\2\u01ef-\3\2\2\2\u01f0\u01f1\5"+
		"\60\31\2\u01f1\u01f2\b\30\1\2\u01f2\u01f7\3\2\2\2\u01f3\u01f4\5\62\32"+
		"\2\u01f4\u01f5\b\30\1\2\u01f5\u01f7\3\2\2\2\u01f6\u01f0\3\2\2\2\u01f6"+
		"\u01f3\3\2\2\2\u01f7/\3\2\2\2\u01f8\u01f9\7F\2\2\u01f9\u01fa\5\34\17\2"+
		"\u01fa\u01fb\5,\27\2\u01fb\u01fc\b\31\1\2\u01fc\61\3\2\2\2\u01fd\u01fe"+
		"\7H\2\2\u01fe\u01ff\5,\27\2\u01ff\u0200\b\32\1\2\u0200\63\3\2\2\2\25\67"+
		">Vgv\u0093\u00dc\u00ea\u0104\u0112\u014d\u0198\u019a\u01ac\u01b1\u01d7"+
		"\u01e1\u01ee\u01f6";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}