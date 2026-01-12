package engine

func Score(
	pid int,
	state *State,
) float64 {
	var score float64

	// Alignment with current Condition
	state.Processes[pid].ExpectedDelta.Each(func(i int, val *float64) {
		if val == nil || state.Condition.Vector == nil {
			return
		}
		if av := state.Condition.Vector.Get(i); av != nil {
			score += *val * *av
		}
	})

	// Residual causal load (fatigue)
	fatigue := state.History.RecentMagnitude(state.Processes[pid].ID, 0.6)
	score -= fatigue

	return score
}
