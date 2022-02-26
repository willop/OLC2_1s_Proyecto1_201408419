//gramatica.g4
grammar gramatica;


@parser::header{
    import "proyecto1/Interfaces"
}

// insertar atributos en la clase generada
@parser::members{
}

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


//espacios en blanco
WHITESPACE: [ \r\n\t]+ -> skip;
TK_comentario_multi: '/*' .*? '*/' -> skip;
TK_comentario_lineal: '//' ~[\r\n]* -> skip;


//Palabras reservadas
TKR_numericos_enteros: 'i64';
TKR_numericos_flotantes: 'f64';
TKR_pow: 'pow';  //creo que les falta
TKR_vec: 'vec';  //creo que les falta
TKR_powf: 'powf';
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


//reglas gramaticales
TK_entero: [0-9]+;
TK_decimal: TK_entero('.'TK_entero+|'.');
TK_id: Letra(Letra|TK_entero)*;
Letra : 'a'..'z'|'A'..'Z'|'_' ;	

TK_cadena:  '"'~["]*'"';
TK_caracter:'\''~["]*'\'';
//Rules




start : instrucciones* EOF                                                                              //{fmt.Println("Inicio de gramatica: "+ $expresiones.value)}
;

instrucciones: expresion                                                                {fmt.Println("mensaje en instrucciones: ",$expresion.valorexpresion)}
            |impresion
            |declaracion                                                                {fmt.Println("mensaje en declaracion: ",$declaracion.decla)}
            |identificadores
;

declaracion returns[interface{} decla]: TKR_let TKR_mut TK_id TK_dosPuntos tipovariable igualacion TK_pcoma    
        |TKR_let TKR_mut TK_id igualacion TK_pcoma //mutables
        |TKR_let TK_id TK_dosPuntos tipovariable igualacion TK_pcoma
        |TKR_let idd=TK_id TK_igual expresion TK_pcoma                       {$decla = Interfaces.ConstructorSimbolo($idd.text,50,false,2)}
;

tipovariable: TKR_numericos_enteros
            |TKR_numericos_flotantes
            |TKR_String
            |TKR_bool
            |TKR_char
            |TKR_usize
;

igualacion: TK_igual expresion 
;

identificadores: TK_id igualacion TK_pcoma
;

valores: TK_entero    
        |TK_decimal
        |TK_cadena
        |TK_caracter
        |TKR_true
        |TKR_false
        |TK_id
;

expresion returns [string valorexpresion]
        : TK_menos expresion  
        |e1=expresion op=TK_suma e2=expresion                                      {$valorexpresion=Interfaces.OperacionAritmetica($e1.text,$op.text,$e2.text)}
        |e1=expresion op=TK_menos e2=expresion                                     {$valorexpresion=Interfaces.OperacionAritmetica($e1.text,$op.text,$e2.text)}
        |e1=expresion op=TK_por e2=expresion                                       {$valorexpresion=Interfaces.OperacionAritmetica($e1.text,$op.text,$e2.text)}
        |e1=expresion op=TK_diagonal e2=expresion                                  {$valorexpresion=Interfaces.OperacionAritmetica($e1.text,$op.text,$e2.text)}
        |TKR_pow TK_par_apertura expresion TK_coma expresion TK_par_cierre
        |TKR_powf TK_par_apertura expresion TK_coma expresion TK_par_cierre
        |expresion TK_porcentaje expresion
        |expresion TK_menor expresion
        |expresion TK_mayor expresion
        |expresion TK_mayor_igual expresion
        |expresion TK_menor_igual expresion
        |expresion TK_igualacion expresion
        |expresion TK_diferente expresion
        |expresion TK_or expresion
        |expresion TK_and expresion
        |expresion TK_sig_admiracion expresion
        |TK_par_apertura expresion TK_par_cierre
        |valores TKR_as TKR_numericos_enteros 
        |valores TKR_as TKR_numericos_flotantes 
        |valores
; 

impresion: TKR_println TK_par_apertura expresion TK_par_cierre TK_pcoma                                 {fmt.Println("Impresion")}
        |TKR_println TK_par_apertura expresion impresioncomas TK_pcoma
;

impresioncomas: impresioncomas TK_coma expresion TK_par_cierre TK_pcoma                 {fmt.Println("Impresion especial")}
                |TK_coma expresion TK_par_cierre TK_pcoma                 {fmt.Println("Impresion especial")}

;

//expresiones returns[string value] : TKR_println TK_par_apertura TK_entero TK_par_cierre               {$value = $TK_entero.text;}
//;