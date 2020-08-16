package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type ItemQuantity struct {
	Type     Item
	Quantity int
}

type DisplayImage struct {
	Image string
	Show  bool
	Story string
}

type Leveling struct {
	Rates                   *Specifics
	NextBase, NextRank, Exp int
	achievedLevelsChain     []int
}

type Character struct {
	Welcome, Name, Details, From            string
	Health, Evasion, Strength, Boost, Skill int
	BaseHealth, Crit, LVL, ExpValue         int
	Alive, Npc                              bool
	Inventory                               map[string]*ItemQuantity
	Display                                 *DisplayImage
	LevelUp                                 *Leveling

	CurrentLocation []int
}

func (player *Character) SetPlayerRoom() *Location {
	var loc *Location
	x := player.CurrentLocation[1]
	y := player.CurrentLocation[0]
	loc = WorldMap[y][x]
	if !player.Npc {
		WorldMap[y][x].Visited = true
	}
	return loc
}

func (player *Character) setImage() {
	player.Display = AsciiArts.makeImage(player.Name)
}

func (player *Character) getImage() {
	if player.Display.Show {
		Output("red", player.Display.Image)
		if !player.Npc {
			Output("blue", player.Display.Story)
			time.Sleep(2 * time.Second)
		}
		player.Display.Show = false
	}
}

func (player *Character) MoveTo(direction string) bool {
	ok := false
	switch strings.ToLower(direction) {
	case strings.ToLower(directions.West):
		fallthrough
	case Initial(directions.West):
		if player.CurrentLocation[1] > 0 {
			player.CurrentLocation[1]--
			ok = true
			player.From = directions.East
		}
		break
	case strings.ToLower(directions.East):
		fallthrough
	case Initial(directions.East):
		if player.CurrentLocation[1] < X-1 {
			player.CurrentLocation[1]++
			ok = true
			player.From = directions.West
		}
		break
	case strings.ToLower(directions.North):
		fallthrough
	case Initial(directions.North):
		if player.CurrentLocation[0] > 0 {
			player.CurrentLocation[0]--
			ok = true
			player.From = directions.South
		}
		break
	case strings.ToLower(directions.South):
		fallthrough
	case Initial(directions.South):
		if player.CurrentLocation[0] < Y-1 {
			player.CurrentLocation[0]++
			ok = true
			player.From = directions.North
		}
		break
	default:
		if !player.Npc {
			Output("red", translate(CanTGoTR)+vowelOrNot(direction, directionsArticlesVowelsTR, directionsArticlesConsonantTR)+direction)
		}
	}
	return ok
}

func (p *Character) escapeBattle(b bool) {
	if b {
		if p.From != "" {
			// escaped!
			Output("green", escapeCases[p.Name][escapeResults.OK])
			time.Sleep(1 * time.Second)
			p.MoveTo(p.From)
		} else {
			// escaped randomly
			canGo := p.SetPlayerRoom().CanGoTo
			p.MoveTo(canGo[rand.Intn(len(canGo))])
			Output("green", escapeCases[p.Name][escapeResults.RAND])
			time.Sleep(1 * time.Second)
		}
	} else {
		// missed escape
		Output("green", escapeCases[p.Name][escapeResults.KO])
		time.Sleep(1 * time.Second)
	}
}

func (p *Character) addItemTypeToInventory(n string, i int) {
	if !p.Npc {
		SCORE.scoreItems(n, i)
	}
	if p.hasItemInInventory(n) {
		p.Inventory[n].Quantity += i
	} else {
		p.Inventory[n] = &ItemQuantity{Type: *ItemList[n], Quantity: i}
	}
}

func (p *Character) useItem(name string, enemyInArr ...interface{}) bool {
	item := p.Inventory[name]
	if !p.Npc && (!p.hasItemInInventory(name) || item.Quantity < 1) {
		Output("red", translate(youDontHaveATR), name)
		return false
	}
	if !p.Npc && (!UsableItems[name]) && p.Alive {
		Output("red", translate(youCanTUseATR), name)
		return !UsableItems[name]
	}

	switch name {
	case itemNames.Key:
		if UseKey(p) {
			p.Inventory[name].Quantity--
		}
		break
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
		Output("green", translate(moonstoneUsedTR))
		Output("green", Tab+CalculateSpaceAlign(translate(strengthBoostAddTR))+strconv.Itoa(p.Strength+p.Boost)+"\n")
		break
	case itemNames.Scroll:
		enemy := enemyInArr[0].(*Character)
		enemy.Health -= item.Type.Effect
		p.Inventory[name].Quantity--
		break
	case itemNames.Doll:
		p.Health = 30
		p.Alive = true
		p.Inventory[name].Quantity--
		Output("green", translate(dollUsedTR))
	}
	return true
}

func (player *Character) showHP() {
	color := "green"
	if player.Npc {
		color = "yellow"
	}
	Output(color, Tab+CalculateSpaceAlign(player.Name+": ")+"["+strconv.Itoa(player.Health)+" / "+strconv.Itoa(player.BaseHealth)+" HP]")
}

func (player *Character) attack(enemy *Character) {
	if !evadeAttack(player, enemy) {
		dmg := player.calculateDammage(enemy)
		SCORE.scoreDammages(player.Npc, dmg)
		enemy.Health -= dmg
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
	// fmt.Println("calc ", calc)
	// fmt.Println("dmg ", dmg)
	// fmt.Println("Boost ", player.Boost)
	if rand.Intn(100) < player.Crit {
		dmg = Abs(dmg + (dmg * player.Crit / 100) + (dmg * (player.Strength / 100)))
		Output("red", "\t"+player.Name+translate(doesTR)+strconv.Itoa(dmg)+translate(critDMGTR)+enemy.Name)
		return dmg
	}
	Output("white", "\t"+player.Name+translate(doesTR)+strconv.Itoa(dmg)+translate(dmgToTR)+enemy.Name)
	if (player.Name == heroesList.Thieve) && (PercentChances(60)) {
		extra := ((dmg * 6) / 10)
		Output("white", "\t"+player.Name+translate(doesTR)+strconv.Itoa(extra)+translate(dmgToTR)+enemy.Name)
		dmg += extra
	}
	return dmg
}

func (player *Character) showHealth() {
	loc := player.SetPlayerRoom()
	if loc.HasSeller {
		var concat string = ""
		for name, element := range loc.Item {
			concat += DoubleTab + strconv.Itoa(element.Quantity) + " - " + name + translate(forTR) + strconv.Itoa(element.Type.Price) + translate(forCoinsTR)
		}
		Output("yellow", loc.Seller)
		Output("yellow", translate(HasSellerTR)+concat)
	}
	if loc.HasEnemy && loc.Enemy.isAlive() {
		Output("red", translate(HasEnemyOrSellerTR0)+Article(loc.Enemy.Name+" LVL "+strconv.Itoa(loc.Enemy.LVL))+translate(HasEnemyTR1))
		loc.Enemy.showHP()
	}
	player.showHP()
}

func (player *Character) isAlive() bool {
	player.Alive = player.Health > 0
	if player.Name == enemiesList.DRAGON && !player.Alive {
		loc := player.SetPlayerRoom()
		loc.Ephemeral = ""
	}
	return player.Alive
}

func (player *Character) DisplayInvetory() {
	Output("green", translate(yourInventoryTR))
	for _, item := range player.Inventory {
		if item.Quantity > 0 {
			Output("green", Tab+CustomSpaceAlign(item.Type.Name+": "+item.Type.Description, inventorySpace)+translate(inventoryQuantityTR), item.Quantity)
		}
	}
	fmt.Println()
}

func (player *Character) DisplayItems() string {
	var items []string
	for _, item := range player.Inventory {
		if UsableItems[item.Type.Name] && item.Quantity > 0 {
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
		Output("green", translate(getEnemyItemsTR))
		for name, item := range enemy.Inventory {
			Output("green", Tab+CalculateSpaceAlign(name+": "+item.Type.Description+" -> "), item.Quantity)
			player.addItemTypeToInventory(name, item.Quantity)
		}
	} else {
		Output("green", translate(nothingYouCouldGetTR))
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
		Output("green", DoubleTab+player.Name+translate(missedTR))
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
					Output("green", translate(youBaughtTR)+Article(name))
					result = true
				}
			}
		}
	}
	return result
}

// ***************************************************************************************************
//                                            LEVEL UP
// ***************************************************************************************************

func (player *Character) calcLVL() {
	next := player.LevelUp.NextRank
	xp := player.getXPtoNext()
	if xp-next >= 0 {
		player.passLvl()
		// Output LEVEL UP
	} else {
		// Output ENEMY LEVEL
		return
	}
	player.calcLVL()
}

func (player *Character) passLvl() {
	next := player.getTheNext()
	player.LVL++
	player.setNewNext()
	player.LevelUp.achievedLevelsChain = append(player.LevelUp.achievedLevelsChain, next)
	player.applyRates()
}

func (player *Character) getXPtoNext() (res int) {
	allCaps := 0
	for i := 0; i < len(player.LevelUp.achievedLevelsChain); i++ {
		allCaps += player.LevelUp.achievedLevelsChain[i]
	}
	res = player.LevelUp.Exp - allCaps
	return res
}

func (player *Character) getTheNext() (res int) {
	res = player.LevelUp.NextBase
	for i := 0; i < player.LVL; i++ {
		res = res + res*20/100
	}
	return res
}

func (player *Character) setNewNext() {
	player.LevelUp.NextRank = player.LevelUp.NextRank + player.LevelUp.NextRank*20/100
}

func (player *Character) applyRates() {
	for i := 0; i < 2; i++ {
		operator := lvlRatesMap[rand.Intn(len(lvlRatesMap))]
		switch operator {
		case LevelingNames.Health:
			HPMaxAdd := player.LevelUp.Rates.Health()
			player.BaseHealth += HPMaxAdd
			if !player.Npc {
				Output("white", Tab+CalculateSpaceAlign(translate(HealthUP))+"+ "+strconv.Itoa(HPMaxAdd))
			}
			break
		case LevelingNames.Crit:
			player.Crit += player.LevelUp.Rates.Crit
			if !player.Npc {
				Output("white", Tab+CalculateSpaceAlign(translate(CritsUP))+"+ "+strconv.Itoa(player.LevelUp.Rates.Crit))
			}
			break
		case LevelingNames.Evasion:
			player.Evasion += player.LevelUp.Rates.Evasion
			if !player.Npc {
				Output("white", Tab+CalculateSpaceAlign(translate(EvasionUP))+"+ "+strconv.Itoa(player.LevelUp.Rates.Evasion))
			}
			break
		case LevelingNames.Skill:
			player.Skill += player.LevelUp.Rates.Skill
			if !player.Npc {
				Output("white", Tab+CalculateSpaceAlign(translate(SkillsUP))+"+ "+strconv.Itoa(player.LevelUp.Rates.Skill))
			}
			break
		case LevelingNames.Strength:
			player.Strength += player.LevelUp.Rates.Strength
			if !player.Npc {
				Output("white", Tab+CalculateSpaceAlign(translate(StrengthUP))+"+ "+strconv.Itoa(player.LevelUp.Rates.Strength))
			}
			break
		}
	}
}

var lvlRatesMap []string = []string{LevelingNames.Health, LevelingNames.Crit, LevelingNames.Evasion, LevelingNames.Skill, LevelingNames.Strength}

type LevelingNamesStruct struct {
	Health   string
	Crit     string
	Evasion  string
	Skill    string
	Strength string
}

var LevelingNames *LevelingNamesStruct = &LevelingNamesStruct{
	Health:   "Health",
	Crit:     "Crit",
	Evasion:  "Evasion",
	Skill:    "Skill",
	Strength: "Strength",
}

func (player *Character) DisplayExpGauge() {
	next := player.LevelUp.NextRank
	xp := player.getXPtoNext()
	length := xp * gaugeSize / next
	if length > gaugeSize {
		length = gaugeSize
	}
	partA := strings.Repeat(expChar, length)
	partB := strings.Repeat(emptyGauge, gaugeSize-length)

	Output("green", Tab+"LVL: "+strconv.Itoa(player.LVL)+strings.Repeat(" ", 3-utf8.RuneCountInString(strconv.Itoa(player.LVL)))+" ["+partA+partB+"]")
}

func (player *Character) DisplayStats() {
	exp := player.getXPtoNext()
	r := 0
	tm := 0
	for _, y := range WorldMap {
		for _, loc := range y {
			tm++
			if loc.Visited {
				r++
			}
		}
	}
	Output("cyan", "\n"+DoubleTab+"================= "+translate(Status)+" =================\n")
	Output("stats", Tab+CalculateSpaceAlign(translate(Health)+":")+strconv.Itoa(player.Health)+"/"+strconv.Itoa(player.BaseHealth)+"  "+
		Tab+CalculateSpaceAlign(translate(CritsUP))+strconv.Itoa(player.Crit))
	Output("stats", Tab+CalculateSpaceAlign(translate(StrengthUP))+strconv.Itoa(player.Strength)+"  "+
		Tab+CalculateSpaceAlign(translate(BoostTR)+":")+strconv.Itoa(player.Boost))
	Output("stats", Tab+CalculateSpaceAlign(translate(Level)+":")+strconv.Itoa(player.LVL)+"  "+
		Tab+CalculateSpaceAlign(translate(EvasionUP))+strconv.Itoa(player.Evasion))
	Output("stats", Tab+CalculateSpaceAlign(translate(Exp)+":")+strconv.Itoa(exp)+"/"+strconv.Itoa(player.LevelUp.NextRank)+"  "+
		Tab+CalculateSpaceAlign(translate(Rooms)+":")+strconv.Itoa(r)+"/"+strconv.Itoa(tm))
	Output("stats", Tab+CalculateSpaceAlign(translate(Enemies)+":")+strconv.Itoa(EnemiesKilled)+"/"+strconv.Itoa(EnemiesCount)+"  "+
		Tab+CalculateSpaceAlign(translate(SkillsUP))+strconv.Itoa(player.Skill))
	Output("stats", Tab+CalculateSpaceAlign(translate(Skill)+":")+translate(heroesSkillDescription[player.Name]))

	fmt.Println()
}

// sante, crit, strength, evade,
