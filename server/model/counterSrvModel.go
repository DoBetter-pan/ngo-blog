/**
* @file counterSrvModel.go
* @brief counter service model
* @author yingx
* @date 2015-02-27
 */

package model

import (
	_ "fmt"
    "time"
    _ "database/sql"
    _ "github.com/go-sql-driver/mysql"
	dbwrapper "go-angular/server/datawrapper"
)

var counterSqls map[string] string = map[string] string {
    "querycount":"select count(*) as counter from ng_counter_visitor where url=?",
    "insert":"insert into ng_counter_visitor(username, ip, url, new) values( ?, ?, ?, 1)",
    "lastdate":"select max(visitTime) as lastdate from ng_counter_visitor where ip=? and url=?",
}

type CounterSrvModel struct {
}

func NewCounterSrvModel() *CounterSrvModel {
	return &CounterSrvModel{}
}

func (model *CounterSrvModel) FindCounter(url string) int64 {
    var count int64 = 0

    dbconnection := dbwrapper.GetDatabaseConnection()
    err := dbconnection.DB.QueryRow(counterSqls["querycount"], url).Scan(&count)
    if err != nil {
        count = 0
    }

    return count
}

func (model *CounterSrvModel) Insert(name, ip, url string) error {
    lastdate := ""

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return err
    }
    err = tx.QueryRow(counterSqls["lastdate"], ip, url).Scan(&lastdate)
    if err != nil {
        lastdate = "2006-01-02 15:04:05"
    }

    t, _ := time.Parse("2006-01-02 15:04:05", lastdate)
    if(time.Now().Sub(t) > 7200) {
        _, err := tx.Exec(counterSqls["insert"], name, ip, url)
        if err != nil {
            tx.Rollback()
            return err
        }
    }
   
    tx.Commit()

    return nil
}

