package models

const StriveAllGuard = "ALL"
const StriveLowGuard = "LOW"
const StriveHighGuard = "HIGH"

type StriveCharacterData struct {
	Revision    int                                  `json:"revision"`
	NormalMoves map[string]StriveCharacterNormalMove `json:"normal_moves"`
}

type StriveCharacterNormalMove struct {
	ImageURL       string   `json:"image_url"`
	HitboxURL      string   `json:"hitbox_url"`
	Damage         []string `json:"damage"`
	ChargeDamage   string   `json:"charge_damage,omitempty"`
	Guard          string   `json:"guard"`
	Startup        string   `json:"startup"`
	Active         string   `json:"active"`
	Recovery       string   `json:"recovery"`
	OnBlock        string   `json:"on_block"`
	GatlingOptions []string `json:"gatling_options"`
}
