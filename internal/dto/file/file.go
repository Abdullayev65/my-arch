package file

import (
	"mime/multipart"
)

type List struct {
	Id       int
	MindId   int `json:",omitempty"`
	Url      string
	Name     string
	HashedId *int
	Access   int
	Size     int64
}

type Detail struct {
	Id       int
	MindId   int `json:",omitempty"`
	Url      string
	Name     string
	HashedId *int
	Access   int
	Size     int64
}

type Update struct {
	Id       int
	MindId   int `json:",omitempty"`
	Url      string
	Name     string
	HashedId *int
	Access   int
	Size     int64
}

type CreateWithMind struct {
	Files     []*multipart.FileHeader
	CreatedBy *int `form:"-"`
	MindId    *int
	Access    int
}
type Create struct {
	Files     []*multipart.FileHeader
	CreatedBy *int `form:"-"`
	MindId    *int
	Access    int
}

type MindFile struct {
	MindId int
	FileId int
}

type DeleteMind struct {
	MindId    int
	FileId    int
	UserId    int
	DeletedBy int
}
type Delete struct {
	MindId    int
	FileId    int
	UserId    int
	DeletedBy int
}

type Filter struct {
	MindId    int
	FileId    int
	UserId    int
	DeletedBy int
}
