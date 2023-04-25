package model

import (
	"mindstore/internal/object/model/submodel"
	"mindstore/pkg/hash-types"
	"time"
)

type Mind struct {
	submodel.BasicModel

	Topic     *string
	Caption   *string
	ParentId  *int
	Access    int
	HashedId  *hash.Int
	UpdatedAt time.Time
}

// access
// 33 - private
// 66 - friends - beta
// 99 - public
