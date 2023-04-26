package model

import (
	"my-arch/internal/model/submodel"
	"time"
)

type Mind struct {
	submodel.BasicModel

	Topic     *string
	Caption   *string
	ParentId  *int
	Access    int
	HashedId  *int
	UpdatedAt time.Time
}

// access
// 33 - private
// 66 - friends - beta
// 99 - public
