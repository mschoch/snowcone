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

package snowcone

import (
	"reflect"
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	// DebugLexer = true
	// Logger = log.New(os.Stdout, "", log.LstdFlags)

	tests := []struct {
		in            string
		outTokenTypes []int
		outTokens     []yySymType
	}{
		{
			in:            "what",
			outTokenTypes: []int{tNAME},
			outTokens: []yySymType{
				{
					s: "what",
				},
			},
		},
		{
			in:            "32",
			outTokenTypes: []int{tNUMBER},
			outTokens: []yySymType{
				{
					n: 32,
				},
			},
		},
		{
			in:            "'alit'",
			outTokenTypes: []int{tLITERAL},
			outTokens: []yySymType{
				{
					s: "alit",
				},
			},
		},
		{
			in:            "strings",
			outTokenTypes: []int{tSTRINGS},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "integers",
			outTokenTypes: []int{tINTEGERS},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "booleans",
			outTokenTypes: []int{tBOOLEANS},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "routines",
			outTokenTypes: []int{tROUTINES},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "externals",
			outTokenTypes: []int{tEXTERNALS},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "groupings",
			outTokenTypes: []int{tGROUPINGS},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "(",
			outTokenTypes: []int{tLPAREN},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            ")",
			outTokenTypes: []int{tRPAREN},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "define",
			outTokenTypes: []int{tDEFINE},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "as",
			outTokenTypes: []int{tAS},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "+",
			outTokenTypes: []int{tPLUS},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "-",
			outTokenTypes: []int{tMINUS},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "*",
			outTokenTypes: []int{tMULT},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "/",
			outTokenTypes: []int{tDIV},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "maxint",
			outTokenTypes: []int{tMAXINT},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "minint",
			outTokenTypes: []int{tMININT},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "cursor",
			outTokenTypes: []int{tCURSOR},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "limit",
			outTokenTypes: []int{tLIMIT},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "size",
			outTokenTypes: []int{tSIZE},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "sizeof",
			outTokenTypes: []int{tSIZEOF},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "len",
			outTokenTypes: []int{tLEN},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "lenof",
			outTokenTypes: []int{tLENOF},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "$",
			outTokenTypes: []int{tDOLLAR},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "=",
			outTokenTypes: []int{tASSIGN},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "+=",
			outTokenTypes: []int{tPLUSASSIGN},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "-=",
			outTokenTypes: []int{tMINUSASSIGN},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "*=",
			outTokenTypes: []int{tMULTASSIGN},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "/=",
			outTokenTypes: []int{tDIVASSIGN},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "==",
			outTokenTypes: []int{tEQ},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "!=",
			outTokenTypes: []int{tNEQ},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            ">",
			outTokenTypes: []int{tGT},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "<",
			outTokenTypes: []int{tLT},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            ">=",
			outTokenTypes: []int{tGTEQ},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "<=",
			outTokenTypes: []int{tLTEQ},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "or",
			outTokenTypes: []int{tOR},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "and",
			outTokenTypes: []int{tAND},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "not",
			outTokenTypes: []int{tNOT},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "test",
			outTokenTypes: []int{tTEST},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "try",
			outTokenTypes: []int{tTRY},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "do",
			outTokenTypes: []int{tDO},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "fail",
			outTokenTypes: []int{tFAIL},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "goto",
			outTokenTypes: []int{tGOTO},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "gopast",
			outTokenTypes: []int{tGOPAST},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "repeat",
			outTokenTypes: []int{tREPEAT},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "loop",
			outTokenTypes: []int{tLOOP},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "atleast",
			outTokenTypes: []int{tATLEAST},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "insert",
			outTokenTypes: []int{tINSERT},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "<+",
			outTokenTypes: []int{tINSERT},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "attach",
			outTokenTypes: []int{tATTACH},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "<-",
			outTokenTypes: []int{tREPLACE},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "delete",
			outTokenTypes: []int{tDELETE},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "hop",
			outTokenTypes: []int{tHOP},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "next",
			outTokenTypes: []int{tNEXT},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "=>",
			outTokenTypes: []int{tASSIGNR},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "[",
			outTokenTypes: []int{tLBRACKET},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "]",
			outTokenTypes: []int{tRBRACKET},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "->",
			outTokenTypes: []int{tMOVER},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "setmark",
			outTokenTypes: []int{tSETMARK},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "tomark",
			outTokenTypes: []int{tTOMARK},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "atmark",
			outTokenTypes: []int{tATMARK},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "tolimit",
			outTokenTypes: []int{tTOLIMIT},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "atlimit",
			outTokenTypes: []int{tATLIMIT},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "setlimit",
			outTokenTypes: []int{tSETLIMIT},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "for",
			outTokenTypes: []int{tFOR},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "backwards",
			outTokenTypes: []int{tBACKWARDS},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "reverse",
			outTokenTypes: []int{tREVERSE},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "substring",
			outTokenTypes: []int{tSUBSTRING},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "among",
			outTokenTypes: []int{tAMONG},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "set",
			outTokenTypes: []int{tSET},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "unset",
			outTokenTypes: []int{tUNSET},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "non",
			outTokenTypes: []int{tNON},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "true",
			outTokenTypes: []int{tTRUE},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "false",
			outTokenTypes: []int{tFALSE},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "backwardmode",
			outTokenTypes: []int{tBACKWARDMODE},
			outTokens: []yySymType{
				{},
			},
		},
		{
			in:            "what the 47 'bang'",
			outTokenTypes: []int{tNAME, tNAME, tNUMBER, tLITERAL},
			outTokens: []yySymType{
				{
					s: "what",
				},
				{
					s: "the",
				},
				{
					n: 47,
				},
				{
					s: "bang",
				},
			},
		},
		{
			in:            "strings (bob wood)",
			outTokenTypes: []int{tSTRINGS, tLPAREN, tNAME, tNAME, tRPAREN},
			outTokens: []yySymType{
				{},
				{},
				{
					s: "bob",
				},
				{
					s: "wood",
				},
				{},
			},
		},
		{
			in:            "define stem as",
			outTokenTypes: []int{tDEFINE, tNAME, tAS},
			outTokens: []yySymType{
				{},
				{
					s: "stem",
				},
				{},
			},
		},
		{
			in:            "stringescapes {}",
			outTokenTypes: []int{tSTRINGESCAPES},
			outTokens: []yySymType{
				{
					s: "{}",
				},
			},
		},
		{
			in: `
			/* a multi
			line comment */
			strings (bob wood)`,
			outTokenTypes: []int{tSTRINGS, tLPAREN, tNAME, tNAME, tRPAREN},
			outTokens: []yySymType{
				{},
				{},
				{
					s: "bob",
				},
				{
					s: "wood",
				},
				{},
			},
		},
	}

	for _, test := range tests {
		var tokenTypes []int
		var tokens []yySymType
		r := strings.NewReader(test.in)
		l := newSnowConeLex(r)
		var lval yySymType
		rv := l.Lex(&lval)
		for rv > 0 {
			tokenTypes = append(tokenTypes, rv)
			tokens = append(tokens, lval)
			lval.s = ""
			lval.n = 0
			rv = l.Lex(&lval)
		}
		if !reflect.DeepEqual(test.outTokenTypes, tokenTypes) {
			t.Errorf("expected token types %v, got %v, for input '%s'", test.outTokenTypes, tokenTypes, test.in)
		}
		if !reflect.DeepEqual(test.outTokens, tokens) {
			t.Errorf("expected tokens %v, got %v, for input '%s'", test.outTokens, tokens, test.in)
		}
	}
}
