package routers

import (
	"test-farmacare/app/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)


func InitRouter(){
	PokemonHandler := handlers.NewPokemonHandler()
	BattleHandler := handlers.NewBattleHandler()
	BattlePokemonHandler := handlers.NewBattlePokemonHandler()
	r := gin.Default()
	api :=r.Group("/api/v1")
	api.GET("/health", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message":"I am ready..!",
		})
	})
	api.GET("/health/pokemon", PokemonHandler.HealthPokemon)
	api.GET("/pokemon", PokemonHandler.GetPokemonList)
	api.GET("/health/battle", BattleHandler.HealthBattle)
	api.POST("/battle/auto", BattleHandler.CreateBattleAuto)
	api.POST("/battle/manual", BattleHandler.CreateBattleManual)
	api.GET("/battle", BattleHandler.GetListBattle)
	api.GET("/health/battlepokemon", BattlePokemonHandler.HealthBattlePokemon)
	api.GET("/battlepokemon/score", BattlePokemonHandler.GetTotalScore)
	api.PUT("/battlepokemon/annulled", BattlePokemonHandler.AnnulledPosition)

	r.Run(":8000")
}