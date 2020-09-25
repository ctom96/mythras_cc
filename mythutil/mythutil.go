package mythutil

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Roll take a string of the form "NdX{+-}M", and outputs the result
// N, X are required in the string
func Roll(dieString string) (int, error) {
	dieStringParts := strings.Split(dieString, "d")

	multiple, err := strconv.Atoi(dieStringParts[0])
	if err != nil {
		return 0, errors.New("String format incorrect. Missing 'd'")
	}

	dieType, err := strconv.Atoi(string([]rune(dieStringParts[1])[0]))
	if err != nil {
		return 0, errors.New("String format incorrect. Missing 'd'")
	}

	additive := 0
	if strings.Contains(dieString, "+") {
		dieStringParts := strings.Split(dieString, "+")
		additive, err = strconv.Atoi(string([]rune(dieStringParts[1])[0]))
		if err != nil {
			return 0, errors.New("String format incorrect. Number after + unreadable")
		}
	} else if strings.Contains(dieString, "-") {
		dieStringParts := strings.Split(dieString, "-")
		additive, err = strconv.Atoi(string([]rune(dieStringParts[1])[0]))
		if err != nil {
			return 0, errors.New("String format incorrect. Number after - unreadable")
		}
		additive = additive * -1
	}

	sum := 0
	for i := 0; i < multiple; i++ {
		roll := rollDie(dieType)
		sum += roll
	}
	sum += additive

	return sum, nil
}

func rollDie(sides int) int {
	return rand.Intn(sides) + 1
}
