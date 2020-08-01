package main

func UseKey(p *Character) bool {
	b := false
	loc := p.SetPlayerRoom()
	if loc.HasGate {
		loc.Gate.Warp(p)
		b = true
	}

	return b
}
