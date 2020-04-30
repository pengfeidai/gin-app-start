package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Yaml struct {
	Server `yaml:"server"`
	Redis  `yaml:"redis"`
	Mysql  `yaml:"mysql"`
	Mongo  `yaml:"mongo"`
	Log    `yaml:"log"`
}

type Server struct {
	Port         string `yaml:"port"`
	Mode         string `yaml:"mode"`
	ReadTimeout  int    `yaml:"readTimeout"`
	WriteTimeout int    `yaml:"writeTimeout"`
	LimitNum     int    `yaml:"limitNum"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
	UseRedis bool   `yaml:"useRedis"`
}

type Mysql struct {
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Database     string `yaml:"database"`
	Config       string `yaml:"config"`
	Driver       string `yaml:"drive"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	Log          bool   `yaml:"log"`
}

type Mongo struct {
	Database string `yaml:"database"`
	Url      string `yaml:"url"`
}

type Log struct {
	AccessLogName string `yaml:"accessLogName"`
	ErrorLogName  string `yaml:"errorLogName"`
}

var Conf *Yaml

const defaultConfigFile = "config/config.yaml"

func Init() {
	c := &Yaml{}
	configFile := flag.String("c", defaultConfigFile, "help config path")
	flag.Parse()
	yamlFile, err := ioutil.ReadFile(*configFile)
	if err != nil {
		panic(fmt.Errorf("get yamlFile error: %s", err))
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Printf("config yamlFile load Init success.")
	Conf = c
}
