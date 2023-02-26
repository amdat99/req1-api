package authModel

import "database/sql"
import "req-api/database"


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
	Orgs []map[string]interface{} `json:"orgs"`
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

//Perfomr uinsert user + 0rg_user transaction
func InsertUserTransaction(db *sql.DB , userQuery database.AddRowReturn ,orgUserQuery database.AddRowReturn) error {

	//Perform transaction with the two queries
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_ , err = tx.Exec(userQuery.Query, userQuery.Values...)
	if err != nil {
		tx.Rollback()
		return err
	}
	_ ,err = tx.Exec(orgUserQuery.Query, orgUserQuery.Values...)
	if err != nil {
		tx.Rollback()
		return err
	}	
	err = tx.Commit()

	if err != nil {
		return err 
	}
	return nil
}

func InsetContactOrgIdTransaction(db *sql.DB, body *UserType ) error{

	tx2, err := db.Begin()
	if err != nil {
		return err
	}
	//Insert org_id/private_id into org_ids table. This will be used as the primary key for the org_id field for the org data. Having it in a seperate table will facilitate on delete cascade for the speciic database the org data is stored in.
	_, err =  tx2.Exec("INSERT INTO org_ids (org_id) VALUES ($1)", body.Private_id)
	if err != nil {
		tx2.Rollback()
		return err
	}

	//Create a contact record for the user for the specific org
	var contactLabel string = body.First_name + " " + body.Last_name + " (" + body.Username + ")"
	_, err = tx2.Exec("INSERT INTO contact_"+body.Table_key+" (org_id, email, first_name, last_name, username, label, internal) VALUES ($1, $2, $3, $4, $5, $6, $7)", body.Private_id, body.Email, body.First_name, body.Last_name, body.Username, contactLabel,true)
	if err != nil {
		tx2.Rollback()
		return err
	}
	err = tx2.Commit()
	if err != nil {
		return err
	}
	return nil
}

