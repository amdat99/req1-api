
package utils

import 	"golang.org/x/crypto/bcrypt"
import 	"crypto/rand"
import 	"encoding/base64"
import 	"time"
import   "github.com/gofiber/fiber/v2"
import   "strconv"
import "req-api/configs"
import  "encoding/json"


func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

func ComparePassword(password string, hashedPassword string) error {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    if err != nil {
        return err
    }
    return nil
}


//___Session___ 

//Generate Cookie 
func GenerateSession(c *fiber.Ctx, id int ) (string, error) {
    b,err :=  rand.Prime(rand.Reader, 128)
    if err != nil {
        return "", err
    }

    token := "req1:" + strconv.Itoa(id) + ":" + base64.StdEncoding.EncodeToString(b.Bytes())

    cookie := new(fiber.Cookie)
    cookie.Name = "Req1S"
    cookie.Value = token
    cookie.Expires = time.Now().Add(24 * time.Hour) // 1 day
    cookie.HTTPOnly = true
    cookie.SameSite = "lax"
    
    c.Cookie(cookie)

    return token, nil
}

//Get session from cookie
func GetSession(c *fiber.Ctx) (Session, error) {
    sessionStruct := Session{}

    cookie := c.Cookies("Req1S")
    if cookie == "" {
        return sessionStruct, NewErr("No cookie")
    }
    
    //Query redis for session
	val, err := configs.Redis.Get(c.Context(), cookie).Result()
	if err != nil {
		return sessionStruct, err
	}

	//Unmarshal session
	err = json.Unmarshal([]byte(val), &sessionStruct)
	if err != nil {
		return sessionStruct, err
	}

	return sessionStruct, nil
}

//Delete session from cookie
func DeleteSession(c *fiber.Ctx) error {
    cookie := c.Cookies("Req1S")
    if cookie == "" {
        return NewErr("No cookie")
    }
    
    //Delete session from redis
    err := configs.Redis.Del(c.Context(), cookie).Err()
    if err != nil {
        return err
    }

    return nil
}