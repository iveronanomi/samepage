package domain

// Axis represents a phenomenological dimension of experience.
//
// IMPORTANT INVARIANTS:
//
// - Axes are dimensions, not traits.
// - Values are continuous and transient.
// - Axes describe conditioned experience (dhamma), not a self.
// - Axes must remain stable across engine evolution.
//
// Axes correspond to *spaces of observation*, not conclusions.
type Axis int8
type AxisMultiplication [_axis]Axis

const (
	AxisUncertainty Axis = iota

	// ──────────────── Klesha / Purity Space ────────────────

	AxisAffliction // intensity of greed / aversion / delusion presence
	AxisClarity    // non-delusion / lucidity / absence of confusion

	// ──────────────── Expression Space ────────────────

	AxisManifestation  // thought ↔ speech ↔ action (externalization)
	AxisIntentionality // reactive ↔ deliberate ↔ spontaneous

	// ──────────────── Temporal Space ────────────────

	AxisTemporal // arising ↔ abiding ↔ fading
	AxisMomentum // inertia / persistence of current state

	// ──────────────── Relational Space ────────────────

	AxisRelation // self ↔ other ↔ authority
	AxisBoundary // enmeshment ↔ separation

	// ──────────────── Capacity Space ────────────────

	AxisCapacity  // readiness / openness
	AxisStability // ability to remain without collapse

	// ──────────────── Method / Skillful Means ────────────────

	AxisMethod // insight / silence / restraint
	AxisEffort // forcing ↔ balanced ↔ relinquishing

	_axis //
)

// String returns a stable, human-readable axis name.
// Used for debugging, visualization, and documentation.
func (a Axis) String() string {
	switch a {
	case AxisAffliction:
		return "affliction"
	case AxisClarity:
		return "clarity"
	case AxisManifestation:
		return "manifestation"
	case AxisIntentionality:
		return "intentionality"
	case AxisTemporal:
		return "temporal"
	case AxisMomentum:
		return "momentum"
	case AxisRelation:
		return "relation"
	case AxisBoundary:
		return "boundary"
	case AxisCapacity:
		return "capacity"
	case AxisStability:
		return "stability"
	case AxisMethod:
		return "method"
	case AxisEffort:
		return "effort"
	case AxisUncertainty:
		return "uncertainty"
	default:
		return "{undefined}"
	}
}
