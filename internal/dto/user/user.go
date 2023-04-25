package user

import (
	"mime/multipart"
	"mindstore/internal/object/dto"
	"mindstore/pkg/hash-types"
	"time"
)

type UserCreate struct {
	Username     *string
	Email        *string
	Password     *string
	FirstName    *string
	MiddleName   *string
	LastName     *string
	BirthDate    *time.Time `json:"-"`
	BirthDateStr *string    `json:"birth_date"`
}

type UserUpdate struct {
	Id           hash.Int `json:"-" form:"-"`
	Username     *string
	Email        *string
	Password     *string
	FirstName    *string
	MiddleName   *string
	LastName     *string
	Avatar       *multipart.FileHeader
	AvatarId     *hash.Int
	BirthDate    *time.Time `json:"-" form:"-"`
	BirthDateStr *string    `json:"birth_date" form:"birth_date"`
}

type UserDetail struct {
	Id           hash.Int
	Username     *string
	Email        *string
	MindId       *hash.Int
	FirstName    *string
	MiddleName   *string
	LastName     *string
	AvatarUrl    *string
	AvatarId     *hash.Int  `json:"-"`
	BirthDate    *time.Time `json:"-"`
	BirthDateStr *string    `json:"birth_date"`
}

type UserList struct {
	Id         hash.Int
	Username   *string
	MindId     *hash.Int
	FirstName  *string
	MiddleName *string
	LastName   *string
}

type Filter struct {
	dto.Filter

	Username *string
	Email    *string
}

type UserSearch struct {
	Username string
	Limit    int
	Offset   int
}
