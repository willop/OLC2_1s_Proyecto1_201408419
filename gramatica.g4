//gramatica.g4
grammar gramatica;


@parser::header{
    import "proyecto1/Interfaces"
    import "proyecto1/Expresion"
    import "proyecto1/Instruccion"
    //import arrayList "github.com/colegno/arraylist"
    //import "proyecto1/Expresion"
}

// insertar atributos en la clase generada
@parser::members{
        var temporal = Interfaces.SINTIPO
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




start //returns [*arrayList.List listainstrucciones]
:instrucciones  EOF                                                                    //{$listainstrucciones = $instrucciones.lista}                                                        
;

instrucciones //returns [*arrayList.List lista]
/*@init{
        $lista = arrayList.New()
}*/
        :e+=instruccion*                                                                /*{listaInst := localctx.(*InstruccionesContext).GetE() 
                                                                                                for _, e:= range listaInst {
                                                                                                        $lista.Add(e.GetInstr())
                                                                                                }
                                                                                        }*/
;

instruccion returns[Interfaces.Instruccion inst]
            : expresion                                                                {fmt.Println("mensaje en instrucciones: ",$expresion.valorexpresion)}
            |impresion
            |declaracion                                                                {fmt.Println("mensaje en declaracion: ",$declaracion.decla)}
            //|identificadores
            //|condicionales
;

declaracion returns[interface{} decla]
        //mutables                                                                                                                      nuevadeclaracion   
        : TKR_let TKR_mut idd=TK_id TK_dosPuntos tipovariable TK_igual expresion  TK_pcoma              {$decla = Instruccion.NuevaDeclaracion($idd.text,$tipovariable.tipovar,$expresion.valorexpresion,true,false,false) }
        |TKR_let TKR_mut idd=TK_id TK_igual expresion  TK_pcoma                                         {$decla = $expresion.valorexpresion}    
        //no mutables
        |TKR_let idd=TK_id TK_dosPuntos tipovariable TK_igual expresion  TK_pcoma                       {$decla = Instruccion.NuevaDeclaracion($idd.text,$tipovariable.tipovar,$expresion.valorexpresion,false,false,false) }
        |TKR_let idd=TK_id TK_igual expresion TK_pcoma                                                  {$decla = $expresion.valorexpresion}
;

tipovariable returns[Interfaces.Tipoexpresion tipovar]
            : TKR_numericos_enteros                             {$tipovar = Interfaces.INTEGER}
            |TKR_numericos_flotantes                            {$tipovar = Interfaces.FLOAT}
            |TKR_String                                         {$tipovar = Interfaces.STRING}
            |TKR_bool                                           {$tipovar = Interfaces.BOOLEAN}
            |TKR_char                                           {$tipovar = Interfaces.CHAR}
            |TKR_usize                                          {$tipovar = Interfaces.USIZE}
;

/* 
identificadores: TK_id TK_igual expresion  TK_pcoma                                                                          

;
*/

valores returns[Interfaces.Expresion lit]
        :TK_entero             {
                numero, err := strconv.Atoi($TK_entero.text)
                if err != nil {
                        fmt.Println(err)
                }
                $lit = Expresion.NuevoPrimitivo(numero, Interfaces.INTEGER)}//se convierte a entero el texto, el mensaje en error de traduccion, y se agrega el primitivo
        |TK_decimal            {decimal, err := strconv.ParseFloat($TK_decimal.text,8)
                if err != nil {
                        fmt.Println(err)
                }
                $lit = Expresion.NuevoPrimitivo(decimal, Interfaces.FLOAT)}//se convierte a entero el texto, el mensaje en error de traduccion, y se agrega el primitivo
        |TK_cadena             {
                str:= $TK_cadena.text[1:len($TK_cadena.text)-1]
                $lit = Expresion.NuevoPrimitivo(str,Interfaces.STRING)}
        |TK_caracter           
                {str:= $TK_caracter.text[1:len($TK_caracter.text)-1]
                $lit = Expresion.NuevoPrimitivo(str,Interfaces.CHAR)}
        |TKR_true              {$lit = Expresion.NuevoPrimitivo(true, Interfaces.BOOLEAN)}
        |TKR_false             {$lit = Expresion.NuevoPrimitivo(false,Interfaces.BOOLEAN)}
        |TK_id
;

expresion returns [Interfaces.Expresion valorexpresion]
        : TK_menos expresion  
        |e1=expresion op=TK_por e2=expresion                                                                                  
        |e1=expresion op=TK_menos e2=expresion                                     
        |e1=expresion op=TK_diagonal e2=expresion                                  
        |e1=expresion op=TK_suma e2=expresion                                      
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
        |vall=valores                                                                   {$valorexpresion = $vall.lit
                                                                                        fmt.Println($valorexpresion)}
; 



impresion returns [Interfaces.Instruccion impr]
        :TKR_println TK_par_apertura expresion TK_par_cierre TK_pcoma                                 {fmt.Println("Impresion")
                                                                                                        $impr = Instruccion.NuevoPrint($expresion.valorexpresion)}
        |TKR_println TK_par_apertura expresion impresioncomas TK_pcoma
;

impresioncomas: impresioncomas TK_coma expresion TK_par_cierre TK_pcoma                 {fmt.Println("Impresion especial")}
                |TK_coma expresion TK_par_cierre TK_pcoma                 {fmt.Println("Impresion especial")}
;

//expresiones returns[string value] : TKR_println TK_par_apertura TK_entero TK_par_cierre               {$value = $TK_entero.text;}
//;