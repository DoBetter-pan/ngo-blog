/** 
* @file util.go
* @brief controller util
* @author yingx
* @date 2016-09-04
*/

package controller

import (
	"net/http"
	"fmt"
    "strings"
    "strconv"
)

func GetId(r *http.Request) int64 {
	var id int64 = 1

	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")

    if len(parts) > 1 {
        num, err := strconv.ParseInt(parts[1], 10, 64)
        if err != nil {
            num = 1
        }
        id = num
    }

    return id
}

func SendBack(w http.ResponseWriter, data string) {
    //set header
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, data)
}

func GetError(err error) string {
    str := ""
    if err == nil {
        str = `{ "status": 0, "message": "successful!" }`
    } else {
        str = `{ "status": 1, "message": "unsuccessful!" }`
    }

    return str
}

func ValidateUserCookie(r *http.Request) {
}
