package main

import (
    "os"
	"fmt"
	"req-api/migrations"
	"database/sql"
		_ "github.com/lib/pq"
	"sort"
	"github.com/fatih/color"
	"github.com/redis/go-redis/v9"
	"context"
	"time"
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
		err := migrateTables()
		if err != nil {
			panic(err)
		}
    }
    if arg1 == "rollback" {
		err := rollbackTable()
		if err != nil {
			panic(err)
		}
    }

	if arg1 == "refresh" {
		err := rollbackTable()
		if err != nil {
			panic(err)
		}
		err = migrateTables()
		if err != nil {
			panic(err)
		}

		err = DeleteAllRedisKeys()
		if err != nil {
			panic(err)
		}
	}

}



func migrateTables() error  {

	// Connect to the database A
	DbA, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", host, port, user, password, dbA))
	if err != nil {
		DbA.Close()
		return err
	}
	var queries string = ""
	//loop through the migrations and add them to the queries array
	for _, query := range migrations.Migrations {
		queries += query
	}

	color.Green("Migrating tables...")
	color.Yellow(queries)
	_,err = DbA.Query(queries)
	if err != nil {
		color.Red("Error migrating tables", err)
		DbA.Close()
		return err
	 }else {
		color.Green("Successfully migrated "+fmt.Sprint(len(migrations.Migrations))+" tables")
	}

	DbA.Close()
	return  nil
}

func rollbackTable() error {
	// Connect to the database A
	DbA, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", host, port, user, password, dbA))
	if err != nil {
		DbA.Close()
		return err
	}

	rollbackQueries :=  migrations.Rollback
	
	ReverseSlice(rollbackQueries)

	

	fmt.Println(rollbackQueries)

	var queries string = ""

	//loop through the migrations and add them to the queries array
	for _, query := range rollbackQueries{
		queries += query
	}
	
	color.Green("Rolling back tables...")
	color.Yellow(queries)
	_,err = DbA.Query(queries)
	if err != nil {
		color.Red("Failed to rollback tables", err)
		DbA.Close()
		return err
	}else {
		color.Green("Successfully rolled back "+fmt.Sprint(len(migrations.Migrations))+" tables")
		}	
	
	DbA.Close()
	return nil
}


//Delete all redis keys
func DeleteAllRedisKeys() error {

	//connect to redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	color.Green("Deleting redis keys...")

	context, _ := context.WithTimeout(context.Background(), 5*time.Second)

	//Delete all keys with the prefix "req"
	_, err := client.Del(context, "req*").Result()
	if err != nil {
		client.Close()
		color.Red("Failed to delete redis keys", err)
		return err
	}

	client.Close()
	color.Green("Successfully deleted redis keys")
	return nil
}


func ReverseSlice[T comparable](s []T) {
    sort.SliceStable(s, func(i, j int) bool {
        return i > j
    })
}