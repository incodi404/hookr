package repositories

import (
	"context"

	"github.com/incodi404/hookr.git/pkg/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository interface {
	CreateUser(
		ctx context.Context,
		user_id pgtype.UUID,
		user_name pgtype.Text,
		email_id string,
		is_verified bool,
		created_at pgtype.Timestamptz,
		verified_at pgtype.Timestamptz,
	) (*sqlc.User, error)

	GetUserByUserId(
		ctx context.Context,
		user_id pgtype.UUID,
	) (*sqlc.User, error)

	GetUserByEmailId(
		ctx context.Context,
		email_id string,
	) (*sqlc.User, error)
}

// Implementation of AuthRepository
type PGAuthRepository struct {
	Db      *pgxpool.Pool
	Queries *sqlc.Queries
}

func NewPGAuthRepository(db *pgxpool.Pool, queries *sqlc.Queries) *PGAuthRepository {
	return &PGAuthRepository{
		Db:      db,
		Queries: queries,
	}
}

func (p *PGAuthRepository) CreateUser(
	ctx context.Context,
	user_id pgtype.UUID,
	user_name pgtype.Text,
	email_id string,
	is_verified bool,
	created_at pgtype.Timestamptz,
	verified_at pgtype.Timestamptz,
) (*sqlc.User, error) {
	createUserParams := sqlc.CreateUserParams{
		UserID:     user_id,
		UserName:   user_name,
		EmailID:    email_id,
		IsVerified: is_verified,
		CreatedAt:  created_at,
		VerifiedAt: verified_at,
	}

	// creating user
	user, err := p.Queries.CreateUser(ctx, createUserParams)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *PGAuthRepository) GetUserByUserId(
	ctx context.Context,
	user_id pgtype.UUID,
) (*sqlc.User, error) {

	user, err := p.Queries.GetUserByUserId(ctx, user_id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *PGAuthRepository) GetUserByEmailId(
	ctx context.Context,
	email_id string,
) (*sqlc.User, error) {

	user, err := p.Queries.GetUserByEmailId(ctx, email_id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
