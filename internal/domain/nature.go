package domain

type Nature struct {
	Name   string
	Effect string
}

var Natures = []Nature{
	{Name: "Hardy", Effect: "No effect"},
	{Name: "Lonely", Effect: "Attack up, Defense down"},
	{Name: "Brave", Effect: "Attack up, Speed down"},
	{Name: "Adamant", Effect: "Attack up	, Special Attack down"},
	{Name: "Naughty", Effect: "Attack up, Special Defense down"},
	{Name: "Bold", Effect: "Defense up, Attack down"},
	{Name: "Docile", Effect: "No effect"},
	{Name: "Relaxed", Effect: "Defense up, Speed down"},
	{Name: "Impish", Effect: "Defense up, Special Attack down"},
	{Name: "Lax", Effect: "Defense up, Special Defense down"},
	{Name: "Timid", Effect: "Speed up, Attack down"},
	{Name: "Hasty", Effect: "Speed up, Defense down"},
	{Name: "Serious", Effect: "No effect"},
	{Name: "Jolly", Effect: "Speed up, Special Attack down"},
	{Name: "Naive", Effect: "Speed up, Special Defense down"},
	{Name: "Modest", Effect: "Special Attack up, Attack down"},
	{Name: "Mild", Effect: "Special Attack up, Defense down"},
	{Name: "Quiet", Effect: "Special Attack up, Speed down"},
	{Name: "Bashful", Effect: "No effect"},
	{Name: "Rash", Effect: "Special Attack up, Special Defense down"},
	{Name: "Calm", Effect: "Special Defense up, Attack down"},
	{Name: "Gentle", Effect: "Special Defense up, Defense down"},
	{Name: "Sassy", Effect: "Special Defense up, Speed down"},
	{Name: "Careful", Effect: "Special Defense up, Special Attack down"},
	{Name: "Quirky", Effect: "No effect"},
}
