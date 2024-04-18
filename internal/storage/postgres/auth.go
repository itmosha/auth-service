package storage

import (
	"github.com/itmosha/auth-service/pkg/postgres"
)

type StoragePostgres struct {
	cli *postgres.PostgresClient
}

func NewStoragePostgres(cli *postgres.PostgresClient) *StoragePostgres {
	return &StoragePostgres{cli}
}
