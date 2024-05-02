package utils

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	orgn "github.com/sample-crud-app/repositories/organisation/models"
	"github.com/sample-crud-app/repositories/user/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)


var DB *bun.DB

func DBConnect() {
	dsn := "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	DB = bun.NewDB(pgdb, pgdialect.New())
	DB.NewCreateTable().Model((*models.User)(nil)).Exec(context.Background())
	DB.NewCreateTable().Model((*orgn.Organization)(nil)).Exec(context.Background())
}
