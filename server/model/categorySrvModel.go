/**
* @file categorySrvModel.go
* @brief category service model
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

type Category struct {
     Id int64 `json:"id"`
     Name  string `json:"name"`
     Url  string `json:"url"`
     SectionId  string `json:"sectionId"`     
}

var categorySqls map[string] string = map[string] string {
    "query":"select id, name, url, sectionId from ng_blog_category",
    "querybysec":"select id, name, url from ng_blog_category where sectionId=?",
    "queryone":"select id, name, url, sectionId from ng_blog_category where id=?",
    "insert":"insert into ng_blog_category( name, url, sectionId) values( ?, ?, ?)",
    "update":"update ng_blog_category set name=?, url=?, sectionId=? where id=?",
    "delete":"delete from ng_blog_category where id=?",
}

type CategorySrvModel struct {
}

func NewCategorySrvModel() *CategorySrvModel {
	return &CategorySrvModel{}
}

func (model *CategorySrvModel) FindAll() (string, error) {
    dbconnection := dbwrapper.GetDatabaseConnection()
    rows, err := dbconnection.DB.Query(categorySqls["query"])
    if err != nil {
        return "", err
    }
    defer rows.Close()

    categoryList := make([]Category, 0, 10)
    for rows.Next() {
        var category Category
        err = rows.Scan( &category.Id, &category.Name, &category.Url, &category.SectionId)
        if err == nil {
            categoryList = append(categoryList, category)
        }
    }

    //check error
    if err = rows.Err(); err != nil {
        return "", err
    }

    data, err := json.Marshal(categoryList)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *CategorySrvModel) Find(id int64) (string, error) {
    var category Category
    dbconnection := dbwrapper.GetDatabaseConnection()
    err := dbconnection.DB.QueryRow(categorySqls["queryone"], id).Scan( &category.Id, &category.Name, &category.Url, &category.SectionId)
    if err != nil {
        return "", err
    }

    data, err := json.Marshal(category)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *CategorySrvModel) Insert(str string) (string, error) {
    var category Category

    err := json.Unmarshal([]byte(str), &category)
    if err != nil {
        return "", err
    }

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    res, err := tx.Exec(categorySqls["insert"],  category.Name, category.Url, category.SectionId, category.Id)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    categoryid, err := res.LastInsertId()
    if err != nil {
        tx.Rollback()
        return "", err
    }

    tx.Commit()

    category.Id = categoryid
    data, err := json.Marshal(category)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *CategorySrvModel) Update(id int64, str string) (string, error) {
    var category Category

    err := json.Unmarshal([]byte(str), &category)
    if err != nil {
        return "", err
    }

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    //just update, not check if it is same before updating. It may be supported in future
    _, err = tx.Exec(categorySqls["update"],  category.Name, category.Url, category.SectionId, category.Id)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    tx.Commit()

    return str, nil
}

func (model *CategorySrvModel) Delete(id int64) error {
    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return err
    }

    _, err = tx.Exec(categorySqls["delete"], id)
    if err != nil {
        tx.Rollback()
        return err
    }

    tx.Commit()

    return nil
}
