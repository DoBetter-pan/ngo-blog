/** 
* @file configuration.go
* @brief configuration
* @author yingx
* @date 2016-11-28
*/

package configuration

import (
    _ "fmt"
    "log"
    "encoding/json"
    "io/ioutil"
)

type DbConfig struct {
    Protocal     string  `json:"protocal"`
    Host         string  `json:"host"`
    Port         string  `json:"port"`
    User         string  `json:"user"`
    Password     string  `json:"password"`
    Database     string  `json:"database"`
    Options      string  `json:"options"`
}

type NgoBlogConfig struct {
    Version  string               `json:"v"`
    Db       DbConfig             `json:"Db"`
}


var ngoBlogConfig *NgoBlogConfig = &NgoBlogConfig{}
var configFile string = "config.json"

func init() {
    data, err := ioutil.ReadFile(configFile)
    if err == nil {
        err = json.Unmarshal(data, ngoBlogConfig)
    }
    if err != nil {
        log.Fatal("ngoBlogConfig: ", err)
    }
}


func GetNgoBlogConfig() *NgoBlogConfig {
    return ngoBlogConfig
}


