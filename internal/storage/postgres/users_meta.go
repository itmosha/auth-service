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

type UsersMetaStoragePostgres struct {
	cli *postgres.PostgresClient
}

func NewUsersMetaStoragePostgres(cli *postgres.PostgresClient) *UsersMetaStoragePostgres {
	return &UsersMetaStoragePostgres{cli}
}

func (s *UsersMetaStoragePostgres) Insert(ctx *context.Context, insertedUserMeta *entity.UserMeta) (createdUserMeta *entity.UserMeta, err error) {
	stmt, err := s.cli.PreparexContext(*ctx, `INSERT INTO "users_meta" ("phonenumber") VALUES ($1) RETURNING *;`)
	if err != nil {
		return
	}
	createdUserMeta = &entity.UserMeta{}
	err = stmt.QueryRowxContext(*ctx, insertedUserMeta.Phonenumber).StructScan(createdUserMeta)
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code == "23505" { // unique violation
			err = common.ErrUserMetaAlreadyExist
		}
	}
	return
}

func (s *UsersMetaStoragePostgres) SelectByPhonenumber(ctx *context.Context, phonenumber string) (userMeta *entity.UserMeta, err error) {
	stmt, err := s.cli.PreparexContext(*ctx, `SELECT * FROM "users_meta" WHERE "phonenumber" = $1 LIMIT 1;`)
	if err != nil {
		return
	}
	userMeta = &entity.UserMeta{}
	err = stmt.QueryRowxContext(*ctx, phonenumber).StructScan(userMeta)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = common.ErrUserMetaNotFound
		}
	}
	return
}

func (s *UsersMetaStoragePostgres) SelectByUid(ctx *context.Context, uid string) (userMeta *entity.UserMeta, err error) {
	stmt, err := s.cli.PreparexContext(*ctx, `SELECT * FROM "users_meta" WHERE "uid" = $1 LIMIT 1;`)
	if err != nil {
		return
	}
	userMeta = &entity.UserMeta{}
	err = stmt.QueryRowxContext(*ctx, uid).StructScan(userMeta)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = common.ErrUidNotFound
		}
	}
	return
}

func (s *UsersMetaStoragePostgres) UpdateFields(ctx *context.Context, uid string, fields map[string]interface{}) (err error) {
	query := `UPDATE "users_meta" SET uid = uid`
	for field := range fields {
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

func (s *UsersMetaStoragePostgres) DeleteUnregistered(ctx *context.Context, delta time.Duration) (err error) {
	stmt, err := s.cli.PreparexContext(*ctx, `DELETE FROM "users_meta" WHERE "is_registered" = false AND "created_at" < $1;`)
	if err != nil {
		return
	}
	_, err = stmt.ExecContext(*ctx, time.Now().Add(-delta))
	return
}
