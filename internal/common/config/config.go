package config

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/naoina/toml"
)

// DEBUG 環境參數值
const DEBUG = "on"

// conf 全域設定變數
var conf *Config

// Config 設定檔
type Config struct {
	// Servers Servers `toml:"servers"`
}

// Servers 服務器設定
// type Servers struct {
// 	Host string `toml:"host"`
// 	Port string `toml:"port"`
// }

func init() {
	load()
}

// GetConfig 取得設定檔
func GetConfig() Config {
	return *conf
}

// load 讀取設定檔
func load() *Config {
	if conf != nil {
		return conf
	}

	var path string
	var err error

	path, err = GetPathOfConfig()
	if err != nil {
		log.Fatalf("讀取 Config路徑錯誤： %v", err)
	}

	var data []byte
	data, err = ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("讀取 Config錯誤： %v", err)
	}

	err = toml.Unmarshal(data, &conf)
	if err != nil {
		log.Fatalf("載入 Config錯誤： %v", err)
	}

	return conf
}

// GetAppRoot 取得專案的根目錄
func GetAppRoot() (string, error) {
	root, err := os.Getwd()
	return root, err
}

// GetAppEnv 取得環境變數
func GetAppEnv() string {
	return os.Getenv("PROJECT_ENV")
}

// GetAppSite 取得服務名稱
func GetAppSite() string {
	return os.Getenv("PROJECT_SITE")
}

// GetAppDebug 取得是否開啟 debug
func GetAppDebug() string {
	return os.Getenv("PROJECT_DEBUG")
}

// GetAppMysqlUser 取得 mysql帳號
func GetAppMysqlUser() string {
	return os.Getenv("PROJECT_MYSQL_USER")
}

// GetAppMysqlPwd 取得 mysql密碼
func GetAppMysqlPwd() string {
	return os.Getenv("PROJECT_MYSQL_ROOT_PASSWORD")
}

// GetAppMysqlDB 取得 mysql db
func GetAppMysqlDB() string {
	return os.Getenv("PROJECT_MYSQL_DATABASE")
}

// GetAppMysqlHost 取得 mysql host
func GetAppMysqlHost() string {
	return os.Getenv("PROJECT_MYSQL_HOST")
}

// GetAppMysqlPort 取得 mysql port
func GetAppMysqlPort() string {
	return os.Getenv("PROJECT_MYSQL_PORT")
}

// GetAppHost 取得服務 host
func GetAppHost() string {
	return os.Getenv("PROJECT_APP_HOST")
}

// GetAppPort 取得服務 port
func GetAppPort() string {
	return os.Getenv("PROJECT_APP_PORT")
}

// GetPathOfConfig 取得專案的設定檔案
func GetPathOfConfig() (string, error) {
	root, err := GetAppRoot()
	if err != nil {
		return "", err
	}

	// TODO 改寫 設定檔讀取方式
	/*
		可以用 ImportDir改寫, 避免寫死目錄與檔案名稱名稱

		pkgInfo, err := build.ImportDir(".", 0)
		if err != nil {
			log.Fatalf("讀取 config路徑錯誤： %v", err)
		}
		pkgInfo...
	*/

	sp := string(os.PathSeparator)
	path := root + sp
	path += "configs" + sp
	path += GetAppEnv() + sp
	path += GetAppSite() + ".toml"

	return path, err
}
