package main

type QuestTypes struct {
	KILL, RETRIEVE, SAVE string
}

var questProbabilityPercentage = 5
var questTypes = &QuestTypes{KILL: "kill", RETRIEVE: "retrieve", SAVE: "save"}

// {Name} npc name, {Dialog} dialog for the quest, {QuestType} [kill, objects, save]
type NPC struct {
	Name, Dialog, QuestType string
	Conditions              *Conditions
	resolution              string
	QuestName               string
	Quest                   string
	Icon                    string // people
	Active                  bool
}

type Conditions struct {
	QuestType string
	Target    string
	Quantity  int
	ExpValue  int
	Item      *ItemQuantity
	Answer    string
}

type NPCTypes struct {
	Elf, DarkElf, Villager, VillageChief, Gnoll string
}

var npcTypes = &NPCTypes{
	DarkElf:      translate(DarkElfTR),
	Elf:          translate(ElfTR),
	Gnoll:        translate(GnollTR),
	VillageChief: translate(VillagerChiefTR),
	Villager:     translate(VillagerTR),
}

type Quest struct{}

type PlayerQuestsMap struct {
	activeQuestsByType map[string]*Quest
}

var playerQuestsMap = &PlayerQuestsMap{
	activeQuestsByType: map[string]*Quest{},
}
