package domain

import (
	"time"
)

type Mon struct {
	ID          string     `json:"id" gorm:"primaryKey"`
	Name        MonNames   `json:"name" gorm:"embedded"`
	Bst         int        `json:"bst"`
	Hp          string     `json:"hp"`
	Atk         string     `json:"atk"`
	Def         string     `json:"def"`
	Spa         string     `json:"spa"`
	Spd         string     `json:"spd"`
	Spe         string     `json:"spe"`
	Lvl         int        `json:"lvl"`
	Type1       Type       `json:"Type1" gorm:"embedded"`
	Type2       Type       `json:"Type2" gorm:"embedded"`
	Nature      string     `json:"nature"`
	Ability     string     `json:"ability"`
	Move1       Move       `json:"move1" gorm:"embedded"`
	Move2       Move       `json:"move2" gorm:"embedded"`
	Move3       Move       `json:"move3" gorm:"embedded"`
	Move4       Move       `json:"move4" gorm:"embedded"`
	Resistency  Resistency `json:"resistency" gorm:"embedded"`
	Weakness    Weakness   `json:"weakness" gorm:"embedded"`
	UpdatededAt time.Time  `json:"updatededAt"`
	CreatedAt   time.Time  `json:"createdAt"`

	// Evoluções
	BaseEvoID *string `json:"baseEvoId"`
	BaseEvo   *Mon    `json:"baseEvo" gorm:"foreignKey:BaseEvoID"`

	SecondEvoID *string `json:"secondEvoId"`
	SecondEvo   *Mon    `json:"secondEvo" gorm:"foreignKey:SecondEvoID"`

	ThirdEvoID *string `json:"thirdEvoId"`
	ThirdEvo   *Mon    `json:"thirdEvo" gorm:"foreignKey:ThirdEvoID"`
}
