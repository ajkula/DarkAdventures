package main

import (
	"testing"
)

// Mocking la struct Character pour le test. Assurez-vous d'ajuster ceci en fonction de votre implémentation réelle.

// Les tests

func TestEnemyDecisionHeal(t *testing.T) {
	player := &Character{
		Name:       "Player",
		Health:     100,
		BaseHealth: 100,
		Skill:      50,
	}

	enemy := &Character{
		Name:       "Enemy",
		Health:     40, // santé faible pour encourager la guérison
		BaseHealth: 100,
		Skill:      50,
		Npc:        true,
		Inventory:  map[string]*ItemQuantity{"potion": &ItemQuantity{Quantity: 1}}, // donner une potion à l'ennemi
	}

	MakeEnemyDecision(player, enemy)

	if enemy.Health == 40 {
		t.Errorf("Expected the enemy to heal but it didn't. Current health: %d", enemy.Health)
	}
}

func TestEnemyDecisionAttack(t *testing.T) {
	player := &Character{
		Name:       "Player",
		Health:     100,
		BaseHealth: 100,
		Skill:      50,
	}

	enemy := &Character{
		Name:       "Enemy",
		Health:     100,
		BaseHealth: 100,
		Skill:      50,
		Npc:        true,
		Inventory:  map[string]*ItemQuantity{}, // aucun objet
	}

	MakeEnemyDecision(player, enemy)

	if player.Health == 100 {
		t.Errorf("Expected the enemy to attack but it didn't. Player's current health: %d", player.Health)
	}
}
