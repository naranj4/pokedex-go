package pkmn

type VersionGroupID uint8
type Stat uint8
type Level uint8

type Pokemon struct {
	VersionGroup VersionGroupID `json:"version_group"`

	ID   uint   `json:"id"`
	Name string `json:"name"`

	Abilities []PokemonAbility `json:"abilities"`
	Types     []PokemonType    `json:"types"`
	Stats     Stats            `json:"stats"`

	// TODO: break this out separately so this isn't returned alongside base pokemon metadata
	Moves []PokemonMove `json:"moves"`
}

type Stats struct {
	HP    Stat `json:"hp"`
	Atk   Stat `json:"attack"`
	Def   Stat `json:"defense"`
	SpAtk Stat `json:"special_attack"`
	SpDef Stat `json:"special_defense"`
	Spd   Stat `json:"speed"`
}

type PokemonAbility struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`

	Slot     uint `json:"slot"`
	IsHidden bool `json:"is_hidden"`
}

type PokemonType struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`

	Slot uint `json:"slot"`
}

type PokemonMove struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`

	LearnOptions []PokeMoveLearnOption `json:"learn_options"`
}

type PokeMoveLearnOption struct {
	LearnMethod    LearnMethod `json:"learn_method"`
	LevelLearnedAt Level       `json:"level_learned_at"` // only meaningful for level-up method
}
