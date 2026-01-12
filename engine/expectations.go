package engine

// CheckExpectation returns true if the active process
// is still progressing in the expected direction.
//
// This checks causal coherence, not success or completion.
func CheckExpectation(state *State) bool {
	// No previous delta → cannot evaluate yet
	var delta = state.History.RecentDelta()
	if delta.IsZero() {
		return true
	}

	// Expected direction vs actual recent movement
	// If movement contradicts expectation, abort
	if state.Active.ExpectedDelta.Dot(delta) <= 0 {
		return false
	}

	// Optional:
	// Ensure effect is not vanishing too fast
	if delta.Magnitude() < 0.001 {
		return false
	}

	return true
}
