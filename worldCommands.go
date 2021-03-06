package main

import "fmt"

func DisplayWorldMap(p *Character) {
	defer ResetTurns()
	var result [][]string = create2DStringArray(Y, X, Unseen)

	for y, line := range Grid {
		for x := range line {

			if !WorldMap[y][x].Visited {
				result[y][x] = Unseen
			} else {
				result[y][x] = displayMapIcons[Grid[y][x]] // displayMapIcons[Grid[y][x]] Grid[y][x]

				if WorldMap[y][x].HasSeller {
					result[y][x] = Shop
				}

				if WorldMap[y][x].HasGate {
					result[y][x] = Root
				}
				if WorldMap[y][x].HasEnemy {
					result[y][x] = WorldMap[y][x].Enemy.Icon
				}
				if WorldMap[y][x].HasNPC {
					result[y][x] = NPCPosition
				}
			}
			if WorldMap[y][x].HasEnemy && Difficulty == 0 {
				if i := indexOf(giants, WorldMap[y][x].Enemy.Name); i != -1 {
					result[y][x] = WorldMap[y][x].Enemy.Icon
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

func create2DStringArray(Y, X int, Unseen string) [][]string {
	total := 0
	var tempo []string
	var w [][]string

	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			tempo = append(tempo, Unseen)
		}
		total += X
		w = append(w, tempo[total-X:total])
	}
	return w
}
