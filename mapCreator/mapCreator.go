package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var letters = []string{gridLetters.Desert, gridLetters.Plains, gridLetters.Forest}
var castle string = gridLetters.Castle

var roots int = 12
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
	clear()
	channel := make(chan string, roots)
	for i := 0; i < roots; i++ {
		i := i
		room := getItemForArrayByBigLoop(letters, i)
		y := rand.Intn(Y)
		x := rand.Intn(X)
		newCoordsArr := getAllAdjacents(y, x)
		writeAt(x, y, room)

		go fillerWorker(y, x, room, newCoordsArr, channel)
	}
	for i := 0; i < roots; i++ {
		select {
		case <-channel:
		}
	}
	text := ""
	lr := "\n"
	for y, line := range worldMap {
		for x, elem := range line {
			if elem == randomizedLandscape {
				worldMap[y][x] = castle
				writeAt(x, y, castle)
			}
		}
		if y == len(worldMap)-1 {
			lr = ""
		}
		text += strings.Join(line, "") + lr
		// log(y, len(worldMap), line, "\n")
	}
	writeRandomizedLandscape(text)
}

func showProgress(s string, coordsArr []*Coords) {
	log(s, "\n")
	if len(coordsArr) > 0 {
		str := "[ "
		for _, dir := range coordsArr {
			str += "X: " + strconv.Itoa(dir.X) + " Y: " + strconv.Itoa(dir.Y) + ", "
		}
		str += "]\n"
		log(str)
	}
	for _, line := range worldMap {
		log(line, "\n")
	}
}

func fillerWorker(y, x int, room string, coordsArr []*Coords, ch chan string) {
	worldMap[y][x] = room
	writeAt(x, y, room)
	newCoordsArr := getAllAdjacents(y, x)

	for len(coordsArr) > 0 {
		// Monitor creation progress
		// showProgress(room, coordsArr)
		newCoordsArr = newCoordsArr[0:0]
		for _, coords := range coordsArr {
			worldMap[coords.Y][coords.X] = room
			writeAt(coords.X, coords.Y, room)
			coordsArr = append(newCoordsArr, getAllAdjacents(coords.Y, coords.X)...)
		}
	}
	ch <- room
}

func writeRandomizedLandscape(text string) {
	file, err := os.Create("landscape.txt")
	Check(err)
	defer file.Close()

	num, e := file.WriteString(text)
	Check(e)
	file.Sync()

	writeAt(X/2, Y+1, "Wrote random Landscape WorldMap: "+strconv.Itoa(num)+" bytes wrote to disk\n\n")
	_, e = ioutil.ReadFile("landscape.txt")
	Check(e)
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

func clear() {
	fmt.Print("\033[H\033[2J")
}

func writeAt(x, y int, str string) {
	fmt.Printf("\033[" + strconv.Itoa(y) + ";" + strconv.Itoa(x) + "H")
	fmt.Print(str)
	time.Sleep(100 * time.Millisecond)
}
