package repositories

import (
	"test-farmacare/config"
	"test-farmacare/app/models"
)

type BattlePokemonRepository interface {
	SaveBattlePokemon(BattlePokemon []models.BattlePokemon)([]models.BattlePokemon,error)
}

func NewBattlePokemonRepository() BattlePokemonRepository {
	return &dbConnection{
		connection: config.ConnectDB(),
	}
}

func(db *dbConnection) SaveBattlePokemon(BattlePokemon []models.BattlePokemon)([]models.BattlePokemon,error){
	err := db.connection.Create(BattlePokemon).Error
	if err != nil {
		return BattlePokemon, err
	}
	return BattlePokemon, nil
}