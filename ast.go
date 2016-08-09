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

import "fmt"

type node interface {
	Accept(visitor)
}

type visitor interface {
	Visit(node)
}

// Type represents the snowball type of an name
type stype int

func (s stype) String() string {
	switch s {
	case sstring:
		return "string"
	case sinteger:
		return "integer"
	case sboolean:
		return "boolean"
	case sroutine:
		return "routine"
	case sexternal:
		return "external"
	case sgrouping:
		return "grouping"
	}
	return "unknown"
}

// Snowball types
const (
	sstring stype = iota
	sinteger
	sboolean
	sroutine
	sexternal
	sgrouping
)

type prog struct {
	decls       decls
	groupdefs   []*grouping
	routinedefs []*routine
}

func (p *prog) Accept(v visitor) {
	v.Visit(p)
}

func (p *prog) String() string {
	return "prog"
}

func (p *prog) Declare(d *decl) {
	p.decls = append(p.decls, d)
}

func (p *prog) DefineGroup(g *grouping) {
	p.groupdefs = append(p.groupdefs, g)
}

func (p *prog) DefineRoutine(r *routine) {
	p.routinedefs = append(p.routinedefs, r)
}

func (p *prog) SetBackwardMode() {
	for _, r := range p.routinedefs {
		r.backwardmode = true
	}
}

func (p *prog) Combine(other *prog) {
	for _, d := range other.decls {
		p.Declare(d)
	}
	for _, g := range other.groupdefs {
		p.DefineGroup(g)
	}
	for _, r := range other.routinedefs {
		p.DefineRoutine(r)
	}
}

type decls []*decl

type decl struct {
	name string
	typ  stype
}

func (d *decl) Accept(v visitor) {
	v.Visit(d)
}

func (d *decl) String() string {
	return fmt.Sprintf("declare %s %s", d.typ, d.name)
}

type routine struct {
	name         string
	comm         node
	backwardmode bool
}

func (r *routine) Accept(v visitor) {
	v.Visit(r)
}

func (r *routine) String() string {
	if r.backwardmode {
		return fmt.Sprintf("define backwardmode routine %s", r.name)
	}
	return fmt.Sprintf("define routine %s ", r.name)
}

type groupitems []*groupitem

func (g groupitems) String() string {
	rv := ""
	for i, gi := range g {
		if i != 0 {
			if gi.minus {
				rv += " - "
			} else {
				rv += " + "
			}
		}
		rv += fmt.Sprintf("%v", gi.item)
	}
	return rv
}

func (g groupitems) Accept(v visitor) {
	v.Visit(g)
}

type groupitem struct {
	minus bool
	item  node
}

type grouping struct {
	name     string
	children groupitems
}

func (g *grouping) Accept(v visitor) {
	v.Visit(g)
}

func (g *grouping) String() string {
	return fmt.Sprintf("define grouping %s", g.name)
}

type sliteral struct {
	val string
}

func (s *sliteral) Accept(v visitor) {
	v.Visit(s)
}

func (s *sliteral) String() string {
	return fmt.Sprintf("'%s'", s.val)
}

type nliteral struct {
	val int
}

func (n *nliteral) Accept(v visitor) {
	v.Visit(n)
}

func (n *nliteral) String() string {
	return fmt.Sprintf("%d", n.val)
}

type bliteral struct {
	val bool
}

func (b *bliteral) Accept(v visitor) {
	v.Visit(b)
}

func (b *bliteral) String() string {
	return fmt.Sprintf("%t", b.val)
}

type name struct {
	val string
}

func (n *name) Accept(v visitor) {
	v.Visit(n)
}

func (n *name) String() string {
	return fmt.Sprintf("%s", n.val)
}

type commands []command

func (c commands) String() string {
	return "command sequence"
}

func (c commands) Accept(v visitor) {
	v.Visit(c)
}

type command interface {
	node
}

type non struct {
	minus bool
	gname string
}

func (n *non) Accept(v visitor) {
	v.Visit(n)
}

func (n *non) String() string {
	if n.minus {
		return fmt.Sprintf("non - %s", n.gname)
	}
	return fmt.Sprintf("non %s", n.gname)
}

type amongitems []*amongitem

func (a amongitems) String() string {
	rv := ""
	for _, ai := range a {
		rv += fmt.Sprintf("%v", ai)
	}
	return rv
}

func (a amongitems) Accept(v visitor) {
	v.Visit(a)
}

type amongitem struct {
	slit  *sliteral
	rname string
	comm  command
}

func (a *amongitem) String() string {
	rv := ""
	if a.slit != nil {
		rv += fmt.Sprintf("%v", a.slit)
	}
	if a.rname != "" {
		rv += fmt.Sprintf("(call routine %s)", a.rname)
	}
	if a.comm != nil {
		rv += fmt.Sprintf("command %v", a.comm)
	}
	return rv
}

type among struct {
	children amongitems
}

func (a *among) Accept(v visitor) {
	v.Visit(a)
}

func (a *among) String() string {
	return fmt.Sprintf("among")
}

type unaryCommand struct {
	command        string
	operandCommand command
	operandName    *name
	operandAe      ae
}

func (u *unaryCommand) Accept(v visitor) {
	v.Visit(u)
}

func (u *unaryCommand) String() string {
	return fmt.Sprintf("%s", u.command)
}

type iCommand struct {
	name     *name
	operator string
	operand  ae
}

func (i *iCommand) Accept(v visitor) {
	v.Visit(i)
}

func (i *iCommand) String() string {
	return fmt.Sprintf("$%s %v", i.name, i.operator)
}

type sCommand struct {
	name    *name
	operand command
}

func (s *sCommand) Accept(v visitor) {
	v.Visit(s)
}

func (s *sCommand) String() string {
	return fmt.Sprintf("$%s", s.name)
}

type loop struct {
	n       ae
	operand command
	extra   bool
}

func (l *loop) Accept(v visitor) {
	v.Visit(l)
}

func (l *loop) String() string {
	if !l.extra {
		return fmt.Sprintf("loop %v", l.n)
	}
	return fmt.Sprintf("atleast %v", l.n)
}

type nilaryCommand struct {
	operator string
}

func (n *nilaryCommand) Accept(v visitor) {
	v.Visit(n)
}

func (n *nilaryCommand) String() string {
	return fmt.Sprintf("%s", n.operator)
}

type binaryCommand struct {
	left     command
	operator string
	right    command
}

func (b *binaryCommand) Accept(v visitor) {
	v.Visit(b)
}

func (b *binaryCommand) String() string {
	return fmt.Sprintf("%s", b.operator)
}

// *** ae ***

type ae interface {
	node
}

type binaryAe struct {
	operator string
	left     ae
	right    ae
}

func (b *binaryAe) Accept(v visitor) {
	v.Visit(b)
}

func (b *binaryAe) String() string {
	return fmt.Sprintf("(%v %s %v)", b.left, b.operator, b.right)
}

type unaryAe struct {
	operator string
	operand  ae
}

func (u *unaryAe) Accept(v visitor) {
	v.Visit(u)
}

func (u *unaryAe) String() string {
	return fmt.Sprintf("%s %v", u.operator, u.operand)
}

type nilaryAe struct {
	operator string
}

func (n *nilaryAe) Accept(v visitor) {
	v.Visit(n)
}

func (n *nilaryAe) String() string {
	return fmt.Sprintf("%s", n.operator)
}
