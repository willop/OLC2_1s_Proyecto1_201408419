//gramatica.g4
grammar gramatica;

//Tokens
//retorno up presedencia
TK_flecha: '->';

//condicionales up presedencia
TK_or: '||';
TK_and: '&&';
TK_igualacion: '==';
TK_diferente: '!=';

//signos de comparacion up presedencia
TK_mayor_igual: '>=';
TK_menor_igual: '<=';

TK_corchete_apertura: '{';
TK_corchete_cierre: '}';
TK_llave_apertura: '[';
TK_llave_cierre: ']';
TK_dosPuntos: ':';
TK_coma: ',';
TK_pcoma: ';';
TK_menor: '<';
TK_mayor: '>';
TK_punto: '.';
TK_par_apertura: '(';
TK_par_cierre: ')';
TK_igual: '=';
TK_suma: '+';
TK_menos: '-';
TK_por: '*';
TK_diagonal: '/';
TK_porcentaje: '%';
TK_linea: '|';
TK_amp: '&';
TK_sig_admiracion: '!';
TK_sig_interrogacion: '?';


//reglas gramaticales
TK_entero: [0-9]+;
TK_decimal: TK_entero('.'TK_entero+|'.');
TK_id: Letra (Letra | TK_entero)*;
Letra : 'a'..'z'|'A'..'Z'|'_' ;	

TK_cadena:  '"'~["]*'"';
TK_caracter:'\''~["]*'\'';


//espacios en blanco
WHITESPACE: [\r\n\t]+ -> skip;
TK_comentario_multi: '/*' .*? '*/' -> skip;
TK_comentario_lineal: '//' ~[\r\n]* -> skip;


//Palabras reservadas
TKR_numericos_enteros: 'i64';
TKR_numericos_flotantes: 'f64';
TKR_bool: 'bool';
TKR_char: 'char';
TKR_String: '&str'|'String';
TKR_usize: 'usize';
TKR_let: 'let';
TKR_mut: 'mut';
TKR_struct: 'struct';
TKR_as: 'as';
TKR_println: 'println!';
TKR_true: 'true';
TKR_false: 'false';
TKR_fn: 'fn';
TKR_return: 'return';
TKR_abs: 'abs';
TKR_sqrt: 'sqrt';
TKR_to_string:'to_string()';
TKR_clone: 'clone()';
TKR_new: 'new';
TKR_len: 'len';
TKR_push: 'push';
TKR_remove: 'remove';
TKR_contains: 'contains';
TKR_insert: 'insert';
TKR_capacity: 'capacity';
TKR_with_capacity: 'witch_capacity';


//Rules
start : expresiones EOF;

expresiones: TKR_println TK_par_apertura TK_entero TK_par_cierre
;

