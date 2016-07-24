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
	"sort"
	"strings"
)

type keywordList []string

func (kwl keywordList) StartWith(prefix string) keywordList {
	rv := keywordList{}
	pos := sort.SearchStrings(kwl, prefix)
	if pos >= len(kwl) {
		return rv
	}
	for _, kw := range kwl[pos:] {
		if strings.HasPrefix(kw, prefix) {
			rv = append(rv, kw)
		} else {
			return rv
		}
	}
	return rv
}

func (kwl keywordList) HasExact(name string) bool {
	pos := sort.SearchStrings(kwl, name)
	if pos < len(kwl) && kwl[pos] == name {
		return true
	}
	return false
}

// snowConeKeyWords must be kept in sorted order
var snowConeKeyWords = keywordList{
	"!=", "$", "(", ")", "*", "*=", "+", "+=", "-", "-=", "->", "/", "/=", "<",
	"<+", "<-", "<=", "=", "==", "=>", ">", ">=", "?", "[", "]", "among", "and",
	"as", "atleast", "atlimit", "atmark", "attach", "backwardmode", "backwards",
	"booleans", "cursor", "decimal", "define", "delete", "do", "externals",
	"fail", "false", "for", "gopast", "goto", "groupings", "hex", "hop",
	"insert", "integers", "len", "lenof", "limit", "loop", "maxint", "minint",
	"next", "non", "not", "or", "repeat", "reverse", "routines", "set",
	"setlimit", "setmark", "size", "sizeof", "stringdef", "stringescapes",
	"strings", "substring", "test", "tolimit", "tomark", "true", "try", "unset"}

var keywordTokenTypes = map[string]int{
	"!=":            tNEQ,
	"$":             tDOLLAR,
	"(":             tLPAREN,
	")":             tRPAREN,
	"*":             tMULT,
	"*=":            tMULTASSIGN,
	"+":             tPLUS,
	"+=":            tPLUSASSIGN,
	"-":             tMINUS,
	"-=":            tMINUSASSIGN,
	"->":            tMOVER,
	"/":             tDIV,
	"/=":            tDIVASSIGN,
	"<":             tLT,
	"<+":            tINSERT,
	"<-":            tREPLACE,
	"<=":            tLTEQ,
	"=":             tASSIGN,
	"==":            tEQ,
	"=>":            tASSIGNR,
	">":             tGT,
	">=":            tGTEQ,
	"?":             tQUESTION,
	"[":             tLBRACKET,
	"]":             tRBRACKET,
	"among":         tAMONG,
	"and":           tAND,
	"as":            tAS,
	"atleast":       tATLEAST,
	"atlimit":       tATLIMIT,
	"atmark":        tATMARK,
	"attach":        tATTACH,
	"backwardmode":  tBACKWARDMODE,
	"backwards":     tBACKWARDS,
	"booleans":      tBOOLEANS,
	"cursor":        tCURSOR,
	"decimal":       tDECIMAL,
	"define":        tDEFINE,
	"delete":        tDELETE,
	"do":            tDO,
	"externals":     tEXTERNALS,
	"fail":          tFAIL,
	"false":         tFALSE,
	"for":           tFOR,
	"gopast":        tGOPAST,
	"goto":          tGOTO,
	"groupings":     tGROUPINGS,
	"hex":           tHEX,
	"hop":           tHOP,
	"insert":        tINSERT,
	"integers":      tINTEGERS,
	"len":           tLEN,
	"lenof":         tLENOF,
	"limit":         tLIMIT,
	"loop":          tLOOP,
	"maxint":        tMAXINT,
	"minint":        tMININT,
	"next":          tNEXT,
	"non":           tNON,
	"not":           tNOT,
	"or":            tOR,
	"repeat":        tREPEAT,
	"reverse":       tREVERSE,
	"routines":      tROUTINES,
	"set":           tSET,
	"setlimit":      tSETLIMIT,
	"setmark":       tSETMARK,
	"size":          tSIZE,
	"sizeof":        tSIZEOF,
	"stringdef":     tSTRINGDEF,
	"stringescapes": tSTRINGESCAPES,
	"strings":       tSTRINGS,
	"substring":     tSUBSTRING,
	"test":          tTEST,
	"tolimit":       tTOLIMIT,
	"tomark":        tTOMARK,
	"true":          tTRUE,
	"try":           tTRY,
	"unset":         tUNSET,
}
