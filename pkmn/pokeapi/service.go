package pokeapi

import (
	"context"
	"fmt"
	"math"

	"github.com/sanka047/pokedex-go/pkmn"
)

type PokeAPIService struct {
	api IPokeAPI
}

func NewPokeAPIService(api IPokeAPI) *PokeAPIService {
	return &PokeAPIService{api: api}
}

func (s *PokeAPIService) GetPokemon(
	ctx context.Context,
	vg_id pkmn.Identifier,
	pk_id pkmn.Identifier,
) (pkmn.Pokemon, error) {
	encoded_pk_id, err := pk_id.Encode()
	if err != nil {
		return pkmn.Pokemon{}, err
	}
	encoded_vg_id, err := vg_id.Encode()
	if err != nil {
		return pkmn.Pokemon{}, err
	}

	// TODO: these should be parallelized (using a waitgroup?)
	pk, err := s.api.GetPokemon(ctx, encoded_pk_id)
	if err != nil {
		return pkmn.Pokemon{}, err
	}
	vg, err := s.api.GetVersionGroup(ctx, encoded_vg_id)
	if err != nil {
		return pkmn.Pokemon{}, err
	}

	// Filter values by the provided version group
	pk.Types = determinePokeTypesForGen(vg.Generation.ID, pk.Types, pk.PastTypes)
	pk.Moves = determinePokeMovesForVG(vg, pk.Moves)

	return exportPokemon(vg, pk)
}

func determinePokeTypesForGen(gen_id apiID, types []pokeType, past_types []pokeTypePast) []pokeType {
	// NOTE: The past_types slice contains entries for each generation _before_ the pokemon types
	// were updated.
	//
	// For example, in Gen 6, the Fairy type was added and various existing pokemon had their types
	// updated. A past_type entry is added for _gen 5_ with the previous types as it was the _last_
	// known generation with those types. This means that gen 5 _and below_ all share the old type
	// and the types were updated in gen 6.
	//
	// Given that Pokemon don't normally change types across generations (barring when Fairy was
	// released), I'd expect that past_types is either empty or has just a single entry.
	min_gen := apiID(math.MaxUint)
	for _, past_type := range past_types {
		g := past_type.Generation.ID
		// If the past type's generation is _between_ the requested generation and the min past type
		// generation, then override the types and update min generation.
		if gen_id <= g && g < min_gen {
			min_gen = g
			types = past_type.Types
		}
	}
	return types
}

// NOTE: This purely exists because I _can_. Apparently, this is not idiomatic in Go. I'm going to
// keep this anyway so I can play around with generics. - sanka047
//
// This will create a new slice, so don't use this if that's undesired overhead. Once iterators are
// available, I'd like to try out using them here.
func ffilter[T any](xs []T, f func(T) bool) []T {
	// NOTE: We're guaranteed to return _at least_ the same number of items, so make sure to
	// allocate enough space upfront, so we aren't required to resize the slice later. Yes, this
	// creates some wasted memory.
	ret := make([]T, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			ret = append(ret, x)
		}
	}
	return ret
}

func determinePokeMovesForVG(vg versionGroup, moves []pokeMove) []pokeMove {
	new_moves := make([]pokeMove, 0, len(moves))
	for _, mv := range moves {
		// filter out all details that don't correspond to the requested version
		mv.VersionGroupDetails = ffilter[pokeMoveVersion](
			mv.VersionGroupDetails,
			func(pmv pokeMoveVersion) bool { return pmv.VersionGroup.ID == vg.ID },
		)
		// if there are no details left, we trim out the move entirely as it wasn't learnable for
		// the requested version group
		if len(mv.VersionGroupDetails) > 0 {
			new_moves = append(new_moves, mv)
		}
	}
	return new_moves
}

// NOTE: This purely exists because I _can_. Apparently, this is not idiomatic in Go. I'm going to
// keep this anyway so I can play around with generics. - sanka047
//
// This will create a new slice, so don't use this if that's undesired overhead. Once iterators are
// available, I'd like to try out using them here.
func fmap[T1, T2 any](xs []T1, f func(T1) T2) []T2 {
	// NOTE: We're guaranteed to return the same number of items, so make sure to allocate enough
	// space upfront, so we aren't required to resize the slice later.
	ret := make([]T2, 0, len(xs))
	for _, x := range xs {
		ret = append(ret, f(x))
	}
	return ret
}

// Convert the stats slice received from pokeAPI into the exported container type.
//
// TODO: Figure something out. This sucks, but I can't be bothered to figure out how to deal with
// this sanely for now. I _could_ use reflection, but I'm curious if there's a way around doing
// that. It's also likely that I could (and should) just push this straight into the deserialization
// logic.
func convertPokeStats(stats []pokeStat) (pkmn.Stats, error) {
	var res pkmn.Stats
	for _, stat := range stats {
		switch stat.Stat {
		case stat_HP:
			res.HP = pkmn.Stat(stat.BaseStat)
		case stat_Atk:
			res.Atk = pkmn.Stat(stat.BaseStat)
		case stat_Def:
			res.Def = pkmn.Stat(stat.BaseStat)
		case stat_SpAtk:
			res.SpAtk = pkmn.Stat(stat.BaseStat)
		case stat_SpDef:
			res.SpDef = pkmn.Stat(stat.BaseStat)
		case stat_Spd:
			res.Spd = pkmn.Stat(stat.BaseStat)
		default:
			return res, fmt.Errorf("Undefined stat: %v", stat)
		}
	}
	return res, nil
}

// Takes the internal PokeAPI representation of a Pokemon and converts it to the modeled
// representation that we export from the module.
//
// NOTE: This could be cleaned up, but, meh, everything it's doing seems clear enough.
func exportPokemon(vg versionGroup, internal pokemon) (pkmn.Pokemon, error) {
	res := pkmn.Pokemon{
		VersionGroup: pkmn.VersionGroupID(vg.ID),
		ID:           uint(internal.ID),
		Name:         internal.Name,
	}

	stats, err := convertPokeStats(internal.Stats)
	if err != nil {
		return res, err
	}
	res.Stats = stats
	res.Abilities = fmap[pokeAbility, pkmn.PokemonAbility](
		internal.Abilities,
		func(pa pokeAbility) pkmn.PokemonAbility {
			return pkmn.PokemonAbility{
				ID:       uint(pa.Ability.ID),
				Name:     pa.Ability.Name,
				Slot:     uint(pa.Slot),
				IsHidden: pa.IsHidden,
			}
		},
	)
	res.Types = fmap[pokeType, pkmn.PokemonType](
		internal.Types,
		func(pt pokeType) pkmn.PokemonType {
			return pkmn.PokemonType{
				ID:   uint(pt.Type.ID),
				Name: pt.Type.Name,
				Slot: uint(pt.Slot),
			}
		},
	)
	res.Moves = fmap[pokeMove, pkmn.PokemonMove](
		internal.Moves,
		func(pm pokeMove) pkmn.PokemonMove {
			return pkmn.PokemonMove{
				ID:   uint(pm.Move.ID),
				Name: pm.Move.Name,
				LearnOptions: fmap[pokeMoveVersion, pkmn.PokeMoveLearnOption](
					pm.VersionGroupDetails,
					func(d pokeMoveVersion) pkmn.PokeMoveLearnOption {
						return pkmn.PokeMoveLearnOption{
							LearnMethod:    pkmn.LearnMethod(d.MoveLearnMethod.ID),
							LevelLearnedAt: pkmn.Level(d.LevelLearnedAt),
						}
					},
				),
			}
		},
	)
	return res, nil
}
