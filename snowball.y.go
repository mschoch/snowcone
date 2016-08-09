package snowcone

import __yyfmt__ "fmt"

//line snowball.y:2
import (
	"fmt"
	"unicode/utf8"
)

func logDebugGrammar(format string, v ...interface{}) {
	if DebugParser {
		Logger.Printf(fmt.Sprintf("PARSER %s", format), v...)
	}
}

//line snowball.y:16
type yySymType struct {
	yys          int
	s            string
	n            int
	strings      []string
	declarations decls
	gitems       groupitems
	g            *grouping
	no           node
	b            *bliteral
	r            *routine
	a            *amongitem
	ai           amongitems
	aexpr        ae
	ic           *iCommand
	sc           *sCommand
	p            *prog
	cl           commands
}

const tLITERAL = 57346
const tNUMBER = 57347
const tNAME = 57348
const tSTRINGS = 57349
const tINTEGERS = 57350
const tBOOLEANS = 57351
const tROUTINES = 57352
const tEXTERNALS = 57353
const tGROUPINGS = 57354
const tLPAREN = 57355
const tRPAREN = 57356
const tDEFINE = 57357
const tAS = 57358
const tPLUS = 57359
const tMINUS = 57360
const tMULT = 57361
const tDIV = 57362
const tMAXINT = 57363
const tMININT = 57364
const tCURSOR = 57365
const tLIMIT = 57366
const tSIZE = 57367
const tSIZEOF = 57368
const tLEN = 57369
const tLENOF = 57370
const tDOLLAR = 57371
const tASSIGN = 57372
const tPLUSASSIGN = 57373
const tMINUSASSIGN = 57374
const tMULTASSIGN = 57375
const tDIVASSIGN = 57376
const tEQ = 57377
const tNEQ = 57378
const tGT = 57379
const tLT = 57380
const tGTEQ = 57381
const tLTEQ = 57382
const tOR = 57383
const tAND = 57384
const tNOT = 57385
const tTEST = 57386
const tTRY = 57387
const tDO = 57388
const tFAIL = 57389
const tGOTO = 57390
const tGOPAST = 57391
const tREPEAT = 57392
const tLOOP = 57393
const tATLEAST = 57394
const tINSERT = 57395
const tATTACH = 57396
const tREPLACE = 57397
const tDELETE = 57398
const tHOP = 57399
const tNEXT = 57400
const tASSIGNR = 57401
const tLBRACKET = 57402
const tRBRACKET = 57403
const tMOVER = 57404
const tSETMARK = 57405
const tTOMARK = 57406
const tATMARK = 57407
const tTOLIMIT = 57408
const tATLIMIT = 57409
const tSETLIMIT = 57410
const tFOR = 57411
const tBACKWARDS = 57412
const tREVERSE = 57413
const tSUBSTRING = 57414
const tAMONG = 57415
const tSET = 57416
const tUNSET = 57417
const tNON = 57418
const tTRUE = 57419
const tFALSE = 57420
const tBACKWARDMODE = 57421
const tQUESTION = 57422
const tSTRINGESCAPES = 57423
const tSTRINGDEF = 57424
const tHEX = 57425
const tDECIMAL = 57426
const tUMINUS = 57427

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"tLITERAL",
	"tNUMBER",
	"tNAME",
	"tSTRINGS",
	"tINTEGERS",
	"tBOOLEANS",
	"tROUTINES",
	"tEXTERNALS",
	"tGROUPINGS",
	"tLPAREN",
	"tRPAREN",
	"tDEFINE",
	"tAS",
	"tPLUS",
	"tMINUS",
	"tMULT",
	"tDIV",
	"tMAXINT",
	"tMININT",
	"tCURSOR",
	"tLIMIT",
	"tSIZE",
	"tSIZEOF",
	"tLEN",
	"tLENOF",
	"tDOLLAR",
	"tASSIGN",
	"tPLUSASSIGN",
	"tMINUSASSIGN",
	"tMULTASSIGN",
	"tDIVASSIGN",
	"tEQ",
	"tNEQ",
	"tGT",
	"tLT",
	"tGTEQ",
	"tLTEQ",
	"tOR",
	"tAND",
	"tNOT",
	"tTEST",
	"tTRY",
	"tDO",
	"tFAIL",
	"tGOTO",
	"tGOPAST",
	"tREPEAT",
	"tLOOP",
	"tATLEAST",
	"tINSERT",
	"tATTACH",
	"tREPLACE",
	"tDELETE",
	"tHOP",
	"tNEXT",
	"tASSIGNR",
	"tLBRACKET",
	"tRBRACKET",
	"tMOVER",
	"tSETMARK",
	"tTOMARK",
	"tATMARK",
	"tTOLIMIT",
	"tATLIMIT",
	"tSETLIMIT",
	"tFOR",
	"tBACKWARDS",
	"tREVERSE",
	"tSUBSTRING",
	"tAMONG",
	"tSET",
	"tUNSET",
	"tNON",
	"tTRUE",
	"tFALSE",
	"tBACKWARDMODE",
	"tQUESTION",
	"tSTRINGESCAPES",
	"tSTRINGDEF",
	"tHEX",
	"tDECIMAL",
	"tUMINUS",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 109
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 629

var yyAct = [...]int{

	109, 99, 145, 39, 31, 52, 20, 21, 95, 96,
	95, 96, 184, 154, 155, 152, 153, 154, 155, 152,
	153, 154, 155, 100, 93, 94, 137, 197, 33, 34,
	35, 36, 37, 175, 40, 2, 161, 45, 138, 17,
	10, 11, 12, 13, 14, 15, 147, 42, 16, 41,
	149, 147, 50, 49, 29, 148, 144, 30, 48, 38,
	148, 97, 51, 47, 46, 44, 43, 27, 26, 122,
	25, 24, 23, 22, 126, 123, 124, 125, 18, 177,
	162, 130, 131, 101, 102, 103, 104, 105, 106, 107,
	108, 42, 158, 41, 157, 139, 136, 140, 141, 40,
	40, 135, 150, 129, 128, 127, 32, 28, 132, 133,
	134, 156, 7, 19, 8, 9, 1, 3, 54, 142,
	143, 53, 159, 146, 5, 6, 4, 0, 0, 0,
	0, 0, 0, 151, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 160, 0, 0, 176,
	179, 0, 0, 180, 181, 182, 183, 0, 0, 0,
	0, 0, 0, 163, 0, 186, 187, 188, 189, 190,
	191, 192, 193, 194, 195, 196, 0, 42, 0, 41,
	0, 0, 0, 0, 0, 185, 58, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 92, 164, 165, 166, 167, 168, 169, 170,
	171, 172, 173, 174, 0, 0, 59, 60, 61, 62,
	63, 64, 65, 66, 67, 68, 69, 70, 71, 72,
	73, 74, 75, 76, 77, 78, 79, 80, 81, 82,
	83, 84, 0, 85, 86, 87, 55, 88, 89, 90,
	56, 57, 42, 91, 41, 0, 0, 0, 0, 0,
	0, 58, 120, 119, 0, 152, 153, 154, 155, 0,
	121, 0, 0, 0, 0, 110, 0, 92, 111, 112,
	113, 114, 115, 116, 117, 118, 0, 0, 0, 0,
	0, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 69, 70, 71, 72, 73, 74, 75, 76, 77,
	78, 79, 80, 81, 82, 83, 84, 0, 85, 86,
	87, 55, 88, 89, 90, 56, 57, 42, 91, 41,
	0, 0, 0, 0, 0, 0, 58, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 92, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 95, 96, 59, 60, 61, 62,
	63, 64, 65, 66, 67, 68, 69, 70, 71, 72,
	73, 74, 75, 76, 77, 78, 79, 80, 81, 82,
	83, 84, 0, 85, 86, 87, 55, 88, 89, 90,
	56, 57, 42, 91, 41, 0, 0, 0, 0, 0,
	0, 58, 178, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 92, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 69, 70, 71, 72, 73, 74, 75, 76, 77,
	78, 79, 80, 81, 82, 83, 84, 0, 85, 86,
	87, 55, 88, 89, 90, 56, 57, 42, 91, 41,
	0, 0, 0, 0, 0, 0, 58, 98, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 92, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 59, 60, 61, 62,
	63, 64, 65, 66, 67, 68, 69, 70, 71, 72,
	73, 74, 75, 76, 77, 78, 79, 80, 81, 82,
	83, 84, 0, 85, 86, 87, 55, 88, 89, 90,
	56, 57, 42, 91, 41, 0, 0, 0, 0, 0,
	0, 58, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 92, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 69, 70, 71, 72, 73, 74, 75, 76, 77,
	78, 79, 80, 81, 82, 83, 84, 0, 85, 86,
	87, 55, 88, 89, 90, 56, 57, 0, 91,
}
var yyPact = [...]int{

	33, -1000, -1000, 33, -1000, -1000, -1000, 65, -1000, -77,
	60, 59, 58, 57, 55, 54, 101, -1000, 33, 53,
	-1000, -1000, 100, 100, 100, 100, 100, 100, 43, 52,
	-1000, 51, 100, 50, 49, 44, 39, 38, 548, -1000,
	7, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -31, -1000, -1000, -1000, 48, -1000, -1000, 473, 548,
	548, 548, 548, 548, 548, 548, 548, 257, 257, 87,
	87, 87, -1000, 257, -1000, 99, -1000, -1000, 98, 97,
	257, 257, -1000, -1000, 548, 548, 548, -1000, 95, 90,
	20, -1000, 89, 87, 87, 548, 548, 42, -1000, 36,
	323, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 248,
	257, -1000, -1000, -1000, -1000, -1000, 88, -1000, 86, -1000,
	-1000, 257, 248, -1000, -1000, -1000, 2, -1000, -1000, -1000,
	2, 2, -33, -1000, -1000, -1000, -1000, -1000, 74, 173,
	-1000, -1000, -1000, -1000, -1000, 19, 47, 73, 398, -1000,
	-1000, -1000, 257, 257, 257, 257, -1000, -1000, -1000, -2,
	-1000, 548, -1000, -1000, 257, 257, 257, 257, 257, 257,
	257, 257, 257, 257, 257, -1000, -1000, -1000, -1000, 13,
	-6, -6, -1000, -1000, -1000, -1000, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, -1000,
}
var yyPgo = [...]int{

	0, 4, 126, 3, 5, 125, 23, 124, 123, 2,
	1, 0, 121, 118, 117, 35, 116, 113,
}
var yyR1 = [...]int{

	0, 16, 15, 15, 14, 14, 14, 14, 14, 14,
	17, 17, 17, 2, 2, 2, 2, 2, 2, 7,
	4, 4, 3, 3, 3, 5, 10, 10, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 9, 9, 8, 8, 8, 8, 13,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 1, 1,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 1, 1, 1, 4, 1, 3,
	0, 1, 1, 4, 4, 4, 4, 4, 4, 4,
	1, 1, 1, 3, 3, 3, 1, 2, 1, 1,
	1, 3, 4, 1, 1, 2, 3, 2, 2, 2,
	2, 2, 2, 2, 2, 3, 3, 2, 2, 2,
	1, 2, 1, 2, 1, 1, 2, 2, 2, 2,
	1, 1, 4, 2, 2, 1, 2, 2, 2, 3,
	1, 3, 3, 1, 2, 1, 2, 2, 3, 3,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 3, 3, 3, 3, 2, 1, 1, 1, 1,
	1, 2, 1, 2, 1, 1, 3, 1, 2,
}
var yyChk = [...]int{

	-1000, -16, -15, -14, -2, -7, -5, 79, 81, 82,
	7, 8, 9, 10, 11, 12, 15, -15, 13, -17,
	83, 84, 13, 13, 13, 13, 13, 13, 6, -15,
	4, -1, 6, -1, -1, -1, -1, -1, 16, -3,
	-4, 6, 4, 14, 14, -1, 14, 14, 14, 14,
	14, -6, -4, -12, -13, 73, 77, 78, 13, 43,
	44, 45, 46, 47, 48, 49, 50, 51, 52, 53,
	54, 55, 56, 57, 58, 59, 60, 61, 62, 63,
	64, 65, 66, 67, 68, 70, 71, 72, 74, 75,
	76, 80, 29, 17, 18, 41, 42, 13, 14, -10,
	-6, -6, -6, -6, -6, -6, -6, -6, -6, -11,
	18, 21, 22, 23, 24, 25, 26, 27, 28, 6,
	5, 13, -11, -4, -4, -4, -11, 6, 6, 6,
	-11, -11, -6, -6, -6, 6, 6, 6, 18, 6,
	-3, -3, -6, -6, 14, -9, -8, 4, 13, 14,
	-10, -6, 17, 18, 19, 20, -11, 6, 6, -11,
	-6, 69, 6, -6, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 14, -9, 6, 14, -10,
	-11, -11, -11, -11, 14, -6, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, 14,
}
var yyDef = [...]int{

	0, -2, 1, 2, 4, 5, 6, 0, 8, 10,
	0, 0, 0, 0, 0, 0, 0, 3, 0, 0,
	11, 12, 0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 107, 0, 0, 0, 0, 0, 0, 25,
	22, 20, 21, 7, 13, 108, 14, 15, 16, 17,
	18, 19, 28, 29, 30, 0, 33, 34, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 50, 0, 52, 0, 54, 55, 0, 0,
	0, 0, 60, 61, 0, 0, 0, 65, 0, 0,
	0, 70, 0, 0, 0, 0, 0, 0, 35, 0,
	26, 37, 38, 39, 40, 41, 42, 43, 44, 0,
	0, 96, 97, 98, 99, 100, 0, 102, 0, 104,
	105, 0, 0, 47, 48, 49, 51, 53, 56, 57,
	58, 59, 0, 63, 64, 66, 67, 68, 0, 0,
	23, 24, 71, 72, 31, 0, 73, 75, 0, 36,
	27, 45, 0, 0, 0, 0, 95, 101, 103, 0,
	46, 0, 69, 79, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 32, 74, 76, 77, 0,
	91, 92, 93, 94, 106, 62, 80, 81, 82, 83,
	84, 85, 86, 87, 88, 89, 90, 78,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79, 80, 81,
	82, 83, 84, 85,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:83
		{
			logDebugGrammar("INPUT - %v", yylex.(*lexerWrapper).p)
			yylex.(*lexerWrapper).p = yyDollar[1].p
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:90
		{
			logDebugGrammar("PROGRAM - single")
			yyVAL.p = yyDollar[1].p
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:96
		{
			logDebugGrammar("PROGRAM - multi")
			yyDollar[1].p.Combine(yyDollar[2].p)
			yyVAL.p = yyDollar[1].p
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:104
		{
			p := &prog{}
			logDebugGrammar("P - decl")
			for _, decl := range yyDollar[1].declarations {
				p.Declare(decl)
			}
			yyVAL.p = p
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:114
		{
			p := &prog{}
			logDebugGrammar("P - rdef")
			p.DefineRoutine(yyDollar[1].r)
			yyVAL.p = p
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:122
		{
			p := &prog{}
			logDebugGrammar("P - gdef")
			p.DefineGroup(yyDollar[1].g)
			yyVAL.p = p
		}
	case 7:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:130
		{
			logDebugGrammar("P - backwardmode")
			yyDollar[3].p.SetBackwardMode()
			yyVAL.p = yyDollar[3].p
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:137
		{
			if utf8.RuneCountInString(yyDollar[1].s) == 2 {
				logDebugGrammar("P - stringescapes")
				first, len := utf8.DecodeRuneInString(yyDollar[1].s)
				second, len := utf8.DecodeRuneInString(yyDollar[1].s[len:])
				yylex.(*lexerWrapper).lex.(*snowConeLex).SetStringEscapes(first, second)
				yylex.(*lexerWrapper).sd.SetStringEscapes(first, second)
			} else {
				logDebugGrammar("P - stringescapes rune count NOT 2!!!")
			}
			yyVAL.p = &prog{}
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:151
		{
			logDebugGrammar("P - stringedef")
			replacedLiteral := yylex.(*lexerWrapper).sd.ReplaceInLiteral(yyDollar[3].s)
			yylex.(*lexerWrapper).sd.Define(yyDollar[1].s, replacedLiteral)
			yyVAL.p = &prog{}
		}
	case 10:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line snowball.y:160
		{

		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:165
		{
			logDebugGrammar("STRINGDEFLITERALTYPE - hex")
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:170
		{
			logDebugGrammar("STRINGDEFLITERALTYPE - decimal")
		}
	case 13:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:176
		{
			logDebugGrammar("DECLARATION - strings named: %v", yyDollar[3].strings)
			for _, name := range yyDollar[3].strings {
				yyVAL.declarations = append(yyVAL.declarations, &decl{
					name: name,
					typ:  sstring,
				})
			}
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:187
		{
			logDebugGrammar("DECLARATION - integers")
			for _, name := range yyDollar[3].strings {
				yyVAL.declarations = append(yyVAL.declarations, &decl{
					name: name,
					typ:  sinteger,
				})
			}
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:198
		{
			logDebugGrammar("DECLARATION - booleans")
			logDebugGrammar("DECLARATION - integers")
			for _, name := range yyDollar[3].strings {
				yyVAL.declarations = append(yyVAL.declarations, &decl{
					name: name,
					typ:  sboolean,
				})
			}
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:210
		{
			logDebugGrammar("DECLARATION - routines")
			for _, name := range yyDollar[3].strings {
				yyVAL.declarations = append(yyVAL.declarations, &decl{
					name: name,
					typ:  sroutine,
				})
			}
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:221
		{
			logDebugGrammar("DECLARATION - externals")
			for _, name := range yyDollar[3].strings {
				yyVAL.declarations = append(yyVAL.declarations, &decl{
					name: name,
					typ:  sexternal,
				})
			}
		}
	case 18:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:232
		{
			logDebugGrammar("DECLARATION - groupings")
			for _, name := range yyDollar[3].strings {
				yyVAL.declarations = append(yyVAL.declarations, &decl{
					name: name,
					typ:  sgrouping,
				})
			}
		}
	case 19:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:245
		{
			logDebugGrammar("RDEF")
			yyVAL.r = &routine{name: yyDollar[2].s, comm: yyDollar[4].no}
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:252
		{
			logDebugGrammar("NAMEORLITERAL - name")
			yyVAL.no = &name{val: yyDollar[1].s}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:258
		{
			logDebugGrammar("NAMEORLITERAL - literal")
			replacedLiteral := yylex.(*lexerWrapper).sd.ReplaceInLiteral(yyDollar[1].s)
			yyVAL.no = &sliteral{val: replacedLiteral}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:266
		{
			logDebugGrammar("GPLUSMINUSLIST - single")
			yyVAL.gitems = groupitems{&groupitem{item: yyDollar[1].no}}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:272
		{
			logDebugGrammar("GPLUSMINUSLIST - multi")
			//$$ = append($3, &groupitem{item: $1})
			yyVAL.gitems = append(yyDollar[3].gitems, nil)
			copy(yyVAL.gitems[1:], yyVAL.gitems[0:])
			yyVAL.gitems[0] = &groupitem{item: yyDollar[1].no}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:281
		{
			logDebugGrammar("GPLUSMINUSLIST - MINUS multi")
			//$$ = append($3, &groupitem{item: $1, minus:true})
			yyVAL.gitems = append(yyDollar[3].gitems, nil)
			copy(yyVAL.gitems[1:], yyVAL.gitems[0:])
			yyVAL.gitems[0] = &groupitem{item: yyDollar[1].no}
			yyVAL.gitems[1].minus = true
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:292
		{
			yyVAL.g = &grouping{name: yyDollar[2].s, children: yyDollar[3].gitems}
			logDebugGrammar("GDEF - %v", yyVAL.g)
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:299
		{
			logDebugGrammar("COMMANDS - single")
			yyVAL.cl = commands{yyDollar[1].no}
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:305
		{
			logDebugGrammar("COMMANDS - multi")
			yyVAL.cl = append(yyDollar[2].cl, nil)
			copy(yyVAL.cl[1:], yyVAL.cl[0:])
			yyVAL.cl[0] = yyDollar[1].no
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:314
		{
			logDebugGrammar("COMMANDFACTOR - s")
			yyVAL.no = yyDollar[1].no
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:320
		{
			logDebugGrammar("COMMAND - icommand")
			yyVAL.no = yyDollar[1].ic
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:326
		{
			logDebugGrammar("COMMAND - scommand")
			yyVAL.no = yyDollar[1].sc
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:332
		{
			logDebugGrammar("COMMAND - among empty")
			yyVAL.no = &among{}
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:338
		{
			logDebugGrammar("COMMAND - among list")
			yyVAL.no = &among{children: yyDollar[3].ai}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:344
		{
			logDebugGrammar("COMMANDFACTOR - true")
			yyVAL.no = &bliteral{val: true}
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:350
		{
			logDebugGrammar("COMMANDFACTOR - false")
			yyVAL.no = &bliteral{val: false}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:356
		{
			logDebugGrammar("COMMANDFACTOR - paren empty")
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:361
		{
			logDebugGrammar("COMMANDFACTOR - paren commands")
			yyVAL.no = yyDollar[2].cl
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:367
		{
			logDebugGrammar("COMMAND - not")
			yyVAL.no = &unaryCommand{command: "not", operandCommand: yyDollar[2].no}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:373
		{
			logDebugGrammar("COMMAND - test")
			yyVAL.no = &unaryCommand{command: "test", operandCommand: yyDollar[2].no}
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:379
		{
			logDebugGrammar("COMMAND - try")
			yyVAL.no = &unaryCommand{command: "try", operandCommand: yyDollar[2].no}
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:385
		{
			logDebugGrammar("COMMAND - do")
			yyVAL.no = &unaryCommand{command: "do", operandCommand: yyDollar[2].no}
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:391
		{
			logDebugGrammar("COMMAND - unary fail")
			yyVAL.no = &unaryCommand{command: "fail", operandCommand: yyDollar[2].no}
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:397
		{
			logDebugGrammar("COMMAND - goto")
			yyVAL.no = &unaryCommand{command: "goto", operandCommand: yyDollar[2].no}
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:403
		{
			logDebugGrammar("COMMAND - unary gopast")
			yyVAL.no = &unaryCommand{command: "gopast", operandCommand: yyDollar[2].no}
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:409
		{
			logDebugGrammar("COMMAND - unary repeat")
			yyVAL.no = &unaryCommand{command: "repeat", operandCommand: yyDollar[2].no}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:415
		{
			logDebugGrammar("COMMAND - loop ae")
			yyVAL.no = &loop{n: yyDollar[2].aexpr, operand: yyDollar[3].no}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:421
		{
			logDebugGrammar("COMMAND - loop ae")
			yyVAL.no = &loop{n: yyDollar[2].aexpr, operand: yyDollar[3].no, extra: true}
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:427
		{
			logDebugGrammar("COMMAND - insert")
			yyVAL.no = &unaryCommand{command: "insert", operandCommand: yyDollar[2].no}
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:433
		{
			logDebugGrammar("COMMAND - attach")
			yyVAL.no = &unaryCommand{command: "attach", operandCommand: yyDollar[2].no}
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:439
		{
			logDebugGrammar("COMMAND - replace")
			yyVAL.no = &unaryCommand{command: "replace", operandCommand: yyDollar[2].no}
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:445
		{
			logDebugGrammar("COMMAND - delete")
			yyVAL.no = &nilaryCommand{operator: "delete"}
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:451
		{
			logDebugGrammar("COMMAND - hop")
			yyVAL.no = &unaryCommand{command: "hop", operandCommand: yyDollar[2].aexpr}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:457
		{
			logDebugGrammar("COMMAND - next")
			yyVAL.no = &nilaryCommand{operator: "next"}
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:463
		{
			logDebugGrammar("COMMAND - assign right")
			yyVAL.no = &unaryCommand{command: "assignr", operandName: &name{val: yyDollar[2].s}}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:469
		{
			logDebugGrammar("COMMAND - lbracket")
			yyVAL.no = &nilaryCommand{operator: "["}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:475
		{
			logDebugGrammar("COMMAND - rbracket")
			yyVAL.no = &nilaryCommand{operator: "]"}
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:481
		{
			logDebugGrammar("COMMAND - move right")
			yyVAL.no = &unaryCommand{command: "mover", operandName: &name{val: yyDollar[2].s}}
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:487
		{
			logDebugGrammar("COMMAND - setmark")
			yyVAL.no = &unaryCommand{command: "setmark", operandName: &name{val: yyDollar[2].s}}
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:493
		{
			logDebugGrammar("COMMAND - tomark")
			yyVAL.no = &unaryCommand{command: "tomark", operandAe: yyDollar[2].aexpr}
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:499
		{
			logDebugGrammar("COMMAND - atmark")
			yyVAL.no = &unaryCommand{command: "atmark", operandAe: yyDollar[2].aexpr}
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:505
		{
			logDebugGrammar("COMMAND - tolimit")
			yyVAL.no = &nilaryCommand{operator: "tolimit"}
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:511
		{
			logDebugGrammar("COMMAND - atlimit")
			yyVAL.no = &nilaryCommand{operator: "atlimit"}
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:517
		{
			logDebugGrammar("COMMAND - setlimit")
			yyVAL.no = &binaryCommand{left: yyDollar[2].no, operator: "setlimitfor", right: yyDollar[4].no}
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:523
		{
			logDebugGrammar("COMMAND - backwards")
			yyVAL.no = &unaryCommand{command: "backwards", operandCommand: yyDollar[2].no}
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:529
		{
			logDebugGrammar("COMMAND - reverse")
			yyVAL.no = &unaryCommand{command: "reverse", operandCommand: yyDollar[2].no}
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:535
		{
			logDebugGrammar("COMMAND - substring")
			yyVAL.no = &nilaryCommand{operator: "substring"}
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:541
		{
			logDebugGrammar("COMMAND - set")
			yyVAL.no = &unaryCommand{command: "set", operandName: &name{val: yyDollar[2].s}}
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:547
		{
			logDebugGrammar("COMMAND - unset")
			yyVAL.no = &unaryCommand{command: "unset", operandName: &name{val: yyDollar[2].s}}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:553
		{
			logDebugGrammar("COMMAND - non")
			yyVAL.no = &non{gname: yyDollar[2].s}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:559
		{
			logDebugGrammar("COMMAND - non minus")
			yyVAL.no = &non{gname: yyDollar[3].s, minus: true}
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:565
		{
			logDebugGrammar("COMMAND - question")
			yyVAL.no = &nilaryCommand{operator: "?"}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:571
		{
			logDebugGrammar("COMMANDTERM - or")
			yyVAL.no = &binaryCommand{left: yyDollar[1].no, operator: "or", right: yyDollar[3].no}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:577
		{
			logDebugGrammar("COMMANDTERM - and")
			yyVAL.no = &binaryCommand{left: yyDollar[1].no, operator: "and", right: yyDollar[3].no}
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:585
		{
			logDebugGrammar("AMONGLIST - single")
			yyVAL.ai = amongitems{yyDollar[1].a}
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:590
		{
			logDebugGrammar("AMONGLIST - multi")
			yyVAL.ai = append(yyDollar[2].ai, nil)
			copy(yyVAL.ai[1:], yyVAL.ai[0:])
			yyVAL.ai[0] = yyDollar[1].a
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:599
		{
			logDebugGrammar("AMONGITEM - literal")
			replacedLiteral := yylex.(*lexerWrapper).sd.ReplaceInLiteral(yyDollar[1].s)
			yyVAL.a = &amongitem{slit: &sliteral{val: replacedLiteral}}
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:606
		{
			logDebugGrammar("AMONGITEM - literal name")
			replacedLiteral := yylex.(*lexerWrapper).sd.ReplaceInLiteral(yyDollar[1].s)
			yyVAL.a = &amongitem{slit: &sliteral{val: replacedLiteral}, rname: yyDollar[2].s}
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:613
		{
			logDebugGrammar("AMONGITEM - paren empty")
			yyVAL.a = &amongitem{}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:619
		{
			logDebugGrammar("AMONGITEM - paren command")
			yyVAL.a = &amongitem{comm: yyDollar[2].cl}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:626
		{
			logDebugGrammar("SCOMMAND")
			yyVAL.sc = &sCommand{name: &name{val: yyDollar[2].s}, operand: yyDollar[3].no}
		}
	case 80:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:633
		{
			logDebugGrammar("ICOMMAND - assign")
			yyVAL.ic = &iCommand{name: &name{val: yyDollar[2].s}, operator: "=", operand: yyDollar[4].aexpr}
		}
	case 81:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:639
		{
			logDebugGrammar("ICOMMAND - plus assign")
			yyVAL.ic = &iCommand{name: &name{val: yyDollar[2].s}, operator: "+=", operand: yyDollar[4].aexpr}
		}
	case 82:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:645
		{
			logDebugGrammar("ICOMMAND - minus assign")
			yyVAL.ic = &iCommand{name: &name{val: yyDollar[2].s}, operator: "-=", operand: yyDollar[4].aexpr}
		}
	case 83:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:651
		{
			logDebugGrammar("ICOMMAND - mult assign")
			yyVAL.ic = &iCommand{name: &name{val: yyDollar[2].s}, operator: "*=", operand: yyDollar[4].aexpr}
		}
	case 84:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:657
		{
			logDebugGrammar("ICOMMAND - div assign")
			yyVAL.ic = &iCommand{name: &name{val: yyDollar[2].s}, operator: "/=", operand: yyDollar[4].aexpr}
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:663
		{
			logDebugGrammar("ICOMMAND - eq")
			yyVAL.ic = &iCommand{name: &name{val: yyDollar[2].s}, operator: "==", operand: yyDollar[4].aexpr}
		}
	case 86:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:669
		{
			logDebugGrammar("ICOMMAND - neq")
			yyVAL.ic = &iCommand{name: &name{val: yyDollar[2].s}, operator: "!=", operand: yyDollar[4].aexpr}
		}
	case 87:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:675
		{
			logDebugGrammar("ICOMMAND - greater than")
			yyVAL.ic = &iCommand{name: &name{val: yyDollar[2].s}, operator: ">", operand: yyDollar[4].aexpr}
		}
	case 88:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:681
		{
			logDebugGrammar("ICOMMAND - less than")
			yyVAL.ic = &iCommand{name: &name{val: yyDollar[2].s}, operator: "<", operand: yyDollar[4].aexpr}
		}
	case 89:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:687
		{
			logDebugGrammar("ICOMMAND - greater than or eq")
			yyVAL.ic = &iCommand{name: &name{val: yyDollar[2].s}, operator: ">=", operand: yyDollar[4].aexpr}
		}
	case 90:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:693
		{
			logDebugGrammar("ICOMMAND - less than or eq")
			yyVAL.ic = &iCommand{name: &name{val: yyDollar[2].s}, operator: "<=", operand: yyDollar[4].aexpr}
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:700
		{
			logDebugGrammar("AE - plus")
			yyVAL.aexpr = &binaryAe{left: yyDollar[1].aexpr, operator: "+", right: yyDollar[3].aexpr}
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:706
		{
			logDebugGrammar("AE - minus")
			yyVAL.aexpr = &binaryAe{left: yyDollar[1].aexpr, operator: "-", right: yyDollar[3].aexpr}
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:712
		{
			logDebugGrammar("TERM - mult")
			yyVAL.aexpr = &binaryAe{left: yyDollar[1].aexpr, operator: "*", right: yyDollar[3].aexpr}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:718
		{
			logDebugGrammar("TERM - div")
			yyVAL.aexpr = &binaryAe{left: yyDollar[1].aexpr, operator: "/", right: yyDollar[3].aexpr}
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:724
		{
			logDebugGrammar("AE - unary minus")
			yyVAL.aexpr = &unaryAe{operator: "uminus", operand: yyDollar[2].aexpr}
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:730
		{
			logDebugGrammar("AE - maxint")
			yyVAL.aexpr = &nilaryAe{operator: "maxint"}
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:736
		{
			logDebugGrammar("AE - minint")
			yyVAL.aexpr = &nilaryAe{operator: "minint"}
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:742
		{
			logDebugGrammar("AE - cursor")
			yyVAL.aexpr = &nilaryAe{operator: "cursor"}
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:748
		{
			logDebugGrammar("AE - limit")
			yyVAL.aexpr = &nilaryAe{operator: "limit"}
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:754
		{
			logDebugGrammar("AE - size")
			yyVAL.aexpr = &nilaryAe{operator: "size"}
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:760
		{
			logDebugGrammar("AE - sizeof name")
			yyVAL.aexpr = &unaryAe{operator: "sizeof", operand: &name{val: yyDollar[2].s}}
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:766
		{
			logDebugGrammar("AE - len")
			yyVAL.aexpr = &nilaryAe{operator: "len"}
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:772
		{
			logDebugGrammar("AE - leno name")
			yyVAL.aexpr = &unaryAe{operator: "lenof", operand: &name{val: yyDollar[2].s}}
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:778
		{
			logDebugGrammar("FACTOR - name")
			yyVAL.aexpr = &name{val: yyDollar[1].s}
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:784
		{
			logDebugGrammar("FACTOR - number")
			yyVAL.aexpr = &nliteral{val: yyDollar[1].n}
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:790
		{
			logDebugGrammar("FACTOR - parens")
			yyVAL.aexpr = yyDollar[2].aexpr
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:798
		{
			logDebugGrammar("NAMES - single")
			yyVAL.strings = []string{yyDollar[1].s}
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:804
		{
			logDebugGrammar("NAMEs - multi")
			yyVAL.strings = append(yyDollar[2].strings, "")
			copy(yyVAL.strings[1:], yyVAL.strings[0:])
			yyVAL.strings[0] = yyDollar[1].s
		}
	}
	goto yystack /* stack new state and value */
}
