package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type ItemQuantity struct {
	Type     Item
	Quantity int
}

type Character struct {
	Welcome, Name                                             string
	Health, Evasion, Strength, Boost, Skill, BaseHealth, Crit int
	Alive, Npc                                                bool
	Inventory                                                 map[string]*ItemQuantity

	CurrentLocation []int
}

func (player *Character) SetPlayerRoom() *Location {

	x := player.CurrentLocation[1]
	y := player.CurrentLocation[0]
	loc := WorldMap[y][x]
	WorldMap[y][x].Visited = true
	return loc
}

func (player *Character) MoveTo(direction string) bool {

	ok := false
	switch strings.ToLower(direction) {
	case "west":
		fallthrough
	case "w":
		if player.CurrentLocation[1] > 0 {
			player.CurrentLocation[1]--
			ok = true
		}
		break
	case "east":
		fallthrough
	case "e":
		if player.CurrentLocation[1] < 9 {
			player.CurrentLocation[1]++
			ok = true
		}
		break
	case "north":
		fallthrough
	case "n":
		if player.CurrentLocation[0] > 0 {
			player.CurrentLocation[0]--
			ok = true
		}
		break
	case "south":
		fallthrough
	case "s":
		if player.CurrentLocation[0] < 9 {
			player.CurrentLocation[0]++
			ok = true
		}
		break
	default:
		Output("red", "\tCan't go ", direction)
	}
	// player.SetPlayerRoom()
	return ok
}

// func (player *Character) create() {
// 	Outputf("blue", "\t%s", "Welcome to Adventures, choose a hero:")
// 	cmd := UserInputln()
// 	fmt.Print(cmd)
// }

func (p *Character) addItemTypeToInventory(n string, i int) {

	if p.hasItemInInventory(n) {
		p.Inventory[n].Quantity += i
	} else {
		p.Inventory[n] = &ItemQuantity{Type: *ItemList[n], Quantity: i}
	}
}

func (p *Character) useItem(name string, enemyInArr ...interface{}) bool {
	item := p.Inventory[name]
	if !p.Npc && (!p.hasItemInInventory(name) || item.Quantity < 1) {
		Output("red", "You don't have a ", name)
		return false
	}
	if !p.Npc && (!UsableItems[name]) && p.Alive {
		Output("red", "You can't use a ", name)
		return !UsableItems[name]
	}

	switch name {
	case itemNames.Potion:
		if (p.Health + item.Type.Effect) > p.BaseHealth {
			p.Health = p.BaseHealth
			p.Inventory[name].Quantity--
			break
		}
		p.Health += item.Type.Effect
		p.Inventory[name].Quantity--
		break
	case itemNames.Moonstone:
		p.Boost += item.Type.Effect
		p.Inventory[name].Quantity--
		break
	case itemNames.Scroll:
		enemy := enemyInArr[0].(*Character)
		enemy.Health -= item.Type.Effect
		p.Inventory[name].Quantity--
		break
	case itemNames.Doll:
		p.Health = 30
		p.Alive = true
		Output("green", Tab+"A dark force is devouring your body")
		Output("green", Tab+"A chance has been given to youm or is it?")
		Output("green", Tab+"You died... and revived.")
		Output("green", Tab+"Health +30 HP")
	}
	return true
}

func (player *Character) showHP() {
	color := "green"
	if player.Npc {
		color = "yellow"
	}
	Output(color, Tab+CustomSpaceAlign(player.Name+": ", inventorySpace)+strconv.Itoa(player.Health)+" / "+strconv.Itoa(player.BaseHealth)+" HP")
}

func (player *Character) attack(enemy *Character) {
	if !evadeAttack(player, enemy) {
		enemy.Health -= player.calculateDammage(enemy)
	}
}

func (player *Character) calculateDammage(enemy *Character) int {
	// calc := enemy.BaseHealth*player.Strength/100*(player.Strength+player.Boost)/10 - (rand.Intn(10) + 5)
	calc := ((player.Strength * 80) / ((100 + enemy.Evasion) - (rand.Intn(10) + 5) - player.Boost)) * 7 / 10
	if calc > 20 {
		calc = (calc * 8) / 10
	}
	if enemy.Name == enemiesList.ORC {
		calc = calc / 2
	}
	dmg := Abs(calc)
	// ICI
	fmt.Println("calc ", calc)
	fmt.Println("dmg ", dmg)
	fmt.Println("Boost ", player.Boost)
	if rand.Intn(100) < player.Crit {
		dmg = Abs(dmg + (dmg * player.Crit / 100) + (dmg * (player.Strength / 100)))
		Output("red", "\t"+player.Name+" does "+strconv.Itoa(dmg)+" Critical DMG to "+enemy.Name)
		return dmg
	}
	Output("white", "\t"+player.Name+" does "+strconv.Itoa(dmg)+" DMG to "+enemy.Name)
	if (player.Name == heroesList.Thieve) && (PercentChances(60)) {
		extra := ((dmg * 6) / 10)
		Output("white", "\t"+player.Name+" does "+strconv.Itoa(extra)+" DMG to "+enemy.Name)
		dmg += extra
	}
	return dmg
}

func (player *Character) isAlive() bool {
	player.Alive = player.Health > 0
	return player.Alive
}

func (player *Character) DisplayInvetory() {
	Output("green", DoubleTab+"Your inventory:")
	for _, item := range player.Inventory {
		if item.Quantity > 0 {
			Output("green", Tab+CustomSpaceAlign(item.Type.Name+": "+item.Type.Description, inventorySpace)+"Quatity: ", item.Quantity)
		}
	}
	player.showHP()
	fmt.Println()
}

func (player *Character) DisplayItems() string {
	var items []string
	for _, item := range player.Inventory {
		if UsableItems[item.Type.Name] {
			items = append(items, item.Type.Name+"("+strconv.Itoa(item.Quantity)+")")
		}
	}
	return "[" + strings.Join(items, " ") + "]"
}

func (player *Character) hasItemInInventory(name string) bool {
	_, ok := player.Inventory[name]
	if ok {
		ok = player.Inventory[name].Quantity >= 1
	}
	return ok
}

func (player *Character) getEnemyItems(enemy *Character) {
	if InventoryHasItem(enemy.Inventory) {
		Output("green", DoubleTab+"You get:")
	} else {
		Output("green", Tab+"Enemy had nothing you could use...")
	}
	for name, item := range enemy.Inventory {
		Output("green", Tab+CalculateSpaceAlign(name+": "+item.Type.Description+" -> "), item.Quantity)
		player.addItemTypeToInventory(name, item.Quantity)
	}
}

func (enemy *Character) createEnemyInventory() {
	anchor := rand.Intn(101)
	A := "a"
	B := 0

	for key, val := range ItemChancesByEnemyName[enemy.Name] {
		if anchor >= B && val > B && val < anchor {
			A = key
			B = val
		}
	}
	if A != "a" {
		if A == itemNames.Coins {
			Q := rand.Intn(5) + rand.Intn(5) + rand.Intn(5) + rand.Intn(3)
			enemy.addItemTypeToInventory(A, Q)
			aggregate[A] += Q
			itemsByEnemy[enemy.Name][A] += Q
		} else {
			enemy.addItemTypeToInventory(A, 1)
			aggregate[A]++
			itemsByEnemy[enemy.Name][A]++
		}
	}
}

func evadeAttack(player, enemy *Character) bool {
	if PercentChances(enemy.Evasion) {
		Output("green", DoubleTab+player.Name+" MISSED!!")
		return true
	}
	return false
}

func (player *Character) getPurse() int {
	if coins, ok := player.Inventory[itemNames.Coins]; ok {
		return coins.Quantity
	}
	return 0
}

func (player *Character) spendMoney(amount int) bool {
	if player.getPurse() >= amount {
		player.Inventory[itemNames.Coins].Quantity -= amount
		return true
	}
	return false
}

func (player *Character) BuyFromShop(name string) bool {
	result := false
	loc := player.SetPlayerRoom()

	if loc.HasSeller {
		if itemQ, ok := loc.Item[name]; ok {
			if itemQ.Quantity >= 1 {
				if ok := player.spendMoney(itemQ.Type.Price); ok {
					player.addItemTypeToInventory(name, 1)
					loc.RemoveItem(name)
					Output("green", Tab+"You baught "+Article(name))
					result = true
				}
			}
		}
	}
	return result
}
