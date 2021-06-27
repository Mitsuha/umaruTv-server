package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type app struct {
	ServerAddr string `yaml:"ServerAddr"`
	Env string `yaml:"Env"`
	Key string `yaml:"Key"`
	Debug bool `yaml:"Debug"`
}

type db struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
	Database string `yaml:"Database"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
}

type redis struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
	Password string `yaml:"Password"`
}

var (
	App app
	DB db
	Redis redis
)

func init() {
	configFile, err := os.Open("./config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	var config = struct{
		App app `yaml:"App"`
		DB db `yaml:"DB"`
		Redis redis `yaml:"Redis"`
	}{}

	data, _ := ioutil.ReadAll(configFile)

	if err = yaml.Unmarshal(data, &config); err != nil {
		log.Fatalln(err)
	}

	App, DB, Redis = config.App, config.DB, config.Redis
}

func (d *db) GetDSN() string {
	fmt.Println(d)
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.Database,
	)
}