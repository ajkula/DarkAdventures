package main

type iCharacter interface {
	addItemTypeToInventory(n string, i int)

	useItem(name string, enemyInArr ...interface{}) bool

	showHP()

	attack(enemy *iCharacter)

	calculateDammage(enemy *iCharacter) int

	isAlive() bool

	DisplayInvetory()

	DisplayItems() string

	hasItemInInventory(name string) bool

	getEnemyItems(enemy *iCharacter)

	createEnemyInventory()

	getPurse() int

	spendMoney(amount int) bool

	BuyFromShop(name string) bool
}
