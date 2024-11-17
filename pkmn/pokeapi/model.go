package pokeapi

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

type apiID uint

type pokemon struct {
	ID   apiID  `json:"id"`
	Name string `json:"name"`

	Abilities []pokeAbility  `json:"abilities"`
	Types     []pokeType     `json:"types"`
	PastTypes []pokeTypePast `json:"past_types"`
	Stats     []pokeStat     `json:"stats"`

	Moves []pokeMove `json:"moves"`
}

type pokeAbility struct {
	Ability  namedAPIResource `json:"ability"`
	Slot     uint8            `json:"slot"`
	IsHidden bool             `json:"is_hidden"`
}

type pokeType struct {
	Type namedAPIResource `json:"type"`
	Slot uint8            `json:"slot"`
}

type pokeTypePast struct {
	Generation namedAPIResource `json:"generation"`
	Types      []pokeType       `json:"types"`
}

type pokeStat struct {
	Stat         stat             `json:"-"` // derived from StatResource (don't serialize)
	StatResource namedAPIResource `json:"stat"`
	BaseStat     uint8            `json:"base_stat"`
}

// Parse out the stat into an enum (avoids having to do it later)
func (s *pokeStat) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	// original shape of the data that we'll be using for decoding
	var raw struct {
		StatResource namedAPIResource `json:"stat"`
		BaseStat     uint8            `json:"base_stat"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("Failed to unmarshal `namedAPIResource`: %w", err)
	}

	// unmarshal data into the actual struct
	*s = pokeStat{
		Stat:         stat(raw.StatResource.ID),
		StatResource: raw.StatResource,
		BaseStat:     raw.BaseStat,
	}
	return nil
}

type pokeMove struct {
	Move                namedAPIResource  `json:"move"`
	VersionGroupDetails []pokeMoveVersion `json:"version_group_details"`
}

type pokeMoveVersion struct {
	MoveLearnMethod namedAPIResource `json:"move_learn_method"`
	VersionGroup    namedAPIResource `json:"version_group"`
	LevelLearnedAt  uint8            `json:"level_learned_at"`
}

type versionGroup struct {
	ID         apiID            `json:"id"`
	Name       string           `json:"name"`
	Generation namedAPIResource `json:"generation"`
}

type namedAPIResource struct {
	ID   apiID  `json:"-"` // derived from URL (don't serialize)
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Somewhat janky deserialization logic so that we parse the ID out from the API resource URL. This
// kinda sucks, but it prevents me from needing a bunch of error handling logic all over the place
// after we deserialize (while avoiding needing to write conversion code to a separate struct).
func (r *namedAPIResource) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	// original shape of the data that we'll be using for decoding
	var raw struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("Failed to unmarshal `namedAPIResource`: %w", err)
	}

	id, err := getIDFromURL(raw.URL)
	if err != nil {
		return err
	}

	// unmarshal data into the actual struct
	*r = namedAPIResource{ID: id, Name: raw.Name, URL: raw.URL}
	return nil
}

func getIDFromURL(url string) (apiID, error) {
	// Remove any trailing slashes and grab the index of the last slash
	url = filepath.Clean(url)
	idx := strings.LastIndex(url, "/")
	if idx < 0 {
		return 0, fmt.Errorf("Malformed resource URL (%v)", url)
	}
	// The last part of the URL _should_ be the resource id, so convert it to an int
	id, err := strconv.Atoi(string(url[idx+1:]))
	if err != nil {
		return 0, fmt.Errorf("Malformed resource URL (%v), received (%w)", url, err)
	}
	return apiID(id), nil
}
