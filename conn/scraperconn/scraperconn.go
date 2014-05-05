package scraperconn

import (
	"bytes"
	"log"
	"net/http"
)

var url = "http://localhost:10081/scraper"
var bodyType = ""

func PostHero(bHero []byte) {
	postToEndpoint(bHero, "/heroes")
}

func PostItem(bItem []byte) {
	postToEndpoint(bItem, "/items")
}

func postToEndpoint(b []byte, endpoint string) {
	body := bytes.NewReader(b)
	_, err := http.Post(url+endpoint, bodyType, body)
	if err != nil {
		log.Printf("Post error: %+v", err)
	}
}
