package pkmn

import (
	"context"
	"fmt"
)

type Identifier interface {
	// TODO: remove error return value (errors should be surfaced on construction rather than
	// encoding)
	Encode() (string, error)
}

type nameval string

func (n nameval) Encode() (string, error) {
	return string(n), nil
}

func Name(name string) Identifier {
	return nameval(name)
}

type idval uint

func (i idval) Encode() (string, error) {
	return fmt.Sprint(i), nil
}

func Id(id uint) Identifier {
	return idval(id)
}

type IPokeService interface {
	// TODO: need to pass additional context like the selected version group id
	GetPokemon(ctx context.Context, vg_id Identifier, pk_id Identifier) (Pokemon, error)
}
