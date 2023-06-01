package handlers

import (
	"net/http"
	"net/http/httptest"
	"test-farmacare/app/models"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"fmt"
)


type BattlePokemonRepository interface {
	SaveBattlePokemon(BattlePokemon []models.BattlePokemon)([]models.BattlePokemon,error)
	GetTotalScore()([]map[string]interface{}, error)
	GetPokemonBattleByUuuid(uuid string)(models.BattlePokemon, error)
	GetPokemonBattleByPositionBattleId(position uint, battle_id uint)(models.BattlePokemon, error)
	UpdatePosition(id uint, update_data map[string]interface{})(error)
}

type MockBattlePokemonRepository struct {
	mock.Mock
}

func (m *MockBattlePokemonRepository) GetTotalScore()([]map[string]interface{}, error){
	args := m.Called()
	return args.Get(0).([]map[string]interface{}), nil
}

func (m *MockBattlePokemonRepository) SaveBattlePokemon(BattlePokemon []models.BattlePokemon)([]models.BattlePokemon,error){
	args := m.Called(BattlePokemon)
	return args.Get(0).([]models.BattlePokemon), nil
}

func (m *MockBattlePokemonRepository) GetPokemonBattleByPositionBattleId(position uint, battle_id uint)(models.BattlePokemon, error){
	args := m.Called(position, battle_id)
	return args.Get(0).(models.BattlePokemon), nil
}

func (m *MockBattlePokemonRepository) UpdatePosition(id uint, update_data map[string]interface{})(error){
	args := m.Called(id, update_data)
	return args.Error(0)
}

func (m *MockBattlePokemonRepository)GetPokemonBattleByUuuid(uuid string)(models.BattlePokemon, error){
	args := m.Called(uuid)
	return args.Get(0).(models.BattlePokemon), nil
}

func TestGetTotalScore_Success(t *testing.T){
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/battlepokemon/score", nil)
	
	h := &BattlePokemonHandler{
		repo_battlepokemon: &MockBattlePokemonRepository{},
	}
	response_data_rep := []map[string]interface{}{
		{
			"pokemon_name":"pikachu",
			"total_score" : 30,
		},
	}
	mockRepo := h.repo_battlepokemon.(*MockBattlePokemonRepository)
	mockRepo.On("GetTotalScore").Return(response_data_rep, nil)
	h.GetTotalScore(c)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(),`"message":"Success","code":200`)
	assert.Contains(t, w.Body.String(), `"data":[{"pokemon_name":"pikachu","total_score":30}]}`)
	fmt.Println(w.Body.String())
}

func TestAnnulledPosition_Success(t *testing.T){
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPut, "/battlepokemon/annulled", nil)
	queryParams := c.Request.URL.Query()
	queryParams.Set("uuid_pokemon", "uuui-pokemon-testtt")
	c.Request.URL.RawQuery = queryParams.Encode()
	
	h := &BattlePokemonHandler{
		repo_battlepokemon: &MockBattlePokemonRepository{},
	}
	DataGetByUUid := models.BattlePokemon{
		ID:1,
		PokemonName: "pikachu",
		Position: 1,
		Score: 5,
	}
	DataGetByPossition := models.BattlePokemon{
		ID:2,
		PokemonName: "ditto",
		Position: 2,
		Score: 4,
	}
	update_data_poke1 := map[string]interface{}{
		"position": uint(2),
		"score": uint(4),
	}
	update_data_poke2 := map[string]interface{}{
		"position": uint(1),
		"score": uint(5),
	}

	mockRepo := h.repo_battlepokemon.(*MockBattlePokemonRepository)
	mockRepo.On("GetPokemonBattleByUuuid", "uuui-pokemon-testtt").Return(DataGetByUUid, nil)
	mockRepo.On("GetPokemonBattleByPositionBattleId", uint(2), uint(0)).Return(DataGetByPossition, nil)
	mockRepo.On("UpdatePosition",uint(1),update_data_poke1). Return(nil)
	mockRepo.On("UpdatePosition",uint(2),update_data_poke2). Return(nil)
	h.AnnulledPosition(c)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(),`"message":"Success","code":200`)
	assert.Contains(t, w.Body.String(), `"data":{"message":"Success annulled pokemon position"}`)
}

