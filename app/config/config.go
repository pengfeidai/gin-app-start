package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Yaml struct {
	Server  `yaml:"server"`
	Redis   `yaml:"redis"`
	Mysql   `yaml:"mysql"`
	Mongo   `yaml:"mongo"`
	Session `yaml:"session"`
	Log     `yaml:"log"`
	Url     `yaml:"url"`
}

type Server struct {
	Port      int    `yaml:"port"`
	Mode      string `yaml:"mode"`
	LimitNum  int    `yaml:"limitNum"`
	UserMongo bool   `yaml:"useMongo"`
	UserRedis bool   `yaml:"useRedis"`
}

type Redis struct {
	Addr     string `yaml:"Addr"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type Mysql struct {
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Path         string `yaml:"path"`
	Database     string `yaml:"database"`
	Config       string `yaml:"config"`
	Driver       string `yaml:"driver"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	Log          bool   `yaml:"log"`
}

type Mongo struct {
	Database string `yaml:"database"`
	Url      string `yaml:"url"`
}

type Session struct {
	Key      string `yaml:"key"`
	Size     int    `yaml:"size"`
	MaxAge   int    `yaml:"maxAge"`
	Path     string `yaml:"path"`
	Domain   string `yaml:"domain"`
	HttpOnly bool   `yaml:"httpOnly"`
}

type Log struct {
	AccessLogFile string `yaml:"accessLogFile"`
	ErrorLogFile  string `yaml:"errorLogFile"`
}

type Url struct {
	Prefix string `yaml:"prefix"`
}

var Conf *Yaml

const defaultConfigFile = "config.yaml"

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
