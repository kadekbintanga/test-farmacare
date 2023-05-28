package repositories

import (
	"test-farmacare/config"
	"test-farmacare/app/models"
)

type BattlePokemonRepository interface {
	SaveBattlePokemon(BattlePokemon []models.BattlePokemon)([]models.BattlePokemon,error)
	GetTotalScore()([]map[string]interface{}, error)
	GetPokemonBattleByUuuid(uuid string)(models.BattlePokemon, error)
	GetPokemonBattleByPositionBattleId(position uint, battle_id uint)(models.BattlePokemon, error)
	UpdatePosition(id uint, update_data map[string]interface{})(error)
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

func(db *dbConnection) GetPokemonBattleByUuuid(uuid string)(models.BattlePokemon, error){
	var BattlePokemon models.BattlePokemon
	err := db.connection.Where("uuid = ?", uuid).Find(&BattlePokemon).Error
	if err != nil {
		return BattlePokemon, err
	}
	return BattlePokemon, nil
}

func(db *dbConnection) GetPokemonBattleByPositionBattleId(position uint, battle_id uint)(models.BattlePokemon, error){
	var BattlePokemon models.BattlePokemon
	err := db.connection.Where("position = ? AND battle_id = ?", position, battle_id).Find(&BattlePokemon).Error
	if err != nil {
		return BattlePokemon, err
	}
	return BattlePokemon, nil
}

func(db *dbConnection) UpdatePosition(id uint, update_data map[string]interface{})(error){
	err := db.connection.Table("battle_pokemons").Where("id = ?", id).Updates(update_data).Error
	if err != nil {
		return err
	}
	return nil
}