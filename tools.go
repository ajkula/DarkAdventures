package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	lang "github.com/cloudfoundry/jibber_jabber"
)

var WorldMap [][]*Location

func getSysLang() (language string) {
	language = englishLang
	if lang, err := lang.DetectLanguage(); err == nil {
		i := indexOf(supportedLanguages, lang)
		if i != -1 {
			language = supportedLanguages[i]
		}
	}
	return language
}

// WhereCanYouGo location => []string
func WhereCanYouGo(yourPlace *Location) []string {
	var youCanGo []string

	if yourPlace.X >= 1 {
		youCanGo = append(youCanGo, directions.West)
	}
	if yourPlace.X <= X-2 {
		youCanGo = append(youCanGo, directions.East)
	}

	if yourPlace.Y >= 1 {
		youCanGo = append(youCanGo, directions.North)
	}
	if yourPlace.Y <= Y-2 {
		youCanGo = append(youCanGo, directions.South)
	}
	return youCanGo
}

func Intro() {
	Output("blue", gameIntro)
	Output("blue", translate(easyTR))
	Output("blue", translate(meddiumTR))
	Output("blue", translate(hardTR))
	UserInput(&Difficulty)
	if Difficulty > 3 || Difficulty < 1 {
		Intro()
	} else {
		Difficulty = Difficulty - 1
	}
}

func ChooseHero() {
	Output("blue", translate(chooseHeroTR))
	for index, name := range indexedHeroes {
		Output("blue", Tab+CustomSpaceAlign(strconv.Itoa(index+1)+" - "+name, heroesDetailsSpacing)+heroesDetails[name])
	}
	UserInput(&Hero)
	if Hero > 4 || Hero < 1 {
		ChooseHero()
	} else {
		Hero = Hero - 1
	}
}

var aggregateItems = map[string]int{
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

// ICI
func AnalyzeItemsRepartition() {
	for e, items := range itemsByEnemy {
		Output("green", DoubleTab+e+":")
		for in, iqty := range items {
			if iqty > 0 {
				Output("yellow", Tab+CustomSpaceAlign(in, 22-len(strconv.Itoa(iqty))), iqty)
			}
		}
	}
}

func getRandomArrayString(arr []string) string {
	return arr[rand.Intn(len(arr))]
}

func makeWorldMapSizes(Y, X int) [][]*Location {
	// fmt.Println(Y, X) // ICI
	var w [][]*Location
	total := 0
	var tempo []*Location
	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			hasItem := false
			if rand.Intn(100) < 20 {
				hasItem = true
			}

			item := getItem(hasItem)
			tempo = append(tempo, &Location{
				Description: translate(youAreTR) + RoomFromLandscape[Grid[y][x]][rand.Intn(len(RoomFromLandscape[Grid[y][x]]))] + Tab + Ambiance[rand.Intn(len(Ambiance))],
				HasSeller:   hasItem,
				Item:        item,
				X:           x,
				Y:           y,
			})
		}
		total += X
		w = append(w, tempo[total-X:total])
	}
	return w
}

func CreateMap() [][]*Location {
	WorldMap = makeWorldMapSizes(Y, X)
	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			if WorldMap[y][x].HasSeller {
				WorldMap[y][x].Seller = translate(HasEnemyOrSellerTR0) + getSeller()
			}
			WorldMap[y][x].CanGoTo = WhereCanYouGo(WorldMap[y][x])
			if !((y == Y-1) && (x == ((X / 2) - (X % 2)))) {
				enemyProbability(WorldMap[y][x])
			}
			if WorldMap[y][x].HasEnemy {
				EnemiesCount++
				WorldMap[y][x].Enemy.createEnemyInventory()
			}
			populateTheSlices(WorldMap[y][x])
		}
	}
	for b := 0; b < Y; b++ {
		for a := 0; a < X; a++ {
			addOrcProximity(WorldMap[b][a])
		}
	}
	// ICI
	// showGameItems()
	// showGameEnemies()
	// AnalyzeItemsRepartition()

	Output("white", translate(difficultyTR), difficultyIndex[Difficulty], "\n")
	return WorldMap
}

func makeEnemiesLVL(enemi *Character) {
	enemi.calcLVL()
	enemi.Health = enemi.BaseHealth
}

func addOrcProximity(loc *Location) {
	if loc.HasEnemy {
		if loc.Enemy.Name == enemiesList.ORC {
			if near := strings.Contains(loc.Description, translate(nearOrcSearchTR)); !near {
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
// Output("white", Tab+CustomSpaceAlign("TOTAL SCORE:", 22-len(total))+total)

func showGameItems() {
	fmt.Println(DoubleTab + CalculateSpaceAlign("Items:"))
	for a, b := range aggregateItems {
		fmt.Println(Tab+CustomSpaceAlign(a, 22-len(strconv.Itoa(b))), b)
	}
}

func showGameEnemies() {
	fmt.Println(DoubleTab + CalculateSpaceAlign("Enemies:"))
	for i := 0; i < Y; i++ {
		for j := 0; j < X; j++ {
			if WorldMap[i][j].HasEnemy {
				aggregateEnemies[WorldMap[i][j].Enemy.Name]++
				totalEnemiesMinusQuestsObject[WorldMap[i][j].Enemy.Name]++
			}
		}
	}
	for a, b := range aggregateEnemies {
		fmt.Println(Tab+CustomSpaceAlign(a, 22-len(strconv.Itoa(b))), b)
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

func getSeller() string {
	var str string
	str = SellerList[rand.Intn(3)]
	return str
}

func Initial(s string) string {
	if s == "" {
		return ""
	}
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
	if item != "" {
		for _, elem := range arr {
			if Initial(elem) == Initial(item) {
				boolean = true
			}
		}
	}
	return boolean
}

var heroFromName = func(s string) *Character {
	// json.Unmarshal([]byte(`{"Name": "Thief", "Health": 10}`), &player)
	var hero Character
	switch Initial(s) {
	case Initial(heroesList.Thief):
		hero = Character{
			Name:            heroesList.Thief,
			Alive:           true,
			CurrentLocation: []int{Y - 1, (X / 2) - (X % 2)},
			Evasion:         30,
			Health:          rand.Intn(15) + 25,
			Skill:           0,
			Strength:        18,
			Crit:            30,
			ExpValue:        8,
			Inventory:       map[string]*ItemQuantity{},
			LevelUp: &Leveling{
				NextRank:            5,
				NextBase:            5,
				Exp:                 0,
				achievedLevelsChain: []int{},
				Rates: &Specifics{
					Health: func() int {
						return rand.Intn(4) + 1
					},
					Crit:     1,
					Evasion:  1,
					Skill:    1,
					Strength: 1,
				},
			},
		}
		hero.addItemTypeToInventory(ItemIndexList[0], rand.Intn(3)+1)
		break
	case Initial(heroesList.Paladin):
		hero = Character{
			Name:            heroesList.Paladin,
			Alive:           true,
			CurrentLocation: []int{Y - 1, (X / 2) - (X % 2)},
			Evasion:         20,
			Health:          rand.Intn(30) + 30,
			Skill:           2,
			Strength:        25,
			Crit:            20,
			ExpValue:        12,
			Inventory:       map[string]*ItemQuantity{},
			LevelUp: &Leveling{
				NextRank:            5,
				NextBase:            5,
				Exp:                 0,
				achievedLevelsChain: []int{},
				Rates: &Specifics{
					Health: func() int {
						return rand.Intn(3) + 1
					},
					Crit:     2,
					Evasion:  1,
					Skill:    1,
					Strength: 1,
				},
			},
		}
		hero.addItemTypeToInventory(ItemIndexList[0], rand.Intn(3)+3)
		break
	case Initial(heroesList.Wizard):
		hero = Character{
			Name:            heroesList.Wizard,
			Alive:           true,
			CurrentLocation: []int{Y - 1, (X / 2) - (X % 2)},
			Evasion:         15,
			Health:          rand.Intn(15) + rand.Intn(10) + 25,
			Skill:           3,
			Strength:        20,
			Crit:            15,
			ExpValue:        10,
			Inventory:       map[string]*ItemQuantity{},
			LevelUp: &Leveling{
				NextRank:            5,
				NextBase:            5,
				Exp:                 0,
				achievedLevelsChain: []int{},
				Rates: &Specifics{
					Health: func() int {
						return rand.Intn(5) + 1
					},
					Crit:     2,
					Evasion:  2,
					Skill:    1,
					Strength: 1,
				},
			},
		}
		hero.addItemTypeToInventory(ItemIndexList[0], rand.Intn(3)+2)
		hero.addItemTypeToInventory(ItemIndexList[1], rand.Intn(4)+2)
		break
	case Initial(heroesList.Barbarian):
		hero = Character{
			Name:            heroesList.Barbarian,
			Alive:           true,
			CurrentLocation: []int{Y - 1, (X / 2) - (X % 2)},
			Evasion:         10,
			Health:          rand.Intn(30) + 40,
			Skill:           0,
			Strength:        30,
			Crit:            20,
			ExpValue:        12,
			Inventory:       map[string]*ItemQuantity{},
			LevelUp: &Leveling{
				NextRank:            5,
				NextBase:            5,
				Exp:                 0,
				achievedLevelsChain: []int{},
				Rates: &Specifics{
					Health: func() int {
						return rand.Intn(4) + 1
					},
					Crit:     1,
					Evasion:  2,
					Skill:    2,
					Strength: 1,
				},
			},
		}
		hero.addItemTypeToInventory(ItemIndexList[0], rand.Intn(2)+1)
		break
	}
	hero.BaseHealth = hero.Health
	hero.setImage()
	hero.StatusEffects = &StatusEffectsBlueprint{
		AllStatus: []*Blueprint{},
	}
	hero.Icon = YourPosition
	SCORE.removeBaseInventory(hero.Inventory)
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

func makeNewEnemy() *Character {
	name := enemyNameByChances()
	base := enemiesSpecificsValues[name]
	HP := base.Health()

	var enemy = Character{
		Health:     HP,
		BaseHealth: HP,
		Strength:   base.Strength,
		Evasion:    base.Evasion,
		Crit:       base.Crit,
		ExpValue:   base.ExpValue,
		Skill:      rand.Intn(3),
		Name:       name,
		Alive:      true,
		Npc:        true,
		Icon:       enemyPos,
		Inventory:  map[string]*ItemQuantity{},
		LevelUp: &Leveling{
			NextRank:            5,
			NextBase:            5,
			Exp:                 base.Exp,
			achievedLevelsChain: []int{},
			Rates: &Specifics{
				Health: func() int {
					return rand.Intn(2) + 1
				},
				Crit:     1,
				Evasion:  1,
				Skill:    1,
				Strength: 1,
			},
		},
		StatusEffects: &StatusEffectsBlueprint{
			AllStatus: []*Blueprint{},
		},
	}
	enemy.setImage()
	return &enemy
}

func enemyProbability(loc *Location) {
	if !loc.HasSeller {
		if rand.Intn(100) <= GameDifficulty[difficultyIndex[Difficulty]] {
			loc.HasEnemy = true
			loc.Enemy = makeNewEnemy()
			pile.PushCharacters(loc.Enemy)
		}
	}
}

func getRandPrice() int {
	return rand.Intn(15) + 5
}

func PercentChances(n int) bool {
	r := (rand.Intn(100) + 1)
	// fmt.Println("percentChances: ", r, n)
	return r < n
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
	enemiesList.NIGHTWALKER: map[string]int{
		itemNames.Moonstone: 0,
		itemNames.Doll:      0,
		itemNames.Coins:     0,
		itemNames.Key:       0,
		itemNames.Potion:    0,
		itemNames.Scroll:    0,
	},
	enemiesList.NECROMANCER: map[string]int{
		itemNames.Moonstone: 0,
		itemNames.Doll:      0,
		itemNames.Coins:     0,
		itemNames.Key:       0,
		itemNames.Potion:    0,
		itemNames.Scroll:    0,
	},
}
