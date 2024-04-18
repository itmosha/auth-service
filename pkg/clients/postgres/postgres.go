package postgres

import (
	"fmt"

	"github.com/itmosha/auth-service/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresClient struct {
	DB *sqlx.DB
}

func NewPostgresClient(dbCfg *config.DB) (pgClient *PostgresClient, err error) {
	pgClient = &PostgresClient{}
	pgClient.DB, err = sqlx.Open("postgres",
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
			dbCfg.User, dbCfg.Pass, dbCfg.Host, dbCfg.Name))
	if err != nil {
		return
	}
	pgClient.DB.SetMaxIdleConns(dbCfg.MaxIdleConns)
	pgClient.DB.SetMaxOpenConns(dbCfg.MaxOpenConns)
	err = pgClient.DB.Ping()
	return
}
