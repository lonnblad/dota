package openhttpapi

import (
	"dota/businesslogic"
	"dota/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type httpapi struct {
	listenurl string
	logic     businesslogic.OpenAPILogic
}

var endpoints = []endpoint{
	endpoint{"heroes", getHeroesHandler},
	endpoint{"heroes/{id:[0-9]+}", getHeroByIdHandler},
	endpoint{"items", getItemsHandler},
	endpoint{"items/{id:[0-9]+}", getItemByIdHandler},
}

type endpoint struct {
	name    string
	handler func(http.ResponseWriter, *http.Request)
}

var api httpapi

func listenAndServe() {
	err := http.ListenAndServe(api.listenurl, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func CreateAPI(listenurl string, logic businesslogic.OpenAPILogic) {
	api = httpapi{listenurl, logic}
}

func Run() {
	router := mux.NewRouter()
	for _, endpoint := range endpoints {
		router.HandleFunc("/rest/"+endpoint.name, endpoint.handler).Methods("GET")
	}
	http.Handle("/rest/", router)
	listenAndServe()
}

func getHeroesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	heroes, _ := api.logic.GetHeroes()
	jHeroes := json.CreateJSON(mapToAPIHeroes(heroes))
	w.Write([]byte(`{"data":{"heroes":` + jHeroes + `}}`))
}

func getHeroByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	vars := mux.Vars(r)
	hero, err := api.logic.GetHeroById(vars["id"])
	if err != nil {
		http.Error(w, "404", 404)
		return
	}
	jHero := json.CreateJSON(mapToAPIHero(hero))
	w.Write([]byte(`{"data":{"heroes":[` + jHero + `]}}`))
}

func getItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	items, _ := api.logic.GetItems()
	jItems := json.CreateJSON(mapToAPIItems(items))
	w.Write([]byte(`{"data":{"items":` + jItems + `}}`))
}

func getItemByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	vars := mux.Vars(r)
	item, err := api.logic.GetItemById(vars["id"])
	if err != nil {
		http.Error(w, "404", 404)
		return
	}
	jItem := json.CreateJSON(mapToAPIItem(item))
	w.Write([]byte(`{"data":{"items":[` + jItem + `]}}`))
}
