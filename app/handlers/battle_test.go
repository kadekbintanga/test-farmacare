package handlers

import (
	"net/http"
	"net/http/httptest"
	"test-farmacare/app/models"
	"testing"
	"time"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


type BattleRepository interface {
	GetListBattle(start, end string, limit, offset int) ([]models.Battle, int, error)
	GetBattleById(id int) (models.Battle, error)
}



type MockBattleRepository struct {
	mock.Mock
}


func (m *MockBattleRepository) GetListBattle(start_date string, end_date string, limit, offset int) ([]models.Battle, int64, error) {
	args := m.Called(start_date, end_date, limit, offset)
	return args.Get(0).([]models.Battle),1, nil
}

func (m *MockBattleRepository) GetBattleById(id uint) (models.Battle, error) {
	args := m.Called(id)
	return args.Get(0).(models.Battle), args.Error(2)
}

func(m *MockBattleRepository) SaveBattle(Battle models.Battle)(models.Battle, error){
	args := m.Called(Battle)
	return args.Get(0).(models.Battle), nil
}


func TestGetListBattle_Success(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest(http.MethodGet, "/battles", nil)
	queryParams := c.Request.URL.Query()
	queryParams.Set("start_date", "2023-05-28 14:46:00")
	queryParams.Set("end_date", "2023-05-30 14:46:00")
	queryParams.Set("page", "1")
	queryParams.Set("limit", "10")
	c.Request.URL.RawQuery = queryParams.Encode()

	h := &BattleHandler{
		repo_battle: &MockBattleRepository{},
	}

	mockBattles := []models.Battle{
		{
			ID:         1,
			UUID:       uuid.New(),
			BattleName: "Battle 1",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			BattlePokemon: []models.BattlePokemon{
				{
					ID:           1,
					UUID:         uuid.New(),
					PokemonName:  "Pokemon 1",
					Position:     1,
					Score:        4,
				},
			},
		},
	}

	mockTotal := int(1)
	mockRepo := h.repo_battle.(*MockBattleRepository)
	mockRepo.On("GetListBattle", "2023-05-28 14:46:00", "2023-05-30 14:46:00", 10, 0).Return(mockBattles,mockTotal,nil)

	h.GetListBattle(c)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(),`"message":"Success","code":200`)
	assert.Contains(t, w.Body.String(), `"page":1,"limit":10,"total":1`)
	assert.Contains(t, w.Body.String(), `"battle_name":"Battle 1"`)
	assert.Contains(t, w.Body.String(), `"id":1,"pokemon_name":"Pokemon 1","position":1,"score":4`)
}


// func TestCreateBattleAuto_Success(t *testing.T){
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	data := map[string]interface{}{
// 		"battle_name":"Battle Poke Campion",
//    		"pokemons":[]string{"spearow", "fearow", "ekans", "pikachu","nidoqueen"},
// 	}
// 	jsonValue,_ := json.Marshal(data)

// 	c.Request = httptest.NewRequest(http.MethodPost, "/battle/auto", bytes.NewBuffer(jsonValue))
// 	c.Request.Header.Set("Content-Type", "application/json")

// 	h := &BattleHandler{
// 		repo_battle: &MockBattleRepository{},
// 	}
// 	mockBattles := models.Battle{
// 			ID:1,
// 			UUID:       uuid.New(),
// 			BattleName: "Battle Poke Campion",
// 			CreatedAt:  time.Now(),
// 			UpdatedAt:  time.Now(),
// 			BattlePokemon: []models.BattlePokemon{
// 				{
// 					ID:           1,
// 					UUID:         uuid.New(),
// 					PokemonName:  "phikachu",
// 					Position:     2,
// 					Score:        4,
// 				},
// 			},
// 		}

// 	battle := models.Battle{
// 		BattleName:"Battle Poke Campion",
// 	}
// 	mockRepo := h.repo_battle.(*MockBattleRepository)
// 	mockRepo.On("SaveBattle",battle).Return(mockBattles, nil)
// 	h.CreateBattleAuto(c)
// 	fmt.Println(w.Body)
// }
