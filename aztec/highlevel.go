package aztec

import (
	"github.com/boombuler/barcode/utils"
)

func highlevelEncode(data []byte) *utils.BitList { _ = "STUB: not implemented"; return nil }

// We have one of the four special PUNCT pairs.  Treat them specially.
// Get a new set of states for the two new characters.

// Get a new set of states for the new character.

func simplifyStates(states stateSlice) stateSlice {
	_ = "STUB: not implemented"
	return *new(stateSlice)
}

// We update a set of states for a new character by updating each state
// for the new character, merging the results, and then removing the
// non-optimal states.
func updateStateListForChar(states stateSlice, data []byte, index int) stateSlice {
	_ = "STUB: not implemented"
	return *new(stateSlice)
}

// Return a set of states that represent the possible ways of updating this
// state for the next character.  The resulting set of states are added to
// the "result" list.
func updateStateForChar(s *state, data []byte, index int) stateSlice {
	_ = "STUB: not implemented"
	return *new(stateSlice)
}

// Only create stateNoBinary the first time it's required.

// Try generating the character by latching to its mode

// If the character is in the current table, we don't want to latch to
// any other mode except possibly digit (which uses only 4 bits).  Any
// other latch would be equally successful *after* this character, and
// so wouldn't save any bits.

// Try generating the character by switching to its mode.

// It never makes sense to temporarily shift to another mode if the
// character exists in the current mode.  That can never save bits.

// It's never worthwhile to go into binary shift mode if you're not already
// in binary shift mode, and the character exists in your current mode.
// That can never save bits over just outputting the char in the current mode.

// We update a set of states for a new character by updating each state
// for the new character, merging the results, and then removing the
// non-optimal states.
func updateStateListForPair(states stateSlice, data []byte, index int, pairCode int) stateSlice {
	_ = "STUB: not implemented"
	return *new(stateSlice)
}

func updateStateForPair(s *state, data []byte, index int, pairCode int) stateSlice {
	_ = "STUB: not implemented"
	return *new(stateSlice)
}

// Possibility 1.  Latch to MODE_PUNCT, and then append this code

// Possibility 2.  Shift to MODE_PUNCT, and then append this code.
// Every state except MODE_PUNCT (handled above) can shift

// both characters are in DIGITS.  Sometimes better to just add two digits

// period or comma in DIGIT
// space in DIGIT

// It only makes sense to do the characters as binary if we're already
// in binary mode.
