package model

import (
	"time"
)

type User struct {
	Id         int
	Username   string
	Email      *string
	MindId     *int
	Password   string
	FirstName  *string
	MiddleName *string
	LastName   *string
	BirthDate  *time.Time
	AvatarId   *int
	// hidden fields
	CreatedBy *int
	DeletedBy *int
	CreatedAt time.Time
	DeletedAt *time.Time
}
