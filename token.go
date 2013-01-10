// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

type tokenType uint8

// Known token types.
const (
	tokEof tokenType = iota
	tokError
	tokStitch
	tokNumber
	tokQuantifier
	tokModifier
	tokGroupStart
	tokGroupEnd
	tokRow
)

func (t tokenType) String() string {
	switch t {
	case tokEof:
		return "EOF"
	case tokError:
		return "ERROR"
	case tokStitch:
		return "STITCH"
	case tokNumber:
		return "NUMBER"
	case tokQuantifier:
		return "QUANT"
	case tokModifier:
		return "MOD"
	case tokGroupStart:
		return "GROUPS"
	case tokGroupEnd:
		return "GROUPE"
	case tokRow:
		return "ROW"
	}

	panic("unreachable")
}

// A token represents a single parsed pattern token.
type token struct {
	Type tokenType
	Data string
	Line int
	Col  int
}
