package main

import (
	"math/rand"
)

type Walker struct {
	*Character
	PreviousLocation []int
	Freeze           bool
}

func (walker *Walker) walkerMoves() {
	// walker.Encounter.reset()
	if !walker.Freeze {
		if ok := walker.isAlive(); ok {
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
	var youCanGo []string
	var yourPlace *Location
	yourPlace = walker.SetPlayerRoom()

	if yourPlace.X >= 1 {
		if canWalkerMoveThatWay(yourPlace.Y, yourPlace.X-1) {
			youCanGo = append(youCanGo, directions.West)
		}
	}
	if yourPlace.X <= X-2 {
		if canWalkerMoveThatWay(yourPlace.Y, yourPlace.X+1) {
			youCanGo = append(youCanGo, directions.East)
		}
	}

	if yourPlace.Y >= 1 {
		if canWalkerMoveThatWay(yourPlace.Y-1, yourPlace.X) {
			youCanGo = append(youCanGo, directions.North)
		}
	}
	if yourPlace.Y <= Y-2 {
		if canWalkerMoveThatWay(yourPlace.Y+1, yourPlace.X) {
			youCanGo = append(youCanGo, directions.South)
		}
	}
	return youCanGo
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
