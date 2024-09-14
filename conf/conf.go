package conf

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"sync"
)

var (
	conf   *Config
	vipers *viper.Viper
	once   sync.Once
)

type Config struct {
	Env      string
	Service  Service  `yaml:"service"`
	Registry Registry `yaml:"registry"`
}

type Registry struct {
	RegistryAddress []string `yaml:"RegistryAddress"`
	Username        string   `yaml:"UserName"`
	Password        string   `yaml:"Password"`
}

type Service struct {
	ServiceName string `yaml:"ServiceName"`
	Address     string `yaml:"Address"`
	ToolVersion string `yaml:"ToolVersion"`
}

func initConf() {
	prefix := "conf"
	confFileRelPath := filepath.Join(prefix, filepath.Join(GetEnv(), "conf.yaml"))

	// 初始化 viper
	vipers = viper.New()
	vipers.SetConfigFile(confFileRelPath)
	vipers.SetConfigType("yaml")
	if err := vipers.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}

	// 监听配置文件
	vipers.WatchConfig()
	vipers.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := vipers.Unmarshal(&conf); err != nil {
			klog.Errorf("unmarshal conf failed, err: %v", err)
		}
	})
	// 将配置赋值给全局变量
	if err := vipers.Unmarshal(&conf); err != nil {
		klog.Errorf("unmarshal conf failed, err: %v", err)
	}
}

func GetEnv() string {
	e := os.Getenv("GO_ENV")
	if len(e) == 0 {
		return "test"
	}
	return e
}

func GetConfig() *Config {
	once.Do(initConf)
	return conf
}
