package main

var Lang string = getSysLang()

func translate(text map[string]string) string {
	if _, ok := text[Lang]; ok {
		return text[Lang]
	}
	return text[englishLang]
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

var youSeeTR = map[string]string{
	englishLang: "\tYou see ",
	frenchLang:  "\tVous voyez ",
}

var deadOnTheGroundTR = map[string]string{
	englishLang: "dead on the ground\n",
	frenchLang:  "mort sur le sol\n",
}

var someoneWasHereTR = map[string]string{
	englishLang: Tab + "There was someone here, it's empty now.",
	frenchLang:  Tab + "Il y avait quelqu'un ici, il n'est plus là.",
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
	englishLang: "Trees are burned the soil is ash...\n",
	frenchLang:  "Les arbres sont brûlés le sol est couvert de cendres...\n",
}

var landTR = map[string]string{
	englishLang: "The air carries ashes flying in the wind...\n",
	frenchLang:  "L'air est chargé de cendres porté par le vent...\n",
}

var desertTR = map[string]string{
	englishLang: "It's hotter than usual and so dry...\n",
	frenchLang:  "Il fait très sec et trop chaud pour ce lieu...\n",
}

var castleTR = map[string]string{
	englishLang: "It smells like burning from all directions...\n",
	frenchLang:  "Ca sent le brûlé de toutes les directions...\n",
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
	frenchLang:  " prêt d'un vieux fort ou ce qui en a été un il y a longtemps,\n",
}

var introCastleTR2 = map[string]string{
	englishLang: " below a huge tower, on top of which float an old flag,\n",
	frenchLang:  " en bas d'une tour énorme, avec un vieux drapeau,\n",
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
	frenchLang:  "Vous ramène avec 30 HP",
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
	frenchLang:  "poupee",
}

var keyNameTR = map[string]string{
	englishLang: "key",
	frenchLang:  "cle",
}

var moonstoneNameTR = map[string]string{
	englishLang: "moonstone",
	frenchLang:  "moonstone",
}

var coinsNameTR = map[string]string{
	englishLang: "coins",
	frenchLang:  "pieces",
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
	frenchLang:  "La pierre de lune encercle et fusionne à votre bras, c'est brûlant!\n",
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

var heroStoryThieveTR0 = map[string]string{
	englishLang: Tab + "Member of the Thieves Guild of Novigrad, you've been hired\n" +
		Tab + "by a mysterious contractor to retrieve a so-called treasure here...",
	frenchLang: Tab + "Membre de la Guilde des voleurs de Novigrad, on vous a engagé\n" +
		Tab + "par un homme mystérieux pour retrouver un soi-disant trésor ici...",
}

var heroStoryThieveTR1 = map[string]string{
	englishLang: Tab + "Behind every Pirates story, there's a treasure map...\n" +
		Tab + "The same goes for you, unfortunately...",
	frenchLang: Tab + "Pour toute histoire de Pirates, il y a une carte au trésor...\n" +
		Tab + "Il en va de même pour la votre, malheureusement...",
}

var heroStoryThieveTR2 = map[string]string{
	englishLang: Tab + "A Rogue from the Hanging City of Szyan who left, with the final goal\n" +
		Tab + "the Thieves Guild of Novigrad. Until your path lead you here...",
	frenchLang: Tab + "Voyou ayant quitté votre cité suspendue Szyan, avec pour destination\n" +
		Tab + "la Guilde des voleurs de Novigrad. votre chemin vous a mené ici...",
}

var heroStoryThieveTR3 = map[string]string{
	englishLang: Tab + "As a memeber of the Leaf Walkers Order of female elfic guards,\n" +
		Tab + "you go wherever the forest needs help and heard the trees calling.",
	frenchLang: Tab + "En tant que membre de l'ordre elfique des Marcheuses de feuilles,\n" +
		Tab + "vous allez partout où la forêt à besoin et les arbres ont appelé.",
}

var heroStoryPaladinTR0 = map[string]string{
	englishLang: Tab + "Amazon warrior from the Ionos Archipel, you have seen the Oracle,\n" +
		Tab + "the constellations are clear, Darkness are rising. You must act.",
	frenchLang: Tab + "Guerrière amazonne de l'Archipel Ionos, l'Oracle à parlé,\n" +
		Tab + "les étoiles sont clair, Les ténèbres s’élèvent. Vous devez agir.",
}

var heroStoryPaladinTR1 = map[string]string{
	englishLang: Tab + "Templar of the Crimson Star Fellowship, an oath a creed to follow,\n" +
		Tab + "you will fight the darkness wherever they are, they sent you here.",
	frenchLang: Tab + "Templier de la confrérie de l'étoile écarlate, un serment vous lie,\n" +
		Tab + "vous combattrez les ténèbres où qu'elles soient, ils vous ont envoyé là.",
}

var heroStoryPaladinTR2 = map[string]string{
	englishLang: Tab + "Retired veteran of the Risen Wars, the increasing rumors of evil\n" +
		Tab + "reach you and pulled you out of retirement. Time to draw the sword.",
	frenchLang: Tab + "Vétéran des guerres des relevés, les rumeurs d'un mal croissant\n" +
		Tab + "vous ont atteintes et quitté votre retraite. Le moment de l'épée.",
}

var heroStoryPaladinTR3 = map[string]string{
	englishLang: Tab + "Knight of the 3 Towers Citadel your duty: stand as a wall,\n" +
		Tab + "the rampart to protect the city of Eelring, a menace is near...",
	frenchLang: Tab + "Chevalier de la Citadelle des 3 tours votre devoir: faire mur,\n" +
		Tab + "le rempart protégeant la cité d'Eelring, une menace approche...",
}

var heroStoryWizardTR0 = map[string]string{
	englishLang: Tab + "The Elder scholar from Krispin Academy of Magic himself,\n" +
		Tab + "you came for the rumors of a rising evil in these forsaken lands.",
	frenchLang: Tab + "L’érudit aîné de l'Academy de Magie Krispin lui-même,\n" +
		Tab + "venu suite aux rumeurs d'un mal croissant dans ces terres abandonnées.",
}

var heroStoryWizardTR1 = map[string]string{
	englishLang: Tab + "A Tiefling Wizard girl, speeking the lower planes language,\n" +
		Tab + "that give you immense powers. You're seeking something here...",
	frenchLang: Tab + "Magicienne Tiefling, parlant la langue des plans inférieurs,\n" +
		Tab + "cela vous donne d'immenses pouvoirs. Vous êtes venue chercher quelque chose...",
}

var heroStoryWizardTR2 = map[string]string{
	englishLang: Tab + "A beautiful but extremely powerful wizard girl in a white clothing,\n" +
		Tab + "everyone has a strong reason to come there, yours? Vengence.",
	frenchLang: Tab + "Très belle mais extrêmement puissante magicienne en robe blanche,\n" +
		Tab + "chacun a une bonne raison de venir, la votre? La vengence.",
}

var heroStoryWizardTR3 = map[string]string{
	englishLang: Tab + "An old war veteran, who served under Tzar Krvovoj after leaving\n" +
		Tab + "Krispin Academy of Magic with the highest honors. Why coming here?",
	frenchLang: Tab + "Vétéran de l'ancienne guerre, ayant servi sous le Tzar Krvovoj après avoir\n" +
		Tab + "quitté l'Academie de magie Krispin avec les honneurs. Que faites vous là?",
}

var heroStoryBarbarianTR0 = map[string]string{
	englishLang: Tab + "A warrior women from the mount Zanarkan you seek treasures,\n" +
		Tab + "at the Golden Horse Tavern you heard some rumors... Richnesses!",
	frenchLang: Tab + "Guerrière du Mont Zanarkan à la recherche de trésors,\n" +
		Tab + "vous avez entendu une rumeur à la Tavern du cheval doré... Richesses!",
}

var heroStoryBarbarianTR1 = map[string]string{
	englishLang: Tab + "A proud Warrior from the Dark Moon Gnoll clan, Zorik the Shaman\n" +
		Tab + "asked you to search the last moon's hunters that never came back.",
	frenchLang: Tab + "Fière guerrier Gnoll du clan de la Lune Noire, Zorik le Shaman\n" +
		Tab + "vous envois rechercher les chasseurs disparus de la dernière lune.",
}

var heroStoryBarbarianTR2 = map[string]string{
	englishLang: Tab + "Fierce hunter of the Ionos Archipel, you seek glory...\n" +
		Tab + "The ultimate prey is the one you fell prey to, and vanquisehd.",
	frenchLang: Tab + "Féroce chasseur de l'Archipel Ionos, vous cherchez la gloire...\n" +
		Tab + "La proie ultime est celle dont vous étiez la proie avant la victoire.",
}

var heroStoryBarbarianTR3 = map[string]string{
	englishLang: Tab + "After 20 years fighting in the Coliseum of Styr, you are free.\n" +
		Tab + "Looking for a place to settle, you arrived a weird foreign place..",
	frenchLang: Tab + "Après 20 ans dans le Colisée de Styr, vous êtes libre.\n" +
		Tab + "Cherchant un lieu pour s’installer, vous tombé sur ces terres étranges..",
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

// var NAME = map[string]string{
// 	englishLang: ,
// 	frenchLang: ,
// }

// var NAME = map[string]string{
// 	englishLang: ,
// 	frenchLang: ,
// }
