grammar SwiftGrammar;

options {
	tokenVocab = SwiftLexer;
}

@header {
    import "Backend/interfaces"
    import "Backend/environment"
    import "Backend/expressions"
    import "Backend/instructions/datoscompuestos"
    import "Backend/instructions/datosprimitivos"
    import "Backend/instructions/funciones"
    import "Backend/instructions/sentencias"
    import "strings"
   
}

s returns [[]interface{} code]
: block EOF
    {   
        $code = $block.blk
    }
;

block returns [[]interface{} blk]
@init{
    $blk = []interface{}{}
    var listInt []IInstructionContext
  }
: ins+=instruction+
    {
        listInt = localctx.(*BlockContext).GetIns()
        for _, e := range listInt {
            $blk = append($blk, e.GetInst())
        }
    }
;

// LISTA DE INSTRUCCIONES GENERALES O GLOBALES
instruction returns [interfaces.Instruction inst]
: declavarible (PUNTOCOMA)? { $inst = $declavarible.decvbl}
| declaconstante (PUNTOCOMA)? { $inst = $declaconstante.deccon}
| asignacionvariable (PUNTOCOMA)? { $inst = $asignacionvariable.asgvbl}
| sentenciaifelse { $inst = $sentenciaifelse.myIfElse}
| switchcontrol { $inst = $switchcontrol.mySwitch}
| whilecontrol { $inst = $whilecontrol.whict}
| forcontrol { $inst = $forcontrol.forct}
| guardcontrol { $inst = $guardcontrol.guct}
| vectorcontrol (PUNTOCOMA)? { $inst = $vectorcontrol.vect }
| vectoragregar  { $inst = $vectoragregar.veadct }
| vectorremover  { $inst = $vectorremover.vermct }
| printstmt (PUNTOCOMA)? { $inst = $printstmt.prnt}
| matrizcontrol (PUNTOCOMA)? { $inst = $matrizcontrol.matct}
// | structcontrol { $inst = $structcontrol.struck}
| funciondeclaracioncontrol { $inst = $funciondeclaracioncontrol.fdc}
| funcionllamadacontrol { $inst = $funcionllamadacontrol.flctl}
// | structexpr (PUNTOCOMA)? { $inst = $structexpr.strexpr}
// | asignacionparametrostruct (PUNTOCOMA)? { $inst = $asignacionparametrostruct.llmstruasig}
// | llamadafuncionstructcontrol (PUNTOCOMA)? { $inst = $llamadafuncionstructcontrol.llmstrufun}
;

// LISTA DE INSTRUCCIONES LOCALES
blockinterno returns [[]interface{} blkint]
@init{
    $blkint = []interface{}{}
    var listInt []IInstructionintContext
  }
: insint+=instructionint+
    {
        listInt = localctx.(*BlockinternoContext).GetInsint()
        for _, e := range listInt {
            $blkint = append($blkint, e.GetInsint())
        }
    }
;


// LISTA DE INSTRUCCIONES LOCALES
instructionint returns [interfaces.Instruction insint]
: declavarible (PUNTOCOMA)? { $insint = $declavarible.decvbl}
| declaconstante (PUNTOCOMA)? { $insint = $declaconstante.deccon}
| asignacionvariable (PUNTOCOMA)? { $insint = $asignacionvariable.asgvbl}
| sentenciaifelse { $insint = $sentenciaifelse.myIfElse}
| switchcontrol { $insint = $switchcontrol.mySwitch}
| whilecontrol { $insint = $whilecontrol.whict}
| forcontrol { $insint = $forcontrol.forct}
| guardcontrol { $insint = $guardcontrol.guct}
| continuee (PUNTOCOMA)? { $insint = $continuee.coct}
| breakk (PUNTOCOMA)? { $insint = $breakk.brkct}
| retornos (PUNTOCOMA)? { $insint = $retornos.rect }
| vectorcontrol (PUNTOCOMA)? { $insint = $vectorcontrol.vect }
| vectoragregar  (PUNTOCOMA)? { $insint = $vectoragregar.veadct }
| vectorremover (PUNTOCOMA)? { $insint = $vectorremover.vermct }
| printstmt (PUNTOCOMA)? { $insint = $printstmt.prnt}
| matrizcontrol (PUNTOCOMA)? { $insint = $matrizcontrol.matct}
| funcionllamadacontrol { $insint = $funcionllamadacontrol.flctl}
// | structexpr (PUNTOCOMA)? { $insint = $structexpr.strexpr}
// | asignacionparametrostruct (PUNTOCOMA)? { $insint = $asignacionparametrostruct.llmstruasig}
// | llamadafuncionstructcontrol (PUNTOCOMA)? { $insint = $llamadafuncionstructcontrol.llmstrufun}
;

/////////////////////////
/////////////////////////
///     GENERALES     ///
/////////////////////////
/////////////////////////

// DECLARACION DE VARIABLES
declavarible returns [interfaces.Instruction decvbl]
: VAR ID_VALIDO DOS_PUNTOS tipodato IG expr { $decvbl = datosprimitivos.NewVariableDeclaration($VAR.line, $VAR.pos, $ID_VALIDO.text, $tipodato.tipo, $expr.e)}
| VAR ID_VALIDO IG expr {$decvbl = datosprimitivos.NewVariableDeclaracionSinTipo($VAR.line, $VAR.pos, $ID_VALIDO.text, $expr.e)}
| VAR ID_VALIDO DOS_PUNTOS tipodato CIERRE_INTE {$decvbl = datosprimitivos.NewVariableDeclaracionSinExp($VAR.line, $VAR.pos, $ID_VALIDO.text, $tipodato.tipo)}
;

// DECLARACION DE CONSTANTES
declaconstante returns [interfaces.Instruction deccon]
: LET ID_VALIDO DOS_PUNTOS tipodato IG expr {$deccon = datosprimitivos.NewConstanteDeclaration($LET.line, $LET.pos, $ID_VALIDO.text, $tipodato.tipo, $expr.e)}
| LET ID_VALIDO IG expr {$deccon = datosprimitivos.NewConstanteDeclaracionSinTipo($LET.line, $LET.pos, $ID_VALIDO.text, $expr.e)}
;

// ASIGNACION DE VARIABLES
asignacionvariable returns [interfaces.Instruction asgvbl]
: ID_VALIDO IG expr { $asgvbl = sentencias.NewAsignacionVariable($ID_VALIDO.line, $ID_VALIDO.pos, $ID_VALIDO.text, $expr.e)}
| ID_VALIDO SUMA expr { $asgvbl = sentencias.NewAsignacionSuma($ID_VALIDO.line, $ID_VALIDO.pos, $ID_VALIDO.text, $expr.e)}
| ID_VALIDO RESTA expr { $asgvbl = sentencias.NewAsignacionResta($ID_VALIDO.line, $ID_VALIDO.pos, $ID_VALIDO.text, $expr.e)}
;

// TIPOS DE DATOS
tipodato returns [environment.TipoExpresion tipo]
: INT { $tipo = environment.INTEGER }
| FLOAT { $tipo = environment.FLOAT }
| STRING { $tipo = environment.STRING }
| BOOL { $tipo = environment.BOOLEAN }
| CHARACT { $tipo = environment.CHARACTER };

// EXPRESION
expr returns [interfaces.Expression e]
: op=NOT right=expr { $e = expressions.NewOperation($right.start.GetLine(), $right.start.GetColumn(), $right.e, $op.text, $right.e) }
| left=expr op=MODULO right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MUL|DIV) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(ADD|SUB) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MAY_IG|MAYOR) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MEN_IG|MENOR) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(IG_IG|DIF) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=AND right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=OR right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| PARIZQ expr PARDER { $e = $expr.e }
| SUB NUMBER                             
    {
        if (strings.Contains($NUMBER.text,".")){
            num,err := strconv.ParseFloat($NUMBER.text, 64);
            if err!= nil{
                fmt.Println(err)
            }
	        num2 := fmt.Sprintf("%.6f", num)
            num3,err := strconv.ParseFloat(num2, 64);
            if err!= nil{
                fmt.Println(err)
            }
            $e = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,-num3,environment.FLOAT)
        }else{
            num,err := strconv.Atoi($NUMBER.text)
            if err!= nil{
                fmt.Println(err)
            }
            $e = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,-num,environment.INTEGER)
        }
    }
| NUMBER                             
    {
        if (strings.Contains($NUMBER.text,".")){
            num,err := strconv.ParseFloat($NUMBER.text, 64);
            if err!= nil{
                fmt.Println(err)
            }
	        num2 := fmt.Sprintf("%.6f", num)
            num3,err := strconv.ParseFloat(num2, 64);
            if err!= nil{
                fmt.Println(err)
            }
	        $e = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num3,environment.FLOAT)
        }else{
            num,err := strconv.Atoi($NUMBER.text)
            if err!= nil{
                fmt.Println(err)
            }            
            $e = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.INTEGER)
        }
    }
| CADENA
    {
        str := $CADENA.text
        $e = expressions.NewPrimitive($CADENA.line, $CADENA.pos, str[1:len(str)-1],environment.STRING)
    }                        
| TRU { $e = expressions.NewPrimitive($TRU.line, $TRU.pos, true, environment.BOOLEAN) }
| FAL { $e = expressions.NewPrimitive($FAL.line, $FAL.pos, false, environment.BOOLEAN) }
| CHARACTER 
    { 
        str := $CHARACTER.text
        $e = expressions.NewPrimitive($CHARACTER.line, $CHARACTER.pos, str[1:len(str)-1], environment.CHARACTER) 
    }
|ID_VALIDO
    {
        id := $ID_VALIDO.text
        $e = sentencias.NewCallid($ID_VALIDO.line,$ID_VALIDO.pos,id)
    }
|NULO {$e = expressions.NewPrimitive($NULO.line, $NULO.pos, $NULO.text,environment.NULL)}
| vectorvacio { $e = $vectorvacio.veemct}
| vectorcount { $e = $vectorcount.vecnct}
| vectoraccess { $e = $vectoraccess.vepposct}
| intembebida { $e = $intembebida.intemb}
| floatembebida { $e = $floatembebida.floemb}
| stringembebida { $e = $stringembebida.stremb}
// | funcionllamadacontrolConRetorno { $e = $funcionllamadacontrolConRetorno.flctlret}
// | llamadastruct { $e = $llamadastruct.llmstru}
// | llamadafuncionstructcontrolret { $e = $llamadafuncionstructcontrolret.llmstrufunret}
;

// CREACION DE IF-ELSE
sentenciaifelse returns [interfaces.Instruction myIfElse]
: IF expr LLAVEIZQ blockinterno LLAVEDER { $myIfElse = sentencias.NewSentenciaIf($IF.line, $IF.pos, $expr.e, $blockinterno.blkint)}
| IF expr LLAVEIZQ ifop=blockinterno LLAVEDER ELSE LLAVEIZQ elseop=blockinterno LLAVEDER { $myIfElse = sentencias.NewSentenciaIfElse($IF.line, $IF.pos, $expr.e, $ifop.blkint , $elseop.blkint)}
| IF expr LLAVEIZQ blockinterno LLAVEDER ELSE sentenciaifelse { $myIfElse = sentencias.NewSentenciaIfElseIf($IF.line, $IF.pos, $expr.e, $blockinterno.blkint, $sentenciaifelse.myIfElse)}
;

// CREACION DEL SWITCH
switchcontrol returns [interfaces.Instruction mySwitch]
: SWITCH expr LLAVEIZQ blockcase (DEFAULT DOS_PUNTOS blockinterno)? LLAVEDER 
{
    if ($DEFAULT != nil) {
        $mySwitch = sentencias.NewSentenciaSwitchDefault($SWITCH.line, $SWITCH.pos, $expr.e, $blockcase.blkcase, $blockinterno.blkint)
    } else {
        $mySwitch = sentencias.NewSentenciaSwitch($SWITCH.line, $SWITCH.pos, $expr.e, $blockcase.blkcase)
    }
}
;

blockcase returns [[]interface{} blkcase]
@init{
    $blkcase = []interface{}{}
    var listInt []IBloquecaseContext
}
: blocas+=bloquecase+
{
    listInt = localctx.(*BlockcaseContext).GetBlocas()
    for _, e := range listInt {
        $blkcase = append($blkcase, e.GetBlocas())
    }
}
;

bloquecase returns [interfaces.Instruction blocas]
: CASE expr DOS_PUNTOS blockinterno 
{
    $blocas=sentencias.NewSentenciaSwitchCase($CASE.line ,$CASE.pos, $expr.e, $blockinterno.blkint)
}
;

// CREACION DEL WHILE
whilecontrol returns [interfaces.Instruction whict]
: WHILE expr LLAVEIZQ blockinterno LLAVEDER { $whict = sentencias.NewSentenciaWhile($WHILE.line, $WHILE.pos, $expr.e, $blockinterno.blkint)};

//CREACION DEL FOR
forcontrol returns [interfaces.Instruction forct]
: FOR ID_VALIDO IN left=expr RANGO right=expr LLAVEIZQ blockinterno LLAVEDER { $forct = sentencias.NewSentenciaForRango($FOR.line, $FOR.pos, $ID_VALIDO.text, $left.e, $right.e,$blockinterno.blkint)}
| FOR op1=ID_VALIDO IN op2=ID_VALIDO LLAVEIZQ blockinterno LLAVEDER { $forct = sentencias.NewSentenciaForId($FOR.line, $FOR.pos, $op1.text, $op2.text, $blockinterno.blkint)}
| FOR ID_VALIDO IN expr LLAVEIZQ blockinterno LLAVEDER { $forct = sentencias.NewSentenciaForCadena($FOR.line, $FOR.pos, $ID_VALIDO.text, $expr.e, $blockinterno.blkint)}
;
 
 //CREACION DE GUARD
guardcontrol returns [interfaces.Instruction guct]
:GUARD expr ELSE LLAVEIZQ blockinterno LLAVEDER  
{ 
    $guct = sentencias.NewSentenciaGuard($GUARD.line, $GUARD.pos, $expr.e, $blockinterno.blkint)
}
;

//CREACION DEL CONTINUE
continuee returns [interfaces.Instruction coct]
: CONTINUE {$coct = sentencias.NewTransferenciaContinue($CONTINUE.line, $CONTINUE.pos)};

//CREACION DEL BREAK
breakk returns [interfaces.Instruction brkct]
: BREAK { $brkct = sentencias.NewTransferenciaBreak($BREAK.line, $BREAK.pos)};

//CREACION DEL RETURN
retornos returns [interfaces.Instruction rect]
: RETURN op=expr
{
    $rect = sentencias.NewTransferenciaReturnExp($RETURN.line, $RETURN.pos, $op.e);
}
|RETURN
{
    $rect = sentencias.NewTransferenciaReturn($RETURN.line, $RETURN.pos);
}
;


//CREACION DEL VECTOR 
vectorcontrol returns [interfaces.Instruction vect]
: VAR ID_VALIDO DOS_PUNTOS CORCHIZQ tipodato CORCHDER IG CORCHIZQ blockparams CORCHDER { $vect = datoscompuestos.NewArregloDeclaracionLista($VAR.line ,$VAR.pos, $ID_VALIDO.text , $tipodato.tipo, $blockparams.blkpar)}
| VAR ID_VALIDO DOS_PUNTOS CORCHIZQ tipodato CORCHDER IG CORCHIZQ CORCHDER { $vect = datoscompuestos.NewArregloDeclaracionSinLista($VAR.line ,$VAR.pos, $ID_VALIDO.text , $tipodato.tipo)}
| VAR prin=ID_VALIDO DOS_PUNTOS CORCHIZQ tipodato CORCHDER IG secu=ID_VALIDO { $vect = datoscompuestos.NewArregloDeclaracionId($VAR.line ,$VAR.pos, $prin.text , $tipodato.tipo, $secu.text)}
;

blockparams returns [[]interface{} blkpar]
@init{
    $blkpar = []interface{}{}
    var listInt []IBloqueparamsContext
}
: blopas+=bloqueparams+
{
    listInt = localctx.(*BlockparamsContext).GetBlopas()
    for _, e := range listInt {
        $blkpar = append($blkpar, e.GetBlopas())
    }
}
;

bloqueparams returns [interfaces.Expression blopas]
: COMA expr 
{
    $blopas = datoscompuestos.NewArregloParametros($COMA.line ,$COMA.pos, $expr.e)
}
| expr 
{
    $blopas = datoscompuestos.NewArregloParametro($expr.e)
};

vectoragregar returns [interfaces.Instruction veadct]
: ID_VALIDO PUNTO APPEND PARIZQ expr PARDER { $veadct = datoscompuestos.NewArregloAppend($ID_VALIDO.text , $expr.e)}
| prin=ID_VALIDO CORCHIZQ pop=expr CORCHDER IG secu=ID_VALIDO CORCHIZQ sop=expr CORCHDER { $veadct = datoscompuestos.NewArregloAppendArreglo($prin.text , $pop.e, $secu.text, $sop.e)}
| ID_VALIDO CORCHIZQ op1=expr CORCHDER CORCHIZQ op2=expr CORCHDER listamatrizaddsubs IG op3=expr
{ $veadct = datoscompuestos.NewMatrizAsignacionList($ID_VALIDO.text, $op1.e, $op2.e, $listamatrizaddsubs.blklimatas, $op3.e) }
| ID_VALIDO CORCHIZQ op1=expr CORCHDER CORCHIZQ op2=expr CORCHDER IG op3=expr
{ $veadct = datoscompuestos.NewMatrizAsignacion($ID_VALIDO.text, $op1.e, $op2.e, $op3.e) } 
|ID_VALIDO CORCHIZQ pop=expr CORCHDER IG sop=expr { $veadct = datoscompuestos.NewArregloAppendExp($ID_VALIDO.text , $pop.e, $sop.e)}
;

vectorremover returns [interfaces.Instruction vermct]
: ID_VALIDO PUNTO REMOVELAST PARIZQ PARDER  { $vermct = datoscompuestos.NewArregloRemoveLast($PUNTO.line, $PUNTO.pos, $ID_VALIDO.text)}
| ID_VALIDO PUNTO REMOVE PARIZQ AT DOS_PUNTOS expr PARDER { $vermct = datoscompuestos.NewArregloRemovePos($PUNTO.line, $PUNTO.pos, $ID_VALIDO.text, $expr.e)}
;

vectorvacio returns [interfaces.Expression veemct]
: ID_VALIDO PUNTO ISEMPTY { $veemct = datoscompuestos.NewArregloIsEmpty($PUNTO.line, $PUNTO.pos, $ID_VALIDO.text)};

vectorcount returns [interfaces.Expression vecnct]
: ID_VALIDO PUNTO COUNT { $vecnct = datoscompuestos.NewArregloCount($PUNTO.line, $PUNTO.pos, $ID_VALIDO.text)};

vectoraccess returns [interfaces.Expression vepposct]
: ID_VALIDO CORCHIZQ op1=expr CORCHDER CORCHIZQ op2=expr CORCHDER listamatrizaddsubs
{ $vepposct = datoscompuestos.NewMatrizObtencionList($ID_VALIDO.text, $op1.e, $op2.e, $listamatrizaddsubs.blklimatas) }
| ID_VALIDO CORCHIZQ op1=expr CORCHDER CORCHIZQ op2=expr CORCHDER 
{ $vepposct = datoscompuestos.NewMatrizObtencion($ID_VALIDO.text, $op1.e, $op2.e) } 
| ID_VALIDO CORCHIZQ expr CORCHDER { $vepposct = datoscompuestos.NewArregloAccess($CORCHDER.line, $CORCHDER.pos, $ID_VALIDO.text, $expr.e)}
;

//CREACION DE MATRICES
matrizcontrol returns [interfaces.Instruction matct]
: VAR ID_VALIDO (DOS_PUNTOS tipomatriz)? IG defmatriz
{
    if ($DOS_PUNTOS != nil) {
        $matct = datoscompuestos.NewMatrizDeclaracion($VAR.line, $VAR.pos, $ID_VALIDO.text ,$tipomatriz.tipomat, $defmatriz.defmat)
    } else {
        fmt.Println("Nada")
        //$matct = datoscompuestos.NewMatrizDeclaracionSinTipo($VAR.line, $VAR.pos, $ID_VALIDO.text , $defmatriz.defmat)
    }
}
;

tipomatriz returns [interfaces.Expression tipomat]
: CORCHIZQ tipomatriz CORCHDER 
{ 
    $tipomat = datoscompuestos.NewMatrizDimension($CORCHIZQ.line, $CORCHIZQ.pos, $tipomatriz.tipomat)
}
| CORCHIZQ tipodato CORCHDER 
{ 
    $tipomat = datoscompuestos.NewMatrizTipo($CORCHIZQ.line, $CORCHIZQ.pos, $tipodato.tipo)
}
;

defmatriz returns [interfaces.Instruction defmat]
: listavaloresmat { $defmat = $listavaloresmat.listvlamat}
;

listavaloresmat returns [interfaces.Instruction listvlamat]
: CORCHIZQ listavaloresmat2 CORCHDER { $listvlamat = $listavaloresmat2.mylisttmatt}
| simplematriz { $listvlamat = $simplematriz.simmat}
;

listavaloresmat2 returns [interfaces.Instruction mylisttmatt]
: op=listavaloresmat2 COMA listavaloresmat { $mylisttmatt = datoscompuestos.NewMatrizListaExpresionList($op.mylisttmatt, $listavaloresmat.listvlamat)}
| listavaloresmat { $mylisttmatt = datoscompuestos.NewMatrizListaNivel($listavaloresmat.listvlamat)}
| listaexpresions { $mylisttmatt = datoscompuestos.NewMatrizListaExpresion($listaexpresions.blkparf)}
;

listaexpresions returns [[]interface{} blkparf]
@init{
    $blkparf = []interface{}{}
    var listInt []IListaexpresionContext
}
: funpar+=listaexpresion+
{
    listInt = localctx.(*ListaexpresionsContext).GetFunpar()
    for _, e := range listInt {
        $blkparf = append($blkparf, e.GetFunpar())
    }
}
;

listaexpresion returns [interfaces.Expression funpar]
: COMA expr 
{
    $funpar = datoscompuestos.NewArregloParametros($COMA.line ,$COMA.pos, $expr.e)
}
| expr 
{
    $funpar = datoscompuestos.NewArregloParametro($expr.e)
}
;

simplematriz returns [interfaces.Instruction simmat]
: tipomatriz PARIZQ REPEATING DOS_PUNTOS op=simplematriz COMA COUNT DOS_PUNTOS NUMBER PARDER 
{ $simmat = datoscompuestos.NewMatrizSimpleUno($tipomatriz.tipomat, $op.simmat, $NUMBER.text, $NUMBER.line,$NUMBER.pos)}
| tipomatriz PARIZQ REPEATING DOS_PUNTOS expr COMA COUNT DOS_PUNTOS NUMBER PARDER 
{ $simmat = datoscompuestos.NewMatrizSimpleDos($tipomatriz.tipomat, $expr.e, $NUMBER.text, $NUMBER.line,$NUMBER.pos)}
;

listamatrizaddsubs returns [[]interface{} blklimatas]
@init{
    $blklimatas = []interface{}{}
    var listInt []IListamatrizaddsubContext
}
: lmas+=listamatrizaddsub+
{
    listInt = localctx.(*ListamatrizaddsubsContext).GetLmas()
    for _, e := range listInt {
        $blklimatas = append($blklimatas, e.GetLmas())
    }
}
;

listamatrizaddsub returns [interfaces.Expression lmas]
: CORCHIZQ expr CORCHDER 
{
    $lmas = datoscompuestos.NewArregloParametros($CORCHIZQ.line ,$CORCHIZQ.pos, $expr.e)
}
;

// //CREACION DEL STRUCT

// structcontrol returns [interfaces.Instruction struck]
// : STRUCT ID_VALIDO LLAVEIZQ listaatributos LLAVEDER 
// {
//     $struck = instructions.NewStruck($STRUCT.line, $STRUCT.pos, $ID_VALIDO.text, $listaatributos.blkstlt);
// };

// listaatributos returns [[]interface{} blkstlt]
// @init{
//     $blkstlt = []interface{}{}
//     var listInt []IListaatributoContext
// }
// : listatstr+=listaatributo+
// {
//     listInt = localctx.(*ListaatributosContext).GetListatstr()
//     for _, e := range listInt {
//         $blkstlt = append($blkstlt, e.GetListatstr())
//     }
// }
// ;

// listaatributo returns [interfaces.Instruction listatstr]
// : tip1=(LET | VAR) tip4=ID_VALIDO DOS_PUNTOS (tip2=tipodato|tip3=ID_VALIDO) (IG expr)? (PUNTOCOMA)? 
// {
//     if $IG != nil{
//         if $tip3.text != "" {
//             $listatstr = instructions.NewStructAtributosConTE2($tip1.line, $tip1.pos, $tip1.text, $tip4.text, $tip3.text, $expr.e)
//         }else{                        
//             $listatstr = instructions.NewStructAtributosConTE($tip1.line, $tip1.pos, $tip1.text, $tip4.text, $tip2.tipo, $expr.e)
//         }        
//     }else{ 
//         if $tip3.text != "" {                        
//             $listatstr = instructions.NewStructAtributosConT2($tip1.line, $tip1.pos, $tip1.text, $tip4.text, $tip3.text) 
//         }else{            
//             $listatstr = instructions.NewStructAtributosConT($tip1.line, $tip1.pos, $tip1.text, $tip4.text, $tip2.tipo) 
//         }
//     }
// }
// | tipo=(LET | VAR) ID_VALIDO (IG expr)? (PUNTOCOMA)? 
// {
//     if $IG != nil{
//         $listatstr = instructions.NewStructAtributosConE($tipo.line, $tipo.pos, $tipo.text, $ID_VALIDO.text, $expr.e)
//     }else{
//         $listatstr = instructions.NewStructAtributos($tipo.line, $tipo.pos, $tipo.text, $ID_VALIDO.text)
//     }
// }
// | (MUTATING)?  funciondeclaracioncontrol 
// {
//     if $MUTATING != nil{
//         $listatstr = instructions.NewStruckFunctionMutating($funciondeclaracioncontrol.fdc)
//     } else {
//         $listatstr = instructions.NewStruckFunction($funciondeclaracioncontrol.fdc)
//     }
// }
// ;

// //asignacion del struct
// structexpr returns [interfaces.Instruction strexpr]
// : op1=ID_VALIDO DOS_PUNTOS op=ID_VALIDO op2=ID_VALIDO PARIZQ ldupla PARDER
// {
//     $strexpr = instructions.NewStruckVariable($op1.line, $op1.pos, $op.text, $op1.text, $op2.text, $ldupla.lduplist, true)
// }
// ;

// ldupla returns [interfaces.Instruction lduplist]
// : ID_VALIDO DOS_PUNTOS expr COMA op=ldupla
// { 
//     $lduplist = instructions.NewStructListDuple($ID_VALIDO.text, $expr.e, $op.lduplist, true)
// }
// | ID_VALIDO DOS_PUNTOS expr
// {
//     $lduplist = instructions.NewStructDuple($ID_VALIDO.text, $expr.e, false)  
// }
// ;

// //asignacion y obtencio de struct
// llamadastruct returns [interfaces.Expression llmstru]
// : op=ID_VALIDO PUNTO op1=ID_VALIDO
// {
//     $llmstru = instructions.NewStruckLlamadaExp($op.line, $op.pos, $op.text, $op1.text)
// }
// ;

// asignacionparametrostruct returns [interfaces.Instruction llmstruasig]
// : op=ID_VALIDO PUNTO op1=ID_VALIDO IG expr
// {
//     $llmstruasig = instructions.NewStruckAsignacionExpr($op.line, $op.pos, $op.text, $op1.text, $expr.e)
// }
// ;

// // funciones struc
// llamadafuncionstructcontrol returns [interfaces.Instruction llmstrufun]
// : op=ID_VALIDO PUNTO op1=ID_VALIDO PARIZQ listaparametrosllamada PARDER 
// {
//     $llmstrufun = instructions.NewStruckFuncionesControlP($op.line, $op.pos, $op.text, $op1.text, $listaparametrosllamada.lpll)
// }
// | op=ID_VALIDO PUNTO op1=ID_VALIDO PARIZQ PARDER 
// {
//     $llmstrufun = instructions.NewStruckFuncionesControl($op.line, $op.pos, $op.text, $op1.text)
// };

// llamadafuncionstructcontrolret returns [interfaces.Expression llmstrufunret]
// : op=ID_VALIDO PUNTO op1=ID_VALIDO PARIZQ listaparametrosllamada PARDER 
// {
//     $llmstrufunret = instructions.NewStruckFuncionesControlPR($op.line, $op.pos, $op.text, $op1.text, $listaparametrosllamada.lpll)
// }
// | op=ID_VALIDO PUNTO op1=ID_VALIDO PARIZQ PARDER 
// {
//     $llmstrufunret = instructions.NewStruckFuncionesControlR($op.line, $op.pos, $op.text, $op1.text)
// };


//CREACION DE FUNCIONES
funciondeclaracioncontrol returns [interfaces.Instruction fdc]
// : FUNCION ID_VALIDO PARIZQ listaparametro PARDER RETORNO tipodato LLAVEIZQ blockinterno LLAVEDER 
// {
//     $fdc = funciones.NewFuncionesDeclaracionRP($ID_VALIDO.line, $ID_VALIDO.pos, $ID_VALIDO.text, $listaparametro.listparfun, $tipodato.tipo, $blockinterno.blkint)
// }
: FUNCION ID_VALIDO PARIZQ  PARDER RETORNO tipodato LLAVEIZQ blockinterno LLAVEDER 
{
    $fdc = funciones.NewFuncionesDeclaracionR($ID_VALIDO.line, $ID_VALIDO.pos, $ID_VALIDO.text, $tipodato.tipo, $blockinterno.blkint)
}
// | FUNCION ID_VALIDO PARIZQ listaparametro PARDER LLAVEIZQ blockinterno LLAVEDER 
// {
//    $fdc = funciones.NewFuncionesDeclaracionP($ID_VALIDO.line, $ID_VALIDO.pos, $ID_VALIDO.text, $listaparametro.listparfun, $blockinterno.blkint)
// }
| FUNCION ID_VALIDO PARIZQ PARDER LLAVEIZQ blockinterno LLAVEDER 
{
    $fdc = funciones.NewFuncionesDeclaracion($ID_VALIDO.line, $ID_VALIDO.pos, $ID_VALIDO.text, $blockinterno.blkint)
}
;

listaparametro returns [interfaces.Instruction listparfun]
: op=(ID_VALIDO | GUIONBAJO)? op2=ID_VALIDO DOS_PUNTOS INOUT? tipodato COMA op3=listaparametro 
{
    if $op != nil{
        if $INOUT != nil{
            $listparfun = funciones.NewFuncionesListaParametro($op2.line, $op2.pos, $op.text, $op2.text, $tipodato.tipo, true, true, $op3.listparfun )
        }else {
            $listparfun = funciones.NewFuncionesListaParametro($op2.line, $op2.pos, $op.text, $op2.text, $tipodato.tipo, false, true, $op3.listparfun )
        } 
    }else{
        if $INOUT != nil{
            $listparfun = funciones.NewFuncionesListaParametro($op2.line, $op2.pos, "", $op2.text, $tipodato.tipo, true, false, $op3.listparfun )
        }else {
            $listparfun = funciones.NewFuncionesListaParametro($op2.line, $op2.pos, "", $op2.text, $tipodato.tipo, false, false,$op3.listparfun )
        } 
    }      
}
| op=(ID_VALIDO | GUIONBAJO)? op2=ID_VALIDO DOS_PUNTOS INOUT? tipodato 
{
    if $op != nil{
        if $INOUT != nil{
            $listparfun = funciones.NewFuncionesParametro($op2.line, $op2.pos, $op.text, $op2.text, $tipodato.tipo, true , true)
        }else {
            $listparfun = funciones.NewFuncionesParametro($op2.line, $op2.pos, $op.text, $op2.text, $tipodato.tipo, false, true)
        } 
    }else{
        if $INOUT != nil{
            $listparfun = funciones.NewFuncionesParametro($op2.line, $op2.pos, "", $op2.text, $tipodato.tipo, true, false)
        }else {
            $listparfun = funciones.NewFuncionesParametro($op2.line, $op2.pos, "", $op2.text, $tipodato.tipo, false, false)
    } 
    }
    
}
;

funcionllamadacontrol returns [interfaces.Instruction flctl]
// : ID_VALIDO PARIZQ listaparametrosllamada PARDER 
// {
//     $flctl = funciones.NewFuncionesControlP($ID_VALIDO.line, $ID_VALIDO.pos, $ID_VALIDO.text, $listaparametrosllamada.lpll)
// }
: ID_VALIDO PARIZQ PARDER 
{
    $flctl = funciones.NewFuncionesControl($ID_VALIDO.line, $ID_VALIDO.pos, $ID_VALIDO.text )
}
;

// funcionllamadacontrolConRetorno returns [interfaces.Expression flctlret]
// : ID_VALIDO PARIZQ listaparametrosllamada PARDER 
// {
//     $flctlret = instructions.NewFuncionesControlPR($ID_VALIDO.line, $ID_VALIDO.pos, $ID_VALIDO.text, $listaparametrosllamada.lpll)
// }
// | ID_VALIDO PARIZQ PARDER 
// {
//     $flctlret = instructions.NewFuncionesControlR($ID_VALIDO.line, $ID_VALIDO.pos, $ID_VALIDO.text )
// };

// listaparametrosllamada returns [interfaces.Instruction lpll]
// : DIRME ID_VALIDO COMA op2=listaparametrosllamada 
// {
//     $lpll = instructions.NewFuncionesLlamadaList1($DIRME.line, $DIRME.pos, $ID_VALIDO.text, $op2.lpll)    
// }
// | DIRME ID_VALIDO
// {
//     $lpll = instructions.NewFuncionesLlamadaList2($DIRME.line, $DIRME.pos, $ID_VALIDO.text)    
// }
// | (ID_VALIDO op=DOS_PUNTOS)? expr COMA op2=listaparametrosllamada 
// {
//     if $op != nil{
//         $lpll = instructions.NewFuncionesLlamadaList3($ID_VALIDO.line, $ID_VALIDO.pos, $ID_VALIDO.text, $expr.e, $op2.lpll)
//     }else{
//         $lpll = instructions.NewFuncionesLlamadaList4($COMA.line, $COMA.pos, $expr.e, $op2.lpll)
//     }
// }
// | (ID_VALIDO op=DOS_PUNTOS)? expr
// {
//     if $op != nil{
//         $lpll = instructions.NewFuncionesLlamadaList5($ID_VALIDO.line, $ID_VALIDO.pos, $ID_VALIDO.text, $expr.e)
//     }else{
//         $lpll = instructions.NewFuncionesLlamadaList6($expr.e)
//     }     
// };

// //CREACION DE EMBEBIDAS

// // FUNCION PRINT

// FUNCION PRINT
printstmt returns [interfaces.Instruction prnt]
: PRINT PARIZQ listaexpresions PARDER { $prnt = funciones.NewPrint($PRINT.line,$PRINT.pos,$listaexpresions.blkparf)};

intembebida returns [interfaces.Expression intemb]
: INT PARIZQ expr PARDER { $intemb = funciones.NewFuncionIntEmbebida($expr.e)};

floatembebida returns [interfaces.Expression floemb]
: FLOAT PARIZQ expr PARDER { $floemb = funciones.NewFuncionFloatEmbebida($expr.e)};

stringembebida returns [interfaces.Expression stremb]
: STRING PARIZQ expr PARDER { $stremb = funciones.NewFuncionStringEmbebida($expr.e)};