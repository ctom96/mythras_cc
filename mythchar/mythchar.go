package mythchar

import (
	"strconv"
	"strings"
)

// MythChar is the whole Character Type
type MythChar struct {
	Name   string
	Stats  Characteristics
	Attrs  Attributes
	Skills []Skill
}

// CreateHuman creates a basic human character
func CreateHuman(name string) MythChar {
	newChar := MythChar{
		Name: name,
	}
	newChar.Stats = RollHuman()
	newChar.Attrs = GenerateAttributes(newChar.Stats)
	newChar.GenerateStandardSkills()

	return newChar
}

// ----------------------------------------------------------------------------
// Util Methods
// ----------------------------------------------------------------------------

func (m MythChar) statsString() string {
	var ret strings.Builder

	ret.WriteString("Characteristics:\n")
	ret.WriteString("Str: " + strconv.Itoa(m.Stats.Str.Current) + "\n")
	ret.WriteString("Con: " + strconv.Itoa(m.Stats.Con.Current) + "\n")
	ret.WriteString("Siz: " + strconv.Itoa(m.Stats.Siz.Current) + "\n")
	ret.WriteString("Dex: " + strconv.Itoa(m.Stats.Dex.Current) + "\n")
	ret.WriteString("Int: " + strconv.Itoa(m.Stats.Int.Current) + "\n")
	ret.WriteString("Pow: " + strconv.Itoa(m.Stats.Pow.Current) + "\n")
	ret.WriteString("Cha: " + strconv.Itoa(m.Stats.Cha.Current) + "\n")

	return ret.String()
}

// PrintAttrs prints the attrs of a character
func (m MythChar) attrsString() string {
	var ret strings.Builder

	ret.WriteString("Attributes:\n")
	ret.WriteString("ActionPoints: " + strconv.Itoa(m.Attrs.ActionPoints) + "\n")
	ret.WriteString("DamageModifier: " + m.Attrs.DamageModifier + "\n")
	ret.WriteString("ExperienceModifier: " + strconv.Itoa(m.Attrs.ExperienceModifier) + "\n")
	ret.WriteString("HealingRate: " + strconv.Itoa(m.Attrs.HealingRate) + "\n")
	ret.WriteString("Initiative: " + strconv.Itoa(m.Attrs.Initiative) + "\n")
	ret.WriteString("LuckPoints: " + strconv.Itoa(m.Attrs.LuckPoints) + "\n")
	ret.WriteString("MovementRate: " + strconv.Itoa(m.Attrs.MovementRate) + "\n")

	return ret.String()
}

func (m MythChar) skillsString() string {
	var ret strings.Builder

	ret.WriteString("Skills:\n")
	for _, skill := range m.Skills {
		ret.WriteString(skill.Name + ": " + strconv.Itoa(skill.Percent) + "%\n")
	}

	return ret.String()
}

func (m MythChar) String() string {
	var ret strings.Builder
	ret.WriteString("------------------------------\n")
	ret.WriteString("Name: " + m.Name + "\n\n")
	ret.WriteString(m.statsString())
	ret.WriteString("\n")
	ret.WriteString(m.attrsString())
	ret.WriteString("\n")
	ret.WriteString(m.skillsString())
	ret.WriteString("------------------------------\n")

	return ret.String()

}
