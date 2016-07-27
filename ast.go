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
	children []node
}

func (p *prog) Accept(v visitor) {
	v.Visit(p)
}

func (p *prog) String() string {
	rv := "prog\n"
	for _, n := range p.children {
		rv += fmt.Sprintf("%v\n", n)
	}
	return rv
}

func (p *prog) Declare(d *decl) {
	p.children = append(p.children, d)
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
	name string
	comm node
}

func (r *routine) Accept(v visitor) {
	v.Visit(r)
}

func (r *routine) String() string {
	return fmt.Sprintf("define routine %s as %v", r.name, r.comm)
}

type groupitem struct {
	minus bool
	name  string
	slit  string
	next  *groupitem
}

type grouping struct {
	name  string
	child *groupitem
}
