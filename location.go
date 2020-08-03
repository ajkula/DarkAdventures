package main

type Location struct {
	Description, Ephemeral, treasure                string
	Item, Chest                                     map[string]*ItemQuantity
	HasSeller, HasEnemy, Visited, HasGate, HasChest bool
	Seller                                          string
	Enemy                                           *Character // ICI
	X, Y                                            int
	Gate                                            *Gate
	CanGoTo                                         []string
}

func (loc *Location) RemoveBattle() {
	if loc.HasEnemy && !loc.Enemy.isAlive() {
		loc.HasEnemy = false
		loc.Description += "\tYou see " + Article(loc.Enemy.Name) + "dead on the ground\n"
	}
}

func (loc *Location) showImage() {
	if loc.HasEnemy {
		loc.Enemy.getImage()
	}
}

// To Do: make a shop inventory and remove item from inv func
func (loc *Location) RemoveItem(n string) {
	b := false
	if _, ok := loc.Item[n]; ok {
		loc.Item[n].Quantity--
		// loc.Item[n].Type.Price
	}
	for _, item := range loc.Item {
		if item.Quantity >= 1 {
			b = true
		}
	}
	if !b {
		loc.HasSeller = false
		loc.Description += Tab + "There was someone here, it's empty now."
	}
}

func (loc *Location) addDescriptionToAdjacentRooms(add string) {
	// Output("green", loc.CanGoTo, " ", loc.Y, " ", loc.X)
	for _, dir := range loc.CanGoTo {
		// Output("green", dir)
		switch dir {
		case directions.North:
			WorldMap[loc.Y-1][loc.X].Description += Tab + add
			break
		case directions.South:
			WorldMap[loc.Y+1][loc.X].Description += Tab + add
			break
		case directions.East:
			WorldMap[loc.Y][loc.X+1].Description += Tab + add
			break
		case directions.West:
			WorldMap[loc.Y][loc.X-1].Description += Tab + add
			break
		}
	}
}

func (loc *Location) ClearEphemeral() {
	loc.HasEnemy = false
	loc.Enemy = &NullifiedEnemy
	// WorldMap[loc.Y][loc.X].Ephemeral = ""
	// Output("red", "ClearEphemeral() ", WorldMap[loc.Y][loc.X] == loc)
	for _, dir := range loc.CanGoTo {
		// Output("green", dir)
		switch dir {
		case directions.North:
			WorldMap[loc.Y-1][loc.X].Ephemeral = ""
			break
		case directions.South:
			WorldMap[loc.Y+1][loc.X].Ephemeral = ""
			break
		case directions.East:
			WorldMap[loc.Y][loc.X+1].Ephemeral = ""
			break
		case directions.West:
			WorldMap[loc.Y][loc.X-1].Ephemeral = ""
			break
		}
	}
}

func (loc *Location) AddEphemeral() {
	loc.HasEnemy = true
	loc.Enemy = dragon.Character
	WorldMap[loc.Y][loc.X].Ephemeral = Tab + dragonProximity["x"]
	// Output("red", "AddEphemeral() ", WorldMap[loc.Y][loc.X] == loc)
	// fmt.Printf("\ny: %v x: %v  EPHEMERAL: %v", loc.Y, loc.X, WorldMap[loc.Y][loc.X].Ephemeral)
	for _, dir := range loc.CanGoTo {
		// Output("green", dir)
		switch dir {
		case directions.North:
			WorldMap[loc.Y-1][loc.X].Ephemeral = Tab + dragonProximity[Grid[loc.Y-1][loc.X]]
			// fmt.Printf("\ny: %v x: %v  EPHEMERAL: %v", loc.Y-1, loc.X, WorldMap[loc.Y-1][loc.X].Ephemeral)
			break
		case directions.South:
			WorldMap[loc.Y+1][loc.X].Ephemeral = Tab + dragonProximity[Grid[loc.Y+1][loc.X]]
			// fmt.Printf("\ny: %v x: %v  EPHEMERAL: %v", loc.Y+1, loc.X, WorldMap[loc.Y+1][loc.X].Ephemeral)
			break
		case directions.East:
			WorldMap[loc.Y][loc.X+1].Ephemeral = Tab + dragonProximity[Grid[loc.Y][loc.X+1]]
			// fmt.Printf("\ny: %v x: %v  EPHEMERAL: %v", loc.Y, loc.X+1, WorldMap[loc.Y][loc.X+1].Ephemeral)
			break
		case directions.West:
			WorldMap[loc.Y][loc.X-1].Ephemeral = Tab + dragonProximity[Grid[loc.Y][loc.X-1]]
			// fmt.Printf("\ny: %v x: %v  EPHEMERAL: %v", loc.Y, loc.X-1, WorldMap[loc.Y][loc.X-1].Ephemeral)
			break
		}
	}
	// y := dragon.CurrentLocation[0]
	// x := dragon.CurrentLocation[1]
	// fmt.Printf("\nDRAGON POSITION:  y: %v x: %v  EPHEMERAL: %v", y, x, WorldMap[y][x].Ephemeral)
}
