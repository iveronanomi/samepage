package engine

import (
	"strings"

	"github.com/iveronanomi/samepage/domain"
)

// Observe observes the user's current expressed state
// and encodes it as a transient condition vector.
//
// INVARIANTS:
//
// - This is observation, not interpretation.
// - This is not intent detection.
// - This is not advice.
// - This does not infer identity or continuity.
//
// The output represents what has arisen *now* (paccuppanna-dhamma).
func Observe(expression string) *domain.Condition {
	v := domain.ZeroVector()

	normalized := strings.ToLower(strings.TrimSpace(expression))

	evaluateAffliction(normalized, &v)
	evaluateManifestation(normalized, &v)
	evaluateCapacity(&v)

	return &domain.Condition{
		Vector:      &v,
		Uncertainty: baseUncertainty(normalized),
	}
}

// evaluateAffliction detects coarse emotional valence.
// It never assigns meaning, only intensity.
func evaluateAffliction(expr string, v *domain.Vector) {
	// Observe affective pressure, not intent.
	if strings.Contains(expr, "angry") {
		v.Set(int(domain.AxisAffliction), 0.7)
		v.Set(int(domain.AxisClarity), -0.3)
	}

	// Neutral / friendly contact slightly reduces affliction
	if strings.Contains(expr, "hello") ||
		strings.Contains(expr, "welcome") {
		v.Set(int(domain.AxisAffliction), -0.1)
	}
}

// evaluateManifestation estimates outward expression.
func evaluateManifestation(expr string, v *domain.Vector) {
	// Detect outward expression tendency (very coarse)
	if strings.Contains(expr, "!") {
		v.Set(int(domain.AxisManifestation), 0.4)
	}
}

// evaluateCapacity sets a safe default readiness.
func evaluateCapacity(v *domain.Vector) {
	// Neutral openness unless constrained later
	v.Set(int(domain.AxisCapacity), 0.6)
}

// baseUncertainty assigns uncertainty conservatively.
//
// More text ≠ more certainty.
// Certainty only decreases via observation over time.
func baseUncertainty(expr string) float64 {
	if expr == "" {
		return 0.9
	}
	return 0.5
}
