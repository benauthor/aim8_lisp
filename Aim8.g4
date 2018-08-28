grammar Aim8;

/*
"The S-expressions are formed according to the following
recursive rules.

1.   The atomic symbols p₁ p₂ etc are S-expressions.
2.   A null expression ⋀ is also admitted.
3.   If e is an S-expression so is  (e).
4.   If e₁ and (e₂) are S-expressions so is (e₁,e₂)."

N.B. by 1960's LISP I, we see the replacement of commas in raw S-expressions
with dots, and introduction of list notation such that

  (A B C) -> (A . (B . (C . nil)))

However we will not implement these later refinements here.
*/

sexpr
    : null
    | atom
    | enclosd_sexpr
    ;

enclosd_sexpr
    : LPAREN sexpr (COMMA sexpr)* RPAREN
    ;

/*
AIM 8 included, without naming as such, the idea of M-expressions
(meta expressions). These were primitive operations on S-expressions.
It was quickly seen that M-expressions could be expressed *as*
S-expressions and initial implementations used that syntax. According
to McCarthy:

"The project of defining M-expressions precisely and compiling them or
at least translating them into S-expressions was neither finalized nor
explicitly abandoned. It just receded into the indefinite future, and
a new generation of programmers appeared who preferred internal
notation to any FORTRAN-like or ALGOL-like notation that could be
devised."

Well, let's do it now.
*/

mexpr
    : NAME LBRACKET expr (SEMICOLON expr)* RBRACKET
    ;

expr
    : sexpr
    | mexpr
    | conditional_expr
    ;

conditional_expr
    : LBRACKET conditional_pair (SEMICOLON conditional_pair)* RBRACKET
    ;

conditional_pair
    : predicate arrow expr
    ;

predicate
    : mexpr
    | proposition
    ;

proposition
    : mexpr connective mexpr
    ;

connective
    : AND
    | OR
    | NOT
    ;

arrow
    : ARROW
    | ALT_ARROW
    ;

null
    : NULL
    | ALT_NULL
    ;

atom
    : ATOM
    ;

/* lexer rules */
LPAREN
    : '('
    ;

RPAREN
    : ')'
    ;

LBRACKET
    : '['
    ;

RBRACKET
    : ']'
    ;

COMMA
    : ','
    ;

SEMICOLON
    : ';'
    ;

// capital lambda U+039B
NULL
    : 'Λ'
    ;

// lowercase lambda U+03BB
LAMBDA
    : 'λ'
    ;

// long rightwards arrow U+27F6
ARROW
    : '⟶'
    ;

// logical and U+2227
AND
    : '∧'
    ;

// logical or U+2228
OR
    : '∨'
    ;

// those are a pain to type...
ALT_ARROW
    : '->'
    ;

ALT_NULL
    : 'NIL'
    ;

NOT
    : '~'
    ;

TRUE
    : '1'
    ;

FALSE
    : '0'
    ;

/*
"In what follows we shall use sequences of capital Latin letters
as atomic symbols."
*/
ATOM
    : [A-Z]+
    ;

NAME
    : [a-z]+
    ;
