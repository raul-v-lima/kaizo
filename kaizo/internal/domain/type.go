package domain

type Type struct {
	Name string `json:"name"`
}

var Types = []Type{
	{Name: "normal"},
	{Name: "fire"},
	{Name: "water"},
	{Name: "electric"},
	{Name: "grass"},
	{Name: "ice"},
	{Name: "fighting"},
	{Name: "poison"},
	{Name: "ground"},
	{Name: "flying"},
	{Name: "psychic"},
	{Name: "bug"},
	{Name: "rock"},
	{Name: "ghost"},
	{Name: "dragon"},
	{Name: "dark"},
	{Name: "steel"},
	{Name: "fairy"},
}
