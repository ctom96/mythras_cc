package mythchar

import (
	"fmt"
	"strconv"
)

// MythChar is the whole Character Type
type MythChar struct {
	Stats Characteristics
	Attrs Attributes
}

// PrintStats prints the stats of a character
func (m MythChar) PrintStats() {
	fmt.Println("Str: " + strconv.Itoa(m.Stats.Str.Current))
	fmt.Println("Con: " + strconv.Itoa(m.Stats.Con.Current))
	fmt.Println("Siz: " + strconv.Itoa(m.Stats.Siz.Current))
	fmt.Println("Dex: " + strconv.Itoa(m.Stats.Dex.Current))
	fmt.Println("Int: " + strconv.Itoa(m.Stats.Int.Current))
	fmt.Println("Pow: " + strconv.Itoa(m.Stats.Pow.Current))
	fmt.Println("Cha: " + strconv.Itoa(m.Stats.Cha.Current))
}

// PrintAttrs prints the attrs of a character
func (m MythChar) PrintAttrs() {
	fmt.Println("ActionPoints: " + strconv.Itoa(m.Attrs.ActionPoints))
	fmt.Println("DamageModifier: " + m.Attrs.DamageModifier)
	fmt.Println("ExperienceModifier: " + strconv.Itoa(m.Attrs.ExperienceModifier))
	fmt.Println("HealingRate: " + strconv.Itoa(m.Attrs.HealingRate))
	fmt.Println("Initiative: " + strconv.Itoa(m.Attrs.Initiative))
	fmt.Println("LuckPoints: " + strconv.Itoa(m.Attrs.LuckPoints))
	fmt.Println("MovementRate: " + strconv.Itoa(m.Attrs.MovementRate))
}
