package engine

import (
	"bytes"

	"github.com/iveronanomi/samepage/domain"
)

// State is the dependent-origination state.
// No truth is stored here — only flow.
type State struct {
	Condition *domain.Condition
	Turn      int

	Active    *domain.Process
	Processes []*domain.Process
	History   *domain.History
}

func NewState(condition *domain.Condition) *State {
	return &State{
		Condition: condition,
		Processes: DefaultProcesses,
		History:   domain.Ptr(domain.NewHistory()),
	}
}

// Debug state, shows causal relations (Right View).
func (s State) Debug() string {
	if s.Active == nil {
		return ""
	}

	var b bytes.Buffer
	b.WriteString("[condition] --selects--> [process:")
	b.WriteString(s.Active.ID)
	b.WriteString("]\n[process:")
	b.WriteString(s.Active.ID)
	b.WriteString("] --applies--> [condition]\n")
	b.WriteString(s.Active.ExpectedDelta.Debug())

	return b.String()
}
