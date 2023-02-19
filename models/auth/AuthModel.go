package authModel

import "database/sql"


type UserType struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password,omitempty" validate:"required,min=8"`
	Username   string `json:"username" validate:"required"`
	First_name string `json:"first_name" validate:"required"`
	Last_name  string `json:"last_name" validate:"required"`
	Private_id string  `json:"private_id,omitempty"`
	Id int `json:"id"`
	Db string `json:"db,omitempty"`
	Table_key string `json:"table_key,omitempty"`
}

type LoginType struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type EnterOrgType struct {
	Name string `json:"name"`
	Id int `json:"id" validate:"required"`
	Email string `json:"email"`
	Org_name string `json:"org_name"`
	Org_id sql.NullString `json:"org_id"`
	Role string `json:"role"`
	Db string `json:"db"`
	Table_key string `json:"table_key"`
	Disabled bool `json:"disabled"`
	Personal_id sql.NullString `json:"personal_id"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
} 


// Format User f
func User(u *UserType) map[string]interface{} {
	return map[string]interface{}{
		"email":      u.Email,
		"first_name": u.First_name,
		"last_name":  u.Last_name,
		"password":   u.Password,
		"private_id": u.Private_id,
		"username":   u.Username,
	}
}


// Format Auth user
func OrgUser(u *UserType) map[string]interface{} {
	return map[string]interface{}{
		"name":  u.Username,
		"email": u.Email,
		"role": "admin",
		"db": u.Db,
		"table_key": u.Table_key,
		"personal_id": u.Private_id,	
	}
}

