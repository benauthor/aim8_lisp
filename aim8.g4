grammar aim8;

/*
The S-expressions are formed according to the following
recursive rules.
          1.   The atomic symbols p₁ p₂ etc are S-expressions.
          2.   A null expression ⋀ is also admitted.
          3.   If e is an S-expression so is  (e).
          4.   If e₁ and (e₂) are S-expressions so is (e₁,e₂).
*/

LPAREN
    : "("
    ;

RPAREN
    : ")"
    ;

NULL
    : "⋀"
    ;

sexpr
    : ATOM
    | NULL
    | LPAREN sexpr RPAREN
    | LPAREN sexpr COMMA sexpr RPAREN
    ;

/*
In what follows we shall use sequences of capital Latin
letters as atomic symbols.
*/

ATOM
    : [A-Z]+
    ;
