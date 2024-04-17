package storage

import (
	"github.com/itmosha/auth-service/pkg/postgres"
)

type AuthStoragePostgres struct {
	cli *postgres.PostgresClient
}

func NewAuthStoragePostgres(cli *postgres.PostgresClient) *AuthStoragePostgres {
	return &AuthStoragePostgres{cli}
}
