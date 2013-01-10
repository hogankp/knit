// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

import "strings"

type StitchKind uint8

// Known stitch kinds.
const (
	UnknownStitch StitchKind = iota
	KnitStitch
	PurlStitch
	KnitSlip
	PurlSlip
	CastOn
	BindOff
	Increase
	Decrease
	YarnOver
	K2Tog
	K3Tog
	K4Tog
	P2Tog
	P3Tog
	P4Tog
)

// Listing of builtin stitch types.
var stitches map[string]StitchKind

func init() {
	stitches = make(map[string]StitchKind)
	stitches["inc"] = Increase
	stitches["dec"] = Decrease
	stitches["tog"] = Decrease
	stitches["yo"] = YarnOver
	stitches["ks"] = KnitSlip
	stitches["ps"] = PurlSlip
	stitches["co"] = CastOn
	stitches["bo"] = BindOff
	stitches["k"] = KnitStitch
	stitches["p"] = PurlStitch
	stitches["p2tog"] = P2Tog
	stitches["p3tog"] = P3Tog
	stitches["p4tog"] = P4Tog
	stitches["k2tog"] = K2Tog
	stitches["k3tog"] = K3Tog
	stitches["k4tog"] = K4Tog
}

// getStitchName returns the string equivalent of the given stitch kind.
func getStitchName(k StitchKind) string {
	switch k {
	case UnknownStitch:
		return "Unknown"
	case KnitStitch:
		return "K"
	case PurlStitch:
		return "P"
	case KnitSlip:
		return "Ks"
	case PurlSlip:
		return "Ps"
	case CastOn:
		return "Co"
	case BindOff:
		return "Bo"
	case Increase:
		return "Inc"
	case Decrease:
		return "Dec"
	case YarnOver:
		return "Yo"
	case K2Tog:
		return "K2Tog"
	case K3Tog:
		return "K3Tog"
	case K4Tog:
		return "K4Tog"
	case P2Tog:
		return "P2Tog"
	case P3Tog:
		return "P3Tog"
	case P4Tog:
		return "K4Tog"
	}

	panic("unreachable")
}

// getStitchKind returns the kind of stitch represented by the
// supplied string.
func getStitchKind(s string) StitchKind {
	for k, v := range stitches {
		if strings.EqualFold(k, s) {
			return v
		}
	}

	return UnknownStitch
}

// A stich defines a specific kind of stitch to perform.
type Stitch struct {
	line int
	col  int
	Kind StitchKind
}

// Line returns the original pattern source line number for this node.
func (s *Stitch) Line() int { return s.line }

// Col returns the original pattern source column number for this node.
func (s *Stitch) Col() int { return s.col }
