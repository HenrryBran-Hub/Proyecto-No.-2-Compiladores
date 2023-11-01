lexer grammar C3DLexer; 

// --------------- TOKENS TIPOS DE DATOS
INT: 		'int';
FLOAT:		'float';
DOUBLE: 	'double';
CHAR:		'char';
VOID: 		'void';
INCLUDE: 	'#include';
STDIO: 		'<stdio.h>';
HEAP: 		'heap';
STACK: 		'stack';
IF: 		'if';
GOTO: 		'goto';
RETURN: 	'return';
PRINTF: 	'printf';
PHEAD: 		'H';
PSTACK: 	'P';

// primitives
NUMERO: 			('-')?[0-9]+ ('.' [0-9]+)?;
CADENA: 			'"' ( ~'"')* '"';
ID_VALIDO: 			([a-zA-Z_]) [a-zA-Z0-9_]*;
CHARACTER: 			'\'' ( ESCAPE | ~['\\\r\n]) '\'';
fragment ESCAPE: 	'\\' [\\'tnr];
WS: [ \t\r\n]+ -> skip;

// SIMBOLOS

COMMA: 			',';
DOS_PUNTOS: 	':';
PUNTOCOMA: 		';';

IG: 			'=';
DIF: 			'!=';
IG_IG: 			'==';
NOT: 			'!';

MAY_IG: '		>=';
MEN_IG: 		'<=';
MAYOR: 			'>';
MENOR: 			'<';

MODULO: 		'%';
MUL: 			'*';
DIV: 			'/';
ADD: 			'+';
SUB: 			'-';


PARIZQ: 		'(';
PARDER: 		')';

LLAVEIZQ: 		'{';
LLAVEDER: 		'}';

CORCHIZQ: 		'[';
CORCHDER: 		']';


// skip
WHITESPACE: 	[ \\\r\n\t]+ -> skip;
COMMENT:		'/*' .*? '*/' -> skip;
LINE_COMMENT: 	'//' ~[\r\n]* -> skip;

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

