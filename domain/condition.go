package domain

import (
	"bytes"
	"fmt"
	"strings"
)

// Condition aggregates (Skandhas) a transient, impersonal
// state of mind at a single moment.
// contains no identity, no history, no narrative
type Condition struct {
	*Vector
	Uncertainty float64 // tolerated ambiguity (not an error)
}

func (c Condition) Debug() string {
	var b bytes.Buffer
	b.WriteString("condition:\n")
	b.WriteString(fmt.Sprintf("  %-14s  %+.2f\n", "uncertainty", c.Uncertainty))
	for _, str := range strings.Split(c.Vector.Debug(), "\n") {
		b.WriteString("\n  " + str)
	}
	return b.String()
}
