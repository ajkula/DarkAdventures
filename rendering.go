package main

import (
	"strings"
	"time"
	"unicode/utf8"
)

var started bool = false

func applyStatus(player *Character) { player.applyStatusesEffect() }

func PresentScene(p *Character) {
	loc := p.SetPlayerRoom()
	loc.RemoveBattle()
	// ICI
	t := getTurns()
	if started {
		loc.showImage()
		if t == 1 {
			Output("yellow", loc.Description)
		}
		if loc.HasGate {
			if t == 1 {
				Output("yellow", rootBell[p.hasItemInInventory(itemNames.Key)])
			}
		}
		sayIt := dragonLanding.shouldSayIt()
		if loc.HasEnemy && loc.Enemy.Name == enemiesList.DRAGON && sayIt {
			Output("yellow", loc.Ephemeral)
			dragonLanding.saidIt()
			time.Sleep(1 * time.Second)
		}
		// Output("red", loc.HasEnemy)
		// Output("green", getTurns())
		// Output("red", "dragon.Freeze "+strconv.FormatBool(dragon.Freeze))
		// ICI
		pile.forEachCharacter(applyStatus)
		// if t == 1 {
		p.showHealth()
		p.DisplayExpGauge()
		// }
	}
	// if loc.HasGate {
	// 	y, x := loc.Gate.Warp()
	// 	Output("red", " y: ", y, " x: ", x)
	// }

	if loc.HasEnemy {
		if loc.Enemy.Alive {
			// if t == 1 {
			showActions(p, battleCommands)
			showUniversalCmds()
			// }
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
			Output("red", translate(cantDoThatTR))
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
			Output("red", translate(cantDoThatTR))
		}
	}
	started = true
}

func showWhereCanGo(loc *Location) {
	Output("blue", Tab+CalculateSpaceAlign(translate(youCanGoTR)), WhereCanYouGo(loc))
}

func showActions(p *Character, arr []string) {
	Output("white", Tab+CalculateSpaceAlign(translate(youCanTR)), arr)
	Output("white", Tab+CalculateSpaceAlign(translate(youCanUseTR)), p.DisplayItems())
}

func showUniversalCmds() {
	Output("white", Tab+CalculateSpaceAlign(translate(orTR))+ArrayToString(universalCommands), "\n")
}

func CalculateSpaceAlign(str string) string {
	var length int = 1
	if Tabulation-utf8.RuneCountInString(str) > 1 {
		length = Tabulation - utf8.RuneCountInString(str)
	}
	spaces := strings.Repeat(" ", length)
	return str + spaces
}

func CustomSpaceAlign(str string, i int) string {
	var length int = 1
	if i-utf8.RuneCountInString(str) > 1 {
		length = i - utf8.RuneCountInString(str)
	}
	spaces := strings.Repeat(" ", length)
	return str + spaces
}

func Article(str string) string {
	if Lang == frenchLang {
		return str + " "
	}
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
