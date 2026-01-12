package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/iveronanomi/samepage/engine"
)

// Entry point.
// This is Right Intention: initiate observation, not control.
func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome back hunter, what is your desire?")
	input, _ := reader.ReadString('\n')

	// Something has already arisen
	// we are observing it
	condition := engine.Observe(input)

	state := engine.NewState(condition)

	engine.Run(state, 3)
}
