
package utils

type Model[T any] struct {
    Data []T
}

type RouterType struct {
	Name string
	Session Session
}

type AddRowData struct {
    Id string `json:"id"`
}

type MultiSearchType struct {
	Like map[string]interface{} `json:"like"`
	Equal map[string]interface{} `json:"equal"`
}

type Session struct {
    Name string      `json:"name"`
    Email string     `json:"email"`
    Id int `json:"id"`
    Token string `json:"token"`
    Private_id string `json:"private_id"`
    Org_id string `json:"org_id"`
    Contact_id int `json:"contact_id"`
    Db string `json:"db"`
    Table_key string `json:"table_key"`
    Role string `json:"role"`
}
