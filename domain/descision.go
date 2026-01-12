package domain

type Decision struct {
	Selected  *Process
	Score     float64
	Threshold float64
	Rejected  []ScoredProcess
	Reason    DecisionReason
}

type ScoredProcess struct {
	*Process
	Score float64
}

type DecisionReason string

const (
	ReasonBelowThreshold DecisionReason = "below_threshold"
	ReasonNoBetterOption DecisionReason = "no_better_option"
	ReasonContinuation   DecisionReason = "continuation"
)
