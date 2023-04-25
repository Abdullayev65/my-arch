package mind

import (
	"mindstore/internal/object/dto/file"
	"mindstore/internal/object/dto/mind"
	"mindstore/pkg/ctx"
	"mindstore/pkg/hash-types"
)

type Mind interface {
	Create(c ctx.Ctx, input *mind.Create) (*hash.Int, error)
	Update(c ctx.Ctx, input *mind.Update) error
	ChildrenById(c ctx.Ctx, id *mind.ChildrenFilter, getOwnSelf bool) ([]mind.List, error)
}

type File interface {
	GetByMindIds(c ctx.Ctx, mindIds []hash.Int) (map[hash.Int][]file.List, error)
}
