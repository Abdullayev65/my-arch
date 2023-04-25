package mind

import (
	"mindstore/internal/object/dto/file"
)

type Create struct {
	Topic     *string
	Caption   *string
	ParentId  *hash.Int
	Access    int
	HashedId  *hash.Int
	CreatedBy *hash.Int `json:"-"`
}

type Update struct {
	Id        hash.Int `json:"-"`
	Topic     *string
	Caption   *string
	ParentId  *hash.Int
	Access    *int
	HashedId  *hash.Int
	CreatedBy *hash.Int `json:"-"`
}

type List struct {
	Id       hash.Int
	Topic    *string
	Caption  *string
	ParentId *hash.Int
	Access   int
	HashedId *hash.Int
	Files    []file.List
	Children []List
}

type ChildrenFilter struct {
	MindId hash.Int
	UserId *hash.Int
}
