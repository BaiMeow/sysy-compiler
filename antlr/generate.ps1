Remove-Item -Recurse -Force ../parser
python -m pipenv run antlr4 -Dlanguage=Go -o ../parser -package parser -visitor -no-listener sysyLexer.g4 sysyParser.g4