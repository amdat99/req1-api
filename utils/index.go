package utils

import (
	"fmt"
	"strings"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"encoding/json"
	"errors"
)

func Router (c *fiber.Ctx, authorise bool)  RouterType {
	var method = c.Method()
	var secPath = strings.Split(c.Path(), "/")[2]

	//Check if route is authorised
	if authorise {
    	session, err := GetSession(c)
		if err != nil {
			return RouterType{ Name: "unauthorised"}
		}
		return RouterType{ Name: method + "_" + secPath, Session: session}
	} 

	return RouterType{ Name: method + "_" + secPath}
}


//Set headerss
func setHeaders (c *fiber.Ctx) {
	c.Set("Access-Control-Allow-Origin", "*")
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

}

//Return error on
func ResError(c *fiber.Ctx, message string, status int ) error {
	setHeaders(c)
	return c.Status(status).JSON(fiber.Map{
		"errors" : true,
		"message": message,
	})
}

func UniqueError(c *fiber.Ctx, err string, field string, message string ) error {
//Check if error is postgres unique error
	setHeaders(c)
	if strings.Contains(err, "duplicate key value violates unique constraint") {
		return ResError(c, field + " already exists", 400)
	}

	return ResError(c, message, 500)
}

//Body error
func ResBodyError(c *fiber.Ctx) error {
	setHeaders(c)
	return c.Status(400).JSON(fiber.Map{
		"errors" : true,
		"message": "Invalid JSON",
	})
}

func ResAuthError(c *fiber.Ctx) error {
	setHeaders(c)
	return c.Status(401).JSON(fiber.Map{
		"errors" : true,
		"message": "Unauthorised",
	})
}

//Route error
func ResRouteError(c *fiber.Ctx,name any) error {
	setHeaders(c)
	if(name == "unauthorised"){
		return ResAuthError(c)
	}
	return c.Status(404).JSON(fiber.Map{
		"errors" : true,
		"message": "Route not found with this method",
	})
}

//Return JSON 
func ResJSON(c *fiber.Ctx, data any) error {
	setHeaders(c)
	return c.Status(200).JSON(data)
}


func ResAddSuccess(c *fiber.Ctx, data AddRowData) error {
	setHeaders(c)
	return c.Status(200).JSON(fiber.Map{
		"success" : true,
		"id": data.Id,
	})
}

func ResSuccess(c *fiber.Ctx) error {
	setHeaders(c)
	return c.Status(200).JSON(fiber.Map{
		"success" : true,
	})
}

//Format body from bytes to Struct type
func ParseBody(c *fiber.Ctx,body any) any{
	if err := c.BodyParser(body); err != nil {
		fmt.Println(err)
		ResBodyError(c)
		return nil	
	}
	return body

}

//Validate body
func ValidateBody(c *fiber.Ctx, body any) error {
	validate := validator.New()
	err := validate.Struct(body)
	if (err != nil) {
		fmt.Println(err)
		c.Status(400).JSON(fiber.Map{
			"errors" : true,
			"message": err.Error(),
		})
		return nil
	}
	return c.BodyParser(body)

}

func UUID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

func MapToJsonString(m map[string]interface{}) (string, error) {
    jsonBytes, err := json.Marshal(m)
    if err != nil {
        return "", err
    }
    return string(jsonBytes), nil
}

func StructToJsonString(s interface{}) (string, error) {
    jsonBytes, err := json.Marshal(s)
    if err != nil {
        return "", err
    }
    return string(jsonBytes), nil
}

func NewErr(a string) error {
    return errors.New(a)
}

func JsonArr(field string) string {
 if(field == "") {return "[]"} else {return field}
}

func JsonObj(field string) string {
 if(field == "") {return "{}"} else {return field}
}

var EmptyIntfMap = map[string]interface{}{}

