package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type LunarGateway struct {
	Conf Config `yaml:"lunar-gateway"`
}

type Config struct {
	ConfVersion  string            `yaml:"version"`
	ServerConf   ServerConfig      `yaml:"server"`
	ServicesConf map[string]string `yaml:"services"`
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

	conf := Config{
		ConfVersion: "v1",
		ServerConf:  server,
		ServicesConf: map[string]string{
			"celestial": "localhost:8000",
			"blackhole": "localhost:8001",
		},
	}
	return &conf
}
