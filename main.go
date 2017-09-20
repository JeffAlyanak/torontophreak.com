package main

import (
	"fmt"
	"toronto-phreak/config"
	"toronto-phreak/ui"
	"toronto-phreak/model"
	"toronto-phreak/daemon"
)

func main() {
	if err := config.InitConf(); err != nil {
		fmt.Println(err)
		return
	}

	if err := ui.CacheTemplates(); err != nil {
		fmt.Println(err)
		return
	}
	if err := model.LoadEntries(); err != nil {
		fmt.Println(err)
		return
	}

	if err := daemon.Start(); err != nil {
		fmt.Println(err)
		return
	}
}
