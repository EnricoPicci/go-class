package packagea

import (
	"github.com/EnricoPicci/go-class/src/qualified-names/packageb"
	packagecba "github.com/EnricoPicci/go-class/src/qualified-names/packagec"
)

func HalloFromA() {
	packageb.HalloFromB()
	packagecba.HalloFromCBA()
}
