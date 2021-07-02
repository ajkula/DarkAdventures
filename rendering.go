package main

import (
	"strings"
	"time"
	"unicode/utf8"
)

var started bool = true
var navTutoBool bool = true
var fightTutoBool bool = true
var rootTutoBool bool = true
var oneTimeQuestEvents = &OneTimeQuestEvents{str: "", show: false}

type OneTimeQuestEvents struct {
	str  string
	show bool
}

func PresentScene(p *Character) {
	loc := p.SetPlayerRoom()
	var t int8 = 1
	loc.RemoveBattle()
	cleanResolvedQuests()
	if loc.HasEnemy {
		t = getTurns()
	}
	if started {
		loc.showImage()
		if t == 1 {
			Output("yellow", loc.Description)
			if (loc.HasEnemy && indexOf(giants, loc.Enemy.Name) == -1) || !loc.HasEnemy {
				Output("yellow", loc.Ephemeral)
				if loc.HasNPC {
					if loc.NPC.Quest.Rewarded == true {
						loc.HasNPC = false
					}
					oneTimeQuestEvents.addEventString(dialogFromConditions(loc.NPC.Quest))
				}
			}
		}
		oneTimeQuestEvents.display()
		if loc.HasGate {
			if t == 1 {
				Output("yellow", rootBell[p.hasItemInInventory(itemNames.Key)])
				showRootTuto(p.hasItemInInventory(itemNames.Key))
			}
		}
		if navTutoBool {
			Output("red", translate(navTutoTR)+Initial(loc.CanGoTo[0])+" ->"+translate(forTR)+loc.CanGoTo[0]+"\n")
			navTutoBool = false
		}

		if loc.HasEnemy {
			if indexGiants := indexOf(giants, loc.Enemy.Name); indexGiants != -1 {
				sayIt := loc.Enemy.Encounter.shouldSayIt()
				if sayIt {
					Output("yellow", loc.Ephemeral)
					loc.Enemy.Encounter.saidIt()
					time.Sleep(1 * time.Second)
				}
			}
		}
		pile.forEachCharacter(applyStatus)
		p.showHealth()
		p.DisplayExpGauge()
	}

	// That's awful, but I am looking for another way
	switch true {
	case !loc.HasEnemy && !loc.HasSeller && !loc.HasNPC:
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
		break
	case loc.HasEnemy:
		if loc.Enemy.Alive {
			showActions(p, battleCommands)
			showUniversalCmds()
			Battle(p, loc.Enemy)
		}
		break
	case loc.HasSeller:
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
		break
	case loc.HasNPC:
		if started {
			showActions(p, npcCommands)
			showWhereCanGo(loc)
			showUniversalCmds()
		}
		cmd := UserInputln()
		if ok := arrayIncludesCommand(npcCommands, cmd); ok {
			ProcessCommands(p, cmd)
		} else {
			Output("red", translate(cantDoThatTR))
		}
		break
	default:
		break
	}

	started = true
}

func (otqe *OneTimeQuestEvents) addEventString(str string) {
	otqe.str = str
	otqe.show = true
}

func (otqe *OneTimeQuestEvents) display() {
	// if otqe.show {
	// loc := hero.SetPlayerRoom()
	// if loc.HasNPC {
	Output("blue", otqe.str)
	// }
	// }
	otqe.str = ""
	otqe.show = false
}

func applyStatus(player *Character) { player.applyStatusesEffect() }
func showRootTuto(b bool) {
	if b && rootTutoBool {
		Output("red", translate(rootTutoTR))
		rootTutoBool = false
	}
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
