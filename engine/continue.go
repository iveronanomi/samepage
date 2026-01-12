package engine

import "github.com/iveronanomi/samepage/domain"

const MinDecisionScore = 0.05

// Continue Process advances the currently active process by one turn.
// It applies the expected delta, records the effect, and decides
// whether the process should continue or be aborted.
func Continue(state *State) {
	state.Turn++

	// Select chooses the next process to apply.
	// It does not mutate state.
	// It returns nil if no process exceeds the decision threshold.
	var best *domain.Process
	var bestScore float64

	for pid := range state.Processes {
		if score := Score(pid, state); score > bestScore {
			best = state.Processes[pid]
			bestScore = score
		}
	}

	// No active process → nothing to do
	if bestScore < MinDecisionScore {
		state.Active = domain.Nothing()
		return
	}
	state.Active = best

	// Capture previous condition vector
	prev := state.Condition.Vector.Clone()

	// Apply expected delta (with decay)
	state.Condition.Vector.Add(
		state.Active.ExpectedDelta.Scaled(state.Active.DecayRate),
	)

	// 4. Normalize condition
	state.Condition.Vector.Normalize()

	// 5. Record causal delta in history
	actual := state.Condition.Vector.Sub(prev)
	state.History.Record(state.Active.ID, actual)

	// 6. Expectation check (after MinTurns)
	if state.Turn-state.Active.StartedAt >= state.Active.MinTurns {
		if !CheckExpectation(state) {
			state.Active = nil
		}
	}

	// 7. Normal exit on MaxTurns
	if state.Turn-state.Active.StartedAt >= state.Active.MaxTurns {
		state.Active = nil
	}
}
