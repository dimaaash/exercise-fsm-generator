package main

import (
	"exercise-fsm-generator/fsm"
	"fmt"
	"os"
)

var (
	S0 fsm.State = "S0"
	S1 fsm.State = "S1"
	S2 fsm.State = "S2"
)

// Mod3Transition implements the transition function for the mod-three FSM.
func Mod3Transition(current fsm.State, input fsm.Input) fsm.State {
	switch current {
	case S0:
		if input == '0' {
			return S0
		}
		return S1
	case S1:
		if input == '0' {
			return S2
		}
		return S0
	case S2:
		if input == '0' {
			return S1
		}
		return S2
	}
	return current
}

func stateToRemainder(state fsm.State) int {
	switch state {
	case S0:
		return 0
	case S1:
		return 1
	case S2:
		return 2
	}
	return -1 // invalid state
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mod3 <binary string of length 2 to 16>")
		return
	}
	inputStr := os.Args[1]
	if len(inputStr) < 2 || len(inputStr) > 16 {
		fmt.Println("Error: Input must be a binary string (1s and 0s) of length between 2 and 16")
		return
	}
	var inputs []fsm.Input
	for _, ch := range inputStr {
		if ch == '0' || ch == '1' {
			inputs = append(inputs, fsm.Input(ch))
		} else {
			fmt.Printf("Error: Invalid input character: %c\n", ch)
			return
		}
	}

	mod3 := fsm.NewFSM(
		[]fsm.State{S0, S1, S2},
		[]fsm.Input{'0', '1'},
		S0,
		[]fsm.State{S0, S1, S2},
		Mod3Transition,
	)

	finalState := mod3.Run(inputs)
	remainder := stateToRemainder(finalState)
	fmt.Printf("Input: %s\n", inputStr)
	fmt.Printf("Final State: %s\n", finalState)
	fmt.Printf("Remainder (mod 3): %d\n", remainder)
}
