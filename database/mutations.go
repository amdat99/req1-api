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
    
    var columns []string
    var values []interface{}

    fmt.Println(data)

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
    fmt.Println(query)
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

    //Get ids from ids query param
    ids := c.Query("ids")
    idsArr := strings.Split(ids, ",")

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


