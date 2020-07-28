package main

import "math/rand"

const DefaultLandscape = "ffffffffff\nffffddffff\nffdddddfff\nffdddlllff\nffddllllff\nfffllllfff\nfffffcclll\nfffllcclll\nffffllllll\nllllllllll"
const gameIntro = "\tWelcome to Dark Adventures\n\tSelect a difficulty:"
const Tabulation int = 22
const EnemyChances int = 35
const Tab string = "\t"
const DoubleTab string = "\t\t"
const Unseen string = " "
const Shop string = "#"
const YourPosition string = "@"
const LegendSpace int = 8
const inventorySpace int = 38
const heroesDetailsSpacing int = 14
const NearORC string = "There's an odd smell in this place...\n"

var heroesDetails = map[string]string{
	heroesList.Thieve:    " - High evasion, crits, can attack x2",
	heroesList.Paladin:   " - High strength, health, more potions",
	heroesList.Wizzard:   " - Potions and Scrolls from the start",
	heroesList.Barbarian: " - High health, strength",
}

var NullifiedEnemy Character

var LegendArray []string = []string{
	CustomSpaceAlign("You:", LegendSpace) + YourPosition,
	CustomSpaceAlign("Shops:", LegendSpace) + Shop,
	CustomSpaceAlign("Forest:", LegendSpace) + LettersFromLandscape["forest"],
	CustomSpaceAlign("Plains:", LegendSpace) + LettersFromLandscape["plains"],
	CustomSpaceAlign("Desert:", LegendSpace) + LettersFromLandscape["desert"],
	CustomSpaceAlign("Castle:", LegendSpace) + LettersFromLandscape["castle"],
}

type Directions struct{ North, South, East, West string }

var directions = Directions{
	North: "North",
	East:  "East",
	West:  "West",
	South: "South",
}

type DifficultyNames struct{ Easy, Meddium, Hard string }

var difficultyNames = DifficultyNames{
	Easy:    "Easy",
	Meddium: "Meddium",
	Hard:    "Hard",
}

type Commands struct {
	Attack, Use, Go, Map, Buy, Inv, Help, Quit string
}

var commands = &Commands{
	Attack: "Attack",
	Use:    "Use",
	Go:     "Go",
	Map:    "Map",
	Buy:    "Buy",
	Inv:    "Inv",
	Help:   "Help",
	Quit:   "Quit",
}

var allCommands = []string{commands.Attack, commands.Use, commands.Go, commands.Buy, commands.Help, commands.Quit}
var battleCommands = []string{commands.Attack, commands.Use}
var worldCommands = []string{commands.Go, commands.Use}
var sellerCommands = []string{commands.Go, commands.Use, commands.Buy}
var universalCommands = []string{commands.Map, commands.Inv, commands.Help, commands.Quit}
var difficultyIndex = map[int]string{0: difficultyNames.Easy, 1: difficultyNames.Meddium, 2: difficultyNames.Hard}
var GameDifficulty = map[string]int{difficultyNames.Easy: 15, difficultyNames.Meddium: 30, difficultyNames.Hard: 45}

type HeroesList struct{ Thieve, Paladin, Wizzard, Barbarian string }
type EnemiesList struct{ SKELETON, GOBLIN, SORCERER, ORC, DRAGON string }

var indexedHeroes = []string{heroesList.Thieve, heroesList.Paladin, heroesList.Wizzard, heroesList.Barbarian}
var heroesList = HeroesList{
	Thieve:    "Thieve",
	Paladin:   "Paladin",
	Wizzard:   "Wizzard",
	Barbarian: "Barbarian",
}

var indexedEnemiesForRandomization = []string{enemiesList.SKELETON, enemiesList.GOBLIN, enemiesList.SORCERER, enemiesList.ORC}
var enemiesList = EnemiesList{
	SKELETON: "SKELETON",
	GOBLIN:   "GOBLIN",
	SORCERER: "SORCERER",
	ORC:      "ORC",
	DRAGON:   "DRAGON",
}

// nnnnnnnnnif (i > 98) return "doll";
//     else if (i > 96) return "moonstone";
//     else if (i > 94) return "scroll";
//     else if (i > 92) return "potions";
//     else if (i > 90) return "key";
//     nnnnnnnnnnnnelse return "coins"

var EnemyChancesByName map[string]int = map[string]int{
	enemiesList.GOBLIN:   0,
	enemiesList.SKELETON: 40,
	enemiesList.SORCERER: 70,
	enemiesList.ORC:      92,
}

var ItemChancesByEnemyName map[string]map[string]int = map[string]map[string]int{
	enemiesList.SKELETON: map[string]int{
		itemNames.Doll:      99,
		itemNames.Moonstone: 98,
		itemNames.Scroll:    94,
		itemNames.Potion:    92,
		itemNames.Key:       80,
		itemNames.Coins:     5,
	},
	enemiesList.GOBLIN: map[string]int{
		itemNames.Doll:      99,
		itemNames.Moonstone: 98,
		itemNames.Scroll:    96,
		itemNames.Potion:    70,
		itemNames.Key:       60,
		itemNames.Coins:     1,
	},
	enemiesList.SORCERER: map[string]int{
		itemNames.Doll:      99,
		itemNames.Moonstone: 98,
		itemNames.Potion:    92,
		itemNames.Key:       80,
		itemNames.Scroll:    50,
		itemNames.Coins:     10,
	},
	enemiesList.ORC: map[string]int{
		itemNames.Doll:      99,
		itemNames.Scroll:    98,
		itemNames.Potion:    94,
		itemNames.Key:       90,
		itemNames.Moonstone: 50,
		itemNames.Coins:     1,
	},
	enemiesList.DRAGON: map[string]int{
		itemNames.Doll:      60,
		itemNames.Scroll:    55,
		itemNames.Potion:    50,
		itemNames.Key:       45,
		itemNames.Moonstone: 5,
		itemNames.Coins:     1,
	},
}

type Specifics struct {
	Health   func() int
	Strength int
	Evasion  int
	Skill    int
	Crit     int
}

var enemiesSpecificsValues = map[string]Specifics{
	enemiesList.SKELETON: {
		Health:   func() int { return rand.Intn(15) + 10 },
		Strength: 20,
		Evasion:  5,
		Crit:     35,
	},
	enemiesList.GOBLIN: {
		Health:   func() int { return rand.Intn(20) + 10 },
		Strength: 15,
		Evasion:  15,
		Crit:     10,
	},
	enemiesList.SORCERER: {
		Health:   func() int { return rand.Intn(30) + 10 },
		Strength: 15,
		Evasion:  10,
		Skill:    2,
		Crit:     15,
	},
	enemiesList.ORC: {
		Health:   func() int { return rand.Intn(25) + 35 },
		Strength: 25,
		Evasion:  3,
		Crit:     25,
	},
}

var dragonProximity = map[string]string{
	"f": "Trees are burned the soil is ash...",
	"l": "The air carries ashes flying in the wind...",
	"d": "It's hotter than usual and so dry...",
	"c": "It smells like burning from all directions...",
	"x": "NO LUCK, A strong wind bursts all around the place,\n" +
		Tab + "The sunlight dims before you heard the loudest of noises\n" +
		Tab + "Humongous, wings deployed its scream tearing the sky,\n" +
		Tab + "Here it is. The mightiest of all foes...",
}

// You are
var introPlains = map[int]string{
	0: " in an old foggy village, there's no soul here,\n",
	1: " in the heath, you hear a weird music,\n" + Tab + "let's not waste any time here,\n",
	2: " on a long road between green hills and a river,\n",
}

var introDesert = map[int]string{
	0: " in the wasteland, everything is dead and dry here,\n",
	1: " on the swamp, nauseous and poisonous,\n" + Tab + "something is lurking here,\n",
	2: " in the middle of dust..\n" + Tab + "of a long gone empire and a storm is at the horizon,\n",
}

var introCastle = map[int]string{
	0: " in front of a castle ruin's gate, it barely stands,\n",
	1: " at an old fort or what might have been one long ago,\n",
	2: " below a huge tower, on top of which float an old flag,\n",
}

var introForest = map[int]string{
	0: " near a forest, the trees seem to move by their own will,\n",
	1: " unfortunately at the edge of the thorns wood,\n" + Tab + "no one comes back from it,\n",
	2: " in a part of the forest all trees are rotten\n" + Tab + "and covered by poisonous mushrooms,\n",
}

var RoomFromLandscape = map[string]map[int]string{
	"f": introForest,
	"l": introPlains,
	"d": introDesert,
	"c": introCastle,
}
var LettersFromLandscape = map[string]string{
	"forest": "f",
	"plains": "l",
	"desert": "d",
	"castle": "c",
}

var Ambiance = map[int]string{
	0: "there are nobody around, only the wind.\n",
	1: "it's getting dark and you can see shadows moving..\n",
	2: "all is silent, there's not even wind!\n",
	3: "you don't feel safe but have to keep going on..\n",
	4: "suddenly you feel shivers, a noise, voice or wind?\n",
	5: "many noises around you, but can't see anyone...\n",
}

// There is
var SellerList = map[int]string{
	0: " a dwarf, with a bag full of goods\n",
	1: " an elf, he holds something in his hand\n",
	2: " a troll, he drops something in front of you\n",
}

var RoomTypeList = map[string]string{
	"f": "forest",
	"d": "desert",
	"l": "plains",
	"c": "castle",
}

var Event = map[int]string{0: "chest", 1: "enemy", 2: "seller"}

type ItemNames struct{ Potion, Scroll, Doll, Key, Moonstone, Coins string }

var itemNames = &ItemNames{
	Potion:    "potion",
	Scroll:    "scroll",
	Doll:      "doll",
	Key:       "key",
	Moonstone: "moonstone",
	Coins:     "coins",
}

var ItemIndexList = map[int]string{
	0: itemNames.Potion,
	1: itemNames.Scroll,
	2: itemNames.Doll,
	3: itemNames.Key,
	4: itemNames.Moonstone,
	5: itemNames.Coins,
}

var UsableItems = map[string]bool{
	itemNames.Potion:    true,
	itemNames.Scroll:    true,
	itemNames.Key:       true,
	itemNames.Doll:      false,
	itemNames.Moonstone: true,
}

var ItemList = map[string]*Item{
	itemNames.Doll: {
		Name:        itemNames.Doll,
		Description: "Will revive you with 30 HP",
		Effect:      30,
	},
	itemNames.Moonstone: {
		Name:        itemNames.Moonstone,
		Description: "Increase your dammage by 5",
		Effect:      5,
	},
	itemNames.Scroll: {
		Name:        itemNames.Scroll,
		Description: "20 Dammage to one enemy",
		Effect:      20,
	},
	itemNames.Potion: {
		Name:        itemNames.Potion,
		Description: "Heal 20 HP",
		Effect:      20,
	},
	itemNames.Key: {
		Name:        itemNames.Key,
		Description: "To open locks, chests",
		Effect:      1,
	},
	itemNames.Coins: {
		Name:        itemNames.Coins,
		Description: "Golden coins",
		Effect:      1,
	},
}
