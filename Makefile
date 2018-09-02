.PHONY: clean grammar test

clean:
	rm -f *.tokens *.interp

grammar: clean
	antlr4 \
	-Dlanguage=Python3 \
	-o aim8/generated/ \
	-package aim8 \
	Aim8.g4

test:
	for t in test/*.py; do python $$t; done
