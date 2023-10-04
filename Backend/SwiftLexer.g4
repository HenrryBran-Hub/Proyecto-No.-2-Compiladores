lexer grammar SwiftLexer; 

// --------------- TOKENS TIPOS DE DATOS
INT: 'Int';
FLOAT: 'Float';
STRING: 'String';
BOOL: 'Bool';
CHARACT: 'Character';

// reserved words
TRU: 'true';
FAL: 'false';
VAR: 'var';
LET: 'let';
NULO: 'nil';
IF: 'if';
ELSE: 'else';
SWITCH: 'switch';
CASE: 'case';
DEFAULT: 'default';
BREAK: 'break';
CONTINUE: 'continue';
FOR: 'for';
IN: 'in';
RANGO: '...';
WHILE: 'while';
GUARD: 'guard';
RETURN: 'return';
FUNCION: 'func';
PRINT: 'print';
INOUT: 'inout';
APPEND: 'append';
REMOVE: 'remove';
REMOVELAST: 'removeLast';
COUNT: 'count';
ISEMPTY: 'isEmpty';
AT: 'at';
REPEATING: 'repeating';
STRUCT: 'struct';
MUTATING: 'mutating';
SELF: 'self';

// primitives
NUMBER: [0-9]+ ('.' [0-9]+)?;
CADENA: '"' ( ~'"')* '"';
ID_VALIDO: ([a-zA-Z_]) [a-zA-Z0-9_]*;
CHARACTER: '\'' ( ESCAPE | ~['\\\r\n]) '\'';
fragment ESCAPE: '\\' [\\'tnr];
WS: [ \t\r\n]+ -> skip;

// SIMBOLOS

IG: '=';
DOS_PUNTOS: ':';
PUNTOCOMA: ';';
CIERRE_INTE: '?';
PARIZQ: '(';
PARDER: ')';

DIF: '!=';
IG_IG: '==';
NOT: '!';
OR: '||';
AND: '&&';

MAY_IG: '>=';
MEN_IG: '<=';
MAYOR: '>';
MENOR: '<';

MODULO: '%';
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';

SUMA: '+=';
RESTA: '-=';

LLAVEIZQ: '{';
LLAVEDER: '}';
RETORNO: '->';
COMA: ',';
PUNTO: '.';
GUIONBAJO: '_';

CORCHIZQ: '[';
CORCHDER: ']';

DIRME: '&';

// skip
WHITESPACE: [ \\\r\n\t]+ -> skip;
COMMENT: '/*' .*? '*/' -> skip;
LINE_COMMENT: '//' ~[\r\n]* -> skip;

fragment ESC_SEQ:
	'\\' (
		'\\'
		| '@'
		| '['
		| ']'
		| '.'
		| '#'
		| '+'
		| '-'
		| '!'
		| ':'
		| ' '
	);

