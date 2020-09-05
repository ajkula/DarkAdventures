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

var X = 10
var Y = 10

var Grid [][]string
var player Character
var Difficulty int
var Hero int
var hero *Character

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
	showTitle()
	Intro()
	var lines = strings.Split(string(file), "\n")
	cols := strings.Split(lines[0], "")
	var temporary []string
	temporary = append(temporary, cols...)
	Y = len(lines)
	X = len(cols)
	for index := range lines {
		cols := strings.Split(lines[index], "")
		// for rank, element := range cols {
		temporary = cols
		Grid = append(Grid, temporary)
		// Grid[index] = append(Grid[index], cols...)
		// Grid[index][rank] = element
		// }
	}
	// for _, l := range Grid {
	// 	fmt.Println("grid", l)
	// }
	CreateMap()
	makeEnemyTiers()
	InitGates()
}

func main() {
	// player = *new(Character)
	ChooseHero()
	hero = heroFromName(indexedHeroes[Hero])
	hero.SetPlayerRoom()
	hero.getImage()
	pile.PushCharacters(hero)
	// hero.addStatus(&Blueprint{Name: statuses.Dark})
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

func showTitle() {
	Output("blue", title)
	time.Sleep(2 * time.Second)
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
		col = color.HiRedString
	case "blue":
		col = color.HiCyanString
	case "yellow":
		col = color.YellowString
	case "white":
		col = color.HiWhiteString
	case "stats":
		col = color.HiYellowString
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

func makeEnemyTiers() {
	tierExp1 := 2 * expByDifficulty[Difficulty]
	tierExp2 := expByDifficulty[Difficulty]
	tier1 := Y / 3
	tier2 := tier1 * 2
	var exp int = 0
	for y := 0; y < Y; y++ {
		if y <= tier2 {
			exp = tierExp2
		}
		if y <= tier1 {
			exp = tierExp1
		}
		for x := 0; x < X; x++ {
			if WorldMap[y][x].HasEnemy {
				WorldMap[y][x].Enemy.LevelUp.Exp += exp
			}
		}
	}
	pile.forEachEnemy(makeEnemiesLVL)
}
