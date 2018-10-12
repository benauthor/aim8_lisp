.PHONY: clean grammar test

clean:
	rm -f *.tokens *.interp

grammar: clean
	antlr4 \
	-o internal/parser/ \
	-Dlanguage=Go \
	Aim8.g4
