package db

import (
	"database/sql"
	_ "fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"hvalfangst/golang-gin-api-with-bun/src/configuration"
)

func CreateDB(config configuration.Db) (*bun.DB, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(config.DataSourceName)))
	sqldb.SetMaxIdleConns(config.MaxIdleConns)
	sqldb.SetMaxOpenConns(config.MaxOpenConns)
	db := bun.NewDB(sqldb, pgdialect.New())
	return db, nil
}
