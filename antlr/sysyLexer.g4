lexer grammar sysyLexer;

Const: 'const';

fragment
Int: 'int';

fragment
Float: 'float';

Void: 'void';

Type
    : Int
    | Float
    ;

If: 'if';
Else: 'else';
While: 'while';
Break: 'break';
Continue: 'continue';
Return: 'return';

LeftParen: '(';
RightParen: ')';
LeftBracket: '[';
RightBracket: ']';
LeftBrace: '{';
RightBrace: '}';
Comma: ',';
Semicolon: ';';
Assign: '=';
Add: '+';
Sub: '-';
Mul: '*';
Div: '/';
Mod: '%';
Less: '<';
Greater: '>';
Equal: '==';
NotEqual: '!=';
LessEqual: '<=';
GreaterEqual: '>=';
And: '&&';
Or: '||';
Not: '!';

Identifier
    : NonDigit
     (  NonDigit
     | Digit
     )*
    ;

fragment
NonDigit
    : [a-zA-Z_]
    ;

fragment
Digit
    : [0-9]
    ;

IntegerConstant
    : OctalConstant
    | HexadecimalConstant
    | DecimalConstant
    ;

fragment
DecimalConstant
    : NonZeroDigit
      Digit*
    | '0'
    ;

fragment
NonZeroDigit
    : [1-9]
    ;

fragment
HexadecimalConstant
    : '0' [xX] HexadecimalDigit+
    ;

fragment
HexadecimalDigit
    : Digit
    | [a-fA-F]
    ;

fragment
OctalConstant
    : '0' OctalDigit*
    ;

fragment
OctalDigit
    : [0-7]
    ;

FloatConstant
    :  DecimalConstant '.' Digit*
    | '.' Digit+
    ;

WS
    : [ \t\r\n]+ -> skip
    ;

LineComment
    : '//' ~[\r\n]* -> skip
    ;

BlockComment
    : '/*' .*? '*/' -> skip
    ;

Preprocessor
    : '#' ~[\r\n]* -> skip
    ;