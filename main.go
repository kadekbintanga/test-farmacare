package main

import (
	"fmt"
	"test-farmacare/app/routers"
	"test-farmacare/config"
	"gorm.io/gorm"
)

var(
	db *gorm.DB = config.ConnectDB()
)

func main(){
	fmt.Println("-------------------- SERVICE STARTED --------------------")
	config.LoadEnv()
	config.MigrateDatabase(db)
	defer config.DisconnectDB(db)
	routers.InitRouter()
}