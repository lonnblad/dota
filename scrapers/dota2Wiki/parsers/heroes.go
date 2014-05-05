package parsers

import (
	"errors"
	"log"
	"regexp"
	"strconv"
)

func parseSingleFloat(reg, data string) (float64, error) {
	re := regexp.MustCompile(reg)
	if matches := re.FindAllStringSubmatch(data, -1); matches != nil {
		res, err := strconv.ParseFloat(matches[0][1], 64)
		if err != nil {
			return 0.0, err
		}
		return res, nil
	}
	return 0.0, errors.New("This reg didn't work:\n" + reg)
}

func parseTripleFloat(reg, data string) ([]float64, error) {
	re := regexp.MustCompile(reg)
	if matches := re.FindAllStringSubmatch(data, -1); matches != nil {
		res1, err := strconv.ParseFloat(matches[0][1], 64)
		if err != nil {
			return []float64{}, err
		}
		res2, err := strconv.ParseFloat(matches[0][2], 64)
		if err != nil {
			return []float64{}, err
		}
		res3, err := strconv.ParseFloat(matches[0][3], 64)
		if err != nil {
			return []float64{}, err
		}
		return []float64{res1, res2, res3}, nil
	}
	return []float64{}, errors.New("This reg didn't work:\n" + reg)
}

func parseSingleString(reg, data string) (string, error) {
	re := regexp.MustCompile(reg)
	if matches := re.FindAllStringSubmatch(data, -1); matches != nil {
		return matches[0][1], nil
	}
	return "", errors.New("This reg didn't work:\n" + reg)
}

func parseTripleString(reg, data string) ([]string, error) {
	re := regexp.MustCompile(reg)
	if matches := re.FindAllStringSubmatch(data, -1); matches != nil {
		return []string{matches[0][1], matches[0][2], matches[0][3]}, nil
	}
	return []string{}, errors.New("This reg didn't work:\n" + reg)
}

func ParseHeroes(data string) ([]string, error) {
	re := regexp.MustCompile(`src="http://hydra-media.cursecdn.com/dota2.gamepedia.com/thumb/./../([_a-zA-Z]+).png`)
	if matches := re.FindAllStringSubmatch(data, -1); matches != nil {
		for _, elem := range matches {
			log.Printf("%+v\n", elem[1])
		}
	}
	return []string{}, errors.New("no heroes found")
}

func ParseHeroName(data string) (string, error) {
	re := regexp.MustCompile(`<th align="center"> ([a-zA-Z\'\- ]+)\n`)
	if matches := re.FindAllStringSubmatch(data, -1); matches != nil {
		return matches[0][1], nil
	}
	return "", errors.New("no hero name found")
}

func ParseHeroFaction(data string) (string, error) {
	radiant := regexp.MustCompile(`The Radiant`)
	if radiant.MatchString(data) {
		return "The Radiant", nil
	}
	radiant = regexp.MustCompile(`The Dire`)
	if radiant.MatchString(data) {
		return "The Dire", nil
	}
	return "", errors.New("no faction found")
}

func ParseHeroStrength(data string) ([]float64, error) {
	return parseHeroAttribute(data, "Str")
}

func ParseHeroAgility(data string) ([]float64, error) {
	return parseHeroAttribute(data, "Agi")
}

func ParseHeroIntelligence(data string) ([]float64, error) {
	return parseHeroAttribute(data, "Int")
}

func parseHeroAttribute(data, attr string) ([]float64, error) {

	re := regexp.MustCompile(attr + `([_a-z]*).png 2x" /></a><br /> <b>([0-9]+)</b> . ([0-9\.]+)\n`)
	if matches := re.FindAllStringSubmatch(data, -1); matches != nil {
		main, err := strconv.ParseFloat(matches[0][2], 64)
		if err != nil {
			return []float64{}, err
		}
		add, err := strconv.ParseFloat(matches[0][3], 64)
		if err != nil {
			return []float64{}, err
		}
		return []float64{main, add}, nil
	} else {
		return []float64{}, errors.New(attr + " not found")
	}
}

func ParseHeroPrimary(data string) (string, error) {
	re := regexp.MustCompile(`Icon_([a-zA-Z]+)_selected.png`)
	if matches := re.FindAllStringSubmatch(data, -1); matches != nil {
		return matches[0][1], nil
	}
	return "", errors.New("Primary not found")
}

func ParseHeroMovementSpeed(data string) (float64, error) {
	reg := `<b>Movement Speed</b> </td>\n<td> ([0-9]+)`
	return parseSingleFloat(reg, data)
}

func ParseHeroTurnRate(data string) (float64, error) {
	reg := `<b>Turn Rate</b> </td>\n<td> ([0-9]\.[0-9])`
	return parseSingleFloat(reg, data)
}

func ParseHeroSightRange(data string) ([]float64, error) {
	var day, night float64
	var err error

	regDay := `title="During the day">([0-9]+)</span>`
	day, err = parseSingleFloat(regDay, data)
	if err != nil {
		return []float64{}, err
	}

	regNight := `title="During the night">([0-9]+)</span>`
	night, err = parseSingleFloat(regNight, data)
	if err != nil {
		return []float64{}, err
	}

	return []float64{day, night}, nil
}

func ParseHeroAttackRange(data string) (string, error) {
	reg := `<b>Attack Range</b> </td>\n<td> ([0-9]+)`
	if ar, err := parseSingleString(reg, data); err == nil {
		return ar, nil
	} else {
		return "Melee", nil
	}
}

func ParseHeroMissileSpeed(data string) (string, error) {
	reg := `<b>Missile Speed</b> </td>\n<td>  ([0-9]+)`
	if ar, err := parseSingleString(reg, data); err == nil {
		return ar, nil
	} else {
		return "Instant", nil
	}
}

func ParseHeroAttackDuration(data string) ([]float64, error) {
	var point, backswing float64
	var err error

	regPoint := `title="Attack point">([0-9]+\.{0,1}[0-9]*)</span>`
	point, err = parseSingleFloat(regPoint, data)
	if err != nil {
		return []float64{}, err
	}

	regBackswing := `title="Attack backswing">([0-9]+\.{0,1}[0-9]*)</span>`
	backswing, err = parseSingleFloat(regBackswing, data)
	if err != nil {
		return []float64{}, err
	}

	return []float64{point, backswing}, nil
}

func ParseHeroCastDuration(data string) ([]float64, error) {
	var point, backswing float64
	var err error

	regPoint := `title="Cast point">([0-9]+\.{0,1}[0-9]*)</span>`
	point, err = parseSingleFloat(regPoint, data)
	if err != nil {
		return []float64{}, err
	}

	regBackswing := `title="Cast backswing">([0-9]+\.{0,1}[0-9]*)</span>`
	backswing, err = parseSingleFloat(regBackswing, data)
	if err != nil {
		return []float64{}, err
	}

	return []float64{point, backswing}, nil
}

func ParseHeroBaseAttackTime(data string) (float64, error) {
	reg := `<b>Base Attack Time</b> </td>\n<td>([0-9]\.[0-9])`
	return parseSingleFloat(reg, data)
}

func ParseHeroAttacksPerSecond(data string) ([]float64, error) {
	attr := "Attacks / Second"
	reg := `<b>` + attr + `</b>\s*</td>\n<td>\s*(\-{0,1}[0-9]+\.{0,1}[0-9]*)\s*</td>\n<td>\s*([0-9]+\.{0,1}[0-9]*)\s*</td>\n<td>\s*([0-9]+\.{0,1}[0-9]*)`
	return parseTripleFloat(reg, data)
}

func ParseHeroArmor(data string) ([]float64, error) {
	attr := "Armor"
	reg := `<b>` + attr + `</b>\s*</td>\n<td>\s*(\-{0,1}[0-9]+\.{0,1}[0-9]*)\s*</td>\n<td>\s*([0-9]+\.{0,1}[0-9]*)\s*</td>\n<td>\s*([0-9]+\.{0,1}[0-9]*)`
	return parseTripleFloat(reg, data)
}

func ParseHeroDamage(data string) ([]string, error) {
	reg := `<b>Damage</b> </td>\n<td> ([0-9]+\-[0-9]+) </td>\n<td> ([0-9]+\-[0-9]+) </td>\n<td> ([0-9]+\-[0-9]+)`
	return parseTripleString(reg, data)
}

func ParseHeroMana(data string) ([]float64, error) {
	reg := `<b>Mana</b> </td>\n<td> ([0-9]+) </td>\n<td> ([0-9]+) </td>\n<td> ([0-9]+)`
	return parseTripleFloat(reg, data)
}

func ParseHeroHP(data string) ([]float64, error) {
	reg := `<b>Hit Points</b> </td>\n<td> ([0-9]+) </td>\n<td> ([0-9]+) </td>\n<td> ([0-9]+)`
	return parseTripleFloat(reg, data)
}

func ParseHeroProfilePicture(hero, data string) (string, error) {
	reg := `<img alt="` + hero + `.png" src="([a-zA-Z0-9\:\/\-\.\_\%]+)?`
	return parseSingleString(reg, data)
}
