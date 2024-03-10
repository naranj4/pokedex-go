package pokeapi

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/sanka047/pokedex-go/testing/assert"
)

func TestPokemonSerialization(t *testing.T) {
	dat, err := os.ReadFile("testdata/pokemon-response.json")
	assert.Ok(t, err)

	var pk Pokemon
	err = json.Unmarshal(dat, &pk)
	assert.Ok(t, err)

	exp, err := os.ReadFile("testdata/pokemon-parsed.json")
	assert.Ok(t, err)

	act, err := json.MarshalIndent(pk, "", "  ")
	assert.Ok(t, err)
	assert.Equals(t, string(act), strings.Trim(string(exp), "\n"))
}
