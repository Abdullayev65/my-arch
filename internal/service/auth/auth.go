package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"mindstore/internal/object/dto/auth"
	"mindstore/internal/object/dto/user"
	"mindstore/internal/object/model"
	"mindstore/pkg/ctx"
	"mindstore/pkg/encoder"
	"mindstore/pkg/hash-types"
	"mindstore/pkg/util/timeutil"
	"regexp"
	"strings"
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
	n.emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	n.jwtKey = []byte("jwt-secret-key")
	n.tokenExpiring = time.Hour << 7

	return n
}

func (s *Service) SignUp(c ctx.Ctx, input *user.UserCreate) error {
	var errStr string
	switch {
	case input.Email == nil:
		errStr = "email is required"
	case !s.IsValidEmail(*input.Email):
		errStr = "email is not valid"
	case input.Username == nil:
		errStr = "username is required"
	case s.IsValidUsername(*input.Username) != nil:
		errStr = s.IsValidUsername(*input.Username).Error()
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

	password, err := encoder.HashPassword(*input.Password)
	if err != nil {
		return err
	}

	input.Password = &password
	timeutil.Parse(input.BirthDateStr, &input.BirthDate)

	_, err = s.User.CreateWithMind(c, input)

	return err
}

func (s *Service) LogIn(c ctx.Ctx, data *auth.LogIn) (*auth.Token, error) {
	if data.Identifier == nil || data.Password == nil {
		return nil, errors.New("identifier and password is required")
	}
	var m *model.User
	var err error
	if s.IsValidEmail(*data.Identifier) {
		m, err = s.User.GetByEmail(c, *data.Identifier)
	} else {
		m, err = s.User.GetByUsername(c, *data.Identifier)
	}

	if err != nil {
		return nil, err
	}

	token, err := s.GenerateToken(m.Id)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *Service) Available(c ctx.Ctx, input *auth.Available) (bool, error) {
	var column string
	switch input.Type {
	default:
		return false, errors.New("invalid type: type can be: username(1), email(2)")
	case 1:
		if err := s.IsValidUsername(input.Value); err != nil {
			return false, err
		}
		column = "username"
	case 2:
		if !s.IsValidEmail(input.Value) {
			return false, errors.New("email is invalid")
		}
		column = "email"
	}

	if input.Value == "" {
		return false, errors.New("value not given")
	}

	return s.User.Available(c, column, input.Value)
}

// specific functions
func (s *Service) IsValidEmail(email string) bool {
	return s.emailRegex.MatchString(email)
}

func (s *Service) IsValidUsername(username string) error {
	if len(username) < 3 || len(username) > 26 {
		return errors.New("username length should be between 3 and 26")
	}
	index := strings.IndexFunc(username, func(r rune) bool {
		switch {
		default:
			return true
		case 'A' >= r && r >= 'Z':
		case 'a' >= r && r >= 'z':
		case '0' >= r && r >= '9':
		case r == '.' || r == '_':
		}
		return false
	})
	if index != 0 {
		return fmt.Errorf("email should not conatein %c", username[index])
	}

	return nil
}

func (s *Service) GenerateToken(id hash.Int) (*auth.Token, error) {
	claims := &Claims{
		ID: &id,
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

func (s *Service) UserIdFromToken(tokenStr string) (*hash.Int, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return s.jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims.ID, err
}
