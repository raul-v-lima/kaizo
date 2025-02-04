package domain

type WeaknessAndResistencies struct {
	Type         Type   `json:"type"`
	Weaknessess  []Type `json:"weaknessess"`
	Resistencies []Type `json:"resistencies"`
}

var Weaknesses = map[string]WeaknessAndResistencies{
	"normal": {
		Type:        Type{Name: "normal"},
		Weaknessess: []Type{{Name: "fighting"}},
	},
	"fire": {
		Type:        Type{Name: "fire"},
		Weaknessess: []Type{{Name: "water"}, {Name: "rock"}, {Name: "dragon"}},
	},
	"water": {
		Type:        Type{Name: "water"},
		Weaknessess: []Type{{Name: "electric"}, {Name: "grass"}},
	},
	"electric": {
		Type:        Type{Name: "electric"},
		Weaknessess: []Type{{Name: "ground"}},
	},
	"grass": {
		Type:        Type{Name: "grass"},
		Weaknessess: []Type{{Name: "fire"}, {Name: "ice"}, {Name: "poison"}, {Name: "flying"}, {Name: "bug"}},
	},
	"ice": {
		Type:        Type{Name: "ice"},
		Weaknessess: []Type{{Name: "fire"}, {Name: "fighting"}, {Name: "rock"}, {Name: "steel"}},
	},
	"fighting": {
		Type:        Type{Name: "fighting"},
		Weaknessess: []Type{{Name: "flying"}, {Name: "psychic"}, {Name: "fairy"}},
	},
	"poison": {
		Type:        Type{Name: "poison"},
		Weaknessess: []Type{{Name: "ground"}, {Name: "psychic"}},
	},
	"ground": {
		Type:        Type{Name: "ground"},
		Weaknessess: []Type{{Name: "water"}, {Name: "grass"}, {Name: "ice"}},
	},
	"flying": {
		Type:        Type{Name: "flying"},
		Weaknessess: []Type{{Name: "electric"}, {Name: "ice"}, {Name: "rock"}},
	},
	"psychic": {
		Type:        Type{Name: "psychic"},
		Weaknessess: []Type{{Name: "bug"}, {Name: "ghost"}, {Name: "dark"}},
	},
	"bug": {
		Type:        Type{Name: "bug"},
		Weaknessess: []Type{{Name: "fire"}, {Name: "flying"}, {Name: "rock"}},
	},
	"rock": {
		Type:        Type{Name: "rock"},
		Weaknessess: []Type{{Name: "water"}, {Name: "grass"}, {Name: "fighting"}, {Name: "ground"}, {Name: "steel"}},
	},
	"ghost": {
		Type:        Type{Name: "ghost"},
		Weaknessess: []Type{{Name: "ghost"}, {Name: "dark"}},
	},
	"dragon": {
		Type:        Type{Name: "dragon"},
		Weaknessess: []Type{{Name: "ice"}, {Name: "dragon"}, {Name: "fairy"}},
	},
	"dark": {
		Type:        Type{Name: "dark"},
		Weaknessess: []Type{{Name: "fighting"}, {Name: "bug"}, {Name: "fairy"}},
	},
	"steel": {
		Type:        Type{Name: "steel"},
		Weaknessess: []Type{{Name: "fire"}, {Name: "fighting"}, {Name: "ground"}},
	},
	"fairy": {
		Type:        Type{Name: "fairy"},
		Weaknessess: []Type{{Name: "poison"}, {Name: "steel"}},
	},
}

var Resistencies = map[string]WeaknessAndResistencies{
	"normal": {
		Type:         Type{Name: "normal"},
		Resistencies: []Type{{Name: "fighting"}},
	},
	"fire": {
		Type:         Type{Name: "fire"},
		Resistencies: []Type{{Name: "fire"}, {Name: "grass"}, {Name: "ice"}, {Name: "bug"}, {Name: "steel"}, {Name: "fairy"}},
	},
	"water": {
		Type:         Type{Name: "water"},
		Resistencies: []Type{{Name: "fire"}, {Name: "water"}, {Name: "ice"}, {Name: "steel"}},
	},
	"electric": {
		Type:         Type{Name: "electric"},
		Resistencies: []Type{{Name: "electric"}, {Name: "flying"}, {Name: "steel"}},
	},
	"grass": {
		Type:         Type{Name: "grass"},
		Resistencies: []Type{{Name: "water"}, {Name: "electric"}, {Name: "grass"}, {Name: "	ground"}},
	},
	"ice": {
		Type:         Type{Name: "ice"},
		Resistencies: []Type{{Name: "ice"}},
	},
	"fighting": {
		Type:         Type{Name: "fighting"},
		Resistencies: []Type{{Name: "bug"}, {Name: "rock"}, {Name: "dark"}},
	},
	"poison": {
		Type:         Type{Name: "poison"},
		Resistencies: []Type{{Name: "grass"}, {Name: "fighting"}, {Name: "poison"}, {Name: "bug"}, {Name: "fairy"}},
	},
	"ground": {
		Type:         Type{Name: "ground"},
		Resistencies: []Type{{Name: "poison"}, {Name: "rock"}, {Name: "electric"}, {Name: "steel"}, {Name: "fairy"}},
	},
	"flying": {
		Type:         Type{Name: "flying"},
		Resistencies: []Type{{Name: "fighting"}, {Name: "bug"}, {Name: "grass"}},
	},
	"psychic": {
		Type:         Type{Name: "psychic"},
		Resistencies: []Type{{Name: "fighting"}, {Name: "psychic"}},
	},
	"bug": {
		Type:         Type{Name: "bug"},
		Resistencies: []Type{{Name: "grass"}, {Name: "fighting"}, {Name: "ground"}},
	},
	"rock": {
		Type:         Type{Name: "rock"},
		Resistencies: []Type{{Name: "normal"}, {Name: "fire"}, {Name: "poison"}, {Name: "flying"}},
	},
	"ghost": {
		Type:         Type{Name: "ghost"},
		Resistencies: []Type{{Name: "poison"}, {Name: "bug"}},
	},
	"dragon": {
		Type:         Type{Name: "dragon"},
		Resistencies: []Type{{Name: "fire"}, {Name: "water"}, {Name: "electric"}, {Name: "grass"}},
	},
	"dark": {
		Type:         Type{Name: "dark"},
		Resistencies: []Type{{Name: "ghost"}, {Name: "dark"}},
	},
	"steel": {
		Type:         Type{Name: "steel"},
		Resistencies: []Type{{Name: "normal"}, {Name: "grass"}, {Name: "ice"}, {Name: "flying"}, {Name: "psychic"}, {Name: "bug"}, {Name: "rock"}, {Name: "dragon"}, {Name: "steel"}, {Name: "fairy"}},
	},
	"fairy": {
		Type:         Type{Name: "fairy"},
		Resistencies: []Type{{Name: "fighting"}, {Name: "bug"}, {Name: "dark"}},
	},
}
