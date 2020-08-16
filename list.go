package main

type Pile struct {
	Enemies []*Character
	Gates   []*Gate
}

var pile *Pile = &Pile{}

func (pile *Pile) PushCharacters(args ...*Character) {
	players := args
	pile.Enemies = append(pile.Enemies, players...)
}

func (pile *Pile) PopCharacter() *Character {
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
		apply(enemi)
	}
}
