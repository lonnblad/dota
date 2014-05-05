package main

import (
	"dota/scrapers/dota2Wiki"
	"github.com/BurntSushi/toml"
	"log"
)

func main() {
	var conf config

	if _, err := toml.DecodeFile(
		"/etc/dota/dota-scraper-service/conf.toml",
		&conf); err != nil {
		log.Println(err)
		return
	}
	log.Println(conf)

	dota2Wiki.GetHeroes(conf.Heroes)
	dota2Wiki.GetItems(conf.Items)
}

type config struct {
	Heroes []string
	Items  []string
}
