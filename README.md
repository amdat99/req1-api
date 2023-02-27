# req1-api

## Commands

Update the .env.example file with your database credentials and rename it to .env

To install the dependencies run `go mod download`

To run the api run `go run .`

To migrate the database run `go run commands/main.go migrate`

When starting the api for the first time, run `go run commands/main.go migrate`

To rollback the database run `go run commands/main.go rollback`

To rollback, migrate and seed and delete redis keys run `go run commands/main.go refresh`
