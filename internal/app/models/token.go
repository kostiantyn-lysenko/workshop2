package models

type Token struct {
	Type  string
	Value string
}

const TokenTypeAccess = "access"
const TokenTypeRefresh = "refresh"
