package main

var Lang string = getSysLang()

func translate(text map[string]string) string {
	return text[Lang]
}

var gameintroTR = map[string]string{
	englishLang: "\tWelcome to Dark Adventures\n\tSelect a difficulty:",
	frenchLang:  "\tBienvenu dans Dark Adventures\n\tChoisissez la difficulté:",
}

var NearORCTR = map[string]string{
	englishLang: "There's an odd smell in this place...\n",
	frenchLang:  "Il y a une odeur bizarre ici...\n",
}

var nearOrcSearchTR = map[string]string{
	englishLang: "odd smell",
	frenchLang:  "odeur bizarre",
}

var rootBellTRUE = map[string]string{
	englishLang: Tab + "A big deep-rooted tree stump is in front of you\n" +
		Tab + "You hear a bell sound, something is glowing from your bag..\n",
	frenchLang: Tab + "Une grande souche aux racines profondes devant vous\n" +
		Tab + "Un son de cloche, quelque chose brille dans votre sac..\n",
}

var rootBellFALSE = map[string]string{
	englishLang: Tab + "A big deep-rooted tree stump is in front of you\n",
	frenchLang:  Tab + "Une grande souche aux racines profondes devant vous\n",
}

var warpTextTR = map[string]string{
	englishLang: Tab + "All the roots from around gather and climb your legs,\n" +
		Tab + "a pain bites you and you faint as a white light engulfs everything\n" +
		Tab + "\"wake up!\" a voice says in your mind... you are somewhere else~",
	frenchLang: Tab + "Des racines viennent de partout et enlacent vos jambes,\n" +
		Tab + "soudain, une douleur pinçante et une forte lumière vous enveloppe\n" +
		Tab + "\"réveillez-vous!\" vous dit une voix fantomatique... vous êtes ailleur~",
}

var noComandsReceivedTR = map[string]string{
	englishLang: Tab + "No command received.",
	frenchLang:  Tab + "Pas de commande reçu.",
}

var cantDoThatTR = map[string]string{
	englishLang: Tab + "You can't do that here...",
	frenchLang:  Tab + "Vous ne pouvez pas faire ça...",
}

var cantBuyTR = map[string]string{
	englishLang: Tab + "Can't buy ",
	frenchLang:  Tab + "Vous ne pouvez pas acheter de ",
}

var youCanGoTR = map[string]string{
	englishLang: "You can Go:",
	frenchLang:  "Vous pouvez aller:",
}

var CanTGoTR = map[string]string{
	englishLang: "\tCan't go",
	frenchLang:  "\tVous ne pouvez pas aller",
}

var fromHereTR = map[string]string{
	englishLang: " from here.",
	frenchLang:  " d'ici.",
}

var whatDoYouMeanTR = map[string]string{
	englishLang: "What do you mean? ",
	frenchLang:  "Que voulez-vous dire? ",
}

var youCanTR = map[string]string{
	englishLang: "You can: ",
	frenchLang:  "Vous pouvez: ",
}

var youCanUseTR = map[string]string{
	englishLang: "You can use:",
	frenchLang:  "Vous pouvez utiliser:",
}

var youCanTUseATR = map[string]string{
	englishLang: "You can't use a ",
	frenchLang:  "Vous ne pouvez pas utiliser de ",
}

var orTR = map[string]string{
	englishLang: "...or: ",
	frenchLang:  "...ou: ",
}

var hasBeenSlainTR = map[string]string{
	englishLang: " has been slain",
	frenchLang:  " a été vaincu",
}

var youDiedTR = map[string]string{
	englishLang: "\tYou died\n",
	frenchLang:  "\tVous êtes mort\n",
}

var chooseHeroTR = map[string]string{
	englishLang: DoubleTab + "Select your hero:",
	frenchLang:  DoubleTab + "Choisissez votre hero:",
}

var heroesDetailsTHIEVE = map[string]string{
	englishLang: " - High evasion, crits, can attack x2",
	frenchLang:  " - Evasion, crits élevée, chances de double attaque",
}

var heroesDetailsPALADIN = map[string]string{
	englishLang: " - High strength, health, more potions",
	frenchLang:  " - Force, shanté élevée, plus de potions",
}

var heroesDetailsWIZARD = map[string]string{
	englishLang: " - Potions and Scrolls from the start",
	frenchLang:  " - Potions et Scrolls dès le départ",
}

var heroesDetailsBARBARIAN = map[string]string{
	englishLang: " - High health, strength",
	frenchLang:  " - Santé, force élevée",
}

var easyTR = map[string]string{
	englishLang: DoubleTab + "1 - Easy",
	frenchLang:  DoubleTab + "1 - Facile",
}

var difficultyTR = map[string]string{
	englishLang: "\n" + Tab + "Difficulty",
	frenchLang:  "\n" + Tab + "Difficulté",
}

var meddiumTR = map[string]string{
	englishLang: DoubleTab + "2 - Meddium",
	frenchLang:  DoubleTab + "2 - Moyen",
}

var hardTR = map[string]string{
	englishLang: DoubleTab + "3 - Hard",
	frenchLang:  DoubleTab + "3 - Difficile",
}

var ThieveNAME = map[string]string{
	englishLang: "Thieve",
	frenchLang:  "Voleur",
}

var PaladinNAME = map[string]string{
	englishLang: "Paladin",
	frenchLang:  "Paladin",
}

var WizardNAME = map[string]string{
	englishLang: "Wizard",
	frenchLang:  "Mage",
}

var BarbarianNAME = map[string]string{
	englishLang: "Barbarian",
	frenchLang:  "Barbare",
}

var skeletonNAME = map[string]string{
	englishLang: "SKELETON",
	frenchLang:  "SQUELETTE",
}

var goblinNAME = map[string]string{
	englishLang: "GOBLIN",
	frenchLang:  "GOBLIN",
}

var sorcererNAME = map[string]string{
	englishLang: "SORCERER",
	frenchLang:  "SORCIER",
}

var orcNAME = map[string]string{
	englishLang: "ORC",
	frenchLang:  "ORQUE",
}

var dragonNAME = map[string]string{
	englishLang: "DRAGON",
	frenchLang:  "DRAGON",
}

var forestTR = map[string]string{
	englishLang: "Trees are burned the soil is ash...",
	frenchLang:  "Les arbres sont brûlés le sol est couvert de cendres...",
}

var landTR = map[string]string{
	englishLang: "The air carries ashes flying in the wind...",
	frenchLang:  "L'air est chargé de cendres porté par le vent...",
}

var desertTR = map[string]string{
	englishLang: "It's hotter than usual and so dry...",
	frenchLang:  "Il fait très sec et trop chaud pour ce lieu...",
}

var castleTR = map[string]string{
	englishLang: "It smells like burning from all directions...",
	frenchLang:  "Ca sent le brûlé de toutes les directions...",
}

var xTR = map[string]string{
	englishLang: "NO LUCK, A strong wind bursts all around the place,\n" +
		Tab + "The sunlight dims before you hear the loudest of noises\n" +
		Tab + "Humongous, wings deployed its scream tearing the sky,\n" +
		Tab + "Here it is. The mightiest of all foes...",
	frenchLang: "MALHEUR, Un vent souffle très fortement,\n" +
		Tab + "On croirait une éclipse lorsqu'un bruit assourdissant retenti\n" +
		Tab + "Colossal, ailes déployées un cri déchirant les cieux,\n" +
		Tab + "Le voilà. Le plus redoutable des ennemis...",
}

var youAreTR = map[string]string{
	englishLang: Tab + "You are",
	frenchLang:  Tab + "Vous êtes",
}

var introPlainsTR0 = map[string]string{
	englishLang: " in an old foggy village, there's no soul here,\n",
	frenchLang:  " dans un ancient village brumeux, il n'y a pas âme qui vive,\n",
}

var introPlainsTR1 = map[string]string{
	englishLang: " in the heath, you hear a weird music,\n" + Tab +
		"let's not waste any time here,\n",
	frenchLang: " dans les plaines, on entend une musique étrange,\n" + Tab +
		"ne perdons pas de temps,\n",
}

var introPlainsTR2 = map[string]string{
	englishLang: " on a long road between green hills and a river,\n",
	frenchLang:  " sur une longue route qui sépare plaines vertes et rivière,\n",
}

var introDesertTR0 = map[string]string{
	englishLang: " in the wasteland, everything is dead and dry here,\n",
	frenchLang:  " dans les terres désolées, tout ici est mort et sec,\n",
}

var introDesertTR1 = map[string]string{
	englishLang: " on the swamp, nauseous and poisonous,\n" + Tab +
		"something is lurking here,\n",
	frenchLang: " dans le marais, nauséabon et vénéneux,\n" + Tab +
		"il y a quelque chose qui rôde par ici,\n",
}

var introDesertTR2 = map[string]string{
	englishLang: " in the middle of dust..\n" + Tab +
		"of a long gone empire and a storm is at the horizon,\n",
	frenchLang: " dans un tas de poussière..\n" + Tab +
		"celui des restes d'un empire disparu, une tempête à l'horizon,\n",
}

var introCastleTR0 = map[string]string{
	englishLang: " in front of a castle ruin's gate, it barely stands,\n",
	frenchLang:  " devant les portes d'un chateau en ruines, prêt à s'écrouler,\n",
}

var introCastleTR1 = map[string]string{
	englishLang: " at an old fort or what might have been one long ago,\n",
	frenchLang:  " prêt d'un vieu fort ou ce qui devait en être un il y a longtemps,\n",
}

var introCastleTR2 = map[string]string{
	englishLang: " below a huge tower, on top of which float an old flag,\n",
	frenchLang:  " en bas d'une tour très énorme, sur laquelle flotte un vieu drapeau,\n",
}

var introForestTR0 = map[string]string{
	englishLang: " near a forest, the trees seem to move by their own will,\n",
	frenchLang:  " à la lisière d'une forêt, les arbres semblent mû de volonté,\n",
}

var introForestTR1 = map[string]string{
	englishLang: " unfortunately at the edge of the thorns wood,\n" + Tab +
		"no one comes back from it,\n",
	frenchLang: " malheureusement a la frontière de la forêt d'épines,\n" + Tab +
		"personne n'en ait jamais revenu,\n",
}

var introForestTR2 = map[string]string{
	englishLang: " in a part of the forest all trees are rotten\n" + Tab +
		"and covered by poisonous mushrooms,\n",
	frenchLang: " dans une part de la forêt ou tout est moisi\n" + Tab +
		"et recouvert de champignons vénéneux,\n",
}

var AmbianceTR0 = map[string]string{
	englishLang: "there are nobody around, only the wind.\n",
	frenchLang:  "il n'y a personne ici, seulement le vent.\n",
}

var AmbianceTR1 = map[string]string{
	englishLang: "it's getting dark and you can see shadows moving..\n",
	frenchLang:  "il commence à faire sombre et vous appercevez des ombres bouger..\n",
}

var AmbianceTR2 = map[string]string{
	englishLang: "all is silent, there's not even wind!\n",
	frenchLang:  "tout est silencieux, pas même le son du vent!\n",
}

var AmbianceTR3 = map[string]string{
	englishLang: "you don't feel safe but have to keep going on..\n",
	frenchLang:  "vous ne vous sentez pas en sécurité mais il faut continuer..\n",
}

var AmbianceTR4 = map[string]string{
	englishLang: "suddenly you feel shivers, a noise, voice or wind?\n",
	frenchLang:  "vous avez la chair de poule, un bruit, le vent ou des voix?\n",
}

var AmbianceTR5 = map[string]string{
	englishLang: "many noises around you, but can't see anyone...\n",
	frenchLang:  "beaucoup de bruits autour, mais vous ne voyez personne...\n",
}

var SellerListTR0 = map[string]string{
	englishLang: " a dwarf, with a bag full of goods\n",
	frenchLang:  " un nain, avec une besace pleine de choses\n",
}

var SellerListTR1 = map[string]string{
	englishLang: " an elf, he holds something in his hand\n",
	frenchLang:  " un elf, il tend la main montrant quelque chose\n",
}

var SellerListTR2 = map[string]string{
	englishLang: " a troll, he drops something in front of you\n",
	frenchLang:  " un troll, il laisse tomber un objet devant vous\n",
}

var DollTR = map[string]string{
	englishLang: "Will revive you with 30 HP",
	frenchLang:  "Vous ramène à la vie avec 30 HP",
}

var MoonstoneTR = map[string]string{
	englishLang: "Increase your strength by 5",
	frenchLang:  "Augmente votre force de 5",
}

var ScrollTR = map[string]string{
	englishLang: "20 Dammage to one enemy",
	frenchLang:  "20 points de dégats",
}

var PotionTR = map[string]string{
	englishLang: "Heal 20 HP",
	frenchLang:  "Soigne 20 HP",
}

var KeyTR = map[string]string{
	englishLang: "To open locks, chests",
	frenchLang:  "Ouvre des mechanismes et coffres",
}

var CoinsTR = map[string]string{
	englishLang: "Golden coins",
	frenchLang:  "Pièces d'or",
}

var potionNameTR = map[string]string{
	englishLang: "potion",
	frenchLang:  "potion",
}

var scrollNameTR = map[string]string{
	englishLang: "scroll",
	frenchLang:  "parchemin",
}

var dollNameTR = map[string]string{
	englishLang: "doll",
	frenchLang:  "poupée",
}

var keyNameTR = map[string]string{
	englishLang: "key",
	frenchLang:  "clé",
}

var moonstoneNameTR = map[string]string{
	englishLang: "moonstone",
	frenchLang:  "pierre de lune",
}

var coinsNameTR = map[string]string{
	englishLang: "coins",
	frenchLang:  "pièces",
}

var chestEventNameTR = map[string]string{
	englishLang: "chest",
	frenchLang:  "coffre",
}

var enemyEventNameTR = map[string]string{
	englishLang: "enemy",
	frenchLang:  "ennemi",
}

var sellerEventNameTR = map[string]string{
	englishLang: "seller",
	frenchLang:  "marchand",
}

var HasEnemyOrSellerTR0 = map[string]string{
	englishLang: Tab + "There is ",
	frenchLang:  Tab + "Il y a un ",
}

var HasEnemyTR1 = map[string]string{
	englishLang: "ready to fight you!\n",
	frenchLang:  " prêt à combattre!\n",
}

var HasSellerTR = map[string]string{
	englishLang: Tab + "He's proposing:\n",
	frenchLang:  Tab + "Il vous propose:\n",
}

var forTR = map[string]string{
	englishLang: " for ",
	frenchLang:  " pour ",
}

var forCoinsTR = map[string]string{
	englishLang: " coins.\n",
	frenchLang:  " pièces.\n",
}

var yourInventoryTR = map[string]string{
	englishLang: DoubleTab + "Your inventory:",
	frenchLang:  DoubleTab + "Votre inventaire:",
}

var inventoryQuantityTR = map[string]string{
	englishLang: "Quantity: ",
	frenchLang:  "Quantité: ",
}

var getEnemyItemsTR = map[string]string{
	englishLang: DoubleTab + "You get:",
	frenchLang:  DoubleTab + "Vous récupérez:",
}

var nothingYouCouldGetTR = map[string]string{
	englishLang: Tab + "Enemy had nothing you could get...",
	frenchLang:  Tab + "Il n'y a rien à récupérer...",
}

var missedTR = map[string]string{
	englishLang: " MISSED!!",
	frenchLang:  " A RATE!!",
}

var youBaughtTR = map[string]string{
	englishLang: Tab + "You baught ",
	frenchLang:  Tab + "Vous avez acheté ",
}

var doesTR = map[string]string{
	englishLang: " does ",
	frenchLang:  " fait ",
}

var critDMGTR = map[string]string{
	englishLang: " Critical DMG to ",
	frenchLang:  " DEGATS CRITIQUES à ",
}

var dmgToTR = map[string]string{
	englishLang: " DMG to ",
	frenchLang:  " dégats à ",
}

var dollUsedTR = map[string]string{
	englishLang: Tab + "A dark force is devouring your body\n" +
		Tab + "A chance has been given to you or is it?\n" +
		Tab + "You died... and revived.\n" +
		Tab + "Health +30 HP",
	frenchLang: Tab + "Une force obscure enveloppe votre corp\n" +
		Tab + "Une chance vous a été offerte ou en est-ce vraiment une?\n" +
		Tab + "Vous êtes mort... et avez réssucité.\n" +
		Tab + "Santé +30 HP",
}

var moonstoneUsedTR = map[string]string{
	englishLang: "The moonstone suddenly wraps and fuse in your arms, it's burning!\n",
	frenchLang:  "La moonstone encercle et fusionne à votre bras, c'est brûlant!\n",
}

var strengthBoostAddTR = map[string]string{
	englishLang: "Strength +5 ->",
	frenchLang:  "FORCE +5 ->",
}

var youDontHaveATR = map[string]string{
	englishLang: "You don't have a ",
	frenchLang:  "Vous n'avez pas de ",
}

var youTR = map[string]string{
	englishLang: "You:",
	frenchLang:  "Vous:",
}

var shopsTR = map[string]string{
	englishLang: "Shops:",
	frenchLang:  "Boutique:",
}

var rootsTR = map[string]string{
	englishLang: "Roots:",
	frenchLang:  "Racines:",
}

var northTR = map[string]string{
	englishLang: "North",
	frenchLang:  "Nord",
}

var eastTR = map[string]string{
	englishLang: "East",
	frenchLang:  "Est",
}

var westTR = map[string]string{
	englishLang: "West",
	frenchLang:  "Ouest",
}

var southTR = map[string]string{
	englishLang: "South",
	frenchLang:  "Sud",
}

var directionsArticlesVowelsTR = map[string]string{
	englishLang: " ",
	frenchLang:  " à l'",
}

var directionsArticlesConsonantTR = map[string]string{
	englishLang: " ",
	frenchLang:  " au ",
}

func vowelOrNot(word string, with, without map[string]string) string {
	if hasVowel := VowelNextNun(word); hasVowel {
		return translate(with)
	}
	return translate(without)
}

func VowelNextNun(str string) bool {
	vowels := []string{"a", "e", "i", "o", "y", "u"}
	if InitialsIndexOf(vowels, str) {
		return true
	}
	return false
}

// var NAME = map[string]string{
// 	englishLang: ,
// 	frenchLang: ,
// }

// var NAME = map[string]string{
// 	englishLang: ,
// 	frenchLang: ,
// }

// var NAME = map[string]string{
// 	englishLang: ,
// 	frenchLang: ,
// }
