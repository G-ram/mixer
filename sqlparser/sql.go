//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
import "bytes"

func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func SetAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func ForceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

var (
	SHARE        = []byte("share")
	MODE         = []byte("mode")
	IF_BYTES     = []byte("if")
	VALUES_BYTES = []byte("values")
)

//line sql.y:31
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	smTableExpr SimpleTableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
	tuple       Tuple
	valExprs    ValExprs
	values      Values
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr
}

const LEX_ERROR = 57346
const SELECT = 57347
const INSERT = 57348
const UPDATE = 57349
const DELETE = 57350
const FROM = 57351
const WHERE = 57352
const GROUP = 57353
const HAVING = 57354
const ORDER = 57355
const BY = 57356
const LIMIT = 57357
const FOR = 57358
const ALL = 57359
const DISTINCT = 57360
const AS = 57361
const EXISTS = 57362
const IN = 57363
const IS = 57364
const LIKE = 57365
const BETWEEN = 57366
const NULL = 57367
const ASC = 57368
const DESC = 57369
const VALUES = 57370
const INTO = 57371
const DUPLICATE = 57372
const KEY = 57373
const DEFAULT = 57374
const SET = 57375
const LOCK = 57376
const ID = 57377
const STRING = 57378
const NUMBER = 57379
const VALUE_ARG = 57380
const COMMENT = 57381
const LE = 57382
const GE = 57383
const NE = 57384
const NULL_SAFE_EQUAL = 57385
const UNION = 57386
const MINUS = 57387
const EXCEPT = 57388
const INTERSECT = 57389
const JOIN = 57390
const STRAIGHT_JOIN = 57391
const LEFT = 57392
const RIGHT = 57393
const INNER = 57394
const OUTER = 57395
const CROSS = 57396
const NATURAL = 57397
const USE = 57398
const FORCE = 57399
const ON = 57400
const AND = 57401
const OR = 57402
const NOT = 57403
const UNARY = 57404
const CASE = 57405
const WHEN = 57406
const THEN = 57407
const ELSE = 57408
const END = 57409
const BEGIN = 57410
const COMMIT = 57411
const ROLLBACK = 57412
const NAMES = 57413
const REPLACE = 57414
const ADMIN = 57415
const SHOW = 57416
const DATABASES = 57417
const TABLES = 57418
const PROXY = 57419
const VARIABLES = 57420
const FULL = 57421
const COLUMNS = 57422
const CREATE = 57423
const ALTER = 57424
const DROP = 57425
const RENAME = 57426
const TABLE = 57427
const INDEX = 57428
const VIEW = 57429
const TO = 57430
const IGNORE = 57431
const IF = 57432
const UNIQUE = 57433
const USING = 57434

var yyToknames = []string{
	"LEX_ERROR",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"IN",
	"IS",
	"LIKE",
	"BETWEEN",
	"NULL",
	"ASC",
	"DESC",
	"VALUES",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"COMMENT",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"'('",
	"'='",
	"'<'",
	"'>'",
	"'~'",
	"UNION",
	"MINUS",
	"EXCEPT",
	"INTERSECT",
	"','",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"AND",
	"OR",
	"NOT",
	"'&'",
	"'|'",
	"'^'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'.'",
	"UNARY",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"BEGIN",
	"COMMIT",
	"ROLLBACK",
	"NAMES",
	"REPLACE",
	"ADMIN",
	"SHOW",
	"DATABASES",
	"TABLES",
	"PROXY",
	"VARIABLES",
	"FULL",
	"COLUMNS",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 224
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 600

var yyAct = []int{

	114, 330, 151, 397, 73, 105, 365, 193, 111, 141,
	282, 238, 322, 273, 275, 122, 75, 216, 406, 112,
	203, 208, 100, 101, 406, 34, 35, 36, 37, 168,
	169, 91, 87, 62, 194, 3, 406, 163, 80, 298,
	299, 300, 301, 302, 77, 303, 304, 82, 328, 44,
	84, 46, 76, 163, 88, 47, 289, 144, 163, 117,
	97, 376, 375, 236, 121, 236, 83, 127, 64, 106,
	374, 81, 265, 408, 78, 118, 119, 120, 51, 407,
	347, 349, 140, 109, 267, 98, 78, 125, 356, 94,
	148, 405, 355, 351, 143, 49, 152, 50, 150, 155,
	95, 274, 161, 327, 165, 309, 108, 160, 317, 167,
	123, 124, 195, 315, 191, 192, 196, 128, 266, 274,
	235, 320, 348, 52, 53, 54, 137, 154, 181, 182,
	183, 202, 77, 199, 70, 77, 206, 74, 212, 211,
	76, 132, 139, 76, 126, 168, 169, 255, 213, 168,
	169, 210, 63, 134, 371, 323, 161, 285, 226, 147,
	358, 233, 234, 373, 323, 372, 106, 244, 212, 86,
	227, 136, 345, 248, 246, 247, 253, 254, 242, 257,
	258, 259, 260, 261, 262, 263, 264, 243, 229, 256,
	245, 249, 56, 58, 59, 57, 61, 222, 156, 344,
	106, 106, 384, 385, 130, 77, 77, 133, 343, 278,
	209, 134, 236, 76, 280, 284, 220, 286, 341, 223,
	269, 271, 382, 342, 89, 149, 339, 287, 281, 77,
	360, 340, 241, 291, 292, 293, 209, 76, 277, 294,
	162, 240, 290, 392, 176, 177, 178, 179, 180, 181,
	182, 183, 308, 296, 311, 312, 242, 228, 18, 295,
	204, 391, 277, 179, 180, 181, 182, 183, 205, 129,
	310, 205, 390, 106, 34, 35, 36, 37, 153, 134,
	321, 219, 221, 218, 163, 200, 319, 198, 241, 197,
	99, 326, 329, 316, 63, 307, 325, 240, 166, 78,
	270, 352, 117, 403, 350, 334, 333, 121, 337, 338,
	127, 306, 242, 242, 63, 225, 354, 104, 118, 119,
	120, 404, 379, 357, 224, 207, 109, 71, 145, 77,
	125, 362, 142, 138, 363, 366, 135, 361, 298, 299,
	300, 301, 302, 367, 303, 304, 85, 381, 131, 108,
	359, 18, 90, 123, 124, 102, 377, 69, 314, 92,
	128, 378, 176, 177, 178, 179, 180, 181, 182, 183,
	231, 158, 93, 161, 276, 387, 410, 389, 388, 386,
	380, 67, 232, 159, 394, 366, 214, 126, 396, 395,
	268, 398, 398, 398, 77, 399, 400, 146, 401, 18,
	117, 65, 76, 331, 370, 121, 332, 411, 127, 283,
	369, 412, 336, 413, 117, 104, 118, 119, 120, 121,
	209, 409, 127, 38, 109, 96, 72, 393, 125, 78,
	118, 119, 120, 250, 18, 251, 252, 18, 109, 39,
	60, 230, 125, 40, 41, 42, 43, 108, 157, 17,
	16, 123, 124, 102, 55, 15, 14, 121, 128, 13,
	127, 108, 12, 215, 45, 123, 124, 78, 118, 119,
	120, 121, 128, 288, 127, 217, 153, 48, 79, 279,
	125, 78, 118, 119, 120, 126, 18, 19, 20, 21,
	153, 402, 383, 364, 125, 368, 335, 318, 201, 126,
	272, 116, 113, 123, 124, 115, 171, 175, 173, 174,
	128, 324, 110, 170, 22, 107, 346, 123, 124, 239,
	297, 237, 103, 305, 128, 187, 188, 189, 190, 164,
	184, 185, 186, 66, 33, 353, 68, 126, 176, 177,
	178, 179, 180, 181, 182, 183, 11, 10, 9, 8,
	7, 126, 172, 176, 177, 178, 179, 180, 181, 182,
	183, 6, 5, 4, 27, 28, 29, 2, 30, 32,
	31, 1, 0, 0, 0, 0, 0, 23, 24, 26,
	25, 313, 0, 0, 176, 177, 178, 179, 180, 181,
	182, 183, 176, 177, 178, 179, 180, 181, 182, 183,
}
var yyPact = []int{

	481, -1000, -1000, 225, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -51, -7, -22, 23, -1000, -1000, -1000,
	-1000, 102, 259, 429, 384, -1000, -1000, -1000, 363, -1000,
	328, 292, 417, 51, -67, -30, 259, -1000, -34, 259,
	-1000, 311, -73, 259, -73, 323, 349, 349, 416, 259,
	-10, -1000, 246, -1000, -1000, -1000, 380, -1000, 230, 292,
	315, 65, 292, 158, 301, -1000, 126, -1000, 50, 298,
	75, 259, -1000, 297, -1000, -46, 293, 377, 95, 259,
	292, -1000, 39, 446, -1000, 349, 446, 416, 362, 446,
	231, -1000, -1000, 279, 33, 84, 485, -1000, 39, 394,
	-1000, -1000, -1000, 446, 245, 243, -1000, 241, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 446, -1000,
	227, 264, 290, 410, 264, -1000, 446, 259, -1000, 366,
	-90, -1000, 184, -1000, 289, -1000, -1000, 280, -1000, 224,
	84, 485, 524, 432, -1000, 524, 349, 361, 446, 446,
	12, 524, 197, 380, -1000, -1000, 259, 117, 39, 39,
	446, 234, 412, 446, 446, 122, 446, 446, 446, 446,
	446, 446, 446, 446, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -36, 10, -24, 485, -1000, 282, 380, -1000,
	429, 22, 524, 346, 264, 264, 226, -1000, 396, 39,
	-1000, 524, -1000, -1000, -1000, 93, 259, -1000, -47, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 346, 264, -1000,
	-1000, 446, 446, 524, 524, -1000, 446, 200, 284, 276,
	253, 29, -1000, -1000, -1000, -1000, -1000, -1000, 524, -1000,
	234, 446, 446, 524, 516, -1000, 333, 192, 192, 192,
	55, 55, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 5,
	380, 0, 40, -1000, 39, 91, 234, 225, 100, -5,
	-1000, 396, 388, 392, 84, 271, -1000, -1000, 270, -1000,
	-1000, 158, 524, 524, 524, 401, 197, 197, -1000, -1000,
	172, 164, 154, 145, 118, 18, -1000, 269, -15, 266,
	-1000, 524, 470, 446, -1000, -1000, -16, -1000, 6, -1000,
	446, 80, -1000, 320, 177, -1000, -1000, -1000, 264, 388,
	-1000, 446, 446, -1000, -1000, 398, 390, 284, 90, -1000,
	111, -1000, 109, -1000, -1000, -1000, -1000, -31, -39, -40,
	-1000, -1000, -1000, 446, 524, -1000, -1000, 524, 446, 291,
	234, -1000, -1000, 294, 169, -1000, 176, -1000, 396, 39,
	446, 39, -1000, -1000, 228, 217, 199, 524, 524, 420,
	-1000, 446, 446, -1000, -1000, -1000, 388, 84, 159, 84,
	259, 259, 259, 264, 524, -1000, 287, -17, -1000, -29,
	-35, 158, -1000, 414, 355, -1000, 259, -1000, -1000, -1000,
	259, -1000, 259, -1000,
}
var yyPgo = []int{

	0, 571, 567, 34, 563, 562, 561, 550, 549, 548,
	547, 546, 423, 536, 534, 533, 22, 23, 529, 523,
	522, 521, 11, 520, 519, 134, 516, 3, 21, 5,
	515, 513, 14, 512, 2, 19, 7, 511, 505, 15,
	502, 8, 501, 500, 13, 498, 497, 496, 495, 10,
	493, 6, 492, 1, 491, 20, 479, 12, 4, 16,
	169, 478, 477, 475, 473, 464, 463, 0, 9, 462,
	459, 456, 455, 450, 449, 100, 31, 448, 441, 440,
	439,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	4, 4, 72, 72, 5, 6, 7, 7, 69, 70,
	71, 74, 77, 77, 78, 78, 78, 79, 79, 73,
	73, 73, 73, 73, 8, 8, 8, 9, 9, 9,
	10, 11, 11, 11, 80, 12, 13, 13, 14, 14,
	14, 14, 14, 15, 15, 16, 16, 17, 17, 17,
	20, 20, 18, 18, 18, 21, 21, 22, 22, 22,
	22, 19, 19, 19, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 24, 24, 24, 25, 25, 26, 26,
	26, 26, 27, 27, 28, 28, 76, 76, 76, 75,
	75, 29, 29, 29, 29, 29, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 31, 31, 31, 31,
	31, 31, 31, 32, 32, 37, 37, 35, 35, 39,
	36, 36, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 38,
	38, 40, 40, 40, 42, 45, 45, 43, 43, 44,
	46, 46, 41, 41, 33, 33, 33, 33, 47, 47,
	48, 48, 49, 49, 50, 50, 51, 52, 52, 52,
	53, 53, 53, 54, 54, 54, 55, 55, 56, 56,
	57, 57, 58, 58, 59, 60, 60, 61, 61, 62,
	62, 63, 63, 63, 63, 63, 64, 64, 65, 65,
	66, 66, 67, 68,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 12, 3,
	7, 7, 6, 6, 8, 7, 3, 4, 1, 1,
	1, 5, 2, 2, 0, 2, 2, 0, 1, 3,
	3, 4, 5, 5, 5, 8, 4, 6, 7, 4,
	5, 4, 5, 5, 0, 2, 0, 2, 1, 2,
	1, 1, 1, 0, 1, 1, 3, 1, 2, 3,
	1, 1, 0, 1, 2, 1, 3, 3, 3, 3,
	5, 0, 1, 2, 1, 1, 2, 3, 2, 3,
	2, 2, 2, 1, 3, 1, 1, 3, 0, 5,
	5, 5, 1, 3, 0, 2, 0, 2, 2, 0,
	2, 1, 3, 3, 2, 3, 3, 3, 4, 3,
	4, 5, 6, 3, 4, 2, 1, 1, 1, 1,
	1, 1, 1, 2, 1, 1, 3, 3, 1, 3,
	1, 3, 1, 1, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 3, 4, 5, 4, 1, 1,
	1, 1, 1, 1, 5, 0, 1, 1, 2, 4,
	0, 2, 1, 3, 1, 1, 1, 1, 0, 3,
	0, 2, 0, 3, 1, 3, 2, 0, 1, 1,
	0, 2, 4, 0, 2, 4, 0, 3, 1, 3,
	0, 5, 1, 3, 3, 0, 2, 0, 3, 0,
	1, 1, 1, 1, 1, 1, 0, 1, 0, 1,
	0, 2, 1, 0,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, -72, -73, -74, 5, 6,
	7, 8, 33, 96, 97, 99, 98, 83, 84, 85,
	87, 89, 88, -14, 49, 50, 51, 52, -12, -80,
	-12, -12, -12, -12, 100, -65, 102, 106, -62, 102,
	104, 100, 100, 101, 102, -12, 90, 93, 91, 92,
	-79, 94, -67, 35, -3, 17, -15, 18, -13, 29,
	-25, 35, 9, -58, 86, -59, -41, -67, 35, -61,
	105, 101, -67, 100, -67, 35, -60, 105, -67, -60,
	29, -76, 10, 23, -76, -75, 9, -67, 95, 44,
	-16, -17, 73, -20, 35, -29, -34, -30, 67, 44,
	-33, -41, -35, -40, -67, -38, -42, 20, 36, 37,
	38, 25, -39, 71, 72, 48, 105, 28, 78, 39,
	-25, 33, 76, -25, 53, 35, 45, 76, 35, 67,
	-67, -68, 35, -68, 103, 35, 20, 64, -67, -25,
	-29, -34, -34, 44, -76, -34, -75, -77, 9, 21,
	-36, -34, 9, 53, -18, -67, 19, 76, 65, 66,
	-31, 21, 67, 23, 24, 22, 68, 69, 70, 71,
	72, 73, 74, 75, 45, 46, 47, 40, 41, 42,
	43, -29, -29, -36, -3, -34, -34, 44, 44, -39,
	44, -45, -34, -55, 33, 44, -58, 35, -28, 10,
	-59, -34, -67, -68, 20, -66, 107, -63, 99, 97,
	32, 98, 13, 35, 35, 35, -68, -55, 33, -76,
	-78, 9, 21, -34, -34, 108, 53, -21, -22, -24,
	44, 35, -39, -17, -67, 73, -29, -29, -34, -35,
	21, 23, 24, -34, -34, 25, 67, -34, -34, -34,
	-34, -34, -34, -34, -34, 108, 108, 108, 108, -16,
	18, -16, -43, -44, 79, -32, 28, -3, -58, -56,
	-41, -28, -49, 13, -29, 64, -67, -68, -64, 103,
	-32, -58, -34, -34, -34, -28, 53, -23, 54, 55,
	56, 57, 58, 60, 61, -19, 35, 19, -22, 76,
	-35, -34, -34, 65, 25, 108, -16, 108, -46, -44,
	81, -29, -57, 64, -37, -35, -57, 108, 53, -49,
	-53, 15, 14, 35, 35, -47, 11, -22, -22, 54,
	59, 54, 59, 54, 54, 54, -26, 62, 104, 63,
	35, 108, 35, 65, -34, 108, 82, -34, 80, 30,
	53, -41, -53, -34, -50, -51, -34, -68, -48, 12,
	14, 64, 54, 54, 101, 101, 101, -34, -34, 31,
	-35, 53, 53, -52, 26, 27, -49, -29, -36, -29,
	44, 44, 44, 7, -34, -51, -53, -27, -67, -27,
	-27, -58, -54, 16, 34, 108, 53, 108, 108, 7,
	21, -67, -67, -67,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 54, 54,
	54, 54, 54, 218, 209, 0, 0, 28, 29, 30,
	54, 37, 0, 0, 58, 60, 61, 62, 63, 56,
	0, 0, 0, 0, 207, 0, 0, 219, 0, 0,
	210, 0, 205, 0, 205, 0, 106, 106, 109, 0,
	0, 38, 0, 222, 19, 59, 0, 64, 55, 0,
	0, 96, 0, 26, 0, 202, 0, 172, 222, 0,
	0, 0, 223, 0, 223, 0, 0, 0, 0, 0,
	0, 39, 0, 0, 40, 106, 0, 109, 0, 0,
	17, 65, 67, 72, 222, 70, 71, 111, 0, 0,
	142, 143, 144, 0, 172, 0, 158, 0, 174, 175,
	176, 177, 138, 161, 162, 163, 159, 160, 165, 57,
	196, 0, 0, 104, 0, 27, 0, 0, 223, 0,
	220, 46, 0, 49, 0, 51, 206, 0, 223, 196,
	107, 0, 108, 0, 41, 110, 106, 34, 0, 0,
	0, 140, 0, 0, 68, 73, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 126, 127, 128, 129, 130, 131,
	132, 114, 0, 0, 0, 140, 153, 0, 0, 125,
	0, 0, 166, 0, 0, 0, 104, 97, 182, 0,
	203, 204, 173, 44, 208, 0, 0, 223, 216, 211,
	212, 213, 214, 215, 50, 52, 53, 0, 0, 42,
	43, 0, 0, 32, 33, 31, 0, 104, 75, 81,
	0, 93, 95, 66, 74, 69, 112, 113, 116, 117,
	0, 0, 0, 119, 0, 123, 0, 145, 146, 147,
	148, 149, 150, 151, 152, 115, 137, 139, 154, 0,
	0, 0, 170, 167, 0, 200, 0, 134, 200, 0,
	198, 182, 190, 0, 105, 0, 221, 47, 0, 217,
	22, 23, 35, 36, 141, 178, 0, 0, 84, 85,
	0, 0, 0, 0, 0, 98, 82, 0, 0, 0,
	118, 120, 0, 0, 124, 155, 0, 157, 0, 168,
	0, 0, 20, 0, 133, 135, 21, 197, 0, 190,
	25, 0, 0, 223, 48, 180, 0, 76, 79, 86,
	0, 88, 0, 90, 91, 92, 77, 0, 0, 0,
	83, 78, 94, 0, 121, 156, 164, 171, 0, 0,
	0, 199, 24, 191, 183, 184, 187, 45, 182, 0,
	0, 0, 87, 89, 0, 0, 0, 122, 169, 0,
	136, 0, 0, 186, 188, 189, 190, 181, 179, 80,
	0, 0, 0, 0, 192, 185, 193, 0, 102, 0,
	0, 201, 18, 0, 0, 99, 0, 100, 101, 194,
	0, 103, 0, 195,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 75, 68, 3,
	44, 108, 73, 71, 53, 72, 76, 74, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	46, 45, 47, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 70, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 69, 3, 48,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 49, 50, 51, 52, 54, 55, 56, 57,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	77, 78, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106,
	107,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
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

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
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
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
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
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
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

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line sql.y:171
		{
			SetParseTree(yylex, yyS[yypt-0].statement)
		}
	case 2:
		//line sql.y:177
		{
			yyVAL.statement = yyS[yypt-0].selStmt
		}
	case 3:
		yyVAL.statement = yyS[yypt-0].statement
	case 4:
		yyVAL.statement = yyS[yypt-0].statement
	case 5:
		yyVAL.statement = yyS[yypt-0].statement
	case 6:
		yyVAL.statement = yyS[yypt-0].statement
	case 7:
		yyVAL.statement = yyS[yypt-0].statement
	case 8:
		yyVAL.statement = yyS[yypt-0].statement
	case 9:
		yyVAL.statement = yyS[yypt-0].statement
	case 10:
		yyVAL.statement = yyS[yypt-0].statement
	case 11:
		yyVAL.statement = yyS[yypt-0].statement
	case 12:
		yyVAL.statement = yyS[yypt-0].statement
	case 13:
		yyVAL.statement = yyS[yypt-0].statement
	case 14:
		yyVAL.statement = yyS[yypt-0].statement
	case 15:
		yyVAL.statement = yyS[yypt-0].statement
	case 16:
		yyVAL.statement = yyS[yypt-0].statement
	case 17:
		//line sql.y:197
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyS[yypt-2].bytes2), Distinct: yyS[yypt-1].str, SelectExprs: yyS[yypt-0].selectExprs}
		}
	case 18:
		//line sql.y:201
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-10].bytes2), Distinct: yyS[yypt-9].str, SelectExprs: yyS[yypt-8].selectExprs, From: yyS[yypt-6].tableExprs, Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 19:
		//line sql.y:205
		{
			yyVAL.selStmt = &Union{Type: yyS[yypt-1].str, Left: yyS[yypt-2].selStmt, Right: yyS[yypt-0].selStmt}
		}
	case 20:
		//line sql.y:212
		{
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 21:
		//line sql.y:216
		{
			cols := make(Columns, 0, len(yyS[yypt-1].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-1].updateExprs))
			for _, col := range yyS[yypt-1].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 22:
		//line sql.y:228
		{
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: yyS[yypt-1].columns, Rows: yyS[yypt-0].insRows}
		}
	case 23:
		//line sql.y:232
		{
			cols := make(Columns, 0, len(yyS[yypt-0].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-0].updateExprs))
			for _, col := range yyS[yypt-0].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 24:
		//line sql.y:245
		{
			yyVAL.statement = &Update{Comments: Comments(yyS[yypt-6].bytes2), Table: yyS[yypt-5].tableName, Exprs: yyS[yypt-3].updateExprs, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 25:
		//line sql.y:251
		{
			yyVAL.statement = &Delete{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 26:
		//line sql.y:257
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-1].bytes2), Exprs: yyS[yypt-0].updateExprs}
		}
	case 27:
		//line sql.y:261
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal(yyS[yypt-0].bytes)}}}
		}
	case 28:
		//line sql.y:267
		{
			yyVAL.statement = &Begin{}
		}
	case 29:
		//line sql.y:273
		{
			yyVAL.statement = &Commit{}
		}
	case 30:
		//line sql.y:279
		{
			yyVAL.statement = &Rollback{}
		}
	case 31:
		//line sql.y:285
		{
			yyVAL.statement = &Admin{Name: yyS[yypt-3].bytes, Values: yyS[yypt-1].valExprs}
		}
	case 32:
		//line sql.y:291
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 33:
		//line sql.y:295
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 34:
		//line sql.y:300
		{
			yyVAL.valExpr = nil
		}
	case 35:
		//line sql.y:304
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 36:
		//line sql.y:308
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 37:
		//line sql.y:313
		{
			yyVAL.str = AST_SHOW_NO_MOD
		}
	case 38:
		//line sql.y:317
		{
			yyVAL.str = AST_SHOW_FULL
		}
	case 39:
		//line sql.y:324
		{
			yyVAL.statement = &Show{Section: "databases", LikeOrWhere: yyS[yypt-0].expr}
		}
	case 40:
		//line sql.y:328
		{
			yyVAL.statement = &Show{Section: "variables", LikeOrWhere: yyS[yypt-0].expr}
		}
	case 41:
		//line sql.y:332
		{
			yyVAL.statement = &Show{Section: "tables", From: yyS[yypt-1].valExpr, LikeOrWhere: yyS[yypt-0].expr}
		}
	case 42:
		//line sql.y:336
		{
			yyVAL.statement = &Show{Section: "proxy", Key: string(yyS[yypt-2].bytes), From: yyS[yypt-1].valExpr, LikeOrWhere: yyS[yypt-0].expr}
		}
	case 43:
		//line sql.y:340
		{
			yyVAL.statement = &Show{Section: "columns", From: yyS[yypt-1].valExpr, Modifier: yyS[yypt-3].str, DBFilter: yyS[yypt-0].valExpr}
		}
	case 44:
		//line sql.y:346
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 45:
		//line sql.y:350
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 46:
		//line sql.y:355
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 47:
		//line sql.y:361
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 48:
		//line sql.y:365
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 49:
		//line sql.y:370
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 50:
		//line sql.y:376
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 51:
		//line sql.y:382
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-0].bytes}
		}
	case 52:
		//line sql.y:386
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 53:
		//line sql.y:391
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-1].bytes}
		}
	case 54:
		//line sql.y:396
		{
			SetAllowComments(yylex, true)
		}
	case 55:
		//line sql.y:400
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 56:
		//line sql.y:406
		{
			yyVAL.bytes2 = nil
		}
	case 57:
		//line sql.y:410
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 58:
		//line sql.y:416
		{
			yyVAL.str = AST_UNION
		}
	case 59:
		//line sql.y:420
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 60:
		//line sql.y:424
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 61:
		//line sql.y:428
		{
			yyVAL.str = AST_EXCEPT
		}
	case 62:
		//line sql.y:432
		{
			yyVAL.str = AST_INTERSECT
		}
	case 63:
		//line sql.y:437
		{
			yyVAL.str = ""
		}
	case 64:
		//line sql.y:441
		{
			yyVAL.str = AST_DISTINCT
		}
	case 65:
		//line sql.y:447
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 66:
		//line sql.y:451
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 67:
		//line sql.y:457
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 68:
		//line sql.y:461
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 69:
		//line sql.y:465
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 70:
		//line sql.y:471
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 71:
		//line sql.y:475
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 72:
		//line sql.y:480
		{
			yyVAL.bytes = nil
		}
	case 73:
		//line sql.y:484
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 74:
		//line sql.y:488
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 75:
		//line sql.y:494
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 76:
		//line sql.y:498
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 77:
		//line sql.y:504
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 78:
		//line sql.y:508
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 79:
		//line sql.y:512
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 80:
		//line sql.y:516
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 81:
		//line sql.y:521
		{
			yyVAL.bytes = nil
		}
	case 82:
		//line sql.y:525
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 83:
		//line sql.y:529
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 84:
		//line sql.y:535
		{
			yyVAL.str = AST_JOIN
		}
	case 85:
		//line sql.y:539
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 86:
		//line sql.y:543
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 87:
		//line sql.y:547
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 88:
		//line sql.y:551
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 89:
		//line sql.y:555
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 90:
		//line sql.y:559
		{
			yyVAL.str = AST_JOIN
		}
	case 91:
		//line sql.y:563
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 92:
		//line sql.y:567
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 93:
		//line sql.y:573
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 94:
		//line sql.y:577
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 95:
		//line sql.y:581
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 96:
		//line sql.y:587
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 97:
		//line sql.y:591
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 98:
		//line sql.y:596
		{
			yyVAL.indexHints = nil
		}
	case 99:
		//line sql.y:600
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 100:
		//line sql.y:604
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 101:
		//line sql.y:608
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 102:
		//line sql.y:614
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 103:
		//line sql.y:618
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 104:
		//line sql.y:623
		{
			yyVAL.boolExpr = nil
		}
	case 105:
		//line sql.y:627
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 106:
		//line sql.y:632
		{
			yyVAL.expr = nil
		}
	case 107:
		//line sql.y:636
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 108:
		//line sql.y:640
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 109:
		//line sql.y:645
		{
			yyVAL.valExpr = nil
		}
	case 110:
		//line sql.y:649
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 111:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 112:
		//line sql.y:656
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 113:
		//line sql.y:660
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 114:
		//line sql.y:664
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 115:
		//line sql.y:668
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 116:
		//line sql.y:674
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 117:
		//line sql.y:678
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].tuple}
		}
	case 118:
		//line sql.y:682
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].tuple}
		}
	case 119:
		//line sql.y:686
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 120:
		//line sql.y:690
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 121:
		//line sql.y:694
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 122:
		//line sql.y:698
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 123:
		//line sql.y:702
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 124:
		//line sql.y:706
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 125:
		//line sql.y:710
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 126:
		//line sql.y:716
		{
			yyVAL.str = AST_EQ
		}
	case 127:
		//line sql.y:720
		{
			yyVAL.str = AST_LT
		}
	case 128:
		//line sql.y:724
		{
			yyVAL.str = AST_GT
		}
	case 129:
		//line sql.y:728
		{
			yyVAL.str = AST_LE
		}
	case 130:
		//line sql.y:732
		{
			yyVAL.str = AST_GE
		}
	case 131:
		//line sql.y:736
		{
			yyVAL.str = AST_NE
		}
	case 132:
		//line sql.y:740
		{
			yyVAL.str = AST_NSE
		}
	case 133:
		//line sql.y:746
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 134:
		//line sql.y:750
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 135:
		//line sql.y:756
		{
			yyVAL.values = Values{yyS[yypt-0].tuple}
		}
	case 136:
		//line sql.y:760
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].tuple)
		}
	case 137:
		//line sql.y:766
		{
			yyVAL.tuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 138:
		//line sql.y:770
		{
			yyVAL.tuple = yyS[yypt-0].subquery
		}
	case 139:
		//line sql.y:776
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 140:
		//line sql.y:782
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 141:
		//line sql.y:786
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 142:
		//line sql.y:792
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 143:
		//line sql.y:796
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 144:
		//line sql.y:800
		{
			yyVAL.valExpr = yyS[yypt-0].tuple
		}
	case 145:
		//line sql.y:804
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 146:
		//line sql.y:808
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 147:
		//line sql.y:812
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 148:
		//line sql.y:816
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 149:
		//line sql.y:820
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 150:
		//line sql.y:824
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 151:
		//line sql.y:828
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 152:
		//line sql.y:832
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 153:
		//line sql.y:836
		{
			if num, ok := yyS[yypt-0].valExpr.(NumVal); ok {
				switch yyS[yypt-1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
			}
		}
	case 154:
		//line sql.y:851
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 155:
		//line sql.y:855
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 156:
		//line sql.y:859
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 157:
		//line sql.y:863
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 158:
		//line sql.y:867
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 159:
		//line sql.y:873
		{
			yyVAL.bytes = IF_BYTES
		}
	case 160:
		//line sql.y:877
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 161:
		//line sql.y:883
		{
			yyVAL.byt = AST_UPLUS
		}
	case 162:
		//line sql.y:887
		{
			yyVAL.byt = AST_UMINUS
		}
	case 163:
		//line sql.y:891
		{
			yyVAL.byt = AST_TILDA
		}
	case 164:
		//line sql.y:897
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 165:
		//line sql.y:902
		{
			yyVAL.valExpr = nil
		}
	case 166:
		//line sql.y:906
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 167:
		//line sql.y:912
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 168:
		//line sql.y:916
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 169:
		//line sql.y:922
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 170:
		//line sql.y:927
		{
			yyVAL.valExpr = nil
		}
	case 171:
		//line sql.y:931
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 172:
		//line sql.y:937
		{
			yyVAL.colName = &ColName{Name: yyS[yypt-0].bytes}
		}
	case 173:
		//line sql.y:941
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 174:
		//line sql.y:947
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 175:
		//line sql.y:951
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 176:
		//line sql.y:955
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 177:
		//line sql.y:959
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 178:
		//line sql.y:964
		{
			yyVAL.valExprs = nil
		}
	case 179:
		//line sql.y:968
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 180:
		//line sql.y:973
		{
			yyVAL.boolExpr = nil
		}
	case 181:
		//line sql.y:977
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 182:
		//line sql.y:982
		{
			yyVAL.orderBy = nil
		}
	case 183:
		//line sql.y:986
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 184:
		//line sql.y:992
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 185:
		//line sql.y:996
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 186:
		//line sql.y:1002
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 187:
		//line sql.y:1007
		{
			yyVAL.str = AST_ASC
		}
	case 188:
		//line sql.y:1011
		{
			yyVAL.str = AST_ASC
		}
	case 189:
		//line sql.y:1015
		{
			yyVAL.str = AST_DESC
		}
	case 190:
		//line sql.y:1020
		{
			yyVAL.limit = nil
		}
	case 191:
		//line sql.y:1024
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 192:
		//line sql.y:1028
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 193:
		//line sql.y:1033
		{
			yyVAL.str = ""
		}
	case 194:
		//line sql.y:1037
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 195:
		//line sql.y:1041
		{
			if !bytes.Equal(yyS[yypt-1].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyS[yypt-0].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 196:
		//line sql.y:1054
		{
			yyVAL.columns = nil
		}
	case 197:
		//line sql.y:1058
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 198:
		//line sql.y:1064
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 199:
		//line sql.y:1068
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 200:
		//line sql.y:1073
		{
			yyVAL.updateExprs = nil
		}
	case 201:
		//line sql.y:1077
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 202:
		//line sql.y:1083
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 203:
		//line sql.y:1087
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 204:
		//line sql.y:1093
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 205:
		//line sql.y:1098
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		//line sql.y:1100
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		//line sql.y:1103
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		//line sql.y:1105
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		//line sql.y:1108
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		//line sql.y:1110
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		//line sql.y:1114
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		//line sql.y:1116
		{
			yyVAL.empty = struct{}{}
		}
	case 213:
		//line sql.y:1118
		{
			yyVAL.empty = struct{}{}
		}
	case 214:
		//line sql.y:1120
		{
			yyVAL.empty = struct{}{}
		}
	case 215:
		//line sql.y:1122
		{
			yyVAL.empty = struct{}{}
		}
	case 216:
		//line sql.y:1125
		{
			yyVAL.empty = struct{}{}
		}
	case 217:
		//line sql.y:1127
		{
			yyVAL.empty = struct{}{}
		}
	case 218:
		//line sql.y:1130
		{
			yyVAL.empty = struct{}{}
		}
	case 219:
		//line sql.y:1132
		{
			yyVAL.empty = struct{}{}
		}
	case 220:
		//line sql.y:1135
		{
			yyVAL.empty = struct{}{}
		}
	case 221:
		//line sql.y:1137
		{
			yyVAL.empty = struct{}{}
		}
	case 222:
		//line sql.y:1141
		{
			yyVAL.bytes = bytes.ToLower(yyS[yypt-0].bytes)
		}
	case 223:
		//line sql.y:1146
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
