package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type LunarGateway struct {
	Conf Config `yaml:"lunar-gateway"`
}

type Config struct {
	ConfVersion  string         `yaml:"version"`
	ServerConf   ServerConfig   `yaml:"server"`
	ServicesConf ServicesConfig `yaml:"services"`
}

type ServerSecurity struct {
	Cert    string `yaml:"cert"`
	Key     string `yaml:"key"`
	Trusted string `yaml:"trusted"`
}

type ServerConfig struct {
	Security ServerSecurity `yaml:"security"`

	Address        string `yaml:"address"`
	DialTimeout    int    `yaml:"dialtimeout"`
	RequestTimeout int    `yaml:"requesttimeout"`
}

type ServicesConfig struct {
	Celestial Address `yaml:"celestial"`
}

type Address struct {
	Addr string `yaml:"address"`
}

func ConfigFile(n ...string) (*Config, error) {
	path := "config.yml"
	if len(n) > 0 {
		path = n[0]
	}

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var conf LunarGateway
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		return nil, err
	}

	return &conf.Conf, nil
}

func DefaultConfig() *Config {
	sec := ServerSecurity{
		Cert:    "",
		Key:     "",
		Trusted: "",
	}

	server := ServerConfig{
		Security:       sec,
		Address:        "localhost:8080",
		DialTimeout:    2,
		RequestTimeout: 10,
	}

	celaddr := Address{
		Addr: "localhost:8080",
	}

	conn := ServicesConfig{
		Celestial: celaddr,
	}

	conf := Config{
		ConfVersion:  "v1",
		ServerConf:   server,
		ServicesConf: conn,
	}

	return &conf
}
