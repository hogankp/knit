// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

// Represents stitch modifiers.
type StitchMod uint8

// Known stitch modifier flags.
const (
	BackLoop StitchMod = 1 << iota
)

func (m StitchMod) String() string {
	switch m {
	case BackLoop:
		return "@"
	}

	panic("unreachable")
}

// getModKind returns the appropriate modifier flag for the given input string.
func getModKind(v string) StitchMod {
	switch v {
	case "@":
		return BackLoop
	}

	return 0
}
