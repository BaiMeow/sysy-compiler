parser grammar sysyParser;

options {
    tokenVocab=sysyLexer;
}

comp:
    compUnit* EOF;

compUnit
    : decl
    | funcDef
    ;

decl
    : constDecl
    | varDecl
    ;

constDecl
    : Const Type constDef (Comma constDef)* Semicolon
    ;

constDef:
    Identifier (LeftBracket constExp RightBracket)* Assign constinitVal
    ;

constinitVal
    : constExp
    | LeftBrace (constExp (Comma constExp)*)? RightBrace
    ;

varDecl
    : Type varDef (Comma varDef)* Semicolon
    ;

varDef
    : Identifier (LeftBracket constExp RightBracket)* (Assign initVal)?
    ;

initVal
    : exp
    | LeftBrace (exp (Comma exp)*)? RightBrace
    ;

funcDef
    : funcType Identifier LeftParen funcFParams? RightParen block
    ;

funcType
    : Type
    | Void
    ;

funcFParams
    : funcFParam (Comma funcFParam)*
    ;

funcFParam
    : Type Identifier (LeftBracket RightBracket (LeftBracket constExp RightBracket)*)?
    ;

block
    : LeftBrace blockItem* RightBrace
    ;

blockItem
    : decl
    | stmt
    ;

stmt
    : lVal Assign exp Semicolon
    | exp? Semicolon
    | block
    | If LeftParen cond RightParen stmt (Else stmt)?
    | While LeftParen cond RightParen stmt
    | Break Semicolon
    | Continue Semicolon
    | Return exp? Semicolon
    ;

exp
    : addExp
    ;

cond
    : lOrExp
    ;

lVal
    : Identifier (LeftBracket exp RightBracket)*
    ;
    
primaryExp
    : LeftParen exp RightParen
    | lVal
    | number
    ;

number
    : IntegerConstant
    | FloatConstant
    ;

unaryExp
    : primaryExp
    | Identifier LeftParen funcRParams RightParen
    | unaryOp unaryExp
    ;

unaryOp
    : Add
    | Sub
    | Not
    ;

funcRParams
    : exp (Comma exp)*
    ;

mulExp
    : unaryExp
    ((Mul
    | Div
    | Mod
    ) mulExp )?
    ;

addExp
    : mulExp
    ((Add
    | Sub
    ) addExp )?
    ;

relExp
    : addExp
    ((Less
    | Greater
    | LessEqual
    | GreaterEqual
    ) relExp)?
    ;

eqExp
    : relExp
    ((Equal
    | NotEqual
    ) eqExp )?
    ;

lAndExp
    : eqExp (And lAndExp)?
    ;

lOrExp
    : lAndExp (Or lOrExp)?
    ;

constExp
    : addExp
    ;

