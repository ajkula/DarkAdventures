package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

var Out *os.File
var In *os.File

const X = 10
const Y = 10

var Grid [10][10]string
var player Character
var Difficulty int
var Hero int

func Check(e error) {
	if e != nil {
		fmt.Printf("Error retrieving data: %s\n", e)
	}
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	Out = os.Stdout
	In = os.Stdin
	ResetTurns()
	Difficulty = 0

	file, err := ioutil.ReadFile("landscape.txt")
	if err != nil {
		f, error := os.Create("landscape.txt")
		Check(error)
		defer f.Close()

		num, e := f.WriteString(DefaultLandscape)
		Check(e)
		f.Sync()
		// w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		// fmt.Fprintf(w, "Wrote default Landscape WorldMap: %v bytes wrote to disk\n\n", num)
		fmt.Fprintf(os.Stdout, "Wrote default Landscape WorldMap: %v bytes wrote to disk\n\n", num)
		file, e = ioutil.ReadFile("landscape.txt")
		Check(e)
	}

	Intro()
	var lines = strings.Split(string(file), "\n")
	for index := range lines {
		cols := strings.Split(lines[index], "")
		for rank, element := range cols {

			Grid[index][rank] = element
		}
	}

	// for _, l := range Grid {
	// 	fmt.Println("grid", l)
	// }
	CreateMap()
	InitGates()
}

func main() {
	// player = *new(Character)
	ChooseHero()
	hero := heroFromName(indexedHeroes[Hero])
	hero.SetPlayerRoom()
	hero.getImage()
	// hero.MoveTo("n")
	for {
		loc := hero.SetPlayerRoom()
		if !loc.HasEnemy {
			dragon.dragonMoves()
		}
		PresentScene(hero)
	}
	// DisplayWorldMap(hero)
	// hero.DisplayInvetory()
}

func Outputf(c string, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	Output(c, s)
}

// Output("green", "Combat round ", round, " begins...")
func Output(c string, args ...interface{}) {
	s := fmt.Sprint(args...)
	col := color.WhiteString
	switch c {
	case "green":
		col = color.GreenString
	case "red":
		col = color.RedString
	case "blue":
		col = color.CyanString
	case "yellow":
		col = color.YellowString
	}
	fmt.Fprintln(Out, col(s))
}

// UserInput(&playerAction)
func UserInput(i *int) {
	fmt.Fscan(In, i)
}

// cmd := UserInputln()
func UserInputln() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n >>> ")
	text, _ := reader.ReadString('\n')
	return text
}
