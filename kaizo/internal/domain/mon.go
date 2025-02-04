package domain

import (
	"time"
)

type Mon struct {
	ID          string                    `json:"id"`
	Name        MonNames                  `json:"name"`
	Bst         int                       `json:"bst"`
	Hp          string                    `json:"hp"`
	Atk         string                    `json:"atk"`
	Def         string                    `json:"def"`
	Spa         string                    `json:"spa"`
	Spd         string                    `json:"spd"`
	Spe         string                    `json:"spe"`
	Lvl         int                       `json:"lvl"`
	Type1       Type                      `json:"Type1"`
	Type2       Type                      `json:"Type2"`
	Nature      string                    `json:"nature"`
	Ability     string                    `json:"ability"`
	Move1       Move                      `json:"move1"`
	Move2       Move                      `json:"move2"`
	Move3       Move                      `json:"move3"`
	Move4       Move                      `json:"move4"`
	Resistency  []WeaknessAndResistencies `json:"resistency"`
	Weakness    []WeaknessAndResistencies `json:"weakness"`
	UpdatededAt time.Time                 `json:"updatededAt"`
	CreatedAt   time.Time                 `json:"createdAt"`
	BaseEvo     *Mon                      `json:"baseEvo"`
	SecondEvo   *Mon                      `json:"secondEvo"`
	ThirdEvo    *Mon                      `json:"thirdEvo"`
}
