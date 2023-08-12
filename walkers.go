package main

import (
	"math/rand"
)

type Walker struct {
	*Character
	PreviousLocation []int
	Freeze           bool
}

type Direction struct {
	Name   string
	DeltaX int
	DeltaY int
}

var allDirections = []Direction{
	{Name: directions.North, DeltaX: 0, DeltaY: -1},
	{Name: directions.South, DeltaX: 0, DeltaY: 1},
	{Name: directions.East, DeltaX: 1, DeltaY: 0},
	{Name: directions.West, DeltaX: -1, DeltaY: 0},
}

func (walker *Walker) walkerMoves() {
	// walker.Encounter.reset()
	if !walker.Freeze {
		if isWalkerAlive := walker.isAlive(); isWalkerAlive {
			loc := walker.SetPlayerRoom()
			walker.PreviousLocation = []int{loc.Y, loc.X}

			possibleWays := walker.walkerPossibleWays()
			if len(possibleWays) > 0 {
				goTo := possibleWays[rand.Intn(len(possibleWays))]
				loc.ClearEphemeral()
				walker.MoveTo(goTo)
				Output("red", walker.Name+" moves to "+goTo)
			}
			newLoc := walker.SetPlayerRoom()
			newLoc.AddEphemeral(walker)
		}
	}
}

func (walker *Walker) shouldFreeze(str string) {
	walker.Freeze = InitialsIndexOf([]string{commands.Map}, str)
}

func (walker *Walker) walkerPossibleWays() []string {
	youCanGo := []string{}
	currentLocation := walker.SetPlayerRoom()

	for _, dir := range allDirections {
		newX := currentLocation.X + dir.DeltaX
		newY := currentLocation.Y + dir.DeltaY
		if isInsideMap(newY, newX) && canWalkerMoveThatWay(newY, newX) {
			youCanGo = append(youCanGo, dir.Name)
		}
	}
	return youCanGo
}

// check if a given x, y is inside the map boundaries
func isInsideMap(y, x int) bool {
	return x >= 0 && x < X && y >= 0 && y < Y
}

func canWalkerMoveThatWay(y, x int) bool {
	var loc *Location = WorldMap[y][x]
	if loc.HasEnemy {
		return false
	}
	if loc.HasSeller {
		return false
	}
	if loc.HasNPC {
		return false
	}
	return true
}

func getAvailableRoom(yi, xi int) []int {
	walkerStartPosition := []int{rand.Intn(yi), rand.Intn(xi)}
	y := walkerStartPosition[0]
	x := walkerStartPosition[1]
	if !(WorldMap[y][x].HasEnemy || WorldMap[y][x].HasSeller) {
		return walkerStartPosition
	}
	return getAvailableRoom(yi, xi)
}
