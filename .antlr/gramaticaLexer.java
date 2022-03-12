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
		TKR_capacity=62, TKR_with_capacity=63, TKR_main=64, TK_entero=65, TK_decimal=66, 
		TK_id=67, Letra=68, TK_cadena=69, TK_caracter=70;
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
			"TKR_with_capacity", "TKR_main", "TK_entero", "TK_decimal", "TK_id", 
			"Letra", "TK_cadena", "TK_caracter"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2H\u01da\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\3\2\3\2\3"+
		"\2\3\3\3\3\3\3\3\4\3\4\3\4\3\5\3\5\3\5\3\6\3\6\3\6\3\7\3\7\3\7\3\b\3\b"+
		"\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3\20"+
		"\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3\24\3\24\3\25\3\25\3\26\3\26\3\27"+
		"\3\27\3\30\3\30\3\31\3\31\3\32\3\32\3\33\3\33\3\34\3\34\3\35\3\35\3\36"+
		"\3\36\3\37\6\37\u00d2\n\37\r\37\16\37\u00d3\3\37\3\37\3 \3 \3 \3 \7 \u00dc"+
		"\n \f \16 \u00df\13 \3 \3 \3 \3 \3 \3!\3!\3!\3!\7!\u00ea\n!\f!\16!\u00ed"+
		"\13!\3!\3!\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3$\3$\3$\3$\3%\3%\3%\3%\3&\3&\3"+
		"&\3&\3&\3\'\3\'\3\'\3\'\3\'\3(\3(\3(\3(\3(\3)\3)\3)\3)\3)\3*\3*\3*\3*"+
		"\3*\3*\3*\3+\3+\3+\3+\3+\3+\3,\3,\3,\3,\3-\3-\3-\3-\3.\3.\3.\3.\3.\3."+
		"\3.\3/\3/\3/\3\60\3\60\3\60\3\60\3\60\3\60\3\60\3\60\3\60\3\61\3\61\3"+
		"\61\3\61\3\61\3\62\3\62\3\62\3\62\3\62\3\62\3\63\3\63\3\63\3\64\3\64\3"+
		"\64\3\64\3\64\3\64\3\64\3\65\3\65\3\65\3\65\3\66\3\66\3\66\3\66\3\66\3"+
		"\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\38\38\38\3"+
		"8\38\38\38\38\39\39\39\39\3:\3:\3:\3:\3;\3;\3;\3;\3;\3<\3<\3<\3<\3<\3"+
		"<\3<\3=\3=\3=\3=\3=\3=\3=\3=\3=\3>\3>\3>\3>\3>\3>\3>\3?\3?\3?\3?\3?\3"+
		"?\3?\3?\3?\3@\3@\3@\3@\3@\3@\3@\3@\3@\3@\3@\3@\3@\3@\3@\3A\3A\3A\3A\3"+
		"A\3B\6B\u01b1\nB\rB\16B\u01b2\3C\3C\3C\6C\u01b8\nC\rC\16C\u01b9\3C\5C"+
		"\u01bd\nC\3D\3D\3D\7D\u01c2\nD\fD\16D\u01c5\13D\3E\3E\3F\3F\7F\u01cb\n"+
		"F\fF\16F\u01ce\13F\3F\3F\3G\3G\7G\u01d4\nG\fG\16G\u01d7\13G\3G\3G\3\u00dd"+
		"\2H\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35"+
		"\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36"+
		";\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67"+
		"m8o9q:s;u<w=y>{?}@\177A\u0081B\u0083C\u0085D\u0087E\u0089F\u008bG\u008d"+
		"H\3\2\7\5\2\13\f\17\17\"\"\4\2\f\f\17\17\3\2\62;\5\2C\\aac|\3\2$$\2\u01e3"+
		"\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2"+
		"\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2"+
		"\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2"+
		"\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2"+
		"\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3"+
		"\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2"+
		"\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2"+
		"U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3"+
		"\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2"+
		"\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y\3\2\2\2\2"+
		"{\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2\2\u0081\3\2\2\2\2\u0083\3\2\2\2\2\u0085"+
		"\3\2\2\2\2\u0087\3\2\2\2\2\u0089\3\2\2\2\2\u008b\3\2\2\2\2\u008d\3\2\2"+
		"\2\3\u008f\3\2\2\2\5\u0092\3\2\2\2\7\u0095\3\2\2\2\t\u0098\3\2\2\2\13"+
		"\u009b\3\2\2\2\r\u009e\3\2\2\2\17\u00a1\3\2\2\2\21\u00a4\3\2\2\2\23\u00a6"+
		"\3\2\2\2\25\u00a8\3\2\2\2\27\u00aa\3\2\2\2\31\u00ac\3\2\2\2\33\u00ae\3"+
		"\2\2\2\35\u00b0\3\2\2\2\37\u00b2\3\2\2\2!\u00b4\3\2\2\2#\u00b6\3\2\2\2"+
		"%\u00b8\3\2\2\2\'\u00ba\3\2\2\2)\u00bc\3\2\2\2+\u00be\3\2\2\2-\u00c0\3"+
		"\2\2\2/\u00c2\3\2\2\2\61\u00c4\3\2\2\2\63\u00c6\3\2\2\2\65\u00c8\3\2\2"+
		"\2\67\u00ca\3\2\2\29\u00cc\3\2\2\2;\u00ce\3\2\2\2=\u00d1\3\2\2\2?\u00d7"+
		"\3\2\2\2A\u00e5\3\2\2\2C\u00f0\3\2\2\2E\u00f4\3\2\2\2G\u00f8\3\2\2\2I"+
		"\u00fc\3\2\2\2K\u0100\3\2\2\2M\u0105\3\2\2\2O\u010a\3\2\2\2Q\u010f\3\2"+
		"\2\2S\u0114\3\2\2\2U\u011b\3\2\2\2W\u0121\3\2\2\2Y\u0125\3\2\2\2[\u0129"+
		"\3\2\2\2]\u0130\3\2\2\2_\u0133\3\2\2\2a\u013c\3\2\2\2c\u0141\3\2\2\2e"+
		"\u0147\3\2\2\2g\u014a\3\2\2\2i\u0151\3\2\2\2k\u0155\3\2\2\2m\u015a\3\2"+
		"\2\2o\u0166\3\2\2\2q\u016e\3\2\2\2s\u0172\3\2\2\2u\u0176\3\2\2\2w\u017b"+
		"\3\2\2\2y\u0182\3\2\2\2{\u018b\3\2\2\2}\u0192\3\2\2\2\177\u019b\3\2\2"+
		"\2\u0081\u01aa\3\2\2\2\u0083\u01b0\3\2\2\2\u0085\u01b4\3\2\2\2\u0087\u01be"+
		"\3\2\2\2\u0089\u01c6\3\2\2\2\u008b\u01c8\3\2\2\2\u008d\u01d1\3\2\2\2\u008f"+
		"\u0090\7/\2\2\u0090\u0091\7@\2\2\u0091\4\3\2\2\2\u0092\u0093\7~\2\2\u0093"+
		"\u0094\7~\2\2\u0094\6\3\2\2\2\u0095\u0096\7(\2\2\u0096\u0097\7(\2\2\u0097"+
		"\b\3\2\2\2\u0098\u0099\7?\2\2\u0099\u009a\7?\2\2\u009a\n\3\2\2\2\u009b"+
		"\u009c\7#\2\2\u009c\u009d\7?\2\2\u009d\f\3\2\2\2\u009e\u009f\7@\2\2\u009f"+
		"\u00a0\7?\2\2\u00a0\16\3\2\2\2\u00a1\u00a2\7>\2\2\u00a2\u00a3\7?\2\2\u00a3"+
		"\20\3\2\2\2\u00a4\u00a5\7}\2\2\u00a5\22\3\2\2\2\u00a6\u00a7\7\177\2\2"+
		"\u00a7\24\3\2\2\2\u00a8\u00a9\7]\2\2\u00a9\26\3\2\2\2\u00aa\u00ab\7_\2"+
		"\2\u00ab\30\3\2\2\2\u00ac\u00ad\7<\2\2\u00ad\32\3\2\2\2\u00ae\u00af\7"+
		".\2\2\u00af\34\3\2\2\2\u00b0\u00b1\7=\2\2\u00b1\36\3\2\2\2\u00b2\u00b3"+
		"\7>\2\2\u00b3 \3\2\2\2\u00b4\u00b5\7@\2\2\u00b5\"\3\2\2\2\u00b6\u00b7"+
		"\7\60\2\2\u00b7$\3\2\2\2\u00b8\u00b9\7*\2\2\u00b9&\3\2\2\2\u00ba\u00bb"+
		"\7+\2\2\u00bb(\3\2\2\2\u00bc\u00bd\7?\2\2\u00bd*\3\2\2\2\u00be\u00bf\7"+
		"-\2\2\u00bf,\3\2\2\2\u00c0\u00c1\7/\2\2\u00c1.\3\2\2\2\u00c2\u00c3\7,"+
		"\2\2\u00c3\60\3\2\2\2\u00c4\u00c5\7\61\2\2\u00c5\62\3\2\2\2\u00c6\u00c7"+
		"\7\'\2\2\u00c7\64\3\2\2\2\u00c8\u00c9\7~\2\2\u00c9\66\3\2\2\2\u00ca\u00cb"+
		"\7(\2\2\u00cb8\3\2\2\2\u00cc\u00cd\7#\2\2\u00cd:\3\2\2\2\u00ce\u00cf\7"+
		"A\2\2\u00cf<\3\2\2\2\u00d0\u00d2\t\2\2\2\u00d1\u00d0\3\2\2\2\u00d2\u00d3"+
		"\3\2\2\2\u00d3\u00d1\3\2\2\2\u00d3\u00d4\3\2\2\2\u00d4\u00d5\3\2\2\2\u00d5"+
		"\u00d6\b\37\2\2\u00d6>\3\2\2\2\u00d7\u00d8\7\61\2\2\u00d8\u00d9\7,\2\2"+
		"\u00d9\u00dd\3\2\2\2\u00da\u00dc\13\2\2\2\u00db\u00da\3\2\2\2\u00dc\u00df"+
		"\3\2\2\2\u00dd\u00de\3\2\2\2\u00dd\u00db\3\2\2\2\u00de\u00e0\3\2\2\2\u00df"+
		"\u00dd\3\2\2\2\u00e0\u00e1\7,\2\2\u00e1\u00e2\7\61\2\2\u00e2\u00e3\3\2"+
		"\2\2\u00e3\u00e4\b \2\2\u00e4@\3\2\2\2\u00e5\u00e6\7\61\2\2\u00e6\u00e7"+
		"\7\61\2\2\u00e7\u00eb\3\2\2\2\u00e8\u00ea\n\3\2\2\u00e9\u00e8\3\2\2\2"+
		"\u00ea\u00ed\3\2\2\2\u00eb\u00e9\3\2\2\2\u00eb\u00ec\3\2\2\2\u00ec\u00ee"+
		"\3\2\2\2\u00ed\u00eb\3\2\2\2\u00ee\u00ef\b!\2\2\u00efB\3\2\2\2\u00f0\u00f1"+
		"\7k\2\2\u00f1\u00f2\78\2\2\u00f2\u00f3\7\66\2\2\u00f3D\3\2\2\2\u00f4\u00f5"+
		"\7h\2\2\u00f5\u00f6\78\2\2\u00f6\u00f7\7\66\2\2\u00f7F\3\2\2\2\u00f8\u00f9"+
		"\7r\2\2\u00f9\u00fa\7q\2\2\u00fa\u00fb\7y\2\2\u00fbH\3\2\2\2\u00fc\u00fd"+
		"\7x\2\2\u00fd\u00fe\7g\2\2\u00fe\u00ff\7e\2\2\u00ffJ\3\2\2\2\u0100\u0101"+
		"\7r\2\2\u0101\u0102\7q\2\2\u0102\u0103\7y\2\2\u0103\u0104\7h\2\2\u0104"+
		"L\3\2\2\2\u0105\u0106\7d\2\2\u0106\u0107\7q\2\2\u0107\u0108\7q\2\2\u0108"+
		"\u0109\7n\2\2\u0109N\3\2\2\2\u010a\u010b\7e\2\2\u010b\u010c\7j\2\2\u010c"+
		"\u010d\7c\2\2\u010d\u010e\7t\2\2\u010eP\3\2\2\2\u010f\u0110\7(\2\2\u0110"+
		"\u0111\7u\2\2\u0111\u0112\7v\2\2\u0112\u0113\7t\2\2\u0113R\3\2\2\2\u0114"+
		"\u0115\7U\2\2\u0115\u0116\7v\2\2\u0116\u0117\7t\2\2\u0117\u0118\7k\2\2"+
		"\u0118\u0119\7p\2\2\u0119\u011a\7i\2\2\u011aT\3\2\2\2\u011b\u011c\7w\2"+
		"\2\u011c\u011d\7u\2\2\u011d\u011e\7k\2\2\u011e\u011f\7|\2\2\u011f\u0120"+
		"\7g\2\2\u0120V\3\2\2\2\u0121\u0122\7n\2\2\u0122\u0123\7g\2\2\u0123\u0124"+
		"\7v\2\2\u0124X\3\2\2\2\u0125\u0126\7o\2\2\u0126\u0127\7w\2\2\u0127\u0128"+
		"\7v\2\2\u0128Z\3\2\2\2\u0129\u012a\7u\2\2\u012a\u012b\7v\2\2\u012b\u012c"+
		"\7t\2\2\u012c\u012d\7w\2\2\u012d\u012e\7e\2\2\u012e\u012f\7v\2\2\u012f"+
		"\\\3\2\2\2\u0130\u0131\7c\2\2\u0131\u0132\7u\2\2\u0132^\3\2\2\2\u0133"+
		"\u0134\7r\2\2\u0134\u0135\7t\2\2\u0135\u0136\7k\2\2\u0136\u0137\7p\2\2"+
		"\u0137\u0138\7v\2\2\u0138\u0139\7n\2\2\u0139\u013a\7p\2\2\u013a\u013b"+
		"\7#\2\2\u013b`\3\2\2\2\u013c\u013d\7v\2\2\u013d\u013e\7t\2\2\u013e\u013f"+
		"\7w\2\2\u013f\u0140\7g\2\2\u0140b\3\2\2\2\u0141\u0142\7h\2\2\u0142\u0143"+
		"\7c\2\2\u0143\u0144\7n\2\2\u0144\u0145\7u\2\2\u0145\u0146\7g\2\2\u0146"+
		"d\3\2\2\2\u0147\u0148\7h\2\2\u0148\u0149\7p\2\2\u0149f\3\2\2\2\u014a\u014b"+
		"\7t\2\2\u014b\u014c\7g\2\2\u014c\u014d\7v\2\2\u014d\u014e\7w\2\2\u014e"+
		"\u014f\7t\2\2\u014f\u0150\7p\2\2\u0150h\3\2\2\2\u0151\u0152\7c\2\2\u0152"+
		"\u0153\7d\2\2\u0153\u0154\7u\2\2\u0154j\3\2\2\2\u0155\u0156\7u\2\2\u0156"+
		"\u0157\7s\2\2\u0157\u0158\7t\2\2\u0158\u0159\7v\2\2\u0159l\3\2\2\2\u015a"+
		"\u015b\7v\2\2\u015b\u015c\7q\2\2\u015c\u015d\7a\2\2\u015d\u015e\7u\2\2"+
		"\u015e\u015f\7v\2\2\u015f\u0160\7t\2\2\u0160\u0161\7k\2\2\u0161\u0162"+
		"\7p\2\2\u0162\u0163\7i\2\2\u0163\u0164\7*\2\2\u0164\u0165\7+\2\2\u0165"+
		"n\3\2\2\2\u0166\u0167\7e\2\2\u0167\u0168\7n\2\2\u0168\u0169\7q\2\2\u0169"+
		"\u016a\7p\2\2\u016a\u016b\7g\2\2\u016b\u016c\7*\2\2\u016c\u016d\7+\2\2"+
		"\u016dp\3\2\2\2\u016e\u016f\7p\2\2\u016f\u0170\7g\2\2\u0170\u0171\7y\2"+
		"\2\u0171r\3\2\2\2\u0172\u0173\7n\2\2\u0173\u0174\7g\2\2\u0174\u0175\7"+
		"p\2\2\u0175t\3\2\2\2\u0176\u0177\7r\2\2\u0177\u0178\7w\2\2\u0178\u0179"+
		"\7u\2\2\u0179\u017a\7j\2\2\u017av\3\2\2\2\u017b\u017c\7t\2\2\u017c\u017d"+
		"\7g\2\2\u017d\u017e\7o\2\2\u017e\u017f\7q\2\2\u017f\u0180\7x\2\2\u0180"+
		"\u0181\7g\2\2\u0181x\3\2\2\2\u0182\u0183\7e\2\2\u0183\u0184\7q\2\2\u0184"+
		"\u0185\7p\2\2\u0185\u0186\7v\2\2\u0186\u0187\7c\2\2\u0187\u0188\7k\2\2"+
		"\u0188\u0189\7p\2\2\u0189\u018a\7u\2\2\u018az\3\2\2\2\u018b\u018c\7k\2"+
		"\2\u018c\u018d\7p\2\2\u018d\u018e\7u\2\2\u018e\u018f\7g\2\2\u018f\u0190"+
		"\7t\2\2\u0190\u0191\7v\2\2\u0191|\3\2\2\2\u0192\u0193\7e\2\2\u0193\u0194"+
		"\7c\2\2\u0194\u0195\7r\2\2\u0195\u0196\7c\2\2\u0196\u0197\7e\2\2\u0197"+
		"\u0198\7k\2\2\u0198\u0199\7v\2\2\u0199\u019a\7{\2\2\u019a~\3\2\2\2\u019b"+
		"\u019c\7y\2\2\u019c\u019d\7k\2\2\u019d\u019e\7v\2\2\u019e\u019f\7e\2\2"+
		"\u019f\u01a0\7j\2\2\u01a0\u01a1\7a\2\2\u01a1\u01a2\7e\2\2\u01a2\u01a3"+
		"\7c\2\2\u01a3\u01a4\7r\2\2\u01a4\u01a5\7c\2\2\u01a5\u01a6\7e\2\2\u01a6"+
		"\u01a7\7k\2\2\u01a7\u01a8\7v\2\2\u01a8\u01a9\7{\2\2\u01a9\u0080\3\2\2"+
		"\2\u01aa\u01ab\7o\2\2\u01ab\u01ac\7c\2\2\u01ac\u01ad\7k\2\2\u01ad\u01ae"+
		"\7p\2\2\u01ae\u0082\3\2\2\2\u01af\u01b1\t\4\2\2\u01b0\u01af\3\2\2\2\u01b1"+
		"\u01b2\3\2\2\2\u01b2\u01b0\3\2\2\2\u01b2\u01b3\3\2\2\2\u01b3\u0084\3\2"+
		"\2\2\u01b4\u01bc\5\u0083B\2\u01b5\u01b7\7\60\2\2\u01b6\u01b8\5\u0083B"+
		"\2\u01b7\u01b6\3\2\2\2\u01b8\u01b9\3\2\2\2\u01b9\u01b7\3\2\2\2\u01b9\u01ba"+
		"\3\2\2\2\u01ba\u01bd\3\2\2\2\u01bb\u01bd\7\60\2\2\u01bc\u01b5\3\2\2\2"+
		"\u01bc\u01bb\3\2\2\2\u01bd\u0086\3\2\2\2\u01be\u01c3\5\u0089E\2\u01bf"+
		"\u01c2\5\u0089E\2\u01c0\u01c2\5\u0083B\2\u01c1\u01bf\3\2\2\2\u01c1\u01c0"+
		"\3\2\2\2\u01c2\u01c5\3\2\2\2\u01c3\u01c1\3\2\2\2\u01c3\u01c4\3\2\2\2\u01c4"+
		"\u0088\3\2\2\2\u01c5\u01c3\3\2\2\2\u01c6\u01c7\t\5\2\2\u01c7\u008a\3\2"+
		"\2\2\u01c8\u01cc\7$\2\2\u01c9\u01cb\n\6\2\2\u01ca\u01c9\3\2\2\2\u01cb"+
		"\u01ce\3\2\2\2\u01cc\u01ca\3\2\2\2\u01cc\u01cd\3\2\2\2\u01cd\u01cf\3\2"+
		"\2\2\u01ce\u01cc\3\2\2\2\u01cf\u01d0\7$\2\2\u01d0\u008c\3\2\2\2\u01d1"+
		"\u01d5\7)\2\2\u01d2\u01d4\n\6\2\2\u01d3\u01d2\3\2\2\2\u01d4\u01d7\3\2"+
		"\2\2\u01d5\u01d3\3\2\2\2\u01d5\u01d6\3\2\2\2\u01d6\u01d8\3\2\2\2\u01d7"+
		"\u01d5\3\2\2\2\u01d8\u01d9\7)\2\2\u01d9\u008e\3\2\2\2\r\2\u00d3\u00dd"+
		"\u00eb\u01b2\u01b9\u01bc\u01c1\u01c3\u01cc\u01d5\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}