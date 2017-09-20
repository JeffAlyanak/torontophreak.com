package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

var Conf *Config

type Config struct {
	ServeHttp	bool
	HttpPort	int64

	ServeSsl	bool
	ForceSsl	bool
	SslPort		int64
	ServerCert	string
	ServerKey	string

	AssetDir	string
}

func InitConf() error {
	if _, err := toml.DecodeFile("torontophreak.conf", &Conf); err != nil {
		fmt.Println("The torontophreak.conf config isn't valid.")
		fmt.Println(err)
		return err
	}
	return nil
}
