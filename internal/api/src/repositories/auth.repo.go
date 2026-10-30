package repositories

import "context"

type AuthRepository interface {
	CreateUser(
		ctx context.Context,
		user_id string,
		user_name string,
		email_id string,
	)

	GetUserByUserId(
		ctx context.Context,
		user_id string,
	)

	GetUserByEmailId(
		ctx context.Context,
		email_id string,
	)
}

// Implementation of AuthRepository
type PGAuthRepository interface {
}
