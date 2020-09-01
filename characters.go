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
	Race  string
}

type Leveling struct {
	Rates                   *Specifics
	NextBase, NextRank, Exp int
	achievedLevelsChain     []int
}

type Special struct {
	Action     func(*Character)
	ShowAction func()
}

type Blueprint struct {
	Name      string
	Counter   int
	Timestamp time.Duration
}

type StatusEffectsBlueprint struct {
	AllStatus []*Blueprint
}

type Character struct {
	Welcome, Name, Details, From            string
	Health, Evasion, Strength, Boost, Skill int
	BaseHealth, Crit, LVL, ExpValue         int
	Alive, Npc                              bool
	Inventory                               map[string]*ItemQuantity
	Display                                 *DisplayImage
	LevelUp                                 *Leveling
	Special                                 *Special
	StatusEffects                           *StatusEffectsBlueprint

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

func (player *Character) showAction() {
	AsciiArts.showSkillAction(player.Name)
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

func (player *Character) didYourActionKillThatEnemy(e *Character) {
	loc := player.SetPlayerRoom()
	if loc.HasEnemy {
		if loc.Enemy == e {
			return
		}
	}
	if !e.isAlive() && e.Npc {
		Output("green", "\t", e.Name, translate(hasBeenSlainTR))
		SCORE.scoreKills(e.Name)
		EnemiesKilled++
		player.LevelUp.Exp += e.ExpValue
		player.getEnemyItems(e)
		player.calcLVL()
	}
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
		heal := item.Type.Effect
		if p.Name == heroesList.Wizard {
			heal += int(float32(heal) * .30)
		}
		if (p.Health + heal) > p.BaseHealth {
			p.Health = p.BaseHealth
			p.Inventory[name].Quantity--
			healing := strconv.Itoa(+heal)
			if p.Npc {
				Output("white", Tab+p.Name+translate(usePotionTR)+healing+translate(HPTR))
			} else {
				Output("green", Tab+p.Name+translate(usePotionTR)+healing+translate(HPTR))
			}
			break
		}
		p.Health += heal
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
		dmg := item.Type.Effect
		if p.Name == heroesList.Wizard {
			enemy.StatusEffects.Add(enemy, &Blueprint{
				Name:    statuses.Blight,
				Counter: 3,
			})
			// DoubleTab + "Enemy gets " + translate(blightStatusTR)
		}
		enemy.Health -= dmg
		p.Inventory[name].Quantity--
		if p.Npc {
			Output("white", Tab+p.Name+translate(useScrollTR)+Tab+translate(hitTR)+strconv.Itoa(+item.Type.Effect)+translate(dmgTR))
		} else {
			Output("green", Tab+p.Name+translate(useScrollTR)+Tab+translate(hitTR)+strconv.Itoa(+item.Type.Effect)+translate(dmgTR))
		}
		break
	case itemNames.Doll:
		p.Health = 30
		p.Alive = true
		p.Inventory[name].Quantity--
		Output("green", translate(dollUsedTR))
		time.Sleep(2 * time.Second)
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

func (player *Character) getSpecialAttackOnDragon() int {
	specialAttackOnDragon = map[string]map[string]int{
		heroesList.Thieve:    {races.Human: 0},
		heroesList.Paladin:   {races.Human: 0},
		heroesList.Wizard:    {races.Human: 0, races.Tiefling: 0},
		heroesList.Barbarian: {races.Human: 5 + specialOnDragonByDifficultyMap[Difficulty], races.Gnoll: 15 + specialOnDragonByDifficultyMap[Difficulty]},
	}
	return specialAttackOnDragon[player.Name][player.Display.Race]
}

func (player *Character) calculateDammage(enemy *Character) int {
	var modifier float32 = .0
	if enemy.Name == enemiesList.DRAGON {
		if player.Name == heroesList.Barbarian {
			if PercentChances(player.getSpecialAttackOnDragon()) {
				dmg := int(float32(dragon.BaseHealth) * .35)
				Output("white", translate(BarbarianLuckDragonTR)+strconv.Itoa(dmg)+translate(DamageTR))
				time.Sleep(1 * time.Second)
				return dmg
			}
		}
		if player.Name == heroesList.Wizard {
			if player.Display.Race == races.Tiefling {
				modifier = .02
			}
		}
	}

	calc := ((player.Strength * 80) / ((100 + enemy.Evasion) - (rand.Intn(10) + 5) - player.Boost)) * 7 / 10
	if calc > 20 {
		calc = (calc * 8) / 10
	}
	if enemy.Name == enemiesList.ORC {
		calc = calc / 2
	}
	dmg := Abs(calc)
	if rand.Intn(100) < player.Crit {
		dmg = Abs(dmg + (dmg * player.Crit / 100) + (dmg * (player.Strength / 100)))
		Output("red", "\t"+player.Name+translate(doesTR)+strconv.Itoa(dmg)+translate(critDMGTR)+enemy.Name)
		dmg = dmg + int(float32(dmg)*modifier)
		return dmg
	}
	Output("white", "\t"+player.Name+translate(doesTR)+strconv.Itoa(dmg)+translate(dmgToTR)+enemy.Name)
	if (player.Name == heroesList.Thieve) && (PercentChances(60)) {
		extra := ((dmg * 6) / 10)
		Output("white", "\t"+player.Name+translate(doesTR)+strconv.Itoa(extra)+translate(dmgToTR)+enemy.Name)
		dmg += extra
	}
	dmg = dmg + int(float32(dmg)*modifier)
	return dmg
}

func (player *Character) getAreaRooms() (locArr []*Location) {
	dirs := []string{}
	loc := player.SetPlayerRoom()
	if loc.X >= 1 {
		dirs = append(dirs, directions.West)
	}
	if loc.X < X-1 {
		dirs = append(dirs, directions.East)
	}

	if loc.Y >= 1 {
		dirs = append(dirs, directions.North)
	}
	if loc.Y < Y-1 {
		dirs = append(dirs, directions.South)
	}
	x := loc.X
	y := loc.Y

	if indexOf(dirs, directions.North) != -1 {
		locArr = append(locArr, WorldMap[y-1][x])
	}
	if indexOf(dirs, directions.South) != -1 {
		locArr = append(locArr, WorldMap[y+1][x])
	}

	if indexOf(dirs, directions.West) != -1 {
		locArr = append(locArr, WorldMap[y][x-1])

		if indexOf(dirs, directions.North) != -1 {
			locArr = append(locArr, WorldMap[y-1][x-1])
		}
		if indexOf(dirs, directions.South) != -1 {
			locArr = append(locArr, WorldMap[y+1][x-1])
		}
	}

	if indexOf(dirs, directions.East) != -1 {
		locArr = append(locArr, WorldMap[y][x+1])

		if indexOf(dirs, directions.North) != -1 {
			locArr = append(locArr, WorldMap[y-1][x+1])
		}
		if indexOf(dirs, directions.South) != -1 {
			locArr = append(locArr, WorldMap[y+1][x+1])
		}
	}
	return locArr
}

func (player *Character) showHealth() {
	loc := player.SetPlayerRoom()
	if loc.HasSeller {
		var concat string = ""
		for name, element := range loc.Item {
			concat += DoubleTab + name + "(" + strconv.Itoa(element.Quantity) + ")" + translate(forTR) + strconv.Itoa(element.Type.Price) + translate(forCoinsTR)
		}
		Output("yellow", loc.Seller)
		Output("yellow", translate(HasSellerTR)+concat)
	}
	if loc.HasEnemy && loc.Enemy.isAlive() {
		Output("red", translate(HasEnemyOrSellerTR0)+Article(loc.Enemy.Name+" lvl."+strconv.Itoa(loc.Enemy.LVL))+translate(HasEnemyTR1))
		loc.Enemy.showHP()
		loc.Enemy.showPlayerAfflictions()
	}
	player.showHP()
	player.showPlayerAfflictions()
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
	} else {
		ok = false
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
		// Output LEVEL UP
		if !player.Npc {
			Output("cyan", DoubleTab+CalculateSpaceAlign(translate(LevelUPTR))+"LVL "+strconv.Itoa(player.LVL+1)+" !!")
		}
		player.passLvl()
	} else {
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
	// var odd int = player.LVL % 2
	// To activate depending on how hard mode behaves
	// if player.Npc {
	// 	odd = player.LVL % 3
	// 	if odd == 2 {
	// 		odd = 0
	// 	}
	// }
	// if odd == 1 {
	// 	player.Skill += player.LevelUp.Rates.Skill
	// 	if !player.Npc {
	// 		Output("white", Tab+CalculateSpaceAlign(translate(SkillsUP))+"+ "+strconv.Itoa(player.LevelUp.Rates.Skill))
	// 	}
	// }
	player.ExpValue = player.ExpValue + int(float32(player.ExpValue)*.2)
	// for i := 0; i < (2 - odd); i++ {
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

	Output("blue", Tab+"lvl. "+strconv.Itoa(player.LVL)+strings.Repeat(" ", 3-utf8.RuneCountInString(strconv.Itoa(player.LVL)))+" ["+partA+partB+"]")
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
	Output("cyan", "\n"+DoubleTab+"================= "+translate(StatusTR)+" =================\n")
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

	Output("cyan", "\n"+DoubleTab+"================= "+translate(Skill)+" ================\n")
	Output("stats", translate(heroesSkillDescription[player.Name]))

	Output("cyan", "\n"+DoubleTab+"============== Status Effects ============\n")
	var allStatus []string
	for _, bp := range player.StatusEffects.AllStatus {
		allStatus = append(allStatus, bp.Name)
	}
	Output("stats", Tab+ArrayToString(allStatus))
	fmt.Println()
}

// **************************************************************************************
//                                    Skills
// **************************************************************************************

func (player *Character) oneOfRandItem() (string, *ItemQuantity) {
	for name, itemQ := range player.Inventory {
		ok := PercentChances(30)
		if ok {
			if itemQ.Quantity == 0 {
				continue
			}
			return name, itemQ
		}
	}
	var nameR string
	var itemQR *ItemQuantity
	for name, itemQ := range player.Inventory {
		nameR, itemQR = name, itemQ
		break
	}
	return nameR, itemQR
}

func (player *Character) useSkillSet(e *Character) {
	can := player.checkSkills()
	if !can {
		return
	}
	player.Skill--
	player.showAction()

	switch player.Name {

	case enemiesList.GOBLIN:
		fallthrough
	case heroesList.Thieve:
		name, itemQ := e.oneOfRandItem()
		ok := e.hasItemInInventory(name)
		if ok {
			ok = itemQ.Quantity > 0
		}
		if !ok && !e.Npc {
			Output("red", Tab+translate(TheTR)+player.Name+translate(StealFailTR))
			return
		}
		if name == itemNames.Coins {
			player.addItemTypeToInventory(name, itemQ.Quantity)
			e.Inventory[name].Quantity = 0
		}
		player.addItemTypeToInventory(name, 1)
		e.Inventory[name].Quantity--

		Output(playerEnemyColor[player.Npc], Tab+player.Name+translate(StealSuccessTR)+name+"\n")
		break

	case heroesList.Paladin:
		if e.Display.Race == races.Undead {
			e.Health = 0
			Output(playerEnemyColor[player.Npc], Tab+player.Name+translate(HolySuccessTR)+"\n")
			break
		}
		if e.Name == enemiesList.DRAGON {
			dmg := e.Health - int(float32(e.Health)*.75)
			e.Health = int(float32(e.Health) * .75)
			Output(playerEnemyColor[player.Npc], Tab+player.Name+translate(HolyHugeTR)+strconv.Itoa(dmg)+translate(HPTR)+"\n")
			break
		}
		e.BaseHealth = e.BaseHealth / 2
		e.Health = e.Health / 2
		Output(playerEnemyColor[player.Npc], Tab+player.Name+translate(HolyMitigatedTR)+"\n")
		break

	case heroesList.Wizard:
		locArr := player.getAreaRooms()
		Output(playerEnemyColor[player.Npc], translate(MagisterSkillTR))
		hit := 0
		for _, loc := range locArr {
			if loc.HasEnemy {
				loc.Enemy.Health = loc.Enemy.Health - 10
				player.didYourActionKillThatEnemy(loc.Enemy)
				hit++
			}
		}
		e.StatusEffects.Add(e, &Blueprint{
			Name:    statuses.Blight,
			Counter: 3,
		})
		Output(playerEnemyColor[player.Npc], DoubleTab+strconv.Itoa(hit)+translate(AreaHits))
		break

	case heroesList.Barbarian:
		locArr := player.getAreaRooms()
		for _, loc := range locArr {
			loc.Visited = true
		}
		e.StatusEffects.Add(e, &Blueprint{
			Name:    statuses.Fright,
			Counter: 2,
		})
		Output(playerEnemyColor[player.Npc], translate(DazbogRushSkillTR))
		break

	case enemiesList.SORCERER:
		reducer := .50
		firstLine := ""
		finalSentence := translate(DarkEnergyNormalTR)
		if e.Name == heroesList.Paladin {
			reducer = .25
			finalSentence = translate(DarkEnergyOnPaladinTR)
			firstLine = translate(GraceProtectsYouTR)
		}
		e.Health = e.Health - int(float64(e.Health)*reducer)
		Output("red", firstLine+translate(DarkEnergyTR)+finalSentence)
		break

	case enemiesList.DRAGON:
		dmg := rand.Intn(10) + 15
		Output("red", translate(DragonSkillFireTR)+strconv.Itoa(dmg)+translate(HPTR))
		// Apply reducers here
		if e.Name == heroesList.Wizard {
			Output("green", translate(SorcererDragonFireTR))
			dmg = dmg - int(float32(dmg)*.30)
		}
		e.Health = e.Health - dmg
		time.Sleep(1 * time.Second)
		break
	}
	time.Sleep(2 * time.Second)
}

func (player *Character) checkSkills() bool {
	b := true
	if player.Skill < 1 {
		b = false
		if !player.Npc {
			Output("red", Tab+translate(NoSkillTR))
		} else {
			Output("yellow", Tab+player.Name+translate(StealFailTR))
		}
	}
	return b
}

// ***********************************************************************************
//																	Status Effects
// ***********************************************************************************

func (ste *StatusEffectsBlueprint) Add(p *Character, args ...*Blueprint) {
	statuses := args
	for _, status := range statuses {
		ste.remove(status.Name)
		p.logStatusAfflicted(status.Name)
	}
	ste.AllStatus = append(ste.AllStatus, statuses...)
}

func (player *Character) logStatusAfflicted(name string) {
	NPC := player.Npc

	switch name {
	case statuses.Blight:
		text := translate(heroGotBlightTR)
		if NPC {
			text = translate(EnemiGotBlightTR)
		}
		Output(playerEnemyColor[!NPC], text)
		break
	case statuses.Fright:
		text := translate(heroGotFrightTR)
		if NPC {
			text = translate(EnemiGotFrightTR)
		}
		Output(playerEnemyColor[!NPC], text)
		break
	}
}

func (ste *StatusEffectsBlueprint) remove(name string) {
	i := indexOfBlueprint(ste.AllStatus, name)
	if i >= 0 {
		ste.AllStatus = append(ste.AllStatus[:i], ste.AllStatus[i+1:]...)
	}
}

func indexOfBlueprint(arr []*Blueprint, item string) int {
	for index, elem := range arr {
		if elem.Name == item {
			return index
		}
	}
	return -1
}

func (player *Character) showPlayerAfflictions() {
	if len(player.StatusEffects.AllStatus) > 0 {
		list := []string{}
		for _, status := range player.StatusEffects.AllStatus {
			list = append(list, status.Name)
		}
		Output(playerEnemyColor[!player.Npc], Tab+CalculateSpaceAlign(translate(statusEffectsTR))+ArrayToString(list))
	}
}

func (player *Character) applyStatusesEffect() {
	if !player.isAlive() {
		player.StatusEffects.AllStatus = []*Blueprint{}
		return
	}
	for _, status := range player.StatusEffects.AllStatus {
		// fmt.Printf("%s: %+v  %+v\n", status.Name, status.Counter, status.Timestamp.Seconds())
		player.Affliction(status)
	}
}

func (player *Character) Affliction(status *Blueprint) {
	switch status.Name {
	case statuses.Blight:
		status.Counter--
		flames := rand.Intn(5) + 1
		player.Health -= flames
		Output(playerEnemyColor[!player.Npc], DoubleTab+player.Name+translate(burnsTR)+strconv.Itoa(flames)+translate(dmgTR))
		hero.didYourActionKillThatEnemy(player)
		if status.Counter <= 0 {
			player.StatusEffects.remove(statuses.Blight)
		}
		break
	case statuses.Fright:
		if status.Counter <= 0 {
			player.StatusEffects.remove(statuses.Fright)
			break
		}
		status.Counter--
		Output(playerEnemyColor[!player.Npc], DoubleTab+player.Name+translate(cantMoveTR))
		setTurnsFrightStatus(player.Npc)
		// ResetTurns()
		break
	}

	checkPlayers()
}

// panic: runtime error: invalid memory address or nil pointer dereference
// [signal 0xc0000005 code=0x1 addr=0x0 pc=0x4d4353]

// goroutine 1 [running]:
// main.(*ScoreSchema).scoreItems(...)
//         C:/dev/GO/DarkAdventures/scores.go:52
// main.(*Character).addItemTypeToInventory(...)
//         C:/dev/GO/DarkAdventures/characters.go:158
// main.(*Character).useSkillSet(0xc000158270, 0xc0001581a0)
//         C:/dev/GO/DarkAdventures/characters.go:582 +0xe53
// main.ProcessCommands(0xc000158270, 0xc0000ed033, 0x3, 0xc00006fe68, 0x1, 0x1)
//         C:/dev/GO/DarkAdventures/commands.go:51 +0x17ea
// main.Battle(0xc000158270, 0xc0001581a0)
//         C:/dev/GO/DarkAdventures/battles.go:43 +0xb14
// main.PresentScene(0xc000158270)
//         C:/dev/GO/DarkAdventures/rendering.go:36 +0x6fc
// main.main()
//         C:/dev/GO/DarkAdventures/main.go:91 +0xea
