package domain

type Resistency struct {
	ResName []Type `json:"resname"`
}

var ResistenciesByType = map[string]Resistency{
	"normal": {
		[]Type{{Name: "fighting"}},
	},
	"fire": {
		[]Type{{Name: "fire"}, {Name: "grass"}, {Name: "ice"}, {Name: "bug"}, {Name: "steel"}, {Name: "fairy"}},
	},
	"water": {
		[]Type{{Name: "fire"}, {Name: "water"}, {Name: "ice"}, {Name: "steel"}},
	},
	"electric": {
		[]Type{{Name: "electric"}, {Name: "flying"}, {Name: "steel"}},
	},
	"grass": {
		[]Type{{Name: "water"}, {Name: "electric"}, {Name: "grass"}, {Name: "	ground"}},
	},
	"ice": {
		[]Type{{Name: "ice"}},
	},
	"fighting": {
		[]Type{{Name: "bug"}, {Name: "rock"}, {Name: "dark"}},
	},
	"poison": {
		[]Type{{Name: "grass"}, {Name: "fighting"}, {Name: "poison"}, {Name: "bug"}, {Name: "fairy"}},
	},
	"ground": {
		[]Type{{Name: "poison"}, {Name: "rock"}, {Name: "electric"}, {Name: "steel"}, {Name: "fairy"}},
	},
	"flying": {
		[]Type{{Name: "fighting"}, {Name: "bug"}, {Name: "grass"}},
	},
	"psychic": {
		[]Type{{Name: "fighting"}, {Name: "psychic"}},
	},
	"bug": {
		[]Type{{Name: "grass"}, {Name: "fighting"}, {Name: "ground"}},
	},
	"rock": {
		[]Type{{Name: "normal"}, {Name: "fire"}, {Name: "poison"}, {Name: "flying"}},
	},
	"ghost": {
		[]Type{{Name: "poison"}, {Name: "bug"}},
	},
	"dragon": {
		[]Type{{Name: "fire"}, {Name: "water"}, {Name: "electric"}, {Name: "grass"}},
	},
	"dark": {
		[]Type{{Name: "ghost"}, {Name: "dark"}},
	},
	"steel": {
		[]Type{{Name: "normal"}, {Name: "grass"}, {Name: "ice"}, {Name: "flying"}, {Name: "psychic"}, {Name: "bug"}, {Name: "rock"}, {Name: "dragon"}, {Name: "steel"}, {Name: "fairy"}},
	},
	"fairy": {
		[]Type{{Name: "fighting"}, {Name: "bug"}, {Name: "dark"}},
	},
}

func GetResistencies(monResistencies string) Resistency {

	monRes := ResistenciesByType[monResistencies]
	return monRes
}
