package main

import "math/rand"

type ASCII struct {
	GOBLIN   []string
	SKELETON []string
	SORCERER []string
	ORC      []string
	DRAGON   []string
	HERO     []string
}

var AsciiArts = &ASCII{
	GOBLIN:   []string{goblinAscii, goblinAsciiB},
	SKELETON: []string{skeletonAscii, skeletonAsciiB},
	SORCERER: []string{sorcererAscii, sorcererAsciiB},
	ORC:      []string{orcAscii, orcAsciiB},
	DRAGON:   []string{dragonAscii, dragonAsciiB},
	HERO:     []string{heroAscii},
}

func (a *ASCII) setImage(name string) *DisplayImage {
	var image *DisplayImage
	switch name {
	case enemiesList.GOBLIN:
		image = &DisplayImage{
			Image: a.GOBLIN[rand.Intn(len(a.GOBLIN))],
			Show:  true,
		}
	case enemiesList.SKELETON:
		image = &DisplayImage{
			Image: a.SKELETON[rand.Intn(len(a.SKELETON))],
			Show:  true,
		}
	case enemiesList.SORCERER:
		image = &DisplayImage{
			Image: a.SORCERER[rand.Intn(len(a.SORCERER))],
			Show:  true,
		}
	case enemiesList.ORC:
		image = &DisplayImage{
			Image: a.ORC[rand.Intn(len(a.ORC))],
			Show:  true,
		}
	case enemiesList.DRAGON:
		image = &DisplayImage{
			Image: a.DRAGON[rand.Intn(len(a.DRAGON))],
			Show:  true,
		}
	default:
		image = &DisplayImage{
			Image: a.HERO[rand.Intn(len(a.HERO))],
			Show:  true,
		}
	}
	return image
}

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

const heroAscii string = DoubleTab + "........................................\n" + DoubleTab +
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
