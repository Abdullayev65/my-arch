package auth

type LogIn struct {
	Identifier *string
	Password   *string
	//identifier = email or username
}

type Token struct {
	Token string
}

type Available struct {
	Type  int
	Value string
	//type:
	//1 = username
	//2 = email
}
