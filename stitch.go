// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

import (
	"fmt"
	"strings"
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
	Kind StitchKind // Type of stitch.
	Mod  StitchMod  // Stitch modifier.
}

// Line returns the original pattern source line number for this node.
func (s *Stitch) Line() int { return s.line }

// Col returns the original pattern source column number for this node.
func (s *Stitch) Col() int { return s.col }

func (s *Stitch) String() string {
	if s.Mod == 0 {
		return s.Kind.String()
	}

	return fmt.Sprintf("%s%s", s.Mod, s.Kind)
}
