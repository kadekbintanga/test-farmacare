package models

import(
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
)


type Battle struct{
	gorm.Model
	ID				uint			`json:"id" gorm:"primary_key"`
	UUID			uuid.UUID		`gorm:"type:uuid;default:uuid_generate_v4();index:uuid_unique;unique;"`
	BattleName		string			`json:"battle_name"`
	CreatedAt		time.Time		`json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt		time.Time		`json:"updated_at" gorm:"autoCreateTime,autoUpdateTime"`
	BattlePokemon	[]BattlePokemon	`gorm:"foreignKey:BattleId"`
}


type BattlePokemon struct{
	gorm.Model
	ID				uint			`json:"id" gorm:"primary_key"`
	UUID			uuid.UUID		`gorm:"type:uuid;default:uuid_generate_v4();index:uuid_unique;unique;"`
	BattleId		uint			`json:"battle_uuid"`
	Battle			Battle			`gorm:"foreignKey:BattleId"`
	PokemonName		string			`json:"pokemon_name"`
	Position		uint			`json:"position"`
	Score			uint			`json:"score"`
	CreatedAt		time.Time		`json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt		time.Time		`json:"updated_at" gorm:"autoCreateTime,autoUpdateTime"`
}


