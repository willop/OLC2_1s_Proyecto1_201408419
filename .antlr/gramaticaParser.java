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
		TKR_else=67, TK_entero=68, TK_decimal=69, TK_id=70, Letra=71, TK_cadena=72, 
		TK_caracter=73;
	public static final int
		RULE_start = 0, RULE_instrucciones = 1, RULE_instruccion = 2, RULE_declaracion = 3, 
		RULE_tipovariable = 4, RULE_identificadores = 5, RULE_valores = 6, RULE_expresion = 7, 
		RULE_impresion = 8, RULE_impresioncomas = 9, RULE_condicionales = 10, 
		RULE_funcionif = 11, RULE_funcionelseif = 12, RULE_listaelseif = 13, RULE_bloque = 14;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "instrucciones", "instruccion", "declaracion", "tipovariable", 
			"identificadores", "valores", "expresion", "impresion", "impresioncomas", 
			"condicionales", "funcionif", "funcionelseif", "listaelseif", "bloque"
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
			"'main'", "'if'", "'else if'", "'else'"
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
			"TK_entero", "TK_decimal", "TK_id", "Letra", "TK_cadena", "TK_caracter"
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
			setState(30);
			match(TKR_fn);
			setState(31);
			match(TKR_main);
			setState(32);
			match(TK_par_apertura);
			setState(33);
			match(TK_par_cierre);
			setState(34);
			match(TK_corchete_apertura);
			setState(35);
			((StartContext)_localctx).instrucciones = instrucciones();
			setState(36);
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
			setState(42);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 18)) & ~0x3f) == 0 && ((1L << (_la - 18)) & ((1L << (TK_par_apertura - 18)) | (1L << (TK_menos - 18)) | (1L << (TK_sig_admiracion - 18)) | (1L << (TKR_pow - 18)) | (1L << (TKR_powf - 18)) | (1L << (TKR_Str - 18)) | (1L << (TKR_let - 18)) | (1L << (TKR_println - 18)) | (1L << (TKR_true - 18)) | (1L << (TKR_false - 18)) | (1L << (TKR_if - 18)) | (1L << (TK_entero - 18)) | (1L << (TK_decimal - 18)) | (1L << (TK_id - 18)) | (1L << (TK_cadena - 18)) | (1L << (TK_caracter - 18)))) != 0)) {
				{
				{
				setState(39);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).e.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(44);
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
		public InstruccionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion; }
	}

	public final InstruccionContext instruccion() throws RecognitionException {
		InstruccionContext _localctx = new InstruccionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruccion);
		try {
			setState(62);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(47);
				((InstruccionContext)_localctx).expresion = expresion(0);
				fmt.Println("mensaje en instrucciones: ",((InstruccionContext)_localctx).expresion.exp)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(50);
				((InstruccionContext)_localctx).impresion = impresion();
				_localctx.inst = ((InstruccionContext)_localctx).impresion.inst
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(53);
				((InstruccionContext)_localctx).declaracion = declaracion();
				_localctx.inst = ((InstruccionContext)_localctx).declaracion.inst
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(56);
				((InstruccionContext)_localctx).identificadores = identificadores();
				_localctx.inst = ((InstruccionContext)_localctx).identificadores.inst
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(59);
				((InstruccionContext)_localctx).condicionales = condicionales();
				_localctx.inst = ((InstruccionContext)_localctx).condicionales.inst
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
			setState(98);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(64);
				match(TKR_let);
				setState(65);
				match(TKR_mut);
				setState(66);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(67);
				match(TK_dosPuntos);
				setState(68);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(69);
				match(TK_igual);
				setState(70);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(71);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar,((DeclaracionContext)_localctx).expresion.exp,false,false,true) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(74);
				match(TKR_let);
				setState(75);
				match(TKR_mut);
				setState(76);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(77);
				match(TK_igual);
				setState(78);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(79);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),Interfaces.SINTIPO,((DeclaracionContext)_localctx).expresion.exp,false,false,true) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(82);
				match(TKR_let);
				setState(83);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(84);
				match(TK_dosPuntos);
				setState(85);
				((DeclaracionContext)_localctx).tipovariable = tipovariable();
				setState(86);
				match(TK_igual);
				setState(87);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(88);
				match(TK_pcoma);
				_localctx.inst = Instruccion.NuevaDeclaracion((((DeclaracionContext)_localctx).idd!=null?((DeclaracionContext)_localctx).idd.getText():null),((DeclaracionContext)_localctx).tipovariable.tipovar,((DeclaracionContext)_localctx).expresion.exp,false,false,false) 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(91);
				match(TKR_let);
				setState(92);
				((DeclaracionContext)_localctx).idd = match(TK_id);
				setState(93);
				match(TK_igual);
				setState(94);
				((DeclaracionContext)_localctx).expresion = expresion(0);
				setState(95);
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
			setState(112);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TKR_numericos_enteros:
				enterOuterAlt(_localctx, 1);
				{
				setState(100);
				match(TKR_numericos_enteros);
				_localctx.tipovar = Interfaces.INTEGER
				}
				break;
			case TKR_numericos_flotantes:
				enterOuterAlt(_localctx, 2);
				{
				setState(102);
				match(TKR_numericos_flotantes);
				_localctx.tipovar = Interfaces.FLOAT
				}
				break;
			case TKR_String:
				enterOuterAlt(_localctx, 3);
				{
				setState(104);
				match(TKR_String);
				_localctx.tipovar = Interfaces.STRING
				}
				break;
			case TKR_bool:
				enterOuterAlt(_localctx, 4);
				{
				setState(106);
				match(TKR_bool);
				_localctx.tipovar = Interfaces.BOOLEAN
				}
				break;
			case TKR_char:
				enterOuterAlt(_localctx, 5);
				{
				setState(108);
				match(TKR_char);
				_localctx.tipovar = Interfaces.CHAR
				}
				break;
			case TKR_usize:
				enterOuterAlt(_localctx, 6);
				{
				setState(110);
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
		enterRule(_localctx, 10, RULE_identificadores);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(114);
			((IdentificadoresContext)_localctx).e1 = match(TK_id);
			setState(115);
			match(TK_igual);
			setState(116);
			((IdentificadoresContext)_localctx).e2 = expresion(0);
			setState(117);
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
		enterRule(_localctx, 12, RULE_valores);
		try {
			setState(136);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_entero:
				enterOuterAlt(_localctx, 1);
				{
				setState(120);
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
				setState(122);
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
				setState(124);
				((ValoresContext)_localctx).TK_cadena = match(TK_cadena);

				                str:= (((ValoresContext)_localctx).TK_cadena!=null?((ValoresContext)_localctx).TK_cadena.getText():null)[1:len((((ValoresContext)_localctx).TK_cadena!=null?((ValoresContext)_localctx).TK_cadena.getText():null))-1]
				                _localctx.exp = Expresion.NuevoPrimitivo(str,Interfaces.STRING)
				}
				break;
			case TK_caracter:
				enterOuterAlt(_localctx, 4);
				{
				setState(126);
				((ValoresContext)_localctx).TK_caracter = match(TK_caracter);
				str:= (((ValoresContext)_localctx).TK_caracter!=null?((ValoresContext)_localctx).TK_caracter.getText():null)[1:len((((ValoresContext)_localctx).TK_caracter!=null?((ValoresContext)_localctx).TK_caracter.getText():null))-1]
				                _localctx.exp = Expresion.NuevoPrimitivo(str,Interfaces.CHAR)
				}
				break;
			case TKR_Str:
				enterOuterAlt(_localctx, 5);
				{
				setState(128);
				((ValoresContext)_localctx).TKR_Str = match(TKR_Str);
				str:= (((ValoresContext)_localctx).TKR_Str!=null?((ValoresContext)_localctx).TKR_Str.getText():null)[1:len((((ValoresContext)_localctx).TKR_Str!=null?((ValoresContext)_localctx).TKR_Str.getText():null))-1]
				                _localctx.exp = Expresion.NuevoPrimitivo(str,Interfaces.STR)
				}
				break;
			case TKR_true:
				enterOuterAlt(_localctx, 6);
				{
				setState(130);
				match(TKR_true);
				_localctx.exp = Expresion.NuevoPrimitivo(true, Interfaces.BOOLEAN)
				}
				break;
			case TKR_false:
				enterOuterAlt(_localctx, 7);
				{
				setState(132);
				match(TKR_false);
				_localctx.exp = Expresion.NuevoPrimitivo(false,Interfaces.BOOLEAN)
				}
				break;
			case TK_id:
				enterOuterAlt(_localctx, 8);
				{
				setState(134);
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
		int _startState = 14;
		enterRecursionRule(_localctx, 14, RULE_expresion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(178);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				{
				setState(139);
				match(TK_menos);
				setState(140);
				((ExpresionContext)_localctx).e1 = expresion(18);
				numero := -1
				                                                                                                        numm := Expresion.NuevoPrimitivo(numero, Interfaces.INTEGER)
				                                                                                                                _localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp, numm ,Interfaces.MULTIPLICACION)
				                                                                                                
				}
				break;
			case 2:
				{
				setState(143);
				match(TKR_pow);
				setState(144);
				match(TK_par_apertura);
				setState(145);
				((ExpresionContext)_localctx).e1 = expresion(0);
				setState(146);
				match(TK_coma);
				setState(147);
				((ExpresionContext)_localctx).e2 = expresion(0);
				setState(148);
				match(TK_par_cierre);
				_localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.POW)
				}
				break;
			case 3:
				{
				setState(151);
				match(TKR_powf);
				setState(152);
				match(TK_par_apertura);
				setState(153);
				((ExpresionContext)_localctx).e1 = expresion(0);
				setState(154);
				match(TK_coma);
				setState(155);
				((ExpresionContext)_localctx).e2 = expresion(0);
				setState(156);
				match(TK_par_cierre);
				_localctx.exp = Expresiones.NuevaAritmetica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.POWF)
				}
				break;
			case 4:
				{
				setState(159);
				match(TK_sig_admiracion);
				setState(160);
				((ExpresionContext)_localctx).e1 = expresion(5);
				_localctx.exp = Expresiones.NuevaLogica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e1.exp,Interfaces.NOT)
				}
				break;
			case 5:
				{
				setState(163);
				match(TK_par_apertura);
				setState(164);
				expresion(0);
				setState(165);
				match(TK_par_cierre);
				}
				break;
			case 6:
				{
				setState(167);
				valores();
				setState(168);
				match(TKR_as);
				setState(169);
				match(TKR_numericos_enteros);
				}
				break;
			case 7:
				{
				setState(171);
				valores();
				setState(172);
				match(TKR_as);
				setState(173);
				match(TKR_numericos_flotantes);
				}
				break;
			case 8:
				{
				setState(175);
				((ExpresionContext)_localctx).vall = valores();
				_localctx.exp = ((ExpresionContext)_localctx).vall.exp
				                                                                                        fmt.Println(_localctx.exp)
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(232);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(230);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
					case 1:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(180);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(181);
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
						setState(182);
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
						setState(185);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(186);
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
						setState(187);
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
						setState(190);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(191);
						match(TK_menor);
						setState(192);
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
						setState(195);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(196);
						match(TK_mayor);
						setState(197);
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
						setState(200);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(201);
						match(TK_mayor_igual);
						setState(202);
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
						setState(205);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(206);
						match(TK_menor_igual);
						setState(207);
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
						setState(210);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(211);
						match(TK_igualacion);
						setState(212);
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
						setState(215);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(216);
						match(TK_diferente);
						setState(217);
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
						setState(220);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(221);
						match(TK_or);
						setState(222);
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
						setState(225);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(226);
						match(TK_and);
						setState(227);
						((ExpresionContext)_localctx).e2 = expresion(7);
						_localctx.exp = Expresiones.NuevaLogica(((ExpresionContext)_localctx).e1.exp,((ExpresionContext)_localctx).e2.exp,Interfaces.AND)
						}
						break;
					}
					} 
				}
				setState(234);
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
		public TerminalNode TKR_println() { return getToken(gramaticaParser.TKR_println, 0); }
		public TerminalNode TK_par_apertura() { return getToken(gramaticaParser.TK_par_apertura, 0); }
		public TerminalNode TK_par_cierre() { return getToken(gramaticaParser.TK_par_cierre, 0); }
		public TerminalNode TK_pcoma() { return getToken(gramaticaParser.TK_pcoma, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
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
			setState(248);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(235);
				match(TKR_println);
				setState(236);
				match(TK_par_apertura);
				setState(237);
				((ImpresionContext)_localctx).e1 = expresion(0);
				setState(238);
				match(TK_par_cierre);
				setState(239);
				match(TK_pcoma);
				fmt.Println("Impresion")
				                                                                                                        _localctx.inst = Instruccion.NuevoPrint(((ImpresionContext)_localctx).e1.exp)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(242);
				match(TKR_println);
				setState(243);
				match(TK_par_apertura);
				setState(244);
				expresion(0);
				setState(245);
				impresioncomas(0);
				setState(246);
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
			setState(251);
			match(TK_coma);
			setState(252);
			expresion(0);
			setState(253);
			match(TK_par_cierre);
			setState(254);
			match(TK_pcoma);
			fmt.Println("Impresion especial")
			}
			_ctx.stop = _input.LT(-1);
			setState(266);
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
					setState(257);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(258);
					match(TK_coma);
					setState(259);
					expresion(0);
					setState(260);
					match(TK_par_cierre);
					setState(261);
					match(TK_pcoma);
					fmt.Println("Impresion especial")
					}
					} 
				}
				setState(268);
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
		enterRule(_localctx, 20, RULE_condicionales);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(269);
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
		enterRule(_localctx, 22, RULE_funcionif);
		try {
			setState(292);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(272);
				match(TKR_if);
				setState(273);
				((FuncionifContext)_localctx).e1 = expresion(0);
				setState(274);
				((FuncionifContext)_localctx).ee = bloque();
				_localctx.inst = Instruccion.NewIf(((FuncionifContext)_localctx).e1.exp,((FuncionifContext)_localctx).ee.lista,nil,nil )
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(277);
				match(TKR_if);
				setState(278);
				((FuncionifContext)_localctx).e1 = expresion(0);
				setState(279);
				((FuncionifContext)_localctx).e5 = bloque();
				setState(280);
				match(TKR_else);
				setState(281);
				((FuncionifContext)_localctx).b2 = bloque();
				_localctx.inst = Instruccion.NewIf(((FuncionifContext)_localctx).e1.exp,((FuncionifContext)_localctx).e5.lista,nil,((FuncionifContext)_localctx).b2.lista)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(284);
				match(TKR_if);
				setState(285);
				((FuncionifContext)_localctx).e1 = expresion(0);
				setState(286);
				((FuncionifContext)_localctx).b1 = bloque();
				setState(287);
				((FuncionifContext)_localctx).listaelseif = listaelseif();
				setState(288);
				match(TKR_else);
				setState(289);
				((FuncionifContext)_localctx).b2 = bloque();
				_localctx.inst = Instruccion.NewIf(((FuncionifContext)_localctx).e1.exp,((FuncionifContext)_localctx).b1.lista,((FuncionifContext)_localctx).listaelseif.lista,((FuncionifContext)_localctx).b2.lista)
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
		enterRule(_localctx, 24, RULE_funcionelseif);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(294);
			match(TKR_elseif);
			setState(295);
			((FuncionelseifContext)_localctx).e1 = expresion(0);
			setState(296);
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
		public FuncionelseifContext funcionelseif() {
			return getRuleContext(FuncionelseifContext.class,0);
		}
		public ListaelseifContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listaelseif; }
	}

	public final ListaelseifContext listaelseif() throws RecognitionException {
		ListaelseifContext _localctx = new ListaelseifContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_listaelseif);
		 _localctx.lista = arrayList.New()
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(299);
			((ListaelseifContext)_localctx).funcionelseif = funcionelseif();
			((ListaelseifContext)_localctx).list.add(((ListaelseifContext)_localctx).funcionelseif);

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
		enterRule(_localctx, 28, RULE_bloque);
		try {
			setState(310);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(302);
				match(TK_corchete_apertura);
				setState(303);
				((BloqueContext)_localctx).instrucciones = instrucciones();
				setState(304);
				match(TK_corchete_cierre);
				_localctx.lista = ((BloqueContext)_localctx).instrucciones.lista
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(307);
				match(TK_corchete_apertura);
				setState(308);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3K\u013b\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\3\2\3\2\3\2\3\2\3\2"+
		"\3\2\3\2\3\2\3\2\3\3\7\3+\n\3\f\3\16\3.\13\3\3\3\3\3\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4A\n\4\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5"+
		"\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\5\5e\n\5\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6s\n\6\3\7\3\7\3\7\3\7\3\7\3\7"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\5\b\u008b"+
		"\n\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\5\t\u00b5\n\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\3\t\7\t\u00e9\n\t\f\t\16\t\u00ec\13\t\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n\u00fb\n\n\3\13\3\13\3"+
		"\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\7\13\u010b"+
		"\n\13\f\13\16\13\u010e\13\13\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3"+
		"\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\5\r\u0127\n\r\3\16"+
		"\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\5\20\u0139\n\20\3\20\2\4\20\24\21\2\4\6\b\n\f\16\20\22\24\26\30"+
		"\32\34\36\2\4\3\2\31\33\3\2\27\30\2\u0155\2 \3\2\2\2\4,\3\2\2\2\6@\3\2"+
		"\2\2\bd\3\2\2\2\nr\3\2\2\2\ft\3\2\2\2\16\u008a\3\2\2\2\20\u00b4\3\2\2"+
		"\2\22\u00fa\3\2\2\2\24\u00fc\3\2\2\2\26\u010f\3\2\2\2\30\u0126\3\2\2\2"+
		"\32\u0128\3\2\2\2\34\u012d\3\2\2\2\36\u0138\3\2\2\2 !\7\64\2\2!\"\7B\2"+
		"\2\"#\7\24\2\2#$\7\25\2\2$%\7\n\2\2%&\5\4\3\2&\'\7\13\2\2\'(\b\2\1\2("+
		"\3\3\2\2\2)+\5\6\4\2*)\3\2\2\2+.\3\2\2\2,*\3\2\2\2,-\3\2\2\2-/\3\2\2\2"+
		".,\3\2\2\2/\60\b\3\1\2\60\5\3\2\2\2\61\62\5\20\t\2\62\63\b\4\1\2\63A\3"+
		"\2\2\2\64\65\5\22\n\2\65\66\b\4\1\2\66A\3\2\2\2\678\5\b\5\289\b\4\1\2"+
		"9A\3\2\2\2:;\5\f\7\2;<\b\4\1\2<A\3\2\2\2=>\5\26\f\2>?\b\4\1\2?A\3\2\2"+
		"\2@\61\3\2\2\2@\64\3\2\2\2@\67\3\2\2\2@:\3\2\2\2@=\3\2\2\2A\7\3\2\2\2"+
		"BC\7-\2\2CD\7.\2\2DE\7H\2\2EF\7\16\2\2FG\5\n\6\2GH\7\26\2\2HI\5\20\t\2"+
		"IJ\7\20\2\2JK\b\5\1\2Ke\3\2\2\2LM\7-\2\2MN\7.\2\2NO\7H\2\2OP\7\26\2\2"+
		"PQ\5\20\t\2QR\7\20\2\2RS\b\5\1\2Se\3\2\2\2TU\7-\2\2UV\7H\2\2VW\7\16\2"+
		"\2WX\5\n\6\2XY\7\26\2\2YZ\5\20\t\2Z[\7\20\2\2[\\\b\5\1\2\\e\3\2\2\2]^"+
		"\7-\2\2^_\7H\2\2_`\7\26\2\2`a\5\20\t\2ab\7\20\2\2bc\b\5\1\2ce\3\2\2\2"+
		"dB\3\2\2\2dL\3\2\2\2dT\3\2\2\2d]\3\2\2\2e\t\3\2\2\2fg\7#\2\2gs\b\6\1\2"+
		"hi\7$\2\2is\b\6\1\2jk\7+\2\2ks\b\6\1\2lm\7(\2\2ms\b\6\1\2no\7)\2\2os\b"+
		"\6\1\2pq\7,\2\2qs\b\6\1\2rf\3\2\2\2rh\3\2\2\2rj\3\2\2\2rl\3\2\2\2rn\3"+
		"\2\2\2rp\3\2\2\2s\13\3\2\2\2tu\7H\2\2uv\7\26\2\2vw\5\20\t\2wx\7\20\2\2"+
		"xy\b\7\1\2y\r\3\2\2\2z{\7F\2\2{\u008b\b\b\1\2|}\7G\2\2}\u008b\b\b\1\2"+
		"~\177\7J\2\2\177\u008b\b\b\1\2\u0080\u0081\7K\2\2\u0081\u008b\b\b\1\2"+
		"\u0082\u0083\7*\2\2\u0083\u008b\b\b\1\2\u0084\u0085\7\62\2\2\u0085\u008b"+
		"\b\b\1\2\u0086\u0087\7\63\2\2\u0087\u008b\b\b\1\2\u0088\u0089\7H\2\2\u0089"+
		"\u008b\b\b\1\2\u008az\3\2\2\2\u008a|\3\2\2\2\u008a~\3\2\2\2\u008a\u0080"+
		"\3\2\2\2\u008a\u0082\3\2\2\2\u008a\u0084\3\2\2\2\u008a\u0086\3\2\2\2\u008a"+
		"\u0088\3\2\2\2\u008b\17\3\2\2\2\u008c\u008d\b\t\1\2\u008d\u008e\7\30\2"+
		"\2\u008e\u008f\5\20\t\24\u008f\u0090\b\t\1\2\u0090\u00b5\3\2\2\2\u0091"+
		"\u0092\7%\2\2\u0092\u0093\7\24\2\2\u0093\u0094\5\20\t\2\u0094\u0095\7"+
		"\17\2\2\u0095\u0096\5\20\t\2\u0096\u0097\7\25\2\2\u0097\u0098\b\t\1\2"+
		"\u0098\u00b5\3\2\2\2\u0099\u009a\7\'\2\2\u009a\u009b\7\24\2\2\u009b\u009c"+
		"\5\20\t\2\u009c\u009d\7\17\2\2\u009d\u009e\5\20\t\2\u009e\u009f\7\25\2"+
		"\2\u009f\u00a0\b\t\1\2\u00a0\u00b5\3\2\2\2\u00a1\u00a2\7\36\2\2\u00a2"+
		"\u00a3\5\20\t\7\u00a3\u00a4\b\t\1\2\u00a4\u00b5\3\2\2\2\u00a5\u00a6\7"+
		"\24\2\2\u00a6\u00a7\5\20\t\2\u00a7\u00a8\7\25\2\2\u00a8\u00b5\3\2\2\2"+
		"\u00a9\u00aa\5\16\b\2\u00aa\u00ab\7\60\2\2\u00ab\u00ac\7#\2\2\u00ac\u00b5"+
		"\3\2\2\2\u00ad\u00ae\5\16\b\2\u00ae\u00af\7\60\2\2\u00af\u00b0\7$\2\2"+
		"\u00b0\u00b5\3\2\2\2\u00b1\u00b2\5\16\b\2\u00b2\u00b3\b\t\1\2\u00b3\u00b5"+
		"\3\2\2\2\u00b4\u008c\3\2\2\2\u00b4\u0091\3\2\2\2\u00b4\u0099\3\2\2\2\u00b4"+
		"\u00a1\3\2\2\2\u00b4\u00a5\3\2\2\2\u00b4\u00a9\3\2\2\2\u00b4\u00ad\3\2"+
		"\2\2\u00b4\u00b1\3\2\2\2\u00b5\u00ea\3\2\2\2\u00b6\u00b7\f\23\2\2\u00b7"+
		"\u00b8\t\2\2\2\u00b8\u00b9\5\20\t\24\u00b9\u00ba\b\t\1\2\u00ba\u00e9\3"+
		"\2\2\2\u00bb\u00bc\f\22\2\2\u00bc\u00bd\t\3\2\2\u00bd\u00be\5\20\t\23"+
		"\u00be\u00bf\b\t\1\2\u00bf\u00e9\3\2\2\2\u00c0\u00c1\f\17\2\2\u00c1\u00c2"+
		"\7\21\2\2\u00c2\u00c3\5\20\t\20\u00c3\u00c4\b\t\1\2\u00c4\u00e9\3\2\2"+
		"\2\u00c5\u00c6\f\16\2\2\u00c6\u00c7\7\22\2\2\u00c7\u00c8\5\20\t\17\u00c8"+
		"\u00c9\b\t\1\2\u00c9\u00e9\3\2\2\2\u00ca\u00cb\f\r\2\2\u00cb\u00cc\7\b"+
		"\2\2\u00cc\u00cd\5\20\t\16\u00cd\u00ce\b\t\1\2\u00ce\u00e9\3\2\2\2\u00cf"+
		"\u00d0\f\f\2\2\u00d0\u00d1\7\t\2\2\u00d1\u00d2\5\20\t\r\u00d2\u00d3\b"+
		"\t\1\2\u00d3\u00e9\3\2\2\2\u00d4\u00d5\f\13\2\2\u00d5\u00d6\7\6\2\2\u00d6"+
		"\u00d7\5\20\t\f\u00d7\u00d8\b\t\1\2\u00d8\u00e9\3\2\2\2\u00d9\u00da\f"+
		"\n\2\2\u00da\u00db\7\7\2\2\u00db\u00dc\5\20\t\13\u00dc\u00dd\b\t\1\2\u00dd"+
		"\u00e9\3\2\2\2\u00de\u00df\f\t\2\2\u00df\u00e0\7\4\2\2\u00e0\u00e1\5\20"+
		"\t\n\u00e1\u00e2\b\t\1\2\u00e2\u00e9\3\2\2\2\u00e3\u00e4\f\b\2\2\u00e4"+
		"\u00e5\7\5\2\2\u00e5\u00e6\5\20\t\t\u00e6\u00e7\b\t\1\2\u00e7\u00e9\3"+
		"\2\2\2\u00e8\u00b6\3\2\2\2\u00e8\u00bb\3\2\2\2\u00e8\u00c0\3\2\2\2\u00e8"+
		"\u00c5\3\2\2\2\u00e8\u00ca\3\2\2\2\u00e8\u00cf\3\2\2\2\u00e8\u00d4\3\2"+
		"\2\2\u00e8\u00d9\3\2\2\2\u00e8\u00de\3\2\2\2\u00e8\u00e3\3\2\2\2\u00e9"+
		"\u00ec\3\2\2\2\u00ea\u00e8\3\2\2\2\u00ea\u00eb\3\2\2\2\u00eb\21\3\2\2"+
		"\2\u00ec\u00ea\3\2\2\2\u00ed\u00ee\7\61\2\2\u00ee\u00ef\7\24\2\2\u00ef"+
		"\u00f0\5\20\t\2\u00f0\u00f1\7\25\2\2\u00f1\u00f2\7\20\2\2\u00f2\u00f3"+
		"\b\n\1\2\u00f3\u00fb\3\2\2\2\u00f4\u00f5\7\61\2\2\u00f5\u00f6\7\24\2\2"+
		"\u00f6\u00f7\5\20\t\2\u00f7\u00f8\5\24\13\2\u00f8\u00f9\7\20\2\2\u00f9"+
		"\u00fb\3\2\2\2\u00fa\u00ed\3\2\2\2\u00fa\u00f4\3\2\2\2\u00fb\23\3\2\2"+
		"\2\u00fc\u00fd\b\13\1\2\u00fd\u00fe\7\17\2\2\u00fe\u00ff\5\20\t\2\u00ff"+
		"\u0100\7\25\2\2\u0100\u0101\7\20\2\2\u0101\u0102\b\13\1\2\u0102\u010c"+
		"\3\2\2\2\u0103\u0104\f\4\2\2\u0104\u0105\7\17\2\2\u0105\u0106\5\20\t\2"+
		"\u0106\u0107\7\25\2\2\u0107\u0108\7\20\2\2\u0108\u0109\b\13\1\2\u0109"+
		"\u010b\3\2\2\2\u010a\u0103\3\2\2\2\u010b\u010e\3\2\2\2\u010c\u010a\3\2"+
		"\2\2\u010c\u010d\3\2\2\2\u010d\25\3\2\2\2\u010e\u010c\3\2\2\2\u010f\u0110"+
		"\5\30\r\2\u0110\u0111\b\f\1\2\u0111\27\3\2\2\2\u0112\u0113\7C\2\2\u0113"+
		"\u0114\5\20\t\2\u0114\u0115\5\36\20\2\u0115\u0116\b\r\1\2\u0116\u0127"+
		"\3\2\2\2\u0117\u0118\7C\2\2\u0118\u0119\5\20\t\2\u0119\u011a\5\36\20\2"+
		"\u011a\u011b\7E\2\2\u011b\u011c\5\36\20\2\u011c\u011d\b\r\1\2\u011d\u0127"+
		"\3\2\2\2\u011e\u011f\7C\2\2\u011f\u0120\5\20\t\2\u0120\u0121\5\36\20\2"+
		"\u0121\u0122\5\34\17\2\u0122\u0123\7E\2\2\u0123\u0124\5\36\20\2\u0124"+
		"\u0125\b\r\1\2\u0125\u0127\3\2\2\2\u0126\u0112\3\2\2\2\u0126\u0117\3\2"+
		"\2\2\u0126\u011e\3\2\2\2\u0127\31\3\2\2\2\u0128\u0129\7D\2\2\u0129\u012a"+
		"\5\20\t\2\u012a\u012b\5\36\20\2\u012b\u012c\b\16\1\2\u012c\33\3\2\2\2"+
		"\u012d\u012e\5\32\16\2\u012e\u012f\b\17\1\2\u012f\35\3\2\2\2\u0130\u0131"+
		"\7\n\2\2\u0131\u0132\5\4\3\2\u0132\u0133\7\13\2\2\u0133\u0134\b\20\1\2"+
		"\u0134\u0139\3\2\2\2\u0135\u0136\7\n\2\2\u0136\u0137\7\13\2\2\u0137\u0139"+
		"\b\20\1\2\u0138\u0130\3\2\2\2\u0138\u0135\3\2\2\2\u0139\37\3\2\2\2\16"+
		",@dr\u008a\u00b4\u00e8\u00ea\u00fa\u010c\u0126\u0138";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}