package main

type PlayerQuestsMap struct {
	activeQuestsByType map[string]*Quest
}

var playerQuestsMap = &PlayerQuestsMap{
	activeQuestsByType: map[string]*Quest{},
}

var questProbabilityPercentage = 5

type QuestTypes struct{ KILL, RETRIEVE, SAVE string }

var questTypes = &QuestTypes{KILL: "kill", RETRIEVE: "retrieve", SAVE: "save"}

var noEnemyRooms = []*Location{}
var enemyRooms = []*Location{}

func populateTheSlices(loc *Location) {
	if loc.HasEnemy {
		enemyRooms = append(enemyRooms, loc)
	} else {
		noEnemyRooms = append(noEnemyRooms, loc)
	}
}

var conditions = map[string]*Condition{}

// QuestTypes
var baseConditionsMap = map[string][]*Condition{
	questTypes.KILL: {
		&Condition{Target: enemiesList.GOBLIN, Quantity: 4},
		&Condition{Target: enemiesList.ORC, Quantity: 1},
		&Condition{Target: enemiesList.SORCERER, Quantity: 2},
		&Condition{Target: enemiesList.NIGHTWALKER, Quantity: 2},
		&Condition{Target: enemiesList.DRAGON, Quantity: 1},
	},
	questTypes.SAVE: {
		&Condition{Target: translate(RescueTargetTR), Quantity: 1},
		&Condition{Target: translate(ShamanTR), Quantity: 1},
		&Condition{Target: translate(VillagerTR), Quantity: 2},
	},
	questTypes.RETRIEVE: {
		&Condition{Target: itemNames.Scroll, Quantity: 1},
		&Condition{Target: itemNames.Moonstone, Quantity: 1},
		&Condition{Target: itemNames.Potion, Quantity: 1},
	},
}

type Quest struct {
	ID        string
	QuestType string
	Condition *Condition
	Resolved  bool
	ExpValue  int
	Answer    string
}

type Condition struct {
	Target   string
	Quantity int
}

func createQuest() {

}

func checkCondition(t string, c *Condition) (ok bool) {
	ok = false
	switch t {
	case questTypes.KILL:
		ok = aggregateEnemies[c.Target] >= c.Quantity
		break
	case questTypes.RETRIEVE:
		ok = aggregate[c.Target] >= c.Quantity
		break
	default:
		ok = true
		break
	}
	return ok
}
