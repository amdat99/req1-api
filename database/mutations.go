package database

import (
	"fmt"
	"strconv"
	"strings"
	"database/sql"
	"github.com/gofiber/fiber/v2"
    "req-api/utils"
)

//___Add Row_____
//Add a row in a table, return the id of the row, takes in a map of column name and value
func AddRow(db *sql.DB, tableName string, session utils.Session, data map[string]interface{}) (utils.AddRowData, error) {
    var columns []string
    var values []interface{}

    //Add created by field to data 
    data["created_by"] = session.Contact_id

    for columnName, value := range data {
        columns = append(columns, columnName)
        values = append(values, value)
    }

    var placeholders []string
    for i := 0; i < len(values); i++ {
        placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
    }

    tableKeyName := tableName + "_" + session.Table_key
    query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id", tableKeyName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
    var id int64
    err := db.QueryRow(query, values...).Scan(&id)
    if err != nil {
        return utils.AddRowData{}, err
    }

    return utils.AddRowData{Id: strconv.FormatInt(id, 10)}, nil
}

//MultiAddRow, add multiple rows in a table, return the id of the row, takes in a map of column name and value
func MultiAddRow(db *sql.DB, tableName string, session utils.Session, dataSlice []map[string]interface{}) (utils.AddRowsData, error) {
    

    var columns []string
    var values [][]interface{}

    for _, data := range dataSlice {
        //Add created by field to data 
        data["created_by"] = session.Contact_id

        var rowValues []interface{}
        for _, value := range data {
            rowValues = append(rowValues, value)
        }
        values = append(values, rowValues)
    }

    for columnName := range dataSlice[0] {
        columns = append(columns, columnName)
    }

    var placeholders []string
    var ids []string
    for i := 0; i < len(columns); i++ {
        placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
    }

    tableKeyName := tableName + "_" + session.Table_key

    tx, err := db.Begin()
    if err != nil {
        return utils.AddRowsData{}, err
    }
    defer tx.Rollback()



    for _, rowValues := range values {
        query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id", tableKeyName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
        var id int64
        err := tx.QueryRow(query, rowValues...).Scan(&id)
        if err != nil {
            tx.Rollback()
            return utils.AddRowsData{}, err
        }
        ids = append(ids, strconv.FormatInt(id, 10))
    }

    err = tx.Commit()
    if err != nil {
        return utils.AddRowsData{}, err
    }

    return utils.AddRowsData{Ids: ids}, nil
}



//retrn onterface of query and values
func AddRowQuery(table string, data map[string]interface{}) AddRowReturn {
    var columns []string
    var values []interface{}

    for columnName, value := range data {
        columns = append(columns, columnName)
        values = append(values, value)
    }

    var placeholders []string
    for i := 0; i < len(values); i++ {
        placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
    }

    query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id", table, strings.Join(columns, ", "), strings.Join(placeholders, ", "))

    return AddRowReturn{Query: query, Values: values}
}

type AddRowReturn struct {
    Query  string
    Values []interface{}
}

//___Update Row_____

//Update a row in a table, return the id of the row, takes in a map of column name and value
func UpdateRow(db *sql.DB, tableName string,session utils.Session, data map[string]interface{},id int64) error {

    if(session.Org_id == ""){
        return utils.NewErr("No org_id")
    } 

    if(id == 0){
        return utils.NewErr("No id")
    }

    
    var columns []string
    var values []interface{}

    for columnName, value := range data {
        columns = append(columns, columnName)
        values = append(values, value)
    }

    var placeholders []string
    for i := 0; i < len(values); i++ {
        placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
    }

    values = append(values, id)

    tableKeyName := tableName + "_" + session.Table_key
    query := fmt.Sprintf("UPDATE %s SET (%s) = (%s) WHERE id = $%d", tableKeyName, strings.Join(columns, ", "), strings.Join(placeholders, ", "), len(values))
    _,err := db.Query(query, values...)
    if err != nil {
        return err
    }
    return nil

}

//___Delete Row_____

//___Delete Rows by id in transaction_____
func DeleteRows( db *sql.DB, table string, c *fiber.Ctx, session utils.Session) error {

    //Get ids from ids query param
    ids := c.Query("ids")
    idsArr := strings.Split(ids, ",")

    //check if ids is empty
    if len(idsArr) == 0 {
        return utils.NewErr("No ids to delete")
    } 

    //check if org_id is empty
    if(session.Org_id == ""){
        return utils.NewErr("No org_id")
    }

    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()
    var query string
    for _, id := range idsArr {
        //Delete row where id = id and org_id = session.Org_id
        query = fmt.Sprintf("DELETE FROM %s WHERE org_id = '%s' AND id = $1", table+"_"+session.Table_key,session.Org_id)
        //Cnvert id to int
        intId, err := strconv.Atoi(id)
        if err != nil {
            return err
        }
        _, err = tx.Exec(query, intId)
        if err != nil {
            tx.Rollback()
            return err
        }
    }

    err = tx.Commit()
    if err != nil {
        return err
    }

    return nil
}

//Update deleted_at row where id = id and org_id = session.Org_id
func SoftDeleteRows( db *sql.DB, table string, c *fiber.Ctx, session utils.Session) error {


     //check if org_id is empty
    if(session.Org_id == ""){
        return utils.NewErr("No org_id")
    }


    //Get ids from ids query param
    ids := c.Query("ids")
    idsArr := strings.Split(ids, ",")

    //check if ids is empty
    if len(idsArr) == 0 {
        return utils.NewErr("No ids to delete")
    }



    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()
    var query string
    for _, id := range idsArr {
        //update delted_at row where id = id and org_id = session.Org_id
        query = fmt.Sprintf("UPDATE %s SET deleted_at = NOW() WHERE org_id = '%s' AND id = $1", table+"_"+session.Table_key,session.Org_id)
        //Cnvert id to int
        intId, err := strconv.Atoi(id)
        if err != nil {
            return err
        }
        _, err = tx.Exec(query, intId)
        if err != nil {
            tx.Rollback()
            return err
        }
    }
    err = tx.Commit()
    if err != nil {
        return err
    }

    return nil
}


