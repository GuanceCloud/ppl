// Code generated by goyacc -o gram_y.go gram.y. DO NOT EDIT.

//line gram.y:7
package parser

import __yyfmt__ "fmt"

//line gram.y:7

import (
	ast "github.com/GuanceCloud/platypus/pkg/ast"
)

//line gram.y:15
type yySymType struct {
	yys      int
	aststmts ast.Stmts
	astblock *ast.BlockStmt

	ifitem *ast.IfStmtElem
	iflist []*ast.IfStmtElem
	node   *ast.Node
	nodes  []*ast.Node
	item   Item
}

const SEMICOLON = 57346
const COMMA = 57347
const COMMENT = 57348
const DOT = 57349
const EOF = 57350
const ERROR = 57351
const ID = 57352
const NUMBER = 57353
const LEFT_PAREN = 57354
const LEFT_BRACKET = 57355
const LEFT_BRACE = 57356
const RIGHT_BRACE = 57357
const RIGHT_PAREN = 57358
const RIGHT_BRACKET = 57359
const SPACE = 57360
const STRING = 57361
const QUOTED_STRING = 57362
const MULTILINE_STRING = 57363
const FOR = 57364
const IN = 57365
const WHILE = 57366
const BREAK = 57367
const CONTINUE = 57368
const RETURN = 57369
const EOL = 57370
const COLON = 57371
const STR = 57372
const INT = 57373
const FLOAT = 57374
const BOOL = 57375
const LIST = 57376
const MAP = 57377
const operatorsStart = 57378
const ADD = 57379
const DIV = 57380
const GTE = 57381
const GT = 57382
const LT = 57383
const LTE = 57384
const MOD = 57385
const MUL = 57386
const NEQ = 57387
const EQ = 57388
const EQEQ = 57389
const SUB = 57390
const operatorsEnd = 57391
const keywordsStart = 57392
const TRUE = 57393
const FALSE = 57394
const IDENTIFIER = 57395
const AND = 57396
const OR = 57397
const NIL = 57398
const NULL = 57399
const IF = 57400
const ELIF = 57401
const ELSE = 57402
const keywordsEnd = 57403
const startSymbolsStart = 57404
const START_STMTS = 57405
const startSymbolsEnd = 57406

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"SEMICOLON",
	"COMMA",
	"COMMENT",
	"DOT",
	"EOF",
	"ERROR",
	"ID",
	"NUMBER",
	"LEFT_PAREN",
	"LEFT_BRACKET",
	"LEFT_BRACE",
	"RIGHT_BRACE",
	"RIGHT_PAREN",
	"RIGHT_BRACKET",
	"SPACE",
	"STRING",
	"QUOTED_STRING",
	"MULTILINE_STRING",
	"FOR",
	"IN",
	"WHILE",
	"BREAK",
	"CONTINUE",
	"RETURN",
	"EOL",
	"COLON",
	"STR",
	"INT",
	"FLOAT",
	"BOOL",
	"LIST",
	"MAP",
	"operatorsStart",
	"ADD",
	"DIV",
	"GTE",
	"GT",
	"LT",
	"LTE",
	"MOD",
	"MUL",
	"NEQ",
	"EQ",
	"EQEQ",
	"SUB",
	"operatorsEnd",
	"keywordsStart",
	"TRUE",
	"FALSE",
	"IDENTIFIER",
	"AND",
	"OR",
	"NIL",
	"NULL",
	"IF",
	"ELIF",
	"ELSE",
	"keywordsEnd",
	"startSymbolsStart",
	"START_STMTS",
	"startSymbolsEnd",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line gram.y:507

//line yacctab:1
var yyExca = [...]int16{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 157,
	1, 46,
	4, 46,
	8, 46,
	15, 46,
	28, 46,
	-2, 94,
}

const yyPrivate = 57344

const yyLast = 903

var yyAct = [...]uint8{
	21, 170, 135, 3, 23, 107, 105, 65, 63, 81,
	78, 81, 162, 30, 82, 80, 82, 80, 179, 67,
	40, 79, 36, 160, 145, 167, 8, 163, 88, 84,
	164, 100, 86, 101, 60, 137, 137, 164, 4, 164,
	83, 66, 99, 94, 61, 78, 81, 70, 71, 74,
	75, 82, 80, 76, 69, 77, 79, 103, 1, 85,
	14, 95, 73, 72, 2, 84, 108, 15, 62, 111,
	113, 114, 115, 116, 117, 118, 119, 120, 121, 122,
	123, 124, 125, 126, 106, 129, 104, 45, 136, 138,
	54, 16, 149, 132, 97, 5, 142, 127, 144, 130,
	55, 147, 146, 139, 96, 35, 128, 106, 131, 7,
	151, 152, 140, 155, 95, 150, 59, 98, 34, 156,
	78, 81, 70, 71, 74, 75, 82, 80, 76, 106,
	77, 79, 33, 157, 32, 37, 161, 73, 24, 39,
	78, 81, 70, 71, 74, 75, 82, 80, 76, 169,
	77, 79, 25, 29, 173, 175, 26, 44, 171, 178,
	174, 176, 177, 180, 89, 91, 42, 181, 28, 87,
	88, 27, 106, 182, 157, 157, 106, 90, 9, 183,
	184, 89, 185, 13, 12, 11, 87, 88, 92, 186,
	10, 43, 133, 157, 106, 17, 106, 109, 15, 64,
	22, 45, 6, 106, 54, 52, 41, 38, 56, 102,
	53, 0, 0, 48, 55, 49, 18, 0, 0, 20,
	19, 0, 16, 0, 0, 0, 0, 0, 0, 0,
	0, 57, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 58, 0, 0, 46, 47, 0, 0, 0,
	50, 51, 31, 15, 0, 0, 45, 0, 0, 54,
	52, 41, 38, 56, 0, 0, 0, 0, 48, 55,
	49, 18, 0, 0, 20, 19, 0, 16, 0, 0,
	0, 0, 0, 0, 0, 0, 57, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 58, 0, 0,
	46, 47, 0, 0, 0, 50, 51, 31, 45, 0,
	0, 54, 52, 41, 38, 56, 0, 0, 0, 0,
	48, 55, 49, 18, 0, 0, 20, 19, 0, 0,
	45, 0, 0, 54, 52, 41, 38, 56, 57, 134,
	0, 0, 48, 55, 49, 0, 0, 0, 0, 58,
	0, 137, 46, 47, 0, 0, 0, 50, 51, 31,
	57, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 58, 0, 0, 46, 47, 0, 0, 153, 50,
	51, 45, 0, 0, 54, 52, 41, 38, 56, 0,
	0, 0, 0, 48, 55, 49, 0, 0, 0, 0,
	0, 0, 45, 0, 0, 54, 52, 41, 38, 56,
	143, 57, 0, 0, 48, 55, 49, 0, 0, 0,
	0, 0, 58, 0, 0, 46, 47, 0, 0, 0,
	50, 51, 57, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 58, 0, 0, 46, 47, 0, 0,
	112, 50, 51, 45, 0, 0, 54, 52, 41, 38,
	56, 0, 0, 0, 0, 48, 55, 49, 0, 0,
	0, 0, 0, 0, 45, 0, 0, 54, 52, 41,
	38, 56, 102, 57, 0, 0, 48, 55, 49, 0,
	0, 0, 0, 0, 58, 0, 0, 46, 47, 0,
	0, 0, 50, 51, 57, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 58, 0, 0, 46, 47,
	0, 0, 68, 50, 51, 45, 0, 0, 54, 52,
	41, 38, 56, 0, 0, 0, 0, 48, 55, 49,
	0, 0, 0, 0, 0, 0, 45, 0, 0, 54,
	52, 41, 38, 158, 0, 57, 0, 0, 48, 55,
	49, 0, 0, 0, 0, 0, 58, 0, 0, 46,
	47, 0, 0, 0, 50, 51, 57, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 107, 58, 0, 0,
	46, 47, 0, 0, 45, 50, 51, 54, 52, 41,
	38, 56, 0, 0, 0, 0, 48, 55, 49, 78,
	81, 70, 71, 74, 75, 82, 80, 76, 69, 77,
	79, 0, 0, 0, 57, 0, 73, 72, 0, 0,
	0, 172, 0, 0, 0, 58, 0, 148, 46, 47,
	0, 0, 0, 50, 51, 78, 81, 70, 71, 74,
	75, 82, 80, 76, 69, 77, 79, 0, 0, 0,
	0, 0, 73, 72, 78, 81, 70, 71, 74, 75,
	82, 80, 76, 69, 77, 79, 168, 0, 0, 0,
	0, 73, 72, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 78, 81, 70, 71,
	74, 75, 82, 80, 76, 69, 77, 79, 166, 0,
	0, 0, 0, 73, 72, 165, 78, 81, 70, 71,
	74, 75, 82, 80, 76, 69, 77, 79, 0, 0,
	0, 0, 0, 73, 72, 78, 81, 70, 71, 74,
	75, 82, 80, 76, 69, 77, 79, 159, 0, 0,
	0, 0, 73, 72, 154, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 78, 81, 70,
	71, 74, 75, 82, 80, 76, 69, 77, 79, 110,
	0, 0, 0, 0, 73, 72, 0, 78, 81, 70,
	71, 74, 75, 82, 80, 76, 69, 77, 79, 0,
	0, 0, 0, 0, 73, 72, 0, 0, 0, 0,
	0, 0, 78, 81, 70, 71, 74, 75, 82, 80,
	76, 69, 77, 79, 0, 0, 0, 0, 0, 73,
	72, 78, 81, 70, 71, 74, 75, 82, 80, 76,
	69, 77, 79, 54, 52, 0, 0, 0, 73, 72,
	141, 0, 48, 55, 49, 54, 52, 0, 0, 0,
	0, 0, 93, 0, 48, 55, 49, 0, 0, 0,
	57, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 58, 57, 0, 46, 47, 0, 0, 0, 50,
	51, 0, 0, 58, 0, 0, 46, 47, 0, 0,
	0, 50, 51,
}

var yyPact = [...]int16{
	1, 30, 249, -1000, -1000, -1000, 301, 63, 40, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -52, 518, -1000,
	-1000, 794, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 33,
	52, 587, -1000, -1000, -1000, -1000, 157, 160, 845, 89,
	-1000, 587, -1000, -1000, -1000, 18, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 22, -1000, -1000, 467, -1000, -1000, 63,
	40, -1000, -1000, -9, -1000, 587, 174, 775, 446, 587,
	587, 587, 587, 587, 587, 587, 587, 587, 587, 587,
	587, 587, 587, 80, 587, 80, 572, 323, 587, 80,
	-1000, 833, -1000, -1000, -1000, -1000, -1000, 395, -1000, 8,
	587, -1000, -1000, 608, 40, -1000, -1000, 194, 572, 587,
	374, 750, 539, 794, -27, -27, 83, 103, -27, -27,
	-27, -27, -29, -29, -1000, -1000, -1000, 16, 15, 730,
	16, 15, -1000, 7, -1000, 11, 794, -1000, 698, 16,
	15, -1000, -1000, -1000, 679, -1000, 9, 659, 587, -14,
	-1000, 572, 627, 539, 539, 572, -1000, -1000, 194, -1000,
	-1000, 2, 587, -1000, -1000, -1000, 587, -1000, -1000, 794,
	-1000, -1000, 539, 572, -1000, 572, -1000, -1000, 608, -1000,
	794, 794, 572, -1000, -1000, -1000, -1000,
}

var yyPgo = [...]uint8{
	0, 210, 6, 20, 92, 202, 200, 199, 195, 192,
	109, 191, 190, 185, 184, 183, 178, 171, 22, 168,
	166, 157, 156, 13, 153, 0, 152, 139, 138, 135,
	4, 134, 132, 118, 105, 60, 58, 26, 2,
}

var yyR1 = [...]int8{
	0, 37, 37, 37, 37, 36, 36, 36, 4, 4,
	4, 5, 5, 5, 10, 10, 10, 10, 10, 10,
	35, 25, 25, 25, 25, 25, 25, 25, 25, 15,
	14, 12, 13, 13, 13, 13, 13, 13, 13, 13,
	16, 16, 6, 8, 8, 7, 2, 2, 3, 17,
	17, 17, 17, 9, 9, 9, 19, 19, 19, 11,
	20, 20, 20, 20, 20, 20, 20, 20, 21, 21,
	21, 21, 21, 22, 22, 38, 38, 23, 23, 23,
	24, 24, 24, 24, 24, 24, 28, 28, 28, 29,
	29, 29, 26, 26, 26, 27, 27, 27, 30, 30,
	30, 30, 30, 31, 31, 32, 32, 33, 33, 34,
	34, 18, 18, 1, 1,
}

var yyR2 = [...]int8{
	0, 1, 1, 2, 2, 2, 2, 1, 2, 1,
	1, 2, 1, 3, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 5, 7, 6, 6, 5, 6, 5, 5, 4,
	1, 3, 3, 1, 2, 3, 1, 3, 2, 4,
	3, 5, 4, 3, 2, 1, 1, 1, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 1, 2, 4, 4, 4,
	3, 3, 3, 3, 3, 3, 2, 3, 2, 2,
	3, 2, 2, 3, 1, 4, 5, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 1, 1, 1, 1,
}

var yyChk = [...]int16{
	-1000, -36, 63, 2, 8, -4, -5, -10, -37, -16,
	-12, -13, -14, -15, -35, 4, 28, -8, 22, 26,
	25, -25, -6, -30, -28, -26, -22, -17, -19, -24,
	-23, 58, -31, -32, -33, -34, -18, -29, 13, -27,
	-3, 12, -20, -11, -21, 7, 51, 52, 19, 21,
	56, 57, 11, -1, 10, 20, 14, 37, 48, -10,
	-37, 4, 28, 60, -7, 59, -18, -25, 4, 46,
	39, 40, 55, 54, 41, 42, 45, 47, 37, 48,
	44, 38, 43, 7, 13, 7, -25, 12, 13, 7,
	17, 5, 28, 17, -30, -18, 15, 5, 28, -25,
	13, 11, 15, -25, -37, -2, -3, 14, -25, 23,
	4, -25, 4, -25, -25, -25, -25, -25, -25, -25,
	-25, -25, -25, -25, -25, -25, -25, -23, -18, -25,
	-23, -18, -2, -9, 16, -38, -25, 28, -25, -23,
	-18, 17, -30, 15, -25, 16, -38, -25, 29, -4,
	-2, -25, -25, 4, 4, -25, -2, -3, 14, 17,
	16, -38, 5, 16, 28, 17, 29, 16, 17, -25,
	15, -2, 4, -25, -2, -25, -2, -2, -25, 16,
	-25, -25, -25, -2, -2, -2, -2,
}

var yyDef = [...]int8{
	0, -2, 0, 7, 6, 5, 9, 10, 12, 14,
	15, 16, 17, 18, 19, 1, 2, 40, 0, 30,
	29, 20, 43, 21, 22, 23, 24, 25, 26, 27,
	28, 0, 98, 99, 100, 101, 102, 0, 0, 0,
	94, 0, 56, 57, 58, 0, 103, 104, 105, 106,
	107, 108, 109, 0, 111, 112, 0, 113, 114, 8,
	11, 3, 4, 0, 44, 0, 102, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	86, 0, 91, 88, 89, 102, 92, 0, 97, 0,
	0, 110, 48, 0, 13, 41, 46, 0, 0, 0,
	0, 0, 0, 59, 60, 61, 62, 63, 64, 65,
	66, 67, 68, 69, 70, 71, 72, 84, 85, 0,
	82, 83, 42, 0, 50, 0, 55, 75, 0, 80,
	81, 87, 90, 93, 0, 73, 0, 0, 0, 0,
	45, 0, 0, 0, 0, 0, 39, -2, 0, 79,
	49, 0, 54, 52, 76, 77, 0, 74, 78, 95,
	47, 31, 0, 0, 37, 0, 35, 38, 20, 51,
	53, 96, 0, 33, 36, 34, 32,
}

var yyTok1 = [...]int8{
	1,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:129
		{
			yylex.(*parser).parseResult = yyDollar[2].aststmts
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:132
		{
			yylex.(*parser).unexpected("", "")
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:137
		{
			s := yyDollar[1].aststmts
			s = append(s, yyDollar[2].node)
			yyVAL.aststmts = s
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:144
		{
			yyVAL.aststmts = ast.Stmts{yyDollar[1].node}
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:148
		{
			yyVAL.aststmts = ast.Stmts{yyDollar[1].node}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:150
		{
			yyVAL.aststmts = ast.Stmts{}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:152
		{
			s := yyDollar[1].aststmts
			s = append(s, yyDollar[2].node)
			yyVAL.aststmts = s
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:176
		{
			yyVAL.node = yylex.(*parser).newBreakStmt(yyDollar[1].item.Pos)
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:180
		{
			yyVAL.node = yylex.(*parser).newContinueStmt(yyDollar[1].item.Pos)
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:189
		{
			yyVAL.node = yylex.(*parser).newForInStmt(yyDollar[2].node, yyDollar[4].node, yyDollar[5].astblock, yyDollar[1].item, yyDollar[3].item)
		}
	case 32:
		yyDollar = yyS[yypt-7 : yypt+1]
//line gram.y:200
		{
			yyVAL.node = yylex.(*parser).newForStmt(yyDollar[2].node, yyDollar[4].node, yyDollar[6].node, yyDollar[7].astblock)
		}
	case 33:
		yyDollar = yyS[yypt-6 : yypt+1]
//line gram.y:202
		{
			yyVAL.node = yylex.(*parser).newForStmt(yyDollar[2].node, yyDollar[4].node, nil, yyDollar[6].astblock)
		}
	case 34:
		yyDollar = yyS[yypt-6 : yypt+1]
//line gram.y:204
		{
			yyVAL.node = yylex.(*parser).newForStmt(nil, yyDollar[3].node, yyDollar[5].node, yyDollar[6].astblock)
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:206
		{
			yyVAL.node = yylex.(*parser).newForStmt(nil, yyDollar[3].node, nil, yyDollar[5].astblock)
		}
	case 36:
		yyDollar = yyS[yypt-6 : yypt+1]
//line gram.y:209
		{
			yyVAL.node = yylex.(*parser).newForStmt(yyDollar[2].node, nil, yyDollar[5].node, yyDollar[6].astblock)
		}
	case 37:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:211
		{
			yyVAL.node = yylex.(*parser).newForStmt(yyDollar[2].node, nil, nil, yyDollar[5].astblock)
		}
	case 38:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:213
		{
			yyVAL.node = yylex.(*parser).newForStmt(nil, nil, yyDollar[4].node, yyDollar[5].astblock)
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:215
		{
			yyVAL.node = yylex.(*parser).newForStmt(nil, nil, nil, yyDollar[4].astblock)
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:219
		{
			yyVAL.node = yylex.(*parser).newIfElifStmt(yyDollar[1].iflist)
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:223
		{
			yyVAL.node = yylex.(*parser).newIfElifelseStmt(yyDollar[1].iflist, yyDollar[2].item, yyDollar[3].astblock)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:229
		{
			yyVAL.ifitem = yylex.(*parser).newIfElem(yyDollar[1].item, yyDollar[2].node, yyDollar[3].astblock)
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:233
		{
			yyVAL.iflist = []*ast.IfStmtElem{yyDollar[1].ifitem}
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:235
		{
			yyVAL.iflist = append(yyDollar[1].iflist, yyDollar[2].ifitem)
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:239
		{
			yyVAL.ifitem = yylex.(*parser).newIfElem(yyDollar[1].item, yyDollar[2].node, yyDollar[3].astblock)
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:245
		{
			yyVAL.astblock = yylex.(*parser).newBlockStmt(yyDollar[1].item, yyDollar[2].aststmts, yyDollar[3].item)
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:249
		{
			yyVAL.astblock = yylex.(*parser).newBlockStmt(yyDollar[1].item, ast.Stmts{}, yyDollar[2].item)
		}
	case 49:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:254
		{
			yyVAL.node = yylex.(*parser).newCallExpr(yyDollar[1].node, yyDollar[3].nodes, yyDollar[2].item, yyDollar[4].item)
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:258
		{
			yyVAL.node = yylex.(*parser).newCallExpr(yyDollar[1].node, nil, yyDollar[2].item, yyDollar[3].item)
		}
	case 51:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:262
		{
			yyVAL.node = yylex.(*parser).newCallExpr(yyDollar[1].node, yyDollar[3].nodes, yyDollar[2].item, yyDollar[5].item)
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:266
		{
			yyVAL.node = yylex.(*parser).newCallExpr(yyDollar[1].node, nil, yyDollar[2].item, yyDollar[4].item)
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:273
		{
			yyVAL.nodes = append(yyVAL.nodes, yyDollar[3].node)
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:278
		{
			yyVAL.nodes = []*ast.Node{yyDollar[1].node}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:285
		{
			yyVAL.node = yylex.(*parser).newAssignmentExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:289
		{
			yyVAL.node = yylex.(*parser).newConditionalExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:291
		{
			yyVAL.node = yylex.(*parser).newConditionalExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:293
		{
			yyVAL.node = yylex.(*parser).newConditionalExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:295
		{
			yyVAL.node = yylex.(*parser).newConditionalExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:297
		{
			yyVAL.node = yylex.(*parser).newConditionalExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:299
		{
			yyVAL.node = yylex.(*parser).newConditionalExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:301
		{
			yyVAL.node = yylex.(*parser).newConditionalExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:303
		{
			yyVAL.node = yylex.(*parser).newConditionalExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:308
		{
			yyVAL.node = yylex.(*parser).newArithmeticExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:310
		{
			yyVAL.node = yylex.(*parser).newArithmeticExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:312
		{
			yyVAL.node = yylex.(*parser).newArithmeticExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:314
		{
			yyVAL.node = yylex.(*parser).newArithmeticExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:316
		{
			yyVAL.node = yylex.(*parser).newArithmeticExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[2].item)
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:321
		{
			yyVAL.node = yylex.(*parser).newParenExpr(yyDollar[1].item, yyDollar[2].node, yyDollar[3].item)
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:323
		{
			yyVAL.node = yylex.(*parser).newParenExpr(yyDollar[1].item, yyDollar[2].node, yyDollar[4].item)
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:332
		{
			yyVAL.node = yylex.(*parser).newIndexExpr(yyDollar[1].node, yyDollar[2].item, yyDollar[3].node, yyDollar[4].item)
		}
	case 78:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:335
		{
			yyVAL.node = yylex.(*parser).newIndexExpr(nil, yyDollar[2].item, yyDollar[3].node, yyDollar[4].item)
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:337
		{
			yyVAL.node = yylex.(*parser).newIndexExpr(yyDollar[1].node, yyDollar[2].item, yyDollar[3].node, yyDollar[4].item)
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:344
		{
			yyVAL.node = yylex.(*parser).newAttrExpr(yyDollar[1].node, yyDollar[3].node)
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:348
		{
			yyVAL.node = yylex.(*parser).newAttrExpr(yyDollar[1].node, yyDollar[3].node)
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:352
		{
			yyVAL.node = yylex.(*parser).newAttrExpr(yyDollar[1].node, yyDollar[3].node)
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:356
		{
			yyVAL.node = yylex.(*parser).newAttrExpr(yyDollar[1].node, yyDollar[3].node)
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:360
		{
			yyVAL.node = yylex.(*parser).newAttrExpr(yyDollar[1].node, yyDollar[3].node)
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:364
		{
			yyVAL.node = yylex.(*parser).newAttrExpr(yyDollar[1].node, yyDollar[3].node)
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:371
		{
			yyVAL.node = yylex.(*parser).newListInitEndExpr(yyVAL.node, yyDollar[2].item.Pos)
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:375
		{
			yyVAL.node = yylex.(*parser).newListInitEndExpr(yyVAL.node, yyDollar[2].item.Pos)
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:379
		{
			yyVAL.node = yylex.(*parser).newListInitStartExpr(yyDollar[1].item.Pos)
			yyVAL.node = yylex.(*parser).newListInitEndExpr(yyVAL.node, yyDollar[2].item.Pos)

		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:387
		{
			yyVAL.node = yylex.(*parser).newListInitStartExpr(yyDollar[1].item.Pos)
			yyVAL.node = yylex.(*parser).newListInitAppendExpr(yyVAL.node, yyDollar[2].node)
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:392
		{
			yyVAL.node = yylex.(*parser).newListInitAppendExpr(yyVAL.node, yyDollar[3].node)
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:400
		{
			yyVAL.node = yylex.(*parser).newMapInitEndExpr(yyVAL.node, yyDollar[2].item.Pos)
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:404
		{
			yyVAL.node = yylex.(*parser).newMapInitEndExpr(yyVAL.node, yyDollar[3].item.Pos)
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:408
		{
			yyVAL.node = yylex.(*parser).newMapInitStartExpr(yyDollar[1].astblock.LBracePos.Pos)
			yyVAL.node = yylex.(*parser).newMapInitEndExpr(yyVAL.node, yyDollar[1].astblock.RBracePos.Pos)
		}
	case 95:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:415
		{
			yyVAL.node = yylex.(*parser).newMapInitStartExpr(yyDollar[1].item.Pos)
			yyVAL.node = yylex.(*parser).newMapInitAppendExpr(yyVAL.node, yyDollar[2].node, yyDollar[4].node)
		}
	case 96:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:420
		{
			yyVAL.node = yylex.(*parser).newMapInitAppendExpr(yyDollar[1].node, yyDollar[3].node, yyDollar[5].node)
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:444
		{
			yyVAL.node = yylex.(*parser).newBoolLiteral(yyDollar[1].item.Pos, true)
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:446
		{
			yyVAL.node = yylex.(*parser).newBoolLiteral(yyDollar[1].item.Pos, false)
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:451
		{
			yyDollar[1].item.Val = yylex.(*parser).unquoteString(yyDollar[1].item.Val)
			yyVAL.node = yylex.(*parser).newStringLiteral(yyDollar[1].item)
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:456
		{
			yyDollar[1].item.Val = yylex.(*parser).unquoteMultilineString(yyDollar[1].item.Val)
			yyVAL.node = yylex.(*parser).newStringLiteral(yyDollar[1].item)
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:464
		{
			yyVAL.node = yylex.(*parser).newNilLiteral(yyDollar[1].item.Pos)
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:466
		{
			yyVAL.node = yylex.(*parser).newNilLiteral(yyDollar[1].item.Pos)
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:472
		{
			yyVAL.node = yylex.(*parser).newNumberLiteral(yyDollar[1].item)
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:474
		{
			num := yylex.(*parser).newNumberLiteral(yyDollar[2].item)
			switch yyDollar[1].item.Typ {
			case ADD: // pass
			case SUB:
				if num.NodeType == ast.TypeFloatLiteral {
					num.FloatLiteral.Val = -num.FloatLiteral.Val
					num.FloatLiteral.Start = yylex.(*parser).posCache.LnCol(yyDollar[1].item.Pos)
				} else {
					num.IntegerLiteral.Val = -num.IntegerLiteral.Val
					num.IntegerLiteral.Start = yylex.(*parser).posCache.LnCol(yyDollar[1].item.Pos)

				}
			}
			yyVAL.node = num
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:493
		{
			yyVAL.node = yylex.(*parser).newIdentifierLiteral(yyDollar[1].item)
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:497
		{
			yyDollar[1].item.Val = yylex.(*parser).unquoteString(yyDollar[1].item.Val)
			yyVAL.node = yylex.(*parser).newIdentifierLiteral(yyDollar[1].item)
		}
	}
	goto yystack /* stack new state and value */
}
