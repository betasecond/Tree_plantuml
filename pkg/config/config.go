package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config 定义了配置文件的结构
type Config struct {
	SourcePath string `json:"sourcePath"`
	Features   struct {
		ContextMenu      bool     `json:"contextMenu"`
		ClipboardSupport bool     `json:"clipboardSupport"`
		OutputFormats    []string `json:"outputFormats"`
	} `json:"features"`
	OutputDirectory string `json:"outputDirectory"`
	RegisterKeys    []struct {
		Name    string   `json:"name"`
		Command string   `json:"command"`
		Formats []string `json:"formats"`
	} `json:"registerKeys"`
}

// LoadConfig 从指定路径加载配置文件
func LoadConfig(configPath string) (*Config, error) {
	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	bytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
