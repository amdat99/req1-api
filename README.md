# req1-api


## Commands

To run the api run  `go run .`


To migrate the database run  `go run commands/main.go migrate`

To rollback the database run  `go run commands/main.go rollback`

To rollback, migrate and seed and delete redis keys run  `go run commands/main.go refresh`