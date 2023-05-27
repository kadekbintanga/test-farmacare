package repositories

import (
	"test-farmacare/config"
	"test-farmacare/app/models"
)

type BattleRepository interface {
	SaveBattle(Battle models.Battle)(models.Battle, error)
	GetBattleById(id uint)(models.Battle, error)
}

func NewBattleRepository() BattleRepository {
	return &dbConnection{
		connection: config.ConnectDB(),
	}
}

func(db *dbConnection) SaveBattle(Battle models.Battle)(models.Battle, error){
	err := db.connection.Save(&Battle).Error
	if err != nil {
		return Battle, err
	}
	return Battle, nil
}


func(db *dbConnection) GetBattleById(id uint)(models.Battle, error){
	var Battle models.Battle
	connection := db.connection.Where("id = ?", id).Preload("BattlePokemon").Find(&Battle)
	err := connection.Error
	if err != nil {
		return Battle, err
	}
	return Battle, nil
}