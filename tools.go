package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

var WorldMap [Y][X]*Location

// WhereCanYouGo location => []string
func WhereCanYouGo(yourPlace *Location) []string {
	var youCanGo []string

	if yourPlace.X >= 1 {
		youCanGo = append(youCanGo, directions.West)
	}
	if yourPlace.X <= 8 {
		youCanGo = append(youCanGo, directions.East)
	}

	if yourPlace.Y >= 1 {
		youCanGo = append(youCanGo, directions.North)
	}
	if yourPlace.Y <= 8 {
		youCanGo = append(youCanGo, directions.South)
	}
	return youCanGo
}

func Intro() {
	Output("blue", gameIntro)
	Output("blue", DoubleTab+"1 - Easy")
	Output("blue", DoubleTab+"2 - Meddium")
	Output("blue", DoubleTab+"3 - Hard")
	UserInput(&Difficulty)
	if Difficulty > 3 || Difficulty < 1 {
		Intro()
	} else {
		Difficulty = Difficulty - 1
	}
}

func ChooseHero() {
	Output("blue", Tab+"Select your hero:")
	for index, name := range indexedHeroes {
		Output("blue", DoubleTab+strconv.Itoa(index+1)+" - "+name)
	}
	UserInput(&Hero)
	if Hero > 4 || Hero < 1 {
		ChooseHero()
	} else {
		Hero = Hero - 1
	}
}

var aggregate = map[string]int{
	itemNames.Moonstone: 0,
	itemNames.Doll:      0,
	itemNames.Coins:     0,
	itemNames.Key:       0,
	itemNames.Potion:    0,
	itemNames.Scroll:    0,
}
var aggregateEnemies = map[string]int{
	enemiesList.GOBLIN:   0,
	enemiesList.SKELETON: 0,
	enemiesList.ORC:      0,
	enemiesList.SORCERER: 0,
	enemiesList.DRAGON:   0,
}

func AnalyzeItemsRepartition() {
	for e, items := range itemsByEnemy {
		Output("green", DoubleTab+e+":")
		for in, iq := range items {
			if iq > 0 {
				Output("yellow", Tab+CalculateSpaceAlign(in), iq)
			}
		}
	}
}

func CreateMap() [Y][X]*Location {
	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			hasItem := false
			if rand.Intn(100) < 20 {
				hasItem = true
			}

			item := getItem(hasItem)
			WorldMap[y][x] = &Location{
				Description: Tab + "You are" + RoomFromLandscape[Grid[y][x]][rand.Intn(len(RoomFromLandscape[Grid[y][x]]))] + Tab + Ambiance[rand.Intn(len(Ambiance))],
				HasSeller:   hasItem,
				Item:        item,
				X:           x,
				Y:           y,
			}
			if hasItem {
				WorldMap[y][x].Seller = Tab + "There is" + getSeller(hasItem)
			}
			WorldMap[y][x].CanGoTo = WhereCanYouGo(WorldMap[y][x])
			if !((y == 9) && (x == 4)) {
				enemyProbability(WorldMap[y][x])
			}
			if WorldMap[y][x].HasEnemy {
				WorldMap[y][x].Enemy.createEnemyInventory()
			}
		}
	}
	for b := 0; b < Y; b++ {
		for a := 0; a < Y; a++ {
			addOrcProximity(WorldMap[a][b])
		}
	}
	CreateDragon()
	showGameItems()
	showGameEnemies()
	AnalyzeItemsRepartition()

	fmt.Println("Difficulty", Difficulty+1, difficultyIndex[Difficulty])
	return WorldMap
}

func addOrcProximity(loc *Location) {
	if loc.HasEnemy {
		if loc.Enemy.Name == enemiesList.ORC {
			if near := strings.Contains(loc.Description, "odd smell"); !near {
				loc.addDescriptionToAdjacentRooms(NearORC)
			}
		}
	}
}

// fmt.Printf("\nBEFORE: %+v\n", loc.Enemy)
// fmt.Printf("\nBEFORE: %+v\n", loc.Enemy.Inventory)
// fmt.Println(loc.HasEnemy)
// loc.Enemy.createEnemyInventory()
// fmt.Printf("\nAFTER: %+v\n", loc.Enemy.Inventory)

func showGameItems() {
	fmt.Println(DoubleTab + CalculateSpaceAlign("Items:"))
	for a, b := range aggregate {
		fmt.Println(Tab+CalculateSpaceAlign(a), b)
	}
}

func showGameEnemies() {
	fmt.Println(DoubleTab + CalculateSpaceAlign("Enemies:"))
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			if WorldMap[i][j].HasEnemy {
				aggregateEnemies[WorldMap[i][j].Enemy.Name]++
			}
		}
	}
	for a, b := range aggregateEnemies {
		fmt.Println(Tab+CalculateSpaceAlign(a), b)
	}
}

func getItem(b bool) map[string]*ItemQuantity {
	var result Item
	if b {
		result = *ItemList[ItemIndexList[rand.Intn(5)]]
		result.Price = getRandPrice()
	}
	return map[string]*ItemQuantity{
		result.Name: {
			Type:     result,
			Quantity: itemAmounts(result.Name),
		},
	}
}

func itemAmounts(n string) int {
	switch n {
	case itemNames.Key:
		fallthrough
	case itemNames.Doll:
		fallthrough
	case itemNames.Moonstone:
		return 1
	case itemNames.Scroll:
		fallthrough
	case itemNames.Potion:
		return rand.Intn(2) + 1
	case itemNames.Coins:
		return (rand.Intn(11) + 3)
	default:
		return 0
	}
}

func getSeller(b bool) string {
	var str string
	if b {
		str = SellerList[rand.Intn(3)]
	}
	return str
}

func Initial(s string) string {
	return strings.ToLower(string(s[0]))
}

func InventoryHasItem(arr map[string]*ItemQuantity) bool {
	for name, item := range arr {
		if name != "" && item.Quantity > 0 {
			return true
		}
	}
	return false
}

func indexOf(arr []string, item string) int {
	for index, elem := range arr {
		if elem == item {
			return index
		}
	}
	return -1
}

func InitialsIndexOf(arr []string, item string) bool {
	boolean := false
	for _, elem := range arr {
		if Initial(elem) == Initial(item) {
			boolean = true
		}
	}
	return boolean
}

var heroFromName = func(s string) *Character {
	// json.Unmarshal([]byte(`{"Name": "Thieve", "Health": 10}`), &player)
	var hero Character
	switch Initial(s) {
	case "t":
		hero = Character{
			Name:            heroesList.Thieve,
			Alive:           true,
			CurrentLocation: []int{9, 4},
			Evasion:         30,
			Health:          rand.Intn(15) + 25,
			Skill:           0,
			Strength:        18,
			Crit:            25,
			Inventory:       map[string]*ItemQuantity{},
		}
		hero.addItemTypeToInventory(ItemIndexList[0], rand.Intn(3)+1)
		break
	case "p":
		hero = Character{
			Name:            heroesList.Paladin,
			Alive:           true,
			CurrentLocation: []int{9, 4},
			Evasion:         20,
			Health:          rand.Intn(30) + 30,
			Skill:           2,
			Strength:        25,
			Crit:            20,
			Inventory:       map[string]*ItemQuantity{},
		}
		hero.addItemTypeToInventory(ItemIndexList[0], rand.Intn(5)+1)
		break
	case "w":
		hero = Character{
			Name:            heroesList.Wizzard,
			Alive:           true,
			CurrentLocation: []int{9, 4},
			Evasion:         15,
			Health:          rand.Intn(25) + 25,
			Skill:           3,
			Strength:        15,
			Crit:            10,
			Inventory:       map[string]*ItemQuantity{},
		}
		hero.addItemTypeToInventory(ItemIndexList[0], rand.Intn(3)+1)
		hero.addItemTypeToInventory(ItemIndexList[1], rand.Intn(5)+1)
		break
	case "b":
		hero = Character{
			Name:            heroesList.Barbarian,
			Alive:           true,
			CurrentLocation: []int{9, 4},
			Evasion:         10,
			Health:          rand.Intn(30) + 40,
			Skill:           0,
			Strength:        30,
			Crit:            20,
			Inventory:       map[string]*ItemQuantity{},
		}
		hero.addItemTypeToInventory(ItemIndexList[0], rand.Intn(2)+1)
		break
	}
	hero.BaseHealth = hero.Health
	return &hero
}

func enemyNameByChances() string {
	anchor := rand.Intn(100) + 1
	A := enemiesList.GOBLIN
	B := 0

	for key, val := range EnemyChancesByName {
		if anchor >= B && val > B && val < anchor {
			A = key
			B = val
		}
	}
	return A
}

func makeNewEnemy() Character {
	name := enemyNameByChances()
	base := enemiesSpecificsValues[name]
	HP := base.Health()

	var enemy = Character{
		Health:     HP,
		BaseHealth: HP,
		Strength:   base.Strength,
		Evasion:    base.Evasion,
		Crit:       base.Crit,
		Name:       name,
		Alive:      true,
		Npc:        true,
		Inventory:  map[string]*ItemQuantity{},
	}

	return enemy
}

func enemyProbability(loc *Location) {
	if !loc.HasSeller {
		if rand.Intn(100) <= GameDifficulty[difficultyIndex[Difficulty]] {
			loc.HasEnemy = true
			loc.Enemy = makeNewEnemy()
		}
	}
}

func getRandPrice() int {
	return rand.Intn(15) + 5
}

func PercentChances(n int) bool {
	return rand.Intn(101) < n
}

func GetAPercentageOfB(a, b int) float32 {
	return float32(a * 100 / b)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ArrayToString(arr []string) string {
	return "[" + strings.Join(arr, " ") + "]"
}

func OffsetLegend(arr []string, offset, index int) string {
	if index >= offset && index < len(LegendArray)+offset {
		return arr[index-offset]
	}
	return ""
}

var itemsByEnemy = map[string]map[string]int{
	enemiesList.GOBLIN: map[string]int{
		itemNames.Moonstone: 0,
		itemNames.Doll:      0,
		itemNames.Coins:     0,
		itemNames.Key:       0,
		itemNames.Potion:    0,
		itemNames.Scroll:    0,
	},
	enemiesList.SKELETON: map[string]int{
		itemNames.Moonstone: 0,
		itemNames.Doll:      0,
		itemNames.Coins:     0,
		itemNames.Key:       0,
		itemNames.Potion:    0,
		itemNames.Scroll:    0,
	},
	enemiesList.ORC: map[string]int{
		itemNames.Moonstone: 0,
		itemNames.Doll:      0,
		itemNames.Coins:     0,
		itemNames.Key:       0,
		itemNames.Potion:    0,
		itemNames.Scroll:    0,
	},
	enemiesList.SORCERER: map[string]int{
		itemNames.Moonstone: 0,
		itemNames.Doll:      0,
		itemNames.Coins:     0,
		itemNames.Key:       0,
		itemNames.Potion:    0,
		itemNames.Scroll:    0,
	},
	enemiesList.DRAGON: map[string]int{
		itemNames.Moonstone: 0,
		itemNames.Doll:      0,
		itemNames.Coins:     0,
		itemNames.Key:       0,
		itemNames.Potion:    0,
		itemNames.Scroll:    0,
	},
}
