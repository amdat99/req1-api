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
	"database/sql"
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

	fmt.Println("Registering user: ", body.Email)

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

	//Create transaction dor inserting user and org_user records for the main db
	err = authModel.InsertUserTransaction(configs.DB["A"],userQuery, orgUserQuery)
	if err != nil {
		fmt.Println("Error inserting user: ", err)
		return utils.UniqueError(c, err.Error(), "Email/Username", "Error registering")
	}

	//Create transaction with org db (currently is A, but will be different if new org data sjould be stored in a different db)
	err = authModel.InsetContactOrgIdTransaction(configs.DB[body.Db],body);
	if err != nil {
		fmt.Println("Error inserting user: ", err)
		err := deleteUserOnFail(configs.DB["A"], c, body.Email)
		if err != nil {
			return  nil
		}
		return utils.ResError(c, "Error registering", 500)
	}

	utils.ResSuccess(c)
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
		return utils.ResError(c, "Invalid credentials", 400)
	}
	//Get org users record for user
	orgs,err := database.SelectRows(configs.DB["A"], "org_user", c, []string{"name","id","org_name"}, map[string]interface{}{"email": user.Email},100)
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
 	user.Orgs = orgs

	//Return user object
	return utils.ResJSON(c,user)
}

//Logout a user, deletes the session from redis
func Logout(c *fiber.Ctx) error {

 	cookie := c.Cookies("Req1S")
    if cookie == "" {
        return utils.ResError(c, "Invalid session", 400)
    }

	err := configs.Redis.Del(c.Context(), cookie).Err()
	if err != nil {
		return utils.ResError(c, "Error logging out", 500)
	}
	c.Cookies(cookie, "")
	utils.ResSuccess(c)
	return nil
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


	//Delete user and org_user records if query fails
func deleteUserOnFail(db *sql.DB, c *fiber.Ctx, email string) error {
	_, err := db.Exec("DELETE FROM users WHERE email = $1", email)
	if err != nil {
		fmt.Println("Error deleting user on fail: ", err)
		return utils.ResError(c, "Error registering", 500)
	}
	_, err = db.Exec("DELETE FROM org_user WHERE email = $1", email)
	if err != nil {
		fmt.Println("Error deleting org_user on fail: ", err)
		return utils.ResError(c, "Error registering", 500)
	}
	return nil
}
