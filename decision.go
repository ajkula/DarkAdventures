package main

type Node struct {
	Leaf     bool
	Label    string
	Name     string
	Value    int
	Action   func(p, e *Character)
	Children []Node
}

func MakeEnemyDecision(p1, p2 *Character) {
	var p, e *Character
	if p2.Npc {
		p = p1
		e = p2
	} else {
		p = p2
		e = p1
	}

	var enemyOrDragon int = 20
	if e.Name == enemiesList.DRAGON {
		enemyOrDragon = 35
	}

	Tree := Node{
		Leaf:  false,
		Label: "health",
		Name:  e.Name,
		Value: 50,
		Children: []Node{
			{
				Leaf:  false,
				Label: "random",
				Name:  e.Name,
				Value: 40,
				Children: []Node{
					{
						Leaf:   true,
						Label:  "health",
						Name:   e.Name,
						Action: doHeal,
					},
					{
						Leaf:   true,
						Label:  "attack",
						Name:   e.Name,
						Action: doAttack,
					},
				},
			},
			{
				Leaf:  false,
				Label: "skill",
				Name:  e.Name,
				Value: enemyOrDragon,
				Children: []Node{
					{
						Leaf:   true,
						Label:  "attack",
						Name:   e.Name,
						Action: doAttack,
					},
					{
						Leaf:  false,
						Label: "skill",
						Name:  e.Name,
						Value: enemyOrDragon,
						Children: []Node{
							{
								Leaf:   true,
								Label:  "attack",
								Name:   e.Name,
								Action: doAttack,
							},
							{
								Leaf:   true,
								Label:  "skill",
								Name:   e.Name,
								Action: doSkill,
							},
						},
					},
				},
			},
		},
	}

	TreeVector(Tree, p, e)
}

func doHeal(p, e *Character) {
	if e.hasItemInInventory("potion") {
		e.useItem("potion")
	}
}

func doSkill(p, e *Character) {
	e.useSkillSet(p)
}

func doAttack(p, e *Character) {
	e.attack(p)
}

func TreeVector(tree Node, p, e *Character) {
	if tree.Leaf {
		tree.Action(p, e)
		return
	}

	TreeVector(tree.Children[getTreeIndexNavigation(tree, e)], p, e)
}

func getTreeIndexNavigation(tree Node, e *Character) int {
	switch tree.Label {
	case "health":
		if GetAPercentageOfB(e.Health, e.BaseHealth) >= float32(tree.Value) || !e.hasItemInInventory("potion") {
			return 1
		}
	case "random":
		if PercentChances(tree.Value) {
			return 1
		}
	case "skill":
		if e.Skill > 0 {
			if PercentChances(tree.Value) {
				return 1
			}
		}
	}
	return 0
}
