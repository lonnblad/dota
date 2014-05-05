package parsers

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

var pl_data string
var dp_data string

func init() {
	pl_data = getDota2Wiki("Phantom_Lancer")
	dp_data = getDota2Wiki("Death_Prophet")
	// log.Printf("\n%+v\n", data)
}

func getDota2Wiki(key string) string {
	res, err := http.Get("http://dota2.gamepedia.com/" + key)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	return string(result)
}

func Test_GetHeroName_Err(t *testing.T) {
	_, err := ParseHeroName("")
	if err.Error() != "no hero name found" {
		t.Errorf("%+v", err)
	}
}

func Test_GetHeroName_PL(t *testing.T) {
	ans, err := ParseHeroName(pl_data)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if ans != "Phantom Lancer" {
		t.Errorf("%+v", ans)
	}
}

func Test_GetHeroName_DP(t *testing.T) {
	ans, err := ParseHeroName(dp_data)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if ans != "Death Prophet" {
		t.Errorf("%+v", ans)
	}
}

func Test_GetFaction_Err(t *testing.T) {
	_, err := ParseHeroFaction("")
	if err.Error() != "no faction found" {
		t.Errorf("%+v", err)
	}
}

func Test_GetFaction_PL(t *testing.T) {
	ans, err := ParseHeroFaction(pl_data)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if ans != "The Radiant" {
		t.Errorf("%+v", ans)
	}
}

func Test_GetFaction_DP(t *testing.T) {
	ans, err := ParseHeroFaction(dp_data)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if ans != "The Dire" {
		t.Errorf("%+v", ans)
	}
}

func Test_GetStrength_Err(t *testing.T) {
	_, err := ParseHeroStrength("")
	if err.Error() != "Str not found" {
		t.Errorf("%+v", err)
	}
}

func Test_GetStrength_PL(t *testing.T) {
	answer, err := ParseHeroStrength(pl_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 2 || answer[0] != 18 || answer[1] != 1.7 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetStrength_DP(t *testing.T) {
	answer, err := ParseHeroStrength(dp_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 2 || answer[0] != 19 || answer[1] != 2.2 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetAgility_Err(t *testing.T) {
	_, err := ParseHeroAgility("")
	if err.Error() != "Agi not found" {
		t.Errorf("%+v", err)
	}
}

func Test_GetAgility_PL(t *testing.T) {
	answer, err := ParseHeroAgility(pl_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 2 || answer[0] != 23 || answer[1] != 4.2 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetAgility_DP(t *testing.T) {
	answer, err := ParseHeroAgility(dp_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 2 || answer[0] != 14 || answer[1] != 1.4 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetIntelligence_Err(t *testing.T) {
	_, err := ParseHeroIntelligence("")
	if err.Error() != "Int not found" {
		t.Errorf("%+v", err)
	}
}

func Test_GetIntelligence_PL(t *testing.T) {
	answer, err := ParseHeroIntelligence(pl_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 2 || answer[0] != 21 || answer[1] != 2 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetIntelligence_DP(t *testing.T) {
	answer, err := ParseHeroIntelligence(dp_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 2 || answer[0] != 20 || answer[1] != 3 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetPrimary_Err(t *testing.T) {
	_, err := ParseHeroPrimary("")
	if err.Error() != "Primary not found" {
		t.Errorf("%+v", err)
	}
}

func Test_GetPrimary_PL(t *testing.T) {
	answer, err := ParseHeroPrimary(pl_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if answer != "Agi" {
		t.Errorf("%+v", answer)
	}
}

func Test_GetPrimary_DP(t *testing.T) {
	answer, err := ParseHeroPrimary(dp_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if answer != "Int" {
		t.Errorf("%+v", answer)
	}
}
