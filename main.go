package main

import (
	"bufio"
	"fmt"
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
var Difficulty int
var Hero int
var hero *Character
var msgDragonDead string = ""

func Check(e error) {
	if e != nil {
		fmt.Printf("Error retrieving data: %s\n", e)
	}
}

func init() {
	lastVer := getRepos()
	updateTitle(lastVer)
	// rand.Seed(time.Now().UTC().UnixNano())
	Out = os.Stdout
	In = os.Stdin
	ResetTurns()
	Difficulty = 0

	file, err := os.ReadFile("landscape.txt")
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
		file, e = os.ReadFile("landscape.txt")
		Check(e)
	}
	showTitle()
	Intro()
	var lines = strings.Split(string(file), "\n")
	cols := strings.Split(lines[0], "")
	Y = len(lines)
	X = len(cols)
	for index := range lines {
		cols := strings.Split(lines[index], "")
		Grid = append(Grid, cols)
	}

	CreateMap()
	CreateDragon()
	initializeQuests()

	makeEnemyTiers()
	InitGates()
}

func main() {
	var sigOnce int = 0
	ChooseHero()
	hero = heroFromName(indexedHeroes[Hero])
	hero.SetPlayerRoom()
	hero.getImage()
	pile.PushCharacters(hero)

	for {
		if msgDragonDead == enemiesList.DRAGON && sigOnce == 0 {
			nightWalkerA = CreateNightWalker()
			nightWalkerB = CreateNightWalker()
			sigOnce = 1
		}
		loc := hero.SetPlayerRoom()
		if !loc.HasEnemy {
			dragon.walkerMoves()
			if msgDragonDead == enemiesList.DRAGON {
				nightWalkerA.walkerMoves()
				nightWalkerB.walkerMoves()
			}
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

func updateTitle(lastVer *Repo) {
	// fmt.Printf("\n%+v\n", lastVer.Name) // ligne 29 title
	titleAsArray := strings.Split(title, "\n")
	ver := " " + releaseVersion + " "

	if lastVer.Name != releaseVersion {
		titleAsArray[28] = titleAsArray[28][:28] + updatedAvailable + titleAsArray[28][28+len(updatedAvailable):]
	}
	titleAsArray[32] = titleAsArray[32][:4] + ver + titleAsArray[32][4+len(ver):]

	title = strings.Join(titleAsArray, "\n")
}

func printQuestEvent(args ...interface{}) {
	s := fmt.Sprint(args...)
	col := color.HiGreenString
	fmt.Fprintln(Out, col(s))
}
