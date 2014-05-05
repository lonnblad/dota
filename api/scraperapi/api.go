package scraperapi

import (
	"dota/businesslogic"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type httpapi struct {
	listenurl string
	logic     businesslogic.ScraperAPILogic
}

type endpoint struct {
	name    string
	handler func(http.ResponseWriter, *http.Request)
}

var endpoints = []endpoint{
	endpoint{"heroes", addHeroHandler},
	endpoint{"items", addItemHandler},
}

var api httpapi

func CreateAPI(listenurl string, logic businesslogic.ScraperAPILogic) {
	api = httpapi{listenurl, logic}
}

func listenAndServe() {
	err := http.ListenAndServe(api.listenurl, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Run() {
	router := mux.NewRouter()
	for _, endpoint := range endpoints {
		router.HandleFunc("/scraper/"+endpoint.name, endpoint.handler).Methods("POST")
	}
	http.Handle("/scraper/", router)
	listenAndServe()
}

func addHeroHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("addHeroHandler")
	var sHero Hero
	readAndDecodeReq(req, &sHero)
	api.logic.AddHero(mapFromAPIHeroToHero(sHero))
}

func addItemHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("addItemHandler")
	var sItem Item
	readAndDecodeReq(req, &sItem)
	api.logic.AddItem(mapFromAPIItemToItem(sItem))
}

func readAndDecodeReq(req *http.Request, element interface{}) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("error:", err)
	}
	defer req.Body.Close()

	err = json.Unmarshal([]byte(data), element)
	if err != nil {
		log.Println("error:", err)
	}
}
