package engine

import (
	"fmt"
)

// Run executes dependent-origination cycles.
func Run(state *State, cycles int) {
	for i := 0; i < cycles; i++ {
		Continue(state)

		fmt.Printf("Cycle klesha[%d]\n", state.Turn)
		fmt.Println(state.History.Debug())
		fmt.Println(state.Condition.Debug())
		fmt.Println(state.Debug())
	}
}
