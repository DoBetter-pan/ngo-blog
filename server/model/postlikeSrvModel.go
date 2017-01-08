/**
* @file postlikeSrvModel.go
* @brief postlike service model
* @author yingx
* @date 2015-02-27
 */

package model

import (
	_ "fmt"
    "time"
    _ "database/sql"
    _ "github.com/go-sql-driver/mysql"
	dbwrapper "ngo-blog/server/datawrapper"
)

var postlikeSqls map[string] string = map[string] string {
    "querylikecount":"select sum(value) from ng_post_like where value>=0 and postId=?",
    "queryunlikecount":"select sum(value) from ng_post_like where value<=0 and postId=?",
    "insert":"insert into ng_post_like(postId, value, dateTime, ip, cookie, userId) values( ?, ?, ?, ?, ?, ?)",
    "querycount":"select count(*) from ng_post_like where postId=? and ((userId!=0 and userId=?) or (cookie!='' and cookie=?))",
}

type PostlikeSrvModel struct {
}

func NewPostlikeSrvModel() *PostlikeSrvModel {
	return &PostlikeSrvModel{}
}

func (model *PostlikeSrvModel) GetLikeCount(postId int64) int64 {
    var count int64 = 0

    dbconnection := dbwrapper.GetDatabaseConnection()
    err := dbconnection.DB.QueryRow(postlikeSqls["querylikecount"], postId).Scan(&count)
    if err != nil {
        count = 0
    }

    return count
}

func (model *PostlikeSrvModel) GetUnlikeCount(postId int64) int64 {
    var count int64 = 0

    dbconnection := dbwrapper.GetDatabaseConnection()
    err := dbconnection.DB.QueryRow(postlikeSqls["queryunlikecount"], postId).Scan(&count)
    if err != nil {
        count = 0
    }

    return count
}

func (model *PostlikeSrvModel) GetLikeUnlikeCount(postId, userId int64, cookie string) int64 {
    var count int64 = 0

    dbconnection := dbwrapper.GetDatabaseConnection()
    err := dbconnection.DB.QueryRow(postlikeSqls["querycount"], postId, userId, cookie).Scan(&count)
    if err != nil {
        count = 0
    }

    return count
}

func (model *PostlikeSrvModel) Insert(postId int64, value int64, ip, cookie string, userId int64) error {
    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {        
        return err
    }
    t := time.Now()
    strTime := t.Format("2006-01-02 15:04:05")
    _, err = tx.Exec(postlikeSqls["insert"], postId, value, strTime, ip, cookie, userId)
    if err != nil {
        tx.Rollback()
        return err
    }

    tx.Commit()
    return nil
}


