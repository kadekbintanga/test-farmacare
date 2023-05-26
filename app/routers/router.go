package routers

import (
	"test-farmacare/app/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)


func InitRouter(){
	PokemonHandler := handlers.NewPokemonHandler()
	r := gin.Default()
	api :=r.Group("/api/v1")
	api.GET("/health", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message":"I am ready..!",
		})
	})
	api.GET("/health/pokemon", PokemonHandler.HealthPokemon)
	api.GET("/pokemon", PokemonHandler.GetPokemonList)

	r.Run(":8000")
}