package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type ConfigFile struct {
	ServerPort string `json:"server-port"`
	ProxyHosts []struct {
		Domain   string `json:"domain"`
		Protocol string `json:"protocol"`
		Host     string `json:"host"`
		Port     string `json:"port"`
	} `json:"proxy-hosts"`
}

type Config struct {
	File ConfigFile
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) GetServerPort() string {
	return c.File.ServerPort
}

func (c *Config) LoadConfigFile(path string) error {
	// Open the JSON file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	// Read the file content
	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	// Unmarshal the JSON content into a Config struct
	var configFile ConfigFile
	err = json.Unmarshal(byteValue, &configFile)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return err
	}
	c.File = configFile
	return nil
}
