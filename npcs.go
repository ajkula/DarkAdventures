package main

import "math/rand"

// {Name} npc name, {Dialog} dialog for the quest, {QuestType} [kill, objects, save]
type NPC struct {
	Name, Dialog, QuestType string
	QuestID                 string
	Quest                   *Quest
	Icon                    string // people
	Active                  bool
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

var npcTypesList = []string{
	npcTypes.DarkElf,
	npcTypes.Elf,
	npcTypes.Gnoll,
	npcTypes.VillageChief,
	npcTypes.Villager,
}

func createNPC() (npc *NPC) {
	name := npcTypesList[rand.Intn(len(npcTypesList))]
	npc = &NPC{
		Name:      name,
		Dialog:    "",
		QuestID:   "",
		Quest:     &Quest{},
		QuestType: "",
		Icon:      "",
		Active:    false,
	}
	return npc
}
