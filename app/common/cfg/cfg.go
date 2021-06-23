package cfg

import (
	"fmt"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	instance *config
	once     sync.Once
)

//获取配置文档实例
func Instance() *config {
	once.Do(func() {
		instance = initConfig()
	})

	return instance
}

func initConfig() *config {
	var conf config
	config := "./conf/config.yaml"
	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&conf); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&conf); err != nil {
		fmt.Println(err)
	}
	return &conf
}

type config struct {
	Status   status
	Admin    admin
	Api      api
	Task     task
	Database database
	Logger   logger
	Gen      gen
}

type status struct {
	Admin bool `yaml:"admin"`
	Api   bool `yaml:"api"`
}

type admin struct {
	Address    string `yaml:"address"`
	ServerRoot string `yaml:"serverRoot"`
	Swagger    string `yaml:"swagger"`
}

type jwt struct {
	Timeout    int    `yaml:"timeout"`
	Refresh    int    `yaml:"refresh"`
	EncryptKey string `yaml:"encryptKey"`
}

type api struct {
	Address    string `yaml:"address"`
	ServerRoot string `yaml:"serverRoot"`
	Jwt        jwt
}

type task struct {
	WorkPoolSize int `yaml:"workPoolSize"`
}

type database struct {
	Master  string `yaml:"master"`
	Slave   string `yaml:"slave"`
	Debug   bool   `yaml:"debug"`
	Log     string `yaml:"log"`
	LogMode bool   `yaml:"logMode"`
}

type logger struct {
	Path  string `yaml:"path"`
	Level string `yaml:"level"`
}

type gen struct {
	Author        string `yaml:"author"`
	ModuleName    string `yaml:"moduleName"`
	PackageName   string `yaml:"packageName"`
	AutoRemovePre bool   `yaml:"autoRemovePre"`
	TablePrefix   string `yaml:"tablePrefix"`
}
