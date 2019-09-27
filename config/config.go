package config

import (
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/linnv/logx"
	yaml "gopkg.in/yaml.v2"

	"github.com/linnv/logx/version"
)

type Configuration struct {
	AppName    string `yaml:"appName"`
	ServerPort string `yaml:"serverPort"`
	DevMode    bool   `yaml:"devMode"`
	LogDir     string `yaml:"logDir"`
}

func GetDefaultConfigPath() string {
	configFile := "./config/config.yaml"
	return configFile
}

const copyright = " Copyright ©2018-%d jialinwu.com 版权所有"

var flagOnce sync.Once

var ConfigFile *string

var defaultConfigPath = GetDefaultConfigPath()

func InitFlag() {
	flagOnce.Do(func() {
		ConfigFile = flag.String("c", defaultConfigPath, "absolute path of config.yaml")
	})
}

func initConfig() (config *Configuration, err error) {
	version.COYPRIGHT = copyright
	version.ReadBuildInfo()

	if ConfigFile == nil {
		ConfigFile = &defaultConfigPath
	}
	configFile := *ConfigFile

	logx.Debugf("load configFile: %+v\n", configFile)
	bs, err := ioutil.ReadFile(configFile)
	if err != nil {
		logx.Warnf("err: %+v\ntry /abs/appName -c /abs/config.yaml\n ", err)
		return nil, err
	}

	config = new(Configuration)
	err = yaml.Unmarshal([]byte(bs), &config)
	if err != nil {
		logx.Warnf("err: %+v\n", err)
		return nil, err
	}

	cantEmpty := func(str, name string) {
		str = strings.TrimSpace(str)
		if str == "" {
			panic(name + " can't be empty")
		}
	}

	if config.AppName == "" {
		config.AppName = "tplgo"
	}
	if config.LogDir == "" {
		curDir := path.Join(defaultConfigPath, "..", "..", "log")
		config.LogDir = path.Join(curDir, "../log")
		logx.Debugf("using defalut logDir: %+v\n", config.LogDir)
	} else {
		if err := os.MkdirAll(config.LogDir, 0777); err != nil {
			logx.Errorf("err: %+v\n", err)
		}
	}
	cantEmpty(config.LogDir, "LogDir")

	hostname, _ := os.Hostname()
	logPrefix = "AN=go-" + config.AppName + "@" + hostname

	logx.Warnf("config: %#v\n", config)

	if !config.DevMode {
		logx.EnableDevMode(false)
	}
	rootConfig = config

	httpClient = &http.Client{
		Timeout: time.Second * 60,
	}

	return
}

var once sync.Once

var logPrefix string
var rootConfig *Configuration
var httpClient *http.Client

func InitConfig() *Configuration {
	once.Do(func() {
		_, err := initConfig()
		if err != nil {
			time.Sleep(time.Millisecond * 600)
			panic(err.Error())
		}
	})
	return Config()
}

func Config() *Configuration {
	if rootConfig == nil {
		panic("init config first")
	}
	return rootConfig
}

func GetHttpClient() *http.Client {
	if httpClient == nil {
		panic("init config first")
	}
	return httpClient
}
