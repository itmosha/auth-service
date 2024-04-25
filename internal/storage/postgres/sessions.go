package storage

import (
	"context"

	"github.com/itmosha/auth-service/internal/entity"
	"github.com/itmosha/auth-service/pkg/clients/postgres"
)

type SessionsStoragePostgres struct {
	cli *postgres.PostgresClient
}

func NewSessionsStoragePostgres(cli *postgres.PostgresClient) *SessionsStoragePostgres {
	return &SessionsStoragePostgres{cli}
}

func (s *SessionsStoragePostgres) Insert(ctx *context.Context, insertedSession *entity.Session) (createdSession *entity.Session, err error) {
	stmt, err := s.cli.PreparexContext(*ctx, `INSERT INTO "sessions" ("user_uid", "token", "issued_at", "expires_at") VALUES ($1, $2, $3, $4) RETURNING *;`)
	if err != nil {
		return
	}
	createdSession = &entity.Session{}
	err = stmt.QueryRowxContext(*ctx, insertedSession.UserUid, insertedSession.Token, insertedSession.IssuedAt, insertedSession.ExpiresAt).StructScan(createdSession)
	return
}
