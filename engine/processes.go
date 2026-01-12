package engine

import "github.com/iveronanomi/samepage/domain"

// DefaultProcesses returns the canonical set of available processes.
//
// These represent core Buddhist interventions as state transformations.
// Order does not imply priority.
var DefaultProcesses = []*domain.Process{
	domain.Silence(),
	domain.Insight(),
	domain.Restraint(),
	domain.Investigation(),
	domain.Grounding(),
}
