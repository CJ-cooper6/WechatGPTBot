package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

// Configuration 项目配置
type Configuration struct {
	URL         string `json:"url"`
	AccessToken string `json:"access_token"`
}

var config *Configuration
var once sync.Once

// LoadConfig 加载配置
func LoadConfig() *Configuration {
	once.Do(func() {
		// 从文件中读取
		config = &Configuration{}
		f, err := os.Open("config.json")
		if err != nil {
			log.Fatalf("open config err: %v", err)
			return
		}
		defer f.Close()
		encoder := json.NewDecoder(f)
		err = encoder.Decode(config)
		if err != nil {
			log.Fatalf("decode config err: %v", err)
			return
		}
	})
	return config
}
