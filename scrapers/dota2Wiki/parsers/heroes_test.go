package parsers

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

var pl_data string
var dp_data string
var np_data string
var io_data string
var medusa_data string
var tw_data string
var tiny_data string
var am_data string

func init() {
	pl_data = getDota2Wiki("Phantom_Lancer")
	dp_data = getDota2Wiki("Death_Prophet")
	np_data = getDota2Wiki("Nature's_Prophet")
	io_data = getDota2Wiki("Io")
	medusa_data = getDota2Wiki("Medusa")
	tw_data = getDota2Wiki("Troll_Warlord")
	tiny_data = getDota2Wiki("Tiny")
	am_data = getDota2Wiki("Anti-Mage")
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

func Test_GetHeroName_NP(t *testing.T) {
	ans, err := ParseHeroName(np_data)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if ans != "Nature's Prophet" {
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

func Test_GetHeroName_AM(t *testing.T) {
	ans, err := ParseHeroName(am_data)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if ans != "Anti-Mage" {
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

func Test_GetMovementSpeed_PL(t *testing.T) {
	answer, err := ParseHeroMovementSpeed(pl_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if answer != 290 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetMovementSpeed_DP(t *testing.T) {
	answer, err := ParseHeroMovementSpeed(dp_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if answer != 280 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetTurnRate_PL(t *testing.T) {
	answer, err := ParseHeroTurnRate(pl_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if answer != 0.6 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetTurnRate_DP(t *testing.T) {
	answer, err := ParseHeroTurnRate(dp_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if answer != 0.5 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetSightRange_PL(t *testing.T) {
	answer, err := ParseHeroSightRange(pl_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 2 || answer[0] != 1800 || answer[1] != 800 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetAttackRange_DP(t *testing.T) {
	answer, err := ParseHeroAttackRange(dp_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if answer != "600" {
		t.Errorf("%+v", answer)
	}
}

func Test_GetAttackRange_PL(t *testing.T) {
	answer, err := ParseHeroAttackRange(pl_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if answer != "Melee" {
		t.Errorf("%+v", answer)
	}
}

func Test_GetMissileSpeed_PL(t *testing.T) {
	answer, err := ParseHeroMissileSpeed(pl_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if answer != "Instant" {
		t.Errorf("%+v", answer)
	}
}

func Test_GetMissileSpeed_DP(t *testing.T) {
	answer, err := ParseHeroMissileSpeed(dp_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if answer != "1000" {
		t.Errorf("%+v", answer)
	}
}

func Test_GetAttackDuration_PL(t *testing.T) {
	answer, err := ParseHeroAttackDuration(pl_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 2 || answer[0] != 0.5 || answer[1] != 0.5 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetAttackDuration_DP(t *testing.T) {
	answer, err := ParseHeroAttackDuration(dp_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 2 || answer[0] != 0.56 || answer[1] != 0.51 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetAttackDuration_Tiny(t *testing.T) {
	answer, err := ParseHeroAttackDuration(tiny_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 2 ||
		answer[0] != 0.49 ||
		answer[1] != 1 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetCastDuration_PL(t *testing.T) {
	answer, err := ParseHeroCastDuration(pl_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 2 || answer[0] != 0.3 || answer[1] != 0.51 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetCastDuration_DP(t *testing.T) {
	answer, err := ParseHeroCastDuration(dp_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 2 || answer[0] != 0.5 || answer[1] != 0.83 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetCastDuration_Tiny(t *testing.T) {
	answer, err := ParseHeroCastDuration(tiny_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 2 ||
		answer[0] != 0.001 ||
		answer[1] != 0 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetBaseAttackTime_PL(t *testing.T) {
	answer, err := ParseHeroBaseAttackTime(pl_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if answer != 1.7 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetBaseAttackTime_DP(t *testing.T) {
	answer, err := ParseHeroBaseAttackTime(dp_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if answer != 1.7 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetAttacksPerSecond_PL(t *testing.T) {
	answer, err := ParseHeroAttacksPerSecond(pl_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 3 ||
		answer[0] != 0.72 ||
		answer[1] != 1.1 ||
		answer[2] != 1.42 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetAttacksPerSecond_DP(t *testing.T) {
	answer, err := ParseHeroAttacksPerSecond(dp_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 3 ||
		answer[0] != 0.67 ||
		answer[1] != 0.8 ||
		answer[2] != 0.98 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetArmor_PL(t *testing.T) {
	answer, err := ParseHeroArmor(pl_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 3 ||
		answer[0] != 3.22 ||
		answer[1] != 12.32 ||
		answer[2] != 20.13 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetArmor_DP(t *testing.T) {
	answer, err := ParseHeroArmor(dp_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 3 ||
		answer[0] != 2.96 ||
		answer[1] != 6.18 ||
		answer[2] != 10.46 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetArmor_Io(t *testing.T) {
	answer, err := ParseHeroArmor(io_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 3 ||
		answer[0] != -0.04 ||
		answer[1] != 3.6 ||
		answer[2] != 8.14 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetArmor_Medusa(t *testing.T) {
	answer, err := ParseHeroArmor(medusa_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 3 ||
		answer[0] != 1.8 ||
		answer[1] != 7.33 ||
		answer[2] != 13 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetArmor_TW(t *testing.T) {
	answer, err := ParseHeroArmor(tw_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 3 ||
		answer[0] != 1.94 ||
		answer[1] != 8 ||
		answer[2] != 13.98 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetDamage_PL(t *testing.T) {
	answer, err := ParseHeroDamage(pl_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 3 ||
		answer[0] != "45-67" ||
		answer[1] != "110-132" ||
		answer[2] != "165-187" {
		t.Errorf("%+v", answer)
	}
}

func Test_GetDamage_DP(t *testing.T) {
	answer, err := ParseHeroDamage(dp_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 3 ||
		answer[0] != "44-56" ||
		answer[1] != "91-103" ||
		answer[2] != "136-148" {
		t.Errorf("%+v", answer)
	}
}

func Test_GetMana_PL(t *testing.T) {
	answer, err := ParseHeroMana(pl_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 3 ||
		answer[0] != 273 ||
		answer[1] != 689 ||
		answer[2] != 1157 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetMana_DP(t *testing.T) {
	answer, err := ParseHeroMana(dp_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 3 ||
		answer[0] != 260 ||
		answer[1] != 871 ||
		answer[2] != 1456 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetHP_PL(t *testing.T) {
	answer, err := ParseHeroHP(pl_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 3 ||
		answer[0] != 492 ||
		answer[1] != 1005 ||
		answer[2] != 1632 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetHP_DP(t *testing.T) {
	answer, err := ParseHeroHP(dp_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if len(answer) != 3 ||
		answer[0] != 511 ||
		answer[1] != 1176 ||
		answer[2] != 1879 {
		t.Errorf("%+v", answer)
	}
}

func Test_GetProfilePicture_PL(t *testing.T) {
	answer, err := ParseHeroProfilePicture("Phantom Lancer", pl_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if answer != "http://hydra-media.cursecdn.com/dota2.gamepedia.com/5/51/Phantom_Lancer.png" {
		t.Errorf("%+v", answer)
	}
}

func Test_GetProfilePicture_DP(t *testing.T) {
	answer, err := ParseHeroProfilePicture("Death Prophet", dp_data)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if answer != "http://hydra-media.cursecdn.com/dota2.gamepedia.com/5/52/Death_Prophet.png" {
		t.Errorf("%+v", answer)
	}
}
