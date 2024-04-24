package storage

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/itmosha/auth-service/internal/entity"
	common "github.com/itmosha/auth-service/internal/storage"
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
			err = common.ErrPhonenumberAlreadyExist
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
			err = common.ErrPhonenumberNotFound
		}
	}
	return
}

func (s *StoragePostgres) SelectByUid(ctx *context.Context, uid string) (authData *entity.AuthData, err error) {
	stmt, err := s.cli.PreparexContext(*ctx, `SELECT * FROM "phonenumbers" WHERE "uid" = $1 LIMIT 1;`)
	if err != nil {
		return
	}
	authData = &entity.AuthData{}
	err = stmt.QueryRowxContext(*ctx, uid).StructScan(authData)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = common.ErrUidNotFound
		}
	}
	return
}

func (s *StoragePostgres) UpdateFields(ctx *context.Context, uid string, fields map[string]interface{}) (err error) {
	query := `UPDATE "phonenumbers" SET uid = uid`
	for field, _ := range fields {
		query += `, ` + field + ` = :` + field
	}
	query += ` WHERE uid = :uid;`
	values := fields
	values["uid"] = uid

	stmt, err := s.cli.PrepareNamedContext(*ctx, query)
	if err != nil {
		return
	}
	res, err := stmt.ExecContext(*ctx, values)
	if err != nil {
		return
	}
	cntRows, err := res.RowsAffected()
	if err != nil {
		return
	} else if cntRows == 0 {
		err = common.ErrUidNotFound
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
