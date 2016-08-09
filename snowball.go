//  Copyright (c) 2016 Marty Schoch
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the
//  License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing,
//  software distributed under the License is distributed on an "AS
//  IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
//  express or implied. See the License for the specific language
//  governing permissions and limitations under the License.

//go:generate go tool yacc -o snowball.y.go snowball.y
//go:generate sed -i "" -e 1d snowball.y.go

// Package snowcone provides the ability to parse snowball source files.
// See http://snowballstem.org/
// In the future, this package will be able to generate Go source capable of
// executing the programs described by the snowball source.
package snowcone

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

// DebugLexer enables debug output from the lexer
var DebugLexer bool

// DebugParser enables debug output from the parser
var DebugParser bool

// Logger lets you provide a custom log endpoint in your application
var Logger = log.New(ioutil.Discard, "bleve", log.LstdFlags)

// Parse takes the snowball source pointed at by the provided reader
// and attemptes to parse it
// Currently this only returns an error, in the future it will also
// return an AST built during parsing.
func Parse(program io.Reader) (*prog, error) {
	lex := newLexerWrapper(newSnowConeLex(program))
	p := doParse(lex)

	if len(lex.errs) > 0 {
		return nil, fmt.Errorf(strings.Join(lex.errs, "\n"))
	}

	return p, nil
}

func doParse(lex *lexerWrapper) *prog {
	defer func() {
		r := recover()
		if r != nil {
			lex.Error("Errors while parsing.")
		}
	}()

	yyParse(lex)
	return lex.p
}

type lexerWrapper struct {
	lex  yyLexer
	errs []string
	p    *prog
	sd   *stringdefs
}

func newLexerWrapper(lex yyLexer) *lexerWrapper {
	return &lexerWrapper{
		lex:  lex,
		errs: []string{},
		p:    &prog{},
		sd:   &stringdefs{},
	}
}

func (l *lexerWrapper) Lex(lval *yySymType) int {
	return l.lex.Lex(lval)
}

func (l *lexerWrapper) Error(s string) {
	l.errs = append(l.errs, s)
}
