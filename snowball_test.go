package snowcone

import (
	"reflect"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	// DebugLexer = true
	// DebugParser = true
	// Logger = log.New(os.Stdout, "", log.LstdFlags)
	tests := []struct {
		in string
	}{
		{
			in: `
		  strings ( p1 p2 )
		  `,
		},
		{
			in: `
		  define cat as $bob = 1+1
		  `,
		},
		{
			in: `
      strings ( p1 p2 )
      define cat as $bob = 1+1
      `,
		},
		{
			in: porterStemmerSrc,
		},
	}

	for _, test := range tests {
		err := Parse(strings.NewReader(test.in))
		if err != nil {
			t.Errorf("parse error: %v", err)
		}
	}
}

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
	}

	for _, test := range tests {
		var tokenTypes []int
		var tokens []yySymType
		r := strings.NewReader(test.in)
		l := newLexer(r)
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

var porterStemmerSrc = `
integers ( p1 p2 )
booleans ( Y_found )

routines (
   shortv
   R1 R2
   Step_1a Step_1b Step_1c Step_2 Step_3 Step_4 Step_5a Step_5b
)

externals ( stem )

groupings ( v v_WXY )

define v        'aeiouy'
define v_WXY    v + 'wxY'

backwardmode (

    define shortv as ( non-v_WXY v non-v )

    define R1 as $p1 <= cursor
    define R2 as $p2 <= cursor

    define Step_1a as (
        [substring] among (
            'sses' (<-'ss')
            'ies'  (<-'i')
            'ss'   ()
            's'    (delete)
        )
    )

    define Step_1b as (
        [substring] among (
            'eed'  (R1 <-'ee')
            'ed'
            'ing' (
                test gopast v  delete
                test substring among(
                    'at' 'bl' 'iz'
                         (<+ 'e')
                    'bb' 'dd' 'ff' 'gg' 'mm' 'nn' 'pp' 'rr' 'tt'
                    // ignoring double c, h, j, k, q, v, w, and x
                         ([next]  delete)
                    ''   (atmark p1  test shortv  <+ 'e')
                )
            )
        )
    )

    define Step_1c as (
        ['y' or 'Y']
        gopast v
        <-'i'
    )

    define Step_2 as (
        [substring] R1 among (
            'tional'  (<-'tion')
            'enci'    (<-'ence')
            'anci'    (<-'ance')
            'abli'    (<-'able')
            'entli'   (<-'ent')
            'eli'     (<-'e')
            'izer' 'ization'
                      (<-'ize')
            'ational' 'ation' 'ator'
                      (<-'ate')
            'alli'    (<-'al')
            'alism' 'aliti'
                      (<-'al')
            'fulness' (<-'ful')
            'ousli' 'ousness'
                      (<-'ous')
            'iveness' 'iviti'
                      (<-'ive')
            'biliti'  (<-'ble')
        )
    )

    define Step_3 as (
        [substring] R1 among (
            'alize'   (<-'al')
            'icate' 'iciti' 'ical'
                      (<-'ic')
            'ative' 'ful' 'ness'
                      (delete)
        )
    )

    define Step_4 as (
        [substring] R2 among (
            'al' 'ance' 'ence' 'er' 'ic' 'able' 'ible' 'ant' 'ement'
            'ment' 'ent' 'ou' 'ism' 'ate' 'iti' 'ous' 'ive' 'ize'
                      (delete)
            'ion'     ('s' or 't' delete)
        )
    )

    define Step_5a as (
        ['e']
        R2 or (R1 not shortv)
        delete
    )

    define Step_5b as (
        ['l']
        R2 'l'
        delete
    )
)

define stem as (

    unset Y_found
    do ( ['y'] <-'Y' set Y_found)
    do repeat(goto (v ['y']) <-'Y' set Y_found)

    $p1 = limit
    $p2 = limit
    do(
        gopast v  gopast non-v  setmark p1
        gopast v  gopast non-v  setmark p2
    )

    backwards (
        do Step_1a
        do Step_1b
        do Step_1c
        do Step_2
        do Step_3
        do Step_4
        do Step_5a
        do Step_5b
    )

    do(Y_found  repeat(goto (['Y']) <-'y'))

)`
