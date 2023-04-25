package mw

type MiddleWere struct {
	auth Auth
}

func New(auth Auth) *MiddleWere {
	n := new(MiddleWere)

	n.auth = auth

	return n
}
