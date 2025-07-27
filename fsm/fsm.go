package fsm

// State represents a state in the FSM.
type State string

// Input represents an input symbol.
type Input rune

// TransitionFunc defines the transition function δ.
type TransitionFunc func(current State, input Input) State

// FSM defines the finite state machine structure.
type FSM struct {
	States       []State
	Alphabet     []Input
	InitialState State
	FinalStates  []State
	Transition   TransitionFunc
	currentState State
	maxSteps     int // max steps to prevent infinite loops
}

// NewFSM creates a new FSM instance.
func NewFSM(states []State, alphabet []Input, initial State, finals []State, delta TransitionFunc) *FSM {
	return &FSM{
		States:       states,
		Alphabet:     alphabet,
		InitialState: initial,
		FinalStates:  finals,
		Transition:   delta,
		currentState: initial,
		maxSteps:     10000, // default max steps
	}
}

// isValidState checks if a state is in the FSM's state set.
func (f *FSM) isValidState(s State) bool {
	for _, st := range f.States {
		if st == s {
			return true
		}
	}
	return false
}

// isValidInput checks if an input is in the FSM's alphabet.
func (f *FSM) isValidInput(i Input) bool {
	for _, in := range f.Alphabet {
		if in == i {
			return true
		}
	}
	return false
}

// Reset sets the FSM back to its initial state.
func (f *FSM) Reset() {
	f.currentState = f.InitialState
}

// Step processes a single input symbol.
func (f *FSM) Step(input Input) State {
	if !f.isValidInput(input) {
		panic("FSM: invalid input symbol")
	}
	if !f.isValidState(f.currentState) {
		panic("FSM: invalid current state")
	}
	var next State
	defer func() {
		if r := recover(); r != nil {
			panic("FSM: panic in transition function")
		}
	}()
	next = f.Transition(f.currentState, input)
	if !f.isValidState(next) {
		panic("FSM: transition to invalid state")
	}
	f.currentState = next
	return f.currentState
}

// Run processes a sequence of input symbols.
func (f *FSM) Run(inputs []Input) State {
	steps := 0
	for _, in := range inputs {
		if steps > f.maxSteps {
			panic("FSM: exceeded maximum allowed steps (possible infinite loop)")
		}
		f.Step(in)
		steps++
	}
	return f.currentState
}

// IsFinal returns true if the current state is a final state.
func (f *FSM) IsFinal() bool {
	for _, s := range f.FinalStates {
		if s == f.currentState {
			return true
		}
	}
	return false
}
