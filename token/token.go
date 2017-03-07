package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	CONSTANT          = "CONSTANT"
	IDENT             = "IDENT"
	INSTANCE_VARIABLE = "INSTANCE_VAR"
	INT               = "INT"
	STRING            = "STRING"
	COMMENT           = "COMMENT"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	DOT      = "."
	INCR     = "++"
	DECR     = "--"

	LT = "<"
	GT = ">"

	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	EQ     = "=="
	NOT_EQ = "!="

	CLASS  = "CLASS"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
	DEF    = "DEF"
	SELF   = "SELF"
)

var keyworkds = map[string]TokenType{
	"def":    DEF,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"self":   SELF,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keyworkds[ident]; ok {
		return tok
	}
	return IDENT
}
