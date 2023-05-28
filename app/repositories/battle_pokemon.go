package repositories

import (
	"test-farmacare/config"
	"test-farmacare/app/models"
)

type BattlePokemonRepository interface {
	SaveBattlePokemon(BattlePokemon []models.BattlePokemon)([]models.BattlePokemon,error)
	GetTotalScore()([]map[string]interface{}, error)
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

func(db *dbConnection) GetTotalScore()([]map[string]interface{}, error){
	var BattlePokemon []models.BattlePokemon
	var result []map[string]interface{}
	connection := db.connection.Model(&BattlePokemon).Select("pokemon_name, sum(score) as total_score").Group("pokemon_name").Order("total_score desc").Find(&result)
	err := connection.Error
	if err != nil {
		return result, err
	}
	return result, nil
}