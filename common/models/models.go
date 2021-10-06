package models

const StriveAllGuard = "ALL"
const StriveLowGuard = "LOW"
const StriveHighGuard = "HIGH"

type StriveCharacterData struct {
	Revision                int                                   `bson:"revision" json:"revision"`
	NormalMoves             map[string]StriveCharacterNormalMove  `bson:"normal_moves" json:"normal_moves"`
	Throws                  map[string]StriveCharacterThrowMove   `json:"throws" bson:"throws"`
	SpecialMoves            map[string]StriveCharacterSpecialMove `json:"special_moves" bson:"special_moves"`
	SuperMoves              map[string]StriveCharacterSuperMove   `json:"super_moves" bson:"super_moves"`
	CharacterName           string                                `bson:"character_name" json:"character_name"`
	CharacterNameSimplified string                                `bson:"character_name_simplified" json:"character_name_simplified"`
	GameName                string                                `bson:"game_name" json:"game_name"`
	GameNameSimplified      string                                `bson:"game_name_simplified" json:"game_name_simplified"`
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

type StriveCharacterThrowMove struct {
	ImageURL  string   `json:"image_url" bson:"image_url"`
	HitboxURL string   `json:"hitbox_url" bson:"hitbox_url"`
	Damage    []string `json:"damage" bson:"damage"`
	Recovery  string   `json:"recovery" bson:"recovery"`
	Startup   []string `json:"startup" bson:"startup"`
}

type StriveCharacterSuperMove struct {
	ImageURL        string   `json:"image_url" bson:"image_url"`
	HitboxURL       string   `json:"hitbox_url" bson:"hitbox_url"`
	Damage          []string `json:"damage" bson:"damage"`
	OnBlock         string   `json:"on_block,omitempty" bson:"on_block,omitempty"`
	Active          []string `json:"active,omitempty" bson:"active,omitempty"`
	Recovery        string   `json:"recovery" bson:"recovery"`
	Startup         []string `json:"startup" bson:"startup"`
	Name            string   `json:"name" bson:"name"`
	NameSimplified  string   `json:"name_simplified" bson:"name_simplified"`
	Notes           string   `json:"notes,omitempty" bson:"notes,omitempty"`
	Invulnerability string   `json:"invulnerability,omitempty" bson:"invulnerability,omitempty"`
	MeterCost       string   `json:"meter_cost" bson:"meter_cost"`
	MinimumDamage   string   `json:"minimum_damage,omitempty" bson:"minimum_damage,omitempty"`
}

type StriveCharacterSpecialMove struct {
	ImageURL       string   `json:"image_url" bson:"image_url"`
	HitboxURL      string   `json:"hitbox_url" bson:"hitbox_url"`
	Damage         []string `json:"damage" bson:"damage"`
	OnBlock        string   `json:"on_block,omitempty" bson:"on_block,omitempty"`
	Active         []string `json:"active,omitempty" bson:"active,omitempty"`
	Recovery       string   `json:"recovery" bson:"recovery"`
	Startup        []string `json:"startup" bson:"startup"`
	Name           string   `json:"name" bson:"name"`
	NameSimplified string   `json:"name_simplified" bson:"name_simplified"`
	Notes          string   `json:"notes,omitempty" bson:"notes,omitempty"`
}
