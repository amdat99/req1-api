package configs

import (
	"database/sql"
	"fmt"
	"github.com/redis/go-redis/v9"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbA   = "req1"
)


//DB is a map of databases
var DB map[string]*sql.DB = make(map[string]*sql.DB)


var Redis *redis.Client 

func Init() {

	//Connect to the database A
	DbA, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", host, port, user, password, dbA))
	if err != nil {
		panic(err)
	}
	//Add database A to the map of databases
	DB["A"] = DbA
	DB["A"].SetMaxOpenConns(10)


	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	fmt.Println("Successfully connected!")
}
