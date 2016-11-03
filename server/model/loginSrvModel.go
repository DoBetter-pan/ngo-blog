/**
* @file loginSrvModel.go
* @brief login service model
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

type Login struct {
     Id int64 `json:"id"`
     Name  string `json:"name"`
     Password string `json:"password"`
     Role string `json:"role"`
     Nickname string `json:"nickname"`
     Email string `json:"email"`
     LastAccess string `json:"lastAccess"`
     Nonce string `json:"nonce"`
     Stored int64 `json:"stored"`
}

var loginSqls map[string] string = map[string] string {
    "query":"select id, name, password, role, nickname, email, lastAccess, nonce from ng_blog_user",
    "queryone":"select id, name, password, role, nickname, email, lastAccess, nonce from ng_blog_user where id=?",
    "queryonebyname":"select id, name, password, role, nickname, email, lastAccess, nonce from ng_blog_user where name=?",
    "insert":"insert into ng_blog_user(name, password, role, nickname, email, lastAccess, nonce) values( ?, ?, ?, ?, ?, ?, ?)",
    "update":"update ng_blog_user set name=?, password=?, role=?, nickname=?, email=?, lastAccess=?, nonce=? where id=?",
    "delete":"delete from ng_blog_user where id=?",
}

type LoginSrvModel struct {
}

func NewLoginSrvModel() *LoginSrvModel {
	return &LoginSrvModel{}
}

func (model *LoginSrvModel) FindAll() (string, error) {
    dbconnection := dbwrapper.GetDatabaseConnection()
    rows, err := dbconnection.DB.Query(loginSqls["query"])
    if err != nil {
        return "", err
    }
    defer rows.Close()

    loginList := make([]Login, 0, 10)
    for rows.Next() {
        var login Login
        err = rows.Scan( &login.Id, &login.Name, &login.Password, &login.Role, &login.Nickname, &login.Email, &login.LastAccess, &login.Nonce)
        if err == nil {
            loginList = append(loginList, login)
        }
    }

    //check error
    if err = rows.Err(); err != nil {
        return "", err
    }

    data, err := json.Marshal(loginList)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *LoginSrvModel) Find(id int64) (string, error) {
    var login Login
    dbconnection := dbwrapper.GetDatabaseConnection()
    err := dbconnection.DB.QueryRow(loginSqls["queryone"], id).Scan(&login.Id, &login.Name, &login.Password, &login.Role, &login.Nickname, &login.Email, &login.LastAccess, &login.Nonce)
    if err != nil {
        return "", err
    }

    data, err := json.Marshal(login)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *LoginSrvModel) Insert(str string) (string, error) {
    var login Login

    err := json.Unmarshal([]byte(str), &login)
    if err != nil {
        return "", err
    }

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    res, err := tx.Exec(loginSqls["insert"], login.Name, login.Password, login.Role, login.Nickname, login.Email, login.LastAccess, login.Nonce, login.Id)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    loginid, err := res.LastInsertId()
    if err != nil {
        tx.Rollback()
        return "", err
    }

    tx.Commit()

    login.Id = loginid
    data, err := json.Marshal(login)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *LoginSrvModel) Update(id int64, str string) (string, error) {
    var login Login

    err := json.Unmarshal([]byte(str), &login)
    if err != nil {
        return "", err
    }

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    //just update, not check if it is same before updating. It may be supported in future
    _, err = tx.Exec(loginSqls["update"], login.Name, login.Password, login.Role, login.Nickname, login.Email, login.LastAccess, login.Nonce, login.Id)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    tx.Commit()

    return str, nil
}

func (model *LoginSrvModel) Delete(id int64) error {
    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return err
    }

    _, err = tx.Exec(loginSqls["delete"], id)
    if err != nil {
        tx.Rollback()
        return err
    }

    tx.Commit()

    return nil
}

func (model *LoginSrvModel) FindObject(id int64) (Login, error) {
    var login Login
    dbconnection := dbwrapper.GetDatabaseConnection()
    err := dbconnection.DB.QueryRow(loginSqls["queryone"], id).Scan(&login.Id, &login.Name, &login.Password, &login.Role, &login.Nickname, &login.Email, &login.LastAccess, &login.Nonce)

    return login, err
}

func (model *LoginSrvModel) FindObjectByName(name string) (Login, error) {
    var login Login
    dbconnection := dbwrapper.GetDatabaseConnection()
    err := dbconnection.DB.QueryRow(loginSqls["queryonebyname"], name).Scan(&login.Id, &login.Name, &login.Password, &login.Role, &login.Nickname, &login.Email, &login.LastAccess, &login.Nonce)

    return login, err
}