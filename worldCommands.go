package main

import "fmt"

func DisplayWorldMap(p *Character) {
	defer ResetTurns()
	var result [10][10]string

	for y, line := range Grid {
		for x := range line {

			if !WorldMap[y][x].Visited {
				result[y][x] = Unseen
			} else {
				result[y][x] = Grid[y][x]

				if WorldMap[y][x].HasSeller {
					result[y][x] = Shop
				}

				if WorldMap[y][x].HasGate {
					result[y][x] = Root
				}
			}
			if WorldMap[y][x].HasEnemy && Difficulty == 0 {
				if WorldMap[y][x].Enemy.Name == enemiesList.DRAGON {
					result[y][x] = "V"
				}
			}

			if p.CurrentLocation[1] == x && p.CurrentLocation[0] == y {
				result[y][x] = YourPosition
			}
		}
	}

	Output("green", DoubleTab+"World Map:")
	for index, l := range result {
		Output("yellow", Tab+CustomSpaceAlign(OffsetLegend(LegendArray, 2, index), 20), l)
	}
	fmt.Println()
}
