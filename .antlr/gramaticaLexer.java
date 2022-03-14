// Generated from e:\WillOP\u005Cu\Compi2\Laboratorio\Proyecto1_201408419\gramatica.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class gramaticaLexer extends Lexer {
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
		TKR_else=67, TKR_while=68, TKR_break=69, TK_entero=70, TK_decimal=71, 
		TK_id=72, Letra=73, TK_cadena=74, TK_caracter=75;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"TK_flecha", "TK_or", "TK_and", "TK_igualacion", "TK_diferente", "TK_mayor_igual", 
			"TK_menor_igual", "TK_corchete_apertura", "TK_corchete_cierre", "TK_llave_apertura", 
			"TK_llave_cierre", "TK_dosPuntos", "TK_coma", "TK_pcoma", "TK_menor", 
			"TK_mayor", "TK_punto", "TK_par_apertura", "TK_par_cierre", "TK_igual", 
			"TK_suma", "TK_menos", "TK_por", "TK_diagonal", "TK_porcentaje", "TK_linea", 
			"TK_amp", "TK_sig_admiracion", "TK_sig_interrogacion", "WHITESPACE", 
			"TK_comentario_multi", "TK_comentario_lineal", "TKR_numericos_enteros", 
			"TKR_numericos_flotantes", "TKR_pow", "TKR_vec", "TKR_powf", "TKR_bool", 
			"TKR_char", "TKR_Str", "TKR_String", "TKR_usize", "TKR_let", "TKR_mut", 
			"TKR_struct", "TKR_as", "TKR_println", "TKR_true", "TKR_false", "TKR_fn", 
			"TKR_return", "TKR_abs", "TKR_sqrt", "TKR_to_string", "TKR_clone", "TKR_new", 
			"TKR_len", "TKR_push", "TKR_remove", "TKR_contains", "TKR_insert", "TKR_capacity", 
			"TKR_with_capacity", "TKR_main", "TKR_if", "TKR_elseif", "TKR_else", 
			"TKR_while", "TKR_break", "TK_entero", "TK_decimal", "TK_id", "Letra", 
			"TK_cadena", "TK_caracter"
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
			"'main'", "'if'", "'else if'", "'else'", "'while'", "'break'"
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
			"TKR_while", "TKR_break", "TK_entero", "TK_decimal", "TK_id", "Letra", 
			"TK_cadena", "TK_caracter"
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


	public gramaticaLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "gramatica.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2M\u0200\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\4J\tJ\4K\tK\4L\tL\3\2\3\2\3\2\3\3\3\3\3\3\3\4\3\4\3\4\3\5\3\5\3\5"+
		"\3\6\3\6\3\6\3\7\3\7\3\7\3\b\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f"+
		"\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3"+
		"\24\3\24\3\25\3\25\3\26\3\26\3\27\3\27\3\30\3\30\3\31\3\31\3\32\3\32\3"+
		"\33\3\33\3\34\3\34\3\35\3\35\3\36\3\36\3\37\6\37\u00dc\n\37\r\37\16\37"+
		"\u00dd\3\37\3\37\3 \3 \3 \3 \7 \u00e6\n \f \16 \u00e9\13 \3 \3 \3 \3 "+
		"\3 \3!\3!\3!\3!\7!\u00f4\n!\f!\16!\u00f7\13!\3!\3!\3\"\3\"\3\"\3\"\3#"+
		"\3#\3#\3#\3$\3$\3$\3$\3%\3%\3%\3%\3&\3&\3&\3&\3&\3\'\3\'\3\'\3\'\3\'\3"+
		"(\3(\3(\3(\3(\3)\3)\3)\3)\3)\3*\3*\3*\3*\3*\3*\3*\3+\3+\3+\3+\3+\3+\3"+
		",\3,\3,\3,\3-\3-\3-\3-\3.\3.\3.\3.\3.\3.\3.\3/\3/\3/\3\60\3\60\3\60\3"+
		"\60\3\60\3\60\3\60\3\60\3\60\3\61\3\61\3\61\3\61\3\61\3\62\3\62\3\62\3"+
		"\62\3\62\3\62\3\63\3\63\3\63\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\65\3"+
		"\65\3\65\3\65\3\66\3\66\3\66\3\66\3\66\3\67\3\67\3\67\3\67\3\67\3\67\3"+
		"\67\3\67\3\67\3\67\3\67\3\67\38\38\38\38\38\38\38\38\39\39\39\39\3:\3"+
		":\3:\3:\3;\3;\3;\3;\3;\3<\3<\3<\3<\3<\3<\3<\3=\3=\3=\3=\3=\3=\3=\3=\3"+
		"=\3>\3>\3>\3>\3>\3>\3>\3?\3?\3?\3?\3?\3?\3?\3?\3?\3@\3@\3@\3@\3@\3@\3"+
		"@\3@\3@\3@\3@\3@\3@\3@\3@\3A\3A\3A\3A\3A\3B\3B\3B\3C\3C\3C\3C\3C\3C\3"+
		"C\3C\3D\3D\3D\3D\3D\3E\3E\3E\3E\3E\3E\3F\3F\3F\3F\3F\3F\3G\6G\u01d7\n"+
		"G\rG\16G\u01d8\3H\3H\3H\6H\u01de\nH\rH\16H\u01df\3H\5H\u01e3\nH\3I\3I"+
		"\3I\7I\u01e8\nI\fI\16I\u01eb\13I\3J\3J\3K\3K\7K\u01f1\nK\fK\16K\u01f4"+
		"\13K\3K\3K\3L\3L\7L\u01fa\nL\fL\16L\u01fd\13L\3L\3L\3\u00e7\2M\3\3\5\4"+
		"\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22"+
		"#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C"+
		"#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67m8o9q:s;u<w"+
		"=y>{?}@\177A\u0081B\u0083C\u0085D\u0087E\u0089F\u008bG\u008dH\u008fI\u0091"+
		"J\u0093K\u0095L\u0097M\3\2\7\5\2\13\f\17\17\"\"\4\2\f\f\17\17\3\2\62;"+
		"\5\2C\\aac|\3\2$$\2\u0209\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2"+
		"\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2"+
		"\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3"+
		"\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2"+
		"\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67"+
		"\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2"+
		"\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2"+
		"\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]"+
		"\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2"+
		"\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2"+
		"\2w\3\2\2\2\2y\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2\2\u0081\3\2"+
		"\2\2\2\u0083\3\2\2\2\2\u0085\3\2\2\2\2\u0087\3\2\2\2\2\u0089\3\2\2\2\2"+
		"\u008b\3\2\2\2\2\u008d\3\2\2\2\2\u008f\3\2\2\2\2\u0091\3\2\2\2\2\u0093"+
		"\3\2\2\2\2\u0095\3\2\2\2\2\u0097\3\2\2\2\3\u0099\3\2\2\2\5\u009c\3\2\2"+
		"\2\7\u009f\3\2\2\2\t\u00a2\3\2\2\2\13\u00a5\3\2\2\2\r\u00a8\3\2\2\2\17"+
		"\u00ab\3\2\2\2\21\u00ae\3\2\2\2\23\u00b0\3\2\2\2\25\u00b2\3\2\2\2\27\u00b4"+
		"\3\2\2\2\31\u00b6\3\2\2\2\33\u00b8\3\2\2\2\35\u00ba\3\2\2\2\37\u00bc\3"+
		"\2\2\2!\u00be\3\2\2\2#\u00c0\3\2\2\2%\u00c2\3\2\2\2\'\u00c4\3\2\2\2)\u00c6"+
		"\3\2\2\2+\u00c8\3\2\2\2-\u00ca\3\2\2\2/\u00cc\3\2\2\2\61\u00ce\3\2\2\2"+
		"\63\u00d0\3\2\2\2\65\u00d2\3\2\2\2\67\u00d4\3\2\2\29\u00d6\3\2\2\2;\u00d8"+
		"\3\2\2\2=\u00db\3\2\2\2?\u00e1\3\2\2\2A\u00ef\3\2\2\2C\u00fa\3\2\2\2E"+
		"\u00fe\3\2\2\2G\u0102\3\2\2\2I\u0106\3\2\2\2K\u010a\3\2\2\2M\u010f\3\2"+
		"\2\2O\u0114\3\2\2\2Q\u0119\3\2\2\2S\u011e\3\2\2\2U\u0125\3\2\2\2W\u012b"+
		"\3\2\2\2Y\u012f\3\2\2\2[\u0133\3\2\2\2]\u013a\3\2\2\2_\u013d\3\2\2\2a"+
		"\u0146\3\2\2\2c\u014b\3\2\2\2e\u0151\3\2\2\2g\u0154\3\2\2\2i\u015b\3\2"+
		"\2\2k\u015f\3\2\2\2m\u0164\3\2\2\2o\u0170\3\2\2\2q\u0178\3\2\2\2s\u017c"+
		"\3\2\2\2u\u0180\3\2\2\2w\u0185\3\2\2\2y\u018c\3\2\2\2{\u0195\3\2\2\2}"+
		"\u019c\3\2\2\2\177\u01a5\3\2\2\2\u0081\u01b4\3\2\2\2\u0083\u01b9\3\2\2"+
		"\2\u0085\u01bc\3\2\2\2\u0087\u01c4\3\2\2\2\u0089\u01c9\3\2\2\2\u008b\u01cf"+
		"\3\2\2\2\u008d\u01d6\3\2\2\2\u008f\u01da\3\2\2\2\u0091\u01e4\3\2\2\2\u0093"+
		"\u01ec\3\2\2\2\u0095\u01ee\3\2\2\2\u0097\u01f7\3\2\2\2\u0099\u009a\7/"+
		"\2\2\u009a\u009b\7@\2\2\u009b\4\3\2\2\2\u009c\u009d\7~\2\2\u009d\u009e"+
		"\7~\2\2\u009e\6\3\2\2\2\u009f\u00a0\7(\2\2\u00a0\u00a1\7(\2\2\u00a1\b"+
		"\3\2\2\2\u00a2\u00a3\7?\2\2\u00a3\u00a4\7?\2\2\u00a4\n\3\2\2\2\u00a5\u00a6"+
		"\7#\2\2\u00a6\u00a7\7?\2\2\u00a7\f\3\2\2\2\u00a8\u00a9\7@\2\2\u00a9\u00aa"+
		"\7?\2\2\u00aa\16\3\2\2\2\u00ab\u00ac\7>\2\2\u00ac\u00ad\7?\2\2\u00ad\20"+
		"\3\2\2\2\u00ae\u00af\7}\2\2\u00af\22\3\2\2\2\u00b0\u00b1\7\177\2\2\u00b1"+
		"\24\3\2\2\2\u00b2\u00b3\7]\2\2\u00b3\26\3\2\2\2\u00b4\u00b5\7_\2\2\u00b5"+
		"\30\3\2\2\2\u00b6\u00b7\7<\2\2\u00b7\32\3\2\2\2\u00b8\u00b9\7.\2\2\u00b9"+
		"\34\3\2\2\2\u00ba\u00bb\7=\2\2\u00bb\36\3\2\2\2\u00bc\u00bd\7>\2\2\u00bd"+
		" \3\2\2\2\u00be\u00bf\7@\2\2\u00bf\"\3\2\2\2\u00c0\u00c1\7\60\2\2\u00c1"+
		"$\3\2\2\2\u00c2\u00c3\7*\2\2\u00c3&\3\2\2\2\u00c4\u00c5\7+\2\2\u00c5("+
		"\3\2\2\2\u00c6\u00c7\7?\2\2\u00c7*\3\2\2\2\u00c8\u00c9\7-\2\2\u00c9,\3"+
		"\2\2\2\u00ca\u00cb\7/\2\2\u00cb.\3\2\2\2\u00cc\u00cd\7,\2\2\u00cd\60\3"+
		"\2\2\2\u00ce\u00cf\7\61\2\2\u00cf\62\3\2\2\2\u00d0\u00d1\7\'\2\2\u00d1"+
		"\64\3\2\2\2\u00d2\u00d3\7~\2\2\u00d3\66\3\2\2\2\u00d4\u00d5\7(\2\2\u00d5"+
		"8\3\2\2\2\u00d6\u00d7\7#\2\2\u00d7:\3\2\2\2\u00d8\u00d9\7A\2\2\u00d9<"+
		"\3\2\2\2\u00da\u00dc\t\2\2\2\u00db\u00da\3\2\2\2\u00dc\u00dd\3\2\2\2\u00dd"+
		"\u00db\3\2\2\2\u00dd\u00de\3\2\2\2\u00de\u00df\3\2\2\2\u00df\u00e0\b\37"+
		"\2\2\u00e0>\3\2\2\2\u00e1\u00e2\7\61\2\2\u00e2\u00e3\7,\2\2\u00e3\u00e7"+
		"\3\2\2\2\u00e4\u00e6\13\2\2\2\u00e5\u00e4\3\2\2\2\u00e6\u00e9\3\2\2\2"+
		"\u00e7\u00e8\3\2\2\2\u00e7\u00e5\3\2\2\2\u00e8\u00ea\3\2\2\2\u00e9\u00e7"+
		"\3\2\2\2\u00ea\u00eb\7,\2\2\u00eb\u00ec\7\61\2\2\u00ec\u00ed\3\2\2\2\u00ed"+
		"\u00ee\b \2\2\u00ee@\3\2\2\2\u00ef\u00f0\7\61\2\2\u00f0\u00f1\7\61\2\2"+
		"\u00f1\u00f5\3\2\2\2\u00f2\u00f4\n\3\2\2\u00f3\u00f2\3\2\2\2\u00f4\u00f7"+
		"\3\2\2\2\u00f5\u00f3\3\2\2\2\u00f5\u00f6\3\2\2\2\u00f6\u00f8\3\2\2\2\u00f7"+
		"\u00f5\3\2\2\2\u00f8\u00f9\b!\2\2\u00f9B\3\2\2\2\u00fa\u00fb\7k\2\2\u00fb"+
		"\u00fc\78\2\2\u00fc\u00fd\7\66\2\2\u00fdD\3\2\2\2\u00fe\u00ff\7h\2\2\u00ff"+
		"\u0100\78\2\2\u0100\u0101\7\66\2\2\u0101F\3\2\2\2\u0102\u0103\7r\2\2\u0103"+
		"\u0104\7q\2\2\u0104\u0105\7y\2\2\u0105H\3\2\2\2\u0106\u0107\7x\2\2\u0107"+
		"\u0108\7g\2\2\u0108\u0109\7e\2\2\u0109J\3\2\2\2\u010a\u010b\7r\2\2\u010b"+
		"\u010c\7q\2\2\u010c\u010d\7y\2\2\u010d\u010e\7h\2\2\u010eL\3\2\2\2\u010f"+
		"\u0110\7d\2\2\u0110\u0111\7q\2\2\u0111\u0112\7q\2\2\u0112\u0113\7n\2\2"+
		"\u0113N\3\2\2\2\u0114\u0115\7e\2\2\u0115\u0116\7j\2\2\u0116\u0117\7c\2"+
		"\2\u0117\u0118\7t\2\2\u0118P\3\2\2\2\u0119\u011a\7(\2\2\u011a\u011b\7"+
		"u\2\2\u011b\u011c\7v\2\2\u011c\u011d\7t\2\2\u011dR\3\2\2\2\u011e\u011f"+
		"\7U\2\2\u011f\u0120\7v\2\2\u0120\u0121\7t\2\2\u0121\u0122\7k\2\2\u0122"+
		"\u0123\7p\2\2\u0123\u0124\7i\2\2\u0124T\3\2\2\2\u0125\u0126\7w\2\2\u0126"+
		"\u0127\7u\2\2\u0127\u0128\7k\2\2\u0128\u0129\7|\2\2\u0129\u012a\7g\2\2"+
		"\u012aV\3\2\2\2\u012b\u012c\7n\2\2\u012c\u012d\7g\2\2\u012d\u012e\7v\2"+
		"\2\u012eX\3\2\2\2\u012f\u0130\7o\2\2\u0130\u0131\7w\2\2\u0131\u0132\7"+
		"v\2\2\u0132Z\3\2\2\2\u0133\u0134\7u\2\2\u0134\u0135\7v\2\2\u0135\u0136"+
		"\7t\2\2\u0136\u0137\7w\2\2\u0137\u0138\7e\2\2\u0138\u0139\7v\2\2\u0139"+
		"\\\3\2\2\2\u013a\u013b\7c\2\2\u013b\u013c\7u\2\2\u013c^\3\2\2\2\u013d"+
		"\u013e\7r\2\2\u013e\u013f\7t\2\2\u013f\u0140\7k\2\2\u0140\u0141\7p\2\2"+
		"\u0141\u0142\7v\2\2\u0142\u0143\7n\2\2\u0143\u0144\7p\2\2\u0144\u0145"+
		"\7#\2\2\u0145`\3\2\2\2\u0146\u0147\7v\2\2\u0147\u0148\7t\2\2\u0148\u0149"+
		"\7w\2\2\u0149\u014a\7g\2\2\u014ab\3\2\2\2\u014b\u014c\7h\2\2\u014c\u014d"+
		"\7c\2\2\u014d\u014e\7n\2\2\u014e\u014f\7u\2\2\u014f\u0150\7g\2\2\u0150"+
		"d\3\2\2\2\u0151\u0152\7h\2\2\u0152\u0153\7p\2\2\u0153f\3\2\2\2\u0154\u0155"+
		"\7t\2\2\u0155\u0156\7g\2\2\u0156\u0157\7v\2\2\u0157\u0158\7w\2\2\u0158"+
		"\u0159\7t\2\2\u0159\u015a\7p\2\2\u015ah\3\2\2\2\u015b\u015c\7c\2\2\u015c"+
		"\u015d\7d\2\2\u015d\u015e\7u\2\2\u015ej\3\2\2\2\u015f\u0160\7u\2\2\u0160"+
		"\u0161\7s\2\2\u0161\u0162\7t\2\2\u0162\u0163\7v\2\2\u0163l\3\2\2\2\u0164"+
		"\u0165\7v\2\2\u0165\u0166\7q\2\2\u0166\u0167\7a\2\2\u0167\u0168\7u\2\2"+
		"\u0168\u0169\7v\2\2\u0169\u016a\7t\2\2\u016a\u016b\7k\2\2\u016b\u016c"+
		"\7p\2\2\u016c\u016d\7i\2\2\u016d\u016e\7*\2\2\u016e\u016f\7+\2\2\u016f"+
		"n\3\2\2\2\u0170\u0171\7e\2\2\u0171\u0172\7n\2\2\u0172\u0173\7q\2\2\u0173"+
		"\u0174\7p\2\2\u0174\u0175\7g\2\2\u0175\u0176\7*\2\2\u0176\u0177\7+\2\2"+
		"\u0177p\3\2\2\2\u0178\u0179\7p\2\2\u0179\u017a\7g\2\2\u017a\u017b\7y\2"+
		"\2\u017br\3\2\2\2\u017c\u017d\7n\2\2\u017d\u017e\7g\2\2\u017e\u017f\7"+
		"p\2\2\u017ft\3\2\2\2\u0180\u0181\7r\2\2\u0181\u0182\7w\2\2\u0182\u0183"+
		"\7u\2\2\u0183\u0184\7j\2\2\u0184v\3\2\2\2\u0185\u0186\7t\2\2\u0186\u0187"+
		"\7g\2\2\u0187\u0188\7o\2\2\u0188\u0189\7q\2\2\u0189\u018a\7x\2\2\u018a"+
		"\u018b\7g\2\2\u018bx\3\2\2\2\u018c\u018d\7e\2\2\u018d\u018e\7q\2\2\u018e"+
		"\u018f\7p\2\2\u018f\u0190\7v\2\2\u0190\u0191\7c\2\2\u0191\u0192\7k\2\2"+
		"\u0192\u0193\7p\2\2\u0193\u0194\7u\2\2\u0194z\3\2\2\2\u0195\u0196\7k\2"+
		"\2\u0196\u0197\7p\2\2\u0197\u0198\7u\2\2\u0198\u0199\7g\2\2\u0199\u019a"+
		"\7t\2\2\u019a\u019b\7v\2\2\u019b|\3\2\2\2\u019c\u019d\7e\2\2\u019d\u019e"+
		"\7c\2\2\u019e\u019f\7r\2\2\u019f\u01a0\7c\2\2\u01a0\u01a1\7e\2\2\u01a1"+
		"\u01a2\7k\2\2\u01a2\u01a3\7v\2\2\u01a3\u01a4\7{\2\2\u01a4~\3\2\2\2\u01a5"+
		"\u01a6\7y\2\2\u01a6\u01a7\7k\2\2\u01a7\u01a8\7v\2\2\u01a8\u01a9\7e\2\2"+
		"\u01a9\u01aa\7j\2\2\u01aa\u01ab\7a\2\2\u01ab\u01ac\7e\2\2\u01ac\u01ad"+
		"\7c\2\2\u01ad\u01ae\7r\2\2\u01ae\u01af\7c\2\2\u01af\u01b0\7e\2\2\u01b0"+
		"\u01b1\7k\2\2\u01b1\u01b2\7v\2\2\u01b2\u01b3\7{\2\2\u01b3\u0080\3\2\2"+
		"\2\u01b4\u01b5\7o\2\2\u01b5\u01b6\7c\2\2\u01b6\u01b7\7k\2\2\u01b7\u01b8"+
		"\7p\2\2\u01b8\u0082\3\2\2\2\u01b9\u01ba\7k\2\2\u01ba\u01bb\7h\2\2\u01bb"+
		"\u0084\3\2\2\2\u01bc\u01bd\7g\2\2\u01bd\u01be\7n\2\2\u01be\u01bf\7u\2"+
		"\2\u01bf\u01c0\7g\2\2\u01c0\u01c1\7\"\2\2\u01c1\u01c2\7k\2\2\u01c2\u01c3"+
		"\7h\2\2\u01c3\u0086\3\2\2\2\u01c4\u01c5\7g\2\2\u01c5\u01c6\7n\2\2\u01c6"+
		"\u01c7\7u\2\2\u01c7\u01c8\7g\2\2\u01c8\u0088\3\2\2\2\u01c9\u01ca\7y\2"+
		"\2\u01ca\u01cb\7j\2\2\u01cb\u01cc\7k\2\2\u01cc\u01cd\7n\2\2\u01cd\u01ce"+
		"\7g\2\2\u01ce\u008a\3\2\2\2\u01cf\u01d0\7d\2\2\u01d0\u01d1\7t\2\2\u01d1"+
		"\u01d2\7g\2\2\u01d2\u01d3\7c\2\2\u01d3\u01d4\7m\2\2\u01d4\u008c\3\2\2"+
		"\2\u01d5\u01d7\t\4\2\2\u01d6\u01d5\3\2\2\2\u01d7\u01d8\3\2\2\2\u01d8\u01d6"+
		"\3\2\2\2\u01d8\u01d9\3\2\2\2\u01d9\u008e\3\2\2\2\u01da\u01e2\5\u008dG"+
		"\2\u01db\u01dd\7\60\2\2\u01dc\u01de\5\u008dG\2\u01dd\u01dc\3\2\2\2\u01de"+
		"\u01df\3\2\2\2\u01df\u01dd\3\2\2\2\u01df\u01e0\3\2\2\2\u01e0\u01e3\3\2"+
		"\2\2\u01e1\u01e3\7\60\2\2\u01e2\u01db\3\2\2\2\u01e2\u01e1\3\2\2\2\u01e3"+
		"\u0090\3\2\2\2\u01e4\u01e9\5\u0093J\2\u01e5\u01e8\5\u0093J\2\u01e6\u01e8"+
		"\5\u008dG\2\u01e7\u01e5\3\2\2\2\u01e7\u01e6\3\2\2\2\u01e8\u01eb\3\2\2"+
		"\2\u01e9\u01e7\3\2\2\2\u01e9\u01ea\3\2\2\2\u01ea\u0092\3\2\2\2\u01eb\u01e9"+
		"\3\2\2\2\u01ec\u01ed\t\5\2\2\u01ed\u0094\3\2\2\2\u01ee\u01f2\7$\2\2\u01ef"+
		"\u01f1\n\6\2\2\u01f0\u01ef\3\2\2\2\u01f1\u01f4\3\2\2\2\u01f2\u01f0\3\2"+
		"\2\2\u01f2\u01f3\3\2\2\2\u01f3\u01f5\3\2\2\2\u01f4\u01f2\3\2\2\2\u01f5"+
		"\u01f6\7$\2\2\u01f6\u0096\3\2\2\2\u01f7\u01fb\7)\2\2\u01f8\u01fa\n\6\2"+
		"\2\u01f9\u01f8\3\2\2\2\u01fa\u01fd\3\2\2\2\u01fb\u01f9\3\2\2\2\u01fb\u01fc"+
		"\3\2\2\2\u01fc\u01fe\3\2\2\2\u01fd\u01fb\3\2\2\2\u01fe\u01ff\7)\2\2\u01ff"+
		"\u0098\3\2\2\2\r\2\u00dd\u00e7\u00f5\u01d8\u01df\u01e2\u01e7\u01e9\u01f2"+
		"\u01fb\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}