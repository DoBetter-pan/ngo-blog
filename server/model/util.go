/** 
* @file util.go
* @brief controller util
* @author yingx
* @date 2016-09-04
*/

package model

import (
	"fmt"
)

const (
    BLOG_SECTION = 0
    BLOG_CATEGORY = 1
    BLOG_ARTICLE = 2
)

func CreateUrl(t int, id int64, p int64) string{
    str := ""
    switch t {
    //section
    case BLOG_SECTION:
        str = fmt.Sprintf("#/list?s=%d&p=%d", id, p)
    //category
    case BLOG_CATEGORY:
        str = fmt.Sprintf("#/list?c=%d&p=%d", id, p)
    //article
    case BLOG_ARTICLE:
        str = fmt.Sprintf("#/view/%d", id)
    }

    return str
}

func CreateUrlByKeyValue(key string, value, p int64) string {
    str := fmt.Sprintf("#/list?%s=%d&p=%d", key, value, p)
    return str
}

