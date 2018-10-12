// Code generated from Aim8.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Aim8

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAim8Listener is a complete listener for a parse tree produced by Aim8Parser.
type BaseAim8Listener struct{}

var _ Aim8Listener = &BaseAim8Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAim8Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAim8Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAim8Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAim8Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSexpr is called when production sexpr is entered.
func (s *BaseAim8Listener) EnterSexpr(ctx *SexprContext) {}

// ExitSexpr is called when production sexpr is exited.
func (s *BaseAim8Listener) ExitSexpr(ctx *SexprContext) {}

// EnterEnclosd_sexpr is called when production enclosd_sexpr is entered.
func (s *BaseAim8Listener) EnterEnclosd_sexpr(ctx *Enclosd_sexprContext) {}

// ExitEnclosd_sexpr is called when production enclosd_sexpr is exited.
func (s *BaseAim8Listener) ExitEnclosd_sexpr(ctx *Enclosd_sexprContext) {}

// EnterMexpr is called when production mexpr is entered.
func (s *BaseAim8Listener) EnterMexpr(ctx *MexprContext) {}

// ExitMexpr is called when production mexpr is exited.
func (s *BaseAim8Listener) ExitMexpr(ctx *MexprContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseAim8Listener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseAim8Listener) ExitExpr(ctx *ExprContext) {}

// EnterConditional_expr is called when production conditional_expr is entered.
func (s *BaseAim8Listener) EnterConditional_expr(ctx *Conditional_exprContext) {}

// ExitConditional_expr is called when production conditional_expr is exited.
func (s *BaseAim8Listener) ExitConditional_expr(ctx *Conditional_exprContext) {}

// EnterConditional_pair is called when production conditional_pair is entered.
func (s *BaseAim8Listener) EnterConditional_pair(ctx *Conditional_pairContext) {}

// ExitConditional_pair is called when production conditional_pair is exited.
func (s *BaseAim8Listener) ExitConditional_pair(ctx *Conditional_pairContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseAim8Listener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseAim8Listener) ExitPredicate(ctx *PredicateContext) {}

// EnterProposition is called when production proposition is entered.
func (s *BaseAim8Listener) EnterProposition(ctx *PropositionContext) {}

// ExitProposition is called when production proposition is exited.
func (s *BaseAim8Listener) ExitProposition(ctx *PropositionContext) {}

// EnterConnective is called when production connective is entered.
func (s *BaseAim8Listener) EnterConnective(ctx *ConnectiveContext) {}

// ExitConnective is called when production connective is exited.
func (s *BaseAim8Listener) ExitConnective(ctx *ConnectiveContext) {}

// EnterArrow is called when production arrow is entered.
func (s *BaseAim8Listener) EnterArrow(ctx *ArrowContext) {}

// ExitArrow is called when production arrow is exited.
func (s *BaseAim8Listener) ExitArrow(ctx *ArrowContext) {}

// EnterNull is called when production null is entered.
func (s *BaseAim8Listener) EnterNull(ctx *NullContext) {}

// ExitNull is called when production null is exited.
func (s *BaseAim8Listener) ExitNull(ctx *NullContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseAim8Listener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseAim8Listener) ExitAtom(ctx *AtomContext) {}
