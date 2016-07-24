package snowcone

import __yyfmt__ "fmt"

//line snowball.y:2
import "fmt"

func logDebugGrammar(format string, v ...interface{}) {
	if DebugParser {
		Logger.Printf(fmt.Sprintf("PARSER %s", format), v...)
	}
}

//line snowball.y:13
type yySymType struct {
	yys int
	s   string
	n   int
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

const yyNprod = 114
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 482

var yyAct = [...]int{

	106, 142, 157, 107, 155, 94, 39, 10, 11, 12,
	13, 14, 15, 20, 21, 16, 138, 139, 151, 152,
	146, 147, 119, 118, 31, 96, 97, 2, 119, 118,
	120, 17, 200, 183, 40, 108, 120, 143, 109, 110,
	111, 112, 113, 114, 115, 116, 29, 136, 33, 34,
	35, 36, 37, 159, 42, 133, 41, 45, 181, 137,
	185, 175, 50, 156, 27, 121, 38, 186, 160, 49,
	125, 122, 123, 124, 48, 47, 51, 128, 129, 7,
	46, 8, 9, 44, 43, 26, 25, 24, 23, 22,
	18, 150, 149, 98, 99, 100, 101, 102, 103, 104,
	105, 40, 144, 42, 140, 41, 135, 134, 127, 148,
	126, 32, 28, 159, 30, 117, 89, 158, 130, 131,
	132, 153, 54, 53, 52, 95, 19, 6, 5, 4,
	3, 1, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 145, 176, 0, 177, 178, 0,
	0, 0, 0, 0, 0, 179, 180, 0, 0, 154,
	0, 184, 0, 0, 0, 187, 188, 189, 190, 191,
	192, 193, 194, 195, 196, 197, 161, 162, 163, 0,
	42, 0, 41, 0, 0, 0, 0, 0, 199, 93,
	0, 0, 0, 182, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 90, 164, 165, 166, 167,
	168, 169, 170, 171, 172, 173, 174, 0, 0, 55,
	56, 57, 58, 59, 60, 61, 62, 63, 64, 65,
	66, 67, 68, 69, 70, 71, 72, 73, 74, 75,
	76, 77, 78, 79, 80, 0, 81, 82, 83, 84,
	85, 86, 87, 91, 92, 42, 88, 41, 0, 0,
	0, 0, 0, 0, 93, 198, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	90, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 55, 56, 57, 58, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
	71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
	0, 81, 82, 83, 84, 85, 86, 87, 91, 92,
	42, 88, 41, 0, 0, 0, 0, 0, 0, 93,
	141, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 90, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 55,
	56, 57, 58, 59, 60, 61, 62, 63, 64, 65,
	66, 67, 68, 69, 70, 71, 72, 73, 74, 75,
	76, 77, 78, 79, 80, 0, 81, 82, 83, 84,
	85, 86, 87, 91, 92, 42, 88, 41, 0, 0,
	0, 0, 0, 0, 93, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	90, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 55, 56, 57, 58, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
	71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
	0, 81, 82, 83, 84, 85, 86, 87, 91, 92,
	0, 88,
}
var yyPact = [...]int{

	0, -1000, -1000, 0, -1000, -1000, -1000, 77, -1000, -70,
	76, 75, 74, 73, 72, 51, 106, -1000, 0, 110,
	-1000, -1000, 105, 105, 105, 105, 105, 105, 50, 70,
	-1000, 69, 105, 66, 61, 60, 55, 48, 401, -1000,
	8, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 401, 401, 401, 401, 401,
	401, 401, 401, 17, 17, 99, 99, 99, -1000, 17,
	-1000, -1000, -1000, -1000, 104, 102, 17, 17, -1000, -1000,
	401, 401, 401, -1000, 42, 101, 100, 41, -1000, -25,
	98, -1000, -1000, 326, -1000, 99, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 401, 3, 17, -1000,
	-1000, -1000, -1000, -1000, 86, -1000, 85, -1, -1000, -1000,
	17, 401, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-65, -1000, -1000, 49, -1000, -1000, -1000, 62, 401, 401,
	176, -1000, 47, 401, -1000, -1000, 17, 17, -1000, -1000,
	-1000, 23, 23, 44, -1000, 401, -1000, 19, 109, 54,
	-1000, -1000, -1000, -1000, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 251, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 18,
	-1000,
}
var yyPgo = [...]int{

	0, 131, 27, 130, 129, 128, 127, 126, 24, 37,
	5, 125, 6, 1, 124, 123, 122, 0, 2, 117,
	116, 3, 115,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 3, 3, 3, 3, 3, 3,
	7, 7, 7, 4, 4, 4, 4, 4, 4, 5,
	10, 10, 11, 11, 12, 12, 6, 13, 13, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 18, 18,
	19, 19, 19, 19, 14, 14, 14, 20, 20, 20,
	20, 20, 16, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 21, 21, 21, 22,
	22, 22, 8, 8,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 1, 1, 1, 4, 1, 3,
	0, 1, 1, 4, 4, 4, 4, 4, 4, 4,
	1, 1, 1, 1, 1, 3, 3, 1, 2, 1,
	1, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	3, 3, 2, 2, 2, 1, 2, 1, 1, 1,
	1, 2, 2, 2, 2, 1, 1, 4, 2, 2,
	1, 3, 4, 2, 2, 2, 3, 1, 1, 2,
	1, 2, 3, 4, 1, 3, 3, 1, 1, 2,
	3, 1, 3, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 3, 3, 1, 2, 1, 1,
	1, 1, 1, 2, 1, 2, 1, 3, 3, 1,
	1, 3, 1, 2,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, 79, 81, 82,
	7, 8, 9, 10, 11, 12, 15, -2, 13, -7,
	83, 84, 13, 13, 13, 13, 13, 13, 6, -2,
	4, -8, 6, -8, -8, -8, -8, -8, 16, -12,
	-10, 6, 4, 14, 14, -8, 14, 14, 14, 14,
	14, -9, -14, -15, -16, 43, 44, 45, 46, 47,
	48, 49, 50, 51, 52, 53, 54, 55, 56, 57,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 70, 71, 72, 73, 74, 75, 76, 80, -20,
	29, 77, 78, 13, -10, -11, 17, 18, -9, -9,
	-9, -9, -9, -9, -9, -9, -17, -21, 18, 21,
	22, 23, 24, 25, 26, 27, 28, -22, 6, 5,
	13, -17, -10, -10, -10, -17, 6, 6, -17, -17,
	-9, -9, -9, 13, 6, 6, 6, 18, 41, 42,
	6, 14, -13, -9, -12, -9, 17, 18, -17, 6,
	6, 19, 20, -17, -9, 69, 14, -18, -19, 4,
	6, -9, -9, -9, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 14, -13, -17, -17, -21,
	-21, 14, -9, 14, -18, 6, 13, -17, -17, -17,
	-17, -17, -17, -17, -17, -17, -17, -17, 14, -13,
	14,
}
var yyDef = [...]int{

	0, -2, 1, 2, 4, 5, 6, 0, 8, 10,
	0, 0, 0, 0, 0, 0, 0, 3, 0, 0,
	11, 12, 0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 112, 0, 0, 0, 0, 0, 0, 26,
	24, 20, 21, 7, 13, 113, 14, 15, 16, 17,
	18, 19, 29, 30, 31, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 45, 0,
	47, 48, 49, 50, 0, 0, 0, 0, 55, 56,
	0, 0, 0, 60, 0, 0, 0, 0, 67, 74,
	0, 77, 78, 0, 81, 0, 22, 23, 32, 33,
	34, 35, 36, 37, 38, 39, 0, 96, 0, 98,
	99, 100, 101, 102, 0, 104, 0, 106, 109, 110,
	0, 0, 42, 43, 44, 46, 51, 52, 53, 54,
	0, 58, 59, 0, 63, 64, 65, 0, 0, 0,
	0, 79, 0, 27, 25, 40, 0, 0, 97, 103,
	105, 0, 0, 0, 41, 0, 61, 0, 68, 70,
	66, 75, 76, 82, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 80, 28, 94, 95, 107,
	108, 111, 57, 62, 69, 71, 0, 83, 84, 85,
	86, 87, 88, 89, 90, 91, 92, 93, 72, 0,
	73,
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
	82, 83, 84,
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
		//line snowball.y:37
		{
			logDebugGrammar("INPUT")
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:43
		{
			logDebugGrammar("PROGRAM - single")
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:48
		{
			logDebugGrammar("PROGRAM - multi")
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:54
		{
			logDebugGrammar("P - decl")
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:59
		{
			logDebugGrammar("P - rdef")
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:64
		{
			logDebugGrammar("P - gdef")
		}
	case 7:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:69
		{
			logDebugGrammar("P - backwardmode")
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:74
		{
			logDebugGrammar("P - stringescapes")
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:79
		{
			logDebugGrammar("P - stringedef")
		}
	case 10:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line snowball.y:85
		{

		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:90
		{
			logDebugGrammar("STRINGDEFLITERALTYPE - hex")
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:95
		{
			logDebugGrammar("STRINGDEFLITERALTYPE - decimal")
		}
	case 13:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:101
		{
			logDebugGrammar("DECLARATION - strings")
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:106
		{
			logDebugGrammar("DECLARATION - integers")
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:111
		{
			logDebugGrammar("DECLARATION - booleans")
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:116
		{
			logDebugGrammar("DECLARATION - routines")
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:121
		{
			logDebugGrammar("DECLARATION - externals")
		}
	case 18:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:126
		{
			logDebugGrammar("DECLARATION - groupings")
		}
	case 19:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:133
		{
			logDebugGrammar("RDEF")
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:139
		{
			logDebugGrammar("NAMEORLITERAL - name")
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:144
		{
			logDebugGrammar("NAMEORLITERAL - literal")
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:150
		{
			logDebugGrammar("PLUSORMINUS - plus")
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:155
		{
			logDebugGrammar("PLUSORMINUS - minus")
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:161
		{
			logDebugGrammar("GPLUSMINUSLIST - single")
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:166
		{
			logDebugGrammar("GPLUSMINUSLIST - multi")
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:172
		{
			logDebugGrammar("GDEF")
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:178
		{
			logDebugGrammar("COMMANDS - single")
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:183
		{
			logDebugGrammar("COMMANDS - multi")
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:189
		{
			logDebugGrammar("COMMAND - command term")
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:194
		{
			logDebugGrammar("COMMAND - icommand")
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:199
		{
			logDebugGrammar("COMMAND - scommand")
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:204
		{
			logDebugGrammar("COMMAND - unary not")
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:209
		{
			logDebugGrammar("COMMAND - unary test")
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:214
		{
			logDebugGrammar("COMMAND - unary try")
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:219
		{
			logDebugGrammar("COMMAND - unary do")
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:224
		{
			logDebugGrammar("COMMAND - unary fail")
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:229
		{
			logDebugGrammar("COMMAND - unary goto")
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:234
		{
			logDebugGrammar("COMMAND - unary gopast")
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:239
		{
			logDebugGrammar("COMMAND - unary repeat")
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:244
		{
			logDebugGrammar("COMMAND - loop ae")
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:249
		{
			logDebugGrammar("COMMAND - loop ae")
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:254
		{
			logDebugGrammar("COMMAND - insert")
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:259
		{
			logDebugGrammar("COMMAND - attach")
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:264
		{
			logDebugGrammar("COMMAND - replace")
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:269
		{
			logDebugGrammar("COMMAND - delete")
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:274
		{
			logDebugGrammar("COMMAND - hop")
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:279
		{
			logDebugGrammar("COMMAND - next")
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:284
		{
			logDebugGrammar("COMMAND - assign right")
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:289
		{
			logDebugGrammar("COMMAND - lbracket")
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:294
		{
			logDebugGrammar("COMMAND - rbracket")
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:299
		{
			logDebugGrammar("COMMAND - move right")
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:304
		{
			logDebugGrammar("COMMAND - setmark")
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:309
		{
			logDebugGrammar("COMMAND - tomark")
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:314
		{
			logDebugGrammar("COMMAND - atmark")
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:319
		{
			logDebugGrammar("COMMAND - tolimit")
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:324
		{
			logDebugGrammar("COMMAND - atlimit")
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:329
		{
			logDebugGrammar("COMMAND - setlimit")
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:334
		{
			logDebugGrammar("COMMAND - backwards")
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:339
		{
			logDebugGrammar("COMMAND - reverse")
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:344
		{
			logDebugGrammar("COMMAND - substring")
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:349
		{
			logDebugGrammar("COMMAND - among empty")
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:354
		{
			logDebugGrammar("COMMAND - among list")
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:359
		{
			logDebugGrammar("COMMAND - set")
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:364
		{
			logDebugGrammar("COMMAND - unset")
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:369
		{
			logDebugGrammar("COMMAND - non")
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:374
		{
			logDebugGrammar("COMMAND - non minus")
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:379
		{
			logDebugGrammar("COMMAND - question")
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:386
		{
			logDebugGrammar("AMONGLIST - single")
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:390
		{
			logDebugGrammar("AMONGLIST - multi")
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:396
		{
			logDebugGrammar("AMONGITEM - literal")
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:401
		{
			logDebugGrammar("AMONGITEM - literal name")
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:406
		{
			logDebugGrammar("AMONGITEM - paren empty")
		}
	case 73:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:411
		{
			logDebugGrammar("AMONGITEM - paren command")
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:417
		{
			logDebugGrammar("COMMANDTERM - command factor")
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:422
		{
			logDebugGrammar("COMMANDTERM - or")
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:427
		{
			logDebugGrammar("COMMANDTERM - and")
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:433
		{
			logDebugGrammar("COMMANDFACTOR - true")
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:438
		{
			logDebugGrammar("COMMANDFACTOR - false")
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:443
		{
			logDebugGrammar("COMMANDFACTOR - paren empty")
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:448
		{
			logDebugGrammar("COMMANDFACTOR - paren commands")
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:453
		{
			logDebugGrammar("COMMANDFACTOR - s")
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:460
		{
			logDebugGrammar("SCOMMAND")
		}
	case 83:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:466
		{
			logDebugGrammar("ICOMMAND - assign")
		}
	case 84:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:471
		{
			logDebugGrammar("ICOMMAND - plus assign")
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:476
		{
			logDebugGrammar("ICOMMAND - minus assign")
		}
	case 86:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:481
		{
			logDebugGrammar("ICOMMAND - mult assign")
		}
	case 87:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:486
		{
			logDebugGrammar("ICOMMAND - div assign")
		}
	case 88:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:491
		{
			logDebugGrammar("ICOMMAND - eq")
		}
	case 89:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:496
		{
			logDebugGrammar("ICOMMAND - neq")
		}
	case 90:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:501
		{
			logDebugGrammar("ICOMMAND - greater than")
		}
	case 91:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:506
		{
			logDebugGrammar("ICOMMAND - less than")
		}
	case 92:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:511
		{
			logDebugGrammar("ICOMMAND - greater than or eq")
		}
	case 93:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line snowball.y:516
		{
			logDebugGrammar("ICOMMAND - less than or eq")
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:522
		{
			logDebugGrammar("AE - plus")
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:527
		{
			logDebugGrammar("AE - minus")
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:532
		{
			logDebugGrammar("AE - term")
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:537
		{
			logDebugGrammar("AE - unary minus")
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:542
		{
			logDebugGrammar("AE - maxint")
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:547
		{
			logDebugGrammar("AE - minint")
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:552
		{
			logDebugGrammar("AE - cursor")
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:557
		{
			logDebugGrammar("AE - limit")
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:562
		{
			logDebugGrammar("AE - size")
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:567
		{
			logDebugGrammar("AE - sizeof name")
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:572
		{
			logDebugGrammar("AE - len")
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:577
		{
			logDebugGrammar("AE - leno name")
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:584
		{
			logDebugGrammar("TERM - factor")
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:589
		{
			logDebugGrammar("TERM - mult")
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:594
		{
			logDebugGrammar("TERM - div")
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:600
		{
			logDebugGrammar("FACTOR - name")
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:605
		{
			logDebugGrammar("FACTOR - number")
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line snowball.y:610
		{
			logDebugGrammar("FACTOR - parens")
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line snowball.y:616
		{
			logDebugGrammar("NAMES - single")
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line snowball.y:621
		{
			logDebugGrammar("NAMEs - multi")
		}
	}
	goto yystack /* stack new state and value */
}
