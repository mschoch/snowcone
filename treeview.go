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
	"bytes"
	"fmt"
	"io"
	"os"
)

type treeview struct {
	indent int
	w      io.Writer
}

func (t *treeview) printIndent() {
	for i := 0; i < t.indent; i++ {
		fmt.Fprintf(t.w, "  ")
	}
}

func (t *treeview) Visit(n node) {
	t.printIndent()
	fmt.Fprintf(t.w, "- %v\n", n)
	t.indent++
	switch n := n.(type) {
	case *prog:
		for _, d := range n.decls {
			d.Accept(t)
		}
		for _, g := range n.groupdefs {
			g.Accept(t)
		}
		for _, r := range n.routinedefs {
			r.Accept(t)
		}
	case *routine:
		n.comm.Accept(t)
	case *grouping:
		n.children.Accept(t)
	case *among:
		n.children.Accept(t)
	case *unaryCommand:
		if n.operandCommand != nil {
			n.operandCommand.Accept(t)
		} else if n.operandAe != nil {
			n.operandAe.Accept(t)
		} else if n.operandName != nil {
			n.operandName.Accept(t)
		}
	case *binaryCommand:
		n.left.Accept(t)
		n.right.Accept(t)
	case *loop:
		n.operand.Accept(t)
	case *sCommand:
		n.operand.Accept(t)
	case *iCommand:
		n.operand.Accept(t)
	case commands:
		for _, ci := range n {
			ci.Accept(t)
		}
	}
	t.indent--
}

func PrintTreeView(p *prog) {
	t := &treeview{
		w: os.Stdout,
	}
	p.Accept(t)
}

func PrintTreeViewString(p *prog) string {
	b := bytes.Buffer{}
	t := &treeview{
		w: &b,
	}
	p.Accept(t)
	return b.String()
}
