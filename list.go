package main

type Pile struct {
	Enemies       []*Character
	Gates         []*Gate
	AllCharacters []*Character
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

func (pile *Pile) PopCharacter(player *Character) *Character {

	enemy := pile.AllCharacters[:1][0]
	pile.Enemies = pile.AllCharacters[1:]
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
