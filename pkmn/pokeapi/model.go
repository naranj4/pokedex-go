package pokeapi

type Pokemon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	Abilities []PokemonAbility  `json:"abilities"`
	Types     []PokemonType     `json:"types"`
	PastTypes []PokemonTypePast `json:"past_types"`
	Stats     []PokemonStat     `json:"stats"`

	Moves []PokemonMove `json:"moves"`
}

type PokemonAbility struct {
	Ability  NamedAPIResource `json:"ability"`
	Slot     int              `json:"slot"`
	IsHidden bool             `json:"is_hidden"`
}

type PokemonType struct {
	Type NamedAPIResource `json:"type"`
	Slot int              `json:"slot"`
}

type PokemonTypePast struct {
	Generation NamedAPIResource `json:"generation"`
	Types      []PokemonType    `json:"types"`
}

type PokemonStat struct {
	Stat     NamedAPIResource `json:"stat"`
	BaseStat int              `json:"base_stat"`
}

type PokemonMove struct {
	Move                NamedAPIResource     `json:"move"`
	VersionGroupDetails []PokemonMoveVersion `json:"version_group_details"`
}

type PokemonMoveVersion struct {
	MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
	VersionGroup    NamedAPIResource `json:"version_group"`
	LevelLearnedAt  int              `json:"level_learned_at"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
