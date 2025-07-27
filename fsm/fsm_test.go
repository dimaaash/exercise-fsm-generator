package fsm

import (
	"testing"
)

func TestFSM_Mod3Transitions(t *testing.T) {
	S0 := State("S0")
	S1 := State("S1")
	S2 := State("S2")

	mod3Transition := func(current State, input Input) State {
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

	fsm := NewFSM(
		[]State{S0, S1, S2},
		[]Input{'0', '1'},
		S0,
		[]State{S0, S1, S2},
		mod3Transition,
	)

	tests := []struct {
		inputs   []Input
		expected State
	}{
		{[]Input{'0'}, S0},
		{[]Input{'1'}, S1},
		{[]Input{'1', '0'}, S2},
		{[]Input{'1', '1'}, S0},
		{[]Input{'1', '0', '1'}, S2},
		{[]Input{'1', '0', '1', '1'}, S2},
	}

	for i, test := range tests {
		fsm.Reset()
		result := fsm.Run(test.inputs)
		if result != test.expected {
			t.Errorf("Test %d: inputs %v, expected %s, got %s", i, test.inputs, test.expected, result)
		}
	}
}

func TestFSM_InvalidInputPanics(t *testing.T) {
	S0 := State("S0")
	S1 := State("S1")
	S2 := State("S2")
	mod3Transition := func(current State, input Input) State {
		return S0
	}
	fsm := NewFSM(
		[]State{S0, S1, S2},
		[]Input{'0', '1'},
		S0,
		[]State{S0, S1, S2},
		mod3Transition,
	)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid input, but did not panic")
		}
	}()
	fsm.Step('2') // invalid input
}

func TestFSM_TransitionToInvalidStatePanics(t *testing.T) {
	S0 := State("S0")
	S1 := State("S1")
	S2 := State("S2")
	S3 := State("S3") // not in FSM states
	mod3Transition := func(current State, input Input) State {
		return S3 // always returns invalid state
	}
	fsm := NewFSM(
		[]State{S0, S1, S2},
		[]Input{'0', '1'},
		S0,
		[]State{S0, S1, S2},
		mod3Transition,
	)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for transition to invalid state, but did not panic")
		}
	}()
	fsm.Step('0')
}

func TestFSM_InfiniteLoopDetection(t *testing.T) {
	S0 := State("S0")
	modLoop := func(current State, input Input) State {
		return current // never changes state
	}
	fsm := NewFSM(
		[]State{S0},
		[]Input{'0'},
		S0,
		[]State{S0},
		modLoop,
	)
	fsm.maxSteps = 5 // set low max steps for test
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for infinite loop detection, but did not panic")
		}
	}()
	fsm.Run([]Input{'0', '0', '0', '0', '0', '0', '0'})
}
