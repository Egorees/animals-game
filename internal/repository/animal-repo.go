package repository

import "animals-game/internal/animals"

func (repo *Repository) CreateAnimal(animal *animals.Animal) error {
	request := `INSERT INTO animals(name, type, exp, owner_id) VALUES($1, $2, $3, $4) RETURNING id`

	var animalId int

	err := repo.db.QueryRow(request, animal.Name, animal.Type, animal.Exp, animal.OwnerId).Scan(&animalId)

	if err != nil {
		return err
	}

	err = repo.SetNewCurrentAnimal(animal.OwnerId, animalId)

	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) SetNewCurrentAnimal(userId int, animalId int) error {
	request := `UPDATE users SET animal_id=$1 WHERE id = $2`

	_, err := repo.db.Exec(request, animalId, userId)

	if err != nil {
		return err
	}

	return nil
}
