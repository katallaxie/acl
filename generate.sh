#!/bin/sh

alias antlr4='antlr4'
antlr4 -Dlanguage=Go -no-visitor -package parsing -o pkg/parser *.g4