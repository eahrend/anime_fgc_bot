package models

const StriveAllGuard = "ALL"
const StriveLowGuard = "LOW"
const StriveHighGuard = "HIGH"

type StriveCharacterData struct {
	Revision                int                                  `bson:"revision" json:"revision"`
	NormalMoves             map[string]StriveCharacterNormalMove `bson:"normal_moves" json:"normal_moves"`
	CharacterName           string                               `bson:"character_name" json:"character_name"`
	CharacterNameSimplified string                               `bson:"character_name_simplified" json:"character_name_simplified"`
	GameName                string                               `bson:"game_name" json:"game_name"`
	GameNameSimplified      string                               `bson:"game_name_simplified" json:"game_name_simplified"`
}

type StriveCharacterNormalMove struct {
	ImageURL       string   `json:"image_url" bson:"image_url"`
	HitboxURL      string   `json:"hitbox_url" bson:"hitbox_url"`
	Damage         []string `json:"damage" bson:"damage"`
	ChargeDamage   []string `json:"charge_damage,omitempty" bson:"charge_damage,omitempty"`
	Guard          string   `json:"guard" bson:"guard"`
	Startup        []string `json:"startup" bson:"startup"`
	ChargeStartup  string   `json:"charge_startup,omitempty" bson:"charge_startup,omitempty"`
	Active         []string `json:"active" bson:"active"`
	Recovery       string   `json:"recovery" bson:"recovery"`
	OnBlock        string   `json:"on_block" bson:"on_block"`
	ChargeOnBlock  string   `json:"charge_on_block,omitempty" bson:"charge_on_block,omitempty"`
	GatlingOptions []string `json:"gatling_options,omitempty" bson:"gatling_options,omitempty"`
}
