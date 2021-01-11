package main

import (
	"fmt"

	"github.com/ctom96/mythras_cc/mythchar"
	"github.com/ctom96/mythras_cc/mythserv"
)

func main() {
	mythserv.StartServer()

	player := mythchar.CreateHuman("Duncan Mac Duffy")

	fmt.Println(player)
}
