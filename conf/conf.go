package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	WeChat WeChat `mapstructure:"wechat" json:"wechat"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql"`
	JWT		JWT   `mapstructure:"jwt" json:"mysql"`
}

// jwt
type JWT struct {
	SigningKey string `mapstructure:"signing-key" json:"signing_key" yaml:"signing-key"`
}
// 微信
type WeChat struct {
	AppId     string `mapstructure:"appid" json:"appid"`
	AppSecret string `mapstructure:"appsecret" json:"appsecret"`
}

// mysql
type Mysql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	LogZap       bool   `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
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