package mythchar

import "errors"

// Skill All skills look the same
type Skill struct {
	Name    string
	Percent int
}

var standardSkillBasePercentages = map[string][]string{
	"Athletics":    {"str", "dex"},
	"Boating":      {"str", "con"},
	"Brawn":        {"str", "siz"},
	"Conceal":      {"dex", "pow"},
	"Customs":      {"int", "int"},
	"Dance":        {"dex", "cha"},
	"Deceit":       {"int", "cha"},
	"Drive":        {"dex", "pow"},
	"Endurance":    {"con", "con"},
	"Evade":        {"dex", "dex"},
	"FirstAid":     {"int", "dex"},
	"Influence":    {"cha", "cha"},
	"Insight":      {"int", "pow"},
	"Locale":       {"int", "int"},
	"NativeTongue": {"int", "cha"},
	"Perception":   {"int", "pow"},
	"Ride":         {"dex", "pow"},
	"Sing":         {"cha", "pow"},
	"Stealth":      {"dex", "int"},
	"Swim":         {"str", "con"},
	"Unarmed":      {"str", "dex"},
	"Willpower":    {"pow", "pow"},
}

var professionalSkillBasePercentages = map[string][]string{
	"Acting":       {"cha", "cha"},
	"Acrobatics":   {"cha", "cha"},
	"Art":          {"pow", "cha"},
	"Binding":      {"pow", "cha"},
	"Bureaucracy":  {"int", "int"},
	"Commerce":     {"int", "cha"},
	"Courtsey":     {"int", "cha"},
	"Craft":        {"dex", "int"},
	"Culture":      {"int", "int"},
	"Devotion":     {"pow", "cha"},
	"Disguise":     {"int", "cha"},
	"Engineering":  {"int", "int"},
	"Exhort":       {"int", "cha"},
	"Folk Magic":   {"pow", "cha"},
	"Gambling":     {"int", "pow"},
	"Healing":      {"int", "pow"},
	"Invocation":   {"int", "int"},
	"Language":     {"int", "cha"},
	"Literacy":     {"int", "int"},
	"Lockpicking":  {"dex", "dex"},
	"Lore":         {"int", "int"},
	"Mechanisms":   {"dex", "int"},
	"Meditation":   {"int", "con"},
	"Musicianship": {"dex", "cha"},
	"Mysticism":    {"pow", "cha"},
	"Navigation":   {"int", "pow"},
	"Oratory":      {"pow", "cha"},
	"Seamanship":   {"int", "con"},
	"Seduction":    {"int", "cha"},
	"Shaping":      {"int", "pow"},
	"Sleight":      {"dex", "cha"},
	"Streetwise":   {"pow", "cha"},
	"Survival":     {"con", "pow"},
	"Teach":        {"int", "cha"},
	"Track":        {"int", "con"},
	"Trance":       {"pow", "con"},
}

// GenerateStandardSkills creates the base standard skills for a character
func (m *MythChar) GenerateStandardSkills() {
	for skill, bases := range standardSkillBasePercentages {
		m.Skills = append(m.Skills, m.generateSkillFromBases(skill, bases))
	}
}

func (m *MythChar) generateSkillFromBases(name string, bases []string) Skill {
	newSkill := Skill{Name: name}
	for _, base := range bases {
		switch base {
		case "str":
			newSkill.Percent += m.Stats.Str.Current
		case "con":
			newSkill.Percent += m.Stats.Con.Current
		case "siz":
			newSkill.Percent += m.Stats.Siz.Current
		case "dex":
			newSkill.Percent += m.Stats.Dex.Current
		case "int":
			newSkill.Percent += m.Stats.Int.Current
		case "pow":
			newSkill.Percent += m.Stats.Pow.Current
		case "cha":
			newSkill.Percent += m.Stats.Cha.Current
		}
	}

	return newSkill
}

// CreateProfessionalSkill takes a professional skill name (caps matter) and adds the skill to the MythChar
func (m *MythChar) CreateProfessionalSkill(skill string) error {
	if base, ok := professionalSkillBasePercentages[skill]; ok {
		m.Skills = append(m.Skills, m.generateSkillFromBases(skill, base))
		return nil
	}
	return errors.New("Skill was not in the list of professional skills")
}
