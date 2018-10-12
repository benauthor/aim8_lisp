// Code generated from Aim8.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Aim8

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Aim8Listener is a complete listener for a parse tree produced by Aim8Parser.
type Aim8Listener interface {
	antlr.ParseTreeListener

	// EnterSexpr is called when entering the sexpr production.
	EnterSexpr(c *SexprContext)

	// EnterEnclosd_sexpr is called when entering the enclosd_sexpr production.
	EnterEnclosd_sexpr(c *Enclosd_sexprContext)

	// EnterMexpr is called when entering the mexpr production.
	EnterMexpr(c *MexprContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterConditional_expr is called when entering the conditional_expr production.
	EnterConditional_expr(c *Conditional_exprContext)

	// EnterConditional_pair is called when entering the conditional_pair production.
	EnterConditional_pair(c *Conditional_pairContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterProposition is called when entering the proposition production.
	EnterProposition(c *PropositionContext)

	// EnterConnective is called when entering the connective production.
	EnterConnective(c *ConnectiveContext)

	// EnterArrow is called when entering the arrow production.
	EnterArrow(c *ArrowContext)

	// EnterNull is called when entering the null production.
	EnterNull(c *NullContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// ExitSexpr is called when exiting the sexpr production.
	ExitSexpr(c *SexprContext)

	// ExitEnclosd_sexpr is called when exiting the enclosd_sexpr production.
	ExitEnclosd_sexpr(c *Enclosd_sexprContext)

	// ExitMexpr is called when exiting the mexpr production.
	ExitMexpr(c *MexprContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitConditional_expr is called when exiting the conditional_expr production.
	ExitConditional_expr(c *Conditional_exprContext)

	// ExitConditional_pair is called when exiting the conditional_pair production.
	ExitConditional_pair(c *Conditional_pairContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitProposition is called when exiting the proposition production.
	ExitProposition(c *PropositionContext)

	// ExitConnective is called when exiting the connective production.
	ExitConnective(c *ConnectiveContext)

	// ExitArrow is called when exiting the arrow production.
	ExitArrow(c *ArrowContext)

	// ExitNull is called when exiting the null production.
	ExitNull(c *NullContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)
}
