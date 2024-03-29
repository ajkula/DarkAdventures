package main

import "math/rand"

const releaseVersion = "v1.3.1"
const updatedAvailable = " New version available "
const DefaultLandscape = "ffffffffff\nffffddffff\nffdddddfff\nffdddlllff\nffddllllff\nfffllllfff\nfffffcclll\nfffllcclll\nffffllllll\nllllllllll"
const englishLang string = "en"
const frenchLang string = "fr"
const Tabulation int = 22
const EnemyChances int = 35
const Tab string = "\t"
const DoubleTab string = "\t\t"
const Unseen string = " "
const Shop string = "#"
const Root string = "л"
const YourPosition string = "○"
const LegendSpace int = 10
const inventorySpace int = 38
const heroesDetailsSpacing int = 14
const gaugeSize int = 40
const expChar string = "="
const emptyGauge string = " "
const dragonPos string = "V"
const nightWalkerPos string = "●"
const enemyPos string = "Д"
const NPCPosition string = "P"

var playerEnemyColor = map[bool]string{
	false: "green",
	true:  "red",
}
var EnemiesCount int = 1
var EnemiesKilled int = 0

var expByDifficulty map[int]int = map[int]int{
	0: 5,
	1: 15,
	2: 25,
}
var supportedLanguages = []string{englishLang, frenchLang}
var gameIntro string = translate(gameintroTR)
var NearORC string = translate(NearORCTR)

var rootBell = map[bool]string{
	false: translate(rootBellFALSE),
	true:  translate(rootBellTRUE),
}

var warpText string = translate(warpTextTR)

var heroesDetails = map[string]string{
	heroesList.Thief:     translate(heroesDetailsTHIEF),
	heroesList.Paladin:   translate(heroesDetailsPALADIN),
	heroesList.Wizard:    translate(heroesDetailsWIZARD),
	heroesList.Barbarian: translate(heroesDetailsBARBARIAN),
}

var heroesSkillDescription = map[string]map[string]string{
	heroesList.Thief:     ThiefSkill,
	heroesList.Paladin:   PaladinSkill,
	heroesList.Wizard:    WizardSkill,
	heroesList.Barbarian: BarbarianSkill,
}

type Statuses struct {
	Blight, Dark, Plague, Light, Fright string
}

var statuses = &Statuses{
	Blight: translate(blightStatusTR),
	Dark:   translate(darkStatusTR),
	Light:  translate(lightStatusTR),
	Plague: translate(plagueStatusTR),
	Fright: translate(frightStatusTR),
}

const CORPSE string = "CORPSE"

var specialAttackOnDragon map[string]map[string]int
var specialOnDragonByDifficultyMap = map[int]int{
	0: 5,
	1: 3,
	2: 0,
}

// var CEnemy = &Character{Name: CORPSE}
var NullifiedEnemy Character

var LegendArray []string = []string{
	CustomSpaceAlign(translate(youTR), LegendSpace) + YourPosition,
	CustomSpaceAlign(translate(shopsTR), LegendSpace) + Shop,
	CustomSpaceAlign(translate(rootsTR), LegendSpace) + Root,
	CustomSpaceAlign(roomTypes.FOREST+":", LegendSpace) + displayMapIcons[LettersFromLandscape[roomTypes.FOREST]],
	CustomSpaceAlign(roomTypes.PLAINS+":", LegendSpace) + displayMapIcons[LettersFromLandscape[roomTypes.PLAINS]],
	CustomSpaceAlign(roomTypes.DESERT+":", LegendSpace) + displayMapIcons[LettersFromLandscape[roomTypes.DESERT]],
	CustomSpaceAlign(roomTypes.CASTLE+":", LegendSpace) + displayMapIcons[LettersFromLandscape[roomTypes.CASTLE]],
}

type Directions struct{ North, South, East, West string }

var directions = Directions{
	North: translate(northTR),
	East:  translate(eastTR),
	West:  translate(westTR),
	South: translate(southTR),
}

type DifficultyNames struct{ Easy, Meddium, Hard string }

var difficultyNames = DifficultyNames{
	Easy:    translate(easyTR),
	Meddium: translate(meddiumTR),
	Hard:    translate(hardTR),
}

type Commands struct {
	Attack, Skill, Use, Escape, Go, Map, Buy, Inv, Stats, Yes, Help, Quit string
}

var commands = &Commands{
	Attack: "Attack",
	Skill:  "Skill",
	Use:    "Use",
	Escape: "Escape",
	Go:     "Go",
	Map:    "Map",
	Buy:    "Buy",
	Inv:    "Inv",
	Stats:  "Props",
	Yes:    "Yes",
	Help:   "Help",
	Quit:   "Quit",
}

var allCommands = []string{commands.Attack, commands.Use, commands.Go, commands.Buy, commands.Help, commands.Quit}
var battleCommands = []string{commands.Attack, commands.Use, commands.Escape, commands.Skill}
var worldCommands = []string{commands.Go, commands.Use}
var sellerCommands = []string{commands.Go, commands.Use, commands.Buy}
var npcCommands = []string{commands.Go, commands.Use, commands.Yes}
var universalCommands = []string{commands.Map, commands.Inv, commands.Stats, commands.Help, commands.Quit}
var difficultyIndex = map[int]string{0: difficultyNames.Easy, 1: difficultyNames.Meddium, 2: difficultyNames.Hard}
var GameDifficulty = map[string]int{difficultyNames.Easy: 15, difficultyNames.Meddium: 30, difficultyNames.Hard: 45}

type HeroesList struct{ Thief, Paladin, Wizard, Barbarian string }
type EnemiesList struct{ SKELETON, GOBLIN, SORCERER, ORC, DRAGON, NIGHTWALKER, NECROMANCER string }

var heroesList = HeroesList{
	Thief:     translate(ThiefNAME),
	Paladin:   translate(PaladinNAME),
	Wizard:    translate(WizardNAME),
	Barbarian: translate(BarbarianNAME),
}
var indexedHeroes = []string{heroesList.Thief, heroesList.Paladin, heroesList.Wizard, heroesList.Barbarian}

var indexedEnemiesForRandomization = []string{enemiesList.SKELETON, enemiesList.GOBLIN, enemiesList.SORCERER, enemiesList.ORC}
var enemiesList = EnemiesList{
	SKELETON:    translate(skeletonNAME),
	GOBLIN:      translate(goblinNAME),
	SORCERER:    translate(sorcererNAME),
	ORC:         translate(orcNAME),
	DRAGON:      translate(dragonNAME),
	NIGHTWALKER: translate(nightWalkerNAME),
	NECROMANCER: translate(necromancerNAME),
}

var EnemyChancesByName map[string]int = map[string]int{
	enemiesList.GOBLIN:   0,
	enemiesList.SKELETON: 40,
	enemiesList.SORCERER: 70,
	enemiesList.ORC:      92,
}

var ItemChancesByEnemyName map[string]map[string]int = map[string]map[string]int{
	enemiesList.SKELETON: {
		itemNames.Doll:      99,
		itemNames.Moonstone: 98,
		itemNames.Scroll:    94,
		itemNames.Potion:    92,
		itemNames.Key:       80,
		itemNames.Coins:     5,
	},
	enemiesList.GOBLIN: {
		itemNames.Doll:      99,
		itemNames.Moonstone: 98,
		itemNames.Scroll:    96,
		itemNames.Potion:    70,
		itemNames.Key:       60,
		itemNames.Coins:     1,
	},
	enemiesList.SORCERER: {
		itemNames.Doll:      99,
		itemNames.Moonstone: 98,
		itemNames.Potion:    92,
		itemNames.Key:       80,
		itemNames.Scroll:    50,
		itemNames.Coins:     10,
	},
	enemiesList.ORC: {
		itemNames.Doll:      99,
		itemNames.Scroll:    98,
		itemNames.Potion:    94,
		itemNames.Key:       90,
		itemNames.Moonstone: 50,
		itemNames.Coins:     1,
	},
	enemiesList.DRAGON: {
		itemNames.Doll:      60,
		itemNames.Scroll:    55,
		itemNames.Potion:    50,
		itemNames.Key:       45,
		itemNames.Moonstone: 5,
		itemNames.Coins:     1,
	},
	enemiesList.NECROMANCER: {
		itemNames.Doll:      50,
		itemNames.Scroll:    25,
		itemNames.Potion:    15,
		itemNames.Key:       10,
		itemNames.Moonstone: 5,
		itemNames.Coins:     1,
	},
	enemiesList.NIGHTWALKER: {
		itemNames.Doll:      35,
		itemNames.Scroll:    20,
		itemNames.Potion:    15,
		itemNames.Key:       10,
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
	ExpValue int
	Exp      int
}

var enemiesSpecificsValues = map[string]Specifics{
	enemiesList.SKELETON: {
		Health:   func() int { return rand.Intn(15) + 10 },
		Strength: 20,
		Evasion:  5,
		Crit:     35,
		ExpValue: 3,
		Exp:      rand.Intn(15) + rand.Intn(10) + expByDifficulty[Difficulty],
	},
	enemiesList.GOBLIN: {
		Health:   func() int { return rand.Intn(20) + 10 },
		Strength: 15,
		Evasion:  15,
		Crit:     10,
		ExpValue: 5,
		Exp:      rand.Intn(15) + rand.Intn(15) + expByDifficulty[Difficulty],
	},
	enemiesList.SORCERER: {
		Health:   func() int { return rand.Intn(30) + 10 },
		Strength: 15,
		Evasion:  10,
		Skill:    2,
		Crit:     15,
		ExpValue: 4,
		Exp:      rand.Intn(15) + rand.Intn(15) + expByDifficulty[Difficulty],
	},
	enemiesList.ORC: {
		Health:   func() int { return rand.Intn(25) + 35 },
		Strength: 18,
		Evasion:  3,
		Crit:     25,
		ExpValue: 10,
		Exp:      rand.Intn(15) + rand.Intn(5) + expByDifficulty[Difficulty],
	},
}

var walkerProximity = map[string]map[string]string{
	enemiesList.DRAGON: {
		LettersFromLandscape[roomTypes.FOREST]: translate(forestTR),
		LettersFromLandscape[roomTypes.PLAINS]: translate(landTR),
		LettersFromLandscape[roomTypes.DESERT]: translate(desertTR),
		LettersFromLandscape[roomTypes.CASTLE]: translate(castleTR),
		"x":                                    translate(xTR),
	},
	enemiesList.NIGHTWALKER: {
		LettersFromLandscape[roomTypes.FOREST]: translate(nwProxTR),
		LettersFromLandscape[roomTypes.PLAINS]: translate(nwProxTR),
		LettersFromLandscape[roomTypes.DESERT]: translate(nwProxTR),
		LettersFromLandscape[roomTypes.CASTLE]: translate(nwProxTR),
		"x":                                    translate(nwEncxTR),
	},
}

// You are
var introPlains = map[int]string{
	0: translate(introPlainsTR0),
	1: translate(introPlainsTR1),
	2: translate(introPlainsTR2),
}

var introDesert = map[int]string{
	0: translate(introDesertTR0),
	1: translate(introDesertTR1),
	2: translate(introDesertTR2),
}

var introCastle = map[int]string{
	0: translate(introCastleTR0),
	1: translate(introCastleTR1),
	2: translate(introCastleTR2),
}

var introForest = map[int]string{
	0: translate(introForestTR0),
	1: translate(introForestTR1),
	2: translate(introForestTR2),
}

type RoomTypes struct {
	FOREST string
	PLAINS string
	DESERT string
	CASTLE string
}

var roomTypes = &RoomTypes{
	FOREST: translate(forestNameTR),
	PLAINS: translate(plainsNameTR),
	DESERT: translate(desertNameTR),
	CASTLE: translate(castleNameTR),
}

type GridLetters struct {
	Forest, Plains, Desert, Castle string
}

var gridLetters = &GridLetters{
	Castle: "c",
	Desert: "d",
	Forest: "f",
	Plains: "l",
}

const people = "𐒕" // ᡶ ᡷ ᢜ ᢤ ♾ 𐒉 𐒕

// ICI
var displayMapIcons = map[string]string{
	gridLetters.Forest: "ф", //🌳 🌴 🌲 ⺦ ф Ф
	gridLetters.Plains: "_", // ⏚ 🌬 ◌ ◒ ◛ ◡ ⚆ ㊀
	gridLetters.Desert: "^", //🏜 🏝
	gridLetters.Castle: "H", // 🏯 🏰 ⛫
}

var RoomFromLandscape = map[string]map[int]string{
	gridLetters.Forest: introForest,
	gridLetters.Plains: introPlains,
	gridLetters.Desert: introDesert,
	gridLetters.Castle: introCastle,
}
var LettersFromLandscape = map[string]string{
	roomTypes.FOREST: gridLetters.Forest,
	roomTypes.PLAINS: gridLetters.Plains,
	roomTypes.DESERT: gridLetters.Desert,
	roomTypes.CASTLE: gridLetters.Castle,
}

var Ambiance = map[int]string{
	0: translate(AmbianceTR0),
	1: translate(AmbianceTR1),
	2: translate(AmbianceTR2),
	3: translate(AmbianceTR3),
	4: translate(AmbianceTR4),
	5: translate(AmbianceTR5),
}

// There is
var SellerList = map[int]string{
	0: translate(SellerListTR0),
	1: translate(SellerListTR1),
	2: translate(SellerListTR2),
}

var RoomTypeList = map[string]string{
	gridLetters.Forest: roomTypes.FOREST,
	gridLetters.Desert: roomTypes.DESERT,
	gridLetters.Plains: roomTypes.PLAINS,
	gridLetters.Castle: roomTypes.CASTLE,
}

type EventNames struct {
	CHEST  string
	ENEMY  string
	SELLER string
}

var eventNames *EventNames = &EventNames{
	CHEST:  translate(chestEventNameTR),
	ENEMY:  translate(enemyEventNameTR),
	SELLER: translate(sellerEventNameTR),
}

var Event = map[int]string{0: eventNames.CHEST, 1: eventNames.ENEMY, 2: eventNames.SELLER}

type ItemNames struct{ Potion, Scroll, Doll, Key, Moonstone, Coins string }

var itemNames = &ItemNames{
	Potion:    translate(potionNameTR),
	Scroll:    translate(scrollNameTR),
	Doll:      translate(dollNameTR),
	Key:       translate(keyNameTR),
	Moonstone: translate(moonstoneNameTR),
	Coins:     translate(coinsNameTR),
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
		Description: translate(DollTR),
		Effect:      30,
	},
	itemNames.Moonstone: {
		Name:        itemNames.Moonstone,
		Description: translate(MoonstoneTR),
		Effect:      5,
	},
	itemNames.Scroll: {
		Name:        itemNames.Scroll,
		Description: translate(ScrollTR),
		Effect:      20,
	},
	itemNames.Potion: {
		Name:        itemNames.Potion,
		Description: translate(PotionTR) + " 20 HP",
		Effect:      20,
	},
	itemNames.Key: {
		Name:        itemNames.Key,
		Description: translate(KeyTR),
		Effect:      1,
	},
	itemNames.Coins: {
		Name:        itemNames.Coins,
		Description: translate(CoinsTR),
		Effect:      1,
	},
}

type EscapeResults struct{ OK, RAND, KO string }

var escapeResults = &EscapeResults{OK: "ok", RAND: "rand", KO: "ko"}
var escapeCases map[string]map[string]string = map[string]map[string]string{
	heroesList.Thief: {
		escapeResults.OK:   translate(ThiefEscapeOK),
		escapeResults.RAND: translate(ThiefEscapeRAND),
		escapeResults.KO:   translate(ThiefEscapeKO),
	},
	heroesList.Paladin: {
		escapeResults.OK:   translate(PaladinEscapeOK),
		escapeResults.RAND: translate(PaladinEscapeRAND),
		escapeResults.KO:   translate(PaladinEscapeKO),
	},
	heroesList.Wizard: {
		escapeResults.OK:   translate(WizardEscapeOK),
		escapeResults.RAND: translate(WizardEscapeRAND),
		escapeResults.KO:   translate(WizardEscapeKO),
	},
	heroesList.Barbarian: {
		escapeResults.OK:   translate(BarbarianEscapeOK),
		escapeResults.RAND: translate(BarbarianEscapeRAND),
		escapeResults.KO:   translate(BarbarianEscapeKO),
	},
}
