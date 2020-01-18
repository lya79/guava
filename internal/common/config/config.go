package config

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/naoina/toml"
)

// Config 設定檔
type Config struct {
	Chatroom Addr  `toml:"chatroom"` // chatroom服務
	MySQL    MySQL `toml:"mysql"`    // mysql服務
	Redis    Redis `toml:"redis"`    // redis服務
}

// Addr 服務器設定
type Addr struct {
	Host string `toml:"host"`
	Port int64  `toml:"port"`
}

// MySQL mysql資料庫設定
type MySQL struct {
	Host            string `toml:"host"`
	Port            int64  `toml:"port"`
	User            string `toml:"user"`
	Passwrod        string `toml:"password"`
	Database        string `toml:"database"`
	LogMode         bool   `toml:"logMode"`
	MaxIdleConns    int64  `toml:"maxIdleConns"`
	MaxOpenConns    int64  `toml:"maxOpenConns"`
	ConnMaxLifetime int64  `toml:"connMaxLifetime"`
}

// Redis redis設定
type Redis struct {
	Host        string `toml:"host"`
	Port        int64  `toml:"port"`
	Passwrod    string `toml:"password"`
	Database    string `toml:"database"`
	PoolSize    int64  `toml:"pool_size"`
	IdleTimeout int64  `toml:"idle_timeout"`
}

// conf 全域設定變數
var conf *Config

func init() {
	conf = load()
}

// GetConfig 取得設定檔
func GetConfig() Config {
	return *conf
}

// IsLocalEnv 是否為本機端環境
func IsLocalEnv() bool {
	if getAppEnv() == "local" {
		return true
	}
	return false
}

// getAppRoot 取得專案的根目錄
func getAppRoot() (string, error) {
	var err error
	root := os.Getenv("PROJECT_ROOT")
	if root == "" {
		root, err = os.Getwd()
		if err != nil {
			return root, err
		}
	}
	return root, err
}

// GetAppSite 取得服務名稱
func GetAppSite() string {
	return os.Getenv("PROJECT_SITE")
}

// getAppEnv 取得環境變數
func getAppEnv() string {
	return os.Getenv("PROJECT_ENV")
}

// load 讀取設定檔
func load() *Config {
	if conf != nil {
		return conf
	}

	var path string
	var err error

	path, err = getPathOfConfig()
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

// GetPathOfConfig 取得專案的設定檔案
func getPathOfConfig() (string, error) {
	root, err := getAppRoot()
	if err != nil {
		return "", err
	}

	sp := string(os.PathSeparator)
	path := root + sp
	path += "configs" + sp
	path += getAppEnv() + sp
	path += GetAppSite() + ".toml"

	return path, err
}
