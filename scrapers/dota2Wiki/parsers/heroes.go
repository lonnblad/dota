package parsers

import (
	"errors"
	"log"
	"regexp"
	"strconv"
)

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
	re := regexp.MustCompile(`<th align="center"> ([a-zA-Z ]+)\n`)
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
