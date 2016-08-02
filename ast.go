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
	rv := "prog\n"
	for _, d := range p.decls {
		rv += fmt.Sprintf("%v\n", d)
	}
	for _, g := range p.groupdefs {
		rv += fmt.Sprintf("%v\n", g)
	}
	for _, r := range p.routinedefs {
		rv += fmt.Sprintf("%v\n", r)
	}
	return rv
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
	return fmt.Sprintf("declare %s as %s", d.name, d.typ)
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
	return fmt.Sprintf("define routine %s as %v", r.name, r.comm)
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
	return fmt.Sprintf("define grouping %s as %v", g.name, g.children)
}

type sliteral struct {
	val string
}

func (s *sliteral) Accept(v visitor) {
	v.Visit(s)
}

func (s *sliteral) String() string {
	return fmt.Sprintf("literal string '%s'", s.val)
}

type nliteral struct {
	val int
}

func (n *nliteral) Accept(v visitor) {
	v.Visit(n)
}

func (n *nliteral) String() string {
	return fmt.Sprintf("literal number '%d'", n.val)
}

type name struct {
	val string
}

func (n *name) Accept(v visitor) {
	v.Visit(n)
}

func (n *name) String() string {
	return fmt.Sprintf("name '%s'", n.val)
}

type command interface {
	node
}

type bliteral struct {
	val bool
}

func (b *bliteral) Accept(v visitor) {
	v.Visit(b)
}

func (b *bliteral) String() string {
	return fmt.Sprintf("boolean literal '%t'", b.val)
}

type question struct{}

func (q *question) Accept(v visitor) {
	v.Visit(q)
}

func (q *question) String() string {
	return fmt.Sprintf("question")
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

type set struct {
	bname string
}

func (s *set) Accept(v visitor) {
	v.Visit(s)
}

func (s *set) String() string {
	return fmt.Sprintf("set %s", s.bname)
}

type unset struct {
	bname string
}

func (u *unset) Accept(v visitor) {
	v.Visit(u)
}

func (u *unset) String() string {
	return fmt.Sprintf("unset %s", u.bname)
}

type amongitems []*amongitem

func (a amongitems) String() string {
	rv := ""
	for _, ai := range a {
		rv += fmt.Sprintf("%v", ai)
	}
	return rv
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
	return fmt.Sprintf("among %v", a.children)
}

type substring struct{}

func (s *substring) Accept(v visitor) {
	v.Visit(s)
}

func (s *substring) String() string {
	return fmt.Sprintf("substring")
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
	if u.operandCommand != nil {
		return fmt.Sprintf("%s %v", u.command, u.operandCommand)
	} else if u.operandName != nil {
		return fmt.Sprintf("%s %s", u.command, u.operandName)
	}
	return fmt.Sprintf("%s %v", u.command, u.operandAe)
}

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

type iCommand struct {
	name     *name
	operator string
	operand  ae
}

func (i *iCommand) Accept(v visitor) {
	v.Visit(i)
}

func (i *iCommand) String() string {
	return fmt.Sprintf("$%s %v %v", i.name, i.operator, i.operand)
}

type sCommand struct {
	name    *name
	operand command
}

func (s *sCommand) Accept(v visitor) {
	v.Visit(s)
}

func (s *sCommand) String() string {
	return fmt.Sprintf("$%s %v", s.name, s.operand)
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
		return fmt.Sprintf("loop %v %v", l.n, l.operand)
	}
	return fmt.Sprintf("atleast %v %v", l.n, l.operand)
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
	return fmt.Sprintf("%v %s %v", b.left, b.operator, b.right)
}
