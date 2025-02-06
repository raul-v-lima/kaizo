package domain

import "fmt"

type Weakness struct {
	WeaknessName []Type `json:"weaknessname"`
}

var Weaknesses = map[string]Weakness{
	"normal": {
		[]Type{{Name: "fighting"}},
	}, "fire": {
		[]Type{{Name: "water"}, {Name: "rock"}, {Name: "dragon"}},
	}, "water": {
		[]Type{{Name: "electric"}, {Name: "grass"}},
	}, "electric": {
		[]Type{{Name: "ground"}},
	}, "grass": {
		[]Type{{Name: "fire"}, {Name: "ice"}, {Name: "poison"}, {Name: "flying"}, {Name: "bug"}},
	}, "ice": {
		[]Type{{Name: "fire"}, {Name: "fighting"}, {Name: "rock"}, {Name: "steel"}},
	}, "fighting": {
		[]Type{{Name: "flying"}, {Name: "psychic"}, {Name: "fairy"}},
	}, "poison": {
		[]Type{{Name: "ground"}, {Name: "psychic"}},
	}, "ground": {
		[]Type{{Name: "water"}, {Name: "grass"}, {Name: "ice"}},
	}, "flying": {
		[]Type{{Name: "electric"}, {Name: "ice"}, {Name: "rock"}},
	}, "psychic": {
		[]Type{{Name: "bug"}, {Name: "ghost"}, {Name: "dark"}},
	}, "bug": {
		[]Type{{Name: "fire"}, {Name: "flying"}, {Name: "rock"}},
	}, "rock": {
		[]Type{{Name: "water"}, {Name: "grass"}, {Name: "fighting"}, {Name: "ground"}, {Name: "steel"}},
	}, "ghost": {
		[]Type{{Name: "ghost"}, {Name: "dark"}},
	}, "dragon": {
		[]Type{{Name: "ice"}, {Name: "dragon"}, {Name: "fairy"}},
	}, "dark": {
		[]Type{{Name: "fighting"}, {Name: "bug"}, {Name: "fairy"}},
	}, "steel": {
		[]Type{{Name: "fire"}, {Name: "fighting"}, {Name: "ground"}},
	}, "fairy": {
		[]Type{{Name: "poison"}, {Name: "steel"}},
	}}

func GetWeakness(monType string) Weakness {

	monweaknesses, exists := Weaknesses[monType]
	if !exists {
		fmt.Println("Weaknesses not found")
	}
	return monweaknesses
}
