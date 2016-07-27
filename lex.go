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
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

type snowConeLex struct {
	in                *bufio.Reader
	buf               string
	possibleKeywords  keywordList
	nextToken         *yySymType
	nextTokenType     int
	nextRune          rune
	nextRuneSize      int
	pos               int
	line              int
	currState         lexState
	currConsumed      bool
	atEOF             bool
	stringEscapeStart rune
	stringEscapeEnd   rune
	stringEscapesSeen int
	insideEscapeSeq   bool
}

func (l *snowConeLex) reset() {
	l.buf = ""
	l.possibleKeywords = snowConeKeyWords
	l.stringEscapesSeen = 0
}

type lexState func(l *snowConeLex, next rune, eof bool) (lexState, bool)

func startState(l *snowConeLex, next rune, eof bool) (lexState, bool) {
	if eof {
		return nil, false
	}
	// fast cases up front
	switch {
	case next == '\'':
		l.buf += string(next)
		return inLiteralState, true
	case unicode.IsDigit(next):
		l.buf += string(next)
		return inNumberState, true
	}

	var maybeComment, maybeName, maybeKeyword bool
	l.possibleKeywords = l.possibleKeywords.StartWith(string(next))

	if next == '/' {
		maybeComment = true
	}
	if unicode.IsLetter(next) {
		maybeName = true
	}
	if len(l.possibleKeywords) > 0 {
		maybeKeyword = true
	}

	// a few cases are not checked below because they are not possible
	// maybeComment=true by itself cannot happen because / and /= are keywords
	// mabyeComment=true and maybeName=true cannot happen since / is not valid in name
	if maybeName && maybeKeyword {
		l.buf += string(next)
		return maybeNameKeywordState, true
	} else if maybeComment && maybeKeyword {
		l.buf += string(next)
		return maybeKeywordCommentState, true
	} else if maybeName {
		l.buf += string(next)
		return maybeNameState, true
	} else if maybeKeyword {
		l.buf += string(next)
		return maybeKeywordState, true
	}

	// doesnt look like anything, just eat it and stay here
	l.reset()
	return startState, true
}

func inLiteralState(l *snowConeLex, next rune, eof bool) (lexState, bool) {
	if eof {
		return nil, false
	}
	// FIXME handle string escapes
	if next == '\'' && !l.insideEscapeSeq {
		l.nextTokenType = tLITERAL
		l.nextToken = &yySymType{
			s: l.buf[1:],
		}
		logDebugTokens("LITERAL - '%s'", l.nextToken.s)
		l.reset()
		return startState, true
	} else if next == l.stringEscapeStart {
		l.insideEscapeSeq = true
	} else if next == l.stringEscapeEnd {
		l.insideEscapeSeq = false
	}
	l.buf += string(next)
	return inLiteralState, true
}

func inNumberState(l *snowConeLex, next rune, eof bool) (lexState, bool) {
	if !unicode.IsDigit(next) || eof {
		n, _ := strconv.Atoi(l.buf)
		l.nextTokenType = tNUMBER
		l.nextToken = &yySymType{
			n: n,
		}
		logDebugTokens("NUMBER - '%d'", l.nextToken.n)
		l.reset()
		return startState, false // didn't consume char
	}
	l.buf += string(next)
	return inNumberState, true
}

func maybeKeywordCommentState(l *snowConeLex, next rune, eof bool) (lexState, bool) {
	if !eof && l.buf == "/" {
		switch next {
		case '/':
			return inLineComment, true
		case '*':
			return inMultiLineComment, true
		}
	}

	return maybeKeywordState, false
}

func maybeNameKeywordState(l *snowConeLex, next rune, eof bool) (lexState, bool) {
	if !eof {
		var maybeName, maybeKeyword bool
		if unicode.IsLetter(next) || unicode.IsDigit(next) || next == '_' {
			maybeName = true
		}
		possibleKeywords := l.possibleKeywords.StartWith(l.buf + string(next))
		if len(possibleKeywords) > 0 {
			maybeKeyword = true
		}
		if maybeName && maybeKeyword {
			l.buf += string(next)
			l.possibleKeywords = possibleKeywords
			return maybeNameKeywordState, true
		} else if maybeName {
			l.buf += string(next)
			return maybeNameState, true
		}
		// checking maybeKeyword ONLY is not needed
		// this would be the case where upon seeing a new letter
		// something could have been a name or keyword, now can ONLY be
		// a keyword - but, that cannot happen with the current keywords
		// all keywords that start as a valid name, remain valid names
		// for example, if there was a keyword 'magic/' that would
		// have triggered this case because the / character makes it no
		// longer possible to be a name, and ONLY a keyword
	}

	// still might be a keyword, needs to be exact match with first possible
	if l.possibleKeywords[0] == l.buf {
		logDebugTokens("EOF first match")
		return finishKeyword(l)
	}
	logDebugTokens("EOF first NOT match")
	return finishName(l)
}

// we've already found the stringescapes token, looking for 2 characters to finish
func finishStringEscapesState(l *snowConeLex, next rune, eof bool) (lexState, bool) {
	if !eof {
		if l.stringEscapesSeen == 0 {
			if unicode.IsSpace(next) {
				// haven't seen start yet, eat whitespace
				return finishStringEscapesState, true
			}
			l.buf += string(next)
			l.stringEscapesSeen = 1
			return finishStringEscapesState, true
		}
		l.nextTokenType = keywordTokenTypes["stringescapes"]
		l.nextToken = &yySymType{
			s: l.buf + string(next),
		}
		logDebugTokens("STRINGESCAPES - '%s'", l.nextToken.s)
		l.reset()
		return startState, false
	}
	return nil, true
}

func finishKeyword(l *snowConeLex) (lexState, bool) {
	if l.possibleKeywords[0] == "stringescapes" {
		// just kidding not really done
		l.buf = ""
		return finishStringEscapesState, true
	}
	l.nextTokenType = keywordTokenTypes[l.possibleKeywords[0]]
	l.nextToken = &yySymType{}
	logDebugTokens("%s", strings.ToUpper(l.possibleKeywords[0]))
	l.reset()
	return startState, false
}

func finishName(l *snowConeLex) (lexState, bool) {
	l.nextTokenType = tNAME
	l.nextToken = &yySymType{
		s: l.buf,
	}
	logDebugTokens("NAME - %s", l.nextToken.s)
	l.reset()
	return startState, false
}

func maybeNameState(l *snowConeLex, next rune, eof bool) (lexState, bool) {
	if !eof {
		if unicode.IsLetter(next) || unicode.IsDigit(next) || next == '_' {
			l.buf += string(next)
			return maybeNameState, true
		}
	}

	// current buffer must be a valid name, finish it
	return finishName(l)
}

func maybeKeywordState(l *snowConeLex, next rune, eof bool) (lexState, bool) {
	if !eof {
		possibleKeywords := l.possibleKeywords.StartWith(l.buf + string(next))
		if len(possibleKeywords) > 0 {
			l.buf += string(next)
			l.possibleKeywords = possibleKeywords
			return maybeKeywordState, true
		}
	}

	// current buffer must be a keyword (since we're in this state)
	// so close it out
	return finishKeyword(l)
}

func inLineComment(l *snowConeLex, next rune, eof bool) (lexState, bool) {
	if next == '\r' || next == '\n' || eof {
		l.buf = ""
		return startState, true
	}
	return inLineComment, true
}

func inMultiLineComment(l *snowConeLex, next rune, eof bool) (lexState, bool) {
	if (next == '/' && l.buf[len(l.buf)-1] == '*') || eof {
		l.buf = ""
		return startState, true
	}
	l.buf += string(next)
	return inMultiLineComment, true
}

func newSnowConeLex(in io.Reader) *snowConeLex {
	return &snowConeLex{
		in:               bufio.NewReader(in),
		possibleKeywords: snowConeKeyWords,
		currState:        startState,
		currConsumed:     true,
	}
}

func (l *snowConeLex) Lex(lval *yySymType) int {
	var err error

	for l.nextToken == nil {
		if l.currConsumed {
			l.nextRune, l.nextRuneSize, err = l.in.ReadRune()
			if err != nil && err == io.EOF {
				l.nextRune = 0
				l.atEOF = true
			} else if err != nil {
				return 0
			}
		}
		l.currState, l.currConsumed = l.currState(l, l.nextRune, l.atEOF)
		if l.currState == nil {
			return 0
		}
	}

	*lval = *l.nextToken
	rv := l.nextTokenType
	l.nextToken = nil
	l.nextTokenType = 0
	return rv
}

func logDebugTokens(format string, v ...interface{}) {
	if DebugLexer {
		Logger.Printf(fmt.Sprintf("LEXER %s", format), v...)
	}
}

func (l *snowConeLex) Error(msg string) {
	panic(msg)
}

func (l *snowConeLex) SetStringEscapes(start, end rune) {
	l.stringEscapeStart = start
	l.stringEscapeEnd = end
	logDebugTokens("updated string escapes to %s - %s", string(l.stringEscapeStart), string(l.stringEscapeEnd))
}
