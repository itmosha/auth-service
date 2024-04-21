package storage

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/itmosha/auth-service/internal/entity"
	"github.com/itmosha/auth-service/pkg/clients/postgres"
	"github.com/lib/pq"
)

type StoragePostgres struct {
	cli *postgres.PostgresClient
}

func NewStoragePostgres(cli *postgres.PostgresClient) *StoragePostgres {
	return &StoragePostgres{cli}
}

func (s *StoragePostgres) Insert(ctx *context.Context, insertedAuthData *entity.AuthData) (createdAuthData *entity.AuthData, err error) {
	stmt, err := s.cli.PreparexContext(*ctx, `INSERT INTO "phonenumbers" ("phonenumber") VALUES ($1) RETURNING *;`)
	if err != nil {
		return
	}
	createdAuthData = &entity.AuthData{}
	err = stmt.QueryRowxContext(*ctx, insertedAuthData.Phonenumber).StructScan(createdAuthData)
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code == "23505" { // unique violation
			err = ErrPhonenumberAlreadyExist
		}
	}
	return
}

func (s *StoragePostgres) SelectByPhonenumber(ctx *context.Context, phonenumber string) (authData *entity.AuthData, err error) {
	stmt, err := s.cli.PreparexContext(*ctx, `SELECT * FROM "phonenumbers" WHERE "phonenumber" = $1 LIMIT 1;`)
	if err != nil {
		return
	}
	authData = &entity.AuthData{}
	err = stmt.QueryRowxContext(*ctx, phonenumber).StructScan(authData)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrPhonenumberNotFound
		}
	}
	return
}

func (s *StoragePostgres) DeleteUnregistered(ctx *context.Context, delta time.Duration) (err error) {
	stmt, err := s.cli.PreparexContext(*ctx, `DELETE FROM "phonenumbers" WHERE "is_registered" = false AND "created_at" < $1;`)
	if err != nil {
		return
	}
	_, err = stmt.ExecContext(*ctx, time.Now().Add(-delta))
	return
}
