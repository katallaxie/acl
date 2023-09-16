#!/bin/sh

alias antlr4='antlr4'

rm -rf pkg/parser
antlr4 -Dlanguage=Go -no-visitor -package parsing -o pkg/parser *.g4