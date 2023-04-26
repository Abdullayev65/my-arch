package model

import (
	"my-arch/internal/model/submodel"
)

type MindFile struct {
	MindId int
	FileId int
	submodel.BasicTimeModel
}
