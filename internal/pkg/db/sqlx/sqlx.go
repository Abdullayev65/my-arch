package sqlx

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx"
	"mindstore/pkg/config"
)

func New() (*sqlx.DB, error) {
	dbConfig := config.GetDB()
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s port=%s sslmode=disable",
		dbConfig.Username, dbConfig.DbName, dbConfig.Password, dbConfig.Port)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.MapperFunc(func(s string) string {
		return strcase.ToSnake(s)
	})

	return db, nil
}
