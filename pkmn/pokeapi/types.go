package pokeapi

type stat uint8

//go:generate stringer -type=stat -linecomment
const (
	stat_Undefined stat = iota // undefined
	stat_HP                    // hp
	stat_Atk                   // attack
	stat_Def                   // defense
	stat_SpAtk                 // special-attack
	stat_SpDef                 // special-defense
	stat_Spd                   // speed
)
