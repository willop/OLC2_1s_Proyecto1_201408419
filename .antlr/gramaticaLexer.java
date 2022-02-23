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
		TK_entero=30, TK_decimal=31, TK_id=32, Letra=33, TK_cadena=34, TK_caracter=35, 
		WHITESPACE=36, TK_comentario_multi=37, TK_comentario_lineal=38, TKR_numericos_enteros=39, 
		TKR_numericos_flotantes=40, TKR_bool=41, TKR_char=42, TKR_String=43, TKR_usize=44, 
		TKR_let=45, TKR_mut=46, TKR_struct=47, TKR_as=48, TKR_println=49, TKR_true=50, 
		TKR_false=51, TKR_fn=52, TKR_return=53, TKR_abs=54, TKR_sqrt=55, TKR_to_string=56, 
		TKR_clone=57, TKR_new=58, TKR_len=59, TKR_push=60, TKR_remove=61, TKR_contains=62, 
		TKR_insert=63, TKR_capacity=64, TKR_with_capacity=65;
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
			"TK_amp", "TK_sig_admiracion", "TK_sig_interrogacion", "TK_entero", "TK_decimal", 
			"TK_id", "Letra", "TK_cadena", "TK_caracter", "WHITESPACE", "TK_comentario_multi", 
			"TK_comentario_lineal", "TKR_numericos_enteros", "TKR_numericos_flotantes", 
			"TKR_bool", "TKR_char", "TKR_String", "TKR_usize", "TKR_let", "TKR_mut", 
			"TKR_struct", "TKR_as", "TKR_println", "TKR_true", "TKR_false", "TKR_fn", 
			"TKR_return", "TKR_abs", "TKR_sqrt", "TKR_to_string", "TKR_clone", "TKR_new", 
			"TKR_len", "TKR_push", "TKR_remove", "TKR_contains", "TKR_insert", "TKR_capacity", 
			"TKR_with_capacity"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'->'", "'||'", "'&&'", "'=='", "'!='", "'>='", "'<='", "'{'", 
			"'}'", "'['", "']'", "':'", "','", "';'", "'<'", "'>'", "'.'", "'('", 
			"')'", "'='", "'+'", "'-'", "'*'", "'/'", "'%'", "'|'", "'&'", "'!'", 
			"'?'", null, null, null, null, null, null, null, null, null, "'i64'", 
			"'f64'", "'bool'", "'char'", null, "'usize'", "'let'", "'mut'", "'struct'", 
			"'as'", "'println!'", "'true'", "'false'", "'fn'", "'return'", "'abs'", 
			"'sqrt'", "'to_string()'", "'clone()'", "'new'", "'len'", "'push'", "'remove'", 
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
			"TK_linea", "TK_amp", "TK_sig_admiracion", "TK_sig_interrogacion", "TK_entero", 
			"TK_decimal", "TK_id", "Letra", "TK_cadena", "TK_caracter", "WHITESPACE", 
			"TK_comentario_multi", "TK_comentario_lineal", "TKR_numericos_enteros", 
			"TKR_numericos_flotantes", "TKR_bool", "TKR_char", "TKR_String", "TKR_usize", 
			"TKR_let", "TKR_mut", "TKR_struct", "TKR_as", "TKR_println", "TKR_true", 
			"TKR_false", "TKR_fn", "TKR_return", "TKR_abs", "TKR_sqrt", "TKR_to_string", 
			"TKR_clone", "TKR_new", "TKR_len", "TKR_push", "TKR_remove", "TKR_contains", 
			"TKR_insert", "TKR_capacity", "TKR_with_capacity"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2C\u01be\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\3\2\3\2\3\2\3\3\3\3\3\3\3\4\3\4\3\4\3\5"+
		"\3\5\3\5\3\6\3\6\3\6\3\7\3\7\3\7\3\b\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13"+
		"\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23"+
		"\3\23\3\24\3\24\3\25\3\25\3\26\3\26\3\27\3\27\3\30\3\30\3\31\3\31\3\32"+
		"\3\32\3\33\3\33\3\34\3\34\3\35\3\35\3\36\3\36\3\37\6\37\u00c8\n\37\r\37"+
		"\16\37\u00c9\3 \3 \3 \6 \u00cf\n \r \16 \u00d0\3 \5 \u00d4\n \3!\3!\3"+
		"!\7!\u00d9\n!\f!\16!\u00dc\13!\3\"\3\"\3#\3#\7#\u00e2\n#\f#\16#\u00e5"+
		"\13#\3#\3#\3$\3$\7$\u00eb\n$\f$\16$\u00ee\13$\3$\3$\3%\6%\u00f3\n%\r%"+
		"\16%\u00f4\3%\3%\3&\3&\3&\3&\7&\u00fd\n&\f&\16&\u0100\13&\3&\3&\3&\3&"+
		"\3&\3\'\3\'\3\'\3\'\7\'\u010b\n\'\f\'\16\'\u010e\13\'\3\'\3\'\3(\3(\3"+
		"(\3(\3)\3)\3)\3)\3*\3*\3*\3*\3*\3+\3+\3+\3+\3+\3,\3,\3,\3,\3,\3,\3,\3"+
		",\3,\3,\5,\u012e\n,\3-\3-\3-\3-\3-\3-\3.\3.\3.\3.\3/\3/\3/\3/\3\60\3\60"+
		"\3\60\3\60\3\60\3\60\3\60\3\61\3\61\3\61\3\62\3\62\3\62\3\62\3\62\3\62"+
		"\3\62\3\62\3\62\3\63\3\63\3\63\3\63\3\63\3\64\3\64\3\64\3\64\3\64\3\64"+
		"\3\65\3\65\3\65\3\66\3\66\3\66\3\66\3\66\3\66\3\66\3\67\3\67\3\67\3\67"+
		"\38\38\38\38\38\39\39\39\39\39\39\39\39\39\39\39\39\3:\3:\3:\3:\3:\3:"+
		"\3:\3:\3;\3;\3;\3;\3<\3<\3<\3<\3=\3=\3=\3=\3=\3>\3>\3>\3>\3>\3>\3>\3?"+
		"\3?\3?\3?\3?\3?\3?\3?\3?\3@\3@\3@\3@\3@\3@\3@\3A\3A\3A\3A\3A\3A\3A\3A"+
		"\3A\3B\3B\3B\3B\3B\3B\3B\3B\3B\3B\3B\3B\3B\3B\3B\3\u00fe\2C\3\3\5\4\7"+
		"\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22"+
		"#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C"+
		"#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67m8o9q:s;u<w"+
		"=y>{?}@\177A\u0081B\u0083C\3\2\7\3\2\62;\5\2C\\aac|\3\2$$\4\2\13\f\17"+
		"\17\4\2\f\f\17\17\2\u01c8\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2"+
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
		"\2\2\2\u0083\3\2\2\2\3\u0085\3\2\2\2\5\u0088\3\2\2\2\7\u008b\3\2\2\2\t"+
		"\u008e\3\2\2\2\13\u0091\3\2\2\2\r\u0094\3\2\2\2\17\u0097\3\2\2\2\21\u009a"+
		"\3\2\2\2\23\u009c\3\2\2\2\25\u009e\3\2\2\2\27\u00a0\3\2\2\2\31\u00a2\3"+
		"\2\2\2\33\u00a4\3\2\2\2\35\u00a6\3\2\2\2\37\u00a8\3\2\2\2!\u00aa\3\2\2"+
		"\2#\u00ac\3\2\2\2%\u00ae\3\2\2\2\'\u00b0\3\2\2\2)\u00b2\3\2\2\2+\u00b4"+
		"\3\2\2\2-\u00b6\3\2\2\2/\u00b8\3\2\2\2\61\u00ba\3\2\2\2\63\u00bc\3\2\2"+
		"\2\65\u00be\3\2\2\2\67\u00c0\3\2\2\29\u00c2\3\2\2\2;\u00c4\3\2\2\2=\u00c7"+
		"\3\2\2\2?\u00cb\3\2\2\2A\u00d5\3\2\2\2C\u00dd\3\2\2\2E\u00df\3\2\2\2G"+
		"\u00e8\3\2\2\2I\u00f2\3\2\2\2K\u00f8\3\2\2\2M\u0106\3\2\2\2O\u0111\3\2"+
		"\2\2Q\u0115\3\2\2\2S\u0119\3\2\2\2U\u011e\3\2\2\2W\u012d\3\2\2\2Y\u012f"+
		"\3\2\2\2[\u0135\3\2\2\2]\u0139\3\2\2\2_\u013d\3\2\2\2a\u0144\3\2\2\2c"+
		"\u0147\3\2\2\2e\u0150\3\2\2\2g\u0155\3\2\2\2i\u015b\3\2\2\2k\u015e\3\2"+
		"\2\2m\u0165\3\2\2\2o\u0169\3\2\2\2q\u016e\3\2\2\2s\u017a\3\2\2\2u\u0182"+
		"\3\2\2\2w\u0186\3\2\2\2y\u018a\3\2\2\2{\u018f\3\2\2\2}\u0196\3\2\2\2\177"+
		"\u019f\3\2\2\2\u0081\u01a6\3\2\2\2\u0083\u01af\3\2\2\2\u0085\u0086\7/"+
		"\2\2\u0086\u0087\7@\2\2\u0087\4\3\2\2\2\u0088\u0089\7~\2\2\u0089\u008a"+
		"\7~\2\2\u008a\6\3\2\2\2\u008b\u008c\7(\2\2\u008c\u008d\7(\2\2\u008d\b"+
		"\3\2\2\2\u008e\u008f\7?\2\2\u008f\u0090\7?\2\2\u0090\n\3\2\2\2\u0091\u0092"+
		"\7#\2\2\u0092\u0093\7?\2\2\u0093\f\3\2\2\2\u0094\u0095\7@\2\2\u0095\u0096"+
		"\7?\2\2\u0096\16\3\2\2\2\u0097\u0098\7>\2\2\u0098\u0099\7?\2\2\u0099\20"+
		"\3\2\2\2\u009a\u009b\7}\2\2\u009b\22\3\2\2\2\u009c\u009d\7\177\2\2\u009d"+
		"\24\3\2\2\2\u009e\u009f\7]\2\2\u009f\26\3\2\2\2\u00a0\u00a1\7_\2\2\u00a1"+
		"\30\3\2\2\2\u00a2\u00a3\7<\2\2\u00a3\32\3\2\2\2\u00a4\u00a5\7.\2\2\u00a5"+
		"\34\3\2\2\2\u00a6\u00a7\7=\2\2\u00a7\36\3\2\2\2\u00a8\u00a9\7>\2\2\u00a9"+
		" \3\2\2\2\u00aa\u00ab\7@\2\2\u00ab\"\3\2\2\2\u00ac\u00ad\7\60\2\2\u00ad"+
		"$\3\2\2\2\u00ae\u00af\7*\2\2\u00af&\3\2\2\2\u00b0\u00b1\7+\2\2\u00b1("+
		"\3\2\2\2\u00b2\u00b3\7?\2\2\u00b3*\3\2\2\2\u00b4\u00b5\7-\2\2\u00b5,\3"+
		"\2\2\2\u00b6\u00b7\7/\2\2\u00b7.\3\2\2\2\u00b8\u00b9\7,\2\2\u00b9\60\3"+
		"\2\2\2\u00ba\u00bb\7\61\2\2\u00bb\62\3\2\2\2\u00bc\u00bd\7\'\2\2\u00bd"+
		"\64\3\2\2\2\u00be\u00bf\7~\2\2\u00bf\66\3\2\2\2\u00c0\u00c1\7(\2\2\u00c1"+
		"8\3\2\2\2\u00c2\u00c3\7#\2\2\u00c3:\3\2\2\2\u00c4\u00c5\7A\2\2\u00c5<"+
		"\3\2\2\2\u00c6\u00c8\t\2\2\2\u00c7\u00c6\3\2\2\2\u00c8\u00c9\3\2\2\2\u00c9"+
		"\u00c7\3\2\2\2\u00c9\u00ca\3\2\2\2\u00ca>\3\2\2\2\u00cb\u00d3\5=\37\2"+
		"\u00cc\u00ce\7\60\2\2\u00cd\u00cf\5=\37\2\u00ce\u00cd\3\2\2\2\u00cf\u00d0"+
		"\3\2\2\2\u00d0\u00ce\3\2\2\2\u00d0\u00d1\3\2\2\2\u00d1\u00d4\3\2\2\2\u00d2"+
		"\u00d4\7\60\2\2\u00d3\u00cc\3\2\2\2\u00d3\u00d2\3\2\2\2\u00d4@\3\2\2\2"+
		"\u00d5\u00da\5C\"\2\u00d6\u00d9\5C\"\2\u00d7\u00d9\5=\37\2\u00d8\u00d6"+
		"\3\2\2\2\u00d8\u00d7\3\2\2\2\u00d9\u00dc\3\2\2\2\u00da\u00d8\3\2\2\2\u00da"+
		"\u00db\3\2\2\2\u00dbB\3\2\2\2\u00dc\u00da\3\2\2\2\u00dd\u00de\t\3\2\2"+
		"\u00deD\3\2\2\2\u00df\u00e3\7$\2\2\u00e0\u00e2\n\4\2\2\u00e1\u00e0\3\2"+
		"\2\2\u00e2\u00e5\3\2\2\2\u00e3\u00e1\3\2\2\2\u00e3\u00e4\3\2\2\2\u00e4"+
		"\u00e6\3\2\2\2\u00e5\u00e3\3\2\2\2\u00e6\u00e7\7$\2\2\u00e7F\3\2\2\2\u00e8"+
		"\u00ec\7)\2\2\u00e9\u00eb\n\4\2\2\u00ea\u00e9\3\2\2\2\u00eb\u00ee\3\2"+
		"\2\2\u00ec\u00ea\3\2\2\2\u00ec\u00ed\3\2\2\2\u00ed\u00ef\3\2\2\2\u00ee"+
		"\u00ec\3\2\2\2\u00ef\u00f0\7)\2\2\u00f0H\3\2\2\2\u00f1\u00f3\t\5\2\2\u00f2"+
		"\u00f1\3\2\2\2\u00f3\u00f4\3\2\2\2\u00f4\u00f2\3\2\2\2\u00f4\u00f5\3\2"+
		"\2\2\u00f5\u00f6\3\2\2\2\u00f6\u00f7\b%\2\2\u00f7J\3\2\2\2\u00f8\u00f9"+
		"\7\61\2\2\u00f9\u00fa\7,\2\2\u00fa\u00fe\3\2\2\2\u00fb\u00fd\13\2\2\2"+
		"\u00fc\u00fb\3\2\2\2\u00fd\u0100\3\2\2\2\u00fe\u00ff\3\2\2\2\u00fe\u00fc"+
		"\3\2\2\2\u00ff\u0101\3\2\2\2\u0100\u00fe\3\2\2\2\u0101\u0102\7,\2\2\u0102"+
		"\u0103\7\61\2\2\u0103\u0104\3\2\2\2\u0104\u0105\b&\2\2\u0105L\3\2\2\2"+
		"\u0106\u0107\7\61\2\2\u0107\u0108\7\61\2\2\u0108\u010c\3\2\2\2\u0109\u010b"+
		"\n\6\2\2\u010a\u0109\3\2\2\2\u010b\u010e\3\2\2\2\u010c\u010a\3\2\2\2\u010c"+
		"\u010d\3\2\2\2\u010d\u010f\3\2\2\2\u010e\u010c\3\2\2\2\u010f\u0110\b\'"+
		"\2\2\u0110N\3\2\2\2\u0111\u0112\7k\2\2\u0112\u0113\78\2\2\u0113\u0114"+
		"\7\66\2\2\u0114P\3\2\2\2\u0115\u0116\7h\2\2\u0116\u0117\78\2\2\u0117\u0118"+
		"\7\66\2\2\u0118R\3\2\2\2\u0119\u011a\7d\2\2\u011a\u011b\7q\2\2\u011b\u011c"+
		"\7q\2\2\u011c\u011d\7n\2\2\u011dT\3\2\2\2\u011e\u011f\7e\2\2\u011f\u0120"+
		"\7j\2\2\u0120\u0121\7c\2\2\u0121\u0122\7t\2\2\u0122V\3\2\2\2\u0123\u0124"+
		"\7(\2\2\u0124\u0125\7u\2\2\u0125\u0126\7v\2\2\u0126\u012e\7t\2\2\u0127"+
		"\u0128\7U\2\2\u0128\u0129\7v\2\2\u0129\u012a\7t\2\2\u012a\u012b\7k\2\2"+
		"\u012b\u012c\7p\2\2\u012c\u012e\7i\2\2\u012d\u0123\3\2\2\2\u012d\u0127"+
		"\3\2\2\2\u012eX\3\2\2\2\u012f\u0130\7w\2\2\u0130\u0131\7u\2\2\u0131\u0132"+
		"\7k\2\2\u0132\u0133\7|\2\2\u0133\u0134\7g\2\2\u0134Z\3\2\2\2\u0135\u0136"+
		"\7n\2\2\u0136\u0137\7g\2\2\u0137\u0138\7v\2\2\u0138\\\3\2\2\2\u0139\u013a"+
		"\7o\2\2\u013a\u013b\7w\2\2\u013b\u013c\7v\2\2\u013c^\3\2\2\2\u013d\u013e"+
		"\7u\2\2\u013e\u013f\7v\2\2\u013f\u0140\7t\2\2\u0140\u0141\7w\2\2\u0141"+
		"\u0142\7e\2\2\u0142\u0143\7v\2\2\u0143`\3\2\2\2\u0144\u0145\7c\2\2\u0145"+
		"\u0146\7u\2\2\u0146b\3\2\2\2\u0147\u0148\7r\2\2\u0148\u0149\7t\2\2\u0149"+
		"\u014a\7k\2\2\u014a\u014b\7p\2\2\u014b\u014c\7v\2\2\u014c\u014d\7n\2\2"+
		"\u014d\u014e\7p\2\2\u014e\u014f\7#\2\2\u014fd\3\2\2\2\u0150\u0151\7v\2"+
		"\2\u0151\u0152\7t\2\2\u0152\u0153\7w\2\2\u0153\u0154\7g\2\2\u0154f\3\2"+
		"\2\2\u0155\u0156\7h\2\2\u0156\u0157\7c\2\2\u0157\u0158\7n\2\2\u0158\u0159"+
		"\7u\2\2\u0159\u015a\7g\2\2\u015ah\3\2\2\2\u015b\u015c\7h\2\2\u015c\u015d"+
		"\7p\2\2\u015dj\3\2\2\2\u015e\u015f\7t\2\2\u015f\u0160\7g\2\2\u0160\u0161"+
		"\7v\2\2\u0161\u0162\7w\2\2\u0162\u0163\7t\2\2\u0163\u0164\7p\2\2\u0164"+
		"l\3\2\2\2\u0165\u0166\7c\2\2\u0166\u0167\7d\2\2\u0167\u0168\7u\2\2\u0168"+
		"n\3\2\2\2\u0169\u016a\7u\2\2\u016a\u016b\7s\2\2\u016b\u016c\7t\2\2\u016c"+
		"\u016d\7v\2\2\u016dp\3\2\2\2\u016e\u016f\7v\2\2\u016f\u0170\7q\2\2\u0170"+
		"\u0171\7a\2\2\u0171\u0172\7u\2\2\u0172\u0173\7v\2\2\u0173\u0174\7t\2\2"+
		"\u0174\u0175\7k\2\2\u0175\u0176\7p\2\2\u0176\u0177\7i\2\2\u0177\u0178"+
		"\7*\2\2\u0178\u0179\7+\2\2\u0179r\3\2\2\2\u017a\u017b\7e\2\2\u017b\u017c"+
		"\7n\2\2\u017c\u017d\7q\2\2\u017d\u017e\7p\2\2\u017e\u017f\7g\2\2\u017f"+
		"\u0180\7*\2\2\u0180\u0181\7+\2\2\u0181t\3\2\2\2\u0182\u0183\7p\2\2\u0183"+
		"\u0184\7g\2\2\u0184\u0185\7y\2\2\u0185v\3\2\2\2\u0186\u0187\7n\2\2\u0187"+
		"\u0188\7g\2\2\u0188\u0189\7p\2\2\u0189x\3\2\2\2\u018a\u018b\7r\2\2\u018b"+
		"\u018c\7w\2\2\u018c\u018d\7u\2\2\u018d\u018e\7j\2\2\u018ez\3\2\2\2\u018f"+
		"\u0190\7t\2\2\u0190\u0191\7g\2\2\u0191\u0192\7o\2\2\u0192\u0193\7q\2\2"+
		"\u0193\u0194\7x\2\2\u0194\u0195\7g\2\2\u0195|\3\2\2\2\u0196\u0197\7e\2"+
		"\2\u0197\u0198\7q\2\2\u0198\u0199\7p\2\2\u0199\u019a\7v\2\2\u019a\u019b"+
		"\7c\2\2\u019b\u019c\7k\2\2\u019c\u019d\7p\2\2\u019d\u019e\7u\2\2\u019e"+
		"~\3\2\2\2\u019f\u01a0\7k\2\2\u01a0\u01a1\7p\2\2\u01a1\u01a2\7u\2\2\u01a2"+
		"\u01a3\7g\2\2\u01a3\u01a4\7t\2\2\u01a4\u01a5\7v\2\2\u01a5\u0080\3\2\2"+
		"\2\u01a6\u01a7\7e\2\2\u01a7\u01a8\7c\2\2\u01a8\u01a9\7r\2\2\u01a9\u01aa"+
		"\7c\2\2\u01aa\u01ab\7e\2\2\u01ab\u01ac\7k\2\2\u01ac\u01ad\7v\2\2\u01ad"+
		"\u01ae\7{\2\2\u01ae\u0082\3\2\2\2\u01af\u01b0\7y\2\2\u01b0\u01b1\7k\2"+
		"\2\u01b1\u01b2\7v\2\2\u01b2\u01b3\7e\2\2\u01b3\u01b4\7j\2\2\u01b4\u01b5"+
		"\7a\2\2\u01b5\u01b6\7e\2\2\u01b6\u01b7\7c\2\2\u01b7\u01b8\7r\2\2\u01b8"+
		"\u01b9\7c\2\2\u01b9\u01ba\7e\2\2\u01ba\u01bb\7k\2\2\u01bb\u01bc\7v\2\2"+
		"\u01bc\u01bd\7{\2\2\u01bd\u0084\3\2\2\2\16\2\u00c9\u00d0\u00d3\u00d8\u00da"+
		"\u00e3\u00ec\u00f4\u00fe\u010c\u012d\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}