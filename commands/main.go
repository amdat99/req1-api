package main

import (
    "os"
	"fmt"
	"req-api/migrations"
	"database/sql"
		_ "github.com/lib/pq"
	"sort"
	"github.com/fatih/color"
)


const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbA   = "req1"
)


func main() {
    arg1 := os.Args[1]

    if arg1 == "migrate" {

		// Connect to the database A
		DbA, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", host, port, user, password, dbA))
		if err != nil {
			panic(err)
		}
		var queries string = ""
		//loop through the migrations and add them to the queries array
		for _, query := range migrations.Migrations {
			queries += query
		}

		color.Green("Migrating tables...")

		_,err = DbA.Query(queries)
		if err != nil {
			fmt.Println("Error migrating tables", err)
			color.Red("Error migrating tables", err)
			panic(err)
		}else {
			color.Green("Successfully migrated "+fmt.Sprint(len(migrations.Migrations))+" tables")
		}
    }
    if arg1 == "rollback" {
		// Connect to the database A
		DbA, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", host, port, user, password, dbA))
		if err != nil {
			panic(err)
		}

		var rollbackQueries []string = migrations.Rollback
		ReverseSlice(rollbackQueries)

		var queries string = ""

		//loop through the migrations and add them to the queries array
		for _, query := range rollbackQueries{
			queries += query
		}
		color.Green("Rolling back tables...")
		_,err = DbA.Query(queries)
		if err != nil {
			color.Red("Failed to rollback tables", err)
			panic(err)
		}else {
			color.Green("Successfully rolled back "+fmt.Sprint(len(migrations.Migrations))+" tables")
		}
    }
}
func ReverseSlice[T comparable](s []T) {
    sort.SliceStable(s, func(i, j int) bool {
        return i > j
    })
}