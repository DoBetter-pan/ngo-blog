/**
* @file basicSrvModel.go
* @brief basic service model
* @author yingx
* @date 2015-02-27
 */

package model

import (
	_ "fmt"
    "encoding/json"
    _ "database/sql"
    _ "github.com/go-sql-driver/mysql"
	dbwrapper "go-angular/server/datawrapper"
)

type Basic struct {
##Item##
}

var basicSqls map[string] string = map[string] string {
    "query":"##query##",
    "queryone":"##queryone##",
    "insert":"##insert##",
    "update":"##update##",
    "delete":"##delete##",
}

type BasicSrvModel struct {
}

func NewBasicSrvModel() *BasicSrvModel {
	return &BasicSrvModel{}
}

func (model *BasicSrvModel) FindAll() (string, error) {
    dbconnection := dbwrapper.GetDatabaseConnection()
    rows, err := dbconnection.DB.Query(basicSqls["query"])
    if err != nil {
        return "", err
    }
    defer rows.Close()

    basicList := make([]Basic, 0, 10)
    for rows.Next() {
        var basic Basic
        err = rows.Scan(##scanitemlist##)
        if err == nil {
            basicList = append(basicList, basic)
        }
    }

    //check error
    if err = rows.Err(); err != nil {
        return "", err
    }

    data, err := json.Marshal(basicList)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *BasicSrvModel) Find(id int64) (string, error) {
    var basic Basic
    dbconnection := dbwrapper.GetDatabaseConnection()
    err := dbconnection.DB.QueryRow(basicSqls["queryone"], id).Scan(##scanitemlist##)
    if err != nil {
        return "", err
    }

    data, err := json.Marshal(basic)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *BasicSrvModel) Insert(str string) (string, error) {
    var basic Basic

    err := json.Unmarshal([]byte(str), &basic)
    if err != nil {
        return "", err
    }

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    res, err := tx.Exec(basicSqls["insert"], ##insertitemlist##)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    basicid, err := res.LastInsertId()
    if err != nil {
        tx.Rollback()
        return "", err
    }

    tx.Commit()

    basic.Id = basicid
    data, err := json.Marshal(basic)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *BasicSrvModel) Update(id int64, str string) (string, error) {
    var basic Basic

    err := json.Unmarshal([]byte(str), &basic)
    if err != nil {
        return "", err
    }

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    //just update, not check if it is same before updating. It may be supported in future
    _, err = tx.Exec(basicSqls["update"], ##updateitemlist##)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    tx.Commit()

    return str, nil
}

func (model *BasicSrvModel) Delete(id int64) error {
    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return err
    }

    _, err = tx.Exec(basicSqls["delete"], id)
    if err != nil {
        tx.Rollback()
        return err
    }

    tx.Commit()

    return nil
}
