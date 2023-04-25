package submodel

import (
	"time"
)

type BasicModel struct {
	Id hash.Int `bun:"id,pk,autoincrement"`
	BasicTimeModel
	CreatedBy hash.Int  `bun:"created_by"`
	DeletedBy *hash.Int `bun:"deleted_by"`
}

type BasicTimeModel struct {
	CreatedAt time.Time  `bun:"created_at,default:now(),notnull"`
	DeletedAt *time.Time `bun:"deleted_at,soft_delete"`
}
