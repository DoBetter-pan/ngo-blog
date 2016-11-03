/**
* @file sectionSrvModel.go
* @brief section service model
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

type Section struct {
     Id int64 `json:"id"`
     Name  string `json:"name"`
     Url  string `json:"url"`
     Authority int64 `json:"authority"`
}

var sectionSqls map[string] string = map[string] string {
    "query":"select id, name, url, authority from ng_blog_section",
    "queryone":"select id, name, url, authority from ng_blog_section where id=?",
    "insert":"insert into ng_blog_section( name, url, authority) values( ?, ?, ?)",
    "update":"update ng_blog_section set name=?, url=?, authority=? where id=?",
    "delete":"delete from ng_blog_section where id=?",
}

type SectionSrvModel struct {
}

func NewSectionSrvModel() *SectionSrvModel {
	return &SectionSrvModel{}
}

func (model *SectionSrvModel) FindAll() (string, error) {
    dbconnection := dbwrapper.GetDatabaseConnection()
    rows, err := dbconnection.DB.Query(sectionSqls["query"])
    if err != nil {
        return "", err
    }
    defer rows.Close()

    sectionList := make([]Section, 0, 10)
    for rows.Next() {
        var section Section
        err = rows.Scan( &section.Id, &section.Name, &section.Url, &section.Authority)
        if err == nil {
            sectionList = append(sectionList, section)
        }
    }

    //check error
    if err = rows.Err(); err != nil {
        return "", err
    }

    data, err := json.Marshal(sectionList)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *SectionSrvModel) Find(id int64) (string, error) {
    var section Section
    dbconnection := dbwrapper.GetDatabaseConnection()
    err := dbconnection.DB.QueryRow(sectionSqls["queryone"], id).Scan( &section.Id, &section.Name, &section.Url, &section.Authority)
    if err != nil {
        return "", err
    }

    data, err := json.Marshal(section)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *SectionSrvModel) Insert(str string) (string, error) {
    var section Section

    err := json.Unmarshal([]byte(str), &section)
    if err != nil {
        return "", err
    }

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    res, err := tx.Exec(sectionSqls["insert"],  section.Name, section.Url, section.Authority, section.Id)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    sectionid, err := res.LastInsertId()
    if err != nil {
        tx.Rollback()
        return "", err
    }

    tx.Commit()

    section.Id = sectionid
    data, err := json.Marshal(section)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *SectionSrvModel) Update(id int64, str string) (string, error) {
    var section Section

    err := json.Unmarshal([]byte(str), &section)
    if err != nil {
        return "", err
    }

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    //just update, not check if it is same before updating. It may be supported in future
    _, err = tx.Exec(sectionSqls["update"],  section.Name, section.Url, section.Authority, section.Id)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    tx.Commit()

    return str, nil
}

func (model *SectionSrvModel) Delete(id int64) error {
    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return err
    }

    _, err = tx.Exec(sectionSqls["delete"], id)
    if err != nil {
        tx.Rollback()
        return err
    }

    tx.Commit()

    return nil
}
