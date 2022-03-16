//gramatica.g4
grammar gramatica;


@parser::header{
    import "proyecto1/Interfaces"
    import "proyecto1/Expresion"
    import "proyecto1/Instruccion"
    import "proyecto1/Expresiones"
    import arrayList "github.com/colegno/arraylist"
    //import "proyecto1/Operaciones"
    //import "proyecto1/Expresion"
}

// insertar atributos en la clase generada
@parser::members{
        var temporal = Interfaces.SINTIPO
}


start returns [*arrayList.List listainstrucciones]
        :TKR_fn TKR_main TK_par_apertura TK_par_cierre TK_corchete_apertura instrucciones TK_corchete_cierre   {$listainstrucciones = $instrucciones.lista}

;


instrucciones returns [*arrayList.List lista]
@init{
        $lista = arrayList.New()
}
        :e+=instruccion*        {listaInst := localctx.(*InstruccionesContext).GetE() 
                                    for _, e:= range listaInst {
                                        $lista.Add(e.GetInst())
                                    }
                                }
;


instruccion returns[Interfaces.Instruccion inst]
            : expresion                                                                 {fmt.Println("mensaje en instrucciones: ",$expresion.exp)}
            |impresion                                                                  {$inst = $impresion.inst}
            |declaracion                                                                {$inst = $declaracion.inst}
            |identificadores                                                            {$inst = $identificadores.inst}
            |condicionales                                                              {$inst = $condicionales.inst}
            |bucles                                                                     {$inst = $bucles.inst}
            |control                                                                    {$inst = $control.inst}
;

control returns [Interfaces.Instruccion inst]
        :TKR_break  TK_pcoma                                                                {$inst = Instruccion.NewBreak()}
;

declaracion returns[Interfaces.Instruccion inst]
        //mutables                                                                                                                      nuevadeclaracion   
        : TKR_let TKR_mut idd=TK_id TK_dosPuntos tipovariable TK_igual expresion  TK_pcoma              {$inst = Instruccion.NuevaDeclaracion($idd.text,$tipovariable.tipovar,$expresion.exp,false,false,true) }
        |TKR_let TKR_mut idd=TK_id TK_igual expresion  TK_pcoma                                         {$inst = Instruccion.NuevaDeclaracion($idd.text,Interfaces.SINTIPO,$expresion.exp,false,false,true) }
        //no mutables
        |TKR_let idd=TK_id TK_dosPuntos tipovariable TK_igual expresion  TK_pcoma                       {$inst = Instruccion.NuevaDeclaracion($idd.text,$tipovariable.tipovar,$expresion.exp,false,false,false) }
        |TKR_let idd=TK_id TK_igual expresion TK_pcoma                                                  {$inst = Instruccion.NuevaDeclaracion($idd.text,Interfaces.SINTIPO,$expresion.exp,false,false,false) }
;

tipovariable returns[Interfaces.Tipoexpresion tipovar]
            : TKR_numericos_enteros                             {$tipovar = Interfaces.INTEGER}
            |TKR_numericos_flotantes                            {$tipovar = Interfaces.FLOAT}
            |TKR_String                                         {$tipovar = Interfaces.STRING}
            |TKR_bool                                           {$tipovar = Interfaces.BOOLEAN}
            |TKR_char                                           {$tipovar = Interfaces.CHAR}
            |TKR_usize                                          {$tipovar = Interfaces.USIZE}
;

 
identificadores returns[Interfaces.Instruccion inst]
                : e1=TK_id TK_igual e2=expresion  TK_pcoma             {$inst = Instruccion.NewAsignacion($e1.text,$e2.exp)}                                                             

;


valores returns[Interfaces.Expresion exp]
        :TK_entero             {
                numero, err := strconv.Atoi($TK_entero.text)
                if err != nil {
                        fmt.Println(err)
                }
                $exp = Expresion.NuevoPrimitivo(numero, Interfaces.INTEGER)}//se convierte a entero el texto, el mensaje en error de traduccion, y se agrega el primitivo
        |TK_decimal            {decimal, err := strconv.ParseFloat($TK_decimal.text,8)
                if err != nil {
                        fmt.Println(err)
                }
                $exp = Expresion.NuevoPrimitivo(decimal, Interfaces.FLOAT)}//se convierte a entero el texto, el mensaje en error de traduccion, y se agrega el primitivo
        |TK_cadena             {
                str:= $TK_cadena.text[1:len($TK_cadena.text)-1]
                $exp = Expresion.NuevoPrimitivo(str,Interfaces.STRING)}
        |TK_caracter           
                {str:= $TK_caracter.text[1:len($TK_caracter.text)-1]
                $exp = Expresion.NuevoPrimitivo(str,Interfaces.CHAR)}
        |TKR_Str           
                {str:= $TKR_Str.text[1:len($TKR_Str.text)-1]
                $exp = Expresion.NuevoPrimitivo(str,Interfaces.STR)}
        |TKR_true              {$exp = Expresion.NuevoPrimitivo(true, Interfaces.BOOLEAN)}
        |TKR_false             {$exp = Expresion.NuevoPrimitivo(false,Interfaces.BOOLEAN)}
        |TK_id                 {$exp = Expresiones.NewLlamarvariable($TK_id.text)}
;

expresion returns [Interfaces.Expresion exp]
        : TK_menos e1=expresion                                                                 {numero := -1
                                                                                                        numm := Expresion.NuevoPrimitivo(numero, Interfaces.INTEGER)
                                                                                                                $exp = Expresiones.NuevaAritmetica($e1.exp, numm ,Interfaces.MULTIPLICACION)
                                                                                                }
        |e1=expresion op=(TK_por|TK_diagonal|TK_porcentaje) e2=expresion                        {
                                                                                                        if $op.text == "*" {
                                                                                                                $exp = Expresiones.NuevaAritmetica($e1.exp,$e2.exp,Interfaces.MULTIPLICACION)
                                                                                                        }else if $op.text == "%" {
                                                                                                                $exp = Expresiones.NuevaAritmetica($e1.exp,$e2.exp,Interfaces.MODULO)
                                                                                                        }else{
                                                                                                                $exp = Expresiones.NuevaAritmetica($e1.exp,$e2.exp,Interfaces.DIVISION)
                                                                                                        }
                                                                                                }
        |e1=expresion op=(TK_menos|TK_suma) e2=expresion                                        {
                                                                                                        if $op.text == "+" {
                                                                                                                $exp = Expresiones.NuevaAritmetica($e1.exp,$e2.exp,Interfaces.SUMA)
                                                                                                        }else{
                                                                                                                $exp = Expresiones.NuevaAritmetica($e1.exp,$e2.exp,Interfaces.RESTA)
                                                                                                        }
                                                                                                }
        |TKR_pow TK_par_apertura e1=expresion TK_coma e2=expresion TK_par_cierre                      {$exp = Expresiones.NuevaAritmetica($e1.exp,$e2.exp,Interfaces.POW)}
        |TKR_powf TK_par_apertura e1=expresion TK_coma e2=expresion TK_par_cierre                     {$exp = Expresiones.NuevaAritmetica($e1.exp,$e2.exp,Interfaces.POWF)}
        |e1=expresion TK_menor e2=expresion                                                           {$exp = Expresiones.NuevaRelacional($e1.exp,$e2.exp,Interfaces.MENOR_QUE)}
        |e1=expresion TK_mayor e2=expresion                                                           {$exp = Expresiones.NuevaRelacional($e1.exp,$e2.exp,Interfaces.MAYOR_QUE)}
        |e1=expresion TK_mayor_igual e2=expresion                                                     {$exp = Expresiones.NuevaRelacional($e1.exp,$e2.exp,Interfaces.MAYOR_IGUAL)}
        |e1=expresion TK_menor_igual e2=expresion                                                     {$exp = Expresiones.NuevaRelacional($e1.exp,$e2.exp,Interfaces.MENOR_IGUAL)}
        |e1=expresion TK_igualacion e2=expresion                                                      {$exp = Expresiones.NuevaRelacional($e1.exp,$e2.exp,Interfaces.IGUALDAD)}
        |e1=expresion TK_diferente e2=expresion                                                       {$exp = Expresiones.NuevaRelacional($e1.exp,$e2.exp,Interfaces.DESIGUAL)}
        |e1=expresion TK_or e2=expresion                                                              {$exp = Expresiones.NuevaLogica($e1.exp,$e2.exp,Interfaces.OR)}
        |e1=expresion TK_and e2=expresion                                                             {$exp = Expresiones.NuevaLogica($e1.exp,$e2.exp,Interfaces.AND)}
        |TK_sig_admiracion e1=expresion                                                               {$exp = Expresiones.NuevaLogica($e1.exp,$e1.exp,Interfaces.NOT)}
        |TK_par_apertura va=expresion TK_par_cierre                                                   {$exp = $va.exp}
        |val=valores TKR_as TKR_numericos_enteros                                                     {$exp = Expresiones.NewAsi64($val.exp)}  
        |val=valores TKR_as TKR_numericos_flotantes                                                   {$exp = Expresiones.NewAsf64($val.exp)}
        |vll=expresion TK_punto TKR_abs TK_par_apertura TK_par_cierre                                 {$exp = Expresiones.Newabs($vll.exp)}
        |vll=expresion TK_punto TKR_sqrt
        |vll=expresion TK_punto TKR_to_string
        |vll=expresion TK_punto TKR_clone
        |vall=valores                                                                   {$exp = $vall.exp
                                                                                        fmt.Println($exp)}
; 


impresion returns [Interfaces.Instruccion inst]
        :TKR_println TK_par_apertura e1=expresion TK_par_cierre TK_pcoma                                {$inst = Instruccion.NuevoPrint($e1.exp,nil)}
        |TKR_println TK_par_apertura e2=expresion  li=impresionexpresion TK_pcoma                          {$inst = Instruccion.NuevoPrint($e2.exp,$li.lista)}
;

impresionexpresion returns [*arrayList.List lista]
@init {$lista = arrayList.New()}
                        :list+=expcoma*                                         {
                                                                                        listInt:= localctx.(*ImpresionexpresionContext).GetList()
                                                                                        for _,e := range listInt{
                                                                                                $lista.Add(e.GetExp())
                                                                                        }
                                                                                }
;

expcoma returns[Interfaces.Expresion exp]
        : TK_coma e=expresion                                                           {$exp = $e.exp}


;
condicionales returns[Interfaces.Instruccion inst]
                : funcionif                                                           {$inst = $funcionif.inst}
                //|match
;
/*
match:

;
  */      
        
funcionif returns [Interfaces.Instruccion inst]
        :TKR_if e1=expresion ee=bloque                                                {$inst = Instruccion.NewIf($e1.exp,$ee.lista,nil,nil )}
        |TKR_if e1=expresion e5=bloque TKR_else b2=bloque                             {$inst = Instruccion.NewIf($e1.exp,$e5.lista,nil,$b2.lista)}          
        |TKR_if e1=expresion b1=bloque listaelseif TKR_else b2=bloque                 {$inst = Instruccion.NewIf($e1.exp,$b1.lista,$listaelseif.lista,$b2.lista)}
        |TKR_if e1=expresion b1=bloque listaelseif                                    {$inst = Instruccion.NewIf($e1.exp,$b1.lista,$listaelseif.lista,nil)}
;

funcionelseif returns [Interfaces.Instruccion inst]
        : TKR_elseif e1=expresion bloque                                      {$inst = Instruccion.NewIf($e1.exp,$bloque.lista,nil,nil)}
;

listaelseif returns [*arrayList.List lista]
@init{ $lista = arrayList.New()}
        :list += funcionelseif*{
                                                        listInt := localctx.(*ListaelseifContext).GetList()
                                                        for _,e := range listInt{
                                                                $lista.Add(e.GetInst())
                                                        }
        }
;


bloque returns [*arrayList.List lista]
        :TK_corchete_apertura instrucciones TK_corchete_cierre   {$lista = $instrucciones.lista}
        |TK_corchete_apertura TK_corchete_cierre                 {$lista = arrayList.New()}
;


bucles returns [Interfaces.Instruccion inst]
        :fwhile                                                 {$inst = $fwhile.inst}
        |floop                                                  {$inst = $floop.inst}
;

fwhile returns [Interfaces.Instruccion inst]
        :TKR_while e1=expresion bl=bloque                       {$inst = Instruccion.NewWhile($e1.exp, $bl.lista)}
;


floop returns [Interfaces.Instruccion inst]
        :TKR_loop bb=bloque                                    {$inst = Instruccion.NewLoop($bb.lista)}
;
/* 
ffor:

;
*/





//*********************************************************************************
//****************************Gramatica lexica*************************************
//*********************************************************************************
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
TKR_Str: '&str';
TKR_String: 'String';
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
TKR_main: 'main';
TKR_if: 'if';
TKR_elseif: 'else if';
TKR_else: 'else';
TKR_while: 'while';
TKR_break: 'break';
TKR_loop: 'loop';


//reglas gramaticales
TK_entero: [0-9]+;
TK_decimal: TK_entero('.'TK_entero+|'.');
TK_id: Letra(Letra|TK_entero)*;
Letra : 'a'..'z'|'A'..'Z'|'_' ;	

TK_cadena:  '"'~["]*'"';
TK_caracter:'\''~["]*'\'';
//Rules
