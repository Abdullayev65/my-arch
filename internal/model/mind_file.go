package model

import (
	"mindstore/internal/object/model/submodel"
	"mindstore/pkg/hash-types"
)

type MindFile struct {
	MindId hash.Int
	FileId hash.Int
	submodel.BasicTimeModel
}
