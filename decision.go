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

	Tree := Node{
		Leaf:  false,
		Label: "health",
		Name:  e.Name,
		Value: 50,
		Children: []Node{
			Node{
				Leaf:  false,
				Label: "random",
				Name:  e.Name,
				Value: 60,
				Children: []Node{
					Node{
						Leaf:   true,
						Label:  "health",
						Name:   e.Name,
						Action: doHeal,
					},
					Node{
						Leaf:   true,
						Label:  "attack",
						Name:   e.Name,
						Action: doAttack,
					},
				},
			},
			Node{
				Leaf:  false,
				Label: "random",
				Name:  e.Name,
				Value: 80,
				Children: []Node{
					Node{
						Leaf:   true,
						Label:  "attack",
						Name:   e.Name,
						Action: doAttack,
					},
					Node{
						Leaf:  false,
						Label: "skill",
						Name:  e.Name,
						Children: []Node{
							Node{
								Leaf:   true,
								Label:  "attack",
								Name:   e.Name,
								Action: doAttack,
							},
							Node{
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
	if e.hasItemInInventory("potion") {
		e.useItem("potion")
	}
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
	index := 0
	switch tree.Label {
	case "health":
		if GetAPercentageOfB(e.Health, e.BaseHealth) >= float32(tree.Value) || !e.hasItemInInventory("potion") {
			index = 1
		}
		break
	case "random":
		if !PercentChances(tree.Value) {
			index = 1
		}
		break
	case "skill":
		if e.Skill > 0 {
			index = 1
		}
		break
	default:
		return 0
	}
	// ICI
	Output("red", Tab+tree.Label+" ", index, " leaf: ", tree.Leaf)
	return index
}