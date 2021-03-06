package main

type Pile struct {
	Enemies       []*Character
	Gates         []*Gate
	AllCharacters []*Character
	NPCS          []*NPC
	NPCsRefList   []*NPC
}

var pile *Pile = &Pile{}

func (pile *Pile) PushCharacters(player *Character) {
	if player.Npc {
		pile.Enemies = append(pile.Enemies, player)
	}
	pile.AllCharacters = append(pile.AllCharacters, player)
}

func (pile *Pile) PopEnemy() *Character {
	enemy := pile.Enemies[:1][0]
	pile.Enemies = pile.Enemies[1:]
	return enemy
}

func (pile *Pile) PushGates(args ...*Gate) {
	gate := args
	pile.Gates = append(pile.Gates, gate...)
}

func (pile *Pile) PopGate() *Gate {
	gate := pile.Gates[:1][0]
	pile.Gates = pile.Gates[1:]
	return gate
}

func (pile *Pile) forEachEnemy(apply func(*Character)) {
	for _, enemi := range pile.Enemies {
		if enemi.isAlive() {
			apply(enemi)
		}
	}
}

func (pile *Pile) forEachCharacter(apply func(*Character)) {
	for _, player := range pile.AllCharacters {
		if player.isAlive() {
			apply(player)
		}
	}
}

func (pile *Pile) checkIfDead(player *Character) {
	for i, enemi := range pile.Enemies {
		if !enemi.isAlive() {
			pile.ejectCharacter(i)
		}
	}
}

func (pile *Pile) removeCharacter(player *Character) {
	i := pile.indexOfCharacter(player)
	pile.ejectCharacter(i)
}

func (pile *Pile) indexOfCharacter(player *Character) (i int) {
	i = -1
	for index, char := range pile.Enemies {
		if player == char {
			i = index
		}
	}
	return i
}

func (pile *Pile) ejectCharacter(i int) {
	pile.Enemies = append(pile.Enemies[:i], pile.Enemies[i+1:]...)
}

func (pile *Pile) PushNPC(args ...*NPC) {
	npc := args
	pile.NPCS = append(pile.NPCS, npc...)
	pile.NPCsRefList = append(pile.NPCsRefList, npc...)
}

func (pile *Pile) unshiftNPC() *NPC {
	npc := pile.NPCS[:1][0]
	pile.NPCS = pile.NPCS[1:]
	return npc
}

func (pile *Pile) find(id string) (n *NPC) {
	for id, npc := range pile.NPCsRefList {
		if id == id {
			return npc
		}
	}
	return nil
}
