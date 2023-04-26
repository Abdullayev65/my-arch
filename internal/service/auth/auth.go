package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"my-arch/internal/dto/auth"
	"my-arch/internal/dto/user"
	"my-arch/internal/model"
	"my-arch/internal/tools/password_tool"
	"my-arch/internal/tools/time_tool"
	"my-arch/internal/tools/valid_tool"
	"my-arch/pkg/ctx"
	"regexp"
	"time"
)

type Service struct {
	User          User
	emailRegex    *regexp.Regexp
	jwtKey        []byte
	tokenExpiring time.Duration
}

func New(user User) *Service {
	n := new(Service)

	n.User = user
	n.jwtKey = []byte("jwt-secret-key")
	n.tokenExpiring = time.Hour << 7

	return n
}

func (s *Service) SignUp(c ctx.Ctx, input *user.UserCreate) error {
	var errStr string
	switch {
	case input.Email == nil:
		errStr = "email is required"
	case !valid_tool.IsValidEmail(*input.Email):
		errStr = "email is not valid"
	case input.Username == nil:
		errStr = "username is required"
	case valid_tool.IsValidUsername(*input.Username) != nil:
		errStr = valid_tool.IsValidUsername(*input.Username).Error()
	case len(*input.Username) < 3 || len(*input.Username) > 26:
		errStr = "username length should be between 3 and 26"
	case input.Password == nil:
		errStr = "password is required"
	case len(*input.Password) < 1 || len(*input.Password) > 30:
		errStr = "password length should be between 1 and 30"
	case input.FirstName == nil:
		errStr = "first_name is required"
	}
	if errStr != "" {
		return errors.New(errStr)
	}

	password, err := password_tool.HashPassword(*input.Password)
	if err != nil {
		return err
	}

	input.Password = &password
	time_tool.Parse(input.BirthDateStr, &input.BirthDate)

	err = s.User.Create(c, input)

	return err
}

func (s *Service) LogIn(c ctx.Ctx, data *auth.LogIn) (*auth.Token, error) {
	if data.Identifier == nil || data.Password == nil {
		return nil, errors.New("identifier and password is required")
	}
	var m *model.User
	var err error
	if valid_tool.IsValidEmail(*data.Identifier) {
		m, err = s.User.GetByEmail(c, *data.Identifier)
	} else {
		m, err = s.User.GetByUsername(c, *data.Identifier)
	}

	if err != nil || !password_tool.ComparePassword(m.Password, *data.Password) {
		return nil, errors.New("identifier or password gone wrong")
	}

	token, err := s.GenerateToken(m.Id)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// specific functions

func (s *Service) GenerateToken(id int) (*auth.Token, error) {
	claims := &Claims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.tokenExpiring).Unix(),
		},
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString(s.jwtKey)

	if err != nil {
		return nil, err
	}

	token := new(auth.Token)
	token.Token = tokenString

	return token, nil
}

func (s *Service) UserIdFromToken(tokenStr string) (int, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return s.jwtKey, nil
	})
	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	return claims.ID, err
}
