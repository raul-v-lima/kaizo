package persistence

import (
	"database/sql"
	"errors"
	"fmt"
	"kaizo/kaizo/internal/domain"
	"time"

	"golang.org/x/exp/rand"
)

type MonRepositoryPostgres struct {
	db *sql.DB
}

func NewMonRepositoryPostgres(db *sql.DB) *MonRepositoryPostgres {
	return &MonRepositoryPostgres{db: db}
}
func (r *MonRepositoryPostgres) GetAll() ([]domain.Mon, error) {
	rows, err := r.db.Query("SELECT id, name, bst, hp, atk, def, spa, spd, spe, lvl, type, type2, nature, ability, move1, move2, move3, move4, immunity, weakness, resistance, updatededAt, createdAt, baseEvo, secondEvo, thirdEvo FROM mons")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mons []domain.Mon
	for rows.Next() {
		var mon domain.Mon
		if err := rows.Scan(&mon.ID, &mon.Name, &mon.Bst, &mon.Hp, &mon.Atk, &mon.Def, &mon.Spa, &mon.Spd, &mon.Spe, &mon.Lvl, &mon.Type1, &mon.Type2, &mon.Nature, &mon.Ability, &mon.Move1, &mon.Move2, &mon.Move3, &mon.Move4, &mon.Weakness, &mon.Resistency, &mon.UpdatededAt, &mon.CreatedAt, &mon.BaseEvo, &mon.SecondEvo, &mon.ThirdEvo); err != nil {
			return nil, err
		}
		mons = append(mons, mon)
	}
	return mons, nil
}

func (r *MonRepositoryPostgres) GetByID(id int) (domain.Mon, error) {

	var mon domain.Mon
	err := r.db.QueryRow("SELECT id, name, bst, hp, atk, def, spa, spd, spe, lvl, type, type2, nature, ability, move1, move2, move3, move4, immunity, weakness, resistency, updatededAt, createdAt, baseEvo, secondEvo, thirdEvo FROM mons WHERE id = $1", id).Scan(&mon.ID, &mon.Name, &mon.Bst, &mon.Hp, &mon.Atk, &mon.Def, &mon.Spa, &mon.Spd, &mon.Spe, &mon.Lvl, &mon.Type1, &mon.Type2, &mon.Nature, &mon.Ability, &mon.Move1, &mon.Move2, &mon.Move3, &mon.Move4, &mon.Weakness, &mon.Resistency, &mon.UpdatededAt, &mon.CreatedAt, &mon.BaseEvo, &mon.SecondEvo, &mon.ThirdEvo)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Mon{}, errors.New("mon not found")
		}
		return domain.Mon{}, err
	}
	return mon, nil
}

func (r *MonRepositoryPostgres) AddMon(mon domain.Mon) error {
	// Inicializando o gerador de números aleatórios
	rand.NewSource(uint64(time.Now().UnixNano()))

	mon.Name = domain.MNames[rand.Intn(len(domain.MNames))]
	mon.Type1 = domain.Types[rand.Intn(len(domain.Types))]
	mon.Type2 = domain.Types[rand.Intn(len(domain.Types))]
	mon.CreatedAt = time.Now()
	mon.UpdatededAt = time.Now()
	mon.Bst = rand.Intn(500) + 200 // Exemplo de valor aleatório entre 200 e 700
	mon.Hp = fmt.Sprintf("%d", rand.Intn(100))
	mon.Atk = fmt.Sprintf("%d", rand.Intn(100))
	mon.Def = fmt.Sprintf("%d", rand.Intn(100))
	mon.Spa = fmt.Sprintf("%d", rand.Intn(100))
	mon.Spd = fmt.Sprintf("%d", rand.Intn(100))
	mon.Spe = fmt.Sprintf("%d", rand.Intn(100))
	mon.Lvl = rand.Intn(100) + 1
	mon.Nature = domain.Natures[rand.Intn(len(domain.Natures))].Name
	mon.Ability = domain.Abilities[rand.Intn(len(domain.Abilities))].Name

	mon.Move1 = domain.Moves[mon.Type1.Name][rand.Intn(len(domain.Moves[mon.Type1.Name]))]
	mon.Move2 = domain.Moves[mon.Type1.Name][rand.Intn(len(domain.Moves[mon.Type1.Name]))]
	mon.Move3 = domain.Moves[mon.Type1.Name][rand.Intn(len(domain.Moves[mon.Type1.Name]))]
	mon.Move4 = domain.Moves[mon.Type1.Name][rand.Intn(len(domain.Moves[mon.Type1.Name]))]
	// mon.Move1 = domain.Moves[mon.Type1.Name][rand.Intn(len(domain.Moves[mon.Type1.Name]))]
	// mon.Move2 = domain.Moves[mon.Move2.Type][rand.Intn(len(domain.Moves[mon.Move2.Type]))]
	// mon.Move3 = domain.Moves[mon.Move3.Type][rand.Intn(len(domain.Moves[mon.Move3.Type]))]
	// mon.Move4 = domain.Moves[mon.Move4.Type][rand.Intn(len(domain.Moves[mon.Move4.Type]))]

	_, err := r.db.Exec("INSERT INTO mons (id, name, bst, hp, atk, def, spa, spd, spe, lvl, type, type2, nature, ability, move1, move2, move3, move4, immunity, weakness, resistency, updatededAt, createdAt, baseEvo, secondEvo, thirdEvo) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24)", mon.ID, mon.Name, mon.Bst, mon.Hp, mon.Atk, mon.Def, mon.Spa, mon.Spd, mon.Spe, mon.Lvl, mon.Type1, mon.Type2, mon.Nature, mon.Ability, mon.Move1, mon.Move2, mon.Move3, mon.Move4, mon.Weakness, mon.Resistency, mon.UpdatededAt, mon.CreatedAt, mon.BaseEvo, mon.SecondEvo, mon.ThirdEvo)
	if err != nil {
		return err
	}
	return nil

}

func (r *MonRepositoryPostgres) AddRandomMons(count int) error {
	for i := 0; i < count; i++ {
		var mon domain.Mon
		err := r.AddMon(mon)
		if err != nil {
			return err
		}
	}
	return nil
}
