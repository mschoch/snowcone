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

import "strings"

type stringdefs struct {
	defs    map[string]string
	start   string
	end     string
	initted bool
}

func (s *stringdefs) SetStringEscapes(start, end rune) {
	if !s.initted {
		s.defs = make(map[string]string)
		s.initted = true
	}
	s.start = string(start)
	s.end = string(end)
	s.defs["'"] = "'"
	s.defs[s.start] = s.start
}

func (s *stringdefs) ReplaceInLiteral(literal string) string {
	// until the stringescapes have been defined, don't do any replacements
	if !s.initted {
		return literal
	}
	// replace macro definitions
	for def, rep := range s.defs {
		literal = strings.Replace(literal, s.start+def+s.end, rep, -1)
	}
	// remove whitespace only escapes
	literal = s.removeWhitespaceMacros(literal)
	return literal
}

func (s *stringdefs) Define(def, replace string) {
	s.defs[def] = replace
}

func (s *stringdefs) removeWhitespaceMacros(literal string) string {
	maybeRemoved := s.removeNextWhitespaceMacro(literal)
	for maybeRemoved != literal {
		// something was removed, run it again to see if there is more
		literal = maybeRemoved
		maybeRemoved = s.removeNextWhitespaceMacro(literal)
	}
	return literal
}

func (s *stringdefs) removeNextWhitespaceMacro(literal string) string {
	check := literal
	ws := strings.Index(check, s.start)
	for ws > 0 && ws < len(check)-1 {
		rest := check[ws+1:]
		// found the start rune, now look for the next end rune
		we := strings.Index(rest, s.end)
		if we > 0 {
			if strings.TrimSpace(rest[:we]) == "" {
				// macro was entirely whitespace
				return strings.Replace(check, s.start+rest[:we]+s.end, "", -1)
			}
			check = rest[we:]
		} else {
			check = rest
		}
		ws = strings.Index(check, "{")
	}
	return literal
}
