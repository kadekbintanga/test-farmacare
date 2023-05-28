package repositories

import (
	"test-farmacare/config"
	"test-farmacare/app/models"
	"fmt"
	"time"
)

type BattleRepository interface {
	SaveBattle(Battle models.Battle)(models.Battle, error)
	GetBattleById(id uint)(models.Battle, error)
	GetListBattle(start_date string, end_date string, limit int, offset int)([]models.Battle, int64, error)
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

func(db *dbConnection) GetListBattle(start_date string, end_date string, limit int, offset int)([]models.Battle, int64, error){
	var Battle []models.Battle
	var count int64
	connection := db.connection.Debug().Model(&Battle)
	if start_date != ""{
		time_loc,_:= time.LoadLocation("Asia/Jakarta")
		startDate,_ := time.ParseInLocation("2006-01-02 15:04:05", start_date, time_loc)
		fmt.Println(startDate)
		connection = connection.Where("created_at >= ?", startDate)
	}
	if end_date != ""{
		time_loc,_:= time.LoadLocation("Asia/Jakarta")
		endDate,_ := time.ParseInLocation("2006-01-02 15:04:05", end_date, time_loc)
		connection = connection.Where("created_at <= ?", endDate)
	}
	connection.Count(&count)
	connection = connection.Preload("BattlePokemon").Offset(offset).Limit(limit).Find(&Battle)
	err := connection.Error
	if err != nil {
		return Battle, 0, err
	}
	return Battle, count, nil
}