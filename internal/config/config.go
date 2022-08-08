package config

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	FilePaths    []string `yaml:"filepaths"`
	WebFolder    string   `yaml:"webFolder"`
	DatabasePath string   `yaml:"databasepath"`
}

// 建立 Config 並從設定檔中讀取設定
// 回傳 Config
func NewConfig(configPath string) (*Config, error) {
	// 建立 Config
	config := &Config{}

	// 開啟設定檔
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 使用 yaml 解碼器
	d := yaml.NewDecoder(file)

	// 開始讀取設定
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	// 回傳 Config
	return config, nil
}

// 取得設定檔的路徑，並且嘗試從命令列參數中取得使用者自訂的路徑
// 命令列參數 "-config [設定檔路徑]"
// 設定檔路徑預設： config.yml
func GetConfigPath() (string, error) {
	var configPath string

	// 建置 flag 字串變數
	// CLI 參數： "-config"
	// 預設： "./config.yml"
	// 使用說明： "設定檔路徑"
	flag.StringVar(&configPath, "config", "./config.yml", "設定檔路徑")

	// 嘗試從 CLI 取得參數
	flag.Parse()

	// 檢查檔案路徑是否存在
	s, err := os.Stat(configPath)
	if err != nil {
		return "", err
	}

	// 判斷是否為資料夾
	if s.IsDir() {
		return "", fmt.Errorf("'%s' 是資料夾，並不能作為設定檔的路徑。", configPath)
	}

	// 回傳設定檔路徑
	return configPath, nil
}
