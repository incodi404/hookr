## Migrate SQL in Golang

Installation: `scoop install migrate`

#### Commands:

Run all pending migrations:\
 `migrate -path ./migrations -database "postgres://postgres:password@localhost:5432/mydb?sslmode=disable" up`

Rollback the latest migration:\
`migrate -path ./migrations -database "postgres://postgres:password@localhost:5432/mydb?sslmode=disable" down 1`

Check current migration version:\
`migrate -path ./migrations -database "postgres://postgres:password@localhost:5432/mydb?sslmode=disable" version`
