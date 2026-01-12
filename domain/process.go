package domain

// Process represents a skillful means (Upāya).
// It is optional, conditional, and impermanent.
type Process struct {
	ID string

	ExpectedDelta *Vector // Direction of suffering reduction
	DecayRate     float64 // Impermanence speed

	MinTurns int
	MaxTurns int

	StartedAt int
	LastTurn  int
}

func (p Process) String() string {
	return p.ID
}

const NothingID = "nothing"

// Nothing represents non-intervention.
// It applies no delta and consumes a turn.
func Nothing() *Process {
	return &Process{
		ID:            NothingID,
		ExpectedDelta: Ptr(ZeroVector()),
		DecayRate:     0,
		MinTurns:      1,
		MaxTurns:      1,
	}
}

const SilenceID = "silence"

// Silence returns the Noble Silence process.
//
// Noble Silence represents the intentional absence of reaction.
// It reduces affliction by not feeding it and increases method
// capacity through non-interference.
//
// In Buddhist terms:
// - This is not suppression.
// - This is not avoidance.
// - This is cessation of proliferation (papañca).
func Silence() *Process {
	return &Process{
		ID: SilenceID,

		ExpectedDelta: &Vector{
			_vec: [_axis]*float64{
				AxisAffliction: Ptr(-.10),
				AxisMethod:     Ptr(.10),
			},
		},

		DecayRate: 1.0,
		MinTurns:  1,
		MaxTurns:  3,
	}
}

const InsightID = "insight"

// Insight returns the Insight process.
//
// Insight represents clear seeing of conditions as they are.
// It reduces affliction by increasing clarity, not by force.
//
// This process should only be effective when capacity is sufficient.
func Insight() *Process {
	return &Process{
		ID: InsightID,

		ExpectedDelta: &Vector{
			_vec: _vec{
				AxisAffliction: Ptr(-0.15),
				AxisClarity:    Ptr(0.15),
			},
		},

		DecayRate: 0.8,
		MinTurns:  2,
		MaxTurns:  4,
	}
}

const RestraintID = "restraint"

// Restraint returns the Restraint process.
//
// Restraint prevents escalation into speech or action.
// It does not remove affliction directly,
// but blocks manifestation pathways.
func Restraint() *Process {
	return &Process{
		ID: RestraintID,

		ExpectedDelta: &Vector{
			_vec: [_axis]*float64{
				AxisManifestation: Ptr(-0.20),
				AxisBoundary:      Ptr(0.10),
			},
		},

		DecayRate: 1.0,
		MinTurns:  1,
		MaxTurns:  3,
	}
}

const InvestigationID = "investigation"

// Investigation returns the Investigation process.
//
// Investigation increases clarity in the presence of uncertainty.
// It does NOT resolve uncertainty immediately;
// it prepares conditions for insight.
func Investigation() *Process {
	return &Process{
		ID: InvestigationID,

		ExpectedDelta: &Vector{
			_vec: [_axis]*float64{
				AxisClarity:     Ptr(0.10),
				AxisUncertainty: Ptr(-0.10),
			},
		},

		DecayRate: 0.7,
		MinTurns:  2,
		MaxTurns:  5,
	}
}

const GroundingID = "grounding"

// Grounding returns the Grounding process.
//
// Grounding stabilizes capacity and reduces momentum.
// This is used when the system is overstimulated.
func Grounding() *Process {
	return &Process{
		ID: GroundingID,

		ExpectedDelta: &Vector{
			_vec: [_axis]*float64{
				AxisCapacity: Ptr(0.15),
				AxisMomentum: Ptr(-0.15),
			},
		},

		DecayRate: 1.0,
		MinTurns:  1,
		MaxTurns:  3,
	}
}
