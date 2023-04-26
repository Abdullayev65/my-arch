package repo

import (
	"my-arch/internal/pkg/db"
	"my-arch/internal/repository/file"
	"my-arch/internal/repository/user"
)

var User = new(user.Repo)
var File = new(file.Repo)

func init() {
	*User = *user.New(db.Sqlx)
	*File = *file.New(db.Sqlx)
}
