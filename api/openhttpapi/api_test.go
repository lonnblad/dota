package openhttpapi

import (
	"dota/businesslogic"
	"dota/datalayer/memorymap"
	"dota/json"
	"dota/testutils"
	"io/ioutil"
	"net/http"
	"testing"
)

var url = "http://localhost:20080/rest"

func init() {
	data := memorymap.NewMemory()
	data.LoadData()

	logic := &businesslogic.Logic{}
	logic.SetDataMap(&data)

	CreateAPI(":20080", logic)
	go Run()
}

func Test_Get_Heroes(t *testing.T) {
	hero0 := json.CreateJSON(Hero{Id: "0", Type: "hero"})
	hero1 := json.CreateJSON(Hero{Id: "1", Type: "hero"})
	res, _ := http.Get(url + "/heroes")
	result, _ := ioutil.ReadAll(res.Body)
	testutils.Assert(string(result), `{"data":{"heroes":[`+hero0+`,`+hero1+`]}}`, t)
}

func Test_Get_HeroById_Success(t *testing.T) {
	hero := json.CreateJSON(Hero{Id: "1", Type: "hero"})
	res, _ := http.Get(url + "/heroes/1")
	result, _ := ioutil.ReadAll(res.Body)
	testutils.Assert(string(result), `{"data":{"heroes":[`+hero+`]}}`, t)
}

func Test_Get_HeroById_Error(t *testing.T) {
	res, _ := http.Get(url + "/heroes/01")
	result, _ := ioutil.ReadAll(res.Body)
	testutils.Assert(string(result), "404\n", t)
}

func Test_Get_Items(t *testing.T) {
	item := json.CreateJSON(Item{Id: "0", Type: "item"})
	res, _ := http.Get(url + "/items")
	result, _ := ioutil.ReadAll(res.Body)
	testutils.Assert(string(result), `{"data":{"items":[`+item+`]}}`, t)
}

func Test_Get_ItemsById_Success(t *testing.T) {
	item := json.CreateJSON(Item{Id: "0", Type: "item"})
	res, _ := http.Get(url + "/items/0")
	result, _ := ioutil.ReadAll(res.Body)
	testutils.Assert(string(result), `{"data":{"items":[`+item+`]}}`, t)
}

func Test_Get_ItemsById_Error(t *testing.T) {
	res, _ := http.Get(url + "/items/01")
	result, _ := ioutil.ReadAll(res.Body)
	testutils.Assert(string(result), "404\n", t)
}
