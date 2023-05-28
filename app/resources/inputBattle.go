package resources

type InputBattle struct {
	BattleName string   `json:"battle_name" binding:"required"`
	Pokemons   []string `json:"pokemons" binding:"required"`
}

type InputBattleManual struct {
	BattleName string   `json:"battle_name" binding:"required"`
	Position   []string `json:"position" binding:"required"`
}