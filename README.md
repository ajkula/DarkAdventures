# DarkAdventures
A simple text adventure game I build in Go for fun...

## Installation
It's as simple as clone/compile/run...

You might want to **go get**:
 - github.com/fatih/color
 - github.com/cloudfoundry/jibber_jabber

There are a few steps to make the experience much better however:
 - GDrive URL to download the game:
  - **https://preview.tinyurl.com/y4v69bsn** (Windows)
  - **https://preview.tinyurl.com/** (Mac)
  - **https://preview.tinyurl.com/y38x2d7f** (Linux)
 - I recommend using a terminal emulator on windows such as: 
 **https://conemu.github.io/en/TableOfContents.html**
 - The software will create the default *landscape layout*, it has only a weak contextual influence for the moment but will have more depth later.
 - If the software was started *directly* it might close the console abruptly on exit, I therefore recommend opening the terminal and start the application from it.
 - By default the game will get the O/S language

## How to use

**The game is a text based adventure**

> Select the difficulty level, will have an influence on enemies probability:
>  1. Will have a 15% **Enemy Chance**
>  2. Will have a 30% **Enemy Chance**
>  3. Will have a 45% **Enemy Chance**

**There are 4 Hero classes to choose from**

> Select the hero you want, they have different attributes and item chances:
>  1. **Thieve** has the best avoid chances, has a chance to attack twice with a second attack dealing 60% of regular damages.
>  2. **Paladin** has more potions chances, a good attack and avoid chances. He's the most balanced.
>  3. **Wizard** doesn't have a good resistance and attack but starts with potions and scrolls.
>   4. **Barbarian** has a good attack, more health, but less potions chances.

### Great tools to your disposal

You will come across many sellers with variable deals during your trip, along with enemies dropping their items or gold coins and experience upon defeat.
As you kill enemies, you might rank up LVLs, but if enemies kill you and you get back up, know they'll acquire your exp value and then might LVL UP accordingl as well!

Collect the best items to make it to through alive:

- **Gold Coins** to buy goods from the many shops.
- **Scrolls** are powerful spells you can use on enemies, they go to dust on cast.
- **Potions** will instantly give you back a maximum of **20 HP**, it won't give you more health than your maximum.
- **Keys** to open chests or other mechanisms, will break on use.
- **Moonstones** once used will instantly infuse into your arms and raise your strength by **5 Points**.
- **Dolls** are a weird unexplained find, owners of that unworldly curiosity have reported to wake up after they thought they died. Even the weakest of them felt stronger for a while! Will suddenly disappear and revive you instantly with **30 HP** no matter your base health.
 - **Skills** are a set of various techniques each Hero can do to vanquish enemis, you earn them randomly on lvl up or at least one each 2 levels. Currently only the Thief has his Skill available, regarding enemies: Dragon, Goblin, Sorcerer have theirs.

### More dangers than it seems

There aren't only enemy hits and Skills you should worry about, but a few status effects are making their ways into the game, most of them reducing max health, HP or dealing DMG over time...
Generaly killing the responsible foe would cure them, but not in all cases.
Those are about to come ingame soon, the skeleton mechanics are set, now all is needed are the different statuses and their effects.

### The mightiest of foes

A **Dragon** roams the North parts of the World Map, you will need to be well prepared to take down this enemy.