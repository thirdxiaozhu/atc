package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/beego/beego/v2/core/logs"
)

type Config struct {
	SqlUser     string `json:"sql_user"`
	SqlPassword string `json:"sql_password"`
	SqlDatabase string `json:"sql_database"`
	SqlPort     int    `json:"sql_port"`
}

var logger *log.Logger

func init() {
	logger = logs.GetLogger()
}

func ReadConfig(configPath string) Config {
	configFile, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
	}
	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)

	var config Config
	err = json.Unmarshal([]byte(byteValue), &config)

	return config
}
