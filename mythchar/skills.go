package mythchar

// All skills look the same
type Skill struct {
	Name  string
	Base1 Stat
	Base2 Stat
}

// Skills for a specific character
type Skills struct {
	Athletics Skill,
	Boating Skill,
	Brawn Skill,
	Conceal Skill,
	Customs Skill,
	Dance Skill,
	Deceit Skill,
	Drive Skill,
	Endurance Skill,
	Evade Skill,
	FirstAid Skill,
	Influence Skill,
	Insight Skill,
	Locale Skill,
	NativeTongue Skill,
}
