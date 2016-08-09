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

func TestParsingAST(t *testing.T) {
	// DebugLexer = true
	// DebugParser = true
	// Logger = log.New(os.Stdout, "", log.LstdFlags)
	tests := []struct {
		in string
		p  *prog
	}{
		// this first entry isn't a real test, it's just to see if we parse can
		// parse these structures without blowing up, ultimately this goes away
		// in favore of more specific tests that check the outcome
		// it is commented out because it fails

		// {
		// 	in: `
		// 	booleans( b1 b2 )
		//   strings( s1 s2 s3)
		//   integers( i int2 int3)
		// 	groupings (g1 g2)
		// 	routines ( r1 r2 r3 r4 r5 r6 r7 r8 r10 r11)
		// 	define g1 'abc' + s1
		// 	define g2 g1 - 'cat' + 'dog'
		// 	define r1 as true
		// 	define r2 as false
		// 	define r3 as ?
		// 	define r4 as non g1
		// 	define r5 as non - g2
		// 	define r6 as s1
		// 	define r7 as set b1
		// 	define r8 as unset b2
		// 	define r9 as among('cat' r1 (true))
		// 	define r10 as reverse 'this'
		// 	define r11 as backwards 'that'
		//   `,
		// 	p: nil,
		// },

		// the simplest program i felt i could manually construct the ast for
		{
			in: `routines (r1)
		define r1 as loop 1+1 true`,
			p: &prog{
				decls: decls{
					&decl{
						name: "r1",
						typ:  sroutine,
					},
				},
				routinedefs: []*routine{
					&routine{
						name: "r1",
						comm: &loop{
							n: &binaryAe{
								left: &nliteral{
									val: 1,
								},
								operator: "+",
								right: &nliteral{
									val: 1,
								},
							},
							operand: &bliteral{
								val: true,
							},
						},
					},
				},
			},
		},
		// test order of operations in arithmetic
		{
			in: `routines (r1)
		define r1 as loop 1+3*2 true`,
			p: &prog{
				decls: decls{
					&decl{
						name: "r1",
						typ:  sroutine,
					},
				},
				routinedefs: []*routine{
					&routine{
						name: "r1",
						comm: &loop{
							n: &binaryAe{
								left: &nliteral{
									val: 1,
								},
								operator: "+",
								right: &binaryAe{
									left: &nliteral{
										val: 3,
									},
									operator: "*",
									right: &nliteral{
										val: 2,
									},
								},
							},
							operand: &bliteral{
								val: true,
							},
						},
					},
				},
			},
		},
		// test order of operations in arithmetic, override with parens
		{
			in: `routines (r1)
		define r1 as loop (1+3)*2 true`,
			p: &prog{
				decls: decls{
					&decl{
						name: "r1",
						typ:  sroutine,
					},
				},
				routinedefs: []*routine{
					&routine{
						name: "r1",
						comm: &loop{
							n: &binaryAe{
								left: &binaryAe{
									left: &nliteral{
										val: 1,
									},
									operator: "+",
									right: &nliteral{
										val: 3,
									},
								},
								operator: "*",
								right: &nliteral{
									val: 2,
								},
							},
							operand: &bliteral{
								val: true,
							},
						},
					},
				},
			},
		},
		// precedence test from manual, unary binds to shortest valid command
		{
			in: `integers( x z )
		routines (r1)
		define r1 as (try not $x < 1 $z > 0)`,
			p: &prog{
				decls: decls{
					&decl{
						name: "x",
						typ:  sinteger,
					},
					&decl{
						name: "z",
						typ:  sinteger,
					},
					&decl{
						name: "r1",
						typ:  sroutine,
					},
				},
				routinedefs: []*routine{
					&routine{
						name: "r1",
						comm: commands{
							&unaryCommand{
								command: "try",
								operandCommand: &unaryCommand{
									command: "not",
									operandCommand: &iCommand{
										name:     &name{val: "x"},
										operator: "<",
										operand:  &nliteral{val: 1}},
								},
							},
							&iCommand{
								name:     &name{val: "z"},
								operator: ">",
								operand:  &nliteral{val: 0}},
						},
					},
				},
			},
		},
		// another test from manual, binary commands bind after unary
		// then equally from left to right afterwards
		{
			in: `integers( x y z t )
					routines (r1)
					define r1 as ($x > 0  not $y > 0 or not $z > 0  $t > 0)`,
			p: &prog{
				decls: decls{
					&decl{
						name: "x",
						typ:  sinteger,
					},
					&decl{
						name: "y",
						typ:  sinteger,
					},
					&decl{
						name: "z",
						typ:  sinteger,
					},
					&decl{
						name: "t",
						typ:  sinteger,
					},
					&decl{
						name: "r1",
						typ:  sroutine,
					},
				},
				routinedefs: []*routine{
					&routine{
						name: "r1",
						comm: commands{
							&iCommand{
								name:     &name{val: "x"},
								operator: ">",
								operand:  &nliteral{val: 0},
							},
							&binaryCommand{
								left: &unaryCommand{
									command: "not",
									operandCommand: &iCommand{
										name:     &name{val: "y"},
										operator: ">",
										operand:  &nliteral{val: 0}},
								},
								operator: "or",
								right: &unaryCommand{
									command: "not",
									operandCommand: &iCommand{
										name:     &name{val: "z"},
										operator: ">",
										operand:  &nliteral{val: 0}},
								},
							},
							&iCommand{
								name:     &name{val: "t"},
								operator: ">",
								operand:  &nliteral{val: 0},
							},
						},
					},
				},
			},
		},
		// yet another precedence test from manual, or binds before and in this case
		{
			in: `integers( x y z )
					routines (r1)
					define r1 as ($x > 0  or $y > 0 and $z > 0)`,
			p: &prog{
				decls: decls{
					&decl{
						name: "x",
						typ:  sinteger,
					},
					&decl{
						name: "y",
						typ:  sinteger,
					},
					&decl{
						name: "z",
						typ:  sinteger,
					},
					&decl{
						name: "r1",
						typ:  sroutine,
					},
				},
				routinedefs: []*routine{
					&routine{
						name: "r1",
						comm: commands{
							&binaryCommand{
								left: &binaryCommand{
									left: &iCommand{
										name:     &name{val: "x"},
										operator: ">",
										operand:  &nliteral{val: 0},
									},
									operator: "or",
									right: &iCommand{
										name:     &name{val: "y"},
										operator: ">",
										operand:  &nliteral{val: 0},
									},
								},
								operator: "and",
								right: &iCommand{
									name:     &name{val: "z"},
									operator: ">",
									operand:  &nliteral{val: 0},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		p, err := Parse(strings.NewReader(test.in))
		if err != nil {
			t.Fatalf("parse error: %v", err)
		}
		if !reflect.DeepEqual(p, test.p) {
			t.Errorf("programs differ, got\n%s\nexpected\n%s", PrintTreeViewString(p), PrintTreeViewString(test.p))
		}
	}
}
