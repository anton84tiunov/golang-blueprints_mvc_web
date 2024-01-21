package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	// services "../services"
)

type Config struct {
	Server struct {
		Debug bool   `yaml:"debug"`
		Host  string `yaml:"host"`
		Port  int    `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Db       string `yaml:"db"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Dbname   string `yaml:"dbname"`
	} `yaml:"database"`
	SmtpServer struct {
		From     string `yaml:"from"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
	} `yaml:"smtpServer"`
}

var GLOBAL_CONFIG Config

func readYaml(file string, c *Config) {
	f, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		// services.L.Warn(err)
	}
	err = yaml.Unmarshal(f, &c)
	if err != nil {
		fmt.Println(err)
		// services.L.Warn(err)
	}

}

func ReadConfig() Config {
	readYaml("configs/conf.yaml", &GLOBAL_CONFIG)
	// fmt.Println(GLOBAL_CONFIG.Database)
	return GLOBAL_CONFIG
}
