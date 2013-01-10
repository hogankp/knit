// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

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
	Cable
)

// String returns the string equivalent of the given stitch kind.
func (k StitchKind) String() string {
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
	case Cable:
		return "Ca"
	}

	panic("unreachable")
}
