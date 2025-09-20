package dndcharacter

import (
    "math/rand"
    "time"
    "sort"
    "math"
)
type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	modifier := int(math.Floor(float64(score - 10) / 2.0))
    return modifier
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
    rand.Seed(time.Now().UnixNano())
    score := 0
    random := make([]int, 4)
    for i := 0; i < 4; i++ {
    	random[i] = rand.Intn(6) + 1
    }
    sort.Ints(random)
    for i := 3;i >= 1; i-- {
        score += random[i]
    }
    return score
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
    constitution := Ability()
	return Character {
        Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: constitution,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    Modifier(constitution) + 10,
    }
}
