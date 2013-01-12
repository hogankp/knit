// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

// Represents stitch modifiers.
type StitchMod uint8

// Known stitch modifier flags.
const (
	BackLoop StitchMod = 1 << iota
	DeepKnit
	YarnForward
	YarnBackward
)

func (m StitchMod) String() string {
	switch m {
	case BackLoop:
		return "@"
	case DeepKnit:
		return "^"
	case YarnForward:
		return ">"
	case YarnBackward:
		return "<"
	}

	panic("unreachable")
}

// isMod returns true if the given byte represents a known modifier.
func isMod(b byte) bool {
	switch b {
	case '@', '^', '<', '>':
		return true
	}

	return false
}

// getModKind returns the appropriate modifier flag for the given input string.
func getModKind(v string) StitchMod {
	switch v {
	case "@":
		return BackLoop
	case "^":
		return DeepKnit
	case ">":
		return YarnForward
	case "<":
		return YarnBackward
	}

	return 0
}
