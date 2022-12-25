package config

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/geekr-dev/go-rest-api/model"
	"github.com/geekr-dev/go-rest-api/pkg/log"
	"github.com/spf13/viper"
)

var Data *AppConfig

type Config struct {
	Name string
}

type AppConfig struct {
	Name string
	Mode string
	Addr string
	URL  string
	Log  *log.Config
	Db   *model.Config
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	if err := c.initConfig(); err != nil {
		return err
	}

	c.watchConfig()

	return nil
}

// 初始化配置文件
func (c *Config) initConfig() error {
	// 如果用户未指定配置文件，则使用 conf/config.yaml
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	// 允许从环境变量读取配置
	viper.AutomaticEnv()
	viper.SetEnvPrefix("RESTAPI")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	// 读取所有配置并解析
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.Unmarshal(&Data)
	return nil
}

// 监听配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		viper.Unmarshal(&Data) // 重新反序列化配置数据
		// fmt.Printf("Config file changed: %s\n", e.Name)
	})
}
