package parsers

import (
	"errors"
	"regexp"
	"strconv"
)

func ParseItemName(data string) (string, error) {
	re := regexp.MustCompile(`<meta name="og:title" content="([a-zA-Z ]+)"`)
	if matches := re.FindAllStringSubmatch(data, -1); matches != nil {
		return matches[0][1], nil
	}
	return "", errors.New("no item name found")
}

func ParseItemCost(data string) (int, error) {
	re := regexp.MustCompile(`Cost<br />([0-9]+)`)
	if matches := re.FindAllStringSubmatch(data, -1); matches != nil {
		cost, err := strconv.Atoi(matches[0][1])
		if err != nil {
			return 0, err
		}
		return cost, nil
	}
	return 0, errors.New("no item cost found")
}

func ParseItemDisassemble(data string) (bool, error) {
	no := regexp.MustCompile(`<td colspan="3"> No`)
	if no.MatchString(data) {
		return false, nil
	}
	yes := regexp.MustCompile(`<td colspan="3"> Yes`)
	if yes.MatchString(data) {
		return true, nil
	}
	return false, errors.New("no item disassemble found")
}
