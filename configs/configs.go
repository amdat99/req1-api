package configs

import (
	"database/sql"
	"fmt"
	"github.com/redis/go-redis/v9"
	_ "github.com/lib/pq"
	"net/smtp"
	"github.com/gofiber/fiber/v2"
	"os"
)




const (
	host     = os.Getenv("DB_HOST")
	port     = 5432
	user     = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	dbA   = os.Getenv("DB_DATABASE")
)


//DB is a map of databases
var DB map[string]*sql.DB = make(map[string]*sql.DB)

var Redis *redis.Client 

var MailAuth =  smtp.PlainAuth("", "sender@example.com", "password", "smtp.example.com")


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


func SetHeaders(c *fiber.Ctx) error {
  	c.Set("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Request-With")
	c.Set("Access-Control-Allow-Credentials", "true")
	c.Set("Content-Security-Policy", "upgrade-insecure-requests")
	c.Set("Cross-Origin-Resource-Policy", "same-site")
	c.Set("Cross-Origin-Embedder-Policy", "require-corp")
	c.Set("Origin-Agent-Cluster", "?1")
	c.Set("Referrer-Policy", "no-referrer")
	c.Set("Strict-Transport-Security", "max-age=15552000; includeSubDomains; preload")
	c.Set("X-Content-Type-Options", "nosniff")
	c.Set("X-Frame-Options", "SAMEORIGIN")
	c.Set("X-XSS-Protection", "0")
	c.Set("x-DNS-Prefetch-Control", "off")
	c.Set("x-download-options", "noopen")
	c.Set("x-permitted-cross-domain-policies", "none")
	return c.Next()
}