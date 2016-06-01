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
const SOME = 57359
const ANY = 57360
const ALL = 57361
const DISTINCT = 57362
const PRECISION = 57363
const AS = 57364
const EXISTS = 57365
const IN = 57366
const IS = 57367
const LIKE = 57368
const BETWEEN = 57369
const NULL = 57370
const ASC = 57371
const DESC = 57372
const VALUES = 57373
const INTO = 57374
const DUPLICATE = 57375
const KEY = 57376
const DEFAULT = 57377
const SET = 57378
const LOCK = 57379
const ID = 57380
const STRING = 57381
const NUMBER = 57382
const VALUE_ARG = 57383
const COMMENT = 57384
const LE = 57385
const GE = 57386
const NE = 57387
const NULL_SAFE_EQUAL = 57388
const DATE = 57389
const DATETIME = 57390
const TIME = 57391
const TIMESTAMP = 57392
const YEAR = 57393
const UNION = 57394
const MINUS = 57395
const EXCEPT = 57396
const INTERSECT = 57397
const JOIN = 57398
const STRAIGHT_JOIN = 57399
const LEFT = 57400
const RIGHT = 57401
const INNER = 57402
const OUTER = 57403
const CROSS = 57404
const NATURAL = 57405
const USE = 57406
const FORCE = 57407
const ON = 57408
const AND = 57409
const OR = 57410
const NOT = 57411
const UNARY = 57412
const CASE = 57413
const WHEN = 57414
const THEN = 57415
const ELSE = 57416
const END = 57417
const BEGIN = 57418
const COMMIT = 57419
const ROLLBACK = 57420
const NAMES = 57421
const REPLACE = 57422
const ADMIN = 57423
const SHOW = 57424
const DATABASES = 57425
const TABLES = 57426
const PROXY = 57427
const VARIABLES = 57428
const FULL = 57429
const COLUMNS = 57430
const CREATE = 57431
const ALTER = 57432
const DROP = 57433
const RENAME = 57434
const TABLE = 57435
const INDEX = 57436
const VIEW = 57437
const TO = 57438
const IGNORE = 57439
const IF = 57440
const UNIQUE = 57441
const USING = 57442

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
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
	"SOME",
	"ANY",
	"ALL",
	"DISTINCT",
	"PRECISION",
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
	"DATE",
	"DATETIME",
	"TIME",
	"TIMESTAMP",
	"YEAR",
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
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 196,
	73, 72,
	74, 72,
	-2, 149,
	-1, 411,
	73, 71,
	74, 71,
	-2, 82,
}

const yyNprod = 240
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 676

var yyAct = [...]int{

	114, 350, 106, 419, 73, 103, 387, 105, 111, 192,
	141, 122, 298, 245, 342, 112, 289, 215, 291, 194,
	100, 193, 3, 75, 210, 101, 34, 35, 36, 37,
	280, 428, 223, 62, 91, 314, 315, 316, 317, 318,
	87, 319, 320, 80, 77, 428, 428, 82, 162, 305,
	84, 348, 76, 162, 88, 64, 162, 262, 261, 260,
	97, 144, 49, 243, 50, 398, 243, 243, 121, 397,
	44, 127, 46, 52, 53, 54, 47, 396, 78, 118,
	119, 120, 140, 367, 369, 282, 430, 152, 81, 371,
	148, 125, 94, 95, 83, 143, 151, 51, 150, 154,
	429, 427, 160, 377, 166, 378, 347, 98, 337, 159,
	78, 335, 196, 380, 190, 195, 203, 191, 334, 123,
	124, 281, 242, 290, 63, 368, 128, 325, 168, 206,
	153, 209, 77, 137, 372, 77, 213, 70, 219, 218,
	76, 248, 132, 76, 56, 58, 59, 57, 61, 220,
	247, 139, 290, 126, 340, 160, 18, 393, 217, 233,
	240, 241, 180, 181, 182, 343, 74, 258, 256, 219,
	254, 255, 259, 249, 234, 268, 269, 301, 272, 273,
	274, 275, 276, 277, 278, 279, 263, 257, 252, 248,
	236, 155, 270, 229, 134, 374, 147, 160, 247, 373,
	164, 165, 251, 361, 283, 343, 250, 130, 362, 359,
	133, 395, 77, 77, 360, 227, 294, 394, 230, 365,
	76, 296, 300, 364, 302, 285, 287, 363, 149, 134,
	161, 297, 293, 243, 404, 303, 77, 382, 216, 271,
	307, 308, 309, 216, 76, 86, 310, 235, 136, 414,
	251, 413, 412, 306, 250, 211, 293, 286, 212, 249,
	117, 324, 311, 167, 264, 121, 212, 160, 127, 330,
	331, 207, 326, 327, 328, 104, 118, 119, 120, 63,
	121, 329, 162, 127, 109, 226, 228, 225, 125, 312,
	78, 118, 119, 120, 134, 205, 195, 204, 341, 152,
	89, 99, 129, 125, 63, 339, 323, 336, 345, 346,
	349, 78, 108, 425, 164, 165, 123, 124, 102, 406,
	407, 370, 322, 128, 249, 249, 357, 358, 401, 354,
	353, 123, 124, 232, 426, 376, 231, 214, 128, 34,
	35, 36, 37, 379, 178, 179, 180, 181, 182, 77,
	126, 384, 71, 284, 385, 388, 145, 383, 142, 138,
	135, 85, 131, 381, 389, 126, 175, 176, 177, 178,
	179, 180, 181, 182, 18, 90, 69, 265, 399, 266,
	267, 375, 333, 400, 175, 176, 177, 178, 179, 180,
	181, 182, 92, 221, 238, 160, 432, 409, 402, 195,
	292, 411, 410, 408, 18, 253, 416, 388, 93, 239,
	418, 417, 146, 420, 420, 420, 77, 421, 422, 157,
	423, 67, 117, 65, 76, 299, 351, 121, 392, 433,
	127, 352, 391, 434, 158, 435, 356, 78, 118, 119,
	120, 403, 216, 96, 72, 431, 109, 18, 415, 18,
	125, 198, 201, 199, 200, 202, 175, 176, 177, 178,
	179, 180, 181, 182, 39, 314, 315, 316, 317, 318,
	121, 319, 320, 127, 108, 60, 237, 156, 123, 124,
	78, 118, 119, 120, 17, 128, 16, 15, 14, 152,
	13, 12, 197, 125, 198, 201, 199, 200, 202, 222,
	45, 117, 304, 224, 48, 79, 121, 295, 424, 127,
	405, 386, 126, 390, 18, 117, 104, 118, 119, 120,
	121, 123, 124, 127, 355, 109, 338, 208, 128, 125,
	78, 118, 119, 120, 288, 116, 113, 121, 115, 109,
	127, 344, 110, 125, 169, 107, 366, 78, 118, 119,
	120, 246, 313, 108, 244, 126, 152, 123, 124, 102,
	125, 321, 163, 66, 128, 33, 68, 108, 11, 10,
	9, 123, 124, 18, 19, 20, 21, 332, 128, 8,
	175, 176, 177, 178, 179, 180, 181, 182, 123, 124,
	7, 126, 6, 5, 4, 128, 2, 1, 0, 170,
	174, 172, 173, 0, 22, 126, 175, 176, 177, 178,
	179, 180, 181, 182, 63, 38, 0, 0, 186, 187,
	188, 189, 126, 183, 184, 185, 0, 0, 198, 201,
	199, 200, 202, 0, 0, 40, 41, 42, 43, 0,
	0, 0, 0, 0, 0, 0, 55, 0, 0, 0,
	171, 175, 176, 177, 178, 179, 180, 181, 182, 27,
	28, 29, 0, 30, 32, 31, 0, 0, 0, 0,
	0, 0, 23, 24, 26, 25,
}
var yyPact = [...]int{

	568, -1000, -1000, 282, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -38, -48, -11, -35, -1000, -1000, -1000,
	-1000, 46, 266, 444, 404, -1000, -1000, -1000, 401, -1000,
	344, 314, 435, 72, -70, -21, 266, -1000, -14, 266,
	-1000, 323, -73, 266, -73, 343, 382, 382, 434, 266,
	4, -1000, 254, -1000, -1000, -1000, 478, -1000, 260, 314,
	326, 58, 314, 168, 322, -1000, 200, -1000, 49, 321,
	76, 266, -1000, 320, -1000, -50, 318, 389, 124, 266,
	314, -1000, 492, 252, -1000, 382, 252, 434, 410, 252,
	221, -1000, -1000, 241, 44, -1000, 575, -1000, 492, 399,
	-1000, -1000, -1000, 252, 250, 248, -1000, 224, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 252, -1000,
	219, 273, 299, 432, 273, -1000, 252, 266, -1000, 370,
	-83, -1000, 180, -1000, 298, -1000, -1000, 295, -1000, 211,
	127, 530, 442, -1000, 530, 382, 385, 252, 252, 6,
	530, 103, 478, 384, 492, 492, -1000, 576, 86, 40,
	217, 353, 252, 252, 164, 252, 252, 252, 252, 252,
	252, 252, 252, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -86, 5, -31, 252, 127, 575, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 237, 478, -1000, 444, 36, 530,
	369, 273, 273, 233, -1000, 412, 492, -1000, 530, -1000,
	-1000, -1000, 105, 266, -1000, -62, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 369, 273, -1000, -1000, 252, 252,
	530, 530, -1000, 252, 228, 403, 284, 151, 43, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 530,
	224, 224, 224, -1000, 509, 217, 252, 252, 530, 504,
	-1000, 354, 265, 265, 265, 81, 81, -1000, -1000, -1000,
	-1000, -1000, -1000, 2, -1000, -5, 478, -8, 65, -1000,
	492, 93, 217, 282, 133, -10, -1000, 412, 411, 417,
	127, 292, -1000, -1000, 291, -1000, -1000, 168, 530, 530,
	530, 425, 103, 103, -1000, -1000, 147, 141, 165, 161,
	157, 13, -1000, 283, -27, 96, -1000, -1000, -1000, -1000,
	530, 308, 252, -1000, -1000, -1000, -13, -1000, 15, -1000,
	252, 25, -1000, 330, 176, -1000, -1000, -1000, 273, 411,
	-1000, 252, 252, -1000, -1000, 420, 414, 403, 85, -1000,
	155, -1000, 149, -1000, -1000, -1000, -1000, -32, -40, -44,
	-1000, -1000, -1000, -1000, -1000, 252, 530, -1000, -1000, 530,
	252, 294, 217, -1000, -1000, 380, 173, -1000, 290, -1000,
	412, 492, 252, 492, -1000, -1000, 205, 204, 202, 530,
	530, 441, -1000, 252, 252, -1000, -1000, -1000, 411, 127,
	172, -1000, 266, 266, 266, 273, 530, -1000, 297, -15,
	-1000, -16, -30, 168, -1000, 438, 372, -1000, 266, -1000,
	-1000, -1000, 266, -1000, 266, -1000,
}
var yyPgo = [...]int{

	0, 597, 596, 21, 594, 593, 592, 590, 579, 570,
	569, 568, 615, 566, 565, 563, 20, 25, 562, 561,
	5, 554, 13, 552, 551, 137, 546, 3, 17, 7,
	545, 544, 18, 542, 2, 15, 9, 541, 538, 11,
	536, 8, 535, 534, 16, 527, 526, 524, 513, 12,
	511, 6, 510, 1, 508, 24, 507, 14, 4, 23,
	245, 505, 504, 503, 502, 500, 499, 0, 19, 492,
	10, 491, 490, 488, 487, 486, 484, 93, 34, 477,
	476, 475, 464,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	4, 4, 74, 74, 5, 6, 7, 7, 71, 72,
	73, 76, 79, 79, 80, 80, 80, 81, 81, 75,
	75, 75, 75, 75, 8, 8, 8, 9, 9, 9,
	10, 11, 11, 11, 82, 12, 13, 13, 14, 14,
	14, 14, 14, 15, 15, 16, 16, 17, 17, 17,
	17, 20, 20, 18, 18, 18, 18, 21, 21, 22,
	22, 22, 22, 19, 19, 19, 23, 23, 23, 23,
	23, 23, 23, 23, 23, 24, 24, 24, 24, 24,
	24, 24, 25, 25, 26, 26, 26, 26, 27, 27,
	28, 28, 78, 78, 78, 77, 77, 29, 29, 29,
	29, 29, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 31, 31, 31, 31, 31,
	31, 31, 32, 32, 37, 37, 35, 35, 39, 36,
	36, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 38,
	38, 40, 40, 40, 42, 45, 45, 43, 43, 44,
	46, 46, 41, 41, 33, 33, 33, 33, 47, 47,
	48, 48, 49, 49, 50, 50, 51, 52, 52, 52,
	53, 53, 53, 54, 54, 54, 55, 55, 56, 56,
	57, 57, 58, 58, 59, 60, 60, 61, 61, 62,
	62, 63, 63, 63, 63, 63, 64, 64, 65, 65,
	66, 66, 67, 68, 69, 69, 69, 69, 69, 70,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 12, 3,
	7, 7, 6, 6, 8, 7, 3, 4, 1, 1,
	1, 5, 2, 2, 0, 2, 2, 0, 1, 3,
	3, 4, 5, 5, 5, 8, 4, 6, 7, 4,
	5, 4, 5, 5, 0, 2, 0, 2, 1, 2,
	1, 1, 1, 0, 1, 1, 3, 1, 2, 3,
	3, 1, 1, 0, 1, 2, 2, 1, 3, 3,
	3, 3, 5, 0, 1, 2, 1, 1, 2, 3,
	2, 3, 2, 2, 2, 1, 3, 1, 1, 1,
	3, 3, 1, 3, 0, 5, 5, 5, 1, 3,
	0, 2, 0, 2, 2, 0, 2, 1, 3, 3,
	2, 3, 3, 3, 4, 4, 4, 4, 3, 4,
	5, 6, 3, 4, 2, 1, 1, 1, 1, 1,
	1, 1, 2, 1, 1, 3, 3, 1, 3, 1,
	3, 1, 1, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 3, 4, 5, 4, 4, 1, 1,
	1, 1, 1, 1, 5, 0, 1, 1, 2, 4,
	0, 2, 1, 3, 1, 1, 1, 1, 0, 3,
	0, 2, 0, 3, 1, 3, 2, 0, 1, 1,
	0, 2, 4, 0, 2, 4, 0, 3, 1, 3,
	0, 5, 1, 3, 3, 0, 2, 0, 3, 0,
	1, 1, 1, 1, 1, 1, 0, 1, 0, 1,
	0, 2, 1, 1, 1, 1, 1, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -71, -72, -73, -74, -75, -76, 5, 6,
	7, 8, 36, 104, 105, 107, 106, 91, 92, 93,
	95, 97, 96, -14, 57, 58, 59, 60, -12, -82,
	-12, -12, -12, -12, 108, -65, 110, 114, -62, 110,
	112, 108, 108, 109, 110, -12, 98, 101, 99, 100,
	-81, 102, -67, 38, -3, 19, -15, 20, -13, 32,
	-25, 38, 9, -58, 94, -59, -41, -67, 38, -61,
	113, 109, -67, 108, -67, 38, -60, 113, -67, -60,
	32, -78, 10, 26, -78, -77, 9, -67, 103, 47,
	-16, -17, 81, -20, 38, -29, -34, -30, 75, 47,
	-33, -41, -35, -40, -67, -38, -42, 23, 39, 40,
	41, 28, -39, 79, 80, 51, 113, 31, 86, 42,
	-25, 36, 84, -25, 61, 38, 48, 84, 38, 75,
	-67, -70, 38, -70, 111, 38, 23, 72, -67, -25,
	-20, -34, 47, -78, -34, -77, -79, 9, 24, -36,
	-34, 9, 61, -18, 73, 74, -67, 22, 84, -31,
	24, 75, 26, 27, 25, 76, 77, 78, 79, 80,
	81, 82, 83, 48, 49, 50, 43, 44, 45, 46,
	-20, -29, -36, -3, -68, -20, -34, -69, 52, 54,
	55, 53, 56, -34, 47, 47, -39, 47, -45, -34,
	-55, 36, 47, -58, 38, -28, 10, -59, -34, -67,
	-70, 23, -66, 115, -63, 107, 105, 35, 106, 13,
	38, 38, 38, -70, -55, 36, -78, -80, 9, 24,
	-34, -34, 116, 61, -21, -22, -24, 47, 38, -39,
	103, 99, -17, 21, -20, -20, -67, -68, 81, -34,
	19, 18, 17, -35, 47, 24, 26, 27, -34, -34,
	28, 75, -34, -34, -34, -34, -34, -34, -34, -34,
	116, 116, 116, -36, 116, -16, 20, -16, -43, -44,
	87, -32, 31, -3, -58, -56, -41, -28, -49, 13,
	-20, 72, -67, -70, -64, 111, -32, -58, -34, -34,
	-34, -28, 61, -23, 62, 63, 64, 65, 66, 68,
	69, -19, 38, 22, -22, 84, -39, -39, -39, -35,
	-34, -34, 73, 28, 116, 116, -16, 116, -46, -44,
	89, -29, -57, 72, -37, -35, -57, 116, 61, -49,
	-53, 15, 14, 38, 38, -47, 11, -22, -22, 62,
	67, 62, 67, 62, 62, 62, -26, 70, 112, 71,
	38, 116, 38, 103, 99, 73, -34, 116, 90, -34,
	88, 33, 61, -41, -53, -34, -50, -51, -34, -70,
	-48, 12, 14, 72, 62, 62, 109, 109, 109, -34,
	-34, 34, -35, 61, 61, -52, 29, 30, -49, -20,
	-36, -29, 47, 47, 47, 7, -34, -51, -53, -27,
	-67, -27, -27, -58, -54, 16, 37, 116, 61, 116,
	116, 7, 24, -67, -67, -67,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 54, 54,
	54, 54, 54, 228, 219, 0, 0, 28, 29, 30,
	54, 37, 0, 0, 58, 60, 61, 62, 63, 56,
	0, 0, 0, 0, 217, 0, 0, 229, 0, 0,
	220, 0, 215, 0, 215, 0, 112, 112, 115, 0,
	0, 38, 0, 232, 19, 59, 0, 64, 55, 0,
	0, 102, 0, 26, 0, 212, 0, 182, 232, 0,
	0, 0, 239, 0, 239, 0, 0, 0, 0, 0,
	0, 39, 0, 0, 40, 112, 0, 115, 0, 0,
	17, 65, 67, 73, 232, 71, 72, 117, 0, 0,
	151, 152, 153, 0, 182, 0, 168, 0, 184, 185,
	186, 187, 147, 171, 172, 173, 169, 170, 175, 57,
	206, 0, 0, 110, 0, 27, 0, 0, 239, 0,
	230, 46, 0, 49, 0, 51, 216, 0, 239, 206,
	113, 114, 0, 41, 116, 112, 34, 0, 0, 0,
	149, 0, 0, 68, 0, 0, 74, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 135, 136, 137, 138, 139, 140, 141,
	120, 71, 0, 0, 0, 0, -2, 233, 234, 235,
	236, 237, 238, 162, 0, 0, 134, 0, 0, 176,
	0, 0, 0, 110, 103, 192, 0, 213, 214, 183,
	44, 218, 0, 0, 239, 226, 221, 222, 223, 224,
	225, 50, 52, 53, 0, 0, 42, 43, 0, 0,
	32, 33, 31, 0, 110, 77, 83, 0, 95, 97,
	98, 99, 66, 69, 118, 119, 75, 76, 70, 122,
	0, 0, 0, 123, 0, 0, 0, 0, 128, 0,
	132, 0, 154, 155, 156, 157, 158, 159, 160, 161,
	121, 146, 148, 0, 163, 0, 0, 0, 180, 177,
	0, 210, 0, 143, 210, 0, 208, 192, 200, 0,
	111, 0, 231, 47, 0, 227, 22, 23, 35, 36,
	150, 188, 0, 0, 86, 87, 0, 0, 0, 0,
	0, 104, 84, 0, 0, 0, 124, 125, 126, 127,
	129, 0, 0, 133, 166, 164, 0, 167, 0, 178,
	0, 71, 20, 0, 142, 144, 21, 207, 0, 200,
	25, 0, 0, 239, 48, 190, 0, 78, 81, 88,
	0, 90, 0, 92, 93, 94, 79, 0, 0, 0,
	85, 80, 96, 100, 101, 0, 130, 165, 174, 181,
	0, 0, 0, 209, 24, 201, 193, 194, 197, 45,
	192, 0, 0, 0, 89, 91, 0, 0, 0, 131,
	179, 0, 145, 0, 0, 196, 198, 199, 200, 191,
	189, -2, 0, 0, 0, 0, 202, 195, 203, 0,
	108, 0, 0, 211, 18, 0, 0, 105, 0, 106,
	107, 204, 0, 109, 0, 205,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 83, 76, 3,
	47, 116, 81, 79, 61, 80, 84, 82, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	49, 48, 50, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 78, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 77, 3, 51,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 52, 53, 54, 55, 56,
	57, 58, 59, 60, 62, 63, 64, 65, 66, 67,
	68, 69, 70, 71, 72, 73, 74, 75, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106,
	107, 108, 109, 110, 111, 112, 113, 114, 115,
}
var yyTok3 = [...]int{
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
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
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
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
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
	yyn = yyPact[yystate]
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
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
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
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
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
			if yyn < 0 || yyn == yytoken {
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

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
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
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:173
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:179
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:199
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs}
		}
	case 18:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:203
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].expr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].expr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:207
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 20:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:214
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 21:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:218
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 22:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:230
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 23:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:234
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 24:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:247
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].expr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:253
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].expr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:259
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:263
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal(yyDollar[4].bytes)}}}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:269
		{
			yyVAL.statement = &Begin{}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:275
		{
			yyVAL.statement = &Commit{}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:281
		{
			yyVAL.statement = &Rollback{}
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:287
		{
			yyVAL.statement = &Admin{Name: yyDollar[2].bytes, Values: yyDollar[4].valExprs}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:293
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:297
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 34:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:302
		{
			yyVAL.valExpr = nil
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:306
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:310
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:315
		{
			yyVAL.str = AST_SHOW_NO_MOD
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:319
		{
			yyVAL.str = AST_SHOW_FULL
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:326
		{
			yyVAL.statement = &Show{Section: "databases", LikeOrWhere: yyDollar[3].expr}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:330
		{
			yyVAL.statement = &Show{Section: "variables", LikeOrWhere: yyDollar[3].expr}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:334
		{
			yyVAL.statement = &Show{Section: "tables", From: yyDollar[3].valExpr, LikeOrWhere: yyDollar[4].expr}
		}
	case 42:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:338
		{
			yyVAL.statement = &Show{Section: "proxy", Key: string(yyDollar[3].bytes), From: yyDollar[4].valExpr, LikeOrWhere: yyDollar[5].expr}
		}
	case 43:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:342
		{
			yyVAL.statement = &Show{Section: "columns", From: yyDollar[4].valExpr, Modifier: yyDollar[2].str, DBFilter: yyDollar[5].valExpr}
		}
	case 44:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:348
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 45:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:352
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 46:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:357
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 47:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:363
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 48:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:367
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 49:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:372
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:378
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:384
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 52:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:388
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 53:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:393
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:398
		{
			SetAllowComments(yylex, true)
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:402
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:408
		{
			yyVAL.bytes2 = nil
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:412
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:418
		{
			yyVAL.str = AST_UNION
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:422
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:426
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:430
		{
			yyVAL.str = AST_EXCEPT
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:434
		{
			yyVAL.str = AST_INTERSECT
		}
	case 63:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:439
		{
			yyVAL.str = ""
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:443
		{
			yyVAL.str = AST_DISTINCT
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:449
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:453
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:459
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:463
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:467
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:471
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:477
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:481
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 73:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:486
		{
			yyVAL.bytes = nil
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:490
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:494
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:498
		{
			yyVAL.bytes = []byte(yyDollar[2].str)
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:505
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:509
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:515
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:519
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:523
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 82:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:527
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 83:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:532
		{
			yyVAL.bytes = nil
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:536
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:540
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:546
		{
			yyVAL.str = AST_JOIN
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:550
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:554
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:558
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:562
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:566
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:570
		{
			yyVAL.str = AST_JOIN
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:574
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:578
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:584
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:588
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:592
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:596
		{
			yyVAL.smTableExpr = &TableName{Name: []byte("columns")}
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:600
		{
			yyVAL.smTableExpr = &TableName{Name: []byte("tables")}
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:604
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: []byte("columns")}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:608
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: []byte("tables")}
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:614
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:618
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 104:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:623
		{
			yyVAL.indexHints = nil
		}
	case 105:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:627
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 106:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:631
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 107:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:635
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:641
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:645
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 110:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:650
		{
			yyVAL.expr = nil
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:654
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 112:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:659
		{
			yyVAL.expr = nil
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:663
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 114:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:667
		{
			yyVAL.expr = yyDollar[2].valExpr
		}
	case 115:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:672
		{
			yyVAL.valExpr = nil
		}
	case 116:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:676
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:683
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:687
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:691
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].expr}
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:695
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:701
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:705
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 124:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:709
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, SubqueryOperator: AST_ALL, Right: yyDollar[4].subquery}
		}
	case 125:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:713
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, SubqueryOperator: AST_ANY, Right: yyDollar[4].subquery}
		}
	case 126:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:717
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, SubqueryOperator: AST_SOME, Right: yyDollar[4].subquery}
		}
	case 127:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:721
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:725
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 129:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:729
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 130:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:733
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 131:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:737
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:741
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 133:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:745
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:749
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:755
		{
			yyVAL.str = AST_EQ
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:759
		{
			yyVAL.str = AST_LT
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:763
		{
			yyVAL.str = AST_GT
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:767
		{
			yyVAL.str = AST_LE
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:771
		{
			yyVAL.str = AST_GE
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:775
		{
			yyVAL.str = AST_NE
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:779
		{
			yyVAL.str = AST_NSE
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:785
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:789
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:795
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:799
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:805
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:809
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:815
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:821
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:825
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:831
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:835
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:839
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:843
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:847
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:851
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:855
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:859
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:863
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:867
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:871
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 162:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:875
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				switch yyDollar[1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
			}
		}
	case 163:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:890
		{
			yyVAL.valExpr = &FuncExpr{Name: bytes.ToLower(yyDollar[1].bytes)}
		}
	case 164:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:894
		{
			yyVAL.valExpr = &FuncExpr{Name: bytes.ToLower(yyDollar[1].bytes), Exprs: yyDollar[3].selectExprs}
		}
	case 165:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:898
		{
			yyVAL.valExpr = &FuncExpr{Name: bytes.ToLower(yyDollar[1].bytes), Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 166:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:902
		{
			yyVAL.valExpr = &CtorExpr{Name: yyDollar[2].str, Exprs: yyDollar[3].valExprs}
		}
	case 167:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:906
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:910
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:916
		{
			yyVAL.bytes = IF_BYTES
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:920
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:926
		{
			yyVAL.byt = AST_UPLUS
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:930
		{
			yyVAL.byt = AST_UMINUS
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:934
		{
			yyVAL.byt = AST_TILDA
		}
	case 174:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:940
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 175:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:945
		{
			yyVAL.valExpr = nil
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:949
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:955
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 178:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:959
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 179:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:965
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 180:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:970
		{
			yyVAL.valExpr = nil
		}
	case 181:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:974
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:980
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:984
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:990
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:994
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:998
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1002
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 188:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1007
		{
			yyVAL.valExprs = nil
		}
	case 189:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1011
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 190:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1016
		{
			yyVAL.expr = nil
		}
	case 191:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1020
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 192:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1025
		{
			yyVAL.orderBy = nil
		}
	case 193:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1029
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1035
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 195:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1039
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1045
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 197:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1050
		{
			yyVAL.str = AST_ASC
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1054
		{
			yyVAL.str = AST_ASC
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1058
		{
			yyVAL.str = AST_DESC
		}
	case 200:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1063
		{
			yyVAL.limit = nil
		}
	case 201:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1067
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 202:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1071
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 203:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1076
		{
			yyVAL.str = ""
		}
	case 204:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1080
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 205:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1084
		{
			if !bytes.Equal(yyDollar[3].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyDollar[4].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 206:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1097
		{
			yyVAL.columns = nil
		}
	case 207:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1101
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1107
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 209:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1111
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 210:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1116
		{
			yyVAL.updateExprs = nil
		}
	case 211:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1120
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1126
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 213:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1130
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 214:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1136
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 215:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1141
		{
			yyVAL.empty = struct{}{}
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1143
		{
			yyVAL.empty = struct{}{}
		}
	case 217:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1146
		{
			yyVAL.empty = struct{}{}
		}
	case 218:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1148
		{
			yyVAL.empty = struct{}{}
		}
	case 219:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1151
		{
			yyVAL.empty = struct{}{}
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1153
		{
			yyVAL.empty = struct{}{}
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1157
		{
			yyVAL.empty = struct{}{}
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1159
		{
			yyVAL.empty = struct{}{}
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1161
		{
			yyVAL.empty = struct{}{}
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1163
		{
			yyVAL.empty = struct{}{}
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1165
		{
			yyVAL.empty = struct{}{}
		}
	case 226:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1168
		{
			yyVAL.empty = struct{}{}
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1170
		{
			yyVAL.empty = struct{}{}
		}
	case 228:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1173
		{
			yyVAL.empty = struct{}{}
		}
	case 229:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1175
		{
			yyVAL.empty = struct{}{}
		}
	case 230:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1178
		{
			yyVAL.empty = struct{}{}
		}
	case 231:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1180
		{
			yyVAL.empty = struct{}{}
		}
	case 232:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1184
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 233:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1190
		{
			yyVAL.str = yyDollar[1].str
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1196
		{
			yyVAL.str = AST_DATE
		}
	case 235:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1200
		{
			yyVAL.str = AST_TIME
		}
	case 236:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1204
		{
			yyVAL.str = AST_TIMESTAMP
		}
	case 237:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1208
		{
			yyVAL.str = AST_DATETIME
		}
	case 238:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1212
		{
			yyVAL.str = AST_YEAR
		}
	case 239:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1217
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
