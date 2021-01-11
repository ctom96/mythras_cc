package mythchar

// Skill All skills look the same
type Skill struct {
	Name    string
	Percent int
}

var standardSkillBasePercentages = map[string][]string{
	"Athletics":    {"Str", "Dex"},
	"Boating":      {"Str", "Con"},
	"Brawn":        {"Str", "Siz"},
	"Conceal":      {"Dex", "Pow"},
	"Customs":      {"Int", "Int"},
	"Dance":        {"Dex", "Cha"},
	"Deceit":       {"Int", "Cha"},
	"Drive":        {"Dex", "Pow"},
	"Endurance":    {"Con", "Con"},
	"Evade":        {"Dex", "Dex"},
	"FirstAid":     {"Int", "Dex"},
	"Influence":    {"Cha", "Cha"},
	"Insight":      {"Int", "Pow"},
	"Locale":       {"Int", "Int"},
	"NativeTongue": {"Int", "Cha"},
	"Perception":   {"Int", "Pow"},
	"Ride":         {"Dex", "Pow"},
	"Sing":         {"Cha", "Pow"},
	"Stealth":      {"Dex", "Int"},
	"Swim":         {"Str", "Con"},
	"Unarmed":      {"Str", "Dex"},
	"Willpower":    {"Pow", "Pow"},
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
		case "Str":
			newSkill.Percent += m.Stats.Str.Current
		case "Con":
			newSkill.Percent += m.Stats.Con.Current
		case "Siz":
			newSkill.Percent += m.Stats.Siz.Current
		case "Dex":
			newSkill.Percent += m.Stats.Dex.Current
		case "Int":
			newSkill.Percent += m.Stats.Int.Current
		case "Pow":
			newSkill.Percent += m.Stats.Pow.Current
		case "Cha":
			newSkill.Percent += m.Stats.Cha.Current
		}
	}

	return newSkill
}
