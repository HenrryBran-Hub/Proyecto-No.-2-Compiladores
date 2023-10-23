grammar C3DGrammar;

options {
	tokenVocab = C3DLexer;
}

@header {
    import "Optimizar/interfacesc3d"    
    import "Optimizar/instructionsc3d"
    // import "Optimizar/environmentc3d"
    // import "Optimizar/expressionsc3d"
    // import "strings"
   
}

// Rules
z returns [[]interface{} code]
:   {    
        var mySlice []interface{}
        mySlice = make([]interface{}, 0) // Inicializa el slice vacío
    }
    encabezadoa (temporales PUNTOCOMA)? (blockfuncions)? funcionmain EOF 
    {
        mySlice = append(mySlice, $encabezadoa.encaa)
        if $PUNTOCOMA != nil {
            mySlice = append(mySlice, $temporales.tinst)
        }

        for _, item := range $blockfuncions.blkfunc {
            mySlice = append(mySlice, item)
        }

        mySlice = append(mySlice, $funcionmain.funmain)

        $code = mySlice
    }   
;

encabezadoa returns [interfacesc3d.Instruction encaa]
: INCLUDE STDIO DOUBLE HEAP CORCHIZQ h=NUMERO CORCHDER PUNTOCOMA DOUBLE STACK CORCHIZQ s=NUMERO CORCHDER PUNTOCOMA DOUBLE PSTACK PUNTOCOMA DOUBLE PHEAD  PUNTOCOMA
    {
        $encaa = instructionsc3d.NewAcumuladorEncabezado($h.text,$s.text)
    }
;

temporales returns [interfacesc3d.Instruction tinst]
: DOUBLE blocktemporales
    {
        $tinst = instructionsc3d.NewEjecucionTemporales($blocktemporales.blktmps)
    }
;

blocktemporales returns [[]interface{} blktmps]
@init{
    $blktmps = []interface{}{}
    var listTemp []IBloquetempsContext
}
: temps+=bloquetemps+
{
    listTemp = localctx.(*BlocktemporalesContext).GetTemps()
    for _, e := range listTemp {
        $blktmps = append($blktmps, e.GetTemps())
    }
}
;

bloquetemps returns [interfacesc3d.Instruction temps]
: COMMA ID_VALIDO { $temps = instructionsc3d.NewArregloTemporales($ID_VALIDO.text)}
| ID_VALIDO { $temps = instructionsc3d.NewArregloTemporales($ID_VALIDO.text)}
;

blockfuncions returns [[]interface{} blkfunc]
@init{
    $blkfunc = []interface{}{}
    var listFunc []IBloquefuncionesContext
}
: func+=bloquefunciones+
{
    listFunc = localctx.(*BlockfuncionsContext).GetFunc_()
    for _, e := range listFunc {
        $blkfunc = append($blkfunc, e.GetFunc_())
    }
}
;

bloquefunciones returns [interfacesc3d.Instruction func]
: VOID ID_VALIDO PARIZQ PARDER LLAVEIZQ block LLAVEDER
{
    $func = instructionsc3d.NewFuncionVoid($ID_VALIDO.text,$block.blk)
}
;

funcionmain returns[interfacesc3d.Instruction funmain]
: INT ID_VALIDO PARIZQ PARDER LLAVEIZQ block RETURN NUMERO PUNTOCOMA LLAVEDER 
{
    $funmain = instructionsc3d.NewFuncionMain($block.blk)
}
;

block returns [[]interface{} blk]
@init{
    $blk = []interface{}{}
    var listBlo []IInstructionContext
  }
: ins+=instruction+
    {
        listBlo = localctx.(*BlockContext).GetIns()
        for _, e := range listBlo {
            $blk = append($blk, e.GetInstr())
        }
    }
;


instruction returns [interfacesc3d.Instruction instr]
: head_op { $instr = $head_op.heapop}
| stack_op { $instr = $stack_op.stackop}
| printff { $instr = $printff.prtff } 
| operaritme { $instr = $operaritme.oparit }  
;

head_op returns [interfacesc3d.Instruction heapop]
: PHEAD IG NUMERO PUNTOCOMA
{
    $heapop = instructionsc3d.NewHeapop1($NUMERO.text)
}
| ID_VALIDO IG PHEAD PUNTOCOMA
{
    $heapop = instructionsc3d.NewHeapop2($ID_VALIDO.text)
}
| HEAP CORCHIZQ embebida PHEAD CORCHDER IG op=(NUMERO|ID_VALIDO) PUNTOCOMA
{
    $heapop = instructionsc3d.NewHeapop3($embebida.embe,$op.text)
}
| PHEAD IG PHEAD ADD NUMERO PUNTOCOMA
{
    $heapop = instructionsc3d.NewHeapop4($NUMERO.text)
}
;

stack_op returns [interfacesc3d.Instruction stackop]
: PSTACK IG NUMERO PUNTOCOMA
{
    $stackop = instructionsc3d.NewStackop1($NUMERO.text)
}
| ID_VALIDO IG PSTACK PUNTOCOMA
{
    $stackop = instructionsc3d.NewStackop2($ID_VALIDO.text)
}
| STACK CORCHIZQ (op1=NUMERO|(tip1=embebida op2=(PSTACK|ID_VALIDO|NUMERO))) CORCHDER IG (PARIZQ tipodata PARDER)? op3=(NUMERO|ID_VALIDO) PUNTOCOMA
{
    if $op1 != nil {
        if $PARIZQ != nil {
             $stackop = instructionsc3d.NewStackop31($op1.text, $tipodata.tipdat, $op3.text)
        }else{
            $stackop = instructionsc3d.NewStackop32($op1.text, $op3.text)
        }       
    }else{
         if $PARIZQ != nil {
             $stackop = instructionsc3d.NewStackop33($tip1.embe, $op2.text, $tipodata.tipdat, $op3.text)
        }else{
            $stackop = instructionsc3d.NewStackop34($tip1.embe, $op2.text, $op3.text)
        } 
    }    
}
| PSTACK IG PSTACK op=(ADD|SUB) NUMERO PUNTOCOMA
{
    $stackop = instructionsc3d.NewStackop4($op.text,$NUMERO.text)
}
;

printff returns [interfacesc3d.Instruction prtff]
: PRINTF PARIZQ CADENA COMMA NUMERO PARDER PUNTOCOMA
{
    $prtff = instructionsc3d.NewPrint1($CADENA.text, $NUMERO.text)
}
| PRINTF PARIZQ CADENA COMMA embebida op=(NUMERO|ID_VALIDO) PARDER PUNTOCOMA
{
    $prtff = instructionsc3d.NewPrint2($CADENA.text, $embebida.embe, $op.text)
}
;

embebida returns [string embe]
: PARIZQ INT PARDER
{
    str := "(int)"
    $embe = str
}
|  PARIZQ FLOAT PARDER
{
    str := "(float)"
    $embe = str
}
|  PARIZQ CHAR PARDER
{
    str := "(char)"
    $embe = str
}
;

tipodata returns [string tipdat]
: INT
{
    str := "int"
    $tipdat = str
}
|  FLOAT 
{
    str := "float"
    $tipdat = str
}
|  CHAR
{
    str := "char"
    $tipdat = str
}
;
 
operaritme returns [interfacesc3d.Instruction oparit]
: ID_VALIDO IG NUMERO PUNTOCOMA
{
    $heapop = instructionsc3d.NewOparit1($NUMERO.text)
}
;
