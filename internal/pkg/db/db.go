package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	sqlx2 "mindstore/internal/db/sqlx"
)

var Sqlx = new(sqlx.DB)

func init() {
	sqlxDB, err := sqlx2.New()
	if err != nil {
		panic(fmt.Errorf("err initioling sqlxdb: #%v", err))
	}

	*Sqlx = *sqlxDB
}
