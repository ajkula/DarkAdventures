package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type PlayerQuestsMap struct {
	activeQuestsByID    map[string]*Quest
	activeQuestByTarget map[string]map[string]*Quest
	activeQuestByType   map[string][]*Quest
}

var playerQuestsMap = &PlayerQuestsMap{
	activeQuestsByID:    map[string]*Quest{},
	activeQuestByTarget: map[string]map[string]*Quest{},
	activeQuestByType:   map[string][]*Quest{},
}

func (pqm *PlayerQuestsMap) addQuest(q *Quest) {
	pqm.activeQuestByTarget[q.Condition.Target] = map[string]*Quest{}
	pqm.activeQuestByTarget[q.Condition.Target][q.ID] = q
	pqm.activeQuestsByID[q.ID] = q
	pqm.activeQuestByType[q.QuestType] = append(pqm.activeQuestByType[q.QuestType], q)
}

var questProbabilityPercentage = 5

type QuestTypes struct{ KILL, RETRIEVE, SAVE string }

var questTypes = &QuestTypes{KILL: translate(killTR), RETRIEVE: translate(retrieveTR), SAVE: translate(saveTR)}
var questTypesList = []string{questTypes.KILL, questTypes.RETRIEVE, questTypes.SAVE}

var noEnemyRooms = []*Location{}
var enemyRooms = []*Location{}

func populateTheSlices(loc *Location) {
	notThisOne := WorldMap[Y-1][(X/2)-(X%2)]
	if loc.HasEnemy {
		enemyRooms = append(enemyRooms, loc)
	}
	if !loc.HasEnemy && !loc.HasSeller && loc != notThisOne {
		noEnemyRooms = append(noEnemyRooms, loc)
	}
}

var conditions = map[string][]*Condition{}

// QuestTypes
var baseConditionsMap = map[string][]*Condition{
	questTypes.KILL: {
		&Condition{Target: enemiesList.GOBLIN, Quantity: 4, ExpValue: 15},
		&Condition{Target: enemiesList.SKELETON, Quantity: 8, ExpValue: 25},
		&Condition{Target: enemiesList.ORC, Quantity: 1, ExpValue: 20},
		&Condition{Target: enemiesList.SORCERER, Quantity: 2, ExpValue: 15},
		&Condition{Target: enemiesList.NIGHTWALKER, Quantity: 2, ExpValue: 30},
		&Condition{Target: enemiesList.DRAGON, Quantity: 1, ExpValue: 30},
	},
	questTypes.SAVE: {
		&Condition{Target: translate(RescueFriendTR), Quantity: 1, ExpValue: 10},
		&Condition{Target: translate(RescueShamanTR), Quantity: 1, ExpValue: 10},
		&Condition{Target: translate(RescueVillagerTR), Quantity: 2, ExpValue: 10},
	},
	questTypes.RETRIEVE: {
		&Condition{Target: itemNames.Scroll, Quantity: 1, ExpValue: 13},
		&Condition{Target: itemNames.Moonstone, Quantity: 1, ExpValue: 19},
		&Condition{Target: itemNames.Potion, Quantity: 1, ExpValue: 10},
	},
}

type Quest struct {
	ID        string
	X, Y      int
	QuestType string
	Condition *Condition
	Active    bool
	Resolved  bool
	Rewarded  bool
	Dialogs   *QuestDialogs
}

type QuestDialogs struct {
	Greetings string
	Request   string
	Answer    string
	Accepted  string
}

func createQuest() *Quest {
	condition, name := getRandomCondition()
	id := makeID()
	quest := &Quest{
		ID:        id,
		QuestType: name,
		Condition: condition,
		Active:    false,
		Resolved:  false,
		Rewarded:  false,
		Dialogs: &QuestDialogs{
			Greetings: "",
			Request:   "",
			Accepted:  "",
			Answer:    "",
		},
	}
	return quest
}

type Condition struct {
	Target   string
	Quantity int
	ExpValue int
}

func dialogFromConditions(q *Quest) (str string) {
	if !q.Active {
		str = q.Dialogs.Greetings + q.Dialogs.Request
	}
	if q.Active && !q.Resolved && !q.Rewarded {
		str = q.Dialogs.Accepted
	}
	if q.Active && q.Resolved && !q.Rewarded {
		str = q.Dialogs.Answer
		q.Rewarded = true
	}
	return str
}

func getRandomCondition() (*Condition, string) {
	name := getRandomArrayString(questTypesList)
	arr := conditions[name]
	c := arr[rand.Intn(len(arr))]
	if indexOfConditions(selectedQuestConditions, c) {
		return getRandomCondition()
	}
	selectedQuestConditions = append(selectedQuestConditions, c)
	return c, name
}

func validateAvailableConditions() {
	for _, questType := range questTypesList {
		for _, c := range baseConditionsMap[questType] {
			if checkCondition(questType, c) {
				conditions[questType] = append(conditions[questType], c)
			}
			if c.Target == enemiesList.NIGHTWALKER {
				conditions[questType] = append(conditions[questType], c)
			}
		}
	}
}

var totalItemsMinusQuestsObject = map[string]int{
	itemNames.Moonstone: 0,
	itemNames.Doll:      0,
	itemNames.Coins:     0,
	itemNames.Key:       0,
	itemNames.Potion:    0,
	itemNames.Scroll:    0,
}
var totalEnemiesMinusQuestsObject = map[string]int{
	enemiesList.GOBLIN:   0,
	enemiesList.SKELETON: 0,
	enemiesList.ORC:      0,
	enemiesList.SORCERER: 0,
	enemiesList.DRAGON:   0,
}

var selectedQuestConditions = make([]*Condition, 3+Difficulty)

func checkCondition(t string, c *Condition) (ok bool) {
	ok = false
	switch t {
	case questTypes.KILL:
		ok = (aggregateEnemies[c.Target] >= c.Quantity) && (totalEnemiesMinusQuestsObject[c.Target] >= c.Quantity)
		totalEnemiesMinusQuestsObject[c.Target] -= c.Quantity
		break
	case questTypes.RETRIEVE:
		ok = (aggregateItems[c.Target] >= c.Quantity) && (totalItemsMinusQuestsObject[c.Target] >= c.Quantity)
		totalItemsMinusQuestsObject[c.Target] -= c.Quantity
		break
	default:
		ok = true
		break
	}
	return ok
}

func makeID() string {
	return strconv.FormatInt(int64(math.Abs(float64(time.Now().UTC().UnixNano()+rand.Int63()+rand.Int63()))), 16)
}

func indexOfConditions(arr []*Condition, item *Condition) bool {
	for _, elem := range arr {
		if elem == item {
			return true
		}
	}
	return false
}

var allQuests map[string]*Quest = make(map[string]*Quest)

func initializeQuests() {
	length := 3 + Difficulty
	// ***************************************************
	// ICI Quest Checks
	validateAvailableConditions()
	for n, a := range conditions {
		for _, c := range a {
			fmt.Printf("%s: %+v\n", n, c)
		}
		fmt.Println()
	}

	for i := 0; i < length; i++ {
		n := createNPC(createQuest)
		fmt.Printf("%+v\n", n.Quest)                // ICI
		fmt.Printf("%s\n", n.Quest.Dialogs.Request) // ICI
		allQuests[n.Quest.ID] = n.Quest
		pile.PushNPC(n)
	}

	var locs []*Location
	var rescueQuests []*Quest
	freeRoomsLength := len(noEnemyRooms)
	locs = noEnemyRooms
	for i := 0; i < length; i++ {
		index := rand.Intn(freeRoomsLength)
		r := locs[index]
		r.HasNPC = true
		r.NPC = pile.unshiftNPC()
		r.NPC.Quest.X = r.X
		r.NPC.Quest.Y = r.Y
		if r.NPC.QuestType == questTypes.SAVE {
			rescueQuests = append(rescueQuests, r.NPC.Quest)
		}
		locs = skipOneLocationByIndex(locs, index)
		freeRoomsLength = len(locs)
	}

	if len(rescueQuests) > 0 {
		var enemiesLocs []*Location
		enemyRoomsLength := len(enemyRooms)
		enemiesLocs = enemyRooms
		for _, q := range rescueQuests {
			for i := 0; i < q.Condition.Quantity; i++ {
				index := rand.Intn(enemyRoomsLength)
				r := enemiesLocs[index]
				r.Enemy.HasHostage = true
				r.Enemy.Hostage = &Hostage{Name: q.Condition.Target, QuestID: q.ID}
				enemiesLocs = skipOneLocationByIndex(enemiesLocs, index)
				enemyRoomsLength = len(enemiesLocs)
			}
		}
	}
	// ***************************************************
}

func cleanResolvedQuests() {
	length := 3 + Difficulty
	IDbox := make(map[string]*Quest, length)
	TARGETbox := make(map[string]map[string]*Quest, length)
	for id, quest := range playerQuestsMap.activeQuestsByID {
		if quest.Resolved {
			hero.LevelUp.Exp += quest.Condition.ExpValue
			hero.calcLVL()
			// Output("blue", quest.Dialogs.Answer)
			WorldMap[quest.Y][quest.X].HasNPC = false
			WorldMap[quest.Y][quest.X].Description += translate(someoneWasHereTR)
		}
		if !quest.Resolved {
			IDbox[id] = quest
			if _, ok := TARGETbox[quest.QuestType]; !ok {
				TARGETbox[quest.QuestType] = make(map[string]*Quest) // ICI
			}
			TARGETbox[quest.QuestType][id] = quest // ICI
		}
	}
	playerQuestsMap.activeQuestsByID = IDbox
	playerQuestsMap.activeQuestByTarget = TARGETbox
}

func skipOneLocationByIndex(arr []*Location, i int) []*Location {
	return append(append([]*Location{}, arr[:i]...), arr[i+1:]...)
}

func followKillQuestsEvolution(enemy *Character) {
	for _, questByID := range playerQuestsMap.activeQuestByTarget {
		for _, quest := range questByID {
			if quest.QuestType == questTypes.KILL {
				if quest.Condition.Target == enemy.Name {
					quest.Condition.Quantity--
					if quest.Condition.Quantity == 0 {
						quest.Resolved = true
						oneTimeQuestEvents.addEventString(dialogFromConditions(quest))
					}
				}
			}
		}
	}
}

func followSaveQuests(e *Character) {
	if e.HasHostage {
		if q, ok := playerQuestsMap.activeQuestsByID[e.Hostage.QuestID]; ok {
			q.Condition.Quantity--
			if q.Condition.Quantity == 0 {
				q.Resolved = true
				oneTimeQuestEvents.addEventString(dialogFromConditions(q))
			}
		} else {
			q := allQuests[e.Hostage.QuestID]
			q.Condition.Quantity--
			if q.Condition.Quantity == 0 {
				q.Resolved = true
			}
		}
	}
}

func followRetieveQuests() {
	loc := hero.SetPlayerRoom()
	quest := loc.NPC.Quest
	if !quest.Active {
		playerQuestsMap.addQuest(quest)
		quest.Active = true
		return
	}
	if q, ok := playerQuestsMap.activeQuestsByID[quest.ID]; ok {
		if q.QuestType == questTypes.RETRIEVE {
			requested := q.Condition.Target
			num := q.Condition.Quantity
			if hero.hasItemInInventory(requested) {
				if hero.Inventory[requested].Quantity >= num {
					hero.Inventory[requested].Quantity -= num
					q.Resolved = true
					oneTimeQuestEvents.addEventString(dialogFromConditions(q))
				}
			}
		}
	}
}

func followQuests(enemy *Character) {
	followKillQuestsEvolution(enemy)
	followSaveQuests(enemy)
}
