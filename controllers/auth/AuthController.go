package authcontroller

import (
	"req-api/configs"
	authModel "req-api/models/auth"
	"req-api/utils"
	"req-api/database"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"fmt"
	"strconv"
)

var validate = validator.New()

//Register a new user, creates new user, a peosnal org for that user and a contact record for that org for the user
func Register(c *fiber.Ctx) error {

	var body *authModel.UserType = new(authModel.UserType)
	utils.ParseBody(c, body)
	err := validate.Struct(body)
	if err != nil {
		return utils.ResError(c, err.Error(), 400)
	}

	//Hash password
	body.Password, err = utils.HashPassword(body.Password)
	if err != nil {
		return utils.ResError(c, "Error registering", 500)
	}
	body.Private_id = utils.UUID()
	body.Db = "A"  //New org data will be stored in this database
	body.Table_key = "1" //New org data will have this table key ending

	var userQuery database.AddRowReturn = database.AddRowQuery("users", authModel.User(body))
	var orgUserQuery database.AddRowReturn = database.AddRowQuery("org_user", authModel.OrgUser(body))
	var results []interface{}

	//Perform transaction with the two queries
	tx, err := configs.DB["A"].Begin()
	if err != nil {
		return utils.ResError(c, "Error registering", 500)
	}
	res , err := tx.Exec(userQuery.Query, userQuery.Values...)
	if err != nil {
		tx.Rollback()
		return utils.UniqueError(c, err.Error(), "Email/Username", "Error registering")
	}
	results = append(results, res)
	res , err = tx.Exec(orgUserQuery.Query, orgUserQuery.Values...)
	if err != nil {
		tx.Rollback()
		return utils.ResError(c, "Error registering", 500)
	}	
	results = append(results, res)
	err = tx.Commit()

	if err != nil {
		return utils.ResError(c, "Error", 500)
	}
	//Create transaction with org db (currently is A, but will be different if new org data sjould be stored in a different db)
	tx2, err := configs.DB["A"].Begin()
	if err != nil {
		return utils.ResError(c, "Error registering", 500)
	}

	//Insert org_id/private_id into org_ids table. This will be used as the primary key for the org_id field for the org data. Having it in a seperate table will facilitate on delete cascade for the speciic database the org data is stored in.
	res, err = tx2.Exec("INSERT INTO org_ids (org_id) VALUES ($1)", body.Private_id)
	if err != nil {
		tx2.Rollback()
		deleteUserOnFail(c, body.Email)  //Delete user and org_user records if query fails
		return utils.ResError(c, "Error registering", 500)
	}
	results = append(results, res)

	//Create a contact record for the user for the specific org
	var contactLabel string = body.First_name + " " + body.Last_name + " (" + body.Username + ")"
	res, err = tx2.Exec("INSERT INTO contact_"+body.Table_key+" (org_id, email, first_name, last_name, username, label, internal) VALUES ($1, $2, $3, $4, $5, $6, $7)", body.Private_id, body.Email, body.First_name, body.Last_name, body.Username, contactLabel,true)
	if err != nil {
		tx2.Rollback()
		deleteUserOnFail(c, body.Email)
		return utils.ResError(c, "Error registering", 500)
	}
	results = append(results, res)
	err = tx2.Commit()
	if err != nil {
		deleteUserOnFail(c, body.Email) //Delete user and org_user records if query fails
		return utils.ResError(c, "Error registering", 500)
	}

	utils.ResSuccess(c)

	return nil
}

//Delete user and org_user records if query fails
func deleteUserOnFail(c *fiber.Ctx, email string) error {
	_, err := configs.DB["A"].Exec("DELETE FROM users WHERE id = $1", email)
	if err != nil {
		fmt.Println("Error deleting user on fail: ", err)
		return utils.ResError(c, "Error registering", 500)
	}
	_, err = configs.DB["A"].Exec("DELETE FROM org_user WHERE email = $1", email)
	if err != nil {
		fmt.Println("Error deleting org_user on fail: ", err)
		return utils.ResError(c, "Error registering", 500)
	}
	return nil
}


//Login a user, returns user data, and the orgs the user is a part of
func Login(c *fiber.Ctx) error {

	var body *authModel.LoginType = new(authModel.LoginType)

	utils.ParseBody(c, body)
	err := validate.Struct(body)
	if err != nil {
		return utils.ResError(c, err.Error(), 400)
	}

	var user authModel.UserType
	err = configs.DB["A"].QueryRow("SELECT email, password, username, first_name, last_name, id, private_id FROM users WHERE email = $1", body.Email).Scan(&user.Email, &user.Password, &user.Username, &user.First_name, &user.Last_name, &user.Id, &user.Private_id)
	if err != nil || user.Email == "" {
		return utils.ResError(c, "Invalid credentials", 400)
	}

	if err := utils.ComparePassword(body.Password, user.Password); err != nil {
		//fmt.Println("Error: ", err)
		return utils.ResError(c, "Invalid credentials", 400)
	}

	//Get org users for user
 	var results []map[string]interface{}
	 	err = database.Select(configs.DB["A"], "org_user", []string{"name","id","org_name"},&results,c, 100, map[string]interface{}{"email": body.Email}, map[string]interface{}{})
	 	if err != nil {
	 		return utils.ResError(c, err.Error(), 400)
 	}	

	//Generate session token
	token,err := utils.GenerateSession(c, user.Id)
	if err != nil {
		return utils.ResError(c, "Error Logging in", 500)
	}

	//Create json string for the session object
	sessionString, err := utils.MapToJsonString(map[string]interface{}{
		"email": user.Email,
		"name":  user.First_name + " " + user.Last_name,
		"id":    user.Id,
		"private_id": user.Private_id,
		"token": token,
	})
	if err != nil {
		return utils.ResError(c, "Error Logging in", 500)
	}

	//Add json session to redis
	err = configs.Redis.Set(c.Context(), token, sessionString, 0).Err()
	if err != nil {
		//fmt.Println("Error: ", err)
		c.Cookies(token, "")
		return utils.ResError(c, "Error Logging in", 500)
	}
	//Delete password from user
	user.Password = ""
	user.Private_id = ""

	//Return user object
	return utils.ResJSON(c, map[string]interface{}{
		"user": user,
		"orgs": results,
	})
}

	//EnterOrg, quueries the specific org the user wants to enter, updates the user session with the org_id,role, db, table_key for the org
	func EnterOrg(c *fiber.Ctx, session utils.Session) error {

		id, err := strconv.Atoi(c.Query("id"))
    	if err != nil {
			return utils.ResError(c, "Invalid id", 400)
		}

		var orgUser authModel.EnterOrgType 
		orgUser.Id = id

		err = configs.DB["A"].QueryRow("SELECT * FROM org_user WHERE email = $1 AND id = $2", session.Email, orgUser.Id).Scan(&orgUser.Id,&orgUser.Name, &orgUser.Email, &orgUser.Org_id, &orgUser.Org_name, &orgUser.Role, &orgUser.Db, &orgUser.Table_key, &orgUser.Disabled, &orgUser.Personal_id, &orgUser.Created_at, &orgUser.Updated_at)
		if err != nil {
			return utils.ResError(c, "Error", 500)
		}

		if(orgUser.Disabled){
			return utils.ResError(c, "Your account has been disabled in this org", 400)
		} 

    	if orgUser.Org_id.Valid {
		session.Org_id = orgUser.Org_id.String
    	} else {
        session.Org_id = orgUser.Personal_id.String
    	}	
				
		//If org user found query the contact-record id for this org
		var contactId int
		err = configs.DB["A"].QueryRow("SELECT id FROM contact_"+orgUser.Table_key+" WHERE org_id = $1 AND email = $2", session.Org_id, session.Email).Scan(&contactId)
		if err != nil {
			return utils.ResError(c, "Error", 500)
		}
		fmt.Println("Contact id: ", contactId)
		session.Contact_id = contactId

		if orgUser.Org_name != "" {
			session.Name = orgUser.Org_name
		} else {
			session.Name = orgUser.Name
		}

		session.Db = orgUser.Db
		session.Table_key = orgUser.Table_key
		session.Role = orgUser.Role

		//Update session
		sessionString, err := utils.StructToJsonString(session)
		if err != nil {
			return utils.ResError(c, "Error entering org", 500)
		}
		err = configs.Redis.Set(c.Context(), session.Token, sessionString, 0).Err()
		if err != nil {
			return utils.ResError(c, "Error entering org", 500)
		}

		return utils.ResSuccess(c)
	}



