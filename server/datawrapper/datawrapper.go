/** 
* @file datawrapper.go
* @brief data wrapper. go get github.com/go-sql-driver/mysql
* @author yingx
* @date 2016-07-26
*/

package datawrapper

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type DbConfig struct {
    Protocal string
    Host string
    Port string
    User string
    Password string
    Database string
    Options string
}

type DbConnection struct {
    DB *sql.DB
}

var WrapperConfig *DbConfig = &DbConfig{"tcp", "127.0.0.1", "3306", "test", "123456", "ngo-blog", "charset=utf8"}

var WrapperConn *DbConnection = &DbConnection{}

func init() {
    var err error
    config := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?%s", WrapperConfig.User, WrapperConfig.Password, WrapperConfig.Protocal, WrapperConfig.Host, WrapperConfig.Port, WrapperConfig.Database, WrapperConfig.Options)
    WrapperConn.DB, err = sql.Open("mysql", config)
    if err != nil {
        fmt.Println("Cann't connect to mysql!!!", err)
        return
    }

    err = WrapperConn.DB.Ping()
    if err != nil {
        fmt.Println("Failed to ping to mysql!!!", err)
    }
}

func GetDatabaseConnection() *DbConnection {
    return WrapperConn
}

