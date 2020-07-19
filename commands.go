package main

import (
	"os"
	"strings"
)

func ProcessCommands(player *Character, input string, args ...interface{}) {
	Output("yellow", "======================================================================")
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		Output("red", "No command received.")
		ResetTurns()
		return
	}

	command := strings.ToLower(tokens[0])
	itemName := ""
	if len(tokens) > 1 {
		itemName = tokens[1]
	}
	Output("yellow", "tokens ", tokens)
	Output("yellow", "command ", command)
	Output("yellow", "itemName ", itemName)
	loc := player.SetPlayerRoom()
	switch Initial(command) {
	case Initial(commands.Go):
		ResetTurns()
		if itemName != "" {
			if InitialsIndexOf(loc.CanGoTo, itemName) && !loc.Enemy.isAlive() {
				player.MoveTo(itemName)
			} else {
				Output("red", "Can't go to "+itemName+" from here.")
			}
		} else {
			Output("red", "What do you mean? ", command, " ", itemName)
		}

	case Initial(commands.Attack):
		enemy := &loc.Enemy
		if enemy != nil {
			player.attack(enemy)
		}

	case Initial(commands.Use): // to do
		Output("red", "in the use case")
		if itemName == itemNames.Scroll && player.hasItemInInventory(itemName) {
			Output("red", "in the use if", args[0])
			enemy := &loc.Enemy
			Output("red", "enemy: ", enemy)
			Output("yellow", "has "+itemName+": ", player.hasItemInInventory(itemName), " quantity: ", player.Inventory[itemName].Quantity)
			ok := player.useItem(itemName, enemy)
			if ok != true {
				Output("red", "Couldn't use the "+itemName)
				ResetTurns()
			}
		} else {
			ok := player.useItem(itemName)
			if ok != true {
				ResetTurns()
			}
		}
	case Initial(commands.Map):
		ResetTurns()
		DisplayWorldMap(player)
	case Initial(commands.Inv):
		ResetTurns()
		player.DisplayInvetory()
	case Initial(commands.Buy):
		if ok := player.BuyFromShop(itemName); !ok {
			// ICI
			Output("red", Tab+"Can't buy "+Article(itemName))
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
		Output("blue", Tab+CustomSpaceAlign("go <Location Name>", 25)+"- Go to the new location")
		Output("blue", Tab+CustomSpaceAlign("attack", 25)+"- Attacks opponent(s)")
		// Output("blue", "\tparry - Attemp to Parry incoming attack")
		// Output("blue", "\trun - Escape attack")
		// Output("blue", "\tget <Item Name> - Pick up item")
		// Output("blue", "\topen <Item Name> - Open an iten if it can be opened")
		Output("blue", Tab+CustomSpaceAlign("inv", 25)+"- Display your inventory\n\n")
	case Initial(commands.Quit):
		Output("green", "Goodbye...")
		os.Exit(0)
	default:
		Output("red", "I didn't understand -> ", command)
	}
}
