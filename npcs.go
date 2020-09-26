package main

import "strconv"

// {Name} npc name, {Dialog} dialog for the quest, {QuestType} [kill, objects, save]
type NPC struct {
	Name, QuestType string
	QuestID         string
	Quest           *Quest
	Icon            string // people
}

type NPCTypes struct {
	Elf, DarkElf, Villager, VillageChief, Gnoll string
}

var npcTypes = &NPCTypes{
	DarkElf:      translate(DarkElfNPCTR),
	Elf:          translate(ElfNPCTR),
	Gnoll:        translate(GnollNPCTR),
	VillageChief: translate(VillagerChiefNPCTR),
	Villager:     translate(VillagerNPCTR),
}

var npcTypesList = []string{
	npcTypes.DarkElf,
	npcTypes.Elf,
	npcTypes.Gnoll,
	npcTypes.VillageChief,
	npcTypes.Villager,
}

func createNPC(f func() (q *Quest)) (npc *NPC) {
	q := f()
	name := getRandomArrayString(npcTypesList)
	npc = &NPC{
		Name:      name,
		QuestID:   q.ID,
		Quest:     q,
		QuestType: q.QuestType,
		Icon:      "",
	}
	npc.Quest.Dialogs = makeNPCDialogs(npc)
	return npc
}

func makeNPCDialogs(npc *NPC) *QuestDialogs {
	/*
					Hey there, I'm [a Dark Elf].
				{
					Would you help me to [kill] [4] [goblin]?
					Could you bring me [1] [potion]?
					Please help me [save] [1] [Friend][""|"s"]
				}
		******************* Answer *******************
	*/
	var retrieveAcceptedEndSentence string = ""
	if questTypes.RETRIEVE == npc.QuestType {
		retrieveAcceptedEndSentence = strconv.Itoa(npc.Quest.Condition.Quantity) + " " + npc.Quest.Condition.Target + "?"
	}

	return &QuestDialogs{
		Greetings: translate(GreetingsTR) + npc.Name + "\n",
		Request: Tab + questTypesRequestsDialog[npc.QuestType] +
			strconv.Itoa(npc.Quest.Condition.Quantity) + " " +
			npc.Quest.Condition.Target +
			plurial(npc.Quest.Condition.Quantity) + "?\n",
		Accepted: Tab + questsAcceptedDialog[npc.QuestType] + retrieveAcceptedEndSentence + "\n",
		Answer: DoubleTab + translate(questAnswerTR) + "+" +
			strconv.Itoa(npc.Quest.Condition.ExpValue) + " EXP\n",
	}
}

func plurial(i int) (s string) {
	if i > 1 {
		s = "s"
	}
	return s
}

var questTypesRequestsDialog = map[string]string{
	questTypes.KILL:     translate(killRequestTR),
	questTypes.RETRIEVE: translate(retrieveRequestTR),
	questTypes.SAVE:     translate(saveRequestTR),
}

var questsAcceptedDialog = map[string]string{
	questTypes.KILL:     translate(killsAcceptTR),
	questTypes.RETRIEVE: translate(retrieveAcceptTR),
	questTypes.SAVE:     translate(saveAcceptTR),
}
