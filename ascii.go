package main

import (
	"math/rand"
	"time"
)

type ASCII struct {
	GOBLIN      []string
	SKELETON    []string
	SORCERER    []string
	ORC         []string
	DRAGON      []string
	NIGHTWALKER []string
	NECROMANCER []string
	HERO        map[string][]string
	SKILLS      map[string][]string
}

var AsciiArts = &ASCII{
	GOBLIN:      []string{goblinAscii, goblinAsciiB},
	SKELETON:    []string{skeletonAscii, skeletonAsciiB},
	SORCERER:    []string{sorcererAscii, sorcererAsciiB},
	ORC:         []string{orcAscii, orcAsciiB},
	DRAGON:      []string{dragonAscii, dragonAsciiB},
	NIGHTWALKER: []string{nightWalkerAsciiA},
	NECROMANCER: []string{necromancerAsciiA},
	HERO: map[string][]string{
		heroesList.Thief:     {thiefAscii, thiefAsciiB, thiefAsciiC, thiefAsciiD},
		heroesList.Paladin:   {paladinAscii, paladinAsciiB, paladinAsciiC, paladinAsciiD},
		heroesList.Wizard:    {wizardAscii, wizardAsciiB, wizardAsciiC, wizardAsciiD},
		heroesList.Barbarian: {barbarianAscii, barbarianAsciiB, barbarianAsciiC, barbarianAsciiD},
	},
	SKILLS: map[string][]string{
		enemiesList.DRAGON:      {dragonSkillFireAsciiA, dragonSkillFireAsciiB},
		enemiesList.NIGHTWALKER: {nightWilkerSkillClawAsciiA},
	},
}

var giants = []string{enemiesList.DRAGON, enemiesList.NECROMANCER, enemiesList.NIGHTWALKER}

func (a *ASCII) makeImage(name string) *DisplayImage {
	var image *DisplayImage
	switch name {
	case enemiesList.GOBLIN:
		image = &DisplayImage{
			Image: a.GOBLIN[rand.Intn(len(a.GOBLIN))],
			Show:  true,
			Race:  races.Darkling,
		}
	case enemiesList.SKELETON:
		image = &DisplayImage{
			Image: a.SKELETON[rand.Intn(len(a.SKELETON))],
			Show:  true,
			Race:  races.Undead,
		}
	case enemiesList.SORCERER:
		image = &DisplayImage{
			Image: a.SORCERER[rand.Intn(len(a.SORCERER))],
			Show:  true,
			Race:  races.Human,
		}
	case enemiesList.ORC:
		image = &DisplayImage{
			Image: a.ORC[rand.Intn(len(a.ORC))],
			Show:  true,
			Race:  races.Darkling,
		}
	case enemiesList.DRAGON:
		image = &DisplayImage{
			Image: a.DRAGON[rand.Intn(len(a.DRAGON))],
			Show:  true,
		}
	case enemiesList.NIGHTWALKER:
		image = &DisplayImage{
			Image: a.NIGHTWALKER[rand.Intn(len(a.NIGHTWALKER))],
			Show:  true,
		}
	default:
		index := rand.Intn(len(a.HERO[name]))
		race := races.Human
		if index == 1 {
			if name == heroesList.Wizard {
				race = races.Tiefling
			}
			if name == heroesList.Barbarian {
				race = races.Gnoll
			}
		}
		if index == 3 {
			if name == heroesList.Thief {
				race = races.Elf
			}
		}
		image = &DisplayImage{
			Image: a.HERO[name][index],
			Story: "\n" + storyFromImage(name, index),
			Show:  true,
			Race:  race,
		}
	}
	return image
}

type Races struct {
	Gnoll, Tiefling, Human, Undead, Darkling, Elf string
}

var races = &Races{
	Gnoll:    translate(GnollTR),
	Elf:      translate(ElfTR),
	Human:    translate(HumanTR),
	Tiefling: translate(TieflingTR),
	Undead:   translate(UndeadTR),
	Darkling: translate(DarklingTR),
}

func storyFromImage(name string, index int) string {
	return heroStoryAndImages[name][index]
}

var heroStoryAndImages = map[string]map[int]string{
	heroesList.Thief: {
		0: translate(heroStoryThiefTR0),
		1: translate(heroStoryThiefTR1),
		2: translate(heroStoryThiefTR2),
		3: translate(heroStoryThiefTR3),
	},
	heroesList.Paladin: {
		0: translate(heroStoryPaladinTR0),
		1: translate(heroStoryPaladinTR1),
		2: translate(heroStoryPaladinTR2),
		3: translate(heroStoryPaladinTR3),
	},
	heroesList.Wizard: {
		0: translate(heroStoryWizardTR0),
		1: translate(heroStoryWizardTR1),
		2: translate(heroStoryWizardTR2),
		3: translate(heroStoryWizardTR3),
	},
	heroesList.Barbarian: {
		0: translate(heroStoryBarbarianTR0),
		1: translate(heroStoryBarbarianTR1),
		2: translate(heroStoryBarbarianTR2),
		3: translate(heroStoryBarbarianTR3),
	},
}

func (a *ASCII) showSkillAction(name string) {
	switch name {
	case enemiesList.DRAGON:
		index := rand.Intn(len(a.SKILLS[name]))
		Output("red", a.SKILLS[name][index])
		time.Sleep(1 * time.Second)
		break
	case enemiesList.NIGHTWALKER:
		index := rand.Intn(len(a.SKILLS[name]))
		Output("red", a.SKILLS[name][index])
		time.Sleep(1 * time.Second)
	default:
		break
	}
}

var title string = "88888888888888888888888888888888888888888888888888888888888888888888888888888888\n" +
	"88888888888888888888888888888888888888888888888888Z$8888888888888888888888888888\n" +
	"8888888888888888888888888888OO8D88888888888888888I,:O888888888888888888888888888\n" +
	"8888888888888888888888888$7I?+~?O8888888D88D88888~.+D$IO888888888888888888888888\n" +
	"88888888888888888888888O==+~+7~.=888O$Z$7OO?=I~:$,,??~7O888888888888888888888888\n" +
	"8888888888888888888888O:~I:.?OI,,O8+:=~.,O7.,+~=+.,:,,?8888888888888888888888888\n" +
	"8888888888888888888888$.II,,$D7.:O=.~7~.~8+.~$77~.~?:.?D888888888888888888888888\n" +
	"8888888888888888888888O=Z+.~OD+.=I.:7Z:.I8:.?88Z,.IZ:.=Z888888888888888888888888\n" +
	"8888888888888888888888O$Z:.?DZ,,I=.~O?.,I+.,$8DI.:Z8=:=Z888888888888888888888888\n" +
	"88888888888888888888888D7.,7Z:,=Z+.,~~::~~~=O887+I887I$O888888888888888888888888\n" +
	"888888888888888888888888+,:~,:+787==?7I?77I$888Z$Z888O888888888888888888DD888888\n" +
	"888888888888888888888888I+???I$88O$$Z8OZOOO8888DD8888888888888D8888OO888I+788888\n" +
	"888888888888888888888888ZZOOZO8888888888888888Z++O8D8888888ZZZI$D$==~7DI..+D8888\n" +
	"8888888888888D888888888O888888888888888D8D8888+.:OZI7DI~?D?.:+.=I.:I,I8~:,+88888\n" +
	"88888888888DOI~~O8888$::O888888888ZZ88$I7?=I8Z,,=Z=.+8:.?8:.=7++,,=~~Z=?+.=88888\n" +
	"8888888888DI~~.:O8D88+.~8Z7ZZ~7D$~=~+Z..~:.:8$.,I$,,IZ,,$Z,,7OZ+.:?II~:$=.+88888\n" +
	"88888888D8?:I~,+8O?=+:.?8:.IZ,?$.:I:+I.:7+.=D?.~8I.:O+.~O+.:O8D?.,+==~:~:~I88888\n" +
	"8888888$$+.+?,,7Z,,++.,$$.,$8~+:.~=~$=.=O~.IO:.+Z:.:?:.:=,,+8887~~~+77??I7O88888\n" +
	"888887:=~.=I~.:Z:.=Z=.~8?.~8$:+,,?II+,.7O,.~~,,:~=::=?==+++$888O77$ZO8ZZO8888888\n" +
	"88888=:?,~7Z~.?7.,$O:.?Z,.=I:?7,,~~==,:ZO+~=??+?7$II$Z7$Z$ZO88888888888888888888\n" +
	"888887I:,I8$,.$7.,+=,,:=~,::+ZZ+==+77?I88$7$OZ$ZO8OO8888888888888888888888888888\n" +
	"88888OI.~ZD?,:O$~~~??=+II??7Z88Z$$ZOOZO88888888888888888888888888888888888888888\n" +
	"88888D+.?88I+I8O7I7ZZ$$OOZOO8888888888888888888888888888888888888888888888888888\n" +
	"888888I:788Z$O888O88888888888888888888888888888888888888888888888888888888888888\n" +
	"888888ZIZ88888888888888888888888888888888888888888888888888888888888888888888888\n" +
	"8888888O888888888888888888888888888888888888888888888888888888888888888888888888\n" +
	"88888888888888888888888888888888888888888888888888888888888888888888888888888888\n" +
	"88888888888888888888888888888888888888888888888888888888888888888888888888888888\n" +
	"88888888888888888888888888888888888888888888888888888888888888888888888888888888\n" +
	"88888888888888888888888888888888888888888888888888888888888888888888888888888888\n" +
	"88888888888888888888888888888888888888888888888888888888888888888888888888888888\n" +
	"88888888888888888888888888888888888888888888888888888888888888888888888888888888\n" +
	"888888888888888888888888888888888888888888888 a game by Greg Dazbog 888888888888\n" +
	"88888888888888888888888888888888888888888888888888888888888888888888888888888888\n" +
	"88888888888888888888888888888888888888888888888888888888888888888888888888888888\n"

const goblinAscii string = DoubleTab + "...................................\n" + DoubleTab +
	"...............,:,.................\n" + DoubleTab +
	"..............=+++=................\n" + DoubleTab +
	"...........:II$I+?7II~.............\n" + DoubleTab +
	"...........~IZ$$I$$$+:.............\n" + DoubleTab +
	"..........=II$ZZOZ$777=............\n" + DoubleTab +
	".........~ZOZZZZ7$$Z7$ZI:..........\n" + DoubleTab +
	"........:$Z~+Z7$$$Z?..:$7~.........\n" + DoubleTab +
	"........?$:.=Z$Z$$O=...,?I,........\n" + DoubleTab +
	".......~I:..:Z$Z$$$?.....~+........\n" + DoubleTab +
	"......,+.....?$77II$:.....=+.......\n" + DoubleTab +
	"......??=:..~I$77I$$I:..:=+?,......\n" + DoubleTab +
	"......+?:,.=++7OZOZ$I?..,.:7~......\n" + DoubleTab +
	".....,?=..+?+=~I$OZ~??+..,+I~......\n" + DoubleTab +
	"......=?:.7$:..+Z7$=.?$?,.,,.......\n" + DoubleTab +
	"..........I7~..+Z$7$,.+Z7..........\n" + DoubleTab +
	"......,.,.=I:,,=Z$77~,,~I=.,,......\n" + DoubleTab +
	".,,,,,,,::=?+~==7??+===~=I=:,,,,,,.\n" + DoubleTab +
	",,,,,::=?II7?====+++====+77I~,,,,..\n" + DoubleTab +
	".....,,~=~::,,,.,,,......,::,......\n" + DoubleTab +
	"...................................\n"

const skeletonAscii string = DoubleTab + "MMMMMMMMMMMOMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMZMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMD7IDMMDMMD78MMMMMMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMM8...:$+:+$DDMDZDMMMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMM8,........=D8D8MMMMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMO=~=:......?MMMMMMMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMD$+OM?.,....:$D$=$MMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMOOMMM=,~.~?8?,IMMMMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMD8ODMMM8~D==:+D=.,ZMMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMDMMMMMI$M=...~M8+.~$ZMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMM7?D~.....=MMO+:=ZMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMM7?M+.,,,...8MO$MDI+$DMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMM$8M:.7=,,..$MZMMMMD$I78MMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMO+,DO~+~.:MOMMMMMMMOI78MMMMMMMM\n" + DoubleTab +
	"MMMMMMMMM7..ZD7MM+,ODMMMMMMMMM877OMMMMMM\n" + DoubleTab +
	"MMMMMMMMMD..8MZOM$..7MMMMMMMMMMMDZ7OMMMM\n" + DoubleTab +
	"MMMMMMMMMMI.OMM8MM$.:MMMMMMMMMMMMMMO$OMM\n" + DoubleTab +
	"MMMMMMMMMMM,~8DMMMMO.7DDMMMMMMMMMMMMMDMM\n" + DoubleTab +
	"MMMMMMMMOI+~=88888OM=.I88MMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMD88MMMMMMMMM7=OMMMMMMMMMMMMMMMMM\n"

const sorcererAscii string = DoubleTab + "+++++++++??++++++++=++++++++++\n" + DoubleTab +
	"????????????????II?????II?I???\n" + DoubleTab +
	"IIIIIIIIIIIII77I777III7777IIII\n" + DoubleTab +
	"777777$$7777$I~:I$7777$$$$7777\n" + DoubleTab +
	"$ZZ$$$ZZZZ$Z$,~~.$Z$ZZZZZZ$$$$\n" + DoubleTab +
	"ZZZZZZZZZZOZ=.,..~$OZOOZZZZZZZ\n" + DoubleTab +
	"OOOOOZZZZZZ=,......7OZOZZZZOOO\n" + DoubleTab +
	"888OOOOZZZ+~,.,,::.,OOZOOOOOOO\n" + DoubleTab +
	"8888OOO7$~:,..,..:,.:ZOZZ88888\n" + DoubleTab +
	"D88888Z?~,,......,,..:7$ZO8888\n" + DoubleTab +
	"DD8888Z?,,..:,,~,,:...:I8DDDDD\n" + DoubleTab +
	"DD8888D$,...:,.::.~...I88DDDDD\n" + DoubleTab +
	"DDDDD8D8:....:.:,.,..=MDDDDDDD\n" + DoubleTab +
	"DDDDDDD8,....,.:..:.~8MDMDDDDD\n" + DoubleTab +
	"DDDDDDMD??:....,..:.IMDMDMMMMM\n" + DoubleTab +
	"MMMMMMMM8O:,,,,.,..=8DMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMD:,:,,.,,.?MDMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMM+.:,.,,:,7DMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMM8~,:,,,,~:?OZMMMMMMMM\n" + DoubleTab +
	"MMMMMMDDM$:~:::::~~:+?7$ZODDMD\n" + DoubleTab +
	"DDDDDDDD$::~~:::~~~~:=7ZO8DDDD\n" + DoubleTab +
	"DDDDDDDD$77777I?I7777$DDDD8888\n" + DoubleTab +
	"DDDMMMMMMMMMMDDDDDDDDD8888OOOO\n"

const orcAscii string = DoubleTab + "MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMMZ?78MMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMO,:~=DMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMM$~?I++7I777$8MMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMD=:++,,::=~~~ZMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMM8:~=:,,:,:++=~?$MMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMZ:,:~~:,::,:~~~::IMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMI,:~~=?+++:,==::=~=OMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMM$.:=~:=+++~,~=~,:?+=MM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMM$~~+==???+~,+~.~+=~DM\n" + DoubleTab +
	"MMMMMMDDDDDDMMMMMD8D~:=???I??+~I~::::+MM\n" + DoubleTab +
	"MMMMMMMMMMMMMDD8Z7.,,.:+777?++=Z?=+~?MMM\n" + DoubleTab +
	"MMMMMMMMMDDMMMMMMDI+II,~++I+~~:+=+=IDMMM\n" + DoubleTab +
	"MMMMMMMMMMDOOO8DMMMMM$.:::=:,,:,:~+MMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMD8OOZZZZ+~:~=?~:::~~~DMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMDD=~~:~+~:,:=~?$8MMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMMM=+??=,,:::~IODDMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMMZ,~:~~:.,~=:OMMMMMMM\n"

const dragonAscii string = Tab + ".........,..................................................\n" + Tab +
	"...:~==I7$$$7?=,............................................\n" + Tab +
	"...,,,:?I7$O888Z+........................................=..\n" + Tab +
	"...........,~?ZDDO?~,...................................~+..\n" + Tab +
	"...............~~?$Z$I=:................................I:..\n" + Tab +
	"...................,+7$$I~.............................??...\n" + Tab +
	"......................:?$$I+:.......,,,...............II....\n" + Tab +
	"........................~+$Z7??I77I?7$7II?=~:,....,:=I=.....\n" + Tab +
	".........................?$7I?I77I7I:...,:~~=+++++==~.......\n" + Tab +
	".......................~$OZ7777$7777=:,:,...................\n" + Tab +
	"......................=O$?ZZ$7O$+IZOO77ZII~.................\n" + Tab +
	".....................,I7:..=Z$,....::..=?$7?+,..............\n" + Tab +
	".....................+I+,...,7,..........:8O$$?:............\n" + Tab +
	".....................=Z~....+Z,...........=DZI7I?:..........\n" + Tab +
	"......................,.....,~:............7+~,,~I~.........\n" + Tab +
	"..................................................I:........\n" + Tab +
	"..................................................,?,.......\n" + Tab +
	"...................................................,~.......\n"

const dragonAsciiB string = DoubleTab + "MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMDDMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMMDZI77I?++?$Z8DMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMM8MMMMMMMMMMMMMMMMZ?~:::::~7ODMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMZOMMOMMMMMMMMMMZ?~,,:~~==++ZMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMOI887DMMMMMMM8+..,:::::,,,::=$ODMMMMMMMMMM\n" + DoubleTab +
	"MDMMMMMMM8+::~+IOMMDI:,::::::~~=I$Z8DDDMMMMMMMMMMM\n" + DoubleTab +
	"MDZ===I7DO~,::::~+?,.,,,,,,::~~+8MMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MDOZZZ$+~:.,:,,::,+:..,:::::::,,:7MMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMM?,:.,?+7OI$M$=,.,~~~==+?I7Z8MMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMM7:~,,=IOZ7?++~,.,::=7ODMMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMM7::~:::::,....,:~=7DMMMMMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMO,:,:,~:,,:,~,,,:DMMM8?=+DMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMM$:~:::::~~:,..,~,~$8MMM8?~MMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMO.:~:,,,.:::,...,:,.:=?7?+.I8MMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMM8=.,:::,..::~:,...,:?~,....,?ZMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMZ:,,7I:,:,,::~,:~~:,+MMDDO$7OMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMM~,,,:~~:::::,,7?::::,~8MMMMMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMD?++=++7$8OZ$77ZMMDDD~.:ZMMMMMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMDDMMMMMMDMMMMMMMMMMMMD7I8MMMMMMMMMMMMMMMMMMMMMM\n"

const orcAsciiB string = DoubleTab + "MMMMMMMMMMMMMMMMMMMDD88MMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMO7Z=?+?MMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMM$++$MI,:~+=II7O8DDMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMDZ~,DMM+,,.,..,==:IOMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMO8$,$Z+..=:~==I$ODMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMDDO$8~+==+..~:.OMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMDDZI?=~??+??ZMM=,,,:+III7=ZDMMMMMMM\n" + DoubleTab +
	"MMMMD+?+7+,.~=:~==7I~==,,:7OOD=:MMMMMMMM\n" + DoubleTab +
	"MM$DD,~?I77=+=,,=~~~:I+..=$DMD,IMMMMMMMM\n" + DoubleTab +
	"MM$?8=,,~~=7=,=+::~::I~,,~:~+?$DDZMMMMMM\n" + DoubleTab +
	"MMMI+~.~?=+~,=+=?+:,,:..::==~+Z7$=MMMMMM\n" + DoubleTab +
	"MMMMDI=::~:=?~=I??=:,~..,~:::,~IZIDMMMMM\n" + DoubleTab +
	"MMMMO7=.:~,::+II7=.::,..,:,::.:DDMMMMMMM\n" + DoubleTab +
	"MMMMMDO?,:+::I7++=::....,..,:~:7MMMMMMMM\n" + DoubleTab +
	"MMMMD?~,,=OM7=:::ZMDO+$8M$~,,::=??8MMMMM\n" + DoubleTab +
	"MMMMM877DMMMMDOOMMMMMMMMMMMDOOO=.~OMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM~.=MMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMD$7OMMMMMM\n"

const sorcererAsciiB string = Tab + "++++?I77IIZOO888OOO$ZZZ$$ZZZZZZZZ$777IIIIII???++++\n" + Tab +
	"+++??I7$77Z88888888I78OOZ$ZOOOZZZ$$Z$$77I7II???+++\n" + Tab +
	"++??II7ZZ$$8DDDDDDD$I8888OZZOOOOOZ$ZZZ$7777III????\n" + Tab +
	"++??II7ZOOZZ8DDDDD8O+Z8O8DOZ88OOOZZZZOZZZ$77II????\n" + Tab +
	"++??II$ZOOOOZOO88O8DI+88II7OD8O88O$OZZ$$$777II????\n" + Tab +
	"+????I7ZZ$$ZOOZ8DDD8+=I7???$88DDDDOZ8$ZZ$$77II????\n" + Tab +
	"=++?II7$II$$$ZOOO8DO?==++++7$8D888D$7Z8ZZ$$$77II??\n" + Tab +
	"+??IIII?+7ZZZZZOZO88D$==+++7OOZOOOZZ7O$$$77I?I7III\n" + Tab +
	"+??I?+II=?7777$OZZZ$$Z7?+++=$OZZO8ZZ$7$OOZZ$I+IIII\n" + Tab +
	"+??+=I$I++IZZ$I77777777I+++=+7$Z$77$Z7II?7Z$7?+III\n" + Tab +
	"++==+II==+?77??I777$ZZI+++++++?$O$I$Z$I~~?$7II??++\n" + Tab +
	"+=::+?=~=???+:~II?I77I?++++++I7IIII7$$+~~~+?+++=~=\n" + Tab +
	"+~,,~+~~===?=~~=??=+?+++???+=?I+??I?=+=~~:=??+=~~~\n" + Tab +
	"=~:,:~~==:~+=~~==+~~+=++??+=~++=+=++==+~~~=??+~~==\n" + Tab +
	"==:::====:~=======~~=+==+?+===+++==+++=~~~+?+=====\n" + Tab +
	"::::,:::::::::::::::~~::~~~~=~~~~:~~~~::::~~~:~:::\n" + Tab +
	",,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,\n"

const skeletonAsciiB string = DoubleTab + "...............,,.......................\n" + DoubleTab +
	".............,,.........................\n" + DoubleTab +
	"...........,~~=~??=.,,..................\n" + DoubleTab +
	"............~+I:ONDZ,.....,,............\n" + DoubleTab +
	"..........,.I$Z$Z8DN8~....,,............\n" + DoubleTab +
	"...........?NDDO,,?I=+=,......~.........\n" + DoubleTab +
	"...........~ODD7,..::I8=......II+:,.....\n" + DoubleTab +
	"..........,..$N$I:,,~?~:,..:?7$$I,......\n" + DoubleTab +
	"..........,...?I+==++=,:+=++?=:+O+,.....\n" + DoubleTab +
	"..........:=I777ZO+=~~+II~,...=I=,......\n" + DoubleTab +
	".........?DNNDNDO8O$I+?+~,..............\n" + DoubleTab +
	"........:Z88OZOO88DZ8+:,,,::.,..........\n" + DoubleTab +
	"........~+?:=$DNNNND8=,,,,.,:...........\n" + DoubleTab +
	".........+~,..:IZ$ZD8O~,,,.,:...........\n" + DoubleTab +
	"........,,~:.......?I=:..,.,::.,,.......\n" + DoubleTab +
	"..........:~::,.....,,,:..,:::,,,,......\n" + DoubleTab +
	".............,...,,..:+=:,..,:,:,,,.....\n" + DoubleTab +
	"......................:~,,,.....,,......\n"

const goblinAsciiB string = DoubleTab + "MMMMMMMMMMMMMMMMMM88DMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMM$++7II$OMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMM8Z7=~==7I+~=7DMMMM\n" + DoubleTab +
	"MMMMMMMMMDMMM87=,~==~=7$ZMMMMM\n" + DoubleTab +
	"MMMMMMMM7$ZI7+=+++??+=7DMMMMMD\n" + DoubleTab +
	"MMMMMMMD=O?,=+II$II$O$IOMM8$Z8\n" + DoubleTab +
	"MMMMMMMM+$~,::~:~=8MMMDMMZI78M\n" + DoubleTab +
	"MMMMMMMD~$~:~~:+$OMMMMM87I8MMM\n" + DoubleTab +
	"MMMMMMMI.+I=,=?I$ZMM8O7ZZMMMMM\n" + DoubleTab +
	"MMMMMMM+.:7$?=?++II+~?$MMMMMMM\n" + DoubleTab +
	"MMMMMMM?,~+II=~=I8~,~7DMMMMMMM\n" + DoubleTab +
	"MMMMD7+,,~?ZD8=~~=7O8MMMMMMMMM\n" + DoubleTab +
	"MMDI~~:=+=ZMMMZ~:~IMMMMMMMMMMM\n" + DoubleTab +
	"M8:,~IIZ8MMMMM::~I8MMMMMMMMMMM\n" + DoubleTab +
	"M8$I?++=7DMMM7I$8MMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMDZ7I??I~~?$OOZDMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMD88OZ$$ZOOOOMMMMMMMMMMM\n"

const nightWalkerAsciiA string = Tab + "MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMDMMMMMDOZO8DMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMM7?MMO7+~?DMDOZ77ZODMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMD8DMMM$,,87:.:~..~=IODMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMM8OZ7Z?,:+=.?$M8?IDMMMMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMD$+~==:,O$8MOMMMMMMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMZ,~=,,:788IO8DMDMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMM+,~,~:,++?ZMMMDMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMO..,~::,=?7Z88MMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMD=+=,,:,~I7Z7O8MMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMMMO,,,,,:+I7ZZMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMMMI,:,,::=++IOMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMM$,~:,:~:.+7IZMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMM=~~:,::,.~ZO8MMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMM~:~:,~,,..7IDM8DMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMMI.,.:~:,:$O7DMDDMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMM$.,,=::.+DZ+$DDMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMM~,.:~:,.IODD$DDMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMI.,.:,,.~$8MMMMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMM~.,:::,:?ZDMMMMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMZ.,=~,~.II+ZDMMMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMD88OMMMMMMD,,=~,:..:87~$8DMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMD887=$ZDMM7,=:.....:8M77ZODMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMM88O7,?7+~:.....,.ID8O7?ZDMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMD8D8$:,::......,,.+Z8MDZZDMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMM8ZZO7DZ.....,::,~78MMM8DMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMM7...,+.:~,:Z8DMMMMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMM7,,.:8~,=:,ODDDMMMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMI,,.+8Z,~:,ZMDMMMMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMM?,,.7OM7.~,+MMMMMMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMI,,.ZMDM?.:,?8MMMMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMZ,,,ZDMMMI,,,,+DMMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMM:,.?OOMMMZ~...~DMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMDMMMMMMMMMMMMM$..,$ZZ8MMM?...=DMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMO?OMMMMMMMMMMMMZ..+$$ODZ8MO~..:7DMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMD+DMMMMMMMMMMMM7.:,I$O77DMM$...=ODMMMMMMMMMMMMMMM\n" + Tab +
	"MD+7OD$,$MMMMMMMMMMMDO.,~??II8MMMMZ..:$OMMMMMMMMMMMMMMM\n" + Tab +
	"MZ:,:,.+MMMMMMMMMMMMMM,.?$IZ8MMMMMMI.,$ZDMMMMMMMMMMMMMM\n" + Tab +
	"MM8~,:~7IZMMMMMMMMMMMD,.?OO8DMMDMMMM:.+77ODMMMM8MMMMMMM\n" + Tab +
	"MMMD=~==~8MMMMMMMMMMM8,,+$Z8MMM8MMDMZ.,I7$$8DMMDMMMMMMM\n" + Tab +
	"MMMMM=:::~OMMMMMMMMMMZ.:+?7Z8ODDZO8MM+.=Z$ZZ$$OODDDMMMM\n" + Tab +
	"MMMMM=:,,~~IDMMMMMMMM~.:?+?I77ZD8ZO$Z+.,?$77I7OO8DDMMMM\n" + Tab +
	"MMMMM7,?,,+,ZMMMMMDO?.,,+??I7I7Z$Z$7$~..~I?+?I$$ZDMMMMM\n" + Tab +
	"MMMMM7,:=::$MD8ZOZ=...,~::~~=+?=~=I77+..,++=++I7$ODDMMM\n" + Tab +
	"D8O$I~::?::I+=~::::::::,,..,,,,,:,,~~:...:~~====I$ZOOMM\n" + Tab +
	"MDOI,=I==~==~~:::~:~~:::::~::=~==~~:,~=+++++=:,~+++7ODM\n" + Tab +
	"MD8OO8OO8DDMDMDDMMMMMMMDDDD88D8888888DMMMMDMMDDDD888MMM\n"

const nightWilkerSkillClawAsciiA string = Tab + "MMMMMMMMMMMMMMMMMMMMMMMMMZ?..:ZMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMM..~:MMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMM?...DMMMMMMMMMMMMMMMDDDMMMMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMZ....$MMMMMMM$77?=+?:,:=+?MMMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMMMO??....,.7MMMD$+?I==+I7Z?,8M++MMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMM$:,.....+..=M$=+~+?Z88M8?=8MMD:MMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMM+.....,......:.~:~I7I~..=ZMMMMO~MMMMMMM\n" + Tab +
	"MMMMMMMMMMMM8$I7777+::~.7I........,.$MMMMM8.DMZ+~OMO~MMMMMMM\n" + Tab +
	"MMMMMMMMMMMMZ?+,.....:?8MM~....,..,.ZMMMMD~$Z,?O8MM$~DDDDMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMOZZZODMMMMM7=.......+MMMM~IM=$MMMMM+~II77I?+\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMM7.~$+,I78MMM7~MMO7.:MMM~ZMOZ$$8O\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMD,8$?ZMMMMMMO,DZ..=78MMD,DZ,$ZZMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMM7=M+$M$DMMMM:ZM:+8MMMMM+?M~ZMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMZ:M7=MI=Z+8M:MMZ7I?7MM$,MO:MMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMM+ID=MM:+~=Z,7+?ZDD=IZ~MMI?MMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMMD,DMMMODM?.+$DMMDMI,+I7$:$MMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMZ:OMMMMZ:8MMMM8:M:OM8ZZ,8MMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMO,8MMO:OMM7ZMM:8:MMMMD:MMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM8,DM,8MD=+DMO~7=MMMMO~MMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM+?Z=M$:ZMMM+=?IMMMMZ~MMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMD~.OM?DZ~M8.7:MMMMM7=MMMMMMM\n" + Tab +
	"MMMMMMMMMMMMMMMMMDDDMMD8888MMMMMMZ.DMMM,OM=,+IMMMMM??MMMMMMM\n" + Tab +
	"MMMMMMMMMMMMDDOOOOO7IZOOOOZ7O8D88D.DOMZ~MD..+MMMMMM??MMMMMMM\n" + Tab +
	"MMMMMMD8OZ77$O8OOZ7??7$7IIII??IIO+?M?Z$$Z=~:DMMMMD8I?DMMMMMM\n" + Tab +
	"MMMM8$ZZ7I7ZDD7I??I$$I?IIIII?I?$O:M7?????ZM,DMMOZ$?+7DMMMMMM\n" + Tab +
	"MMMM?778O77IIZ?$IIIIIZO77$?IZMDI:7$$?7ZIDM:7MO$OMMD$$MMMMMMM\n" + Tab +
	"MMMZ7O$O$?I?7I?I?????I7$$MDMMZ7=+OZ77IM8O=.$$OMOZZZ7?MMMDD8M\n" + Tab +
	"MMM$8?I??II?II?I??I77?7$ODZI?7?IO8?IMZ77.+7?I$$I?????7$$Z77M\n" + Tab +
	"MMDI7IO7I$I???I?II+II+?:,.~7$+=II$.+D8$I77+:+=:=?II??=+==::?\n" + Tab +
	"O??=~+$$$I?I??I+?I+=+7I==~,=7O7Z8+.8M~IZDMM8:II++I77Z$O887IO\n" + Tab +
	"MMMMMDOZII?ZIIZI?777D$I??7OMM8ZO~.DMZ,+??IZMZ,D87IIII77$Z7MM\n" + Tab +
	"MMMMMMMM7$7II??I7$$7I.+?III??I7.:DZ==ZODDO7???8MZI7I777+??OM\n" + Tab +
	"MMMMMMMMD8M8I?II$7$77??Z888OZ8M?:~=7$II$ZO$7I7II77I$777Z$7DM\n" + Tab +
	"MMMMMMMMMM$M77I777$I???IIII?I?I$?IZ$??II?++II$77$I?I?I7DMMMM\n" + Tab +
	"MMMMMMMMMM?DZ$??I?I?IO8?7?$$?II7$7?II??I7$$$OO$I$O7ZZ$$$MMMM\n" + Tab +
	"MMMMMMMMMM?I8+~=O$?$$OO7ZO$I7IIIMMO??I7$$MMD8DZOM$ZMMMMMMMMM\n" + Tab +
	"MMMMMMMMMM+==+$$?$?ZMDDMMMMOZO8MMMMOZ$Z$DMMMMM$Z$8MMMMMMMMMM\n" + Tab +
	"MMMMMMMMMI~==+8?=ZM+MMMMMMMMMMMMMMMMMMDMMMMMMMMDMMMMMMMMMMMM\n" + Tab +
	"MMMMMMMD7II=??++$D7$MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" + Tab +
	"MMMMMMM$+?I~ZMMZZ87MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n"

const GameOverAscii string = "\nNNNNNNNNNNNNNNNNNNNNNNNNNNNNMMMMNNMMNNNMMNNMMMMMNMMMMMMMMNNNNNNNNNNNNNNNNNNNNNNN\n" +
	"NNNNNNNNNNNNNNNNNNNNNNNNNNDDMMNNNNDNMNN88NMNN8DNNNMMMMMNNNNNNNNNNNNNNNNNNNNNNNNN\n" +
	"NNNNNNNNNNNNNNNNNNNNNNNNO=,,=ONND=~?$NDI,:7DDI,=ZI++?IONNNNNNNNNNNNNNNNNNNNNNNNN\n" +
	"NNNNNNNNNNNNNNNNNNNNNNNN?.IZ,.ZND:.~:8NZ,.,8NI..+O,:+~INNNNNNNNNNNNNNNNNNNNNNNNN\n" +
	"NNNNNNNNNNNNNNNNNNNNNNNN$.INZ+$ND=,~.ZMO,..IN=..+8,~NMMNNNNNNNNNNNNNNNNNNNNNNNNN\n" +
	"NNNNNNNNNNNNNNNNNNNNNNNN8,~O+=?DN+,Z::DN+.~.Z?,:.O?.?$ONNNNNNNNNNNNNNNNNNNNNNNNN\n" +
	"NNNNNNNNNNNNNNNNNNNNNNNNZ.+D+.=DN+.I+.IN$.?~:+:?.7Z..::ONNNNNNNNNNNNNNNNNNNNNNNN\n" +
	"NNNNNNNNNNNNNNNNNNNNNNNND=,ON+,$N+.,:.,88,~Z,.:Z,=8::8DDNNNNNNNNNNNNNNNNNNNNNNNN\n" +
	"NNNNNNNNNNNNNNNNNNNNNNNNMO.,7=.IN~.$N?.?D:.8I.~8:,8+.ZNDNNNNNNNNNNNNNNNNNNNNNNNN\n" +
	"NNNNNNNNNNNNNNNNNNNNNNNNNNO+:~IDD+=8M8=?D7+8D+IN?:OI.,:,+NNNNNNNNNNNNNNNNNNNNNNN\n" +
	"NNNNNNNNNNNNNNNNNNNNNNNNNNMMNNMNNNNNNNMNNMMNNMNNNDDDOZZ$ZNNNNNNNNNNNNNNNNNNNNNNN\n" +
	"NNNNNNNNNNNNNNNNNNNNNNNNNNNNMMMMNNMMNNNMMNNMMMMMNMMMMMMMMNNNNNNNNNNNNNNNNNNNNNNN\n" +
	"NNNNNNNNNNNNNNNNNNNNNNNNNNN8I=+$NN7I8NN7+OO++++IDO$ZODNMNNNNNNNNNNNNNNNNNNNNNNNN\n" +
	"NNNNNNNNNNNNNNNNNNNNNNNNNN8:,+,.7N~.$MO,,DI.~+==8+.=+,+$DNNNNNNNNNNNNNNNNNNNNNNN\n" +
	"NNNNNNNNNNNNNNNNNNNNNNNNNNI.7MI.?N+.$N?.7N=.OMMND:,8DO~.ZMNNNNNNNNNNNNNNNNNNNNNN\n" +
	"NNNNNNNNNNNNNNNNNNNNNNNNNN=.ON=.$N?.$D,:DD:.7Z8NO.=ND=~$DNNNNNNNNNNNNNNNNNNNNNNN\n" +
	"NNNNNNNNNNNNNNNNNNNNNNNNNN=:8N=,ONI.77.INO,.,,7M$.:I+.?MNNNNNNNNNNNNNNNNNNNNNNNN\n" +
	"NNNNNNNNNNNNNNNNNNNNNNNNO:?DZ~IO8N$.+~,8NZ.+DDDNI.~:.+DNNNNNNNNNNNNNNNNNNNNNNNNN\n" +
	"NNNNNNNNNNNNNNNNNNNNNNNND$,~OI,~DNO,,.+NN$.?NNNN?.Z8.~NNNNNNNNNNNNNNNNNNNNNNNNNN\n" +
	"NNNNNNNNNNNNNNNNNNNNNNNNNMI.,:.INND~..$NN7..::~8+.ZN:.ZMNNNNNNNNNNNNNNNNNNNNNNNN\n" +
	"NNNNNNNNNNNNNNNNNNNNNNNNNNN8$$8NNNNO$$DNN8ZZ$$$DZI8N$~$NNNNNNNNNNNNNNNNNNNNNNNNN\n" +
	"NNNNNNNNNNNNNNNNNNNNNNNNNNNMMMMNNNNMMMNNNMMMMMMNMMNNNNNNNNNNNNNNNNNNNNNNNNNNNNNN\n" +
	"DDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDNDDDDDDDDDDDDDDDDDDD\n"

const wizardAscii string = DoubleTab + "MMMMMMM8DMMMMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMI$MMMMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMM8O8?7MDMDMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMM$~7=~?Z==+?OMMMMMDDMMMMMMM\n" + DoubleTab +
	"MMMMM8ZI,ZZ:++:,MMMMMDO8MMMMMM\n" + DoubleTab +
	"MMMMMMMM:+=~+I=,ZMMMMOZOMMMMMM\n" + DoubleTab +
	"MMMMMMMM7,=~=7+,~MMMM88D8MMMMM\n" + DoubleTab +
	"MMMMMMM8?$++I=?:~$MMM88IOMMMMM\n" + DoubleTab +
	"MMMMMM8~II~=?+~:~.?M87$8MMMMMM\n" + DoubleTab +
	"MMMMMD=+I.:.??~::.~?~IMMMMMMMM\n" + DoubleTab +
	"MMMMMD==:.:,,,....~,.,IMMMMMMM\n" + DoubleTab +
	"MMMMMM+,,.:,......,,...~MMMMMM\n" + DoubleTab +
	"MMMMMO~:~,,:.,......,::.=DMMMM\n" + DoubleTab +
	"MMMMM?:~.=,::,.......~?ZZZMMMM\n" + DoubleTab +
	"MMMMM~:~.+=.~,........,:I8MMMM\n" + DoubleTab +
	"MMMMD::~,==.~:,.......,...?MMM\n" + DoubleTab +
	"MMMMO,::,=~.:~:,.......:I7?7MM\n" + DoubleTab +
	"MMMMO,::.~~,.:=~...,...,8MMMMM\n" + DoubleTab +
	"MMMMO:~~..,,.:=+:..:....,8MMMM\n" + DoubleTab +
	"MMMM8:~~:.,..~~+=:.:,....,MMMM\n" + DoubleTab +
	"MMMM8........~==~~..:...:7MMMM\n" + DoubleTab +
	"MMMM8........:=:=~:..,$8MMMMMM\n" + DoubleTab +
	"MMMMM,.,..,..,=:~:,~$=MMMMMMMM\n" + DoubleTab +
	"MMMMM$~:.,O.~I,,:?IOMDMMMMMMMM\n" + DoubleTab +
	"MMMMMMD:.7M8MM~:,~?ZMMMMMMMMMM\n" + DoubleTab +
	"MMMMMM7:?MMMMMMM7OO8MMMMMMMMMM\n"

const wizardAsciiB string = DoubleTab + "MMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMM8$8M8MMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMZ,:~+MMMMMMMMDMMM\n" + DoubleTab +
	"MMMMMMMMMMDO$~~+:IMMMMMMM8OMMM\n" + DoubleTab +
	"MMMMMMDMMD?,..==,ZMMMMMMD~7MMM\n" + DoubleTab +
	"MMMMM8MOMD7:::+~~$8MMMMM8?DMMM\n" + DoubleTab +
	"MMMMMI?,$MZ:~~=+=,=MMMMM7OMMMM\n" + DoubleTab +
	"MMMMMM8?.7=::::~::=MMMMMIMMMMM\n" + DoubleTab +
	"MMMMMMMM~.,,,::~:,,OMMMZ$MMMMM\n" + DoubleTab +
	"MMMMMMMM8:,,,,,:,.,:?ZO+DMMMMM\n" + DoubleTab +
	"MMMMMMMMZ:,,,,,,,,:$$I++MMMMMM\n" + DoubleTab +
	"MMMMMMMM+:,,::::~:.~MMZ8MMMMMM\n" + DoubleTab +
	"MMMMMMMZ~:,,:~~~,=:.?MIMMMMMMM\n" + DoubleTab +
	"MMMMMMD=~,,,,,::.~:.+OZMMMMMMM\n" + DoubleTab +
	"MMMMMM+~:,,,,,,,,,..+$DMMMMMMM\n" + DoubleTab +
	"MMMMD+~=:,,,,,,,.,,.~ZMMMMMMMM\n" + DoubleTab +
	"MMMO=~=:,,,.,.,,.,..~DMMMMMMMM\n" + DoubleTab +
	"MMMD?~~,,:,,,.,.,,..+MMMMMMMMM\n" + DoubleTab +
	"MMMO~~:,,,,:,,,,::..7MMMMMMMMM\n" + DoubleTab +
	"MMM$:,,,,...,,,...:IMMMMMMMMMM\n" + DoubleTab +
	"MMMMM7,,....,,,,..+ZMMMMMMMMMM\n" + DoubleTab +
	"MMMMMM8+.,~==:.,.:78MMMMMMMMMM\n" + DoubleTab +
	"MMMMMMM~,8MMM8ZZ.+$MMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMO.7MMMMMMD.=8MMMMMMMMMMM\n" + DoubleTab +
	"MMMMMM+.$MMMMMMZ.:MMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMM+,ZMMMMMMZ.~MMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n"

const wizardAsciiC string = DoubleTab + "ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ\n" + DoubleTab +
	"ZZZZZZZZZZZZZZ$I7ZZZZZZZZZZZZZ\n" + DoubleTab +
	"ZZZZZZZZZZZZZ7,,,IZZZZZZZZZZZZ\n" + DoubleTab +
	"ZZZZZZZZZZZZZ~~O?,$ZZZZZZZZZZZ\n" + DoubleTab +
	"ZZZZZZZZZZZZZ::7+.?ZZZZZZZZZZZ\n" + DoubleTab +
	"ZZZZZZZZZZZZ$:.~,.:$ZZZZZZZZZZ\n" + DoubleTab +
	"ZZZZZZZZZZZZII$Z$:~?ZOOZZZZZZZ\n" + DoubleTab +
	"ZZZZZZZZZZZZ7O7$$?,,:===ZZZZZZ\n" + DoubleTab +
	"OZOZZZOZZZO7?$8ZI??:~~,7OZOOOO\n" + DoubleTab +
	"OOOOZZZZOO$?I=I?~78$I$ZOOOOOOO\n" + DoubleTab +
	"OOOZ$777$77I=I=$I~~Z8OOOOOOOOO\n" + DoubleTab +
	"OOO$III??IZ?IO:O8~:78OOOOOOOOO\n" + DoubleTab +
	"88O$III?II$ZO$:7M$$Z8888888888\n" + DoubleTab +
	"88OZ7IIII7O$Z7:?DO8$8888888888\n" + DoubleTab +
	"DD88Z$$ZZO8$Z$:+OO8$OD8DDDDD8D\n" + DoubleTab +
	"DDDDD888DD8O$7:+$O8ZODDDDDDDDD\n" + DoubleTab +
	"DDDDDDDDDDOO?+~~+O8ZZ8DDDDDDDD\n" + DoubleTab +
	"DDDDDDDDDDO8+,:,,8D$ZODDDDDDDD\n" + DoubleTab +
	"MMMMMMMMM8O8=,~~~Z8Z$O8MMMMMMM\n" + DoubleTab +
	"MMMMMMMMM$I?::=7+7+=?ZZDMMMMMM\n" + DoubleTab +
	"MMMMMMMMMO++::~I~?+,,:7DMMMMMM\n" + DoubleTab +
	"MMMMMMMMMM$=,,:,,,.:=I8MMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMO+:.,:.IMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMD8Z$?,.,:.?7$ODMMMMMM\n" + DoubleTab +
	"MMMMMMMMMM8ZI++??++I$O8DMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n"

const wizardAsciiD string = DoubleTab + "ZZZZZZZZZZZZZZZZZZZZ\n" + DoubleTab +
	"ZZZZZOZZZZZOZZZZZZZZ\n" + DoubleTab +
	"ZZZZZ$OZZOZ7ZZZZZZZZ\n" + DoubleTab +
	"ZZZZO=+ZZ$7~~ZZZZZZZ\n" + DoubleTab +
	"ZZZO$:+ZZ=++,?OZZZZZ\n" + DoubleTab +
	"ZZZO$:$OZ+II=,7OZZZZ\n" + DoubleTab +
	"ZZZO$=O$+?Z7+:~=ZOZZ\n" + DoubleTab +
	"ZZZOI=O=~+~====~IZZZ\n" + DoubleTab +
	"ZZZO?IO==+=??+??=IOZ\n" + DoubleTab +
	"ZZZZ==?=I++?+?=:=IOZ\n" + DoubleTab +
	"ZZZO==~~++=+++:I=:$O\n" + DoubleTab +
	"ZZZZ=ZO=:~~=?=:I+.7O\n" + DoubleTab +
	"ZZOZ=ZI~~~~=++:+I,$O\n" + DoubleTab +
	"ZZO$=7::~~+?++~=?:$O\n" + DoubleTab +
	"ZZO7+I:~==++++~~=?ZZ\n" + DoubleTab +
	"ZZOI+O+:+=++++~:=?OZ\n" + DoubleTab +
	"ZZOI?O=:~++=?~=~:IOZ\n" + DoubleTab +
	"ZZO?IO~,==~=+:::=ZZZ\n" + DoubleTab +
	"ZZO?7O=,~::~=~::=ZZZ\n" + DoubleTab +
	"ZZO+7O~:~,,~=::~:ZZZ\n" + DoubleTab +
	"ZZO+$Z~:::~=:,:~:$OZ\n" + DoubleTab +
	"ZZO+$7::::,,:::~:?OZ\n" + DoubleTab +
	"ZZO+ZZ~,~,,::~::,+OZ\n" + DoubleTab +
	"ZZZ+ZO$=:.+?:~,,?ZZZ\n" + DoubleTab +
	"ZZZ=$$?::=?7II~.~ZZZ\n" + DoubleTab +
	"ZZZZZZ77ZOZZZZZ7IZZZ\n"

const thiefAscii string = "MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMMMD8MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMMO~,7MMMMMMMMMMMMMMMOI?OMMMMMMMMMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMMI.,7MMMMMMMMMMM8$Z7=~::IMMMMMMMMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMM$..:IDMMMMMOII7?+I++=~~:MMMMMMMMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMM+$=..~IO88?~~~:,+?::I7::OMMMMMMMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMO?MM=..:=~:,,::,.:,.:=:.~~=OMMMMMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMM7$MMMO+,....,,..::,,,..,:=~8MMMMMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMM?MMMMMMD$7$$8,.,,~::,,,~~=7MMMMMMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMDZMMMMMMMMMMO$=,:,:~::,:~:::$MMMMMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMDMMMMMMMMMMMO+:,,,:~::::,:.=MMMMMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMMMMMMMMMMMMMM8==~,,:,,.,,::+MMMMMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMMMMMMMMMMMM8?I,,==::~:,:~~~+DMMMMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMMMMMMMMM8$7?:.,,,.,,,:,...,~DMMMMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMMMMMMMMD+~,....,,,....,=??78MMMMMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMMMMMMMMM=,::::,,.....,,:IODMMMMMMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMMMMMMMI7+=::::,,,.,..,,,,,:=ZMMMMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMMMMMM7.,:~:.,,........,,,,,,:IDMMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMMMMMM:.,,,,,..~$I$7:.....,:,,.,8MMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMMMMMM~.,,...:OMMMMMD++=~,...,,.IMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMMMMMM:.,,~Z8MMMMMMMMM$?I+$7,...:MMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMMMMMD,.,.=MMMMMMMMMMMMMDMMMI....7MMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMMMMM8..,~OMMMMMMMMMMMMMMMMMMO+...DMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMMMMMZ..:DMMMMMMMMMMMMMMMMMMMMM$..~DMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMMMMM?.,OMMMMMMMMMMMMMMMMMMMMMMMZ...7MMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMMMMO..?MMMMMMMMMMMMMMMMMMMMMMMMM:....?8MMMMMMMM\n" +
	"MMMMMMMMMMMMMMM$,..?MMMMMMMMMMMMMMMMMMMMMMMMMDOZ+:,~ZMMMMMMM\n" +
	"MMMMMMMMMMMMM7....~8MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMM8++I7ZMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n"

const thiefAsciiB string = DoubleTab + "MMMMMMMMMMMMMMM8DMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMD~:=ZMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMM?.==:8MMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMZ~.:~:$MMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMM8=,..,:+~IMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMM=,=::~::~=MMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMD:~:.,,,:::OMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMI.~,.,,.,,.=MMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMM:.:,..,,..,:DMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMM?.~,..,:,.,,D8ODMMMMM\n" + DoubleTab +
	"MMMMMMMM8D=:,,,::,~~,I=,7DMMMM\n" + DoubleTab +
	"MMMMMMMM87?,:.,:,,,~?ODOMMMMMM\n" + DoubleTab +
	"MMMMMMMMMD+.:..::~,IMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMO$I+?+::~=,:7MMMMMMMMMMM\n" + DoubleTab +
	"MMMMMM8OO8MMZ=7?,.OMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMZ.~II~OMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMM~,~,,~?DMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMD,,:,.?O7ODMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMM=.:+.:ZDO$ODMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMI.~M8?~ZMMDDMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMM=,=MMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMI.,ZMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMM$~7MMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n"

const thiefAsciiC string = DoubleTab + "MMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMO$$OMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMM$~7?=DMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMI?$+$MMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMI:~.+$ODMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMDI=~~:,.,:$MMMMMM\n" + DoubleTab +
	"MMMMMMMMMMM7,~~~:,,:.?MMMMMM\n" + DoubleTab +
	"MMMMMMMMMMM+.,,,....:+MMMMMM\n" + DoubleTab +
	"MMMMMMMMMMD,......?~.,8MMMMM\n" + DoubleTab +
	"MMMMMMMMM8=.,,...?MO..IMMMMM\n" + DoubleTab +
	"MMMMMMMMM+,,,,...IMZ.,~MMMMM\n" + DoubleTab +
	"MMMMMMMMM~.,:~:~~:DZ,:ZMMMMM\n" + DoubleTab +
	"MMMMMMMM$......,,.O7:+ZZMMMM\n" + DoubleTab +
	"MMM888O$:..,,:....:,,?OMMMMM\n" + DoubleTab +
	"MMMO7ZZZZ~.,::.....:$MMMMMMM\n" + DoubleTab +
	"MMMMMMMM8..::,.....:MMMMMMMM\n" + DoubleTab +
	"MMMMMMMMI.::,.,.,.,,8MMMMMMM\n" + DoubleTab +
	"MMMMMMM$,:::...+:.:,+MMMMMMM\n" + DoubleTab +
	"MMMMMMM=,:~:.,:78,..$MMMMMMM\n" + DoubleTab +
	"MMMMMMMZ,,$~..+MMZ..=DMMMMMM\n" + DoubleTab +
	"MMMMMMMM?~M=..=MMD.,:ZMMMMMM\n" + DoubleTab +
	"MMMMMMMMD8M7..$MMM?.,7MMMMMM\n" + DoubleTab +
	"MMMMMMMMMMM$.~MMMMM:.7MMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMI.+MMMMM$.+MMMMMM\n" + DoubleTab +
	"MMMMMMMMMMO~~=MMMMMD,~8MMMMM\n" + DoubleTab +
	"MMMMMMDO7=,..~ZOOOZ7~:8MMMMM\n" + DoubleTab +
	"MMMMMM8Z$$$ZO888DD8+:=8MMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMM\n"

const thiefAsciiD string = DoubleTab + "MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMOZ8MMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMM?~~=$MMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMO=+77?DMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMDI:7$$IOO8MMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMM7~::+7?+?=?MMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMM$~:::~=I?~::+ZMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMD~?=:~+7I$++~+,~$MMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMZ,:~=?I=7ZI+?=~7?:IDMMMMMM\n" + DoubleTab +
	"MMMMMMMMM8~~==~:::~::::?DM8?78MMMMM\n" + DoubleTab +
	"MMMMMMMMMOI7+:,,,,.,~=:+OMDOO?DMMMM\n" + DoubleTab +
	"MMMMMM8$77?:,,,,:,,:~7?=$7IMMDMMMMM\n" + DoubleTab +
	"MMMMDZ$OMM+:::,:::,:~+?~=ZMMMMMMMMM\n" + DoubleTab +
	"MMMMDDMMMO:::::,,=+??:~~=ZMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMM?:~:::,,,~~,~:??~MMMMMMMMM\n" + DoubleTab +
	"MMMMMMMM8:~:~~~~,,,.,~~+?:8MMMMMMMM\n" + DoubleTab +
	"MMMMMMMM+,~==~::,,:,,:+I$~8MMMMMMMM\n" + DoubleTab +
	"MMMMMMMZ:~=+=:,,,:::=:~+77MMMMMMMMM\n" + DoubleTab +
	"MMMMMMD=~?=:,,,::::,78.=I?DMMMMMMMM\n" + DoubleTab +
	"MMMMMM7:~:,,,,::~~~:7M=:Z7IMMMMMMMM\n" + DoubleTab +
	"MMMMM8::::,,::::==+:$M8=~+=+8MMMMMM\n" + DoubleTab +
	"MMMMM=,,:,,::::~=++~ZMMMI.~+=OMMMMM\n" + DoubleTab +
	"MMMMI,,,,,:~:~~=+?+=OMMMM~.:~+MMMMM\n" + DoubleTab +
	"MDD$::,,,~~~~~~+??++OMMMMD+.::DMMMM\n" + DoubleTab +
	"M87==:,,:~~~~~=++=~=?Z8MMMMI.,OMMMM\n" + DoubleTab +
	"MMMD7:.,7$$77I777$Z8DMMMMMMM?.+MMMM\n" + DoubleTab +
	"MM8+::+7ZZZZOOOZOO88DDMMMMMM8,,OMMM\n" + DoubleTab +
	"MMD$ODMMMMMMMMMDD88OOOOOOOOOO::?MMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMMMMMMDD8OO=,=MMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMO=IMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n"

const paladinAscii string = DoubleTab + "........................................\n" + DoubleTab +
	"...................::,..................\n" + DoubleTab +
	"..................=???..................\n" + DoubleTab +
	"..................+I?Z~.................\n" + DoubleTab +
	".................,?$Z$?.................\n" + DoubleTab +
	"................,=+?7$I~................\n" + DoubleTab +
	"..............,::7===?+7,...............\n" + DoubleTab +
	".............,~~:~7?++~=+,..............\n" + DoubleTab +
	".............=~~~,7?==~,~=~:,...........\n" + DoubleTab +
	"............:=~~:+I?==~~..~,............\n" + DoubleTab +
	"............~=~~~++?+?==~...............\n" + DoubleTab +
	"............===~=I=+?7+??,..............\n" + DoubleTab +
	"............=====I77I7?I?:..............\n" + DoubleTab +
	"............~+==..7+77+,,...............\n" + DoubleTab +
	".............~~...~=?I,.................\n" + DoubleTab +
	"..................,=++..................\n" + DoubleTab +
	"..................:~=,..................\n" + DoubleTab +
	"..................:~,...................\n" + DoubleTab +
	"..................~=,...................\n" + DoubleTab +
	"...................~~...................\n" + DoubleTab +
	"...................==...................\n" + DoubleTab +
	"...................:,...................\n" + DoubleTab +
	"........................................\n"

const paladinAsciiB string = DoubleTab + "MMMMMMMMMMMMMM8ODMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMD$OM7,.I,$MMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMD.IM7,.=.=MMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMM.$MO..,..8MMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMD.$MD:.~,,7Z87IOMMMMMM\n" + DoubleTab +
	"MMMMMMMDZ??OO7:.~7OZZ7~ZMMMMMM\n" + DoubleTab +
	"MMMMMMZ~=,.IZZ?==~?Z+:IMMMMMMM\n" + DoubleTab +
	"MMMOO$=::::ZZ8Z:I?I=:.=MMMMMMM\n" + DoubleTab +
	"MMMMM7..,MI.:=7.:=Z7?,8MMMMMMM\n" + DoubleTab +
	"MMMMMD:.,M7...~~ZODM$7MMMMMMMM\n" + DoubleTab +
	"MMMMMM8+~M$..:,=I7$+8MMMMMMMMM\n" + DoubleTab +
	"MMMMMMMM?DZ..7,~,.+=DMM8MMMMMM\n" + DoubleTab +
	"MMMMMMMM=DZ.II~~?7IZMD+?$DMMMM\n" + DoubleTab +
	"MMMMMMMD,DZ.=8+~$M,.?~.,?DMMMM\n" + DoubleTab +
	"MMMMMMMO.DO..D+~DZ:$~,,$MMMMMM\n" + DoubleTab +
	"MMMMMMM?.8O..D+~M+.M7.ZMMMMMMM\n" + DoubleTab +
	"MMMMMMM=.8O..8+~M:.8I.DMMMMMMM\n" + DoubleTab +
	"MMMMMMO$.O8..O+~D:.87.DMMMMMMM\n" + DoubleTab +
	"MMMMMMM8.Z8..O+~M~..7ZZMMMMMMM\n" + DoubleTab +
	"MMMMMDMO.$D..O+~M=...OMMMMMMMM\n" + DoubleTab +
	"MMMMMDM$.$D..D=~M=...~8$$MMMMM\n" + DoubleTab +
	"MMMMMMM=.7D.:M=~M?.Z?.?++MMMMM\n" + DoubleTab +
	"MMMMMMD..?M.~M~~DOOMZ.+M7MMMMM\n" + DoubleTab +
	"MMMMDM=,+7D..ZI?8MMMM.?MZMMMMM\n" + DoubleTab +
	"MMMMDZ+DMMD..OMMMMMMM.+MZMMMMM\n" + DoubleTab +
	"MMMMDZMMMMD.IMMMMMMMM+:M7MMMMM\n" + DoubleTab +
	"MMMMDMMM7~8.$MMMMMMMMI.7?8MMMM\n" + DoubleTab +
	"MMMMM8+,..I.IMMMMMMMM=~O8ZDMMM\n" + DoubleTab +
	"MMMMMDO88DDDMMMMMMMMMMMMMMMMMM\n"

const paladinAsciiC string = DoubleTab + "OOZZOZ777I??7$$I?????+,,,,,7?====+??+?+=\n" + DoubleTab +
	"ZZ$$ZZZZZZZZOOOOZ777I:,....=7+++++?++?+=\n" + DoubleTab +
	"Z$7$777I7$ZO888D8OOOO~.....+$????+++???+\n" + DoubleTab +
	"ZOOOOOOOOOOO8888DDDDO+,.,,:+IIII??++?+==\n" + DoubleTab +
	"OOZZZZOZZ888888DO$7+~:::,,,~==~=++II??++\n" + DoubleTab +
	"OZ$I7ZZ$$$O88DD8+,:,,..~,..,,,,,,,?Z$???\n" + DoubleTab +
	"777I?++77I$ZO8OI~.,,,........,,..,:+$7??\n" + DoubleTab +
	"++?III~??II$$I?=~:..........,...:,,~?7II\n" + DoubleTab +
	"+??II7.=Z$$Z=:I=:~=:.........,:~:.~~,$II\n" + DoubleTab +
	"?IIII7,~OZZZ~.~~.::~:.:,...+?=~...,..+$I\n" + DoubleTab +
	"III7$$,~O$ZI.,.:,.,.~,=I:.:=+~:,.....~$I\n" + DoubleTab +
	"I7$ZOO.~8Z7=,,.~~.::+:..,....,,......?$7\n" + DoubleTab +
	"7$$OO:.~I+,::..=+.,:,~=?,..~=,.....,,+O8\n" + DoubleTab +
	"O88D$.....,..:7OZ=.,:,~I~.:??=,..Z~.,~8D\n" + DoubleTab +
	"OZ7II=....,:?DD88Z..............IM=.:,ZD\n" + DoubleTab +
	"OZZO8Z.:++78M8OZO=.............~=MI.,.78\n" + DoubleTab +
	"Z8OOO$.IDD8OZZZZI.,:,......,,,.~,8$...7O\n" + DoubleTab +
	"$ZZZO$.?ZZZZZOOZ,.,,=+~~,......,.7$..,$Z\n" + DoubleTab +
	"77$$Z7.?OO88888I..,.,,~::.,~,:...++...I7\n" + DoubleTab +
	"7777$7.?OOO8888:..,...,,,:::+7?:.:~,.,+?\n" + DoubleTab +
	"$77$Z7.+888D8MI...,.,,,,,~:::=:=:,...:+=\n" + DoubleTab +
	"ZZ$7$7.7DDDDDO,.....,,,....+::~~......==\n" + DoubleTab +
	"$$$$77,?888OO?......,.,....:=::,,..,:~=+\n" + DoubleTab +
	"I7$$77,+Z$Z7I....,:.,,,,....~::,,..~I?++\n" + DoubleTab +
	"=?II7I,=I?I7I=,.,:~.,.:,.....:.,:~,,+===\n" + DoubleTab +
	":~~~==,~=~~+?=+.,:~...:.,....:~:~~~.=:,:\n" + DoubleTab +
	",,,,::,~~::~~~~.,::.....,.....::,,=::~,,\n" + DoubleTab +
	",,,,,,~~::::~::.,,,.....,........,~:.~,.\n"

const paladinAsciiD string = DoubleTab + "MMMMMMMMMMMMMMM8O8MMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMO+?=ZMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMM?:~.+MMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMM7~=,?MMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMZ?::~~+?DMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMM7+:::,:,,+7OMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMI:~~7++,,,:77ZMMMMMM\n" + DoubleTab +
	"MMMMMMMMM$~:=?Z??,.,,=778MMMMM\n" + DoubleTab +
	"MMMMMMMMD+=,::=,,~=+?====DMMMM\n" + DoubleTab +
	"MMMMMMMMO?.:,..~??~.~==~7MMMMM\n" + DoubleTab +
	"MMMMMMM8?:7$.:::~,..DMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMI+=D$,~:....:$DMMMMMMMM\n" + DoubleTab +
	"MMMMMM$~:.=?:...,.:~?MMMMMMMMM\n" + DoubleTab +
	"MMMMMM?=:..,~,,,,,,..$MMMMMMMM\n" + DoubleTab +
	"MMMMMM77:,,,:~,.,.::..8MMMMMMM\n" + DoubleTab +
	"MMMMMM+=,,,,:~:.,::~,,~DMMMMMM\n" + DoubleTab +
	"MMMMMDII,,,::~:,,::~,:?=MMMMMM\n" + DoubleTab +
	"MMMMMD?=,::~~:..,:,~~.$+?MMMMM\n" + DoubleTab +
	"MMMMM8=~,:::~,,=7Z..,:OM:ZMMMM\n" + DoubleTab +
	"MMMMMO~~:,:,:.7MMM:,~:DM8=DMMM\n" + DoubleTab +
	"MMMMMZ~~,,:::.ZMMMI.=:ZMMZ+MMM\n" + DoubleTab +
	"MMMMM7:~,:::,.OMMM8,~:?MMMIIMM\n" + DoubleTab +
	"MMMMMI~:.:::.:MMMMM+:~=MMMMZMM\n" + DoubleTab +
	"MMMMM+:,,,:~.7MMMMMZ:~=MMMMMMM\n" + DoubleTab +
	"MMMMM=,,,~=~:$8OZZZ7~==O88DMMM\n" + DoubleTab +
	"MMMMM=,,,I?~~?I????I+++7O88MMM\n" + DoubleTab +
	"MMMMM+..:~:,=I7Z8OOOI?~$MMDDMM\n" + DoubleTab +
	"MMMMMO?I7ZZ8DMMMMMMM$?+OMMMMMM\n"

const barbarianAscii string = DoubleTab + "MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMD7$O77$8MMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMZ~?I?+I7OMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMM8$?~??$77OMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMZ+~:+?777?OMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMZI$?=III7+=7$ZDODMMMMMMMMMMM\n" + DoubleTab +
	"MMMD$II7O8O7::?7++.:$8MMMMMMMMM\n" + DoubleTab +
	"MMMMMMI???+:,~+IID$.+7?I7ZMMMMM\n" + DoubleTab +
	"MMMMMMI,::..:=+==??:7MD7?$DMMMM\n" + DoubleTab +
	"MMMMMM?.::..,:~.,,,.$MMMMMMMMMM\n" + DoubleTab +
	"MMMMMM:,...,:~=.~,$$DMMMMMMMMMM\n" + DoubleTab +
	"MMMMM8.,..,,,,~...IMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMM7.:,..,~:,....DMMMMMMMMMMM\n" + DoubleTab +
	"MMMMM:..:,.,+?~:...?MMMMMMMMMMM\n" + DoubleTab +
	"MMMMZ...,,,,:~~,....8MMMMMMMMMM\n" + DoubleTab +
	"MMMMI.........,:....=MMMMMMMMMM\n" + DoubleTab +
	"MMMM+..........::....$MMMMMMMMM\n" + DoubleTab +
	"MMMM:...........:....,DMMMMMMMM\n" + DoubleTab +
	"MMMO........,...,,....=MMMMMMMM\n" + DoubleTab +
	"MMM?......,,:,,,,......IMMMMMMM\n" + DoubleTab +
	"MMD,.......,:.,~+~......+MMMMMM\n" + DoubleTab +
	"MM7........,...,.........=MMMMM\n" + DoubleTab +
	"MD:.......,,..,,....,:~::.?MMMM\n" + DoubleTab +
	"MO........,...,...~7DMMMMM8DMMM\n" + DoubleTab +
	"MM7:,..:Z,.....:ZDMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMM8IDMZ$$=..,?8MMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMD$7+~~$MMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n"

const barbarianAsciiB string = DoubleTab + "MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMMMMMMZ7IOMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMMMM8?,,$MM8$MMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMMMM$,,:+?O?:DMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMDMD8OZOM7=I7~:::,~DMMM\n" + DoubleTab +
	"MMMMMMMMMMMDI++=~,.IM=ZM?=$~,~$MMMM\n" + DoubleTab +
	"MMMMMMMMMM8=,:::...,=ZZI+D$=OOMMMMM\n" + DoubleTab +
	"MMMMMMMMM$~,::::~~=~,=~:=+?OMMMMMMM\n" + DoubleTab +
	"MMMMMMMO+~~:~=?+??=~:=~~~:~ZMMMMMMM\n" + DoubleTab +
	"MMMMMM8:~~::?Z$I:::~~,.:~=~:ZMMMMMM\n" + DoubleTab +
	"MMMMMM+~==I?=?I:,:~~:,::~===+7MMMMM\n" + DoubleTab +
	"MMMMD$?+7$I7I~:,.~:,,,,:::,~~,ZMMMM\n" + DoubleTab +
	"MMMM8~?ZZ+:+~:.,~~..........,=OMMMM\n" + DoubleTab +
	"MMMMD??7I~::,:.~:.,....~8?:+OMMMMMM\n" + DoubleTab +
	"MMMD?~~~::,,,.~~.,,...=DMM7DMMMMMMM\n" + DoubleTab +
	"MMMM?=~~=:,,,,,,,,....=MMMDMMMMMMMM\n" + DoubleTab +
	"MMMM=??7=:::,:,.,.,....$MMMMMMMMMMM\n" + DoubleTab +
	"MMMM7:==~=I:.,,,::,...,+8MMMMMMMMMM\n" + DoubleTab +
	"MMMMO::,,Z=,~,.,~:...,.~MMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMD$$$,?=,.,:,...,:.OMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMDI:+~.....:+.,,.8MMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMD~..,.=?=8M...IMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMO~..:8MMM+..?DMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMO=,,.~+~~~,,.:=78MMMMMMM\n" + DoubleTab +
	"MMMMMMMZ7Z$~,,,,,,,,:,.,..:7DMMMMMM\n" + DoubleTab +
	"MMMMMD=:..,,.,::~==++???7O8MMMMMMMM\n" + DoubleTab +
	"MMMMM88M$Z+$OO888DDDDMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n"

const barbarianAsciiC string = DoubleTab + "MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMD$$DMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMM+=+IMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMI77IMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMI$?,78DMMODMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMZ~.~,.,~=+IDMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMM?..,:,:==~~~7MMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMM=,::+==::+~=7MMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMM?::+~:~::.=+IDMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMM$,::,=.,.I?~+OMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMM?....~..IM7.:+MMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMM:...~:,.?M7..+MMMMM\n" + DoubleTab +
	"MMMMMMMMMMMM?....,:..:D8..:?MMMM\n" + DoubleTab +
	"MMMMMMMMMMM$:,.:,...,:OD.,:ZMMMM\n" + DoubleTab +
	"MMMMMMMMMMM==7~=:...,,7I:IOMMMMM\n" + DoubleTab +
	"MMMMMMMMM8?$MO........:,+OMMMMMM\n" + DoubleTab +
	"MMMMMMMMI+8MM7.,:..,,,.IMMMMMMMM\n" + DoubleTab +
	"MMMMMMO=+DMMMI.....,.:,=7MMMMMMM\n" + DoubleTab +
	"MMMMM7~IMMMMM$.........,IMMMMMMM\n" + DoubleTab +
	"MMM8?7DMMMMMMM=...?Z...:?DMMMMMM\n" + DoubleTab +
	"MM8$8MMMMMMMMM?...OMZ...=DMMMMMM\n" + DoubleTab +
	"MMDMMMMMMMMMMMZ..:MMM:..?MMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMM,.=MMM+..$MMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMM8..=MMM$.,ZMDMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMDI...=MMM7.,DMDMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMM7:=7DMMMD..,8MMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMMD7IOMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n"

const barbarianAsciiD string = DoubleTab + "MMMMMMMMMMMMMMMDMMMMMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMMO$ZO8DMMMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMZ++?I7ZODMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMO??=??+78MMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMD??=??=+$DMMMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMM$+=+?~=I7$OMMMMMM\n" + DoubleTab +
	"MMMMMMMMMMMMMO~++?==?==?$DMMMM\n" + DoubleTab +
	"MMMMMMMMMMMD7?:~++:+7?=:~~IDMM\n" + DoubleTab +
	"MMMMMMMMD8D7==?=~~:=?+:~~,==7D\n" + DoubleTab +
	"MMMMMMM8=I+I=+~~~++==~:7M7,:+I\n" + DoubleTab +
	"MMMMMMM$~=:,7MI,~=++??+DO~:=?O\n" + DoubleTab +
	"MMMMMMD+::,~MZ=,~====::MZ7?:$M\n" + DoubleTab +
	"MMMMMD+:,.=OI:::,,,.~==7MMMODM\n" + DoubleTab +
	"MMMM8=,,,7D+,,...,,.:=:,OMMMMM\n" + DoubleTab +
	"MMMMD~,,?DI,.,..::,:::,,?MMMMM\n" + DoubleTab +
	"MMMMMI~+Z=,,....::~~:~:?Z8MMMM\n" + DoubleTab +
	"MMMDI:~DI:,,...,:,,:~=~:$8MMMM\n" + DoubleTab +
	"MMMM7:7M?=:,.,,.,,,,::==?DMMMM\n" + DoubleTab +
	"MMMM?+MMI,,.....~+=,.,,:7ZDMMM\n" + DoubleTab +
	"MMM8~ZMMO,.....:=+=$Z8O788DMMM\n" + DoubleTab +
	"MMMI+MMMMDO?:.,==:7MMMMMMMMMMM\n" + DoubleTab +
	"MM8=$MMMMMMM$.,:~,+MMMMMMMMMMM\n" + DoubleTab +
	"MM7~MMMMMMMM+:~,.,IMMMMMMMMMMM\n" + DoubleTab +
	"MM~7MMMMMMMMO~:,,:MMMMMMMMMMMM\n" + DoubleTab +
	"M7~MMMMMMMMMI,,..+MMMMMMMMMMMM\n" + DoubleTab +
	"D:7MMMMMMMMD~:,.?MMMMMMMMMMMMM\n" + DoubleTab +
	"+:DMMMMMMMMM+:,,?MMMMMMMMMMMMM\n" + DoubleTab +
	"I8MMMMMMMMMM8$$$$MMMMMMMMMMMMM\n"

const dragonSkillFireAsciiA string = "II77I77II+=:::::~~~~~~~::,,,,,,:+7$$$$$Z$$$$$$$$$$ZZ7?~:,,:::::::~~====++=+?I77I\n" +
	"?I7$7I?+=~~::~~~~~~~~::::::::,,,,:+7$ZZZ$$$$$$$$$$7+~:::::::::::~~==+I7777I77$$$\n" +
	"?I$$7I?+=~~~=======~~:::~~~~::::::::?$$$$$$$$$$$7+:,,:::::::~::::~=+7$77I77$$777\n" +
	"?I7$77I77??II7I?++==~~~~~=~~:~~~~~~::~IZZZZ$$$$I~,,::::::::::=++=~=?777IIII77777\n" +
	"?I7$77$$Z$7$$$777?==~~=+===~:~~~~~~:::,~?7I++?=~,,::::::::::=777I?+?IIII??IIIII7\n" +
	"I7$$$$ZZ$$$7777I77?~=?7II+=~:~~==+==~:,,,~~::~:,,,,::~~~:::~IIIIIIIIIIIII7777777\n" +
	"7$$$$$$777$$7IIII7I?I77I7I+~~~=?I77II+~,,:::=~::=+????+=~::+II?IIIIIIII77$$$$$$7\n" +
	"$77777777777$7II777I77IIII?~~+I7III?I7I+:::~=~=II7777777?+:+III??II77$$$$$$$$$$$\n" +
	"7II7IIIIIIII77IIII7777II?I?=+I7IIIIII777+:~===III????IIIII?+II777777$$$$7$$$77$$\n" +
	"III7IIII?????II??III77IIII?+???I7IIIIIII+:=$?~??????????????III77777777777777777\n" +
	"III7I???????????????IIIIII????????????I?~~IO$?????????+???III77II77777IIIIIIIIII\n" +
	"??IIIIII??+????????????II???I??????????+~+7ZZ$III????III?I77IIII??I????IIIIIIIII\n" +
	"??????IIII?+++++??????????????II????????++$O$ZI=?777777IIII????I?????IIIII77IIII\n" +
	"??????????????????????????????II????IIII?=7Z$Z$?I77III?????????7I???IIIII777IIII\n" +
	"???????????????I?IIIII?????IIIII????IIII+=I7$ZO$IIIIIII??IIIIII?+?7IIIIIIIIIIII$\n" +
	"???????????????????IIIIIIIIIIIIIIIIIIIIIII7I$$Z$77III?+==+I7777II7IIIIIIIIII7$$$\n" +
	"???????I????IIIIIIIIIIIIIIIIIIIIIIIIIIIIII77$I$$IIIIII???I$777$$7$7?IIIIIII7$Z$$\n" +
	"II?IIIIIIIII?IIIIIIIIIIIIIIIIIIIIIIIIIIII7$ZZI$$III777777$$$$$$$$$7IIIIIIII7$$$$\n" +
	"??IIIIIIIIIIIIIIIIIIIIIIIIIIIIIII????IIIII7$$77$77$$Z$ZZ$$$7$ZZ$$$III7777II$$$$$\n" +
	"==+???IIIII777II77IIIIIIIIIIII777I??I7III777I777$7$$$7$$7$Z$$$$77II7II7$$$$777Z$\n" +
	"===+++?I7777$$$7$$7I++?77IIII77$$$7$$7IIII7$777?7$Z$II77II7$7$7777777III7$7III7Z\n" +
	"===+??I$$7777$$$$77I++?7777777$7$$$$$77777$$7$II7$ZZ$7$$$77$$$$$$$$7II??7Z7+??7$\n" +
	"=~=++I777$$777$$$777777$$$$$$7$$$$$$7I777$ZZZ$77$$ZZZZ$$$$7$77$$$$$III???7$I?I7Z\n" +
	"~~=+I7??+I7I7$77I??77III7$$77$$$$$7II?I7$ZZOOZZZ$$$$77$$77I77$$$ZZ$7II?++?7III?I\n" +
	",:~++++?+??II7+~~::~=+==?I777$$$Z$7III7$ZZZ$Z$$$$$$7777$7777$$$$7777?+?III?=~:,:\n" +
	",::~:~==~~~~~::,,:~~~~====+??II777II??777$$7777$$$77$$7777I?++++===~:::==~::,,::\n" +
	"::,:,,~=::::~=++~:::~~~:~~::~==~~~=~~+?IIIIIII?+++======~~:~:::,:::,,,:,,,,,,,::\n" +
	"::::~::=~~::=?77=,:::,:::::::~~:::::~=~~==~====~~::::::::::~~~~~:::::~:::,::::,,\n" +
	":,,::::::::,:~~~:::::::::::::::::~~::::~~~~~~~~~~:::::~:::::::::::::::~::,,,,,,,\n" +
	",::,.,~~~~:,,,,,:::::::::::::::::~:,,::::::::::~::::::::::::::::::~=~:::,:,.....\n"

const dragonSkillFireAsciiB string = "....................................,,,,........................................\n" +
	"~:::::~=:.,::....................,:~~:,.........................................\n" +
	"????++~=?I=:,~~:,.............,~===~:,.....,....,,......,.......................\n" +
	"?++??++==+77?~::~~,.......,::~====:,.....,,......,......,,......................\n" +
	"+===++?+==+II$$?+=~~~::::~==+===~:,....,,,.......,,.....,,,....,,,..............\n" +
	"+======+=~~~==?$$O777OO$7?=~==~:,,....,,,,.......,,......,,,,..,,,,,,...........\n" +
	"++=+=+++=++~~:,:?7$$OO8D8$$=,.,,..,,,,,,,........,.........,,,,,,,,,,,,.........\n" +
	"===+=+++=+~+=~:...:?77$OD888Z+..,.,..,,,.........,,........,,,,,,,,,,,,,,.......\n" +
	"+=~==+==+=+=:~~~,.,::?+??IZ7$O?.....,,:,.........,,.........,,,,,,:::::,,,,,.,,,\n" +
	"=+?=+=:::==~~===~.,=...~==?$=+O,....,::..........,,,...,,,,,,,,,,,:::::,,,,,,,,:\n" +
	"II+=~~~:,~=:,.:~~~~,,...,~+I7$$?....,:,..........,,,..,,,,,,,,,,:::::::,,,,,,:~~\n" +
	"?~:,.,,::,~?==~,::::,,,::,~=?$O8I...,,,..,,,,.,..,,,,,,,,,,,,,,,,::::::::,,::~==\n" +
	"+:,.,,.,,..~?=~=+??+~,=+:,.,~+7OMD$~.,...,,,,,,,.,,::,,,,:::,,,,,:::::~::::~=+??\n" +
	"=:,.,,,...,,:::,:?77$OMMDZ+,...,7OO8?.,.,,:,,,,,,,,:::,,,:::::::,,,::~~~~~==+?I7\n" +
	"+::,,:,..~~~:.,.,.:IZDMMMMDO+,...:+I7,,,,,,,,,,,:,,:::,,,::::::::,,,:~~==+??II77\n" +
	"++~~:,.:~~:~~:...,.:7MMMMMDDZI+++=?I=,.,,,,,,,,,,,,,::::::::::::::::~=+??IIII777\n" +
	"====:.:I~:~::~:.....:ZMMMMMDDD8$I?+=~:,,,,,,:,,,,,,:,,:::::::::~=+++???IIIIII777\n" +
	"?==~,,++~~~~::::,....~$8DMMMDDMMD$?+?7+::,,:,,,,,,,,::,::::::~=+????IIIIIIII7777\n" +
	"II?=::I====~~:~~~~,...=78DDDMMMDMMDZ7I?+=~,,,,,,::::,~::::::~=????III777IIII777$\n" +
	"II?=~=$====~~==+=~~:..:?Z8DD8DMMMMMM87II=+?:,::::::::::::~~~=??IIII7777IIIIII77$\n" +
	"?+?+=+7=+++++++?++==~:.:7O8D88DDMMMMMDZ$II7+=~:~~::::~~=::~=?IIII777777I7IIII77$\n" +
	"77??+?7+II?II77III??+?=,,I88OOO8DDDMMMM8D8OZ7I=:::~~~~~~::~+??III77777I77IIII77$\n" +
	"7I+=~+$77777$$$7777IIII?::?D8$Z8DDDDMMMMDDD8OZ7=~~~~~====:~=?IIII77777I77IIII7$$\n" +
	"II?==+$77$7$Z$$$77$$IIIII+,+$$$O8DDDMMMMMDMMDD87??++???I?===?7III77III7777777$O8\n" +
	"I?+?=+$777$$7$77777$777III?+IZZZO8DDDMMMMMMMDMM8$7?I77III??+?IIIIIII777$$$ZZZ8DD\n" +
	"I?++==?II777777$7$$$$$$77II777ZZZO8DDDMMMMMMDMMM8OZ77I7777I?IIIIIII77$$ZZ8DDDDDD\n" +
	"II?+==7$77I77$$$Z$$ZZ$$ZZZ$$$77Z$$ODDDMMMMMMMMMMMMMD8O$$7IIIIIII7$$ZZOOO8DDDDDDD\n" +
	"??++==Z$I77II$$$$$$ZZ$$$Z$$$$$$$$$ZO8DMMMMMMMMMMMMMMMMD8Z7II7777$ZO8888DDDDDDDDD\n" +
	"+?==~~$I~=+++?7777$ZZZ$$$$ZZZZZZZZO888DMMMMMMMMMMMMMMMMMD8O88OOOOO88888DDDDDDDMM\n" +
	"???+=~I7:~==++??II7$ZZZZZZZOOOOOZODDD8DMMDMMMMMDDMMMMMMDDDDDDDD8D88DDDDDDDDDDMMM\n" +
	"??+~~=IZ~~~=++???I7$$$ZZZZZZZOOZO88DDDDMMMDDDMDDDDDDMMDDDDDDDDDDDDDDDDDDMMMDMMMM\n"

const necromancerAsciiA string = "I???????IIII??I????????????I???I???+=++=+==IIIIIIIIIIIIIIIIIIII??II?????????????\n" +
	"????????IIII?????III?????????II?I?+:::,:,::~I??IIIIIII??III??II?????????????????\n" +
	"IIII???II?III??IIII?IIIIIII+++===:::,....,:.=I?III????II?III??III???????????????\n" +
	"IIIIII??IIIIIIIIIII?+??=+III+:,,,,..,:~?~,,,~III?IIIIIII??IIIIIIIIII???????II?II\n" +
	"IIIIIIIIIIIIIIIIIIII?++++???~,::,:,.:~=+=:,~+??=?I???IIIIII?~?I?II??IIIIIIIIIIII\n" +
	"IIIIIIIII??IIIIIIIIIIIIII??++~=:,:,,:~+?+:=?+???II??IIIII??I+=+?I?+?IIIIIIIIIIII\n" +
	"IIIIIIIII??II?IIIIIIIIII?I?I77?+=::,,....~?=~??III??+??II??I?I?III++?+?IIIIIIIII\n" +
	"IIIIIIIIIIIII++IIIIII?+++I~:++~?+=?~....,==~=+=?III???IIII?+?III???++?IIIIIIIIII\n" +
	"IIIIIIIIIIIIII+??II?IIIIIII?++?:~=:,,..,::~~~:==I?IIIIIIIIIIIIII+?IIIIIIIIIIIIII\n" +
	"IIIII7II7IIII777III~~~=II+II??=:...,,.,,,,,~~~~=++?II?7IIIIIIIIIIIIIIIIIIIIIIIII\n" +
	"777I??+I777777I+I7I~::~=+II+?=+=~,,,,,,,,.,,,.=++II7?II7?$I7777II7IIIIIIIIIIIIII\n" +
	"7777777777777777777I::::=??++::::,,,,,,,,,,,,,~=?77I77=?+?7777I?I777777777777777\n" +
	"77777777II77I7$?I77?::::::+=~~~:,::::~~::,,.,::=77???=::~?7777777777777777777777\n" +
	"777777777?77I+=+777?,?=::~+=~=::::::::::~=+=:,~~:~~,,,,,:I77I77$I777777777777777\n" +
	"777777$77$I7$I??I?=+=7I+==,~+~~~::,,,,,,,:?7$I=:::,,:,,=I7I7I7I7I77777777$$$7$77\n" +
	"777777$7$?:7$II7777$77$I?::++=+=~:::,,,,:::=7Z$7I?:,+++I7II7I77777777$$7$$$$$$$$\n" +
	"777777++?=I$I+=7$7I?7I?I::?II?+=~~~::,:::=~=$8I?~7+::7777777?777$77$$$$$$$$$$$$$\n" +
	"77777777I7$7I7I77I+I$+?I++77I?+==+=~:~~~===??O8Z+==+?7?7$$$$?+77$77I7$$$$$$$$$$$\n" +
	"77777$77$$$7$$$$$7$7$7?$$$$7$7I?I?++==++?++~=$888$?7I?I77I77+:I777$7$$$$$$$$$$$$\n" +
	"7777$7$$7$777II777$$7$II77I7$7I?II?I??I+?I?+II88DO+:=77II7II777777$$$$$$$$$$$$$$\n" +
	"77777$$777$7+I77$$$$$$$$$$7?III7III+:+IIIII7I?Z8DZ7+,=I?II7$$$$7$7$$$$$$$$$$$777\n" +
	"77777$$$$$$$$$$$77$$$$II$$$$77$$77+,..:==+??7I+7$77$~~+?++77$$$$$$$$7IIII777?=::\n" +
	"+++??IIIII77$$$$I+I$7II$$ZZ$$$?++=,,,.:++I7$$I=?$7$$$$$$7$I=$7$$$$7?~:::~~=::,,,\n" +
	"::::::~~====?I7$$$7I7$Z$$77Z$$7?~.,:,,,=I$$$7$$?$Z$$$77I7=:7IIII7I+:,,,,,,,,,,,,\n" +
	":::::::::~===+++???I77777$I$$$ZZZ=,=+=:+$$ZZOOOOZZZZZ7?++=??+=~~~~,,,,,,,,,,,,,,\n" +
	",,,:,,,:::~~~~=+++++?I???I??II7$7=I$ZZI~$OOOOZZOOZ7I?++====~:,,,,,,,,,,,,,,,,,,,\n" +
	",,,,,,,,,,,,,::::::::::::~~~~=+?+?I?II$?+7I7?+=??=:,,,,,,,,,,,,,,,,,,,,,,,,,,,,,\n" +
	",,,,,,,,,,,,,,,,,,,,,,,,,,,,,,::~~~~~=~=~~~:::,,,,,,,,,,,,,,,,,,,,,,,,,,,.,,,,,,\n"
