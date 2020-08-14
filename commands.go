package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func ProcessCommands(player *Character, input string, args ...interface{}) {
	Output("yellow", "======================================================================")
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		Output("red", translate(noComandsReceivedTR))
		ResetTurns()
		return
	}

	command := strings.ToLower(tokens[0])
	itemName := ""
	if len(tokens) > 1 {
		itemName = tokens[1]
	}
	// Output("yellow", "tokens ", tokens)
	// Output("yellow", "command ", command)
	// Output("yellow", "itemName ", itemName)
	loc := player.SetPlayerRoom()
	dragon.shouldFreeze(Initial(command))
	switch Initial(command) {
	case Initial(commands.Go):
		ResetTurns()
		if itemName != "" {
			if InitialsIndexOf(loc.CanGoTo, itemName) && (!loc.HasEnemy || (loc.HasEnemy && !loc.Enemy.isAlive())) {
				player.MoveTo(itemName)
			} else {
				Output("red", translate(CanTGoTR)+vowelOrNot(itemName, directionsArticlesVowelsTR, directionsArticlesConsonantTR)+itemName+translate(fromHereTR))
			}
		} else {
			Output("red", translate(whatDoYouMeanTR), command, " ", itemName)
		}

	case Initial(commands.Attack):
		enemy := loc.Enemy
		if enemy != nil {
			player.attack(enemy)
		}

	case Initial(commands.Use): // to do
		if itemName == itemNames.Scroll && player.hasItemInInventory(itemName) {
			enemy := loc.Enemy
			// Output("yellow", "has "+itemName+": ", player.hasItemInInventory(itemName), " quantity: ", player.Inventory[itemName].Quantity)
			ok := player.useItem(itemName, enemy)
			if ok != true {
				Output("red", translate(youCanTUseATR)+itemName)
				ResetTurns()
			}
		} else {
			ok := player.useItem(itemName)
			if ok != true {
				ResetTurns()
			}
		}
	case Initial(commands.Escape):
		enemy := loc.Enemy
		fmt.Println("calcul evade chances: ", int(GetAPercentageOfB(minusOrSquash(enemy.Strength, player.Evasion), minusOrSquash(player.Strength, enemy.Evasion)))+player.Evasion+rand.Intn(player.Evasion))
		canEvade := PercentChances(int(GetAPercentageOfB(minusOrSquash(enemy.Strength, player.Evasion), minusOrSquash(player.Strength, enemy.Evasion))) + player.Evasion + rand.Intn(player.Evasion))
		player.escapeBattle(canEvade)
		// ResetTurns()
	case Initial(commands.Map):
		ResetTurns()
		DisplayWorldMap(player)
	case Initial(commands.Inv):
		ResetTurns()
		player.DisplayInvetory()
	case Initial(commands.Buy):
		if ok := player.BuyFromShop(itemName); !ok {
			Output("red", translate(cantBuyTR)+Article(itemName))
		}
	// case "get":
	// 	err, index, itm := FindItemByName(itemName)
	// 	//Make sure we do not pick it up twice
	// 	if err == nil && itm.ItemInRoom(loc) && !itm.ItemOnPlayer(player) {
	// 		player.Items = append(player.Items, index) // Add Item to Player's bag
	// 		itm.RemoveItemFromRoom(loc)
	// 	} else {
	// 		Output("Could not get " + itemName)
	// 	}
	// case "open":
	// 	OpenItem(player, itemName)
	// case "inv":
	// 	Output("yellow", "Your Inventory: ")
	// 	for _, itm := range player.Items {
	// 		Output("yellow", "\t"+Items[itm].Name)
	// 	}
	case Initial(commands.Help):
		ResetTurns()
		Output("blue", DoubleTab+"Commands:")
		Output("blue", Tab+CustomSpaceAlign("[g]o <Location Name>", 25)+"- Go to the new location")
		Output("blue", Tab+CustomSpaceAlign("[a]ttack", 25)+"- Attacks opponent(s)")
		Output("blue", Tab+CustomSpaceAlign("[b]uy <Item Name>", 25)+"- Buy ONE OF proposed items")
		Output("blue", Tab+CustomSpaceAlign("[u]se <Item Name>", 25)+"- Use item")
		// Output("blue", "\tparry - Attemp to Parry incoming attack")
		// Output("blue", "\trun - Escape attack")
		// Output("blue", "\tget <Item Name> - Pick up item")
		// Output("blue", "\topen <Item Name> - Open an iten if it can be opened")
		Output("blue", Tab+CustomSpaceAlign("[i]nv", 25)+"- Display your inventory")
		Output("blue", Tab+CustomSpaceAlign("[m]ap", 25)+"- Display unveiled World map rooms\n\n")
	case Initial(commands.Quit):
		Output("green", "Goodbye...")
		SCORE.getSCORE()
		os.Exit(0)
	default:
		Output("red", "I didn't understand -> ", command)
	}
}

func minusOrSquash(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return 1
}
