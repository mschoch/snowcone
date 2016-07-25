# snowcone

Snowcone is a [snowball](http://snowballstem.org/) compiler for Go.  Currently, it is only capable of lexing and parsing the source.  It cannot yet generate the corresponding Go code.

## Status

**EXPERIMENTAL** – the API will change once we add the ability to generate code.

[![Build Status](https://travis-ci.org/mschoch/snowcone.svg?branch=master)](https://travis-ci.org/mschoch/snowcone)
[![GoDoc](https://godoc.org/github.com/mschoch/snowcone?status.svg)](https://godoc.org/github.com/mschoch/snowcone)
[![Go Report Card](https://goreportcard.com/badge/github.com/mschoch/snowcone)](https://goreportcard.com/report/github.com/mschoch/snowcone)
[![codebeat badge](https://codebeat.co/badges/6d6bab00-8112-48e6-bceb-f8545a9bda45)](https://codebeat.co/projects/github-com-mschoch-snowcone)

# TODO
- Parser
    - `get` directives
    - stringdef
    - review grammar - rule expanding `C` to  `= S`, I encountered a reduce/reduce conflict that i haven't yet resolved
- Build AST
- Generate Go code
