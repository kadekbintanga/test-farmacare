package handlers

import (
	"test-farmacare/app/helpers"
	"test-farmacare/app/repositories"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type BattlePokemonHandler struct {
	repo_battlepokemon repositories.BattlePokemonRepository
}

func NewBattlePokemonHandler() *BattlePokemonHandler{
	return &BattlePokemonHandler{
		repositories.NewBattlePokemonRepository(),
	}
}

func(h *BattlePokemonHandler) HealthBattlePokemon(c *gin.Context){
	log.Info("Health Battle Pokemon Handler is Hitted")
	response := helpers.APIResponse("Success", http.StatusOK, gin.H{"message": "Success hit Health Battle Pokemon Handler"})
	c.JSON(http.StatusOK, response)
	return
}

func(h *BattlePokemonHandler) GetTotalScore(c *gin.Context){
	repo_battlepokemon := h.repo_battlepokemon
	res, err := repo_battlepokemon.GetTotalScore()
	if err != nil {
		log.Error(err)
		errorMessage := gin.H{"error message": "Something went wrong"}
		response := helpers.APIResponse("Internal Server Error", http.StatusInternalServerError, errorMessage)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response := helpers.APIResponse("Success", http.StatusOK, res)
	c.JSON(http.StatusOK, response)
	return
}


func(h *BattlePokemonHandler) AnnulledPosition(c *gin.Context){
	uuid_pokemon := c.DefaultQuery("uuid_pokemon","")
	if uuid_pokemon == ""{
		log.Error("Param uuid_pokemon is empty string")
		errorMessage := gin.H{"error message": "uuid_pokemon param is required"}
		response := helpers.APIResponse("Bad Request", http.StatusBadRequest, errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	repo_battlepokemon := h.repo_battlepokemon

	data_pokemon1, err := repo_battlepokemon.GetPokemonBattleByUuuid(uuid_pokemon)
	if err != nil {
		log.Error(err)
		errorMessage := gin.H{"error message": "Something went wrong"}
		response := helpers.APIResponse("Internal Server Error", http.StatusInternalServerError, errorMessage)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	if data_pokemon1.ID == 0{
		log.Error("Uuid pokemon not found")
		errorMessage := gin.H{"error message": "uuid_pokemon not found, Please check the uuid_pokemon"}
		response := helpers.APIResponse("Bad Request", http.StatusBadRequest, errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if data_pokemon1.Position == 5{
		log.Error("position data pokemon 1 is the last order")
		errorMessage := gin.H{"error message": "uuid_pokemon cannot be annulled. Because the pokemon got last position of the battle"}
		response := helpers.APIResponse("Bad Request", http.StatusBadRequest, errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	position_pokemon2 := data_pokemon1.Position + 1
	data_pokemon2, err := repo_battlepokemon.GetPokemonBattleByPositionBattleId(position_pokemon2, data_pokemon1.BattleId)
	if err != nil {
		log.Error(err)
		errorMessage := gin.H{"error message": "Something went wrong"}
		response := helpers.APIResponse("Internal Server Error", http.StatusInternalServerError, errorMessage)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	update_data_poke1 := map[string]interface{}{
		"position": data_pokemon2.Position,
		"score":data_pokemon2.Score,
	}
	update_data_poke2 := map[string]interface{}{
		"position": data_pokemon1.Position,
		"score":data_pokemon1.Score,
	}
	err = repo_battlepokemon.UpdatePosition(data_pokemon1.ID, update_data_poke1)
	if err != nil {
		log.Error(err)
		errorMessage := gin.H{"error message": "Something went wrong"}
		response := helpers.APIResponse("Internal Server Error", http.StatusInternalServerError, errorMessage)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	err = repo_battlepokemon.UpdatePosition(data_pokemon2.ID, update_data_poke2)
	if err != nil {
		log.Error(err)
		errorMessage := gin.H{"error message": "Something went wrong"}
		response := helpers.APIResponse("Internal Server Error", http.StatusInternalServerError, errorMessage)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response := helpers.APIResponse("Success", http.StatusOK,gin.H{"message":"Success annulled pokemon position"})
	c.JSON(http.StatusOK, response)
	return
}