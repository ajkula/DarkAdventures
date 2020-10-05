# DarkAdventures
A simple text adventure game I build in Go for fun...

![A simple text adventure game I build in Go for fun...](https://github.com/ajkula/DarkAdventures/blob/master/scr/title.png)

## Installation
It's as simple as clone/compile/run...

You might want to **go get**:
 - github.com/fatih/color
 - github.com/cloudfoundry/jibber_jabber

There are a few steps to make the experience much better however:
 - GDrive URL to download the game:
    * [**Windows 64**](https://tinyurl.com/y4y4avp5)
    * [**Mac 64**](https://tinyurl.com/y57ktnbh)
    * [**Linux 64**](https://tinyurl.com/y3odpymo)
 - I recommend using a terminal emulator on windows such as: 
 **https://conemu.github.io/en/TableOfContents.html**
 - The software will create the default *landscape layout*, it has only a weak contextual influence for the moment but will have more depth later.
 - The mapCreator tool will generate a new random world map for you replacing the existing one.
 - If the software was started *directly* it might close the console abruptly on exit, I therefore recommend opening the terminal and start the application from it.
 - By default the game will get the O/S language

## How to use

**The game is a text based adventure**

> Select the difficulty level, will have an influence on enemies probability:
>  1. Will have a 15% **Enemy Chance**
>  2. Will have a 30% **Enemy Chance**
>  3. Will have a 45% **Enemy Chance**

**There are 4 Hero classes to choose from**

> Select the hero you want, they each have their attributes, passive contextual and item chances:
>  1. **Thief** has the best avoid chances, has a chance to attack twice with a second attack dealing 60% of regular damages.
>  2. **Paladin** has more potions chances, a good attack and avoid chances. He's the most balanced.
>  3. **Wizard** doesn't have a good resistance and attack but starts with potions and scrolls.
>   4. **Barbarian** has a good attack, more health, but less potions chances.

![The Thief](https://github.com/ajkula/DarkAdventures/blob/master/scr/character.png)

### Great tools to your disposal

You will come across many sellers with variable deals during your trip, along with enemies dropping their items or gold coins and experience upon defeat.
As you kill enemies, you might rank up LVLs, but if enemies kill you and you get back up, know they'll acquire your exp value and then might LVL UP accordingly as well!

Collect the best items to make it to through alive:

- **Gold Coins** to buy goods from the many shops.
- **Scrolls** are powerful spells you can use on enemies, scrolls go to dust on cast.
- **Potions** will instantly give you back a maximum of **20 HP**, it won't give you more health than your maximum.
- **Keys** to open chests or other mechanisms, will break on use.
- **Moonstones** once used will instantly infuse into your arms and raise your strength by **5 Points**.
- **Dolls** are a weird unexplained find, owners of that unworldly curiosity have reported to wake up after they thought they died. Even the weakest of them felt stronger for a while! Will suddenly disappear and revive you instantly with **30 HP** no matter your base health.
- **Skills** are a set of various techniques each Hero can do to vanquish enemis, you earn them randomly on lvl up. Currently all heroes have their Skills available.
 
###### Regarding enemies skills available:
 - [x] Dragon
 - [x] Goblin
 - [x] Sorcerer
 - [x] Skeleton
 - [ ] Orc
 - [ ] NightWalkers

### Help NPCs Quests

Some NPCs will have some requests for you, help them to get EXP rewards.
Quests are of these categories: kill, retrieve, save.
They work like so:
- **kill** just kill the requested targets to bring some peace to this world.
- **retrieve** bring the NPC what they need.
- **save** some monsters on the world map hold the NPC's friends, find and kill these monsters to release the hostages.

### Status effects and Passive capacities

The status effects are reducing max health, HP or dealing DMG over time...
You'll see different ones that you can apply to enemies or get from them.
Characters have some passive capacities attack or defense based, by either classes or races.

### The mightiest of foes

A **Dragon** roams the North parts of the World Map, you will need to be well prepared to take down this enemy.

![The Dragon](https://github.com/ajkula/DarkAdventures/blob/master/scr/dragon.png)

### Side tools

mapCreator is a utility to create random landscape.txt (WorldMap) for you.
It will replace the landscape.txt file in the same directory.
The goal is to have random maps filed with different types of rooms, with random sizes from 10x10 to 20x20.