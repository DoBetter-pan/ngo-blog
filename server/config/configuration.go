/** 
* @file configuration.go
* @brief configuration
* @author yingx
* @date 2016-11-28
*/

package configuration

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

type DatbaseConfig struct {
    Protocal     string  `json:"protocal"`
    Host         string  `json:"host"`
    Port         int64  `json:"port"`
    User         string  `json:"user"`
    Password     string `json:"password"`
    Db           string  `json:"Db"`
    Option       string  `json:"option"`
}

type NgoBlogConfig struct {
    Version  string               `json:"v"`
    Db       DatbaseConfig        `json:"Db"`
}

var ngoBlogConfig *NgoBlogConfig = &NgoBlogConfig{}
var configFile string = "ngoblog.json"

func init() {
    data, err := ioutil.ReadFile(configFile)
    if err == nil {
        err = json.Unmarshal(data, config)
    }
    if err != nil {
        log.Fatal("ngoBlogConfig: ", err)
    }
    return config, err
}


func GetNgoBlogConfig() *NgoBlogConfig {
    return ngoBlogConfig
}


