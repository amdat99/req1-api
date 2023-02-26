package database


import (
	"fmt"
	"strconv"
	"strings"
	"database/sql"
	"github.com/gofiber/fiber/v2"
    "encoding/json"
    "req-api/utils"
)


type PaginateType struct {
    Total int `json:"total"`
    Data []map[string]interface{} `json:"data"`
    Page int `json:"page"`
    Limit int `json:"limit"`
}

//___Pagination___
//func Paginate(db *sql.DB, tableName string, fields[]string,c *fiber.Ctx, session Session, conditions map[string]interface{},)  (PaginateType, error) {
func Paginate(db *sql.DB, tableName string, fields[]string,c *fiber.Ctx, session utils.Session)  (PaginateType, error) {

    fmt.Println("session",session)

    if(session.Org_id == "") {
        return PaginateType{}, fmt.Errorf("No org_id")
    }

    var dest []map[string]interface{} = make([]map[string]interface{}, 0)
    var whereClause string = "WHERE org_id = '"+session.Org_id+"'"    
    var total int

	//get limit and page and calculate offset
	limit, err := strconv.Atoi(c.Query("limit", "25"))
    if err != nil {
        limit = 20
    }

    page, err := strconv.Atoi(c.Query("page", "1"))
    if err != nil {
        page = 1
    }

    //Calculate offset (Offest and limit are int type checked so no need to paramaterise)
    offset := (page - 1) * limit


    fmt.Println("limit",limit)
    limitClause := fmt.Sprintf("LIMIT %d OFFSET %d", limit, offset)

    //Order by clause

    orderClause := fmt.Sprintf("ORDER BY id asc")
    
	//Pagination query
    tableKeyName := tableName + "_" + session.Table_key  //Order by id
    query := fmt.Sprintf("SELECT %s FROM %s %s %s %s ; SELECT COUNT(*) FROM %s %s" , strings.Join(fields, ","),tableKeyName,  whereClause,orderClause, limitClause , tableKeyName, whereClause )
    fmt.Println("query",query)
    rows, err := db.Query(query)
    if err != nil {
        fmt.Println(err)
        // return PaginateType{}, err
    }
    defer rows.Close()

    fmt.Println("runs...")




    //Get columns from rows
    columns, err := rows.Columns()
    if err != nil {
        return PaginateType{}, err
    }

    fmt.Println("columns",columns)

	//Create slice of interfaces for values
    values := make([]interface{}, len(columns))

	//Create slice of interface for pointers to values
    valuePtrs := make([]interface{}, len(columns))
    

    fmt.Println("values",values)
    fmt.Println("valuePtrs",valuePtrs)

	//Loop through columns and assign pointers to values
    for i := range columns {
        valuePtrs[i] = &values[i]
    }

    //Loop through first result set
    for rows.Next() {
        //Scan rows into values
        err = rows.Scan(valuePtrs...)
        if err != nil {
            return PaginateType{}, err
        }

        row := make(map[string]interface{})
        //Loop through columns
        for i, col := range columns {
            val := values[i]
            if val == nil {
                row[col] = nil
            } else {
                //Check if value is is json/jsonb and convert to json arr
                if _, ok := val.([]byte); ok {
                
                jsonArr := make([]interface{}, 0)
                    err := json.Unmarshal(val.([]byte), &jsonArr)
                    if err != nil {
                     return PaginateType{}, err
                                fmt.Println(err)
                    }
                row[col] = jsonArr

                } else {
                    row[col] = val
                }

            }
        }
        dest = append(dest, row)
    }

    //Get total count from 2nd result set
    if(rows.NextResultSet()){
        if(rows.Next()){
            if err := rows.Scan(&total); err != nil {
			    return PaginateType{}, err
			}
        }
    }
    err = rows.Err()
    if err != nil {
        return PaginateType{}, err
    }


    //return data
    return PaginateType{
        Total: total,
        Data: dest,
        Page: page,
        Limit: limit,
    }, nil

}

//Get single row
func Single(db *sql.DB, tableName string, fields[]string, c *fiber.Ctx,session utils.Session)  (map[string]interface{},error){
    var result map[string]interface{} = make(map[string]interface{})

    id, err := strconv.Atoi(c.Query("id"))
    if err != nil {
        return result, err
    }

    var query string = fmt.Sprintf("SELECT %s FROM %s WHERE id = $1 AND org_id = $2", strings.Join(fields, ","), tableName + "_" + session.Table_key)
    row,err := db.Query(query, id, session.Org_id)
    if err != nil {
        return result, err
    }
    
    defer row.Close()
    columns, err := row.Columns()
    if err != nil {
        return result, err
    }

    values := make([]interface{}, len(columns))
    valuePtrs := make([]interface{}, len(columns))
    for i := range columns {
        valuePtrs[i] = &values[i]
    }

    for row.Next() {
        err = row.Scan(valuePtrs...)
        if err != nil {
            return result, err
        }

        for i, col := range columns {
            val := values[i]
            if val == nil {
                result[col] = nil
            } else {
                //Check if value is is json/jsonb and convert to arr
                if _, ok := val.([]byte); ok {

                jsonArr := make([]interface{}, 0)
                    err := json.Unmarshal(val.([]byte), &jsonArr)
                    if err != nil {
                    fmt.Println(err)
                    }
                result[col] = jsonArr

                } else {
                    result[col] = val
                }
            }
        }
    }
    err = row.Err()
    if err != nil {
        return result, err
    }

    return result, nil
}

var DefaultFields = []string{"id", "label", "created_at", "updated_at","created_by"}


//Search
func Search(db *sql.DB, tableName string, c *fiber.Ctx,session utils.Session ,fields[]string ) ( []map[string]interface{}, error) {

 	var results []map[string]interface{}
	//Get search query
	query := c.Query("query", "")
	if query == "" {
		return results, utils.NewErr("Query is required")
	}

	var tableKeyName string = tableName + "_" + session.Table_key
	err := Select(db, tableKeyName, fields, &results, c, 120, map[string]interface{}{"org_id": session.Org_id}, map[string]interface{}{"label": query},utils.EmptyIntfMap)
	if err != nil {
		return results, err
	}
	return results, nil
}

//Multi Search
func MultiSearch(db *sql.DB, tableName string, c *fiber.Ctx,session utils.Session ,conditions *utils.MultiSearchType, fields[]string  ) ( []map[string]interface{}, error) {

	var results []map[string]interface{}
	var tableKeyName string = tableName + "_" + session.Table_key

	var equalConditions map[string]interface{} = conditions.Equal
	
	//Add org_id to equal conditions
	equalConditions["org_id"] = session.Org_id

	err := Select(db, tableKeyName, fields, &results, c, 120, equalConditions, conditions.Like, conditions.In)
	if err != nil {
		return results, err
	}
	return results, nil
}



//SlectRows
func SelectRows(db *sql.DB, tableName string, c *fiber.Ctx, fields[]string, conditions map[string]interface{},limit int ) ([]map[string]interface{}, error) {
    var results []map[string]interface{}
    err := Select(db, tableName, fields, &results, c, limit, conditions, utils.EmptyIntfMap, utils.EmptyIntfMap)
    if err != nil {
        return results, err
    }
    return results, nil
}


//___Get Rows With optional where, in and like conditions
func Select(db *sql.DB, tableName string, fields[]string, dest *[]map[string]interface{},c *fiber.Ctx, limit int, conditions map[string]interface{}, likeConditions map[string]interface{}, inConditions map[string]interface{}  ) error {

    //var whereClause string = "WHERE org_id = 1"
	var whereClause string = ""
    var args []interface{}

	//get where clause
	if len(conditions) > 0 {
    for columnName, value := range conditions {
        if whereClause != "" {
            whereClause += " AND "
        } else {
            whereClause = "WHERE "
        }
        whereClause += fmt.Sprintf("%s = $%d", columnName, len(args)+1)
        args = append(args, value)
    }
}

	//Like conditions
	if len(likeConditions) > 0 {
			for columnName, value := range likeConditions {
				if whereClause != "" {
					whereClause += " AND "
				} else {
					whereClause = "WHERE "
				}
				//Check if lowercase column name is in fields
				whereClause += fmt.Sprintf("%s ILIKE $%d", columnName, len(args)+1)
				args = append(args, "%"+value.(string)+"%")
			}
		}

    //Where in conditions
    if len(inConditions) > 0 {
        for columnName, value := range inConditions {
            if whereClause != "" {
                whereClause += " AND "
            } else {
                whereClause = "WHERE"
            }
            //join values with comma
            inValues := strings.Join(value.([]string), ",")
            whereClause += fmt.Sprintf("%s IN ($%d)", columnName, len(args)+1)
            args = append(args, inValues)
        }
    }


    //Add order by id asc

    orderBy := "ORDER BY id ASC"

    query := fmt.Sprintf("SELECT %s FROM %s %s %s LIMIT %d", strings.Join(fields, ","), tableName, whereClause,orderBy, limit)
    rows, err := db.Query(query, args...)
    if err != nil {
        return err
    }
    defer rows.Close()

    columns, err := rows.Columns()
    if err != nil {
        return err
    }
    values := make([]interface{}, len(columns))
    valuePtrs := make([]interface{}, len(columns))
    for i := range columns {
        valuePtrs[i] = &values[i]
    }

    for rows.Next() {
        err = rows.Scan(valuePtrs...)
        if err != nil {
            return err
        }

        row := make(map[string]interface{})
        for i, col := range columns {
            val := values[i]
            if val == nil {
                row[col] = nil
            } else {
                //Check if value is is json/jsonb and convert to arr
                if _, ok := val.([]byte); ok {

                jsonArr := make([]interface{}, 0)
                    err := json.Unmarshal(val.([]byte), &jsonArr)
                    if err != nil {
                        fmt.Println(err)
                    }
                row[col] = jsonArr

                } else {
                    row[col] = val
                }
            }
        }
        *dest = append(*dest, row)
    }

    err = rows.Err()
    if err != nil {
        return err
    }

    return nil
}
























// func prepareQueryPlaceholders(start, quantity int) string {
// 	placeholders := make([]string, 0, quantity)
// 	end := start + quantity
// 	for i := start; i < end; i++ {
// 		placeholders = append(placeholders, strings.Join([]string{"$", strconv.Itoa(i)}, ""))
// 	}
// 	return strings.Join(placeholders, ",")
// }


// //Multi Query, takes in an slice of queries and returns a interface of slices
// func MultiQuery(db *sql.DB, queries []string) ([][]interface{}, error) {
//     // Join the queries into a single string with semicolons between them
//     query := strings.Join(queries, "; ")

//     // Execute the combined query
//     rows, err := db.Query(query)
//     if err != nil {
//         return nil, err
//     }
//     defer rows.Close()

//     // Get the column names for each result set
//     resultSets := make([][][]string, len(queries))
//     for i := range queries {
//         columns, err := rows.Columns()
//         if err != nil {
//             return nil, err
//         }
//         resultSets[i] = make([][]string, len(columns))
//         for j := range columns {
//             resultSets[i][j] = []string{columns[j]}
//         }
//     }

//     // Copy the results into slices of interface{} values
//     results := make([][]interface{}, len(queries))
//     for i := range queries {
//         for rows.Next() {
//             values := make([]interface{}, len(resultSets[i]))
//             scanArgs := make([]interface{}, len(values))
//             for j := range values {
//                 scanArgs[j] = &values[j]
//             }
//             err := rows.Scan(scanArgs...)
//             if err != nil {
//                 return nil, err
//             }
//             for j, value := range values {
//                 resultSets[i][j] = append(resultSets[i][j], fmt.Sprintf("%v", value))
//             }
//         }
//         if err := rows.Err(); err != nil {
//             return nil, err
//         }
//         // Convert the slices of string values to slices of interface{} values
//         for _, row := range resultSets[i] {
//             rowData := make([]interface{}, len(row))
//             for j, value := range row {
//                 rowData[j] = value
//             }
//             results[i] = append(results[i], rowData)
//         }
//     }
//     return results, nil
// }




// //Similar to MultiQuery but it's a transaction
// func ExecuteStatements(tx *sql.Tx, stmts []*sql.Stmt, args ...[]interface{}) ([]sql.Result, error) {
//     results := make([]sql.Result, 0, len(stmts))

//     for i, stmt := range stmts {
//         result, err := stmt.Exec(args[i]...)
//         if err != nil {
//             return nil, err
//         }

//         results = append(results, result)
//     }

//     return results, nil
// }
// func Transaction(db *sql.DB, statements []*sql.Stmt, args []any)  ([]sql.Result, error)  {
//   tx, err := db.Begin()
//     if err != nil {
//         return nil, err
//     }
//     defer tx.Rollback()

//     results := make([]sql.Result, 0, len(statements))

//     for i, stmt := range statements {
//         result, err := stmt.Exec()
//         if err != nil {
//             return nil, err
//         }
//         results = append(results, result)
//     }

//     err = tx.Commit()
//     if err != nil {
//         return nil, err
//     }

//     return results, nil
// }