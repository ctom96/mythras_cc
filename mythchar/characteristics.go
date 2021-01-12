package mythchar

import (
	"fmt"
	"math"

	"github.com/ctom96/mythras_cc/mythutil"
)

// Characteristic is a base stat, like Strength
type Characteristic struct {
	Name     string
	DieRoll  string
	Original int
	Max      int
	Current  int
}

func initCharacteristic(name string, dieRoll string) Characteristic {
	c := Characteristic{
		Name:    name,
		DieRoll: dieRoll,
	}

	roll, err := mythutil.Roll(c.DieRoll)
	if err != nil {
		c.Original = 0
		c.Max = 0
		c.Current = 0
		fmt.Println("Error:", err)
	}
	c.Original = roll
	c.Max = roll
	c.Current = roll

	return c
}

// Characteristics is a set of stats in a character
type Characteristics struct {
	Str Characteristic
	Con Characteristic
	Siz Characteristic
	Dex Characteristic
	Int Characteristic
	Pow Characteristic
	Cha Characteristic
}

// Attributes are the derived traits of a character that are not skills, like Healing Rate
type Attributes struct {
	ActionPoints       int
	DamageModifier     string
	ExperienceModifier int
	HealingRate        int
	Initiative         int
	LuckPoints         int
	MovementRate       int
}

// RollHuman randomly creates characteristics following the rolling rules
// of rolling up a human (SIZ, INT are 2d6+6)
func RollHuman() Characteristics {
	return Characteristics{
		Str: initCharacteristic("str", "3d6"),
		Con: initCharacteristic("con", "3d6"),
		Siz: initCharacteristic("siz", "2d6+6"),
		Dex: initCharacteristic("dex", "3d6"),
		Int: initCharacteristic("int", "2d6+6"),
		Pow: initCharacteristic("pow", "3d6"),
		Cha: initCharacteristic("cha", "3d6"),
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

func determineLuckPoints(pow Characteristic) int {
	return int(math.Ceil(float64(pow.Current) / 6.0))
}

func determineInitiative(dex Characteristic, intel Characteristic) int {
	return int(math.Floor(float64(dex.Current+intel.Current) / 2.0))
}

func determineHealingRate(con Characteristic) int {
	return int(math.Ceil(float64(con.Current) / 6.0))
}

func determineExperienceModifier(cha Characteristic) int {
	return int(float64(cha.Current)/6.0) - 1
}

func determineDamageModifier(str Characteristic, siz Characteristic) string {
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

func determineActionPoints(intel Characteristic, dex Characteristic) int {
	return int(math.Ceil(float64(intel.Current+dex.Current) / 12.0))
}
