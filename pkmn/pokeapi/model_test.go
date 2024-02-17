package pokeapi

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func TestPokemonSerialization(t *testing.T) {
	dat, err := os.ReadFile("testdata/pokemon-response.json")
	ok(t, err)

	var pk Pokemon
	err = json.Unmarshal(dat, &pk)
	ok(t, err)

	exp, err := os.ReadFile("testdata/pokemon-parsed.json")
	ok(t, err)

	act, err := json.MarshalIndent(pk, "", "  ")
	ok(t, err)
	equals(t, string(act), strings.Trim(string(exp), "\n"))
}
