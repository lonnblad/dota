package main

import (
	"dota/api/openhttpapi"
	"dota/api/scraperapi"
	"dota/businesslogic"
	"dota/datalayer/memorymap"
	"github.com/BurntSushi/toml"
	"log"
)

func main() {
	var conf config

	if _, err := toml.DecodeFile("/etc/dota/dota-backend-service/conf.toml", &conf); err != nil {
		log.Println(err)
		return
	}

	log.Println(conf)

	data := memorymap.NewMemory()

	logic := &businesslogic.Logic{}
	logic.SetDataMap(&data)

	openhttpapi.CreateAPI(conf.OpenHTTPAPI.ListenURL, logic)
	go openhttpapi.Run()

	scraperapi.CreateAPI(conf.ScraperAPI.ListenURL, logic)
	scraperapi.Run()
}

type config struct {
	OpenHTTPAPI openHTTPAPI
	ScraperAPI  scraperAPI
}

type openHTTPAPI struct {
	ListenURL string
}

type scraperAPI struct {
	ListenURL string
}
