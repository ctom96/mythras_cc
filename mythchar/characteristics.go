package mythchar

import (
	"math"

	"github.com/ctom96/mythras_cc/mythutil"
)

// Stat is for a specific stat
type Stat struct {
	Original int
	Max      int
	Current  int
}

// Characteristics is a set of stats in a character
type Characteristics struct {
	Str Stat
	Con Stat
	Siz Stat
	Dex Stat
	Int Stat
	Pow Stat
	Cha Stat
}

type Attributes struct {
	ActionPoints       int
	DamageModifier     string
	ExperienceModifier int
	HealingRate        int
	Initiative         int
	LuckPoints         int
	MovementRate       int
}

// ----------
// Useful Generator Methods
// ----------

// CreateStat takes a die string and returns a stat
func CreateStat(dieString string) Stat {
	roll, err := mythutil.Roll(dieString)
	if err != nil {
		return Stat{Original: -1, Max: -1, Current: -1}
	}

	return Stat{Original: roll, Max: roll, Current: roll}
}

// RollHuman randomly creates characteristics following the rolling rules
// of rolling up a human (SIZ, INT are 2d6+6)
func RollHuman() Characteristics {
	return Characteristics{
		Str: CreateStat("3d6"),
		Con: CreateStat("3d6"),
		Siz: CreateStat("2d6+6"),
		Dex: CreateStat("3d6"),
		Int: CreateStat("2d6+6"),
		Pow: CreateStat("3d6"),
		Cha: CreateStat("3d6"),
	}
}

// GenerateAttributes creates a set of attributes based on the
// given characteristics
func GenerateAttributes(chars Characteristics) Attributes {
	retAttributes := Attributes{}
	retAttributes.ActionPoints = determineActionPoints(chars.Int, chars.Dex)
	retAttributes.DamageModifier = determineDamageModifier(chars.Str, chars.Siz)
	retAttributes.ExperienceModifier = determineExperienceModifier(chars.Cha)
	retAttributes.HealingRate = determineHealingRate(chars.Con)
	retAttributes.Initiative = determineInitiative(chars.Dex, chars.Int)
	retAttributes.LuckPoints = determineLuckPoints(chars.Pow)
	retAttributes.MovementRate = 6

	return retAttributes
}

func determineLuckPoints(pow Stat) int {
	return int(math.Ceil(float64(pow.Current) / 6.0))
}

func determineInitiative(dex Stat, intel Stat) int {
	return int(math.Floor(float64(dex.Current+intel.Current) / 2.0))
}

func determineHealingRate(con Stat) int {
	return int(math.Ceil(float64(con.Current) / 6.0))
}

func determineExperienceModifier(cha Stat) int {
	return int(float64(cha.Current)/6.0) - 1
}

func determineDamageModifier(str Stat, siz Stat) string {
	total := str.Current + siz.Current
	switch {
	case total <= 5:
		return "-1d8"
	case total <= 10:
		return "-1d6"
	case total <= 15:
		return "-1d4"
	case total <= 20:
		return "-1d2"
	case total <= 25:
		return "+0"
	case total <= 30:
		return "+1d2"
	case total <= 35:
		return "+1d4"
	case total <= 40:
		return "+1d6"
	case total <= 45:
		return "+1d8"
	case total <= 50:
		return "+1d10"
	case total <= 60:
		return "+1d12"
	case total <= 70:
		return "+2d6"
	case total <= 80:
		return "+1d8+1d6"
	case total <= 90:
		return "+2d8"
	case total <= 100:
		return "+1d10+1d8"
	case total <= 110:
		return "+2d10"
	case total <= 120:
		return "+2d10+1d2"
	case total <= 130:
		return "+2d10+1d4"
	default:
		return "0"
	}
}

func determineActionPoints(intel Stat, dex Stat) int {
	return int(math.Ceil(float64(intel.Current+dex.Current) / 12.0))
}
