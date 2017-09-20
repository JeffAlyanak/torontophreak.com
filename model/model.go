package model

import (
	"io/ioutil"
	"github.com/BurntSushi/toml"
	"toronto-phreak/config"
)

type Article struct {
	Id		string
	Title	string
	Desc	string
	Full	string
	Thumb	string
}

type ArticleList []Article

var Articles ArticleList

func LoadEntries() error {
	entries, err	:= ioutil.ReadDir(config.Conf.AssetDir + "pages/")
	if err != nil {
		return err
	}
	for _, entry := range entries {
		// fmt.Println(entries.Name())
		var a Article
		dir	:= config.Conf.AssetDir + "pages/" + entry.Name() + "/text.txt"
		if _, err := toml.DecodeFile(dir, &a); err != nil {
			return err
		}
		a.Full		= config.Conf.AssetDir + "pages/" + a.Id + "/" + a.Full
		a.Thumb		= config.Conf.AssetDir + "pages/" + a.Id + "/" + a.Thumb
		Articles	= append(Articles, a)
	}
	return nil
}
