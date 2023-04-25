package model

import (
	"mindstore/pkg/hash-types"
	"time"
)

type User struct {
	Id         hash.Int
	Username   string
	Email      *string
	MindId     *int
	Password   string
	FirstName  *string
	MiddleName *string
	LastName   *string
	BirthDate  *time.Time
	AvatarId   *hash.Int
	// hidden fields
	CreatedBy *int
	DeletedBy *int
	CreatedAt time.Time
	DeletedAt *time.Time
}
