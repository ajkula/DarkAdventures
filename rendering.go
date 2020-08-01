package main

import (
	"strconv"
	"strings"
)

var started bool = false

func PresentScene(p *Character) {
	loc := p.SetPlayerRoom()
	loc.RemoveBattle()
	// ICI
	if started {
		loc.showImage()
		Output("yellow", loc.Description)
		Output("yellow", loc.Ephemeral)
		Output("red", loc.HasEnemy)
		Output("green", getTurns())
		Output("red", "dragon.Freeze "+strconv.FormatBool(dragon.Freeze))
		p.showHealth()
	}

	// if loc.HasGate {
	// 	y, x := loc.Gate.Warp()
	// 	Output("red", " y: ", y, " x: ", x)
	// }

	if loc.HasEnemy {
		if loc.Enemy.Alive {
			showActions(p, battleCommands)
			showUniversalCmds()
			Battle(p, loc.Enemy)
		}
	}

	if !loc.HasEnemy && !loc.HasSeller {
		if started {
			showActions(p, worldCommands)
			showWhereCanGo(loc)
			showUniversalCmds()
		}
		cmd := UserInputln()
		if ok := arrayIncludesCommand(worldCommands, cmd); ok {
			ProcessCommands(p, cmd)
		} else {
			Output("red", Tab+"You can't do that here...")
		}
	}

	if loc.HasSeller {
		if started {
			showActions(p, sellerCommands)
			showWhereCanGo(loc)
			showUniversalCmds()
		}
		cmd := UserInputln()
		if ok := arrayIncludesCommand(sellerCommands, cmd); ok {
			ProcessCommands(p, cmd)
		} else {
			Output("red", Tab+"You can't do that here...")
		}
	}
	started = true
}

func showWhereCanGo(loc *Location) {
	Output("blue", Tab+CalculateSpaceAlign("You can Go:"), WhereCanYouGo(loc))
}

func showActions(p *Character, arr []string) {
	Output("white", Tab+CalculateSpaceAlign("You can: "), arr)
	Output("white", Tab+CalculateSpaceAlign("You can use: "), p.DisplayItems())
}

func showUniversalCmds() {
	Output("white", Tab+CalculateSpaceAlign("...or: ")+ArrayToString(universalCommands), "\n")
}

func CalculateSpaceAlign(str string) string {
	var length int = 1
	if Tabulation-len(str) > 1 {
		length = Tabulation - len(str)
	}
	spaces := strings.Repeat(" ", length)
	return str + spaces
}

func CustomSpaceAlign(str string, i int) string {
	var length int = 1
	if i-len(str) > 1 {
		length = i - len(str)
	}
	spaces := strings.Repeat(" ", length)
	return str + spaces
}

func Article(str string) string {
	vowels := []string{"a", "e", "i", "o"}
	if InitialsIndexOf(vowels, str) {
		return "an " + str + " "
	}
	return "a " + str + " "
}

func extractCommand(input string) string {
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		return ""
	}
	return strings.ToLower(tokens[0])
}

func arrayIncludesCommand(comArr []string, input string) bool {
	if str := extractCommand(input); str != "" {
		test := append(comArr[:], universalCommands[:]...)
		if ok := InitialsIndexOf(test, str); ok {
			return true
		}
	}
	return false
}
