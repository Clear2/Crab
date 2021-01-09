package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	WeChat WeChat `mapstructure:"wechat" json:"wechat"`
}

type WeChat struct {
	AppId     string `mapstructure:"appid" json:"appid"`
	AppSecret string `mapstructure:"appsecret" json:"appsecret"`
}

var (
	C = new(Config)
	once sync.Once
)


func init() {
	once.Do(func() {
		v := viper.New()
		v.SetConfigFile("./config.yaml")
		err := v.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error load config file: %s\n", err))
		}
		v.WatchConfig()
		v.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("config file changed: ", e.Name)

			if err := v.Unmarshal(&C); err != nil {
				fmt.Println("config load Fail", err)
			}
		})

		if err := v.Unmarshal(&C); err != nil {
			fmt.Println(err)
		}
	})
}