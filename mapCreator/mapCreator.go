package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

var letters = []string{"d", "l", "f"}
var castle string = "c"

var file []byte
var err error
var Y int
var X int
var randomizedLandscape string = " "
var fullSizeMap int
var maxCaps *Coords
var worldMap [][]string
var log = fmt.Print

func Check(e error) {
	if e != nil {
		fmt.Printf("Error retrieving data: %s\n", e)
	}
}

type Coords struct{ Y, X int }

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	Y = rand.Intn(10) + 10
	X = rand.Intn(10) + 10
	maxCaps = &Coords{
		Y: Y,
		X: X,
	}
	fullSizeMap = Y * X
	worldMap = makeWorldMapSizes(Y, X)
}

func main() {
	for i := 0; i < 18; i++ {
		i := i
		channel := make(chan string)
		room := getItemForArrayByBigLoop(letters, i)
		y := rand.Intn(Y)
		x := rand.Intn(X)
		newCoordsArr := getAllAdjacents(y, x)

		go fillerWorker(y, x, room, newCoordsArr, channel)

		select {
		case <-channel:
		}
	}
	text := ""
	for y, line := range worldMap {
		for x, elem := range line {
			if elem == randomizedLandscape {
				worldMap[y][x] = castle
			}
		}
		text += strings.Join(line, "\n")
		log(line, "\n")
	}
	writeRandomizedLandscape(text)
}

func fillerWorker(y, x int, room string, coordsArr []*Coords, ch chan string) {
	worldMap[y][x] = room
	newCoordsArr := getAllAdjacents(y, x)

	for len(coordsArr) > 0 {
		room := room
		newCoordsArr = newCoordsArr[0:0]
		for _, coords := range coordsArr {
			room := room
			worldMap[coords.Y][coords.X] = room
			coordsArr = append(newCoordsArr, getAllAdjacents(coords.Y, coords.X)...)
		}
	}
	ch <- room
}

func writeRandomizedLandscape(text string) {
	file, err = ioutil.ReadFile("landscape.txt")
	if err != nil {
		f, e := os.Create("landscape.txt")
		Check(e)
		defer f.Close()

		num, e := f.WriteString(text)
		Check(e)
		f.Sync()
		fmt.Fprintf(os.Stdout, "Wrote random Landscape WorldMap: %v bytes wrote to disk\n\n", num)
		file, e = ioutil.ReadFile("landscape.txt")
		Check(e)
	}
}

func makeWorldMapSizes(Y, X int) [][]string {
	var w [][]string
	total := 0
	var tempo []string
	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			tempo = append(tempo, randomizedLandscape)
		}
		total += X
		w = append(w, tempo[total-X:total])
	}
	return w
}

func getItemForArrayByBigLoop(arr []string, arg int) string {
	var res string
	res = arr[(arg)%len(arr)]
	return res
}

func getAllAdjacents(y, x int) (res []*Coords) {
	if y > 0 {
		if isFree(y-1, x) {
			res = append(res, &Coords{Y: y - 1, X: x})
		}
		if x < X-1 {
			if isFree(y-1, x+1) {
				res = append(res, &Coords{Y: y - 1, X: x + 1})
			}
		}
		if x > 0 {
			if isFree(y-1, x-1) {
				res = append(res, &Coords{Y: y - 1, X: x - 1})
			}
		}
	}
	if y < Y-1 {
		if isFree(y+1, x) {
			res = append(res, &Coords{Y: y + 1, X: x})
		}
		if x < X-1 {
			if isFree(y+1, x+1) {
				res = append(res, &Coords{Y: y + 1, X: x + 1})
			}
		}
		if x > 0 {
			if isFree(y+1, x-1) {
				res = append(res, &Coords{Y: y + 1, X: x - 1})
			}
		}
	}
	if x < X-1 {
		if isFree(y, x+1) {
			res = append(res, &Coords{Y: y, X: x + 1})
		}
	}
	if x > 0 {
		if isFree(y, x-1) {
			res = append(res, &Coords{Y: y, X: x - 1})
		}
	}
	return res
}

func isFree(y, x int) bool {
	b := false
	if worldMap[y][x] == randomizedLandscape {
		b = true
	}
	return b
}
