/**
* @file linkSrvModel.go
* @brief link service model
* @author yingx
* @date 2015-02-27
 */

package model

import (
    "encoding/json"
    _ "database/sql"
    _ "github.com/go-sql-driver/mysql"
	dbwrapper "go-angular/server/datawrapper"
)

type Link struct {
     Id int64 `json:"id"`
     Name  string `json:"name"`
     Url string `json:"url"`
     External int64 `json:"external"`
}

var linkSqls map[string] string = map[string] string {
    "query":"select id, name, url, external from ng_blog_link",
    "queryone":"select id, name, url, external from ng_blog_link where id=?",
    "insert":"insert into ng_blog_link( name, url, external) values( ?, ?, ?)",
    "update":"update ng_blog_link set name=?, url=?, external=? where id=?",
    "delete":"delete from ng_blog_link where id=?",
}

type LinkSrvModel struct {
}

func NewLinkSrvModel() *LinkSrvModel {
	return &LinkSrvModel{}
}

func (model *LinkSrvModel) FindAllLinks() ([]Link, error) {
    dbconnection := dbwrapper.GetDatabaseConnection()
    rows, err := dbconnection.DB.Query(linkSqls["query"])
    linkList := make([]Link, 0, 10)
    if err != nil {
        return linkList, err
    }
    defer rows.Close()

    for rows.Next() {
        var link Link
        err = rows.Scan( &link.Id, &link.Name, &link.Url, &link.External)
        if err == nil {
            linkList = append(linkList, link)
        }
    }

    //check error
    if err = rows.Err(); err != nil {
        return linkList, err
    }


    return linkList, nil
}

func (model *LinkSrvModel) FindAll() (string, error) {
    linkList, err := model.FindAllLinks()
    if err != nil {
        return "", err
    }

    data, err := json.Marshal(linkList)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *LinkSrvModel) Find(id int64) (string, error) {
    var link Link
    dbconnection := dbwrapper.GetDatabaseConnection()
    err := dbconnection.DB.QueryRow(linkSqls["queryone"], id).Scan( &link.Id, &link.Name, &link.Url, &link.External)
    if err != nil {
        return "", err
    }

    data, err := json.Marshal(link)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *LinkSrvModel) Insert(str string) (string, error) {
    var link Link

    err := json.Unmarshal([]byte(str), &link)
    if err != nil {
        return "", err
    }

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    res, err := tx.Exec(linkSqls["insert"],  link.Name, link.Url, link.External, link.Id)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    linkid, err := res.LastInsertId()
    if err != nil {
        tx.Rollback()
        return "", err
    }

    tx.Commit()

    link.Id = linkid
    data, err := json.Marshal(link)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *LinkSrvModel) Update(id int64, str string) (string, error) {
    var link Link

    err := json.Unmarshal([]byte(str), &link)
    if err != nil {
        return "", err
    }

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    //just update, not check if it is same before updating. It may be supported in future
    _, err = tx.Exec(linkSqls["update"],  link.Name, link.Url, link.External, link.Id)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    tx.Commit()

    return str, nil
}

func (model *LinkSrvModel) Delete(id int64) error {
    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return err
    }

    _, err = tx.Exec(linkSqls["delete"], id)
    if err != nil {
        tx.Rollback()
        return err
    }

    tx.Commit()

    return nil
}
