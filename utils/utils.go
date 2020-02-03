package utils

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	}

	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Dbname   string `yaml:"dbname"`
		DbPort   string `yaml:"dbPort"`
		Sslmode  string `yaml:"Sslmode"`
		Conn     string `yaml:"Conn"`
	}

	Security struct {
		PrivateKey string `yaml:"privateKey"`
		Hash       string `yaml:"hash"`
	}
}

func Conf() Config {
	t, err := os.Getwd()
	if err != nil {
		logrus.Error("Error get working directory")
	}

	f, err := os.Open(t + "/conf/config.yml")
	if err != nil {
		logrus.Error("Error open  files")
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		logrus.Error("Error get working directory")
	}

	return cfg
}
