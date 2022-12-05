package token

const (
	_ Token = iota

	// Control.
	ILLEGAL
	EOF
	COMMENT
	KEYWORD
	// Types.
	STRING
	BOOLEAN
	NULL
	NUMBER
	IDENTIFIER
	// Maths.
	PLUS      // +
	MINUS     // -
	MULTIPLY  // *
	SLASH     // /
	REMAINDER // %
	// Logical and bitwise operators.
	AND                  // &
	OR                   // |
	EXCLUSIVE_OR         // ^
	SHIFT_LEFT           // <<
	SHIFT_RIGHT          // >>
	UNSIGNED_SHIFT_RIGHT // >>>
	AND_NOT              // &^
	// Math assignments.
	ADD_ASSIGN       // +=
	SUBTRACT_ASSIGN  // -=
	MULTIPLY_ASSIGN  // *=
	QUOTIENT_ASSIGN  // /=
	REMAINDER_ASSIGN // %=
	// Math and bitwise assignments.
	AND_ASSIGN                  // &=
	OR_ASSIGN                   // |=
	EXCLUSIVE_OR_ASSIGN         // ^=
	SHIFT_LEFT_ASSIGN           // <<=
	SHIFT_RIGHT_ASSIGN          // >>=
	UNSIGNED_SHIFT_RIGHT_ASSIGN // >>>=
	AND_NOT_ASSIGN              // &^=
	// Logical operators and decrement / increment.
	LOGICAL_AND // &&
	LOGICAL_OR  // ||
	INCREMENT   // ++
	DECREMENT   // --
	// Comparison operators.
	EQUAL        // ==
	STRICT_EQUAL // ===
	LESS         // <
	GREATER      // >
	ASSIGN       // =
	NOT          // !
	// Bitwise not.
	BITWISE_NOT // ~
	// Comparison operators.
	NOT_EQUAL        // !=
	STRICT_NOT_EQUAL // !==
	LESS_OR_EQUAL    // <=
	GREATER_OR_EQUAL // >=
	// Left operators.
	LEFT_PARENTHESIS // (
	LEFT_BRACKET     // [
	LEFT_BRACE       // {
	COMMA            // ,
	PERIOD           // .
	// Right operators.
	RIGHT_PARENTHESIS // )
	RIGHT_BRACKET     // ]
	RIGHT_BRACE       // }
	SEMICOLON         // ;
	COLON             // :
	QUESTION_MARK     // ?
	// Basic flow - keywords below here.
	_
	IF
	IN
	DO
	// Declarations.
	VAR
	FOR
	NEW
	TRY
	// Advanced flow.
	THIS
	ELSE
	CASE
	VOID
	WITH
	// Loops.
	WHILE
	BREAK
	CATCH
	THROW
	// Functions.
	RETURN
	TYPEOF
	DELETE
	SWITCH
	// Fallback identifiers.
	DEFAULT
	FINALLY
	// Miscellaneous.
	FUNCTION
	CONTINUE
	DEBUGGER
	// Instance of.
	INSTANCEOF
)

var token2string = [...]string{
	ILLEGAL:                     "ILLEGAL",
	EOF:                         "EOF",
	COMMENT:                     "COMMENT",
	KEYWORD:                     "KEYWORD",
	STRING:                      "STRING",
	BOOLEAN:                     "BOOLEAN",
	NULL:                        "NULL",
	NUMBER:                      "NUMBER",
	IDENTIFIER:                  "IDENTIFIER",
	PLUS:                        "+",
	MINUS:                       "-",
	MULTIPLY:                    "*",
	SLASH:                       "/",
	REMAINDER:                   "%",
	AND:                         "&",
	OR:                          "|",
	EXCLUSIVE_OR:                "^",
	SHIFT_LEFT:                  "<<",
	SHIFT_RIGHT:                 ">>",
	UNSIGNED_SHIFT_RIGHT:        ">>>",
	AND_NOT:                     "&^",
	ADD_ASSIGN:                  "+=",
	SUBTRACT_ASSIGN:             "-=",
	MULTIPLY_ASSIGN:             "*=",
	QUOTIENT_ASSIGN:             "/=",
	REMAINDER_ASSIGN:            "%=",
	AND_ASSIGN:                  "&=",
	OR_ASSIGN:                   "|=",
	EXCLUSIVE_OR_ASSIGN:         "^=",
	SHIFT_LEFT_ASSIGN:           "<<=",
	SHIFT_RIGHT_ASSIGN:          ">>=",
	UNSIGNED_SHIFT_RIGHT_ASSIGN: ">>>=",
	AND_NOT_ASSIGN:              "&^=",
	LOGICAL_AND:                 "&&",
	LOGICAL_OR:                  "||",
	INCREMENT:                   "++",
	DECREMENT:                   "--",
	EQUAL:                       "==",
	STRICT_EQUAL:                "===",
	LESS:                        "<",
	GREATER:                     ">",
	ASSIGN:                      "=",
	NOT:                         "!",
	BITWISE_NOT:                 "~",
	NOT_EQUAL:                   "!=",
	STRICT_NOT_EQUAL:            "!==",
	LESS_OR_EQUAL:               "<=",
	GREATER_OR_EQUAL:            ">=",
	LEFT_PARENTHESIS:            "(",
	LEFT_BRACKET:                "[",
	LEFT_BRACE:                  "{",
	COMMA:                       ",",
	PERIOD:                      ".",
	RIGHT_PARENTHESIS:           ")",
	RIGHT_BRACKET:               "]",
	RIGHT_BRACE:                 "}",
	SEMICOLON:                   ";",
	COLON:                       ":",
	QUESTION_MARK:               "?",
	IF:                          "if",
	IN:                          "in",
	DO:                          "do",
	VAR:                         "var",
	FOR:                         "for",
	NEW:                         "new",
	TRY:                         "try",
	THIS:                        "this",
	ELSE:                        "else",
	CASE:                        "case",
	VOID:                        "void",
	WITH:                        "with",
	WHILE:                       "while",
	BREAK:                       "break",
	CATCH:                       "catch",
	THROW:                       "throw",
	RETURN:                      "return",
	TYPEOF:                      "typeof",
	DELETE:                      "delete",
	SWITCH:                      "switch",
	DEFAULT:                     "default",
	FINALLY:                     "finally",
	FUNCTION:                    "function",
	CONTINUE:                    "continue",
	DEBUGGER:                    "debugger",
	INSTANCEOF:                  "instanceof",
}

var keywordTable = map[string]keyword{
	"if": {
		token: IF,
	},
	"in": {
		token: IN,
	},
	"do": {
		token: DO,
	},
	"var": {
		token: VAR,
	},
	"for": {
		token: FOR,
	},
	"new": {
		token: NEW,
	},
	"try": {
		token: TRY,
	},
	"this": {
		token: THIS,
	},
	"else": {
		token: ELSE,
	},
	"case": {
		token: CASE,
	},
	"void": {
		token: VOID,
	},
	"with": {
		token: WITH,
	},
	"while": {
		token: WHILE,
	},
	"break": {
		token: BREAK,
	},
	"catch": {
		token: CATCH,
	},
	"throw": {
		token: THROW,
	},
	"return": {
		token: RETURN,
	},
	"typeof": {
		token: TYPEOF,
	},
	"delete": {
		token: DELETE,
	},
	"switch": {
		token: SWITCH,
	},
	"default": {
		token: DEFAULT,
	},
	"finally": {
		token: FINALLY,
	},
	"function": {
		token: FUNCTION,
	},
	"continue": {
		token: CONTINUE,
	},
	"debugger": {
		token: DEBUGGER,
	},
	"instanceof": {
		token: INSTANCEOF,
	},
	"const": {
		token:         KEYWORD,
		futureKeyword: true,
	},
	"class": {
		token:         KEYWORD,
		futureKeyword: true,
	},
	"enum": {
		token:         KEYWORD,
		futureKeyword: true,
	},
	"export": {
		token:         KEYWORD,
		futureKeyword: true,
	},
	"extends": {
		token:         KEYWORD,
		futureKeyword: true,
	},
	"import": {
		token:         KEYWORD,
		futureKeyword: true,
	},
	"super": {
		token:         KEYWORD,
		futureKeyword: true,
	},
	"implements": {
		token:         KEYWORD,
		futureKeyword: true,
		strict:        true,
	},
	"interface": {
		token:         KEYWORD,
		futureKeyword: true,
		strict:        true,
	},
	"let": {
		token:         KEYWORD,
		futureKeyword: true,
		strict:        true,
	},
	"package": {
		token:         KEYWORD,
		futureKeyword: true,
		strict:        true,
	},
	"private": {
		token:         KEYWORD,
		futureKeyword: true,
		strict:        true,
	},
	"protected": {
		token:         KEYWORD,
		futureKeyword: true,
		strict:        true,
	},
	"public": {
		token:         KEYWORD,
		futureKeyword: true,
		strict:        true,
	},
	"static": {
		token:         KEYWORD,
		futureKeyword: true,
		strict:        true,
	},
}
