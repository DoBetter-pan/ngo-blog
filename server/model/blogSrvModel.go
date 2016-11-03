/**
* @file blogSrvModel.go
* @brief blog service model
* @author yingx
* @date 2015-02-27
 */

package model

import (
	"fmt"
    "time"
    "encoding/json"
    "strconv"
    _ "database/sql"
    _ "github.com/go-sql-driver/mysql"
	dbwrapper "go-angular/server/datawrapper"
)

type Article struct {
    Id int64 `json:"id"`
    Author  string `json:"author"`
    Title  string `json:"title"`
    TitleHtml  string `json:"titleHtml"`
    Content  string `json:"content"`
    ContentHtml  string `json:"contentHtml"`
    Excerpt  string `json:"excerpt"`
    ExcerptHtml  string `json:"excerptHtml"`
    Url  string `json:"url"`
    SectionId int64 `json:"sectionId"`    
    SectionName string `json:"sectionName"`
    SectionUrl string `json:"sectionUrl"`
    CategoryId int64 `json:"categoryId"`       
    CategoryName string `json:"categoryName"`
    CategoryUrl string `json:"categoryUrl"`    
    CommentCount int64 `json:"commentCount"`
    Status int64 `json:"status"`
    Posted  string `json:"posted"`
    LastMod string `json:"lastMod"`
    Expires string `json:"expires"`
}

type ArticleList struct {
    Page  int64 `json:"page"`
    Total int64 `json:"total"`
    HasPrevious  int64 `json:"hasPrevious"`
    PreviousUrl string `json:"previousUrl"`
    HasNext  int64 `json:"hasNext"`
    NextUrl string `json:"nextUrl"`
    Articles []Article `json:"articles"`
}

type ArticleView struct {
    PreviousId  int64 `json:"previousId"`
    PreviousTitle string `json:"previousTitle"`    
    PreviousUrl string `json:"previousUrl"`
    NextId  int64 `json:"nextId"`
    NextTitle string `json:"nextTitle"`        
    NextUrl string `json:"nextUrl"`
    Art Article `json:"article"`
}

type ArticleDesc struct {
    Id int64 `json:"id"`
    Title  string `json:"title"`
    TitleHtml  string `json:"titleHtml"`
    Url  string `json:"url"`
}

type ArticlesInCategory struct {
    Id int64 `json:"id"`
    Name  string `json:"name"`
    Url  string `json:"url"`
    Description  string `json:"description"`
    Articles []ArticleDesc `json:"articles"`
}

type ArticleWriter struct {
    Id int64 `json:"id"`
    Author  int64 `json:"author"`
    Title  string `json:"title"`
    TitleHtml  string `json:"titleHtml"`
    Content  string `json:"content"`
    ContentHtml  string `json:"contentHtml"`
    Excerpt  string `json:"excerpt"`
    ExcerptHtml  string `json:"excerptHtml"`
    Url  string `json:"url"`
    Section int64 `json:"section"`
    Category int64 `json:"category"`
    CommentCount int64 `json:"commentCount"`
    Status int64 `json:"status"`
    Posted  string `json:"posted"`
    LastMod string `json:"lastMod"`
    Expires string `json:"expires"`
}


var blogSqls map[string] string = map[string] string {
    "queryindex":"select id, title, titleHtml from ng_blog_article where categoryId=? order by posted desc limit ?",
    "querylist":"select article.id, user.name, title, titleHtml, content, contentHtml, IFNULL(excerpt, '') as excerpt, IFNULL(excerptHtml, '') as excerptHtml, section.id as sectionId, section.name as sectionName, section.url as sectionUrl, category.id as categoryId, category.name as categoryName, category.url as categoryUrl, commentsCount, status, posted, lastMod, IFNULL(expires, '') as expires from ng_blog_article article, ng_blog_section section, ng_blog_category category, ng_blog_user user where article.sectionId=section.id and categoryId=category.id and authorId=user.id",
    "querycount":"select count(*) as total from ng_blog_article where 1",
    "queryone":"select article.id, user.name, title, titleHtml, content, contentHtml, IFNULL(excerpt, '') as excerpt, IFNULL(excerptHtml, '') as excerptHtml, section.id as sectionId, section.name as sectionName, section.url as sectionUrl, category.id as categoryId, category.name as categoryName, category.url as categoryUrl, commentsCount, status, posted, lastMod, IFNULL(expires, '') as expires from ng_blog_article article, ng_blog_section section, ng_blog_category category, ng_blog_user user where article.sectionId=section.id and categoryId=category.id and authorId=user.id and article.id=?",
    "insert":"insert into ng_blog_article( authorId, title, titleHtml, content, contentHtml, sectionId, categoryId, commentsCount, status, posted, lastMod) values( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
    "update":"update blog set name=?, age=?, sex=? where id=?",
    "delete":"delete from blog where id=?",
    "category":"select id, name, url, description from ng_blog_category where 1",
    "queryprevious":"select id, titleHtml from ng_blog_article where categoryId=? and posted<? order by posted desc limit 1",
    "querynext":"select id, titleHtml from ng_blog_article where categoryId=? and posted>? order by posted asc limit 1",        
}

type BlogSrvModel struct {
}

func NewBlogSrvModel() *BlogSrvModel {
	return &BlogSrvModel{}
}

func (model *BlogSrvModel) FindAllByKeyValue(key string, value, page int64) (string, error) {
    articleList := ArticleList{}
    articleList.Articles = make([]Article, 0, 10)

    condition := ""
    conditionList := ""
    offset := page * 10
    if key == "s" {
        condition = " and sectionId=?"
        conditionList = fmt.Sprintf(" and article.sectionId=? order by posted desc limit 10 offset %d", offset)
    } else if key == "c" {
        condition = " and categoryId=?"
        conditionList = fmt.Sprintf(" and article.categoryId=? order by posted desc limit 10 offset %d", offset)
    }
    sqlCount := blogSqls["querycount"] + condition

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    err = tx.QueryRow(sqlCount, value).Scan(&articleList.Total)
    if err != nil {
        tx.Rollback()
        return "", err
    }
    articleList.Page = page
    if page > 0 {
        articleList.HasPrevious = 1
        articleList.PreviousUrl = CreateUrlByKeyValue(key, value, page - 1)
    } else {
        articleList.HasPrevious = 0
        articleList.PreviousUrl = ""
    }
    if articleList.Total - (page + 1) * 10 > 0 {
        articleList.HasNext = 1
        articleList.NextUrl = CreateUrlByKeyValue(key, value, page + 1)
    } else {
        articleList.HasNext = 0
        articleList.NextUrl = ""
    }

    sqlList := blogSqls["querylist"] + conditionList
    rowsArticle, err := tx.Query(sqlList, value)
    if err != nil {
        tx.Rollback()
        return "", err
    }
    defer rowsArticle.Close()

    for rowsArticle.Next() {
        var art Article

        err = rowsArticle.Scan(&art.Id, &art.Author, &art.Title, &art.TitleHtml, &art.Content, &art.ContentHtml, &art.Excerpt, &art.ExcerptHtml, &art.SectionId, &art.SectionName, &art.SectionUrl, &art.CategoryId, &art.CategoryName, &art.CategoryUrl, &art.CommentCount, &art.Status, &art.Posted, &art.LastMod, &art.Expires)
        if err == nil {
            if len(art.SectionUrl) == 0 {
                art.SectionUrl = CreateUrl(BLOG_SECTION, art.SectionId, 0)
            }
            if len(art.CategoryUrl) == 0 {
                art.CategoryUrl = CreateUrl(BLOG_CATEGORY, art.CategoryId, 0)
            }            
            art.Url = CreateUrl(BLOG_ARTICLE, art.Id, 0)
            articleList.Articles = append(articleList.Articles, art)
        }
    }

    //check error
    if err = rowsArticle.Err(); err != nil {
        tx.Rollback()
        return "", err
    }

    tx.Commit()

    //change into array
    articleListArray := make([]ArticleList, 0, 1)
    articleListArray = append(articleListArray, articleList)
    data, err := json.Marshal(articleListArray)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *BlogSrvModel) FindAllByCount(count int) (string, error) {
    artsInCatList := make([]ArticlesInCategory, 0, 20)

    sql := blogSqls["category"] + " and isPage=1"
    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    rows, err := tx.Query(sql)
    if err != nil {
        tx.Rollback()
        return "", err
    }
    defer rows.Close()

    for rows.Next() {
        var artsInCat ArticlesInCategory
        artsInCat.Articles = make([]ArticleDesc, 0, 5)
        err = rows.Scan(&artsInCat.Id, &artsInCat.Name, &artsInCat.Url, &artsInCat.Description)
        if err == nil {
            if len(artsInCat.Url) == 0 {
                artsInCat.Url = CreateUrl(BLOG_CATEGORY, artsInCat.Id, 0)    
            }            
            artsInCatList = append(artsInCatList, artsInCat)
        }
    }

    //check error
    if err = rows.Err(); err != nil {
        tx.Rollback()
        return "", err
    }

    artsInCatLen := len(artsInCatList)
    for i := 0; i < artsInCatLen; i++ {
        rowsArticle, err := tx.Query(blogSqls["queryindex"], artsInCatList[i].Id, count)
        if err != nil {
            tx.Rollback()
            return "", err
        }
        defer rowsArticle.Close()

        for rowsArticle.Next() {
            var art ArticleDesc

            err = rowsArticle.Scan(&art.Id, &art.Title, &art.TitleHtml )
            if err == nil {
                art.Url = CreateUrl(BLOG_ARTICLE, art.Id, 0)
                artsInCatList[i].Articles = append(artsInCatList[i].Articles, art)
            }
        }

        //check error
        if err = rows.Err(); err != nil {
            tx.Rollback()
            return "", err
        }
    }
    tx.Commit()

    data, err := json.Marshal(artsInCatList)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *BlogSrvModel) Find(id int64) (string, error) {    
    var artView ArticleView
    var art Article
    dbconnection := dbwrapper.GetDatabaseConnection()
    err := dbconnection.DB.QueryRow(blogSqls["queryone"], id).Scan(&art.Id, &art.Author, &art.Title, &art.TitleHtml, &art.Content, &art.ContentHtml, &art.Excerpt, &art.ExcerptHtml, &art.SectionId, &art.SectionName, &art.SectionUrl, &art.CategoryId, &art.CategoryName, &art.CategoryUrl, &art.CommentCount, &art.Status, &art.Posted, &art.LastMod, &art.Expires)
    if err != nil {
        return "", err
    }
    if len(art.SectionUrl) == 0 {
        art.SectionUrl = CreateUrl(BLOG_SECTION, art.SectionId, 0)
    }
    if len(art.CategoryUrl) == 0 {
        art.CategoryUrl = CreateUrl(BLOG_CATEGORY, art.CategoryId, 0)
    }         
    art.Url = CreateUrl(BLOG_ARTICLE, art.Id, 0)
    artView.Art = art
    //previous url
    err = dbconnection.DB.QueryRow(blogSqls["queryprevious"], art.CategoryId, art.Posted).Scan(&artView.PreviousId, &artView.PreviousTitle)
    if err != nil {
        artView.PreviousId = -1
    } else {
        artView.PreviousUrl = CreateUrl(BLOG_ARTICLE, artView.PreviousId, 0)
    } 
    //next url
    err = dbconnection.DB.QueryRow(blogSqls["querynext"], art.CategoryId, art.Posted).Scan(&artView.NextId, &artView.NextTitle)
    if err != nil {
        artView.NextId = -1
    } else {
        artView.NextUrl = CreateUrl(BLOG_ARTICLE, artView.NextId, 0)
    }

    data, err := json.Marshal(artView)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *BlogSrvModel) Insert(id, str string) (string, error) {
    var blog ArticleWriter

    err := json.Unmarshal([]byte(str), &blog)
    if err != nil {
        return "", err
    }

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    //don't need check id because it must be number
    idInt, _ := strconv.ParseInt(id, 10, 64)
    blog.Author = idInt
    now := time.Now()
    nowString := now.Format("2006-01-02 15:04:05")
    res, err := tx.Exec(blogSqls["insert"],  blog.Author, blog.Title, blog.TitleHtml, blog.Content, blog.ContentHtml, blog.Section, blog.Category, 0, blog.Status, nowString, nowString)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    blogid, err := res.LastInsertId()
    if err != nil {
        tx.Rollback()
        return "", err
    }

    tx.Commit()

    blog.Id = blogid
    data, err := json.Marshal(blog)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *BlogSrvModel) Update(id int64, str string) (string, error) {
    /*
    var blog Blog

    err := json.Unmarshal([]byte(str), &blog)
    if err != nil {
        return "", err
    }

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    //just update, not check if it is same before updating. It may be supported in future
    _, err = tx.Exec(blogSqls["update"],  blog.Name, blog.Age, blog.Sex, blog.Id)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    tx.Commit()

    return str, nil
    */
    return "", nil
}

func (model *BlogSrvModel) Delete(id int64) error {
    /*
    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return err
    }

    _, err = tx.Exec(blogSqls["delete"], id)
    if err != nil {
        tx.Rollback()
        return err
    }

    tx.Commit()
    */

    return nil
}
