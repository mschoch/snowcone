package main

import (
	"flag"
	"log"
	"os"

	"github.com/mschoch/snowcone"
)

var debugLexer = flag.Bool("debugLexer", false, "debug lexer")
var debugParser = flag.Bool("debugParser", false, "debug parser")

func main() {
	flag.Parse()

	snowcone.Logger = log.New(os.Stdout, "", log.LstdFlags)

	if flag.NArg() < 1 {
		log.Fatalf("you must specify the path to snowball file to compile")
	}

	if *debugLexer {
		snowcone.DebugLexer = true
	}

	if *debugParser {
		snowcone.DebugParser = true
	}

	f, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatalf("error opening snowball source file %s: %v", flag.Arg(0), err)
	}

	_, err = snowcone.Parse(f)
	if err != nil {
		log.Fatalf("parse error: %v", err)
	}

}
