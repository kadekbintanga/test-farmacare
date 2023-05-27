package helpers

import(
	"strings"
	"fmt"
)

func FormatValidationErrorInput(err error) string{
	fmt.Println("Input Validation Error")
	if strings.Contains(err.Error(), "BattleName"){
		return "battle_name is required"
	}else if strings.Contains(err.Error(), "Pokemons"){
		return "pokemons is required"
	}else if strings.Contains(err.Error(), "Position"){
		return "position is required"
	}else if strings.Contains(err.Error(), "unmarshal"){
		if  strings.Contains(err.Error(), "InputBattle.battle_name"){
			return "battle_name must has string type"
		}else if strings.Contains(err.Error(), "InputBattle.pokemons"){
			return "pokemons must an array with string value"
		}else if strings.Contains(err.Error(), "InputBattleManual.position"){
			return "position must an array with string value"
		}else{
			return "Something went wrong"
		}
	}else{
		return "Something went wrong"
	}
}
