## Notes
### Pokemon
Pokemon DB: Store historical information? Lookup pokemon by the latest generation (before current)

Identifiers:
- [PK] **required** `vg_id`: uint
- [PK_1] `id`: uint
- [PK_2] `name`: pokemon name

Attributes
- stats (hp, attack, defense, special-attack, special-defense, speed): uint8
- `type_1..2`: struct
    - `id`: uint
    - `name`: string
- `ability_1..3`: struct
    - `id`: uint
    - `name`: string

### Pokemon-Type
NOTE: This needs to store historical types (by generation)

Identifiers:
- [PK] **required** `vg_id`: uint

  Should there be a latest sentinel to help with lookups and correctness as new games are released?
  A: Probably not, since you can just query by the latest `vg_id` <= current version group (use
  window functions)
- [PK_1] `pkmn_id`: uint
- [PK_1] `slot`: uint
- [PK_2] `type_id`: uint
- [PK_3] `type_name`: string

Attributes
- is_hidden: bool

### Pokemon-Ability
Maps pokemon to their abilities. Alternatively, there are a set number of slots, so this could be
in-lined. Leaning towards that for now since there's not much reason to care about schema breaking
changes for a cache that I can just nuke later.

Identifiers:
- [PK] **required** `vg_id`: uint (should there be a latest sentinel to help with lookups
  and correctness as new games are released?)
- [PK_1] `pkmn_id`: uint
- [PK_1] `slot`: uint
- [PK_2] `ab_id`: uint

Attributes:
- `is_hidden`: bool

### Abilities
Identifiers:
- [PK] **required** `vg_id`: uint (no sentinel required)
- [PK_1] `id`: uint
- [PK_2] `name`: string

Attributes:
- `effect`: string

### Pokemon-Move
Maps pokemon to their movesets. This must be split out by version groups because learnsets can be
unique for each version group.

Identifiers:
- [PK] **required** `vg_id`: uint (no sentinel required)
- [PK_1] `pkmn_id`: uint
- [PK_2] `mv_id`: uint

Attributes:
- `method`: enum ([11 total](https://pokeapi.co/api/v2/move-learn-method/))
- `level_learned_at`: uint8 (only meaningful for "level-up")

### Moves
Identifiers:
- [PK] **required** `vg_id`: uint (no sentinel required)
- [PK_1] `id`: uint
- [PK_2] `name`: string

Attributes
- `type_id`: uint
- `type_name`: string
- `damage_class`: "status" | "special" | "physical"
- `pp`: uint8
- `priority`: int8
- `accuracy`: uint8 | nil
- `power`: uint8 | nil
- `effect_chance`: uint8 | nil
- `target`: string

## Open Questions
- Should generation and version group be expanded to versions? Different attributes are modeled
  differently in PokeAPI.

  A: From above, it looks like all the attributes we need to store are stored by version-group. I
  think that's a sane default to handle all the look-ups. Not interested as much in the individual
  versions since they tend to be sister games (with barely any changes).
