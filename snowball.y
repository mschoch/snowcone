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
n int
strings []string
declarations decls
gitems groupitems
g *grouping
no node
b *bliteral
r *routine
a *amongitem
ai amongitems
aexpr ae
ic *iCommand
sc *sCommand
p *prog
cl commands
}

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
tSTRINGESCAPES tSTRINGDEF tHEX tDECIMAL tUMINUS

%type <s>                tLITERAL
%type <n>                tNUMBER
%type <s>                tNAME
%type <s>                tSTRINGESCAPES
%type <strings>          names
%type <declarations>     declaration
%type <gitems>           gplusminuslist
%type <no>               nameorliteral
%type <g>                gdef
%type <b>                tTRUE
%type <b>                tFALSE
%type <no>               command
%type <r>                rdef
%type <no>                tQUESTION
%type <a>                amongitem
%type <ai>               amonglist
%type <cl>               commands
%type <aexpr>            ae
%type <ic>               icommand
%type <sc>               scommand
%type <p>                p
%type <p>                program
%type <s>                tSTRINGDEF


%left tOR tAND
%left tNOT tTEST tTRY tDO tFAIL tGOTO tGOPAST tREPEAT tBACKWARDS tREVERSE tLOOP
tATLEAST tNAME tFOR
%left tPLUS tMINUS
%left tMULT tDIV
%right tUMINUS

%%

input:
program
{
        logDebugGrammar("INPUT - %v", yylex.(*lexerWrapper).p)
        yylex.(*lexerWrapper).p = $1
};

program:
p
{
        logDebugGrammar("PROGRAM - single")
        $$ = $1
}
|
p program
{
        logDebugGrammar("PROGRAM - multi")
        $1.Combine($2)
        $$ = $1
};

p:
declaration
{
        p := &prog{}
        logDebugGrammar("P - decl")
        for _, decl := range $1 {
          p.Declare(decl)
        }
        $$ = p
}
|
rdef
{
        p := &prog{}
        logDebugGrammar("P - rdef")
        p.DefineRoutine($1)
        $$ = p
}
|
gdef
{
        p := &prog{}
        logDebugGrammar("P - gdef")
        p.DefineGroup($1)
        $$ = p
}
|
tBACKWARDMODE tLPAREN program tRPAREN
{
        logDebugGrammar("P - backwardmode")
        $3.SetBackwardMode()
        $$ = $3
}
|
tSTRINGESCAPES
{
        if utf8.RuneCountInString($1) == 2 {
          logDebugGrammar("P - stringescapes")
          first, len := utf8.DecodeRuneInString($1)
          second, len := utf8.DecodeRuneInString($1[len:])
          yylex.(*lexerWrapper).lex.(*snowConeLex).SetStringEscapes(first, second)
          yylex.(*lexerWrapper).sd.SetStringEscapes(first, second)
        } else {
          logDebugGrammar("P - stringescapes rune count NOT 2!!!")
        }
        $$ = &prog{}
}
|
tSTRINGDEF stringdefliteraltype tLITERAL
{
        logDebugGrammar("P - stringedef")
        replacedLiteral := yylex.(*lexerWrapper).sd.ReplaceInLiteral($3)
        yylex.(*lexerWrapper).sd.Define($1, replacedLiteral)
        $$ = &prog{}
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
        logDebugGrammar("DECLARATION - strings named: %v", $3)
        for _, name := range $3 {
          $$ = append($$, &decl{
            name: name,
            typ: sstring,
          })
        }
}
|
tINTEGERS tLPAREN names tRPAREN
{
        logDebugGrammar("DECLARATION - integers")
        for _, name := range $3 {
          $$ = append($$, &decl{
            name: name,
            typ: sinteger,
          })
        }
}
|
tBOOLEANS tLPAREN names tRPAREN
{
        logDebugGrammar("DECLARATION - booleans")
        logDebugGrammar("DECLARATION - integers")
        for _, name := range $3 {
          $$ = append($$, &decl{
            name: name,
            typ: sboolean,
          })
        }
}
|
tROUTINES tLPAREN names tRPAREN
{
        logDebugGrammar("DECLARATION - routines")
        for _, name := range $3 {
          $$ = append($$, &decl{
            name: name,
            typ: sroutine,
          })
        }
}
|
tEXTERNALS tLPAREN names tRPAREN
{
        logDebugGrammar("DECLARATION - externals")
        for _, name := range $3 {
          $$ = append($$, &decl{
            name: name,
            typ: sexternal,
          })
        }
}
|
tGROUPINGS tLPAREN names tRPAREN
{
        logDebugGrammar("DECLARATION - groupings")
        for _, name := range $3 {
          $$ = append($$, &decl{
            name: name,
            typ: sgrouping,
          })
        }
}
;

rdef:
tDEFINE tNAME tAS command
{
        logDebugGrammar("RDEF")
        $$ = &routine{name:$2, comm:$4}
};

nameorliteral:
tNAME
{
        logDebugGrammar("NAMEORLITERAL - name")
        $$ = &name{val:$1}
}
|
tLITERAL
{
        logDebugGrammar("NAMEORLITERAL - literal")
        replacedLiteral := yylex.(*lexerWrapper).sd.ReplaceInLiteral($1)
        $$ = &sliteral{val:replacedLiteral}
};

gplusminuslist:
nameorliteral
{
        logDebugGrammar("GPLUSMINUSLIST - single")
        $$ = groupitems{&groupitem{item: $1}}
}
|
nameorliteral tPLUS gplusminuslist
{
        logDebugGrammar("GPLUSMINUSLIST - multi")
        //$$ = append($3, &groupitem{item: $1})
        $$ = append($3, nil)
        copy($$[1:], $$[0:])
        $$[0] = &groupitem{item: $1}
}
|
nameorliteral tMINUS gplusminuslist
{
        logDebugGrammar("GPLUSMINUSLIST - MINUS multi")
        //$$ = append($3, &groupitem{item: $1, minus:true})
        $$ = append($3, nil)
        copy($$[1:], $$[0:])
        $$[0] = &groupitem{item: $1}
        $$[1].minus = true
};

gdef:
tDEFINE tNAME gplusminuslist
{
        $$ = &grouping{name: $2, children: $3}
        logDebugGrammar("GDEF - %v", $$)
};

commands:
command
{
        logDebugGrammar("COMMANDS - single")
        $$ = commands{$1}
}
|
command commands
{
        logDebugGrammar("COMMANDS - multi")
        $$ = append($2, nil)
        copy($$[1:], $$[0:])
        $$[0] = $1
};

command:
nameorliteral
{
        logDebugGrammar("COMMANDFACTOR - s")
        $$ = $1
}
|
icommand
{
        logDebugGrammar("COMMAND - icommand")
        $$ = $1
}
|
scommand
{
        logDebugGrammar("COMMAND - scommand")
        $$ = $1
}
|
tAMONG tLPAREN tRPAREN
{
        logDebugGrammar("COMMAND - among empty")
        $$ = &among{}
}
|
tAMONG tLPAREN amonglist tRPAREN
{
        logDebugGrammar("COMMAND - among list")
        $$ = &among{children:$3}
}
|
tTRUE
{
        logDebugGrammar("COMMANDFACTOR - true")
        $$ = &bliteral{val: true}
}
|
tFALSE
{
        logDebugGrammar("COMMANDFACTOR - false")
        $$ = &bliteral{val: false}
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
        $$ = $2
}
|
tNOT command
{
        logDebugGrammar("COMMAND - not")
        $$ = &unaryCommand{command:"not", operandCommand: $2}
}
|
tTEST command
{
        logDebugGrammar("COMMAND - test")
        $$ = &unaryCommand{command:"test", operandCommand: $2}
}
|
tTRY command
{
        logDebugGrammar("COMMAND - try")
        $$ = &unaryCommand{command:"try", operandCommand: $2}
}
|
tDO command
{
        logDebugGrammar("COMMAND - do")
        $$ = &unaryCommand{command:"do", operandCommand: $2}
}
|
tFAIL command
{
        logDebugGrammar("COMMAND - unary fail")
        $$ = &unaryCommand{command:"fail", operandCommand: $2}
}
|
tGOTO command
{
        logDebugGrammar("COMMAND - goto")
        $$ = &unaryCommand{command:"goto", operandCommand: $2}
}
|
tGOPAST command
{
        logDebugGrammar("COMMAND - unary gopast")
        $$ = &unaryCommand{command:"gopast", operandCommand: $2}
}
|
tREPEAT command
{
        logDebugGrammar("COMMAND - unary repeat")
        $$ = &unaryCommand{command:"repeat", operandCommand: $2}
}
|
tLOOP ae command
{
        logDebugGrammar("COMMAND - loop ae")
        $$ = &loop{n: $2, operand: $3}
}
|
tATLEAST ae command
{
        logDebugGrammar("COMMAND - loop ae")
        $$ = &loop{n: $2, operand: $3, extra:true}
}
|
tINSERT nameorliteral
{
        logDebugGrammar("COMMAND - insert")
        $$ = &unaryCommand{command:"insert", operandCommand: $2}
}
|
tATTACH nameorliteral
{
        logDebugGrammar("COMMAND - attach")
        $$ = &unaryCommand{command:"attach", operandCommand: $2}
}
|
tREPLACE nameorliteral
{
        logDebugGrammar("COMMAND - replace")
        $$ = &unaryCommand{command:"replace", operandCommand: $2}
}
|
tDELETE
{
        logDebugGrammar("COMMAND - delete")
        $$ = &nilaryCommand{operator:"delete"}
}
|
tHOP ae
{
        logDebugGrammar("COMMAND - hop")
        $$ = &unaryCommand{command:"hop", operandCommand: $2}
}
|
tNEXT
{
        logDebugGrammar("COMMAND - next")
        $$ = &nilaryCommand{operator:"next"}
}
|
tASSIGNR tNAME
{
        logDebugGrammar("COMMAND - assign right")
        $$ = &unaryCommand{command:"assignr", operandName: &name{val:$2}}
}
|
tLBRACKET
{
        logDebugGrammar("COMMAND - lbracket")
        $$ = &nilaryCommand{operator:"["}
}
|
tRBRACKET
{
        logDebugGrammar("COMMAND - rbracket")
        $$ = &nilaryCommand{operator:"]"}
}
|
tMOVER tNAME
{
        logDebugGrammar("COMMAND - move right")
        $$ = &unaryCommand{command:"mover", operandName: &name{val:$2}}
}
|
tSETMARK tNAME
{
        logDebugGrammar("COMMAND - setmark")
        $$ = &unaryCommand{command:"setmark", operandName: &name{val:$2}}
}
|
tTOMARK ae
{
        logDebugGrammar("COMMAND - tomark")
        $$ = &unaryCommand{command:"tomark", operandAe: $2}
}
|
tATMARK ae
{
        logDebugGrammar("COMMAND - atmark")
        $$ = &unaryCommand{command:"atmark", operandAe: $2}
}
|
tTOLIMIT
{
        logDebugGrammar("COMMAND - tolimit")
        $$ = &nilaryCommand{operator:"tolimit"}
}
|
tATLIMIT
{
        logDebugGrammar("COMMAND - atlimit")
        $$ = &nilaryCommand{operator:"atlimit"}
}
|
tSETLIMIT command tFOR command
{
        logDebugGrammar("COMMAND - setlimit")
        $$ = &binaryCommand{left:$2, operator:"setlimitfor", right:$4}
}
|
tBACKWARDS command
{
        logDebugGrammar("COMMAND - backwards")
        $$ = &unaryCommand{command:"backwards", operandCommand: $2}
}
|
tREVERSE command
{
        logDebugGrammar("COMMAND - reverse")
        $$ = &unaryCommand{command:"reverse", operandCommand: $2}
}
|
tSUBSTRING
{
        logDebugGrammar("COMMAND - substring")
        $$ = &nilaryCommand{operator:"substring"}
}
|
tSET tNAME
{
        logDebugGrammar("COMMAND - set")
        $$ = &unaryCommand{command:"set", operandName:&name{val:$2}}
}
|
tUNSET tNAME
{
        logDebugGrammar("COMMAND - unset")
        $$ = &unaryCommand{command:"unset", operandName:&name{val:$2}}
}
|
tNON tNAME
{
        logDebugGrammar("COMMAND - non")
        $$ = &non{gname:$2}
}
|
tNON tMINUS tNAME
{
        logDebugGrammar("COMMAND - non minus")
        $$ = &non{gname:$3, minus:true}
}
|
tQUESTION
{
        logDebugGrammar("COMMAND - question")
        $$ = &nilaryCommand{operator:"?"}
}
|
command tOR command
{
        logDebugGrammar("COMMANDTERM - or")
        $$ = &binaryCommand{left:$1, operator:"or", right:$3}
}
|
command tAND command
{
        logDebugGrammar("COMMANDTERM - and")
        $$ = &binaryCommand{left:$1, operator:"and", right:$3}
}
;

amonglist:
amongitem
{
        logDebugGrammar("AMONGLIST - single")
        $$ = amongitems{$1}
}
|
amongitem amonglist {
        logDebugGrammar("AMONGLIST - multi")
        $$ = append($2, nil)
        copy($$[1:], $$[0:])
        $$[0] = $1
};

amongitem:
tLITERAL
{
        logDebugGrammar("AMONGITEM - literal")
        replacedLiteral := yylex.(*lexerWrapper).sd.ReplaceInLiteral($1)
        $$ = &amongitem{slit:&sliteral{val:replacedLiteral}}
}
|
tLITERAL tNAME
{
        logDebugGrammar("AMONGITEM - literal name")
        replacedLiteral := yylex.(*lexerWrapper).sd.ReplaceInLiteral($1)
        $$ = &amongitem{slit:&sliteral{val:replacedLiteral}, rname:$2}
}
|
tLPAREN tRPAREN
{
        logDebugGrammar("AMONGITEM - paren empty")
        $$ = &amongitem{}
}
|
tLPAREN commands tRPAREN
{
        logDebugGrammar("AMONGITEM - paren command")
        $$ = &amongitem{comm:$2}
};

scommand:
tDOLLAR tNAME command
{
        logDebugGrammar("SCOMMAND")
        $$ = &sCommand{name:&name{val:$2}, operand:$3}
};

icommand:
tDOLLAR tNAME tASSIGN ae
{
        logDebugGrammar("ICOMMAND - assign")
        $$ = &iCommand{name:&name{val:$2}, operator:"=", operand:$4}
}
|
tDOLLAR tNAME tPLUSASSIGN ae
{
        logDebugGrammar("ICOMMAND - plus assign")
        $$ = &iCommand{name:&name{val:$2}, operator:"+=", operand:$4}
}
|
tDOLLAR tNAME tMINUSASSIGN ae
{
        logDebugGrammar("ICOMMAND - minus assign")
        $$ = &iCommand{name:&name{val:$2}, operator:"-=", operand:$4}
}
|
tDOLLAR tNAME tMULTASSIGN ae
{
        logDebugGrammar("ICOMMAND - mult assign")
        $$ = &iCommand{name:&name{val:$2}, operator:"*=", operand:$4}
}
|
tDOLLAR tNAME tDIVASSIGN ae
{
        logDebugGrammar("ICOMMAND - div assign")
        $$ = &iCommand{name:&name{val:$2}, operator:"/=", operand:$4}
}
|
tDOLLAR tNAME tEQ ae
{
        logDebugGrammar("ICOMMAND - eq")
        $$ = &iCommand{name:&name{val:$2}, operator:"==", operand:$4}
}
|
tDOLLAR tNAME tNEQ ae
{
        logDebugGrammar("ICOMMAND - neq")
        $$ = &iCommand{name:&name{val:$2}, operator:"!=", operand:$4}
}
|
tDOLLAR tNAME tGT ae
{
        logDebugGrammar("ICOMMAND - greater than")
        $$ = &iCommand{name:&name{val:$2}, operator:">", operand:$4}
}
|
tDOLLAR tNAME tLT ae
{
        logDebugGrammar("ICOMMAND - less than")
        $$ = &iCommand{name:&name{val:$2}, operator:"<", operand:$4}
}
|
tDOLLAR tNAME tGTEQ ae
{
        logDebugGrammar("ICOMMAND - greater than or eq")
        $$ = &iCommand{name:&name{val:$2}, operator:">=", operand:$4}
}
|
tDOLLAR tNAME tLTEQ ae
{
        logDebugGrammar("ICOMMAND - less than or eq")
        $$ = &iCommand{name:&name{val:$2}, operator:"<=", operand:$4}
};

ae:
ae tPLUS ae
{
        logDebugGrammar("AE - plus")
        $$ = &binaryAe{left:$1, operator:"+",right:$3}
}
|
ae tMINUS ae
{
        logDebugGrammar("AE - minus")
        $$ = &binaryAe{left:$1, operator:"-",right:$3}
}
|
ae tMULT ae
{
        logDebugGrammar("TERM - mult")
        $$ = &binaryAe{left:$1, operator:"*",right:$3}
}
|
ae tDIV ae
{
        logDebugGrammar("TERM - div")
        $$ = &binaryAe{left:$1, operator:"/",right:$3}
}
|
tMINUS ae %prec tUMINUS
{
        logDebugGrammar("AE - unary minus")
        $$ = &unaryAe{operator:"uminus", operand: $2}
}
|
tMAXINT
{
        logDebugGrammar("AE - maxint")
        $$ = &nilaryAe{operator:"maxint"}
}
|
tMININT
{
        logDebugGrammar("AE - minint")
        $$ = &nilaryAe{operator:"minint"}
}
|
tCURSOR
{
        logDebugGrammar("AE - cursor")
        $$ = &nilaryAe{operator:"cursor"}
}
|
tLIMIT
{
        logDebugGrammar("AE - limit")
        $$ = &nilaryAe{operator:"limit"}
}
|
tSIZE
{
        logDebugGrammar("AE - size")
        $$ = &nilaryAe{operator:"size"}
}
|
tSIZEOF tNAME
{
        logDebugGrammar("AE - sizeof name")
        $$ = &unaryAe{operator:"sizeof", operand: &name{val:$2}}
}
|
tLEN
{
        logDebugGrammar("AE - len")
        $$ = &nilaryAe{operator:"len"}
}
|
tLENOF tNAME
{
        logDebugGrammar("AE - leno name")
        $$ = &unaryAe{operator:"lenof", operand: &name{val:$2}}
}
|
tNAME
{
        logDebugGrammar("FACTOR - name")
        $$ = &name{val:$1}
}
|
tNUMBER
{
        logDebugGrammar("FACTOR - number")
        $$ = &nliteral{val:$1}
}
|
tLPAREN ae tRPAREN
{
        logDebugGrammar("FACTOR - parens")
        $$ = $2
}
;

names:
tNAME
{
        logDebugGrammar("NAMES - single")
        $$ = []string{$1}
}
|
tNAME names
{
        logDebugGrammar("NAMEs - multi")
        $$ = append($2, "")
        copy($$[1:], $$[0:])
        $$[0] = $1
};
