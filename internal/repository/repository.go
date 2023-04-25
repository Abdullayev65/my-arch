package repo

import (
	"mindstore/internal/db"
	"mindstore/internal/repository/file"
	"mindstore/internal/repository/mind"
	"mindstore/internal/repository/user"
)

var User = new(user.Repo)
var Mind = new(mind.Repo)
var File = new(file.Repo)

func init() {
	*User = *user.New(db.Sqlx)
	*Mind = *mind.New(db.Sqlx)
	*File = *file.New(db.Sqlx)
}
