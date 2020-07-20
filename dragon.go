package main

type Dragon struct {
	*Character
	PreviousLocation [2]int
}

func (dragon *Dragon) dragonMoves() {
	loc := dragon.SetPlayerRoom()
	dragon.PreviousLocation[0] = loc.Y
	dragon.PreviousLocation[1] = loc.X

}

func CreateDragon() {
	dragon = &Dragon{
		// Character: ,
		// Name:            heroesList.Thieve,
		// Alive:           true,
		// CurrentLocation: []int{9, 4},
		// Evasion:         30,
		// Health:          rand.Intn(45) + 10,
		// Skill:           0,
		// Strength:        18,
		// Crit:            25,
		// Inventory:       map[string]*ItemQuantity{},
	}
}
