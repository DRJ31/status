package service

import (
	"encoding/json"
	"github.com/DRJ31/status/model"
	"log"
	"os"
	"sync"
)

var cfg *model.Config
var once sync.Once

func GetConfig() *model.Config {
	once.Do(func() {
		cfg = &model.Config{}
		jsonFile, err := os.Open("config.json")
		if err != nil {
			log.Fatal("[Error] config.json 配置文件不存在")
		}
		defer jsonFile.Close()

		err = json.NewDecoder(jsonFile).Decode(cfg)
		if err != nil {
			log.Fatal("[Error] 配置文件解析失败")
		}
	})
	return cfg
}
