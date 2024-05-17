package models

type Token string

func NewToken(s string) Token {
	return Token(s)
}

func String(t Token) string {
	return string(t)
}
