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
			in: `define shortv as ( non-v_WXY v non-v )`,
		},
		{
			in: `backwardmode (

			    define shortv as ( non-v_WXY v non-v )

			    define shortv as ( non-v_WXY v non-v )
			    )`,
		},
		{
			in: `    define Step_1a as (
		        [substring] among (
		            'sses' (<-'ss')
		            'ies'  (<-'i')
		            'ss'   ()
		            's'    (delete)
		        )
		    )`,
		},
		{
			in: `    define Step_1b as (
						[substring] among (
								'eed'  (R1 R2)
						)
				)`,
		},
		{
			in: `    define Step_1b as (
		        [substring] among (
		            'eed'  (R1 <-'ee')
		        )
		    )`,
		},
		{
			in: `    define Step_1c as (
		        ['y' or 'Y']
		        gopast v
		        <-'i'
		    )`,
		},
		{
			in: porterStemmerSrc,
		},
		{
			in: `stringescapes {}
			define v 'aeiou{a'}{e'}{i'}{o'}{u'}{u"}'`,
		},
	}

	for _, test := range tests {
		_, err := Parse(strings.NewReader(test.in))
		if err != nil {
			t.Errorf("parse error: %v", err)
		}
	}
}

func TestParserFail(t *testing.T) {
	// DebugLexer = true
	// DebugParser = true
	// Logger = log.New(os.Stdout, "", log.LstdFlags)
	tests := []struct {
		in string
	}{
		{
			// incomplete definition
			in: `define`,
		},
		{
			// did not call string escapes to define escape characters
			in: `define v 'aeiou{a'}{e'}{i'}{o'}{u'}{u"}'`,
		},
	}

	for _, test := range tests {
		_, err := Parse(strings.NewReader(test.in))
		if err == nil {
			t.Errorf("expected error got nil for %s", test.in)
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
