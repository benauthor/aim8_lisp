.PHONY: grammar

clean:
	rm -f *.tokens *.interp

grammar: clean
	antlr4 -Dlanguage=Python3 Aim8.g4
