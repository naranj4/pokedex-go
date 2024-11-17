package pkmn

type LearnMethod uint

//go:generate stringer -type=LearnMethod -linecomment
const (
	LearnMethod_Undefined LearnMethod = iota // undefined

	// common learn methods

	LearnMethod_LevelUp // level-up
	LearnMethod_Egg     // egg
	LearnMethod_Tutor   // tutor
	LearnMethod_Machine // machine

	// special snowflake learn methods for certain pokemon

	LearnMethod_StadiumSurfingPikachu // stadium-surfing-pikachu
	LearnMethod_LightBallEgg          // light-ball-egg
	LearnMethod_ColosseumPurification // colosseum-purification
	LearnMethod_XDShadow              // xd-shadow
	LearnMethod_XDPurification        // xd-purification
	LearnMethod_FormChange            // form-change
	LearnMethod_ZygardeCube           // zygarde-cube
)
