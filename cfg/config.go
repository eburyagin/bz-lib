package cfg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

type Config struct {
	Datastore DatastoreConfig `json:"datastore"`
	Bus       BusConfig       `json:"bus"`
	Http      HttpConfig      `json:"http"`
}

type DatastoreConfig struct {
	Type     string `json:"type"`
	Addr     string `json:"addr"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type BusConfig struct {
	Type       string `json:"type"`
	Urls       string `json:"urls"`
	Timeout    string `json:"timeout"`
	TimeoutVal time.Duration
}

type HttpConfig struct {
	Port string `json:"port"`
}

func Load(cf string) (*Config, error) {
	var config Config
	raw, err := ioutil.ReadFile(cf)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	err = json.Unmarshal(raw, &config)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if config.Bus.Timeout != "" {
		config.Bus.TimeoutVal, err = time.ParseDuration(config.Bus.Timeout)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &config, nil
}
