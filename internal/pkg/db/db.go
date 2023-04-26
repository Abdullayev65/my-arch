package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	sqlxDb "my-arch/internal/pkg/db/sqlx"
)

var Sqlx = new(sqlx.DB)

func init() {
	sqlxDB, err := sqlxDb.New()
	if err != nil {
		panic(fmt.Errorf("err initioling sqlxdb: #%v", err))
	}

	*Sqlx = *sqlxDB
}
