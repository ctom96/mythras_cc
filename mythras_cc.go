package main

import (
	"github.com/ctom96/mythras_cc/mythchar"
	"github.com/ctom96/mythras_cc/mythserv"
)

func main() {
	mythserv.StartServer()

	playerChars := mythchar.RollHuman()
	playerAttrs := mythchar.GenerateAttributes(playerChars)

	player := mythchar.MythChar{Stats: playerChars, Attrs: playerAttrs}

	player.PrintStats()
	player.PrintAttrs()
}
