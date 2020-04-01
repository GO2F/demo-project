package config

import (
	"os"
)

// TypeAppConfig Type:基础配置
type TypeAppConfig struct {
	Name string
	Port string
	Host string
	// 当前进程运行路径, 非文件所在路径, 这个需要分清楚
	RunAtPath string
	// Cache路径
	CachePathURI string
	// Log路径
	LogPathURI string
}

type typeSystemIniConfigBase struct {
	IDC               string `ini:"IDC"`
	PORT              string `ini:"PORT"`
	IPADDR            string `ini:"IPADDR"`
	HOSTNAME          string `ini:"HOSTNAME"`
	MatrixCacheDir    string `ini:"CACHE_DIR"`
	MatrixPrivdataDir string `ini:"PRIVDATA_DIR"`
	MatrixApplogsDir  string `ini:"APPLOGS_DIR"`
}

type typeSystemIniConfig struct {
	Base typeSystemIniConfigBase `ini:"base"`
}

// App 基础配置
var App TypeAppConfig

func appInit() {
	currentPathURI, _ := os.Getwd()
	iniConfig := typeSystemIniConfig{
		Base: typeSystemIniConfigBase{
			PORT:             "5060",
			HOSTNAME:         "http://localhost",
			MatrixCacheDir:   currentPathURI,
			MatrixApplogsDir: currentPathURI,
		},
	}

	switch Current {
	case Dev:
		App = TypeAppConfig{
			Name:         "dev",
			Port:         "5060",
			Host:         "http://localhost",
			RunAtPath:    currentPathURI,
			CachePathURI: currentPathURI + "/cache/",
			LogPathURI:   currentPathURI + "/log/",
		}
	case Test:
		App = TypeAppConfig{
			Name:         "test",
			Port:         "5060",
			Host:         "http://m.ke.com",
			RunAtPath:    currentPathURI,
			CachePathURI: currentPathURI + "/cache/",
			LogPathURI:   currentPathURI + "/log/",
		}
	case Prod:
		fallthrough
	default:
		App = TypeAppConfig{
			Name:         "prod",
			Port:         iniConfig.Base.PORT,
			Host:         iniConfig.Base.HOSTNAME,
			RunAtPath:    currentPathURI,
			CachePathURI: iniConfig.Base.MatrixCacheDir,
			LogPathURI:   iniConfig.Base.MatrixApplogsDir,
		}
	}

	// 创建文件路径
	os.MkdirAll(App.CachePathURI, os.ModePerm)
	os.MkdirAll(App.LogPathURI, os.ModePerm)
}
