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

import "testing"

func TestStringdefs(t *testing.T) {

	sd := &stringdefs{}

	// before setting the string escapes, always get back what you pass in
	tests := []struct {
		in  string
		out string
	}{
		{
			in:  `something`,
			out: `something`,
		},
		{
			in:  `something{'}else`,
			out: `something{'}else`,
		},
	}

	for _, test := range tests {
		actual := sd.ReplaceInLiteral(test.in)
		if actual != test.out {
			t.Errorf("expected '%s', got '%s', for '%s'", test.out, actual, test.in)
		}
	}

	// as soon as we set string escapes, three are automatically defined
	// the first is the ' character,
	// the second is the start escape rune,
	// the third is whitespace only macros
	sd.SetStringEscapes('{', '}')
	tests = []struct {
		in  string
		out string
	}{
		{
			in:  `something{'}else`,
			out: `something'else`,
		},
		{
			in:  `something{{}else`,
			out: `something{else`,
		},
		// contains tab, newline and spaces inside macro
		{
			in: `something{
           }else`,
			out: `somethingelse`,
		},
		// unterminated escape
		{
			in:  `something{`,
			out: `something{`,
		},
		// empty escape
		{
			in:  `something{}`,
			out: `something{}`,
		},
	}

	for _, test := range tests {
		actual := sd.ReplaceInLiteral(test.in)
		if actual != test.out {
			t.Errorf("expected '%s', got '%s', for '%s'", test.out, actual, test.in)
		}
	}

	// additional expansions can then be defined
	sd.Define("a'", "á")
	tests = []struct {
		in  string
		out string
	}{
		{
			in:  `someth{a'}ng`,
			out: `sometháng`,
		},
	}

	for _, test := range tests {
		actual := sd.ReplaceInLiteral(test.in)
		if actual != test.out {
			t.Errorf("expected '%s', got '%s', for '%s'", test.out, actual, test.in)
		}
	}
}
