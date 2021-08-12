package postgres

import (
	"database/sql"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func getPostgresURL() (url string) {
	return "postgresql://" +
		utils.Getenv("POSTGRES_USER", "postgres") + ":" +
		utils.Getenv("POSTGRES_PASSWORD", "postgres") + "@" +
		utils.Getenv("POSTGRES_HOST", "localhost") + ":" +
		utils.Getenv("POSTGRES_PORT", "5432") + "/" +
		utils.Getenv("POSTGRES_DB", "postgres")
}

func GetConnection() *bun.DB {
	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(getPostgresURL())))

	return bun.NewDB(sqlDb, pgdialect.New())
}
