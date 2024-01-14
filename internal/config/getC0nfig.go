package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
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
}

var GLOBAL_CONFIG Config

func readYaml(file string, c *Config) {
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(f, &c)
	if err != nil {
		panic(err)
	}

}

func ReadConfig() Config {
	readYaml("configs/conf.yaml", &GLOBAL_CONFIG)
	// fmt.Println(GLOBAL_CONFIG.Database)
	return GLOBAL_CONFIG
}
