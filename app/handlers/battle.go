package handlers

import (
	"test-farmacare/app/helpers"
	"test-farmacare/app/models"
	"test-farmacare/app/repositories"
	"test-farmacare/app/resources"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	// "encoding/json"
	// "io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"os"
)

type BattleHandler struct {
	repo_battle repositories.BattleRepository
	repo_battlepokemon repositories.BattlePokemonRepository
}

func NewBattleHandler() *BattleHandler{
	return &BattleHandler{
		repositories.NewBattleRepository(),
		repositories.NewBattlePokemonRepository(),
	}
}



func(h *BattleHandler) HealthBattle(c *gin.Context){
	log.Info("Health Battle Handler is Hitted")
	response := helpers.APIResponse("Success", http.StatusOK, gin.H{"message": "Success hit Health Battle Handler"})
	c.JSON(http.StatusOK, response)
	return
}

func(h *BattleHandler) CreateBattleAuto(c *gin.Context){
	var req resources.InputBattle
	err := c.ShouldBind(&req)
	if err != nil {
		log.Error(err)
		errors := helpers.FormatValidationErrorInput(err)
		errorMessage := gin.H{"error message": errors}
		response := helpers.APIResponse("Bad Request", http.StatusBadRequest, errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	check_duplicate := h.CheckDuplicate(req.Pokemons)
	if check_duplicate == true{
		log.Error("Pokemon name is duplicated")
		errorMessage := gin.H{"error message": "Pokemons name is duplicated"}
		response := helpers.APIResponse("Bad Request", http.StatusBadRequest, errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	Battle := models.Battle{
		BattleName:req.BattleName,
	}
	repo_battle := h.repo_battle
	repo_battlepokemon := h.repo_battlepokemon
	save_battle, err := repo_battle.SaveBattle(Battle)
	if err != nil {
		log.Error(err)
		errorMessage := gin.H{"error message": "Something went wrong"}
		response := helpers.APIResponse("Internal Server Error", http.StatusInternalServerError, errorMessage)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	scores := []int{1,2,3,4,5}
	var BattlePokemon []models.BattlePokemon
	if len(req.Pokemons) != 5{
		log.Error("Pokemon input is ", len(req.Pokemons))
		errorMessage := gin.H{"error message": "Pokemons input must contain only 5 pokemons"}
		response := helpers.APIResponse("Bad Request", http.StatusBadRequest, errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
		
	}

	for _,v := range req.Pokemons{
		if v == ""{
			log.Error("Pokemon input is null or ''")
			errorMessage := gin.H{"error message": "Please check your pokemons input"}
			response := helpers.APIResponse("Bad Request", http.StatusBadRequest, errorMessage)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		check_pokemon := h.CheckPokemon(v)
		if check_pokemon == false{
			log.Error("Check pokemon is false")
			errorMessage := gin.H{"error message": "Pokemon "+v+" is not found"}
			response := helpers.APIResponse("Bad Request", http.StatusBadRequest, errorMessage)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		index_score_rand := rand.Intn(len(scores))
		score := scores[index_score_rand]
		position := 5-(score-1)
		scores = append(scores[:index_score_rand], scores[index_score_rand+1:]...)

		data_pokemon := models.BattlePokemon{
			BattleId:save_battle.ID,
			PokemonName:v,
			Position:uint(position),
			Score:uint(score),
		}

		BattlePokemon = append(BattlePokemon,data_pokemon)
	}

	save_pokemon, err := repo_battlepokemon.SaveBattlePokemon(BattlePokemon)
	if err != nil {
		log.Error(err)
		errorMessage := gin.H{"error message": "Something went wrong"}
		response := helpers.APIResponse("Internal Server Error", http.StatusInternalServerError, errorMessage)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	var data_poke []map[string]interface{}
	for _,value := range save_pokemon{
		d := gin.H{
			"pokemon_name":value.PokemonName,
			"position":value.Position,
			"score":value.Score,
		}
		data_poke = append(data_poke,d)
	}

	res := gin.H{
		"uuid":save_battle.UUID,
		"battle_name":save_battle.BattleName,
		"pokemons":data_poke,
	}

	response := helpers.APIResponse("Success", http.StatusOK, res)
	c.JSON(http.StatusOK, response)
	return
}

func(h *BattleHandler) CreateBattleManual(c *gin.Context){
	var req resources.InputBattleManual
	err := c.ShouldBind(&req)
	if err != nil {
		log.Error(err)
		errors := helpers.FormatValidationErrorInput(err)
		errorMessage := gin.H{"error message": errors}
		response := helpers.APIResponse("Bad Request", http.StatusBadRequest, errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	check_duplicate := h.CheckDuplicate(req.Position)
	if check_duplicate == true{
		log.Error("Pokemon name is duplicated")
		errorMessage := gin.H{"error message": "Pokemons name is duplicated"}
		response := helpers.APIResponse("Bad Request", http.StatusBadRequest, errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	Battle := models.Battle{
		BattleName:req.BattleName,
	}
	repo_battle := h.repo_battle
	repo_battlepokemon := h.repo_battlepokemon
	save_battle, err := repo_battle.SaveBattle(Battle)
	if err != nil {
		log.Error(err)
		errorMessage := gin.H{"error message": "Something went wrong"}
		response := helpers.APIResponse("Internal Server Error", http.StatusInternalServerError, errorMessage)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	score := 5
	var BattlePokemon []models.BattlePokemon
	for i,v := range req.Position{
		if v == ""{
			log.Error("Pokemon input is null or ''")
			errorMessage := gin.H{"error message": "Please check your pokemons input"}
			response := helpers.APIResponse("Bad Request", http.StatusBadRequest, errorMessage)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		check_pokemon := h.CheckPokemon(v)
		if check_pokemon == false{
			log.Error("Check pokemon is false")
			errorMessage := gin.H{"error message": "Pokemon "+v+" is not found"}
			response := helpers.APIResponse("Bad Request", http.StatusBadRequest, errorMessage)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		data_pokemon := models.BattlePokemon{
			BattleId:save_battle.ID,
			PokemonName:v,
			Position:uint(i+1),
			Score:uint(score),
		}
		BattlePokemon = append(BattlePokemon,data_pokemon)
		score = score - 1
	}
	save_pokemon, err := repo_battlepokemon.SaveBattlePokemon(BattlePokemon)
	if err != nil {
		log.Error(err)
		errorMessage := gin.H{"error message": "Something went wrong"}
		response := helpers.APIResponse("Internal Server Error", http.StatusInternalServerError, errorMessage)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	var data_poke []map[string]interface{}
	for _,value := range save_pokemon{
		d := gin.H{
			"pokemon_name":value.PokemonName,
			"position":value.Position,
			"score":value.Score,
		}
		data_poke = append(data_poke,d)
	}

	res := gin.H{
		"uuid":save_battle.UUID,
		"battle_name":save_battle.BattleName,
		"pokemons":data_poke,
	}

	response := helpers.APIResponse("Success", http.StatusOK, res)
	c.JSON(http.StatusOK, response)
	return
}

func(h *BattleHandler) GetListBattle(c *gin.Context){
	start_date := c.DefaultQuery("start_date","")
	if start_date != ""{
		check_date,_ := time.Parse("2006-01-02 15:04:05", start_date)
		log.Info(check_date)
		if check_date.String() == "0001-01-01 00:00:00 +0000 UTC"{
			log.Error("Checkdate : ",start_date)
			errorMessage := gin.H{"error message": "Invalid format start_date"}
			response := helpers.APIResponse("Bad Request", http.StatusBadRequest, errorMessage)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}
	end_date := c.DefaultQuery("end_date","")
	if end_date!= ""{
		check_date,_ := time.Parse("2006-01-02 15:04:05", end_date)
		if check_date.String() == "0001-01-01 00:00:00 +0000 UTC"{
			log.Error("Checkdate : ",end_date)
			errorMessage := gin.H{"error message": "Invalid format end_date"}
			response := helpers.APIResponse("Bad Request", http.StatusBadRequest, errorMessage)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}
	page, err := strconv.Atoi(c.DefaultQuery("page","1"))
	if err != nil {
		log.Error("Page got : ",err)
		errorMessage := gin.H{"error message": "Page param must be a number, Please check your page params!"}
		response := helpers.APIResponse("Bad Request", http.StatusBadRequest, errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit","10"))
	if err != nil {
		log.Error("Page got : ",err)
		errorMessage := gin.H{"error message": "Limit param must be a number, Please check your page params!"}
		response := helpers.APIResponse("Bad Request", http.StatusBadRequest, errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	offset := (page-1)*limit
	repo_battle := h.repo_battle
	res, total, err := repo_battle.GetListBattle(start_date, end_date, limit, offset)

	var data []map[string]interface{}
	
	for _,v := range res{
		var data_batle []map[string]interface{}
		for _,v1 := range v.BattlePokemon{
			d1 := gin.H{
				"id":v1.ID,
				"uuid":v1.UUID,
				"pokemon_name":v1.PokemonName,
				"position":v1.Position,
				"score":v1.Score,
			}
			data_batle = append(data_batle,d1)
		}
		d := gin.H{
			"id":v.ID,
			"uuid":v.UUID,
			"battle_name":v.BattleName,
			"created_at":v.CreatedAt,
			"updated_at":v.UpdatedAt,
			"battle_pokemon":data_batle,
		}
		data = append(data,d)
	}

	response := helpers.APIResponse2("Success", http.StatusOK,page,limit,int(total),data)
	c.JSON(http.StatusOK, response)
	return
}




func(h *BattleHandler) CheckPokemon(pokemon_name string)bool{
	request_pokeapi, err := http.NewRequest("GET", os.Getenv("API_POKEMON_URL")+"/"+pokemon_name, nil)
	if err != nil {
		log.Error("Error when hit pokeapi : ",err)
	}
	response_request_pokeapi, err := http.DefaultClient.Do(request_pokeapi)
	if err != nil {
		log.Error("Error when hit pokeapi : ",err)
	}
	status_code := response_request_pokeapi.Status
	if status_code == "200 OK"{
		return true
	}
	return false
}

func(h *BattleHandler) CheckDuplicate(pokemon_names []string)bool{
	var sample []string
	for _,v := range pokemon_names{
		for _,j := range sample{
			if j == v{
				return true
			}
		}
		sample = append(sample,v)
	}
	return false
}