package user

import (
	"mime/multipart"
	"my-arch/internal/dto"
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
	Id           int `json:"-" form:"-"`
	Username     *string
	Email        *string
	Password     *string
	FirstName    *string
	MiddleName   *string
	LastName     *string
	Avatar       *multipart.FileHeader
	AvatarId     *int
	BirthDate    *time.Time `json:"-" form:"-"`
	BirthDateStr *string    `json:"birth_date" form:"birth_date"`
}

type UserDetail struct {
	Id           int
	Username     *string
	Email        *string
	MindId       *int
	FirstName    *string
	MiddleName   *string
	LastName     *string
	AvatarUrl    *string
	AvatarId     *int       `json:"-"`
	BirthDate    *time.Time `json:"-"`
	BirthDateStr *string    `json:"birth_date"`
}

type UserList struct {
	Id         int
	Username   *string
	MindId     *int
	FirstName  *string
	MiddleName *string
	LastName   *string
}

type Filter struct {
	dto.Filter

	Username *string
	Email    *string
}

type Delete struct {
	dto.Filter

	Username *string
	Email    *string
}

type UserSearch struct {
	Username string
	Limit    int
	Offset   int
}
