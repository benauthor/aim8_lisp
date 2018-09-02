from unittest import TestCase

from aim8.interpreter import (
    FALSE,
    NIL,
    TRUE,
    aim8EvalString,
)


class Aim8TestCase(TestCase):
    def test_token(self):
        self.assertEqual("AB", aim8EvalString("AB"))

    def test_sexpr_literal(self):
        self.assertEqual(["AB", "A"], aim8EvalString("(AB,A)"))

    def test_null_literal(self):
        self.assertEqual(NIL, aim8EvalString("NIL"))
        self.assertEqual(NIL, aim8EvalString("Λ"))

    def test_null_func(self):
        self.assertEqual(TRUE, aim8EvalString("null[NIL]"))
        self.assertEqual(FALSE, aim8EvalString("null[ABAB]"))


if __name__ == '__main__':
    import unittest
    unittest.main()
