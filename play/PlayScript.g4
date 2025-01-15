grammar PlayScript;

//options { tokenVocab=CommonLexer; }
import CommonLexer;

expression
    : primary
    | expression bop=('*'|'/'|'%') expression
    | expression bop=('+'|'-') expression
    ;

primary
    : '(' expression ')'
    | integerLiteral
    ;

integerLiteral
    : DECIMAL_LITERAL
    | HEX_LITERAL
    | OCT_LITERAL
    | BINARY_LITERAL
    ;