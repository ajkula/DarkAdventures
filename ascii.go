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
		enemiesList.DRAGON: {dragonSkillFireAsciiA, dragonSkillFireAsciiB},
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
	heroesList.Thief: map[int]string{
		0: translate(heroStoryThiefTR0),
		1: translate(heroStoryThiefTR1),
		2: translate(heroStoryThiefTR2),
		3: translate(heroStoryThiefTR3),
	},
	heroesList.Paladin: map[int]string{
		0: translate(heroStoryPaladinTR0),
		1: translate(heroStoryPaladinTR1),
		2: translate(heroStoryPaladinTR2),
		3: translate(heroStoryPaladinTR3),
	},
	heroesList.Wizard: map[int]string{
		0: translate(heroStoryWizardTR0),
		1: translate(heroStoryWizardTR1),
		2: translate(heroStoryWizardTR2),
		3: translate(heroStoryWizardTR3),
	},
	heroesList.Barbarian: map[int]string{
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

const skeletonAscii string = DoubleTab + "...........~............................\n" + DoubleTab +
	"...........=............................\n" + DoubleTab +
	".....,?I,..,..,?:.......................\n" + DoubleTab +
	"....:NMM8+$8$+,,.,=,....................\n" + DoubleTab +
	"....:DMMNMMMMMZ,:,:.....................\n" + DoubleTab +
	"....~ZOZ8MMMMMM7........................\n" + DoubleTab +
	"...,+$~.7MDMMMN8+,+Z+...................\n" + DoubleTab +
	"....~~...ZDOMO7:7DI.....................\n" + DoubleTab +
	"..,:~,...:O,ZZ8$,ZMD=...................\n" + DoubleTab +
	"...,.....I+.ZMMMO.:$NO+=................\n" + DoubleTab +
	"........?7,ONMMMMZ..~$8Z=...............\n" + DoubleTab +
	".......?7.$MDDDMMN:.~+.,I$+,............\n" + DoubleTab +
	".......+:.8M?ZDDMN+.=....,+I?:..........\n" + DoubleTab +
	".........~$D,~O$OM8.~.......~I?:........\n" + DoubleTab +
	".........?NM=,?..$D~,.........:??~......\n" + DoubleTab +
	".........,NN:.=~.+MM?...........,=?~....\n" + DoubleTab +
	"..........IM~..:..+M8..............~+~..\n" + DoubleTab +
	"...........DO:,....~M?,,.............,..\n" + DoubleTab +
	"........~I$OZ:::::~.ZMI::...............\n" + DoubleTab +
	"........,::.........?Z~.................\n"

const sorcererAscii string = DoubleTab + "$$$$$$$$$77$$$$$$$$Z$$$$$$$$$$\n" + DoubleTab +
	"7777777777777777II77777II7I777\n" + DoubleTab +
	"IIIIIIIIIIIII??I???III????IIII\n" + DoubleTab +
	"??????++????+IO8I+????++++????\n" + DoubleTab +
	"+==+++====+=+DOON+=+======++++\n" + DoubleTab +
	"==========~=ZNDNMO+~=~~=======\n" + DoubleTab +
	"~~~~~======ZDNNMNNN?~=~====~~~\n" + DoubleTab +
	":::~~~~===$ODNDD88ND~~=~~~~~~~\n" + DoubleTab +
	"::::~~~?+O8DNNDNN8DM8=~==:::::\n" + DoubleTab +
	",:::::=7ODDMMNNNNDDMM8?+=~::::\n" + DoubleTab +
	",,::::=7DDNM8DDODD8NNN8I:,,,,,\n" + DoubleTab +
	",,::::,+DNNM8DN88NONNMI::,,,,,\n" + DoubleTab +
	",,,,,:,:8MNMN8N8DNDNMZ.,,,,,,,\n" + DoubleTab +
	",,,,,,,:DMNNNDM8NN8NO:.,.,,,,,\n" + DoubleTab +
	",,,,,,.,778NNNNDNN8MI.,.,.....\n" + DoubleTab +
	"........:~8DDDDNDNMZ:,........\n" + DoubleTab +
	".........,8D8DDNDDN7.,........\n" + DoubleTab +
	"..........$N8DNDD8D?,.........\n" + DoubleTab +
	".........:OD8DDDDO87~=........\n" + DoubleTab +
	"......,,.+8O88888OO8$7?+=~,,.,\n" + DoubleTab +
	",,,,,,,,+88OO888OOOO8Z?=~:,,,,\n" + DoubleTab +
	",,,,,,,,+?????I7I????+,,,,::::\n" + DoubleTab +
	",,,..........,,,,,,,,,::::~~~~\n"

const orcAscii string = DoubleTab + "........................................\n" + DoubleTab +
	"....................=7?:................\n" + DoubleTab +
	"...................~D8OZ,...............\n" + DoubleTab +
	"...................+O7I$$?I???+:........\n" + DoubleTab +
	"...................,Z8$$DD88ZOOO=.......\n" + DoubleTab +
	"...................:8OZ8DD8D8$$ZO7+.....\n" + DoubleTab +
	"..................=8D8OO8D88D8OOO88I....\n" + DoubleTab +
	"..................ID8OOZ7$$$8DZZ88ZOZ~..\n" + DoubleTab +
	"..................+N8ZO8Z$$$ODOZOD87$Z..\n" + DoubleTab +
	"...................+OO$ZZ777$OD$ONO$ZO,.\n" + DoubleTab +
	"......,,,,,,.....,:,O8Z777I77$OIO8888$..\n" + DoubleTab +
	".............,,:=?NDDN8$???7$$Z=7Z$O7...\n" + DoubleTab +
	".........,,......,I$IIDO$$I$OO8$Z$ZI,...\n" + DoubleTab +
	"..........,~~~:,.....+N888Z8DD8D8O$.....\n" + DoubleTab +
	".............,:~~====$O8OZ7O888OOO,.....\n" + DoubleTab +
	"...................,,ZOO8O$O8D8ZO7+:....\n" + DoubleTab +
	".....................Z$77ZDD888OI~,,....\n" + DoubleTab +
	"....................=DO8OO8NDOZ8~.......\n"

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

const dragonAsciiB string = DoubleTab + ".......................................,,.........\n" + DoubleTab +
	"............................,=I??I7$$7+=:,........\n" + DoubleTab +
	".........:................=7O88888O?~,............\n" + DoubleTab +
	"........=~..~..........=7ODD8OOZZ$$=..............\n" + DoubleTab +
	"........~I::?,.......:$NND88888DDD88Z+~,..........\n" + DoubleTab +
	".,.......:$88O$I~..,I8D888888OOZI+=:,,,...........\n" + DoubleTab +
	".,=ZZZI?,~OD8888O$7DNDDDDDD88OO$:.................\n" + DoubleTab +
	".,~===+$O8ND8DD88D$8MND8888888DD8?................\n" + DoubleTab +
	"........7D8ND7$?~I+.+ZDNDOOOZZ$7I?=:..............\n" + DoubleTab +
	"........?8ODDZI~=?7$$ODND88Z?~,...................\n" + DoubleTab +
	"........?88O88888DNNNND8OZ?,......................\n" + DoubleTab +
	"........~D8D8DO8DD8DODDD8,...:7Z$,................\n" + DoubleTab +
	"........+8O88888OO8DNMDODO+:...:7O................\n" + DoubleTab +
	".......~N8O8DDDN888DNMND8DN8Z7?7$NI:..............\n" + DoubleTab +
	"......:ZND888DNN88O8DNNND87ODMMMMD7=..............\n" + DoubleTab +
	".....=8DD?I8D8DD88OD8OO8D$..,,~+?~................\n" + DoubleTab +
	".....ODDD8OO88888DD?78888DO:......................\n" + DoubleTab +
	"...,7$$Z$$?+:~=+??=..,,,OM8=......................\n" + DoubleTab +
	"...,,......,............,?I:......................\n"

const orcAsciiB string = DoubleTab + "...................,,::.................\n" + DoubleTab +
	"................~?=Z7$7.................\n" + DoubleTab +
	"..............+$$+.ID8O$ZII?~:,,........\n" + DoubleTab +
	"............,=OD,..$DDNDNNDZZ8I~........\n" + DoubleTab +
	"............~:+D+=$MNZ8OZZI+~,..........\n" + DoubleTab +
	"..........,,~+:O$ZZ$NNO8N~..............\n" + DoubleTab +
	".....,,=I7ZO77$77=..ZDDD8$III?Z=,.......\n" + DoubleTab +
	"....,$7$?$DNOZ8OZZ?IOZZDD8?~~,Z8........\n" + DoubleTab +
	"..+,,DO7I??Z$ZDDZOOO8I$NMZ+,.,DI........\n" + DoubleTab +
	"..+7:ZDDOOZ?ZDZ$88O88IODDO8O$7+,,=......\n" + DoubleTab +
	"...I$ONO7Z$ODZ$Z7$8DD8NN88ZZO$=?+Z......\n" + DoubleTab +
	"....,IZ88O8Z7OZI77Z8DONNDO888DOI=I,.....\n" + DoubleTab +
	"....~?ZN8OD88$II?ZN88DNND8D88N8,,.......\n" + DoubleTab +
	".....,~7D8$88I?$$Z88NMMNDMMD8O8?........\n" + DoubleTab +
	"....,7ODDZ~.?Z888=.,~$+:.+ODD88Z77:.....\n" + DoubleTab +
	".....:??,....,~~...........,~~~ZNO~.....\n" + DoubleTab +
	"...............................OMZ......\n" + DoubleTab +
	"..............................,+?~......\n"

const sorcererAsciiB string = Tab + "$$$$7I??II=~~:::~~~+===++========+???IIIIII777$$$$\n" + Tab +
	"$$$77I?+??=::::::::I?:~~=+=~~~===++=++??I?II777$$$\n" + Tab +
	"$$77II?==++:,,,,,,,+I::::~==~~~~~=+===+????III7777\n" + Tab +
	"$$77II?=~~==:,,,,,:~$=:~:,~=::~~~====~===+??II7777\n" + Tab +
	"$$77II+=~~~~=~~::~:,I$::II?~,:~::~+~==+++???II7777\n" + Tab +
	"$7777I?==++=~~=:,,,:$ZI?777+::,,,,~=:+==++??II7777\n" + Tab +
	"Z$$7II?+II+++=~~~:,~7ZZ$$$$?+:,:::,+?=:==+++??II77\n" + Tab +
	"$77IIII7$?=====~=~::,+ZZ$$$?~~=~~~==?~+++??I7I?III\n" + Tab +
	"$77I7$IIZ7????+~===++=?7$$$Z+~==~:==+?+~~==+I$IIII\n" + Tab +
	"$77$ZI+I$$I==+I????????I$$$Z$?+=+??+=?II7?=+?7$III\n" + Tab +
	"$$ZZ$IIZZ$7??77I???+==I$$$$$$$7+~+I+=+IOO7+?II77$$\n" + Tab +
	"$Z88$7ZOZ777$8OII7I??I7$$$$$$I?IIII?++$OOO$7$$$ZOZ\n" + Tab +
	"$ODDO$OOZZZ7ZOOZ77Z$7$$$777$Z7I$77I7Z$ZOO8Z77$ZOOO\n" + Tab +
	"ZO8D8OOZZ8O$ZOOZZ$OO$Z$$77$ZO$$Z$Z$$ZZ$OOOZ77$OOZZ\n" + Tab +
	"ZZ888ZZZZ8OZZZZZZZOOZ$ZZ$7$ZZZ$$$ZZ$$$ZOOO$7$ZZZZZ\n" + Tab +
	"8888D888888888888888OO88OOOOZOOOO8OOOO8888OOO8O888\n" + Tab +
	"DDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD\n"

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

const goblinAsciiB string = DoubleTab + "..................::,.........\n" + DoubleTab +
	"................+$$?II+~......\n" + DoubleTab +
	"............:=?ZOZZ?I$OZ?,....\n" + DoubleTab +
	".........,...:?ZDOZZOZ?+=.....\n" + DoubleTab +
	"........?+=I?$Z$$$77$Z?,.....,\n" + DoubleTab +
	".......,Z~7DZ$II+II+~+I~..:+=:\n" + DoubleTab +
	"........$+OD88O8OZ:...,..=I?:.\n" + DoubleTab +
	".......,O+O8OO8$+~.....:?I:...\n" + DoubleTab +
	".......IN$IZDZ7I+=..:~?==.....\n" + DoubleTab +
	".......$N8?+7Z7$$II$O7+.......\n" + DoubleTab +
	".......7DO$IIZOZI:ODO?,.......\n" + DoubleTab +
	"....,?$DDO7=,:ZOOZ?~:.........\n" + DoubleTab +
	"..,IOO8Z$Z=...=O8OI...........\n" + DoubleTab +
	".:8DOII=:.....88OI:...........\n" + DoubleTab +
	".:+I7$$Z?,...?I+:.............\n" + DoubleTab +
	".....,=?I77IOO7+~~=,..........\n" + DoubleTab +
	".......,::~=++=~~~~...........\n"

const nightWalkerAsciiA string = Tab + "..............................,.....,~=~:,.............\n" + Tab +
	"....................?7..~?$O7,.,~=??=~,................\n" + Tab +
	".............,:,...+DD:?8N8OMNOZI~,....................\n" + Tab +
	"..............:~=?=7D8$ZN7+.:7I,.......................\n" + Tab +
	".................,+$OZZ8D~+:.~.........................\n" + Tab +
	"...................=DOZDD8?::I~:,.,....................\n" + Tab +
	"....................$DODO8D$$7=...,....................\n" + Tab +
	"....................~NNDO88DZ7?=::.....................\n" + Tab +
	"....................,Z$ZDD8DOI?=?~:....................\n" + Tab +
	".......................~DDDDD8$I?==....................\n" + Tab +
	".......................ID8DD88Z$$I~....................\n" + Tab +
	"......................+DO8D8O8N$?I=....................\n" + Tab +
	"......................ZOO8D88DNO=~:....................\n" + Tab +
	"......................O8O8DODDNN?I,.:,.................\n" + Tab +
	"......................INDN8O8D8+~?,.,,.................\n" + Tab +
	"......................+NDDZ88M$,=$+,,..................\n" + Tab +
	"......................ODN8O8DMI~,,+,,..................\n" + Tab +
	".....................INDN8DDMO+:.......................\n" + Tab +
	".....................OND888D87=,.......................\n" + Tab +
	"....................=MDZODONII$=,......................\n" + Tab +
	".........,::~......,DDZOD8NM8:?O+:,....................\n" + Tab +
	".........,::?Z+=,..?DZ8NNNNN8:.??=~,...................\n" + Tab +
	"...........::~?D7?$O8NNNNNDNI,:~?7=,...................\n" + Tab +
	"..........,:,:+8D88NNNNNNDDN$=:.,==,...................\n" + Tab +
	"............:==~?,=MNNNND88DO?:...:,...................\n" + Tab +
	"..................?NNND$N8OD8=:,.......................\n" + Tab +
	"..................?DDN8:ODZ8D~,,,......................\n" + Tab +
	"..................IDDN$:=DO8D=.,.......................\n" + Tab +
	"..................7DDN?~.?NOD$.........................\n" + Tab +
	"..................IDDN=.,.7N8D7:.......................\n" + Tab +
	"..................=DDD=,...IDDDD$,.....................\n" + Tab +
	"...................8DN7~~...=OMNNO,....................\n" + Tab +
	".....,.............+NND+==:...7MMMZ,...................\n" + Tab +
	".....~7~............=NN$++~,=:.~OMM8?,.................\n" + Tab +
	"......,$,............?M8DI+~??,..+NMNZ~,...............\n" + Tab +
	".,$?~,+D+...........,~NDO77II:....=MN8+~...............\n" + Tab +
	".=8D8DN$..............DM7+I=:......IMD+=,..............\n" + Tab +
	"..:OD8O?I=...........,DM7~~:,..,....8M$??~,....:.......\n" + Tab +
	"...,ZOZZO:...........:DD$+=:...:..,.=NDI?++:,..,.......\n" + Tab +
	".....Z888O~..........=M8$7?=:~,,=~:..$MZ=+==++~~,,,....\n" + Tab +
	".....Z8DDOOI,........ON87$7I??=,:=~+=$ND7+??I?~~:,,....\n" + Tab +
	".....?D7DD$D=.....,~7NDD$77I?I?=+=+?+ONNOI7$7I++=,.....\n" + Tab +
	".....?D8Z88+.,:=~=ZNMNDO88OOZ$7ZOZI??$NND$$Z$$I?+~,,...\n" + Tab +
	",:~+IO88788I$ZO88888888DDNNDDDDD8DDOO8NNN8OOZZZZI+=~~..\n" + Tab +
	".,~IDZIZZOZZOO888O8OO88888O88ZOZZOO8DOZ$$$$$Z8DO$$$?~,.\n" + Tab +
	".,:~~:~~:,,.,.,,.......,,,,::,:::::::,....,..,,,,:::...\n"

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

const wizardAscii string = DoubleTab + ".......:,.....................\n" + DoubleTab +
	".......I+.....................\n" + DoubleTab +
	"....:~:7?.,.,.................\n" + DoubleTab +
	"....+O?ZO7=ZZ$7~.....,,.......\n" + DoubleTab +
	".....:=ID==8$$8D.....,~:......\n" + DoubleTab +
	"........8$ZO$IZD=....~=~......\n" + DoubleTab +
	"........?DZOZ?$DO....::,:.....\n" + DoubleTab +
	".......:7+$$IZ78O+...::I~.....\n" + DoubleTab +
	"......:OIIOZ7$O8ON7.:?+:......\n" + DoubleTab +
	".....,Z$IN8N77O88NO7OI........\n" + DoubleTab +
	".....,ZZ8M8DDDNNNNODMDI.......\n" + DoubleTab +
	"......$DDN8DNNNNNMDDNNMO......\n" + DoubleTab +
	".....~O8ODD8NDNNNMMND88NZ,....\n" + DoubleTab +
	".....78ONZD88DNNMMMMNO7===....\n" + DoubleTab +
	".....O8ON$ZNODNNMMMMMND8I:....\n" + DoubleTab +
	"....,88ODZZNO8DNMMMMMNDNMN7...\n" + DoubleTab +
	"....~D88DZON8O8DNMNMMNN8I?7?..\n" + DoubleTab +
	"....~D88NOODN8ZONMNDMMND:.....\n" + DoubleTab +
	"....~8OONNDDM8Z$8NN8NMNND:....\n" + DoubleTab +
	"....:8OO8NDNMOO$Z8M8DMMNMD....\n" + DoubleTab +
	"....:NNNNNNNMOZZOONN8NMM8?....\n" + DoubleTab +
	"....:NNNMMMNM8Z8ZO8MND+:......\n" + DoubleTab +
	".....DNDNNDNMDZ8O8DO+Z........\n" + DoubleTab +
	".....+O8ND~NOIDD87I~.,........\n" + DoubleTab +
	"......,8M?.:..O8DO7=..........\n" + DoubleTab +
	"......?87.......?~~:..........\n"

const wizardAsciiB string = DoubleTab + "..............................\n" + DoubleTab +
	".............:+:.:............\n" + DoubleTab +
	".............=D8O$........,...\n" + DoubleTab +
	"..........,~+OO$8I.......:~...\n" + DoubleTab +
	"......,..,7DNNZZD=......,O?...\n" + DoubleTab +
	".....:.~.,?888$OO+:.....:7,...\n" + DoubleTab +
	".....I7D+.=8OOZ$ZDZ.....?~....\n" + DoubleTab +
	"......:7N?Z8888O88Z.....I.....\n" + DoubleTab +
	"........ONDDD88O8DD~...=+.....\n" + DoubleTab +
	"........:8DDDDD8DND87=~$,.....\n" + DoubleTab +
	"........=8DDDDDDDD8++I$$......\n" + DoubleTab +
	"........$8DD8888O8NO..=:......\n" + DoubleTab +
	".......=O8DD8OOODZ8M7.I.......\n" + DoubleTab +
	"......,ZODDDDD88NO8M$~=.......\n" + DoubleTab +
	"......$O8DDDDDDDDDNM$+,.......\n" + DoubleTab +
	"....,$OZ8DDDDDDDNDDMO=........\n" + DoubleTab +
	"...~ZOZ8DDDNDNDDNDNMO,........\n" + DoubleTab +
	"...,7OODD8DDDNDNDDNM$.........\n" + DoubleTab +
	"...~OO8DDDD8DDDD88NN?.........\n" + DoubleTab +
	"...+8DDDDNNNDDDNNN8I..........\n" + DoubleTab +
	".....?DDNNMNDDDDNN$=..........\n" + DoubleTab +
	"......:$MDOZZ8NDN8?:..........\n" + DoubleTab +
	".......OD:...:==N$+...........\n" + DoubleTab +
	"......~M?......,NZ:...........\n" + DoubleTab +
	"......$M+......=M8............\n" + DoubleTab +
	"......$D=......=NO............\n" + DoubleTab +
	"..............................\n"

const wizardAsciiC string = DoubleTab + "==============================\n" + DoubleTab +
	"==============+I?=============\n" + DoubleTab +
	"=============?DDDI============\n" + DoubleTab +
	"=============OO~7D+===========\n" + DoubleTab +
	"=============88?$M7===========\n" + DoubleTab +
	"============+8NODN8+==========\n" + DoubleTab +
	"============II+=+8O7=~~=======\n" + DoubleTab +
	"============?~?++7DD8ZZZ======\n" + DoubleTab +
	"~=~===~===~?7+:=I778OOD?~=~~~~\n" + DoubleTab +
	"~~~~====~~+7IZI7O?:+I+=~~~~~~~\n" + DoubleTab +
	"~~~=+???+??IZIZ+IOO=:~~~~~~~~~\n" + DoubleTab +
	"~~~+III77I=7I~8~:O8?:~~~~~~~~~\n" + DoubleTab +
	"::~+III7II+=~+8?.++=::::::::::\n" + DoubleTab +
	"::~=?IIII?~+=?87,~:+::::::::::\n" + DoubleTab +
	",,::=++==~:+=+8$~~:+~,:,,,,,:,\n" + DoubleTab +
	",,,,,:::,,:~+?8$+~:=~,,,,,,,,,\n" + DoubleTab +
	",,,,,,,,,,~~7$OO$~:==:,,,,,,,,\n" + DoubleTab +
	",,,,,,,,,,~:$D8DD:,+=~,,,,,,,,\n" + DoubleTab +
	".........:~:ZDOOO=:=+~:.......\n" + DoubleTab +
	".........+I788Z?$?$Z7==,......\n" + DoubleTab +
	".........~$$88OIO7$DD8?,......\n" + DoubleTab +
	"..........+ZDD8DDDN8ZI:.......\n" + DoubleTab +
	"............~$8ND8NI..........\n" + DoubleTab +
	".........,:=+7DND8N7?+~,......\n" + DoubleTab +
	"..........:=I$$77$$I+~:,......\n" + DoubleTab +
	"..............................\n"

const wizardAsciiD string = DoubleTab + "====================\n" + DoubleTab +
	"=====~=====~========\n" + DoubleTab +
	"=====+~==~=?========\n" + DoubleTab +
	"====~Z$==+?OO=======\n" + DoubleTab +
	"===~+8$==Z$$D7~=====\n" + DoubleTab +
	"===~+8+~=$IIZD?~====\n" + DoubleTab +
	"===~+Z~+$7=?$8OZ=~==\n" + DoubleTab +
	"===~IZ~ZO$OZZZZOI===\n" + DoubleTab +
	"===~7I~ZZ$Z77$77ZI~=\n" + DoubleTab +
	"====ZZ7ZI$$7$7Z8ZI~=\n" + DoubleTab +
	"===~ZZOO$$Z$$$8IZ8+~\n" + DoubleTab +
	"====Z=~Z8OOZ7Z8I$N?~\n" + DoubleTab +
	"==~=Z=IOOOOZ$$8$ID+~\n" + DoubleTab +
	"==~+Z?88OO$7$$OZ78+~\n" + DoubleTab +
	"==~?$I8OZZ$$$$OOZ7==\n" + DoubleTab +
	"==~I$~$8$Z$$$$O8Z7~=\n" + DoubleTab +
	"==~I7~Z8O$$Z7OZO8I~=\n" + DoubleTab +
	"==~7I~ODZZOZ$888Z===\n" + DoubleTab +
	"==~7?~ZDO88OZO88Z===\n" + DoubleTab +
	"==~$?~O8ODDOZ88O8===\n" + DoubleTab +
	"==~$+=O888OZ8D8O8+~=\n" + DoubleTab +
	"==~$+?8888DD888O87~=\n" + DoubleTab +
	"==~$==ODODD88O88D$~=\n" + DoubleTab +
	"===$=~+Z8N$78ODD7===\n" + DoubleTab +
	"===Z++788Z7?IIONO===\n" + DoubleTab +
	"======??=~=====?I===\n"

const thiefAscii string = "............................................................\n" +
	"............................................................\n" +
	"...............,:...........................................\n" +
	"..............~OD?...............~I7~.......................\n" +
	"..............IND?...........:+=?ZO88I......................\n" +
	"..............+NN8I,.....~II?7$I$$ZOO8......................\n" +
	"..............$+ZMNOI~::7OOO8D$788I?88~.....................\n" +
	".............~7..ZMN8ZO8DD88DN8DN8Z8MOOZ~...................\n" +
	".............?+...~$DNNNNDDNN88DDDNND8ZO:...................\n" +
	".............7......,+?++:DNDDO88DDDOOZ?....................\n" +
	"............,=..........~+ZD8D8O88D8O888+...................\n" +
	".............,...........~$8DDD8O8888D8NZ...................\n" +
	"..........................:ZZODD8DDMDD88$...................\n" +
	"........................:7IDDZZ88O8D8OOO$,..................\n" +
	".....................:+?78NDDDNDDD8DNMNDO,..................\n" +
	"....................,$ODNNNNDDDNNNNDZ77?:...................\n" +
	".....................ZD8888DDNNNMNDD8I~,....................\n" +
	"...................I?$Z8888DDDNDNNDDDDD8Z=..................\n" +
	"..................?ND8O8NDDMMMMMMNNDDDDDD8I,................\n" +
	"..................8NDDDDDMMO+I+?8NNMMND8DDND:...............\n" +
	"..................ONDDMMN8~.....,$$ZODMMNDDMI...............\n" +
	"..................8NDDO=:.........+7I$+?DNNN8...............\n" +
	".................,DNDMZ.............,...IMMNM?..............\n" +
	".................:NNDO~..................~$MMN,.............\n" +
	".................=MN8,.....................+NMO,............\n" +
	".................7MD~.......................=MMN?...........\n" +
	"................~NN7.........................8MMMN7:........\n" +
	"...............+DNM7.........................,~=$8DO=.......\n" +
	".............?NMMNO:........................................\n" +
	"............:$$I?=..........................................\n" +
	"............................................................\n"

const thiefAsciiB string = DoubleTab + "...............:,.............\n" + DoubleTab +
	".............,O8Z=............\n" + DoubleTab +
	".............7MZZ8:...........\n" + DoubleTab +
	"............=OM8O8+...........\n" + DoubleTab +
	"..........:ZDNND8$OI..........\n" + DoubleTab +
	"..........ZDZ88O88OZ..........\n" + DoubleTab +
	".........,8O8NDDD888~.........\n" + DoubleTab +
	".........INODMDDNDDNZ.........\n" + DoubleTab +
	".........8N8DMNDDNND8,........\n" + DoubleTab +
	".........7NODNND8DNDD,:~,.....\n" + DoubleTab +
	"........:,Z8DDD88DOODIZD?,....\n" + DoubleTab +
	"........:?7D8ND8DDDO7~,~......\n" + DoubleTab +
	".........,$N8NN88ODI..........\n" + DoubleTab +
	"......~+I$7$88OZD8?...........\n" + DoubleTab +
	"......:~~:..=Z?7DN~...........\n" + DoubleTab +
	"............=NOIIO~...........\n" + DoubleTab +
	"............ODODDO7,..........\n" + DoubleTab +
	"...........,DD8DM7~?~,........\n" + DoubleTab +
	"............ZN8$N8=,~+~,......\n" + DoubleTab +
	"............IMO.:7O=..,,......\n" + DoubleTab +
	"............ZDZ...............\n" + DoubleTab +
	"...........IND=...............\n" + DoubleTab +
	"...........+O?................\n" + DoubleTab +
	"..............................\n"

const thiefAsciiC string = DoubleTab + "............................\n" + DoubleTab +
	"............~++~............\n" + DoubleTab +
	"...........+O?7Z,...........\n" + DoubleTab +
	"............I7+$+...........\n" + DoubleTab +
	".............I8ON$+~,.......\n" + DoubleTab +
	"...........,IZOO8DND8+......\n" + DoubleTab +
	"...........?DOOO8DD8N7......\n" + DoubleTab +
	"...........$NDDDNNMN8$......\n" + DoubleTab +
	"..........,DMNNNNM7OND:.....\n" + DoubleTab +
	".........:ZNDDNNM7.~NNI.....\n" + DoubleTab +
	".........$DDDDNNNI.=NDO.....\n" + DoubleTab +
	".........OND8O8OO8,=D8=.....\n" + DoubleTab +
	"........+NNNNNNDDN~?8$==....\n" + DoubleTab +
	"...:::~+8MNDD8NNNM8DD7~.....\n" + DoubleTab +
	"...~?====OND88MMNNM8+.......\n" + DoubleTab +
	"........:NN88DMNNNM8........\n" + DoubleTab +
	"........IN88DNDNDNDD:.......\n" + DoubleTab +
	".......+D888NNN$8M8D$.......\n" + DoubleTab +
	".......ZD8O8ND8?:DNN+.......\n" + DoubleTab +
	".......=DD+ONN$..=NNZ,......\n" + DoubleTab +
	"........7O.ZNNZ..,ND8=......\n" + DoubleTab +
	"........,:.?MN+...7MD?......\n" + DoubleTab +
	"...........+MO.....8M?......\n" + DoubleTab +
	"...........IN$.....+M$......\n" + DoubleTab +
	"..........~OOZ.....,DO:.....\n" + DoubleTab +
	"......,~?ZDNNO=~~~=?O8:.....\n" + DoubleTab +
	"......:=+++=~:::,,:$8Z:.....\n" + DoubleTab +
	"............................\n"

const thiefAsciiD string = DoubleTab + "...................................\n" + DoubleTab +
	".................~=:...............\n" + DoubleTab +
	"................7OOZ+..............\n" + DoubleTab +
	"...............~Z$??7,.............\n" + DoubleTab +
	"..............,I8?++I~~:...........\n" + DoubleTab +
	".............?O88$?7$7Z7...........\n" + DoubleTab +
	"...........+O888OZI7O88$=..........\n" + DoubleTab +
	"..........,O7Z8O$?I+$$O$DO+........\n" + DoubleTab +
	".........=D8OZ7IZ?=I$7ZO?78I,......\n" + DoubleTab +
	".........:OOZZO888O88887,.:7?:.....\n" + DoubleTab +
	".........~I?$8DDDDNDOZ8$~.,~~7,....\n" + DoubleTab +
	"......:+??78DDDD8DD8O?7Z+?I..,.....\n" + DoubleTab +
	"....,=+~..$888D888D8O$7OZ=.........\n" + DoubleTab +
	"....,,...~88888DDZ$778OOZ=.........\n" + DoubleTab +
	".........78O888DDDOODO877O.........\n" + DoubleTab +
	"........:8O8OOOODDDNDOO$78:........\n" + DoubleTab +
	"........$DOZZO88DD8DD8$I+O:........\n" + DoubleTab +
	".......=8OZ$Z8DDD888Z8O$??.........\n" + DoubleTab +
	"......,ZO7Z8DDD8888D?:NZI7,........\n" + DoubleTab +
	"......?8O8DDDD88OOO8?.Z8=?I........\n" + DoubleTab +
	".....:8888DD8888ZZ$8+.:ZO$Z$:......\n" + DoubleTab +
	".....ZDD8DD8888OZ$$O=...INO$Z~.....\n" + DoubleTab +
	"....IDDDDD8O8OOZ$7$Z~....ON8O$.....\n" + DoubleTab +
	".,,+88DDDOOOOOO$77$$~....,$N88,....\n" + DoubleTab +
	".:?ZZ8DD8OOOOOZ$$ZOZ7=:....IND~....\n" + DoubleTab +
	"...,?8ND?++??I???+=:,.......7N$....\n" + DoubleTab +
	"..:$88$?====~~~=~~::,,......:DD~...\n" + DoubleTab +
	"..,+~,.........,,::~~~~~~~~~~887...\n" + DoubleTab +
	"........................,,:~~ZDZ...\n" + DoubleTab +
	".............................~ZI...\n" + DoubleTab +
	"...................................\n"

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

const paladinAsciiB string = DoubleTab + "..............:~,.............\n" + DoubleTab +
	"........,+~.?DMID+............\n" + DoubleTab +
	"........,MI.?DNZNZ............\n" + DoubleTab +
	".........M+.~MNDNN:...........\n" + DoubleTab +
	"........,N+.,8MODD?=:?I~......\n" + DoubleTab +
	".......,=77~~?8NO?~==?O=......\n" + DoubleTab +
	"......=OZDNI==7ZZO7=$8I.......\n" + DoubleTab +
	"...~~+Z8888==:=8I7IZ8NZ.......\n" + DoubleTab +
	".....?MMD.IN8Z?M8Z=?7D:.......\n" + DoubleTab +
	".....,8MD.?MMMOO=~,.+?........\n" + DoubleTab +
	"......:$O.+MM8DZI?+$:.........\n" + DoubleTab +
	"........7,=MM?DODM$Z,..:......\n" + DoubleTab +
	"........Z,=MIIOO7?I=.,$7+,....\n" + DoubleTab +
	".......,D,=MZ:$O+.DN7OMD7,....\n" + DoubleTab +
	".......~M,~MN,$O,=8+ODD+......\n" + DoubleTab +
	".......7M:~MN,$O.$N.?M=.......\n" + DoubleTab +
	".......ZM:~MM:$O.8M:IM,.......\n" + DoubleTab +
	"......~+M~:MM~$O,8M:?N,.......\n" + DoubleTab +
	".......:M=:MM~$O.OMN?==.......\n" + DoubleTab +
	".....,.~M+,MM~$O.ZMMM~........\n" + DoubleTab +
	".....,.+M+,MN,ZO.ZMNNO:++.....\n" + DoubleTab +
	".......ZM?,M8.ZO.7M=7M7$$.....\n" + DoubleTab +
	"......,MM7.MO.OO,~~.=M$.?.....\n" + DoubleTab +
	"....,.ZD$?,MN=I7:....N7.=.....\n" + DoubleTab +
	"....,=$,..,MM~.......N$.=.....\n" + DoubleTab +
	"....,=....,MI........$8.?.....\n" + DoubleTab +
	"....,...?O:M+........IM?7:....\n" + DoubleTab +
	".....:$DMMIMI........ZO~:=,...\n" + DoubleTab +
	".....,~::,,,..................\n"

const paladinAsciiC string = DoubleTab + "~~==~=???I77?++I77777$DDDDD?7ZZZZ$77$7$Z\n" + DoubleTab +
	"==++========~~~~=???I8DNNNMZ?$$$$$7$$7$Z\n" + DoubleTab +
	"=+?+???I?+=~:::,:~~~~ONNMNN$+7777$$$777$\n" + DoubleTab +
	"=~~~~~~~~~~~::::,,,,~$DNDD8$IIII77$$7$ZZ\n" + DoubleTab +
	"~~====~==::::::,~+?$O888DDDOZZOZ$$II77$$\n" + DoubleTab +
	"~=+I?==+++~::,,:$D8DDNNODMNDDDDDDD7=+777\n" + DoubleTab +
	"???I7$$??I+=~:~IONDDDNNMMMNNNDDNND8$+?77\n" + DoubleTab +
	"$$7IIIO77II++I7ZO8NNNNNNNNNNDNNN8DDO7?II\n" + DoubleTab +
	"$77II?NZ=++=Z8IZ8OZ8NNNNNNNNND8O8NOOD+II\n" + DoubleTab +
	"7IIII?DO~===ONOON88O8N8DNMN$7ZONMNDMN$+I\n" + DoubleTab +
	"III?++DO~+=INDN8DNDNODZI8M8Z$O8DNMMMMO+I\n" + DoubleTab +
	"I?+=~~NO:=?ZDDMOON88$8MNDMNMMDDNNNMMM7+?\n" + DoubleTab +
	"?++~~8NOI$D88NMZ$ND8DOZ7DMMOZDNNMNNDD$~:\n" + DoubleTab +
	"~::,+MMMMNDNM8?~=ZND8DOIOM877ZDMN=ONDO:,\n" + DoubleTab +
	"~=?IIZMMMMD87,,::=NMNNNMMMNNMNMNI.ZN8D=,\n" + DoubleTab +
	"~==~:=M8$$?:.:~=~ZMNMMMMMMMMMMMOZ.INDN?:\n" + DoubleTab +
	"=:~~~+MI,,:~====IMD8DNMMMMNDDDNOD:+NNN?~\n" + DoubleTab +
	"+===~+N7=====~~=DMDDZ$OODNNMNNNDN?+NND+=\n" + DoubleTab +
	"??++=?N7~~:::::IMMDNDDO88MDOD8NMM$$NNNI?\n" + DoubleTab +
	"????+?N7~~~::::8MNDNNNDDD888$?78N8ODND$7\n" + DoubleTab +
	"+??+=?N$:::,:.IMMNDNDDDDDO888Z8Z8DNMM8$Z\n" + DoubleTab +
	"==+?+?N?,,,,,~DMMNNNDDDNNMN$88OOMNNMMNZZ\n" + DoubleTab +
	"++++??D7:::~~7MMMNNNDNDNNMM8Z88DDMMD8OZ$\n" + DoubleTab +
	"I?++??D$=+=?IMMMND8NDDDDMMMMO88DDNMOI7$$\n" + DoubleTab +
	"Z7II?IDZI7I?IZDMD8ONDN8DMMMMN8ND8ODD$ZZZ\n" + DoubleTab +
	"8OOOZZDOZOO$7Z$ND8ONNN8NDMMMM8O8OOONZ8D8\n" + DoubleTab +
	"DDDD88DOO88OOOOND88NNNNMDMMMMM88DDZ88ODD\n" + DoubleTab +
	"DDDDDDOO8888O88NDDDNNNNNDNMMMMNNMDO8NODN\n"

const paladinAsciiD string = DoubleTab + "...............:~:............\n" + DoubleTab +
	"..............~$7Z=...........\n" + DoubleTab +
	"..............78ON$...........\n" + DoubleTab +
	"..............?OZD7...........\n" + DoubleTab +
	".............=788OO$7,........\n" + DoubleTab +
	"...........?$888D8DD$?~.......\n" + DoubleTab +
	"..........I8OO?$$DDD8??=......\n" + DoubleTab +
	".........+O8Z7=77DNDDZ??:.....\n" + DoubleTab +
	"........,$ZD88ZDDOZ$7ZZZZ,....\n" + DoubleTab +
	"........~7N8DNMO77OMOZZO?.....\n" + DoubleTab +
	".......:78?+N888ODNN,.........\n" + DoubleTab +
	".......I$Z,+DO8NMMN8+,........\n" + DoubleTab +
	"......+O8NZ78NNNDN8O7.........\n" + DoubleTab +
	"......7Z8NNDODDDDDDNN+........\n" + DoubleTab +
	"......??8DDD8ODNDN88NN:.......\n" + DoubleTab +
	"......$ZDDDD8O8ND88ODDO,......\n" + DoubleTab +
	".....,IIDDD88O8DD88OD87Z......\n" + DoubleTab +
	".....,7ZD88OO8NND8DOON+$7.....\n" + DoubleTab +
	".....:ZOD888ODDZ?=NND8~.8=....\n" + DoubleTab +
	".....~OO8D8D8M?...8DO8,.:Z,...\n" + DoubleTab +
	".....=OODD888M=...INZ8=..=$...\n" + DoubleTab +
	".....?8OD888DM~...:DO87...II..\n" + DoubleTab +
	".....IO8N888N8.....$8OZ....=..\n" + DoubleTab +
	".....$8DDD8OM?.....=8OZ.......\n" + DoubleTab +
	".....ZDDDOZO8+:~===?OZZ~::,...\n" + DoubleTab +
	".....ZDDDI7OO7I7777I$$$?~::...\n" + DoubleTab +
	".....$MN8O8DZI?=:~~~I7O+..,,..\n" + DoubleTab +
	".....~7I?==:,.......+7$~......\n"

const barbarianAscii string = DoubleTab + "...............................\n" + DoubleTab +
	"....,?+~??+:...................\n" + DoubleTab +
	".....=O7I7$I?~.................\n" + DoubleTab +
	".....:+7O77+??~................\n" + DoubleTab +
	".....=$O8$7???7~...............\n" + DoubleTab +
	"...=I+7ZIII?$Z?+=,~,...........\n" + DoubleTab +
	"...,+II?~:~?887?$$N8+:.........\n" + DoubleTab +
	"......I777$8DO$II,+N$?7I?=.....\n" + DoubleTab +
	"......ID88MN8Z$ZZ778?.,?7+,....\n" + DoubleTab +
	"......7M88NND8ONDDDM+..........\n" + DoubleTab +
	"......8DNMND8OZNOD++,..........\n" + DoubleTab +
	".....:MDMMDDDDONMMI............\n" + DoubleTab +
	".....?M8DMNDO8DNMMN,...........\n" + DoubleTab +
	".....8MN8DND$7O8MMM7...........\n" + DoubleTab +
	"....=MMMDDDD8OODMMMM:..........\n" + DoubleTab +
	"....IMNNMMNNNND8NMMMZ..........\n" + DoubleTab +
	"....$MMMMMMMMNN88MMMM+.........\n" + DoubleTab +
	"....8MMMMMMMMNNN8NMMMD,........\n" + DoubleTab +
	"...~MMMMMMMMDNMMDDMMNMZ........\n" + DoubleTab +
	"...7MNMMMMDD8DDDDNMMMNMI.......\n" + DoubleTab +
	"..,DMMMMMMND8NDO$OMMMMMM$......\n" + DoubleTab +
	"..?MMMMMMMNDMNNDMMMMMMMMMZ.....\n" + DoubleTab +
	".,8MMMMMMMDDMNDDMMMMD8O88N7....\n" + DoubleTab +
	".~MMMMNMMMDNMNDNMMO?,.....:,...\n" + DoubleTab +
	"..?8DMM8=DMMNNN8=,.............\n" + DoubleTab +
	".....:I,.=++ZMND7:.............\n" + DoubleTab +
	"............,+?$OO+............\n" + DoubleTab +
	"...............................\n"

const barbarianAsciiB string = DoubleTab + "...................................\n" + DoubleTab +
	"........................=?I~.......\n" + DoubleTab +
	"......................:7DD+..:+....\n" + DoubleTab +
	"......................+DD8$7~78,...\n" + DoubleTab +
	".............,.,:~=~.?ZI?O888DO,...\n" + DoubleTab +
	"...........,I$$ZODNI.Z=.7Z+ODO+....\n" + DoubleTab +
	"..........:ZD888NNNDZ==I$,+Z~~.....\n" + DoubleTab +
	".........+OD8888OOZODZO8Z$7~.......\n" + DoubleTab +
	".......~$OO8OZ7$77ZO8ZOOO8O=.......\n" + DoubleTab +
	"......:8OO887=+I888OODN8OZO8=......\n" + DoubleTab +
	"......$OZZI7Z7I8D8OO8D88OZZZ$?.....\n" + DoubleTab +
	"....,+7$?+I?IO8DNO8DDDD888DOOD=....\n" + DoubleTab +
	"....:O7==$8$O8NDOONNNNNNNNNNDZ~....\n" + DoubleTab +
	"....,77?IO88D8NO8NDNNNMO:78$~......\n" + DoubleTab +
	"...,7OOO88DDDNOONDDNNNZ,..?,.......\n" + DoubleTab +
	"....7ZOOZ8DDDDDDDDNNNMZ...,........\n" + DoubleTab +
	"....Z77?Z888D8DNDNDNNNN+...........\n" + DoubleTab +
	"....?8ZZOZI8NDDD88DNNND$:..........\n" + DoubleTab +
	"....~88DD=ZDODNDO8NNNDNO...........\n" + DoubleTab +
	"......,+++D7ZDND8DNNND8N~..........\n" + DoubleTab +
	"........,I8$ONNNNN8$MDDM:..........\n" + DoubleTab +
	".........,ONNDMZ7Z:.NNNI...........\n" + DoubleTab +
	"...........~ONN8:...$MN7,..........\n" + DoubleTab +
	"...........~ZDDNO$OOODDN8Z?:.......\n" + DoubleTab +
	".......=?=+ODDDDDDDD8DNDNN8?,......\n" + DoubleTab +
	".....,Z8NMDDND88OZZ$$777?~:........\n" + DoubleTab +
	".....::.+=$+~~:::,,,,..............\n" + DoubleTab +
	"...................................\n"

const barbarianAsciiC string = DoubleTab + "................................\n" + DoubleTab +
	"................,++,............\n" + DoubleTab +
	"................$Z$I............\n" + DoubleTab +
	"................I??I............\n" + DoubleTab +
	"................I+7D?:,..~,.....\n" + DoubleTab +
	"..............=ONODNDOZ$I,......\n" + DoubleTab +
	".............7NND8D8ZZOOO?......\n" + DoubleTab +
	".............ZD88$ZZ88$OZ?......\n" + DoubleTab +
	".............788$O8O88NZ$I,.....\n" + DoubleTab +
	".............+D88DZNDNI7O$~.....\n" + DoubleTab +
	".............7MNNNONMI.?N8$.....\n" + DoubleTab +
	".............8NNNO8DM7.?MN$.....\n" + DoubleTab +
	"............7MMNND8MN8,:MM87....\n" + DoubleTab +
	"...........+8DN8DMNND8~,ND8=....\n" + DoubleTab +
	"...........ZZ?OZ8MMNDD?I8I~.....\n" + DoubleTab +
	".........:7+.~MNNNNNNN8D$~......\n" + DoubleTab +
	"........I$:..?MD8NNDDDNI........\n" + DoubleTab +
	"......~Z$,...IMNNMNDN8DZ?.......\n" + DoubleTab +
	".....?OI.....+NNNNMNNNNDI.......\n" + DoubleTab +
	"...:7?,.......ZMNN7=NMN87,......\n" + DoubleTab +
	"..:+:.........7MNN~.=MMMZ,......\n" + DoubleTab +
	"..,...........=MN8...8MM7.......\n" + DoubleTab +
	"...............DMZ...$MN+.......\n" + DoubleTab +
	"..............:NMZ...+MD=.,.....\n" + DoubleTab +
	"............,INMMZ...?MD,.,.....\n" + DoubleTab +
	"............?8Z?,...,NMD:.......\n" + DoubleTab +
	"....................,?I~........\n" + DoubleTab +
	"................................\n"

const barbarianAsciiD string = DoubleTab + "...............,..............\n" + DoubleTab +
	"..............~+=~:,..........\n" + DoubleTab +
	".............=$$7I?=~,........\n" + DoubleTab +
	"............~77Z77$?:.........\n" + DoubleTab +
	"............,77Z77Z$+,........\n" + DoubleTab +
	".............+$Z$7OZI?+~......\n" + DoubleTab +
	".............~O$$7ZZ7ZZ7+,....\n" + DoubleTab +
	"...........,?78O$$8$?7Z8OOI,..\n" + DoubleTab +
	"........,:,?ZZ7ZOO8Z7$8OODZZ?,\n" + DoubleTab +
	".......:ZI$IZ$OOO$$ZZO8?.?D8$I\n" + DoubleTab +
	".......+OZ8D?.IDOZ$$77$,~O8Z7~\n" + DoubleTab +
	"......,$88DO.=ZDOZZZZ88.=?78+.\n" + DoubleTab +
	".....,$8DMZ~I888DDDNOZZ?...~,.\n" + DoubleTab +
	"....:ZDDD?,$DDNNNDDN8Z8D~.....\n" + DoubleTab +
	"....,ODD7,IDNDNN88D888DD7.....\n" + DoubleTab +
	".....IO$=ZDDNNMN88OO8O87=:....\n" + DoubleTab +
	"...,I8O,I8DDNNND8DD8OZO8+:....\n" + DoubleTab +
	"....?8?.7Z8DNDDNDDDD88ZZ7,....\n" + DoubleTab +
	"....7$..IDDNMNNNO$ZDMDD8?=,...\n" + DoubleTab +
	"...:O=..~DNMMMM8Z$Z+=:~?::,...\n" + DoubleTab +
	"...I$....,~78MDZZ8?...........\n" + DoubleTab +
	"..:Z+.......+ND8OD$...........\n" + DoubleTab +
	"..?O........$8ODNDI...........\n" + DoubleTab +
	"..O?........~O8DD8............\n" + DoubleTab +
	".?O.........IDDNN$............\n" + DoubleTab +
	",8?........,O8DN7.............\n" + DoubleTab +
	"$8,.........$8DD7.............\n" + DoubleTab +
	"I:..........:++++.............\n"

const dragonSkillFireAsciiA string = "II??I??II$Z88888OOOOOOO88DDDDDD8$?+++++=++++++++++==?7O8DD8888888OOZZZZ$$Z$7I??I\n" +
	"7I?+?I7$ZOO88OOOOOOOO88888888DDDD8$?+===++++++++++?$O88888888888OOZZ$I????I??+++\n" +
	"7I++?I7$ZOOOZZZZZZZOO888OOOO888888887+++++++++++?$8DD8888888O8888OZ$?+??I??++???\n" +
	"7I?+??I??77II?I7$$ZZOOOOOZOO8OOOOOO88OI====++++IODD8888888888Z$$ZOZ7???IIII?????\n" +
	"7I?+??++=+?+++???7ZZOOZ$ZZZO8OOOOOO888DO7?I$$7ZODD8888888888Z???I7$7IIII77IIIII?\n" +
	"I?++++==+++????I??7OZ7?II$ZO8OOZZ$ZZO8DDDOO88O8DDDD88OOO888OIIIIIIIIIIIII???????\n" +
	"?++++++???++?IIII?I7I??I?I$OOOZ7I??II$ODD888ZO88Z$7777$ZO88$II7IIIIIIII??++++++?\n" +
	"+???????????+?II???I??IIII7OO$I?III7I?I$888OZOZII???????7$8$III77II??+++++++++++\n" +
	"?II?IIIIIIII??IIII????II7I7Z$I?IIIIII???$8OZZZIII7777IIIII7$II??????++++?+++??++\n" +
	"III?IIII77777II77III??IIII7$777I?IIIIIII$8Z+7O77777777777777III?????????????????\n" +
	"III?I777777777777777IIIIII777777777777I7OOI~+777777777$777III??II?????IIIIIIIIII\n" +
	"77IIIIII77$777777777777II777I7777777777$O$?==+III7777III7I??IIII77I7777IIIIIIIII\n" +
	"777777IIII7$$$$$77777777777777II77777777$$+~+=IZ7??????IIII7777I77777IIIII??IIII\n" +
	"777777777777777777777777777777II7777IIII7Z?=+=+7I??III777777777?I777IIIII???IIII\n" +
	"777777777777777I7IIIII77777IIIII7777IIII$ZI?+=~+IIIIIII77IIIIII7$7?IIIIIIIIIIII+\n" +
	"7777777777777777777IIIIIIIIIIIIIIIIIIIIIII?I++=+??III7$ZZ$I????II?IIIIIIIIII?+++\n" +
	"7777777I7777IIIIIIIIIIIIIIIIIIIIIIIIIIIIII??+I++IIIIII777I+???++?+?7IIIIIII?+=++\n" +
	"II7IIIIIIIII7IIIIIIIIIIIIIIIIIIIIIIIIIIII?+==I++III??????+++++++++?IIIIIIII?++++\n" +
	"77IIIIIIIIIIIIIIIIIIIIIIIIIIIIIII7777IIIII?++??+??++=+==+++?+==+++III????II+++++\n" +
	"ZZ$777IIIII???II??IIIIIIIIIIII???I77I?III???I???+?+++?++?+=++++??II?II?++++???=+\n" +
	"ZZZ$$$7I????+++?++?I$$7??IIII??+++?++?IIII?+???7?+=+II??II?+?+???????III?+?III?=\n" +
	"ZZZ$77I++????++++??I$$7???????+?+++++?????++?+II?+==+?+++??++++++++?II77?=?$77?+\n" +
	"ZOZ$$I???++???+++??????++++++?++++++?I???+===+??++====++++?+??+++++III777?+I7I?=\n" +
	"OOZ$I?77$I?I?+??I77??III?++??+++++?II7I?+==~~===++++??++??I??+++==+?II7$$7?III7I\n" +
	"D8O$$$$7$77II?$OO88OZ$ZZ7I???+++=+?III?+===+=++++++????+????++++????7$7III7ZO8D8\n" +
	"D88O8OZZOOOOO88DD8OOOOZZZZ$77II???II77???++????+++??++????I7$$$$ZZZO888ZZO88DD88\n" +
	"88D8DDOZ8888OZ$$O888OOO8OO88OZZOOOZOO$7IIIIIII7$$$ZZZZZZOO8O888D888DDD8DDDDDDD88\n" +
	"8888O88ZOO88Z7??ZD888D8888888OO88888OZOOZZOZZZZOO8888888888OOOOO88888O888D8888DD\n" +
	"8DD88888888D8OOO88888888888888888OO8888OOOOOOOOOO88888O888888888888888O88DDDDDDD\n" +
	"D88DNDOOOO8DDDDD88888888888888888O8DD8888888888O888888888888888888OZO888D8DNNNNN\n"

const dragonSkillFireAsciiB string = "NNNNNNNMMMNNMNNMMMMMMMMMMMMMMMMMMMMNDDDDNMMMNNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" +
	"O88888OZ8ND88NNNNNNNNNNNNNNMMMMMND8OO8DNMMNNNNMMNNNMMMNNNMMMMMMMMMMMMMMMMMMMMMMM\n" +
	"7777$$OZ7IZ8DOO8DNNNNNNNNNNNNNDOZZZO8DNNMNNDNNMNDDNNMMNNDNMMMMMMMMMMMMNMMNNMMMMM\n" +
	"7$$77$$ZZ$??7O88OODNMNNNNND88OZZZZ8DNNNNNDDNNNNNNDNNNNNNDDNNNNNNNNNNNNMMMNNMMMMM\n" +
	"$ZZZ$$7$ZZ$II++7$ZOOO8888OZZ$ZZZO8DNNNNDDDNNMNMNNDDNNNNNDDDNNNNDDDNNNNNNNNNNNMMM\n" +
	"$ZZZZZZ$ZOOOZZ7++~???~~+?7ZOZZO8DDNNNNDDDDNMMMMNNDDNNNNMNDDDDNNDDDDDDNNNNNNNNNNN\n" +
	"$$Z$Z$$$Z$$OO8D87?++~~:,:++ZDNDDNNDDDDDDDNNNNNNNNDNNNNNNNNNDDDDDDDDDDDDNNNNNNNNN\n" +
	"ZZZ$Z$$$Z$O$ZO8NMN87??+~,:::=$NNDNDNNDDDNNNNNNNMNDDNNNNNNNNDDDDDDDDDDDDDDNNNNNNN\n" +
	"$ZOZZ$ZZ$Z$Z8OOODMD887$77I=?+~7MNNNNDD8DNNNNNNNNNDDNNNNNNNNNDDDDDD88888DDDDDNDDD\n" +
	"Z$7Z$Z888ZZOOZZZONDZMMNOZZ7+Z$~DMNNND88NNNNNNNNNNDDDNNNDDDDDDDDDDD88888DDDDDDDD8\n" +
	"II$ZOOO8DOZ8DN8OOOODDNNNDO$I?++7MMNND8DNNNNNNNNNNDDDNNDDDDDDDDDD8888888DDDDDD8OO\n" +
	"7O8DNDD88DO7ZZOD8888DDD88DOZ7+~:INMMDDDNNDDDDNDNNDDDDDDDDDDDDDDDD88888888DD88OZZ\n" +
	"$8DNDDNDDNMO7ZOZ$77$ODZ$8DNDO$?~.,+OMDNNNDDDDDDDNDD88DDDD888DDDDD88888O8888OZ$77\n" +
	"Z8DNDDDNNNDD888D87??+~..,=$DNNMD?~~:7NDNDD8DDDDDDDD888DDD8888888DDD88OOOOOZZ$7I?\n" +
	"$88DD8DNNOOO8NDNDN8I=,....,~$DNMM8$I?DDDDDDDDDDD8DD888DDD88888888DDD8OOZZ$77II??\n" +
	"$$OO8DN8OO8OO8NMNDN8?.....,,=I$$$Z7IZDNDDDDDDDDDDDDD8888888888888888OZ$77IIII???\n" +
	"ZZZZ8N8IO8O88O8NMMMN8=.....,,,:+I7$ZO8DDDDDD8DDDDDD8DD888888888OZ$$$777IIIIII???\n" +
	"7ZZODD$$OOOO8888DNMMNO+:,...,,..,+7$7?$88DD8DDDDDDDD88D888888OZ$7777IIIIIIII????\n" +
	"II7Z88IZZZZOO8OOOODMMNZ?:,,,...,..,=?I7$ZODDDDDD8888DO888888OZ7777III???IIII???+\n" +
	"II7ZOZ+ZZZZOOZZ$ZOO8NN87=:,,:,......:?IIZ$78D888888888888OOOZ77IIII????IIIIII??+\n" +
	"7$7$Z$?Z$$$$$$$7$$ZZO8N8?~:,::,,.....,=+II?$ZO8OO8888OOZ88OZ7IIII??????I?IIII??+\n" +
	"??77$7?$II7II??III77$7ZDDI::~~~:,,,....:,:~=?IZ888OOOOOO88O$77III?????I??IIII??+\n" +
	"?I$ZO$+?????+++????IIII7887,:+=:,,,,....,,,:~=?ZOOOOOZZZZ8OZ7IIII?????I??IIII?++\n" +
	"II7ZZ$+??+?+=+++??++IIIII$D$+++~:,,,.....,..,,:?77$$777I7ZZZ7?III??III???????+~:\n" +
	"I7$7Z$+???++?+?????+???III7$I===~:,,,.......,..:+?7I??III77$7IIIIIII???+++===:,,\n" +
	"I7$$ZZ7II??????+?++++++??II???===~:,,,......,...:~=??I????I7IIIIIII??++==:,,,,,,\n" +
	"II7$ZZ?+??I??+++=++==++===+++??=++~,,,.............,:~++?IIIIIII?++==~~~:,,,,,,,\n" +
	"77$$ZZ=+I??II++++++==+++=+++++++++=~:,................,:=?II????+=~::::,,,,,,,,,\n" +
	"$7ZZOO+IOZ$$$7????+===++++========~:::,.................,:~::~~~~~:::::,,,,,,,..\n" +
	"777$ZOI?8OZZ$$77II?+=======~~~~~=~,,,:,..,.....,,......,,,,,,,,:,::,,,,,,,,,,...\n" +
	"77$OOZI=OOOZ$$777I?+++=======~~=~::,,,,...,,,.,,,,,,..,,,,,,,,,,,,,,,,,,...,....\n"

const necromancerAsciiA string = "I7777777IIII77I777777777777I777I777$Z$$Z$ZZIIIIIIIIIIIIIIIIIIII77II7777777777777\n" +
	"77777777IIII77777III777777777II7I7$888D8D88OI77IIIIIII77III77II77777777777777777\n" +
	"IIII777II7III77IIII7IIIIIII$$$ZZZ888DNNNND8NZI7III7777II7III77III777777777777777\n" +
	"IIIIII77IIIIIIIIIII7$77Z$III$8DDDDNND8O7ODDDOIII7IIIIIII77IIIIIIIIII7777777II7II\n" +
	"IIIIIIIIIIIIIIIIIIII7$$$$777OD88D8DN8OZ$Z8DO$77Z7I777IIIIII7O7I7II77IIIIIIIIIIII\n" +
	"IIIIIIIII77IIIIIIIIIIIIII77$$OZ8D8DD8O$7$8Z7$777II77IIIII77I$Z$7I7$7IIIIIIIIIIII\n" +
	"IIIIIIIII77II7IIIIIIIIII7I7I??7$Z88DDNNNNO7ZO77III77$77II77I7I7III$$7$7IIIIIIIII\n" +
	"IIIIIIIIIIIII$$IIIIII7$$$IO8$$O7$Z7ONNNNDZZOZ$Z7III777IIII7$7III777$$7IIIIIIIIII\n" +
	"IIIIIIIIIIIIII$77II7IIIIIII7$$78OZ8DDNND88OOO8ZZI7IIIIIIIIIIIIII$7IIIIIIIIIIIIII\n" +
	"IIIII?II?IIII???IIIOOOZII$II77Z8NNNDDNDDDDDOOOOZ$$7II7?IIIIIIIIIIIIIIIIIIIIIIIII\n" +
	"???I77$I??????I$I?IO88OZ$II$7Z$ZODDDDDDDDNDDDNZ$$II?7II?7+I????II?IIIIIIIIIIIIII\n" +
	"???????????????????I8888Z77$$8888DDDDDDDDDDDDDOZ7??I??Z7$7????I7I???????????????\n" +
	"????????II??I?+7I??7888888$ZOOO8D8888OO88DDND88Z??777Z88O7??????????????????????\n" +
	"?????????7??I$Z$???7D7Z88O$ZOZ8888888888OZ$Z8DOO8OODDDDD8I??I??+I???????????????\n" +
	"??????+??+I?+I77I7Z$Z?I$ZZDO$OOO88DDDDDDD87?+IZ888DD8DDZI?I?I?I?I????????+++?+??\n" +
	"??????+?+78?+II????+??+I788$$Z$ZO888DDDD888Z?=+?I78D$$$I?II?I????????++?++++++++\n" +
	"??????$$7ZI+I$Z?+?I7?I7I887II7$ZOOO88D888ZOZ+:I7O?$88???????7???+??+++++++++++++\n" +
	"????????I?+?I?I??I$I+$7I$$??I7$ZZ$ZO8OOOZZZ77~:=$ZZ$7?7?++++7$??+??I?+++++++++++\n" +
	"?????+??+++?+++++?+?+?7++++?+?I7I7$$ZZ$$7$$OZ+:::+7?I7I??I??$8I???+?++++++++++++\n" +
	"????+?++?+???II???++?+II??I?+?I7II7I77I$7I7$II::,~$8Z??II?II??????++++++++++++++\n" +
	"?????++???+?$I??++++++++++?7III?III$8$IIIII?I7=:,=?$DZI7II?++++?+?+++++++++++???\n" +
	"?????+++++++++++??++++II++++??++??$DNN8ZZ$77?I$?+??+OO$7$$??++++++++?IIII???7Z88\n" +
	"$$$77IIIII??++++I$I+?II++==+++7$$ZDDDN8$$I?++IZ7+?++++++?+IZ+?++++?7O888OOZ88DDD\n" +
	"888888OOZZZZ7I?+++?I?+=++??=++?7OND8DDDZI+++?++7+=+++??I?Z8?IIII?I$8DDDDDDDDDDDD\n" +
	"888888888OZZZ$$$777I?????+I+++===ZDZ$Z8$++==~~~~=====?7$$Z77$ZOOOODDDDDDDDDDDDDD\n" +
	"DDD8DDD888OOOOZ$$$$$7I777I77II?+?ZI+==IO+~~~~==~~=?I7$$ZZZZO8DDDDDDDDDDDDDDDDDDD\n" +
	"DDDDDDDDDDDDD888888888888OOOOZ$7$7I7II+7$?I?7$Z77Z8DDDDDDDDDDDDDDDDDDDDDDDDDDDDD\n" +
	"DDDDDDDDDDDDDDDDDDDDDDDDDDDDDD88OOOOOZOZOOO888DDDDDDDDDDDDDDDDDDDDDDDDDDDNDDDDDD\n"
