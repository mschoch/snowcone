%{
package snowcone

import (
  "fmt"
  "unicode/utf8"
)

func logDebugGrammar(format string, v ...interface{}) {
    if DebugParser {
        Logger.Printf(fmt.Sprintf("PARSER %s", format), v...)
    }
}
%}

%union {
s string
n int}

%token tLITERAL tNUMBER tNAME tSTRINGS tINTEGERS tBOOLEANS tROUTINES
tEXTERNALS tGROUPINGS tLPAREN tRPAREN tDEFINE tAS tPLUS tMINUS
tMULT tDIV tMAXINT tMININT tCURSOR tLIMIT tSIZE tSIZEOF tLEN tLENOF
tDOLLAR tASSIGN tPLUSASSIGN tMINUSASSIGN tMULTASSIGN tDIVASSIGN
tEQ tNEQ tGT tLT tGTEQ tLTEQ tOR tAND tNOT tTEST tTRY tDO tFAIL
tGOTO tGOPAST tREPEAT tLOOP tATLEAST tINSERT tATTACH
tREPLACE tDELETE tHOP tNEXT tASSIGNR tLBRACKET tRBRACKET
tMOVER tSETMARK tTOMARK tATMARK tTOLIMIT tATLIMIT tSETLIMIT
tFOR tBACKWARDS tREVERSE tSUBSTRING tAMONG tSET tUNSET
tNON tTRUE tFALSE tBACKWARDMODE tQUESTION
tSTRINGESCAPES tSTRINGDEF tHEX tDECIMAL

%type <s>                tLITERAL
%type <n>                tNUMBER
%type <s>                tNAME
%type <s>                tSTRINGESCAPES

%left tNOT tTEST tTRY tDO tFAIL tGOTO tGOPAST tREPEAT tBACKWARDS tREVERSE tLOOP
tATLEAST tNAME tFOR
%left tOR tAND

%%

input:
program
{
        logDebugGrammar("INPUT")
};

program:
p
{
        logDebugGrammar("PROGRAM - single")
}
|
p program
{
        logDebugGrammar("PROGRAM - multi")
};

p:
declaration
{
        logDebugGrammar("P - decl")
}
|
rdef
{
        logDebugGrammar("P - rdef")
}
|
gdef
{
        logDebugGrammar("P - gdef")
}
|
tBACKWARDMODE tLPAREN program tRPAREN
{
        logDebugGrammar("P - backwardmode")
}
|
tSTRINGESCAPES
{
        if utf8.RuneCountInString($1) == 2 {
          logDebugGrammar("P - stringescapes")
          first, len := utf8.DecodeRuneInString($1)
          second, len := utf8.DecodeRuneInString($1[len:])
          yylex.(*lexerWrapper).lex.(*snowConeLex).SetStringEscapes(first, second)
        } else {
          logDebugGrammar("P - stringescapes rune count NOT 2!!!")
        }
}
|
tSTRINGDEF stringdefliteraltype tLITERAL
{
        logDebugGrammar("P - stringedef")
};

stringdefliteraltype:
/*empty*/
{

}
|
tHEX
{
        logDebugGrammar("STRINGDEFLITERALTYPE - hex")
}
|
tDECIMAL
{
        logDebugGrammar("STRINGDEFLITERALTYPE - decimal")
};

declaration:
tSTRINGS tLPAREN names tRPAREN
{
        logDebugGrammar("DECLARATION - strings")
}
|
tINTEGERS tLPAREN names tRPAREN
{
        logDebugGrammar("DECLARATION - integers")
}
|
tBOOLEANS tLPAREN names tRPAREN
{
        logDebugGrammar("DECLARATION - booleans")
}
|
tROUTINES tLPAREN names tRPAREN
{
        logDebugGrammar("DECLARATION - routines")
}
|
tEXTERNALS tLPAREN names tRPAREN
{
        logDebugGrammar("DECLARATION - externals")
}
|
tGROUPINGS tLPAREN names tRPAREN
{
        logDebugGrammar("DECLARATION - groupings")
}
;

rdef:
tDEFINE tNAME tAS command
{
        logDebugGrammar("RDEF")
};

nameorliteral:
tNAME
{
        logDebugGrammar("NAMEORLITERAL - name")
}
|
tLITERAL
{
        logDebugGrammar("NAMEORLITERAL - literal")
};

plusorminus:
tPLUS
{
        logDebugGrammar("PLUSORMINUS - plus")
}
|
tMINUS
{
        logDebugGrammar("PLUSORMINUS - minus")
};

gplusminuslist:
nameorliteral
{
        logDebugGrammar("GPLUSMINUSLIST - single")
}
|
nameorliteral plusorminus gplusminuslist
{
        logDebugGrammar("GPLUSMINUSLIST - multi")
};

gdef:
tDEFINE tNAME gplusminuslist
{
        logDebugGrammar("GDEF")
};

commands:
command
{
        logDebugGrammar("COMMANDS - single")
}
|
command commands
{
        logDebugGrammar("COMMANDS - multi")
};

command:
nameorliteral
{
        logDebugGrammar("COMMANDFACTOR - s")
}
|
icommand
{
        logDebugGrammar("COMMAND - icommand")
}
|
scommand
{
        logDebugGrammar("COMMAND - scommand")
}
|
tAMONG tLPAREN tRPAREN
{
        logDebugGrammar("COMMAND - among empty")
}
|
tAMONG tLPAREN amonglist tRPAREN
{
        logDebugGrammar("COMMAND - among list")
}
|
tTRUE
{
        logDebugGrammar("COMMANDFACTOR - true")
}
|
tFALSE
{
        logDebugGrammar("COMMANDFACTOR - false")
}
|
tLPAREN tRPAREN
{
        logDebugGrammar("COMMANDFACTOR - paren empty")
}
|
tLPAREN commands tRPAREN
{
        logDebugGrammar("COMMANDFACTOR - paren commands")
}
|
tNOT command
{
        logDebugGrammar("COMMAND - unary not")
}
|
tTEST command
{
        logDebugGrammar("COMMAND - unary test")
}
|
tTRY command
{
        logDebugGrammar("COMMAND - unary try")
}
|
tDO command
{
        logDebugGrammar("COMMAND - unary do")
}
|
tFAIL command
{
        logDebugGrammar("COMMAND - unary fail")
}
|
tGOTO command
{
        logDebugGrammar("COMMAND - unary goto")
}
|
tGOPAST command
{
        logDebugGrammar("COMMAND - unary gopast")
}
|
tREPEAT command
{
        logDebugGrammar("COMMAND - unary repeat")
}
|
tLOOP ae command
{
        logDebugGrammar("COMMAND - loop ae")
}
|
tATLEAST ae command
{
        logDebugGrammar("COMMAND - loop ae")
}
|
tINSERT nameorliteral
{
        logDebugGrammar("COMMAND - insert")
}
|
tATTACH nameorliteral
{
        logDebugGrammar("COMMAND - attach")
}
|
tREPLACE nameorliteral
{
        logDebugGrammar("COMMAND - replace")
}
|
tDELETE
{
        logDebugGrammar("COMMAND - delete")
}
|
tHOP ae
{
        logDebugGrammar("COMMAND - hop")
}
|
tNEXT
{
        logDebugGrammar("COMMAND - next")
}
|
tASSIGNR
{
        logDebugGrammar("COMMAND - assign right")
}
|
tLBRACKET
{
        logDebugGrammar("COMMAND - lbracket")
}
|
tRBRACKET
{
        logDebugGrammar("COMMAND - rbracket")
}
|
tMOVER tNAME
{
        logDebugGrammar("COMMAND - move right")
}
|
tSETMARK tNAME
{
        logDebugGrammar("COMMAND - setmark")
}
|
tTOMARK ae
{
        logDebugGrammar("COMMAND - tomark")
}
|
tATMARK ae
{
        logDebugGrammar("COMMAND - atmark")
}
|
tTOLIMIT
{
        logDebugGrammar("COMMAND - tolimit")
}
|
tATLIMIT
{
        logDebugGrammar("COMMAND - atlimit")
}
|
tSETLIMIT command tFOR command
{
        logDebugGrammar("COMMAND - setlimit")
}
|
tBACKWARDS command
{
        logDebugGrammar("COMMAND - backwards")
}
|
tREVERSE command
{
        logDebugGrammar("COMMAND - reverse")
}
|
tSUBSTRING
{
        logDebugGrammar("COMMAND - substring")
}
|
tSET tNAME
{
        logDebugGrammar("COMMAND - set")
}
|
tUNSET tNAME
{
        logDebugGrammar("COMMAND - unset")
}
|
tNON tNAME
{
        logDebugGrammar("COMMAND - non")
}
|
tNON tMINUS tNAME
{
        logDebugGrammar("COMMAND - non minus")
}
|
tQUESTION
{
        logDebugGrammar("COMMAND - question")
}
|
command tOR command
{
        logDebugGrammar("COMMANDTERM - or")
}
|
command tAND command
{
        logDebugGrammar("COMMANDTERM - and")
}
;

amonglist:
amongitem
{
        logDebugGrammar("AMONGLIST - single")
}
|
amongitem amonglist {
        logDebugGrammar("AMONGLIST - multi")
};

amongitem:
tLITERAL
{
        logDebugGrammar("AMONGITEM - literal")
}
|
tLITERAL tNAME
{
        logDebugGrammar("AMONGITEM - literal name")
}
|
tLITERAL tLPAREN tRPAREN
{
        logDebugGrammar("AMONGITEM - paren empty")
}
|
tLITERAL tLPAREN commands tRPAREN
{
        logDebugGrammar("AMONGITEM - paren command")
};

scommand:
tDOLLAR tNAME command
{
        logDebugGrammar("SCOMMAND")
};

icommand:
tDOLLAR tNAME tASSIGN ae
{
        logDebugGrammar("ICOMMAND - assign")
}
|
tDOLLAR tNAME tPLUSASSIGN ae
{
        logDebugGrammar("ICOMMAND - plus assign")
}
|
tDOLLAR tNAME tMINUSASSIGN ae
{
        logDebugGrammar("ICOMMAND - minus assign")
}
|
tDOLLAR tNAME tMULTASSIGN ae
{
        logDebugGrammar("ICOMMAND - mult assign")
}
|
tDOLLAR tNAME tDIVASSIGN ae
{
        logDebugGrammar("ICOMMAND - div assign")
}
|
tDOLLAR tNAME tEQ ae
{
        logDebugGrammar("ICOMMAND - eq")
}
|
tDOLLAR tNAME tNEQ ae
{
        logDebugGrammar("ICOMMAND - neq")
}
|
tDOLLAR tNAME tGT ae
{
        logDebugGrammar("ICOMMAND - greater than")
}
|
tDOLLAR tNAME tLT ae
{
        logDebugGrammar("ICOMMAND - less than")
}
|
tDOLLAR tNAME tGTEQ ae
{
        logDebugGrammar("ICOMMAND - greater than or eq")
}
|
tDOLLAR tNAME tLTEQ ae
{
        logDebugGrammar("ICOMMAND - less than or eq")
};

ae:
term tPLUS ae
{
        logDebugGrammar("AE - plus")
}
|
term tMINUS ae
{
        logDebugGrammar("AE - minus")
}
|
term
{
        logDebugGrammar("AE - term")
}
|
tMINUS ae
{
        logDebugGrammar("AE - unary minus")
}
|
tMAXINT
{
        logDebugGrammar("AE - maxint")
}
|
tMININT
{
        logDebugGrammar("AE - minint")
}
|
tCURSOR
{
        logDebugGrammar("AE - cursor")
}
|
tLIMIT
{
        logDebugGrammar("AE - limit")
}
|
tSIZE
{
        logDebugGrammar("AE - size")
}
|
tSIZEOF tNAME
{
        logDebugGrammar("AE - sizeof name")
}
|
tLEN
{
        logDebugGrammar("AE - len")
}
|
tLENOF tNAME
{
        logDebugGrammar("AE - leno name")
}
;

term:
factor
{
        logDebugGrammar("TERM - factor")
}
|
factor tMULT term
{
        logDebugGrammar("TERM - mult")
}
|
factor tDIV term
{
        logDebugGrammar("TERM - div")
};

factor:
tNAME
{
        logDebugGrammar("FACTOR - name")
}
|
tNUMBER
{
        logDebugGrammar("FACTOR - number")
}
|
tLPAREN ae tRPAREN
{
        logDebugGrammar("FACTOR - parens")
};

names:
tNAME
{
        logDebugGrammar("NAMES - single")
}
|
tNAME names
{
        logDebugGrammar("NAMEs - multi")
};
