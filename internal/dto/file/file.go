package file

import (
	"mime/multipart"
	"mindstore/pkg/hash-types"
)

type List struct {
	Id       hash.Int
	MindId   hash.Int `json:",omitempty"`
	Url      string
	Name     string
	HashedId *int
	Access   int
	Size     int64
}

type CreateWithMind struct {
	Files     []*multipart.FileHeader
	CreatedBy *hash.Int `form:"-"`
	MindId    *hash.Int
	Access    int
}

type MindFile struct {
	MindId hash.Int
	FileId hash.Int
}

type DeleteMind struct {
	MindId    hash.Int
	FileId    hash.Int
	UserId    hash.Int
	DeletedBy hash.Int
}
