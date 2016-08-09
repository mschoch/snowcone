# ![snowcone](snowcone.png) snowcone

Snowcone is a [snowball](http://snowballstem.org/) compiler for Go.  Currently, it is only capable of lexing and parsing the source.  It cannot yet generate the corresponding Go code.

## Status

**EXPERIMENTAL** – the API will change once we add the ability to generate code.

[![Build Status](https://travis-ci.org/mschoch/snowcone.svg?branch=master)](https://travis-ci.org/mschoch/snowcone)
[![GoDoc](https://godoc.org/github.com/mschoch/snowcone?status.svg)](https://godoc.org/github.com/mschoch/snowcone)
[![Go Report Card](https://goreportcard.com/badge/github.com/mschoch/snowcone)](https://goreportcard.com/report/github.com/mschoch/snowcone)
[![codebeat badge](https://codebeat.co/badges/6d6bab00-8112-48e6-bceb-f8545a9bda45)](https://codebeat.co/projects/github-com-mschoch-snowcone)
[![Coverage Status](https://coveralls.io/repos/github/mschoch/snowcone/badge.svg?branch=master)](https://coveralls.io/github/mschoch/snowcone?branch=master)

# TODO
- Parser
    - `get` directives
    - review grammar - rule expanding `C` to  `= S`, I encountered a reduce/reduce conflict that i haven't yet resolved
- Build interpreter to evaluate externals on user supplied input
- Generate Go code to execute externals in external apps

# Open Questions
- What does `?` command do?  Manual only includes it in the grammar and says it's a debugging aid.
