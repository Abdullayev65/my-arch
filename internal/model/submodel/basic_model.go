package submodel

import (
	"time"
)

type BasicModel struct {
	Id int `bun:"id,pk,autoincrement"`
	BasicTimeModel
	CreatedBy int  `bun:"created_by"`
	DeletedBy *int `bun:"deleted_by"`
}

type BasicTimeModel struct {
	CreatedAt time.Time  `bun:"created_at,default:now(),notnull"`
	DeletedAt *time.Time `bun:"deleted_at,soft_delete"`
}
