package main

type Location struct {
	Description, Ephemeral       string
	Item                         map[string]*ItemQuantity
	HasSeller, HasEnemy, Visited bool
	Seller                       string
	Enemy                        Character
	X, Y                         int
	CanGoTo                      []string
}

func (loc *Location) RemoveBattle() {
	if loc.HasEnemy && !loc.Enemy.isAlive() {
		loc.HasEnemy = false
		loc.Description += "\tYou see " + Article(loc.Enemy.Name) + "dead on the ground\n"
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
	Output("green", loc.CanGoTo, " ", loc.Y, " ", loc.X)
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
	loc.Enemy = NullifiedEnemy
	loc.Ephemeral = ""
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
	loc.Enemy = *dragon.Character
	loc.Ephemeral += dragonProximity["x"]
	for _, dir := range loc.CanGoTo {
		// Output("green", dir)
		switch dir {
		case directions.North:
			WorldMap[loc.Y-1][loc.X].Ephemeral += Tab + dragonProximity[Grid[loc.Y-1][loc.X]]
			break
		case directions.South:
			WorldMap[loc.Y+1][loc.X].Ephemeral += Tab + dragonProximity[Grid[loc.Y+1][loc.X]]
			break
		case directions.East:
			WorldMap[loc.Y][loc.X+1].Ephemeral += Tab + dragonProximity[Grid[loc.Y][loc.X+1]]
			break
		case directions.West:
			WorldMap[loc.Y][loc.X-1].Ephemeral += Tab + dragonProximity[Grid[loc.Y][loc.X-1]]
			break
		}
	}
}
