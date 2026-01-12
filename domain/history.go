package domain

import "fmt"

// HistoryNode represents one causal step.
// It has no identity beyond its position in the chain.
type HistoryNode struct {
	ProcessID string
	Delta     Vector
	Prev      *HistoryNode
}

// History is a causal chain, newest-first.
// It is traversed backward only.
type History struct {
	Head *HistoryNode
}

// NewHistory with a limit
func NewHistory() History {
	return History{
		Head: nil,
	}
}

// Record appends a new causal step to history.
func (h *History) Record(processID string, delta *Vector) {
	h.Head = &HistoryNode{
		ProcessID: processID,
		Delta:     *delta.Clone(),
		Prev:      h.Head,
	}
}

// RecentMagnitude computes decayed cumulative effect
// of a process over recent history.
func (h *History) RecentMagnitude(processID string, decay float64) float64 {
	var sum float64
	var weight = 1.0

	for node := h.Head; node != nil; node = node.Prev {
		if node.ProcessID == processID {
			sum += node.Delta.Magnitude() * weight
		}
		weight *= decay
		if weight < 0.01 {
			break
		}
	}

	return sum
}

// RecentDelta returns the most recent causal delta,
// or a zero _vec if history is empty.
func (h *History) RecentDelta() *Vector {
	if h.Head == nil {
		return &Vector{}
	}
	return h.Head.Delta.Clone()
}

// Debug prints the recent causal history (most recent first).
func (h *History) Debug() string {
	out := "history:\n"
	node := h.Head
	i := 0

	for node != nil && i < 5 {
		out += fmt.Sprintf(
			"  [%s] Δ\n%s",
			node.ProcessID,
			node.Delta.Debug(),
		)
		node = node.Prev
		i++
	}

	return out
}
