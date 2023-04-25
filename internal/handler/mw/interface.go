package mw

type Auth interface {
	UserIdFromToken(tokenStr string) (int, error)
}
