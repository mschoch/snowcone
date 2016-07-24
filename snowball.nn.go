package snowcone

import (
	"fmt"
	"strconv"
)
import (
	"bufio"
	"io"
	"strings"
)

type frame struct {
	i            int
	s            string
	line, column int
}
type lexer struct {
	// The lexer runs in its own goroutine, and communicates via channel 'ch'.
	ch chan frame
	// We record the level of nesting because the action could return, and a
	// subsequent call expects to pick up where it left off. In other words,
	// we're simulating a coroutine.
	// TODO: Support a channel-based variant that compatible with Go's yacc.
	stack []frame
	stale bool

	// The 'l' and 'c' fields were added for
	// https://github.com/wagerlabs/docker/blob/65694e801a7b80930961d70c69cba9f2465459be/buildfile.nex
	// Since then, I introduced the built-in Line() and Column() functions.
	l, c int

	parseResult interface{}

	// The following line makes it easy for scripts to insert fields in the
	// generated code.
	// [NEX_END_OF_LEXER_STRUCT]
}

// newLexerWithInit creates a new lexer object, runs the given callback on it,
// then returns it.
func newLexerWithInit(in io.Reader, initFun func(*lexer)) *lexer {
	type dfa struct {
		acc          []bool           // Accepting states.
		f            []func(rune) int // Transitions.
		startf, endf []int            // Transitions at start and end of input.
		nest         []dfa
	}
	yylex := new(lexer)
	if initFun != nil {
		initFun(yylex)
	}
	yylex.ch = make(chan frame)
	var scan func(in *bufio.Reader, ch chan frame, family []dfa, line, column int)
	scan = func(in *bufio.Reader, ch chan frame, family []dfa, line, column int) {
		// Index of DFA and length of highest-precedence match so far.
		matchi, matchn := 0, -1
		var buf []rune
		n := 0
		checkAccept := func(i int, st int) bool {
			// Higher precedence match? DFAs are run in parallel, so matchn is at most len(buf), hence we may omit the length equality check.
			if family[i].acc[st] && (matchn < n || matchi > i) {
				matchi, matchn = i, n
				return true
			}
			return false
		}
		var state [][2]int
		for i := 0; i < len(family); i++ {
			mark := make([]bool, len(family[i].startf))
			// Every DFA starts at state 0.
			st := 0
			for {
				state = append(state, [2]int{i, st})
				mark[st] = true
				// As we're at the start of input, follow all ^ transitions and append to our list of start states.
				st = family[i].startf[st]
				if -1 == st || mark[st] {
					break
				}
				// We only check for a match after at least one transition.
				checkAccept(i, st)
			}
		}
		atEOF := false
		for {
			if n == len(buf) && !atEOF {
				r, _, err := in.ReadRune()
				switch err {
				case io.EOF:
					atEOF = true
				case nil:
					buf = append(buf, r)
				default:
					panic(err)
				}
			}
			if !atEOF {
				r := buf[n]
				n++
				var nextState [][2]int
				for _, x := range state {
					x[1] = family[x[0]].f[x[1]](r)
					if -1 == x[1] {
						continue
					}
					nextState = append(nextState, x)
					checkAccept(x[0], x[1])
				}
				state = nextState
			} else {
			dollar: // Handle $.
				for _, x := range state {
					mark := make([]bool, len(family[x[0]].endf))
					for {
						mark[x[1]] = true
						x[1] = family[x[0]].endf[x[1]]
						if -1 == x[1] || mark[x[1]] {
							break
						}
						if checkAccept(x[0], x[1]) {
							// Unlike before, we can break off the search. Now that we're at the end, there's no need to maintain the state of each DFA.
							break dollar
						}
					}
				}
				state = nil
			}

			if state == nil {
				lcUpdate := func(r rune) {
					if r == '\n' {
						line++
						column = 0
					} else {
						column++
					}
				}
				// All DFAs stuck. Return last match if it exists, otherwise advance by one rune and restart all DFAs.
				if matchn == -1 {
					if len(buf) == 0 { // This can only happen at the end of input.
						break
					}
					lcUpdate(buf[0])
					buf = buf[1:]
				} else {
					text := string(buf[:matchn])
					buf = buf[matchn:]
					matchn = -1
					ch <- frame{matchi, text, line, column}
					if len(family[matchi].nest) > 0 {
						scan(bufio.NewReader(strings.NewReader(text)), ch, family[matchi].nest, line, column)
					}
					if atEOF {
						break
					}
					for _, r := range text {
						lcUpdate(r)
					}
				}
				n = 0
				for i := 0; i < len(family); i++ {
					state = append(state, [2]int{i, 0})
				}
			}
		}
		ch <- frame{-1, "", line, column}
	}
	go scan(bufio.NewReader(in), yylex.ch, []dfa{
		// strings
		{[]bool{false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return 1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return 3
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 105:
					return 4
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return 5
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return 6
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return 7
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// integers
		{[]bool{false, false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 103:
					return -1
				case 105:
					return 1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return 2
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 4
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 103:
					return 5
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 6
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return 7
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return 8
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// booleans
		{[]bool{false, false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return 1
				case 101:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				case 111:
					return 2
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				case 111:
					return 3
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 101:
					return -1
				case 108:
					return 4
				case 110:
					return -1
				case 111:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 101:
					return 5
				case 108:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 6
				case 98:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 110:
					return 7
				case 111:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 115:
					return 8
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 115:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// routines
		{[]bool{false, false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 114:
					return 1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 111:
					return 2
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 4
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return 5
				case 110:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return 6
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 7
				case 105:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return 8
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// externals
		{[]bool{false, false, false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return 1
				case 108:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 120:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 3
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return 4
				case 108:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				case 114:
					return 5
				case 115:
					return -1
				case 116:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 110:
					return 6
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 7
				case 101:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 108:
					return 8
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return 9
				case 116:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 120:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// groupings
		{[]bool{false, false, false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 103:
					return 1
				case 105:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return 2
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 111:
					return 3
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return 5
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 105:
					return 6
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return 7
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return 8
				case 105:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return 9
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// define
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 100:
					return 1
				case 101:
					return -1
				case 102:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return 2
				case 102:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return 3
				case 105:
					return -1
				case 110:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 105:
					return 4
				case 110:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 105:
					return -1
				case 110:
					return 5
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return 6
				case 102:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// as
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return 1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 115:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 115:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \(
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 40:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 40:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \)
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 41:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 41:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \+
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 43:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \-
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 45:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \*
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 42:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 42:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \/
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 47:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 47:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// maxint
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return -1
				case 109:
					return 1
				case 110:
					return -1
				case 116:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 2
				case 105:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 116:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 116:
					return -1
				case 120:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return 4
				case 109:
					return -1
				case 110:
					return -1
				case 116:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return -1
				case 109:
					return -1
				case 110:
					return 5
				case 116:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 116:
					return 6
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 116:
					return -1
				case 120:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// minint
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 109:
					return 1
				case 110:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return 2
				case 109:
					return -1
				case 110:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 109:
					return -1
				case 110:
					return 3
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return 4
				case 109:
					return -1
				case 110:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 109:
					return -1
				case 110:
					return 5
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 116:
					return 6
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// cursor
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 99:
					return 1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 111:
					return -1
				case 114:
					return 3
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return 4
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 111:
					return 5
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 111:
					return -1
				case 114:
					return 6
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// limit
		{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 108:
					return 1
				case 109:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return 2
				case 108:
					return -1
				case 109:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return 3
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return 4
				case 108:
					return -1
				case 109:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 116:
					return 5
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

		// size
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return -1
				case 115:
					return 1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return 2
				case 115:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return -1
				case 115:
					return -1
				case 122:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 4
				case 105:
					return -1
				case 115:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return -1
				case 115:
					return -1
				case 122:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// sizeof
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 105:
					return -1
				case 111:
					return -1
				case 115:
					return 1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 105:
					return 2
				case 111:
					return -1
				case 115:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 105:
					return -1
				case 111:
					return -1
				case 115:
					return -1
				case 122:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 4
				case 102:
					return -1
				case 105:
					return -1
				case 111:
					return -1
				case 115:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 105:
					return -1
				case 111:
					return 5
				case 115:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return 6
				case 105:
					return -1
				case 111:
					return -1
				case 115:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 105:
					return -1
				case 111:
					return -1
				case 115:
					return -1
				case 122:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// len
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 108:
					return 1
				case 110:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 2
				case 108:
					return -1
				case 110:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 108:
					return -1
				case 110:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// lenof
		{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 108:
					return 1
				case 110:
					return -1
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 2
				case 102:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 108:
					return -1
				case 110:
					return 3
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				case 111:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return 5
				case 108:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

		// \=
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 61:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 61:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \+\=
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 43:
					return 1
				case 61:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return -1
				case 61:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return -1
				case 61:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \-\=
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 45:
					return 1
				case 61:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 61:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 61:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \*\=
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 42:
					return 1
				case 61:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 42:
					return -1
				case 61:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 42:
					return -1
				case 61:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \/\=
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 47:
					return 1
				case 61:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 47:
					return -1
				case 61:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 47:
					return -1
				case 61:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \$
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 36:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 36:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \=\=
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 61:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 61:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 61:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \!\=
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 33:
					return 1
				case 61:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 61:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 61:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \>
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 62:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 62:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \<
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 60:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 60:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \>\=
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 61:
					return -1
				case 62:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 61:
					return 2
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 61:
					return -1
				case 62:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \<\=
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 60:
					return 1
				case 61:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 60:
					return -1
				case 61:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 60:
					return -1
				case 61:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// or
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 111:
					return 1
				case 114:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 111:
					return -1
				case 114:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 111:
					return -1
				case 114:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// and
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return 1
				case 100:
					return -1
				case 110:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 100:
					return -1
				case 110:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 100:
					return 3
				case 110:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 100:
					return -1
				case 110:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// not
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 110:
					return 1
				case 111:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 110:
					return -1
				case 111:
					return 2
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 110:
					return -1
				case 111:
					return -1
				case 116:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 110:
					return -1
				case 111:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// test
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 115:
					return -1
				case 116:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 2
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 115:
					return 3
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 115:
					return -1
				case 116:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// try
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 114:
					return -1
				case 116:
					return 1
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 114:
					return 2
				case 116:
					return -1
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 114:
					return -1
				case 116:
					return -1
				case 121:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 114:
					return -1
				case 116:
					return -1
				case 121:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// do
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 100:
					return 1
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 111:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 111:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// fail
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 102:
					return 1
				case 105:
					return -1
				case 108:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 2
				case 102:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 102:
					return -1
				case 105:
					return 3
				case 108:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 102:
					return -1
				case 105:
					return -1
				case 108:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 102:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// goto
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 103:
					return 1
				case 111:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 111:
					return 2
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 111:
					return -1
				case 116:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 111:
					return 4
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 103:
					return -1
				case 111:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// gopast
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 103:
					return 1
				case 111:
					return -1
				case 112:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 103:
					return -1
				case 111:
					return 2
				case 112:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 103:
					return -1
				case 111:
					return -1
				case 112:
					return 3
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 4
				case 103:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 103:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 115:
					return 5
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 103:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 115:
					return -1
				case 116:
					return 6
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 103:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// repeat
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 112:
					return -1
				case 114:
					return 1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return 2
				case 112:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 112:
					return 3
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return 4
				case 112:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 5
				case 101:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 116:
					return 6
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// loop
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 108:
					return 1
				case 111:
					return -1
				case 112:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 108:
					return -1
				case 111:
					return 2
				case 112:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 108:
					return -1
				case 111:
					return 3
				case 112:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 108:
					return -1
				case 111:
					return -1
				case 112:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 108:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// atleast
		{[]bool{false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return 1
				case 101:
					return -1
				case 108:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 115:
					return -1
				case 116:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 108:
					return 3
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return 4
				case 108:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 5
				case 101:
					return -1
				case 108:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 115:
					return 6
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 115:
					return -1
				case 116:
					return 7
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// insert
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return 1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return 2
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return 3
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 4
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return 5
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 6
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// \<\+
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 43:
					return -1
				case 60:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return 2
				case 60:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return -1
				case 60:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// attach
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return 1
				case 99:
					return -1
				case 104:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 104:
					return -1
				case 116:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 104:
					return -1
				case 116:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 4
				case 99:
					return -1
				case 104:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return 5
				case 104:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 104:
					return 6
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 104:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// \<\-
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 60:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return 2
				case 60:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 60:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// delete
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 100:
					return 1
				case 101:
					return -1
				case 108:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return 2
				case 108:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 108:
					return 3
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return 4
				case 108:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 116:
					return 5
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return 6
				case 108:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// hop
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 104:
					return 1
				case 111:
					return -1
				case 112:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 104:
					return -1
				case 111:
					return 2
				case 112:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 104:
					return -1
				case 111:
					return -1
				case 112:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 104:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// next
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return 1
				case 116:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 2
				case 110:
					return -1
				case 116:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 116:
					return -1
				case 120:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 116:
					return 4
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 116:
					return -1
				case 120:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// =>
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 61:
					return 1
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 61:
					return -1
				case 62:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 61:
					return -1
				case 62:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \[
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 91:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 91:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \]
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 93:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 93:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \-\>
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 45:
					return 1
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 62:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 62:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// setmark
		{[]bool{false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 114:
					return -1
				case 115:
					return 1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return 2
				case 107:
					return -1
				case 109:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 107:
					return -1
				case 109:
					return 4
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 5
				case 101:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 114:
					return 6
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 107:
					return 7
				case 109:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// tomark
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 116:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 111:
					return 2
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 107:
					return -1
				case 109:
					return 3
				case 111:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 4
				case 107:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return 5
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 107:
					return 6
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// atmark
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return 1
				case 107:
					return -1
				case 109:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 114:
					return -1
				case 116:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 107:
					return -1
				case 109:
					return 3
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 4
				case 107:
					return -1
				case 109:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 114:
					return 5
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 107:
					return 6
				case 109:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// tolimit
		{[]bool{false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 116:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 111:
					return 2
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 108:
					return 3
				case 109:
					return -1
				case 111:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return 4
				case 108:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return 5
				case 111:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return 6
				case 108:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 116:
					return 7
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// atlimit
		{[]bool{false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return 1
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 116:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return -1
				case 108:
					return 3
				case 109:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return 4
				case 108:
					return -1
				case 109:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return 5
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return 6
				case 108:
					return -1
				case 109:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 116:
					return 7
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// setlimit
		{[]bool{false, false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 115:
					return 1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 2
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 115:
					return -1
				case 116:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return -1
				case 108:
					return 4
				case 109:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return 5
				case 108:
					return -1
				case 109:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return 6
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return 7
				case 108:
					return -1
				case 109:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 115:
					return -1
				case 116:
					return 8
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// for
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 102:
					return 1
				case 111:
					return -1
				case 114:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 102:
					return -1
				case 111:
					return 2
				case 114:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 102:
					return -1
				case 111:
					return -1
				case 114:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 102:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// backwards
		{[]bool{false, false, false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return 1
				case 99:
					return -1
				case 100:
					return -1
				case 107:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 2
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 107:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return 3
				case 100:
					return -1
				case 107:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 107:
					return 4
				case 114:
					return -1
				case 115:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 107:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 119:
					return 5
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 6
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 107:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 107:
					return -1
				case 114:
					return 7
				case 115:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return 8
				case 107:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 107:
					return -1
				case 114:
					return -1
				case 115:
					return 9
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 107:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 119:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// reverse
		{[]bool{false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 114:
					return 1
				case 115:
					return -1
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 2
				case 114:
					return -1
				case 115:
					return -1
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 118:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 4
				case 114:
					return -1
				case 115:
					return -1
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 114:
					return 5
				case 115:
					return -1
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 114:
					return -1
				case 115:
					return 6
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 7
				case 114:
					return -1
				case 115:
					return -1
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 118:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// substring
		{[]bool{false, false, false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 98:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return 1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 98:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 98:
					return 3
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 98:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return 4
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 98:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 5
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 98:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return 6
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 98:
					return -1
				case 103:
					return -1
				case 105:
					return 7
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 98:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return 8
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 98:
					return -1
				case 103:
					return 9
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 98:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// among
		{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return 1
				case 103:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 103:
					return -1
				case 109:
					return 2
				case 110:
					return -1
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 103:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 103:
					return -1
				case 109:
					return -1
				case 110:
					return 4
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 103:
					return 5
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 103:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

		// set
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 115:
					return 1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 2
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 115:
					return -1
				case 116:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// unset
		{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return 2
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 115:
					return 3
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 4
				case 110:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 115:
					return -1
				case 116:
					return 5
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

		// non
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 110:
					return 1
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 110:
					return -1
				case 111:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 110:
					return 3
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 110:
					return -1
				case 111:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// true
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 114:
					return -1
				case 116:
					return 1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 114:
					return 2
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				case 117:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 4
				case 114:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// false
		{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 102:
					return 1
				case 108:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 2
				case 101:
					return -1
				case 102:
					return -1
				case 108:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 108:
					return 3
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 108:
					return -1
				case 115:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return 5
				case 102:
					return -1
				case 108:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 108:
					return -1
				case 115:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

		// backwardmode
		{[]bool{false, false, false, false, false, false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return 1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 2
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return 3
				case 100:
					return -1
				case 101:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 107:
					return 4
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 119:
					return 5
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 6
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return 7
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return 8
				case 101:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 107:
					return -1
				case 109:
					return 9
				case 111:
					return -1
				case 114:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 111:
					return 10
				case 114:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return 11
				case 101:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 12
				case 107:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 107:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 119:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// stringescapes.*[A-Za-z0-9!"#$%&'()*+,\-.\/:;<=>?@[\\\]^_`{|}~][A-Za-z0-9!"#$%&'()*+,\-.\/:;<=>?@[\\\]^_`{|}~]
		{[]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 34:
					return -1
				case 35:
					return -1
				case 36:
					return -1
				case 37:
					return -1
				case 38:
					return -1
				case 39:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 47:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				case 63:
					return -1
				case 64:
					return -1
				case 91:
					return -1
				case 92:
					return -1
				case 93:
					return -1
				case 94:
					return -1
				case 95:
					return -1
				case 96:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return 1
				case 116:
					return -1
				case 123:
					return -1
				case 124:
					return -1
				case 125:
					return -1
				case 126:
					return -1
				}
				switch {
				case 44 <= r && r <= 46:
					return -1
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 90:
					return -1
				case 97 <= r && r <= 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 34:
					return -1
				case 35:
					return -1
				case 36:
					return -1
				case 37:
					return -1
				case 38:
					return -1
				case 39:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 47:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				case 63:
					return -1
				case 64:
					return -1
				case 91:
					return -1
				case 92:
					return -1
				case 93:
					return -1
				case 94:
					return -1
				case 95:
					return -1
				case 96:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 2
				case 123:
					return -1
				case 124:
					return -1
				case 125:
					return -1
				case 126:
					return -1
				}
				switch {
				case 44 <= r && r <= 46:
					return -1
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 90:
					return -1
				case 97 <= r && r <= 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 34:
					return -1
				case 35:
					return -1
				case 36:
					return -1
				case 37:
					return -1
				case 38:
					return -1
				case 39:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 47:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				case 63:
					return -1
				case 64:
					return -1
				case 91:
					return -1
				case 92:
					return -1
				case 93:
					return -1
				case 94:
					return -1
				case 95:
					return -1
				case 96:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 112:
					return -1
				case 114:
					return 3
				case 115:
					return -1
				case 116:
					return -1
				case 123:
					return -1
				case 124:
					return -1
				case 125:
					return -1
				case 126:
					return -1
				}
				switch {
				case 44 <= r && r <= 46:
					return -1
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 90:
					return -1
				case 97 <= r && r <= 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 34:
					return -1
				case 35:
					return -1
				case 36:
					return -1
				case 37:
					return -1
				case 38:
					return -1
				case 39:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 47:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				case 63:
					return -1
				case 64:
					return -1
				case 91:
					return -1
				case 92:
					return -1
				case 93:
					return -1
				case 94:
					return -1
				case 95:
					return -1
				case 96:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 103:
					return -1
				case 105:
					return 4
				case 110:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 123:
					return -1
				case 124:
					return -1
				case 125:
					return -1
				case 126:
					return -1
				}
				switch {
				case 44 <= r && r <= 46:
					return -1
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 90:
					return -1
				case 97 <= r && r <= 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 34:
					return -1
				case 35:
					return -1
				case 36:
					return -1
				case 37:
					return -1
				case 38:
					return -1
				case 39:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 47:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				case 63:
					return -1
				case 64:
					return -1
				case 91:
					return -1
				case 92:
					return -1
				case 93:
					return -1
				case 94:
					return -1
				case 95:
					return -1
				case 96:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return 5
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 123:
					return -1
				case 124:
					return -1
				case 125:
					return -1
				case 126:
					return -1
				}
				switch {
				case 44 <= r && r <= 46:
					return -1
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 90:
					return -1
				case 97 <= r && r <= 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 34:
					return -1
				case 35:
					return -1
				case 36:
					return -1
				case 37:
					return -1
				case 38:
					return -1
				case 39:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 47:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				case 63:
					return -1
				case 64:
					return -1
				case 91:
					return -1
				case 92:
					return -1
				case 93:
					return -1
				case 94:
					return -1
				case 95:
					return -1
				case 96:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 103:
					return 6
				case 105:
					return -1
				case 110:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 123:
					return -1
				case 124:
					return -1
				case 125:
					return -1
				case 126:
					return -1
				}
				switch {
				case 44 <= r && r <= 46:
					return -1
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 90:
					return -1
				case 97 <= r && r <= 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 34:
					return -1
				case 35:
					return -1
				case 36:
					return -1
				case 37:
					return -1
				case 38:
					return -1
				case 39:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 47:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				case 63:
					return -1
				case 64:
					return -1
				case 91:
					return -1
				case 92:
					return -1
				case 93:
					return -1
				case 94:
					return -1
				case 95:
					return -1
				case 96:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return 7
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 123:
					return -1
				case 124:
					return -1
				case 125:
					return -1
				case 126:
					return -1
				}
				switch {
				case 44 <= r && r <= 46:
					return -1
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 90:
					return -1
				case 97 <= r && r <= 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 34:
					return -1
				case 35:
					return -1
				case 36:
					return -1
				case 37:
					return -1
				case 38:
					return -1
				case 39:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 47:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				case 63:
					return -1
				case 64:
					return -1
				case 91:
					return -1
				case 92:
					return -1
				case 93:
					return -1
				case 94:
					return -1
				case 95:
					return -1
				case 96:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return 8
				case 116:
					return -1
				case 123:
					return -1
				case 124:
					return -1
				case 125:
					return -1
				case 126:
					return -1
				}
				switch {
				case 44 <= r && r <= 46:
					return -1
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 90:
					return -1
				case 97 <= r && r <= 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 34:
					return -1
				case 35:
					return -1
				case 36:
					return -1
				case 37:
					return -1
				case 38:
					return -1
				case 39:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 47:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				case 63:
					return -1
				case 64:
					return -1
				case 91:
					return -1
				case 92:
					return -1
				case 93:
					return -1
				case 94:
					return -1
				case 95:
					return -1
				case 96:
					return -1
				case 97:
					return -1
				case 99:
					return 9
				case 101:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 123:
					return -1
				case 124:
					return -1
				case 125:
					return -1
				case 126:
					return -1
				}
				switch {
				case 44 <= r && r <= 46:
					return -1
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 90:
					return -1
				case 97 <= r && r <= 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 34:
					return -1
				case 35:
					return -1
				case 36:
					return -1
				case 37:
					return -1
				case 38:
					return -1
				case 39:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 47:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				case 63:
					return -1
				case 64:
					return -1
				case 91:
					return -1
				case 92:
					return -1
				case 93:
					return -1
				case 94:
					return -1
				case 95:
					return -1
				case 96:
					return -1
				case 97:
					return 10
				case 99:
					return -1
				case 101:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 123:
					return -1
				case 124:
					return -1
				case 125:
					return -1
				case 126:
					return -1
				}
				switch {
				case 44 <= r && r <= 46:
					return -1
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 90:
					return -1
				case 97 <= r && r <= 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 34:
					return -1
				case 35:
					return -1
				case 36:
					return -1
				case 37:
					return -1
				case 38:
					return -1
				case 39:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 47:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				case 63:
					return -1
				case 64:
					return -1
				case 91:
					return -1
				case 92:
					return -1
				case 93:
					return -1
				case 94:
					return -1
				case 95:
					return -1
				case 96:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 112:
					return 11
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 123:
					return -1
				case 124:
					return -1
				case 125:
					return -1
				case 126:
					return -1
				}
				switch {
				case 44 <= r && r <= 46:
					return -1
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 90:
					return -1
				case 97 <= r && r <= 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 34:
					return -1
				case 35:
					return -1
				case 36:
					return -1
				case 37:
					return -1
				case 38:
					return -1
				case 39:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 47:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				case 63:
					return -1
				case 64:
					return -1
				case 91:
					return -1
				case 92:
					return -1
				case 93:
					return -1
				case 94:
					return -1
				case 95:
					return -1
				case 96:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return 12
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 123:
					return -1
				case 124:
					return -1
				case 125:
					return -1
				case 126:
					return -1
				}
				switch {
				case 44 <= r && r <= 46:
					return -1
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 90:
					return -1
				case 97 <= r && r <= 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 34:
					return -1
				case 35:
					return -1
				case 36:
					return -1
				case 37:
					return -1
				case 38:
					return -1
				case 39:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 47:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				case 63:
					return -1
				case 64:
					return -1
				case 91:
					return -1
				case 92:
					return -1
				case 93:
					return -1
				case 94:
					return -1
				case 95:
					return -1
				case 96:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return 13
				case 116:
					return -1
				case 123:
					return -1
				case 124:
					return -1
				case 125:
					return -1
				case 126:
					return -1
				}
				switch {
				case 44 <= r && r <= 46:
					return -1
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 90:
					return -1
				case 97 <= r && r <= 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return 14
				case 34:
					return 14
				case 35:
					return 14
				case 36:
					return 14
				case 37:
					return 14
				case 38:
					return 14
				case 39:
					return 14
				case 40:
					return 14
				case 41:
					return 14
				case 42:
					return 14
				case 43:
					return 14
				case 47:
					return 14
				case 58:
					return 14
				case 59:
					return 14
				case 60:
					return 14
				case 61:
					return 14
				case 62:
					return 14
				case 63:
					return 14
				case 64:
					return 14
				case 91:
					return 14
				case 92:
					return 14
				case 93:
					return 14
				case 94:
					return 14
				case 95:
					return 14
				case 96:
					return 14
				case 97:
					return 14
				case 99:
					return 14
				case 101:
					return 14
				case 103:
					return 14
				case 105:
					return 14
				case 110:
					return 14
				case 112:
					return 14
				case 114:
					return 14
				case 115:
					return 14
				case 116:
					return 14
				case 123:
					return 14
				case 124:
					return 14
				case 125:
					return 14
				case 126:
					return 14
				}
				switch {
				case 44 <= r && r <= 46:
					return 14
				case 48 <= r && r <= 57:
					return 14
				case 65 <= r && r <= 90:
					return 14
				case 97 <= r && r <= 122:
					return 14
				}
				return 15
			},
			func(r rune) int {
				switch r {
				case 33:
					return 16
				case 34:
					return 16
				case 35:
					return 16
				case 36:
					return 16
				case 37:
					return 16
				case 38:
					return 16
				case 39:
					return 16
				case 40:
					return 16
				case 41:
					return 16
				case 42:
					return 16
				case 43:
					return 16
				case 47:
					return 16
				case 58:
					return 16
				case 59:
					return 16
				case 60:
					return 16
				case 61:
					return 16
				case 62:
					return 16
				case 63:
					return 16
				case 64:
					return 16
				case 91:
					return 16
				case 92:
					return 16
				case 93:
					return 16
				case 94:
					return 16
				case 95:
					return 16
				case 96:
					return 16
				case 97:
					return 16
				case 99:
					return 16
				case 101:
					return 16
				case 103:
					return 16
				case 105:
					return 16
				case 110:
					return 16
				case 112:
					return 16
				case 114:
					return 16
				case 115:
					return 16
				case 116:
					return 16
				case 123:
					return 16
				case 124:
					return 16
				case 125:
					return 16
				case 126:
					return 16
				}
				switch {
				case 44 <= r && r <= 46:
					return 16
				case 48 <= r && r <= 57:
					return 16
				case 65 <= r && r <= 90:
					return 16
				case 97 <= r && r <= 122:
					return 16
				}
				return 15
			},
			func(r rune) int {
				switch r {
				case 33:
					return 14
				case 34:
					return 14
				case 35:
					return 14
				case 36:
					return 14
				case 37:
					return 14
				case 38:
					return 14
				case 39:
					return 14
				case 40:
					return 14
				case 41:
					return 14
				case 42:
					return 14
				case 43:
					return 14
				case 47:
					return 14
				case 58:
					return 14
				case 59:
					return 14
				case 60:
					return 14
				case 61:
					return 14
				case 62:
					return 14
				case 63:
					return 14
				case 64:
					return 14
				case 91:
					return 14
				case 92:
					return 14
				case 93:
					return 14
				case 94:
					return 14
				case 95:
					return 14
				case 96:
					return 14
				case 97:
					return 14
				case 99:
					return 14
				case 101:
					return 14
				case 103:
					return 14
				case 105:
					return 14
				case 110:
					return 14
				case 112:
					return 14
				case 114:
					return 14
				case 115:
					return 14
				case 116:
					return 14
				case 123:
					return 14
				case 124:
					return 14
				case 125:
					return 14
				case 126:
					return 14
				}
				switch {
				case 44 <= r && r <= 46:
					return 14
				case 48 <= r && r <= 57:
					return 14
				case 65 <= r && r <= 90:
					return 14
				case 97 <= r && r <= 122:
					return 14
				}
				return 15
			},
			func(r rune) int {
				switch r {
				case 33:
					return 16
				case 34:
					return 16
				case 35:
					return 16
				case 36:
					return 16
				case 37:
					return 16
				case 38:
					return 16
				case 39:
					return 16
				case 40:
					return 16
				case 41:
					return 16
				case 42:
					return 16
				case 43:
					return 16
				case 47:
					return 16
				case 58:
					return 16
				case 59:
					return 16
				case 60:
					return 16
				case 61:
					return 16
				case 62:
					return 16
				case 63:
					return 16
				case 64:
					return 16
				case 91:
					return 16
				case 92:
					return 16
				case 93:
					return 16
				case 94:
					return 16
				case 95:
					return 16
				case 96:
					return 16
				case 97:
					return 16
				case 99:
					return 16
				case 101:
					return 16
				case 103:
					return 16
				case 105:
					return 16
				case 110:
					return 16
				case 112:
					return 16
				case 114:
					return 16
				case 115:
					return 16
				case 116:
					return 16
				case 123:
					return 16
				case 124:
					return 16
				case 125:
					return 16
				case 126:
					return 16
				}
				switch {
				case 44 <= r && r <= 46:
					return 16
				case 48 <= r && r <= 57:
					return 16
				case 65 <= r && r <= 90:
					return 16
				case 97 <= r && r <= 122:
					return 16
				}
				return 15
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// \?
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 63:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 63:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// '.*'
		{[]bool{false, false, true, false}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 39:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 39:
					return 2
				}
				return 3
			},
			func(r rune) int {
				switch r {
				case 39:
					return 2
				}
				return 3
			},
			func(r rune) int {
				switch r {
				case 39:
					return 2
				}
				return 3
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// [0-9]+
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch {
				case 48 <= r && r <= 57:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch {
				case 48 <= r && r <= 57:
					return 1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// [a-zA-Z][a-zA-Z0-9_]*
		{[]bool{false, true, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 95:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 90:
					return 1
				case 97 <= r && r <= 122:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 95:
					return 2
				}
				switch {
				case 48 <= r && r <= 57:
					return 2
				case 65 <= r && r <= 90:
					return 2
				case 97 <= r && r <= 122:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 95:
					return 2
				}
				switch {
				case 48 <= r && r <= 57:
					return 2
				case 65 <= r && r <= 90:
					return 2
				case 97 <= r && r <= 122:
					return 2
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// .
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				return 1
			},
			func(r rune) int {
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},
	}, 0, 0)
	return yylex
}

func newLexer(in io.Reader) *lexer {
	return newLexerWithInit(in, nil)
}

// Text returns the matched text.
func (yylex *lexer) Text() string {
	return yylex.stack[len(yylex.stack)-1].s
}

// Line returns the current line number.
// The first line is 0.
func (yylex *lexer) Line() int {
	if len(yylex.stack) == 0 {
		return 0
	}
	return yylex.stack[len(yylex.stack)-1].line
}

// Column returns the current column number.
// The first column is 0.
func (yylex *lexer) Column() int {
	if len(yylex.stack) == 0 {
		return 0
	}
	return yylex.stack[len(yylex.stack)-1].column
}

func (yylex *lexer) next(lvl int) int {
	if lvl == len(yylex.stack) {
		l, c := 0, 0
		if lvl > 0 {
			l, c = yylex.stack[lvl-1].line, yylex.stack[lvl-1].column
		}
		yylex.stack = append(yylex.stack, frame{0, "", l, c})
	}
	if lvl == len(yylex.stack)-1 {
		p := &yylex.stack[lvl]
		*p = <-yylex.ch
		yylex.stale = false
	} else {
		yylex.stale = true
	}
	return yylex.stack[lvl].i
}
func (yylex *lexer) pop() {
	yylex.stack = yylex.stack[:len(yylex.stack)-1]
}
func (yylex lexer) Error(e string) {
	panic(e)
}

// Lex runs the lexer. Always returns 0.
// When the -s option is given, this function is not generated;
// instead, the NN_FUN macro runs the lexer.
func (yylex *lexer) Lex(lval *yySymType) int {
OUTER0:
	for {
		switch yylex.next(0) {
		case 0:
			{
				logDebugTokens("strings")
				return tSTRINGS
			}
		case 1:
			{
				logDebugTokens("integers")
				return tINTEGERS
			}
		case 2:
			{
				logDebugTokens("booleans")
				return tBOOLEANS
			}
		case 3:
			{
				logDebugTokens("routines")
				return tROUTINES
			}
		case 4:
			{
				logDebugTokens("externals")
				return tEXTERNALS
			}
		case 5:
			{
				logDebugTokens("groupings")
				return tGROUPINGS
			}
		case 6:
			{
				logDebugTokens("define")
				return tDEFINE
			}
		case 7:
			{
				logDebugTokens("as")
				return tAS
			}
		case 8:
			{
				logDebugTokens("(")
				return tLPAREN
			}
		case 9:
			{
				logDebugTokens(")")
				return tRPAREN
			}
		case 10:
			{
				logDebugTokens("+")
				return tPLUS
			}
		case 11:
			{
				logDebugTokens("-")
				return tMINUS
			}
		case 12:
			{
				logDebugTokens("*")
				return tMULT
			}
		case 13:
			{
				logDebugTokens("/")
				return tDIV
			}
		case 14:
			{
				logDebugTokens("maxint")
				return tMAXINT
			}
		case 15:
			{
				logDebugTokens("minint")
				return tMININT
			}
		case 16:
			{
				logDebugTokens("cursor")
				return tCURSOR
			}
		case 17:
			{
				logDebugTokens("limit")
				return tLIMIT
			}
		case 18:
			{
				logDebugTokens("size")
				return tSIZE
			}
		case 19:
			{
				logDebugTokens("sizeof")
				return tSIZEOF
			}
		case 20:
			{
				logDebugTokens("len")
				return tLEN
			}
		case 21:
			{
				logDebugTokens("lenof")
				return tLENOF
			}
		case 22:
			{
				logDebugTokens("=")
				return tASSIGN
			}
		case 23:
			{
				logDebugTokens("+=")
				return tPLUSASSIGN
			}
		case 24:
			{
				logDebugTokens("-=")
				return tMINUSASSIGN
			}
		case 25:
			{
				logDebugTokens("*=")
				return tMULTASSIGN
			}
		case 26:
			{
				logDebugTokens("/=")
				return tDIVASSIGN
			}
		case 27:
			{
				logDebugTokens("$")
				return tDOLLAR
			}
		case 28:
			{
				logDebugTokens("==")
				return tEQ
			}
		case 29:
			{
				logDebugTokens("!=")
				return tNEQ
			}
		case 30:
			{
				logDebugTokens(">")
				return tGT
			}
		case 31:
			{
				logDebugTokens("<")
				return tLT
			}
		case 32:
			{
				logDebugTokens(">=")
				return tGTEQ
			}
		case 33:
			{
				logDebugTokens("<=")
				return tLTEQ
			}
		case 34:
			{
				logDebugTokens("or")
				return tOR
			}
		case 35:
			{
				logDebugTokens("and")
				return tAND
			}
		case 36:
			{
				logDebugTokens("not")
				return tNOT
			}
		case 37:
			{
				logDebugTokens("test")
				return tTEST
			}
		case 38:
			{
				logDebugTokens("try")
				return tTRY
			}
		case 39:
			{
				logDebugTokens("do")
				return tDO
			}
		case 40:
			{
				logDebugTokens("fail")
				return tFAIL
			}
		case 41:
			{
				logDebugTokens("goto")
				return tGOTO
			}
		case 42:
			{
				logDebugTokens("gopast")
				return tGOPAST
			}
		case 43:
			{
				logDebugTokens("repeat")
				return tREPEAT
			}
		case 44:
			{
				logDebugTokens("loop")
				return tLOOP
			}
		case 45:
			{
				logDebugTokens("atleast")
				return tATLEAST
			}
		case 46:
			{
				logDebugTokens("insert")
				return tINSERT
			}
		case 47:
			{
				logDebugTokens("insert")
				return tINSERT
			}
		case 48:
			{
				logDebugTokens("attach")
				return tATTACH
			}
		case 49:
			{
				logDebugTokens("<-")
				return tREPLACE
			}
		case 50:
			{
				logDebugTokens("delete")
				return tDELETE
			}
		case 51:
			{
				logDebugTokens("hop")
				return tHOP
			}
		case 52:
			{
				logDebugTokens("next")
				return tNEXT
			}
		case 53:
			{
				logDebugTokens("=>")
				return tASSIGNR
			}
		case 54:
			{
				logDebugTokens("[")
				return tLBRACKET
			}
		case 55:
			{
				logDebugTokens("]")
				return tRBRACKET
			}
		case 56:
			{
				logDebugTokens("->")
				return tMOVER
			}
		case 57:
			{
				logDebugTokens("setmark")
				return tSETMARK
			}
		case 58:
			{
				logDebugTokens("tomark")
				return tTOMARK
			}
		case 59:
			{
				logDebugTokens("atmark")
				return tATMARK
			}
		case 60:
			{
				logDebugTokens("tolimit")
				return tTOLIMIT
			}
		case 61:
			{
				logDebugTokens("atlimit")
				return tATLIMIT
			}
		case 62:
			{
				logDebugTokens("setlimit")
				return tSETLIMIT
			}
		case 63:
			{
				logDebugTokens("for")
				return tFOR
			}
		case 64:
			{
				logDebugTokens("backwards")
				return tBACKWARDS
			}
		case 65:
			{
				logDebugTokens("reverse")
				return tREVERSE
			}
		case 66:
			{
				logDebugTokens("substring")
				return tSUBSTRING
			}
		case 67:
			{
				logDebugTokens("among")
				return tAMONG
			}
		case 68:
			{
				logDebugTokens("set")
				return tSET
			}
		case 69:
			{
				logDebugTokens("unset")
				return tUNSET
			}
		case 70:
			{
				logDebugTokens("non")
				return tNON
			}
		case 71:
			{
				logDebugTokens("true")
				return tTRUE
			}
		case 72:
			{
				logDebugTokens("false")
				return tFALSE
			}
		case 73:
			{
				logDebugTokens("backwardmode")
				return tBACKWARDMODE
			}
		case 74:
			{
				lval.s = yylex.Text()[len(yylex.Text())-2:]
				logDebugTokens("stringescapes")
				return tSTRINGESCAPES
			}
		case 75:
			{
				logDebugTokens("?")
				return tQUESTION
			}
		case 76:
			{
				lval.s = yylex.Text()[1 : len(yylex.Text())-1]
				logDebugTokens("literal - %s", lval.s)
				return tLITERAL
			}
		case 77:
			{
				lval.n, _ = strconv.Atoi(yylex.Text())
				logDebugTokens("number - %d", lval.n)
				return tNUMBER
			}
		case 78:
			{
				lval.s = yylex.Text()
				logDebugTokens("name - %s", lval.s)
				return tNAME
			}
		case 79:
			{
				logDebugTokens("other")
			}
		default:
			break OUTER0
		}
		continue
	}
	yylex.pop()

	return 0
}
func logDebugTokens(format string, v ...interface{}) {
	if DebugLexer {
		Logger.Printf(fmt.Sprintf("LEXER %s", format), v...)
	}
}
