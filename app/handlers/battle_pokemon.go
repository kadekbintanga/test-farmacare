package handlers

import (
	"test-farmacare/app/helpers"
	// "test-farmacare/app/models"
	"test-farmacare/app/repositories"
	// "test-farmacare/app/resources"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	// "math/rand"
	"net/http"
	// "strconv"
	// "time"
	// "os"
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