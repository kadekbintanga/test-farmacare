package handlers

import (
	"test-farmacare/app/helpers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"os"
)

type PokemonHandler struct {

}

func NewPokemonHandler() *PokemonHandler{
	return &PokemonHandler{}
}


func(h *PokemonHandler) HealthPokemon(c *gin.Context){
	log.Info("Health Pokemon Handler is Hitted")
	response := helpers.APIResponse("Success", http.StatusOK, gin.H{"message": "Success hit Health Pokemon Handler"})
	c.JSON(http.StatusOK, response)
	return
}


func(h *PokemonHandler) GetPokemonList(c *gin.Context){
	page_str := c.DefaultQuery("page", "10")
	page_int,err := strconv.Atoi(page_str)
	if err != nil {
		log.Error("Page got : ",err)
		errorMessage := gin.H{"error message": "Page param must be a number, Please check your page params!"}
		response := helpers.APIResponse("Bad Request", http.StatusBadRequest, errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	limit_str := c.DefaultQuery("limit", "10")
	limit_int,err := strconv.Atoi(limit_str)
	if err != nil {
		log.Error("Limit got : ",err)
		errorMessage := gin.H{"error message": "Limit param must be a number, Please check your limit params!"}
		response := helpers.APIResponse("Bad Request", http.StatusBadRequest, errorMessage)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	offset_int := (page_int - 1)*limit_int
	offset_str := strconv.Itoa(offset_int)
	log.Info(offset_str)
	request_pokeapi, err := http.NewRequest("GET", os.Getenv("API_POKEMON_URL"), nil)
	if err != nil {
		log.Error("Error hhtp new request : ",err)
		errorMessage := gin.H{"error message": "Something went wrong"}
		response := helpers.APIResponse("Internal Server Error", http.StatusInternalServerError, errorMessage)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	params := request_pokeapi.URL.Query()
	params.Add("offset",offset_str)
	params.Add("limit", limit_str)
	request_pokeapi.URL.RawQuery = params.Encode()

	response_request_pokeapi, err := http.DefaultClient.Do(request_pokeapi)
	if err != nil {
		log.Error("Error when hit pokeapi : ",err)
		errorMessage := gin.H{"error message": "Something went wrong"}
		response := helpers.APIResponse("Internal Server Error", http.StatusInternalServerError, errorMessage)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response_data,_ := ioutil.ReadAll(response_request_pokeapi.Body)
	log.Info("Response from PokeApi : ", string(response_data))

	var data_map map[string]interface{}
	if err := json.Unmarshal([]byte(string(response_data)), &data_map);err !=nil{
		log.Error("Error unmarshal data json : ", string(err.Error()))
	}

	res := data_map["results"].([]interface{})
	total := int(data_map["count"].(float64))
	response := helpers.APIResponse2("Success", http.StatusOK,page_int,limit_int,total,res)
	c.JSON(http.StatusOK, response)
	return

}

