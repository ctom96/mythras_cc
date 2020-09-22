package mythchar

// "Enum" for gender
type Gender int
const (
	Male Gender = iota
	Female
)
func (g Gender) String() string {
	return [...]string{"Male", "Female"}[g]
}

// Culture "enum"
type Culture int
const (
	Barbarian Culture = iota
	Civilized
	Nomadic
	Primitive
)
func (c Culture) String() string {
	return [...]string{
		"Barbarian", 
		"Civilized", 
		"Nomadic", 
		"Primitive",
	}[c]
}

// Social Class "Enum"
type SocialClass int
const (
	Outcast SocialClass = iota
	Slave 
	Freeman
	Gentry
	Aristocracy
	Ruling
)
func (s SocialClass) String() string {
	return [...]string{
		"Outcast",
		"Slave",
		"Freeman",
		"Gentry",
		"Aristocracy",
		"Ruling",
	}[s]
}

// Characteristics 
type Stat struct {
	Original int
	Max int
	Current int
}

type Attribute stuct {
	Original int
	Current int
}

type Skill struct {
	
}

// ---------------
// MythChar is the whole Character Type
// ---------------
type MythChar struct {
	// General Character statistics
	Player MythPlayer
	Name string
	Species string
	Frame string
	Height int
	Weight int
	Gender Gender
	Culture Culture
	SocialClass SocialClass
	Career string
	
	// Stats
	Str Stat
	Con Stat
	Siz Stat
	Dex Stat
	Int Stat
	Pow Stat
	Cha Stat

	// Attributes
	ActionPoints Attribute
	DamageModifier Attribute
	ExperienceModifier Attribute
	HealingRate Attribute
	Initiative Attribute
	LuckPoints Attribute
	MovementRate Attribute

	// Skills

}
