package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	Config     *Config     `yaml:"config"`
	ServerChan *ServerChan `yaml:"serverChan"`
}

type Config struct {
	Cookie    string `yaml:"cookie"`
	UserAgent string `yaml:"userAgent"`
}

type ServerChan struct {
	SecretKey string `yaml:"secretKey"`
}

func Init(path string) (*Conf, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	conf := &Conf{}
	err = yaml.Unmarshal(data, conf)
	return conf, err
}
