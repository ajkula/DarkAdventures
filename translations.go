package main

var Lang string = getSysLang()

func translate(text map[string]string) string {
	if _, ok := text[Lang]; ok {
		return text[Lang]
	}
	return text[englishLang]
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

var gameintroTR = map[string]string{
	englishLang: Tab + "Welcome to Dark Adventures\n" + Tab + "Select a difficulty:",
	frenchLang:  Tab + "Bienvenue dans Dark Adventures\n" + Tab + "Choisissez la difficulté:",
}

var forestNameTR = map[string]string{
	englishLang: "forest",
	frenchLang:  "forêt",
}

var plainsNameTR = map[string]string{
	englishLang: "plains",
	frenchLang:  "plaines",
}

var desertNameTR = map[string]string{
	englishLang: "desert",
	frenchLang:  "désert",
}

var castleNameTR = map[string]string{
	englishLang: "castle",
	frenchLang:  "ruines",
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
	englishLang: Tab + "You see ",
	frenchLang:  Tab + "Vous voyez ",
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
		Tab + "\"wake up!\" an eerie voice says in your mind.. you are somewhere else~",
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
	englishLang: Tab + "Can't go",
	frenchLang:  Tab + "Vous ne pouvez pas aller",
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
	englishLang: Tab + "You died\n",
	frenchLang:  Tab + "Vous êtes mort\n",
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
	frenchLang:  " - Force, santé élevée, plus de potions",
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

var ThieveEscapeOK = map[string]string{
	englishLang: Tab + "Using tricks and tools, you make a dark suffocating smoke cloud\n" +
		Tab + "Giving you an opportunity to fall back",
	frenchLang: Tab + "Utilisant vos gadgets et tours, vous créez un nuage noir\n" +
		Tab + "de fumée suffocante vous donnant l'opportunité de battre en retraite",
}

var ThieveEscapeRAND = map[string]string{
	englishLang: Tab + "Parrying the enemy's attacks and throw a handful of darts\n" +
		Tab + "they fly straight to their face as you jump wherever you can",
	frenchLang: Tab + "Parrant les attaques ennemies vous jetez une poignée de flechettes\n" +
		Tab + "qui volent droit à leur visage tandis que vous fuiyez où ce peut",
}

var ThieveEscapeKO = map[string]string{
	englishLang: Tab + "After a clumsy parry you trigger agrenade off your arm-canon\n" +
		Tab + "it blocks and your step trips giving the enemy an opportunity to hit",
	frenchLang: Tab + "Après avoir maladroitement parré vous tantez de tirer une grenade\n" +
		Tab + "votre bras-canon s'enraille et vous trébuchez à la mercie de l'ennemi",
}

var PaladinEscapeOK = map[string]string{
	englishLang: Tab + "The god Triglav shows a happy face, you're in their grace\n" +
		Tab + "a huge pillar of light chases darkness around, you fallback safely",
	frenchLang: Tab + "Le dieu Triglav montre un visage souriant, vous êtes en grâce\n" +
		Tab + "un énorme pilier de lumière chasse l'obscurité, battez en retraite",
}

var PaladinEscapeRAND = map[string]string{
	englishLang: Tab + "The god Triglav shows a sad face, parrying from all sides\n" +
		Tab + "you roll dodge a heavy one and run to escape where you can...",
	frenchLang: Tab + "Le dieu Triglav montre un visage triste, parant de tout côtés\n" +
		Tab + "vous roulez évitant un mauvais coup et courrez où vous pouvez...",
}

var PaladinEscapeKO = map[string]string{
	englishLang: Tab + "The god Triglav shows an angry face, you try to run away\n" +
		Tab + "your first step slides on a wet stone bad timing, but why?",
	frenchLang: Tab + "Le dieu Triglav montre un visage fâché, vous tantez de fuir\n" +
		Tab + "glissez sur une pierre humide, mauvais timing, pourquoi?",
}

var WizardEscapeOK = map[string]string{
	englishLang: Tab + "By a slight move the spells in your hand switches\n" +
		Tab + "a magic shield appears in front of you, time to retreat",
	frenchLang: Tab + "D'un geste légé vous changez de sorts en mains\n" +
		Tab + "un bouclier magique apparaît devant vous, il est temp de fuir",
}

var WizardEscapeRAND = map[string]string{
	englishLang: Tab + "No time for sofisticated spell you were off guaard\n" +
		Tab + "throwing a teleport vial rapidly on the ground you vanish",
	frenchLang: Tab + "Pas le temp d'un sort complexe on vous a surpris\n" +
		Tab + "jetant vite une fiole de teleportation vous disparaîssez",
}

var WizardEscapeKO = map[string]string{
	englishLang: Tab + "you start to cast a defense spell but suddenly...\n" +
		Tab + "caughing makes it fail, too late.. the enemy took the chance",
	frenchLang: Tab + "Vous démarrer un sort défensif mais soudain...\n" +
		Tab + "une quinte de toux, échec... trop tard l'ennemi saisi l'ouverture",
}

var BarbarianEscapeOK = map[string]string{
	englishLang: Tab + "Dazbog, Wolf God of the Hunt is here, Ô luck!\n" +
		Tab + "Raising from your guts his scream knocks enemies, run!",
	frenchLang: Tab + "Dazbog, le Dieu Loup de la chasse est là, Ô grace!\n" +
		Tab + "Venant de vos trippes son cri étourdi les ennemis, courez!",
}

var BarbarianEscapeRAND = map[string]string{
	englishLang: Tab + "The fight is unfair, time to escape, an oppening!\n" +
		Tab + "You charge the enemy up front and end up somewhere else...",
	frenchLang: Tab + "Le combat est inégal, il faut fuir, une ouverture!\n" +
		Tab + "Vous chargez l'ennemi plein front et atterissez ailleurs...",
}

var BarbarianEscapeKO = map[string]string{
	englishLang: Tab + "Dazbog, Wolf God of the hunt is not in this place.\n" +
		Tab + "You try to run but no luck you're cornere...d",
	frenchLang: Tab + "Dazbog, le Dieu Loup de la chasse n'est pas ici.\n" +
		Tab + "Vous courez mais pas de chance vous êtes acculé...",
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

var necromancerNAME = map[string]string{
	englishLang: "NECROMANCER",
	frenchLang:  "NECROMANCIEN",
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
	englishLang: "Heal",
	frenchLang:  "Soigne",
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

var sellerTutoTR = map[string]string{
	englishLang: "\n" + Tab + "You can buy one of his items like so:\n" +
		DoubleTab + "b(uy) ",
	frenchLang: "\n" + Tab + "Vous pouvez acheter de cette manière:\n" +
		DoubleTab + "b(uy) ",
}

var HasEnemyOrSellerTR0 = map[string]string{
	englishLang: Tab + "There is ",
	frenchLang:  Tab + "Il y a ",
}

var HasEnemyTR1 = map[string]string{
	englishLang: "ready to fight you!\n",
	frenchLang:  "prêt à combattre!\n",
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

var heroStoryThieveTR0 = map[string]string{
	englishLang: Tab + "Member of the Thieves Guild of Novigrad, you've been hired\n" +
		Tab + "by a mysterious contractor to retrieve a so-called treasure here...",
	frenchLang: Tab + "Membre de la Guilde des voleurs de Novigrad, on vous a engagé\n" +
		Tab + "par un homme mystérieux pour retrouver un soi-disant trésor ici...",
}

var heroStoryThieveTR1 = map[string]string{
	englishLang: Tab + "Behind every Pirates story, there's a treasure map...\n" +
		Tab + "The same goes for yours, unfortunately...",
	frenchLang: Tab + "Pour toute histoire de Pirates, il y a une carte au trésor...\n" +
		Tab + "Vous en avez également trouvé une, malheureusement...",
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
		Tab + "reach you and pulled you out of retirement. Time to draw the sword!",
	frenchLang: Tab + "Vétéran des guerres des relevés, les rumeurs d'un mal croissant\n" +
		Tab + "vous ont atteintes et fait quitter votre retraite. Brandissez l'épée!",
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
		Tab + "vous donnant d'immenses pouvoirs. Qu'êtes-vous êtes venue chercher?",
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
	frenchLang: Tab + "Vétéran de l'ancienne guerre ayant servi le Tzar Krvovoj après avoir\n" +
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

var HealthUP = map[string]string{
	englishLang: "Max Health:",
	frenchLang:  "Santé Max:",
}

var CritsUP = map[string]string{
	englishLang: "Crit Chances:",
	frenchLang:  "Chances Crit:",
}

var EvasionUP = map[string]string{
	englishLang: "Evasion:",
	frenchLang:  "Esquive:",
}

var StrengthUP = map[string]string{
	englishLang: "Strength:",
	frenchLang:  "Force:",
}

var SkillsUP = map[string]string{
	englishLang: "Skills:",
	frenchLang:  "Compétences:",
}

var Health = map[string]string{
	englishLang: "Health",
	frenchLang:  "Santé",
}

var Crit = map[string]string{
	englishLang: "Crit Chances",
	frenchLang:  "Chances Crit",
}

var Evasion = map[string]string{
	englishLang: "Evasion",
	frenchLang:  "Esquive",
}

var Strength = map[string]string{
	englishLang: "Strength",
	frenchLang:  "Force",
}

var Skill = map[string]string{
	englishLang: "Special",
	frenchLang:  "Spécial",
}

var Exp = map[string]string{
	englishLang: "Experience",
	frenchLang:  "Experience",
}

var Level = map[string]string{
	englishLang: "Level",
	frenchLang:  "Niveau",
}

var Rooms = map[string]string{
	englishLang: "Rooms",
	frenchLang:  "Lieux",
}

var Enemies = map[string]string{
	englishLang: "Enemies killed",
	frenchLang:  "Ennemis tués",
}

var BoostTR = map[string]string{
	englishLang: "Moonstones UP",
	frenchLang:  "Gain de Force",
}

var ThiefSkill = map[string]string{
	englishLang: DoubleTab + "The tief, a skilled agent of deception,\n" +
		DoubleTab + "will steal one object of value the enemy carries.",
	frenchLang: DoubleTab + "Le voleur, un compétent agent de la ruse,\n" +
		DoubleTab + "volera un des objet précieux de l'ennemi.",
}

var PaladinSkill = map[string]string{
	englishLang: DoubleTab + "The Paladin Knight is a holy master of combat,\n" +
		DoubleTab + "undead dies instantly, others get HP halfed.",
	frenchLang: DoubleTab + "Le Paladin est un maître du combat sacré,\n" +
		DoubleTab + "les mort-vivants meurent, les autres ont -50% HP.",
}

var WizardSkill = map[string]string{
	englishLang: DoubleTab + "A powerful spell does 10 area dmg near you,\n" +
		DoubleTab + "the enemy you fight gets a fire status for 3 turns.",
	frenchLang: DoubleTab + "Un sort puissant cause 10 dmg de zone,\n" +
		DoubleTab + "l'ennemi prend le statut enflammé pour 3 tours.",
}

var BarbarianSkill = map[string]string{
	englishLang: DoubleTab + "A skilled hunter, the map is revealed around,\n" +
		DoubleTab + "the enemy you fight is terrified for 2 turns.",
	frenchLang: DoubleTab + "Un chasseur hors pairs, carte révélée alentour,\n" +
		DoubleTab + "l'ennemi est terrifié pour 2 tours.",
}

var StatusTR = map[string]string{
	englishLang: "Status",
	frenchLang:  "Statut",
}

var LevelUPTR = map[string]string{
	englishLang: "LEVEL UP!",
	frenchLang:  "Niveau Supérieur!",
}

var NoSkillTR = map[string]string{
	englishLang: "You seem too tired to do that...",
	frenchLang:  "Vous avez l'air trop épuisé pour faire ça...",
}

var StealFailTR = map[string]string{
	englishLang: " couldn't steal anything...",
	frenchLang:  " n'a rien pu voler...",
}

var TheTR = map[string]string{
	englishLang: "The ",
	frenchLang:  "Le ",
}

var StealSuccessTR = map[string]string{
	englishLang: " has successfully stolen ",
	frenchLang:  " a réussi a voler ",
}

var HolySuccessTR = map[string]string{
	englishLang: Tab + "The god Triglav shows a smiling face, your sword glows\n" +
		Tab + "you raise it above your head illuminating everything,\n" +
		DoubleTab + translate(UndeadTR) + " creatures die instantly!",
	frenchLang: Tab + "Le dieu Triglav montre un visage souriant, votre épée brille\n" +
		Tab + "la levant au-dessus de vous, elle illumine tout,\n" +
		DoubleTab + "les " + translate(UndeadTR) + "s meurent instantanément!",
}

var HolyMitigatedTR = map[string]string{
	englishLang: Tab + "The god Triglav shows a sad face, your sword resonates\n" +
		Tab + "you raise it above you the sound crushes all but you,\n" +
		DoubleTab + "enemy HP and HP MAX Halfed...",
	frenchLang: Tab + "Le dieu Triglav montre un visage triste, votre épée vibre\n" +
		Tab + "la levant haut le son qu'elle émet brise tout sauf vous,\n" +
		DoubleTab + "HP et HP MAX ennemi réduits de moitié...",
}

var HolyHugeTR = map[string]string{
	englishLang: Tab + "The god Triglav shows a sad face, your sword resonates\n" +
		Tab + "you raise it above you the sound crushes all but you,\n" +
		DoubleTab + "Enemy -",
	frenchLang: Tab + "Le dieu Triglav montre un visage triste, votre épée vibre\n" +
		Tab + "la levant haut le son qu'elle émet brise tout sauf vous,\n" +
		DoubleTab + "Ennemi -",
}

var MagisterSkillTR = map[string]string{
	englishLang: Tab + "The Neptunium of the Observatory Tower of Krispin\n," +
		Tab + "turns to align and focus the light of the Sirius star.\n" +
		Tab + "The bean hits you wand's crystal dealing an area damage\n" +
		DoubleTab + "Enemies around get -> -10 HP\n" +
		Tab + "Using your spirit's strength you focus a beam forward\n",
	frenchLang: Tab + "Le Neptunium de l'Observatoire de la tour de Krispin\n," +
		Tab + "s'oriente pour refleter vers vous la lumière de l'étoile Sirius.\n" +
		Tab + "Le rayon frappe votre baton causant des dégats de zone\n" +
		DoubleTab + "Les ennemis autour reçoivent -> -10 HP\n" +
		Tab + "Avec la force de l'esprit vous concentrer le rayon\n",
}

var DazbogRushSkillTR = map[string]string{
	englishLang: Tab + "Dazbog, Wolf God of the Hunt's ethereal wolves runs all ways.\n" +
		Tab + "They suddenly jump dive back straight to your mind.\n" +
		Tab + "The map around is revealed to you, a huge ghostly\n" +
		Tab + "wolf jaw bites the enemy from the sky above...\n",
	frenchLang: Tab + "Dazbog, Dieu Loup de la Chasse envois ses loups éthérés.\n" +
		Tab + "Ils reviennent d'un bond dans votre esprit.\n" +
		Tab + "La carte alentour vous est révelée, une énorme\n" +
		Tab + "mâchoire spéctrale mors l'ennemi depuis les cieux...\n",
}

var EnemiGotFrightTR = map[string]string{
	englishLang: DoubleTab + "Enemy gets " + translate(frightStatusTR) + " for 2 turn",
	frenchLang:  DoubleTab + "L'ennemi est affecté par " + translate(frightStatusTR) + " pour 2 tour",
}

var heroGotFrightTR = map[string]string{
	englishLang: DoubleTab + "You got " + translate(blightStatusTR),
	frenchLang:  DoubleTab + "Vous êtes affecté par " + translate(frightStatusTR) + " pour 2 tour",
}

var EnemiGotBlightTR = map[string]string{
	englishLang: DoubleTab + "Enemy gets " + translate(blightStatusTR),
	frenchLang:  DoubleTab + "L'ennemi est affecté par " + translate(blightStatusTR),
}

var heroGotBlightTR = map[string]string{
	englishLang: DoubleTab + "You got " + translate(blightStatusTR),
	frenchLang:  DoubleTab + "Vous êtes affecté par " + translate(blightStatusTR),
}

var AreaHits = map[string]string{
	englishLang: " Enemies Hit!\n",
	frenchLang:  " Ennemis touchés!\n",
}

var DarkEnergyTR = map[string]string{
	englishLang: Tab + "A huge Dark Energy ball descends the sky above you\n" +
		Tab + "Your life is Reduced by ",
	frenchLang: Tab + "Une énorme boule d'énergie noir vous descend dessus\n" +
		Tab + "Votre santé est Réduite de ",
}

var DarkEnergyOnPaladinTR = map[string]string{
	englishLang: "a Quarter!",
	frenchLang:  "un Quart!",
}

var DarkEnergyNormalTR = map[string]string{
	englishLang: "Half!",
	frenchLang:  "Moitié!",
}

var GraceProtectsYouTR = map[string]string{
	englishLang: Tab + "Grace protects you from Darkness!\n",
	frenchLang:  Tab + "La Grâce vous protège des Ténèbres!\n",
}

var DragonSkillFireTR = map[string]string{
	englishLang: Tab + "This is the reason a Dragon is terrible to encounter!\n" +
		Tab + "It suddenly spits fire towards you for a duration...\n" +
		Tab + "You barely find cover, Oh no! Are you still Alive?..\n" +
		DoubleTab + "Cost of the flames blow ",
	frenchLang: Tab + "Voici la raison pour laquelle rencontrer un dragon\n" +
		Tab + "est si terrible!\n" +
		Tab + "Tout d'un coup il crache un torrent de feu vers vous...\n" +
		Tab + "Vous avez à peine le temp de vous protéger, êtes-vous en vie?..\n" +
		DoubleTab + "Coup du déluge de flames ",
}

var HPTR = map[string]string{
	englishLang: " HP",
	frenchLang:  " HP",
}

var SorcererDragonFireTR = map[string]string{
	englishLang: Tab + "Seeing it coming, you quickly whistle a protection spell\n",
	frenchLang:  Tab + "Voyant ce qui arrive, vous souffler un sort de protection\n",
}

var BarbarianLuckDragonTR = map[string]string{
	englishLang: Tab + "Opportunity! You seize the Dragon's wing and jump down,\n" +
		Tab + "using the velocity you dive your weapon first on its neck...\n" +
		DoubleTab + "Dealing -> ",
	frenchLang: Tab + "Opportunité! Vous saisissez l'aile au passage et sautez,\n" +
		Tab + "utilisant la vélocité vous plongez arme en avant sur son cou...\n" +
		DoubleTab + "Causant -> ",
}

var DamageTR = map[string]string{
	englishLang: " Damage!\n",
	frenchLang:  " de Dégats!\n",
}

var GnollTR = map[string]string{
	englishLang: "Gnoll",
	frenchLang:  "Gnoll",
}

var HumanTR = map[string]string{
	englishLang: "Human",
	frenchLang:  "Humain",
}

var TieflingTR = map[string]string{
	englishLang: "Tiefling",
	frenchLang:  "Tiefling",
}

var UndeadTR = map[string]string{
	englishLang: "Undead",
	frenchLang:  "Mort-vivant",
}

var DarklingTR = map[string]string{
	englishLang: "Darkling",
	frenchLang:  "Ténébréen",
}

var usePotionTR = map[string]string{
	englishLang: " use potion +",
	frenchLang:  " utilise potion +",
}

var useScrollTR = map[string]string{
	englishLang: " reads an old scroll, it suddenly turns into ashes,\n" +
		Tab + "as a fireball appears and rush the enemy\n",
	frenchLang: " lit un vieux parchemin qui se change en cendre aussitôt\n" +
		Tab + "alors qu'une boule de feu se rue sur l'ennemi\n",
}

var dmgTR = map[string]string{
	englishLang: " DMG",
	frenchLang:  " DMG",
}

var hitTR = map[string]string{
	englishLang: DoubleTab + "HIT: ",
	frenchLang:  DoubleTab + "HIT: ",
}

var blightStatusTR = map[string]string{
	englishLang: "blight",
	frenchLang:  "Brûlure",
}

var darkStatusTR = map[string]string{
	englishLang: "dark",
	frenchLang:  "Sombre",
}

var plagueStatusTR = map[string]string{
	englishLang: "plague",
	frenchLang:  "Peste",
}

var lightStatusTR = map[string]string{
	englishLang: "light",
	frenchLang:  "Lumière",
}

var frightStatusTR = map[string]string{
	englishLang: "fright",
	frenchLang:  "Effroi",
}

var statusEffectsTR = map[string]string{
	englishLang: "Status:",
	frenchLang:  "Statuts:",
}

var cantMoveTR = map[string]string{
	englishLang: " can't move!",
	frenchLang:  " ne peut pas bouger!",
}

var burnsTR = map[string]string{
	englishLang: " Burns for ",
	frenchLang:  " Brûle pour ",
}

var goHelpTR = map[string]string{
	englishLang: "- Go the selected way",
	frenchLang:  "- Aller dans la direction choisie",
}

var atkHelpTR = map[string]string{
	englishLang: "- Attacks enemies",
	frenchLang:  "- Attaquer les ennemis",
}

var buyHelpTR = map[string]string{
	englishLang: "- Buy ONE OF proposed items",
	frenchLang:  "- Acheter UN DES items proposés",
}

var useHelpTR = map[string]string{
	englishLang: "- Use item",
	frenchLang:  "- Utiliser l'item",
}

var escHelpTR = map[string]string{
	englishLang: "- Try to escape battle",
	frenchLang:  "- Tantative de s'échapper du combat",
}

var invHelpTR = map[string]string{
	englishLang: "- Display your inventory",
	frenchLang:  "- Afficher votre inventaire",
}

var mapHelpTR = map[string]string{
	englishLang: "- Display unveiled World map rooms\n\n",
	frenchLang:  "- Afficher la carte du monde\n\n",
}

var skillHelpTR = map[string]string{
	englishLang: "- Use one of your skills",
	frenchLang:  "- Utiliser une de vos compétences",
}

// var NAME = map[string]string{
// 	englishLang: ,
// 	frenchLang: ,
// }
